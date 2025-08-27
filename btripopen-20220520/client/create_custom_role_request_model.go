// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleId(v string) *CreateCustomRoleRequest
	GetRoleId() *string
	SetRoleName(v string) *CreateCustomRoleRequest
	GetRoleName() *string
}

type CreateCustomRoleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123abc
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
	// This parameter is required.
	RoleName *string `json:"role_name,omitempty" xml:"role_name,omitempty"`
}

func (s CreateCustomRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomRoleRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *CreateCustomRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateCustomRoleRequest) SetRoleId(v string) *CreateCustomRoleRequest {
	s.RoleId = &v
	return s
}

func (s *CreateCustomRoleRequest) SetRoleName(v string) *CreateCustomRoleRequest {
	s.RoleName = &v
	return s
}

func (s *CreateCustomRoleRequest) Validate() error {
	return dara.Validate(s)
}
