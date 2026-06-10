// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIStaffChatEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListAIStaffChatEventsRequest
	GetBizId() *string
	SetChatId(v string) *ListAIStaffChatEventsRequest
	GetChatId() *string
	SetConversationId(v string) *ListAIStaffChatEventsRequest
	GetConversationId() *string
	SetLastEventId(v int32) *ListAIStaffChatEventsRequest
	GetLastEventId() *int32
}

type ListAIStaffChatEventsRequest struct {
	// Business ID
	//
	// example:
	//
	// WS20250731233102000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Chat ID (optional; if not provided, the latest chatId is used)
	//
	// example:
	//
	// 20833ba4-d189-4c50-9a44-a6bcbda2c93b
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// session ID
	//
	// example:
	//
	// 593fe1a2-d0b4-4fde-a2b0-78ad6a438d41
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// Last event ID, used for incremental retrieval
	//
	// example:
	//
	// event-71ece53b7d474e01b755a9b5fa5305e6
	LastEventId *int32 `json:"LastEventId,omitempty" xml:"LastEventId,omitempty"`
}

func (s ListAIStaffChatEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIStaffChatEventsRequest) GoString() string {
	return s.String()
}

func (s *ListAIStaffChatEventsRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListAIStaffChatEventsRequest) GetChatId() *string {
	return s.ChatId
}

func (s *ListAIStaffChatEventsRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListAIStaffChatEventsRequest) GetLastEventId() *int32 {
	return s.LastEventId
}

func (s *ListAIStaffChatEventsRequest) SetBizId(v string) *ListAIStaffChatEventsRequest {
	s.BizId = &v
	return s
}

func (s *ListAIStaffChatEventsRequest) SetChatId(v string) *ListAIStaffChatEventsRequest {
	s.ChatId = &v
	return s
}

func (s *ListAIStaffChatEventsRequest) SetConversationId(v string) *ListAIStaffChatEventsRequest {
	s.ConversationId = &v
	return s
}

func (s *ListAIStaffChatEventsRequest) SetLastEventId(v int32) *ListAIStaffChatEventsRequest {
	s.LastEventId = &v
	return s
}

func (s *ListAIStaffChatEventsRequest) Validate() error {
	return dara.Validate(s)
}
