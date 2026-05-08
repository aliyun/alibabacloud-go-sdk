// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitProjectTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitProjectTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *SubmitProjectTaskResponseBody
	GetTaskId() *string
}

type SubmitProjectTaskResponseBody struct {
	// example:
	//
	// 551FF252-6CFC-5DDA-9F84-9B07302385C2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 313123123
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitProjectTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitProjectTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitProjectTaskResponseBody) SetRequestId(v string) *SubmitProjectTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitProjectTaskResponseBody) SetTaskId(v string) *SubmitProjectTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *SubmitProjectTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
