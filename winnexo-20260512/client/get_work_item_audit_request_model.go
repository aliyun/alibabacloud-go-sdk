// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetWorkItemAuditRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *GetWorkItemAuditRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *GetWorkItemAuditRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetWorkItemAuditRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetWorkItemAuditRequest
	GetTenantId() *string
	SetWorkItemId(v string) *GetWorkItemAuditRequest
	GetWorkItemId() *string
}

type GetWorkItemAuditRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The collaboration workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
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
	// The tenant ID. If not specified, the default tenant of the caller is used.
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
	// wi_0123456789ab4cde
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s GetWorkItemAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemAuditRequest) GoString() string {
	return s.String()
}

func (s *GetWorkItemAuditRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetWorkItemAuditRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkItemAuditRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetWorkItemAuditRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetWorkItemAuditRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkItemAuditRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetWorkItemAuditRequest) SetAppToken(v string) *GetWorkItemAuditRequest {
	s.AppToken = &v
	return s
}

func (s *GetWorkItemAuditRequest) SetCollaborationGroupId(v string) *GetWorkItemAuditRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkItemAuditRequest) SetOpsUserCode(v string) *GetWorkItemAuditRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetWorkItemAuditRequest) SetOpsUserId(v string) *GetWorkItemAuditRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetWorkItemAuditRequest) SetTenantId(v string) *GetWorkItemAuditRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkItemAuditRequest) SetWorkItemId(v string) *GetWorkItemAuditRequest {
	s.WorkItemId = &v
	return s
}

func (s *GetWorkItemAuditRequest) Validate() error {
	return dara.Validate(s)
}
