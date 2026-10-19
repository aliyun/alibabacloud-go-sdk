// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppResourcePermissionRecordsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGranteeId(v string) *ListAppResourcePermissionRecordsShrinkRequest
	GetGranteeId() *string
	SetGranteeTenantId(v string) *ListAppResourcePermissionRecordsShrinkRequest
	GetGranteeTenantId() *string
	SetGranteeType(v string) *ListAppResourcePermissionRecordsShrinkRequest
	GetGranteeType() *string
	SetPage(v int64) *ListAppResourcePermissionRecordsShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListAppResourcePermissionRecordsShrinkRequest
	GetPageSize() *int64
	SetPermissionsShrink(v string) *ListAppResourcePermissionRecordsShrinkRequest
	GetPermissionsShrink() *string
	SetResourceKeyPrefix(v string) *ListAppResourcePermissionRecordsShrinkRequest
	GetResourceKeyPrefix() *string
	SetResourceTenantId(v string) *ListAppResourcePermissionRecordsShrinkRequest
	GetResourceTenantId() *string
	SetResourceType(v string) *ListAppResourcePermissionRecordsShrinkRequest
	GetResourceType() *string
	SetTenantId(v string) *ListAppResourcePermissionRecordsShrinkRequest
	GetTenantId() *string
}

type ListAppResourcePermissionRecordsShrinkRequest struct {
	// 主体 ID 筛选
	//
	// example:
	//
	// 10086
	GranteeId *string `json:"granteeId,omitempty" xml:"granteeId,omitempty"`
	// 跨租户 USER 所属租户筛选
	//
	// example:
	//
	// 10000
	GranteeTenantId *string `json:"granteeTenantId,omitempty" xml:"granteeTenantId,omitempty"`
	// 主体类型筛选
	//
	// example:
	//
	// USER
	GranteeType *string `json:"granteeType,omitempty" xml:"granteeType,omitempty"`
	// 页码
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页条数
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 权限点筛选
	PermissionsShrink *string `json:"permissions,omitempty" xml:"permissions,omitempty"`
	// 应用键或其子路径
	//
	// This parameter is required.
	//
	// example:
	//
	// app.demo
	ResourceKeyPrefix *string `json:"resourceKeyPrefix,omitempty" xml:"resourceKeyPrefix,omitempty"`
	// 资源所属租户 ID；省略为调用人当前租户
	//
	// example:
	//
	// 10000
	ResourceTenantId *string `json:"resourceTenantId,omitempty" xml:"resourceTenantId,omitempty"`
	// 固定为 APP_RESOURCE
	//
	// This parameter is required.
	//
	// example:
	//
	// APP_RESOURCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAppResourcePermissionRecordsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppResourcePermissionRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) GetGranteeId() *string {
	return s.GranteeId
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) GetGranteeTenantId() *string {
	return s.GranteeTenantId
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) GetGranteeType() *string {
	return s.GranteeType
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) GetPermissionsShrink() *string {
	return s.PermissionsShrink
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) GetResourceKeyPrefix() *string {
	return s.ResourceKeyPrefix
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) SetGranteeId(v string) *ListAppResourcePermissionRecordsShrinkRequest {
	s.GranteeId = &v
	return s
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) SetGranteeTenantId(v string) *ListAppResourcePermissionRecordsShrinkRequest {
	s.GranteeTenantId = &v
	return s
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) SetGranteeType(v string) *ListAppResourcePermissionRecordsShrinkRequest {
	s.GranteeType = &v
	return s
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) SetPage(v int64) *ListAppResourcePermissionRecordsShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) SetPageSize(v int64) *ListAppResourcePermissionRecordsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) SetPermissionsShrink(v string) *ListAppResourcePermissionRecordsShrinkRequest {
	s.PermissionsShrink = &v
	return s
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) SetResourceKeyPrefix(v string) *ListAppResourcePermissionRecordsShrinkRequest {
	s.ResourceKeyPrefix = &v
	return s
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) SetResourceTenantId(v string) *ListAppResourcePermissionRecordsShrinkRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) SetResourceType(v string) *ListAppResourcePermissionRecordsShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) SetTenantId(v string) *ListAppResourcePermissionRecordsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListAppResourcePermissionRecordsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
