// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEmployeesToCustomRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleId(v string) *AddEmployeesToCustomRoleRequest
	GetRoleId() *string
	SetUserIdList(v []*string) *AddEmployeesToCustomRoleRequest
	GetUserIdList() []*string
}

type AddEmployeesToCustomRoleRequest struct {
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
	UserIdList []*string `json:"user_id_list,omitempty" xml:"user_id_list,omitempty" type:"Repeated"`
}

func (s AddEmployeesToCustomRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddEmployeesToCustomRoleRequest) GoString() string {
	return s.String()
}

func (s *AddEmployeesToCustomRoleRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *AddEmployeesToCustomRoleRequest) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *AddEmployeesToCustomRoleRequest) SetRoleId(v string) *AddEmployeesToCustomRoleRequest {
	s.RoleId = &v
	return s
}

func (s *AddEmployeesToCustomRoleRequest) SetUserIdList(v []*string) *AddEmployeesToCustomRoleRequest {
	s.UserIdList = v
	return s
}

func (s *AddEmployeesToCustomRoleRequest) Validate() error {
	return dara.Validate(s)
}
