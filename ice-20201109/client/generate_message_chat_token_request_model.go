// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateMessageChatTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIAgentId(v string) *GenerateMessageChatTokenRequest
	GetAIAgentId() *string
	SetExpire(v int32) *GenerateMessageChatTokenRequest
	GetExpire() *int32
	SetRole(v string) *GenerateMessageChatTokenRequest
	GetRole() *string
	SetUserId(v string) *GenerateMessageChatTokenRequest
	GetUserId() *string
}

type GenerateMessageChatTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 19de81b3b3d94abda22******
	AIAgentId *string `json:"AIAgentId,omitempty" xml:"AIAgentId,omitempty"`
	// example:
	//
	// 3600
	Expire *int32 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// YOURUSERID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GenerateMessageChatTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateMessageChatTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateMessageChatTokenRequest) GetAIAgentId() *string {
	return s.AIAgentId
}

func (s *GenerateMessageChatTokenRequest) GetExpire() *int32 {
	return s.Expire
}

func (s *GenerateMessageChatTokenRequest) GetRole() *string {
	return s.Role
}

func (s *GenerateMessageChatTokenRequest) GetUserId() *string {
	return s.UserId
}

func (s *GenerateMessageChatTokenRequest) SetAIAgentId(v string) *GenerateMessageChatTokenRequest {
	s.AIAgentId = &v
	return s
}

func (s *GenerateMessageChatTokenRequest) SetExpire(v int32) *GenerateMessageChatTokenRequest {
	s.Expire = &v
	return s
}

func (s *GenerateMessageChatTokenRequest) SetRole(v string) *GenerateMessageChatTokenRequest {
	s.Role = &v
	return s
}

func (s *GenerateMessageChatTokenRequest) SetUserId(v string) *GenerateMessageChatTokenRequest {
	s.UserId = &v
	return s
}

func (s *GenerateMessageChatTokenRequest) Validate() error {
	return dara.Validate(s)
}
