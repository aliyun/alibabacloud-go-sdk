// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluatorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *ListEvaluatorsRequest
	GetAgentSpace() *string
	SetMaxResults(v int32) *ListEvaluatorsRequest
	GetMaxResults() *int32
	SetName(v string) *ListEvaluatorsRequest
	GetName() *string
	SetNextToken(v string) *ListEvaluatorsRequest
	GetNextToken() *string
	SetSource(v string) *ListEvaluatorsRequest
	GetSource() *string
	SetType(v string) *ListEvaluatorsRequest
	GetType() *string
}

type ListEvaluatorsRequest struct {
	// The AgentSpace name.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-agentspace
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The fuzzy match condition for the evaluator name.
	//
	// example:
	//
	// trace_task_completion
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// eyJsYXN0SWQiOjEyM30=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The evaluator source filter.
	//
	// example:
	//
	// custom
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The evaluator type filter.
	//
	// example:
	//
	// AGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListEvaluatorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluatorsRequest) GoString() string {
	return s.String()
}

func (s *ListEvaluatorsRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *ListEvaluatorsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEvaluatorsRequest) GetName() *string {
	return s.Name
}

func (s *ListEvaluatorsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEvaluatorsRequest) GetSource() *string {
	return s.Source
}

func (s *ListEvaluatorsRequest) GetType() *string {
	return s.Type
}

func (s *ListEvaluatorsRequest) SetAgentSpace(v string) *ListEvaluatorsRequest {
	s.AgentSpace = &v
	return s
}

func (s *ListEvaluatorsRequest) SetMaxResults(v int32) *ListEvaluatorsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEvaluatorsRequest) SetName(v string) *ListEvaluatorsRequest {
	s.Name = &v
	return s
}

func (s *ListEvaluatorsRequest) SetNextToken(v string) *ListEvaluatorsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEvaluatorsRequest) SetSource(v string) *ListEvaluatorsRequest {
	s.Source = &v
	return s
}

func (s *ListEvaluatorsRequest) SetType(v string) *ListEvaluatorsRequest {
	s.Type = &v
	return s
}

func (s *ListEvaluatorsRequest) Validate() error {
	return dara.Validate(s)
}
