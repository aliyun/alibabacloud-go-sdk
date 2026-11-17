// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebFilterRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int64) *ListWebFilterRulesResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *ListWebFilterRulesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListWebFilterRulesResponseBody
	GetRequestId() *string
	SetRules(v []*ListWebFilterRulesResponseBodyRules) *ListWebFilterRulesResponseBody
	GetRules() []*ListWebFilterRulesResponseBodyRules
	SetTotalCount(v int64) *ListWebFilterRulesResponseBody
	GetTotalCount() *int64
}

type ListWebFilterRulesResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of web filtering rules.
	Rules []*ListWebFilterRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of web filtering rules that meet the filter conditions.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWebFilterRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWebFilterRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWebFilterRulesResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListWebFilterRulesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWebFilterRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWebFilterRulesResponseBody) GetRules() []*ListWebFilterRulesResponseBodyRules {
	return s.Rules
}

func (s *ListWebFilterRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWebFilterRulesResponseBody) SetPageNo(v int64) *ListWebFilterRulesResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListWebFilterRulesResponseBody) SetPageSize(v int64) *ListWebFilterRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListWebFilterRulesResponseBody) SetRequestId(v string) *ListWebFilterRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWebFilterRulesResponseBody) SetRules(v []*ListWebFilterRulesResponseBodyRules) *ListWebFilterRulesResponseBody {
	s.Rules = v
	return s
}

func (s *ListWebFilterRulesResponseBody) SetTotalCount(v int64) *ListWebFilterRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWebFilterRulesResponseBody) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWebFilterRulesResponseBodyRules struct {
	// The list of match conditions for the rule. The list contains only one condition.
	Conditions []*ListWebFilterRulesResponseBodyRulesConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The time when the web filtering rule was created. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1731465600
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The action performed after the rule is matched. Valid values:
	//
	// - **Allow**: Allows the request.
	//
	// - **Alert**: Monitors the request.
	//
	// - **Deny**: Blocks the request.
	//
	// example:
	//
	// Deny
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The time when the web filtering rule was last modified. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1731469200
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the web filtering rule.
	//
	// example:
	//
	// wfr-8c1d2e3f4a5b6c7d
	WebFilterRuleId *string `json:"WebFilterRuleId,omitempty" xml:"WebFilterRuleId,omitempty"`
	// The ID of the web filtering template to which the rule belongs.
	//
	// example:
	//
	// wft-3f2a9c1d4b5e6f70
	WebFilterTemplateId *string `json:"WebFilterTemplateId,omitempty" xml:"WebFilterTemplateId,omitempty"`
}

func (s ListWebFilterRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s ListWebFilterRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListWebFilterRulesResponseBodyRules) GetConditions() []*ListWebFilterRulesResponseBodyRulesConditions {
	return s.Conditions
}

func (s *ListWebFilterRulesResponseBodyRules) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListWebFilterRulesResponseBodyRules) GetRuleAction() *string {
	return s.RuleAction
}

func (s *ListWebFilterRulesResponseBodyRules) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListWebFilterRulesResponseBodyRules) GetWebFilterRuleId() *string {
	return s.WebFilterRuleId
}

func (s *ListWebFilterRulesResponseBodyRules) GetWebFilterTemplateId() *string {
	return s.WebFilterTemplateId
}

func (s *ListWebFilterRulesResponseBodyRules) SetConditions(v []*ListWebFilterRulesResponseBodyRulesConditions) *ListWebFilterRulesResponseBodyRules {
	s.Conditions = v
	return s
}

func (s *ListWebFilterRulesResponseBodyRules) SetCreateTime(v int64) *ListWebFilterRulesResponseBodyRules {
	s.CreateTime = &v
	return s
}

func (s *ListWebFilterRulesResponseBodyRules) SetRuleAction(v string) *ListWebFilterRulesResponseBodyRules {
	s.RuleAction = &v
	return s
}

func (s *ListWebFilterRulesResponseBodyRules) SetUpdateTime(v int64) *ListWebFilterRulesResponseBodyRules {
	s.UpdateTime = &v
	return s
}

func (s *ListWebFilterRulesResponseBodyRules) SetWebFilterRuleId(v string) *ListWebFilterRulesResponseBodyRules {
	s.WebFilterRuleId = &v
	return s
}

func (s *ListWebFilterRulesResponseBodyRules) SetWebFilterTemplateId(v string) *ListWebFilterRulesResponseBodyRules {
	s.WebFilterTemplateId = &v
	return s
}

func (s *ListWebFilterRulesResponseBodyRules) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWebFilterRulesResponseBodyRulesConditions struct {
	// The type of the match condition. Valid values: **URL**, which indicates matching by URL.
	//
	// example:
	//
	// URL
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The URL to match. A single asterisk (*) matches all URLs. A value ending with /	- indicates prefix wildcard matching.
	//
	// example:
	//
	// www.example.com/login
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s ListWebFilterRulesResponseBodyRulesConditions) String() string {
	return dara.Prettify(s)
}

func (s ListWebFilterRulesResponseBodyRulesConditions) GoString() string {
	return s.String()
}

func (s *ListWebFilterRulesResponseBodyRulesConditions) GetKey() *string {
	return s.Key
}

func (s *ListWebFilterRulesResponseBodyRulesConditions) GetValues() *string {
	return s.Values
}

func (s *ListWebFilterRulesResponseBodyRulesConditions) SetKey(v string) *ListWebFilterRulesResponseBodyRulesConditions {
	s.Key = &v
	return s
}

func (s *ListWebFilterRulesResponseBodyRulesConditions) SetValues(v string) *ListWebFilterRulesResponseBodyRulesConditions {
	s.Values = &v
	return s
}

func (s *ListWebFilterRulesResponseBodyRulesConditions) Validate() error {
	return dara.Validate(s)
}
