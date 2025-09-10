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
	// The maximum number of entries to return.
	//
	// Valid values: 1 to 200. Default value: 30.
	//
	// example:
	//
	// 4
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
