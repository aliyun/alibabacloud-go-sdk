// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReconnectAppChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *ReconnectAppChatRequest
	GetChatId() *string
	SetConversationId(v string) *ReconnectAppChatRequest
	GetConversationId() *string
	SetLastEventId(v int32) *ReconnectAppChatRequest
	GetLastEventId() *int32
	SetLatestMessageCreateTime(v string) *ReconnectAppChatRequest
	GetLatestMessageCreateTime() *string
}

type ReconnectAppChatRequest struct {
	// The chat record ID.
	//
	// example:
	//
	// 20833ba4-d189-4c50-9a44-a6bcbda2c93b
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// The conversation ID.
	//
	// example:
	//
	// 5b7105a2-2999-430b-ba23-ba09149d5434
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The last event ID.
	LastEventId *int32 `json:"LastEventId,omitempty" xml:"LastEventId,omitempty"`
	// The creation time of the latest message.
	LatestMessageCreateTime *string `json:"LatestMessageCreateTime,omitempty" xml:"LatestMessageCreateTime,omitempty"`
}

func (s ReconnectAppChatRequest) String() string {
	return dara.Prettify(s)
}

func (s ReconnectAppChatRequest) GoString() string {
	return s.String()
}

func (s *ReconnectAppChatRequest) GetChatId() *string {
	return s.ChatId
}

func (s *ReconnectAppChatRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *ReconnectAppChatRequest) GetLastEventId() *int32 {
	return s.LastEventId
}

func (s *ReconnectAppChatRequest) GetLatestMessageCreateTime() *string {
	return s.LatestMessageCreateTime
}

func (s *ReconnectAppChatRequest) SetChatId(v string) *ReconnectAppChatRequest {
	s.ChatId = &v
	return s
}

func (s *ReconnectAppChatRequest) SetConversationId(v string) *ReconnectAppChatRequest {
	s.ConversationId = &v
	return s
}

func (s *ReconnectAppChatRequest) SetLastEventId(v int32) *ReconnectAppChatRequest {
	s.LastEventId = &v
	return s
}

func (s *ReconnectAppChatRequest) SetLatestMessageCreateTime(v string) *ReconnectAppChatRequest {
	s.LatestMessageCreateTime = &v
	return s
}

func (s *ReconnectAppChatRequest) Validate() error {
	return dara.Validate(s)
}
