// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowNodeConfigurationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetResponsibleUserIdsShrink(v string) *UpdateWorkflowNodeConfigurationShrinkRequest
	GetResponsibleUserIdsShrink() *string
	SetTenantId(v string) *UpdateWorkflowNodeConfigurationShrinkRequest
	GetTenantId() *string
}

type UpdateWorkflowNodeConfigurationShrinkRequest struct {
	// 协作空间ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 节点执行人；省略保持，null恢复默认；与负责人、模型至少修改一项
	//
	// example:
	//
	// {"type":"AGENT","id":"agent_1001"}
	ExecutorShrink *string `json:"executor,omitempty" xml:"executor,omitempty"`
	// 读取接口返回的configVersion；非负整数；版本冲突须刷新
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ExpectedVersion *int64 `json:"expectedVersion,omitempty" xml:"expectedVersion,omitempty"`
	// 当前物理计划ID（planId）
	//
	// This parameter is required.
	//
	// example:
	//
	// inst_2081
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 模型模式；quick归一化standard；省略保持，null恢复默认
	//
	// example:
	//
	// standard
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// 业务节点key（nodeKey）
	//
	// This parameter is required.
	//
	// example:
	//
	// node_7f3a2c
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 节点负责人用户ID列表，非空正整数数组；省略保持，null恢复默认
	//
	// example:
	//
	// [10001]
	ResponsibleUserIdsShrink *string `json:"responsibleUserIds,omitempty" xml:"responsibleUserIds,omitempty"`
	// 租户ID；缺省使用调用方默认租户
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

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) GetResponsibleUserIdsShrink() *string {
	return s.ResponsibleUserIdsShrink
}

func (s *UpdateWorkflowNodeConfigurationShrinkRequest) GetTenantId() *string {
	return s.TenantId
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
