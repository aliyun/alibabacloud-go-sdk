// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *UpdateUserRequest
	GetDisplayName() *string
	SetIsActive(v bool) *UpdateUserRequest
	GetIsActive() *bool
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
	RoleCodes []*string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty" type:"Repeated"`
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
	UserGroupIds []*string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty" type:"Repeated"`
	// The ID of the target user (WINNEXO platform user ID).
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

func (s *UpdateUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateUserRequest) GetIsActive() *bool {
	return s.IsActive
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

func (s *UpdateUserRequest) SetDisplayName(v string) *UpdateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateUserRequest) SetIsActive(v bool) *UpdateUserRequest {
	s.IsActive = &v
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
