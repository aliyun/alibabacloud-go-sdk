// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *CancelAsyncTaskRequest
	GetTaskId() *string
}

type CancelAsyncTaskRequest struct {
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CancelAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelAsyncTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelAsyncTaskRequest) SetTaskId(v string) *CancelAsyncTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CancelAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
