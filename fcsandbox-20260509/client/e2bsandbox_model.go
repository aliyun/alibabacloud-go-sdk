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
  SetCpuCount(v int32) *E2BSandbox
  GetCpuCount() *int32 
  SetDiskSizeMB(v int32) *E2BSandbox
  GetDiskSizeMB() *int32 
  SetDomain(v string) *E2BSandbox
  GetDomain() *string 
  SetEndAt(v string) *E2BSandbox
  GetEndAt() *string 
  SetFcFunctionName(v string) *E2BSandbox
  GetFcFunctionName() *string 
  SetFcInstanceID(v string) *E2BSandbox
  GetFcInstanceID() *string 
  SetFcSessionID(v string) *E2BSandbox
  GetFcSessionID() *string 
  SetMemoryMB(v int32) *E2BSandbox
  GetMemoryMB() *int32 
  SetMetadata(v map[string]*string) *E2BSandbox
  GetMetadata() map[string]*string 
  SetResourceGroupID(v string) *E2BSandbox
  GetResourceGroupID() *string 
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
  SetTemplateID(v string) *E2BSandbox
  GetTemplateID() *string 
  SetTemplateName(v string) *E2BSandbox
  GetTemplateName() *string 
  SetUserID(v string) *E2BSandbox
  GetUserID() *string 
}

type E2BSandbox struct {
  AccessEndpoint *string `json:"accessEndpoint,omitempty" xml:"accessEndpoint,omitempty"`
  CpuCount *int32 `json:"cpuCount,omitempty" xml:"cpuCount,omitempty"`
  DiskSizeMB *int32 `json:"diskSizeMB,omitempty" xml:"diskSizeMB,omitempty"`
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  EndAt *string `json:"endAt,omitempty" xml:"endAt,omitempty"`
  FcFunctionName *string `json:"fcFunctionName,omitempty" xml:"fcFunctionName,omitempty"`
  FcInstanceID *string `json:"fcInstanceID,omitempty" xml:"fcInstanceID,omitempty"`
  FcSessionID *string `json:"fcSessionID,omitempty" xml:"fcSessionID,omitempty"`
  MemoryMB *int32 `json:"memoryMB,omitempty" xml:"memoryMB,omitempty"`
  Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
  ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
  SandboxID *string `json:"sandboxID,omitempty" xml:"sandboxID,omitempty"`
  StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
  State *string `json:"state,omitempty" xml:"state,omitempty"`
  TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
  TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
  TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
  TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
  UserID *string `json:"userID,omitempty" xml:"userID,omitempty"`
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

func (s *E2BSandbox) GetFcFunctionName() *string  {
  return s.FcFunctionName
}

func (s *E2BSandbox) GetFcInstanceID() *string  {
  return s.FcInstanceID
}

func (s *E2BSandbox) GetFcSessionID() *string  {
  return s.FcSessionID
}

func (s *E2BSandbox) GetMemoryMB() *int32  {
  return s.MemoryMB
}

func (s *E2BSandbox) GetMetadata() map[string]*string  {
  return s.Metadata
}

func (s *E2BSandbox) GetResourceGroupID() *string  {
  return s.ResourceGroupID
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

func (s *E2BSandbox) GetTemplateID() *string  {
  return s.TemplateID
}

func (s *E2BSandbox) GetTemplateName() *string  {
  return s.TemplateName
}

func (s *E2BSandbox) GetUserID() *string  {
  return s.UserID
}

func (s *E2BSandbox) SetAccessEndpoint(v string) *E2BSandbox {
  s.AccessEndpoint = &v
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

func (s *E2BSandbox) SetMemoryMB(v int32) *E2BSandbox {
  s.MemoryMB = &v
  return s
}

func (s *E2BSandbox) SetMetadata(v map[string]*string) *E2BSandbox {
  s.Metadata = v
  return s
}

func (s *E2BSandbox) SetResourceGroupID(v string) *E2BSandbox {
  s.ResourceGroupID = &v
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

func (s *E2BSandbox) Validate() error {
  return dara.Validate(s)
}

