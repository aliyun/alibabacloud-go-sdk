// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpToolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMcpToolsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMcpToolsRequest
	GetNextToken() *string
}

type ListMcpToolsRequest struct {
	// The maximum number of results to return per request. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 6
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// 2
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListMcpToolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcpToolsRequest) GoString() string {
	return s.String()
}

func (s *ListMcpToolsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpToolsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpToolsRequest) SetMaxResults(v int32) *ListMcpToolsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMcpToolsRequest) SetNextToken(v string) *ListMcpToolsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMcpToolsRequest) Validate() error {
	return dara.Validate(s)
}
