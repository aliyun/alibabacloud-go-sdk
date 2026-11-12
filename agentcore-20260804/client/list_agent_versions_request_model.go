// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAgentVersionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgentVersionsRequest
	GetNextToken() *string
}

type ListAgentVersionsRequest struct {
	// The maximum number of records per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token returned in the previous ListAgentVersions response when you query the next page. Do not specify this parameter for the first query.
	//
	// example:
	//
	// YWdlbnQtdmVyc2lvbnM6MjA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAgentVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentVersionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentVersionsRequest) SetMaxResults(v int32) *ListAgentVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAgentVersionsRequest) SetNextToken(v string) *ListAgentVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAgentVersionsRequest) Validate() error {
	return dara.Validate(s)
}
