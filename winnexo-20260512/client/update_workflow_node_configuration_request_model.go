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
	Executor *UpdateWorkflowNodeConfigurationRequestExecutor `json:"executor,omitempty" xml:"executor,omitempty" type:"Struct"`
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
	// The list of user IDs for the node owners. The value must be a non-empty array of positive integers. If this parameter is omitted, the current value is retained. If this parameter is set to null, the default value is restored.
	//
	// example:
	//
	// [10001]
	ResponsibleUserIds []*int64 `json:"responsibleUserIds,omitempty" xml:"responsibleUserIds,omitempty" type:"Repeated"`
	// The tenant ID. If not specified, the default tenant of the caller is used.
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
	// The executor ID. If the type is USER, the value is a string representing a positive integer platform user ID. If the type is AGENT, the value is the technical name of the digital worker in the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent_1001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The executor type. Valid values: USER (a human user) and AGENT (a digital worker).
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
