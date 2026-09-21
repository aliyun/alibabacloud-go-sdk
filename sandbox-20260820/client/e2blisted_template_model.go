// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BListedTemplate interface {
  dara.Model
  String() string
  GoString() string
  SetBuildStatus(v string) *E2BListedTemplate
  GetBuildStatus() *string 
  SetCategory(v string) *E2BListedTemplate
  GetCategory() *string 
  SetContainerConfiguration(v *ContainerConfiguration) *E2BListedTemplate
  GetContainerConfiguration() *ContainerConfiguration 
  SetCpuCount(v int32) *E2BListedTemplate
  GetCpuCount() *int32 
  SetCreatedAt(v string) *E2BListedTemplate
  GetCreatedAt() *string 
  SetFunctionName(v string) *E2BListedTemplate
  GetFunctionName() *string 
  SetLogConfiguration(v *LogConfiguration) *E2BListedTemplate
  GetLogConfiguration() *LogConfiguration 
  SetMemoryMB(v int32) *E2BListedTemplate
  GetMemoryMB() *int32 
  SetNames(v []*string) *E2BListedTemplate
  GetNames() []*string 
  SetPublic(v bool) *E2BListedTemplate
  GetPublic() *bool 
  SetResourceGroupID(v string) *E2BListedTemplate
  GetResourceGroupID() *string 
  SetStatusReason(v string) *E2BListedTemplate
  GetStatusReason() *string 
  SetTags(v []*E2BTemplateTag) *E2BListedTemplate
  GetTags() []*E2BTemplateTag 
  SetTeamID(v string) *E2BListedTemplate
  GetTeamID() *string 
  SetTeamName(v string) *E2BListedTemplate
  GetTeamName() *string 
  SetTeamPlan(v string) *E2BListedTemplate
  GetTeamPlan() *string 
  SetTemplateID(v string) *E2BListedTemplate
  GetTemplateID() *string 
  SetUpdatedAt(v string) *E2BListedTemplate
  GetUpdatedAt() *string 
  SetUserID(v string) *E2BListedTemplate
  GetUserID() *string 
}

type E2BListedTemplate struct {
  // example:
  // 
  // ready
  BuildStatus *string `json:"buildStatus,omitempty" xml:"buildStatus,omitempty"`
  // example:
  // 
  // custom
  Category *string `json:"category,omitempty" xml:"category,omitempty"`
  ContainerConfiguration *ContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty"`
  // example:
  // 
  // 2
  CpuCount *int32 `json:"cpuCount,omitempty" xml:"cpuCount,omitempty"`
  // example:
  // 
  // 2026-08-20T08:30:00Z
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
  // example:
  // 
  // sandbox-tm-8f3a2c7b5e14d806
  FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
  LogConfiguration *LogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
  // example:
  // 
  // 2048
  MemoryMB *int32 `json:"memoryMB,omitempty" xml:"memoryMB,omitempty"`
  Names []*string `json:"names,omitempty" xml:"names,omitempty" type:"Repeated"`
  Public *bool `json:"public,omitempty" xml:"public,omitempty"`
  // example:
  // 
  // rg-****
  ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
  // example:
  // 
  // 拉取源镜像失败：认证信息无效
  StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
  Tags []*E2BTemplateTag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
  // example:
  // 
  // 5f4a2c18-****
  TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
  // example:
  // 
  // sandbox-dev
  TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
  // example:
  // 
  // std
  TeamPlan *string `json:"teamPlan,omitempty" xml:"teamPlan,omitempty"`
  // example:
  // 
  // tm-8f3a2c7b5e14d806
  TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
  // example:
  // 
  // 2026-08-21T09:15:30Z
  UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
  // example:
  // 
  // 9c1d4e72-****
  UserID *string `json:"userID,omitempty" xml:"userID,omitempty"`
}

func (s E2BListedTemplate) String() string {
  return dara.Prettify(s)
}

func (s E2BListedTemplate) GoString() string {
  return s.String()
}

func (s *E2BListedTemplate) GetBuildStatus() *string  {
  return s.BuildStatus
}

func (s *E2BListedTemplate) GetCategory() *string  {
  return s.Category
}

func (s *E2BListedTemplate) GetContainerConfiguration() *ContainerConfiguration  {
  return s.ContainerConfiguration
}

func (s *E2BListedTemplate) GetCpuCount() *int32  {
  return s.CpuCount
}

func (s *E2BListedTemplate) GetCreatedAt() *string  {
  return s.CreatedAt
}

