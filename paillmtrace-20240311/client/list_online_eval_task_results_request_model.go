// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOnlineEvalTaskResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluationId(v string) *ListOnlineEvalTaskResultsRequest
	GetEvaluationId() *string
	SetMostRecentResultsOnly(v bool) *ListOnlineEvalTaskResultsRequest
	GetMostRecentResultsOnly() *bool
	SetPageNumber(v int32) *ListOnlineEvalTaskResultsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOnlineEvalTaskResultsRequest
	GetPageSize() *int32
	SetTraceIds(v []*string) *ListOnlineEvalTaskResultsRequest
	GetTraceIds() []*string
}

type ListOnlineEvalTaskResultsRequest struct {
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
	TraceIds []*string `json:"TraceIds,omitempty" xml:"TraceIds,omitempty" type:"Repeated"`
}

func (s ListOnlineEvalTaskResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOnlineEvalTaskResultsRequest) GoString() string {
	return s.String()
}

func (s *ListOnlineEvalTaskResultsRequest) GetEvaluationId() *string {
	return s.EvaluationId
}

func (s *ListOnlineEvalTaskResultsRequest) GetMostRecentResultsOnly() *bool {
	return s.MostRecentResultsOnly
}

func (s *ListOnlineEvalTaskResultsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOnlineEvalTaskResultsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOnlineEvalTaskResultsRequest) GetTraceIds() []*string {
	return s.TraceIds
}

func (s *ListOnlineEvalTaskResultsRequest) SetEvaluationId(v string) *ListOnlineEvalTaskResultsRequest {
	s.EvaluationId = &v
	return s
}

func (s *ListOnlineEvalTaskResultsRequest) SetMostRecentResultsOnly(v bool) *ListOnlineEvalTaskResultsRequest {
	s.MostRecentResultsOnly = &v
	return s
}

func (s *ListOnlineEvalTaskResultsRequest) SetPageNumber(v int32) *ListOnlineEvalTaskResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOnlineEvalTaskResultsRequest) SetPageSize(v int32) *ListOnlineEvalTaskResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOnlineEvalTaskResultsRequest) SetTraceIds(v []*string) *ListOnlineEvalTaskResultsRequest {
	s.TraceIds = v
	return s
}

func (s *ListOnlineEvalTaskResultsRequest) Validate() error {
	return dara.Validate(s)
}
