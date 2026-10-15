// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BTemplate interface {
  dara.Model
  String() string
  GoString() string
  SetBuildStatus(v string) *E2BTemplate
  GetBuildStatus() *string 
  SetCategory(v string) *E2BTemplate
  GetCategory() *string 
  SetContainerConfiguration(v *ContainerConfiguration) *E2BTemplate
  GetContainerConfiguration() *ContainerConfiguration 
  SetCpuCount(v int32) *E2BTemplate
  GetCpuCount() *int32 
  SetCreatedAt(v string) *E2BTemplate
  GetCreatedAt() *string 
  SetLogConfiguration(v *LogConfiguration) *E2BTemplate
  GetLogConfiguration() *LogConfiguration 
  SetMemoryMB(v int32) *E2BTemplate
  GetMemoryMB() *int32 
  SetNames(v []*string) *E2BTemplate
  GetNames() []*string 
  SetNetworkConfiguration(v *NetworkConfiguration) *E2BTemplate
  GetNetworkConfiguration() *NetworkConfiguration 
  SetPublic(v bool) *E2BTemplate
  GetPublic() *bool 
  SetResourceGroupID(v string) *E2BTemplate
  GetResourceGroupID() *string 
  SetStatusReason(v string) *E2BTemplate
  GetStatusReason() *string 
  SetTags(v []*E2BTemplateTag) *E2BTemplate
  GetTags() []*E2BTemplateTag 
  SetTeamID(v string) *E2BTemplate
  GetTeamID() *string 
  SetTeamName(v string) *E2BTemplate
  GetTeamName() *string 
  SetTeamPlan(v string) *E2BTemplate
  GetTeamPlan() *string 
  SetTemplateID(v string) *E2BTemplate
  GetTemplateID() *string 
  SetUpdatedAt(v string) *E2BTemplate
  GetUpdatedAt() *string 
  SetUserID(v string) *E2BTemplate
  GetUserID() *string 
}

type E2BTemplate struct {
  // The template build status. Valid values:
  // 
  // - waiting: Waiting to build.
  // 
  // - building: Building in progress.
  // 
  // - ready: Build succeeded.
  // 
  // - error: Build failed.
  // 
  // example:
  // 
  // ready
  BuildStatus *string `json:"buildStatus,omitempty" xml:"buildStatus,omitempty"`
  // The template category. Valid values:
  // 
  // - builtin: Built-in system template.
  // 
  // - custom: User-created custom template.
  // 
  // example:
  // 
  // custom
  Category *string `json:"category,omitempty" xml:"category,omitempty"`
  // The container image configuration of the template.
  ContainerConfiguration *ContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty"`
  // The CPU specification of the sandbox instance, in cores.
  // 
  // example:
  // 
  // 2
  CpuCount *int32 `json:"cpuCount,omitempty" xml:"cpuCount,omitempty"`
  // The creation time in UTC, in RFC 3339 format.
  // 
  // example:
  // 
  // 2026-08-20T08:30:00Z
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
  // The log delivery configuration of the template.
  LogConfiguration *LogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
  // The memory specification of the sandbox instance, in MB.
  // 
  // example:
  // 
  // 2048
  MemoryMB *int32 `json:"memoryMB,omitempty" xml:"memoryMB,omitempty"`
  // The list of template names, including the template name and all its aliases.
  Names []*string `json:"names" xml:"names" type:"Repeated"`
  // The network configuration of the template.
  NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
  // Indicates whether the template is public to all teams.
  Public *bool `json:"public,omitempty" xml:"public,omitempty"`
  // The Alibaba Cloud resource group ID to which the team belongs.
  // 
  // example:
  // 
  // rg-****
  ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
  // The desensitized description of the reason why the template build failed. This value is empty when the build succeeds.
  // 
  // example:
  // 
  // Failed to pull source image: invalid authentication credentials
  StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
  // The list of template tags. Each tag points to a specific build.
  Tags []*E2BTemplateTag `json:"tags" xml:"tags" type:"Repeated"`
  // The ID of the team to which the resource belongs, in UUID format.
  // 
  // example:
  // 
  // 5f4a2c18-****
  TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
  // The team name.
  // 
  // example:
  // 
  // sandbox-dev
  TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
  // The team plan specification. Valid values:
  // 
  // - eco: economy.
  // 
  // - std: standard.
  // 
  // - pro: professional.
  // 
  // - ent: enterprise.
  // 
  // example:
  // 
  // std
  TeamPlan *string `json:"teamPlan,omitempty" xml:"teamPlan,omitempty"`
  // The sandbox template ID.
  // 
  // example:
  // 
  // tm-8f3a2c7b5e14d806
  TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
  // The last update time in UTC, in RFC 3339 format.
  // 
  // example:
  // 
  // 2026-08-21T09:15:30Z
  UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
  // The user ID of the resource creator, in UUID format.
  // 
  // example:
  // 
  // 9c1d4e72-****
  UserID *string `json:"userID,omitempty" xml:"userID,omitempty"`
}

