// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliDingGroupMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListAliDingGroupMessagesRequest
	GetAppToken() *string
	SetChatId(v string) *ListAliDingGroupMessagesRequest
	GetChatId() *string
	SetDirection(v string) *ListAliDingGroupMessagesRequest
	GetDirection() *string
	SetOpsUserCode(v string) *ListAliDingGroupMessagesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListAliDingGroupMessagesRequest
	GetOpsUserId() *string
	SetPageSize(v int32) *ListAliDingGroupMessagesRequest
	GetPageSize() *int32
	SetTenantId(v string) *ListAliDingGroupMessagesRequest
	GetTenantId() *string
	SetTime(v string) *ListAliDingGroupMessagesRequest
	GetTime() *string
}

type ListAliDingGroupMessagesRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The session ID, typically used for JSSDK.
	//
	// This parameter is required.
	//
	// example:
	//
	// cid-example
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// The traffic direction. Valid values:
	//
	// - OutBound: outbound.
	//
	// - InBound: inbound.
	//
	// - Both: bidirectional.
	//
	// example:
	//
	// newer
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
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
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The tenant ID. This is a common parameter. Pass it explicitly through `--tenant-id` in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The relationship information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-08T09:00:00+08:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s ListAliDingGroupMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAliDingGroupMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListAliDingGroupMessagesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListAliDingGroupMessagesRequest) GetChatId() *string {
	return s.ChatId
}

func (s *ListAliDingGroupMessagesRequest) GetDirection() *string {
	return s.Direction
}

func (s *ListAliDingGroupMessagesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListAliDingGroupMessagesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListAliDingGroupMessagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAliDingGroupMessagesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAliDingGroupMessagesRequest) GetTime() *string {
	return s.Time
}

func (s *ListAliDingGroupMessagesRequest) SetAppToken(v string) *ListAliDingGroupMessagesRequest {
	s.AppToken = &v
	return s
}

func (s *ListAliDingGroupMessagesRequest) SetChatId(v string) *ListAliDingGroupMessagesRequest {
	s.ChatId = &v
	return s
}

func (s *ListAliDingGroupMessagesRequest) SetDirection(v string) *ListAliDingGroupMessagesRequest {
	s.Direction = &v
	return s
}

func (s *ListAliDingGroupMessagesRequest) SetOpsUserCode(v string) *ListAliDingGroupMessagesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListAliDingGroupMessagesRequest) SetOpsUserId(v string) *ListAliDingGroupMessagesRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListAliDingGroupMessagesRequest) SetPageSize(v int32) *ListAliDingGroupMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAliDingGroupMessagesRequest) SetTenantId(v string) *ListAliDingGroupMessagesRequest {
	s.TenantId = &v
	return s
}

func (s *ListAliDingGroupMessagesRequest) SetTime(v string) *ListAliDingGroupMessagesRequest {
	s.Time = &v
	return s
}

func (s *ListAliDingGroupMessagesRequest) Validate() error {
	return dara.Validate(s)
}
