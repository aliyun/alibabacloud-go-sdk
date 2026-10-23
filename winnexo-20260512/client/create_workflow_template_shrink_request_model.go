// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *CreateWorkflowTemplateShrinkRequest
	GetAction() *string
	SetCollaborationGroupId(v string) *CreateWorkflowTemplateShrinkRequest
	GetCollaborationGroupId() *string
	SetDefinitionShrink(v string) *CreateWorkflowTemplateShrinkRequest
	GetDefinitionShrink() *string
	SetTenantId(v string) *CreateWorkflowTemplateShrinkRequest
	GetTenantId() *string
}

type CreateWorkflowTemplateShrinkRequest struct {
	// The write action. Valid values: SAVE: Saves the template as a draft without taking effect. SUBMIT: Saves and submits the template and sets it as the only effective version. Default value: SUBMIT.
	//
	// example:
	//
	// SUBMIT
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The ID of the collaboration workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The complete JSON definition of the template, including name, description, and steps.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	DefinitionShrink *string `json:"definition,omitempty" xml:"definition,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
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

func (s *CreateWorkflowTemplateShrinkRequest) GetAction() *string {
	return s.Action
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

func (s *CreateWorkflowTemplateShrinkRequest) SetAction(v string) *CreateWorkflowTemplateShrinkRequest {
	s.Action = &v
	return s
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
