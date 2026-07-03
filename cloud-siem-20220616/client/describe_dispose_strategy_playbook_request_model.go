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
	// The end time of the query. This value is a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The region of the Data Management center for Threat Analysis. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: The assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: The assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. An administrator can use this parameter to switch to the member\\"s view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The start time of the query. This value is a UNIX timestamp in milliseconds.
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
