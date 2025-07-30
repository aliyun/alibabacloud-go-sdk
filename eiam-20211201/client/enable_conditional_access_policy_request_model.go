// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableConditionalAccessPolicyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetConditionalAccessPolicyId(v string) *EnableConditionalAccessPolicyRequest
  GetConditionalAccessPolicyId() *string 
  SetInstanceId(v string) *EnableConditionalAccessPolicyRequest
  GetInstanceId() *string 
}

type EnableConditionalAccessPolicyRequest struct {
  // Conditional Access Policy ID
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cap_11111
  ConditionalAccessPolicyId *string `json:"ConditionalAccessPolicyId,omitempty" xml:"ConditionalAccessPolicyId,omitempty"`
  // Instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableConditionalAccessPolicyRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableConditionalAccessPolicyRequest) GoString() string {
  return s.String()
}

func (s *EnableConditionalAccessPolicyRequest) GetConditionalAccessPolicyId() *string  {
  return s.ConditionalAccessPolicyId
}

func (s *EnableConditionalAccessPolicyRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableConditionalAccessPolicyRequest) SetConditionalAccessPolicyId(v string) *EnableConditionalAccessPolicyRequest {
  s.ConditionalAccessPolicyId = &v
  return s
}

func (s *EnableConditionalAccessPolicyRequest) SetInstanceId(v string) *EnableConditionalAccessPolicyRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableConditionalAccessPolicyRequest) Validate() error {
  return dara.Validate(s)
}

