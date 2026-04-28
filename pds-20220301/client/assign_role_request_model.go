// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentity(v *Identity) *AssignRoleRequest
	GetIdentity() *Identity
	SetManageResourceId(v string) *AssignRoleRequest
	GetManageResourceId() *string
	SetManageResourceType(v string) *AssignRoleRequest
	GetManageResourceType() *string
	SetRoleId(v string) *AssignRoleRequest
	GetRoleId() *string
}

type AssignRoleRequest struct {
	// The unique identifier of a user. The group administrator role can only be assigned to a user.
	//
	// This parameter is required.
	Identity *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
	// The ID of the resource that the role can manage. You can only set this parameter to the ID of a group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 105***b82
	ManageResourceId *string `json:"manage_resource_id,omitempty" xml:"manage_resource_id,omitempty"`
	// The type of the resource that the role can manage. Valid value: RT_Group.
	//
	// This parameter is required.
	//
	// example:
	//
	// RT_Group
	ManageResourceType *string `json:"manage_resource_type,omitempty" xml:"manage_resource_type,omitempty"`
	// The ID of the role that is assigned to a user. Valid value: SystemGroupAdmin.
	//
	// This parameter is required.
	//
	// example:
	//
	// SystemGroupAdmin
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

func (s AssignRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignRoleRequest) GoString() string {
	return s.String()
}

func (s *AssignRoleRequest) GetIdentity() *Identity {
	return s.Identity
}

func (s *AssignRoleRequest) GetManageResourceId() *string {
	return s.ManageResourceId
}

func (s *AssignRoleRequest) GetManageResourceType() *string {
	return s.ManageResourceType
}

func (s *AssignRoleRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *AssignRoleRequest) SetIdentity(v *Identity) *AssignRoleRequest {
	s.Identity = v
	return s
}

func (s *AssignRoleRequest) SetManageResourceId(v string) *AssignRoleRequest {
	s.ManageResourceId = &v
	return s
}

func (s *AssignRoleRequest) SetManageResourceType(v string) *AssignRoleRequest {
	s.ManageResourceType = &v
	return s
}

func (s *AssignRoleRequest) SetRoleId(v string) *AssignRoleRequest {
	s.RoleId = &v
	return s
}

func (s *AssignRoleRequest) Validate() error {
	if s.Identity != nil {
		if err := s.Identity.Validate(); err != nil {
			return err
		}
	}
	return nil
}
