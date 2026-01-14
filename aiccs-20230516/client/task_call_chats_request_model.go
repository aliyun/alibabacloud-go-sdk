// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskCallChatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *TaskCallChatsRequest
	GetAgentId() *int64
	SetAgentTag(v string) *TaskCallChatsRequest
	GetAgentTag() *string
	SetCallId(v string) *TaskCallChatsRequest
	GetCallId() *string
	SetOwnerId(v int64) *TaskCallChatsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TaskCallChatsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TaskCallChatsRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *TaskCallChatsRequest
	GetTaskId() *int64
}

type TaskCallChatsRequest struct {
	// 坐席ID
	//
	// example:
	//
	// 72
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// AA
	AgentTag *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	// 外呼ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 9b2eb6b8-7a27-4357-b5ec-104450086e24
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 任务ID
	//
	// example:
	//
	// 26
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskCallChatsRequest) String() string {
	return dara.Prettify(s)
}

func (s TaskCallChatsRequest) GoString() string {
	return s.String()
}

func (s *TaskCallChatsRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *TaskCallChatsRequest) GetAgentTag() *string {
	return s.AgentTag
}

func (s *TaskCallChatsRequest) GetCallId() *string {
	return s.CallId
}

func (s *TaskCallChatsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TaskCallChatsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TaskCallChatsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TaskCallChatsRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *TaskCallChatsRequest) SetAgentId(v int64) *TaskCallChatsRequest {
	s.AgentId = &v
	return s
}

func (s *TaskCallChatsRequest) SetAgentTag(v string) *TaskCallChatsRequest {
	s.AgentTag = &v
	return s
}

func (s *TaskCallChatsRequest) SetCallId(v string) *TaskCallChatsRequest {
	s.CallId = &v
	return s
}

func (s *TaskCallChatsRequest) SetOwnerId(v int64) *TaskCallChatsRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskCallChatsRequest) SetResourceOwnerAccount(v string) *TaskCallChatsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskCallChatsRequest) SetResourceOwnerId(v int64) *TaskCallChatsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskCallChatsRequest) SetTaskId(v int64) *TaskCallChatsRequest {
	s.TaskId = &v
	return s
}

func (s *TaskCallChatsRequest) Validate() error {
	return dara.Validate(s)
}
