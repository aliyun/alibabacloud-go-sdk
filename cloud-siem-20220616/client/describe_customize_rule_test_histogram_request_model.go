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
	// The ID of the rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
