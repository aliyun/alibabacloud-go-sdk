// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatisticsListByTypeReportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOffset(v int64) *StatisticsListByTypeReportShrinkRequest
	GetOffset() *int64
	SetReportId(v string) *StatisticsListByTypeReportShrinkRequest
	GetReportId() *string
	SetSize(v int64) *StatisticsListByTypeReportShrinkRequest
	GetSize() *int64
	SetTenantContextShrink(v string) *StatisticsListByTypeReportShrinkRequest
	GetTenantContextShrink() *string
	SetType(v int64) *StatisticsListByTypeReportShrinkRequest
	GetType() *int64
}

type StatisticsListByTypeReportShrinkRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 173xxxx
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s StatisticsListByTypeReportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StatisticsListByTypeReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *StatisticsListByTypeReportShrinkRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *StatisticsListByTypeReportShrinkRequest) GetReportId() *string {
	return s.ReportId
}

func (s *StatisticsListByTypeReportShrinkRequest) GetSize() *int64 {
	return s.Size
}

func (s *StatisticsListByTypeReportShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *StatisticsListByTypeReportShrinkRequest) GetType() *int64 {
	return s.Type
}

func (s *StatisticsListByTypeReportShrinkRequest) SetOffset(v int64) *StatisticsListByTypeReportShrinkRequest {
	s.Offset = &v
	return s
}

func (s *StatisticsListByTypeReportShrinkRequest) SetReportId(v string) *StatisticsListByTypeReportShrinkRequest {
	s.ReportId = &v
	return s
}

func (s *StatisticsListByTypeReportShrinkRequest) SetSize(v int64) *StatisticsListByTypeReportShrinkRequest {
	s.Size = &v
	return s
}

func (s *StatisticsListByTypeReportShrinkRequest) SetTenantContextShrink(v string) *StatisticsListByTypeReportShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *StatisticsListByTypeReportShrinkRequest) SetType(v int64) *StatisticsListByTypeReportShrinkRequest {
	s.Type = &v
	return s
}

func (s *StatisticsListByTypeReportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
