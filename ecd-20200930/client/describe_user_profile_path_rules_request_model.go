// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserProfilePathRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopGroupId(v string) *DescribeUserProfilePathRulesRequest
	GetDesktopGroupId() *string
	SetRegionId(v string) *DescribeUserProfilePathRulesRequest
	GetRegionId() *string
	SetRuleType(v string) *DescribeUserProfilePathRulesRequest
	GetRuleType() *string
}

type DescribeUserProfilePathRulesRequest struct {
	// The desktop group ID. This parameter is required when you set RuleType parameter to DesktopGroup.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The rule type that you want to configure for the directory.
	//
	// Valid values:
	//
	// 	- DesktopGroup
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Default
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// DesktopGroup
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeUserProfilePathRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserProfilePathRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserProfilePathRulesRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeUserProfilePathRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserProfilePathRulesRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeUserProfilePathRulesRequest) SetDesktopGroupId(v string) *DescribeUserProfilePathRulesRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeUserProfilePathRulesRequest) SetRegionId(v string) *DescribeUserProfilePathRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserProfilePathRulesRequest) SetRuleType(v string) *DescribeUserProfilePathRulesRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeUserProfilePathRulesRequest) Validate() error {
	return dara.Validate(s)
}
