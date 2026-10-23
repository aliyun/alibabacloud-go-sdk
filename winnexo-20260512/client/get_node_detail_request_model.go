// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *GetNodeDetailRequest
	GetCollaborationGroupId() *string
	SetInstanceId(v string) *GetNodeDetailRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetNodeDetailRequest
	GetNodeId() *string
	SetTenantId(v string) *GetNodeDetailRequest
	GetTenantId() *string
}

type GetNodeDetailRequest struct {
	// The ID of the collaboration workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The physical workflow instance ID (planId).
	//
	// This parameter is required.
	//
	// example:
	//
	// plan_example
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The ID of the node (nodeKey).
	//
	// This parameter is required.
	//
	// example:
	//
	// step_1
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The tenant ID. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetNodeDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeDetailRequest) GoString() string {
	return s.String()
}

func (s *GetNodeDetailRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetNodeDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNodeDetailRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNodeDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetNodeDetailRequest) SetCollaborationGroupId(v string) *GetNodeDetailRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetNodeDetailRequest) SetInstanceId(v string) *GetNodeDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNodeDetailRequest) SetNodeId(v string) *GetNodeDetailRequest {
	s.NodeId = &v
	return s
}

func (s *GetNodeDetailRequest) SetTenantId(v string) *GetNodeDetailRequest {
	s.TenantId = &v
	return s
}

func (s *GetNodeDetailRequest) Validate() error {
	return dara.Validate(s)
}
