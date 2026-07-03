// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizeRuleTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DescribeCustomizeRuleTestRequest
	GetId() *int64
	SetRegionId(v string) *DescribeCustomizeRuleTestRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeCustomizeRuleTestRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeCustomizeRuleTestRequest
	GetRoleType() *int32
}

type DescribeCustomizeRuleTestRequest struct {
	// The ID of the custom rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region of the Data Management center for Threat Analysis. Select a region based on where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in regions outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member account to which the administrator switches.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts within the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeCustomizeRuleTestRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeRuleTestRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeCustomizeRuleTestRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomizeRuleTestRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeCustomizeRuleTestRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeCustomizeRuleTestRequest) SetId(v int64) *DescribeCustomizeRuleTestRequest {
	s.Id = &v
	return s
}

func (s *DescribeCustomizeRuleTestRequest) SetRegionId(v string) *DescribeCustomizeRuleTestRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomizeRuleTestRequest) SetRoleFor(v int64) *DescribeCustomizeRuleTestRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeCustomizeRuleTestRequest) SetRoleType(v int32) *DescribeCustomizeRuleTestRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeCustomizeRuleTestRequest) Validate() error {
	return dara.Validate(s)
}
