package v2action_test

import (
	"errors"
	"time"

	. "code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/actor/v2action/v2actionfakes"
	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Application Instance Actions", func() {
	var (
		actor                     Actor
		fakeCloudControllerClient *v2actionfakes.FakeCloudControllerClient
	)

	BeforeEach(func() {
		fakeCloudControllerClient = new(v2actionfakes.FakeCloudControllerClient)
		actor = NewActor(fakeCloudControllerClient, nil)
	})

	Describe("ApplicationInstance", func() {
		var instance ApplicationInstance

		BeforeEach(func() {
			instance = ApplicationInstance{}
		})

		Describe("TimeSinceCreation", func() {
			It("returns the time the instance started", func() {
				instance.Since = 1485985587.12345
				Expect(instance.TimeSinceCreation()).To(Equal(time.Unix(1485985587, 0)))
			})
		})
	})

	Describe("GetApplicationInstancesByApplication", func() {
		Context("when the application exists", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.GetApplicationInstanceStatusesByApplicationReturns(
					map[int]ccv2.ApplicationInstanceStatus{
						0: {
							ID:          0,
							CPU:         100,
							Memory:      100,
							MemoryQuota: 200,
							Disk:        50,
							DiskQuota:   100,
						},
						1: {ID: 1, CPU: 200},
					},
					ccv2.Warnings{"stats-warning-1", "stats-warning-2"},
					nil)

				fakeCloudControllerClient.GetApplicationInstancesByApplicationReturns(
					map[int]ccv2.ApplicationInstance{
						0: {ID: 0, Details: "hello", Since: 1485985587.12345, State: ccv2.ApplicationInstanceRunning},
						1: {ID: 1, Details: "hi", Since: 1485985587.567},
					},
					ccv2.Warnings{"instance-warning-1", "instance-warning-2"},
					nil)
			})

			It("returns the application instances and all warnings", func() {
				instances, warnings, err := actor.GetApplicationInstancesByApplication("some-app-guid")
				Expect(err).ToNot(HaveOccurred())
				Expect(instances).To(ConsistOf(
					ApplicationInstance{
						ID:          0,
						CPU:         100,
						Memory:      100,
						MemoryQuota: 200,
						Disk:        50,
						DiskQuota:   100,
						Details:     "hello",
						Since:       1485985587.12345,
						State:       ccv2.ApplicationInstanceRunning,
					},
					ApplicationInstance{ID: 1, CPU: 200, Details: "hi", Since: 1485985587.567}))
				Expect(warnings).To(ConsistOf(
					"stats-warning-1",
					"stats-warning-2",
					"instance-warning-1",
					"instance-warning-2"))

				Expect(fakeCloudControllerClient.GetApplicationInstanceStatusesByApplicationCallCount()).To(Equal(1))
				Expect(fakeCloudControllerClient.GetApplicationInstanceStatusesByApplicationArgsForCall(0)).To(Equal("some-app-guid"))
				Expect(fakeCloudControllerClient.GetApplicationInstancesByApplicationCallCount()).To(Equal(1))
				Expect(fakeCloudControllerClient.GetApplicationInstancesByApplicationArgsForCall(0)).To(Equal("some-app-guid"))
			})
		})

		Context("when an error is encountered", func() {
			var expectedErr error

			BeforeEach(func() {
				expectedErr = errors.New("banana")
				fakeCloudControllerClient.GetApplicationInstanceStatusesByApplicationReturns(
					nil,
					ccv2.Warnings{"stats-warning"},
					nil)
				fakeCloudControllerClient.GetApplicationInstancesByApplicationReturns(
					nil,
					ccv2.Warnings{"instances-warning"},
					expectedErr)
			})

			It("returns the error and all warnings", func() {
				_, warnings, err := actor.GetApplicationInstancesByApplication("some-app-guid")
				Expect(err).To(MatchError(expectedErr))
				Expect(warnings).To(ConsistOf("stats-warning", "instances-warning"))
			})

			Context("when the application does not exist", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.GetApplicationInstanceStatusesByApplicationReturns(
						nil,
						nil,
						cloudcontroller.ResourceNotFoundError{})
				})

				It("returns an ApplicationInstancesNotFoundError", func() {
					_, _, err := actor.GetApplicationInstancesByApplication("some-app-guid")
					Expect(err).To(MatchError(ApplicationInstancesNotFoundError{ApplicationGUID: "some-app-guid"}))
				})
			})

			Context("when an instance is missing from stats", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.GetApplicationInstanceStatusesByApplicationReturns(
						map[int]ccv2.ApplicationInstanceStatus{
							0: {ID: 0},
						}, nil, nil)

					fakeCloudControllerClient.GetApplicationInstancesByApplicationReturns(
						map[int]ccv2.ApplicationInstance{
							0: {ID: 0},
							1: {ID: 1, Details: "backend details"},
						}, nil, nil)
				})

				It("sets the detail field to incomplete", func() {
					instances, _, err := actor.GetApplicationInstancesByApplication("some-app-guid")
					Expect(err).ToNot(HaveOccurred())
					Expect(instances).To(ConsistOf(
						ApplicationInstance{ID: 0},
						ApplicationInstance{ID: 1, Details: "backend details (Unable to retrieve information)"},
					))
				})
			})

			Context("when an instance is missing from instances", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.GetApplicationInstanceStatusesByApplicationReturns(
						map[int]ccv2.ApplicationInstanceStatus{
							0: {ID: 0},
							1: {ID: 1},
						}, nil, nil)

					fakeCloudControllerClient.GetApplicationInstancesByApplicationReturns(
						map[int]ccv2.ApplicationInstance{
							0: {ID: 0},
						}, nil, nil)
				})

				It("sets the detail field to incomplete", func() {
					instances, _, err := actor.GetApplicationInstancesByApplication("some-app-guid")
					Expect(err).ToNot(HaveOccurred())
					Expect(instances).To(ConsistOf(
						ApplicationInstance{ID: 0},
						ApplicationInstance{ID: 1, Details: "(Unable to retrieve information)"},
					))
				})
			})
		})
	})
})
