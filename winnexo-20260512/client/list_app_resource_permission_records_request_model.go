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
	// The grantee ID filter.
	//
	// example:
	//
	// 10086
	GranteeId *string `json:"granteeId,omitempty" xml:"granteeId,omitempty"`
	// The tenant ID filter for cross-tenant users.
	//
	// example:
	//
	// 10000
	GranteeTenantId *string `json:"granteeTenantId,omitempty" xml:"granteeTenantId,omitempty"`
	// The grantee type filter.
	//
	// example:
	//
	// USER
	GranteeType *string `json:"granteeType,omitempty" xml:"granteeType,omitempty"`
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
	// The permission filter.
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
	// The application key or its subpath.
	//
	// This parameter is required.
	//
	// example:
	//
	// app.demo
	ResourceKeyPrefix *string `json:"resourceKeyPrefix,omitempty" xml:"resourceKeyPrefix,omitempty"`
	// The ID of the tenant to which the resource belongs. If not specified, the current tenant of the caller is used.
	//
	// example:
	//
	// 10000
	ResourceTenantId *string `json:"resourceTenantId,omitempty" xml:"resourceTenantId,omitempty"`
	// The resource type. Set this parameter to APP_RESOURCE.
	//
	// This parameter is required.
	//
	// example:
	//
	// APP_RESOURCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
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
