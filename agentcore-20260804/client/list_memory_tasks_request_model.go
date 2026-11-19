// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMemoryTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMemoryTasksRequest
	GetNextToken() *string
	SetRuleId(v string) *ListMemoryTasksRequest
	GetRuleId() *string
	SetStatus(v string) *ListMemoryTasksRequest
	GetStatus() *string
}

type ListMemoryTasksRequest struct {
	// The maximum number of results per page.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token returned for the same resource and filter conditions.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwMH0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// ruleId
	//
	// This parameter is required.
	//
	// example:
	//
	// rule-2f5f9c5a8d0e4c3b
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// status
	//
	// example:
	//
	// SUCCEEDED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListMemoryTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryTasksRequest) GoString() string {
	return s.String()
}

func (s *ListMemoryTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemoryTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemoryTasksRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *ListMemoryTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListMemoryTasksRequest) SetMaxResults(v int32) *ListMemoryTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMemoryTasksRequest) SetNextToken(v string) *ListMemoryTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListMemoryTasksRequest) SetRuleId(v string) *ListMemoryTasksRequest {
	s.RuleId = &v
	return s
}

func (s *ListMemoryTasksRequest) SetStatus(v string) *ListMemoryTasksRequest {
	s.Status = &v
	return s
}

func (s *ListMemoryTasksRequest) Validate() error {
	return dara.Validate(s)
}
