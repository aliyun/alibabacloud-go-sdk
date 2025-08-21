// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSceneDefensePolicyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetPolicyId(v string) *EnableSceneDefensePolicyRequest
  GetPolicyId() *string 
}

type EnableSceneDefensePolicyRequest struct {
  // The ID of the policy that you want to enable.
  // 
  // > You can call the [DescribeSceneDefensePolicies](https://help.aliyun.com/document_detail/159382.html) operation to query the IDs of all policies.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 321a-fd31-df51-****
  PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s EnableSceneDefensePolicyRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSceneDefensePolicyRequest) GoString() string {
  return s.String()
}

func (s *EnableSceneDefensePolicyRequest) GetPolicyId() *string  {
  return s.PolicyId
}

func (s *EnableSceneDefensePolicyRequest) SetPolicyId(v string) *EnableSceneDefensePolicyRequest {
  s.PolicyId = &v
  return s
}

func (s *EnableSceneDefensePolicyRequest) Validate() error {
  return dara.Validate(s)
}

