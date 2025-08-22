// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDcdnWafPoliciesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnWafPoliciesResponseBody
	GetPageSize() *int32
	SetPolicies(v []*DescribeDcdnWafPoliciesResponseBodyPolicies) *DescribeDcdnWafPoliciesResponseBody
	GetPolicies() []*DescribeDcdnWafPoliciesResponseBodyPolicies
	SetRequestId(v string) *DescribeDcdnWafPoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDcdnWafPoliciesResponseBody
	GetTotalCount() *int32
}

type DescribeDcdnWafPoliciesResponseBody struct {
	// The page number of the returned page. Valid values: **1*	- to **100000**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of protection policies returned per page. Valid values: an integer from **1*	- to **500**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about protection policies.
	Policies []*DescribeDcdnWafPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 153ca2cd-3c01-44be-2e83-64dbc6c88630
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of protection policies.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnWafPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPoliciesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnWafPoliciesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnWafPoliciesResponseBody) GetPolicies() []*DescribeDcdnWafPoliciesResponseBodyPolicies {
	return s.Policies
}

func (s *DescribeDcdnWafPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafPoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDcdnWafPoliciesResponseBody) SetPageNumber(v int32) *DescribeDcdnWafPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnWafPoliciesResponseBody) SetPageSize(v int32) *DescribeDcdnWafPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafPoliciesResponseBody) SetPolicies(v []*DescribeDcdnWafPoliciesResponseBodyPolicies) *DescribeDcdnWafPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *DescribeDcdnWafPoliciesResponseBody) SetRequestId(v string) *DescribeDcdnWafPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafPoliciesResponseBody) SetTotalCount(v int32) *DescribeDcdnWafPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnWafPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafPoliciesResponseBodyPolicies struct {
	// The type of the protection policy, which is the same as the DefenseScenes field in the QueryArgs parameter.
	//
	// example:
	//
	// custom_acl
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
	// The status of the protection policy, which is the same as the PolicyStatus field in the QueryArgs parameter.
	//
	// example:
	//
	// on
	PolicyStatus *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	// Indicates whether this protection policy is the default policy, which is the same as the PolicyType field in the QueryArgs parameter.
	//
	// example:
	//
	// default
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The number of protection rules in the protection policy.
	//
	// example:
	//
	// 9
	RuleCount *int64 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
}

func (s DescribeDcdnWafPoliciesResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) GetDomainCount() *int32 {
	return s.DomainCount
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) GetPolicyStatus() *string {
	return s.PolicyStatus
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) GetRuleCount() *int64 {
	return s.RuleCount
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) SetDefenseScene(v string) *DescribeDcdnWafPoliciesResponseBodyPolicies {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) SetDomainCount(v int32) *DescribeDcdnWafPoliciesResponseBodyPolicies {
	s.DomainCount = &v
	return s
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) SetGmtModified(v string) *DescribeDcdnWafPoliciesResponseBodyPolicies {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) SetPolicyId(v int64) *DescribeDcdnWafPoliciesResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) SetPolicyName(v string) *DescribeDcdnWafPoliciesResponseBodyPolicies {
	s.PolicyName = &v
	return s
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) SetPolicyStatus(v string) *DescribeDcdnWafPoliciesResponseBodyPolicies {
	s.PolicyStatus = &v
	return s
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) SetPolicyType(v string) *DescribeDcdnWafPoliciesResponseBodyPolicies {
	s.PolicyType = &v
	return s
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) SetRuleCount(v int64) *DescribeDcdnWafPoliciesResponseBodyPolicies {
	s.RuleCount = &v
	return s
}

func (s *DescribeDcdnWafPoliciesResponseBodyPolicies) Validate() error {
	return dara.Validate(s)
}
