// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertsCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeAlertsCountRequest
	GetEndTime() *int64
	SetQueryType(v string) *DescribeAlertsCountRequest
	GetQueryType() *string
	SetRegionId(v string) *DescribeAlertsCountRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeAlertsCountRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeAlertsCountRequest
	GetRoleType() *int32
	SetStartTime(v int64) *DescribeAlertsCountRequest
	GetStartTime() *int64
}

type DescribeAlertsCountRequest struct {
	// The end of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// bySrcProd
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
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
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAlertsCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertsCountRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAlertsCountRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeAlertsCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlertsCountRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeAlertsCountRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeAlertsCountRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAlertsCountRequest) SetEndTime(v int64) *DescribeAlertsCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertsCountRequest) SetQueryType(v string) *DescribeAlertsCountRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeAlertsCountRequest) SetRegionId(v string) *DescribeAlertsCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertsCountRequest) SetRoleFor(v int64) *DescribeAlertsCountRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertsCountRequest) SetRoleType(v int32) *DescribeAlertsCountRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAlertsCountRequest) SetStartTime(v int64) *DescribeAlertsCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertsCountRequest) Validate() error {
	return dara.Validate(s)
}
