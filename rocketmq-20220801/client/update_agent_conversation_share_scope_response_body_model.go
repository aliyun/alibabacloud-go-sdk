// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentConversationShareScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateAgentConversationShareScopeResponseBody
	GetId() *string
	SetShareScope(v string) *UpdateAgentConversationShareScopeResponseBody
	GetShareScope() *string
	SetTitle(v string) *UpdateAgentConversationShareScopeResponseBody
	GetTitle() *string
	SetUpdatedAt(v string) *UpdateAgentConversationShareScopeResponseBody
	GetUpdatedAt() *string
}

type UpdateAgentConversationShareScopeResponseBody struct {
	// example:
	//
	// 0198c8a91f4d7f01a3c21e97
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// tenant_viewer
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

func (s UpdateAgentConversationShareScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentConversationShareScopeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentConversationShareScopeResponseBody) GetId() *string {
	return s.Id
}

func (s *UpdateAgentConversationShareScopeResponseBody) GetShareScope() *string {
	return s.ShareScope
}

func (s *UpdateAgentConversationShareScopeResponseBody) GetTitle() *string {
	return s.Title
}

func (s *UpdateAgentConversationShareScopeResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateAgentConversationShareScopeResponseBody) SetId(v string) *UpdateAgentConversationShareScopeResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateAgentConversationShareScopeResponseBody) SetShareScope(v string) *UpdateAgentConversationShareScopeResponseBody {
	s.ShareScope = &v
	return s
}

func (s *UpdateAgentConversationShareScopeResponseBody) SetTitle(v string) *UpdateAgentConversationShareScopeResponseBody {
	s.Title = &v
	return s
}

func (s *UpdateAgentConversationShareScopeResponseBody) SetUpdatedAt(v string) *UpdateAgentConversationShareScopeResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateAgentConversationShareScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
