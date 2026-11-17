// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListCustomDomainsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCustomDomainsRequest
	GetNextToken() *string
}

type ListCustomDomainsRequest struct {
	// The maximum number of records to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page, used for subsequent queries.
	//
	// example:
	//
	// example-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListCustomDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomDomainsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCustomDomainsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomDomainsRequest) SetMaxResults(v int32) *ListCustomDomainsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCustomDomainsRequest) SetNextToken(v string) *ListCustomDomainsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCustomDomainsRequest) Validate() error {
	return dara.Validate(s)
}
