// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptanceCriteria(v string) *CreateWorkflowInstanceShrinkRequest
	GetAcceptanceCriteria() *string
	SetBusinessId(v string) *CreateWorkflowInstanceShrinkRequest
	GetBusinessId() *string
	SetCollaborationGroupId(v string) *CreateWorkflowInstanceShrinkRequest
	GetCollaborationGroupId() *string
	SetInitiator(v string) *CreateWorkflowInstanceShrinkRequest
	GetInitiator() *string
	SetPriority(v string) *CreateWorkflowInstanceShrinkRequest
	GetPriority() *string
	SetTemplateId(v string) *CreateWorkflowInstanceShrinkRequest
	GetTemplateId() *string
	SetTenantId(v string) *CreateWorkflowInstanceShrinkRequest
	GetTenantId() *string
	SetTitle(v string) *CreateWorkflowInstanceShrinkRequest
	GetTitle() *string
	SetVariablesShrink(v string) *CreateWorkflowInstanceShrinkRequest
	GetVariablesShrink() *string
}

type CreateWorkflowInstanceShrinkRequest struct {
	// The required acceptance criteria for the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// Complete all approval nodes and generate the report
	AcceptanceCriteria *string `json:"acceptanceCriteria,omitempty" xml:"acceptanceCriteria,omitempty"`
	// The unique requirement ID in the workspace. Repeated requests are idempotent.
	//
	// This parameter is required.
	//
	// example:
	//
	// REQ-1001
	BusinessId *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	// The collaboration workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_20261016_abc123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The display label of the external initiator.
	//
	// example:
	//
	// External Initiator Zhang San
	Initiator *string `json:"initiator,omitempty" xml:"initiator,omitempty"`
	// The work item priority.
	//
	// example:
	//
	// HIGH
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// The workflow template ID. Maximum length: 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// wft_20261016_abc123
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The tenant ID. This is a common parameter explicitly passed by using --tenant-id in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The requirement title.
	//
	// This parameter is required.
	//
	// example:
	//
	// Quarterly Business Analysis Report
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The business variable object. The tags field can be a list of string tags.
	//
	// example:
	//
	// tags=Quarterly Report
	VariablesShrink *string `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s CreateWorkflowInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstanceShrinkRequest) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *CreateWorkflowInstanceShrinkRequest) GetBusinessId() *string {
	return s.BusinessId
}

func (s *CreateWorkflowInstanceShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateWorkflowInstanceShrinkRequest) GetInitiator() *string {
	return s.Initiator
}

func (s *CreateWorkflowInstanceShrinkRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateWorkflowInstanceShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateWorkflowInstanceShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateWorkflowInstanceShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateWorkflowInstanceShrinkRequest) GetVariablesShrink() *string {
	return s.VariablesShrink
}

func (s *CreateWorkflowInstanceShrinkRequest) SetAcceptanceCriteria(v string) *CreateWorkflowInstanceShrinkRequest {
	s.AcceptanceCriteria = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetBusinessId(v string) *CreateWorkflowInstanceShrinkRequest {
	s.BusinessId = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetCollaborationGroupId(v string) *CreateWorkflowInstanceShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetInitiator(v string) *CreateWorkflowInstanceShrinkRequest {
	s.Initiator = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetPriority(v string) *CreateWorkflowInstanceShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetTemplateId(v string) *CreateWorkflowInstanceShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetTenantId(v string) *CreateWorkflowInstanceShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetTitle(v string) *CreateWorkflowInstanceShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetVariablesShrink(v string) *CreateWorkflowInstanceShrinkRequest {
	s.VariablesShrink = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
