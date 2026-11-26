// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserVisibleKnowledgeBaseContentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListUserVisibleKnowledgeBaseContentsRequest
	GetAppToken() *string
	SetDirectoryId(v string) *ListUserVisibleKnowledgeBaseContentsRequest
	GetDirectoryId() *string
	SetKeyword(v string) *ListUserVisibleKnowledgeBaseContentsRequest
	GetKeyword() *string
	SetOpsUserCode(v string) *ListUserVisibleKnowledgeBaseContentsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListUserVisibleKnowledgeBaseContentsRequest
	GetOpsUserId() *string
	SetPage(v int64) *ListUserVisibleKnowledgeBaseContentsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListUserVisibleKnowledgeBaseContentsRequest
	GetPageSize() *int64
	SetSortField(v string) *ListUserVisibleKnowledgeBaseContentsRequest
	GetSortField() *string
	SetSortOrder(v string) *ListUserVisibleKnowledgeBaseContentsRequest
	GetSortOrder() *string
	SetSourceTypes(v string) *ListUserVisibleKnowledgeBaseContentsRequest
	GetSourceTypes() *string
	SetTenantId(v string) *ListUserVisibleKnowledgeBaseContentsRequest
	GetTenantId() *string
}

type ListUserVisibleKnowledgeBaseContentsRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dir_example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The search keyword. Fuzzy search on version names is supported.
	//
	// example:
	//
	// Product description
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
	// The page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The field by which to sort the returned results. Valid values:
	//
	// - event_time: the event creation time
	//
	// - event_execute_start_time: the event execution start time
	//
	// - event_execute_finish_time: the event execution completion time
	//
	// example:
	//
	// name
	SortField *string `json:"sortField,omitempty" xml:"sortField,omitempty"`
	// Specifies the sort order. This parameter takes effect only when sortBy is specified. Valid values: ASC and DESC. The values are case-insensitive.
	//
	// example:
	//
	// asc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The list of service source types.
	//
	// example:
	//
	// FILE,WEB_PAGE
	SourceTypes *string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty"`
	// The tenant ID. This is a common parameter. For winnexo-cli, it is explicitly passed by using the --tenant-id option.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListUserVisibleKnowledgeBaseContentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserVisibleKnowledgeBaseContentsRequest) GoString() string {
	return s.String()
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetSourceTypes() *string {
	return s.SourceTypes
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetAppToken(v string) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.AppToken = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetDirectoryId(v string) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetKeyword(v string) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.Keyword = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetOpsUserCode(v string) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetOpsUserId(v string) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetPage(v int64) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.Page = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetPageSize(v int64) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetSortField(v string) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.SortField = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetSortOrder(v string) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetSourceTypes(v string) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.SourceTypes = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetTenantId(v string) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.TenantId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) Validate() error {
	return dara.Validate(s)
}
