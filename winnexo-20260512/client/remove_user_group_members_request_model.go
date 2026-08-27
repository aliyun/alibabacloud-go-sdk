// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserGroupMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *RemoveUserGroupMembersRequest
	GetTenantId() *string
	SetUserGroupId(v string) *RemoveUserGroupMembersRequest
	GetUserGroupId() *string
	SetUserIds(v []*int64) *RemoveUserGroupMembersRequest
	GetUserIds() []*int64
}

type RemoveUserGroupMembersRequest struct {
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
	UserIds []*int64 `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s RemoveUserGroupMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMembersRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RemoveUserGroupMembersRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *RemoveUserGroupMembersRequest) GetUserIds() []*int64 {
	return s.UserIds
}

func (s *RemoveUserGroupMembersRequest) SetTenantId(v string) *RemoveUserGroupMembersRequest {
	s.TenantId = &v
	return s
}

func (s *RemoveUserGroupMembersRequest) SetUserGroupId(v string) *RemoveUserGroupMembersRequest {
	s.UserGroupId = &v
	return s
}

func (s *RemoveUserGroupMembersRequest) SetUserIds(v []*int64) *RemoveUserGroupMembersRequest {
	s.UserIds = v
	return s
}

func (s *RemoveUserGroupMembersRequest) Validate() error {
	return dara.Validate(s)
}
