// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseWorkItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *PauseWorkItemRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *PauseWorkItemRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *PauseWorkItemRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *PauseWorkItemRequest
	GetOpsUserId() *string
	SetTenantId(v string) *PauseWorkItemRequest
	GetTenantId() *string
	SetWorkItemId(v string) *PauseWorkItemRequest
	GetWorkItemId() *string
}

type PauseWorkItemRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The collaboration space ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
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
	// The work item ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// wi_example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s PauseWorkItemRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseWorkItemRequest) GoString() string {
	return s.String()
}

func (s *PauseWorkItemRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *PauseWorkItemRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *PauseWorkItemRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *PauseWorkItemRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *PauseWorkItemRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PauseWorkItemRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *PauseWorkItemRequest) SetAppToken(v string) *PauseWorkItemRequest {
	s.AppToken = &v
	return s
}

func (s *PauseWorkItemRequest) SetCollaborationGroupId(v string) *PauseWorkItemRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *PauseWorkItemRequest) SetOpsUserCode(v string) *PauseWorkItemRequest {
	s.OpsUserCode = &v
	return s
}

func (s *PauseWorkItemRequest) SetOpsUserId(v string) *PauseWorkItemRequest {
	s.OpsUserId = &v
	return s
}

func (s *PauseWorkItemRequest) SetTenantId(v string) *PauseWorkItemRequest {
	s.TenantId = &v
	return s
}

func (s *PauseWorkItemRequest) SetWorkItemId(v string) *PauseWorkItemRequest {
	s.WorkItemId = &v
	return s
}

func (s *PauseWorkItemRequest) Validate() error {
	return dara.Validate(s)
}
