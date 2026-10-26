// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateAgentConversationRequest
	GetClientToken() *string
	SetTitle(v string) *UpdateAgentConversationRequest
	GetTitle() *string
}

type UpdateAgentConversationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// conversation-request-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 消息堆积排查
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateAgentConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentConversationRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentConversationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAgentConversationRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateAgentConversationRequest) SetClientToken(v string) *UpdateAgentConversationRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAgentConversationRequest) SetTitle(v string) *UpdateAgentConversationRequest {
	s.Title = &v
	return s
}

func (s *UpdateAgentConversationRequest) Validate() error {
	return dara.Validate(s)
}
