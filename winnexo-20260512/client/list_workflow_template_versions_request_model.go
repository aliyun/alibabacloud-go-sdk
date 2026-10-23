// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowTemplateVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ListWorkflowTemplateVersionsRequest
	GetCollaborationGroupId() *string
	SetTemplateId(v string) *ListWorkflowTemplateVersionsRequest
	GetTemplateId() *string
	SetTenantId(v string) *ListWorkflowTemplateVersionsRequest
	GetTenantId() *string
}

type ListWorkflowTemplateVersionsRequest struct {
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
}

func (s ListWorkflowTemplateVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTemplateVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowTemplateVersionsRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListWorkflowTemplateVersionsRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListWorkflowTemplateVersionsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListWorkflowTemplateVersionsRequest) SetCollaborationGroupId(v string) *ListWorkflowTemplateVersionsRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListWorkflowTemplateVersionsRequest) SetTemplateId(v string) *ListWorkflowTemplateVersionsRequest {
	s.TemplateId = &v
	return s
}

func (s *ListWorkflowTemplateVersionsRequest) SetTenantId(v string) *ListWorkflowTemplateVersionsRequest {
	s.TenantId = &v
	return s
}

func (s *ListWorkflowTemplateVersionsRequest) Validate() error {
	return dara.Validate(s)
}
