// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemAttachmentAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetWorkItemAttachmentAccessRequest
	GetAppToken() *string
	SetAttachmentId(v string) *GetWorkItemAttachmentAccessRequest
	GetAttachmentId() *string
	SetCollaborationGroupId(v string) *GetWorkItemAttachmentAccessRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *GetWorkItemAttachmentAccessRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetWorkItemAttachmentAccessRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetWorkItemAttachmentAccessRequest
	GetTenantId() *string
	SetWorkItemId(v string) *GetWorkItemAttachmentAccessRequest
	GetWorkItemId() *string
}

type GetWorkItemAttachmentAccessRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The attachment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	AttachmentId *string `json:"attachmentId,omitempty" xml:"attachmentId,omitempty"`
	// The collaboration workspace ID.
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
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
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

func (s GetWorkItemAttachmentAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemAttachmentAccessRequest) GoString() string {
	return s.String()
}

func (s *GetWorkItemAttachmentAccessRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetWorkItemAttachmentAccessRequest) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *GetWorkItemAttachmentAccessRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkItemAttachmentAccessRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetWorkItemAttachmentAccessRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetWorkItemAttachmentAccessRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkItemAttachmentAccessRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetWorkItemAttachmentAccessRequest) SetAppToken(v string) *GetWorkItemAttachmentAccessRequest {
	s.AppToken = &v
	return s
}

func (s *GetWorkItemAttachmentAccessRequest) SetAttachmentId(v string) *GetWorkItemAttachmentAccessRequest {
	s.AttachmentId = &v
	return s
}

func (s *GetWorkItemAttachmentAccessRequest) SetCollaborationGroupId(v string) *GetWorkItemAttachmentAccessRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkItemAttachmentAccessRequest) SetOpsUserCode(v string) *GetWorkItemAttachmentAccessRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetWorkItemAttachmentAccessRequest) SetOpsUserId(v string) *GetWorkItemAttachmentAccessRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetWorkItemAttachmentAccessRequest) SetTenantId(v string) *GetWorkItemAttachmentAccessRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkItemAttachmentAccessRequest) SetWorkItemId(v string) *GetWorkItemAttachmentAccessRequest {
	s.WorkItemId = &v
	return s
}

func (s *GetWorkItemAttachmentAccessRequest) Validate() error {
	return dara.Validate(s)
}
