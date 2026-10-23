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
