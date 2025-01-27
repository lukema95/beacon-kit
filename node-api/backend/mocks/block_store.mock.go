// Code generated by mockery v2.49.0. DO NOT EDIT.

package mocks

import (
	common "github.com/berachain/beacon-kit/primitives/common"
	math "github.com/berachain/beacon-kit/primitives/math"

	mock "github.com/stretchr/testify/mock"
)

// BlockStore is an autogenerated mock type for the BlockStore type
type BlockStore struct {
	mock.Mock
}

type BlockStore_Expecter struct {
	mock *mock.Mock
}

func (_m *BlockStore) EXPECT() *BlockStore_Expecter {
	return &BlockStore_Expecter{mock: &_m.Mock}
}

// GetParentSlotByTimestamp provides a mock function with given fields: timestamp
func (_m *BlockStore) GetParentSlotByTimestamp(timestamp math.U64) (math.U64, error) {
	ret := _m.Called(timestamp)

	if len(ret) == 0 {
		panic("no return value specified for GetParentSlotByTimestamp")
	}

	var r0 math.U64
	var r1 error
	if rf, ok := ret.Get(0).(func(math.U64) (math.U64, error)); ok {
		return rf(timestamp)
	}
	if rf, ok := ret.Get(0).(func(math.U64) math.U64); ok {
		r0 = rf(timestamp)
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	if rf, ok := ret.Get(1).(func(math.U64) error); ok {
		r1 = rf(timestamp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockStore_GetParentSlotByTimestamp_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetParentSlotByTimestamp'
type BlockStore_GetParentSlotByTimestamp_Call struct {
	*mock.Call
}

// GetParentSlotByTimestamp is a helper method to define mock.On call
//   - timestamp math.U64
func (_e *BlockStore_Expecter) GetParentSlotByTimestamp(timestamp interface{}) *BlockStore_GetParentSlotByTimestamp_Call {
	return &BlockStore_GetParentSlotByTimestamp_Call{Call: _e.mock.On("GetParentSlotByTimestamp", timestamp)}
}

func (_c *BlockStore_GetParentSlotByTimestamp_Call) Run(run func(timestamp math.U64)) *BlockStore_GetParentSlotByTimestamp_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(math.U64))
	})
	return _c
}

func (_c *BlockStore_GetParentSlotByTimestamp_Call) Return(_a0 math.U64, _a1 error) *BlockStore_GetParentSlotByTimestamp_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlockStore_GetParentSlotByTimestamp_Call) RunAndReturn(run func(math.U64) (math.U64, error)) *BlockStore_GetParentSlotByTimestamp_Call {
	_c.Call.Return(run)
	return _c
}

// GetSlotByBlockRoot provides a mock function with given fields: root
func (_m *BlockStore) GetSlotByBlockRoot(root common.Root) (math.U64, error) {
	ret := _m.Called(root)

	if len(ret) == 0 {
		panic("no return value specified for GetSlotByBlockRoot")
	}

	var r0 math.U64
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Root) (math.U64, error)); ok {
		return rf(root)
	}
	if rf, ok := ret.Get(0).(func(common.Root) math.U64); ok {
		r0 = rf(root)
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	if rf, ok := ret.Get(1).(func(common.Root) error); ok {
		r1 = rf(root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockStore_GetSlotByBlockRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSlotByBlockRoot'
type BlockStore_GetSlotByBlockRoot_Call struct {
	*mock.Call
}

// GetSlotByBlockRoot is a helper method to define mock.On call
//   - root common.Root
func (_e *BlockStore_Expecter) GetSlotByBlockRoot(root interface{}) *BlockStore_GetSlotByBlockRoot_Call {
	return &BlockStore_GetSlotByBlockRoot_Call{Call: _e.mock.On("GetSlotByBlockRoot", root)}
}

func (_c *BlockStore_GetSlotByBlockRoot_Call) Run(run func(root common.Root)) *BlockStore_GetSlotByBlockRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Root))
	})
	return _c
}

func (_c *BlockStore_GetSlotByBlockRoot_Call) Return(_a0 math.U64, _a1 error) *BlockStore_GetSlotByBlockRoot_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlockStore_GetSlotByBlockRoot_Call) RunAndReturn(run func(common.Root) (math.U64, error)) *BlockStore_GetSlotByBlockRoot_Call {
	_c.Call.Return(run)
	return _c
}

// GetSlotByStateRoot provides a mock function with given fields: root
func (_m *BlockStore) GetSlotByStateRoot(root common.Root) (math.U64, error) {
	ret := _m.Called(root)

	if len(ret) == 0 {
		panic("no return value specified for GetSlotByStateRoot")
	}

	var r0 math.U64
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Root) (math.U64, error)); ok {
		return rf(root)
	}
	if rf, ok := ret.Get(0).(func(common.Root) math.U64); ok {
		r0 = rf(root)
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	if rf, ok := ret.Get(1).(func(common.Root) error); ok {
		r1 = rf(root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockStore_GetSlotByStateRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSlotByStateRoot'
type BlockStore_GetSlotByStateRoot_Call struct {
	*mock.Call
}

// GetSlotByStateRoot is a helper method to define mock.On call
//   - root common.Root
func (_e *BlockStore_Expecter) GetSlotByStateRoot(root interface{}) *BlockStore_GetSlotByStateRoot_Call {
	return &BlockStore_GetSlotByStateRoot_Call{Call: _e.mock.On("GetSlotByStateRoot", root)}
}

func (_c *BlockStore_GetSlotByStateRoot_Call) Run(run func(root common.Root)) *BlockStore_GetSlotByStateRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Root))
	})
	return _c
}

func (_c *BlockStore_GetSlotByStateRoot_Call) Return(_a0 math.U64, _a1 error) *BlockStore_GetSlotByStateRoot_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlockStore_GetSlotByStateRoot_Call) RunAndReturn(run func(common.Root) (math.U64, error)) *BlockStore_GetSlotByStateRoot_Call {
	_c.Call.Return(run)
	return _c
}

// NewBlockStore creates a new instance of BlockStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlockStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlockStore {
	mock := &BlockStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
