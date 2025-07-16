// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRolePrincipal(v string) *GetRoleRequest
	GetRolePrincipal() *string
}

type GetRoleRequest struct {
	// example:
	//
	// acs:dlf::[accountId]:role/role_name
	RolePrincipal *string `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
}

func (s GetRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRoleRequest) GoString() string {
	return s.String()
}

func (s *GetRoleRequest) GetRolePrincipal() *string {
	return s.RolePrincipal
}

func (s *GetRoleRequest) SetRolePrincipal(v string) *GetRoleRequest {
	s.RolePrincipal = &v
	return s
}

func (s *GetRoleRequest) Validate() error {
	return dara.Validate(s)
}
