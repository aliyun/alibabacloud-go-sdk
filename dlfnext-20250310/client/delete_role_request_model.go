// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRolePrincipal(v string) *DeleteRoleRequest
	GetRolePrincipal() *string
}

type DeleteRoleRequest struct {
	// example:
	//
	// acs:dlf::[accountId]:role/role_name
	RolePrincipal *string `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
}

func (s DeleteRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoleRequest) GetRolePrincipal() *string {
	return s.RolePrincipal
}

func (s *DeleteRoleRequest) SetRolePrincipal(v string) *DeleteRoleRequest {
	s.RolePrincipal = &v
	return s
}

func (s *DeleteRoleRequest) Validate() error {
	return dara.Validate(s)
}
