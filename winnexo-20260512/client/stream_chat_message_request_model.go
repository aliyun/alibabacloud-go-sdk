// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStreamChatMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *StreamChatMessageRequest
	GetAppToken() *string
	SetLastEventId(v string) *StreamChatMessageRequest
	GetLastEventId() *string
	SetOpsUserCode(v string) *StreamChatMessageRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *StreamChatMessageRequest
	GetOpsUserId() *string
	SetTenantId(v string) *StreamChatMessageRequest
	GetTenantId() *string
}

type StreamChatMessageRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The last received SSE event ID, used for resuming delivery after a disconnection. If this parameter is not specified, full replay starts from the beginning.
	//
	// example:
	//
	// 5-1683456789012
	LastEventId *string `json:"lastEventId,omitempty" xml:"lastEventId,omitempty"`
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
	// The ID of the tenant for which the subscription takes effect.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s StreamChatMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s StreamChatMessageRequest) GoString() string {
	return s.String()
}

func (s *StreamChatMessageRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *StreamChatMessageRequest) GetLastEventId() *string {
	return s.LastEventId
}

func (s *StreamChatMessageRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *StreamChatMessageRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *StreamChatMessageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *StreamChatMessageRequest) SetAppToken(v string) *StreamChatMessageRequest {
	s.AppToken = &v
	return s
}

func (s *StreamChatMessageRequest) SetLastEventId(v string) *StreamChatMessageRequest {
	s.LastEventId = &v
	return s
}

func (s *StreamChatMessageRequest) SetOpsUserCode(v string) *StreamChatMessageRequest {
	s.OpsUserCode = &v
	return s
}

func (s *StreamChatMessageRequest) SetOpsUserId(v string) *StreamChatMessageRequest {
	s.OpsUserId = &v
	return s
}

func (s *StreamChatMessageRequest) SetTenantId(v string) *StreamChatMessageRequest {
	s.TenantId = &v
	return s
}

func (s *StreamChatMessageRequest) Validate() error {
	return dara.Validate(s)
}
