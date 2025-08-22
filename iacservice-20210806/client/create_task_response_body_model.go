// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateTaskResponseBody
	GetTaskId() *string
}

type CreateTaskResponseBody struct {
	// example:
	//
	// CD478792-6952-5A1C-9F57-78932BF0FAC6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// task-433aead756057fffeaba4828f5195
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTaskResponseBody) SetRequestId(v string) *CreateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskResponseBody) SetTaskId(v string) *CreateTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
