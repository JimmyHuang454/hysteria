// Code generated by mockery v2.43.0. DO NOT EDIT.

package server

import (
	protocol "github.com/apernet/hysteria/core/v2/international/protocol"
	mock "github.com/stretchr/testify/mock"
)

// mockUDPIO is an autogenerated mock type for the udpIO type
type mockUDPIO struct {
	mock.Mock
}

type mockUDPIO_Expecter struct {
	mock *mock.Mock
}

func (_m *mockUDPIO) EXPECT() *mockUDPIO_Expecter {
	return &mockUDPIO_Expecter{mock: &_m.Mock}
}

// ReceiveMessage provides a mock function with given fields:
func (_m *mockUDPIO) ReceiveMessage() (*protocol.UDPMessage, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReceiveMessage")
	}

	var r0 *protocol.UDPMessage
	var r1 error
	if rf, ok := ret.Get(0).(func() (*protocol.UDPMessage, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *protocol.UDPMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protocol.UDPMessage)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockUDPIO_ReceiveMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReceiveMessage'
type mockUDPIO_ReceiveMessage_Call struct {
	*mock.Call
}

// ReceiveMessage is a helper method to define mock.On call
func (_e *mockUDPIO_Expecter) ReceiveMessage() *mockUDPIO_ReceiveMessage_Call {
	return &mockUDPIO_ReceiveMessage_Call{Call: _e.mock.On("ReceiveMessage")}
}

func (_c *mockUDPIO_ReceiveMessage_Call) Run(run func()) *mockUDPIO_ReceiveMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockUDPIO_ReceiveMessage_Call) Return(_a0 *protocol.UDPMessage, _a1 error) *mockUDPIO_ReceiveMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockUDPIO_ReceiveMessage_Call) RunAndReturn(run func() (*protocol.UDPMessage, error)) *mockUDPIO_ReceiveMessage_Call {
	_c.Call.Return(run)
	return _c
}

// SendMessage provides a mock function with given fields: _a0, _a1
func (_m *mockUDPIO) SendMessage(_a0 []byte, _a1 *protocol.UDPMessage) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SendMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, *protocol.UDPMessage) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockUDPIO_SendMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMessage'
type mockUDPIO_SendMessage_Call struct {
	*mock.Call
}

// SendMessage is a helper method to define mock.On call
//   - _a0 []byte
//   - _a1 *protocol.UDPMessage
func (_e *mockUDPIO_Expecter) SendMessage(_a0 interface{}, _a1 interface{}) *mockUDPIO_SendMessage_Call {
	return &mockUDPIO_SendMessage_Call{Call: _e.mock.On("SendMessage", _a0, _a1)}
}

func (_c *mockUDPIO_SendMessage_Call) Run(run func(_a0 []byte, _a1 *protocol.UDPMessage)) *mockUDPIO_SendMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(*protocol.UDPMessage))
	})
	return _c
}

func (_c *mockUDPIO_SendMessage_Call) Return(_a0 error) *mockUDPIO_SendMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockUDPIO_SendMessage_Call) RunAndReturn(run func([]byte, *protocol.UDPMessage) error) *mockUDPIO_SendMessage_Call {
	_c.Call.Return(run)
	return _c
}

// UDP provides a mock function with given fields: reqAddr
func (_m *mockUDPIO) UDP(reqAddr string) (UDPConn, error) {
	ret := _m.Called(reqAddr)

	if len(ret) == 0 {
		panic("no return value specified for UDP")
	}

	var r0 UDPConn
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (UDPConn, error)); ok {
		return rf(reqAddr)
	}
	if rf, ok := ret.Get(0).(func(string) UDPConn); ok {
		r0 = rf(reqAddr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(UDPConn)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(reqAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockUDPIO_UDP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UDP'
type mockUDPIO_UDP_Call struct {
	*mock.Call
}

// UDP is a helper method to define mock.On call
//   - reqAddr string
func (_e *mockUDPIO_Expecter) UDP(reqAddr interface{}) *mockUDPIO_UDP_Call {
	return &mockUDPIO_UDP_Call{Call: _e.mock.On("UDP", reqAddr)}
}

func (_c *mockUDPIO_UDP_Call) Run(run func(reqAddr string)) *mockUDPIO_UDP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *mockUDPIO_UDP_Call) Return(_a0 UDPConn, _a1 error) *mockUDPIO_UDP_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockUDPIO_UDP_Call) RunAndReturn(run func(string) (UDPConn, error)) *mockUDPIO_UDP_Call {
	_c.Call.Return(run)
	return _c
}

// newMockUDPIO creates a new instance of mockUDPIO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockUDPIO(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockUDPIO {
	mock := &mockUDPIO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
