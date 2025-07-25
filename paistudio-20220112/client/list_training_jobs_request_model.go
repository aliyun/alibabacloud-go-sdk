// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmName(v string) *ListTrainingJobsRequest
	GetAlgorithmName() *string
	SetAlgorithmProvider(v string) *ListTrainingJobsRequest
	GetAlgorithmProvider() *string
	SetEndTime(v string) *ListTrainingJobsRequest
	GetEndTime() *string
	SetIsTempAlgo(v bool) *ListTrainingJobsRequest
	GetIsTempAlgo() *bool
	SetLabels(v map[string]interface{}) *ListTrainingJobsRequest
	GetLabels() map[string]interface{}
	SetOrder(v string) *ListTrainingJobsRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListTrainingJobsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListTrainingJobsRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListTrainingJobsRequest
	GetSortBy() *string
	SetStartTime(v string) *ListTrainingJobsRequest
	GetStartTime() *string
	SetStatus(v string) *ListTrainingJobsRequest
	GetStatus() *string
	SetTrainingJobId(v string) *ListTrainingJobsRequest
	GetTrainingJobId() *string
	SetTrainingJobName(v string) *ListTrainingJobsRequest
	GetTrainingJobName() *string
	SetWorkspaceId(v string) *ListTrainingJobsRequest
	GetWorkspaceId() *string
}

type ListTrainingJobsRequest struct {
	// example:
	//
	// llm_train
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// example:
	//
	// 2023-12-27T02:10:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// false
	IsTempAlgo *bool `json:"IsTempAlgo,omitempty" xml:"IsTempAlgo,omitempty"`
	// example:
	//
	// {"project": "sd-s3"}
	Labels map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 2024-06-22T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// trains930928remn
	TrainingJobId *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
	// example:
	//
	// large_language_model_training
	TrainingJobName *string `json:"TrainingJobName,omitempty" xml:"TrainingJobName,omitempty"`
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTrainingJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobsRequest) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsRequest) GetAlgorithmName() *string {
	return s.AlgorithmName
}

func (s *ListTrainingJobsRequest) GetAlgorithmProvider() *string {
	return s.AlgorithmProvider
}

func (s *ListTrainingJobsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTrainingJobsRequest) GetIsTempAlgo() *bool {
	return s.IsTempAlgo
}

func (s *ListTrainingJobsRequest) GetLabels() map[string]interface{} {
	return s.Labels
}

func (s *ListTrainingJobsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListTrainingJobsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListTrainingJobsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTrainingJobsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTrainingJobsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTrainingJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTrainingJobsRequest) GetTrainingJobId() *string {
	return s.TrainingJobId
}

func (s *ListTrainingJobsRequest) GetTrainingJobName() *string {
	return s.TrainingJobName
}

func (s *ListTrainingJobsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListTrainingJobsRequest) SetAlgorithmName(v string) *ListTrainingJobsRequest {
	s.AlgorithmName = &v
	return s
}

func (s *ListTrainingJobsRequest) SetAlgorithmProvider(v string) *ListTrainingJobsRequest {
	s.AlgorithmProvider = &v
	return s
}

func (s *ListTrainingJobsRequest) SetEndTime(v string) *ListTrainingJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobsRequest) SetIsTempAlgo(v bool) *ListTrainingJobsRequest {
	s.IsTempAlgo = &v
	return s
}

func (s *ListTrainingJobsRequest) SetLabels(v map[string]interface{}) *ListTrainingJobsRequest {
	s.Labels = v
	return s
}

func (s *ListTrainingJobsRequest) SetOrder(v string) *ListTrainingJobsRequest {
	s.Order = &v
	return s
}

func (s *ListTrainingJobsRequest) SetPageNumber(v int64) *ListTrainingJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrainingJobsRequest) SetPageSize(v int64) *ListTrainingJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrainingJobsRequest) SetSortBy(v string) *ListTrainingJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListTrainingJobsRequest) SetStartTime(v string) *ListTrainingJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTrainingJobsRequest) SetStatus(v string) *ListTrainingJobsRequest {
	s.Status = &v
	return s
}

func (s *ListTrainingJobsRequest) SetTrainingJobId(v string) *ListTrainingJobsRequest {
	s.TrainingJobId = &v
	return s
}

func (s *ListTrainingJobsRequest) SetTrainingJobName(v string) *ListTrainingJobsRequest {
	s.TrainingJobName = &v
	return s
}

func (s *ListTrainingJobsRequest) SetWorkspaceId(v string) *ListTrainingJobsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListTrainingJobsRequest) Validate() error {
	return dara.Validate(s)
}
