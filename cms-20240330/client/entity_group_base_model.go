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
  // 实体描述。
  // 
  // example:
  // 
  // ECS 实例
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // 实体ID。
  // 
  // example:
  // 
  // eg-1234567890
  EntityGroupId *string `json:"entityGroupId,omitempty" xml:"entityGroupId,omitempty"`
  // 实体名称。
  // 
  // example:
  // 
  // ECS 全部实体
  EntityGroupName *string `json:"entityGroupName,omitempty" xml:"entityGroupName,omitempty"`
  // 实体查询规则集合。
  EntityQueries []*EntityGroupBaseEntityQueries `json:"entityQueries,omitempty" xml:"entityQueries,omitempty" type:"Repeated"`
  // 用于实体发现的规则。
  EntityRules *EntityDiscoverRule `json:"entityRules,omitempty" xml:"entityRules,omitempty"`
  // 地域ID。
  // 
  // example:
  // 
  // cn-heyuan
  RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
  // 用户ID。
  // 
  // example:
  // 
  // 1654218***343050
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
  // 工作空间。
  // 
  // example:
  // 
  // default-cms-1654218***343050-cn-hangzhou
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
  if s.EntityQueries != nil {
    for _, item := range s.EntityQueries {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.EntityRules != nil {
    if err := s.EntityRules.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EntityGroupBaseEntityQueries struct {
  // 实体类型。
  // 
  // example:
  // 
  // acs.ecs.instance
  EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
  // SPL查询语句。
  // 
  // example:
  // 
  // .entity with(type=\\"acs.ecs.instance\\") | where region_id in (\\"cn-beijing\\")
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

