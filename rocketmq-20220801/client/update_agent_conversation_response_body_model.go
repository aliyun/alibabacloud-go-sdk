// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *UpdateAgentConversationResponseBody
	GetConversationId() *string
	SetId(v string) *UpdateAgentConversationResponseBody
	GetId() *string
	SetShareScope(v string) *UpdateAgentConversationResponseBody
	GetShareScope() *string
	SetTitle(v string) *UpdateAgentConversationResponseBody
	GetTitle() *string
	SetUpdatedAt(v string) *UpdateAgentConversationResponseBody
	GetUpdatedAt() *string
}

type UpdateAgentConversationResponseBody struct {
	// example:
	//
	// 0198c8a91f4d7f01a3c21e97
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// example:
	//
	// 0198c8a91f4d7f01a3c21e97
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// private
	ShareScope *string `json:"shareScope,omitempty" xml:"shareScope,omitempty"`
	// example:
	//
	// 消息堆积排查
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2026-08-26T12:00:00.000Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s UpdateAgentConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentConversationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentConversationResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *UpdateAgentConversationResponseBody) GetId() *string {
	return s.Id
}

func (s *UpdateAgentConversationResponseBody) GetShareScope() *string {
	return s.ShareScope
}

func (s *UpdateAgentConversationResponseBody) GetTitle() *string {
	return s.Title
}

func (s *UpdateAgentConversationResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateAgentConversationResponseBody) SetConversationId(v string) *UpdateAgentConversationResponseBody {
	s.ConversationId = &v
	return s
}

func (s *UpdateAgentConversationResponseBody) SetId(v string) *UpdateAgentConversationResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateAgentConversationResponseBody) SetShareScope(v string) *UpdateAgentConversationResponseBody {
	s.ShareScope = &v
	return s
}

func (s *UpdateAgentConversationResponseBody) SetTitle(v string) *UpdateAgentConversationResponseBody {
	s.Title = &v
	return s
}

func (s *UpdateAgentConversationResponseBody) SetUpdatedAt(v string) *UpdateAgentConversationResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateAgentConversationResponseBody) Validate() error {
	return dara.Validate(s)
}
