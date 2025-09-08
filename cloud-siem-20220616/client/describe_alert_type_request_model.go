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
	// The type of rule. Valid values:
	//
	// - predefine: the defined rule by system
	//
	// - customize: the customed rule by user
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
