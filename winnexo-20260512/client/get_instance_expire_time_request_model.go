// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceExpireTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetInstanceExpireTimeRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *GetInstanceExpireTimeRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetInstanceExpireTimeRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetInstanceExpireTimeRequest
	GetTenantId() *string
}

type GetInstanceExpireTimeRequest struct {
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

func (s GetInstanceExpireTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceExpireTimeRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceExpireTimeRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetInstanceExpireTimeRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetInstanceExpireTimeRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetInstanceExpireTimeRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetInstanceExpireTimeRequest) SetAppToken(v string) *GetInstanceExpireTimeRequest {
	s.AppToken = &v
	return s
}

func (s *GetInstanceExpireTimeRequest) SetOpsUserCode(v string) *GetInstanceExpireTimeRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetInstanceExpireTimeRequest) SetOpsUserId(v string) *GetInstanceExpireTimeRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetInstanceExpireTimeRequest) SetTenantId(v string) *GetInstanceExpireTimeRequest {
	s.TenantId = &v
	return s
}

func (s *GetInstanceExpireTimeRequest) Validate() error {
	return dara.Validate(s)
}
