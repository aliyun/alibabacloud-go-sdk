// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *GetAsyncTaskRequest
	GetAsyncTaskId() *string
}

type GetAsyncTaskRequest struct {
	// The ID of the asynchronous task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 000e89fb-cf8f-11e9-8ab4-b6e980803a3b
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
}

func (s GetAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskRequest) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *GetAsyncTaskRequest) SetAsyncTaskId(v string) *GetAsyncTaskRequest {
	s.AsyncTaskId = &v
	return s
}

func (s *GetAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
