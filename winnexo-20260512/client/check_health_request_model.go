// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckHealthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CheckHealthRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *CheckHealthRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CheckHealthRequest
	GetOpsUserId() *string
	SetTenantId(v string) *CheckHealthRequest
	GetTenantId() *string
}

type CheckHealthRequest struct {
	// 平台应用 Token（控制台签发）。传入后以 opsUserId 或 opsUserCode 指定的用户身份执行
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
	// The tenant ID.
	//
	// example:
	//
	// 692318833855074
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CheckHealthRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckHealthRequest) GoString() string {
	return s.String()
}

func (s *CheckHealthRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CheckHealthRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CheckHealthRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CheckHealthRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CheckHealthRequest) SetAppToken(v string) *CheckHealthRequest {
	s.AppToken = &v
	return s
}

func (s *CheckHealthRequest) SetOpsUserCode(v string) *CheckHealthRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CheckHealthRequest) SetOpsUserId(v string) *CheckHealthRequest {
	s.OpsUserId = &v
	return s
}

func (s *CheckHealthRequest) SetTenantId(v string) *CheckHealthRequest {
	s.TenantId = &v
	return s
}

func (s *CheckHealthRequest) Validate() error {
	return dara.Validate(s)
}
