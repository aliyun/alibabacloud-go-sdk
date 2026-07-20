// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleId(v string) *DeleteCustomRoleRequest
	GetRoleId() *string
}

type DeleteCustomRoleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

func (s DeleteCustomRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoleRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *DeleteCustomRoleRequest) SetRoleId(v string) *DeleteCustomRoleRequest {
	s.RoleId = &v
	return s
}

func (s *DeleteCustomRoleRequest) Validate() error {
	return dara.Validate(s)
}
