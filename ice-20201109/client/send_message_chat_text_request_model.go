// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageChatTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIAgentId(v string) *SendMessageChatTextRequest
	GetAIAgentId() *string
	SetMode(v string) *SendMessageChatTextRequest
	GetMode() *string
	SetNeedArchiving(v bool) *SendMessageChatTextRequest
	GetNeedArchiving() *bool
	SetReceiverId(v string) *SendMessageChatTextRequest
	GetReceiverId() *string
	SetSessionId(v string) *SendMessageChatTextRequest
	GetSessionId() *string
	SetText(v string) *SendMessageChatTextRequest
	GetText() *string
	SetType(v string) *SendMessageChatTextRequest
	GetType() *string
}

type SendMessageChatTextRequest struct {
	// The ID of the AI agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	AIAgentId *string `json:"AIAgentId,omitempty" xml:"AIAgentId,omitempty"`
	// The mode of message sending. Valid values:
	//
	// - online
	//
	// - offline
	//
	// Default value: offline.
	//
	// example:
	//
	// online
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// Specifies whether to archive chat records. Default value: true.
	//
	// example:
	//
	// true
	NeedArchiving *bool `json:"NeedArchiving,omitempty" xml:"NeedArchiving,omitempty"`
	// The ID of the user who receives the message. The ID can be up to 64 bytes in length and can contain letters and digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60000042053
	ReceiverId *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	// The ID of the session.
	//
	// This parameter is required.
	//
	// example:
	//
	// f27f9b9be28642a88e18****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The content of the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hello
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The type of the message. Valid values:
	//
	// - announcement: notification.
	//
	// - custom: custom message.
	//
	// This parameter is required.
	//
	// example:
	//
	// announcement
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SendMessageChatTextRequest) String() string {
	return dara.Prettify(s)
}

func (s SendMessageChatTextRequest) GoString() string {
	return s.String()
}

func (s *SendMessageChatTextRequest) GetAIAgentId() *string {
	return s.AIAgentId
}

func (s *SendMessageChatTextRequest) GetMode() *string {
	return s.Mode
}

func (s *SendMessageChatTextRequest) GetNeedArchiving() *bool {
	return s.NeedArchiving
}

func (s *SendMessageChatTextRequest) GetReceiverId() *string {
	return s.ReceiverId
}

func (s *SendMessageChatTextRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SendMessageChatTextRequest) GetText() *string {
	return s.Text
}

func (s *SendMessageChatTextRequest) GetType() *string {
	return s.Type
}

func (s *SendMessageChatTextRequest) SetAIAgentId(v string) *SendMessageChatTextRequest {
	s.AIAgentId = &v
	return s
}

func (s *SendMessageChatTextRequest) SetMode(v string) *SendMessageChatTextRequest {
	s.Mode = &v
	return s
}

func (s *SendMessageChatTextRequest) SetNeedArchiving(v bool) *SendMessageChatTextRequest {
	s.NeedArchiving = &v
	return s
}

func (s *SendMessageChatTextRequest) SetReceiverId(v string) *SendMessageChatTextRequest {
	s.ReceiverId = &v
	return s
}

func (s *SendMessageChatTextRequest) SetSessionId(v string) *SendMessageChatTextRequest {
	s.SessionId = &v
	return s
}

func (s *SendMessageChatTextRequest) SetText(v string) *SendMessageChatTextRequest {
	s.Text = &v
	return s
}

func (s *SendMessageChatTextRequest) SetType(v string) *SendMessageChatTextRequest {
	s.Type = &v
	return s
}

func (s *SendMessageChatTextRequest) Validate() error {
	return dara.Validate(s)
}
