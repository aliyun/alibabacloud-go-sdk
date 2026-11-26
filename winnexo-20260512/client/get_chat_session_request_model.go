// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetChatSessionRequest
	GetAppToken() *string
	SetLimit(v int32) *GetChatSessionRequest
	GetLimit() *int32
	SetOpsUserCode(v string) *GetChatSessionRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetChatSessionRequest
	GetOpsUserId() *string
	SetSessionId(v string) *GetChatSessionRequest
	GetSessionId() *string
	SetTenantId(v string) *GetChatSessionRequest
	GetTenantId() *string
}

type GetChatSessionRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The maximum number of resources to return. If not specified, the default value is 20. The maximum value is 100. The number of returned results may be less than the specified value but will not exceed it.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
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
	// The SessionId value to filter by. If specified, all Active/Expired status information associated with this session is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The effective tenant ID.
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

func (s *GetChatSessionRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetChatSessionRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *GetChatSessionRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetChatSessionRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetChatSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetChatSessionRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetChatSessionRequest) SetAppToken(v string) *GetChatSessionRequest {
	s.AppToken = &v
	return s
}

func (s *GetChatSessionRequest) SetLimit(v int32) *GetChatSessionRequest {
	s.Limit = &v
	return s
}

func (s *GetChatSessionRequest) SetOpsUserCode(v string) *GetChatSessionRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetChatSessionRequest) SetOpsUserId(v string) *GetChatSessionRequest {
	s.OpsUserId = &v
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
