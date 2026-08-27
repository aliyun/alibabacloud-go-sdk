// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateScheduledTaskResponseBody
	GetCode() *string
	SetMessage(v string) *CreateScheduledTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateScheduledTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateScheduledTaskResponseBody
	GetTaskId() *string
}

type CreateScheduledTaskResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// exampleTaskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateScheduledTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateScheduledTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateScheduledTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScheduledTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateScheduledTaskResponseBody) SetCode(v string) *CreateScheduledTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetMessage(v string) *CreateScheduledTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetRequestId(v string) *CreateScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetTaskId(v string) *CreateScheduledTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
