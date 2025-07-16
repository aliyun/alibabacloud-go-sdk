// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatisticsReportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *StatisticsReportShrinkRequest
	GetReportId() *string
	SetTenantContextShrink(v string) *StatisticsReportShrinkRequest
	GetTenantContextShrink() *string
}

type StatisticsReportShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 173xxxx
	ReportId            *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s StatisticsReportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StatisticsReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *StatisticsReportShrinkRequest) GetReportId() *string {
	return s.ReportId
}

func (s *StatisticsReportShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *StatisticsReportShrinkRequest) SetReportId(v string) *StatisticsReportShrinkRequest {
	s.ReportId = &v
	return s
}

func (s *StatisticsReportShrinkRequest) SetTenantContextShrink(v string) *StatisticsReportShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *StatisticsReportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
