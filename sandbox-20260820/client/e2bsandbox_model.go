// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BSandbox interface {
  dara.Model
  String() string
  GoString() string
  SetAccessEndpoint(v string) *E2BSandbox
  GetAccessEndpoint() *string 
  SetAllowInternetAccess(v bool) *E2BSandbox
  GetAllowInternetAccess() *bool 
  SetCpuCount(v int32) *E2BSandbox
  GetCpuCount() *int32 
  SetDiskSizeMB(v int32) *E2BSandbox
  GetDiskSizeMB() *int32 
  SetDomain(v string) *E2BSandbox
  GetDomain() *string 
  SetEndAt(v string) *E2BSandbox
  GetEndAt() *string 
  SetEnvVars(v map[string]*string) *E2BSandbox
  GetEnvVars() map[string]*string 
  SetEnvdAccessToken(v string) *E2BSandbox
  GetEnvdAccessToken() *string 
  SetFcFunctionName(v string) *E2BSandbox
  GetFcFunctionName() *string 
  SetFcInstanceID(v string) *E2BSandbox
  GetFcInstanceID() *string 
  SetFcSessionID(v string) *E2BSandbox
  GetFcSessionID() *string 
  SetGeneration(v int32) *E2BSandbox
  GetGeneration() *int32 
  SetMemoryMB(v int32) *E2BSandbox
  GetMemoryMB() *int32 
  SetMetadata(v map[string]*string) *E2BSandbox
  GetMetadata() map[string]*string 
  SetNetwork(v *E2BNetwork) *E2BSandbox
  GetNetwork() *E2BNetwork 
  SetResourceGroupID(v string) *E2BSandbox
  GetResourceGroupID() *string 
  SetRuntime(v *InnerSandboxRuntimeConfig) *E2BSandbox
  GetRuntime() *InnerSandboxRuntimeConfig 
  SetSandboxID(v string) *E2BSandbox
  GetSandboxID() *string 
  SetStartedAt(v string) *E2BSandbox
  GetStartedAt() *string 
  SetState(v string) *E2BSandbox
  GetState() *string 
  SetTeamID(v string) *E2BSandbox
  GetTeamID() *string 
  SetTeamName(v string) *E2BSandbox
  GetTeamName() *string 
  SetTeamPlan(v string) *E2BSandbox
  GetTeamPlan() *string 
  SetTemplateID(v string) *E2BSandbox
  GetTemplateID() *string 
  SetTemplateName(v string) *E2BSandbox
  GetTemplateName() *string 
  SetUserID(v string) *E2BSandbox
  GetUserID() *string 
  SetVolumeMounts(v *InnerSandboxVolumeMount) *E2BSandbox
  GetVolumeMounts() *InnerSandboxVolumeMount 
}

type E2BSandbox struct {
  // The access endpoint.
  // 
  // example:
  // 
  // sbx-xxxx.ap-southeast-1.sandbox.aliyun.com
  AccessEndpoint *string `json:"accessEndpoint,omitempty" xml:"accessEndpoint,omitempty"`
  // Specifies whether the sandbox is allowed to access the external network.
  AllowInternetAccess *bool `json:"allowInternetAccess,omitempty" xml:"allowInternetAccess,omitempty"`
  // The vCPU specifications.
  // 
  // example:
  // 
  // 2
  CpuCount *int32 `json:"cpuCount,omitempty" xml:"cpuCount,omitempty"`
  // The ephemeral disk specifications.
  // 
  // example:
  // 
  // 10240
  DiskSizeMB *int32 `json:"diskSizeMB,omitempty" xml:"diskSizeMB,omitempty"`
  // The access domain name of the sandbox.
  // 
  // example:
  // 
  // ap-southeast-1.sandbox.aliyun.com
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // The end time.
  // 
  // example:
  // 
  // 2026-09-04T02:03:43.826Z
  EndAt *string `json:"endAt,omitempty" xml:"endAt,omitempty"`
  // The environment variables.
  EnvVars map[string]*string `json:"envVars" xml:"envVars"`
  // The authentication token for accessing the envd daemon process in the sandbox, in UUID format. This token is automatically generated only when secure is set to true.
  // 
  // example:
  // 
  // 4b1e9d07-****
  EnvdAccessToken *string `json:"envdAccessToken,omitempty" xml:"envdAccessToken,omitempty"`
  // The Function Compute function name.
  // 
  // example:
  // 
  // template-d391c80b5d1527de
  FcFunctionName *string `json:"fcFunctionName,omitempty" xml:"fcFunctionName,omitempty"`
  // The Function Compute instance ID.
  // 
  // example:
  // 
  // c-6a9d1501-133e6f48-b3f07782dd47
  FcInstanceID *string `json:"fcInstanceID,omitempty" xml:"fcInstanceID,omitempty"`
  // The Function Compute session ID corresponding to the sandbox, in UUID format. This ID is used to locate the underlying session when pausing, resuming, or deleting the sandbox.
  // 
  // example:
  // 
  // 3a7ecd09-****
  FcSessionID *string `json:"fcSessionID,omitempty" xml:"fcSessionID,omitempty"`
  // The sandbox version.
  // 
  // example:
  // 
  // 2
  Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
  // The memory specifications.
  // 
  // example:
  // 
  // 2
  MemoryMB *int32 `json:"memoryMB,omitempty" xml:"memoryMB,omitempty"`
  // The metadata configuration.
  Metadata map[string]*string `json:"metadata" xml:"metadata"`
  // The network configuration.
  Network *E2BNetwork `json:"network,omitempty" xml:"network,omitempty"`
  // The resource group ID.
  // 
  // example:
  // 
  // rg-****
  ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
  // The runtime configuration.
  Runtime *InnerSandboxRuntimeConfig `json:"runtime,omitempty" xml:"runtime,omitempty"`
  // The unique identifier of the sandbox.
  // 
  // example:
  // 
  // sbx-xxxx
  SandboxID *string `json:"sandboxID,omitempty" xml:"sandboxID,omitempty"`
  // The start time.
  // 
  // example:
  // 
  // 2026-09-04T02:02:43.826Z
  StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
  // The sandbox state. Valid values:
  // 
  // - running: The sandbox is running.
  // 
  // - pausing: The sandbox is pausing.
  // 
  // - paused: The sandbox is paused.
  // 
  // - resuming: The sandbox is resuming.
  // 
  // - pause_fail: The pause operation failed.
  // 
  // - resume_fail: The resume operation failed.
  // 
  // - snapshotting: A snapshot is being created.
  // 
  // - snapshot_failed: Snapshot creation failed.
  // 
  // - terminated: The sandbox is terminated.
  // 
  // example:
  // 
  // running
  State *string `json:"state,omitempty" xml:"state,omitempty"`
  // The unique identifier of the team.
  // 
  // example:
  // 
  // eea033af-****
  TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
  // The team name.
  // 
  // example:
  // 
  // default
  TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
  // The QoS version.
  // 
  // example:
  // 
  // eco
  TeamPlan *string `json:"teamPlan,omitempty" xml:"teamPlan,omitempty"`
  // The unique identifier of the template.
  // 
  // example:
  // 
  // v9vjyzw64bsu83vw0dgs
  TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
  // The template name.
  // 
  // example:
  // 
  // base
  TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
  // The UID of the creator.
  // 
  // example:
  // 
  // 2000****
  UserID *string `json:"userID,omitempty" xml:"userID,omitempty"`
  // The mount configuration.
  VolumeMounts *InnerSandboxVolumeMount `json:"volumeMounts,omitempty" xml:"volumeMounts,omitempty"`
}

