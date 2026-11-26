// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppResourcePermissionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *DeleteAppResourcePermissionsShrinkRequest
	GetAppToken() *string
	SetItemsShrink(v string) *DeleteAppResourcePermissionsShrinkRequest
	GetItemsShrink() *string
	SetOpsUserCode(v string) *DeleteAppResourcePermissionsShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *DeleteAppResourcePermissionsShrinkRequest
	GetOpsUserId() *string
	SetResourceTenantId(v string) *DeleteAppResourcePermissionsShrinkRequest
	GetResourceTenantId() *string
	SetTenantId(v string) *DeleteAppResourcePermissionsShrinkRequest
	GetTenantId() *string
}

type DeleteAppResourcePermissionsShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The scope of permission revocation.
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
	// The scope of permission revocation.
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

func (s DeleteAppResourcePermissionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppResourcePermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppResourcePermissionsShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *DeleteAppResourcePermissionsShrinkRequest) GetItemsShrink() *string {
	return s.ItemsShrink
}

func (s *DeleteAppResourcePermissionsShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *DeleteAppResourcePermissionsShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *DeleteAppResourcePermissionsShrinkRequest) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *DeleteAppResourcePermissionsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteAppResourcePermissionsShrinkRequest) SetAppToken(v string) *DeleteAppResourcePermissionsShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *DeleteAppResourcePermissionsShrinkRequest) SetItemsShrink(v string) *DeleteAppResourcePermissionsShrinkRequest {
	s.ItemsShrink = &v
	return s
}

func (s *DeleteAppResourcePermissionsShrinkRequest) SetOpsUserCode(v string) *DeleteAppResourcePermissionsShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *DeleteAppResourcePermissionsShrinkRequest) SetOpsUserId(v string) *DeleteAppResourcePermissionsShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *DeleteAppResourcePermissionsShrinkRequest) SetResourceTenantId(v string) *DeleteAppResourcePermissionsShrinkRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *DeleteAppResourcePermissionsShrinkRequest) SetTenantId(v string) *DeleteAppResourcePermissionsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteAppResourcePermissionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
