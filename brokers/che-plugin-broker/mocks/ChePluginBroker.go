// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import jsonrpc "github.com/eclipse/che-go-jsonrpc"
import mock "github.com/stretchr/testify/mock"
import model "github.com/eclipse/che-plugin-broker/model"

// ChePluginBroker is an autogenerated mock type for the ChePluginBroker type
type ChePluginBroker struct {
	mock.Mock
}

// ProcessPlugin provides a mock function with given fields: meta
func (_m *ChePluginBroker) ProcessPlugin(meta model.PluginMeta) error {
	ret := _m.Called(meta)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.PluginMeta) error); ok {
		r0 = rf(meta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PushEvents provides a mock function with given fields: tun
func (_m *ChePluginBroker) PushEvents(tun *jsonrpc.Tunnel) {
	_m.Called(tun)
}

// Start provides a mock function with given fields: _a0
func (_m *ChePluginBroker) Start(_a0 []model.PluginMeta) {
	_m.Called(_a0)
}