// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetUserInfoRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *GetUserInfoRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetUserInfoRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetUserInfoRequest
	GetTenantId() *string
}

type GetUserInfoRequest struct {
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
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetUserInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserInfoRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetUserInfoRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetUserInfoRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetUserInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserInfoRequest) SetAppToken(v string) *GetUserInfoRequest {
	s.AppToken = &v
	return s
}

func (s *GetUserInfoRequest) SetOpsUserCode(v string) *GetUserInfoRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetUserInfoRequest) SetOpsUserId(v string) *GetUserInfoRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetUserInfoRequest) SetTenantId(v string) *GetUserInfoRequest {
	s.TenantId = &v
	return s
}

func (s *GetUserInfoRequest) Validate() error {
	return dara.Validate(s)
}
