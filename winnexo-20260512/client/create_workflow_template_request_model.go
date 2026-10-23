// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *CreateWorkflowTemplateRequest
	GetAction() *string
	SetCollaborationGroupId(v string) *CreateWorkflowTemplateRequest
	GetCollaborationGroupId() *string
	SetDefinition(v map[string]interface{}) *CreateWorkflowTemplateRequest
	GetDefinition() map[string]interface{}
	SetTenantId(v string) *CreateWorkflowTemplateRequest
	GetTenantId() *string
}

type CreateWorkflowTemplateRequest struct {
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
	Definition map[string]interface{} `json:"definition,omitempty" xml:"definition,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateWorkflowTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowTemplateRequest) GetAction() *string {
	return s.Action
}

func (s *CreateWorkflowTemplateRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateWorkflowTemplateRequest) GetDefinition() map[string]interface{} {
	return s.Definition
}

func (s *CreateWorkflowTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateWorkflowTemplateRequest) SetAction(v string) *CreateWorkflowTemplateRequest {
	s.Action = &v
	return s
}

func (s *CreateWorkflowTemplateRequest) SetCollaborationGroupId(v string) *CreateWorkflowTemplateRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateWorkflowTemplateRequest) SetDefinition(v map[string]interface{}) *CreateWorkflowTemplateRequest {
	s.Definition = v
	return s
}

func (s *CreateWorkflowTemplateRequest) SetTenantId(v string) *CreateWorkflowTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *CreateWorkflowTemplateRequest) Validate() error {
	return dara.Validate(s)
}
