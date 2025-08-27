// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntityAddRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEntityDOList(v []*EntityAddRequestEntityDOList) *EntityAddRequest
  GetEntityDOList() []*EntityAddRequestEntityDOList 
  SetThirdpartId(v string) *EntityAddRequest
  GetThirdpartId() *string 
}

type EntityAddRequest struct {
  EntityDOList []*EntityAddRequestEntityDOList `json:"entity_d_o_list,omitempty" xml:"entity_d_o_list,omitempty" type:"Repeated"`
  // This parameter is required.
  // 
  // example:
  // 
  // 340049
  ThirdpartId *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
}

func (s EntityAddRequest) String() string {
  return dara.Prettify(s)
}

func (s EntityAddRequest) GoString() string {
  return s.String()
}

func (s *EntityAddRequest) GetEntityDOList() []*EntityAddRequestEntityDOList  {
  return s.EntityDOList
}

func (s *EntityAddRequest) GetThirdpartId() *string  {
  return s.ThirdpartId
}

func (s *EntityAddRequest) SetEntityDOList(v []*EntityAddRequestEntityDOList) *EntityAddRequest {
  s.EntityDOList = v
  return s
}

func (s *EntityAddRequest) SetThirdpartId(v string) *EntityAddRequest {
  s.ThirdpartId = &v
  return s
}

func (s *EntityAddRequest) Validate() error {
  return dara.Validate(s)
}

type EntityAddRequestEntityDOList struct {
  // example:
  // 
  // 12345
  EntityId *string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
  // example:
  // 
  // 1
  EntityType *string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
}

func (s EntityAddRequestEntityDOList) String() string {
  return dara.Prettify(s)
}

func (s EntityAddRequestEntityDOList) GoString() string {
  return s.String()
}

func (s *EntityAddRequestEntityDOList) GetEntityId() *string  {
  return s.EntityId
}

func (s *EntityAddRequestEntityDOList) GetEntityType() *string  {
  return s.EntityType
}

func (s *EntityAddRequestEntityDOList) SetEntityId(v string) *EntityAddRequestEntityDOList {
  s.EntityId = &v
  return s
}

func (s *EntityAddRequestEntityDOList) SetEntityType(v string) *EntityAddRequestEntityDOList {
  s.EntityType = &v
  return s
}

func (s *EntityAddRequestEntityDOList) Validate() error {
  return dara.Validate(s)
}

