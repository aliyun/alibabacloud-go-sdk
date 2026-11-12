// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableFirewallRuleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableFirewallRuleRequest
  GetClientToken() *string 
  SetInstanceId(v string) *EnableFirewallRuleRequest
  GetInstanceId() *string 
  SetRegionId(v string) *EnableFirewallRuleRequest
  GetRegionId() *string 
  SetRemark(v string) *EnableFirewallRuleRequest
  GetRemark() *string 
  SetRuleId(v string) *EnableFirewallRuleRequest
  GetRuleId() *string 
  SetSourceCidrIp(v string) *EnableFirewallRuleRequest
  GetSourceCidrIp() *string 
}

type EnableFirewallRuleRequest struct {
  // The client token that is used to ensure the idempotence of the request. A client token is generated from your client and must be unique among different requests. The value of ClientToken can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
  // 
  // example:
  // 
  // 123e4567-e89b-12d3-a456-426655440000
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // The instance ID of the specified simple application server.
  // 
  // example:
  // 
  // 2ad1ae67295445f598017499dc****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The region ID of the specified simple application server.
  // 
  // You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to query the list of supported Alibaba Cloud regions.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The remarks of the firewall rule.
  // 
  // example:
  // 
  // Custom
  Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
  // The firewall rule ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // eeea34d9867b4d55a4ff8d5fcfbd****
  RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
  // The IP address or CIDR block.
  // 
  // example:
  // 
  // 10.147.33.**
  SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s EnableFirewallRuleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableFirewallRuleRequest) GoString() string {
  return s.String()
}

func (s *EnableFirewallRuleRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableFirewallRuleRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableFirewallRuleRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableFirewallRuleRequest) GetRemark() *string  {
  return s.Remark
}

func (s *EnableFirewallRuleRequest) GetRuleId() *string  {
  return s.RuleId
}

func (s *EnableFirewallRuleRequest) GetSourceCidrIp() *string  {
  return s.SourceCidrIp
}

func (s *EnableFirewallRuleRequest) SetClientToken(v string) *EnableFirewallRuleRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableFirewallRuleRequest) SetInstanceId(v string) *EnableFirewallRuleRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableFirewallRuleRequest) SetRegionId(v string) *EnableFirewallRuleRequest {
  s.RegionId = &v
  return s
}

func (s *EnableFirewallRuleRequest) SetRemark(v string) *EnableFirewallRuleRequest {
  s.Remark = &v
  return s
}

func (s *EnableFirewallRuleRequest) SetRuleId(v string) *EnableFirewallRuleRequest {
  s.RuleId = &v
  return s
}

func (s *EnableFirewallRuleRequest) SetSourceCidrIp(v string) *EnableFirewallRuleRequest {
  s.SourceCidrIp = &v
  return s
}

func (s *EnableFirewallRuleRequest) Validate() error {
  return dara.Validate(s)
}

