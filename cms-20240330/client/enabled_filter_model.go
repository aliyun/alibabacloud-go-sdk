// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnabledFilter interface {
  dara.Model
  String() string
  GoString() string
  SetEq(v bool) *EnabledFilter
  GetEq() *bool 
}

type EnabledFilter struct {
  // 精确匹配
  Eq *bool `json:"eq,omitempty" xml:"eq,omitempty"`
}

func (s EnabledFilter) String() string {
  return dara.Prettify(s)
}

func (s EnabledFilter) GoString() string {
  return s.String()
}

func (s *EnabledFilter) GetEq() *bool  {
  return s.Eq
}

func (s *EnabledFilter) SetEq(v bool) *EnabledFilter {
  s.Eq = &v
  return s
}

func (s *EnabledFilter) Validate() error {
  return dara.Validate(s)
}

