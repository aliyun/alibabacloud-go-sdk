// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvalResultsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluationId(v string) *ListEvalResultsShrinkRequest
	GetEvaluationId() *string
	SetKeyword(v string) *ListEvalResultsShrinkRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListEvalResultsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEvalResultsShrinkRequest
	GetPageSize() *int32
	SetRecordIdsShrink(v string) *ListEvalResultsShrinkRequest
	GetRecordIdsShrink() *string
}

type ListEvalResultsShrinkRequest struct {
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
	RecordIdsShrink *string `json:"RecordIds,omitempty" xml:"RecordIds,omitempty"`
}

func (s ListEvalResultsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEvalResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListEvalResultsShrinkRequest) GetEvaluationId() *string {
	return s.EvaluationId
}

func (s *ListEvalResultsShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListEvalResultsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEvalResultsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEvalResultsShrinkRequest) GetRecordIdsShrink() *string {
	return s.RecordIdsShrink
}

func (s *ListEvalResultsShrinkRequest) SetEvaluationId(v string) *ListEvalResultsShrinkRequest {
	s.EvaluationId = &v
	return s
}

func (s *ListEvalResultsShrinkRequest) SetKeyword(v string) *ListEvalResultsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListEvalResultsShrinkRequest) SetPageNumber(v int32) *ListEvalResultsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEvalResultsShrinkRequest) SetPageSize(v int32) *ListEvalResultsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListEvalResultsShrinkRequest) SetRecordIdsShrink(v string) *ListEvalResultsShrinkRequest {
	s.RecordIdsShrink = &v
	return s
}

func (s *ListEvalResultsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
