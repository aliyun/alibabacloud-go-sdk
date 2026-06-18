// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMemoryNodesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMemoryNodesRequest
	GetNextToken() *string
}

type ListMemoryNodesRequest struct {
	// The maximum number of results to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token returned from a previous request. If specified, the query results are paginated based on the previous request.
	//
	// example:
	//
	// dc270401186b433f975d7e1faaa34e0e
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListMemoryNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryNodesRequest) GoString() string {
	return s.String()
}

func (s *ListMemoryNodesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemoryNodesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemoryNodesRequest) SetMaxResults(v int32) *ListMemoryNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMemoryNodesRequest) SetNextToken(v string) *ListMemoryNodesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMemoryNodesRequest) Validate() error {
	return dara.Validate(s)
}
