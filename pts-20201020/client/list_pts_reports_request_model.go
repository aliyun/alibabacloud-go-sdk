// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPtsReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *ListPtsReportsRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *ListPtsReportsRequest
	GetEndTime() *int64
	SetKeyword(v string) *ListPtsReportsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListPtsReportsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPtsReportsRequest
	GetPageSize() *int32
	SetReportId(v string) *ListPtsReportsRequest
	GetReportId() *string
	SetSceneId(v string) *ListPtsReportsRequest
	GetSceneId() *string
}

type ListPtsReportsRequest struct {
	// The timestamp when the stress testing starts. Unit: ms.
	//
	// example:
	//
	// 1637115303000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The timestamp when the stress testing ends. Unit: ms.
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
	// The number of the page to return. The page number starts from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of reports to return per page.
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
	// 7RLPM3Y2
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the scenario whose report you want to view.
	//
	// example:
	//
	// 1PDAL8H
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ListPtsReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPtsReportsRequest) GoString() string {
	return s.String()
}

func (s *ListPtsReportsRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *ListPtsReportsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListPtsReportsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListPtsReportsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPtsReportsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPtsReportsRequest) GetReportId() *string {
	return s.ReportId
}

func (s *ListPtsReportsRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *ListPtsReportsRequest) SetBeginTime(v int64) *ListPtsReportsRequest {
	s.BeginTime = &v
	return s
}

func (s *ListPtsReportsRequest) SetEndTime(v int64) *ListPtsReportsRequest {
	s.EndTime = &v
	return s
}

func (s *ListPtsReportsRequest) SetKeyword(v string) *ListPtsReportsRequest {
	s.Keyword = &v
	return s
}

func (s *ListPtsReportsRequest) SetPageNumber(v int32) *ListPtsReportsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPtsReportsRequest) SetPageSize(v int32) *ListPtsReportsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPtsReportsRequest) SetReportId(v string) *ListPtsReportsRequest {
	s.ReportId = &v
	return s
}

func (s *ListPtsReportsRequest) SetSceneId(v string) *ListPtsReportsRequest {
	s.SceneId = &v
	return s
}

func (s *ListPtsReportsRequest) Validate() error {
	return dara.Validate(s)
}
