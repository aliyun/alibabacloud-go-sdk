// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeAlertTypeRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeAlertTypeRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeAlertTypeRequest
	GetRoleType() *int32
	SetRuleType(v string) *DescribeAlertTypeRequest
	GetRuleType() *string
}

type DescribeAlertTypeRequest struct {
	// The region where the Data Management center is located. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. An administrator can use this parameter to switch to the perspective of the member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in your enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The type of the rule. Valid values:
	//
	// - predefine: predefined
	//
	// - customize: custom
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeAlertTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlertTypeRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeAlertTypeRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeAlertTypeRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeAlertTypeRequest) SetRegionId(v string) *DescribeAlertTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertTypeRequest) SetRoleFor(v int64) *DescribeAlertTypeRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertTypeRequest) SetRoleType(v int32) *DescribeAlertTypeRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAlertTypeRequest) SetRuleType(v string) *DescribeAlertTypeRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeAlertTypeRequest) Validate() error {
	return dara.Validate(s)
}
