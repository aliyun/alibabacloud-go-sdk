// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowTemplateDraftVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *CreateWorkflowTemplateDraftVersionRequest
	GetCollaborationGroupId() *string
	SetTemplateId(v string) *CreateWorkflowTemplateDraftVersionRequest
	GetTemplateId() *string
	SetTenantId(v string) *CreateWorkflowTemplateDraftVersionRequest
	GetTenantId() *string
	SetVersion(v string) *CreateWorkflowTemplateDraftVersionRequest
	GetVersion() *string
}

type CreateWorkflowTemplateDraftVersionRequest struct {
	// The ID of the collaboration workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
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
	// The version number of the template on which the edit is based. The format is a lowercase v followed by a positive integer, such as v1. A draft version is returned as-is.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateWorkflowTemplateDraftVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowTemplateDraftVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowTemplateDraftVersionRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateWorkflowTemplateDraftVersionRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateWorkflowTemplateDraftVersionRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateWorkflowTemplateDraftVersionRequest) GetVersion() *string {
	return s.Version
}

func (s *CreateWorkflowTemplateDraftVersionRequest) SetCollaborationGroupId(v string) *CreateWorkflowTemplateDraftVersionRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionRequest) SetTemplateId(v string) *CreateWorkflowTemplateDraftVersionRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionRequest) SetTenantId(v string) *CreateWorkflowTemplateDraftVersionRequest {
	s.TenantId = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionRequest) SetVersion(v string) *CreateWorkflowTemplateDraftVersionRequest {
	s.Version = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionRequest) Validate() error {
	return dara.Validate(s)
}
