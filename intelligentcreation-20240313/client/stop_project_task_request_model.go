// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopProjectTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *StopProjectTaskRequest
	GetTaskId() *string
}

type StopProjectTaskRequest struct {
	// example:
	//
	// 1111111
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s StopProjectTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *StopProjectTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopProjectTaskRequest) SetTaskId(v string) *StopProjectTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StopProjectTaskRequest) Validate() error {
	return dara.Validate(s)
}
