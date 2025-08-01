// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOnlineEvalTaskResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListOnlineEvalTaskResultsResponseBody
	GetCode() *string
	SetEvaluationResults(v []*string) *ListOnlineEvalTaskResultsResponseBody
	GetEvaluationResults() []*string
	SetMessage(v string) *ListOnlineEvalTaskResultsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListOnlineEvalTaskResultsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListOnlineEvalTaskResultsResponseBody
	GetTotalCount() *int32
}

type ListOnlineEvalTaskResultsResponseBody struct {
	// Internal error code. Set only when the response has an error.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of evaluation results.
	EvaluationResults []*string `json:"EvaluationResults,omitempty" xml:"EvaluationResults,omitempty" type:"Repeated"`
	// Response error message. Set only when the response has an error.
	//
	// example:
	//
	// must provide trace_id(s) or eval_id
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 22BA9A5A-E3D8-5B4C-90FC-F33F6E5853F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of evaluation results that meet the criteria.
	//
	// example:
	//
	// 123
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOnlineEvalTaskResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOnlineEvalTaskResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOnlineEvalTaskResultsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListOnlineEvalTaskResultsResponseBody) GetEvaluationResults() []*string {
	return s.EvaluationResults
}

func (s *ListOnlineEvalTaskResultsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListOnlineEvalTaskResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOnlineEvalTaskResultsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOnlineEvalTaskResultsResponseBody) SetCode(v string) *ListOnlineEvalTaskResultsResponseBody {
	s.Code = &v
	return s
}

func (s *ListOnlineEvalTaskResultsResponseBody) SetEvaluationResults(v []*string) *ListOnlineEvalTaskResultsResponseBody {
	s.EvaluationResults = v
	return s
}

func (s *ListOnlineEvalTaskResultsResponseBody) SetMessage(v string) *ListOnlineEvalTaskResultsResponseBody {
	s.Message = &v
	return s
}

func (s *ListOnlineEvalTaskResultsResponseBody) SetRequestId(v string) *ListOnlineEvalTaskResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOnlineEvalTaskResultsResponseBody) SetTotalCount(v int32) *ListOnlineEvalTaskResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOnlineEvalTaskResultsResponseBody) Validate() error {
	return dara.Validate(s)
}
