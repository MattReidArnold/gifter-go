// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	app "github.com/mattreidarnold/gifter/app"
	mock "github.com/stretchr/testify/mock"
)

// UnitOfWork is an autogenerated mock type for the UnitOfWork type
type UnitOfWork struct {
	mock.Mock
}

// Groups provides a mock function with given fields:
func (_m *UnitOfWork) Groups() app.GroupRepository {
	ret := _m.Called()

	var r0 app.GroupRepository
	if rf, ok := ret.Get(0).(func() app.GroupRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(app.GroupRepository)
		}
	}

	return r0
}