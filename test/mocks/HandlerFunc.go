// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	context "context"

	app "github.com/mattreidarnold/gifter/app"

	mock "github.com/stretchr/testify/mock"
)

// HandlerFunc is an autogenerated mock type for the HandlerFunc type
type HandlerFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *HandlerFunc) Execute(_a0 context.Context, _a1 app.Message) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, app.Message) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}