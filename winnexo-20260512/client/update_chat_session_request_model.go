// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateChatSessionRequest
	GetAppToken() *string
	SetModel(v string) *UpdateChatSessionRequest
	GetModel() *string
	SetOpsUserCode(v string) *UpdateChatSessionRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateChatSessionRequest
	GetOpsUserId() *string
	SetSessionId(v string) *UpdateChatSessionRequest
	GetSessionId() *string
	SetTenantId(v string) *UpdateChatSessionRequest
	GetTenantId() *string
	SetTitle(v string) *UpdateChatSessionRequest
	GetTitle() *string
}

type UpdateChatSessionRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The abstract model name (model tier). If not specified, the current model of the session is not modified.
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
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
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The new session title.
	//
	// example:
	//
	// Sample title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateChatSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatSessionRequest) GoString() string {
	return s.String()
}

func (s *UpdateChatSessionRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateChatSessionRequest) GetModel() *string {
	return s.Model
}

func (s *UpdateChatSessionRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateChatSessionRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateChatSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *UpdateChatSessionRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateChatSessionRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateChatSessionRequest) SetAppToken(v string) *UpdateChatSessionRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateChatSessionRequest) SetModel(v string) *UpdateChatSessionRequest {
	s.Model = &v
	return s
}

func (s *UpdateChatSessionRequest) SetOpsUserCode(v string) *UpdateChatSessionRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateChatSessionRequest) SetOpsUserId(v string) *UpdateChatSessionRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateChatSessionRequest) SetSessionId(v string) *UpdateChatSessionRequest {
	s.SessionId = &v
	return s
}

func (s *UpdateChatSessionRequest) SetTenantId(v string) *UpdateChatSessionRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateChatSessionRequest) SetTitle(v string) *UpdateChatSessionRequest {
	s.Title = &v
	return s
}

func (s *UpdateChatSessionRequest) Validate() error {
	return dara.Validate(s)
}
