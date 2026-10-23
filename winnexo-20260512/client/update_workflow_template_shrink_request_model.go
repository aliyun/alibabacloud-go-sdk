// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *UpdateWorkflowTemplateShrinkRequest
	GetAction() *string
	SetCollaborationGroupId(v string) *UpdateWorkflowTemplateShrinkRequest
	GetCollaborationGroupId() *string
	SetDefinitionShrink(v string) *UpdateWorkflowTemplateShrinkRequest
	GetDefinitionShrink() *string
	SetSourceVersion(v string) *UpdateWorkflowTemplateShrinkRequest
	GetSourceVersion() *string
	SetTemplateId(v string) *UpdateWorkflowTemplateShrinkRequest
	GetTemplateId() *string
	SetTenantId(v string) *UpdateWorkflowTemplateShrinkRequest
	GetTenantId() *string
	SetVersion(v string) *UpdateWorkflowTemplateShrinkRequest
	GetVersion() *string
	SetVersionDescription(v string) *UpdateWorkflowTemplateShrinkRequest
	GetVersionDescription() *string
}

type UpdateWorkflowTemplateShrinkRequest struct {
	// The write action. Valid values:
	//
	// - SAVE: Saves the template as a draft without taking effect.
	//
	// - SUBMIT: Saves and submits the template, setting it as the only effective version.
	//
	// Default value: SUBMIT.
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
	// The complete JSON definition of the template to overwrite the existing one.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	DefinitionShrink *string `json:"definition,omitempty" xml:"definition,omitempty"`
	// The submitted version number on which the new version is based. This parameter is mutually exclusive with the version parameter. If not specified, the current effective version is used as the fallback.
	//
	// example:
	//
	// v1
	SourceVersion *string `json:"sourceVersion,omitempty" xml:"sourceVersion,omitempty"`
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
	// The draft template version number to write. The format is a lowercase v followed by a positive integer, such as v1. Submitted versions cannot be written. If not specified, a new version is generated.
	//
	// example:
	//
	// v2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The version description. The maximum length is 500 characters. When writing a draft, if this parameter is not specified or is null, the original value is retained. An empty string clears the value.
	//
	// example:
	//
	// Add a review step
	VersionDescription *string `json:"versionDescription,omitempty" xml:"versionDescription,omitempty"`
}

func (s UpdateWorkflowTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowTemplateShrinkRequest) GetAction() *string {
	return s.Action
}

func (s *UpdateWorkflowTemplateShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *UpdateWorkflowTemplateShrinkRequest) GetDefinitionShrink() *string {
	return s.DefinitionShrink
}

func (s *UpdateWorkflowTemplateShrinkRequest) GetSourceVersion() *string {
	return s.SourceVersion
}

func (s *UpdateWorkflowTemplateShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateWorkflowTemplateShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateWorkflowTemplateShrinkRequest) GetVersion() *string {
	return s.Version
}

func (s *UpdateWorkflowTemplateShrinkRequest) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *UpdateWorkflowTemplateShrinkRequest) SetAction(v string) *UpdateWorkflowTemplateShrinkRequest {
	s.Action = &v
	return s
}

func (s *UpdateWorkflowTemplateShrinkRequest) SetCollaborationGroupId(v string) *UpdateWorkflowTemplateShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *UpdateWorkflowTemplateShrinkRequest) SetDefinitionShrink(v string) *UpdateWorkflowTemplateShrinkRequest {
	s.DefinitionShrink = &v
	return s
}

func (s *UpdateWorkflowTemplateShrinkRequest) SetSourceVersion(v string) *UpdateWorkflowTemplateShrinkRequest {
	s.SourceVersion = &v
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

func (s *UpdateWorkflowTemplateShrinkRequest) SetVersion(v string) *UpdateWorkflowTemplateShrinkRequest {
	s.Version = &v
	return s
}

func (s *UpdateWorkflowTemplateShrinkRequest) SetVersionDescription(v string) *UpdateWorkflowTemplateShrinkRequest {
	s.VersionDescription = &v
	return s
}

func (s *UpdateWorkflowTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
