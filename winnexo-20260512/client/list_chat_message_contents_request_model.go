// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatMessageContentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListChatMessageContentsRequest
	GetAppToken() *string
	SetMaxCharsPerAttachment(v int32) *ListChatMessageContentsRequest
	GetMaxCharsPerAttachment() *int32
	SetOpsUserCode(v string) *ListChatMessageContentsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListChatMessageContentsRequest
	GetOpsUserId() *string
	SetPage(v int32) *ListChatMessageContentsRequest
	GetPage() *int32
	SetPageSize(v int32) *ListChatMessageContentsRequest
	GetPageSize() *int32
	SetSourceId(v string) *ListChatMessageContentsRequest
	GetSourceId() *string
	SetTenantId(v string) *ListChatMessageContentsRequest
	GetTenantId() *string
}

type ListChatMessageContentsRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The maximum length of the attachment content.
	//
	// example:
	//
	// 100000
	MaxCharsPerAttachment *int32 `json:"maxCharsPerAttachment,omitempty" xml:"maxCharsPerAttachment,omitempty"`
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
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// source-chat-1
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListChatMessageContentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatMessageContentsRequest) GoString() string {
	return s.String()
}

func (s *ListChatMessageContentsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListChatMessageContentsRequest) GetMaxCharsPerAttachment() *int32 {
	return s.MaxCharsPerAttachment
}

func (s *ListChatMessageContentsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListChatMessageContentsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListChatMessageContentsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListChatMessageContentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChatMessageContentsRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ListChatMessageContentsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListChatMessageContentsRequest) SetAppToken(v string) *ListChatMessageContentsRequest {
	s.AppToken = &v
	return s
}

func (s *ListChatMessageContentsRequest) SetMaxCharsPerAttachment(v int32) *ListChatMessageContentsRequest {
	s.MaxCharsPerAttachment = &v
	return s
}

func (s *ListChatMessageContentsRequest) SetOpsUserCode(v string) *ListChatMessageContentsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListChatMessageContentsRequest) SetOpsUserId(v string) *ListChatMessageContentsRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListChatMessageContentsRequest) SetPage(v int32) *ListChatMessageContentsRequest {
	s.Page = &v
	return s
}

func (s *ListChatMessageContentsRequest) SetPageSize(v int32) *ListChatMessageContentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListChatMessageContentsRequest) SetSourceId(v string) *ListChatMessageContentsRequest {
	s.SourceId = &v
	return s
}

func (s *ListChatMessageContentsRequest) SetTenantId(v string) *ListChatMessageContentsRequest {
	s.TenantId = &v
	return s
}

func (s *ListChatMessageContentsRequest) Validate() error {
	return dara.Validate(s)
}
