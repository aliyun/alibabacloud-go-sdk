// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCustomPrivacyPolicyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCustomPrivacyPolicyId(v string) *EnableCustomPrivacyPolicyRequest
  GetCustomPrivacyPolicyId() *string 
  SetInstanceId(v string) *EnableCustomPrivacyPolicyRequest
  GetInstanceId() *string 
}

type EnableCustomPrivacyPolicyRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // pp_neagxpoznsjdtxxxxx
  CustomPrivacyPolicyId *string `json:"CustomPrivacyPolicyId,omitempty" xml:"CustomPrivacyPolicyId,omitempty"`
  // IDaaS EIAM实例的ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableCustomPrivacyPolicyRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCustomPrivacyPolicyRequest) GoString() string {
  return s.String()
}

func (s *EnableCustomPrivacyPolicyRequest) GetCustomPrivacyPolicyId() *string  {
  return s.CustomPrivacyPolicyId
}

func (s *EnableCustomPrivacyPolicyRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableCustomPrivacyPolicyRequest) SetCustomPrivacyPolicyId(v string) *EnableCustomPrivacyPolicyRequest {
  s.CustomPrivacyPolicyId = &v
  return s
}

func (s *EnableCustomPrivacyPolicyRequest) SetInstanceId(v string) *EnableCustomPrivacyPolicyRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableCustomPrivacyPolicyRequest) Validate() error {
  return dara.Validate(s)
}

