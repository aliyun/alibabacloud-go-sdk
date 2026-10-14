// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOptimizeTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListOptimizeTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListOptimizeTasksRequest
	GetNextToken() *string
	SetStatus(v string) *ListOptimizeTasksRequest
	GetStatus() *string
	SetType(v string) *ListOptimizeTasksRequest
	GetType() *string
}

type ListOptimizeTasksRequest struct {
	// The maximum number of results to return. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// MTIzNDU2Nzg5MA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The task status.
	//
	// example:
	//
	// Enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The task type.
	//
	// example:
	//
	// EfficiencyAnalysis
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListOptimizeTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOptimizeTasksRequest) GoString() string {
	return s.String()
}

func (s *ListOptimizeTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOptimizeTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOptimizeTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListOptimizeTasksRequest) GetType() *string {
	return s.Type
}

func (s *ListOptimizeTasksRequest) SetMaxResults(v int32) *ListOptimizeTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOptimizeTasksRequest) SetNextToken(v string) *ListOptimizeTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListOptimizeTasksRequest) SetStatus(v string) *ListOptimizeTasksRequest {
	s.Status = &v
	return s
}

func (s *ListOptimizeTasksRequest) SetType(v string) *ListOptimizeTasksRequest {
	s.Type = &v
	return s
}

func (s *ListOptimizeTasksRequest) Validate() error {
	return dara.Validate(s)
}