func (s E2BTemplate) String() string {
  return dara.Prettify(s)
}

func (s E2BTemplate) GoString() string {
  return s.String()
}

func (s *E2BTemplate) GetBuildStatus() *string  {
  return s.BuildStatus
}

func (s *E2BTemplate) GetCategory() *string  {
  return s.Category
}

func (s *E2BTemplate) GetContainerConfiguration() *ContainerConfiguration  {
  return s.ContainerConfiguration
}

func (s *E2BTemplate) GetCpuCount() *int32  {
  return s.CpuCount
}

func (s *E2BTemplate) GetCreatedAt() *string  {
  return s.CreatedAt
}

func (s *E2BTemplate) GetLogConfiguration() *LogConfiguration  {
  return s.LogConfiguration
}

func (s *E2BTemplate) GetMemoryMB() *int32  {
  return s.MemoryMB
}

func (s *E2BTemplate) GetNames() []*string  {
  return s.Names
}

func (s *E2BTemplate) GetNetworkConfiguration() *NetworkConfiguration  {
  return s.NetworkConfiguration
}

func (s *E2BTemplate) GetPublic() *bool  {
  return s.Public
}

func (s *E2BTemplate) GetResourceGroupID() *string  {
  return s.ResourceGroupID
}

func (s *E2BTemplate) GetStatusReason() *string  {
  return s.StatusReason
}

func (s *E2BTemplate) GetTags() []*E2BTemplateTag  {
  return s.Tags
}

func (s *E2BTemplate) GetTeamID() *string  {
  return s.TeamID
}

func (s *E2BTemplate) GetTeamName() *string  {
  return s.TeamName
}

func (s *E2BTemplate) GetTeamPlan() *string  {
  return s.TeamPlan
}

func (s *E2BTemplate) GetTemplateID() *string  {
  return s.TemplateID
}

func (s *E2BTemplate) GetUpdatedAt() *string  {
  return s.UpdatedAt
}

func (s *E2BTemplate) GetUserID() *string  {
  return s.UserID
}

func (s *E2BTemplate) SetBuildStatus(v string) *E2BTemplate {
  s.BuildStatus = &v
  return s
}

func (s *E2BTemplate) SetCategory(v string) *E2BTemplate {
  s.Category = &v
  return s
}

func (s *E2BTemplate) SetContainerConfiguration(v *ContainerConfiguration) *E2BTemplate {
  s.ContainerConfiguration = v
  return s
}

func (s *E2BTemplate) SetCpuCount(v int32) *E2BTemplate {
  s.CpuCount = &v
  return s
}

func (s *E2BTemplate) SetCreatedAt(v string) *E2BTemplate {
  s.CreatedAt = &v
  return s
}

func (s *E2BTemplate) SetLogConfiguration(v *LogConfiguration) *E2BTemplate {
  s.LogConfiguration = v
  return s
}

func (s *E2BTemplate) SetMemoryMB(v int32) *E2BTemplate {
  s.MemoryMB = &v
  return s
}

func (s *E2BTemplate) SetNames(v []*string) *E2BTemplate {
  s.Names = v
  return s
}

func (s *E2BTemplate) SetNetworkConfiguration(v *NetworkConfiguration) *E2BTemplate {
  s.NetworkConfiguration = v
  return s
}

func (s *E2BTemplate) SetPublic(v bool) *E2BTemplate {
  s.Public = &v
  return s
}

func (s *E2BTemplate) SetResourceGroupID(v string) *E2BTemplate {
  s.ResourceGroupID = &v
  return s
}

func (s *E2BTemplate) SetStatusReason(v string) *E2BTemplate {
  s.StatusReason = &v
  return s
}

func (s *E2BTemplate) SetTags(v []*E2BTemplateTag) *E2BTemplate {
  s.Tags = v
  return s
}

func (s *E2BTemplate) SetTeamID(v string) *E2BTemplate {
  s.TeamID = &v
  return s
}

func (s *E2BTemplate) SetTeamName(v string) *E2BTemplate {
  s.TeamName = &v
  return s
}

func (s *E2BTemplate) SetTeamPlan(v string) *E2BTemplate {
  s.TeamPlan = &v
  return s
}

func (s *E2BTemplate) SetTemplateID(v string) *E2BTemplate {
  s.TemplateID = &v
  return s
}

func (s *E2BTemplate) SetUpdatedAt(v string) *E2BTemplate {
  s.UpdatedAt = &v
  return s
}

func (s *E2BTemplate) SetUserID(v string) *E2BTemplate {
  s.UserID = &v
  return s
}

func (s *E2BTemplate) Validate() error {
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
  if s.NetworkConfiguration != nil {
    if err := s.NetworkConfiguration.Validate(); err != nil {
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

