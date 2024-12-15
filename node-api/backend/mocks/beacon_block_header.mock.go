// Code generated by mockery v2.49.0. DO NOT EDIT.

package mocks

import (
	common "github.com/berachain/beacon-kit/primitives/common"
	math "github.com/berachain/beacon-kit/primitives/math"
	"github.com/berachain/beacon-kit/consensus-types/types"

	mock "github.com/stretchr/testify/mock"
)

// BeaconBlockHeader is an autogenerated mock type for the BeaconBlockHeader type
type BeaconBlockHeader struct {
	mock.Mock
}

type BeaconBlockHeader_Expecter struct {
	mock *mock.Mock
}

func (_m *BeaconBlockHeader) EXPECT() *BeaconBlockHeader_Expecter {
	return &BeaconBlockHeader_Expecter{mock: &_m.Mock}
}

// GetBodyRoot provides a mock function with given fields:
func (_m *BeaconBlockHeader) GetBodyRoot() common.Root {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBodyRoot")
	}

	var r0 common.Root
	if rf, ok := ret.Get(0).(func() common.Root); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Root)
		}
	}

	return r0
}

// BeaconBlockHeader_GetBodyRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBodyRoot'
type BeaconBlockHeader_GetBodyRoot_Call struct {
	*mock.Call
}

// GetBodyRoot is a helper method to define mock.On call
func (_e *BeaconBlockHeader_Expecter) GetBodyRoot() *BeaconBlockHeader_GetBodyRoot_Call {
	return &BeaconBlockHeader_GetBodyRoot_Call{Call: _e.mock.On("GetBodyRoot")}
}

func (_c *BeaconBlockHeader_GetBodyRoot_Call) Run(run func()) *BeaconBlockHeader_GetBodyRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlockHeader_GetBodyRoot_Call) Return(_a0 common.Root) *BeaconBlockHeader_GetBodyRoot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_GetBodyRoot_Call) RunAndReturn(run func() common.Root) *BeaconBlockHeader_GetBodyRoot_Call {
	_c.Call.Return(run)
	return _c
}

// GetParentBlockRoot provides a mock function with given fields:
func (_m *BeaconBlockHeader) GetParentBlockRoot() common.Root {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetParentBlockRoot")
	}

	var r0 common.Root
	if rf, ok := ret.Get(0).(func() common.Root); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Root)
		}
	}

	return r0
}

// BeaconBlockHeader_GetParentBlockRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetParentBlockRoot'
type BeaconBlockHeader_GetParentBlockRoot_Call struct {
	*mock.Call
}

// GetParentBlockRoot is a helper method to define mock.On call
func (_e *BeaconBlockHeader_Expecter) GetParentBlockRoot() *BeaconBlockHeader_GetParentBlockRoot_Call {
	return &BeaconBlockHeader_GetParentBlockRoot_Call{Call: _e.mock.On("GetParentBlockRoot")}
}

func (_c *BeaconBlockHeader_GetParentBlockRoot_Call) Run(run func()) *BeaconBlockHeader_GetParentBlockRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlockHeader_GetParentBlockRoot_Call) Return(_a0 common.Root) *BeaconBlockHeader_GetParentBlockRoot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_GetParentBlockRoot_Call) RunAndReturn(run func() common.Root) *BeaconBlockHeader_GetParentBlockRoot_Call {
	_c.Call.Return(run)
	return _c
}

// GetProposerIndex provides a mock function with given fields:
func (_m *BeaconBlockHeader) GetProposerIndex() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetProposerIndex")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// BeaconBlockHeader_GetProposerIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProposerIndex'
type BeaconBlockHeader_GetProposerIndex_Call struct {
	*mock.Call
}

// GetProposerIndex is a helper method to define mock.On call
func (_e *BeaconBlockHeader_Expecter) GetProposerIndex() *BeaconBlockHeader_GetProposerIndex_Call {
	return &BeaconBlockHeader_GetProposerIndex_Call{Call: _e.mock.On("GetProposerIndex")}
}

func (_c *BeaconBlockHeader_GetProposerIndex_Call) Run(run func()) *BeaconBlockHeader_GetProposerIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlockHeader_GetProposerIndex_Call) Return(_a0 math.U64) *BeaconBlockHeader_GetProposerIndex_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_GetProposerIndex_Call) RunAndReturn(run func() math.U64) *BeaconBlockHeader_GetProposerIndex_Call {
	_c.Call.Return(run)
	return _c
}

// GetSlot provides a mock function with given fields:
func (_m *BeaconBlockHeader) GetSlot() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSlot")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// BeaconBlockHeader_GetSlot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSlot'
type BeaconBlockHeader_GetSlot_Call struct {
	*mock.Call
}

