// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInnerConvertAsyncResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetInnerConvertAsyncResultRequest
	GetTaskId() *string
}

type GetInnerConvertAsyncResultRequest struct {
	// The task ID that uniquely identifies a task.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetInnerConvertAsyncResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInnerConvertAsyncResultRequest) GoString() string {
	return s.String()
}

func (s *GetInnerConvertAsyncResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetInnerConvertAsyncResultRequest) SetTaskId(v string) *GetInnerConvertAsyncResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetInnerConvertAsyncResultRequest) Validate() error {
	return dara.Validate(s)
}
