// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommentListReportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOffset(v int64) *CommentListReportShrinkRequest
	GetOffset() *int64
	SetReportId(v string) *CommentListReportShrinkRequest
	GetReportId() *string
	SetSize(v int64) *CommentListReportShrinkRequest
	GetSize() *int64
	SetTenantContextShrink(v string) *CommentListReportShrinkRequest
	GetTenantContextShrink() *string
}

type CommentListReportShrinkRequest struct {
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

func (s CommentListReportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CommentListReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *CommentListReportShrinkRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *CommentListReportShrinkRequest) GetReportId() *string {
	return s.ReportId
}

func (s *CommentListReportShrinkRequest) GetSize() *int64 {
	return s.Size
}

func (s *CommentListReportShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CommentListReportShrinkRequest) SetOffset(v int64) *CommentListReportShrinkRequest {
	s.Offset = &v
	return s
}

func (s *CommentListReportShrinkRequest) SetReportId(v string) *CommentListReportShrinkRequest {
	s.ReportId = &v
	return s
}

func (s *CommentListReportShrinkRequest) SetSize(v int64) *CommentListReportShrinkRequest {
	s.Size = &v
	return s
}

func (s *CommentListReportShrinkRequest) SetTenantContextShrink(v string) *CommentListReportShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CommentListReportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
