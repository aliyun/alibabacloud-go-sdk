// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *DeleteChatSessionRequest
	GetSessionId() *string
	SetTenantId(v string) *DeleteChatSessionRequest
	GetTenantId() *string
}

type DeleteChatSessionRequest struct {
	// The ID of the session to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteChatSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatSessionRequest) GoString() string {
	return s.String()
}

func (s *DeleteChatSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DeleteChatSessionRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteChatSessionRequest) SetSessionId(v string) *DeleteChatSessionRequest {
	s.SessionId = &v
	return s
}

func (s *DeleteChatSessionRequest) SetTenantId(v string) *DeleteChatSessionRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteChatSessionRequest) Validate() error {
	return dara.Validate(s)
}
