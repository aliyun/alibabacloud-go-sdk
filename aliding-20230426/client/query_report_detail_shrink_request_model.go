// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReportDetailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *QueryReportDetailShrinkRequest
	GetReportId() *string
	SetTenantContextShrink(v string) *QueryReportDetailShrinkRequest
	GetTenantContextShrink() *string
}

type QueryReportDetailShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 18054XXX
	ReportId            *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s QueryReportDetailShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryReportDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryReportDetailShrinkRequest) GetReportId() *string {
	return s.ReportId
}

func (s *QueryReportDetailShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryReportDetailShrinkRequest) SetReportId(v string) *QueryReportDetailShrinkRequest {
	s.ReportId = &v
	return s
}

func (s *QueryReportDetailShrinkRequest) SetTenantContextShrink(v string) *QueryReportDetailShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryReportDetailShrinkRequest) Validate() error {
	return dara.Validate(s)
}
