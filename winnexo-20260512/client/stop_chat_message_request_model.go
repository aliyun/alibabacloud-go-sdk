// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopChatMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *StopChatMessageRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *StopChatMessageRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *StopChatMessageRequest
	GetOpsUserId() *string
	SetSessionId(v string) *StopChatMessageRequest
	GetSessionId() *string
	SetTenantId(v string) *StopChatMessageRequest
	GetTenantId() *string
}

type StopChatMessageRequest struct {
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
	// The session ID.
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

func (s StopChatMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s StopChatMessageRequest) GoString() string {
	return s.String()
}

func (s *StopChatMessageRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *StopChatMessageRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *StopChatMessageRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *StopChatMessageRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StopChatMessageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *StopChatMessageRequest) SetAppToken(v string) *StopChatMessageRequest {
	s.AppToken = &v
	return s
}

func (s *StopChatMessageRequest) SetOpsUserCode(v string) *StopChatMessageRequest {
	s.OpsUserCode = &v
	return s
}

func (s *StopChatMessageRequest) SetOpsUserId(v string) *StopChatMessageRequest {
	s.OpsUserId = &v
	return s
}

func (s *StopChatMessageRequest) SetSessionId(v string) *StopChatMessageRequest {
	s.SessionId = &v
	return s
}

func (s *StopChatMessageRequest) SetTenantId(v string) *StopChatMessageRequest {
	s.TenantId = &v
	return s
}

func (s *StopChatMessageRequest) Validate() error {
	return dara.Validate(s)
}
