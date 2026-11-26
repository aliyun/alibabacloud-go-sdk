// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListAgentsRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *ListAgentsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListAgentsRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ListAgentsRequest
	GetTenantId() *string
}

type ListAgentsRequest struct {
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
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListAgentsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListAgentsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListAgentsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAgentsRequest) SetAppToken(v string) *ListAgentsRequest {
	s.AppToken = &v
	return s
}

func (s *ListAgentsRequest) SetOpsUserCode(v string) *ListAgentsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListAgentsRequest) SetOpsUserId(v string) *ListAgentsRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListAgentsRequest) SetTenantId(v string) *ListAgentsRequest {
	s.TenantId = &v
	return s
}

func (s *ListAgentsRequest) Validate() error {
	return dara.Validate(s)
}
