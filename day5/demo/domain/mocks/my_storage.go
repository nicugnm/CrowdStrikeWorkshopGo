// Code generated by MockGen. DO NOT EDIT.
// Source: my_storage.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMyStorage is a mock of MyStorage interface.
type MockMyStorage struct {
	ctrl     *gomock.Controller
	recorder *MockMyStorageMockRecorder
}

// MockMyStorageMockRecorder is the mock recorder for MockMyStorage.
type MockMyStorageMockRecorder struct {
	mock *MockMyStorage
}

// NewMockMyStorage creates a new mock instance.
func NewMockMyStorage(ctrl *gomock.Controller) *MockMyStorage {
	mock := &MockMyStorage{ctrl: ctrl}
	mock.recorder = &MockMyStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMyStorage) EXPECT() *MockMyStorageMockRecorder {
	return m.recorder
}

// GetMyContent mocks base method.
func (m *MockMyStorage) GetMyContent(id string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMyContent", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyContent indicates an expected call of GetMyContent.
func (mr *MockMyStorageMockRecorder) GetMyContent(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyContent", reflect.TypeOf((*MockMyStorage)(nil).GetMyContent), id)
}

// WriteContent mocks base method.
func (m *MockMyStorage) WriteMyContent(id, content string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMyContent", id, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteContent indicates an expected call of WriteContent.
func (mr *MockMyStorageMockRecorder) WriteMyContent(id, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMyContent", reflect.TypeOf((*MockMyStorage)(nil).WriteMyContent), id, content)
}
