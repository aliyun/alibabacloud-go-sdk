// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeAppProxyOpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *AuthorizeAppProxyOpsRequest
	GetConversationId() *string
}

type AuthorizeAppProxyOpsRequest struct {
	// The session ID.
	//
	// example:
	//
	// 593fe1a2-d0b4-4fde-a2b0-78ad6a438d41
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
}

func (s AuthorizeAppProxyOpsRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeAppProxyOpsRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeAppProxyOpsRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *AuthorizeAppProxyOpsRequest) SetConversationId(v string) *AuthorizeAppProxyOpsRequest {
	s.ConversationId = &v
	return s
}

func (s *AuthorizeAppProxyOpsRequest) Validate() error {
	return dara.Validate(s)
}
