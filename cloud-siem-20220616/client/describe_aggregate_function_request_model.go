// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAggregateFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeAggregateFunctionRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeAggregateFunctionRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeAggregateFunctionRequest
	GetRoleType() *int32
}

type DescribeAggregateFunctionRequest struct {
	// The region where the data management center of Threat Analysis is located. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: Your assets are in a region outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member. An administrator can switch to the perspective of the member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAggregateFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAggregateFunctionRequest) GoString() string {
	return s.String()
}

func (s *DescribeAggregateFunctionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAggregateFunctionRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeAggregateFunctionRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeAggregateFunctionRequest) SetRegionId(v string) *DescribeAggregateFunctionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAggregateFunctionRequest) SetRoleFor(v int64) *DescribeAggregateFunctionRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAggregateFunctionRequest) SetRoleType(v int32) *DescribeAggregateFunctionRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAggregateFunctionRequest) Validate() error {
	return dara.Validate(s)
}
