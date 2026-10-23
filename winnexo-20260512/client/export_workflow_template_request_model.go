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
  SetVersion(v string) *ExportWorkflowTemplateRequest
  GetVersion() *string 
}

type ExportWorkflowTemplateRequest struct {
  // The collaboration workspace ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // group_example
  CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
  // The template ID.
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
  // The version number of the template to export. The format is a lowercase v followed by a positive integer, such as v1. If not specified, the effective version is used. If no effective version exists, the latest version is used.
  // 
  // example:
  // 
  // v1
  Version *string `json:"version,omitempty" xml:"version,omitempty"`
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

func (s *ExportWorkflowTemplateRequest) GetVersion() *string  {
  return s.Version
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

func (s *ExportWorkflowTemplateRequest) SetVersion(v string) *ExportWorkflowTemplateRequest {
  s.Version = &v
  return s
}

func (s *ExportWorkflowTemplateRequest) Validate() error {
  return dara.Validate(s)
}

