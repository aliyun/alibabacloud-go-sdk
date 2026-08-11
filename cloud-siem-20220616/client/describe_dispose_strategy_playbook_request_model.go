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
	// The end time of the query, in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The region where the data management center of threat detection and response is located. Select the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: the asset belongs to the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: the asset belongs to a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user to which the administrator switches the view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The start time of the query, in milliseconds.
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
