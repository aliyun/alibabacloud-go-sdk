// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventMatchRule interface {
  dara.Model
  String() string
  GoString() string
  SetMatchState(v bool) *EventMatchRule
  GetMatchState() *bool 
  SetName(v string) *EventMatchRule
  GetName() *string 
  SetPrefix(v string) *EventMatchRule
  GetPrefix() *string 
  SetSuffix(v string) *EventMatchRule
  GetSuffix() *string 
}

type EventMatchRule struct {
  // Specifies whether to match.
  // 
  // example:
  // 
  // true
  MatchState *bool `json:"MatchState,omitempty" xml:"MatchState,omitempty"`
  // The full name to match.
  // 
  // example:
  // 
  // testBucket/folder1/file1
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // The prefix to match.
  // 
  // example:
  // 
  // testBucket/prefix1xxx
  Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
  // The suffix to match.
  // 
  // example:
  // 
  // suffix/xxxx.test
  Suffix *string `json:"Suffix,omitempty" xml:"Suffix,omitempty"`
}

func (s EventMatchRule) String() string {
  return dara.Prettify(s)
}

func (s EventMatchRule) GoString() string {
  return s.String()
}

func (s *EventMatchRule) GetMatchState() *bool  {
  return s.MatchState
}

func (s *EventMatchRule) GetName() *string  {
  return s.Name
}

func (s *EventMatchRule) GetPrefix() *string  {
  return s.Prefix
}

func (s *EventMatchRule) GetSuffix() *string  {
  return s.Suffix
}

func (s *EventMatchRule) SetMatchState(v bool) *EventMatchRule {
  s.MatchState = &v
  return s
}

func (s *EventMatchRule) SetName(v string) *EventMatchRule {
  s.Name = &v
  return s
}

func (s *EventMatchRule) SetPrefix(v string) *EventMatchRule {
  s.Prefix = &v
  return s
}

func (s *EventMatchRule) SetSuffix(v string) *EventMatchRule {
  s.Suffix = &v
  return s
}

func (s *EventMatchRule) Validate() error {
  return dara.Validate(s)
}

