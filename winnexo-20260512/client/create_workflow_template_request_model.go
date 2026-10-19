// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *CreateWorkflowTemplateRequest
	GetCollaborationGroupId() *string
	SetDefinition(v map[string]interface{}) *CreateWorkflowTemplateRequest
	GetDefinition() map[string]interface{}
	SetTenantId(v string) *CreateWorkflowTemplateRequest
	GetTenantId() *string
}

type CreateWorkflowTemplateRequest struct {
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 完整模板定义 JSON（name/description/steps）
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Definition map[string]interface{} `json:"definition,omitempty" xml:"definition,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
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

func (s *CreateWorkflowTemplateRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateWorkflowTemplateRequest) GetDefinition() map[string]interface{} {
	return s.Definition
}

func (s *CreateWorkflowTemplateRequest) GetTenantId() *string {
	return s.TenantId
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
