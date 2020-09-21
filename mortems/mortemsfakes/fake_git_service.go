// Code generated by counterfeiter. DO NOT EDIT.
package mortemsfakes

import (
	"sync"

	"github.com/ostenbom/morty/mortems"
)

type FakeGitService struct {
	CommitNewFilesStub        func(*mortems.RepoFiles) error
	commitNewFilesMutex       sync.RWMutex
	commitNewFilesArgsForCall []struct {
		arg1 *mortems.RepoFiles
	}
	commitNewFilesReturns struct {
		result1 error
	}
	commitNewFilesReturnsOnCall map[int]struct {
		result1 error
	}
	GetFilesStub        func() (*mortems.RepoFiles, error)
	getFilesMutex       sync.RWMutex
	getFilesArgsForCall []struct {
	}
	getFilesReturns struct {
		result1 *mortems.RepoFiles
		result2 error
	}
	getFilesReturnsOnCall map[int]struct {
		result1 *mortems.RepoFiles
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGitService) CommitNewFiles(arg1 *mortems.RepoFiles) error {
	fake.commitNewFilesMutex.Lock()
	ret, specificReturn := fake.commitNewFilesReturnsOnCall[len(fake.commitNewFilesArgsForCall)]
	fake.commitNewFilesArgsForCall = append(fake.commitNewFilesArgsForCall, struct {
		arg1 *mortems.RepoFiles
	}{arg1})
	fake.recordInvocation("CommitNewFiles", []interface{}{arg1})
	fake.commitNewFilesMutex.Unlock()
	if fake.CommitNewFilesStub != nil {
		return fake.CommitNewFilesStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.commitNewFilesReturns
	return fakeReturns.result1
}

func (fake *FakeGitService) CommitNewFilesCallCount() int {
	fake.commitNewFilesMutex.RLock()
	defer fake.commitNewFilesMutex.RUnlock()
	return len(fake.commitNewFilesArgsForCall)
}

func (fake *FakeGitService) CommitNewFilesCalls(stub func(*mortems.RepoFiles) error) {
	fake.commitNewFilesMutex.Lock()
	defer fake.commitNewFilesMutex.Unlock()
	fake.CommitNewFilesStub = stub
}

func (fake *FakeGitService) CommitNewFilesArgsForCall(i int) *mortems.RepoFiles {
	fake.commitNewFilesMutex.RLock()
	defer fake.commitNewFilesMutex.RUnlock()
	argsForCall := fake.commitNewFilesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGitService) CommitNewFilesReturns(result1 error) {
	fake.commitNewFilesMutex.Lock()
	defer fake.commitNewFilesMutex.Unlock()
	fake.CommitNewFilesStub = nil
	fake.commitNewFilesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitService) CommitNewFilesReturnsOnCall(i int, result1 error) {
	fake.commitNewFilesMutex.Lock()
	defer fake.commitNewFilesMutex.Unlock()
	fake.CommitNewFilesStub = nil
	if fake.commitNewFilesReturnsOnCall == nil {
		fake.commitNewFilesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.commitNewFilesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitService) GetFiles() (*mortems.RepoFiles, error) {
	fake.getFilesMutex.Lock()
	ret, specificReturn := fake.getFilesReturnsOnCall[len(fake.getFilesArgsForCall)]
	fake.getFilesArgsForCall = append(fake.getFilesArgsForCall, struct {
	}{})
	fake.recordInvocation("GetFiles", []interface{}{})
	fake.getFilesMutex.Unlock()
	if fake.GetFilesStub != nil {
		return fake.GetFilesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getFilesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitService) GetFilesCallCount() int {
	fake.getFilesMutex.RLock()
	defer fake.getFilesMutex.RUnlock()
	return len(fake.getFilesArgsForCall)
}

func (fake *FakeGitService) GetFilesCalls(stub func() (*mortems.RepoFiles, error)) {
	fake.getFilesMutex.Lock()
	defer fake.getFilesMutex.Unlock()
	fake.GetFilesStub = stub
}

func (fake *FakeGitService) GetFilesReturns(result1 *mortems.RepoFiles, result2 error) {
	fake.getFilesMutex.Lock()
	defer fake.getFilesMutex.Unlock()
	fake.GetFilesStub = nil
	fake.getFilesReturns = struct {
		result1 *mortems.RepoFiles
		result2 error
	}{result1, result2}
}

func (fake *FakeGitService) GetFilesReturnsOnCall(i int, result1 *mortems.RepoFiles, result2 error) {
	fake.getFilesMutex.Lock()
	defer fake.getFilesMutex.Unlock()
	fake.GetFilesStub = nil
	if fake.getFilesReturnsOnCall == nil {
		fake.getFilesReturnsOnCall = make(map[int]struct {
			result1 *mortems.RepoFiles
			result2 error
		})
	}
	fake.getFilesReturnsOnCall[i] = struct {
		result1 *mortems.RepoFiles
		result2 error
	}{result1, result2}
}

func (fake *FakeGitService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.commitNewFilesMutex.RLock()
	defer fake.commitNewFilesMutex.RUnlock()
	fake.getFilesMutex.RLock()
	defer fake.getFilesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGitService) recordInvocation(key string, args []interface{}) {
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

var _ mortems.GitService = new(FakeGitService)
