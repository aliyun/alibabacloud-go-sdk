// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *DeleteWorkflowTemplateRequest
	GetCollaborationGroupId() *string
	SetTemplateId(v string) *DeleteWorkflowTemplateRequest
	GetTemplateId() *string
	SetTenantId(v string) *DeleteWorkflowTemplateRequest
	GetTenantId() *string
}

type DeleteWorkflowTemplateRequest struct {
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

func (s DeleteWorkflowTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowTemplateRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *DeleteWorkflowTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteWorkflowTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteWorkflowTemplateRequest) SetCollaborationGroupId(v string) *DeleteWorkflowTemplateRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *DeleteWorkflowTemplateRequest) SetTemplateId(v string) *DeleteWorkflowTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteWorkflowTemplateRequest) SetTenantId(v string) *DeleteWorkflowTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteWorkflowTemplateRequest) Validate() error {
	return dara.Validate(s)
}
