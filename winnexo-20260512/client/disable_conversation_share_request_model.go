// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableConversationShareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *DisableConversationShareRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *DisableConversationShareRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *DisableConversationShareRequest
	GetOpsUserId() *string
	SetSessionId(v string) *DisableConversationShareRequest
	GetSessionId() *string
	SetTenantId(v string) *DisableConversationShareRequest
	GetTenantId() *string
}

type DisableConversationShareRequest struct {
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

func (s *DisableConversationShareRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *DisableConversationShareRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *DisableConversationShareRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *DisableConversationShareRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DisableConversationShareRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DisableConversationShareRequest) SetAppToken(v string) *DisableConversationShareRequest {
	s.AppToken = &v
	return s
}

func (s *DisableConversationShareRequest) SetOpsUserCode(v string) *DisableConversationShareRequest {
	s.OpsUserCode = &v
	return s
}

func (s *DisableConversationShareRequest) SetOpsUserId(v string) *DisableConversationShareRequest {
	s.OpsUserId = &v
	return s
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
