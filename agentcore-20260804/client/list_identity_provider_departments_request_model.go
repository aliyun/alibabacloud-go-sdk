// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProviderDepartmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListIdentityProviderDepartmentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIdentityProviderDepartmentsRequest
	GetNextToken() *string
	SetParentDepartmentId(v string) *ListIdentityProviderDepartmentsRequest
	GetParentDepartmentId() *string
}

type ListIdentityProviderDepartmentsRequest struct {
	// The number of departments to return per page. Valid values: 1 to 50. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The nextToken returned from the previous page. Do not specify this parameter for the first query. When continuing pagination, keep the identity provider and parentDepartmentId unchanged, and pass the token as is. An empty or missing return value indicates that there are no subsequent pages at the current hierarchy level.
	//
	// example:
	//
	// example-next-page-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Queries the direct child departments of the specified department. If this parameter is not specified or is empty, the root hierarchy is queried, and all descendants are not recursively returned. When expanding a node in the list, pass the ID of the node, not its parentId.
	//
	// example:
	//
	// 10001
	ParentDepartmentId *string `json:"parentDepartmentId,omitempty" xml:"parentDepartmentId,omitempty"`
}

func (s ListIdentityProviderDepartmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProviderDepartmentsRequest) GoString() string {
	return s.String()
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

func (s *ListIdentityProviderDepartmentsRequest) Validate() error {
	return dara.Validate(s)
}
