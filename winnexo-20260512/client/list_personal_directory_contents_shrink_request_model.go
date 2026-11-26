// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPersonalDirectoryContentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListPersonalDirectoryContentsShrinkRequest
	GetAppToken() *string
	SetDirectoryId(v string) *ListPersonalDirectoryContentsShrinkRequest
	GetDirectoryId() *string
	SetOperatingObjectName(v string) *ListPersonalDirectoryContentsShrinkRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *ListPersonalDirectoryContentsShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListPersonalDirectoryContentsShrinkRequest
	GetOpsUserId() *string
	SetPage(v int64) *ListPersonalDirectoryContentsShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListPersonalDirectoryContentsShrinkRequest
	GetPageSize() *int64
	SetSortField(v string) *ListPersonalDirectoryContentsShrinkRequest
	GetSortField() *string
	SetSortOrder(v string) *ListPersonalDirectoryContentsShrinkRequest
	GetSortOrder() *string
	SetSourceTypesShrink(v string) *ListPersonalDirectoryContentsShrinkRequest
	GetSourceTypesShrink() *string
	SetTenantId(v string) *ListPersonalDirectoryContentsShrinkRequest
	GetTenantId() *string
}

type ListPersonalDirectoryContentsShrinkRequest struct {
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
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The name of the digital employee.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
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
	// The page number. Default value: 1. Pages start from page 1.
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
	// The field by which the returned results are sorted. Valid values:
	//
	// - event_time: the event creation time.
	//
	// - event_execute_start_time: the event execution time.
	//
	// - event_execute_finish_time: the event completion time.
	//
	// example:
	//
	// name
	SortField *string `json:"sortField,omitempty" xml:"sortField,omitempty"`
	// The sort order. This parameter takes effect when sortBy is specified. Valid values: ASC and DESC (case-insensitive).
	//
	// example:
	//
	// asc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The list of service source types.
	//
	// example:
	//
	// string_value
	SourceTypesShrink *string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListPersonalDirectoryContentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalDirectoryContentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetSourceTypesShrink() *string {
	return s.SourceTypesShrink
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetAppToken(v string) *ListPersonalDirectoryContentsShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetDirectoryId(v string) *ListPersonalDirectoryContentsShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetOperatingObjectName(v string) *ListPersonalDirectoryContentsShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetOpsUserCode(v string) *ListPersonalDirectoryContentsShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetOpsUserId(v string) *ListPersonalDirectoryContentsShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetPage(v int64) *ListPersonalDirectoryContentsShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetPageSize(v int64) *ListPersonalDirectoryContentsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetSortField(v string) *ListPersonalDirectoryContentsShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetSortOrder(v string) *ListPersonalDirectoryContentsShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetSourceTypesShrink(v string) *ListPersonalDirectoryContentsShrinkRequest {
	s.SourceTypesShrink = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetTenantId(v string) *ListPersonalDirectoryContentsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
