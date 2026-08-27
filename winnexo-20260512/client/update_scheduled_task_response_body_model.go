// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduledTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateScheduledTaskResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateScheduledTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateScheduledTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpdateScheduledTaskResponseBody
	GetTaskId() *string
	SetUpdated(v bool) *UpdateScheduledTaskResponseBody
	GetUpdated() *bool
}

type UpdateScheduledTaskResponseBody struct {
	// The business status code. A value of 200 indicates success. A failure returns a backend error code (ERR.	- or InvalidParameter.*).
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description. Empty when the request is successful.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The task ID (echoed back).
	//
	// example:
	//
	// exampleTaskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// Indicates whether an actual update was made.
	//
	// example:
	//
	// true
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateScheduledTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduledTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateScheduledTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateScheduledTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateScheduledTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateScheduledTaskResponseBody) GetUpdated() *bool {
	return s.Updated
}

func (s *UpdateScheduledTaskResponseBody) SetCode(v string) *UpdateScheduledTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateScheduledTaskResponseBody) SetMessage(v string) *UpdateScheduledTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateScheduledTaskResponseBody) SetRequestId(v string) *UpdateScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScheduledTaskResponseBody) SetTaskId(v string) *UpdateScheduledTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateScheduledTaskResponseBody) SetUpdated(v bool) *UpdateScheduledTaskResponseBody {
	s.Updated = &v
	return s
}

func (s *UpdateScheduledTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
