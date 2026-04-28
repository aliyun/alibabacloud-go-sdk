// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAssignRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentity(v *Identity) *CancelAssignRoleRequest
	GetIdentity() *Identity
	SetManageResourceId(v string) *CancelAssignRoleRequest
	GetManageResourceId() *string
	SetManageResourceType(v string) *CancelAssignRoleRequest
	GetManageResourceType() *string
	SetRoleId(v string) *CancelAssignRoleRequest
	GetRoleId() *string
}

type CancelAssignRoleRequest struct {
	// The unique identifier. You can cancel only the role assigned to a user.
	//
	// This parameter is required.
	Identity *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
	// The ID of the resource that the role manages. Set the value to a group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 105***b82
	ManageResourceId *string `json:"manage_resource_id,omitempty" xml:"manage_resource_id,omitempty"`
	// The type of the resource that the role manages. Set the value to RT_Group, which specifies group.
	//
	// This parameter is required.
	//
	// example:
	//
	// RT_Group
	ManageResourceType *string `json:"manage_resource_type,omitempty" xml:"manage_resource_type,omitempty"`
	// The ID of the role to be canceled. Set the value to SystemGroupAdmin, which is the ID of the group administrator role.
	//
	// This parameter is required.
	//
	// example:
	//
	// SystemGroupAdmin
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

func (s CancelAssignRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAssignRoleRequest) GoString() string {
	return s.String()
}

func (s *CancelAssignRoleRequest) GetIdentity() *Identity {
	return s.Identity
}

func (s *CancelAssignRoleRequest) GetManageResourceId() *string {
	return s.ManageResourceId
}

func (s *CancelAssignRoleRequest) GetManageResourceType() *string {
	return s.ManageResourceType
}

func (s *CancelAssignRoleRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *CancelAssignRoleRequest) SetIdentity(v *Identity) *CancelAssignRoleRequest {
	s.Identity = v
	return s
}

func (s *CancelAssignRoleRequest) SetManageResourceId(v string) *CancelAssignRoleRequest {
	s.ManageResourceId = &v
	return s
}

func (s *CancelAssignRoleRequest) SetManageResourceType(v string) *CancelAssignRoleRequest {
	s.ManageResourceType = &v
	return s
}

func (s *CancelAssignRoleRequest) SetRoleId(v string) *CancelAssignRoleRequest {
	s.RoleId = &v
	return s
}

func (s *CancelAssignRoleRequest) Validate() error {
	if s.Identity != nil {
		if err := s.Identity.Validate(); err != nil {
			return err
		}
	}
	return nil
}
