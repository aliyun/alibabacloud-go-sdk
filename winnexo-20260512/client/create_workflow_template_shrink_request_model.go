// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *CreateWorkflowTemplateShrinkRequest
	GetCollaborationGroupId() *string
	SetDefinitionShrink(v string) *CreateWorkflowTemplateShrinkRequest
	GetDefinitionShrink() *string
	SetTenantId(v string) *CreateWorkflowTemplateShrinkRequest
	GetTenantId() *string
}

type CreateWorkflowTemplateShrinkRequest struct {
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 完整模板定义 JSON（name/description/steps）
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	DefinitionShrink *string `json:"definition,omitempty" xml:"definition,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateWorkflowTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowTemplateShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateWorkflowTemplateShrinkRequest) GetDefinitionShrink() *string {
	return s.DefinitionShrink
}

func (s *CreateWorkflowTemplateShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateWorkflowTemplateShrinkRequest) SetCollaborationGroupId(v string) *CreateWorkflowTemplateShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateWorkflowTemplateShrinkRequest) SetDefinitionShrink(v string) *CreateWorkflowTemplateShrinkRequest {
	s.DefinitionShrink = &v
	return s
}

func (s *CreateWorkflowTemplateShrinkRequest) SetTenantId(v string) *CreateWorkflowTemplateShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateWorkflowTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
