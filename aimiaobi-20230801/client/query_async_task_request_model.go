// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *QueryAsyncTaskRequest
	GetAgentKey() *string
	SetTaskId(v string) *QueryAsyncTaskRequest
	GetTaskId() *string
}

type QueryAsyncTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 33a2658aaabf4c24b45d50e575125311_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *QueryAsyncTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryAsyncTaskRequest) SetAgentKey(v string) *QueryAsyncTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *QueryAsyncTaskRequest) SetTaskId(v string) *QueryAsyncTaskRequest {
	s.TaskId = &v
	return s
}

func (s *QueryAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
