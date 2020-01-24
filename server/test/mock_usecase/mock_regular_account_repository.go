// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/regular_account_repository.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/kindaidensan/UMR/domain"
	reflect "reflect"
)

// MockRegularAccountRepository is a mock of RegularAccountRepository interface
type MockRegularAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRegularAccountRepositoryMockRecorder
}

// MockRegularAccountRepositoryMockRecorder is the mock recorder for MockRegularAccountRepository
type MockRegularAccountRepositoryMockRecorder struct {
	mock *MockRegularAccountRepository
}

// NewMockRegularAccountRepository creates a new mock instance
func NewMockRegularAccountRepository(ctrl *gomock.Controller) *MockRegularAccountRepository {
	mock := &MockRegularAccountRepository{ctrl: ctrl}
	mock.recorder = &MockRegularAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegularAccountRepository) EXPECT() *MockRegularAccountRepositoryMockRecorder {
	return m.recorder
}

// TemporaryStore mocks base method
func (m *MockRegularAccountRepository) TemporaryStore(arg0 domain.RegularAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TemporaryStore", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TemporaryStore indicates an expected call of TemporaryStore
func (mr *MockRegularAccountRepositoryMockRecorder) TemporaryStore(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TemporaryStore", reflect.TypeOf((*MockRegularAccountRepository)(nil).TemporaryStore), arg0)
}

// FindByIdFromTemporary mocks base method
func (m *MockRegularAccountRepository) FindByIdFromTemporary(arg0 string) (domain.RegularAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIdFromTemporary", arg0)
	ret0, _ := ret[0].(domain.RegularAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdFromTemporary indicates an expected call of FindByIdFromTemporary
func (mr *MockRegularAccountRepositoryMockRecorder) FindByIdFromTemporary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdFromTemporary", reflect.TypeOf((*MockRegularAccountRepository)(nil).FindByIdFromTemporary), arg0)
}

// Store mocks base method
func (m *MockRegularAccountRepository) Store(arg0 domain.RegularAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store
func (mr *MockRegularAccountRepositoryMockRecorder) Store(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockRegularAccountRepository)(nil).Store), arg0)
}
