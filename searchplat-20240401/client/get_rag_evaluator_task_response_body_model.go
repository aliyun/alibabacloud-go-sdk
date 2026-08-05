// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRagEvaluatorTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRagEvaluatorTaskResponseBody
	GetCode() *string
	SetMessage(v string) *GetRagEvaluatorTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRagEvaluatorTaskResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *GetRagEvaluatorTaskResponseBody
	GetResult() map[string]interface{}
}

type GetRagEvaluatorTaskResponseBody struct {
	// The status code.
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
	// The response result.
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
}

func (s GetRagEvaluatorTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRagEvaluatorTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetRagEvaluatorTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRagEvaluatorTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRagEvaluatorTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRagEvaluatorTaskResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *GetRagEvaluatorTaskResponseBody) SetCode(v string) *GetRagEvaluatorTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetRagEvaluatorTaskResponseBody) SetMessage(v string) *GetRagEvaluatorTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetRagEvaluatorTaskResponseBody) SetRequestId(v string) *GetRagEvaluatorTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRagEvaluatorTaskResponseBody) SetResult(v map[string]interface{}) *GetRagEvaluatorTaskResponseBody {
	s.Result = v
	return s
}

func (s *GetRagEvaluatorTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
