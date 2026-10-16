// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProviderDepartmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderType(v string) *ListIdentityProviderDepartmentsRequest
	GetIdentityProviderType() *string
	SetMaxResults(v int32) *ListIdentityProviderDepartmentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIdentityProviderDepartmentsRequest
	GetNextToken() *string
	SetParentDepartmentId(v string) *ListIdentityProviderDepartmentsRequest
	GetParentDepartmentId() *string
	SetUserPoolName(v string) *ListIdentityProviderDepartmentsRequest
	GetUserPoolName() *string
}

type ListIdentityProviderDepartmentsRequest struct {
	// example:
	//
	// DingTalk
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	// example:
	//
	// 2
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// eyJwYWdlIjoyfQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	ParentDepartmentId *string `json:"ParentDepartmentId,omitempty" xml:"ParentDepartmentId,omitempty"`
	// example:
	//
	// syncscope-example
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s ListIdentityProviderDepartmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProviderDepartmentsRequest) GoString() string {
	return s.String()
}

func (s *ListIdentityProviderDepartmentsRequest) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *ListIdentityProviderDepartmentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIdentityProviderDepartmentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIdentityProviderDepartmentsRequest) GetParentDepartmentId() *string {
	return s.ParentDepartmentId
}

func (s *ListIdentityProviderDepartmentsRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *ListIdentityProviderDepartmentsRequest) SetIdentityProviderType(v string) *ListIdentityProviderDepartmentsRequest {
	s.IdentityProviderType = &v
	return s
}

func (s *ListIdentityProviderDepartmentsRequest) SetMaxResults(v int32) *ListIdentityProviderDepartmentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIdentityProviderDepartmentsRequest) SetNextToken(v string) *ListIdentityProviderDepartmentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIdentityProviderDepartmentsRequest) SetParentDepartmentId(v string) *ListIdentityProviderDepartmentsRequest {
	s.ParentDepartmentId = &v
	return s
}

func (s *ListIdentityProviderDepartmentsRequest) SetUserPoolName(v string) *ListIdentityProviderDepartmentsRequest {
	s.UserPoolName = &v
	return s
}

func (s *ListIdentityProviderDepartmentsRequest) Validate() error {
	return dara.Validate(s)
}
