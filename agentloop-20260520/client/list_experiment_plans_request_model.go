// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentPlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListExperimentPlansRequest
	GetLimit() *int32
	SetMaxResults(v int32) *ListExperimentPlansRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListExperimentPlansRequest
	GetNextToken() *string
	SetOffset(v int32) *ListExperimentPlansRequest
	GetOffset() *int32
	SetPlanName(v string) *ListExperimentPlansRequest
	GetPlanName() *string
	SetStatus(v string) *ListExperimentPlansRequest
	GetStatus() *string
}

type ListExperimentPlansRequest struct {
	// The number of entries to return. Default value: 20.
	//
	// example:
	//
	// 20
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// Optional. Use `offset` and `limit` for pagination instead.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Optional. Use `offset` and `limit` for pagination instead.
	//
	// example:
	//
	// eyJsYXN0SWQiOjEyM30=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The offset. Default value: 0.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// Fuzzy match by plan name.
	//
	// example:
	//
	// arms_agent
	PlanName *string `json:"planName,omitempty" xml:"planName,omitempty"`
	// Filters by exact status.
	//
	// example:
	//
	// pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListExperimentPlansRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentPlansRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentPlansRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListExperimentPlansRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExperimentPlansRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExperimentPlansRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *ListExperimentPlansRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *ListExperimentPlansRequest) GetStatus() *string {
	return s.Status
}

func (s *ListExperimentPlansRequest) SetLimit(v int32) *ListExperimentPlansRequest {
	s.Limit = &v
	return s
}

func (s *ListExperimentPlansRequest) SetMaxResults(v int32) *ListExperimentPlansRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExperimentPlansRequest) SetNextToken(v string) *ListExperimentPlansRequest {
	s.NextToken = &v
	return s
}

func (s *ListExperimentPlansRequest) SetOffset(v int32) *ListExperimentPlansRequest {
	s.Offset = &v
	return s
}

func (s *ListExperimentPlansRequest) SetPlanName(v string) *ListExperimentPlansRequest {
	s.PlanName = &v
	return s
}

func (s *ListExperimentPlansRequest) SetStatus(v string) *ListExperimentPlansRequest {
	s.Status = &v
	return s
}

func (s *ListExperimentPlansRequest) Validate() error {
	return dara.Validate(s)
}
