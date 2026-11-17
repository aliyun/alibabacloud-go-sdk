// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationCustomRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCustomRules(v []*ListApplicationCustomRulesResponseBodyApplicationCustomRules) *ListApplicationCustomRulesResponseBody
	GetApplicationCustomRules() []*ListApplicationCustomRulesResponseBodyApplicationCustomRules
	SetPageNo(v int64) *ListApplicationCustomRulesResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *ListApplicationCustomRulesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListApplicationCustomRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationCustomRulesResponseBody
	GetTotalCount() *int64
}

type ListApplicationCustomRulesResponseBody struct {
	// The list of application custom rules, sorted by priority in ascending order.
	ApplicationCustomRules []*ListApplicationCustomRulesResponseBodyApplicationCustomRules `json:"ApplicationCustomRules,omitempty" xml:"ApplicationCustomRules,omitempty" type:"Repeated"`
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
	// The total number of application custom rules that meet the filter conditions.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationCustomRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationCustomRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationCustomRulesResponseBody) GetApplicationCustomRules() []*ListApplicationCustomRulesResponseBodyApplicationCustomRules {
	return s.ApplicationCustomRules
}

func (s *ListApplicationCustomRulesResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListApplicationCustomRulesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApplicationCustomRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationCustomRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationCustomRulesResponseBody) SetApplicationCustomRules(v []*ListApplicationCustomRulesResponseBodyApplicationCustomRules) *ListApplicationCustomRulesResponseBody {
	s.ApplicationCustomRules = v
	return s
}

func (s *ListApplicationCustomRulesResponseBody) SetPageNo(v int64) *ListApplicationCustomRulesResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListApplicationCustomRulesResponseBody) SetPageSize(v int64) *ListApplicationCustomRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationCustomRulesResponseBody) SetRequestId(v string) *ListApplicationCustomRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationCustomRulesResponseBody) SetTotalCount(v int64) *ListApplicationCustomRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationCustomRulesResponseBody) Validate() error {
	if s.ApplicationCustomRules != nil {
		for _, item := range s.ApplicationCustomRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationCustomRulesResponseBodyApplicationCustomRules struct {
	// The ID of the application custom rule.
	//
	// example:
	//
	// acr-7d2e4f6a8b0c1d3e
	ApplicationCustomRuleId *string `json:"ApplicationCustomRuleId,omitempty" xml:"ApplicationCustomRuleId,omitempty"`
	// The name of the custom rule.
	//
	// example:
	//
	// Custom rule 01
	ApplicationCustomRuleName *string `json:"ApplicationCustomRuleName,omitempty" xml:"ApplicationCustomRuleName,omitempty"`
	// The ID of the application template to which the rule belongs.
	//
	// example:
	//
	// ati-4f0691c56b8d4f06
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The list of applications associated with the rule.
	Applications []*ListApplicationCustomRulesResponseBodyApplicationCustomRulesApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The priority of the rule. Valid values: integers from 1 to 200. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The action performed after the rule is matched. Valid values:
	//
	// - **Allow**: allows the traffic.
	//
	// - **Alert**: monitors the traffic.
	//
	// - **Deny**: blocks the traffic.
	//
	// example:
	//
	// Deny
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
}

func (s ListApplicationCustomRulesResponseBodyApplicationCustomRules) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationCustomRulesResponseBodyApplicationCustomRules) GoString() string {
	return s.String()
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRules) GetApplicationCustomRuleId() *string {
	return s.ApplicationCustomRuleId
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRules) GetApplicationCustomRuleName() *string {
	return s.ApplicationCustomRuleName
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRules) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRules) GetApplications() []*ListApplicationCustomRulesResponseBodyApplicationCustomRulesApplications {
	return s.Applications
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRules) GetPriority() *int64 {
	return s.Priority
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRules) GetRuleAction() *string {
	return s.RuleAction
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRules) SetApplicationCustomRuleId(v string) *ListApplicationCustomRulesResponseBodyApplicationCustomRules {
	s.ApplicationCustomRuleId = &v
	return s
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRules) SetApplicationCustomRuleName(v string) *ListApplicationCustomRulesResponseBodyApplicationCustomRules {
	s.ApplicationCustomRuleName = &v
	return s
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRules) SetApplicationTemplateId(v string) *ListApplicationCustomRulesResponseBodyApplicationCustomRules {
	s.ApplicationTemplateId = &v
	return s
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRules) SetApplications(v []*ListApplicationCustomRulesResponseBodyApplicationCustomRulesApplications) *ListApplicationCustomRulesResponseBodyApplicationCustomRules {
	s.Applications = v
	return s
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRules) SetPriority(v int64) *ListApplicationCustomRulesResponseBodyApplicationCustomRules {
	s.Priority = &v
	return s
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRules) SetRuleAction(v string) *ListApplicationCustomRulesResponseBodyApplicationCustomRules {
	s.RuleAction = &v
	return s
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRules) Validate() error {
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationCustomRulesResponseBodyApplicationCustomRulesApplications struct {
	// The ID of the application in the built-in application signature library.
	//
	// example:
	//
	// 10012150000
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The name of the application. The name is returned in Chinese or English based on the Lang parameter.
	//
	// example:
	//
	// DingTalk
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
}

func (s ListApplicationCustomRulesResponseBodyApplicationCustomRulesApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationCustomRulesResponseBodyApplicationCustomRulesApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRulesApplications) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRulesApplications) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRulesApplications) SetApplicationId(v int64) *ListApplicationCustomRulesResponseBodyApplicationCustomRulesApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRulesApplications) SetApplicationName(v string) *ListApplicationCustomRulesResponseBodyApplicationCustomRulesApplications {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationCustomRulesResponseBodyApplicationCustomRulesApplications) Validate() error {
	return dara.Validate(s)
}
