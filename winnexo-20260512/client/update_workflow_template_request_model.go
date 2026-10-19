// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *UpdateWorkflowTemplateRequest
	GetCollaborationGroupId() *string
	SetDefinition(v map[string]interface{}) *UpdateWorkflowTemplateRequest
	GetDefinition() map[string]interface{}
	SetTemplateId(v string) *UpdateWorkflowTemplateRequest
	GetTemplateId() *string
	SetTenantId(v string) *UpdateWorkflowTemplateRequest
	GetTenantId() *string
}

type UpdateWorkflowTemplateRequest struct {
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 完整模板定义 JSON（整体覆盖）
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Definition map[string]interface{} `json:"definition,omitempty" xml:"definition,omitempty"`
	// 模板 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// tpl_example
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateWorkflowTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowTemplateRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *UpdateWorkflowTemplateRequest) GetDefinition() map[string]interface{} {
	return s.Definition
}

func (s *UpdateWorkflowTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateWorkflowTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateWorkflowTemplateRequest) SetCollaborationGroupId(v string) *UpdateWorkflowTemplateRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetDefinition(v map[string]interface{}) *UpdateWorkflowTemplateRequest {
	s.Definition = v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetTemplateId(v string) *UpdateWorkflowTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetTenantId(v string) *UpdateWorkflowTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) Validate() error {
	return dara.Validate(s)
}
