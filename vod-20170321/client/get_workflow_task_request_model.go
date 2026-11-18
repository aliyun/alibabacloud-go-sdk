// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetWorkflowTaskRequest
	GetTaskId() *string
}

type GetWorkflowTaskRequest struct {
	// The workflow task ID. It is returned by the SubmitWorkflowJob operation and is used to query the execution status and result of the task.
	//
	// example:
	//
	// 70422****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetWorkflowTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowTaskRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetWorkflowTaskRequest) SetTaskId(v string) *GetWorkflowTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetWorkflowTaskRequest) Validate() error {
	return dara.Validate(s)
}
