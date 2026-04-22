// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalInstanceReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListHistoricalInstanceReportRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *ListHistoricalInstanceReportRequest
	GetInstanceIds() []*string
	SetPageNumber(v int32) *ListHistoricalInstanceReportRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHistoricalInstanceReportRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ListHistoricalInstanceReportRequest
	GetStartTime() *int64
}

type ListHistoricalInstanceReportRequest struct {
	// example:
	//
	// 1775131200000
	EndTime     *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1775044800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListHistoricalInstanceReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalInstanceReportRequest) GoString() string {
	return s.String()
}

func (s *ListHistoricalInstanceReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListHistoricalInstanceReportRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListHistoricalInstanceReportRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHistoricalInstanceReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHistoricalInstanceReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListHistoricalInstanceReportRequest) SetEndTime(v int64) *ListHistoricalInstanceReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListHistoricalInstanceReportRequest) SetInstanceIds(v []*string) *ListHistoricalInstanceReportRequest {
	s.InstanceIds = v
	return s
}

func (s *ListHistoricalInstanceReportRequest) SetPageNumber(v int32) *ListHistoricalInstanceReportRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalInstanceReportRequest) SetPageSize(v int32) *ListHistoricalInstanceReportRequest {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalInstanceReportRequest) SetStartTime(v int64) *ListHistoricalInstanceReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListHistoricalInstanceReportRequest) Validate() error {
	return dara.Validate(s)
}
