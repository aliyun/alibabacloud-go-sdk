// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowTemplateVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *DeleteWorkflowTemplateVersionRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *DeleteWorkflowTemplateVersionRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *DeleteWorkflowTemplateVersionRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *DeleteWorkflowTemplateVersionRequest
	GetOpsUserId() *string
	SetTemplateId(v string) *DeleteWorkflowTemplateVersionRequest
	GetTemplateId() *string
	SetTenantId(v string) *DeleteWorkflowTemplateVersionRequest
	GetTenantId() *string
	SetVersion(v string) *DeleteWorkflowTemplateVersionRequest
	GetVersion() *string
}

type DeleteWorkflowTemplateVersionRequest struct {
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
	// The ID of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// tpl_example
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The version number of the template to be deleted. The format is a lowercase v followed by a positive integer, such as v1. The version cannot be the current active version.
	//
	// This parameter is required.
	//
	// example:
	//
	// v2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DeleteWorkflowTemplateVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowTemplateVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowTemplateVersionRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *DeleteWorkflowTemplateVersionRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *DeleteWorkflowTemplateVersionRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *DeleteWorkflowTemplateVersionRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *DeleteWorkflowTemplateVersionRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteWorkflowTemplateVersionRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteWorkflowTemplateVersionRequest) GetVersion() *string {
	return s.Version
}

func (s *DeleteWorkflowTemplateVersionRequest) SetAppToken(v string) *DeleteWorkflowTemplateVersionRequest {
	s.AppToken = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionRequest) SetCollaborationGroupId(v string) *DeleteWorkflowTemplateVersionRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionRequest) SetOpsUserCode(v string) *DeleteWorkflowTemplateVersionRequest {
	s.OpsUserCode = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionRequest) SetOpsUserId(v string) *DeleteWorkflowTemplateVersionRequest {
	s.OpsUserId = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionRequest) SetTemplateId(v string) *DeleteWorkflowTemplateVersionRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionRequest) SetTenantId(v string) *DeleteWorkflowTemplateVersionRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionRequest) SetVersion(v string) *DeleteWorkflowTemplateVersionRequest {
	s.Version = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionRequest) Validate() error {
	return dara.Validate(s)
}
