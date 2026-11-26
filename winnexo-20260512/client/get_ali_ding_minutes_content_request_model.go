// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliDingMinutesContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetAliDingMinutesContentRequest
	GetAppToken() *string
	SetMinutesId(v string) *GetAliDingMinutesContentRequest
	GetMinutesId() *string
	SetOpsUserCode(v string) *GetAliDingMinutesContentRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetAliDingMinutesContentRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetAliDingMinutesContentRequest
	GetTenantId() *string
}

type GetAliDingMinutesContentRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The DingTalk minutes ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 76327569643231383535353939365f3436383537393431335f32
	MinutesId *string `json:"minutesId,omitempty" xml:"minutesId,omitempty"`
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
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetAliDingMinutesContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliDingMinutesContentRequest) GoString() string {
	return s.String()
}

func (s *GetAliDingMinutesContentRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetAliDingMinutesContentRequest) GetMinutesId() *string {
	return s.MinutesId
}

func (s *GetAliDingMinutesContentRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetAliDingMinutesContentRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetAliDingMinutesContentRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetAliDingMinutesContentRequest) SetAppToken(v string) *GetAliDingMinutesContentRequest {
	s.AppToken = &v
	return s
}

func (s *GetAliDingMinutesContentRequest) SetMinutesId(v string) *GetAliDingMinutesContentRequest {
	s.MinutesId = &v
	return s
}

func (s *GetAliDingMinutesContentRequest) SetOpsUserCode(v string) *GetAliDingMinutesContentRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetAliDingMinutesContentRequest) SetOpsUserId(v string) *GetAliDingMinutesContentRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetAliDingMinutesContentRequest) SetTenantId(v string) *GetAliDingMinutesContentRequest {
	s.TenantId = &v
	return s
}

func (s *GetAliDingMinutesContentRequest) Validate() error {
	return dara.Validate(s)
}
