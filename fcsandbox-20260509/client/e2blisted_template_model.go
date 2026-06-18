// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BListedTemplate interface {
  dara.Model
  String() string
  GoString() string
  SetAliases(v []*string) *E2BListedTemplate
  GetAliases() []*string 
  SetBuildStatus(v string) *E2BListedTemplate
  GetBuildStatus() *string 
  SetCpuCount(v int32) *E2BListedTemplate
  GetCpuCount() *int32 
  SetCreatedAt(v string) *E2BListedTemplate
  GetCreatedAt() *string 
  SetLastSpawnedAt(v string) *E2BListedTemplate
  GetLastSpawnedAt() *string 
  SetLogConfiguration(v *LogConfiguration) *E2BListedTemplate
  GetLogConfiguration() *LogConfiguration 
  SetMemoryMB(v int32) *E2BListedTemplate
  GetMemoryMB() *int32 
  SetNames(v []*string) *E2BListedTemplate
  GetNames() []*string 
  SetPublic(v bool) *E2BListedTemplate
  GetPublic() *bool 
  SetSpawnCount(v string) *E2BListedTemplate
  GetSpawnCount() *string 
  SetStatusReason(v string) *E2BListedTemplate
  GetStatusReason() *string 
  SetTags(v []*E2BTemplateTag) *E2BListedTemplate
  GetTags() []*E2BTemplateTag 
  SetTemplateID(v string) *E2BListedTemplate
  GetTemplateID() *string 
  SetUpdatedAt(v string) *E2BListedTemplate
  GetUpdatedAt() *string 
}

type E2BListedTemplate struct {
  Aliases []*string `json:"aliases,omitempty" xml:"aliases,omitempty" type:"Repeated"`
  BuildStatus *string `json:"buildStatus,omitempty" xml:"buildStatus,omitempty"`
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
}

func (s E2BListedTemplate) String() string {
  return dara.Prettify(s)
}

func (s E2BListedTemplate) GoString() string {
  return s.String()
}

func (s *E2BListedTemplate) GetAliases() []*string  {
  return s.Aliases
}

func (s *E2BListedTemplate) GetBuildStatus() *string  {
  return s.BuildStatus
}

func (s *E2BListedTemplate) GetCpuCount() *int32  {
  return s.CpuCount
}

func (s *E2BListedTemplate) GetCreatedAt() *string  {
  return s.CreatedAt
}

func (s *E2BListedTemplate) GetLastSpawnedAt() *string  {
  return s.LastSpawnedAt
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

func (s *E2BListedTemplate) GetSpawnCount() *string  {
  return s.SpawnCount
}

func (s *E2BListedTemplate) GetStatusReason() *string  {
  return s.StatusReason
}

func (s *E2BListedTemplate) GetTags() []*E2BTemplateTag  {
  return s.Tags
}

func (s *E2BListedTemplate) GetTemplateID() *string  {
  return s.TemplateID
}

func (s *E2BListedTemplate) GetUpdatedAt() *string  {
  return s.UpdatedAt
}

func (s *E2BListedTemplate) SetAliases(v []*string) *E2BListedTemplate {
  s.Aliases = v
  return s
}

func (s *E2BListedTemplate) SetBuildStatus(v string) *E2BListedTemplate {
  s.BuildStatus = &v
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

func (s *E2BListedTemplate) SetLastSpawnedAt(v string) *E2BListedTemplate {
  s.LastSpawnedAt = &v
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

func (s *E2BListedTemplate) SetSpawnCount(v string) *E2BListedTemplate {
  s.SpawnCount = &v
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

func (s *E2BListedTemplate) SetTemplateID(v string) *E2BListedTemplate {
  s.TemplateID = &v
  return s
}

func (s *E2BListedTemplate) SetUpdatedAt(v string) *E2BListedTemplate {
  s.UpdatedAt = &v
  return s
}

func (s *E2BListedTemplate) Validate() error {
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

