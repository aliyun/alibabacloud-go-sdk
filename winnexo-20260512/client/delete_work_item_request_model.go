// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *DeleteWorkItemRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *DeleteWorkItemRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *DeleteWorkItemRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *DeleteWorkItemRequest
	GetOpsUserId() *string
	SetTenantId(v string) *DeleteWorkItemRequest
	GetTenantId() *string
	SetWorkItemId(v string) *DeleteWorkItemRequest
	GetWorkItemId() *string
}

type DeleteWorkItemRequest struct {
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

func (s DeleteWorkItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkItemRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *DeleteWorkItemRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *DeleteWorkItemRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *DeleteWorkItemRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *DeleteWorkItemRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteWorkItemRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *DeleteWorkItemRequest) SetAppToken(v string) *DeleteWorkItemRequest {
	s.AppToken = &v
	return s
}

func (s *DeleteWorkItemRequest) SetCollaborationGroupId(v string) *DeleteWorkItemRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *DeleteWorkItemRequest) SetOpsUserCode(v string) *DeleteWorkItemRequest {
	s.OpsUserCode = &v
	return s
}

func (s *DeleteWorkItemRequest) SetOpsUserId(v string) *DeleteWorkItemRequest {
	s.OpsUserId = &v
	return s
}

func (s *DeleteWorkItemRequest) SetTenantId(v string) *DeleteWorkItemRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteWorkItemRequest) SetWorkItemId(v string) *DeleteWorkItemRequest {
	s.WorkItemId = &v
	return s
}

func (s *DeleteWorkItemRequest) Validate() error {
	return dara.Validate(s)
}
