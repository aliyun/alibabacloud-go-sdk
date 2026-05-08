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
	// example:
	//
	// 84657DE0-B68C-508B-AFE7-8ED921854E3C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 837091359375048704
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
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
