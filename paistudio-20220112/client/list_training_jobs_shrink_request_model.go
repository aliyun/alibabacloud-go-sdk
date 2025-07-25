// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmName(v string) *ListTrainingJobsShrinkRequest
	GetAlgorithmName() *string
	SetAlgorithmProvider(v string) *ListTrainingJobsShrinkRequest
	GetAlgorithmProvider() *string
	SetEndTime(v string) *ListTrainingJobsShrinkRequest
	GetEndTime() *string
	SetIsTempAlgo(v bool) *ListTrainingJobsShrinkRequest
	GetIsTempAlgo() *bool
	SetLabelsShrink(v string) *ListTrainingJobsShrinkRequest
	GetLabelsShrink() *string
	SetOrder(v string) *ListTrainingJobsShrinkRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListTrainingJobsShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListTrainingJobsShrinkRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListTrainingJobsShrinkRequest
	GetSortBy() *string
	SetStartTime(v string) *ListTrainingJobsShrinkRequest
	GetStartTime() *string
	SetStatus(v string) *ListTrainingJobsShrinkRequest
	GetStatus() *string
	SetTrainingJobId(v string) *ListTrainingJobsShrinkRequest
	GetTrainingJobId() *string
	SetTrainingJobName(v string) *ListTrainingJobsShrinkRequest
	GetTrainingJobName() *string
	SetWorkspaceId(v string) *ListTrainingJobsShrinkRequest
	GetWorkspaceId() *string
}

type ListTrainingJobsShrinkRequest struct {
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
	LabelsShrink *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
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

func (s ListTrainingJobsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsShrinkRequest) GetAlgorithmName() *string {
	return s.AlgorithmName
}

func (s *ListTrainingJobsShrinkRequest) GetAlgorithmProvider() *string {
	return s.AlgorithmProvider
}

func (s *ListTrainingJobsShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTrainingJobsShrinkRequest) GetIsTempAlgo() *bool {
	return s.IsTempAlgo
}

func (s *ListTrainingJobsShrinkRequest) GetLabelsShrink() *string {
	return s.LabelsShrink
}

func (s *ListTrainingJobsShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListTrainingJobsShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListTrainingJobsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTrainingJobsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTrainingJobsShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTrainingJobsShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTrainingJobsShrinkRequest) GetTrainingJobId() *string {
	return s.TrainingJobId
}

func (s *ListTrainingJobsShrinkRequest) GetTrainingJobName() *string {
	return s.TrainingJobName
}

func (s *ListTrainingJobsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListTrainingJobsShrinkRequest) SetAlgorithmName(v string) *ListTrainingJobsShrinkRequest {
	s.AlgorithmName = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetAlgorithmProvider(v string) *ListTrainingJobsShrinkRequest {
	s.AlgorithmProvider = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetEndTime(v string) *ListTrainingJobsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetIsTempAlgo(v bool) *ListTrainingJobsShrinkRequest {
	s.IsTempAlgo = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetLabelsShrink(v string) *ListTrainingJobsShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetOrder(v string) *ListTrainingJobsShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetPageNumber(v int64) *ListTrainingJobsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetPageSize(v int64) *ListTrainingJobsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetSortBy(v string) *ListTrainingJobsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetStartTime(v string) *ListTrainingJobsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetStatus(v string) *ListTrainingJobsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetTrainingJobId(v string) *ListTrainingJobsShrinkRequest {
	s.TrainingJobId = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetTrainingJobName(v string) *ListTrainingJobsShrinkRequest {
	s.TrainingJobName = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) SetWorkspaceId(v string) *ListTrainingJobsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListTrainingJobsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
