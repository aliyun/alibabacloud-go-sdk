// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntityDeleteShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDelAll(v bool) *EntityDeleteShrinkRequest
  GetDelAll() *bool 
  SetEntityDOListShrink(v string) *EntityDeleteShrinkRequest
  GetEntityDOListShrink() *string 
  SetThirdpartId(v string) *EntityDeleteShrinkRequest
  GetThirdpartId() *string 
}

type EntityDeleteShrinkRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // true
  DelAll *bool `json:"del_all,omitempty" xml:"del_all,omitempty"`
  EntityDOListShrink *string `json:"entity_d_o_list,omitempty" xml:"entity_d_o_list,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 12345
  ThirdpartId *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
}

func (s EntityDeleteShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EntityDeleteShrinkRequest) GoString() string {
  return s.String()
}

func (s *EntityDeleteShrinkRequest) GetDelAll() *bool  {
  return s.DelAll
}

func (s *EntityDeleteShrinkRequest) GetEntityDOListShrink() *string  {
  return s.EntityDOListShrink
}

func (s *EntityDeleteShrinkRequest) GetThirdpartId() *string  {
  return s.ThirdpartId
}

func (s *EntityDeleteShrinkRequest) SetDelAll(v bool) *EntityDeleteShrinkRequest {
  s.DelAll = &v
  return s
}

func (s *EntityDeleteShrinkRequest) SetEntityDOListShrink(v string) *EntityDeleteShrinkRequest {
  s.EntityDOListShrink = &v
  return s
}

func (s *EntityDeleteShrinkRequest) SetThirdpartId(v string) *EntityDeleteShrinkRequest {
  s.ThirdpartId = &v
  return s
}

func (s *EntityDeleteShrinkRequest) Validate() error {
  return dara.Validate(s)
}

