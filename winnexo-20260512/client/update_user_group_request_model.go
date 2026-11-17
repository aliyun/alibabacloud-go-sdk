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
	// The new description of the user group. If this parameter is not specified, the description is not modified.
	//
	// example:
	//
	// string_value
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to move the user group to the root node. This parameter cannot be set at the same time as a non-null parentId.
	//
	// example:
	//
	// false
	MoveToRoot *bool `json:"moveToRoot,omitempty" xml:"moveToRoot,omitempty"`
	// The ID of the new parent user group. If this parameter is not specified, the user group is not moved. If it is explicitly set to null, the user group is moved to the root node.
	//
	// example:
	//
	// string_value
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, it is explicitly passed by using --tenant-id.
	//
	// example:
	//
	// string_value
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the target user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	UserGroupId *string `json:"userGroupId,omitempty" xml:"userGroupId,omitempty"`
	// The new name of the user group. If this parameter is not specified, the name is not modified.
	//
	// example:
	//
	// string_value
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
