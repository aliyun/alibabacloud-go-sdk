// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemClientRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *ListSystemClientRulesResponseBodyPageInfo) *ListSystemClientRulesResponseBody
	GetPageInfo() *ListSystemClientRulesResponseBodyPageInfo
	SetRequestId(v string) *ListSystemClientRulesResponseBody
	GetRequestId() *string
	SetRuleList(v []*ListSystemClientRulesResponseBodyRuleList) *ListSystemClientRulesResponseBody
	GetRuleList() []*ListSystemClientRulesResponseBodyRuleList
}

type ListSystemClientRulesResponseBody struct {
	// The pagination information.
	PageInfo *ListSystemClientRulesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C43CA30F-EF67-51BB-8C95-F31B8303****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the system defense rules.
	RuleList []*ListSystemClientRulesResponseBodyRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s ListSystemClientRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSystemClientRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemClientRulesResponseBody) GetPageInfo() *ListSystemClientRulesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListSystemClientRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSystemClientRulesResponseBody) GetRuleList() []*ListSystemClientRulesResponseBodyRuleList {
	return s.RuleList
}

func (s *ListSystemClientRulesResponseBody) SetPageInfo(v *ListSystemClientRulesResponseBodyPageInfo) *ListSystemClientRulesResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListSystemClientRulesResponseBody) SetRequestId(v string) *ListSystemClientRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemClientRulesResponseBody) SetRuleList(v []*ListSystemClientRulesResponseBodyRuleList) *ListSystemClientRulesResponseBody {
	s.RuleList = v
	return s
}

func (s *ListSystemClientRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSystemClientRulesResponseBodyPageInfo struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 17
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSystemClientRulesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSystemClientRulesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListSystemClientRulesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListSystemClientRulesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSystemClientRulesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSystemClientRulesResponseBodyPageInfo) SetCurrentPage(v int32) *ListSystemClientRulesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListSystemClientRulesResponseBodyPageInfo) SetPageSize(v int32) *ListSystemClientRulesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListSystemClientRulesResponseBodyPageInfo) SetTotalCount(v int32) *ListSystemClientRulesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListSystemClientRulesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListSystemClientRulesResponseBodyRuleList struct {
	// The name of the aggregation type for the system defense rule.
	//
	// example:
	//
	// Initial entry
	AggregationName *string `json:"AggregationName,omitempty" xml:"AggregationName,omitempty"`
	// The description of the system defense rule.
	//
	// example:
	//
	// Supports alerting or blocking of images that have high-risk vulnerabilities\\*\\*\\*\\*
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the OS. Valid values:
	//
	// 	- **windows**: Windows
	//
	// 	- **linux**: Linux
	//
	// 	- **all**: all types
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// An array that consists of policies.
	Policies []*ListSystemClientRulesResponseBodyRuleListPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The ID of the system defense rule.
	//
	// example:
	//
	// 30****
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the system defense rule.
	//
	// example:
	//
	// Rule\\*\\*\\*\\*
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the system defense rule. Valid values:
	//
	// 	- **1**: alihips, process-specific defense
	//
	// 	- **2**: alinet, network-specific defense
	//
	// example:
	//
	// alihips
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The status of the system defense rule. Valid values:
	//
	// 	- **online**: enabled
	//
	// 	- **offline**: disabled
	//
	// example:
	//
	// online
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Whether the current rule switch takes effect. Valid values:
	//
	// 	- **true**: enabled
	//
	// 	- **false**: disabled
	//
	// example:
	//
	// true
	SwitchEnable *bool `json:"SwitchEnable,omitempty" xml:"SwitchEnable,omitempty"`
	// The switch ID of the system defense rule.
	//
	// example:
	//
	// USER-ENABLE-SWITCH-TYPE_****
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
}

func (s ListSystemClientRulesResponseBodyRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListSystemClientRulesResponseBodyRuleList) GoString() string {
	return s.String()
}

func (s *ListSystemClientRulesResponseBodyRuleList) GetAggregationName() *string {
	return s.AggregationName
}

func (s *ListSystemClientRulesResponseBodyRuleList) GetDescription() *string {
	return s.Description
}

