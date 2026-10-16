// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProviderDepartmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDepartments(v []*ListIdentityProviderDepartmentsResponseBodyDepartments) *ListIdentityProviderDepartmentsResponseBody
	GetDepartments() []*ListIdentityProviderDepartmentsResponseBodyDepartments
	SetNextToken(v string) *ListIdentityProviderDepartmentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIdentityProviderDepartmentsResponseBody
	GetRequestId() *string
}

type ListIdentityProviderDepartmentsResponseBody struct {
	Departments []*ListIdentityProviderDepartmentsResponseBodyDepartments `json:"Departments,omitempty" xml:"Departments,omitempty" type:"Repeated"`
	// example:
	//
	// eyJwYWdlIjoyfQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// EXAMPLE-REQUEST-ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIdentityProviderDepartmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProviderDepartmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIdentityProviderDepartmentsResponseBody) GetDepartments() []*ListIdentityProviderDepartmentsResponseBodyDepartments {
	return s.Departments
}

func (s *ListIdentityProviderDepartmentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIdentityProviderDepartmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIdentityProviderDepartmentsResponseBody) SetDepartments(v []*ListIdentityProviderDepartmentsResponseBodyDepartments) *ListIdentityProviderDepartmentsResponseBody {
	s.Departments = v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBody) SetNextToken(v string) *ListIdentityProviderDepartmentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBody) SetRequestId(v string) *ListIdentityProviderDepartmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBody) Validate() error {
	if s.Departments != nil {
		for _, item := range s.Departments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIdentityProviderDepartmentsResponseBodyDepartments struct {
	// example:
	//
	// true
	HasChildren *bool `json:"HasChildren,omitempty" xml:"HasChildren,omitempty"`
	// example:
	//
	// 100
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 12
	MemberCount *int64 `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	// example:
	//
	// 研发部
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s ListIdentityProviderDepartmentsResponseBodyDepartments) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProviderDepartmentsResponseBodyDepartments) GoString() string {
	return s.String()
}

func (s *ListIdentityProviderDepartmentsResponseBodyDepartments) GetHasChildren() *bool {
	return s.HasChildren
}

func (s *ListIdentityProviderDepartmentsResponseBodyDepartments) GetId() *string {
	return s.Id
}

func (s *ListIdentityProviderDepartmentsResponseBodyDepartments) GetMemberCount() *int64 {
	return s.MemberCount
}

func (s *ListIdentityProviderDepartmentsResponseBodyDepartments) GetName() *string {
	return s.Name
}

func (s *ListIdentityProviderDepartmentsResponseBodyDepartments) GetParentId() *string {
	return s.ParentId
}

func (s *ListIdentityProviderDepartmentsResponseBodyDepartments) SetHasChildren(v bool) *ListIdentityProviderDepartmentsResponseBodyDepartments {
	s.HasChildren = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBodyDepartments) SetId(v string) *ListIdentityProviderDepartmentsResponseBodyDepartments {
	s.Id = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBodyDepartments) SetMemberCount(v int64) *ListIdentityProviderDepartmentsResponseBodyDepartments {
	s.MemberCount = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBodyDepartments) SetName(v string) *ListIdentityProviderDepartmentsResponseBodyDepartments {
	s.Name = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBodyDepartments) SetParentId(v string) *ListIdentityProviderDepartmentsResponseBodyDepartments {
	s.ParentId = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBodyDepartments) Validate() error {
	return dara.Validate(s)
}
