// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportWorkflowTemplateRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCollaborationGroupId(v string) *ExportWorkflowTemplateRequest
  GetCollaborationGroupId() *string 
  SetTemplateId(v string) *ExportWorkflowTemplateRequest
  GetTemplateId() *string 
  SetTenantId(v string) *ExportWorkflowTemplateRequest
  GetTenantId() *string 
}

type ExportWorkflowTemplateRequest struct {
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

func (s ExportWorkflowTemplateRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportWorkflowTemplateRequest) GoString() string {
  return s.String()
}

func (s *ExportWorkflowTemplateRequest) GetCollaborationGroupId() *string  {
  return s.CollaborationGroupId
}

func (s *ExportWorkflowTemplateRequest) GetTemplateId() *string  {
  return s.TemplateId
}

func (s *ExportWorkflowTemplateRequest) GetTenantId() *string  {
  return s.TenantId
}

func (s *ExportWorkflowTemplateRequest) SetCollaborationGroupId(v string) *ExportWorkflowTemplateRequest {
  s.CollaborationGroupId = &v
  return s
}

func (s *ExportWorkflowTemplateRequest) SetTemplateId(v string) *ExportWorkflowTemplateRequest {
  s.TemplateId = &v
  return s
}

func (s *ExportWorkflowTemplateRequest) SetTenantId(v string) *ExportWorkflowTemplateRequest {
  s.TenantId = &v
  return s
}

func (s *ExportWorkflowTemplateRequest) Validate() error {
  return dara.Validate(s)
}