// GetSlot is a helper method to define mock.On call
func (_e *BeaconBlockHeader_Expecter) GetSlot() *BeaconBlockHeader_GetSlot_Call {
	return &BeaconBlockHeader_GetSlot_Call{Call: _e.mock.On("GetSlot")}
}

func (_c *BeaconBlockHeader_GetSlot_Call) Run(run func()) *BeaconBlockHeader_GetSlot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlockHeader_GetSlot_Call) Return(_a0 math.U64) *BeaconBlockHeader_GetSlot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_GetSlot_Call) RunAndReturn(run func() math.U64) *BeaconBlockHeader_GetSlot_Call {
	_c.Call.Return(run)
	return _c
}

// GetStateRoot provides a mock function with given fields:
func (_m *BeaconBlockHeader) GetStateRoot() common.Root {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStateRoot")
	}

	var r0 common.Root
	if rf, ok := ret.Get(0).(func() common.Root); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Root)
		}
	}

	return r0
}

// BeaconBlockHeader_GetStateRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStateRoot'
type BeaconBlockHeader_GetStateRoot_Call struct {
	*mock.Call
}

// GetStateRoot is a helper method to define mock.On call
func (_e *BeaconBlockHeader_Expecter) GetStateRoot() *BeaconBlockHeader_GetStateRoot_Call {
	return &BeaconBlockHeader_GetStateRoot_Call{Call: _e.mock.On("GetStateRoot")}
}

func (_c *BeaconBlockHeader_GetStateRoot_Call) Run(run func()) *BeaconBlockHeader_GetStateRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlockHeader_GetStateRoot_Call) Return(_a0 common.Root) *BeaconBlockHeader_GetStateRoot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_GetStateRoot_Call) RunAndReturn(run func() common.Root) *BeaconBlockHeader_GetStateRoot_Call {
	_c.Call.Return(run)
	return _c
}

// HashTreeRoot provides a mock function with given fields:
func (_m *BeaconBlockHeader) HashTreeRoot() common.Root {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HashTreeRoot")
	}

	var r0 common.Root
	if rf, ok := ret.Get(0).(func() common.Root); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Root)
		}
	}

	return r0
}

// BeaconBlockHeader_HashTreeRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HashTreeRoot'
type BeaconBlockHeader_HashTreeRoot_Call struct {
	*mock.Call
}

// HashTreeRoot is a helper method to define mock.On call
func (_e *BeaconBlockHeader_Expecter) HashTreeRoot() *BeaconBlockHeader_HashTreeRoot_Call {
	return &BeaconBlockHeader_HashTreeRoot_Call{Call: _e.mock.On("HashTreeRoot")}
}

func (_c *BeaconBlockHeader_HashTreeRoot_Call) Run(run func()) *BeaconBlockHeader_HashTreeRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlockHeader_HashTreeRoot_Call) Return(_a0 common.Root) *BeaconBlockHeader_HashTreeRoot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_HashTreeRoot_Call) RunAndReturn(run func() common.Root) *BeaconBlockHeader_HashTreeRoot_Call {
	_c.Call.Return(run)
	return _c
}

// MarshalSSZ provides a mock function with given fields:
func (_m *BeaconBlockHeader) MarshalSSZ() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MarshalSSZ")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeaconBlockHeader_MarshalSSZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarshalSSZ'
type BeaconBlockHeader_MarshalSSZ_Call struct {
	*mock.Call
}

// MarshalSSZ is a helper method to define mock.On call
func (_e *BeaconBlockHeader_Expecter) MarshalSSZ() *BeaconBlockHeader_MarshalSSZ_Call {
	return &BeaconBlockHeader_MarshalSSZ_Call{Call: _e.mock.On("MarshalSSZ")}
}

func (_c *BeaconBlockHeader_MarshalSSZ_Call) Run(run func()) *BeaconBlockHeader_MarshalSSZ_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlockHeader_MarshalSSZ_Call) Return(_a0 []byte, _a1 error) *BeaconBlockHeader_MarshalSSZ_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BeaconBlockHeader_MarshalSSZ_Call) RunAndReturn(run func() ([]byte, error)) *BeaconBlockHeader_MarshalSSZ_Call {
	_c.Call.Return(run)
	return _c
}

