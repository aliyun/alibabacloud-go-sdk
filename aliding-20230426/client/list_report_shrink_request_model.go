// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCursor(v int64) *ListReportShrinkRequest
	GetCursor() *int64
	SetEndTime(v int64) *ListReportShrinkRequest
	GetEndTime() *int64
	SetModifiedEndTime(v int64) *ListReportShrinkRequest
	GetModifiedEndTime() *int64
	SetModifiedStartTime(v int64) *ListReportShrinkRequest
	GetModifiedStartTime() *int64
	SetSize(v int64) *ListReportShrinkRequest
	GetSize() *int64
	SetStartTime(v int64) *ListReportShrinkRequest
	GetStartTime() *int64
	SetTemplateName(v string) *ListReportShrinkRequest
	GetTemplateName() *string
	SetTenantContextShrink(v string) *ListReportShrinkRequest
	GetTenantContextShrink() *string
}

type ListReportShrinkRequest struct {
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
	TemplateName        *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s ListReportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListReportShrinkRequest) GetCursor() *int64 {
	return s.Cursor
}

func (s *ListReportShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListReportShrinkRequest) GetModifiedEndTime() *int64 {
	return s.ModifiedEndTime
}

func (s *ListReportShrinkRequest) GetModifiedStartTime() *int64 {
	return s.ModifiedStartTime
}

func (s *ListReportShrinkRequest) GetSize() *int64 {
	return s.Size
}

func (s *ListReportShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListReportShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListReportShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *ListReportShrinkRequest) SetCursor(v int64) *ListReportShrinkRequest {
	s.Cursor = &v
	return s
}

func (s *ListReportShrinkRequest) SetEndTime(v int64) *ListReportShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListReportShrinkRequest) SetModifiedEndTime(v int64) *ListReportShrinkRequest {
	s.ModifiedEndTime = &v
	return s
}

func (s *ListReportShrinkRequest) SetModifiedStartTime(v int64) *ListReportShrinkRequest {
	s.ModifiedStartTime = &v
	return s
}

func (s *ListReportShrinkRequest) SetSize(v int64) *ListReportShrinkRequest {
	s.Size = &v
	return s
}

func (s *ListReportShrinkRequest) SetStartTime(v int64) *ListReportShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListReportShrinkRequest) SetTemplateName(v string) *ListReportShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *ListReportShrinkRequest) SetTenantContextShrink(v string) *ListReportShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *ListReportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
