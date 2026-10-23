// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *UpdateWorkflowTemplateRequest
	GetAction() *string
	SetCollaborationGroupId(v string) *UpdateWorkflowTemplateRequest
	GetCollaborationGroupId() *string
	SetDefinition(v map[string]interface{}) *UpdateWorkflowTemplateRequest
	GetDefinition() map[string]interface{}
	SetSourceVersion(v string) *UpdateWorkflowTemplateRequest
	GetSourceVersion() *string
	SetTemplateId(v string) *UpdateWorkflowTemplateRequest
	GetTemplateId() *string
	SetTenantId(v string) *UpdateWorkflowTemplateRequest
	GetTenantId() *string
	SetVersion(v string) *UpdateWorkflowTemplateRequest
	GetVersion() *string
	SetVersionDescription(v string) *UpdateWorkflowTemplateRequest
	GetVersionDescription() *string
}

type UpdateWorkflowTemplateRequest struct {
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
	Definition map[string]interface{} `json:"definition,omitempty" xml:"definition,omitempty"`
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

func (s UpdateWorkflowTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowTemplateRequest) GetAction() *string {
	return s.Action
}

func (s *UpdateWorkflowTemplateRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *UpdateWorkflowTemplateRequest) GetDefinition() map[string]interface{} {
	return s.Definition
}

func (s *UpdateWorkflowTemplateRequest) GetSourceVersion() *string {
	return s.SourceVersion
}

func (s *UpdateWorkflowTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateWorkflowTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateWorkflowTemplateRequest) GetVersion() *string {
	return s.Version
}

func (s *UpdateWorkflowTemplateRequest) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *UpdateWorkflowTemplateRequest) SetAction(v string) *UpdateWorkflowTemplateRequest {
	s.Action = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetCollaborationGroupId(v string) *UpdateWorkflowTemplateRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetDefinition(v map[string]interface{}) *UpdateWorkflowTemplateRequest {
	s.Definition = v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetSourceVersion(v string) *UpdateWorkflowTemplateRequest {
	s.SourceVersion = &v
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

func (s *UpdateWorkflowTemplateRequest) SetVersion(v string) *UpdateWorkflowTemplateRequest {
	s.Version = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetVersionDescription(v string) *UpdateWorkflowTemplateRequest {
	s.VersionDescription = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) Validate() error {
	return dara.Validate(s)
}
