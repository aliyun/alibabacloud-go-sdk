// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMcpServersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMcpServersRequest
	GetNextToken() *string
	SetQ(v string) *ListMcpServersRequest
	GetQ() *string
	SetVisibility(v []*string) *ListMcpServersRequest
	GetVisibility() []*string
}

type ListMcpServersRequest struct {
	// The maximum number of results to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The next page token from a previous response. Use this token to retrieve the next page of results. Leave this parameter empty for the first request.
	//
	// example:
	//
	// 12345
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search keyword for a fuzzy search on MCP Server names.
	//
	// example:
	//
	// mcp
	Q *string `json:"Q,omitempty" xml:"Q,omitempty"`
	// The visibility level for filtering the results.
	//
	// example:
	//
	// -
	Visibility []*string `json:"Visibility,omitempty" xml:"Visibility,omitempty" type:"Repeated"`
}

func (s ListMcpServersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServersRequest) GoString() string {
	return s.String()
}

func (s *ListMcpServersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpServersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpServersRequest) GetQ() *string {
	return s.Q
}

func (s *ListMcpServersRequest) GetVisibility() []*string {
	return s.Visibility
}

func (s *ListMcpServersRequest) SetMaxResults(v int32) *ListMcpServersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMcpServersRequest) SetNextToken(v string) *ListMcpServersRequest {
	s.NextToken = &v
	return s
}

func (s *ListMcpServersRequest) SetQ(v string) *ListMcpServersRequest {
	s.Q = &v
	return s
}

func (s *ListMcpServersRequest) SetVisibility(v []*string) *ListMcpServersRequest {
	s.Visibility = v
	return s
}

func (s *ListMcpServersRequest) Validate() error {
	return dara.Validate(s)
}
