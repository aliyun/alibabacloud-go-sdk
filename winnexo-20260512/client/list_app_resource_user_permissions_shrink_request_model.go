// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppResourceUserPermissionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int64) *ListAppResourceUserPermissionsShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListAppResourceUserPermissionsShrinkRequest
	GetPageSize() *int64
	SetPermissionsShrink(v string) *ListAppResourceUserPermissionsShrinkRequest
	GetPermissionsShrink() *string
	SetResourceKeyPrefix(v string) *ListAppResourceUserPermissionsShrinkRequest
	GetResourceKeyPrefix() *string
	SetResourceTenantId(v string) *ListAppResourceUserPermissionsShrinkRequest
	GetResourceTenantId() *string
	SetResourceType(v string) *ListAppResourceUserPermissionsShrinkRequest
	GetResourceType() *string
	SetTargetUserId(v int64) *ListAppResourceUserPermissionsShrinkRequest
	GetTargetUserId() *int64
	SetTargetUserTenantId(v string) *ListAppResourceUserPermissionsShrinkRequest
	GetTargetUserTenantId() *string
	SetTenantId(v string) *ListAppResourceUserPermissionsShrinkRequest
	GetTenantId() *string
}

type ListAppResourceUserPermissionsShrinkRequest struct {
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
	// 目标平台用户 ID
	//
	// example:
	//
	// 10086
	TargetUserId *int64 `json:"targetUserId,omitempty" xml:"targetUserId,omitempty"`
	// 目标用户所属租户；跨租户目标必填
	//
	// example:
	//
	// 10000
	TargetUserTenantId *string `json:"targetUserTenantId,omitempty" xml:"targetUserTenantId,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAppResourceUserPermissionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppResourceUserPermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAppResourceUserPermissionsShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListAppResourceUserPermissionsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAppResourceUserPermissionsShrinkRequest) GetPermissionsShrink() *string {
	return s.PermissionsShrink
}

func (s *ListAppResourceUserPermissionsShrinkRequest) GetResourceKeyPrefix() *string {
	return s.ResourceKeyPrefix
}

func (s *ListAppResourceUserPermissionsShrinkRequest) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *ListAppResourceUserPermissionsShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAppResourceUserPermissionsShrinkRequest) GetTargetUserId() *int64 {
	return s.TargetUserId
}

func (s *ListAppResourceUserPermissionsShrinkRequest) GetTargetUserTenantId() *string {
	return s.TargetUserTenantId
}

func (s *ListAppResourceUserPermissionsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAppResourceUserPermissionsShrinkRequest) SetPage(v int64) *ListAppResourceUserPermissionsShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListAppResourceUserPermissionsShrinkRequest) SetPageSize(v int64) *ListAppResourceUserPermissionsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppResourceUserPermissionsShrinkRequest) SetPermissionsShrink(v string) *ListAppResourceUserPermissionsShrinkRequest {
	s.PermissionsShrink = &v
	return s
}

func (s *ListAppResourceUserPermissionsShrinkRequest) SetResourceKeyPrefix(v string) *ListAppResourceUserPermissionsShrinkRequest {
	s.ResourceKeyPrefix = &v
	return s
}

func (s *ListAppResourceUserPermissionsShrinkRequest) SetResourceTenantId(v string) *ListAppResourceUserPermissionsShrinkRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *ListAppResourceUserPermissionsShrinkRequest) SetResourceType(v string) *ListAppResourceUserPermissionsShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListAppResourceUserPermissionsShrinkRequest) SetTargetUserId(v int64) *ListAppResourceUserPermissionsShrinkRequest {
	s.TargetUserId = &v
	return s
}

func (s *ListAppResourceUserPermissionsShrinkRequest) SetTargetUserTenantId(v string) *ListAppResourceUserPermissionsShrinkRequest {
	s.TargetUserTenantId = &v
	return s
}

func (s *ListAppResourceUserPermissionsShrinkRequest) SetTenantId(v string) *ListAppResourceUserPermissionsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListAppResourceUserPermissionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
