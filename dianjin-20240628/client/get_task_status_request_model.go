// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetTaskStatusRequest
	GetTaskId() *string
}

type GetTaskStatusRequest struct {
	// This parameter is required.
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskStatusRequest) SetTaskId(v string) *GetTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
