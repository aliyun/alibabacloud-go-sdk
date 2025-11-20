// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatisticsListByTypeReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOffset(v int64) *StatisticsListByTypeReportRequest
	GetOffset() *int64
	SetReportId(v string) *StatisticsListByTypeReportRequest
	GetReportId() *string
	SetSize(v int64) *StatisticsListByTypeReportRequest
	GetSize() *int64
	SetTenantContext(v *StatisticsListByTypeReportRequestTenantContext) *StatisticsListByTypeReportRequest
	GetTenantContext() *StatisticsListByTypeReportRequestTenantContext
	SetType(v int64) *StatisticsListByTypeReportRequest
	GetType() *int64
}

type StatisticsListByTypeReportRequest struct {
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
	Size          *int64                                          `json:"Size,omitempty" xml:"Size,omitempty"`
	TenantContext *StatisticsListByTypeReportRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 173xxxx
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s StatisticsListByTypeReportRequest) String() string {
	return dara.Prettify(s)
}

func (s StatisticsListByTypeReportRequest) GoString() string {
	return s.String()
}

func (s *StatisticsListByTypeReportRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *StatisticsListByTypeReportRequest) GetReportId() *string {
	return s.ReportId
}

func (s *StatisticsListByTypeReportRequest) GetSize() *int64 {
	return s.Size
}

func (s *StatisticsListByTypeReportRequest) GetTenantContext() *StatisticsListByTypeReportRequestTenantContext {
	return s.TenantContext
}

func (s *StatisticsListByTypeReportRequest) GetType() *int64 {
	return s.Type
}

func (s *StatisticsListByTypeReportRequest) SetOffset(v int64) *StatisticsListByTypeReportRequest {
	s.Offset = &v
	return s
}

func (s *StatisticsListByTypeReportRequest) SetReportId(v string) *StatisticsListByTypeReportRequest {
	s.ReportId = &v
	return s
}

func (s *StatisticsListByTypeReportRequest) SetSize(v int64) *StatisticsListByTypeReportRequest {
	s.Size = &v
	return s
}

func (s *StatisticsListByTypeReportRequest) SetTenantContext(v *StatisticsListByTypeReportRequestTenantContext) *StatisticsListByTypeReportRequest {
	s.TenantContext = v
	return s
}

func (s *StatisticsListByTypeReportRequest) SetType(v int64) *StatisticsListByTypeReportRequest {
	s.Type = &v
	return s
}

func (s *StatisticsListByTypeReportRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StatisticsListByTypeReportRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s StatisticsListByTypeReportRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s StatisticsListByTypeReportRequestTenantContext) GoString() string {
	return s.String()
}

func (s *StatisticsListByTypeReportRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *StatisticsListByTypeReportRequestTenantContext) SetTenantId(v string) *StatisticsListByTypeReportRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *StatisticsListByTypeReportRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
