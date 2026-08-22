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
	// The region of the data management center for threat analysis. Specify the management center region based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: assets in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: assets outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that the administrator switches to for viewing as another member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The rule type. Valid values:
	//
	// - predefine: predefined.
	//
	// - customize: custom.
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
