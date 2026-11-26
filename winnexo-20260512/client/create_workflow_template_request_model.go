// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *CreateWorkflowTemplateRequest
	GetAction() *string
	SetAppToken(v string) *CreateWorkflowTemplateRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *CreateWorkflowTemplateRequest
	GetCollaborationGroupId() *string
	SetDefinition(v map[string]interface{}) *CreateWorkflowTemplateRequest
	GetDefinition() map[string]interface{}
	SetOpsUserCode(v string) *CreateWorkflowTemplateRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateWorkflowTemplateRequest
	GetOpsUserId() *string
	SetTenantId(v string) *CreateWorkflowTemplateRequest
	GetTenantId() *string
}

type CreateWorkflowTemplateRequest struct {
	// The write action. Valid values:
	//
	// - SAVE: Saves the template as a draft. The draft does not take effect. Node names and assignees can be left empty.
	//
	// - SUBMIT: Requires complete nodes, and saves and submits the template as the only effective version.
	//
	// Default value: SUBMIT.
	//
	// example:
	//
	// SUBMIT
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
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
	// The complete JSON definition of the template, including name, description, and steps.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"name":"Customer Complaint Handling Process","schemaVersion":2,"steps":[{"stepId":"collect","sortOrder":1,"name":"Collect Customer Information"}]}
	Definition map[string]interface{} `json:"definition,omitempty" xml:"definition,omitempty"`
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
}

func (s CreateWorkflowTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowTemplateRequest) GetAction() *string {
	return s.Action
}

func (s *CreateWorkflowTemplateRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateWorkflowTemplateRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateWorkflowTemplateRequest) GetDefinition() map[string]interface{} {
	return s.Definition
}

func (s *CreateWorkflowTemplateRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateWorkflowTemplateRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateWorkflowTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateWorkflowTemplateRequest) SetAction(v string) *CreateWorkflowTemplateRequest {
	s.Action = &v
	return s
}

func (s *CreateWorkflowTemplateRequest) SetAppToken(v string) *CreateWorkflowTemplateRequest {
	s.AppToken = &v
	return s
}

func (s *CreateWorkflowTemplateRequest) SetCollaborationGroupId(v string) *CreateWorkflowTemplateRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateWorkflowTemplateRequest) SetDefinition(v map[string]interface{}) *CreateWorkflowTemplateRequest {
	s.Definition = v
	return s
}

func (s *CreateWorkflowTemplateRequest) SetOpsUserCode(v string) *CreateWorkflowTemplateRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateWorkflowTemplateRequest) SetOpsUserId(v string) *CreateWorkflowTemplateRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateWorkflowTemplateRequest) SetTenantId(v string) *CreateWorkflowTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *CreateWorkflowTemplateRequest) Validate() error {
	return dara.Validate(s)
}
