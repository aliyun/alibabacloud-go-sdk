// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResolveWorkItemAttachmentAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ResolveWorkItemAttachmentAccessRequest
	GetAppToken() *string
	SetAttachmentId(v string) *ResolveWorkItemAttachmentAccessRequest
	GetAttachmentId() *string
	SetCollaborationGroupId(v string) *ResolveWorkItemAttachmentAccessRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *ResolveWorkItemAttachmentAccessRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ResolveWorkItemAttachmentAccessRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ResolveWorkItemAttachmentAccessRequest
	GetTenantId() *string
	SetWorkItemId(v string) *ResolveWorkItemAttachmentAccessRequest
	GetWorkItemId() *string
}

type ResolveWorkItemAttachmentAccessRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the attachment.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	AttachmentId *string `json:"attachmentId,omitempty" xml:"attachmentId,omitempty"`
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

func (s ResolveWorkItemAttachmentAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s ResolveWorkItemAttachmentAccessRequest) GoString() string {
	return s.String()
}

func (s *ResolveWorkItemAttachmentAccessRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ResolveWorkItemAttachmentAccessRequest) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *ResolveWorkItemAttachmentAccessRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ResolveWorkItemAttachmentAccessRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ResolveWorkItemAttachmentAccessRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ResolveWorkItemAttachmentAccessRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ResolveWorkItemAttachmentAccessRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *ResolveWorkItemAttachmentAccessRequest) SetAppToken(v string) *ResolveWorkItemAttachmentAccessRequest {
	s.AppToken = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessRequest) SetAttachmentId(v string) *ResolveWorkItemAttachmentAccessRequest {
	s.AttachmentId = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessRequest) SetCollaborationGroupId(v string) *ResolveWorkItemAttachmentAccessRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessRequest) SetOpsUserCode(v string) *ResolveWorkItemAttachmentAccessRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessRequest) SetOpsUserId(v string) *ResolveWorkItemAttachmentAccessRequest {
	s.OpsUserId = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessRequest) SetTenantId(v string) *ResolveWorkItemAttachmentAccessRequest {
	s.TenantId = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessRequest) SetWorkItemId(v string) *ResolveWorkItemAttachmentAccessRequest {
	s.WorkItemId = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessRequest) Validate() error {
	return dara.Validate(s)
}