func (s E2BSandbox) String() string {
  return dara.Prettify(s)
}

func (s E2BSandbox) GoString() string {
  return s.String()
}

func (s *E2BSandbox) GetAccessEndpoint() *string  {
  return s.AccessEndpoint
}

func (s *E2BSandbox) GetAllowInternetAccess() *bool  {
  return s.AllowInternetAccess
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

func (s *E2BSandbox) GetEnvVars() map[string]*string  {
  return s.EnvVars
}

func (s *E2BSandbox) GetEnvdAccessToken() *string  {
  return s.EnvdAccessToken
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

func (s *E2BSandbox) GetGeneration() *int32  {
  return s.Generation
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

func (s *E2BSandbox) GetResourceGroupID() *string  {
  return s.ResourceGroupID
}

func (s *E2BSandbox) GetRuntime() *InnerSandboxRuntimeConfig  {
  return s.Runtime
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

func (s *E2BSandbox) GetTeamID() *string  {
  return s.TeamID
}

func (s *E2BSandbox) GetTeamName() *string  {
  return s.TeamName
}

func (s *E2BSandbox) GetTeamPlan() *string  {
  return s.TeamPlan
}

func (s *E2BSandbox) GetTemplateID() *string  {
  return s.TemplateID
}

func (s *E2BSandbox) GetTemplateName() *string  {
  return s.TemplateName
}

func (s *E2BSandbox) GetUserID() *string  {
  return s.UserID
}

func (s *E2BSandbox) GetVolumeMounts() *InnerSandboxVolumeMount  {
  return s.VolumeMounts
}

func (s *E2BSandbox) SetAccessEndpoint(v string) *E2BSandbox {
  s.AccessEndpoint = &v
  return s
}

func (s *E2BSandbox) SetAllowInternetAccess(v bool) *E2BSandbox {
  s.AllowInternetAccess = &v
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

func (s *E2BSandbox) SetEnvVars(v map[string]*string) *E2BSandbox {
  s.EnvVars = v
  return s
}

func (s *E2BSandbox) SetEnvdAccessToken(v string) *E2BSandbox {
  s.EnvdAccessToken = &v
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

func (s *E2BSandbox) SetGeneration(v int32) *E2BSandbox {
  s.Generation = &v
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

func (s *E2BSandbox) SetResourceGroupID(v string) *E2BSandbox {
  s.ResourceGroupID = &v
  return s
}

func (s *E2BSandbox) SetRuntime(v *InnerSandboxRuntimeConfig) *E2BSandbox {
  s.Runtime = v
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

func (s *E2BSandbox) SetTeamID(v string) *E2BSandbox {
  s.TeamID = &v
  return s
}

func (s *E2BSandbox) SetTeamName(v string) *E2BSandbox {
  s.TeamName = &v
  return s
}

func (s *E2BSandbox) SetTeamPlan(v string) *E2BSandbox {
  s.TeamPlan = &v
  return s
}

func (s *E2BSandbox) SetTemplateID(v string) *E2BSandbox {
  s.TemplateID = &v
  return s
}

func (s *E2BSandbox) SetTemplateName(v string) *E2BSandbox {
  s.TemplateName = &v
  return s
}

func (s *E2BSandbox) SetUserID(v string) *E2BSandbox {
  s.UserID = &v
  return s
}

func (s *E2BSandbox) SetVolumeMounts(v *InnerSandboxVolumeMount) *E2BSandbox {
  s.VolumeMounts = v
  return s
}

func (s *E2BSandbox) Validate() error {
  if s.Network != nil {
    if err := s.Network.Validate(); err != nil {
      return err
    }
  }
  if s.Runtime != nil {
    if err := s.Runtime.Validate(); err != nil {
      return err
    }
  }
  if s.VolumeMounts != nil {
    if err := s.VolumeMounts.Validate(); err != nil {
      return err
    }
  }
  return nil
}

