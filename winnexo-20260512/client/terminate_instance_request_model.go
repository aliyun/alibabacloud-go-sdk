// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *TerminateInstanceRequest
	GetCollaborationGroupId() *string
	SetInstanceId(v string) *TerminateInstanceRequest
	GetInstanceId() *string
	SetReason(v string) *TerminateInstanceRequest
	GetReason() *string
	SetTenantId(v string) *TerminateInstanceRequest
	GetTenantId() *string
}

type TerminateInstanceRequest struct {
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
	// 终止原因
	//
	// This parameter is required.
	//
	// example:
	//
	// 用户终止
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 租户ID，缺省使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s TerminateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminateInstanceRequest) GoString() string {
	return s.String()
}

func (s *TerminateInstanceRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *TerminateInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TerminateInstanceRequest) GetReason() *string {
	return s.Reason
}

func (s *TerminateInstanceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *TerminateInstanceRequest) SetCollaborationGroupId(v string) *TerminateInstanceRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *TerminateInstanceRequest) SetInstanceId(v string) *TerminateInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *TerminateInstanceRequest) SetReason(v string) *TerminateInstanceRequest {
	s.Reason = &v
	return s
}

func (s *TerminateInstanceRequest) SetTenantId(v string) *TerminateInstanceRequest {
	s.TenantId = &v
	return s
}

func (s *TerminateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
