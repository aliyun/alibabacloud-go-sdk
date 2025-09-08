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
