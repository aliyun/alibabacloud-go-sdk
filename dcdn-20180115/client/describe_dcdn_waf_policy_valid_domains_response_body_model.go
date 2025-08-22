// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafPolicyValidDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*DescribeDcdnWafPolicyValidDomainsResponseBodyDomains) *DescribeDcdnWafPolicyValidDomainsResponseBody
	GetDomains() []*DescribeDcdnWafPolicyValidDomainsResponseBodyDomains
	SetPageNumber(v int32) *DescribeDcdnWafPolicyValidDomainsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnWafPolicyValidDomainsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDcdnWafPolicyValidDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDcdnWafPolicyValidDomainsResponseBody
	GetTotalCount() *int32
}

type DescribeDcdnWafPolicyValidDomainsResponseBody struct {
	// The information about the protected domain names.
	Domains []*DescribeDcdnWafPolicyValidDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The page number of the returned page, which is the same as the PageNumber parameter in request parameters.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of domain names returned per page, which is the same as the PageSize parameter in request parameters.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3C6CCEC4-6B88-4D4A-93E4-D47B3D92C630
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of domain names returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnWafPolicyValidDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPolicyValidDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBody) GetDomains() []*DescribeDcdnWafPolicyValidDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBody) SetDomains(v []*DescribeDcdnWafPolicyValidDomainsResponseBodyDomains) *DescribeDcdnWafPolicyValidDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBody) SetPageNumber(v int32) *DescribeDcdnWafPolicyValidDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBody) SetPageSize(v int32) *DescribeDcdnWafPolicyValidDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBody) SetRequestId(v string) *DescribeDcdnWafPolicyValidDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBody) SetTotalCount(v int32) *DescribeDcdnWafPolicyValidDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafPolicyValidDomainsResponseBodyDomains struct {
	// The protected domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The policy that is bound to the domain name.
	Policies []*DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The ID of the protection policy.
	//
	// example:
	//
	// 1000001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the protection policy.
	//
	// example:
	//
	// test1
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// Indicates whether the protection policy is the default policy. Valid values:
	//
	// 	- default: The protection policy is the default policy.
	//
	// 	- custom: The protection policy is not the default policy.
	//
	// example:
	//
	// default
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeDcdnWafPolicyValidDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPolicyValidDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains) GetPolicies() []*DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies {
	return s.Policies
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains) SetDomainName(v string) *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains) SetPolicies(v []*DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies) *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains {
	s.Policies = v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains) SetPolicyId(v int64) *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains {
	s.PolicyId = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains) SetPolicyName(v string) *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains {
	s.PolicyName = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains) SetPolicyType(v string) *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains {
	s.PolicyType = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies struct {
	// The ID of the rule.
	//
	// example:
	//
	// 10000002
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// test2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the policy.
	//
	// example:
	//
	// custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies) GetId() *int64 {
	return s.Id
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies) GetName() *string {
	return s.Name
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies) GetType() *string {
	return s.Type
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies) SetId(v int64) *DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies {
	s.Id = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies) SetName(v string) *DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies {
	s.Name = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies) SetType(v string) *DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies {
	s.Type = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponseBodyDomainsPolicies) Validate() error {
	return dara.Validate(s)
}
