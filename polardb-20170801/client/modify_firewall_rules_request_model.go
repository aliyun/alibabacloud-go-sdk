// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFirewallRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyFirewallRulesRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyFirewallRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyFirewallRulesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyFirewallRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyFirewallRulesRequest
	GetResourceOwnerId() *int64
	SetRuleConfig(v string) *ModifyFirewallRulesRequest
	GetRuleConfig() *string
	SetRuleName(v string) *ModifyFirewallRulesRequest
	GetRuleName() *string
}

type ModifyFirewallRulesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*****************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// A JSON string that contains the configuration parameters of the firewall rule to modify and their values. The parameter values are strings. Example: `{ "id": "test", "enabled": "true", "mode": "Defending", "users": { "applies_to": [] }, "endpoint": "[{"EndpointName":"pe-***************","EndpointType":"Cluster","DBEndpointDescription":"Cluster Endpoint"},{"EndpointName":"pe-***************","EndpointType":"Custom","DBEndpointDescription":"pc-***************"},{"EndpointName":"pe-***************","EndpointType":"Custom","DBEndpointDescription":"pc-***************"}]", "type": "BlackList", "sub_rules": [] }, "RuleName": "test" }`. The parameters are described as follows:
	//
	// - `"id"`: Required. The name of the firewall rule.
	//
	// - `"databases"`: Optional. The names of the databases to which the rule applies. You can specify multiple database names. Separate the names with commas (,). If you leave this parameter empty, the rule applies to all databases in the cluster.
	//
	// - `"tables"`: Optional. The names of the tables to which the rule applies. You can specify multiple table names. Separate the names with commas (,). If you leave this parameter empty, the rule applies to all tables in the cluster.
	//
	// - `"columns"`: Required. The names of the fields to which the rule applies. You can specify multiple field names. Separate the names with commas (,).
	//
	// - `"description"`: Optional. The description of the data masking rule. The description can be up to 64 characters in length.
	//
	// - `"enabled"`: Required. Specifies whether to enable or disable the data masking rule. Valid values: **true*	- (enable) and **false*	- (disable).
	//
	// - `"applies_to"`: The names of the database accounts to which the rule applies. You can specify multiple database account names. Separate the names with commas (,).
	//
	// - `"exempted"`: The names of the database accounts to which the rule does not apply. You can specify multiple database account names. Separate the names with commas (,).
	//
	// > 	- If you specify the `RuleName` parameter, the `RuleConfig` parameter is required.
	//
	// >
	//
	// > 	- You must specify either `"applies_to"` or `"exempted"`.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"id":"test","enabled":"true","mode":"Collecting","users":{"applies_to":[]},"endpoint":"[{"EndpointName":"pe-************","EndpointType":"Cluster","DBEndpointDescription":"Cluster Address"},{"EndpointName":"pe-************","EndpointType":"Custom","DBEndpointDescription":"pc-************"},{"EndpointName":"pe-************","EndpointType":"Custom","DBEndpointDescription":"pc-************K"}]","type":"WhiteList"}
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// The name of the firewall rule. You can specify only one rule name at a time.
	//
	// > - Call the [DescribeFirewallRules](https://help.aliyun.com/document_detail/212573.html) operation to query the details of all firewall rules for the target cluster, including the rule names.
	//
	// >
	//
	// > - If the specified rule name does not exist in the current cluster, the system automatically creates a new firewall rule based on the rule name and the value of `RuleConfig`.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ModifyFirewallRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirewallRulesRequest) GoString() string {
	return s.String()
}

func (s *ModifyFirewallRulesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyFirewallRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyFirewallRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyFirewallRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyFirewallRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyFirewallRulesRequest) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *ModifyFirewallRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyFirewallRulesRequest) SetDBClusterId(v string) *ModifyFirewallRulesRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyFirewallRulesRequest) SetOwnerAccount(v string) *ModifyFirewallRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyFirewallRulesRequest) SetOwnerId(v int64) *ModifyFirewallRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyFirewallRulesRequest) SetResourceOwnerAccount(v string) *ModifyFirewallRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyFirewallRulesRequest) SetResourceOwnerId(v int64) *ModifyFirewallRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyFirewallRulesRequest) SetRuleConfig(v string) *ModifyFirewallRulesRequest {
	s.RuleConfig = &v
	return s
}

func (s *ModifyFirewallRulesRequest) SetRuleName(v string) *ModifyFirewallRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyFirewallRulesRequest) Validate() error {
	return dara.Validate(s)
}
