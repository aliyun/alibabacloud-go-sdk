// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHostAvailabilityListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *DescribeHostAvailabilityListRequest
	GetGroupId() *int64
	SetId(v int64) *DescribeHostAvailabilityListRequest
	GetId() *int64
	SetIds(v string) *DescribeHostAvailabilityListRequest
	GetIds() *string
	SetPageNumber(v int32) *DescribeHostAvailabilityListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHostAvailabilityListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeHostAvailabilityListRequest
	GetRegionId() *string
	SetTaskName(v string) *DescribeHostAvailabilityListRequest
	GetTaskName() *string
}

type DescribeHostAvailabilityListRequest struct {
	// The ID of the application group.
	//
	// example:
	//
	// 12345
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the availability monitoring task.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IDs of the availability monitoring tasks. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// 123456,345678
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Minimum value: 1. Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the availability monitoring task.
	//
	// example:
	//
	// ecs_instance
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeHostAvailabilityListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostAvailabilityListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeHostAvailabilityListRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeHostAvailabilityListRequest) GetIds() *string {
	return s.Ids
}

func (s *DescribeHostAvailabilityListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHostAvailabilityListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHostAvailabilityListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHostAvailabilityListRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeHostAvailabilityListRequest) SetGroupId(v int64) *DescribeHostAvailabilityListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeHostAvailabilityListRequest) SetId(v int64) *DescribeHostAvailabilityListRequest {
	s.Id = &v
	return s
}

func (s *DescribeHostAvailabilityListRequest) SetIds(v string) *DescribeHostAvailabilityListRequest {
	s.Ids = &v
	return s
}

func (s *DescribeHostAvailabilityListRequest) SetPageNumber(v int32) *DescribeHostAvailabilityListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHostAvailabilityListRequest) SetPageSize(v int32) *DescribeHostAvailabilityListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHostAvailabilityListRequest) SetRegionId(v string) *DescribeHostAvailabilityListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHostAvailabilityListRequest) SetTaskName(v string) *DescribeHostAvailabilityListRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeHostAvailabilityListRequest) Validate() error {
	return dara.Validate(s)
}
