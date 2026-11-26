// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *RemoveUserRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *RemoveUserRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *RemoveUserRequest
	GetOpsUserId() *string
	SetTenantId(v string) *RemoveUserRequest
	GetTenantId() *string
	SetWnUserId(v string) *RemoveUserRequest
	GetWnUserId() *string
}

type RemoveUserRequest struct {
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
	// The ID of the target user (WINNEXO platform user ID).
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s RemoveUserRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *RemoveUserRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *RemoveUserRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *RemoveUserRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RemoveUserRequest) GetWnUserId() *string {
	return s.WnUserId
}

func (s *RemoveUserRequest) SetAppToken(v string) *RemoveUserRequest {
	s.AppToken = &v
	return s
}

func (s *RemoveUserRequest) SetOpsUserCode(v string) *RemoveUserRequest {
	s.OpsUserCode = &v
	return s
}

func (s *RemoveUserRequest) SetOpsUserId(v string) *RemoveUserRequest {
	s.OpsUserId = &v
	return s
}

func (s *RemoveUserRequest) SetTenantId(v string) *RemoveUserRequest {
	s.TenantId = &v
	return s
}

func (s *RemoveUserRequest) SetWnUserId(v string) *RemoveUserRequest {
	s.WnUserId = &v
	return s
}

func (s *RemoveUserRequest) Validate() error {
	return dara.Validate(s)
}
