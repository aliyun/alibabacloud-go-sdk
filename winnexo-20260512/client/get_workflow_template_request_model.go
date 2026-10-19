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
}

type GetWorkflowTemplateRequest struct {
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 模板 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// tpl_example
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
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

func (s *GetWorkflowTemplateRequest) Validate() error {
	return dara.Validate(s)
}
