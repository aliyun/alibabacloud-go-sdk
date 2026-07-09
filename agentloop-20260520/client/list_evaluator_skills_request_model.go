// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluatorSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *ListEvaluatorSkillsRequest
	GetAgentSpace() *string
	SetMaxResults(v int32) *ListEvaluatorSkillsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEvaluatorSkillsRequest
	GetNextToken() *string
}

type ListEvaluatorSkillsRequest struct {
	// The AgentSpace name.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-agentspace
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// eyJuZXh0IjoiMjAifQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListEvaluatorSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluatorSkillsRequest) GoString() string {
	return s.String()
}

func (s *ListEvaluatorSkillsRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *ListEvaluatorSkillsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEvaluatorSkillsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEvaluatorSkillsRequest) SetAgentSpace(v string) *ListEvaluatorSkillsRequest {
	s.AgentSpace = &v
	return s
}

func (s *ListEvaluatorSkillsRequest) SetMaxResults(v int32) *ListEvaluatorSkillsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEvaluatorSkillsRequest) SetNextToken(v string) *ListEvaluatorSkillsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEvaluatorSkillsRequest) Validate() error {
	return dara.Validate(s)
}
