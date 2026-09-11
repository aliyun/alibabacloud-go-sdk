// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAliDingGroupChatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCursor(v string) *SearchAliDingGroupChatsRequest
	GetCursor() *string
	SetExcludeMuted(v bool) *SearchAliDingGroupChatsRequest
	GetExcludeMuted() *bool
	SetKeyword(v string) *SearchAliDingGroupChatsRequest
	GetKeyword() *string
	SetPageSize(v int32) *SearchAliDingGroupChatsRequest
	GetPageSize() *int32
	SetTenantId(v string) *SearchAliDingGroupChatsRequest
	GetTenantId() *string
}

type SearchAliDingGroupChatsRequest struct {
	// 分页游标，首页传 0
	//
	// example:
	//
	// 0
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// 是否排除免打扰群聊
	//
	// example:
	//
	// false
	ExcludeMuted *bool `json:"excludeMuted,omitempty" xml:"excludeMuted,omitempty"`
	// 群聊搜索关键词
	//
	// This parameter is required.
	//
	// example:
	//
	// 客户项目
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 每页条数，范围 1-100
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 租户 ID，公共参数；缺省时使用调用方默认租户
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

func (s *SearchAliDingGroupChatsRequest) GetCursor() *string {
	return s.Cursor
}

func (s *SearchAliDingGroupChatsRequest) GetExcludeMuted() *bool {
	return s.ExcludeMuted
}

func (s *SearchAliDingGroupChatsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *SearchAliDingGroupChatsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchAliDingGroupChatsRequest) GetTenantId() *string {
	return s.TenantId
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
