// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetWorkflowTemplateRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *GetWorkflowTemplateRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *GetWorkflowTemplateRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetWorkflowTemplateRequest
	GetOpsUserId() *string
	SetTemplateId(v string) *GetWorkflowTemplateRequest
	GetTemplateId() *string
	SetTenantId(v string) *GetWorkflowTemplateRequest
	GetTenantId() *string
	SetVersion(v string) *GetWorkflowTemplateRequest
	GetVersion() *string
}

type GetWorkflowTemplateRequest struct {
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
	// The ID of the template.
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
	// The version number of the template to read. The format is a lowercase v followed by a positive integer, such as v1. If not specified, the effective version is used. If no effective version exists, the latest version is used.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetWorkflowTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowTemplateRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetWorkflowTemplateRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkflowTemplateRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetWorkflowTemplateRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetWorkflowTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetWorkflowTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkflowTemplateRequest) GetVersion() *string {
	return s.Version
}

func (s *GetWorkflowTemplateRequest) SetAppToken(v string) *GetWorkflowTemplateRequest {
	s.AppToken = &v
	return s
}

func (s *GetWorkflowTemplateRequest) SetCollaborationGroupId(v string) *GetWorkflowTemplateRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkflowTemplateRequest) SetOpsUserCode(v string) *GetWorkflowTemplateRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetWorkflowTemplateRequest) SetOpsUserId(v string) *GetWorkflowTemplateRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetWorkflowTemplateRequest) SetTemplateId(v string) *GetWorkflowTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetWorkflowTemplateRequest) SetTenantId(v string) *GetWorkflowTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkflowTemplateRequest) SetVersion(v string) *GetWorkflowTemplateRequest {
	s.Version = &v
	return s
}

func (s *GetWorkflowTemplateRequest) Validate() error {
	return dara.Validate(s)
}
