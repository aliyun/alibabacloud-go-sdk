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
}

type ListAgentSpacesRequest struct {
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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

func (s *ListAgentSpacesRequest) Validate() error {
	return dara.Validate(s)
}
