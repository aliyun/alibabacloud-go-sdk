// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentCodeRuntimesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAgentCodeRuntimesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgentCodeRuntimesRequest
	GetNextToken() *string
}

type ListAgentCodeRuntimesRequest struct {
	// The maximum number of records per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page.
	//
	// example:
	//
	// d3MtYTFiMmMzZDRlNWY2CmNvZGUtcnVudGltZS1leGFtcGxlCjE
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAgentCodeRuntimesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentCodeRuntimesRequest) GoString() string {
	return s.String()
}

func (s *ListAgentCodeRuntimesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentCodeRuntimesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentCodeRuntimesRequest) SetMaxResults(v int32) *ListAgentCodeRuntimesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAgentCodeRuntimesRequest) SetNextToken(v string) *ListAgentCodeRuntimesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAgentCodeRuntimesRequest) Validate() error {
	return dara.Validate(s)
}
