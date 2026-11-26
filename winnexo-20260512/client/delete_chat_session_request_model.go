// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *DeleteChatSessionRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *DeleteChatSessionRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *DeleteChatSessionRequest
	GetOpsUserId() *string
	SetSessionId(v string) *DeleteChatSessionRequest
	GetSessionId() *string
	SetTenantId(v string) *DeleteChatSessionRequest
	GetTenantId() *string
}

type DeleteChatSessionRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The ID of the session to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The tenant ID that takes effect.
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

func (s *DeleteChatSessionRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *DeleteChatSessionRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *DeleteChatSessionRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *DeleteChatSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DeleteChatSessionRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteChatSessionRequest) SetAppToken(v string) *DeleteChatSessionRequest {
	s.AppToken = &v
	return s
}

func (s *DeleteChatSessionRequest) SetOpsUserCode(v string) *DeleteChatSessionRequest {
	s.OpsUserCode = &v
	return s
}

func (s *DeleteChatSessionRequest) SetOpsUserId(v string) *DeleteChatSessionRequest {
	s.OpsUserId = &v
	return s
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
