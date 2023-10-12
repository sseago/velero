// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	watch "k8s.io/apimachinery/pkg/watch"
)

// Dynamic is an autogenerated mock type for the Dynamic type
type Dynamic struct {
	mock.Mock
}

// Create provides a mock function with given fields: obj
func (_m *Dynamic) Create(obj *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	ret := _m.Called(obj)

	var r0 *unstructured.Unstructured
	var r1 error
	if rf, ok := ret.Get(0).(func(*unstructured.Unstructured) (*unstructured.Unstructured, error)); ok {
		return rf(obj)
	}
	if rf, ok := ret.Get(0).(func(*unstructured.Unstructured) *unstructured.Unstructured); ok {
		r0 = rf(obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*unstructured.Unstructured)
		}
	}

	if rf, ok := ret.Get(1).(func(*unstructured.Unstructured) error); ok {
		r1 = rf(obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name, opts
func (_m *Dynamic) Delete(name string, opts v1.DeleteOptions) error {
	ret := _m.Called(name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, v1.DeleteOptions) error); ok {
		r0 = rf(name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: name, opts
func (_m *Dynamic) Get(name string, opts v1.GetOptions) (*unstructured.Unstructured, error) {
	ret := _m.Called(name, opts)

	var r0 *unstructured.Unstructured
	var r1 error
	if rf, ok := ret.Get(0).(func(string, v1.GetOptions) (*unstructured.Unstructured, error)); ok {
		return rf(name, opts)
	}
	if rf, ok := ret.Get(0).(func(string, v1.GetOptions) *unstructured.Unstructured); ok {
		r0 = rf(name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*unstructured.Unstructured)
		}
	}

	if rf, ok := ret.Get(1).(func(string, v1.GetOptions) error); ok {
		r1 = rf(name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0
func (_m *Dynamic) List(_a0 v1.ListOptions) (*unstructured.UnstructuredList, error) {
	ret := _m.Called(_a0)

	var r0 *unstructured.UnstructuredList
	var r1 error
	if rf, ok := ret.Get(0).(func(v1.ListOptions) (*unstructured.UnstructuredList, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(v1.ListOptions) *unstructured.UnstructuredList); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*unstructured.UnstructuredList)
		}
	}

	if rf, ok := ret.Get(1).(func(v1.ListOptions) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: name, data
func (_m *Dynamic) Patch(name string, data []byte) (*unstructured.Unstructured, error) {
	ret := _m.Called(name, data)

	var r0 *unstructured.Unstructured
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []byte) (*unstructured.Unstructured, error)); ok {
		return rf(name, data)
	}
	if rf, ok := ret.Get(0).(func(string, []byte) *unstructured.Unstructured); ok {
		r0 = rf(name, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*unstructured.Unstructured)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []byte) error); ok {
		r1 = rf(name, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: obj, opts
func (_m *Dynamic) UpdateStatus(obj *unstructured.Unstructured, opts v1.UpdateOptions) (*unstructured.Unstructured, error) {
	ret := _m.Called(obj, opts)

	var r0 *unstructured.Unstructured
	var r1 error
	if rf, ok := ret.Get(0).(func(*unstructured.Unstructured, v1.UpdateOptions) (*unstructured.Unstructured, error)); ok {
		return rf(obj, opts)
	}
	if rf, ok := ret.Get(0).(func(*unstructured.Unstructured, v1.UpdateOptions) *unstructured.Unstructured); ok {
		r0 = rf(obj, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*unstructured.Unstructured)
		}
	}

	if rf, ok := ret.Get(1).(func(*unstructured.Unstructured, v1.UpdateOptions) error); ok {
		r1 = rf(obj, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: _a0
func (_m *Dynamic) Watch(_a0 v1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(_a0)

	var r0 watch.Interface
	var r1 error
	if rf, ok := ret.Get(0).(func(v1.ListOptions) (watch.Interface, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(v1.ListOptions) watch.Interface); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	if rf, ok := ret.Get(1).(func(v1.ListOptions) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDynamic creates a new instance of Dynamic. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDynamic(t interface {
	mock.TestingT
	Cleanup(func())
}) *Dynamic {
	mock := &Dynamic{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
