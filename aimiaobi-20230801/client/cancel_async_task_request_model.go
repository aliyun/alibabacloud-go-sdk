// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CancelAsyncTaskRequest
	GetAgentKey() *string
	SetTaskId(v string) *CancelAsyncTaskRequest
	GetTaskId() *string
}

type CancelAsyncTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelAsyncTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CancelAsyncTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelAsyncTaskRequest) SetAgentKey(v string) *CancelAsyncTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *CancelAsyncTaskRequest) SetTaskId(v string) *CancelAsyncTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CancelAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
