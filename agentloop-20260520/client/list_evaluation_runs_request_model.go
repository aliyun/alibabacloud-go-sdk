// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationRunsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListEvaluationRunsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEvaluationRunsRequest
	GetNextToken() *string
	SetRunType(v string) *ListEvaluationRunsRequest
	GetRunType() *string
	SetStatus(v string) *ListEvaluationRunsRequest
	GetStatus() *string
}

type ListEvaluationRunsRequest struct {
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// eyJsYXN0SWQiOjEwMX0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The run type filter condition.
	//
	// example:
	//
	// backfill
	RunType *string `json:"runType,omitempty" xml:"runType,omitempty"`
	// The run status filter condition.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListEvaluationRunsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationRunsRequest) GoString() string {
	return s.String()
}

func (s *ListEvaluationRunsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEvaluationRunsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEvaluationRunsRequest) GetRunType() *string {
	return s.RunType
}

func (s *ListEvaluationRunsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListEvaluationRunsRequest) SetMaxResults(v int32) *ListEvaluationRunsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEvaluationRunsRequest) SetNextToken(v string) *ListEvaluationRunsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEvaluationRunsRequest) SetRunType(v string) *ListEvaluationRunsRequest {
	s.RunType = &v
	return s
}

func (s *ListEvaluationRunsRequest) SetStatus(v string) *ListEvaluationRunsRequest {
	s.Status = &v
	return s
}

func (s *ListEvaluationRunsRequest) Validate() error {
	return dara.Validate(s)
}
