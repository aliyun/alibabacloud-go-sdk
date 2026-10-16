// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoleAssignmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRoleAssignmentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRoleAssignmentsRequest
	GetNextToken() *string
	SetPrincipalName(v string) *ListRoleAssignmentsRequest
	GetPrincipalName() *string
	SetPrincipalType(v string) *ListRoleAssignmentsRequest
	GetPrincipalType() *string
	SetRoleName(v string) *ListRoleAssignmentsRequest
	GetRoleName() *string
	SetUserPoolName(v string) *ListRoleAssignmentsRequest
	GetUserPoolName() *string
}

type ListRoleAssignmentsRequest struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// TGlzdEFzc2lnbm1lbnRzOjoxMA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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

func (s ListRoleAssignmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRoleAssignmentsRequest) GoString() string {
	return s.String()
}

func (s *ListRoleAssignmentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRoleAssignmentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRoleAssignmentsRequest) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *ListRoleAssignmentsRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListRoleAssignmentsRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *ListRoleAssignmentsRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *ListRoleAssignmentsRequest) SetMaxResults(v int32) *ListRoleAssignmentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRoleAssignmentsRequest) SetNextToken(v string) *ListRoleAssignmentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRoleAssignmentsRequest) SetPrincipalName(v string) *ListRoleAssignmentsRequest {
	s.PrincipalName = &v
	return s
}

func (s *ListRoleAssignmentsRequest) SetPrincipalType(v string) *ListRoleAssignmentsRequest {
	s.PrincipalType = &v
	return s
}

func (s *ListRoleAssignmentsRequest) SetRoleName(v string) *ListRoleAssignmentsRequest {
	s.RoleName = &v
	return s
}

func (s *ListRoleAssignmentsRequest) SetUserPoolName(v string) *ListRoleAssignmentsRequest {
	s.UserPoolName = &v
	return s
}

func (s *ListRoleAssignmentsRequest) Validate() error {
	return dara.Validate(s)
}
