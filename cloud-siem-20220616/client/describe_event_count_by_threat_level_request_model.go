// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventCountByThreatLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeEventCountByThreatLevelRequest
	GetEndTime() *int64
	SetRegionId(v string) *DescribeEventCountByThreatLevelRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeEventCountByThreatLevelRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeEventCountByThreatLevelRequest
	GetRoleType() *int32
	SetStartTime(v int64) *DescribeEventCountByThreatLevelRequest
	GetStartTime() *int64
}

type DescribeEventCountByThreatLevelRequest struct {
	// End time of the query, in milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Region where the Data Management Center for threat analysis is located. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Assets are in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: Assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource directory member account ID.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// View type.
	//
	// - 0: View for the current Alibaba Cloud account.
	//
	// - 1: View for all accounts in your enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// Start time of the query, in milliseconds.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeEventCountByThreatLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventCountByThreatLevelRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventCountByThreatLevelRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeEventCountByThreatLevelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEventCountByThreatLevelRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeEventCountByThreatLevelRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeEventCountByThreatLevelRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeEventCountByThreatLevelRequest) SetEndTime(v int64) *DescribeEventCountByThreatLevelRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventCountByThreatLevelRequest) SetRegionId(v string) *DescribeEventCountByThreatLevelRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEventCountByThreatLevelRequest) SetRoleFor(v int64) *DescribeEventCountByThreatLevelRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeEventCountByThreatLevelRequest) SetRoleType(v int32) *DescribeEventCountByThreatLevelRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeEventCountByThreatLevelRequest) SetStartTime(v int64) *DescribeEventCountByThreatLevelRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEventCountByThreatLevelRequest) Validate() error {
	return dara.Validate(s)
}
