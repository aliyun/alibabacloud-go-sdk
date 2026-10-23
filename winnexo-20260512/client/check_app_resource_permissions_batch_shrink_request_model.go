// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAppResourcePermissionsBatchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChecksShrink(v string) *CheckAppResourcePermissionsBatchShrinkRequest
	GetChecksShrink() *string
	SetResourceTenantId(v string) *CheckAppResourcePermissionsBatchShrinkRequest
	GetResourceTenantId() *string
	SetTargetUserId(v int64) *CheckAppResourcePermissionsBatchShrinkRequest
	GetTargetUserId() *int64
	SetTargetUserTenantId(v string) *CheckAppResourcePermissionsBatchShrinkRequest
	GetTargetUserTenantId() *string
	SetTenantId(v string) *CheckAppResourcePermissionsBatchShrinkRequest
	GetTenantId() *string
}

type CheckAppResourcePermissionsBatchShrinkRequest struct {
	// The check items.
	//
	// This parameter is required.
	ChecksShrink *string `json:"checks,omitempty" xml:"checks,omitempty"`
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

func (s CheckAppResourcePermissionsBatchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAppResourcePermissionsBatchShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) GetChecksShrink() *string {
	return s.ChecksShrink
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) GetTargetUserId() *int64 {
	return s.TargetUserId
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) GetTargetUserTenantId() *string {
	return s.TargetUserTenantId
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) SetChecksShrink(v string) *CheckAppResourcePermissionsBatchShrinkRequest {
	s.ChecksShrink = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) SetResourceTenantId(v string) *CheckAppResourcePermissionsBatchShrinkRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) SetTargetUserId(v int64) *CheckAppResourcePermissionsBatchShrinkRequest {
	s.TargetUserId = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) SetTargetUserTenantId(v string) *CheckAppResourcePermissionsBatchShrinkRequest {
	s.TargetUserTenantId = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) SetTenantId(v string) *CheckAppResourcePermissionsBatchShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
