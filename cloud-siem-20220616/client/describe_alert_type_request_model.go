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
	// The region where the data management center of Threat Analysis is located. Select the region of the management center based on the region where your assets are deployed. Valid values:
	//
	// - cn-hangzhou: Assets are deployed in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Assets are deployed in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose perspective the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The rule type. Valid values:
	//
	// - predefine: Predefined.
	//
	// - customize: Custom.
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
