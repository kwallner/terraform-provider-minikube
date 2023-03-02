// Code generated by MockGen. DO NOT EDIT.
// Source: minikube_client.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	config "k8s.io/minikube/pkg/minikube/config"
	kubeconfig "k8s.io/minikube/pkg/minikube/kubeconfig"
)

// MockClusterClient is a mock of ClusterClient interface.
type MockClusterClient struct {
	ctrl     *gomock.Controller
	recorder *MockClusterClientMockRecorder
}

// MockClusterClientMockRecorder is the mock recorder for MockClusterClient.
type MockClusterClientMockRecorder struct {
	mock *MockClusterClient
}

// NewMockClusterClient creates a new mock instance.
func NewMockClusterClient(ctrl *gomock.Controller) *MockClusterClient {
	mock := &MockClusterClient{ctrl: ctrl}
	mock.recorder = &MockClusterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterClient) EXPECT() *MockClusterClientMockRecorder {
	return m.recorder
}

// ApplyAddons mocks base method.
func (m *MockClusterClient) ApplyAddons(addons []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyAddons", addons)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyAddons indicates an expected call of ApplyAddons.
func (mr *MockClusterClientMockRecorder) ApplyAddons(addons interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyAddons", reflect.TypeOf((*MockClusterClient)(nil).ApplyAddons), addons)
}

// Delete mocks base method.
func (m *MockClusterClient) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockClusterClientMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterClient)(nil).Delete))
}

// GetClusterConfig mocks base method.
func (m *MockClusterClient) GetClusterConfig() *config.ClusterConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterConfig")
	ret0, _ := ret[0].(*config.ClusterConfig)
	return ret0
}

// GetClusterConfig indicates an expected call of GetClusterConfig.
func (mr *MockClusterClientMockRecorder) GetClusterConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterConfig", reflect.TypeOf((*MockClusterClient)(nil).GetClusterConfig))
}

// GetConfig mocks base method.
func (m *MockClusterClient) GetConfig() MinikubeClientConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(MinikubeClientConfig)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockClusterClientMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockClusterClient)(nil).GetConfig))
}

// GetK8sVersion mocks base method.
func (m *MockClusterClient) GetK8sVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetK8sVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetK8sVersion indicates an expected call of GetK8sVersion.
func (mr *MockClusterClientMockRecorder) GetK8sVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetK8sVersion", reflect.TypeOf((*MockClusterClient)(nil).GetK8sVersion))
}

// SetConfig mocks base method.
func (m *MockClusterClient) SetConfig(args MinikubeClientConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConfig", args)
}

// SetConfig indicates an expected call of SetConfig.
func (mr *MockClusterClientMockRecorder) SetConfig(args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockClusterClient)(nil).SetConfig), args)
}

// SetDependencies mocks base method.
func (m *MockClusterClient) SetDependencies(dep MinikubeClientDeps) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDependencies", dep)
}

// SetDependencies indicates an expected call of SetDependencies.
func (mr *MockClusterClientMockRecorder) SetDependencies(dep interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDependencies", reflect.TypeOf((*MockClusterClient)(nil).SetDependencies), dep)
}

// Start mocks base method.
func (m *MockClusterClient) Start() (*kubeconfig.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(*kubeconfig.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Start indicates an expected call of Start.
func (mr *MockClusterClientMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockClusterClient)(nil).Start))
}
