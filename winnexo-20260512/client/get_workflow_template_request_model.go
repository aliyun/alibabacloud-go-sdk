// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *GetWorkflowTemplateRequest
	GetCollaborationGroupId() *string
	SetTemplateId(v string) *GetWorkflowTemplateRequest
	GetTemplateId() *string
	SetTenantId(v string) *GetWorkflowTemplateRequest
	GetTenantId() *string
	SetVersion(v string) *GetWorkflowTemplateRequest
	GetVersion() *string
}

type GetWorkflowTemplateRequest struct {
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
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The version number of the template to read. The format is a lowercase v followed by a positive integer, such as v1. If not specified, the effective version is used. If no effective version exists, the latest version is used.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetWorkflowTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowTemplateRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkflowTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetWorkflowTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkflowTemplateRequest) GetVersion() *string {
	return s.Version
}

func (s *GetWorkflowTemplateRequest) SetCollaborationGroupId(v string) *GetWorkflowTemplateRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkflowTemplateRequest) SetTemplateId(v string) *GetWorkflowTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetWorkflowTemplateRequest) SetTenantId(v string) *GetWorkflowTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkflowTemplateRequest) SetVersion(v string) *GetWorkflowTemplateRequest {
	s.Version = &v
	return s
}

func (s *GetWorkflowTemplateRequest) Validate() error {
	return dara.Validate(s)
}
