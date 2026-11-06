// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowNodeConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *GetWorkflowNodeConfigurationRequest
	GetCollaborationGroupId() *string
	SetInstanceId(v string) *GetWorkflowNodeConfigurationRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetWorkflowNodeConfigurationRequest
	GetNodeId() *string
	SetTenantId(v string) *GetWorkflowNodeConfigurationRequest
	GetTenantId() *string
}

type GetWorkflowNodeConfigurationRequest struct {
	// 协作空间ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
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

func (s GetWorkflowNodeConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowNodeConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowNodeConfigurationRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkflowNodeConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkflowNodeConfigurationRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetWorkflowNodeConfigurationRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkflowNodeConfigurationRequest) SetCollaborationGroupId(v string) *GetWorkflowNodeConfigurationRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkflowNodeConfigurationRequest) SetInstanceId(v string) *GetWorkflowNodeConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetWorkflowNodeConfigurationRequest) SetNodeId(v string) *GetWorkflowNodeConfigurationRequest {
	s.NodeId = &v
	return s
}

func (s *GetWorkflowNodeConfigurationRequest) SetTenantId(v string) *GetWorkflowNodeConfigurationRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkflowNodeConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
