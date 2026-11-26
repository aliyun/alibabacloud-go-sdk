// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAliDingGroupChatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SearchAliDingGroupChatsRequest
	GetAppToken() *string
	SetCursor(v string) *SearchAliDingGroupChatsRequest
	GetCursor() *string
	SetExcludeMuted(v bool) *SearchAliDingGroupChatsRequest
	GetExcludeMuted() *bool
	SetKeyword(v string) *SearchAliDingGroupChatsRequest
	GetKeyword() *string
	SetOpsUserCode(v string) *SearchAliDingGroupChatsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SearchAliDingGroupChatsRequest
	GetOpsUserId() *string
	SetPageSize(v int32) *SearchAliDingGroupChatsRequest
	GetPageSize() *int32
	SetTenantId(v string) *SearchAliDingGroupChatsRequest
	GetTenantId() *string
}

type SearchAliDingGroupChatsRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The cursor for the paged query. Set this parameter to 0 for the first request. For subsequent paging requests, set this parameter to the value of **nextCursor*	- returned in the previous response.
	//
	// example:
	//
	// 0
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// Specifies whether to exclude muted group chats.
	//
	// example:
	//
	// false
	ExcludeMuted *bool `json:"excludeMuted,omitempty" xml:"excludeMuted,omitempty"`
	// The keyword for search. The keyword always matches the name field. If the description field is defined in the schema, the keyword also matches the description field.
	//
	// This parameter is required.
	//
	// example:
	//
	// Customer project
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
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
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SearchAliDingGroupChatsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchAliDingGroupChatsRequest) GoString() string {
	return s.String()
}

func (s *SearchAliDingGroupChatsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SearchAliDingGroupChatsRequest) GetCursor() *string {
	return s.Cursor
}

func (s *SearchAliDingGroupChatsRequest) GetExcludeMuted() *bool {
	return s.ExcludeMuted
}

func (s *SearchAliDingGroupChatsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *SearchAliDingGroupChatsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SearchAliDingGroupChatsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SearchAliDingGroupChatsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchAliDingGroupChatsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SearchAliDingGroupChatsRequest) SetAppToken(v string) *SearchAliDingGroupChatsRequest {
	s.AppToken = &v
	return s
}

func (s *SearchAliDingGroupChatsRequest) SetCursor(v string) *SearchAliDingGroupChatsRequest {
	s.Cursor = &v
	return s
}

func (s *SearchAliDingGroupChatsRequest) SetExcludeMuted(v bool) *SearchAliDingGroupChatsRequest {
	s.ExcludeMuted = &v
	return s
}

func (s *SearchAliDingGroupChatsRequest) SetKeyword(v string) *SearchAliDingGroupChatsRequest {
	s.Keyword = &v
	return s
}

func (s *SearchAliDingGroupChatsRequest) SetOpsUserCode(v string) *SearchAliDingGroupChatsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SearchAliDingGroupChatsRequest) SetOpsUserId(v string) *SearchAliDingGroupChatsRequest {
	s.OpsUserId = &v
	return s
}

func (s *SearchAliDingGroupChatsRequest) SetPageSize(v int32) *SearchAliDingGroupChatsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAliDingGroupChatsRequest) SetTenantId(v string) *SearchAliDingGroupChatsRequest {
	s.TenantId = &v
	return s
}

func (s *SearchAliDingGroupChatsRequest) Validate() error {
	return dara.Validate(s)
}
