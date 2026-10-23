// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowTemplateVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *DeleteWorkflowTemplateVersionRequest
	GetCollaborationGroupId() *string
	SetTemplateId(v string) *DeleteWorkflowTemplateVersionRequest
	GetTemplateId() *string
	SetTenantId(v string) *DeleteWorkflowTemplateVersionRequest
	GetTenantId() *string
	SetVersion(v string) *DeleteWorkflowTemplateVersionRequest
	GetVersion() *string
}

type DeleteWorkflowTemplateVersionRequest struct {
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
	// The version number of the template to be deleted. The format is a lowercase v followed by a positive integer, such as v1. The version cannot be the current active version.
	//
	// This parameter is required.
	//
	// example:
	//
	// v2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DeleteWorkflowTemplateVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowTemplateVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowTemplateVersionRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *DeleteWorkflowTemplateVersionRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteWorkflowTemplateVersionRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteWorkflowTemplateVersionRequest) GetVersion() *string {
	return s.Version
}

func (s *DeleteWorkflowTemplateVersionRequest) SetCollaborationGroupId(v string) *DeleteWorkflowTemplateVersionRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionRequest) SetTemplateId(v string) *DeleteWorkflowTemplateVersionRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionRequest) SetTenantId(v string) *DeleteWorkflowTemplateVersionRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionRequest) SetVersion(v string) *DeleteWorkflowTemplateVersionRequest {
	s.Version = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionRequest) Validate() error {
	return dara.Validate(s)
}
