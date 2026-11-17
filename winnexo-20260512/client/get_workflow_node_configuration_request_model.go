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
	// The ID of the collaboration workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The ID of the current physical plan (planId).
	//
	// This parameter is required.
	//
	// example:
	//
	// inst_2081
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The key of the business node (nodeKey).
	//
	// This parameter is required.
	//
	// example:
	//
	// node_7f3a2c
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The tenant ID. If not specified, the default tenant of the caller is used.
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
