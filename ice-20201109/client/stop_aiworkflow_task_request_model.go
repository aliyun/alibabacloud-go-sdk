// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAIWorkflowTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *StopAIWorkflowTaskRequest
	GetTaskId() *string
}

type StopAIWorkflowTaskRequest struct {
	// The ID of the workflow task.
	//
	// example:
	//
	// ********-266c-4bb8-b20c-6faa********
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopAIWorkflowTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAIWorkflowTaskRequest) GoString() string {
	return s.String()
}

func (s *StopAIWorkflowTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopAIWorkflowTaskRequest) SetTaskId(v string) *StopAIWorkflowTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StopAIWorkflowTaskRequest) Validate() error {
	return dara.Validate(s)
}
