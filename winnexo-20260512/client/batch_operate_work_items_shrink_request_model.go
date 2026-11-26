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
	SetAppToken(v string) *BatchOperateWorkItemsShrinkRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *BatchOperateWorkItemsShrinkRequest
	GetCollaborationGroupId() *string
	SetIdempotencyKey(v string) *BatchOperateWorkItemsShrinkRequest
	GetIdempotencyKey() *string
	SetOpsUserCode(v string) *BatchOperateWorkItemsShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *BatchOperateWorkItemsShrinkRequest
	GetOpsUserId() *string
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
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
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
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
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

func (s *BatchOperateWorkItemsShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *BatchOperateWorkItemsShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *BatchOperateWorkItemsShrinkRequest) GetIdempotencyKey() *string {
	return s.IdempotencyKey
}

func (s *BatchOperateWorkItemsShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *BatchOperateWorkItemsShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *BatchOperateWorkItemsShrinkRequest) SetAppToken(v string) *BatchOperateWorkItemsShrinkRequest {
	s.AppToken = &v
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

func (s *BatchOperateWorkItemsShrinkRequest) SetOpsUserCode(v string) *BatchOperateWorkItemsShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *BatchOperateWorkItemsShrinkRequest) SetOpsUserId(v string) *BatchOperateWorkItemsShrinkRequest {
	s.OpsUserId = &v
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
