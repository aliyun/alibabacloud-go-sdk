// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentConversationShareScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetShareScope(v string) *UpdateAgentConversationShareScopeRequest
	GetShareScope() *string
}

type UpdateAgentConversationShareScopeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tenant_viewer
	ShareScope *string `json:"shareScope,omitempty" xml:"shareScope,omitempty"`
}

func (s UpdateAgentConversationShareScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentConversationShareScopeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentConversationShareScopeRequest) GetShareScope() *string {
	return s.ShareScope
}

func (s *UpdateAgentConversationShareScopeRequest) SetShareScope(v string) *UpdateAgentConversationShareScopeRequest {
	s.ShareScope = &v
	return s
}

func (s *UpdateAgentConversationShareScopeRequest) Validate() error {
	return dara.Validate(s)
}
