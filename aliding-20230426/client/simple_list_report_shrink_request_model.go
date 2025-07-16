// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleListReportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCursor(v int64) *SimpleListReportShrinkRequest
	GetCursor() *int64
	SetEndTime(v int64) *SimpleListReportShrinkRequest
	GetEndTime() *int64
	SetSize(v int64) *SimpleListReportShrinkRequest
	GetSize() *int64
	SetStartTime(v int64) *SimpleListReportShrinkRequest
	GetStartTime() *int64
	SetTemplateName(v string) *SimpleListReportShrinkRequest
	GetTemplateName() *string
	SetTenantContextShrink(v string) *SimpleListReportShrinkRequest
	GetTenantContextShrink() *string
}

type SimpleListReportShrinkRequest struct {
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
	TemplateName        *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s SimpleListReportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SimpleListReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *SimpleListReportShrinkRequest) GetCursor() *int64 {
	return s.Cursor
}

func (s *SimpleListReportShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SimpleListReportShrinkRequest) GetSize() *int64 {
	return s.Size
}

func (s *SimpleListReportShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SimpleListReportShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *SimpleListReportShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *SimpleListReportShrinkRequest) SetCursor(v int64) *SimpleListReportShrinkRequest {
	s.Cursor = &v
	return s
}

func (s *SimpleListReportShrinkRequest) SetEndTime(v int64) *SimpleListReportShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *SimpleListReportShrinkRequest) SetSize(v int64) *SimpleListReportShrinkRequest {
	s.Size = &v
	return s
}

func (s *SimpleListReportShrinkRequest) SetStartTime(v int64) *SimpleListReportShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *SimpleListReportShrinkRequest) SetTemplateName(v string) *SimpleListReportShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *SimpleListReportShrinkRequest) SetTenantContextShrink(v string) *SimpleListReportShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *SimpleListReportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
