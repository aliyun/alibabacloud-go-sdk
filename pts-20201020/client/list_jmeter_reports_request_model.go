// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJMeterReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *ListJMeterReportsRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *ListJMeterReportsRequest
	GetEndTime() *int64
	SetKeyword(v string) *ListJMeterReportsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListJMeterReportsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJMeterReportsRequest
	GetPageSize() *int32
	SetReportId(v string) *ListJMeterReportsRequest
	GetReportId() *string
	SetSceneId(v string) *ListJMeterReportsRequest
	GetSceneId() *string
}

type ListJMeterReportsRequest struct {
	// The beginning of the time range to query. Unit: ms.
	//
	// example:
	//
	// 1637115303000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query.
	//
	// example:
	//
	// 1637115306000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The report keyword.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of the report page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of reports to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The report ID.
	//
	// example:
	//
	// 7R4RE352
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the scenario whose report you want to view.
	//
	// example:
	//
	// 10YPA8H
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ListJMeterReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJMeterReportsRequest) GoString() string {
	return s.String()
}

func (s *ListJMeterReportsRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *ListJMeterReportsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListJMeterReportsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListJMeterReportsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJMeterReportsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJMeterReportsRequest) GetReportId() *string {
	return s.ReportId
}

func (s *ListJMeterReportsRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *ListJMeterReportsRequest) SetBeginTime(v int64) *ListJMeterReportsRequest {
	s.BeginTime = &v
	return s
}

func (s *ListJMeterReportsRequest) SetEndTime(v int64) *ListJMeterReportsRequest {
	s.EndTime = &v
	return s
}

func (s *ListJMeterReportsRequest) SetKeyword(v string) *ListJMeterReportsRequest {
	s.Keyword = &v
	return s
}

func (s *ListJMeterReportsRequest) SetPageNumber(v int32) *ListJMeterReportsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJMeterReportsRequest) SetPageSize(v int32) *ListJMeterReportsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJMeterReportsRequest) SetReportId(v string) *ListJMeterReportsRequest {
	s.ReportId = &v
	return s
}

func (s *ListJMeterReportsRequest) SetSceneId(v string) *ListJMeterReportsRequest {
	s.SceneId = &v
	return s
}

func (s *ListJMeterReportsRequest) Validate() error {
	return dara.Validate(s)
}
