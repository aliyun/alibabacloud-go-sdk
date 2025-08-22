// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchModifyDcdnWafRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *BatchModifyDcdnWafRulesRequest
	GetPolicyId() *int64
	SetRuleConfigs(v string) *BatchModifyDcdnWafRulesRequest
	GetRuleConfigs() *string
}

type BatchModifyDcdnWafRulesRequest struct {
	// The ID of the protection policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The configurations of the protection rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"id":135,"type":"web_sdk","status":"on","config":{"mode":"automatic","crossDomain":"example.com"},"action":""},{"id":149,"type":"intelligence_fake_crawler","status":"on","config":{},"action":"deny"}]
	RuleConfigs *string `json:"RuleConfigs,omitempty" xml:"RuleConfigs,omitempty"`
}

func (s BatchModifyDcdnWafRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchModifyDcdnWafRulesRequest) GoString() string {
	return s.String()
}

func (s *BatchModifyDcdnWafRulesRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *BatchModifyDcdnWafRulesRequest) GetRuleConfigs() *string {
	return s.RuleConfigs
}

func (s *BatchModifyDcdnWafRulesRequest) SetPolicyId(v int64) *BatchModifyDcdnWafRulesRequest {
	s.PolicyId = &v
	return s
}

func (s *BatchModifyDcdnWafRulesRequest) SetRuleConfigs(v string) *BatchModifyDcdnWafRulesRequest {
	s.RuleConfigs = &v
	return s
}

func (s *BatchModifyDcdnWafRulesRequest) Validate() error {
	return dara.Validate(s)
}
