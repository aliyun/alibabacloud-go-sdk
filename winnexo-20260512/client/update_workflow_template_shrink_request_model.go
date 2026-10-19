// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *UpdateWorkflowTemplateShrinkRequest
	GetCollaborationGroupId() *string
	SetDefinitionShrink(v string) *UpdateWorkflowTemplateShrinkRequest
	GetDefinitionShrink() *string
	SetTemplateId(v string) *UpdateWorkflowTemplateShrinkRequest
	GetTemplateId() *string
	SetTenantId(v string) *UpdateWorkflowTemplateShrinkRequest
	GetTenantId() *string
}

type UpdateWorkflowTemplateShrinkRequest struct {
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
	DefinitionShrink *string `json:"definition,omitempty" xml:"definition,omitempty"`
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

func (s UpdateWorkflowTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowTemplateShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *UpdateWorkflowTemplateShrinkRequest) GetDefinitionShrink() *string {
	return s.DefinitionShrink
}

func (s *UpdateWorkflowTemplateShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateWorkflowTemplateShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateWorkflowTemplateShrinkRequest) SetCollaborationGroupId(v string) *UpdateWorkflowTemplateShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *UpdateWorkflowTemplateShrinkRequest) SetDefinitionShrink(v string) *UpdateWorkflowTemplateShrinkRequest {
	s.DefinitionShrink = &v
	return s
}

func (s *UpdateWorkflowTemplateShrinkRequest) SetTemplateId(v string) *UpdateWorkflowTemplateShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateWorkflowTemplateShrinkRequest) SetTenantId(v string) *UpdateWorkflowTemplateShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateWorkflowTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
