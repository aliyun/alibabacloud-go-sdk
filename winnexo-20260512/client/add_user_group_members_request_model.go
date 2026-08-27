// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserGroupMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *AddUserGroupMembersRequest
	GetTenantId() *string
	SetUserGroupId(v string) *AddUserGroupMembersRequest
	GetUserGroupId() *string
	SetUserIds(v []*int64) *AddUserGroupMembersRequest
	GetUserIds() []*int64
}

type AddUserGroupMembersRequest struct {
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
	UserIds []*int64 `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s AddUserGroupMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *AddUserGroupMembersRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *AddUserGroupMembersRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AddUserGroupMembersRequest) GetUserIds() []*int64 {
	return s.UserIds
}

func (s *AddUserGroupMembersRequest) SetTenantId(v string) *AddUserGroupMembersRequest {
	s.TenantId = &v
	return s
}

func (s *AddUserGroupMembersRequest) SetUserGroupId(v string) *AddUserGroupMembersRequest {
	s.UserGroupId = &v
	return s
}

func (s *AddUserGroupMembersRequest) SetUserIds(v []*int64) *AddUserGroupMembersRequest {
	s.UserIds = v
	return s
}

func (s *AddUserGroupMembersRequest) Validate() error {
	return dara.Validate(s)
}
