// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAppResourcePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CheckAppResourcePermissionRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *CheckAppResourcePermissionRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CheckAppResourcePermissionRequest
	GetOpsUserId() *string
	SetPermission(v string) *CheckAppResourcePermissionRequest
	GetPermission() *string
	SetResourceKey(v string) *CheckAppResourcePermissionRequest
	GetResourceKey() *string
	SetResourceTenantId(v string) *CheckAppResourcePermissionRequest
	GetResourceTenantId() *string
	SetResourceType(v string) *CheckAppResourcePermissionRequest
	GetResourceType() *string
	SetTargetUserId(v int64) *CheckAppResourcePermissionRequest
	GetTargetUserId() *int64
	SetTargetUserTenantId(v string) *CheckAppResourcePermissionRequest
	GetTargetUserTenantId() *string
	SetTenantId(v string) *CheckAppResourcePermissionRequest
	GetTenantId() *string
}

type CheckAppResourcePermissionRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The permission to check.
	//
	// This parameter is required.
	//
	// example:
	//
	// APP/VIEW
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// The dot-separated resource key. The first segment is the application key.
	//
	// This parameter is required.
	//
	// example:
	//
	// app.demo.resource_1
	ResourceKey *string `json:"resourceKey,omitempty" xml:"resourceKey,omitempty"`
	// The ID of the tenant to which the resource belongs. If not specified, the current tenant of the caller is used.
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
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CheckAppResourcePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAppResourcePermissionRequest) GoString() string {
	return s.String()
}

func (s *CheckAppResourcePermissionRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CheckAppResourcePermissionRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CheckAppResourcePermissionRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CheckAppResourcePermissionRequest) GetPermission() *string {
	return s.Permission
}

func (s *CheckAppResourcePermissionRequest) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *CheckAppResourcePermissionRequest) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *CheckAppResourcePermissionRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CheckAppResourcePermissionRequest) GetTargetUserId() *int64 {
	return s.TargetUserId
}

func (s *CheckAppResourcePermissionRequest) GetTargetUserTenantId() *string {
	return s.TargetUserTenantId
}

func (s *CheckAppResourcePermissionRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CheckAppResourcePermissionRequest) SetAppToken(v string) *CheckAppResourcePermissionRequest {
	s.AppToken = &v
	return s
}

func (s *CheckAppResourcePermissionRequest) SetOpsUserCode(v string) *CheckAppResourcePermissionRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CheckAppResourcePermissionRequest) SetOpsUserId(v string) *CheckAppResourcePermissionRequest {
	s.OpsUserId = &v
	return s
}

func (s *CheckAppResourcePermissionRequest) SetPermission(v string) *CheckAppResourcePermissionRequest {
	s.Permission = &v
	return s
}

func (s *CheckAppResourcePermissionRequest) SetResourceKey(v string) *CheckAppResourcePermissionRequest {
	s.ResourceKey = &v
	return s
}

func (s *CheckAppResourcePermissionRequest) SetResourceTenantId(v string) *CheckAppResourcePermissionRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *CheckAppResourcePermissionRequest) SetResourceType(v string) *CheckAppResourcePermissionRequest {
	s.ResourceType = &v
	return s
}

func (s *CheckAppResourcePermissionRequest) SetTargetUserId(v int64) *CheckAppResourcePermissionRequest {
	s.TargetUserId = &v
	return s
}

func (s *CheckAppResourcePermissionRequest) SetTargetUserTenantId(v string) *CheckAppResourcePermissionRequest {
	s.TargetUserTenantId = &v
	return s
}

func (s *CheckAppResourcePermissionRequest) SetTenantId(v string) *CheckAppResourcePermissionRequest {
	s.TenantId = &v
	return s
}

func (s *CheckAppResourcePermissionRequest) Validate() error {
	return dara.Validate(s)
}
