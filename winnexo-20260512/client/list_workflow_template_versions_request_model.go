// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowTemplateVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListWorkflowTemplateVersionsRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *ListWorkflowTemplateVersionsRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *ListWorkflowTemplateVersionsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListWorkflowTemplateVersionsRequest
	GetOpsUserId() *string
	SetTemplateId(v string) *ListWorkflowTemplateVersionsRequest
	GetTemplateId() *string
	SetTenantId(v string) *ListWorkflowTemplateVersionsRequest
	GetTenantId() *string
}

type ListWorkflowTemplateVersionsRequest struct {
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

func (s ListWorkflowTemplateVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTemplateVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowTemplateVersionsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListWorkflowTemplateVersionsRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListWorkflowTemplateVersionsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListWorkflowTemplateVersionsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListWorkflowTemplateVersionsRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListWorkflowTemplateVersionsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListWorkflowTemplateVersionsRequest) SetAppToken(v string) *ListWorkflowTemplateVersionsRequest {
	s.AppToken = &v
	return s
}

func (s *ListWorkflowTemplateVersionsRequest) SetCollaborationGroupId(v string) *ListWorkflowTemplateVersionsRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListWorkflowTemplateVersionsRequest) SetOpsUserCode(v string) *ListWorkflowTemplateVersionsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListWorkflowTemplateVersionsRequest) SetOpsUserId(v string) *ListWorkflowTemplateVersionsRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListWorkflowTemplateVersionsRequest) SetTemplateId(v string) *ListWorkflowTemplateVersionsRequest {
	s.TemplateId = &v
	return s
}

func (s *ListWorkflowTemplateVersionsRequest) SetTenantId(v string) *ListWorkflowTemplateVersionsRequest {
	s.TenantId = &v
	return s
}

func (s *ListWorkflowTemplateVersionsRequest) Validate() error {
	return dara.Validate(s)
}
