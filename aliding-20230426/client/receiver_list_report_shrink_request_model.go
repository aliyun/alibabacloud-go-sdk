// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiverListReportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOffset(v int64) *ReceiverListReportShrinkRequest
	GetOffset() *int64
	SetReportId(v string) *ReceiverListReportShrinkRequest
	GetReportId() *string
	SetSize(v int64) *ReceiverListReportShrinkRequest
	GetSize() *int64
	SetTenantContextShrink(v string) *ReceiverListReportShrinkRequest
	GetTenantContextShrink() *string
}

type ReceiverListReportShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 173xxxx
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Size                *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s ReceiverListReportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReceiverListReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReceiverListReportShrinkRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *ReceiverListReportShrinkRequest) GetReportId() *string {
	return s.ReportId
}

func (s *ReceiverListReportShrinkRequest) GetSize() *int64 {
	return s.Size
}

func (s *ReceiverListReportShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *ReceiverListReportShrinkRequest) SetOffset(v int64) *ReceiverListReportShrinkRequest {
	s.Offset = &v
	return s
}

func (s *ReceiverListReportShrinkRequest) SetReportId(v string) *ReceiverListReportShrinkRequest {
	s.ReportId = &v
	return s
}

func (s *ReceiverListReportShrinkRequest) SetSize(v int64) *ReceiverListReportShrinkRequest {
	s.Size = &v
	return s
}

func (s *ReceiverListReportShrinkRequest) SetTenantContextShrink(v string) *ReceiverListReportShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *ReceiverListReportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
