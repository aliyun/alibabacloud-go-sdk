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
	SetCreatedBy(v string) *Role
	GetCreatedBy() *string
	SetDescription(v string) *Role
	GetDescription() *string
	SetDisplayName(v string) *Role
	GetDisplayName() *string
	SetIsPredefined(v string) *Role
	GetIsPredefined() *string
	SetRoleName(v string) *Role
	GetRoleName() *string
	SetRolePrincipal(v string) *Role
	GetRolePrincipal() *string
	SetUpdatedAt(v int64) *Role
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *Role
	GetUpdatedBy() *string
	SetUsers(v []*User) *Role
	GetUsers() []*User
}

type Role struct {
	// example:
	//
	// 1744970111419
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// acs:ram::[accountId]:root
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// example:
	//
	// role_description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// role_display_name
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// true
	IsPredefined *string `json:"isPredefined,omitempty" xml:"isPredefined,omitempty"`
	// example:
	//
	// role_name
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// example:
	//
	// acs:dlf::[accountId]:role/role_name
	RolePrincipal *string `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
	// example:
	//
	// 1744970111419
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// acs:ram::[accountId]:root
	UpdatedBy *string `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
	Users     []*User `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
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

func (s *Role) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *Role) GetDescription() *string {
	return s.Description
}

func (s *Role) GetDisplayName() *string {
	return s.DisplayName
}

func (s *Role) GetIsPredefined() *string {
	return s.IsPredefined
}

func (s *Role) GetRoleName() *string {
	return s.RoleName
}

func (s *Role) GetRolePrincipal() *string {
	return s.RolePrincipal
}

func (s *Role) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Role) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *Role) GetUsers() []*User {
	return s.Users
}

func (s *Role) SetCreatedAt(v int64) *Role {
	s.CreatedAt = &v
	return s
}

func (s *Role) SetCreatedBy(v string) *Role {
	s.CreatedBy = &v
	return s
}

func (s *Role) SetDescription(v string) *Role {
	s.Description = &v
	return s
}

func (s *Role) SetDisplayName(v string) *Role {
	s.DisplayName = &v
	return s
}

func (s *Role) SetIsPredefined(v string) *Role {
	s.IsPredefined = &v
	return s
}

func (s *Role) SetRoleName(v string) *Role {
	s.RoleName = &v
	return s
}

func (s *Role) SetRolePrincipal(v string) *Role {
	s.RolePrincipal = &v
	return s
}

func (s *Role) SetUpdatedAt(v int64) *Role {
	s.UpdatedAt = &v
	return s
}

func (s *Role) SetUpdatedBy(v string) *Role {
	s.UpdatedBy = &v
	return s
}

func (s *Role) SetUsers(v []*User) *Role {
	s.Users = v
	return s
}

func (s *Role) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
