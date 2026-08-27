// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateUserGroupRequest
	GetDescription() *string
	SetMoveToRoot(v bool) *UpdateUserGroupRequest
	GetMoveToRoot() *bool
	SetParentId(v string) *UpdateUserGroupRequest
	GetParentId() *string
	SetTenantId(v string) *UpdateUserGroupRequest
	GetTenantId() *string
	SetUserGroupId(v string) *UpdateUserGroupRequest
	GetUserGroupId() *string
	SetUserGroupName(v string) *UpdateUserGroupRequest
	GetUserGroupName() *string
}

type UpdateUserGroupRequest struct {
	// The new description of the user group. If not specified, the description is not modified.
	//
	// example:
	//
	// South China Sales Organization
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to move the user group to the root node. This parameter cannot be set together with parentId.
	//
	// example:
	//
	// false
	MoveToRoot *bool `json:"moveToRoot,omitempty" xml:"moveToRoot,omitempty"`
	// The ID of the new parent user group. If not specified, the user group is not moved.
	//
	// example:
	//
	// 7ea8973f-7a5c-4e8a-956b-4fe0e7e2eb11
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
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
	// The new name of the user group. If not specified, the name is not modified.
	//
	// example:
	//
	// South China Sales
	UserGroupName *string `json:"userGroupName,omitempty" xml:"userGroupName,omitempty"`
}

func (s UpdateUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateUserGroupRequest) GetMoveToRoot() *bool {
	return s.MoveToRoot
}

func (s *UpdateUserGroupRequest) GetParentId() *string {
	return s.ParentId
}

func (s *UpdateUserGroupRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *UpdateUserGroupRequest) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *UpdateUserGroupRequest) SetDescription(v string) *UpdateUserGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateUserGroupRequest) SetMoveToRoot(v bool) *UpdateUserGroupRequest {
	s.MoveToRoot = &v
	return s
}

func (s *UpdateUserGroupRequest) SetParentId(v string) *UpdateUserGroupRequest {
	s.ParentId = &v
	return s
}

func (s *UpdateUserGroupRequest) SetTenantId(v string) *UpdateUserGroupRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateUserGroupRequest) SetUserGroupId(v string) *UpdateUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *UpdateUserGroupRequest) SetUserGroupName(v string) *UpdateUserGroupRequest {
	s.UserGroupName = &v
	return s
}

func (s *UpdateUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
