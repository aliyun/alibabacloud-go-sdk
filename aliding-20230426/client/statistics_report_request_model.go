// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatisticsReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *StatisticsReportRequest
	GetReportId() *string
	SetTenantContext(v *StatisticsReportRequestTenantContext) *StatisticsReportRequest
	GetTenantContext() *StatisticsReportRequestTenantContext
}

type StatisticsReportRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 173xxxx
	ReportId      *string                               `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	TenantContext *StatisticsReportRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s StatisticsReportRequest) String() string {
	return dara.Prettify(s)
}

func (s StatisticsReportRequest) GoString() string {
	return s.String()
}

func (s *StatisticsReportRequest) GetReportId() *string {
	return s.ReportId
}

func (s *StatisticsReportRequest) GetTenantContext() *StatisticsReportRequestTenantContext {
	return s.TenantContext
}

func (s *StatisticsReportRequest) SetReportId(v string) *StatisticsReportRequest {
	s.ReportId = &v
	return s
}

func (s *StatisticsReportRequest) SetTenantContext(v *StatisticsReportRequestTenantContext) *StatisticsReportRequest {
	s.TenantContext = v
	return s
}

func (s *StatisticsReportRequest) Validate() error {
	return dara.Validate(s)
}

type StatisticsReportRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s StatisticsReportRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s StatisticsReportRequestTenantContext) GoString() string {
	return s.String()
}

func (s *StatisticsReportRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *StatisticsReportRequestTenantContext) SetTenantId(v string) *StatisticsReportRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *StatisticsReportRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
