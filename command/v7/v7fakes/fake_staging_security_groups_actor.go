// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
	"code.cloudfoundry.org/cli/resources"
)

type FakeStagingSecurityGroupsActor struct {
	GetGlobalStagingSecurityGroupsStub        func() ([]resources.SecurityGroup, v7action.Warnings, error)
	getGlobalStagingSecurityGroupsMutex       sync.RWMutex
	getGlobalStagingSecurityGroupsArgsForCall []struct {
	}
	getGlobalStagingSecurityGroupsReturns struct {
		result1 []resources.SecurityGroup
		result2 v7action.Warnings
		result3 error
	}
	getGlobalStagingSecurityGroupsReturnsOnCall map[int]struct {
		result1 []resources.SecurityGroup
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStagingSecurityGroupsActor) GetGlobalStagingSecurityGroups() ([]resources.SecurityGroup, v7action.Warnings, error) {
	fake.getGlobalStagingSecurityGroupsMutex.Lock()
	ret, specificReturn := fake.getGlobalStagingSecurityGroupsReturnsOnCall[len(fake.getGlobalStagingSecurityGroupsArgsForCall)]
	fake.getGlobalStagingSecurityGroupsArgsForCall = append(fake.getGlobalStagingSecurityGroupsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetGlobalStagingSecurityGroups", []interface{}{})
	fake.getGlobalStagingSecurityGroupsMutex.Unlock()
	if fake.GetGlobalStagingSecurityGroupsStub != nil {
		return fake.GetGlobalStagingSecurityGroupsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getGlobalStagingSecurityGroupsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeStagingSecurityGroupsActor) GetGlobalStagingSecurityGroupsCallCount() int {
	fake.getGlobalStagingSecurityGroupsMutex.RLock()
	defer fake.getGlobalStagingSecurityGroupsMutex.RUnlock()
	return len(fake.getGlobalStagingSecurityGroupsArgsForCall)
}

func (fake *FakeStagingSecurityGroupsActor) GetGlobalStagingSecurityGroupsCalls(stub func() ([]resources.SecurityGroup, v7action.Warnings, error)) {
	fake.getGlobalStagingSecurityGroupsMutex.Lock()
	defer fake.getGlobalStagingSecurityGroupsMutex.Unlock()
	fake.GetGlobalStagingSecurityGroupsStub = stub
}

func (fake *FakeStagingSecurityGroupsActor) GetGlobalStagingSecurityGroupsReturns(result1 []resources.SecurityGroup, result2 v7action.Warnings, result3 error) {
	fake.getGlobalStagingSecurityGroupsMutex.Lock()
	defer fake.getGlobalStagingSecurityGroupsMutex.Unlock()
	fake.GetGlobalStagingSecurityGroupsStub = nil
	fake.getGlobalStagingSecurityGroupsReturns = struct {
		result1 []resources.SecurityGroup
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStagingSecurityGroupsActor) GetGlobalStagingSecurityGroupsReturnsOnCall(i int, result1 []resources.SecurityGroup, result2 v7action.Warnings, result3 error) {
	fake.getGlobalStagingSecurityGroupsMutex.Lock()
	defer fake.getGlobalStagingSecurityGroupsMutex.Unlock()
	fake.GetGlobalStagingSecurityGroupsStub = nil
	if fake.getGlobalStagingSecurityGroupsReturnsOnCall == nil {
		fake.getGlobalStagingSecurityGroupsReturnsOnCall = make(map[int]struct {
			result1 []resources.SecurityGroup
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getGlobalStagingSecurityGroupsReturnsOnCall[i] = struct {
		result1 []resources.SecurityGroup
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStagingSecurityGroupsActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getGlobalStagingSecurityGroupsMutex.RLock()
	defer fake.getGlobalStagingSecurityGroupsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStagingSecurityGroupsActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.StagingSecurityGroupsActor = new(FakeStagingSecurityGroupsActor)