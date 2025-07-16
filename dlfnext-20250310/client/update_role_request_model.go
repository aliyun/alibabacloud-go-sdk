// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateRoleRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateRoleRequest
	GetDisplayName() *string
	SetRolePrincipal(v string) *UpdateRoleRequest
	GetRolePrincipal() *string
}

type UpdateRoleRequest struct {
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
	// acs:dlf::[accountId]:role/role_name
	RolePrincipal *string `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
}

func (s UpdateRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRoleRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateRoleRequest) GetRolePrincipal() *string {
	return s.RolePrincipal
}

func (s *UpdateRoleRequest) SetDescription(v string) *UpdateRoleRequest {
	s.Description = &v
	return s
}

func (s *UpdateRoleRequest) SetDisplayName(v string) *UpdateRoleRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateRoleRequest) SetRolePrincipal(v string) *UpdateRoleRequest {
	s.RolePrincipal = &v
	return s
}

func (s *UpdateRoleRequest) Validate() error {
	return dara.Validate(s)
}
