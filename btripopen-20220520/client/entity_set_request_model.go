// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntitySetRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEntityDOList(v []*EntitySetRequestEntityDOList) *EntitySetRequest
  GetEntityDOList() []*EntitySetRequestEntityDOList 
  SetThirdpartId(v string) *EntitySetRequest
  GetThirdpartId() *string 
}

type EntitySetRequest struct {
  EntityDOList []*EntitySetRequestEntityDOList `json:"entity_d_o_list,omitempty" xml:"entity_d_o_list,omitempty" type:"Repeated"`
  // This parameter is required.
  // 
  // example:
  // 
  // 340049
  ThirdpartId *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
}

func (s EntitySetRequest) String() string {
  return dara.Prettify(s)
}

func (s EntitySetRequest) GoString() string {
  return s.String()
}

func (s *EntitySetRequest) GetEntityDOList() []*EntitySetRequestEntityDOList  {
  return s.EntityDOList
}

func (s *EntitySetRequest) GetThirdpartId() *string  {
  return s.ThirdpartId
}

func (s *EntitySetRequest) SetEntityDOList(v []*EntitySetRequestEntityDOList) *EntitySetRequest {
  s.EntityDOList = v
  return s
}

func (s *EntitySetRequest) SetThirdpartId(v string) *EntitySetRequest {
  s.ThirdpartId = &v
  return s
}

func (s *EntitySetRequest) Validate() error {
  return dara.Validate(s)
}

type EntitySetRequestEntityDOList struct {
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

func (s EntitySetRequestEntityDOList) String() string {
  return dara.Prettify(s)
}

func (s EntitySetRequestEntityDOList) GoString() string {
  return s.String()
}

func (s *EntitySetRequestEntityDOList) GetEntityId() *string  {
  return s.EntityId
}

func (s *EntitySetRequestEntityDOList) GetEntityType() *string  {
  return s.EntityType
}

func (s *EntitySetRequestEntityDOList) SetEntityId(v string) *EntitySetRequestEntityDOList {
  s.EntityId = &v
  return s
}

func (s *EntitySetRequestEntityDOList) SetEntityType(v string) *EntitySetRequestEntityDOList {
  s.EntityType = &v
  return s
}

func (s *EntitySetRequestEntityDOList) Validate() error {
  return dara.Validate(s)
}

