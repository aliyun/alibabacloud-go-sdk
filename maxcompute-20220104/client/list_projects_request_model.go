// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListSystemCatalog(v bool) *ListProjectsRequest
	GetListSystemCatalog() *bool
	SetMarker(v string) *ListProjectsRequest
	GetMarker() *string
	SetMaxItem(v int32) *ListProjectsRequest
	GetMaxItem() *int32
	SetPrefix(v string) *ListProjectsRequest
	GetPrefix() *string
	SetQuotaName(v string) *ListProjectsRequest
	GetQuotaName() *string
	SetQuotaNickName(v string) *ListProjectsRequest
	GetQuotaNickName() *string
	SetRegion(v string) *ListProjectsRequest
	GetRegion() *string
	SetSaleTags(v string) *ListProjectsRequest
	GetSaleTags() *string
	SetSortBy(v string) *ListProjectsRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListProjectsRequest
	GetSortOrder() *string
	SetTenantId(v string) *ListProjectsRequest
	GetTenantId() *string
	SetType(v string) *ListProjectsRequest
	GetType() *string
}

type ListProjectsRequest struct {
	// Specifies whether to list the built-in **SYSTEM_CATALOG*	- project. This project provides information such as project metadata and usage history. For more information, see <props="intl">[Information Schema](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/tenant-level-information-schema).
	//
	// example:
	//
	// true
	ListSystemCatalog *bool `json:"listSystemCatalog,omitempty" xml:"listSystemCatalog,omitempty"`
	// The token that specifies the starting point of the query. The results are returned in alphabetical order, starting from the entry that immediately follows the marker.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The prefix of the resource names to query. For example, if you specify `a` for this parameter, only resources whose names start with "a" are returned.
	//
	// example:
	//
	// a
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// The quota name. The system automatically generates this name. To obtain the quota name, log in to the [MaxCompute console](https://maxcompute.console.aliyun.com) and select **Workspace*	- > **Quota*	- **Management*	- from the navigation pane on the left.
	//
	// example:
	//
	// aliyun_5495***3697
	QuotaName *string `json:"quotaName,omitempty" xml:"quotaName,omitempty"`
	// The quota nickname. To obtain the quota nickname, log in to the [MaxCompute console](https://maxcompute.console.aliyun.com) and select **Workspace*	- > **Quota*	- **Management*	- from the navigation pane on the left.
	//
	// example:
	//
	// os_PayAsYouGoQuota
	QuotaNickName *string `json:"quotaNickName,omitempty" xml:"quotaNickName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The instance ID and billing method of the default compute quota.
	//
	// example:
	//
	// {
	//
	//       "resourceId": "b7afb7d1-****-****-****-c393669c307b",
	//
	//       "resourceType": "PayAsYouGo"
	//
	//     }
	SaleTags *string `json:"saleTags,omitempty" xml:"saleTags,omitempty"`
	// The sort field. The only supported value is `createdTime`.
	//
	// example:
	//
	// createdTime
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// The sort order. This parameter takes effect only when `sortBy` is specified. Valid values are `ASC` and `DESC`. The values are case-insensitive.
	//
	// example:
	//
	// DESC
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The tenant ID. To obtain the ID, log in to the [MaxCompute console](https://maxcompute.console.aliyun.com) and select **Tenant Management*	- > **Tenant Properties*	- from the navigation pane on the left.
	//
	// example:
	//
	// 5495****3697
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The project type. Valid values:
	//
	// - **managed**: a managed project.
	//
	// - **external**: an external project.
	//
	// example:
	//
	// managed
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) GetListSystemCatalog() *bool {
	return s.ListSystemCatalog
}

func (s *ListProjectsRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListProjectsRequest) GetMaxItem() *int32 {
	return s.MaxItem
}

func (s *ListProjectsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListProjectsRequest) GetQuotaName() *string {
	return s.QuotaName
}

func (s *ListProjectsRequest) GetQuotaNickName() *string {
	return s.QuotaNickName
}

func (s *ListProjectsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListProjectsRequest) GetSaleTags() *string {
	return s.SaleTags
}

func (s *ListProjectsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListProjectsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListProjectsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListProjectsRequest) GetType() *string {
	return s.Type
}

func (s *ListProjectsRequest) SetListSystemCatalog(v bool) *ListProjectsRequest {
	s.ListSystemCatalog = &v
	return s
}

func (s *ListProjectsRequest) SetMarker(v string) *ListProjectsRequest {
	s.Marker = &v
	return s
}

func (s *ListProjectsRequest) SetMaxItem(v int32) *ListProjectsRequest {
	s.MaxItem = &v
	return s
}

func (s *ListProjectsRequest) SetPrefix(v string) *ListProjectsRequest {
	s.Prefix = &v
	return s
}

func (s *ListProjectsRequest) SetQuotaName(v string) *ListProjectsRequest {
	s.QuotaName = &v
	return s
}

func (s *ListProjectsRequest) SetQuotaNickName(v string) *ListProjectsRequest {
	s.QuotaNickName = &v
	return s
}

func (s *ListProjectsRequest) SetRegion(v string) *ListProjectsRequest {
	s.Region = &v
	return s
}

func (s *ListProjectsRequest) SetSaleTags(v string) *ListProjectsRequest {
	s.SaleTags = &v
	return s
}

func (s *ListProjectsRequest) SetSortBy(v string) *ListProjectsRequest {
	s.SortBy = &v
	return s
}

func (s *ListProjectsRequest) SetSortOrder(v string) *ListProjectsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListProjectsRequest) SetTenantId(v string) *ListProjectsRequest {
	s.TenantId = &v
	return s
}

func (s *ListProjectsRequest) SetType(v string) *ListProjectsRequest {
	s.Type = &v
	return s
}

func (s *ListProjectsRequest) Validate() error {
	return dara.Validate(s)
}
