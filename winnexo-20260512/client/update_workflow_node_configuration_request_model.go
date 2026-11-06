// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowNodeConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *UpdateWorkflowNodeConfigurationRequest
	GetCollaborationGroupId() *string
	SetExecutor(v *UpdateWorkflowNodeConfigurationRequestExecutor) *UpdateWorkflowNodeConfigurationRequest
	GetExecutor() *UpdateWorkflowNodeConfigurationRequestExecutor
	SetExpectedVersion(v int64) *UpdateWorkflowNodeConfigurationRequest
	GetExpectedVersion() *int64
	SetInstanceId(v string) *UpdateWorkflowNodeConfigurationRequest
	GetInstanceId() *string
	SetModel(v string) *UpdateWorkflowNodeConfigurationRequest
	GetModel() *string
	SetNodeId(v string) *UpdateWorkflowNodeConfigurationRequest
	GetNodeId() *string
	SetResponsibleUserIds(v []*int64) *UpdateWorkflowNodeConfigurationRequest
	GetResponsibleUserIds() []*int64
	SetTenantId(v string) *UpdateWorkflowNodeConfigurationRequest
	GetTenantId() *string
}

type UpdateWorkflowNodeConfigurationRequest struct {
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
	Executor *UpdateWorkflowNodeConfigurationRequestExecutor `json:"executor,omitempty" xml:"executor,omitempty" type:"Struct"`
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
	ResponsibleUserIds []*int64 `json:"responsibleUserIds,omitempty" xml:"responsibleUserIds,omitempty" type:"Repeated"`
	// 租户ID；缺省使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateWorkflowNodeConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowNodeConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowNodeConfigurationRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *UpdateWorkflowNodeConfigurationRequest) GetExecutor() *UpdateWorkflowNodeConfigurationRequestExecutor {
	return s.Executor
}

func (s *UpdateWorkflowNodeConfigurationRequest) GetExpectedVersion() *int64 {
	return s.ExpectedVersion
}

func (s *UpdateWorkflowNodeConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateWorkflowNodeConfigurationRequest) GetModel() *string {
	return s.Model
}

func (s *UpdateWorkflowNodeConfigurationRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateWorkflowNodeConfigurationRequest) GetResponsibleUserIds() []*int64 {
	return s.ResponsibleUserIds
}

func (s *UpdateWorkflowNodeConfigurationRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateWorkflowNodeConfigurationRequest) SetCollaborationGroupId(v string) *UpdateWorkflowNodeConfigurationRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationRequest) SetExecutor(v *UpdateWorkflowNodeConfigurationRequestExecutor) *UpdateWorkflowNodeConfigurationRequest {
	s.Executor = v
	return s
}

func (s *UpdateWorkflowNodeConfigurationRequest) SetExpectedVersion(v int64) *UpdateWorkflowNodeConfigurationRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationRequest) SetInstanceId(v string) *UpdateWorkflowNodeConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationRequest) SetModel(v string) *UpdateWorkflowNodeConfigurationRequest {
	s.Model = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationRequest) SetNodeId(v string) *UpdateWorkflowNodeConfigurationRequest {
	s.NodeId = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationRequest) SetResponsibleUserIds(v []*int64) *UpdateWorkflowNodeConfigurationRequest {
	s.ResponsibleUserIds = v
	return s
}

func (s *UpdateWorkflowNodeConfigurationRequest) SetTenantId(v string) *UpdateWorkflowNodeConfigurationRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationRequest) Validate() error {
	if s.Executor != nil {
		if err := s.Executor.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkflowNodeConfigurationRequestExecutor struct {
	// USER 为正整数平台用户ID字符串；AGENT 为空间数字员工技术名称
	//
	// This parameter is required.
	//
	// example:
	//
	// agent_1001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// USER 真人或 AGENT 数字员工
	//
	// This parameter is required.
	//
	// example:
	//
	// AGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateWorkflowNodeConfigurationRequestExecutor) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowNodeConfigurationRequestExecutor) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowNodeConfigurationRequestExecutor) GetId() *string {
	return s.Id
}

func (s *UpdateWorkflowNodeConfigurationRequestExecutor) GetType() *string {
	return s.Type
}

func (s *UpdateWorkflowNodeConfigurationRequestExecutor) SetId(v string) *UpdateWorkflowNodeConfigurationRequestExecutor {
	s.Id = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationRequestExecutor) SetType(v string) *UpdateWorkflowNodeConfigurationRequestExecutor {
	s.Type = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationRequestExecutor) Validate() error {
	return dara.Validate(s)
}
