// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BSandbox interface {
  dara.Model
  String() string
  GoString() string
  SetAlias(v string) *E2BSandbox
  GetAlias() *string 
  SetAllowInternetAccess(v bool) *E2BSandbox
  GetAllowInternetAccess() *bool 
  SetClientID(v string) *E2BSandbox
  GetClientID() *string 
  SetCpuCount(v int32) *E2BSandbox
  GetCpuCount() *int32 
  SetDiskSizeMB(v int32) *E2BSandbox
  GetDiskSizeMB() *int32 
  SetDomain(v string) *E2BSandbox
  GetDomain() *string 
  SetEndAt(v string) *E2BSandbox
  GetEndAt() *string 
  SetEnvdAccessToken(v string) *E2BSandbox
  GetEnvdAccessToken() *string 
  SetEnvdVersion(v string) *E2BSandbox
  GetEnvdVersion() *string 
  SetFcFunctionName(v string) *E2BSandbox
  GetFcFunctionName() *string 
  SetFcInstanceID(v string) *E2BSandbox
  GetFcInstanceID() *string 
  SetFcSessionID(v string) *E2BSandbox
  GetFcSessionID() *string 
  SetLifecycle(v *E2BLifecycle) *E2BSandbox
  GetLifecycle() *E2BLifecycle 
  SetMemoryMB(v int32) *E2BSandbox
  GetMemoryMB() *int32 
  SetMetadata(v map[string]*string) *E2BSandbox
  GetMetadata() map[string]*string 
  SetNetwork(v *E2BNetwork) *E2BSandbox
  GetNetwork() *E2BNetwork 
  SetSandboxID(v string) *E2BSandbox
  GetSandboxID() *string 
  SetStartedAt(v string) *E2BSandbox
  GetStartedAt() *string 
  SetState(v string) *E2BSandbox
  GetState() *string 
  SetTemplateId(v string) *E2BSandbox
  GetTemplateId() *string 
  SetTemplateName(v string) *E2BSandbox
  GetTemplateName() *string 
  SetVolumeMounts(v []*E2BVolumeMount) *E2BSandbox
  GetVolumeMounts() []*E2BVolumeMount 
}

type E2BSandbox struct {
  Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
  AllowInternetAccess *bool `json:"allowInternetAccess,omitempty" xml:"allowInternetAccess,omitempty"`
  ClientID *string `json:"clientID,omitempty" xml:"clientID,omitempty"`
  CpuCount *int32 `json:"cpuCount,omitempty" xml:"cpuCount,omitempty"`
  DiskSizeMB *int32 `json:"diskSizeMB,omitempty" xml:"diskSizeMB,omitempty"`
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  EndAt *string `json:"endAt,omitempty" xml:"endAt,omitempty"`
  EnvdAccessToken *string `json:"envdAccessToken,omitempty" xml:"envdAccessToken,omitempty"`
  EnvdVersion *string `json:"envdVersion,omitempty" xml:"envdVersion,omitempty"`
  FcFunctionName *string `json:"fcFunctionName,omitempty" xml:"fcFunctionName,omitempty"`
  FcInstanceID *string `json:"fcInstanceID,omitempty" xml:"fcInstanceID,omitempty"`
  FcSessionID *string `json:"fcSessionID,omitempty" xml:"fcSessionID,omitempty"`
  Lifecycle *E2BLifecycle `json:"lifecycle,omitempty" xml:"lifecycle,omitempty"`
  MemoryMB *int32 `json:"memoryMB,omitempty" xml:"memoryMB,omitempty"`
  Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
  Network *E2BNetwork `json:"network,omitempty" xml:"network,omitempty"`
  SandboxID *string `json:"sandboxID,omitempty" xml:"sandboxID,omitempty"`
  StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
  State *string `json:"state,omitempty" xml:"state,omitempty"`
  TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
  TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
  VolumeMounts []*E2BVolumeMount `json:"volumeMounts,omitempty" xml:"volumeMounts,omitempty" type:"Repeated"`
}

func (s E2BSandbox) String() string {
  return dara.Prettify(s)
}

func (s E2BSandbox) GoString() string {
  return s.String()
}

func (s *E2BSandbox) GetAlias() *string  {
  return s.Alias
}

func (s *E2BSandbox) GetAllowInternetAccess() *bool  {
  return s.AllowInternetAccess
}

func (s *E2BSandbox) GetClientID() *string  {
  return s.ClientID
}

func (s *E2BSandbox) GetCpuCount() *int32  {
  return s.CpuCount
}

func (s *E2BSandbox) GetDiskSizeMB() *int32  {
  return s.DiskSizeMB
}

func (s *E2BSandbox) GetDomain() *string  {
  return s.Domain
}

func (s *E2BSandbox) GetEndAt() *string  {
  return s.EndAt
}

