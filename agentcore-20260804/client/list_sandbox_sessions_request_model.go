// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSandboxSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSandboxSessionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSandboxSessionsRequest
	GetNextToken() *string
}

type ListSandboxSessionsRequest struct {
	// The maximum number of entries to return per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page. Omit this parameter in the first request. For subsequent requests, use the returned token as is, and keep the workspace, agent, instance, and maxResults unchanged. An empty value indicates that no more pages are available.
	//
	// example:
	//
	// next-token-1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListSandboxSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSandboxSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListSandboxSessionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSandboxSessionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSandboxSessionsRequest) SetMaxResults(v int32) *ListSandboxSessionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSandboxSessionsRequest) SetNextToken(v string) *ListSandboxSessionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSandboxSessionsRequest) Validate() error {
	return dara.Validate(s)
}
