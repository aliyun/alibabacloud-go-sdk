// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatData(v string) *ModelRouterCreateConversationRequest
	GetChatData() *string
	SetModelIds(v string) *ModelRouterCreateConversationRequest
	GetModelIds() *string
	SetTitle(v string) *ModelRouterCreateConversationRequest
	GetTitle() *string
}

type ModelRouterCreateConversationRequest struct {
	// example:
	//
	// {"stream":true,"messages":[{"role":"user","content":"1+1"}],"model":"qwen/qwen-max/r0","stream_options":{"include_usage":true}}
	ChatData *string `json:"chatData,omitempty" xml:"chatData,omitempty"`
	// example:
	//
	// 15
	ModelIds *string `json:"modelIds,omitempty" xml:"modelIds,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ModelRouterCreateConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateConversationRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateConversationRequest) GetChatData() *string {
	return s.ChatData
}

func (s *ModelRouterCreateConversationRequest) GetModelIds() *string {
	return s.ModelIds
}

func (s *ModelRouterCreateConversationRequest) GetTitle() *string {
	return s.Title
}

func (s *ModelRouterCreateConversationRequest) SetChatData(v string) *ModelRouterCreateConversationRequest {
	s.ChatData = &v
	return s
}

func (s *ModelRouterCreateConversationRequest) SetModelIds(v string) *ModelRouterCreateConversationRequest {
	s.ModelIds = &v
	return s
}

func (s *ModelRouterCreateConversationRequest) SetTitle(v string) *ModelRouterCreateConversationRequest {
	s.Title = &v
	return s
}

func (s *ModelRouterCreateConversationRequest) Validate() error {
	return dara.Validate(s)
}
