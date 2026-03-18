// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConversationUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetChatData(v string) *ConversationUpdateCmd
	GetChatData() *string
	SetMessageCount(v int32) *ConversationUpdateCmd
	GetMessageCount() *int32
	SetModelIds(v string) *ConversationUpdateCmd
	GetModelIds() *string
	SetTitle(v string) *ConversationUpdateCmd
	GetTitle() *string
}

type ConversationUpdateCmd struct {
	// example:
	//
	// {"messages":[{"role":"user","content":"你好"}]}
	ChatData *string `json:"chatData,omitempty" xml:"chatData,omitempty"`
	// example:
	//
	// 10
	MessageCount *int32 `json:"messageCount,omitempty" xml:"messageCount,omitempty"`
	// example:
	//
	// [1,2,3]
	ModelIds *string `json:"modelIds,omitempty" xml:"modelIds,omitempty"`
	// example:
	//
	// 我的对话
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ConversationUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s ConversationUpdateCmd) GoString() string {
	return s.String()
}

func (s *ConversationUpdateCmd) GetChatData() *string {
	return s.ChatData
}

func (s *ConversationUpdateCmd) GetMessageCount() *int32 {
	return s.MessageCount
}

func (s *ConversationUpdateCmd) GetModelIds() *string {
	return s.ModelIds
}

func (s *ConversationUpdateCmd) GetTitle() *string {
	return s.Title
}

func (s *ConversationUpdateCmd) SetChatData(v string) *ConversationUpdateCmd {
	s.ChatData = &v
	return s
}

func (s *ConversationUpdateCmd) SetMessageCount(v int32) *ConversationUpdateCmd {
	s.MessageCount = &v
	return s
}

func (s *ConversationUpdateCmd) SetModelIds(v string) *ConversationUpdateCmd {
	s.ModelIds = &v
	return s
}

func (s *ConversationUpdateCmd) SetTitle(v string) *ConversationUpdateCmd {
	s.Title = &v
	return s
}

func (s *ConversationUpdateCmd) Validate() error {
	return dara.Validate(s)
}
