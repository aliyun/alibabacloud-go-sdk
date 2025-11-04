// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIWorkflowTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetAIWorkflowTaskRequest
	GetTaskId() *string
}

type GetAIWorkflowTaskRequest struct {
	// The ID of the workflow task.
	//
	// example:
	//
	// ********-266c-4bb8-b20c-6faa********
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAIWorkflowTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAIWorkflowTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAIWorkflowTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAIWorkflowTaskRequest) SetTaskId(v string) *GetAIWorkflowTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetAIWorkflowTaskRequest) Validate() error {
	return dara.Validate(s)
}
