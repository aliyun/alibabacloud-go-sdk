// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoleAssignmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrincipalName(v string) *DeleteRoleAssignmentRequest
	GetPrincipalName() *string
	SetPrincipalType(v string) *DeleteRoleAssignmentRequest
	GetPrincipalType() *string
	SetRoleName(v string) *DeleteRoleAssignmentRequest
	GetRoleName() *string
	SetUserPoolName(v string) *DeleteRoleAssignmentRequest
	GetUserPoolName() *string
}

type DeleteRoleAssignmentRequest struct {
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

func (s DeleteRoleAssignmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoleAssignmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoleAssignmentRequest) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *DeleteRoleAssignmentRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *DeleteRoleAssignmentRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *DeleteRoleAssignmentRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *DeleteRoleAssignmentRequest) SetPrincipalName(v string) *DeleteRoleAssignmentRequest {
	s.PrincipalName = &v
	return s
}

func (s *DeleteRoleAssignmentRequest) SetPrincipalType(v string) *DeleteRoleAssignmentRequest {
	s.PrincipalType = &v
	return s
}

func (s *DeleteRoleAssignmentRequest) SetRoleName(v string) *DeleteRoleAssignmentRequest {
	s.RoleName = &v
	return s
}

func (s *DeleteRoleAssignmentRequest) SetUserPoolName(v string) *DeleteRoleAssignmentRequest {
	s.UserPoolName = &v
	return s
}

func (s *DeleteRoleAssignmentRequest) Validate() error {
	return dara.Validate(s)
}
