// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppResourcePermissionRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGranteeId(v string) *ListAppResourcePermissionRecordsRequest
	GetGranteeId() *string
	SetGranteeTenantId(v string) *ListAppResourcePermissionRecordsRequest
	GetGranteeTenantId() *string
	SetGranteeType(v string) *ListAppResourcePermissionRecordsRequest
	GetGranteeType() *string
	SetPage(v int64) *ListAppResourcePermissionRecordsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListAppResourcePermissionRecordsRequest
	GetPageSize() *int64
	SetPermissions(v []*string) *ListAppResourcePermissionRecordsRequest
	GetPermissions() []*string
	SetResourceKeyPrefix(v string) *ListAppResourcePermissionRecordsRequest
	GetResourceKeyPrefix() *string
	SetResourceTenantId(v string) *ListAppResourcePermissionRecordsRequest
	GetResourceTenantId() *string
	SetResourceType(v string) *ListAppResourcePermissionRecordsRequest
	GetResourceType() *string
	SetTenantId(v string) *ListAppResourcePermissionRecordsRequest
	GetTenantId() *string
}

type ListAppResourcePermissionRecordsRequest struct {
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
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
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

func (s ListAppResourcePermissionRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppResourcePermissionRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListAppResourcePermissionRecordsRequest) GetGranteeId() *string {
	return s.GranteeId
}

func (s *ListAppResourcePermissionRecordsRequest) GetGranteeTenantId() *string {
	return s.GranteeTenantId
}

func (s *ListAppResourcePermissionRecordsRequest) GetGranteeType() *string {
	return s.GranteeType
}

func (s *ListAppResourcePermissionRecordsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListAppResourcePermissionRecordsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAppResourcePermissionRecordsRequest) GetPermissions() []*string {
	return s.Permissions
}

func (s *ListAppResourcePermissionRecordsRequest) GetResourceKeyPrefix() *string {
	return s.ResourceKeyPrefix
}

func (s *ListAppResourcePermissionRecordsRequest) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *ListAppResourcePermissionRecordsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAppResourcePermissionRecordsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAppResourcePermissionRecordsRequest) SetGranteeId(v string) *ListAppResourcePermissionRecordsRequest {
	s.GranteeId = &v
	return s
}

func (s *ListAppResourcePermissionRecordsRequest) SetGranteeTenantId(v string) *ListAppResourcePermissionRecordsRequest {
	s.GranteeTenantId = &v
	return s
}

func (s *ListAppResourcePermissionRecordsRequest) SetGranteeType(v string) *ListAppResourcePermissionRecordsRequest {
	s.GranteeType = &v
	return s
}

func (s *ListAppResourcePermissionRecordsRequest) SetPage(v int64) *ListAppResourcePermissionRecordsRequest {
	s.Page = &v
	return s
}

func (s *ListAppResourcePermissionRecordsRequest) SetPageSize(v int64) *ListAppResourcePermissionRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppResourcePermissionRecordsRequest) SetPermissions(v []*string) *ListAppResourcePermissionRecordsRequest {
	s.Permissions = v
	return s
}

func (s *ListAppResourcePermissionRecordsRequest) SetResourceKeyPrefix(v string) *ListAppResourcePermissionRecordsRequest {
	s.ResourceKeyPrefix = &v
	return s
}

func (s *ListAppResourcePermissionRecordsRequest) SetResourceTenantId(v string) *ListAppResourcePermissionRecordsRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *ListAppResourcePermissionRecordsRequest) SetResourceType(v string) *ListAppResourcePermissionRecordsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListAppResourcePermissionRecordsRequest) SetTenantId(v string) *ListAppResourcePermissionRecordsRequest {
	s.TenantId = &v
	return s
}

func (s *ListAppResourcePermissionRecordsRequest) Validate() error {
	return dara.Validate(s)
}
