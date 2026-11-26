// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetWorkflowNodeConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ResetWorkflowNodeConfigurationRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *ResetWorkflowNodeConfigurationRequest
	GetCollaborationGroupId() *string
	SetExpectedVersion(v int64) *ResetWorkflowNodeConfigurationRequest
	GetExpectedVersion() *int64
	SetInstanceId(v string) *ResetWorkflowNodeConfigurationRequest
	GetInstanceId() *string
	SetNodeId(v string) *ResetWorkflowNodeConfigurationRequest
	GetNodeId() *string
	SetOpsUserCode(v string) *ResetWorkflowNodeConfigurationRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ResetWorkflowNodeConfigurationRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ResetWorkflowNodeConfigurationRequest
	GetTenantId() *string
}

type ResetWorkflowNodeConfigurationRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The collaboration workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The configuration version returned by the read API. This must be a non-negative integer. Refresh the data if a version conflict occurs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ExpectedVersion *int64 `json:"expectedVersion,omitempty" xml:"expectedVersion,omitempty"`
	// The current physical plan ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// inst_2081
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The business node key.
	//
	// This parameter is required.
	//
	// example:
	//
	// node_7f3a2c
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
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
	// The tenant ID. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ResetWorkflowNodeConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetWorkflowNodeConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ResetWorkflowNodeConfigurationRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ResetWorkflowNodeConfigurationRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ResetWorkflowNodeConfigurationRequest) GetExpectedVersion() *int64 {
	return s.ExpectedVersion
}

func (s *ResetWorkflowNodeConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetWorkflowNodeConfigurationRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ResetWorkflowNodeConfigurationRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ResetWorkflowNodeConfigurationRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ResetWorkflowNodeConfigurationRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ResetWorkflowNodeConfigurationRequest) SetAppToken(v string) *ResetWorkflowNodeConfigurationRequest {
	s.AppToken = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationRequest) SetCollaborationGroupId(v string) *ResetWorkflowNodeConfigurationRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationRequest) SetExpectedVersion(v int64) *ResetWorkflowNodeConfigurationRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationRequest) SetInstanceId(v string) *ResetWorkflowNodeConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationRequest) SetNodeId(v string) *ResetWorkflowNodeConfigurationRequest {
	s.NodeId = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationRequest) SetOpsUserCode(v string) *ResetWorkflowNodeConfigurationRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationRequest) SetOpsUserId(v string) *ResetWorkflowNodeConfigurationRequest {
	s.OpsUserId = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationRequest) SetTenantId(v string) *ResetWorkflowNodeConfigurationRequest {
	s.TenantId = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
