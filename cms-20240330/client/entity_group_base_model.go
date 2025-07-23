// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntityGroupBase interface {
  dara.Model
  String() string
  GoString() string
  SetDescription(v string) *EntityGroupBase
  GetDescription() *string 
  SetEntityGroupId(v string) *EntityGroupBase
  GetEntityGroupId() *string 
  SetEntityGroupName(v string) *EntityGroupBase
  GetEntityGroupName() *string 
  SetEntityQueries(v []*EntityGroupBaseEntityQueries) *EntityGroupBase
  GetEntityQueries() []*EntityGroupBaseEntityQueries 
  SetEntityRules(v *EntityDiscoverRule) *EntityGroupBase
  GetEntityRules() *EntityDiscoverRule 
  SetRegionId(v string) *EntityGroupBase
  GetRegionId() *string 
  SetUserId(v string) *EntityGroupBase
  GetUserId() *string 
  SetWorkspace(v string) *EntityGroupBase
  GetWorkspace() *string 
}

type EntityGroupBase struct {
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  EntityGroupId *string `json:"entityGroupId,omitempty" xml:"entityGroupId,omitempty"`
  EntityGroupName *string `json:"entityGroupName,omitempty" xml:"entityGroupName,omitempty"`
  EntityQueries []*EntityGroupBaseEntityQueries `json:"entityQueries,omitempty" xml:"entityQueries,omitempty" type:"Repeated"`
  EntityRules *EntityDiscoverRule `json:"entityRules,omitempty" xml:"entityRules,omitempty"`
  RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
  Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s EntityGroupBase) String() string {
  return dara.Prettify(s)
}

func (s EntityGroupBase) GoString() string {
  return s.String()
}

func (s *EntityGroupBase) GetDescription() *string  {
  return s.Description
}

func (s *EntityGroupBase) GetEntityGroupId() *string  {
  return s.EntityGroupId
}

func (s *EntityGroupBase) GetEntityGroupName() *string  {
  return s.EntityGroupName
}

func (s *EntityGroupBase) GetEntityQueries() []*EntityGroupBaseEntityQueries  {
  return s.EntityQueries
}

func (s *EntityGroupBase) GetEntityRules() *EntityDiscoverRule  {
  return s.EntityRules
}

func (s *EntityGroupBase) GetRegionId() *string  {
  return s.RegionId
}

func (s *EntityGroupBase) GetUserId() *string  {
  return s.UserId
}

func (s *EntityGroupBase) GetWorkspace() *string  {
  return s.Workspace
}

func (s *EntityGroupBase) SetDescription(v string) *EntityGroupBase {
  s.Description = &v
  return s
}

func (s *EntityGroupBase) SetEntityGroupId(v string) *EntityGroupBase {
  s.EntityGroupId = &v
  return s
}

func (s *EntityGroupBase) SetEntityGroupName(v string) *EntityGroupBase {
  s.EntityGroupName = &v
  return s
}

func (s *EntityGroupBase) SetEntityQueries(v []*EntityGroupBaseEntityQueries) *EntityGroupBase {
  s.EntityQueries = v
  return s
}

func (s *EntityGroupBase) SetEntityRules(v *EntityDiscoverRule) *EntityGroupBase {
  s.EntityRules = v
  return s
}

func (s *EntityGroupBase) SetRegionId(v string) *EntityGroupBase {
  s.RegionId = &v
  return s
}

func (s *EntityGroupBase) SetUserId(v string) *EntityGroupBase {
  s.UserId = &v
  return s
}

func (s *EntityGroupBase) SetWorkspace(v string) *EntityGroupBase {
  s.Workspace = &v
  return s
}

func (s *EntityGroupBase) Validate() error {
  return dara.Validate(s)
}

type EntityGroupBaseEntityQueries struct {
  EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
  Spl *string `json:"spl,omitempty" xml:"spl,omitempty"`
}

func (s EntityGroupBaseEntityQueries) String() string {
  return dara.Prettify(s)
}

func (s EntityGroupBaseEntityQueries) GoString() string {
  return s.String()
}

func (s *EntityGroupBaseEntityQueries) GetEntityType() *string  {
  return s.EntityType
}

func (s *EntityGroupBaseEntityQueries) GetSpl() *string  {
  return s.Spl
}

func (s *EntityGroupBaseEntityQueries) SetEntityType(v string) *EntityGroupBaseEntityQueries {
  s.EntityType = &v
  return s
}

func (s *EntityGroupBaseEntityQueries) SetSpl(v string) *EntityGroupBaseEntityQueries {
  s.Spl = &v
  return s
}

func (s *EntityGroupBaseEntityQueries) Validate() error {
  return dara.Validate(s)
}

