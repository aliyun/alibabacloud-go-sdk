// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizeRuleTestHistogramRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DescribeCustomizeRuleTestHistogramRequest
	GetId() *int64
	SetRegionId(v string) *DescribeCustomizeRuleTestHistogramRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeCustomizeRuleTestHistogramRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeCustomizeRuleTestHistogramRequest
	GetRoleType() *int32
}

type DescribeCustomizeRuleTestHistogramRequest struct {
	// The ID of the custom rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region where the Management Hub of threat analysis is located. Select the region of the Management Hub based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. This parameter is used when an administrator switches to the perspective of a member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in your enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeCustomizeRuleTestHistogramRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeRuleTestHistogramRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestHistogramRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeCustomizeRuleTestHistogramRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomizeRuleTestHistogramRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeCustomizeRuleTestHistogramRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeCustomizeRuleTestHistogramRequest) SetId(v int64) *DescribeCustomizeRuleTestHistogramRequest {
	s.Id = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramRequest) SetRegionId(v string) *DescribeCustomizeRuleTestHistogramRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramRequest) SetRoleFor(v int64) *DescribeCustomizeRuleTestHistogramRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramRequest) SetRoleType(v int32) *DescribeCustomizeRuleTestHistogramRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramRequest) Validate() error {
	return dara.Validate(s)
}
