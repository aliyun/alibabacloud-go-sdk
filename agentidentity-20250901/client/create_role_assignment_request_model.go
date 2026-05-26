// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoleAssignmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrincipalName(v string) *CreateRoleAssignmentRequest
	GetPrincipalName() *string
	SetPrincipalType(v string) *CreateRoleAssignmentRequest
	GetPrincipalType() *string
	SetRoleName(v string) *CreateRoleAssignmentRequest
	GetRoleName() *string
	SetUserPoolName(v string) *CreateRoleAssignmentRequest
	GetUserPoolName() *string
}

type CreateRoleAssignmentRequest struct {
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
	// Analyst
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s CreateRoleAssignmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleAssignmentRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleAssignmentRequest) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *CreateRoleAssignmentRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *CreateRoleAssignmentRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateRoleAssignmentRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *CreateRoleAssignmentRequest) SetPrincipalName(v string) *CreateRoleAssignmentRequest {
	s.PrincipalName = &v
	return s
}

func (s *CreateRoleAssignmentRequest) SetPrincipalType(v string) *CreateRoleAssignmentRequest {
	s.PrincipalType = &v
	return s
}

func (s *CreateRoleAssignmentRequest) SetRoleName(v string) *CreateRoleAssignmentRequest {
	s.RoleName = &v
	return s
}

func (s *CreateRoleAssignmentRequest) SetUserPoolName(v string) *CreateRoleAssignmentRequest {
	s.UserPoolName = &v
	return s
}

func (s *CreateRoleAssignmentRequest) Validate() error {
	return dara.Validate(s)
}
