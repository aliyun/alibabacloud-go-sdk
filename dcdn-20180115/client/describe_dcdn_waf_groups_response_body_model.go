// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDcdnWafGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnWafGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDcdnWafGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDcdnWafGroupsResponseBody
	GetTotalCount() *int32
	SetWafGroups(v []*DescribeDcdnWafGroupsResponseBodyWafGroups) *DescribeDcdnWafGroupsResponseBody
	GetWafGroups() []*DescribeDcdnWafGroupsResponseBodyWafGroups
}

type DescribeDcdnWafGroupsResponseBody struct {
	// The page number of the returned page. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of WAF rule groups.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of WAF rule groups.
	WafGroups []*DescribeDcdnWafGroupsResponseBodyWafGroups `json:"WafGroups,omitempty" xml:"WafGroups,omitempty" type:"Repeated"`
}

func (s DescribeDcdnWafGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnWafGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnWafGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDcdnWafGroupsResponseBody) GetWafGroups() []*DescribeDcdnWafGroupsResponseBodyWafGroups {
	return s.WafGroups
}

func (s *DescribeDcdnWafGroupsResponseBody) SetPageNumber(v int32) *DescribeDcdnWafGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnWafGroupsResponseBody) SetPageSize(v int32) *DescribeDcdnWafGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafGroupsResponseBody) SetRequestId(v string) *DescribeDcdnWafGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafGroupsResponseBody) SetTotalCount(v int32) *DescribeDcdnWafGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnWafGroupsResponseBody) SetWafGroups(v []*DescribeDcdnWafGroupsResponseBodyWafGroups) *DescribeDcdnWafGroupsResponseBody {
	s.WafGroups = v
	return s
}

func (s *DescribeDcdnWafGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafGroupsResponseBodyWafGroups struct {
	// The time when the WAF rule group was modified.
	//
	// example:
	//
	// 2020-04-12 22:25:11
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the custom WAF rule group.
	//
	// example:
	//
	// 30000156
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the WAF rule.
	//
	// example:
	//
	// DCDN-test-operation-reports-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The policy that is associated with the WAF rule group.
	Policies []*DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The number of WAF rules.
	//
	// example:
	//
	// 452
	RuleCount *int32 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	// Indicates whether to enable subscription. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	Subscribe *string `json:"Subscribe,omitempty" xml:"Subscribe,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// 1012
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDcdnWafGroupsResponseBodyWafGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafGroupsResponseBodyWafGroups) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroups) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroups) GetId() *int64 {
	return s.Id
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroups) GetName() *string {
	return s.Name
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroups) GetPolicies() []*DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies {
	return s.Policies
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroups) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroups) GetSubscribe() *string {
	return s.Subscribe
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroups) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroups) SetGmtModified(v string) *DescribeDcdnWafGroupsResponseBodyWafGroups {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroups) SetId(v int64) *DescribeDcdnWafGroupsResponseBodyWafGroups {
	s.Id = &v
	return s
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroups) SetName(v string) *DescribeDcdnWafGroupsResponseBodyWafGroups {
	s.Name = &v
	return s
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroups) SetPolicies(v []*DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies) *DescribeDcdnWafGroupsResponseBodyWafGroups {
	s.Policies = v
	return s
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroups) SetRuleCount(v int32) *DescribeDcdnWafGroupsResponseBodyWafGroups {
	s.RuleCount = &v
	return s
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroups) SetSubscribe(v string) *DescribeDcdnWafGroupsResponseBodyWafGroups {
	s.Subscribe = &v
	return s
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroups) SetTemplateId(v int64) *DescribeDcdnWafGroupsResponseBodyWafGroups {
	s.TemplateId = &v
	return s
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies struct {
	// The ID of the policy.
	//
	// example:
	//
	// 30000165
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// wasm-testmaster
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **custom**: a custom policy
	//
	// 	- **default**: the default policy
	//
	// example:
	//
	// default
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies) GetId() *int64 {
	return s.Id
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies) GetName() *string {
	return s.Name
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies) GetType() *string {
	return s.Type
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies) SetId(v int64) *DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies {
	s.Id = &v
	return s
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies) SetName(v string) *DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies {
	s.Name = &v
	return s
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies) SetType(v string) *DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies {
	s.Type = &v
	return s
}

func (s *DescribeDcdnWafGroupsResponseBodyWafGroupsPolicies) Validate() error {
	return dara.Validate(s)
}
