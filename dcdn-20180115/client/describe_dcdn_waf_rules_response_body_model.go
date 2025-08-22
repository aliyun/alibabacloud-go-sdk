// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDcdnWafRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnWafRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDcdnWafRulesResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeDcdnWafRulesResponseBodyRules) *DescribeDcdnWafRulesResponseBody
	GetRules() []*DescribeDcdnWafRulesResponseBodyRules
	SetTotalCount(v int32) *DescribeDcdnWafRulesResponseBody
	GetTotalCount() *int32
}

type DescribeDcdnWafRulesResponseBody struct {
	// The page number of the returned page. The value of this parameter is the same as that of the PageNumber parameter in the request.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of protection rules returned per page. The value of this parameter is the same as that of the PageSize parameter in the request.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 66A98669-CC6E-4F3E-80A6-3014697B11AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the protection rule.
	Rules []*DescribeDcdnWafRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of protection rules.
	//
	// example:
	//
	// 121
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnWafRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnWafRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnWafRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafRulesResponseBody) GetRules() []*DescribeDcdnWafRulesResponseBodyRules {
	return s.Rules
}

func (s *DescribeDcdnWafRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDcdnWafRulesResponseBody) SetPageNumber(v int32) *DescribeDcdnWafRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnWafRulesResponseBody) SetPageSize(v int32) *DescribeDcdnWafRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafRulesResponseBody) SetRequestId(v string) *DescribeDcdnWafRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafRulesResponseBody) SetRules(v []*DescribeDcdnWafRulesResponseBodyRules) *DescribeDcdnWafRulesResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeDcdnWafRulesResponseBody) SetTotalCount(v int32) *DescribeDcdnWafRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnWafRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafRulesResponseBodyRules struct {
	// The type of the protection policy. The value of this parameter is the same as that of the DefenseScene field in QueryArgst.
	//
	// example:
	//
	// custom_acl
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The time when the protection policy was last modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-29T17:08:45Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the protection policy.
	//
	// example:
	//
	// 200001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The configuration information about the protection rule.
	//
	// example:
	//
	// {\\"action\\":\\"monitor\\",\\"actionExternal\\":\\"{}\\",\\"ccStatus\\":1,\\"conditions\\":[{\\"key\\":\\"URL\\",\\"opValue\\":\\"eq\\",\\"targetKey\\":\\"request_uri\\",\\"values\\":\\"/example\\"},{\\"key\\":\\"Header\\",\\"opValue\\":\\"eq\\",\\"subKey\\":\\"trt\\",\\"targetKey\\":\\"header.trt\\",\\"values\\":\\"3333\\"}],\\"effect\\":\\"service\\",\\"name\\":\\"aaa333\\",\\"origin\\":\\"custom\\",\\"ratelimit\\":{\\"interval\\":5,\\"status\\":{\\"code\\":404,\\"count\\":2,\\"stat\\":{\\"mode\\":\\"count\\",\\"value\\":2.0}},\\"target\\":\\"remote_addr\\",\\"threshold\\":2,\\"ttl\\":1800}}
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// The ID of the protection rule.
	//
	// example:
	//
	// 100001
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the protection rule.
	//
	// example:
	//
	// rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the protection rule. The value of this parameter is the same as that of the RuleStatus field in QueryArgst.
	//
	// example:
	//
	// on
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
}

func (s DescribeDcdnWafRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafRulesResponseBodyRules) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDcdnWafRulesResponseBodyRules) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDcdnWafRulesResponseBodyRules) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeDcdnWafRulesResponseBodyRules) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *DescribeDcdnWafRulesResponseBodyRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeDcdnWafRulesResponseBodyRules) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeDcdnWafRulesResponseBodyRules) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeDcdnWafRulesResponseBodyRules) SetDefenseScene(v string) *DescribeDcdnWafRulesResponseBodyRules {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDcdnWafRulesResponseBodyRules) SetGmtModified(v string) *DescribeDcdnWafRulesResponseBodyRules {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnWafRulesResponseBodyRules) SetPolicyId(v int64) *DescribeDcdnWafRulesResponseBodyRules {
	s.PolicyId = &v
	return s
}

func (s *DescribeDcdnWafRulesResponseBodyRules) SetRuleConfig(v string) *DescribeDcdnWafRulesResponseBodyRules {
	s.RuleConfig = &v
	return s
}

func (s *DescribeDcdnWafRulesResponseBodyRules) SetRuleId(v int64) *DescribeDcdnWafRulesResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeDcdnWafRulesResponseBodyRules) SetRuleName(v string) *DescribeDcdnWafRulesResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *DescribeDcdnWafRulesResponseBodyRules) SetRuleStatus(v string) *DescribeDcdnWafRulesResponseBodyRules {
	s.RuleStatus = &v
	return s
}

func (s *DescribeDcdnWafRulesResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
