// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTuringTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *CheckTuringTaskRequest
	GetTaskId() *string
}

type CheckTuringTaskRequest struct {
	// example:
	//
	// 871509423398305792
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CheckTuringTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckTuringTaskRequest) GoString() string {
	return s.String()
}

func (s *CheckTuringTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CheckTuringTaskRequest) SetTaskId(v string) *CheckTuringTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CheckTuringTaskRequest) Validate() error {
	return dara.Validate(s)
}
