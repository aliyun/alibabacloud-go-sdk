// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEmployeesFromCustomRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleId(v string) *DeleteEmployeesFromCustomRoleRequest
	GetRoleId() *string
	SetUserIdList(v []*string) *DeleteEmployeesFromCustomRoleRequest
	GetUserIdList() []*string
}

type DeleteEmployeesFromCustomRoleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
	// This parameter is required.
	UserIdList []*string `json:"user_id_list,omitempty" xml:"user_id_list,omitempty" type:"Repeated"`
}

func (s DeleteEmployeesFromCustomRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEmployeesFromCustomRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteEmployeesFromCustomRoleRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *DeleteEmployeesFromCustomRoleRequest) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *DeleteEmployeesFromCustomRoleRequest) SetRoleId(v string) *DeleteEmployeesFromCustomRoleRequest {
	s.RoleId = &v
	return s
}

func (s *DeleteEmployeesFromCustomRoleRequest) SetUserIdList(v []*string) *DeleteEmployeesFromCustomRoleRequest {
	s.UserIdList = v
	return s
}

func (s *DeleteEmployeesFromCustomRoleRequest) Validate() error {
	return dara.Validate(s)
}
