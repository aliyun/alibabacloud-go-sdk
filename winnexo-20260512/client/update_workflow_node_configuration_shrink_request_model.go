// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowNodeConfigurationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateWorkflowNodeConfigurationShrinkRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *UpdateWorkflowNodeConfigurationShrinkRequest
	GetCollaborationGroupId() *string
	SetExecutorShrink(v string) *UpdateWorkflowNodeConfigurationShrinkRequest
	GetExecutorShrink() *string
	SetExpectedVersion(v int64) *UpdateWorkflowNodeConfigurationShrinkRequest
	GetExpectedVersion() *int64
	SetInstanceId(v string) *UpdateWorkflowNodeConfigurationShrinkRequest
	GetInstanceId() *string
	SetModel(v string) *UpdateWorkflowNodeConfigurationShrinkRequest
	GetModel() *string
	SetNodeId(v string) *UpdateWorkflowNodeConfigurationShrinkRequest
	GetNodeId() *string
	SetOpsUserCode(v string) *UpdateWorkflowNodeConfigurationShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateWorkflowNodeConfigurationShrinkRequest
	GetOpsUserId() *string
	SetResponsibleUserIdsShrink(v string) *UpdateWorkflowNodeConfigurationShrinkRequest
	GetResponsibleUserIdsShrink() *string
	SetTenantId(v string) *UpdateWorkflowNodeConfigurationShrinkRequest
	GetTenantId() *string
}

type UpdateWorkflowNodeConfigurationShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the collaboration workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The executor of the node. If this parameter is omitted, the current value is retained. If this parameter is set to null, the default value is restored. You must modify at least one of the owner, model, or executor.
	//
	// example:
	//
	// {"type":"AGENT","id":"agent_1001"}
	ExecutorShrink *string `json:"executor,omitempty" xml:"executor,omitempty"`
	// The configuration version returned by the read operation. The value must be a non-negative integer. If a version conflict occurs, you must refresh the version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ExpectedVersion *int64 `json:"expectedVersion,omitempty" xml:"expectedVersion,omitempty"`
	// The ID of the current physical plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// inst_2081
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The model mode. The quick value is normalized to standard. If this parameter is omitted, the current value is retained. If this parameter is set to null, the default value is restored.
	//
	// example:
	//
	// standard
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The key of the business node.
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
	// The list of user IDs for the node owners. The value must be a non-empty array of positive integers. If this parameter is omitted, the current value is retained. If this parameter is set to null, the default value is restored.
	//
	// example:
	//
	// [10001]
	ResponsibleUserIdsShrink *string `json:"responsibleUserIds,omitempty" xml:"responsibleUserIds,omitempty"`
	// The tenant ID. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateWorkflowNodeConfigurationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowNodeConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) GetExecutorShrink() *string {
	return s.ExecutorShrink
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) GetExpectedVersion() *int64 {
	return s.ExpectedVersion
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) GetResponsibleUserIdsShrink() *string {
	return s.ResponsibleUserIdsShrink
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) SetAppToken(v string) *UpdateWorkflowNodeConfigurationShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) SetCollaborationGroupId(v string) *UpdateWorkflowNodeConfigurationShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) SetExecutorShrink(v string) *UpdateWorkflowNodeConfigurationShrinkRequest {
	s.ExecutorShrink = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) SetExpectedVersion(v int64) *UpdateWorkflowNodeConfigurationShrinkRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) SetInstanceId(v string) *UpdateWorkflowNodeConfigurationShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) SetModel(v string) *UpdateWorkflowNodeConfigurationShrinkRequest {
	s.Model = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) SetNodeId(v string) *UpdateWorkflowNodeConfigurationShrinkRequest {
	s.NodeId = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) SetOpsUserCode(v string) *UpdateWorkflowNodeConfigurationShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) SetOpsUserId(v string) *UpdateWorkflowNodeConfigurationShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) SetResponsibleUserIdsShrink(v string) *UpdateWorkflowNodeConfigurationShrinkRequest {
	s.ResponsibleUserIdsShrink = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) SetTenantId(v string) *UpdateWorkflowNodeConfigurationShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
