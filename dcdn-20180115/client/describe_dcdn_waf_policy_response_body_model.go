// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *DescribeDcdnWafPolicyResponseBodyPolicy) *DescribeDcdnWafPolicyResponseBody
	GetPolicy() *DescribeDcdnWafPolicyResponseBodyPolicy
	SetRequestId(v string) *DescribeDcdnWafPolicyResponseBody
	GetRequestId() *string
}

type DescribeDcdnWafPolicyResponseBody struct {
	// The information about the protection policy.
	Policy *DescribeDcdnWafPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 66A98669-CC6E-4F3E-80A6-3014697B11AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnWafPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPolicyResponseBody) GetPolicy() *DescribeDcdnWafPolicyResponseBodyPolicy {
	return s.Policy
}

func (s *DescribeDcdnWafPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafPolicyResponseBody) SetPolicy(v *DescribeDcdnWafPolicyResponseBodyPolicy) *DescribeDcdnWafPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *DescribeDcdnWafPolicyResponseBody) SetRequestId(v string) *DescribeDcdnWafPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafPolicyResponseBodyPolicy struct {
	// The type of the protection policy. Valid values:
	//
	// 	- waf_group: basic web protection
	//
	// 	- custom_acl: custom protection
	//
	// 	- whitelist: whitelist
	//
	// example:
	//
	// waf_group
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The number of domain names that use the protection policy.
	//
	// example:
	//
	// 22
	DomainCount *int32 `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	// The time when the protection policy was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-29T17:08:45Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the protection policy.
	//
	// example:
	//
	// 100001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the protection policy.
	//
	// example:
	//
	// policy_test
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The status of the protection policy. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	PolicyStatus *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	// Indicates whether the current policy is the default policy. Valid values:
	//
	// 	- default
	//
	// 	- custom
	//
	// example:
	//
	// default
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The protection rule configurations corresponding to the protection policy. The configurations only support Bot management. For more information, see [BatchCreateDcdnWafRules](~~BatchCreateDcdnWafRules~~).
	//
	// example:
	//
	// {     "type":"target_type",     "status":"on",     "config":{"target":"app"},     "action":""   }
	RuleConfigs *string `json:"RuleConfigs,omitempty" xml:"RuleConfigs,omitempty"`
	// The number of protection rules in the protection policy.
	//
	// example:
	//
	// 9
	RuleCount *int64 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
}

func (s DescribeDcdnWafPolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) GetDomainCount() *int32 {
	return s.DomainCount
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) GetPolicyStatus() *string {
	return s.PolicyStatus
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) GetRuleConfigs() *string {
	return s.RuleConfigs
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) GetRuleCount() *int64 {
	return s.RuleCount
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) SetDefenseScene(v string) *DescribeDcdnWafPolicyResponseBodyPolicy {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) SetDomainCount(v int32) *DescribeDcdnWafPolicyResponseBodyPolicy {
	s.DomainCount = &v
	return s
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) SetGmtModified(v string) *DescribeDcdnWafPolicyResponseBodyPolicy {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) SetPolicyId(v int64) *DescribeDcdnWafPolicyResponseBodyPolicy {
	s.PolicyId = &v
	return s
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) SetPolicyName(v string) *DescribeDcdnWafPolicyResponseBodyPolicy {
	s.PolicyName = &v
	return s
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) SetPolicyStatus(v string) *DescribeDcdnWafPolicyResponseBodyPolicy {
	s.PolicyStatus = &v
	return s
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) SetPolicyType(v string) *DescribeDcdnWafPolicyResponseBodyPolicy {
	s.PolicyType = &v
	return s
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) SetRuleConfigs(v string) *DescribeDcdnWafPolicyResponseBodyPolicy {
	s.RuleConfigs = &v
	return s
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) SetRuleCount(v int64) *DescribeDcdnWafPolicyResponseBodyPolicy {
	s.RuleCount = &v
	return s
}

func (s *DescribeDcdnWafPolicyResponseBodyPolicy) Validate() error {
	return dara.Validate(s)
}
