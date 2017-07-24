// This file was generated by counterfeiter
package storefakes

import (
	"database/sql"
	"sync"
)

type FakeResult struct {
	LastInsertIdStub        func() (int64, error)
	lastInsertIdMutex       sync.RWMutex
	lastInsertIdArgsForCall []struct{}
	lastInsertIdReturns     struct {
		result1 int64
		result2 error
	}
	RowsAffectedStub        func() (int64, error)
	rowsAffectedMutex       sync.RWMutex
	rowsAffectedArgsForCall []struct{}
	rowsAffectedReturns     struct {
		result1 int64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResult) LastInsertId() (int64, error) {
	fake.lastInsertIdMutex.Lock()
	fake.lastInsertIdArgsForCall = append(fake.lastInsertIdArgsForCall, struct{}{})
	fake.recordInvocation("LastInsertId", []interface{}{})
	fake.lastInsertIdMutex.Unlock()
	if fake.LastInsertIdStub != nil {
		return fake.LastInsertIdStub()
	} else {
		return fake.lastInsertIdReturns.result1, fake.lastInsertIdReturns.result2
	}
}

func (fake *FakeResult) LastInsertIdCallCount() int {
	fake.lastInsertIdMutex.RLock()
	defer fake.lastInsertIdMutex.RUnlock()
	return len(fake.lastInsertIdArgsForCall)
}

func (fake *FakeResult) LastInsertIdReturns(result1 int64, result2 error) {
	fake.LastInsertIdStub = nil
	fake.lastInsertIdReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeResult) RowsAffected() (int64, error) {
	fake.rowsAffectedMutex.Lock()
	fake.rowsAffectedArgsForCall = append(fake.rowsAffectedArgsForCall, struct{}{})
	fake.recordInvocation("RowsAffected", []interface{}{})
	fake.rowsAffectedMutex.Unlock()
	if fake.RowsAffectedStub != nil {
		return fake.RowsAffectedStub()
	} else {
		return fake.rowsAffectedReturns.result1, fake.rowsAffectedReturns.result2
	}
}

func (fake *FakeResult) RowsAffectedCallCount() int {
	fake.rowsAffectedMutex.RLock()
	defer fake.rowsAffectedMutex.RUnlock()
	return len(fake.rowsAffectedArgsForCall)
}

func (fake *FakeResult) RowsAffectedReturns(result1 int64, result2 error) {
	fake.RowsAffectedStub = nil
	fake.rowsAffectedReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeResult) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lastInsertIdMutex.RLock()
	defer fake.lastInsertIdMutex.RUnlock()
	fake.rowsAffectedMutex.RLock()
	defer fake.rowsAffectedMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeResult) recordInvocation(key string, args []interface{}) {
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

var _ sql.Result = new(FakeResult)