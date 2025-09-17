// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListProductsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProductsRequest
	GetNextToken() *string
}

type ListProductsRequest struct {
	// The number of entries per page. Maximum value: 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// ECS
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductsRequest) SetMaxResults(v int32) *ListProductsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductsRequest) SetNextToken(v string) *ListProductsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProductsRequest) Validate() error {
	return dara.Validate(s)
}
