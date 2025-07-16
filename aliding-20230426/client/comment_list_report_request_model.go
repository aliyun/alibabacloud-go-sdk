// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommentListReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOffset(v int64) *CommentListReportRequest
	GetOffset() *int64
	SetReportId(v string) *CommentListReportRequest
	GetReportId() *string
	SetSize(v int64) *CommentListReportRequest
	GetSize() *int64
	SetTenantContext(v *CommentListReportRequestTenantContext) *CommentListReportRequest
	GetTenantContext() *CommentListReportRequestTenantContext
}

type CommentListReportRequest struct {
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
	Size          *int64                                 `json:"Size,omitempty" xml:"Size,omitempty"`
	TenantContext *CommentListReportRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s CommentListReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CommentListReportRequest) GoString() string {
	return s.String()
}

func (s *CommentListReportRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *CommentListReportRequest) GetReportId() *string {
	return s.ReportId
}

func (s *CommentListReportRequest) GetSize() *int64 {
	return s.Size
}

func (s *CommentListReportRequest) GetTenantContext() *CommentListReportRequestTenantContext {
	return s.TenantContext
}

func (s *CommentListReportRequest) SetOffset(v int64) *CommentListReportRequest {
	s.Offset = &v
	return s
}

func (s *CommentListReportRequest) SetReportId(v string) *CommentListReportRequest {
	s.ReportId = &v
	return s
}

func (s *CommentListReportRequest) SetSize(v int64) *CommentListReportRequest {
	s.Size = &v
	return s
}

func (s *CommentListReportRequest) SetTenantContext(v *CommentListReportRequestTenantContext) *CommentListReportRequest {
	s.TenantContext = v
	return s
}

func (s *CommentListReportRequest) Validate() error {
	return dara.Validate(s)
}

type CommentListReportRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CommentListReportRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CommentListReportRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CommentListReportRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CommentListReportRequestTenantContext) SetTenantId(v string) *CommentListReportRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CommentListReportRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
