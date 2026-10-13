// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetAsyncResultRequest
	GetAgentKey() *string
	SetTaskId(v string) *GetAsyncResultRequest
	GetTaskId() *string
}

type GetAsyncResultRequest struct {
	// The workspace key. If not specified, the default workspace is accessed. Obtain the key value from the business management page of the Alibaba Cloud account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The task ID of the asynchronous service.
	//
	// This parameter is required.
	//
	// example:
	//
	// Obtain TASK_ID from the Chat operation
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAsyncResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncResultRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetAsyncResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAsyncResultRequest) SetAgentKey(v string) *GetAsyncResultRequest {
	s.AgentKey = &v
	return s
}

func (s *GetAsyncResultRequest) SetTaskId(v string) *GetAsyncResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetAsyncResultRequest) Validate() error {
	return dara.Validate(s)
}
