// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppResourceUserPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int64) *ListAppResourceUserPermissionsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListAppResourceUserPermissionsRequest
	GetPageSize() *int64
	SetPermissions(v []*string) *ListAppResourceUserPermissionsRequest
	GetPermissions() []*string
	SetResourceKeyPrefix(v string) *ListAppResourceUserPermissionsRequest
	GetResourceKeyPrefix() *string
	SetResourceTenantId(v string) *ListAppResourceUserPermissionsRequest
	GetResourceTenantId() *string
	SetResourceType(v string) *ListAppResourceUserPermissionsRequest
	GetResourceType() *string
	SetTargetUserId(v int64) *ListAppResourceUserPermissionsRequest
	GetTargetUserId() *int64
	SetTargetUserTenantId(v string) *ListAppResourceUserPermissionsRequest
	GetTargetUserTenantId() *string
	SetTenantId(v string) *ListAppResourceUserPermissionsRequest
	GetTenantId() *string
}

type ListAppResourceUserPermissionsRequest struct {
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
	// The permission points to filter.
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
	// The application key or its subpath.
	//
	// This parameter is required.
	//
	// example:
	//
	// app.demo
	ResourceKeyPrefix *string `json:"resourceKeyPrefix,omitempty" xml:"resourceKeyPrefix,omitempty"`
	// The ID of the tenant to which the resource belongs. If you do not specify this parameter, the current tenant of the caller is used.
	//
	// example:
	//
	// 10000
	ResourceTenantId *string `json:"resourceTenantId,omitempty" xml:"resourceTenantId,omitempty"`
	// The resource type. Set the value to APP_RESOURCE.
	//
	// This parameter is required.
	//
	// example:
	//
	// APP_RESOURCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The ID of the target platform user.
	//
	// example:
	//
	// 10086
	TargetUserId *int64 `json:"targetUserId,omitempty" xml:"targetUserId,omitempty"`
	// The tenant to which the target user belongs. This parameter is required for cross-tenant targets.
	//
	// example:
	//
	// 10000
	TargetUserTenantId *string `json:"targetUserTenantId,omitempty" xml:"targetUserTenantId,omitempty"`
	// The tenant ID. This is a common parameter. If you do not specify this parameter, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAppResourceUserPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppResourceUserPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListAppResourceUserPermissionsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListAppResourceUserPermissionsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAppResourceUserPermissionsRequest) GetPermissions() []*string {
	return s.Permissions
}

func (s *ListAppResourceUserPermissionsRequest) GetResourceKeyPrefix() *string {
	return s.ResourceKeyPrefix
}

func (s *ListAppResourceUserPermissionsRequest) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *ListAppResourceUserPermissionsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAppResourceUserPermissionsRequest) GetTargetUserId() *int64 {
	return s.TargetUserId
}

func (s *ListAppResourceUserPermissionsRequest) GetTargetUserTenantId() *string {
	return s.TargetUserTenantId
}

func (s *ListAppResourceUserPermissionsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAppResourceUserPermissionsRequest) SetPage(v int64) *ListAppResourceUserPermissionsRequest {
	s.Page = &v
	return s
}

func (s *ListAppResourceUserPermissionsRequest) SetPageSize(v int64) *ListAppResourceUserPermissionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppResourceUserPermissionsRequest) SetPermissions(v []*string) *ListAppResourceUserPermissionsRequest {
	s.Permissions = v
	return s
}

func (s *ListAppResourceUserPermissionsRequest) SetResourceKeyPrefix(v string) *ListAppResourceUserPermissionsRequest {
	s.ResourceKeyPrefix = &v
	return s
}

func (s *ListAppResourceUserPermissionsRequest) SetResourceTenantId(v string) *ListAppResourceUserPermissionsRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *ListAppResourceUserPermissionsRequest) SetResourceType(v string) *ListAppResourceUserPermissionsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListAppResourceUserPermissionsRequest) SetTargetUserId(v int64) *ListAppResourceUserPermissionsRequest {
	s.TargetUserId = &v
	return s
}

func (s *ListAppResourceUserPermissionsRequest) SetTargetUserTenantId(v string) *ListAppResourceUserPermissionsRequest {
	s.TargetUserTenantId = &v
	return s
}

func (s *ListAppResourceUserPermissionsRequest) SetTenantId(v string) *ListAppResourceUserPermissionsRequest {
	s.TenantId = &v
	return s
}

func (s *ListAppResourceUserPermissionsRequest) Validate() error {
	return dara.Validate(s)
}
