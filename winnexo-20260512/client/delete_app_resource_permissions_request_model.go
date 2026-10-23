// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppResourcePermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DeleteAppResourcePermissionsRequestItems) *DeleteAppResourcePermissionsRequest
	GetItems() []*DeleteAppResourcePermissionsRequestItems
	SetResourceTenantId(v string) *DeleteAppResourcePermissionsRequest
	GetResourceTenantId() *string
	SetTenantId(v string) *DeleteAppResourcePermissionsRequest
	GetTenantId() *string
}

type DeleteAppResourcePermissionsRequest struct {
	// The scope of the permissions to be revoked.
	//
	// This parameter is required.
	Items []*DeleteAppResourcePermissionsRequestItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The ID of the tenant to which the resource belongs. If not specified, the current tenant of the caller is used.
	//
	// example:
	//
	// 10000
	ResourceTenantId *string `json:"resourceTenantId,omitempty" xml:"resourceTenantId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteAppResourcePermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppResourcePermissionsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppResourcePermissionsRequest) GetItems() []*DeleteAppResourcePermissionsRequestItems {
	return s.Items
}

func (s *DeleteAppResourcePermissionsRequest) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *DeleteAppResourcePermissionsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteAppResourcePermissionsRequest) SetItems(v []*DeleteAppResourcePermissionsRequestItems) *DeleteAppResourcePermissionsRequest {
	s.Items = v
	return s
}

func (s *DeleteAppResourcePermissionsRequest) SetResourceTenantId(v string) *DeleteAppResourcePermissionsRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *DeleteAppResourcePermissionsRequest) SetTenantId(v string) *DeleteAppResourcePermissionsRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteAppResourcePermissionsRequest) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteAppResourcePermissionsRequestItems struct {
	// The list of grantee principal IDs.
	GranteeIds []*string `json:"granteeIds,omitempty" xml:"granteeIds,omitempty" type:"Repeated"`
	// The ID of the tenant to which the grantee user belongs. This parameter is optional if the user and the resource belong to the same tenant, and is required for cross-tenant users.
	//
	// example:
	//
	// 10000
	GranteeTenantId *string `json:"granteeTenantId,omitempty" xml:"granteeTenantId,omitempty"`
	// The grantee principal type.
	//
	// example:
	//
	// USER
	GranteeType *string `json:"granteeType,omitempty" xml:"granteeType,omitempty"`
	// Specifies whether to include dot-separated descendant resources for the RESOURCE. Default value: false.
	//
	// example:
	//
	// false
	IncludeDescendants *bool `json:"includeDescendants,omitempty" xml:"includeDescendants,omitempty"`
	// The dot-separated resource key whose first segment is the application key of the APP.
	//
	// This parameter is required.
	//
	// example:
	//
	// app.demo.resource_1
	ResourceKey *string `json:"resourceKey,omitempty" xml:"resourceKey,omitempty"`
	// The resource type. Set the value to APP_RESOURCE.
	//
	// This parameter is required.
	//
	// example:
	//
	// APP_RESOURCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The deletion scope: by resource or by principal.
	//
	// This parameter is required.
	//
	// example:
	//
	// RESOURCE
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s DeleteAppResourcePermissionsRequestItems) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppResourcePermissionsRequestItems) GoString() string {
	return s.String()
}

func (s *DeleteAppResourcePermissionsRequestItems) GetGranteeIds() []*string {
	return s.GranteeIds
}

func (s *DeleteAppResourcePermissionsRequestItems) GetGranteeTenantId() *string {
	return s.GranteeTenantId
}

func (s *DeleteAppResourcePermissionsRequestItems) GetGranteeType() *string {
	return s.GranteeType
}

func (s *DeleteAppResourcePermissionsRequestItems) GetIncludeDescendants() *bool {
	return s.IncludeDescendants
}

func (s *DeleteAppResourcePermissionsRequestItems) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *DeleteAppResourcePermissionsRequestItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteAppResourcePermissionsRequestItems) GetScope() *string {
	return s.Scope
}

func (s *DeleteAppResourcePermissionsRequestItems) SetGranteeIds(v []*string) *DeleteAppResourcePermissionsRequestItems {
	s.GranteeIds = v
	return s
}

func (s *DeleteAppResourcePermissionsRequestItems) SetGranteeTenantId(v string) *DeleteAppResourcePermissionsRequestItems {
	s.GranteeTenantId = &v
	return s
}

func (s *DeleteAppResourcePermissionsRequestItems) SetGranteeType(v string) *DeleteAppResourcePermissionsRequestItems {
	s.GranteeType = &v
	return s
}

func (s *DeleteAppResourcePermissionsRequestItems) SetIncludeDescendants(v bool) *DeleteAppResourcePermissionsRequestItems {
	s.IncludeDescendants = &v
	return s
}

func (s *DeleteAppResourcePermissionsRequestItems) SetResourceKey(v string) *DeleteAppResourcePermissionsRequestItems {
	s.ResourceKey = &v
	return s
}

func (s *DeleteAppResourcePermissionsRequestItems) SetResourceType(v string) *DeleteAppResourcePermissionsRequestItems {
	s.ResourceType = &v
	return s
}

func (s *DeleteAppResourcePermissionsRequestItems) SetScope(v string) *DeleteAppResourcePermissionsRequestItems {
	s.Scope = &v
	return s
}

func (s *DeleteAppResourcePermissionsRequestItems) Validate() error {
	return dara.Validate(s)
}
