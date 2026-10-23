// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportWorkflowTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ImportWorkflowTemplateShrinkRequest
	GetCollaborationGroupId() *string
	SetTemplateShrink(v string) *ImportWorkflowTemplateShrinkRequest
	GetTemplateShrink() *string
	SetTenantId(v string) *ImportWorkflowTemplateShrinkRequest
	GetTenantId() *string
}

type ImportWorkflowTemplateShrinkRequest struct {
	// The collaboration workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The portable template JSON. The stepId and sortOrder fields can be omitted.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	TemplateShrink *string `json:"template,omitempty" xml:"template,omitempty"`
	// The tenant ID. This is a common parameter. If omitted, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ImportWorkflowTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportWorkflowTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportWorkflowTemplateShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ImportWorkflowTemplateShrinkRequest) GetTemplateShrink() *string {
	return s.TemplateShrink
}

func (s *ImportWorkflowTemplateShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ImportWorkflowTemplateShrinkRequest) SetCollaborationGroupId(v string) *ImportWorkflowTemplateShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ImportWorkflowTemplateShrinkRequest) SetTemplateShrink(v string) *ImportWorkflowTemplateShrinkRequest {
	s.TemplateShrink = &v
	return s
}

func (s *ImportWorkflowTemplateShrinkRequest) SetTenantId(v string) *ImportWorkflowTemplateShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ImportWorkflowTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
