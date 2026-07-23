// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentRunsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *ListExperimentRunsRequest
	GetDatasetId() *string
	SetExperimentName(v string) *ListExperimentRunsRequest
	GetExperimentName() *string
	SetMaxResults(v int32) *ListExperimentRunsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListExperimentRunsRequest
	GetNextToken() *string
	SetPage(v int32) *ListExperimentRunsRequest
	GetPage() *int32
	SetPageSize(v int32) *ListExperimentRunsRequest
	GetPageSize() *int32
	SetPlanName(v string) *ListExperimentRunsRequest
	GetPlanName() *string
	SetStatus(v string) *ListExperimentRunsRequest
	GetStatus() *string
}

type ListExperimentRunsRequest struct {
	// Filters results by exact dataset ID.
	//
	// example:
	//
	// arms_customer_agent_level1
	DatasetId *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	// Filters results by fuzzy match on the experiment configuration name.
	//
	// example:
	//
	// experimentA
	ExperimentName *string `json:"experimentName,omitempty" xml:"experimentName,omitempty"`
	// Optional. Use `page` and `pageSize` for pagination instead.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Optional. Use `page` and `pageSize` for pagination instead.
	//
	// example:
	//
	// eyJwYWdlIjoxfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number, starting from 0. Default value: 0.
	//
	// example:
	//
	// 0
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Filters results by fuzzy match on experiment plan name.
	//
	// example:
	//
	// arms_agent_experiment
	PlanName *string `json:"planName,omitempty" xml:"planName,omitempty"`
	// Filters results by status.
	//
	// example:
	//
	// evaluating
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListExperimentRunsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentRunsRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentRunsRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *ListExperimentRunsRequest) GetExperimentName() *string {
	return s.ExperimentName
}

func (s *ListExperimentRunsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExperimentRunsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExperimentRunsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListExperimentRunsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExperimentRunsRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *ListExperimentRunsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListExperimentRunsRequest) SetDatasetId(v string) *ListExperimentRunsRequest {
	s.DatasetId = &v
	return s
}

func (s *ListExperimentRunsRequest) SetExperimentName(v string) *ListExperimentRunsRequest {
	s.ExperimentName = &v
	return s
}

func (s *ListExperimentRunsRequest) SetMaxResults(v int32) *ListExperimentRunsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExperimentRunsRequest) SetNextToken(v string) *ListExperimentRunsRequest {
	s.NextToken = &v
	return s
}

func (s *ListExperimentRunsRequest) SetPage(v int32) *ListExperimentRunsRequest {
	s.Page = &v
	return s
}

func (s *ListExperimentRunsRequest) SetPageSize(v int32) *ListExperimentRunsRequest {
	s.PageSize = &v
	return s
}

func (s *ListExperimentRunsRequest) SetPlanName(v string) *ListExperimentRunsRequest {
	s.PlanName = &v
	return s
}

func (s *ListExperimentRunsRequest) SetStatus(v string) *ListExperimentRunsRequest {
	s.Status = &v
	return s
}

func (s *ListExperimentRunsRequest) Validate() error {
	return dara.Validate(s)
}
