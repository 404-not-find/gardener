// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/client/garden/clientset/versioned/typed/garden/v1beta1 (interfaces: GardenV1beta1Interface,ShootInterface)

// Package v1beta1 is a generated GoMock package.
package v1beta1

import (
	v1beta1 "github.com/gardener/gardener/pkg/apis/garden/v1beta1"
	v1beta10 "github.com/gardener/gardener/pkg/client/garden/clientset/versioned/typed/garden/v1beta1"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	reflect "reflect"
)

// MockGardenV1beta1Interface is a mock of GardenV1beta1Interface interface
type MockGardenV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockGardenV1beta1InterfaceMockRecorder
}

// MockGardenV1beta1InterfaceMockRecorder is the mock recorder for MockGardenV1beta1Interface
type MockGardenV1beta1InterfaceMockRecorder struct {
	mock *MockGardenV1beta1Interface
}

// NewMockGardenV1beta1Interface creates a new mock instance
func NewMockGardenV1beta1Interface(ctrl *gomock.Controller) *MockGardenV1beta1Interface {
	mock := &MockGardenV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockGardenV1beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGardenV1beta1Interface) EXPECT() *MockGardenV1beta1InterfaceMockRecorder {
	return m.recorder
}

// BackupInfrastructures mocks base method
func (m *MockGardenV1beta1Interface) BackupInfrastructures(arg0 string) v1beta10.BackupInfrastructureInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackupInfrastructures", arg0)
	ret0, _ := ret[0].(v1beta10.BackupInfrastructureInterface)
	return ret0
}

// BackupInfrastructures indicates an expected call of BackupInfrastructures
func (mr *MockGardenV1beta1InterfaceMockRecorder) BackupInfrastructures(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackupInfrastructures", reflect.TypeOf((*MockGardenV1beta1Interface)(nil).BackupInfrastructures), arg0)
}

// CloudProfiles mocks base method
func (m *MockGardenV1beta1Interface) CloudProfiles() v1beta10.CloudProfileInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProfiles")
	ret0, _ := ret[0].(v1beta10.CloudProfileInterface)
	return ret0
}

// CloudProfiles indicates an expected call of CloudProfiles
func (mr *MockGardenV1beta1InterfaceMockRecorder) CloudProfiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProfiles", reflect.TypeOf((*MockGardenV1beta1Interface)(nil).CloudProfiles))
}

// Projects mocks base method
func (m *MockGardenV1beta1Interface) Projects() v1beta10.ProjectInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Projects")
	ret0, _ := ret[0].(v1beta10.ProjectInterface)
	return ret0
}

// Projects indicates an expected call of Projects
func (mr *MockGardenV1beta1InterfaceMockRecorder) Projects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Projects", reflect.TypeOf((*MockGardenV1beta1Interface)(nil).Projects))
}

// Quotas mocks base method
func (m *MockGardenV1beta1Interface) Quotas(arg0 string) v1beta10.QuotaInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quotas", arg0)
	ret0, _ := ret[0].(v1beta10.QuotaInterface)
	return ret0
}

// Quotas indicates an expected call of Quotas
func (mr *MockGardenV1beta1InterfaceMockRecorder) Quotas(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quotas", reflect.TypeOf((*MockGardenV1beta1Interface)(nil).Quotas), arg0)
}

// RESTClient mocks base method
func (m *MockGardenV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient
func (mr *MockGardenV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockGardenV1beta1Interface)(nil).RESTClient))
}

// SecretBindings mocks base method
func (m *MockGardenV1beta1Interface) SecretBindings(arg0 string) v1beta10.SecretBindingInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretBindings", arg0)
	ret0, _ := ret[0].(v1beta10.SecretBindingInterface)
	return ret0
}

// SecretBindings indicates an expected call of SecretBindings
func (mr *MockGardenV1beta1InterfaceMockRecorder) SecretBindings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretBindings", reflect.TypeOf((*MockGardenV1beta1Interface)(nil).SecretBindings), arg0)
}

// Seeds mocks base method
func (m *MockGardenV1beta1Interface) Seeds() v1beta10.SeedInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seeds")
	ret0, _ := ret[0].(v1beta10.SeedInterface)
	return ret0
}

// Seeds indicates an expected call of Seeds
func (mr *MockGardenV1beta1InterfaceMockRecorder) Seeds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seeds", reflect.TypeOf((*MockGardenV1beta1Interface)(nil).Seeds))
}

// Shoots mocks base method
func (m *MockGardenV1beta1Interface) Shoots(arg0 string) v1beta10.ShootInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shoots", arg0)
	ret0, _ := ret[0].(v1beta10.ShootInterface)
	return ret0
}

// Shoots indicates an expected call of Shoots
func (mr *MockGardenV1beta1InterfaceMockRecorder) Shoots(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shoots", reflect.TypeOf((*MockGardenV1beta1Interface)(nil).Shoots), arg0)
}

// MockShootInterface is a mock of ShootInterface interface
type MockShootInterface struct {
	ctrl     *gomock.Controller
	recorder *MockShootInterfaceMockRecorder
}

// MockShootInterfaceMockRecorder is the mock recorder for MockShootInterface
type MockShootInterfaceMockRecorder struct {
	mock *MockShootInterface
}

// NewMockShootInterface creates a new mock instance
func NewMockShootInterface(ctrl *gomock.Controller) *MockShootInterface {
	mock := &MockShootInterface{ctrl: ctrl}
	mock.recorder = &MockShootInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockShootInterface) EXPECT() *MockShootInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockShootInterface) Create(arg0 *v1beta1.Shoot) (*v1beta1.Shoot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1beta1.Shoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockShootInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockShootInterface)(nil).Create), arg0)
}

// Delete mocks base method
func (m *MockShootInterface) Delete(arg0 string, arg1 *v1.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockShootInterfaceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockShootInterface)(nil).Delete), arg0, arg1)
}

// DeleteCollection mocks base method
func (m *MockShootInterface) DeleteCollection(arg0 *v1.DeleteOptions, arg1 v1.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockShootInterfaceMockRecorder) DeleteCollection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockShootInterface)(nil).DeleteCollection), arg0, arg1)
}

// Get mocks base method
func (m *MockShootInterface) Get(arg0 string, arg1 v1.GetOptions) (*v1beta1.Shoot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.Shoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockShootInterfaceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockShootInterface)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockShootInterface) List(arg0 v1.ListOptions) (*v1beta1.ShootList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1beta1.ShootList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockShootInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockShootInterface)(nil).List), arg0)
}

// Patch mocks base method
func (m *MockShootInterface) Patch(arg0 string, arg1 types.PatchType, arg2 []byte, arg3 ...string) (*v1beta1.Shoot, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1beta1.Shoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockShootInterfaceMockRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockShootInterface)(nil).Patch), varargs...)
}

// Update mocks base method
func (m *MockShootInterface) Update(arg0 *v1beta1.Shoot) (*v1beta1.Shoot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1beta1.Shoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockShootInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockShootInterface)(nil).Update), arg0)
}

// UpdateStatus mocks base method
func (m *MockShootInterface) UpdateStatus(arg0 *v1beta1.Shoot) (*v1beta1.Shoot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0)
	ret0, _ := ret[0].(*v1beta1.Shoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockShootInterfaceMockRecorder) UpdateStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockShootInterface)(nil).UpdateStatus), arg0)
}

// Watch mocks base method
func (m *MockShootInterface) Watch(arg0 v1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockShootInterfaceMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockShootInterface)(nil).Watch), arg0)
}