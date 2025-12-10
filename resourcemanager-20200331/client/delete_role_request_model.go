// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleName(v string) *DeleteRoleRequest
	GetRoleName() *string
}

type DeleteRoleRequest struct {
	// The name of the role.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, periods (.), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s DeleteRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *DeleteRoleRequest) SetRoleName(v string) *DeleteRoleRequest {
	s.RoleName = &v
	return s
}

func (s *DeleteRoleRequest) Validate() error {
	return dara.Validate(s)
}
