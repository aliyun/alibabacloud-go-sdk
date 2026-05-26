// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoleAssignmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRoleAssignmentResponseBody
	GetRequestId() *string
	SetRoleAssignment(v *CreateRoleAssignmentResponseBodyRoleAssignment) *CreateRoleAssignmentResponseBody
	GetRoleAssignment() *CreateRoleAssignmentResponseBodyRoleAssignment
}

type CreateRoleAssignmentResponseBody struct {
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleAssignment *CreateRoleAssignmentResponseBodyRoleAssignment `json:"RoleAssignment,omitempty" xml:"RoleAssignment,omitempty" type:"Struct"`
}

func (s CreateRoleAssignmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleAssignmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoleAssignmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRoleAssignmentResponseBody) GetRoleAssignment() *CreateRoleAssignmentResponseBodyRoleAssignment {
	return s.RoleAssignment
}

func (s *CreateRoleAssignmentResponseBody) SetRequestId(v string) *CreateRoleAssignmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoleAssignmentResponseBody) SetRoleAssignment(v *CreateRoleAssignmentResponseBodyRoleAssignment) *CreateRoleAssignmentResponseBody {
	s.RoleAssignment = v
	return s
}

func (s *CreateRoleAssignmentResponseBody) Validate() error {
	if s.RoleAssignment != nil {
		if err := s.RoleAssignment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRoleAssignmentResponseBodyRoleAssignment struct {
	// example:
	//
	// 2026-05-07T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// user_xxxxxxxxxxxxxxxxxxxx
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// example:
	//
	// alice
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// example:
	//
	// User
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
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
	// up_xxxxxxxxxxxxxxxxxxxx
	UserPoolId *string `json:"UserPoolId,omitempty" xml:"UserPoolId,omitempty"`
}

func (s CreateRoleAssignmentResponseBodyRoleAssignment) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleAssignmentResponseBodyRoleAssignment) GoString() string {
	return s.String()
}

func (s *CreateRoleAssignmentResponseBodyRoleAssignment) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateRoleAssignmentResponseBodyRoleAssignment) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *CreateRoleAssignmentResponseBodyRoleAssignment) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *CreateRoleAssignmentResponseBodyRoleAssignment) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *CreateRoleAssignmentResponseBodyRoleAssignment) GetRoleId() *string {
	return s.RoleId
}

func (s *CreateRoleAssignmentResponseBodyRoleAssignment) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateRoleAssignmentResponseBodyRoleAssignment) GetUserPoolId() *string {
	return s.UserPoolId
}

func (s *CreateRoleAssignmentResponseBodyRoleAssignment) SetCreateTime(v string) *CreateRoleAssignmentResponseBodyRoleAssignment {
	s.CreateTime = &v
	return s
}

func (s *CreateRoleAssignmentResponseBodyRoleAssignment) SetPrincipalId(v string) *CreateRoleAssignmentResponseBodyRoleAssignment {
	s.PrincipalId = &v
	return s
}

func (s *CreateRoleAssignmentResponseBodyRoleAssignment) SetPrincipalName(v string) *CreateRoleAssignmentResponseBodyRoleAssignment {
	s.PrincipalName = &v
	return s
}

func (s *CreateRoleAssignmentResponseBodyRoleAssignment) SetPrincipalType(v string) *CreateRoleAssignmentResponseBodyRoleAssignment {
	s.PrincipalType = &v
	return s
}

func (s *CreateRoleAssignmentResponseBodyRoleAssignment) SetRoleId(v string) *CreateRoleAssignmentResponseBodyRoleAssignment {
	s.RoleId = &v
	return s
}

func (s *CreateRoleAssignmentResponseBodyRoleAssignment) SetRoleName(v string) *CreateRoleAssignmentResponseBodyRoleAssignment {
	s.RoleName = &v
	return s
}

func (s *CreateRoleAssignmentResponseBodyRoleAssignment) SetUserPoolId(v string) *CreateRoleAssignmentResponseBodyRoleAssignment {
	s.UserPoolId = &v
	return s
}

func (s *CreateRoleAssignmentResponseBodyRoleAssignment) Validate() error {
	return dara.Validate(s)
}
