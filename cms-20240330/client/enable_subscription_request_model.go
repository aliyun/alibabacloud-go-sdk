// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSubscriptionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetWorkspace(v string) *EnableSubscriptionRequest
  GetWorkspace() *string 
}

type EnableSubscriptionRequest struct {
  // example:
  // 
  // workspace-test
  Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s EnableSubscriptionRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSubscriptionRequest) GoString() string {
  return s.String()
}

func (s *EnableSubscriptionRequest) GetWorkspace() *string  {
  return s.Workspace
}

func (s *EnableSubscriptionRequest) SetWorkspace(v string) *EnableSubscriptionRequest {
  s.Workspace = &v
  return s
}

func (s *EnableSubscriptionRequest) Validate() error {
  return dara.Validate(s)
}

