// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOnlineEvalTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateOnlineEvalTaskResponseBody
	GetCode() *string
	SetMessage(v string) *CreateOnlineEvalTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateOnlineEvalTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateOnlineEvalTaskResponseBody
	GetTaskId() *string
}

type CreateOnlineEvalTaskResponseBody struct {
	// The internal error code. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// EvaluationConfig.Answer.SpanName is required.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the created trace evaluation task.
	//
	// example:
	//
	// 711ef9112343286810abbfce04e161ee
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateOnlineEvalTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOnlineEvalTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOnlineEvalTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateOnlineEvalTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateOnlineEvalTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOnlineEvalTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateOnlineEvalTaskResponseBody) SetCode(v string) *CreateOnlineEvalTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOnlineEvalTaskResponseBody) SetMessage(v string) *CreateOnlineEvalTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOnlineEvalTaskResponseBody) SetRequestId(v string) *CreateOnlineEvalTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOnlineEvalTaskResponseBody) SetTaskId(v string) *CreateOnlineEvalTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateOnlineEvalTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
