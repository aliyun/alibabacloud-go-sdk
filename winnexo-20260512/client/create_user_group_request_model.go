// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateUserGroupRequest
	GetDescription() *string
	SetParentId(v string) *CreateUserGroupRequest
	GetParentId() *string
	SetTenantId(v string) *CreateUserGroupRequest
	GetTenantId() *string
	SetUserGroupName(v string) *CreateUserGroupRequest
	GetUserGroupName() *string
}

type CreateUserGroupRequest struct {
	// The description of the user group.
	//
	// example:
	//
	// Sales organization
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the parent user group. If this parameter is not specified, a root node is created.
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
	// The name of the user group. The name must be unique under the same parent node.
	//
	// This parameter is required.
	//
	// example:
	//
	// East China Sales
	UserGroupName *string `json:"userGroupName,omitempty" xml:"userGroupName,omitempty"`
}

func (s CreateUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUserGroupRequest) GetParentId() *string {
	return s.ParentId
}

func (s *CreateUserGroupRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateUserGroupRequest) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *CreateUserGroupRequest) SetDescription(v string) *CreateUserGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateUserGroupRequest) SetParentId(v string) *CreateUserGroupRequest {
	s.ParentId = &v
	return s
}

func (s *CreateUserGroupRequest) SetTenantId(v string) *CreateUserGroupRequest {
	s.TenantId = &v
	return s
}

func (s *CreateUserGroupRequest) SetUserGroupName(v string) *CreateUserGroupRequest {
	s.UserGroupName = &v
	return s
}

func (s *CreateUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
