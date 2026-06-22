// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryAgentlessTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *RetryAgentlessTaskRequest
	GetTaskId() *string
}

type RetryAgentlessTaskRequest struct {
	// The task ID. Call [ListAgentlessTask](~~ListAgentlessTask~~) to obtain the task ID.
	//
	// example:
	//
	// 5347c7b6-c85c-4070-846a-3029e08e****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RetryAgentlessTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryAgentlessTaskRequest) GoString() string {
	return s.String()
}

func (s *RetryAgentlessTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RetryAgentlessTaskRequest) SetTaskId(v string) *RetryAgentlessTaskRequest {
	s.TaskId = &v
	return s
}

func (s *RetryAgentlessTaskRequest) Validate() error {
	return dara.Validate(s)
}
