// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionReportListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInspectionReportListRequest
	GetEndTime() *string
	SetPageNumber(v int64) *DescribeInspectionReportListRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeInspectionReportListRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeInspectionReportListRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeInspectionReportListRequest
	GetStartTime() *string
	SetTaskId(v string) *DescribeInspectionReportListRequest
	GetTaskId() *string
}

type DescribeInspectionReportListRequest struct {
	// The end time of the query. If you specify StartTime, you must also specify EndTime. Specify the time in the ISO 8601 standard in the YYYY-MM-DDTHH:mmZ format. The time must be in UTC and cannot be earlier than the start time.
	//
	// example:
	//
	// 2026-11-04T02:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of the page to return. Pages start from 1. Maximum value: 2147483647. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 30, 50, and 100. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the inspection task.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDTHH:mmZ format. The time must be in UTC and cannot be later than the current time.
	//
	// example:
	//
	// 2026-11-04T01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the inspection task. This parameter is required.
	//
	// example:
	//
	// task-001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeInspectionReportListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionReportListRequest) GoString() string {
	return s.String()
}

func (s *DescribeInspectionReportListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInspectionReportListRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeInspectionReportListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInspectionReportListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInspectionReportListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInspectionReportListRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeInspectionReportListRequest) SetEndTime(v string) *DescribeInspectionReportListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInspectionReportListRequest) SetPageNumber(v int64) *DescribeInspectionReportListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInspectionReportListRequest) SetPageSize(v int64) *DescribeInspectionReportListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInspectionReportListRequest) SetRegionId(v string) *DescribeInspectionReportListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInspectionReportListRequest) SetStartTime(v string) *DescribeInspectionReportListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInspectionReportListRequest) SetTaskId(v string) *DescribeInspectionReportListRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeInspectionReportListRequest) Validate() error {
	return dara.Validate(s)
}