func (s *ListSystemClientRulesResponseBodyRuleList) GetPlatform() *string {
	return s.Platform
}

func (s *ListSystemClientRulesResponseBodyRuleList) GetPolicies() []*ListSystemClientRulesResponseBodyRuleListPolicies {
	return s.Policies
}

func (s *ListSystemClientRulesResponseBodyRuleList) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListSystemClientRulesResponseBodyRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSystemClientRulesResponseBodyRuleList) GetRuleType() *int32 {
	return s.RuleType
}

func (s *ListSystemClientRulesResponseBodyRuleList) GetStatus() *int32 {
	return s.Status
}

func (s *ListSystemClientRulesResponseBodyRuleList) GetSwitchEnable() *bool {
	return s.SwitchEnable
}

func (s *ListSystemClientRulesResponseBodyRuleList) GetSwitchId() *string {
	return s.SwitchId
}

func (s *ListSystemClientRulesResponseBodyRuleList) SetAggregationName(v string) *ListSystemClientRulesResponseBodyRuleList {
	s.AggregationName = &v
	return s
}

func (s *ListSystemClientRulesResponseBodyRuleList) SetDescription(v string) *ListSystemClientRulesResponseBodyRuleList {
	s.Description = &v
	return s
}

func (s *ListSystemClientRulesResponseBodyRuleList) SetPlatform(v string) *ListSystemClientRulesResponseBodyRuleList {
	s.Platform = &v
	return s
}

func (s *ListSystemClientRulesResponseBodyRuleList) SetPolicies(v []*ListSystemClientRulesResponseBodyRuleListPolicies) *ListSystemClientRulesResponseBodyRuleList {
	s.Policies = v
	return s
}

func (s *ListSystemClientRulesResponseBodyRuleList) SetRuleId(v int64) *ListSystemClientRulesResponseBodyRuleList {
	s.RuleId = &v
	return s
}

func (s *ListSystemClientRulesResponseBodyRuleList) SetRuleName(v string) *ListSystemClientRulesResponseBodyRuleList {
	s.RuleName = &v
	return s
}

func (s *ListSystemClientRulesResponseBodyRuleList) SetRuleType(v int32) *ListSystemClientRulesResponseBodyRuleList {
	s.RuleType = &v
	return s
}

func (s *ListSystemClientRulesResponseBodyRuleList) SetStatus(v int32) *ListSystemClientRulesResponseBodyRuleList {
	s.Status = &v
	return s
}

func (s *ListSystemClientRulesResponseBodyRuleList) SetSwitchEnable(v bool) *ListSystemClientRulesResponseBodyRuleList {
	s.SwitchEnable = &v
	return s
}

func (s *ListSystemClientRulesResponseBodyRuleList) SetSwitchId(v string) *ListSystemClientRulesResponseBodyRuleList {
	s.SwitchId = &v
	return s
}

func (s *ListSystemClientRulesResponseBodyRuleList) Validate() error {
	return dara.Validate(s)
}

type ListSystemClientRulesResponseBodyRuleListPolicies struct {
	// The policy key.
	//
	// example:
	//
	// system_auto_****
	PolicyKey *string `json:"PolicyKey,omitempty" xml:"PolicyKey,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// System self-starting task\\*\\*\\*\\*
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s ListSystemClientRulesResponseBodyRuleListPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListSystemClientRulesResponseBodyRuleListPolicies) GoString() string {
	return s.String()
}

func (s *ListSystemClientRulesResponseBodyRuleListPolicies) GetPolicyKey() *string {
	return s.PolicyKey
}

func (s *ListSystemClientRulesResponseBodyRuleListPolicies) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListSystemClientRulesResponseBodyRuleListPolicies) SetPolicyKey(v string) *ListSystemClientRulesResponseBodyRuleListPolicies {
	s.PolicyKey = &v
	return s
}

func (s *ListSystemClientRulesResponseBodyRuleListPolicies) SetPolicyName(v string) *ListSystemClientRulesResponseBodyRuleListPolicies {
	s.PolicyName = &v
	return s
}

func (s *ListSystemClientRulesResponseBodyRuleListPolicies) Validate() error {
	return dara.Validate(s)
}
