// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ListInstanceNodesRequest
	GetCollaborationGroupId() *string
	SetInstanceId(v string) *ListInstanceNodesRequest
	GetInstanceId() *string
	SetTenantId(v string) *ListInstanceNodesRequest
	GetTenantId() *string
}

type ListInstanceNodesRequest struct {
	// The collaboration workspace ID.
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
	// The tenant ID. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListInstanceNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceNodesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceNodesRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListInstanceNodesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceNodesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListInstanceNodesRequest) SetCollaborationGroupId(v string) *ListInstanceNodesRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListInstanceNodesRequest) SetInstanceId(v string) *ListInstanceNodesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceNodesRequest) SetTenantId(v string) *ListInstanceNodesRequest {
	s.TenantId = &v
	return s
}

func (s *ListInstanceNodesRequest) Validate() error {
	return dara.Validate(s)
}
