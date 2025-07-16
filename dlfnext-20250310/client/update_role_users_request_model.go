// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoleUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRolePrincipal(v string) *UpdateRoleUsersRequest
	GetRolePrincipal() *string
	SetUserPrincipals(v []*string) *UpdateRoleUsersRequest
	GetUserPrincipals() []*string
}

type UpdateRoleUsersRequest struct {
	// example:
	//
	// acs:dlf::[accountId]:role/role_name
	RolePrincipal  *string   `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
	UserPrincipals []*string `json:"userPrincipals,omitempty" xml:"userPrincipals,omitempty" type:"Repeated"`
}

func (s UpdateRoleUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoleUsersRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleUsersRequest) GetRolePrincipal() *string {
	return s.RolePrincipal
}

func (s *UpdateRoleUsersRequest) GetUserPrincipals() []*string {
	return s.UserPrincipals
}

func (s *UpdateRoleUsersRequest) SetRolePrincipal(v string) *UpdateRoleUsersRequest {
	s.RolePrincipal = &v
	return s
}

func (s *UpdateRoleUsersRequest) SetUserPrincipals(v []*string) *UpdateRoleUsersRequest {
	s.UserPrincipals = v
	return s
}

func (s *UpdateRoleUsersRequest) Validate() error {
	return dara.Validate(s)
}
