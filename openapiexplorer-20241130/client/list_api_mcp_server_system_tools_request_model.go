// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiMcpServerSystemToolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListApiMcpServerSystemToolsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListApiMcpServerSystemToolsRequest
	GetNextToken() *string
	SetSkip(v int32) *ListApiMcpServerSystemToolsRequest
	GetSkip() *int32
}

type ListApiMcpServerSystemToolsRequest struct {
	// The maximum number of entries to return on each page. The maximum value is 100. The default value is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token to start the next query. Set this parameter to the NextToken value returned by the last API call.
	//
	// > Do not specify this parameter for the first query. If a query does not return all results, pass the NextToken value from the previous response to the next query to continue.
	//
	// example:
	//
	// AAAAAZjtYxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The number of entries to skip.
	//
	// example:
	//
	// 5
	Skip *int32 `json:"skip,omitempty" xml:"skip,omitempty"`
}

func (s ListApiMcpServerSystemToolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServerSystemToolsRequest) GoString() string {
	return s.String()
}

func (s *ListApiMcpServerSystemToolsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApiMcpServerSystemToolsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApiMcpServerSystemToolsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListApiMcpServerSystemToolsRequest) SetMaxResults(v int32) *ListApiMcpServerSystemToolsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApiMcpServerSystemToolsRequest) SetNextToken(v string) *ListApiMcpServerSystemToolsRequest {
	s.NextToken = &v
	return s
}

func (s *ListApiMcpServerSystemToolsRequest) SetSkip(v int32) *ListApiMcpServerSystemToolsRequest {
	s.Skip = &v
	return s
}

func (s *ListApiMcpServerSystemToolsRequest) Validate() error {
	return dara.Validate(s)
}
