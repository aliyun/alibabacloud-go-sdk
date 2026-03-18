// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConversationCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetChatData(v string) *ConversationCreateCmd
	GetChatData() *string
	SetModelIds(v string) *ConversationCreateCmd
	GetModelIds() *string
	SetTitle(v string) *ConversationCreateCmd
	GetTitle() *string
}

type ConversationCreateCmd struct {
	ChatData *string `json:"chatData,omitempty" xml:"chatData,omitempty"`
	ModelIds *string `json:"modelIds,omitempty" xml:"modelIds,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ConversationCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s ConversationCreateCmd) GoString() string {
	return s.String()
}

func (s *ConversationCreateCmd) GetChatData() *string {
	return s.ChatData
}

func (s *ConversationCreateCmd) GetModelIds() *string {
	return s.ModelIds
}

func (s *ConversationCreateCmd) GetTitle() *string {
	return s.Title
}

func (s *ConversationCreateCmd) SetChatData(v string) *ConversationCreateCmd {
	s.ChatData = &v
	return s
}

func (s *ConversationCreateCmd) SetModelIds(v string) *ConversationCreateCmd {
	s.ModelIds = &v
	return s
}

func (s *ConversationCreateCmd) SetTitle(v string) *ConversationCreateCmd {
	s.Title = &v
	return s
}

func (s *ConversationCreateCmd) Validate() error {
	return dara.Validate(s)
}
