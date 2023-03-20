// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entitie "microservice/domain/entitie"

	mock "github.com/stretchr/testify/mock"
)

// TagRepository is an autogenerated mock type for the TagRepository type
type TagRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *TagRepository) Delete(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *TagRepository) GetByID(ctx context.Context, id uint) (entitie.Tag, error) {
	ret := _m.Called(ctx, id)

	var r0 entitie.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (entitie.Tag, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) entitie.Tag); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entitie.Tag)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: ctx, title
func (_m *TagRepository) GetByName(ctx context.Context, title string) (entitie.Tag, error) {
	ret := _m.Called(ctx, title)

	var r0 entitie.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entitie.Tag, error)); ok {
		return rf(ctx, title)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entitie.Tag); ok {
		r0 = rf(ctx, title)
	} else {
		r0 = ret.Get(0).(entitie.Tag)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, a
func (_m *TagRepository) Store(ctx context.Context, a *entitie.Tag) error {
	ret := _m.Called(ctx, a)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entitie.Tag) error); ok {
		r0 = rf(ctx, a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, ar
func (_m *TagRepository) Update(ctx context.Context, ar *entitie.Tag) error {
	ret := _m.Called(ctx, ar)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entitie.Tag) error); ok {
		r0 = rf(ctx, ar)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTagRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTagRepository creates a new instance of TagRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTagRepository(t mockConstructorTestingTNewTagRepository) *TagRepository {
	mock := &TagRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
