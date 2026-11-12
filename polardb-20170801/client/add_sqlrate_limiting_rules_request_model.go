// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSQLRateLimitingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *AddSQLRateLimitingRulesRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *AddSQLRateLimitingRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddSQLRateLimitingRulesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddSQLRateLimitingRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddSQLRateLimitingRulesRequest
	GetResourceOwnerId() *int64
	SetRuleConfig(v string) *AddSQLRateLimitingRulesRequest
	GetRuleConfig() *string
	SetRuleName(v string) *AddSQLRateLimitingRulesRequest
	GetRuleName() *string
}

type AddSQLRateLimitingRulesRequest struct {
	// The cluster ID. Usage notes: SQL throttling is supported only for PostgreSQL and Oracle engines, and is not supported for MySQL engines.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The JSON string of the configuration parameters and their values for the SQL throttling rule to be modified. All parameter values are of the string type. Example: `{"id":"test","enabled":"true","match_mode":"0","template":"dXBkYXRlIHQgc2V0IGEgPSAxIHdoZXJlIGlkID0gMQ==","user":"","database":"","waiting":1024,"endpoint":"[{"EndpointName":"pe-***********","EndpointType":"Cluster","DBEndpointDescription":"Cluster Address"}]","throttle_mode":0,"concurrency":1}`. Parameters:
	//
	// 	- `"id"`: Required. The name of the throttling rule. The name must meet the following requirements:
	//
	//   - The name cannot exceed 30 characters in length.
	//
	//   - The name must consist of uppercase and lowercase letters and digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"id":"test","enabled":"true","match_mode":"0","template":"dXBkYXRlIHQgc2V0IGEgPSAxIHdoZXJlIGlkID0gMQ==","user":"","database":"","waiting":1024,"endpoint":"[{"EndpointName":"pe-***********","EndpointType":"Cluster","DBEndpointDescription":"Cluster Address"}]","throttle_mode":0,"concurrency":1}
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// The name of the SQL throttling rule. You can specify only one rule name at a time.
	//
	// > 	- You can call the [DescribeSQLRateLimitingRules](https://help.aliyun.com/document_detail/212573.html) operation to query the details of all SQL throttling rules in the specified cluster, including the rule names.
	//
	// > 	- If the specified rule name does not exist in the current cluster, the system automatically creates an SQL throttling rule based on the name and the `RuleConfig` value.
	//
	// This parameter is required.
	//
	// example:
	//
	// testrule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s AddSQLRateLimitingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSQLRateLimitingRulesRequest) GoString() string {
	return s.String()
}

func (s *AddSQLRateLimitingRulesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AddSQLRateLimitingRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddSQLRateLimitingRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddSQLRateLimitingRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddSQLRateLimitingRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddSQLRateLimitingRulesRequest) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *AddSQLRateLimitingRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *AddSQLRateLimitingRulesRequest) SetDBClusterId(v string) *AddSQLRateLimitingRulesRequest {
	s.DBClusterId = &v
	return s
}

func (s *AddSQLRateLimitingRulesRequest) SetOwnerAccount(v string) *AddSQLRateLimitingRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddSQLRateLimitingRulesRequest) SetOwnerId(v int64) *AddSQLRateLimitingRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *AddSQLRateLimitingRulesRequest) SetResourceOwnerAccount(v string) *AddSQLRateLimitingRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddSQLRateLimitingRulesRequest) SetResourceOwnerId(v int64) *AddSQLRateLimitingRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddSQLRateLimitingRulesRequest) SetRuleConfig(v string) *AddSQLRateLimitingRulesRequest {
	s.RuleConfig = &v
	return s
}

func (s *AddSQLRateLimitingRulesRequest) SetRuleName(v string) *AddSQLRateLimitingRulesRequest {
	s.RuleName = &v
	return s
}

func (s *AddSQLRateLimitingRulesRequest) Validate() error {
	return dara.Validate(s)
}
