// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRoleResponseBody
	GetRequestId() *string
	SetRole(v *GetRoleResponseBodyRole) *GetRoleResponseBody
	GetRole() *GetRoleResponseBodyRole
}

type GetRoleResponseBody struct {
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Role      *GetRoleResponseBodyRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
}

func (s GetRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRoleResponseBody) GetRole() *GetRoleResponseBodyRole {
	return s.Role
}

func (s *GetRoleResponseBody) SetRequestId(v string) *GetRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoleResponseBody) SetRole(v *GetRoleResponseBodyRole) *GetRoleResponseBody {
	s.Role = v
	return s
}

func (s *GetRoleResponseBody) Validate() error {
	if s.Role != nil {
		if err := s.Role.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRoleResponseBodyRole struct {
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

func (s GetRoleResponseBodyRole) String() string {
	return dara.Prettify(s)
}

func (s GetRoleResponseBodyRole) GoString() string {
	return s.String()
}

func (s *GetRoleResponseBodyRole) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRoleResponseBodyRole) GetDescription() *string {
	return s.Description
}

func (s *GetRoleResponseBodyRole) GetRoleId() *string {
	return s.RoleId
}

func (s *GetRoleResponseBodyRole) GetRoleName() *string {
	return s.RoleName
}

func (s *GetRoleResponseBodyRole) GetType() *string {
	return s.Type
}

func (s *GetRoleResponseBodyRole) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetRoleResponseBodyRole) SetCreateTime(v string) *GetRoleResponseBodyRole {
	s.CreateTime = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetDescription(v string) *GetRoleResponseBodyRole {
	s.Description = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetRoleId(v string) *GetRoleResponseBodyRole {
	s.RoleId = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetRoleName(v string) *GetRoleResponseBodyRole {
	s.RoleName = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetType(v string) *GetRoleResponseBodyRole {
	s.Type = &v
	return s
}

func (s *GetRoleResponseBodyRole) SetUpdateTime(v string) *GetRoleResponseBodyRole {
	s.UpdateTime = &v
	return s
}

func (s *GetRoleResponseBodyRole) Validate() error {
	return dara.Validate(s)
}
