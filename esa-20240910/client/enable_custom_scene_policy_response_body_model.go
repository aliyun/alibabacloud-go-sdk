// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCustomScenePolicyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetPolicyId(v int64) *EnableCustomScenePolicyResponseBody
  GetPolicyId() *int64 
  SetRequestId(v string) *EnableCustomScenePolicyResponseBody
  GetRequestId() *string 
}

type EnableCustomScenePolicyResponseBody struct {
  // The policy ID.
  // 
  // example:
  // 
  // 1
  PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 65C66B7B-671A-8297-9187-2R5477247B76
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCustomScenePolicyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableCustomScenePolicyResponseBody) GoString() string {
  return s.String()
}

func (s *EnableCustomScenePolicyResponseBody) GetPolicyId() *int64  {
  return s.PolicyId
}

func (s *EnableCustomScenePolicyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableCustomScenePolicyResponseBody) SetPolicyId(v int64) *EnableCustomScenePolicyResponseBody {
  s.PolicyId = &v
  return s
}

func (s *EnableCustomScenePolicyResponseBody) SetRequestId(v string) *EnableCustomScenePolicyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableCustomScenePolicyResponseBody) Validate() error {
  return dara.Validate(s)
}

