//  Copyright 2019 Google Inc. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GoogleCloudPlatform/compute-image-tools/cli_tools/gce_vm_image_import/domain (interfaces: HttpClientInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockHttpClientInterface is a mock of HttpClientInterface interface
type MockHttpClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHttpClientInterfaceMockRecorder
}

// MockHttpClientInterfaceMockRecorder is the mock recorder for MockHttpClientInterface
type MockHttpClientInterfaceMockRecorder struct {
	mock *MockHttpClientInterface
}

// NewMockHttpClientInterface creates a new mock instance
func NewMockHttpClientInterface(ctrl *gomock.Controller) *MockHttpClientInterface {
	mock := &MockHttpClientInterface{ctrl: ctrl}
	mock.recorder = &MockHttpClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHttpClientInterface) EXPECT() *MockHttpClientInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockHttpClientInterface) Get(arg0 string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockHttpClientInterfaceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHttpClientInterface)(nil).Get), arg0)
}