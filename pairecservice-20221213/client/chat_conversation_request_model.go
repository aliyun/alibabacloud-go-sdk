// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ChatConversationRequest
	GetConfig() *string
	SetContent(v string) *ChatConversationRequest
	GetContent() *string
	SetConversationId(v string) *ChatConversationRequest
	GetConversationId() *string
	SetInstanceId(v string) *ChatConversationRequest
	GetInstanceId() *string
}

type ChatConversationRequest struct {
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// e47cfae9-c0cc-42e1-91e2-e67cdb0e7b96
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// learn-pairec-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ChatConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatConversationRequest) GoString() string {
	return s.String()
}

func (s *ChatConversationRequest) GetConfig() *string {
	return s.Config
}

func (s *ChatConversationRequest) GetContent() *string {
	return s.Content
}

func (s *ChatConversationRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *ChatConversationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ChatConversationRequest) SetConfig(v string) *ChatConversationRequest {
	s.Config = &v
	return s
}

func (s *ChatConversationRequest) SetContent(v string) *ChatConversationRequest {
	s.Content = &v
	return s
}

func (s *ChatConversationRequest) SetConversationId(v string) *ChatConversationRequest {
	s.ConversationId = &v
	return s
}

func (s *ChatConversationRequest) SetInstanceId(v string) *ChatConversationRequest {
	s.InstanceId = &v
	return s
}

func (s *ChatConversationRequest) Validate() error {
	return dara.Validate(s)
}
