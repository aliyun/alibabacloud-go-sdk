// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduleWorkItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ScheduleWorkItemRequest
	GetCollaborationGroupId() *string
	SetIdempotencyKey(v string) *ScheduleWorkItemRequest
	GetIdempotencyKey() *string
	SetTenantId(v string) *ScheduleWorkItemRequest
	GetTenantId() *string
	SetWorkItemId(v string) *ScheduleWorkItemRequest
	GetWorkItemId() *string
}

type ScheduleWorkItemRequest struct {
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 幂等键
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	IdempotencyKey *string `json:"idempotencyKey,omitempty" xml:"idempotencyKey,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 工作项 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// wi_example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s ScheduleWorkItemRequest) String() string {
	return dara.Prettify(s)
}

func (s ScheduleWorkItemRequest) GoString() string {
	return s.String()
}

func (s *ScheduleWorkItemRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ScheduleWorkItemRequest) GetIdempotencyKey() *string {
	return s.IdempotencyKey
}

func (s *ScheduleWorkItemRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ScheduleWorkItemRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *ScheduleWorkItemRequest) SetCollaborationGroupId(v string) *ScheduleWorkItemRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ScheduleWorkItemRequest) SetIdempotencyKey(v string) *ScheduleWorkItemRequest {
	s.IdempotencyKey = &v
	return s
}

func (s *ScheduleWorkItemRequest) SetTenantId(v string) *ScheduleWorkItemRequest {
	s.TenantId = &v
	return s
}

func (s *ScheduleWorkItemRequest) SetWorkItemId(v string) *ScheduleWorkItemRequest {
	s.WorkItemId = &v
	return s
}

func (s *ScheduleWorkItemRequest) Validate() error {
	return dara.Validate(s)
}
