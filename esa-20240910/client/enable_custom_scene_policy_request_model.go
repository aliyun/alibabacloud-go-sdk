// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCustomScenePolicyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetPolicyId(v int64) *EnableCustomScenePolicyRequest
  GetPolicyId() *int64 
}

type EnableCustomScenePolicyRequest struct {
  // This parameter is required.
  PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s EnableCustomScenePolicyRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCustomScenePolicyRequest) GoString() string {
  return s.String()
}

func (s *EnableCustomScenePolicyRequest) GetPolicyId() *int64  {
  return s.PolicyId
}

func (s *EnableCustomScenePolicyRequest) SetPolicyId(v int64) *EnableCustomScenePolicyRequest {
  s.PolicyId = &v
  return s
}

func (s *EnableCustomScenePolicyRequest) Validate() error {
  return dara.Validate(s)
}

