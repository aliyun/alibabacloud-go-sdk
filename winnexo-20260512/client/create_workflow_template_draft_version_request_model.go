// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowTemplateDraftVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateWorkflowTemplateDraftVersionRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *CreateWorkflowTemplateDraftVersionRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *CreateWorkflowTemplateDraftVersionRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateWorkflowTemplateDraftVersionRequest
	GetOpsUserId() *string
	SetTemplateId(v string) *CreateWorkflowTemplateDraftVersionRequest
	GetTemplateId() *string
	SetTenantId(v string) *CreateWorkflowTemplateDraftVersionRequest
	GetTenantId() *string
	SetVersion(v string) *CreateWorkflowTemplateDraftVersionRequest
	GetVersion() *string
}

type CreateWorkflowTemplateDraftVersionRequest struct {
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
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tpl_9401
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The version number to which the returned content belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateWorkflowTemplateDraftVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowTemplateDraftVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowTemplateDraftVersionRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateWorkflowTemplateDraftVersionRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateWorkflowTemplateDraftVersionRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateWorkflowTemplateDraftVersionRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateWorkflowTemplateDraftVersionRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateWorkflowTemplateDraftVersionRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateWorkflowTemplateDraftVersionRequest) GetVersion() *string {
	return s.Version
}

func (s *CreateWorkflowTemplateDraftVersionRequest) SetAppToken(v string) *CreateWorkflowTemplateDraftVersionRequest {
	s.AppToken = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionRequest) SetCollaborationGroupId(v string) *CreateWorkflowTemplateDraftVersionRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionRequest) SetOpsUserCode(v string) *CreateWorkflowTemplateDraftVersionRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionRequest) SetOpsUserId(v string) *CreateWorkflowTemplateDraftVersionRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionRequest) SetTemplateId(v string) *CreateWorkflowTemplateDraftVersionRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionRequest) SetTenantId(v string) *CreateWorkflowTemplateDraftVersionRequest {
	s.TenantId = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionRequest) SetVersion(v string) *CreateWorkflowTemplateDraftVersionRequest {
	s.Version = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionRequest) Validate() error {
	return dara.Validate(s)
}
