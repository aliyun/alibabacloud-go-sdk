// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEmployeesToCustomRoleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleId(v string) *AddEmployeesToCustomRoleShrinkRequest
	GetRoleId() *string
	SetUserIdListShrink(v string) *AddEmployeesToCustomRoleShrinkRequest
	GetUserIdListShrink() *string
}

type AddEmployeesToCustomRoleShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	UserIdListShrink *string `json:"user_id_list,omitempty" xml:"user_id_list,omitempty"`
}

func (s AddEmployeesToCustomRoleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddEmployeesToCustomRoleShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddEmployeesToCustomRoleShrinkRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *AddEmployeesToCustomRoleShrinkRequest) GetUserIdListShrink() *string {
	return s.UserIdListShrink
}

func (s *AddEmployeesToCustomRoleShrinkRequest) SetRoleId(v string) *AddEmployeesToCustomRoleShrinkRequest {
	s.RoleId = &v
	return s
}

func (s *AddEmployeesToCustomRoleShrinkRequest) SetUserIdListShrink(v string) *AddEmployeesToCustomRoleShrinkRequest {
	s.UserIdListShrink = &v
	return s
}

func (s *AddEmployeesToCustomRoleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
