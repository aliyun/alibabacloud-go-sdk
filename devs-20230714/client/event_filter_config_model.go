// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventFilterConfig interface {
  dara.Model
  String() string
  GoString() string
  SetBranch(v *BranchFilter) *EventFilterConfig
  GetBranch() *BranchFilter 
}

type EventFilterConfig struct {
  Branch *BranchFilter `json:"branch,omitempty" xml:"branch,omitempty"`
}

func (s EventFilterConfig) String() string {
  return dara.Prettify(s)
}

func (s EventFilterConfig) GoString() string {
  return s.String()
}

func (s *EventFilterConfig) GetBranch() *BranchFilter  {
  return s.Branch
}

func (s *EventFilterConfig) SetBranch(v *BranchFilter) *EventFilterConfig {
  s.Branch = v
  return s
}

func (s *EventFilterConfig) Validate() error {
  return dara.Validate(s)
}

