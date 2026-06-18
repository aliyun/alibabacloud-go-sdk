// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BTemplate interface {
  dara.Model
  String() string
  GoString() string
  SetAliases(v []*string) *E2BTemplate
  GetAliases() []*string 
  SetBuildStatus(v string) *E2BTemplate
  GetBuildStatus() *string 
  SetBuilds(v *E2BTemplateBuild) *E2BTemplate
  GetBuilds() *E2BTemplateBuild 
  SetCpuCount(v int32) *E2BTemplate
  GetCpuCount() *int32 
  SetCreatedAt(v string) *E2BTemplate
  GetCreatedAt() *string 
  SetLastSpawnedAt(v string) *E2BTemplate
  GetLastSpawnedAt() *string 
  SetLogConfiguration(v *LogConfiguration) *E2BTemplate
  GetLogConfiguration() *LogConfiguration 
  SetMemoryMB(v int32) *E2BTemplate
  GetMemoryMB() *int32 
  SetNames(v []*string) *E2BTemplate
  GetNames() []*string 
  SetPublic(v bool) *E2BTemplate
  GetPublic() *bool 
  SetSpawnCount(v string) *E2BTemplate
  GetSpawnCount() *string 
  SetStatusReason(v string) *E2BTemplate
  GetStatusReason() *string 
  SetTags(v []*E2BTemplateTag) *E2BTemplate
  GetTags() []*E2BTemplateTag 
  SetTemplateID(v string) *E2BTemplate
  GetTemplateID() *string 
  SetUpdatedAt(v string) *E2BTemplate
  GetUpdatedAt() *string 
  SetUserID(v string) *E2BTemplate
  GetUserID() *string 
}

type E2BTemplate struct {
  Aliases []*string `json:"aliases,omitempty" xml:"aliases,omitempty" type:"Repeated"`
  BuildStatus *string `json:"buildStatus,omitempty" xml:"buildStatus,omitempty"`
  Builds *E2BTemplateBuild `json:"builds,omitempty" xml:"builds,omitempty"`
  CpuCount *int32 `json:"cpuCount,omitempty" xml:"cpuCount,omitempty"`
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
  LastSpawnedAt *string `json:"lastSpawnedAt,omitempty" xml:"lastSpawnedAt,omitempty"`
  LogConfiguration *LogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
  MemoryMB *int32 `json:"memoryMB,omitempty" xml:"memoryMB,omitempty"`
  Names []*string `json:"names,omitempty" xml:"names,omitempty" type:"Repeated"`
  Public *bool `json:"public,omitempty" xml:"public,omitempty"`
  SpawnCount *string `json:"spawnCount,omitempty" xml:"spawnCount,omitempty"`
  StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
  Tags []*E2BTemplateTag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
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

func (s *E2BTemplate) GetAliases() []*string  {
  return s.Aliases
}

func (s *E2BTemplate) GetBuildStatus() *string  {
  return s.BuildStatus
}

func (s *E2BTemplate) GetBuilds() *E2BTemplateBuild  {
  return s.Builds
}

func (s *E2BTemplate) GetCpuCount() *int32  {
  return s.CpuCount
}

func (s *E2BTemplate) GetCreatedAt() *string  {
  return s.CreatedAt
}

func (s *E2BTemplate) GetLastSpawnedAt() *string  {
  return s.LastSpawnedAt
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

func (s *E2BTemplate) GetPublic() *bool  {
  return s.Public
}

func (s *E2BTemplate) GetSpawnCount() *string  {
  return s.SpawnCount
}

func (s *E2BTemplate) GetStatusReason() *string  {
  return s.StatusReason
}

func (s *E2BTemplate) GetTags() []*E2BTemplateTag  {
  return s.Tags
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

func (s *E2BTemplate) SetAliases(v []*string) *E2BTemplate {
  s.Aliases = v
  return s
}

func (s *E2BTemplate) SetBuildStatus(v string) *E2BTemplate {
  s.BuildStatus = &v
  return s
}

func (s *E2BTemplate) SetBuilds(v *E2BTemplateBuild) *E2BTemplate {
  s.Builds = v
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

func (s *E2BTemplate) SetLastSpawnedAt(v string) *E2BTemplate {
  s.LastSpawnedAt = &v
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

func (s *E2BTemplate) SetPublic(v bool) *E2BTemplate {
  s.Public = &v
  return s
}

func (s *E2BTemplate) SetSpawnCount(v string) *E2BTemplate {
  s.SpawnCount = &v
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
  if s.Builds != nil {
    if err := s.Builds.Validate(); err != nil {
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

