// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	fs "io/fs"

	mock "github.com/stretchr/testify/mock"
)

// FileManager is an autogenerated mock type for the FileManager type
type FileManager struct {
	mock.Mock
}

// Remove provides a mock function with given fields: path
func (_m *FileManager) Remove(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveAll provides a mock function with given fields: path
func (_m *FileManager) RemoveAll(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Write provides a mock function with given fields: path, value, mode
func (_m *FileManager) Write(path string, value string, mode fs.FileMode) error {
	ret := _m.Called(path, value, mode)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, fs.FileMode) error); ok {
		r0 = rf(path, value, mode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *FileManager) WriteBytes(path string, value []byte) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}