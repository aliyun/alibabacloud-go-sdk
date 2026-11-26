// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperateWorkItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *BatchOperateWorkItemsRequest
	GetAction() *string
	SetAppToken(v string) *BatchOperateWorkItemsRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *BatchOperateWorkItemsRequest
	GetCollaborationGroupId() *string
	SetIdempotencyKey(v string) *BatchOperateWorkItemsRequest
	GetIdempotencyKey() *string
	SetOpsUserCode(v string) *BatchOperateWorkItemsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *BatchOperateWorkItemsRequest
	GetOpsUserId() *string
	SetTenantId(v string) *BatchOperateWorkItemsRequest
	GetTenantId() *string
	SetWorkItemIds(v []*string) *BatchOperateWorkItemsRequest
	GetWorkItemIds() []*string
}

type BatchOperateWorkItemsRequest struct {
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
	WorkItemIds []*string `json:"workItemIds,omitempty" xml:"workItemIds,omitempty" type:"Repeated"`
}

func (s BatchOperateWorkItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchOperateWorkItemsRequest) GoString() string {
	return s.String()
}

func (s *BatchOperateWorkItemsRequest) GetAction() *string {
	return s.Action
}

func (s *BatchOperateWorkItemsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *BatchOperateWorkItemsRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *BatchOperateWorkItemsRequest) GetIdempotencyKey() *string {
	return s.IdempotencyKey
}

func (s *BatchOperateWorkItemsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *BatchOperateWorkItemsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *BatchOperateWorkItemsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *BatchOperateWorkItemsRequest) GetWorkItemIds() []*string {
	return s.WorkItemIds
}

func (s *BatchOperateWorkItemsRequest) SetAction(v string) *BatchOperateWorkItemsRequest {
	s.Action = &v
	return s
}

func (s *BatchOperateWorkItemsRequest) SetAppToken(v string) *BatchOperateWorkItemsRequest {
	s.AppToken = &v
	return s
}

func (s *BatchOperateWorkItemsRequest) SetCollaborationGroupId(v string) *BatchOperateWorkItemsRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *BatchOperateWorkItemsRequest) SetIdempotencyKey(v string) *BatchOperateWorkItemsRequest {
	s.IdempotencyKey = &v
	return s
}

func (s *BatchOperateWorkItemsRequest) SetOpsUserCode(v string) *BatchOperateWorkItemsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *BatchOperateWorkItemsRequest) SetOpsUserId(v string) *BatchOperateWorkItemsRequest {
	s.OpsUserId = &v
	return s
}

func (s *BatchOperateWorkItemsRequest) SetTenantId(v string) *BatchOperateWorkItemsRequest {
	s.TenantId = &v
	return s
}

func (s *BatchOperateWorkItemsRequest) SetWorkItemIds(v []*string) *BatchOperateWorkItemsRequest {
	s.WorkItemIds = v
	return s
}

func (s *BatchOperateWorkItemsRequest) Validate() error {
	return dara.Validate(s)
}
