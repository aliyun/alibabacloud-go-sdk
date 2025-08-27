// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntitySetShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEntityDOListShrink(v string) *EntitySetShrinkRequest
  GetEntityDOListShrink() *string 
  SetThirdpartId(v string) *EntitySetShrinkRequest
  GetThirdpartId() *string 
}

type EntitySetShrinkRequest struct {
  EntityDOListShrink *string `json:"entity_d_o_list,omitempty" xml:"entity_d_o_list,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 340049
  ThirdpartId *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
}

func (s EntitySetShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EntitySetShrinkRequest) GoString() string {
  return s.String()
}

func (s *EntitySetShrinkRequest) GetEntityDOListShrink() *string  {
  return s.EntityDOListShrink
}

func (s *EntitySetShrinkRequest) GetThirdpartId() *string  {
  return s.ThirdpartId
}

func (s *EntitySetShrinkRequest) SetEntityDOListShrink(v string) *EntitySetShrinkRequest {
  s.EntityDOListShrink = &v
  return s
}

func (s *EntitySetShrinkRequest) SetThirdpartId(v string) *EntitySetShrinkRequest {
  s.ThirdpartId = &v
  return s
}

func (s *EntitySetShrinkRequest) Validate() error {
  return dara.Validate(s)
}

