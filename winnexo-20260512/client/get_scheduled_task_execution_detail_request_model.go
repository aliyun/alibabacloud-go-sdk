// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskExecutionDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetScheduledTaskExecutionDetailRequest
	GetAppToken() *string
	SetExecutionId(v string) *GetScheduledTaskExecutionDetailRequest
	GetExecutionId() *string
	SetOpsUserCode(v string) *GetScheduledTaskExecutionDetailRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetScheduledTaskExecutionDetailRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetScheduledTaskExecutionDetailRequest
	GetTenantId() *string
}

type GetScheduledTaskExecutionDetailRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the execution record.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleExecutionId
	ExecutionId *string `json:"executionId,omitempty" xml:"executionId,omitempty"`
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

func (s GetScheduledTaskExecutionDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskExecutionDetailRequest) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskExecutionDetailRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetScheduledTaskExecutionDetailRequest) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *GetScheduledTaskExecutionDetailRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetScheduledTaskExecutionDetailRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetScheduledTaskExecutionDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetScheduledTaskExecutionDetailRequest) SetAppToken(v string) *GetScheduledTaskExecutionDetailRequest {
	s.AppToken = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailRequest) SetExecutionId(v string) *GetScheduledTaskExecutionDetailRequest {
	s.ExecutionId = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailRequest) SetOpsUserCode(v string) *GetScheduledTaskExecutionDetailRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailRequest) SetOpsUserId(v string) *GetScheduledTaskExecutionDetailRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailRequest) SetTenantId(v string) *GetScheduledTaskExecutionDetailRequest {
	s.TenantId = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailRequest) Validate() error {
	return dara.Validate(s)
}
