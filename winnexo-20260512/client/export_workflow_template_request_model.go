// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportWorkflowTemplateRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppToken(v string) *ExportWorkflowTemplateRequest
  GetAppToken() *string 
  SetCollaborationGroupId(v string) *ExportWorkflowTemplateRequest
  GetCollaborationGroupId() *string 
  SetOpsUserCode(v string) *ExportWorkflowTemplateRequest
  GetOpsUserCode() *string 
  SetOpsUserId(v string) *ExportWorkflowTemplateRequest
  GetOpsUserId() *string 
  SetTemplateId(v string) *ExportWorkflowTemplateRequest
  GetTemplateId() *string 
  SetTenantId(v string) *ExportWorkflowTemplateRequest
  GetTenantId() *string 
  SetVersion(v string) *ExportWorkflowTemplateRequest
  GetVersion() *string 
}

type ExportWorkflowTemplateRequest struct {
  // 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
  // 
  // example:
  // 
  // wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
  AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
  // The ID of the collaboration workspace.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cg_9401
  CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
  // 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
  // 
  // example:
  // 
  // alice
  OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
  // 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
  // 
  // example:
  // 
  // 10001
  OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
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

func (s *ExportWorkflowTemplateRequest) GetAppToken() *string  {
  return s.AppToken
}

func (s *ExportWorkflowTemplateRequest) GetCollaborationGroupId() *string  {
  return s.CollaborationGroupId
}

func (s *ExportWorkflowTemplateRequest) GetOpsUserCode() *string  {
  return s.OpsUserCode
}

func (s *ExportWorkflowTemplateRequest) GetOpsUserId() *string  {
  return s.OpsUserId
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

func (s *ExportWorkflowTemplateRequest) SetAppToken(v string) *ExportWorkflowTemplateRequest {
  s.AppToken = &v
  return s
}

func (s *ExportWorkflowTemplateRequest) SetCollaborationGroupId(v string) *ExportWorkflowTemplateRequest {
  s.CollaborationGroupId = &v
  return s
}

func (s *ExportWorkflowTemplateRequest) SetOpsUserCode(v string) *ExportWorkflowTemplateRequest {
  s.OpsUserCode = &v
  return s
}

func (s *ExportWorkflowTemplateRequest) SetOpsUserId(v string) *ExportWorkflowTemplateRequest {
  s.OpsUserId = &v
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

