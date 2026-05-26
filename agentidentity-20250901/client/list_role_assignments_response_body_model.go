// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoleAssignmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssignments(v []*ListRoleAssignmentsResponseBodyAssignments) *ListRoleAssignmentsResponseBody
	GetAssignments() []*ListRoleAssignmentsResponseBodyAssignments
	SetMaxResults(v int32) *ListRoleAssignmentsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListRoleAssignmentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRoleAssignmentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListRoleAssignmentsResponseBody
	GetTotalCount() *int32
}

type ListRoleAssignmentsResponseBody struct {
	Assignments []*ListRoleAssignmentsResponseBodyAssignments `json:"Assignments,omitempty" xml:"Assignments,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// TGlzdEFzc2lnbm1lbnRzOjoyMA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRoleAssignmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRoleAssignmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoleAssignmentsResponseBody) GetAssignments() []*ListRoleAssignmentsResponseBodyAssignments {
	return s.Assignments
}

func (s *ListRoleAssignmentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRoleAssignmentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRoleAssignmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRoleAssignmentsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRoleAssignmentsResponseBody) SetAssignments(v []*ListRoleAssignmentsResponseBodyAssignments) *ListRoleAssignmentsResponseBody {
	s.Assignments = v
	return s
}

func (s *ListRoleAssignmentsResponseBody) SetMaxResults(v int32) *ListRoleAssignmentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRoleAssignmentsResponseBody) SetNextToken(v string) *ListRoleAssignmentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRoleAssignmentsResponseBody) SetRequestId(v string) *ListRoleAssignmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoleAssignmentsResponseBody) SetTotalCount(v int32) *ListRoleAssignmentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRoleAssignmentsResponseBody) Validate() error {
	if s.Assignments != nil {
		for _, item := range s.Assignments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRoleAssignmentsResponseBodyAssignments struct {
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

func (s ListRoleAssignmentsResponseBodyAssignments) String() string {
	return dara.Prettify(s)
}

func (s ListRoleAssignmentsResponseBodyAssignments) GoString() string {
	return s.String()
}

func (s *ListRoleAssignmentsResponseBodyAssignments) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRoleAssignmentsResponseBodyAssignments) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListRoleAssignmentsResponseBodyAssignments) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *ListRoleAssignmentsResponseBodyAssignments) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListRoleAssignmentsResponseBodyAssignments) GetRoleId() *string {
	return s.RoleId
}

func (s *ListRoleAssignmentsResponseBodyAssignments) GetRoleName() *string {
	return s.RoleName
}

func (s *ListRoleAssignmentsResponseBodyAssignments) GetUserPoolId() *string {
	return s.UserPoolId
}

func (s *ListRoleAssignmentsResponseBodyAssignments) SetCreateTime(v string) *ListRoleAssignmentsResponseBodyAssignments {
	s.CreateTime = &v
	return s
}

func (s *ListRoleAssignmentsResponseBodyAssignments) SetPrincipalId(v string) *ListRoleAssignmentsResponseBodyAssignments {
	s.PrincipalId = &v
	return s
}

func (s *ListRoleAssignmentsResponseBodyAssignments) SetPrincipalName(v string) *ListRoleAssignmentsResponseBodyAssignments {
	s.PrincipalName = &v
	return s
}

func (s *ListRoleAssignmentsResponseBodyAssignments) SetPrincipalType(v string) *ListRoleAssignmentsResponseBodyAssignments {
	s.PrincipalType = &v
	return s
}

func (s *ListRoleAssignmentsResponseBodyAssignments) SetRoleId(v string) *ListRoleAssignmentsResponseBodyAssignments {
	s.RoleId = &v
	return s
}

func (s *ListRoleAssignmentsResponseBodyAssignments) SetRoleName(v string) *ListRoleAssignmentsResponseBodyAssignments {
	s.RoleName = &v
	return s
}

func (s *ListRoleAssignmentsResponseBodyAssignments) SetUserPoolId(v string) *ListRoleAssignmentsResponseBodyAssignments {
	s.UserPoolId = &v
	return s
}

func (s *ListRoleAssignmentsResponseBodyAssignments) Validate() error {
	return dara.Validate(s)
}
