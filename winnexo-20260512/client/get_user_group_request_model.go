// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetUserGroupRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *GetUserGroupRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetUserGroupRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetUserGroupRequest
	GetTenantId() *string
	SetUserGroupId(v string) *GetUserGroupRequest
	GetUserGroupId() *string
}

type GetUserGroupRequest struct {
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
	// The ID of the target user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7ea8973f-7a5c-4e8a-956b-4fe0e7e2eb11
	UserGroupId *string `json:"userGroupId,omitempty" xml:"userGroupId,omitempty"`
}

func (s GetUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupRequest) GoString() string {
	return s.String()
}

func (s *GetUserGroupRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetUserGroupRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetUserGroupRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetUserGroupRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *GetUserGroupRequest) SetAppToken(v string) *GetUserGroupRequest {
	s.AppToken = &v
	return s
}

func (s *GetUserGroupRequest) SetOpsUserCode(v string) *GetUserGroupRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetUserGroupRequest) SetOpsUserId(v string) *GetUserGroupRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetUserGroupRequest) SetTenantId(v string) *GetUserGroupRequest {
	s.TenantId = &v
	return s
}

func (s *GetUserGroupRequest) SetUserGroupId(v string) *GetUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *GetUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
