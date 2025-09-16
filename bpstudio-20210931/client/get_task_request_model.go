// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetTaskRequest
	GetTaskId() *string
}

type GetTaskRequest struct {
	// example:
	//
	// 2615
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskRequest) SetTaskId(v string) *GetTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetTaskRequest) Validate() error {
	return dara.Validate(s)
}
