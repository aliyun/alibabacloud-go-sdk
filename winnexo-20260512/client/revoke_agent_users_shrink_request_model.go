// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAgentUsersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *RevokeAgentUsersShrinkRequest
	GetAppToken() *string
	SetOperatingObjectName(v string) *RevokeAgentUsersShrinkRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *RevokeAgentUsersShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *RevokeAgentUsersShrinkRequest
	GetOpsUserId() *string
	SetTenantId(v string) *RevokeAgentUsersShrinkRequest
	GetTenantId() *string
	SetUserGroupIdsShrink(v string) *RevokeAgentUsersShrinkRequest
	GetUserGroupIdsShrink() *string
	SetUserIdsShrink(v string) *RevokeAgentUsersShrinkRequest
	GetUserIdsShrink() *string
}

type RevokeAgentUsersShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
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
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The list of user group IDs to be revoked (16-character hex strings).
	//
	// example:
	//
	// string_value
	UserGroupIdsShrink *string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty"`
	// The list of user IDs to be revoked.
	//
	// example:
	//
	// 1
	UserIdsShrink *string `json:"userIds,omitempty" xml:"userIds,omitempty"`
}

func (s RevokeAgentUsersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeAgentUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevokeAgentUsersShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *RevokeAgentUsersShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *RevokeAgentUsersShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *RevokeAgentUsersShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *RevokeAgentUsersShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RevokeAgentUsersShrinkRequest) GetUserGroupIdsShrink() *string {
	return s.UserGroupIdsShrink
}

func (s *RevokeAgentUsersShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *RevokeAgentUsersShrinkRequest) SetAppToken(v string) *RevokeAgentUsersShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *RevokeAgentUsersShrinkRequest) SetOperatingObjectName(v string) *RevokeAgentUsersShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *RevokeAgentUsersShrinkRequest) SetOpsUserCode(v string) *RevokeAgentUsersShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *RevokeAgentUsersShrinkRequest) SetOpsUserId(v string) *RevokeAgentUsersShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *RevokeAgentUsersShrinkRequest) SetTenantId(v string) *RevokeAgentUsersShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *RevokeAgentUsersShrinkRequest) SetUserGroupIdsShrink(v string) *RevokeAgentUsersShrinkRequest {
	s.UserGroupIdsShrink = &v
	return s
}

func (s *RevokeAgentUsersShrinkRequest) SetUserIdsShrink(v string) *RevokeAgentUsersShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *RevokeAgentUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
