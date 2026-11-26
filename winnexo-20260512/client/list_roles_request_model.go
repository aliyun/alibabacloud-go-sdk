// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListRolesRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *ListRolesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListRolesRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ListRolesRequest
	GetTenantId() *string
}

type ListRolesRequest struct {
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
	// The tenant ID. This is a common parameter. In winnexo-cli, pass it explicitly with --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListRolesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListRolesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListRolesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListRolesRequest) SetAppToken(v string) *ListRolesRequest {
	s.AppToken = &v
	return s
}

func (s *ListRolesRequest) SetOpsUserCode(v string) *ListRolesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListRolesRequest) SetOpsUserId(v string) *ListRolesRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListRolesRequest) SetTenantId(v string) *ListRolesRequest {
	s.TenantId = &v
	return s
}

func (s *ListRolesRequest) Validate() error {
	return dara.Validate(s)
}
