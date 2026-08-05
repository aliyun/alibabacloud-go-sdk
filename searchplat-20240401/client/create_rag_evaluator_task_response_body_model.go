// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRagEvaluatorTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRagEvaluatorTaskResponseBody
	GetCode() *string
	SetMessage(v string) *CreateRagEvaluatorTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRagEvaluatorTaskResponseBody
	GetRequestId() *string
	SetResult(v *CreateRagEvaluatorTaskResponseBodyResult) *CreateRagEvaluatorTaskResponseBody
	GetResult() *CreateRagEvaluatorTaskResponseBodyResult
}

type CreateRagEvaluatorTaskResponseBody struct {
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
	// 0abb793917165176014887584e28d9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result *CreateRagEvaluatorTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateRagEvaluatorTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRagEvaluatorTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRagEvaluatorTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRagEvaluatorTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRagEvaluatorTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRagEvaluatorTaskResponseBody) GetResult() *CreateRagEvaluatorTaskResponseBodyResult {
	return s.Result
}

func (s *CreateRagEvaluatorTaskResponseBody) SetCode(v string) *CreateRagEvaluatorTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRagEvaluatorTaskResponseBody) SetMessage(v string) *CreateRagEvaluatorTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRagEvaluatorTaskResponseBody) SetRequestId(v string) *CreateRagEvaluatorTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRagEvaluatorTaskResponseBody) SetResult(v *CreateRagEvaluatorTaskResponseBodyResult) *CreateRagEvaluatorTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateRagEvaluatorTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRagEvaluatorTaskResponseBodyResult struct {
	// The task ID.
	//
	// example:
	//
	// 1846389386674049024
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateRagEvaluatorTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateRagEvaluatorTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRagEvaluatorTaskResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateRagEvaluatorTaskResponseBodyResult) SetTaskId(v string) *CreateRagEvaluatorTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *CreateRagEvaluatorTaskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
