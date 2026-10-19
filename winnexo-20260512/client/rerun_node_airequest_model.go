// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunNodeAIRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *RerunNodeAIRequest
	GetCollaborationGroupId() *string
	SetInstanceId(v string) *RerunNodeAIRequest
	GetInstanceId() *string
	SetNodeId(v string) *RerunNodeAIRequest
	GetNodeId() *string
	SetTenantId(v string) *RerunNodeAIRequest
	GetTenantId() *string
}

type RerunNodeAIRequest struct {
	// 协作空间ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 物理工作流实例ID（planId）
	//
	// This parameter is required.
	//
	// example:
	//
	// plan_example
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// FAILED 节点 key（nodeKey）
	//
	// This parameter is required.
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 租户ID，缺省使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RerunNodeAIRequest) String() string {
	return dara.Prettify(s)
}

func (s RerunNodeAIRequest) GoString() string {
	return s.String()
}

func (s *RerunNodeAIRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *RerunNodeAIRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RerunNodeAIRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *RerunNodeAIRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RerunNodeAIRequest) SetCollaborationGroupId(v string) *RerunNodeAIRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *RerunNodeAIRequest) SetInstanceId(v string) *RerunNodeAIRequest {
	s.InstanceId = &v
	return s
}

func (s *RerunNodeAIRequest) SetNodeId(v string) *RerunNodeAIRequest {
	s.NodeId = &v
	return s
}

func (s *RerunNodeAIRequest) SetTenantId(v string) *RerunNodeAIRequest {
	s.TenantId = &v
	return s
}

func (s *RerunNodeAIRequest) Validate() error {
	return dara.Validate(s)
}
