// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	jsonrpc "github.com/ybbus/jsonrpc"
)

// Caller is an autogenerated mock type for the Caller type
type Caller struct {
	mock.Mock
}

// Call provides a mock function with given fields: method, params
func (_m *Caller) Call(method string, params ...string) (*jsonrpc.RPCResponse, error) {
	_va := make([]interface{}, len(params))
	for _i := range params {
		_va[_i] = params[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, method)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *jsonrpc.RPCResponse
	if rf, ok := ret.Get(0).(func(string, ...string) *jsonrpc.RPCResponse); ok {
		r0 = rf(method, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jsonrpc.RPCResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...string) error); ok {
		r1 = rf(method, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}