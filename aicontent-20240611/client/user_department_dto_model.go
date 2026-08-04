// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserDepartmentDTO interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *UserDepartmentDTO
	GetClientId() *int64
	SetClientName(v string) *UserDepartmentDTO
	GetClientName() *string
	SetRoleCode(v string) *UserDepartmentDTO
	GetRoleCode() *string
	SetRoleName(v string) *UserDepartmentDTO
	GetRoleName() *string
}

type UserDepartmentDTO struct {
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// example:
	//
	// 研发部
	ClientName *string `json:"clientName,omitempty" xml:"clientName,omitempty"`
	// example:
	//
	// member
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	// example:
	//
	// 成员
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s UserDepartmentDTO) String() string {
	return dara.Prettify(s)
}

func (s UserDepartmentDTO) GoString() string {
	return s.String()
}

func (s *UserDepartmentDTO) GetClientId() *int64 {
	return s.ClientId
}

func (s *UserDepartmentDTO) GetClientName() *string {
	return s.ClientName
}

func (s *UserDepartmentDTO) GetRoleCode() *string {
	return s.RoleCode
}

func (s *UserDepartmentDTO) GetRoleName() *string {
	return s.RoleName
}

func (s *UserDepartmentDTO) SetClientId(v int64) *UserDepartmentDTO {
	s.ClientId = &v
	return s
}

func (s *UserDepartmentDTO) SetClientName(v string) *UserDepartmentDTO {
	s.ClientName = &v
	return s
}

func (s *UserDepartmentDTO) SetRoleCode(v string) *UserDepartmentDTO {
	s.RoleCode = &v
	return s
}

func (s *UserDepartmentDTO) SetRoleName(v string) *UserDepartmentDTO {
	s.RoleName = &v
	return s
}

func (s *UserDepartmentDTO) Validate() error {
	return dara.Validate(s)
}
