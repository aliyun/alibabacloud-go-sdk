// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrainTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTrainTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateTrainTaskResponseBody
	GetTaskId() *string
}

type CreateTrainTaskResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskId    *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateTrainTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrainTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTrainTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTrainTaskResponseBody) SetRequestId(v string) *CreateTrainTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrainTaskResponseBody) SetTaskId(v string) *CreateTrainTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateTrainTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
