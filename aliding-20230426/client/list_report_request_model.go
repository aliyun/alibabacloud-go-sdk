// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCursor(v int64) *ListReportRequest
	GetCursor() *int64
	SetEndTime(v int64) *ListReportRequest
	GetEndTime() *int64
	SetModifiedEndTime(v int64) *ListReportRequest
	GetModifiedEndTime() *int64
	SetModifiedStartTime(v int64) *ListReportRequest
	GetModifiedStartTime() *int64
	SetSize(v int64) *ListReportRequest
	GetSize() *int64
	SetStartTime(v int64) *ListReportRequest
	GetStartTime() *int64
	SetTemplateName(v string) *ListReportRequest
	GetTemplateName() *string
	SetTenantContext(v *ListReportRequestTenantContext) *ListReportRequest
	GetTenantContext() *ListReportRequestTenantContext
}

type ListReportRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	Cursor *int64 `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1507564800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1507564800000
	ModifiedEndTime *int64 `json:"ModifiedEndTime,omitempty" xml:"ModifiedEndTime,omitempty"`
	// example:
	//
	// 1507564800000
	ModifiedStartTime *int64 `json:"ModifiedStartTime,omitempty" xml:"ModifiedStartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1507564800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 我管理的模版
	TemplateName  *string                         `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TenantContext *ListReportRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s ListReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListReportRequest) GoString() string {
	return s.String()
}

func (s *ListReportRequest) GetCursor() *int64 {
	return s.Cursor
}

func (s *ListReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListReportRequest) GetModifiedEndTime() *int64 {
	return s.ModifiedEndTime
}

func (s *ListReportRequest) GetModifiedStartTime() *int64 {
	return s.ModifiedStartTime
}

func (s *ListReportRequest) GetSize() *int64 {
	return s.Size
}

func (s *ListReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListReportRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListReportRequest) GetTenantContext() *ListReportRequestTenantContext {
	return s.TenantContext
}

func (s *ListReportRequest) SetCursor(v int64) *ListReportRequest {
	s.Cursor = &v
	return s
}

func (s *ListReportRequest) SetEndTime(v int64) *ListReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListReportRequest) SetModifiedEndTime(v int64) *ListReportRequest {
	s.ModifiedEndTime = &v
	return s
}

func (s *ListReportRequest) SetModifiedStartTime(v int64) *ListReportRequest {
	s.ModifiedStartTime = &v
	return s
}

func (s *ListReportRequest) SetSize(v int64) *ListReportRequest {
	s.Size = &v
	return s
}

func (s *ListReportRequest) SetStartTime(v int64) *ListReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListReportRequest) SetTemplateName(v string) *ListReportRequest {
	s.TemplateName = &v
	return s
}

func (s *ListReportRequest) SetTenantContext(v *ListReportRequestTenantContext) *ListReportRequest {
	s.TenantContext = v
	return s
}

func (s *ListReportRequest) Validate() error {
	return dara.Validate(s)
}

type ListReportRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListReportRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s ListReportRequestTenantContext) GoString() string {
	return s.String()
}

func (s *ListReportRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *ListReportRequestTenantContext) SetTenantId(v string) *ListReportRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *ListReportRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
