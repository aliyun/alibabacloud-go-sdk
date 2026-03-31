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
	SetTenantId(v string) *ListProjectsRequest
	GetTenantId() *string
	SetType(v string) *ListProjectsRequest
	GetType() *string
}

type ListProjectsRequest struct {
	// Specifies whether to list the built-in **SYSTEM_CATALOG*	- projects that are used to provide data such as project metadata and historical usage data. For more information, see [Tenant-level Information Schema](https://www.alibabacloud.com/help/zh/maxcompute/user-guide/tenant-level-information-schema).
	//
	// Valid values:
	//
	// 	- true: The built-in SYSTEM_CATALOG projects are listed.
	//
	// 	- false: The built-in SYSTEM_CATALOG projects are not listed.
	//
	// example:
	//
	// true
	ListSystemCatalog *bool `json:"listSystemCatalog,omitempty" xml:"listSystemCatalog,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// Specifies the marker after which the returned list begins.
	//
	// example:
	//
	// a
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// The quota name that is automatically generated. You can log on to the [MaxCompute console](https://maxcompute.console.aliyun.com), choose **Workspace*	- > **Quotas*	- from the left-side navigation pane, and then view the quota name on the **Quotas*	- page.
	//
	// example:
	//
	// "hsajkdgbkaubh"
	QuotaName *string `json:"quotaName,omitempty" xml:"quotaName,omitempty"`
	// The quota nickname. You can log on to the [MaxCompute console](https://maxcompute.console.aliyun.com), choose **Workspace*	- > **Quotas*	- from the left-side navigation pane, and then view the quota nickname on the **Quotas*	- page.
	//
	// example:
	//
	// quotaA
	QuotaNickName *string `json:"quotaNickName,omitempty" xml:"quotaNickName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The instance ID and billing method of the default computing quota.
	//
	// example:
	//
	// "aaaa-bbbb"
	SaleTags *string `json:"saleTags,omitempty" xml:"saleTags,omitempty"`
	// The tenant ID. You can log on to the [MaxCompute console](https://maxcompute.console.aliyun.com), and choose **Tenants*	- > **Tenant Property*	- from the left-side navigation pane to view the tenant ID.
	//
	// example:
	//
	// 549532154333697
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The project type. Valid values:
	//
	// 	- **managed**: internal project
	//
	// 	- **external**: external project
	//
	// example:
	//
	// "managed"
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
