// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntityDeleteRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDelAll(v bool) *EntityDeleteRequest
  GetDelAll() *bool 
  SetEntityDOList(v []*EntityDeleteRequestEntityDOList) *EntityDeleteRequest
  GetEntityDOList() []*EntityDeleteRequestEntityDOList 
  SetThirdpartId(v string) *EntityDeleteRequest
  GetThirdpartId() *string 
}

type EntityDeleteRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // true
  DelAll *bool `json:"del_all,omitempty" xml:"del_all,omitempty"`
  EntityDOList []*EntityDeleteRequestEntityDOList `json:"entity_d_o_list,omitempty" xml:"entity_d_o_list,omitempty" type:"Repeated"`
  // This parameter is required.
  // 
  // example:
  // 
  // 12345
  ThirdpartId *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
}

func (s EntityDeleteRequest) String() string {
  return dara.Prettify(s)
}

func (s EntityDeleteRequest) GoString() string {
  return s.String()
}

func (s *EntityDeleteRequest) GetDelAll() *bool  {
  return s.DelAll
}

func (s *EntityDeleteRequest) GetEntityDOList() []*EntityDeleteRequestEntityDOList  {
  return s.EntityDOList
}

func (s *EntityDeleteRequest) GetThirdpartId() *string  {
  return s.ThirdpartId
}

func (s *EntityDeleteRequest) SetDelAll(v bool) *EntityDeleteRequest {
  s.DelAll = &v
  return s
}

func (s *EntityDeleteRequest) SetEntityDOList(v []*EntityDeleteRequestEntityDOList) *EntityDeleteRequest {
  s.EntityDOList = v
  return s
}

func (s *EntityDeleteRequest) SetThirdpartId(v string) *EntityDeleteRequest {
  s.ThirdpartId = &v
  return s
}

func (s *EntityDeleteRequest) Validate() error {
  return dara.Validate(s)
}

type EntityDeleteRequestEntityDOList struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 123
  EntityId *string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  EntityType *string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
}

func (s EntityDeleteRequestEntityDOList) String() string {
  return dara.Prettify(s)
}

func (s EntityDeleteRequestEntityDOList) GoString() string {
  return s.String()
}

func (s *EntityDeleteRequestEntityDOList) GetEntityId() *string  {
  return s.EntityId
}

func (s *EntityDeleteRequestEntityDOList) GetEntityType() *string  {
  return s.EntityType
}

func (s *EntityDeleteRequestEntityDOList) SetEntityId(v string) *EntityDeleteRequestEntityDOList {
  s.EntityId = &v
  return s
}

func (s *EntityDeleteRequestEntityDOList) SetEntityType(v string) *EntityDeleteRequestEntityDOList {
  s.EntityType = &v
  return s
}

func (s *EntityDeleteRequestEntityDOList) Validate() error {
  return dara.Validate(s)
}

