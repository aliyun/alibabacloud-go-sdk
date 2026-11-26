// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *DeleteWorkflowTemplateRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *DeleteWorkflowTemplateRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *DeleteWorkflowTemplateRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *DeleteWorkflowTemplateRequest
	GetOpsUserId() *string
	SetTemplateId(v string) *DeleteWorkflowTemplateRequest
	GetTemplateId() *string
	SetTenantId(v string) *DeleteWorkflowTemplateRequest
	GetTenantId() *string
}

type DeleteWorkflowTemplateRequest struct {
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
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteWorkflowTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowTemplateRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *DeleteWorkflowTemplateRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *DeleteWorkflowTemplateRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *DeleteWorkflowTemplateRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *DeleteWorkflowTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteWorkflowTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteWorkflowTemplateRequest) SetAppToken(v string) *DeleteWorkflowTemplateRequest {
	s.AppToken = &v
	return s
}

func (s *DeleteWorkflowTemplateRequest) SetCollaborationGroupId(v string) *DeleteWorkflowTemplateRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *DeleteWorkflowTemplateRequest) SetOpsUserCode(v string) *DeleteWorkflowTemplateRequest {
	s.OpsUserCode = &v
	return s
}

func (s *DeleteWorkflowTemplateRequest) SetOpsUserId(v string) *DeleteWorkflowTemplateRequest {
	s.OpsUserId = &v
	return s
}

func (s *DeleteWorkflowTemplateRequest) SetTemplateId(v string) *DeleteWorkflowTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteWorkflowTemplateRequest) SetTenantId(v string) *DeleteWorkflowTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteWorkflowTemplateRequest) Validate() error {
	return dara.Validate(s)
}
