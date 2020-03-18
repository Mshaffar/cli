// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeSpaceUsersActor struct {
	GetOrganizationByNameStub        func(string) (v7action.Organization, v7action.Warnings, error)
	getOrganizationByNameMutex       sync.RWMutex
	getOrganizationByNameArgsForCall []struct {
		arg1 string
	}
	getOrganizationByNameReturns struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}
	getOrganizationByNameReturnsOnCall map[int]struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}
	GetSpaceByNameAndOrganizationStub        func(string, string) (v7action.Space, v7action.Warnings, error)
	getSpaceByNameAndOrganizationMutex       sync.RWMutex
	getSpaceByNameAndOrganizationArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSpaceByNameAndOrganizationReturns struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}
	getSpaceByNameAndOrganizationReturnsOnCall map[int]struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}
	GetSpaceUsersByRoleTypeStub        func(string) (map[constant.RoleType][]v7action.User, v7action.Warnings, error)
	getSpaceUsersByRoleTypeMutex       sync.RWMutex
	getSpaceUsersByRoleTypeArgsForCall []struct {
		arg1 string
	}
	getSpaceUsersByRoleTypeReturns struct {
		result1 map[constant.RoleType][]v7action.User
		result2 v7action.Warnings
		result3 error
	}
	getSpaceUsersByRoleTypeReturnsOnCall map[int]struct {
		result1 map[constant.RoleType][]v7action.User
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpaceUsersActor) GetOrganizationByName(arg1 string) (v7action.Organization, v7action.Warnings, error) {
	fake.getOrganizationByNameMutex.Lock()
	ret, specificReturn := fake.getOrganizationByNameReturnsOnCall[len(fake.getOrganizationByNameArgsForCall)]
	fake.getOrganizationByNameArgsForCall = append(fake.getOrganizationByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrganizationByName", []interface{}{arg1})
	fake.getOrganizationByNameMutex.Unlock()
	if fake.GetOrganizationByNameStub != nil {
		return fake.GetOrganizationByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationByNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSpaceUsersActor) GetOrganizationByNameCallCount() int {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return len(fake.getOrganizationByNameArgsForCall)
}

func (fake *FakeSpaceUsersActor) GetOrganizationByNameCalls(stub func(string) (v7action.Organization, v7action.Warnings, error)) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = stub
}

func (fake *FakeSpaceUsersActor) GetOrganizationByNameArgsForCall(i int) string {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	argsForCall := fake.getOrganizationByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSpaceUsersActor) GetOrganizationByNameReturns(result1 v7action.Organization, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	fake.getOrganizationByNameReturns = struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceUsersActor) GetOrganizationByNameReturnsOnCall(i int, result1 v7action.Organization, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	if fake.getOrganizationByNameReturnsOnCall == nil {
		fake.getOrganizationByNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Organization
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getOrganizationByNameReturnsOnCall[i] = struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceUsersActor) GetSpaceByNameAndOrganization(arg1 string, arg2 string) (v7action.Space, v7action.Warnings, error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	ret, specificReturn := fake.getSpaceByNameAndOrganizationReturnsOnCall[len(fake.getSpaceByNameAndOrganizationArgsForCall)]
	fake.getSpaceByNameAndOrganizationArgsForCall = append(fake.getSpaceByNameAndOrganizationArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetSpaceByNameAndOrganization", []interface{}{arg1, arg2})
	fake.getSpaceByNameAndOrganizationMutex.Unlock()
	if fake.GetSpaceByNameAndOrganizationStub != nil {
		return fake.GetSpaceByNameAndOrganizationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getSpaceByNameAndOrganizationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSpaceUsersActor) GetSpaceByNameAndOrganizationCallCount() int {
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	return len(fake.getSpaceByNameAndOrganizationArgsForCall)
}

func (fake *FakeSpaceUsersActor) GetSpaceByNameAndOrganizationCalls(stub func(string, string) (v7action.Space, v7action.Warnings, error)) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = stub
}

func (fake *FakeSpaceUsersActor) GetSpaceByNameAndOrganizationArgsForCall(i int) (string, string) {
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	argsForCall := fake.getSpaceByNameAndOrganizationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSpaceUsersActor) GetSpaceByNameAndOrganizationReturns(result1 v7action.Space, result2 v7action.Warnings, result3 error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = nil
	fake.getSpaceByNameAndOrganizationReturns = struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceUsersActor) GetSpaceByNameAndOrganizationReturnsOnCall(i int, result1 v7action.Space, result2 v7action.Warnings, result3 error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = nil
	if fake.getSpaceByNameAndOrganizationReturnsOnCall == nil {
		fake.getSpaceByNameAndOrganizationReturnsOnCall = make(map[int]struct {
			result1 v7action.Space
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getSpaceByNameAndOrganizationReturnsOnCall[i] = struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceUsersActor) GetSpaceUsersByRoleType(arg1 string) (map[constant.RoleType][]v7action.User, v7action.Warnings, error) {
	fake.getSpaceUsersByRoleTypeMutex.Lock()
	ret, specificReturn := fake.getSpaceUsersByRoleTypeReturnsOnCall[len(fake.getSpaceUsersByRoleTypeArgsForCall)]
	fake.getSpaceUsersByRoleTypeArgsForCall = append(fake.getSpaceUsersByRoleTypeArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetSpaceUsersByRoleType", []interface{}{arg1})
	fake.getSpaceUsersByRoleTypeMutex.Unlock()
	if fake.GetSpaceUsersByRoleTypeStub != nil {
		return fake.GetSpaceUsersByRoleTypeStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getSpaceUsersByRoleTypeReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSpaceUsersActor) GetSpaceUsersByRoleTypeCallCount() int {
	fake.getSpaceUsersByRoleTypeMutex.RLock()
	defer fake.getSpaceUsersByRoleTypeMutex.RUnlock()
	return len(fake.getSpaceUsersByRoleTypeArgsForCall)
}

func (fake *FakeSpaceUsersActor) GetSpaceUsersByRoleTypeCalls(stub func(string) (map[constant.RoleType][]v7action.User, v7action.Warnings, error)) {
	fake.getSpaceUsersByRoleTypeMutex.Lock()
	defer fake.getSpaceUsersByRoleTypeMutex.Unlock()
	fake.GetSpaceUsersByRoleTypeStub = stub
}

func (fake *FakeSpaceUsersActor) GetSpaceUsersByRoleTypeArgsForCall(i int) string {
	fake.getSpaceUsersByRoleTypeMutex.RLock()
	defer fake.getSpaceUsersByRoleTypeMutex.RUnlock()
	argsForCall := fake.getSpaceUsersByRoleTypeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSpaceUsersActor) GetSpaceUsersByRoleTypeReturns(result1 map[constant.RoleType][]v7action.User, result2 v7action.Warnings, result3 error) {
	fake.getSpaceUsersByRoleTypeMutex.Lock()
	defer fake.getSpaceUsersByRoleTypeMutex.Unlock()
	fake.GetSpaceUsersByRoleTypeStub = nil
	fake.getSpaceUsersByRoleTypeReturns = struct {
		result1 map[constant.RoleType][]v7action.User
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceUsersActor) GetSpaceUsersByRoleTypeReturnsOnCall(i int, result1 map[constant.RoleType][]v7action.User, result2 v7action.Warnings, result3 error) {
	fake.getSpaceUsersByRoleTypeMutex.Lock()
	defer fake.getSpaceUsersByRoleTypeMutex.Unlock()
	fake.GetSpaceUsersByRoleTypeStub = nil
	if fake.getSpaceUsersByRoleTypeReturnsOnCall == nil {
		fake.getSpaceUsersByRoleTypeReturnsOnCall = make(map[int]struct {
			result1 map[constant.RoleType][]v7action.User
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getSpaceUsersByRoleTypeReturnsOnCall[i] = struct {
		result1 map[constant.RoleType][]v7action.User
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceUsersActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	fake.getSpaceUsersByRoleTypeMutex.RLock()
	defer fake.getSpaceUsersByRoleTypeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSpaceUsersActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.SpaceUsersActor = new(FakeSpaceUsersActor)