// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvalResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListEvalResultsResponseBody
	GetCode() *string
	SetEvaluationResults(v []*string) *ListEvalResultsResponseBody
	GetEvaluationResults() []*string
	SetMessage(v string) *ListEvalResultsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEvalResultsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEvalResultsResponseBody
	GetTotalCount() *int32
}

type ListEvalResultsResponseBody struct {
	// The internal error code. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// ExecutionFailure
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The evaluation results.
	EvaluationResults []*string `json:"EvaluationResults,omitempty" xml:"EvaluationResults,omitempty" type:"Repeated"`
	// The error message. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// cannot get data back.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the POP request.
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of results that meet the condition.
	//
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEvalResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEvalResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEvalResultsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListEvalResultsResponseBody) GetEvaluationResults() []*string {
	return s.EvaluationResults
}

func (s *ListEvalResultsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEvalResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEvalResultsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEvalResultsResponseBody) SetCode(v string) *ListEvalResultsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEvalResultsResponseBody) SetEvaluationResults(v []*string) *ListEvalResultsResponseBody {
	s.EvaluationResults = v
	return s
}

func (s *ListEvalResultsResponseBody) SetMessage(v string) *ListEvalResultsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEvalResultsResponseBody) SetRequestId(v string) *ListEvalResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEvalResultsResponseBody) SetTotalCount(v int32) *ListEvalResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEvalResultsResponseBody) Validate() error {
	return dara.Validate(s)
}
