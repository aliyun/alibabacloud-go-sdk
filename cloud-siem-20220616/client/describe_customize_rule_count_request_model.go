// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizeRuleCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeCustomizeRuleCountRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeCustomizeRuleCountRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeCustomizeRuleCountRequest
	GetRoleType() *int32
}

type DescribeCustomizeRuleCountRequest struct {
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
}

func (s DescribeCustomizeRuleCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeRuleCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomizeRuleCountRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeCustomizeRuleCountRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeCustomizeRuleCountRequest) SetRegionId(v string) *DescribeCustomizeRuleCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomizeRuleCountRequest) SetRoleFor(v int64) *DescribeCustomizeRuleCountRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeCustomizeRuleCountRequest) SetRoleType(v int32) *DescribeCustomizeRuleCountRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeCustomizeRuleCountRequest) Validate() error {
	return dara.Validate(s)
}
