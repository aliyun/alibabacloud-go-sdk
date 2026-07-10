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
  SetTemplateID(v string) *E2BTemplate
  GetTemplateID() *string 
  SetUpdatedAt(v string) *E2BTemplate
  GetUpdatedAt() *string 
  SetUserID(v string) *E2BTemplate
  GetUserID() *string 
}

type E2BTemplate struct {
  BuildStatus *string `json:"buildStatus,omitempty" xml:"buildStatus,omitempty"`
  Category *string `json:"category,omitempty" xml:"category,omitempty"`
  CpuCount *int32 `json:"cpuCount,omitempty" xml:"cpuCount,omitempty"`
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
  LogConfiguration *LogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
  MemoryMB *int32 `json:"memoryMB,omitempty" xml:"memoryMB,omitempty"`
  Names []*string `json:"names,omitempty" xml:"names,omitempty" type:"Repeated"`
  NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
  Public *bool `json:"public,omitempty" xml:"public,omitempty"`
  ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
  StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
  Tags []*E2BTemplateTag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
  TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
  TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
  TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
  UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
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

