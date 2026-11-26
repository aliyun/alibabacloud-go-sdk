// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetWorkflowInstanceRequest
	GetAppToken() *string
	SetInstanceId(v string) *GetWorkflowInstanceRequest
	GetInstanceId() *string
	SetOpsUserCode(v string) *GetWorkflowInstanceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetWorkflowInstanceRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetWorkflowInstanceRequest
	GetTenantId() *string
}

type GetWorkflowInstanceRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The stable logical instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// wfi_20261016_abc123
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
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
	// The tenant ID. This is a common parameter. If you use winnexo-cli, pass this parameter explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetWorkflowInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetWorkflowInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkflowInstanceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetWorkflowInstanceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetWorkflowInstanceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkflowInstanceRequest) SetAppToken(v string) *GetWorkflowInstanceRequest {
	s.AppToken = &v
	return s
}

func (s *GetWorkflowInstanceRequest) SetInstanceId(v string) *GetWorkflowInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetWorkflowInstanceRequest) SetOpsUserCode(v string) *GetWorkflowInstanceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetWorkflowInstanceRequest) SetOpsUserId(v string) *GetWorkflowInstanceRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetWorkflowInstanceRequest) SetTenantId(v string) *GetWorkflowInstanceRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkflowInstanceRequest) Validate() error {
	return dara.Validate(s)
}
