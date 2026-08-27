// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserGroupMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *AddUserGroupMembersShrinkRequest
	GetTenantId() *string
	SetUserGroupId(v string) *AddUserGroupMembersShrinkRequest
	GetUserGroupId() *string
	SetUserIdsShrink(v string) *AddUserGroupMembersShrinkRequest
	GetUserIdsShrink() *string
}

type AddUserGroupMembersShrinkRequest struct {
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this parameter explicitly by using `--tenant-id`.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the target user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7ea8973f-7a5c-4e8a-956b-4fe0e7e2eb11
	UserGroupId *string `json:"userGroupId,omitempty" xml:"userGroupId,omitempty"`
	// The list of platform user IDs to add. Supports single or batch input. Duplicate relationships are idempotent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserIdsShrink *string `json:"userIds,omitempty" xml:"userIds,omitempty"`
}

func (s AddUserGroupMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserGroupMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddUserGroupMembersShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *AddUserGroupMembersShrinkRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AddUserGroupMembersShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *AddUserGroupMembersShrinkRequest) SetTenantId(v string) *AddUserGroupMembersShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *AddUserGroupMembersShrinkRequest) SetUserGroupId(v string) *AddUserGroupMembersShrinkRequest {
	s.UserGroupId = &v
	return s
}

func (s *AddUserGroupMembersShrinkRequest) SetUserIdsShrink(v string) *AddUserGroupMembersShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *AddUserGroupMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
