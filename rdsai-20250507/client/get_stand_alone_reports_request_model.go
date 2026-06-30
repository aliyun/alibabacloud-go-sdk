// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandAloneReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetStandAloneReportsRequest
	GetEndTime() *string
	SetPageNumber(v int64) *GetStandAloneReportsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetStandAloneReportsRequest
	GetPageSize() *int64
	SetReportType(v string) *GetStandAloneReportsRequest
	GetReportType() *string
	SetStartTime(v string) *GetStandAloneReportsRequest
	GetStartTime() *string
}

type GetStandAloneReportsRequest struct {
	// The end of the query\\"s time range. The time must be in UTC and in the `YYYY-MM-DDTHH:mm:ssZ` format. If omitted, no end time filter is applied.
	//
	// example:
	//
	// 2026-01-19T02:20:20Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of the page to return. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default: 20. Maximum: 100.
	//
	// example:
	//
	// 10
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The start of the query\\"s time range. The time must be in UTC and in the `YYYY-MM-DDTHH:mm:ssZ` format. If omitted, no start time filter is applied.
	//
	// example:
	//
	// 2025-03-11T02:09:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetStandAloneReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStandAloneReportsRequest) GoString() string {
	return s.String()
}

func (s *GetStandAloneReportsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetStandAloneReportsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetStandAloneReportsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetStandAloneReportsRequest) GetReportType() *string {
	return s.ReportType
}

func (s *GetStandAloneReportsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetStandAloneReportsRequest) SetEndTime(v string) *GetStandAloneReportsRequest {
	s.EndTime = &v
	return s
}

func (s *GetStandAloneReportsRequest) SetPageNumber(v int64) *GetStandAloneReportsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetStandAloneReportsRequest) SetPageSize(v int64) *GetStandAloneReportsRequest {
	s.PageSize = &v
	return s
}

func (s *GetStandAloneReportsRequest) SetReportType(v string) *GetStandAloneReportsRequest {
	s.ReportType = &v
	return s
}

func (s *GetStandAloneReportsRequest) SetStartTime(v string) *GetStandAloneReportsRequest {
	s.StartTime = &v
	return s
}

func (s *GetStandAloneReportsRequest) Validate() error {
	return dara.Validate(s)
}
