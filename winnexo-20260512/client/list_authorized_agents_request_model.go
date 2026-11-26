// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListAuthorizedAgentsRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *ListAuthorizedAgentsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListAuthorizedAgentsRequest
	GetOpsUserId() *string
	SetPermission(v string) *ListAuthorizedAgentsRequest
	GetPermission() *string
	SetTargetUserId(v int64) *ListAuthorizedAgentsRequest
	GetTargetUserId() *int64
	SetTenantId(v string) *ListAuthorizedAgentsRequest
	GetTenantId() *string
}

type ListAuthorizedAgentsRequest struct {
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
	// The userId of the responsible user.
	//
	// example:
	//
	// USE
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// The target user ID.
	//
	// example:
	//
	// 1
	TargetUserId *int64 `json:"targetUserId,omitempty" xml:"targetUserId,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAuthorizedAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAgentsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListAuthorizedAgentsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListAuthorizedAgentsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListAuthorizedAgentsRequest) GetPermission() *string {
	return s.Permission
}

func (s *ListAuthorizedAgentsRequest) GetTargetUserId() *int64 {
	return s.TargetUserId
}

func (s *ListAuthorizedAgentsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAuthorizedAgentsRequest) SetAppToken(v string) *ListAuthorizedAgentsRequest {
	s.AppToken = &v
	return s
}

func (s *ListAuthorizedAgentsRequest) SetOpsUserCode(v string) *ListAuthorizedAgentsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListAuthorizedAgentsRequest) SetOpsUserId(v string) *ListAuthorizedAgentsRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListAuthorizedAgentsRequest) SetPermission(v string) *ListAuthorizedAgentsRequest {
	s.Permission = &v
	return s
}

func (s *ListAuthorizedAgentsRequest) SetTargetUserId(v int64) *ListAuthorizedAgentsRequest {
	s.TargetUserId = &v
	return s
}

func (s *ListAuthorizedAgentsRequest) SetTenantId(v string) *ListAuthorizedAgentsRequest {
	s.TenantId = &v
	return s
}

func (s *ListAuthorizedAgentsRequest) Validate() error {
	return dara.Validate(s)
}
