// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableFirewallRulesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBClusterId(v string) *EnableFirewallRulesRequest
  GetDBClusterId() *string 
  SetEnable(v bool) *EnableFirewallRulesRequest
  GetEnable() *bool 
  SetOwnerAccount(v string) *EnableFirewallRulesRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableFirewallRulesRequest
  GetOwnerId() *int64 
  SetResourceOwnerAccount(v string) *EnableFirewallRulesRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableFirewallRulesRequest
  GetResourceOwnerId() *int64 
  SetRuleNameList(v string) *EnableFirewallRulesRequest
  GetRuleNameList() *string 
}

type EnableFirewallRulesRequest struct {
  // The cluster ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // pc-************
  DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
  // Specifies whether to enable or disable the specified firewall rules. Valid values:
  // 
  // - **true**: Enable.
  // 
  // - **false**: Disable.
  // 
  // > This parameter takes effect only when a value is specified for the **RuleNameList*	- parameter.
  // 
  // example:
  // 
  // true
  Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // A comma-separated list of the firewall rule names to enable or disable.
  // 
  // > You can find rule names under **Security Management****SQL Firewall*	- tab of the cluster.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // test111
  RuleNameList *string `json:"RuleNameList,omitempty" xml:"RuleNameList,omitempty"`
}

func (s EnableFirewallRulesRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableFirewallRulesRequest) GoString() string {
  return s.String()
}

func (s *EnableFirewallRulesRequest) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *EnableFirewallRulesRequest) GetEnable() *bool  {
  return s.Enable
}

func (s *EnableFirewallRulesRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableFirewallRulesRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableFirewallRulesRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableFirewallRulesRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableFirewallRulesRequest) GetRuleNameList() *string  {
  return s.RuleNameList
}

func (s *EnableFirewallRulesRequest) SetDBClusterId(v string) *EnableFirewallRulesRequest {
  s.DBClusterId = &v
  return s
}

func (s *EnableFirewallRulesRequest) SetEnable(v bool) *EnableFirewallRulesRequest {
  s.Enable = &v
  return s
}

func (s *EnableFirewallRulesRequest) SetOwnerAccount(v string) *EnableFirewallRulesRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableFirewallRulesRequest) SetOwnerId(v int64) *EnableFirewallRulesRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableFirewallRulesRequest) SetResourceOwnerAccount(v string) *EnableFirewallRulesRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableFirewallRulesRequest) SetResourceOwnerId(v int64) *EnableFirewallRulesRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableFirewallRulesRequest) SetRuleNameList(v string) *EnableFirewallRulesRequest {
  s.RuleNameList = &v
  return s
}

func (s *EnableFirewallRulesRequest) Validate() error {
  return dara.Validate(s)
}

