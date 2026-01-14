// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallChatListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *CallChatListRequest
	GetCallId() *string
	SetOwnerId(v int64) *CallChatListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CallChatListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CallChatListRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *CallChatListRequest
	GetTaskId() *int64
}

type CallChatListRequest struct {
	// callId
	//
	// example:
	//
	// a
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 任务ID
	//
	// example:
	//
	// 1
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CallChatListRequest) String() string {
	return dara.Prettify(s)
}

func (s CallChatListRequest) GoString() string {
	return s.String()
}

func (s *CallChatListRequest) GetCallId() *string {
	return s.CallId
}

func (s *CallChatListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CallChatListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CallChatListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CallChatListRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CallChatListRequest) SetCallId(v string) *CallChatListRequest {
	s.CallId = &v
	return s
}

func (s *CallChatListRequest) SetOwnerId(v int64) *CallChatListRequest {
	s.OwnerId = &v
	return s
}

func (s *CallChatListRequest) SetResourceOwnerAccount(v string) *CallChatListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CallChatListRequest) SetResourceOwnerId(v int64) *CallChatListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CallChatListRequest) SetTaskId(v int64) *CallChatListRequest {
	s.TaskId = &v
	return s
}

func (s *CallChatListRequest) Validate() error {
	return dara.Validate(s)
}
