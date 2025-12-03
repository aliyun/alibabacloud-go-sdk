// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatMessagesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *ChatMessagesShrinkRequest
	GetApiId() *string
	SetConversationId(v string) *ChatMessagesShrinkRequest
	GetConversationId() *string
	SetInputsShrink(v string) *ChatMessagesShrinkRequest
	GetInputsShrink() *string
	SetParentMessageId(v string) *ChatMessagesShrinkRequest
	GetParentMessageId() *string
	SetQuery(v string) *ChatMessagesShrinkRequest
	GetQuery() *string
}

type ChatMessagesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// app-iBuGU1VxEY42zrQRQfNA****
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// fea7bdca-e848-44dd-b1ae-852472b8****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	InputsShrink   *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// example:
	//
	// 84dc9f9b-424a-404d-9c36-35e9d000****
	ParentMessageId *string `json:"ParentMessageId,omitempty" xml:"ParentMessageId,omitempty"`
	// This parameter is required.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s ChatMessagesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatMessagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChatMessagesShrinkRequest) GetApiId() *string {
	return s.ApiId
}

func (s *ChatMessagesShrinkRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *ChatMessagesShrinkRequest) GetInputsShrink() *string {
	return s.InputsShrink
}

func (s *ChatMessagesShrinkRequest) GetParentMessageId() *string {
	return s.ParentMessageId
}

func (s *ChatMessagesShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *ChatMessagesShrinkRequest) SetApiId(v string) *ChatMessagesShrinkRequest {
	s.ApiId = &v
	return s
}

func (s *ChatMessagesShrinkRequest) SetConversationId(v string) *ChatMessagesShrinkRequest {
	s.ConversationId = &v
	return s
}

func (s *ChatMessagesShrinkRequest) SetInputsShrink(v string) *ChatMessagesShrinkRequest {
	s.InputsShrink = &v
	return s
}

func (s *ChatMessagesShrinkRequest) SetParentMessageId(v string) *ChatMessagesShrinkRequest {
	s.ParentMessageId = &v
	return s
}

func (s *ChatMessagesShrinkRequest) SetQuery(v string) *ChatMessagesShrinkRequest {
	s.Query = &v
	return s
}

func (s *ChatMessagesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
