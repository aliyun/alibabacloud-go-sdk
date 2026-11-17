// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableConversationShareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *DisableConversationShareRequest
	GetSessionId() *string
	SetTenantId(v string) *DisableConversationShareRequest
	GetTenantId() *string
}

type DisableConversationShareRequest struct {
	// The ID of the source one-on-one session.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DisableConversationShareRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableConversationShareRequest) GoString() string {
	return s.String()
}

func (s *DisableConversationShareRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DisableConversationShareRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DisableConversationShareRequest) SetSessionId(v string) *DisableConversationShareRequest {
	s.SessionId = &v
	return s
}

func (s *DisableConversationShareRequest) SetTenantId(v string) *DisableConversationShareRequest {
	s.TenantId = &v
	return s
}

func (s *DisableConversationShareRequest) Validate() error {
	return dara.Validate(s)
}
