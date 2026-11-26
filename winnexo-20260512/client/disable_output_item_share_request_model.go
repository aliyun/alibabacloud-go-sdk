// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableOutputItemShareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *DisableOutputItemShareRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *DisableOutputItemShareRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *DisableOutputItemShareRequest
	GetOpsUserId() *string
	SetOutputItemId(v string) *DisableOutputItemShareRequest
	GetOutputItemId() *string
	SetOutputItemName(v string) *DisableOutputItemShareRequest
	GetOutputItemName() *string
	SetSessionId(v string) *DisableOutputItemShareRequest
	GetSessionId() *string
	SetTenantId(v string) *DisableOutputItemShareRequest
	GetTenantId() *string
}

type DisableOutputItemShareRequest struct {
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
	// The output item ID.
	//
	// example:
	//
	// exampleOutputItemId
	OutputItemId *string `json:"outputItemId,omitempty" xml:"outputItemId,omitempty"`
	// The name of the output item. This parameter must be used together with sessionId.
	//
	// example:
	//
	// Quarterly Review.pptx
	OutputItemName *string `json:"outputItemName,omitempty" xml:"outputItemName,omitempty"`
	// The session ID used for exact filtering. The value must be 1 to 128 characters in length. It applies to common sessions, skill sessions, and the execution ID of task executions.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DisableOutputItemShareRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableOutputItemShareRequest) GoString() string {
	return s.String()
}

func (s *DisableOutputItemShareRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *DisableOutputItemShareRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *DisableOutputItemShareRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *DisableOutputItemShareRequest) GetOutputItemId() *string {
	return s.OutputItemId
}

func (s *DisableOutputItemShareRequest) GetOutputItemName() *string {
	return s.OutputItemName
}

func (s *DisableOutputItemShareRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DisableOutputItemShareRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DisableOutputItemShareRequest) SetAppToken(v string) *DisableOutputItemShareRequest {
	s.AppToken = &v
	return s
}

func (s *DisableOutputItemShareRequest) SetOpsUserCode(v string) *DisableOutputItemShareRequest {
	s.OpsUserCode = &v
	return s
}

func (s *DisableOutputItemShareRequest) SetOpsUserId(v string) *DisableOutputItemShareRequest {
	s.OpsUserId = &v
	return s
}

func (s *DisableOutputItemShareRequest) SetOutputItemId(v string) *DisableOutputItemShareRequest {
	s.OutputItemId = &v
	return s
}

func (s *DisableOutputItemShareRequest) SetOutputItemName(v string) *DisableOutputItemShareRequest {
	s.OutputItemName = &v
	return s
}

func (s *DisableOutputItemShareRequest) SetSessionId(v string) *DisableOutputItemShareRequest {
	s.SessionId = &v
	return s
}

func (s *DisableOutputItemShareRequest) SetTenantId(v string) *DisableOutputItemShareRequest {
	s.TenantId = &v
	return s
}

func (s *DisableOutputItemShareRequest) Validate() error {
	return dara.Validate(s)
}
