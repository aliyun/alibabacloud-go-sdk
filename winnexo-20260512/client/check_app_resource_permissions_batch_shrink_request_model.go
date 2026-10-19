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
	// 检查项
	//
	// This parameter is required.
	ChecksShrink *string `json:"checks,omitempty" xml:"checks,omitempty"`
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
