// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableNotifyPolicyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetWorkspace(v string) *EnableNotifyPolicyRequest
  GetWorkspace() *string 
}

type EnableNotifyPolicyRequest struct {
  // The workspace ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // default
  Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s EnableNotifyPolicyRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableNotifyPolicyRequest) GoString() string {
  return s.String()
}

func (s *EnableNotifyPolicyRequest) GetWorkspace() *string  {
  return s.Workspace
}

func (s *EnableNotifyPolicyRequest) SetWorkspace(v string) *EnableNotifyPolicyRequest {
  s.Workspace = &v
  return s
}

func (s *EnableNotifyPolicyRequest) Validate() error {
  return dara.Validate(s)
}

