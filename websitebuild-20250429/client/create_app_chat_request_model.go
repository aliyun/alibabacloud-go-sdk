// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBotId(v string) *CreateAppChatRequest
	GetBotId() *string
	SetChatId(v string) *CreateAppChatRequest
	GetChatId() *string
	SetConversationId(v string) *CreateAppChatRequest
	GetConversationId() *string
	SetMessages(v string) *CreateAppChatRequest
	GetMessages() *string
	SetSiteId(v string) *CreateAppChatRequest
	GetSiteId() *string
}

type CreateAppChatRequest struct {
	// Bot ID
	//
	// example:
	//
	// Zero2
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// Chat ID to be provided when recovering a conversation after an execution break
	//
	// example:
	//
	// 3b465fe1-6f06-4899-af9f-d43d9338df25
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// Session ID; required from the second turn onward in a multi-turn conversation scenario
	//
	// example:
	//
	// 593fe1a2-d0b4-4fde-a2b0-78ad6a438d41
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// List of conversation messages (in JSON array format)
	//
	// example:
	//
	// [\\r\\n  {\\r\\n    \\"role\\": \\"user\\",\\r\\n    \\"content\\": \\"Prove that there are only five platonic solids, namely the tetrahedron, octahedron, icosahedron, cube, and dodecahedron.\\"\\r\\n  }\\r\\n]
	Messages *string `json:"Messages,omitempty" xml:"Messages,omitempty"`
	// Site ID
	//
	// example:
	//
	// 857240041851344
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CreateAppChatRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppChatRequest) GoString() string {
	return s.String()
}

func (s *CreateAppChatRequest) GetBotId() *string {
	return s.BotId
}

func (s *CreateAppChatRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreateAppChatRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *CreateAppChatRequest) GetMessages() *string {
	return s.Messages
}

func (s *CreateAppChatRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *CreateAppChatRequest) SetBotId(v string) *CreateAppChatRequest {
	s.BotId = &v
	return s
}

func (s *CreateAppChatRequest) SetChatId(v string) *CreateAppChatRequest {
	s.ChatId = &v
	return s
}

func (s *CreateAppChatRequest) SetConversationId(v string) *CreateAppChatRequest {
	s.ConversationId = &v
	return s
}

func (s *CreateAppChatRequest) SetMessages(v string) *CreateAppChatRequest {
	s.Messages = &v
	return s
}

func (s *CreateAppChatRequest) SetSiteId(v string) *CreateAppChatRequest {
	s.SiteId = &v
	return s
}

func (s *CreateAppChatRequest) Validate() error {
	return dara.Validate(s)
}
