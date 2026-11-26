// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduleWorkItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ScheduleWorkItemRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *ScheduleWorkItemRequest
	GetCollaborationGroupId() *string
	SetIdempotencyKey(v string) *ScheduleWorkItemRequest
	GetIdempotencyKey() *string
	SetOpsUserCode(v string) *ScheduleWorkItemRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ScheduleWorkItemRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ScheduleWorkItemRequest
	GetTenantId() *string
	SetWorkItemId(v string) *ScheduleWorkItemRequest
	GetWorkItemId() *string
}

type ScheduleWorkItemRequest struct {
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
	// The ID of the work item.
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

func (s *ScheduleWorkItemRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ScheduleWorkItemRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ScheduleWorkItemRequest) GetIdempotencyKey() *string {
	return s.IdempotencyKey
}

func (s *ScheduleWorkItemRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ScheduleWorkItemRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ScheduleWorkItemRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ScheduleWorkItemRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *ScheduleWorkItemRequest) SetAppToken(v string) *ScheduleWorkItemRequest {
	s.AppToken = &v
	return s
}

func (s *ScheduleWorkItemRequest) SetCollaborationGroupId(v string) *ScheduleWorkItemRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ScheduleWorkItemRequest) SetIdempotencyKey(v string) *ScheduleWorkItemRequest {
	s.IdempotencyKey = &v
	return s
}

func (s *ScheduleWorkItemRequest) SetOpsUserCode(v string) *ScheduleWorkItemRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ScheduleWorkItemRequest) SetOpsUserId(v string) *ScheduleWorkItemRequest {
	s.OpsUserId = &v
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
