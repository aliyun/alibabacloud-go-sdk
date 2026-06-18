// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BListedSandbox interface {
  dara.Model
  String() string
  GoString() string
  SetAlias(v string) *E2BListedSandbox
  GetAlias() *string 
  SetClientID(v string) *E2BListedSandbox
  GetClientID() *string 
  SetCpuCount(v int32) *E2BListedSandbox
  GetCpuCount() *int32 
  SetDiskSizeMB(v int32) *E2BListedSandbox
  GetDiskSizeMB() *int32 
  SetEndAt(v string) *E2BListedSandbox
  GetEndAt() *string 
  SetEnvdVersion(v string) *E2BListedSandbox
  GetEnvdVersion() *string 
  SetMemoryMB(v int32) *E2BListedSandbox
  GetMemoryMB() *int32 
  SetMetadata(v map[string]*string) *E2BListedSandbox
  GetMetadata() map[string]*string 
  SetSandboxID(v string) *E2BListedSandbox
  GetSandboxID() *string 
  SetStartedAt(v string) *E2BListedSandbox
  GetStartedAt() *string 
  SetState(v string) *E2BListedSandbox
  GetState() *string 
  SetTemplateID(v string) *E2BListedSandbox
  GetTemplateID() *string 
  SetTemplateName(v string) *E2BListedSandbox
  GetTemplateName() *string 
  SetVolumeMounts(v []*E2BVolumeMount) *E2BListedSandbox
  GetVolumeMounts() []*E2BVolumeMount 
}

type E2BListedSandbox struct {
  Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
  ClientID *string `json:"clientID,omitempty" xml:"clientID,omitempty"`
  CpuCount *int32 `json:"cpuCount,omitempty" xml:"cpuCount,omitempty"`
  DiskSizeMB *int32 `json:"diskSizeMB,omitempty" xml:"diskSizeMB,omitempty"`
  EndAt *string `json:"endAt,omitempty" xml:"endAt,omitempty"`
  EnvdVersion *string `json:"envdVersion,omitempty" xml:"envdVersion,omitempty"`
  MemoryMB *int32 `json:"memoryMB,omitempty" xml:"memoryMB,omitempty"`
  Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
  SandboxID *string `json:"sandboxID,omitempty" xml:"sandboxID,omitempty"`
  StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
  State *string `json:"state,omitempty" xml:"state,omitempty"`
  TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
  TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
  VolumeMounts []*E2BVolumeMount `json:"volumeMounts,omitempty" xml:"volumeMounts,omitempty" type:"Repeated"`
}

func (s E2BListedSandbox) String() string {
  return dara.Prettify(s)
}

func (s E2BListedSandbox) GoString() string {
  return s.String()
}

func (s *E2BListedSandbox) GetAlias() *string  {
  return s.Alias
}

func (s *E2BListedSandbox) GetClientID() *string  {
  return s.ClientID
}

func (s *E2BListedSandbox) GetCpuCount() *int32  {
  return s.CpuCount
}

func (s *E2BListedSandbox) GetDiskSizeMB() *int32  {
  return s.DiskSizeMB
}

func (s *E2BListedSandbox) GetEndAt() *string  {
  return s.EndAt
}

func (s *E2BListedSandbox) GetEnvdVersion() *string  {
  return s.EnvdVersion
}

func (s *E2BListedSandbox) GetMemoryMB() *int32  {
  return s.MemoryMB
}

func (s *E2BListedSandbox) GetMetadata() map[string]*string  {
  return s.Metadata
}

func (s *E2BListedSandbox) GetSandboxID() *string  {
  return s.SandboxID
}

func (s *E2BListedSandbox) GetStartedAt() *string  {
  return s.StartedAt
}

func (s *E2BListedSandbox) GetState() *string  {
  return s.State
}

func (s *E2BListedSandbox) GetTemplateID() *string  {
  return s.TemplateID
}

func (s *E2BListedSandbox) GetTemplateName() *string  {
  return s.TemplateName
}

func (s *E2BListedSandbox) GetVolumeMounts() []*E2BVolumeMount  {
  return s.VolumeMounts
}

func (s *E2BListedSandbox) SetAlias(v string) *E2BListedSandbox {
  s.Alias = &v
  return s
}

func (s *E2BListedSandbox) SetClientID(v string) *E2BListedSandbox {
  s.ClientID = &v
  return s
}

func (s *E2BListedSandbox) SetCpuCount(v int32) *E2BListedSandbox {
  s.CpuCount = &v
  return s
}

func (s *E2BListedSandbox) SetDiskSizeMB(v int32) *E2BListedSandbox {
  s.DiskSizeMB = &v
  return s
}

func (s *E2BListedSandbox) SetEndAt(v string) *E2BListedSandbox {
  s.EndAt = &v
  return s
}

func (s *E2BListedSandbox) SetEnvdVersion(v string) *E2BListedSandbox {
  s.EnvdVersion = &v
  return s
}

func (s *E2BListedSandbox) SetMemoryMB(v int32) *E2BListedSandbox {
  s.MemoryMB = &v
  return s
}

func (s *E2BListedSandbox) SetMetadata(v map[string]*string) *E2BListedSandbox {
  s.Metadata = v
  return s
}

func (s *E2BListedSandbox) SetSandboxID(v string) *E2BListedSandbox {
  s.SandboxID = &v
  return s
}

func (s *E2BListedSandbox) SetStartedAt(v string) *E2BListedSandbox {
  s.StartedAt = &v
  return s
}

func (s *E2BListedSandbox) SetState(v string) *E2BListedSandbox {
  s.State = &v
  return s
}

func (s *E2BListedSandbox) SetTemplateID(v string) *E2BListedSandbox {
  s.TemplateID = &v
  return s
}

func (s *E2BListedSandbox) SetTemplateName(v string) *E2BListedSandbox {
  s.TemplateName = &v
  return s
}

func (s *E2BListedSandbox) SetVolumeMounts(v []*E2BVolumeMount) *E2BListedSandbox {
  s.VolumeMounts = v
  return s
}

func (s *E2BListedSandbox) Validate() error {
  if s.VolumeMounts != nil {
    for _, item := range s.VolumeMounts {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

