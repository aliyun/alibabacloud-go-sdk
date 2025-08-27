// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleId(v string) *UpdateCustomRoleRequest
	GetRoleId() *string
	SetRoleName(v string) *UpdateCustomRoleRequest
	GetRoleName() *string
}

type UpdateCustomRoleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
	// This parameter is required.
	RoleName *string `json:"role_name,omitempty" xml:"role_name,omitempty"`
}

func (s UpdateCustomRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoleRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *UpdateCustomRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *UpdateCustomRoleRequest) SetRoleId(v string) *UpdateCustomRoleRequest {
	s.RoleId = &v
	return s
}

func (s *UpdateCustomRoleRequest) SetRoleName(v string) *UpdateCustomRoleRequest {
	s.RoleName = &v
	return s
}

func (s *UpdateCustomRoleRequest) Validate() error {
	return dara.Validate(s)
}
