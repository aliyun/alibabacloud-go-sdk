// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAppResourcePermissionsBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChecks(v []*CheckAppResourcePermissionsBatchRequestChecks) *CheckAppResourcePermissionsBatchRequest
	GetChecks() []*CheckAppResourcePermissionsBatchRequestChecks
	SetResourceTenantId(v string) *CheckAppResourcePermissionsBatchRequest
	GetResourceTenantId() *string
	SetTargetUserId(v int64) *CheckAppResourcePermissionsBatchRequest
	GetTargetUserId() *int64
	SetTargetUserTenantId(v string) *CheckAppResourcePermissionsBatchRequest
	GetTargetUserTenantId() *string
	SetTenantId(v string) *CheckAppResourcePermissionsBatchRequest
	GetTenantId() *string
}

type CheckAppResourcePermissionsBatchRequest struct {
	// The check items.
	//
	// This parameter is required.
	Checks []*CheckAppResourcePermissionsBatchRequestChecks `json:"checks,omitempty" xml:"checks,omitempty" type:"Repeated"`
	// The ID of the tenant to which the resource belongs. If not specified, the current tenant of the caller is used.
	//
	// example:
	//
	// 10000
	ResourceTenantId *string `json:"resourceTenantId,omitempty" xml:"resourceTenantId,omitempty"`
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
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CheckAppResourcePermissionsBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAppResourcePermissionsBatchRequest) GoString() string {
	return s.String()
}

func (s *CheckAppResourcePermissionsBatchRequest) GetChecks() []*CheckAppResourcePermissionsBatchRequestChecks {
	return s.Checks
}

func (s *CheckAppResourcePermissionsBatchRequest) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *CheckAppResourcePermissionsBatchRequest) GetTargetUserId() *int64 {
	return s.TargetUserId
}

func (s *CheckAppResourcePermissionsBatchRequest) GetTargetUserTenantId() *string {
	return s.TargetUserTenantId
}

func (s *CheckAppResourcePermissionsBatchRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CheckAppResourcePermissionsBatchRequest) SetChecks(v []*CheckAppResourcePermissionsBatchRequestChecks) *CheckAppResourcePermissionsBatchRequest {
	s.Checks = v
	return s
}

func (s *CheckAppResourcePermissionsBatchRequest) SetResourceTenantId(v string) *CheckAppResourcePermissionsBatchRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchRequest) SetTargetUserId(v int64) *CheckAppResourcePermissionsBatchRequest {
	s.TargetUserId = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchRequest) SetTargetUserTenantId(v string) *CheckAppResourcePermissionsBatchRequest {
	s.TargetUserTenantId = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchRequest) SetTenantId(v string) *CheckAppResourcePermissionsBatchRequest {
	s.TenantId = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchRequest) Validate() error {
	if s.Checks != nil {
		for _, item := range s.Checks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckAppResourcePermissionsBatchRequestChecks struct {
	// The target permission.
	//
	// This parameter is required.
	//
	// example:
	//
	// APP/VIEW
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// The dot-separated resource key whose first segment is the APP application key.
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
}

func (s CheckAppResourcePermissionsBatchRequestChecks) String() string {
	return dara.Prettify(s)
}

func (s CheckAppResourcePermissionsBatchRequestChecks) GoString() string {
	return s.String()
}

func (s *CheckAppResourcePermissionsBatchRequestChecks) GetPermission() *string {
	return s.Permission
}

func (s *CheckAppResourcePermissionsBatchRequestChecks) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *CheckAppResourcePermissionsBatchRequestChecks) GetResourceType() *string {
	return s.ResourceType
}

func (s *CheckAppResourcePermissionsBatchRequestChecks) SetPermission(v string) *CheckAppResourcePermissionsBatchRequestChecks {
	s.Permission = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchRequestChecks) SetResourceKey(v string) *CheckAppResourcePermissionsBatchRequestChecks {
	s.ResourceKey = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchRequestChecks) SetResourceType(v string) *CheckAppResourcePermissionsBatchRequestChecks {
	s.ResourceType = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchRequestChecks) Validate() error {
	return dara.Validate(s)
}
