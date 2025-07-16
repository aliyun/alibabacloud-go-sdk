// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantRoleToUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRolePrincipal(v string) *GrantRoleToUsersRequest
	GetRolePrincipal() *string
	SetUserPrincipals(v []*string) *GrantRoleToUsersRequest
	GetUserPrincipals() []*string
}

type GrantRoleToUsersRequest struct {
	// example:
	//
	// acs:dlf::[accountId]:role/role_name
	RolePrincipal  *string   `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
	UserPrincipals []*string `json:"userPrincipals,omitempty" xml:"userPrincipals,omitempty" type:"Repeated"`
}

func (s GrantRoleToUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantRoleToUsersRequest) GoString() string {
	return s.String()
}

func (s *GrantRoleToUsersRequest) GetRolePrincipal() *string {
	return s.RolePrincipal
}

func (s *GrantRoleToUsersRequest) GetUserPrincipals() []*string {
	return s.UserPrincipals
}

func (s *GrantRoleToUsersRequest) SetRolePrincipal(v string) *GrantRoleToUsersRequest {
	s.RolePrincipal = &v
	return s
}

func (s *GrantRoleToUsersRequest) SetUserPrincipals(v []*string) *GrantRoleToUsersRequest {
	s.UserPrincipals = v
	return s
}

func (s *GrantRoleToUsersRequest) Validate() error {
	return dara.Validate(s)
}
