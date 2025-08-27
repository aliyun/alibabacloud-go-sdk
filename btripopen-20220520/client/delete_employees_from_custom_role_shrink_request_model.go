// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEmployeesFromCustomRoleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleId(v string) *DeleteEmployeesFromCustomRoleShrinkRequest
	GetRoleId() *string
	SetUserIdListShrink(v string) *DeleteEmployeesFromCustomRoleShrinkRequest
	GetUserIdListShrink() *string
}

type DeleteEmployeesFromCustomRoleShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
	// This parameter is required.
	UserIdListShrink *string `json:"user_id_list,omitempty" xml:"user_id_list,omitempty"`
}

func (s DeleteEmployeesFromCustomRoleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEmployeesFromCustomRoleShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteEmployeesFromCustomRoleShrinkRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *DeleteEmployeesFromCustomRoleShrinkRequest) GetUserIdListShrink() *string {
	return s.UserIdListShrink
}

func (s *DeleteEmployeesFromCustomRoleShrinkRequest) SetRoleId(v string) *DeleteEmployeesFromCustomRoleShrinkRequest {
	s.RoleId = &v
	return s
}

func (s *DeleteEmployeesFromCustomRoleShrinkRequest) SetUserIdListShrink(v string) *DeleteEmployeesFromCustomRoleShrinkRequest {
	s.UserIdListShrink = &v
	return s
}

func (s *DeleteEmployeesFromCustomRoleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
