// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantAgentUsersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GrantAgentUsersShrinkRequest
	GetAppToken() *string
	SetExpireDate(v int64) *GrantAgentUsersShrinkRequest
	GetExpireDate() *int64
	SetOperatingObjectName(v string) *GrantAgentUsersShrinkRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *GrantAgentUsersShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GrantAgentUsersShrinkRequest
	GetOpsUserId() *string
	SetPermissionsShrink(v string) *GrantAgentUsersShrinkRequest
	GetPermissionsShrink() *string
	SetTenantId(v string) *GrantAgentUsersShrinkRequest
	GetTenantId() *string
	SetUserGroupIdsShrink(v string) *GrantAgentUsersShrinkRequest
	GetUserGroupIdsShrink() *string
	SetUserIdsShrink(v string) *GrantAgentUsersShrinkRequest
	GetUserIdsShrink() *string
}

type GrantAgentUsersShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The authorization expiration timestamp in milliseconds. If this parameter is not specified, the authorization never expires.
	//
	// example:
	//
	// 1
	ExpireDate *int64 `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// The name of the digital human.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
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
	// The permission items.
	//
	// example:
	//
	// string_value
	PermissionsShrink *string `json:"permissions,omitempty" xml:"permissions,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 676577544219585
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The list of user group IDs.
	//
	// example:
	//
	// string_value
	UserGroupIdsShrink *string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty"`
	// The list of user IDs to be authorized.
	//
	// example:
	//
	// 1
	UserIdsShrink *string `json:"userIds,omitempty" xml:"userIds,omitempty"`
}

func (s GrantAgentUsersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantAgentUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *GrantAgentUsersShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GrantAgentUsersShrinkRequest) GetExpireDate() *int64 {
	return s.ExpireDate
}

func (s *GrantAgentUsersShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *GrantAgentUsersShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GrantAgentUsersShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GrantAgentUsersShrinkRequest) GetPermissionsShrink() *string {
	return s.PermissionsShrink
}

func (s *GrantAgentUsersShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GrantAgentUsersShrinkRequest) GetUserGroupIdsShrink() *string {
	return s.UserGroupIdsShrink
}

func (s *GrantAgentUsersShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *GrantAgentUsersShrinkRequest) SetAppToken(v string) *GrantAgentUsersShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *GrantAgentUsersShrinkRequest) SetExpireDate(v int64) *GrantAgentUsersShrinkRequest {
	s.ExpireDate = &v
	return s
}

func (s *GrantAgentUsersShrinkRequest) SetOperatingObjectName(v string) *GrantAgentUsersShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *GrantAgentUsersShrinkRequest) SetOpsUserCode(v string) *GrantAgentUsersShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GrantAgentUsersShrinkRequest) SetOpsUserId(v string) *GrantAgentUsersShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *GrantAgentUsersShrinkRequest) SetPermissionsShrink(v string) *GrantAgentUsersShrinkRequest {
	s.PermissionsShrink = &v
	return s
}

func (s *GrantAgentUsersShrinkRequest) SetTenantId(v string) *GrantAgentUsersShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *GrantAgentUsersShrinkRequest) SetUserGroupIdsShrink(v string) *GrantAgentUsersShrinkRequest {
	s.UserGroupIdsShrink = &v
	return s
}

func (s *GrantAgentUsersShrinkRequest) SetUserIdsShrink(v string) *GrantAgentUsersShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *GrantAgentUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
