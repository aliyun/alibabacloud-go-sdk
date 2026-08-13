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
	// 会话ID
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
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
