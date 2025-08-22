// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateDcdnWafRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *BatchCreateDcdnWafRulesRequest
	GetPolicyId() *int64
	SetRuleConfigs(v string) *BatchCreateDcdnWafRulesRequest
	GetRuleConfigs() *string
}

type BatchCreateDcdnWafRulesRequest struct {
	// The ID of the protection policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The configuration of the protection rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"name":"ttttt","action":"monitor","conditions":[{"key":"URL","opValue":"match-one","values":"1,2,3,4,5"},{"key":"Header","opValue":"contain-one","subKey":"testheader","values":"6,7,8,9,10"}],"ratelimit":{"target":"header","interval":10,"threshold":5,"ttl":1800,"subKey":"testheadercc","status":{"code":"502","count":5}},"ccStatus":"on","effect":"rule","status":"on"}
	RuleConfigs *string `json:"RuleConfigs,omitempty" xml:"RuleConfigs,omitempty"`
}

func (s BatchCreateDcdnWafRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateDcdnWafRulesRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateDcdnWafRulesRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *BatchCreateDcdnWafRulesRequest) GetRuleConfigs() *string {
	return s.RuleConfigs
}

func (s *BatchCreateDcdnWafRulesRequest) SetPolicyId(v int64) *BatchCreateDcdnWafRulesRequest {
	s.PolicyId = &v
	return s
}

func (s *BatchCreateDcdnWafRulesRequest) SetRuleConfigs(v string) *BatchCreateDcdnWafRulesRequest {
	s.RuleConfigs = &v
	return s
}

func (s *BatchCreateDcdnWafRulesRequest) Validate() error {
	return dara.Validate(s)
}
