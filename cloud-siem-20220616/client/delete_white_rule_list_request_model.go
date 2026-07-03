// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWhiteRuleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteWhiteRuleListRequest
	GetId() *int64
	SetRegionId(v string) *DeleteWhiteRuleListRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DeleteWhiteRuleListRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DeleteWhiteRuleListRequest
	GetRoleType() *int32
}

type DeleteWhiteRuleListRequest struct {
	// The unique ID of the whitelist rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region of the data management center for threat analysis. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: The assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: The assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. An administrator can switch to the perspective of the member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: The view for the current Alibaba Cloud account.
	//
	// - 1: The view for all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DeleteWhiteRuleListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWhiteRuleListRequest) GoString() string {
	return s.String()
}

func (s *DeleteWhiteRuleListRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteWhiteRuleListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteWhiteRuleListRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteWhiteRuleListRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DeleteWhiteRuleListRequest) SetId(v int64) *DeleteWhiteRuleListRequest {
	s.Id = &v
	return s
}

func (s *DeleteWhiteRuleListRequest) SetRegionId(v string) *DeleteWhiteRuleListRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteWhiteRuleListRequest) SetRoleFor(v int64) *DeleteWhiteRuleListRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteWhiteRuleListRequest) SetRoleType(v int32) *DeleteWhiteRuleListRequest {
	s.RoleType = &v
	return s
}

func (s *DeleteWhiteRuleListRequest) Validate() error {
	return dara.Validate(s)
}
