// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/favonia/cloudflare-ddns/internal/api (interfaces: Handle)
//
// Generated by this command:
//
//	mockgen -typed -destination=../mocks/mock_api.go -package=mocks . Handle
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	netip "net/netip"
	reflect "reflect"

	api "github.com/favonia/cloudflare-ddns/internal/api"
	domain "github.com/favonia/cloudflare-ddns/internal/domain"
	ipnet "github.com/favonia/cloudflare-ddns/internal/ipnet"
	pp "github.com/favonia/cloudflare-ddns/internal/pp"
	gomock "go.uber.org/mock/gomock"
)

// MockHandle is a mock of Handle interface.
type MockHandle struct {
	ctrl     *gomock.Controller
	recorder *MockHandleMockRecorder
}

// MockHandleMockRecorder is the mock recorder for MockHandle.
type MockHandleMockRecorder struct {
	mock *MockHandle
}

// NewMockHandle creates a new mock instance.
func NewMockHandle(ctrl *gomock.Controller) *MockHandle {
	mock := &MockHandle{ctrl: ctrl}
	mock.recorder = &MockHandleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandle) EXPECT() *MockHandleMockRecorder {
	return m.recorder
}

// CreateRecord mocks base method.
func (m *MockHandle) CreateRecord(arg0 context.Context, arg1 pp.PP, arg2 domain.Domain, arg3 ipnet.Type, arg4 netip.Addr, arg5 api.TTL, arg6 bool) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecord", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CreateRecord indicates an expected call of CreateRecord.
func (mr *MockHandleMockRecorder) CreateRecord(arg0, arg1, arg2, arg3, arg4, arg5, arg6 any) *HandleCreateRecordCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecord", reflect.TypeOf((*MockHandle)(nil).CreateRecord), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	return &HandleCreateRecordCall{Call: call}
}

// HandleCreateRecordCall wrap *gomock.Call
type HandleCreateRecordCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandleCreateRecordCall) Return(arg0 string, arg1 bool) *HandleCreateRecordCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandleCreateRecordCall) Do(f func(context.Context, pp.PP, domain.Domain, ipnet.Type, netip.Addr, api.TTL, bool) (string, bool)) *HandleCreateRecordCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandleCreateRecordCall) DoAndReturn(f func(context.Context, pp.PP, domain.Domain, ipnet.Type, netip.Addr, api.TTL, bool) (string, bool)) *HandleCreateRecordCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteRecord mocks base method.
func (m *MockHandle) DeleteRecord(arg0 context.Context, arg1 pp.PP, arg2 domain.Domain, arg3 ipnet.Type, arg4 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecord", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DeleteRecord indicates an expected call of DeleteRecord.
func (mr *MockHandleMockRecorder) DeleteRecord(arg0, arg1, arg2, arg3, arg4 any) *HandleDeleteRecordCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecord", reflect.TypeOf((*MockHandle)(nil).DeleteRecord), arg0, arg1, arg2, arg3, arg4)
	return &HandleDeleteRecordCall{Call: call}
}

// HandleDeleteRecordCall wrap *gomock.Call
type HandleDeleteRecordCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandleDeleteRecordCall) Return(arg0 bool) *HandleDeleteRecordCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandleDeleteRecordCall) Do(f func(context.Context, pp.PP, domain.Domain, ipnet.Type, string) bool) *HandleDeleteRecordCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandleDeleteRecordCall) DoAndReturn(f func(context.Context, pp.PP, domain.Domain, ipnet.Type, string) bool) *HandleDeleteRecordCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FlushCache mocks base method.
func (m *MockHandle) FlushCache() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FlushCache")
}

// FlushCache indicates an expected call of FlushCache.
func (mr *MockHandleMockRecorder) FlushCache() *HandleFlushCacheCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushCache", reflect.TypeOf((*MockHandle)(nil).FlushCache))
	return &HandleFlushCacheCall{Call: call}
}

// HandleFlushCacheCall wrap *gomock.Call
type HandleFlushCacheCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandleFlushCacheCall) Return() *HandleFlushCacheCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandleFlushCacheCall) Do(f func()) *HandleFlushCacheCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandleFlushCacheCall) DoAndReturn(f func()) *HandleFlushCacheCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListRecords mocks base method.
func (m *MockHandle) ListRecords(arg0 context.Context, arg1 pp.PP, arg2 domain.Domain, arg3 ipnet.Type) (map[string]netip.Addr, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRecords", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(map[string]netip.Addr)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ListRecords indicates an expected call of ListRecords.
func (mr *MockHandleMockRecorder) ListRecords(arg0, arg1, arg2, arg3 any) *HandleListRecordsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRecords", reflect.TypeOf((*MockHandle)(nil).ListRecords), arg0, arg1, arg2, arg3)
	return &HandleListRecordsCall{Call: call}
}

// HandleListRecordsCall wrap *gomock.Call
type HandleListRecordsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandleListRecordsCall) Return(arg0 map[string]netip.Addr, arg1 bool) *HandleListRecordsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandleListRecordsCall) Do(f func(context.Context, pp.PP, domain.Domain, ipnet.Type) (map[string]netip.Addr, bool)) *HandleListRecordsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandleListRecordsCall) DoAndReturn(f func(context.Context, pp.PP, domain.Domain, ipnet.Type) (map[string]netip.Addr, bool)) *HandleListRecordsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateRecord mocks base method.
func (m *MockHandle) UpdateRecord(arg0 context.Context, arg1 pp.PP, arg2 domain.Domain, arg3 ipnet.Type, arg4 string, arg5 netip.Addr) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRecord", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateRecord indicates an expected call of UpdateRecord.
func (mr *MockHandleMockRecorder) UpdateRecord(arg0, arg1, arg2, arg3, arg4, arg5 any) *HandleUpdateRecordCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRecord", reflect.TypeOf((*MockHandle)(nil).UpdateRecord), arg0, arg1, arg2, arg3, arg4, arg5)
	return &HandleUpdateRecordCall{Call: call}
}

// HandleUpdateRecordCall wrap *gomock.Call
type HandleUpdateRecordCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandleUpdateRecordCall) Return(arg0 bool) *HandleUpdateRecordCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandleUpdateRecordCall) Do(f func(context.Context, pp.PP, domain.Domain, ipnet.Type, string, netip.Addr) bool) *HandleUpdateRecordCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandleUpdateRecordCall) DoAndReturn(f func(context.Context, pp.PP, domain.Domain, ipnet.Type, string, netip.Addr) bool) *HandleUpdateRecordCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
