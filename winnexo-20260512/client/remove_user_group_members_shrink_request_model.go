// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserGroupMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *RemoveUserGroupMembersShrinkRequest
	GetTenantId() *string
	SetUserGroupId(v string) *RemoveUserGroupMembersShrinkRequest
	GetUserGroupId() *string
	SetUserIdsShrink(v string) *RemoveUserGroupMembersShrinkRequest
	GetUserIdsShrink() *string
}

type RemoveUserGroupMembersShrinkRequest struct {
	// The tenant ID. This is a common parameter. You can explicitly pass this parameter in winnexo-cli by using `--tenant-id`.
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
	// The list of platform user IDs to be removed. You can specify one or more IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserIdsShrink *string `json:"userIds,omitempty" xml:"userIds,omitempty"`
}

func (s RemoveUserGroupMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserGroupMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMembersShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RemoveUserGroupMembersShrinkRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *RemoveUserGroupMembersShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *RemoveUserGroupMembersShrinkRequest) SetTenantId(v string) *RemoveUserGroupMembersShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *RemoveUserGroupMembersShrinkRequest) SetUserGroupId(v string) *RemoveUserGroupMembersShrinkRequest {
	s.UserGroupId = &v
	return s
}

func (s *RemoveUserGroupMembersShrinkRequest) SetUserIdsShrink(v string) *RemoveUserGroupMembersShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *RemoveUserGroupMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
