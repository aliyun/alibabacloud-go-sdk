// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRolesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListRolesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRolesResponseBody
	GetRequestId() *string
	SetRoles(v []*ListRolesResponseBodyRoles) *ListRolesResponseBody
	GetRoles() []*ListRolesResponseBodyRoles
	SetTotalCount(v int32) *ListRolesResponseBody
	GetTotalCount() *int32
}

type ListRolesResponseBody struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// TGlzdFJvbGVzOjoyMA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Roles     []*ListRolesResponseBodyRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// example:
	//
	// 8
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRolesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRolesResponseBody) GetRoles() []*ListRolesResponseBodyRoles {
	return s.Roles
}

func (s *ListRolesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRolesResponseBody) SetMaxResults(v int32) *ListRolesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRolesResponseBody) SetNextToken(v string) *ListRolesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRolesResponseBody) SetRequestId(v string) *ListRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRolesResponseBody) SetRoles(v []*ListRolesResponseBodyRoles) *ListRolesResponseBody {
	s.Roles = v
	return s
}

func (s *ListRolesResponseBody) SetTotalCount(v int32) *ListRolesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRolesResponseBody) Validate() error {
	if s.Roles != nil {
		for _, item := range s.Roles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRolesResponseBodyRoles struct {
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

func (s ListRolesResponseBodyRoles) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyRoles) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyRoles) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRolesResponseBodyRoles) GetDescription() *string {
	return s.Description
}

func (s *ListRolesResponseBodyRoles) GetRoleId() *string {
	return s.RoleId
}

func (s *ListRolesResponseBodyRoles) GetRoleName() *string {
	return s.RoleName
}

func (s *ListRolesResponseBodyRoles) GetType() *string {
	return s.Type
}

func (s *ListRolesResponseBodyRoles) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListRolesResponseBodyRoles) SetCreateTime(v string) *ListRolesResponseBodyRoles {
	s.CreateTime = &v
	return s
}

func (s *ListRolesResponseBodyRoles) SetDescription(v string) *ListRolesResponseBodyRoles {
	s.Description = &v
	return s
}

func (s *ListRolesResponseBodyRoles) SetRoleId(v string) *ListRolesResponseBodyRoles {
	s.RoleId = &v
	return s
}

func (s *ListRolesResponseBodyRoles) SetRoleName(v string) *ListRolesResponseBodyRoles {
	s.RoleName = &v
	return s
}

func (s *ListRolesResponseBodyRoles) SetType(v string) *ListRolesResponseBodyRoles {
	s.Type = &v
	return s
}

func (s *ListRolesResponseBodyRoles) SetUpdateTime(v string) *ListRolesResponseBodyRoles {
	s.UpdateTime = &v
	return s
}

func (s *ListRolesResponseBodyRoles) Validate() error {
	return dara.Validate(s)
}
