// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterAppResourcePermissionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *RegisterAppResourcePermissionsShrinkRequest
	GetAppToken() *string
	SetItemsShrink(v string) *RegisterAppResourcePermissionsShrinkRequest
	GetItemsShrink() *string
	SetOpsUserCode(v string) *RegisterAppResourcePermissionsShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *RegisterAppResourcePermissionsShrinkRequest
	GetOpsUserId() *string
	SetResourceTenantId(v string) *RegisterAppResourcePermissionsShrinkRequest
	GetResourceTenantId() *string
	SetTenantId(v string) *RegisterAppResourcePermissionsShrinkRequest
	GetTenantId() *string
}

type RegisterAppResourcePermissionsShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The batch authorization item.
	//
	// This parameter is required.
	ItemsShrink *string `json:"items,omitempty" xml:"items,omitempty"`
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
	// The batch authorization item.
	//
	// example:
	//
	// 10000
	ResourceTenantId *string `json:"resourceTenantId,omitempty" xml:"resourceTenantId,omitempty"`
	// The ID of the tenant to which the resource belongs. If omitted, the current tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RegisterAppResourcePermissionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterAppResourcePermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *RegisterAppResourcePermissionsShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *RegisterAppResourcePermissionsShrinkRequest) GetItemsShrink() *string {
	return s.ItemsShrink
}

func (s *RegisterAppResourcePermissionsShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *RegisterAppResourcePermissionsShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *RegisterAppResourcePermissionsShrinkRequest) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *RegisterAppResourcePermissionsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RegisterAppResourcePermissionsShrinkRequest) SetAppToken(v string) *RegisterAppResourcePermissionsShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *RegisterAppResourcePermissionsShrinkRequest) SetItemsShrink(v string) *RegisterAppResourcePermissionsShrinkRequest {
	s.ItemsShrink = &v
	return s
}

func (s *RegisterAppResourcePermissionsShrinkRequest) SetOpsUserCode(v string) *RegisterAppResourcePermissionsShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *RegisterAppResourcePermissionsShrinkRequest) SetOpsUserId(v string) *RegisterAppResourcePermissionsShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *RegisterAppResourcePermissionsShrinkRequest) SetResourceTenantId(v string) *RegisterAppResourcePermissionsShrinkRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *RegisterAppResourcePermissionsShrinkRequest) SetTenantId(v string) *RegisterAppResourcePermissionsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *RegisterAppResourcePermissionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
