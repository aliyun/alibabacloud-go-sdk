// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BTemplateBuild interface {
  dara.Model
  String() string
  GoString() string
  SetBuildID(v string) *E2BTemplateBuild
  GetBuildID() *string 
  SetCpuCount(v int32) *E2BTemplateBuild
  GetCpuCount() *int32 
  SetCreatedAt(v string) *E2BTemplateBuild
  GetCreatedAt() *string 
  SetDiskSizeMB(v int32) *E2BTemplateBuild
  GetDiskSizeMB() *int32 
  SetEnvdVersion(v string) *E2BTemplateBuild
  GetEnvdVersion() *string 
  SetFinishedAt(v string) *E2BTemplateBuild
  GetFinishedAt() *string 
  SetMemoryMB(v int32) *E2BTemplateBuild
  GetMemoryMB() *int32 
  SetStatus(v string) *E2BTemplateBuild
  GetStatus() *string 
  SetUpdatedAt(v string) *E2BTemplateBuild
  GetUpdatedAt() *string 
}

type E2BTemplateBuild struct {
  BuildID *string `json:"buildID,omitempty" xml:"buildID,omitempty"`
  CpuCount *int32 `json:"cpuCount,omitempty" xml:"cpuCount,omitempty"`
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
  DiskSizeMB *int32 `json:"diskSizeMB,omitempty" xml:"diskSizeMB,omitempty"`
  EnvdVersion *string `json:"envdVersion,omitempty" xml:"envdVersion,omitempty"`
  FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
  MemoryMB *int32 `json:"memoryMB,omitempty" xml:"memoryMB,omitempty"`
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s E2BTemplateBuild) String() string {
  return dara.Prettify(s)
}

func (s E2BTemplateBuild) GoString() string {
  return s.String()
}

func (s *E2BTemplateBuild) GetBuildID() *string  {
  return s.BuildID
}

func (s *E2BTemplateBuild) GetCpuCount() *int32  {
  return s.CpuCount
}

func (s *E2BTemplateBuild) GetCreatedAt() *string  {
  return s.CreatedAt
}

func (s *E2BTemplateBuild) GetDiskSizeMB() *int32  {
  return s.DiskSizeMB
}

func (s *E2BTemplateBuild) GetEnvdVersion() *string  {
  return s.EnvdVersion
}

func (s *E2BTemplateBuild) GetFinishedAt() *string  {
  return s.FinishedAt
}

func (s *E2BTemplateBuild) GetMemoryMB() *int32  {
  return s.MemoryMB
}

func (s *E2BTemplateBuild) GetStatus() *string  {
  return s.Status
}

func (s *E2BTemplateBuild) GetUpdatedAt() *string  {
  return s.UpdatedAt
}

func (s *E2BTemplateBuild) SetBuildID(v string) *E2BTemplateBuild {
  s.BuildID = &v
  return s
}

func (s *E2BTemplateBuild) SetCpuCount(v int32) *E2BTemplateBuild {
  s.CpuCount = &v
  return s
}

func (s *E2BTemplateBuild) SetCreatedAt(v string) *E2BTemplateBuild {
  s.CreatedAt = &v
  return s
}

func (s *E2BTemplateBuild) SetDiskSizeMB(v int32) *E2BTemplateBuild {
  s.DiskSizeMB = &v
  return s
}

func (s *E2BTemplateBuild) SetEnvdVersion(v string) *E2BTemplateBuild {
  s.EnvdVersion = &v
  return s
}

func (s *E2BTemplateBuild) SetFinishedAt(v string) *E2BTemplateBuild {
  s.FinishedAt = &v
  return s
}

func (s *E2BTemplateBuild) SetMemoryMB(v int32) *E2BTemplateBuild {
  s.MemoryMB = &v
  return s
}

func (s *E2BTemplateBuild) SetStatus(v string) *E2BTemplateBuild {
  s.Status = &v
  return s
}

func (s *E2BTemplateBuild) SetUpdatedAt(v string) *E2BTemplateBuild {
  s.UpdatedAt = &v
  return s
}

func (s *E2BTemplateBuild) Validate() error {
  return dara.Validate(s)
}

