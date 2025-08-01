// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvalResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluationId(v string) *ListEvalResultsRequest
	GetEvaluationId() *string
	SetKeyword(v string) *ListEvalResultsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListEvalResultsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEvalResultsRequest
	GetPageSize() *int32
	SetRecordIds(v []*string) *ListEvalResultsRequest
	GetRecordIds() []*string
}

type ListEvalResultsRequest struct {
	// The task ID of the evaluation task to which the trace belongs.
	//
	// example:
	//
	// 0bb05ae2a2dc11ef9757faaa2a1ec0c6
	EvaluationId *string `json:"EvaluationId,omitempty" xml:"EvaluationId,omitempty"`
	// The keyword to query from the evaluation inputs.
	//
	// example:
	//
	// foo
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Page starts from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The trace data IDs.
	RecordIds []*string `json:"RecordIds,omitempty" xml:"RecordIds,omitempty" type:"Repeated"`
}

func (s ListEvalResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEvalResultsRequest) GoString() string {
	return s.String()
}

func (s *ListEvalResultsRequest) GetEvaluationId() *string {
	return s.EvaluationId
}

func (s *ListEvalResultsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListEvalResultsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEvalResultsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEvalResultsRequest) GetRecordIds() []*string {
	return s.RecordIds
}

func (s *ListEvalResultsRequest) SetEvaluationId(v string) *ListEvalResultsRequest {
	s.EvaluationId = &v
	return s
}

func (s *ListEvalResultsRequest) SetKeyword(v string) *ListEvalResultsRequest {
	s.Keyword = &v
	return s
}

func (s *ListEvalResultsRequest) SetPageNumber(v int32) *ListEvalResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEvalResultsRequest) SetPageSize(v int32) *ListEvalResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEvalResultsRequest) SetRecordIds(v []*string) *ListEvalResultsRequest {
	s.RecordIds = v
	return s
}

func (s *ListEvalResultsRequest) Validate() error {
	return dara.Validate(s)
}
