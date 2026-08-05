// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRagEvaluatorTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRagEvaluatorTasksResponseBody
	GetCode() *string
	SetMessage(v string) *ListRagEvaluatorTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRagEvaluatorTasksResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ListRagEvaluatorTasksResponseBody
	GetResult() map[string]interface{}
	SetTotalCount(v int32) *ListRagEvaluatorTasksResponseBody
	GetTotalCount() *int32
}

type ListRagEvaluatorTasksResponseBody struct {
	// The error code.
	//
	// example:
	//
	// not found
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message.
	//
	// example:
	//
	// "xx not found"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1CC93E65-6734-5060-BEF7-0EB0A4862BCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// {
	//
	//     "odps_task_id": 224525243,
	//
	//   "usage" : {
	//
	//     "cu" : 0.000
	//
	//   }
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListRagEvaluatorTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRagEvaluatorTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListRagEvaluatorTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRagEvaluatorTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRagEvaluatorTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRagEvaluatorTasksResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ListRagEvaluatorTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRagEvaluatorTasksResponseBody) SetCode(v string) *ListRagEvaluatorTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListRagEvaluatorTasksResponseBody) SetMessage(v string) *ListRagEvaluatorTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListRagEvaluatorTasksResponseBody) SetRequestId(v string) *ListRagEvaluatorTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRagEvaluatorTasksResponseBody) SetResult(v map[string]interface{}) *ListRagEvaluatorTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListRagEvaluatorTasksResponseBody) SetTotalCount(v int32) *ListRagEvaluatorTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRagEvaluatorTasksResponseBody) Validate() error {
	return dara.Validate(s)
}
