// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBotId(v string) *GetAppConversationRequest
	GetBotId() *string
	SetConversationId(v string) *GetAppConversationRequest
	GetConversationId() *string
}

type GetAppConversationRequest struct {
	// example:
	//
	// Zero2
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// example:
	//
	// 81bc5a34-1d8d-4ef7-a208-7401c51b054b
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
}

func (s GetAppConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppConversationRequest) GoString() string {
	return s.String()
}

func (s *GetAppConversationRequest) GetBotId() *string {
	return s.BotId
}

func (s *GetAppConversationRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetAppConversationRequest) SetBotId(v string) *GetAppConversationRequest {
	s.BotId = &v
	return s
}

func (s *GetAppConversationRequest) SetConversationId(v string) *GetAppConversationRequest {
	s.ConversationId = &v
	return s
}

func (s *GetAppConversationRequest) Validate() error {
	return dara.Validate(s)
}
