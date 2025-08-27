// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntityAddShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEntityDOListShrink(v string) *EntityAddShrinkRequest
  GetEntityDOListShrink() *string 
  SetThirdpartId(v string) *EntityAddShrinkRequest
  GetThirdpartId() *string 
}

type EntityAddShrinkRequest struct {
  EntityDOListShrink *string `json:"entity_d_o_list,omitempty" xml:"entity_d_o_list,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 340049
  ThirdpartId *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
}

func (s EntityAddShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EntityAddShrinkRequest) GoString() string {
  return s.String()
}

func (s *EntityAddShrinkRequest) GetEntityDOListShrink() *string  {
  return s.EntityDOListShrink
}

func (s *EntityAddShrinkRequest) GetThirdpartId() *string  {
  return s.ThirdpartId
}

func (s *EntityAddShrinkRequest) SetEntityDOListShrink(v string) *EntityAddShrinkRequest {
  s.EntityDOListShrink = &v
  return s
}

func (s *EntityAddShrinkRequest) SetThirdpartId(v string) *EntityAddShrinkRequest {
  s.ThirdpartId = &v
  return s
}

func (s *EntityAddShrinkRequest) Validate() error {
  return dara.Validate(s)
}

