// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeEventsListRequest
	GetEndTime() *string
	SetInstanceIdList(v string) *DescribeEventsListRequest
	GetInstanceIdList() *string
	SetPageNumber(v int64) *DescribeEventsListRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeEventsListRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeEventsListRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeEventsListRequest
	GetStartTime() *string
}

type DescribeEventsListRequest struct {
	// example:
	//
	// 2025-08-28 18:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// r-uf6ce0r08lr7xnriq2
	InstanceIdList *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2025-01-01 18:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeEventsListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventsListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEventsListRequest) GetInstanceIdList() *string {
	return s.InstanceIdList
}

func (s *DescribeEventsListRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeEventsListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeEventsListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEventsListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEventsListRequest) SetEndTime(v string) *DescribeEventsListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventsListRequest) SetInstanceIdList(v string) *DescribeEventsListRequest {
	s.InstanceIdList = &v
	return s
}

func (s *DescribeEventsListRequest) SetPageNumber(v int64) *DescribeEventsListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventsListRequest) SetPageSize(v int64) *DescribeEventsListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsListRequest) SetRegionId(v string) *DescribeEventsListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEventsListRequest) SetStartTime(v string) *DescribeEventsListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEventsListRequest) Validate() error {
	return dara.Validate(s)
}
