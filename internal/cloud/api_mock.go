// Code generated by MockGen. DO NOT EDIT.
// Source: ./types.go

// Package cloud is a generated GoMock package.
package cloud

import (
	reflect "reflect"

	types "github.com/api7/cloud-cli/internal/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// ListControlPlanes mocks base method.
func (m *MockAPI) ListControlPlanes(orgID string) ([]*types.ControlPlaneSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListControlPlanes", orgID)
	ret0, _ := ret[0].([]*types.ControlPlaneSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListControlPlanes indicates an expected call of ListControlPlanes.
func (mr *MockAPIMockRecorder) ListControlPlanes(orgID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListControlPlanes", reflect.TypeOf((*MockAPI)(nil).ListControlPlanes), orgID)
}

// Me mocks base method.
func (m *MockAPI) Me() (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Me")
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Me indicates an expected call of Me.
func (mr *MockAPIMockRecorder) Me() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Me", reflect.TypeOf((*MockAPI)(nil).Me))
}
