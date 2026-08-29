// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMcpsRequest
	GetMaxResults() *int32
	SetName(v string) *ListMcpsRequest
	GetName() *string
	SetNextToken(v string) *ListMcpsRequest
	GetNextToken() *string
	SetSearchType(v string) *ListMcpsRequest
	GetSearchType() *string
}

type ListMcpsRequest struct {
	// The maximum number of entries to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// my-mcp-server
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// next-page-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// blur
	SearchType *string `json:"searchType,omitempty" xml:"searchType,omitempty"`
}

func (s ListMcpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsRequest) GoString() string {
	return s.String()
}

func (s *ListMcpsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpsRequest) GetName() *string {
	return s.Name
}

func (s *ListMcpsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpsRequest) GetSearchType() *string {
	return s.SearchType
}

func (s *ListMcpsRequest) SetMaxResults(v int32) *ListMcpsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMcpsRequest) SetName(v string) *ListMcpsRequest {
	s.Name = &v
	return s
}

func (s *ListMcpsRequest) SetNextToken(v string) *ListMcpsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMcpsRequest) SetSearchType(v string) *ListMcpsRequest {
	s.SearchType = &v
	return s
}

func (s *ListMcpsRequest) Validate() error {
	return dara.Validate(s)
}
