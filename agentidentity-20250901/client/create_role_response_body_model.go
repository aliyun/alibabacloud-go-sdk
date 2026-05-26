// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRoleResponseBody
	GetRequestId() *string
	SetRole(v *CreateRoleResponseBodyRole) *CreateRoleResponseBody
	GetRole() *CreateRoleResponseBodyRole
}

type CreateRoleResponseBody struct {
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Role      *CreateRoleResponseBodyRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
}

func (s CreateRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRoleResponseBody) GetRole() *CreateRoleResponseBodyRole {
	return s.Role
}

func (s *CreateRoleResponseBody) SetRequestId(v string) *CreateRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoleResponseBody) SetRole(v *CreateRoleResponseBodyRole) *CreateRoleResponseBody {
	s.Role = v
	return s
}

func (s *CreateRoleResponseBody) Validate() error {
	if s.Role != nil {
		if err := s.Role.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRoleResponseBodyRole struct {
	// example:
	//
	// 2026-05-07T06:19:17Z
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// role_xxxxxxxxxxxxxxxxxxxx
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// example:
	//
	// Analyst
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// Manual
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2026-05-07T06:19:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateRoleResponseBodyRole) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleResponseBodyRole) GoString() string {
	return s.String()
}

func (s *CreateRoleResponseBodyRole) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateRoleResponseBodyRole) GetDescription() *string {
	return s.Description
}

func (s *CreateRoleResponseBodyRole) GetRoleId() *string {
	return s.RoleId
}

func (s *CreateRoleResponseBodyRole) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateRoleResponseBodyRole) GetType() *string {
	return s.Type
}

func (s *CreateRoleResponseBodyRole) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateRoleResponseBodyRole) SetCreateTime(v string) *CreateRoleResponseBodyRole {
	s.CreateTime = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetDescription(v string) *CreateRoleResponseBodyRole {
	s.Description = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetRoleId(v string) *CreateRoleResponseBodyRole {
	s.RoleId = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetRoleName(v string) *CreateRoleResponseBodyRole {
	s.RoleName = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetType(v string) *CreateRoleResponseBodyRole {
	s.Type = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetUpdateTime(v string) *CreateRoleResponseBodyRole {
	s.UpdateTime = &v
	return s
}

func (s *CreateRoleResponseBodyRole) Validate() error {
	return dara.Validate(s)
}
