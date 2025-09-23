// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableLayer7CCRuleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDomain(v string) *EnableLayer7CCRuleRequest
  GetDomain() *string 
  SetResourceGroupId(v string) *EnableLayer7CCRuleRequest
  GetResourceGroupId() *string 
  SetSourceIp(v string) *EnableLayer7CCRuleRequest
  GetSourceIp() *string 
}

type EnableLayer7CCRuleRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // www.aliyun.com
  Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
  // example:
  // 
  // test
  ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
  // example:
  // 
  // 1.1.1.1
  SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s EnableLayer7CCRuleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableLayer7CCRuleRequest) GoString() string {
  return s.String()
}

func (s *EnableLayer7CCRuleRequest) GetDomain() *string  {
  return s.Domain
}

func (s *EnableLayer7CCRuleRequest) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *EnableLayer7CCRuleRequest) GetSourceIp() *string  {
  return s.SourceIp
}

func (s *EnableLayer7CCRuleRequest) SetDomain(v string) *EnableLayer7CCRuleRequest {
  s.Domain = &v
  return s
}

func (s *EnableLayer7CCRuleRequest) SetResourceGroupId(v string) *EnableLayer7CCRuleRequest {
  s.ResourceGroupId = &v
  return s
}

func (s *EnableLayer7CCRuleRequest) SetSourceIp(v string) *EnableLayer7CCRuleRequest {
  s.SourceIp = &v
  return s
}

func (s *EnableLayer7CCRuleRequest) Validate() error {
  return dara.Validate(s)
}