func (s *E2BListedTemplate) GetFunctionName() *string  {
  return s.FunctionName
}

func (s *E2BListedTemplate) GetLogConfiguration() *LogConfiguration  {
  return s.LogConfiguration
}

func (s *E2BListedTemplate) GetMemoryMB() *int32  {
  return s.MemoryMB
}

func (s *E2BListedTemplate) GetNames() []*string  {
  return s.Names
}

func (s *E2BListedTemplate) GetPublic() *bool  {
  return s.Public
}

func (s *E2BListedTemplate) GetResourceGroupID() *string  {
  return s.ResourceGroupID
}

func (s *E2BListedTemplate) GetStatusReason() *string  {
  return s.StatusReason
}

func (s *E2BListedTemplate) GetTags() []*E2BTemplateTag  {
  return s.Tags
}

func (s *E2BListedTemplate) GetTeamID() *string  {
  return s.TeamID
}

func (s *E2BListedTemplate) GetTeamName() *string  {
  return s.TeamName
}

func (s *E2BListedTemplate) GetTeamPlan() *string  {
  return s.TeamPlan
}

func (s *E2BListedTemplate) GetTemplateID() *string  {
  return s.TemplateID
}

func (s *E2BListedTemplate) GetUpdatedAt() *string  {
  return s.UpdatedAt
}

func (s *E2BListedTemplate) GetUserID() *string  {
  return s.UserID
}

func (s *E2BListedTemplate) SetBuildStatus(v string) *E2BListedTemplate {
  s.BuildStatus = &v
  return s
}

func (s *E2BListedTemplate) SetCategory(v string) *E2BListedTemplate {
  s.Category = &v
  return s
}

func (s *E2BListedTemplate) SetContainerConfiguration(v *ContainerConfiguration) *E2BListedTemplate {
  s.ContainerConfiguration = v
  return s
}

func (s *E2BListedTemplate) SetCpuCount(v int32) *E2BListedTemplate {
  s.CpuCount = &v
  return s
}

func (s *E2BListedTemplate) SetCreatedAt(v string) *E2BListedTemplate {
  s.CreatedAt = &v
  return s
}

func (s *E2BListedTemplate) SetFunctionName(v string) *E2BListedTemplate {
  s.FunctionName = &v
  return s
}

func (s *E2BListedTemplate) SetLogConfiguration(v *LogConfiguration) *E2BListedTemplate {
  s.LogConfiguration = v
  return s
}

func (s *E2BListedTemplate) SetMemoryMB(v int32) *E2BListedTemplate {
  s.MemoryMB = &v
  return s
}

func (s *E2BListedTemplate) SetNames(v []*string) *E2BListedTemplate {
  s.Names = v
  return s
}

func (s *E2BListedTemplate) SetPublic(v bool) *E2BListedTemplate {
  s.Public = &v
  return s
}

func (s *E2BListedTemplate) SetResourceGroupID(v string) *E2BListedTemplate {
  s.ResourceGroupID = &v
  return s
}

func (s *E2BListedTemplate) SetStatusReason(v string) *E2BListedTemplate {
  s.StatusReason = &v
  return s
}

func (s *E2BListedTemplate) SetTags(v []*E2BTemplateTag) *E2BListedTemplate {
  s.Tags = v
  return s
}

func (s *E2BListedTemplate) SetTeamID(v string) *E2BListedTemplate {
  s.TeamID = &v
  return s
}

func (s *E2BListedTemplate) SetTeamName(v string) *E2BListedTemplate {
  s.TeamName = &v
  return s
}

func (s *E2BListedTemplate) SetTeamPlan(v string) *E2BListedTemplate {
  s.TeamPlan = &v
  return s
}

func (s *E2BListedTemplate) SetTemplateID(v string) *E2BListedTemplate {
  s.TemplateID = &v
  return s
}

func (s *E2BListedTemplate) SetUpdatedAt(v string) *E2BListedTemplate {
  s.UpdatedAt = &v
  return s
}

func (s *E2BListedTemplate) SetUserID(v string) *E2BListedTemplate {
  s.UserID = &v
  return s
}

func (s *E2BListedTemplate) Validate() error {
  if s.ContainerConfiguration != nil {
    if err := s.ContainerConfiguration.Validate(); err != nil {
      return err
    }
  }
  if s.LogConfiguration != nil {
    if err := s.LogConfiguration.Validate(); err != nil {
      return err
    }
  }
  if s.Tags != nil {
    for _, item := range s.Tags {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

