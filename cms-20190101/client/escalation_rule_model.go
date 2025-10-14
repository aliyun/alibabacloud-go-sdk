// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEscalationRule interface {
  dara.Model
  String() string
  GoString() string
  SetCreateTime(v string) *EscalationRule
  GetCreateTime() *string 
  SetDescription(v string) *EscalationRule
  GetDescription() *string 
  SetEscalations(v []*EscalationRuleEscalations) *EscalationRule
  GetEscalations() []*EscalationRuleEscalations 
  SetName(v string) *EscalationRule
  GetName() *string 
  SetUpdateTime(v string) *EscalationRule
  GetUpdateTime() *string 
  SetUserId(v string) *EscalationRule
  GetUserId() *string 
  SetUuid(v string) *EscalationRule
  GetUuid() *string 
}

type EscalationRule struct {
  CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
  Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
  // This parameter is required.
  Escalations []*EscalationRuleEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Repeated"`
  // This parameter is required.
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
  UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
  Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s EscalationRule) String() string {
  return dara.Prettify(s)
}

func (s EscalationRule) GoString() string {
  return s.String()
}

func (s *EscalationRule) GetCreateTime() *string  {
  return s.CreateTime
}

func (s *EscalationRule) GetDescription() *string  {
  return s.Description
}

func (s *EscalationRule) GetEscalations() []*EscalationRuleEscalations  {
  return s.Escalations
}

func (s *EscalationRule) GetName() *string  {
  return s.Name
}

func (s *EscalationRule) GetUpdateTime() *string  {
  return s.UpdateTime
}

func (s *EscalationRule) GetUserId() *string  {
  return s.UserId
}

func (s *EscalationRule) GetUuid() *string  {
  return s.Uuid
}

func (s *EscalationRule) SetCreateTime(v string) *EscalationRule {
  s.CreateTime = &v
  return s
}

func (s *EscalationRule) SetDescription(v string) *EscalationRule {
  s.Description = &v
  return s
}

func (s *EscalationRule) SetEscalations(v []*EscalationRuleEscalations) *EscalationRule {
  s.Escalations = v
  return s
}

func (s *EscalationRule) SetName(v string) *EscalationRule {
  s.Name = &v
  return s
}

func (s *EscalationRule) SetUpdateTime(v string) *EscalationRule {
  s.UpdateTime = &v
  return s
}

func (s *EscalationRule) SetUserId(v string) *EscalationRule {
  s.UserId = &v
  return s
}

func (s *EscalationRule) SetUuid(v string) *EscalationRule {
  s.Uuid = &v
  return s
}

func (s *EscalationRule) Validate() error {
  if s.Escalations != nil {
    for _, item := range s.Escalations {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EscalationRuleEscalations struct {
  BackupContactGroups []*string `json:"BackupContactGroups,omitempty" xml:"BackupContactGroups,omitempty" type:"Repeated"`
  ContactGroups []*string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Repeated"`
  ContactGroupsByLevel *EscalationRuleEscalationsContactGroupsByLevel `json:"ContactGroupsByLevel,omitempty" xml:"ContactGroupsByLevel,omitempty" type:"Struct"`
  EscalateMin *int64 `json:"EscalateMin,omitempty" xml:"EscalateMin,omitempty"`
}

func (s EscalationRuleEscalations) String() string {
  return dara.Prettify(s)
}

func (s EscalationRuleEscalations) GoString() string {
  return s.String()
}

func (s *EscalationRuleEscalations) GetBackupContactGroups() []*string  {
  return s.BackupContactGroups
}

func (s *EscalationRuleEscalations) GetContactGroups() []*string  {
  return s.ContactGroups
}

func (s *EscalationRuleEscalations) GetContactGroupsByLevel() *EscalationRuleEscalationsContactGroupsByLevel  {
  return s.ContactGroupsByLevel
}

func (s *EscalationRuleEscalations) GetEscalateMin() *int64  {
  return s.EscalateMin
}

func (s *EscalationRuleEscalations) SetBackupContactGroups(v []*string) *EscalationRuleEscalations {
  s.BackupContactGroups = v
  return s
}

func (s *EscalationRuleEscalations) SetContactGroups(v []*string) *EscalationRuleEscalations {
  s.ContactGroups = v
  return s
}

func (s *EscalationRuleEscalations) SetContactGroupsByLevel(v *EscalationRuleEscalationsContactGroupsByLevel) *EscalationRuleEscalations {
  s.ContactGroupsByLevel = v
  return s
}

func (s *EscalationRuleEscalations) SetEscalateMin(v int64) *EscalationRuleEscalations {
  s.EscalateMin = &v
  return s
}

func (s *EscalationRuleEscalations) Validate() error {
  if s.ContactGroupsByLevel != nil {
    if err := s.ContactGroupsByLevel.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EscalationRuleEscalationsContactGroupsByLevel struct {
  Critical []*string `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Repeated"`
  Error []*string `json:"Error,omitempty" xml:"Error,omitempty" type:"Repeated"`
  Info []*string `json:"Info,omitempty" xml:"Info,omitempty" type:"Repeated"`
  Resolve []*string `json:"Resolve,omitempty" xml:"Resolve,omitempty" type:"Repeated"`
  Warning []*string `json:"Warning,omitempty" xml:"Warning,omitempty" type:"Repeated"`
}

func (s EscalationRuleEscalationsContactGroupsByLevel) String() string {
  return dara.Prettify(s)
}

func (s EscalationRuleEscalationsContactGroupsByLevel) GoString() string {
  return s.String()
}

func (s *EscalationRuleEscalationsContactGroupsByLevel) GetCritical() []*string  {
  return s.Critical
}

func (s *EscalationRuleEscalationsContactGroupsByLevel) GetError() []*string  {
  return s.Error
}

func (s *EscalationRuleEscalationsContactGroupsByLevel) GetInfo() []*string  {
  return s.Info
}

func (s *EscalationRuleEscalationsContactGroupsByLevel) GetResolve() []*string  {
  return s.Resolve
}

func (s *EscalationRuleEscalationsContactGroupsByLevel) GetWarning() []*string  {
  return s.Warning
}

func (s *EscalationRuleEscalationsContactGroupsByLevel) SetCritical(v []*string) *EscalationRuleEscalationsContactGroupsByLevel {
  s.Critical = v
  return s
}

func (s *EscalationRuleEscalationsContactGroupsByLevel) SetError(v []*string) *EscalationRuleEscalationsContactGroupsByLevel {
  s.Error = v
  return s
}

func (s *EscalationRuleEscalationsContactGroupsByLevel) SetInfo(v []*string) *EscalationRuleEscalationsContactGroupsByLevel {
  s.Info = v
  return s
}

func (s *EscalationRuleEscalationsContactGroupsByLevel) SetResolve(v []*string) *EscalationRuleEscalationsContactGroupsByLevel {
  s.Resolve = v
  return s
}

func (s *EscalationRuleEscalationsContactGroupsByLevel) SetWarning(v []*string) *EscalationRuleEscalationsContactGroupsByLevel {
  s.Warning = v
  return s
}

func (s *EscalationRuleEscalationsContactGroupsByLevel) Validate() error {
  return dara.Validate(s)
}

