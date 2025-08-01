// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetAgentTaskRequest
	GetTaskId() *string
}

type GetAgentTaskRequest struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetAgentTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAgentTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAgentTaskRequest) SetTaskId(v string) *GetAgentTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetAgentTaskRequest) Validate() error {
	return dara.Validate(s)
}
