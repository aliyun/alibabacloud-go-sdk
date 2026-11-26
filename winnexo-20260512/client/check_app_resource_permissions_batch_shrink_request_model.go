// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAppResourcePermissionsBatchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CheckAppResourcePermissionsBatchShrinkRequest
	GetAppToken() *string
	SetChecksShrink(v string) *CheckAppResourcePermissionsBatchShrinkRequest
	GetChecksShrink() *string
	SetOpsUserCode(v string) *CheckAppResourcePermissionsBatchShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CheckAppResourcePermissionsBatchShrinkRequest
	GetOpsUserId() *string
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
	// The platform application token issued by the console. If specified, the operation is performed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The check item.
	//
	// This parameter is required.
	ChecksShrink *string `json:"checks,omitempty" xml:"checks,omitempty"`
	// The proxy user code, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The ID of the target platform user.
	//
	// example:
	//
	// 10000
	ResourceTenantId *string `json:"resourceTenantId,omitempty" xml:"resourceTenantId,omitempty"`
	// The tenant to which the target user belongs. This parameter is required for cross-tenant targets.
	//
	// example:
	//
	// 10086
	TargetUserId *int64 `json:"targetUserId,omitempty" xml:"targetUserId,omitempty"`
	// The check item.
	//
	// example:
	//
	// 10000
	TargetUserTenantId *string `json:"targetUserTenantId,omitempty" xml:"targetUserTenantId,omitempty"`
	// The ID of the tenant to which the resource belongs. If not specified, the current tenant of the caller is used.
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

func (s *CheckAppResourcePermissionsBatchShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) GetChecksShrink() *string {
	return s.ChecksShrink
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *CheckAppResourcePermissionsBatchShrinkRequest) SetAppToken(v string) *CheckAppResourcePermissionsBatchShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) SetChecksShrink(v string) *CheckAppResourcePermissionsBatchShrinkRequest {
	s.ChecksShrink = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) SetOpsUserCode(v string) *CheckAppResourcePermissionsBatchShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchShrinkRequest) SetOpsUserId(v string) *CheckAppResourcePermissionsBatchShrinkRequest {
	s.OpsUserId = &v
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