func (s *E2BSandbox) GetEnvdAccessToken() *string  {
  return s.EnvdAccessToken
}

func (s *E2BSandbox) GetEnvdVersion() *string  {
  return s.EnvdVersion
}

func (s *E2BSandbox) GetFcFunctionName() *string  {
  return s.FcFunctionName
}

func (s *E2BSandbox) GetFcInstanceID() *string  {
  return s.FcInstanceID
}

func (s *E2BSandbox) GetFcSessionID() *string  {
  return s.FcSessionID
}

func (s *E2BSandbox) GetLifecycle() *E2BLifecycle  {
  return s.Lifecycle
}

func (s *E2BSandbox) GetMemoryMB() *int32  {
  return s.MemoryMB
}

func (s *E2BSandbox) GetMetadata() map[string]*string  {
  return s.Metadata
}

func (s *E2BSandbox) GetNetwork() *E2BNetwork  {
  return s.Network
}

func (s *E2BSandbox) GetSandboxID() *string  {
  return s.SandboxID
}

func (s *E2BSandbox) GetStartedAt() *string  {
  return s.StartedAt
}

func (s *E2BSandbox) GetState() *string  {
  return s.State
}

func (s *E2BSandbox) GetTemplateId() *string  {
  return s.TemplateId
}

func (s *E2BSandbox) GetTemplateName() *string  {
  return s.TemplateName
}

func (s *E2BSandbox) GetVolumeMounts() []*E2BVolumeMount  {
  return s.VolumeMounts
}

func (s *E2BSandbox) SetAlias(v string) *E2BSandbox {
  s.Alias = &v
  return s
}

func (s *E2BSandbox) SetAllowInternetAccess(v bool) *E2BSandbox {
  s.AllowInternetAccess = &v
  return s
}

func (s *E2BSandbox) SetClientID(v string) *E2BSandbox {
  s.ClientID = &v
  return s
}

func (s *E2BSandbox) SetCpuCount(v int32) *E2BSandbox {
  s.CpuCount = &v
  return s
}

func (s *E2BSandbox) SetDiskSizeMB(v int32) *E2BSandbox {
  s.DiskSizeMB = &v
  return s
}

func (s *E2BSandbox) SetDomain(v string) *E2BSandbox {
  s.Domain = &v
  return s
}

func (s *E2BSandbox) SetEndAt(v string) *E2BSandbox {
  s.EndAt = &v
  return s
}

func (s *E2BSandbox) SetEnvdAccessToken(v string) *E2BSandbox {
  s.EnvdAccessToken = &v
  return s
}

func (s *E2BSandbox) SetEnvdVersion(v string) *E2BSandbox {
  s.EnvdVersion = &v
  return s
}

func (s *E2BSandbox) SetFcFunctionName(v string) *E2BSandbox {
  s.FcFunctionName = &v
  return s
}

func (s *E2BSandbox) SetFcInstanceID(v string) *E2BSandbox {
  s.FcInstanceID = &v
  return s
}

func (s *E2BSandbox) SetFcSessionID(v string) *E2BSandbox {
  s.FcSessionID = &v
  return s
}

func (s *E2BSandbox) SetLifecycle(v *E2BLifecycle) *E2BSandbox {
  s.Lifecycle = v
  return s
}

func (s *E2BSandbox) SetMemoryMB(v int32) *E2BSandbox {
  s.MemoryMB = &v
  return s
}

func (s *E2BSandbox) SetMetadata(v map[string]*string) *E2BSandbox {
  s.Metadata = v
  return s
}

func (s *E2BSandbox) SetNetwork(v *E2BNetwork) *E2BSandbox {
  s.Network = v
  return s
}

func (s *E2BSandbox) SetSandboxID(v string) *E2BSandbox {
  s.SandboxID = &v
  return s
}

func (s *E2BSandbox) SetStartedAt(v string) *E2BSandbox {
  s.StartedAt = &v
  return s
}

func (s *E2BSandbox) SetState(v string) *E2BSandbox {
  s.State = &v
  return s
}

func (s *E2BSandbox) SetTemplateId(v string) *E2BSandbox {
  s.TemplateId = &v
  return s
}

func (s *E2BSandbox) SetTemplateName(v string) *E2BSandbox {
  s.TemplateName = &v
  return s
}

func (s *E2BSandbox) SetVolumeMounts(v []*E2BVolumeMount) *E2BSandbox {
  s.VolumeMounts = v
  return s
}

func (s *E2BSandbox) Validate() error {
  if s.Lifecycle != nil {
    if err := s.Lifecycle.Validate(); err != nil {
      return err
    }
  }
  if s.Network != nil {
    if err := s.Network.Validate(); err != nil {
      return err
    }
  }
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

