// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportWorkflowTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ImportWorkflowTemplateRequest
	GetCollaborationGroupId() *string
	SetTemplate(v map[string]interface{}) *ImportWorkflowTemplateRequest
	GetTemplate() map[string]interface{}
	SetTenantId(v string) *ImportWorkflowTemplateRequest
	GetTenantId() *string
}

type ImportWorkflowTemplateRequest struct {
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 可移植模板 JSON（stepId/sortOrder 可省略）
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Template map[string]interface{} `json:"template,omitempty" xml:"template,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ImportWorkflowTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportWorkflowTemplateRequest) GoString() string {
	return s.String()
}

func (s *ImportWorkflowTemplateRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ImportWorkflowTemplateRequest) GetTemplate() map[string]interface{} {
	return s.Template
}

func (s *ImportWorkflowTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ImportWorkflowTemplateRequest) SetCollaborationGroupId(v string) *ImportWorkflowTemplateRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ImportWorkflowTemplateRequest) SetTemplate(v map[string]interface{}) *ImportWorkflowTemplateRequest {
	s.Template = v
	return s
}

func (s *ImportWorkflowTemplateRequest) SetTenantId(v string) *ImportWorkflowTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *ImportWorkflowTemplateRequest) Validate() error {
	return dara.Validate(s)
}
