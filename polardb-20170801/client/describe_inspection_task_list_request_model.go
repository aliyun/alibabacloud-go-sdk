// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeInspectionTaskListRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeInspectionTaskListRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeInspectionTaskListRequest
	GetRegionId() *string
	SetSpaceId(v string) *DescribeInspectionTaskListRequest
	GetSpaceId() *string
}

type DescribeInspectionTaskListRequest struct {
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the region where the task is located.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the inspection workspace, which is used to filter tasks.
	//
	// example:
	//
	// space-123
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeInspectionTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTaskListRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeInspectionTaskListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInspectionTaskListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInspectionTaskListRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *DescribeInspectionTaskListRequest) SetPageNumber(v int64) *DescribeInspectionTaskListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInspectionTaskListRequest) SetPageSize(v int64) *DescribeInspectionTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInspectionTaskListRequest) SetRegionId(v string) *DescribeInspectionTaskListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInspectionTaskListRequest) SetSpaceId(v string) *DescribeInspectionTaskListRequest {
	s.SpaceId = &v
	return s
}

func (s *DescribeInspectionTaskListRequest) Validate() error {
	return dara.Validate(s)
}
