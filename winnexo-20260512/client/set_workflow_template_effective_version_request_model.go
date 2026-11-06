// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkflowTemplateEffectiveVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *SetWorkflowTemplateEffectiveVersionRequest
	GetCollaborationGroupId() *string
	SetTemplateId(v string) *SetWorkflowTemplateEffectiveVersionRequest
	GetTemplateId() *string
	SetTenantId(v string) *SetWorkflowTemplateEffectiveVersionRequest
	GetTenantId() *string
	SetVersion(v string) *SetWorkflowTemplateEffectiveVersionRequest
	GetVersion() *string
}

type SetWorkflowTemplateEffectiveVersionRequest struct {
	// The collaboration workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tpl_9401
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The version number of the returned content.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SetWorkflowTemplateEffectiveVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetWorkflowTemplateEffectiveVersionRequest) GoString() string {
	return s.String()
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) GetVersion() *string {
	return s.Version
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) SetCollaborationGroupId(v string) *SetWorkflowTemplateEffectiveVersionRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) SetTemplateId(v string) *SetWorkflowTemplateEffectiveVersionRequest {
	s.TemplateId = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) SetTenantId(v string) *SetWorkflowTemplateEffectiveVersionRequest {
	s.TenantId = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) SetVersion(v string) *SetWorkflowTemplateEffectiveVersionRequest {
	s.Version = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) Validate() error {
	return dara.Validate(s)
}
