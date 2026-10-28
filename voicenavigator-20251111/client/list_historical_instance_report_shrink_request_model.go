// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalInstanceReportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListHistoricalInstanceReportShrinkRequest
	GetEndTime() *int64
	SetInstanceIdsShrink(v string) *ListHistoricalInstanceReportShrinkRequest
	GetInstanceIdsShrink() *string
	SetPageNumber(v int32) *ListHistoricalInstanceReportShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHistoricalInstanceReportShrinkRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ListHistoricalInstanceReportShrinkRequest
	GetStartTime() *int64
}

type ListHistoricalInstanceReportShrinkRequest struct {
	// The end time for the query. The server rounds this value up to the next hour.
	//
	// example:
	//
	// 1775131200000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// A JSON array of instance IDs. The array can contain up to 100 IDs.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The number of the page to return. The maximum value is 100.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The maximum value is 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time for the query. The server rounds this value up to the next hour.
	//
	// example:
	//
	// 1775044800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListHistoricalInstanceReportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalInstanceReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHistoricalInstanceReportShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListHistoricalInstanceReportShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *ListHistoricalInstanceReportShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHistoricalInstanceReportShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHistoricalInstanceReportShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListHistoricalInstanceReportShrinkRequest) SetEndTime(v int64) *ListHistoricalInstanceReportShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListHistoricalInstanceReportShrinkRequest) SetInstanceIdsShrink(v string) *ListHistoricalInstanceReportShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *ListHistoricalInstanceReportShrinkRequest) SetPageNumber(v int32) *ListHistoricalInstanceReportShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalInstanceReportShrinkRequest) SetPageSize(v int32) *ListHistoricalInstanceReportShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalInstanceReportShrinkRequest) SetStartTime(v int64) *ListHistoricalInstanceReportShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListHistoricalInstanceReportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
