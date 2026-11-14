// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarmPoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWarmPoolsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWarmPoolsRequest
	GetNextToken() *string
}

type ListWarmPoolsRequest struct {
	// The actual maximum number of results returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. An empty value indicates that no more results are available.
	//
	// example:
	//
	// example-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListWarmPoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWarmPoolsRequest) GoString() string {
	return s.String()
}

func (s *ListWarmPoolsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWarmPoolsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWarmPoolsRequest) SetMaxResults(v int32) *ListWarmPoolsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWarmPoolsRequest) SetNextToken(v string) *ListWarmPoolsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWarmPoolsRequest) Validate() error {
	return dara.Validate(s)
}
