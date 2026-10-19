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
	// 检查项
	//
	// This parameter is required.
	Checks []*CheckAppResourcePermissionsBatchRequestChecks `json:"checks,omitempty" xml:"checks,omitempty" type:"Repeated"`
	// 资源所属租户 ID；省略为调用人当前租户
	//
	// example:
	//
	// 10000
	ResourceTenantId *string `json:"resourceTenantId,omitempty" xml:"resourceTenantId,omitempty"`
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
	// 目标权限点
	//
	// This parameter is required.
	//
	// example:
	//
	// APP/VIEW
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// 首段为 APP 应用键的点分资源键
	//
	// This parameter is required.
	//
	// example:
	//
	// app.demo.resource_1
	ResourceKey *string `json:"resourceKey,omitempty" xml:"resourceKey,omitempty"`
	// 固定为 APP_RESOURCE
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
