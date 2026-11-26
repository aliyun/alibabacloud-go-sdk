// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertDirectConversationToWorkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ConvertDirectConversationToWorkRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *ConvertDirectConversationToWorkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ConvertDirectConversationToWorkRequest
	GetOpsUserId() *string
	SetSessionId(v string) *ConvertDirectConversationToWorkRequest
	GetSessionId() *string
	SetTenantId(v string) *ConvertDirectConversationToWorkRequest
	GetTenantId() *string
}

type ConvertDirectConversationToWorkRequest struct {
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
	// The source direct session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ConvertDirectConversationToWorkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertDirectConversationToWorkRequest) GoString() string {
	return s.String()
}

func (s *ConvertDirectConversationToWorkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ConvertDirectConversationToWorkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ConvertDirectConversationToWorkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ConvertDirectConversationToWorkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ConvertDirectConversationToWorkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ConvertDirectConversationToWorkRequest) SetAppToken(v string) *ConvertDirectConversationToWorkRequest {
	s.AppToken = &v
	return s
}

func (s *ConvertDirectConversationToWorkRequest) SetOpsUserCode(v string) *ConvertDirectConversationToWorkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ConvertDirectConversationToWorkRequest) SetOpsUserId(v string) *ConvertDirectConversationToWorkRequest {
	s.OpsUserId = &v
	return s
}

func (s *ConvertDirectConversationToWorkRequest) SetSessionId(v string) *ConvertDirectConversationToWorkRequest {
	s.SessionId = &v
	return s
}

func (s *ConvertDirectConversationToWorkRequest) SetTenantId(v string) *ConvertDirectConversationToWorkRequest {
	s.TenantId = &v
	return s
}

func (s *ConvertDirectConversationToWorkRequest) Validate() error {
	return dara.Validate(s)
}
