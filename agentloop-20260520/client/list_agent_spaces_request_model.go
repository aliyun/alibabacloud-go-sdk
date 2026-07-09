// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSpacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *ListAgentSpacesRequest
	GetAgentSpace() *string
	SetMaxResults(v int32) *ListAgentSpacesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgentSpacesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListAgentSpacesRequest
	GetRegionId() *string
}

type ListAgentSpacesRequest struct {
	// The AgentSpace name.
	//
	// example:
	//
	// test-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The maximum number of results to return.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// pEL20OGYeZQez8NdW7ve
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RegionId  *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListAgentSpacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSpacesRequest) GoString() string {
	return s.String()
}

func (s *ListAgentSpacesRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *ListAgentSpacesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentSpacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentSpacesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAgentSpacesRequest) SetAgentSpace(v string) *ListAgentSpacesRequest {
	s.AgentSpace = &v
	return s
}

func (s *ListAgentSpacesRequest) SetMaxResults(v int32) *ListAgentSpacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAgentSpacesRequest) SetNextToken(v string) *ListAgentSpacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAgentSpacesRequest) SetRegionId(v string) *ListAgentSpacesRequest {
	s.RegionId = &v
	return s
}

func (s *ListAgentSpacesRequest) Validate() error {
	return dara.Validate(s)
}
