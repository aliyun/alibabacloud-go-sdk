// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateUserShrinkRequest
	GetAppToken() *string
	SetDisplayName(v string) *UpdateUserShrinkRequest
	GetDisplayName() *string
	SetIsActive(v bool) *UpdateUserShrinkRequest
	GetIsActive() *bool
	SetOpsUserCode(v string) *UpdateUserShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateUserShrinkRequest
	GetOpsUserId() *string
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
	// The platform application token issued by the console. If specified, the operation is executed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The display name of the user.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Specifies whether the account is activated. Valid values:
	//
	// - **true**: Activated.
	//
	// - **false**: Not activated.
	//
	// example:
	//
	// true
	IsActive *bool `json:"isActive,omitempty" xml:"isActive,omitempty"`
	// The proxy user code, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The new list of system role codes. This list fully replaces the existing roles and must contain at least one role. Valid values: SUPER_ADMIN, SYSTEM_ADMIN, SEMANTIC_ADMIN, SKILL_ADMIN, KB_ADMIN, AGENT_ADMIN, and APPLICATION_USER.
	//
	// example:
	//
	// ["APPLICATION_USER"]
	RoleCodesShrink *string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty"`
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The new list of user group IDs. This list fully replaces the existing user groups. If this parameter is not specified, the user groups are not modified.
	//
	// example:
	//
	// ["string_value"]
	UserGroupIdsShrink *string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty"`
	// The ID of the target user, which is the WINNEXO platform user ID.
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

func (s *UpdateUserShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateUserShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateUserShrinkRequest) GetIsActive() *bool {
	return s.IsActive
}

func (s *UpdateUserShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateUserShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *UpdateUserShrinkRequest) SetAppToken(v string) *UpdateUserShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateUserShrinkRequest) SetDisplayName(v string) *UpdateUserShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateUserShrinkRequest) SetIsActive(v bool) *UpdateUserShrinkRequest {
	s.IsActive = &v
	return s
}

func (s *UpdateUserShrinkRequest) SetOpsUserCode(v string) *UpdateUserShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateUserShrinkRequest) SetOpsUserId(v string) *UpdateUserShrinkRequest {
	s.OpsUserId = &v
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
