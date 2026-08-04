// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDepartmentRoleCmd interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *DepartmentRoleCmd
	GetClientId() *int64
	SetRoleCode(v string) *DepartmentRoleCmd
	GetRoleCode() *string
}

type DepartmentRoleCmd struct {
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// example:
	//
	// member
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
}

func (s DepartmentRoleCmd) String() string {
	return dara.Prettify(s)
}

func (s DepartmentRoleCmd) GoString() string {
	return s.String()
}

func (s *DepartmentRoleCmd) GetClientId() *int64 {
	return s.ClientId
}

func (s *DepartmentRoleCmd) GetRoleCode() *string {
	return s.RoleCode
}

func (s *DepartmentRoleCmd) SetClientId(v int64) *DepartmentRoleCmd {
	s.ClientId = &v
	return s
}

func (s *DepartmentRoleCmd) SetRoleCode(v string) *DepartmentRoleCmd {
	s.RoleCode = &v
	return s
}

func (s *DepartmentRoleCmd) Validate() error {
	return dara.Validate(s)
}
