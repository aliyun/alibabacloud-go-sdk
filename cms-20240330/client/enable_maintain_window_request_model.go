// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMaintainWindowRequest interface {
  dara.Model
  String() string
  GoString() string
  SetWorkspace(v string) *EnableMaintainWindowRequest
  GetWorkspace() *string 
}

type EnableMaintainWindowRequest struct {
  // example:
  // 
  // workspace-test
  Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s EnableMaintainWindowRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableMaintainWindowRequest) GoString() string {
  return s.String()
}

func (s *EnableMaintainWindowRequest) GetWorkspace() *string  {
  return s.Workspace
}

func (s *EnableMaintainWindowRequest) SetWorkspace(v string) *EnableMaintainWindowRequest {
  s.Workspace = &v
  return s
}

func (s *EnableMaintainWindowRequest) Validate() error {
  return dara.Validate(s)
}

