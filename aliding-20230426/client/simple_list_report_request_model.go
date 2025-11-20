// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleListReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCursor(v int64) *SimpleListReportRequest
	GetCursor() *int64
	SetEndTime(v int64) *SimpleListReportRequest
	GetEndTime() *int64
	SetSize(v int64) *SimpleListReportRequest
	GetSize() *int64
	SetStartTime(v int64) *SimpleListReportRequest
	GetStartTime() *int64
	SetTemplateName(v string) *SimpleListReportRequest
	GetTemplateName() *string
	SetTenantContext(v *SimpleListReportRequestTenantContext) *SimpleListReportRequest
	GetTenantContext() *SimpleListReportRequestTenantContext
}

type SimpleListReportRequest struct {
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
	TemplateName  *string                               `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TenantContext *SimpleListReportRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s SimpleListReportRequest) String() string {
	return dara.Prettify(s)
}

func (s SimpleListReportRequest) GoString() string {
	return s.String()
}

func (s *SimpleListReportRequest) GetCursor() *int64 {
	return s.Cursor
}

func (s *SimpleListReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SimpleListReportRequest) GetSize() *int64 {
	return s.Size
}

func (s *SimpleListReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SimpleListReportRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *SimpleListReportRequest) GetTenantContext() *SimpleListReportRequestTenantContext {
	return s.TenantContext
}

func (s *SimpleListReportRequest) SetCursor(v int64) *SimpleListReportRequest {
	s.Cursor = &v
	return s
}

func (s *SimpleListReportRequest) SetEndTime(v int64) *SimpleListReportRequest {
	s.EndTime = &v
	return s
}

func (s *SimpleListReportRequest) SetSize(v int64) *SimpleListReportRequest {
	s.Size = &v
	return s
}

func (s *SimpleListReportRequest) SetStartTime(v int64) *SimpleListReportRequest {
	s.StartTime = &v
	return s
}

func (s *SimpleListReportRequest) SetTemplateName(v string) *SimpleListReportRequest {
	s.TemplateName = &v
	return s
}

func (s *SimpleListReportRequest) SetTenantContext(v *SimpleListReportRequestTenantContext) *SimpleListReportRequest {
	s.TenantContext = v
	return s
}

func (s *SimpleListReportRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SimpleListReportRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SimpleListReportRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s SimpleListReportRequestTenantContext) GoString() string {
	return s.String()
}

func (s *SimpleListReportRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *SimpleListReportRequestTenantContext) SetTenantId(v string) *SimpleListReportRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *SimpleListReportRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