// New provides a mock function with given fields: slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot
func (_m *BeaconBlockHeader) New(slot math.U64, proposerIndex math.U64, parentBlockRoot common.Root, stateRoot common.Root, bodyRoot common.Root) *types.BeaconBlockHeader {
	ret := _m.Called(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)

	if len(ret) == 0 {
		panic("no return value specified for New")
	}

	var r0 *types.BeaconBlockHeader
	if rf, ok := ret.Get(0).(func(math.U64, math.U64, common.Root, common.Root, common.Root) *types.BeaconBlockHeader); ok {
		r0 = rf(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BeaconBlockHeader)
		}
	}

	return r0
}

// BeaconBlockHeader_New_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'New'
type BeaconBlockHeader_New_Call struct {
	*mock.Call
}

// New is a helper method to define mock.On call
//   - slot math.U64
//   - proposerIndex math.U64
//   - parentBlockRoot common.Root
//   - stateRoot common.Root
//   - bodyRoot common.Root
func (_e *BeaconBlockHeader_Expecter) New(slot interface{}, proposerIndex interface{}, parentBlockRoot interface{}, stateRoot interface{}, bodyRoot interface{}) *BeaconBlockHeader_New_Call {
	return &BeaconBlockHeader_New_Call{Call: _e.mock.On("New", slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)}
}

func (_c *BeaconBlockHeader_New_Call) Run(run func(slot math.U64, proposerIndex math.U64, parentBlockRoot common.Root, stateRoot common.Root, bodyRoot common.Root)) *BeaconBlockHeader_New_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(math.U64), args[1].(math.U64), args[2].(common.Root), args[3].(common.Root), args[4].(common.Root))
	})
	return _c
}

func (_c *BeaconBlockHeader_New_Call) Return(_a0 *types.BeaconBlockHeader) *BeaconBlockHeader_New_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_New_Call) RunAndReturn(run func(math.U64, math.U64, common.Root, common.Root, common.Root) *types.BeaconBlockHeader) *BeaconBlockHeader_New_Call {
	_c.Call.Return(run)
	return _c
}

// SetStateRoot provides a mock function with given fields: _a0
func (_m *BeaconBlockHeader) SetStateRoot(_a0 common.Root) {
	_m.Called(_a0)
}

// BeaconBlockHeader_SetStateRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetStateRoot'
type BeaconBlockHeader_SetStateRoot_Call struct {
	*mock.Call
}

// SetStateRoot is a helper method to define mock.On call
//   - _a0 common.Root
func (_e *BeaconBlockHeader_Expecter) SetStateRoot(_a0 interface{}) *BeaconBlockHeader_SetStateRoot_Call {
	return &BeaconBlockHeader_SetStateRoot_Call{Call: _e.mock.On("SetStateRoot", _a0)}
}

func (_c *BeaconBlockHeader_SetStateRoot_Call) Run(run func(_a0 common.Root)) *BeaconBlockHeader_SetStateRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Root))
	})
	return _c
}

func (_c *BeaconBlockHeader_SetStateRoot_Call) Return() *BeaconBlockHeader_SetStateRoot_Call {
	_c.Call.Return()
	return _c
}

func (_c *BeaconBlockHeader_SetStateRoot_Call) RunAndReturn(run func(common.Root)) *BeaconBlockHeader_SetStateRoot_Call {
	_c.Call.Return(run)
	return _c
}

// UnmarshalSSZ provides a mock function with given fields: _a0
func (_m *BeaconBlockHeader) UnmarshalSSZ(_a0 []byte) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UnmarshalSSZ")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BeaconBlockHeader_UnmarshalSSZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnmarshalSSZ'
type BeaconBlockHeader_UnmarshalSSZ_Call struct {
	*mock.Call
}

// UnmarshalSSZ is a helper method to define mock.On call
//   - _a0 []byte
func (_e *BeaconBlockHeader_Expecter) UnmarshalSSZ(_a0 interface{}) *BeaconBlockHeader_UnmarshalSSZ_Call {
	return &BeaconBlockHeader_UnmarshalSSZ_Call{Call: _e.mock.On("UnmarshalSSZ", _a0)}
}

func (_c *BeaconBlockHeader_UnmarshalSSZ_Call) Run(run func(_a0 []byte)) *BeaconBlockHeader_UnmarshalSSZ_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *BeaconBlockHeader_UnmarshalSSZ_Call) Return(_a0 error) *BeaconBlockHeader_UnmarshalSSZ_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_UnmarshalSSZ_Call) RunAndReturn(run func([]byte) error) *BeaconBlockHeader_UnmarshalSSZ_Call {
	_c.Call.Return(run)
	return _c
}

// NewBeaconBlockHeader creates a new instance of BeaconBlockHeader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBeaconBlockHeader(t interface {
	mock.TestingT
	Cleanup(func())
}) *BeaconBlockHeader {
	mock := &BeaconBlockHeader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}