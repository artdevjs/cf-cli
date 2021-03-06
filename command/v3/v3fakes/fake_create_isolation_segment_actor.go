// This file was generated by counterfeiter
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeCreateIsolationSegmentActor struct {
	CreateIsolationSegmentStub        func(name string) (v3action.Warnings, error)
	createIsolationSegmentMutex       sync.RWMutex
	createIsolationSegmentArgsForCall []struct {
		name string
	}
	createIsolationSegmentReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	createIsolationSegmentReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreateIsolationSegmentActor) CreateIsolationSegment(name string) (v3action.Warnings, error) {
	fake.createIsolationSegmentMutex.Lock()
	ret, specificReturn := fake.createIsolationSegmentReturnsOnCall[len(fake.createIsolationSegmentArgsForCall)]
	fake.createIsolationSegmentArgsForCall = append(fake.createIsolationSegmentArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("CreateIsolationSegment", []interface{}{name})
	fake.createIsolationSegmentMutex.Unlock()
	if fake.CreateIsolationSegmentStub != nil {
		return fake.CreateIsolationSegmentStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createIsolationSegmentReturns.result1, fake.createIsolationSegmentReturns.result2
}

func (fake *FakeCreateIsolationSegmentActor) CreateIsolationSegmentCallCount() int {
	fake.createIsolationSegmentMutex.RLock()
	defer fake.createIsolationSegmentMutex.RUnlock()
	return len(fake.createIsolationSegmentArgsForCall)
}

func (fake *FakeCreateIsolationSegmentActor) CreateIsolationSegmentArgsForCall(i int) string {
	fake.createIsolationSegmentMutex.RLock()
	defer fake.createIsolationSegmentMutex.RUnlock()
	return fake.createIsolationSegmentArgsForCall[i].name
}

func (fake *FakeCreateIsolationSegmentActor) CreateIsolationSegmentReturns(result1 v3action.Warnings, result2 error) {
	fake.CreateIsolationSegmentStub = nil
	fake.createIsolationSegmentReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateIsolationSegmentActor) CreateIsolationSegmentReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.CreateIsolationSegmentStub = nil
	if fake.createIsolationSegmentReturnsOnCall == nil {
		fake.createIsolationSegmentReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.createIsolationSegmentReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateIsolationSegmentActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createIsolationSegmentMutex.RLock()
	defer fake.createIsolationSegmentMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCreateIsolationSegmentActor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v3.CreateIsolationSegmentActor = new(FakeCreateIsolationSegmentActor)
