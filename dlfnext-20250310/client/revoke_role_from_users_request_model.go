// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeRoleFromUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRolePrincipal(v string) *RevokeRoleFromUsersRequest
	GetRolePrincipal() *string
	SetUserPrincipals(v []*string) *RevokeRoleFromUsersRequest
	GetUserPrincipals() []*string
}

type RevokeRoleFromUsersRequest struct {
	// example:
	//
	// acs:dlf::[accountId]:role/role_name
	RolePrincipal  *string   `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
	UserPrincipals []*string `json:"userPrincipals,omitempty" xml:"userPrincipals,omitempty" type:"Repeated"`
}

func (s RevokeRoleFromUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeRoleFromUsersRequest) GoString() string {
	return s.String()
}

func (s *RevokeRoleFromUsersRequest) GetRolePrincipal() *string {
	return s.RolePrincipal
}

func (s *RevokeRoleFromUsersRequest) GetUserPrincipals() []*string {
	return s.UserPrincipals
}

func (s *RevokeRoleFromUsersRequest) SetRolePrincipal(v string) *RevokeRoleFromUsersRequest {
	s.RolePrincipal = &v
	return s
}

func (s *RevokeRoleFromUsersRequest) SetUserPrincipals(v []*string) *RevokeRoleFromUsersRequest {
	s.UserPrincipals = v
	return s
}

func (s *RevokeRoleFromUsersRequest) Validate() error {
	return dara.Validate(s)
}
