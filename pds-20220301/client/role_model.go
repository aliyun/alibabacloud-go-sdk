// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRole interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *Role
	GetCreatedAt() *int64
	SetCreator(v string) *Role
	GetCreator() *string
	SetDescription(v string) *Role
	GetDescription() *string
	SetManageResourceType(v string) *Role
	GetManageResourceType() *string
	SetName(v string) *Role
	GetName() *string
	SetPermissions(v []*Permission) *Role
	GetPermissions() []*Permission
	SetRoleId(v string) *Role
	GetRoleId() *string
	SetStatus(v string) *Role
	GetStatus() *string
	SetUpdatedAt(v int64) *Role
	GetUpdatedAt() *int64
}

type Role struct {
	// The time when the role was created. The value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1703648502811
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The ID of the user who created the role.
	//
	// example:
	//
	// a23***
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The description of the role.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The type of the resource on which the role has permissions. Valid value: RT_File.
	//
	// example:
	//
	// RT_File
	ManageResourceType *string `json:"manage_resource_type,omitempty" xml:"manage_resource_type,omitempty"`
	// The name of the role.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The permissions.
	Permissions []*Permission `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
	// The ID of the role.
	//
	// example:
	//
	// f2a***
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
	// The status of the role.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the role was modified. The value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1703648502811
	UpdatedAt *int64 `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s Role) String() string {
	return dara.Prettify(s)
}

func (s Role) GoString() string {
	return s.String()
}

func (s *Role) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Role) GetCreator() *string {
	return s.Creator
}

func (s *Role) GetDescription() *string {
	return s.Description
}

func (s *Role) GetManageResourceType() *string {
	return s.ManageResourceType
}

func (s *Role) GetName() *string {
	return s.Name
}

func (s *Role) GetPermissions() []*Permission {
	return s.Permissions
}

func (s *Role) GetRoleId() *string {
	return s.RoleId
}

func (s *Role) GetStatus() *string {
	return s.Status
}

func (s *Role) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Role) SetCreatedAt(v int64) *Role {
	s.CreatedAt = &v
	return s
}

func (s *Role) SetCreator(v string) *Role {
	s.Creator = &v
	return s
}

func (s *Role) SetDescription(v string) *Role {
	s.Description = &v
	return s
}

func (s *Role) SetManageResourceType(v string) *Role {
	s.ManageResourceType = &v
	return s
}

func (s *Role) SetName(v string) *Role {
	s.Name = &v
	return s
}

func (s *Role) SetPermissions(v []*Permission) *Role {
	s.Permissions = v
	return s
}

func (s *Role) SetRoleId(v string) *Role {
	s.RoleId = &v
	return s
}

func (s *Role) SetStatus(v string) *Role {
	s.Status = &v
	return s
}

func (s *Role) SetUpdatedAt(v int64) *Role {
	s.UpdatedAt = &v
	return s
}

func (s *Role) Validate() error {
	if s.Permissions != nil {
		for _, item := range s.Permissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
