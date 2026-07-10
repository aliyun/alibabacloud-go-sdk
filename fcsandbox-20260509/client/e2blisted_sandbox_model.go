// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BListedSandbox interface {
  dara.Model
  String() string
  GoString() string
  SetCpuCount(v int32) *E2BListedSandbox
  GetCpuCount() *int32 
  SetDiskSizeMB(v int32) *E2BListedSandbox
  GetDiskSizeMB() *int32 
  SetEndAt(v string) *E2BListedSandbox
  GetEndAt() *string 
  SetMemoryMB(v int32) *E2BListedSandbox
  GetMemoryMB() *int32 
  SetMetadata(v map[string]*string) *E2BListedSandbox
  GetMetadata() map[string]*string 
  SetResourceGroupID(v string) *E2BListedSandbox
  GetResourceGroupID() *string 
  SetSandboxID(v string) *E2BListedSandbox
  GetSandboxID() *string 
  SetStartedAt(v string) *E2BListedSandbox
  GetStartedAt() *string 
  SetState(v string) *E2BListedSandbox
  GetState() *string 
  SetTeamID(v string) *E2BListedSandbox
  GetTeamID() *string 
  SetTeamName(v string) *E2BListedSandbox
  GetTeamName() *string 
  SetTemplateID(v string) *E2BListedSandbox
  GetTemplateID() *string 
  SetTemplateName(v string) *E2BListedSandbox
  GetTemplateName() *string 
  SetUserID(v string) *E2BListedSandbox
  GetUserID() *string 
}

type E2BListedSandbox struct {
  CpuCount *int32 `json:"cpuCount,omitempty" xml:"cpuCount,omitempty"`
  DiskSizeMB *int32 `json:"diskSizeMB,omitempty" xml:"diskSizeMB,omitempty"`
  EndAt *string `json:"endAt,omitempty" xml:"endAt,omitempty"`
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

func (s E2BListedSandbox) String() string {
  return dara.Prettify(s)
}

func (s E2BListedSandbox) GoString() string {
  return s.String()
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

func (s *E2BListedSandbox) GetMemoryMB() *int32  {
  return s.MemoryMB
}

func (s *E2BListedSandbox) GetMetadata() map[string]*string  {
  return s.Metadata
}

func (s *E2BListedSandbox) GetResourceGroupID() *string  {
  return s.ResourceGroupID
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

func (s *E2BListedSandbox) GetTeamID() *string  {
  return s.TeamID
}

func (s *E2BListedSandbox) GetTeamName() *string  {
  return s.TeamName
}

func (s *E2BListedSandbox) GetTemplateID() *string  {
  return s.TemplateID
}

func (s *E2BListedSandbox) GetTemplateName() *string  {
  return s.TemplateName
}

func (s *E2BListedSandbox) GetUserID() *string  {
  return s.UserID
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

func (s *E2BListedSandbox) SetMemoryMB(v int32) *E2BListedSandbox {
  s.MemoryMB = &v
  return s
}

func (s *E2BListedSandbox) SetMetadata(v map[string]*string) *E2BListedSandbox {
  s.Metadata = v
  return s
}

func (s *E2BListedSandbox) SetResourceGroupID(v string) *E2BListedSandbox {
  s.ResourceGroupID = &v
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

func (s *E2BListedSandbox) SetTeamID(v string) *E2BListedSandbox {
  s.TeamID = &v
  return s
}

func (s *E2BListedSandbox) SetTeamName(v string) *E2BListedSandbox {
  s.TeamName = &v
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

func (s *E2BListedSandbox) SetUserID(v string) *E2BListedSandbox {
  s.UserID = &v
  return s
}

func (s *E2BListedSandbox) Validate() error {
  return dara.Validate(s)
}

