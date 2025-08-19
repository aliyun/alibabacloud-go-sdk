// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEqualRule interface {
  dara.Model
  String() string
  GoString() string
  SetMatch(v string) *EqualRule
  GetMatch() *string 
  SetReplacement(v string) *EqualRule
  GetReplacement() *string 
}

type EqualRule struct {
  // This parameter is required.
  // 
  // example:
  // 
  // /old
  Match *string `json:"match,omitempty" xml:"match,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // /new
  Replacement *string `json:"replacement,omitempty" xml:"replacement,omitempty"`
}

func (s EqualRule) String() string {
  return dara.Prettify(s)
}

func (s EqualRule) GoString() string {
  return s.String()
}

func (s *EqualRule) GetMatch() *string  {
  return s.Match
}

func (s *EqualRule) GetReplacement() *string  {
  return s.Replacement
}

func (s *EqualRule) SetMatch(v string) *EqualRule {
  s.Match = &v
  return s
}

func (s *EqualRule) SetReplacement(v string) *EqualRule {
  s.Replacement = &v
  return s
}

func (s *EqualRule) Validate() error {
  return dara.Validate(s)
}

