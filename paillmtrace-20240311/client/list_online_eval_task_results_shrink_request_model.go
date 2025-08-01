// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOnlineEvalTaskResultsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluationId(v string) *ListOnlineEvalTaskResultsShrinkRequest
	GetEvaluationId() *string
	SetMostRecentResultsOnly(v bool) *ListOnlineEvalTaskResultsShrinkRequest
	GetMostRecentResultsOnly() *bool
	SetPageNumber(v int32) *ListOnlineEvalTaskResultsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOnlineEvalTaskResultsShrinkRequest
	GetPageSize() *int32
	SetTraceIdsShrink(v string) *ListOnlineEvalTaskResultsShrinkRequest
	GetTraceIdsShrink() *string
}

type ListOnlineEvalTaskResultsShrinkRequest struct {
	// The ID of the evaluation task. At least one of the trace ID or task ID must be set.
	//
	// example:
	//
	// 0bb05ae8888c11ef9757faaa2a1ec0c6
	EvaluationId *string `json:"EvaluationId,omitempty" xml:"EvaluationId,omitempty"`
	// The same trace data may have been evaluated by different tasks. If no task ID is specified and there are multiple evaluation results for the same trace ID, this parameter specifies whether to return only the most recent evaluation result.
	//
	// example:
	//
	// True
	MostRecentResultsOnly *bool `json:"MostRecentResultsOnly,omitempty" xml:"MostRecentResultsOnly,omitempty"`
	// The current page number. Value range: integers greater than 0. Default value: 1.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, default is 10.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specify a set of trace IDs, and only return the evaluation results for these traces. At least one of the trace ID or task ID must be set.
	TraceIdsShrink *string `json:"TraceIds,omitempty" xml:"TraceIds,omitempty"`
}

func (s ListOnlineEvalTaskResultsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOnlineEvalTaskResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) GetEvaluationId() *string {
	return s.EvaluationId
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) GetMostRecentResultsOnly() *bool {
	return s.MostRecentResultsOnly
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) GetTraceIdsShrink() *string {
	return s.TraceIdsShrink
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) SetEvaluationId(v string) *ListOnlineEvalTaskResultsShrinkRequest {
	s.EvaluationId = &v
	return s
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) SetMostRecentResultsOnly(v bool) *ListOnlineEvalTaskResultsShrinkRequest {
	s.MostRecentResultsOnly = &v
	return s
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) SetPageNumber(v int32) *ListOnlineEvalTaskResultsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) SetPageSize(v int32) *ListOnlineEvalTaskResultsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) SetTraceIdsShrink(v string) *ListOnlineEvalTaskResultsShrinkRequest {
	s.TraceIdsShrink = &v
	return s
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
