// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *GetChatSessionRequest
	GetLimit() *int32
	SetSessionId(v string) *GetChatSessionRequest
	GetSessionId() *string
	SetTenantId(v string) *GetChatSessionRequest
	GetTenantId() *string
}

type GetChatSessionRequest struct {
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
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

func (s GetChatSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatSessionRequest) GoString() string {
	return s.String()
}

func (s *GetChatSessionRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *GetChatSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetChatSessionRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetChatSessionRequest) SetLimit(v int32) *GetChatSessionRequest {
	s.Limit = &v
	return s
}

func (s *GetChatSessionRequest) SetSessionId(v string) *GetChatSessionRequest {
	s.SessionId = &v
	return s
}

func (s *GetChatSessionRequest) SetTenantId(v string) *GetChatSessionRequest {
	s.TenantId = &v
	return s
}

func (s *GetChatSessionRequest) Validate() error {
	return dara.Validate(s)
}
