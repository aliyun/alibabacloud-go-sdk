// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserCreditUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetUserCreditUsageRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *GetUserCreditUsageRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetUserCreditUsageRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetUserCreditUsageRequest
	GetTenantId() *string
}

type GetUserCreditUsageRequest struct {
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
}

func (s GetUserCreditUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserCreditUsageRequest) GoString() string {
	return s.String()
}

func (s *GetUserCreditUsageRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetUserCreditUsageRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetUserCreditUsageRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetUserCreditUsageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserCreditUsageRequest) SetAppToken(v string) *GetUserCreditUsageRequest {
	s.AppToken = &v
	return s
}

func (s *GetUserCreditUsageRequest) SetOpsUserCode(v string) *GetUserCreditUsageRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetUserCreditUsageRequest) SetOpsUserId(v string) *GetUserCreditUsageRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetUserCreditUsageRequest) SetTenantId(v string) *GetUserCreditUsageRequest {
	s.TenantId = &v
	return s
}

func (s *GetUserCreditUsageRequest) Validate() error {
	return dara.Validate(s)
}
