// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSQLRateLimitingRulesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBClusterId(v string) *EnableSQLRateLimitingRulesRequest
  GetDBClusterId() *string 
  SetEnable(v bool) *EnableSQLRateLimitingRulesRequest
  GetEnable() *bool 
  SetOwnerAccount(v string) *EnableSQLRateLimitingRulesRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableSQLRateLimitingRulesRequest
  GetOwnerId() *int64 
  SetResourceOwnerAccount(v string) *EnableSQLRateLimitingRulesRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableSQLRateLimitingRulesRequest
  GetResourceOwnerId() *int64 
  SetRuleNameList(v string) *EnableSQLRateLimitingRulesRequest
  GetRuleNameList() *string 
}

type EnableSQLRateLimitingRulesRequest struct {
  // The cluster ID. Notes on invoking: The SQL throttling feature is supported only for clusters that have the database proxy enabled. Currently, the Oracle-compatible engine is verified to support this feature. For MySQL clusters, confirm whether the database proxy and the SQL throttling feature are enabled.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // pc-*****************
  DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
  // Specifies whether to enable or disable the specified SQL throttling rules. Valid values:
  // 
  // - **true**: enables the rules.
  // 
  // - **false**: disables the rules.
  // 
  // > This parameter takes effect only when **RuleNameList*	- is specified.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // true
  Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The names of the SQL throttling rules to enable or disable. You can specify multiple rule names for batch operations. Separate multiple names with commas (,).
  // 
  // > You can view all SQL throttling rule information, including rule names, on the **SQL Firewall*	- tab of the **Security Management*	- page in the cluster console.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // testrule
  RuleNameList *string `json:"RuleNameList,omitempty" xml:"RuleNameList,omitempty"`
}

func (s EnableSQLRateLimitingRulesRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSQLRateLimitingRulesRequest) GoString() string {
  return s.String()
}

func (s *EnableSQLRateLimitingRulesRequest) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *EnableSQLRateLimitingRulesRequest) GetEnable() *bool  {
  return s.Enable
}

func (s *EnableSQLRateLimitingRulesRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableSQLRateLimitingRulesRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableSQLRateLimitingRulesRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableSQLRateLimitingRulesRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableSQLRateLimitingRulesRequest) GetRuleNameList() *string  {
  return s.RuleNameList
}

func (s *EnableSQLRateLimitingRulesRequest) SetDBClusterId(v string) *EnableSQLRateLimitingRulesRequest {
  s.DBClusterId = &v
  return s
}

func (s *EnableSQLRateLimitingRulesRequest) SetEnable(v bool) *EnableSQLRateLimitingRulesRequest {
  s.Enable = &v
  return s
}

func (s *EnableSQLRateLimitingRulesRequest) SetOwnerAccount(v string) *EnableSQLRateLimitingRulesRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableSQLRateLimitingRulesRequest) SetOwnerId(v int64) *EnableSQLRateLimitingRulesRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableSQLRateLimitingRulesRequest) SetResourceOwnerAccount(v string) *EnableSQLRateLimitingRulesRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableSQLRateLimitingRulesRequest) SetResourceOwnerId(v int64) *EnableSQLRateLimitingRulesRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableSQLRateLimitingRulesRequest) SetRuleNameList(v string) *EnableSQLRateLimitingRulesRequest {
  s.RuleNameList = &v
  return s
}

func (s *EnableSQLRateLimitingRulesRequest) Validate() error {
  return dara.Validate(s)
}

