// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatData(v string) *ModelRouterUpdateConversationRequest
	GetChatData() *string
	SetMessageCount(v int32) *ModelRouterUpdateConversationRequest
	GetMessageCount() *int32
	SetModelIds(v string) *ModelRouterUpdateConversationRequest
	GetModelIds() *string
	SetTitle(v string) *ModelRouterUpdateConversationRequest
	GetTitle() *string
}

type ModelRouterUpdateConversationRequest struct {
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

func (s ModelRouterUpdateConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateConversationRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateConversationRequest) GetChatData() *string {
	return s.ChatData
}

func (s *ModelRouterUpdateConversationRequest) GetMessageCount() *int32 {
	return s.MessageCount
}

func (s *ModelRouterUpdateConversationRequest) GetModelIds() *string {
	return s.ModelIds
}

func (s *ModelRouterUpdateConversationRequest) GetTitle() *string {
	return s.Title
}

func (s *ModelRouterUpdateConversationRequest) SetChatData(v string) *ModelRouterUpdateConversationRequest {
	s.ChatData = &v
	return s
}

func (s *ModelRouterUpdateConversationRequest) SetMessageCount(v int32) *ModelRouterUpdateConversationRequest {
	s.MessageCount = &v
	return s
}

func (s *ModelRouterUpdateConversationRequest) SetModelIds(v string) *ModelRouterUpdateConversationRequest {
	s.ModelIds = &v
	return s
}

func (s *ModelRouterUpdateConversationRequest) SetTitle(v string) *ModelRouterUpdateConversationRequest {
	s.Title = &v
	return s
}

func (s *ModelRouterUpdateConversationRequest) Validate() error {
	return dara.Validate(s)
}
