// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisposeStrategyPlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeDisposeStrategyPlaybookRequest
	GetEndTime() *int64
	SetRegionId(v string) *DescribeDisposeStrategyPlaybookRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeDisposeStrategyPlaybookRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeDisposeStrategyPlaybookRequest
	GetRoleType() *int32
	SetStartTime(v int64) *DescribeDisposeStrategyPlaybookRequest
	GetStartTime() *int64
}

type DescribeDisposeStrategyPlaybookRequest struct {
	// The end of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the region in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions inside China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDisposeStrategyPlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisposeStrategyPlaybookRequest) GoString() string {
	return s.String()
}

func (s *DescribeDisposeStrategyPlaybookRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDisposeStrategyPlaybookRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDisposeStrategyPlaybookRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeDisposeStrategyPlaybookRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeDisposeStrategyPlaybookRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDisposeStrategyPlaybookRequest) SetEndTime(v int64) *DescribeDisposeStrategyPlaybookRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookRequest) SetRegionId(v string) *DescribeDisposeStrategyPlaybookRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookRequest) SetRoleFor(v int64) *DescribeDisposeStrategyPlaybookRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookRequest) SetRoleType(v int32) *DescribeDisposeStrategyPlaybookRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookRequest) SetStartTime(v int64) *DescribeDisposeStrategyPlaybookRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookRequest) Validate() error {
	return dara.Validate(s)
}
