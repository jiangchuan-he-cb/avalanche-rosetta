// Code generated by mockery v2.20.2. DO NOT EDIT.

package chain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/coinbase/rosetta-sdk-go/types"
)

// NetworkBackend is an autogenerated mock type for the NetworkBackend type
type NetworkBackend struct {
	mock.Mock
}

// NetworkIdentifier provides a mock function with given fields:
func (_m *NetworkBackend) NetworkIdentifier() *types.NetworkIdentifier {
	ret := _m.Called()

	var r0 *types.NetworkIdentifier
	if rf, ok := ret.Get(0).(func() *types.NetworkIdentifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NetworkIdentifier)
		}
	}

	return r0
}

// NetworkOptions provides a mock function with given fields: ctx, request
func (_m *NetworkBackend) NetworkOptions(ctx context.Context, request *types.NetworkRequest) (*types.NetworkOptionsResponse, *types.Error) {
	ret := _m.Called(ctx, request)

	var r0 *types.NetworkOptionsResponse
	var r1 *types.Error
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkRequest) (*types.NetworkOptionsResponse, *types.Error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkRequest) *types.NetworkOptionsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NetworkOptionsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkRequest) *types.Error); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Error)
		}
	}

	return r0, r1
}

// NetworkStatus provides a mock function with given fields: ctx, request
func (_m *NetworkBackend) NetworkStatus(ctx context.Context, request *types.NetworkRequest) (*types.NetworkStatusResponse, *types.Error) {
	ret := _m.Called(ctx, request)

	var r0 *types.NetworkStatusResponse
	var r1 *types.Error
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkRequest) (*types.NetworkStatusResponse, *types.Error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkRequest) *types.NetworkStatusResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NetworkStatusResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkRequest) *types.Error); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Error)
		}
	}

	return r0, r1
}

// ShouldHandleRequest provides a mock function with given fields: req
func (_m *NetworkBackend) ShouldHandleRequest(req interface{}) bool {
	ret := _m.Called(req)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}) bool); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewNetworkBackend interface {
	mock.TestingT
	Cleanup(func())
}

// NewNetworkBackend creates a new instance of NetworkBackend. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNetworkBackend(t mockConstructorTestingTNewNetworkBackend) *NetworkBackend {
	mock := &NetworkBackend{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
