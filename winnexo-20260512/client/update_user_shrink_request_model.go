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
	// The display name of the user.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Specifies whether the account is activated. Valid values:
	//
	//  - **true**: Activated.
	//
	// - **false**: Not activated.
	//
	// example:
	//
	// true
	IsActive *bool `json:"isActive,omitempty" xml:"isActive,omitempty"`
	// The new list of system role codes (full replacement, must contain at least one role). Valid values: SUPER_ADMIN / SYSTEM_ADMIN / SEMANTIC_ADMIN / SKILL_ADMIN / KB_ADMIN / AGENT_ADMIN / APPLICATION_USER.
	//
	// example:
	//
	// string_value
	RoleCodesShrink *string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty"`
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The new list of user group IDs (full replacement. If not specified, the value is not modified).
	//
	// example:
	//
	// string_value
	UserGroupIdsShrink *string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty"`
	// The ID of the target user (WINNEXO platform user ID).
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
