// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateUserRequest
	GetAppToken() *string
	SetDisplayName(v string) *UpdateUserRequest
	GetDisplayName() *string
	SetIsActive(v bool) *UpdateUserRequest
	GetIsActive() *bool
	SetOpsUserCode(v string) *UpdateUserRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateUserRequest
	GetOpsUserId() *string
	SetRoleCodes(v []*string) *UpdateUserRequest
	GetRoleCodes() []*string
	SetTenantId(v string) *UpdateUserRequest
	GetTenantId() *string
	SetUserGroupIds(v []*string) *UpdateUserRequest
	GetUserGroupIds() []*string
	SetWnUserId(v string) *UpdateUserRequest
	GetWnUserId() *string
}

type UpdateUserRequest struct {
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
	RoleCodes []*string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty" type:"Repeated"`
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
	UserGroupIds []*string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty" type:"Repeated"`
	// The ID of the target user, which is the WINNEXO platform user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateUserRequest) GetIsActive() *bool {
	return s.IsActive
}

func (s *UpdateUserRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateUserRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateUserRequest) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *UpdateUserRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateUserRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *UpdateUserRequest) GetWnUserId() *string {
	return s.WnUserId
}

func (s *UpdateUserRequest) SetAppToken(v string) *UpdateUserRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateUserRequest) SetDisplayName(v string) *UpdateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateUserRequest) SetIsActive(v bool) *UpdateUserRequest {
	s.IsActive = &v
	return s
}

func (s *UpdateUserRequest) SetOpsUserCode(v string) *UpdateUserRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateUserRequest) SetOpsUserId(v string) *UpdateUserRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateUserRequest) SetRoleCodes(v []*string) *UpdateUserRequest {
	s.RoleCodes = v
	return s
}

func (s *UpdateUserRequest) SetTenantId(v string) *UpdateUserRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateUserRequest) SetUserGroupIds(v []*string) *UpdateUserRequest {
	s.UserGroupIds = v
	return s
}

func (s *UpdateUserRequest) SetWnUserId(v string) *UpdateUserRequest {
	s.WnUserId = &v
	return s
}

func (s *UpdateUserRequest) Validate() error {
	return dara.Validate(s)
}
