// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetWorkflowNodeConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ResetWorkflowNodeConfigurationRequest
	GetCollaborationGroupId() *string
	SetExpectedVersion(v int64) *ResetWorkflowNodeConfigurationRequest
	GetExpectedVersion() *int64
	SetInstanceId(v string) *ResetWorkflowNodeConfigurationRequest
	GetInstanceId() *string
	SetNodeId(v string) *ResetWorkflowNodeConfigurationRequest
	GetNodeId() *string
	SetTenantId(v string) *ResetWorkflowNodeConfigurationRequest
	GetTenantId() *string
}

type ResetWorkflowNodeConfigurationRequest struct {
	// 协作空间ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
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
	// 业务节点key（nodeKey）
	//
	// This parameter is required.
	//
	// example:
	//
	// node_7f3a2c
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 租户ID；缺省使用调用方默认租户
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

func (s *ResetWorkflowNodeConfigurationRequest) GetTenantId() *string {
	return s.TenantId
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

func (s *ResetWorkflowNodeConfigurationRequest) SetTenantId(v string) *ResetWorkflowNodeConfigurationRequest {
	s.TenantId = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
