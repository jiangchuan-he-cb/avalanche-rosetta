// Code generated by mockery v2.20.2. DO NOT EDIT.

package chain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/coinbase/rosetta-sdk-go/types"
)

// AccountBackend is an autogenerated mock type for the AccountBackend type
type AccountBackend struct {
	mock.Mock
}

// AccountBalance provides a mock function with given fields: ctx, req
func (_m *AccountBackend) AccountBalance(ctx context.Context, req *types.AccountBalanceRequest) (*types.AccountBalanceResponse, *types.Error) {
	ret := _m.Called(ctx, req)

	var r0 *types.AccountBalanceResponse
	var r1 *types.Error
	if rf, ok := ret.Get(0).(func(context.Context, *types.AccountBalanceRequest) (*types.AccountBalanceResponse, *types.Error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.AccountBalanceRequest) *types.AccountBalanceResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.AccountBalanceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.AccountBalanceRequest) *types.Error); ok {
		r1 = rf(ctx, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Error)
		}
	}

	return r0, r1
}

// AccountCoins provides a mock function with given fields: ctx, req
func (_m *AccountBackend) AccountCoins(ctx context.Context, req *types.AccountCoinsRequest) (*types.AccountCoinsResponse, *types.Error) {
	ret := _m.Called(ctx, req)

	var r0 *types.AccountCoinsResponse
	var r1 *types.Error
	if rf, ok := ret.Get(0).(func(context.Context, *types.AccountCoinsRequest) (*types.AccountCoinsResponse, *types.Error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.AccountCoinsRequest) *types.AccountCoinsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.AccountCoinsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.AccountCoinsRequest) *types.Error); ok {
		r1 = rf(ctx, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Error)
		}
	}

	return r0, r1
}

// ShouldHandleRequest provides a mock function with given fields: req
func (_m *AccountBackend) ShouldHandleRequest(req interface{}) bool {
	ret := _m.Called(req)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}) bool); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewAccountBackend interface {
	mock.TestingT
	Cleanup(func())
}

// NewAccountBackend creates a new instance of AccountBackend. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAccountBackend(t mockConstructorTestingTNewAccountBackend) *AccountBackend {
	mock := &AccountBackend{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}