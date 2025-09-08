// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizeRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteCustomizeRuleRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DeleteCustomizeRuleRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DeleteCustomizeRuleRequest
	GetRoleType() *int32
	SetRuleId(v int64) *DeleteCustomizeRuleRequest
	GetRuleId() *int64
}

type DeleteCustomizeRuleRequest struct {
	// The region in which the service is deployed.
	//
	// example:
	//
	// cn-shanghai
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
	// The ID of the rule.
	//
	// example:
	//
	// 123456789
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteCustomizeRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizeRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomizeRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCustomizeRuleRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteCustomizeRuleRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DeleteCustomizeRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteCustomizeRuleRequest) SetRegionId(v string) *DeleteCustomizeRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomizeRuleRequest) SetRoleFor(v int64) *DeleteCustomizeRuleRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteCustomizeRuleRequest) SetRoleType(v int32) *DeleteCustomizeRuleRequest {
	s.RoleType = &v
	return s
}

func (s *DeleteCustomizeRuleRequest) SetRuleId(v int64) *DeleteCustomizeRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteCustomizeRuleRequest) Validate() error {
	return dara.Validate(s)
}
