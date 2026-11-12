// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostFinishCustomizeRuleTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *PostFinishCustomizeRuleTestRequest
	GetId() *int64
	SetRegionId(v string) *PostFinishCustomizeRuleTestRequest
	GetRegionId() *string
	SetRoleFor(v int64) *PostFinishCustomizeRuleTestRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *PostFinishCustomizeRuleTestRequest
	GetRoleType() *int32
}

type PostFinishCustomizeRuleTestRequest struct {
	// The ID of the custom rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region of the data management center for Threat Analysis. You must select the region of the management center based on the region where your assets are deployed. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Assets are located in regions outside China.
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
}

func (s PostFinishCustomizeRuleTestRequest) String() string {
	return dara.Prettify(s)
}

func (s PostFinishCustomizeRuleTestRequest) GoString() string {
	return s.String()
}

func (s *PostFinishCustomizeRuleTestRequest) GetId() *int64 {
	return s.Id
}

func (s *PostFinishCustomizeRuleTestRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PostFinishCustomizeRuleTestRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *PostFinishCustomizeRuleTestRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *PostFinishCustomizeRuleTestRequest) SetId(v int64) *PostFinishCustomizeRuleTestRequest {
	s.Id = &v
	return s
}

func (s *PostFinishCustomizeRuleTestRequest) SetRegionId(v string) *PostFinishCustomizeRuleTestRequest {
	s.RegionId = &v
	return s
}

func (s *PostFinishCustomizeRuleTestRequest) SetRoleFor(v int64) *PostFinishCustomizeRuleTestRequest {
	s.RoleFor = &v
	return s
}

func (s *PostFinishCustomizeRuleTestRequest) SetRoleType(v int32) *PostFinishCustomizeRuleTestRequest {
	s.RoleType = &v
	return s
}

func (s *PostFinishCustomizeRuleTestRequest) Validate() error {
	return dara.Validate(s)
}
