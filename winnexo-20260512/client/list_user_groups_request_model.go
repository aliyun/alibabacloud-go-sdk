// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListUserGroupsRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *ListUserGroupsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListUserGroupsRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ListUserGroupsRequest
	GetTenantId() *string
}

type ListUserGroupsRequest struct {
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
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this parameter explicitly by using `--tenant-id`.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListUserGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListUserGroupsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListUserGroupsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListUserGroupsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListUserGroupsRequest) SetAppToken(v string) *ListUserGroupsRequest {
	s.AppToken = &v
	return s
}

func (s *ListUserGroupsRequest) SetOpsUserCode(v string) *ListUserGroupsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListUserGroupsRequest) SetOpsUserId(v string) *ListUserGroupsRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListUserGroupsRequest) SetTenantId(v string) *ListUserGroupsRequest {
	s.TenantId = &v
	return s
}

func (s *ListUserGroupsRequest) Validate() error {
	return dara.Validate(s)
}
