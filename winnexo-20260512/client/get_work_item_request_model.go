// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetWorkItemRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *GetWorkItemRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *GetWorkItemRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetWorkItemRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetWorkItemRequest
	GetTenantId() *string
	SetWorkItemId(v string) *GetWorkItemRequest
	GetWorkItemId() *string
}

type GetWorkItemRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the collaboration space.
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

func (s GetWorkItemRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemRequest) GoString() string {
	return s.String()
}

func (s *GetWorkItemRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetWorkItemRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkItemRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetWorkItemRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetWorkItemRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkItemRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetWorkItemRequest) SetAppToken(v string) *GetWorkItemRequest {
	s.AppToken = &v
	return s
}

func (s *GetWorkItemRequest) SetCollaborationGroupId(v string) *GetWorkItemRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkItemRequest) SetOpsUserCode(v string) *GetWorkItemRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetWorkItemRequest) SetOpsUserId(v string) *GetWorkItemRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetWorkItemRequest) SetTenantId(v string) *GetWorkItemRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkItemRequest) SetWorkItemId(v string) *GetWorkItemRequest {
	s.WorkItemId = &v
	return s
}

func (s *GetWorkItemRequest) Validate() error {
	return dara.Validate(s)
}
