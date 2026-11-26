// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetUserRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *GetUserRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetUserRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetUserRequest
	GetTenantId() *string
	SetWnAccountId(v string) *GetUserRequest
	GetWnAccountId() *string
	SetWnUserId(v string) *GetUserRequest
	GetWnUserId() *string
}

type GetUserRequest struct {
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
	// The ID of the tenant to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The WINNEXO logon account. This parameter is a unique identifier and cannot be left empty.
	//
	// example:
	//
	// exampleAccountId
	WnAccountId *string `json:"wnAccountId,omitempty" xml:"wnAccountId,omitempty"`
	// The user ID on the WINNEXO platform. Specify either this parameter or accountId.
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s GetUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetUserRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetUserRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetUserRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserRequest) GetWnAccountId() *string {
	return s.WnAccountId
}

func (s *GetUserRequest) GetWnUserId() *string {
	return s.WnUserId
}

func (s *GetUserRequest) SetAppToken(v string) *GetUserRequest {
	s.AppToken = &v
	return s
}

func (s *GetUserRequest) SetOpsUserCode(v string) *GetUserRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetUserRequest) SetOpsUserId(v string) *GetUserRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetUserRequest) SetTenantId(v string) *GetUserRequest {
	s.TenantId = &v
	return s
}

func (s *GetUserRequest) SetWnAccountId(v string) *GetUserRequest {
	s.WnAccountId = &v
	return s
}

func (s *GetUserRequest) SetWnUserId(v string) *GetUserRequest {
	s.WnUserId = &v
	return s
}

func (s *GetUserRequest) Validate() error {
	return dara.Validate(s)
}
