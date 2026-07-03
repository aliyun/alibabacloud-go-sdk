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
	// The region of the management center for threat analysis. Select a region based on where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Assets are in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. An administrator can use this ID to switch to the member\\"s perspective.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of view. Valid values:
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts in the enterprise.
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
