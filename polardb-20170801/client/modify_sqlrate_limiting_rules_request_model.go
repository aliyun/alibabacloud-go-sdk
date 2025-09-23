// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySQLRateLimitingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifySQLRateLimitingRulesRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifySQLRateLimitingRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySQLRateLimitingRulesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifySQLRateLimitingRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySQLRateLimitingRulesRequest
	GetResourceOwnerId() *int64
	SetRuleConfig(v string) *ModifySQLRateLimitingRulesRequest
	GetRuleConfig() *string
	SetRuleName(v string) *ModifySQLRateLimitingRulesRequest
	GetRuleName() *string
}

type ModifySQLRateLimitingRulesRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// {"id":"test","enabled":"true","match_mode":"0","template":"dXBkYXRlIHQgc2V0IGEgPSAxIHdoZXJlIGlkID0gMQ==","user":"","database":"","waiting":1024,"endpoint":"[{"EndpointName":"pe-***********","EndpointType":"Cluster","DBEndpointDescription":"Cluster Address"}]","throttle_mode":0,"concurrency":1}
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testrule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ModifySQLRateLimitingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySQLRateLimitingRulesRequest) GoString() string {
	return s.String()
}

func (s *ModifySQLRateLimitingRulesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifySQLRateLimitingRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySQLRateLimitingRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySQLRateLimitingRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySQLRateLimitingRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySQLRateLimitingRulesRequest) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *ModifySQLRateLimitingRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifySQLRateLimitingRulesRequest) SetDBClusterId(v string) *ModifySQLRateLimitingRulesRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifySQLRateLimitingRulesRequest) SetOwnerAccount(v string) *ModifySQLRateLimitingRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySQLRateLimitingRulesRequest) SetOwnerId(v int64) *ModifySQLRateLimitingRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySQLRateLimitingRulesRequest) SetResourceOwnerAccount(v string) *ModifySQLRateLimitingRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySQLRateLimitingRulesRequest) SetResourceOwnerId(v int64) *ModifySQLRateLimitingRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySQLRateLimitingRulesRequest) SetRuleConfig(v string) *ModifySQLRateLimitingRulesRequest {
	s.RuleConfig = &v
	return s
}

func (s *ModifySQLRateLimitingRulesRequest) SetRuleName(v string) *ModifySQLRateLimitingRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ModifySQLRateLimitingRulesRequest) Validate() error {
	return dara.Validate(s)
}
