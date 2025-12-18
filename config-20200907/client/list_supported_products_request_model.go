// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportedProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSupportedProductsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSupportedProductsRequest
	GetNextToken() *string
}

type ListSupportedProductsRequest struct {
	// The maximum number of entries to return in a request.
	//
	// Valid values: 1 to 500. Default value: 200.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// D3AjqMNSy0ls7zBNCf3a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListSupportedProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSupportedProductsRequest) GoString() string {
	return s.String()
}

func (s *ListSupportedProductsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSupportedProductsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupportedProductsRequest) SetMaxResults(v int32) *ListSupportedProductsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSupportedProductsRequest) SetNextToken(v string) *ListSupportedProductsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSupportedProductsRequest) Validate() error {
	return dara.Validate(s)
}
