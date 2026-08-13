// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *UpdateUserShrinkRequest
	GetDisplayName() *string
	SetIsActive(v bool) *UpdateUserShrinkRequest
	GetIsActive() *bool
	SetRoleCodesShrink(v string) *UpdateUserShrinkRequest
	GetRoleCodesShrink() *string
	SetTenantId(v string) *UpdateUserShrinkRequest
	GetTenantId() *string
	SetUserGroupIdsShrink(v string) *UpdateUserShrinkRequest
	GetUserGroupIdsShrink() *string
	SetWnUserId(v string) *UpdateUserShrinkRequest
	GetWnUserId() *string
}

type UpdateUserShrinkRequest struct {
	// 新的显示名称（不传不修改，传则不可为空，最多100字）
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 启用/停用状态（不传不修改）。false=停用，true=启用
	//
	// example:
	//
	// true
	IsActive *bool `json:"isActive,omitempty" xml:"isActive,omitempty"`
	// 新的系统角色 code 列表（全量替换，至少包含一个角色）。可选值: SUPER_ADMIN / SYSTEM_ADMIN / SEMANTIC_ADMIN / SKILL_ADMIN / KB_ADMIN / AGENT_ADMIN / APPLICATION_USER
	//
	// example:
	//
	// string_value
	RoleCodesShrink *string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 新的用户组ID列表（全量替换，不传不修改）
	//
	// example:
	//
	// string_value
	UserGroupIdsShrink *string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty"`
	// 目标用户ID（WINNEXO 平台用户ID）
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s UpdateUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateUserShrinkRequest) GetIsActive() *bool {
	return s.IsActive
}

func (s *UpdateUserShrinkRequest) GetRoleCodesShrink() *string {
	return s.RoleCodesShrink
}

func (s *UpdateUserShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateUserShrinkRequest) GetUserGroupIdsShrink() *string {
	return s.UserGroupIdsShrink
}

func (s *UpdateUserShrinkRequest) GetWnUserId() *string {
	return s.WnUserId
}

func (s *UpdateUserShrinkRequest) SetDisplayName(v string) *UpdateUserShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateUserShrinkRequest) SetIsActive(v bool) *UpdateUserShrinkRequest {
	s.IsActive = &v
	return s
}

func (s *UpdateUserShrinkRequest) SetRoleCodesShrink(v string) *UpdateUserShrinkRequest {
	s.RoleCodesShrink = &v
	return s
}

func (s *UpdateUserShrinkRequest) SetTenantId(v string) *UpdateUserShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateUserShrinkRequest) SetUserGroupIdsShrink(v string) *UpdateUserShrinkRequest {
	s.UserGroupIdsShrink = &v
	return s
}

func (s *UpdateUserShrinkRequest) SetWnUserId(v string) *UpdateUserShrinkRequest {
	s.WnUserId = &v
	return s
}

func (s *UpdateUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
