// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperateWorkItemsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *BatchOperateWorkItemsShrinkRequest
	GetAction() *string
	SetCollaborationGroupId(v string) *BatchOperateWorkItemsShrinkRequest
	GetCollaborationGroupId() *string
	SetIdempotencyKey(v string) *BatchOperateWorkItemsShrinkRequest
	GetIdempotencyKey() *string
	SetTenantId(v string) *BatchOperateWorkItemsShrinkRequest
	GetTenantId() *string
	SetWorkItemIdsShrink(v string) *BatchOperateWorkItemsShrinkRequest
	GetWorkItemIdsShrink() *string
}

type BatchOperateWorkItemsShrinkRequest struct {
	// The batch action.
	//
	// This parameter is required.
	//
	// example:
	//
	// SCHEDULE
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The ID of the collaboration workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The idempotency key.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	IdempotencyKey *string `json:"idempotencyKey,omitempty" xml:"idempotencyKey,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The IDs of the target work items. You can specify 1 to 100 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	WorkItemIdsShrink *string `json:"workItemIds,omitempty" xml:"workItemIds,omitempty"`
}

func (s BatchOperateWorkItemsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchOperateWorkItemsShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchOperateWorkItemsShrinkRequest) GetAction() *string {
	return s.Action
}

func (s *BatchOperateWorkItemsShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *BatchOperateWorkItemsShrinkRequest) GetIdempotencyKey() *string {
	return s.IdempotencyKey
}

func (s *BatchOperateWorkItemsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *BatchOperateWorkItemsShrinkRequest) GetWorkItemIdsShrink() *string {
	return s.WorkItemIdsShrink
}

func (s *BatchOperateWorkItemsShrinkRequest) SetAction(v string) *BatchOperateWorkItemsShrinkRequest {
	s.Action = &v
	return s
}

func (s *BatchOperateWorkItemsShrinkRequest) SetCollaborationGroupId(v string) *BatchOperateWorkItemsShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *BatchOperateWorkItemsShrinkRequest) SetIdempotencyKey(v string) *BatchOperateWorkItemsShrinkRequest {
	s.IdempotencyKey = &v
	return s
}

func (s *BatchOperateWorkItemsShrinkRequest) SetTenantId(v string) *BatchOperateWorkItemsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *BatchOperateWorkItemsShrinkRequest) SetWorkItemIdsShrink(v string) *BatchOperateWorkItemsShrinkRequest {
	s.WorkItemIdsShrink = &v
	return s
}

func (s *BatchOperateWorkItemsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
