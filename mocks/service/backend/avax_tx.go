// Code generated by mockery v2.20.2. DO NOT EDIT.

package chain

import (
	ids "github.com/ava-labs/avalanchego/ids"
	mock "github.com/stretchr/testify/mock"
)

// AvaxTx is an autogenerated mock type for the AvaxTx type
type AvaxTx struct {
	mock.Mock
}

// Hash provides a mock function with given fields:
func (_m *AvaxTx) Hash() ids.ID {
	ret := _m.Called()

	var r0 ids.ID
	if rf, ok := ret.Get(0).(func() ids.ID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	return r0
}

// Initialize provides a mock function with given fields:
func (_m *AvaxTx) Initialize() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Marshal provides a mock function with given fields:
func (_m *AvaxTx) Marshal() ([]byte, error) {
	ret := _m.Called()

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

// SigningPayload provides a mock function with given fields:
func (_m *AvaxTx) SigningPayload() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// Unmarshal provides a mock function with given fields: _a0
func (_m *AvaxTx) Unmarshal(_a0 []byte) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAvaxTx interface {
	mock.TestingT
	Cleanup(func())
}

// NewAvaxTx creates a new instance of AvaxTx. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAvaxTx(t mockConstructorTestingTNewAvaxTx) *AvaxTx {
	mock := &AvaxTx{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}