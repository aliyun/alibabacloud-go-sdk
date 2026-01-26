// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSilencePoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *ListSilencePoliciesResponseBodyPageBean) *ListSilencePoliciesResponseBody
	GetPageBean() *ListSilencePoliciesResponseBodyPageBean
	SetRequestId(v string) *ListSilencePoliciesResponseBody
	GetRequestId() *string
}

type ListSilencePoliciesResponseBody struct {
	// The returned pages.
	PageBean *ListSilencePoliciesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 78901766-3806-4E96-8E47-CFEF59E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSilencePoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSilencePoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSilencePoliciesResponseBody) GetPageBean() *ListSilencePoliciesResponseBodyPageBean {
	return s.PageBean
}

func (s *ListSilencePoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSilencePoliciesResponseBody) SetPageBean(v *ListSilencePoliciesResponseBodyPageBean) *ListSilencePoliciesResponseBody {
	s.PageBean = v
	return s
}

func (s *ListSilencePoliciesResponseBody) SetRequestId(v string) *ListSilencePoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSilencePoliciesResponseBody) Validate() error {
	if s.PageBean != nil {
		if err := s.PageBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSilencePoliciesResponseBodyPageBean struct {
	// The number of the page returned.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The queried silence policies.
	SilencePolicies []*ListSilencePoliciesResponseBodyPageBeanSilencePolicies `json:"SilencePolicies,omitempty" xml:"SilencePolicies,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The number of silence policies that were returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSilencePoliciesResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s ListSilencePoliciesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *ListSilencePoliciesResponseBodyPageBean) GetPage() *int64 {
	return s.Page
}

func (s *ListSilencePoliciesResponseBodyPageBean) GetSilencePolicies() []*ListSilencePoliciesResponseBodyPageBeanSilencePolicies {
	return s.SilencePolicies
}

func (s *ListSilencePoliciesResponseBodyPageBean) GetSize() *int64 {
	return s.Size
}

func (s *ListSilencePoliciesResponseBodyPageBean) GetTotal() *int64 {
	return s.Total
}

func (s *ListSilencePoliciesResponseBodyPageBean) SetPage(v int64) *ListSilencePoliciesResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBean) SetSilencePolicies(v []*ListSilencePoliciesResponseBodyPageBeanSilencePolicies) *ListSilencePoliciesResponseBodyPageBean {
	s.SilencePolicies = v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBean) SetSize(v int64) *ListSilencePoliciesResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBean) SetTotal(v int64) *ListSilencePoliciesResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBean) Validate() error {
	if s.SilencePolicies != nil {
		for _, item := range s.SilencePolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSilencePoliciesResponseBodyPageBeanSilencePolicies struct {
	// The effective type. Valid values: PERMANENT: The policy is effective permanently. CYCLE_EFFECT: The policy is effective cyclically. CUSTOM_TIME: The policy is effective during a custom time period.
	//
	// example:
	//
	// PERMANENT
	EffectiveTimeType *string `json:"EffectiveTimeType,omitempty" xml:"EffectiveTimeType,omitempty"`
	// The ID of the silence policy.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The matching rules.
	MatchingRules []*ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules `json:"MatchingRules,omitempty" xml:"MatchingRules,omitempty" type:"Repeated"`
	// The name of the silence policy.
	//
	// example:
	//
	// silencepolicy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the silence policy is enabled. Valid values: enable and disable.
	//
	// example:
	//
	// enable
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The effective time. Valid values: WEEK: weekly DAY: daily
	//
	// example:
	//
	// WEEK
	TimePeriod *string `json:"TimePeriod,omitempty" xml:"TimePeriod,omitempty"`
	// Effective period.
	//
	// example:
	//
	// [{"endTime":"00:00","endWeek":"7","startTime":"00:00","startWeek":"1"}]
	TimeSlots *string `json:"TimeSlots,omitempty" xml:"TimeSlots,omitempty"`
}

func (s ListSilencePoliciesResponseBodyPageBeanSilencePolicies) String() string {
	return dara.Prettify(s)
}

func (s ListSilencePoliciesResponseBodyPageBeanSilencePolicies) GoString() string {
	return s.String()
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) GetEffectiveTimeType() *string {
	return s.EffectiveTimeType
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) GetId() *int64 {
	return s.Id
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) GetMatchingRules() []*ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules {
	return s.MatchingRules
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) GetName() *string {
	return s.Name
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) GetState() *string {
	return s.State
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) GetTimePeriod() *string {
	return s.TimePeriod
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) GetTimeSlots() *string {
	return s.TimeSlots
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) SetEffectiveTimeType(v string) *ListSilencePoliciesResponseBodyPageBeanSilencePolicies {
	s.EffectiveTimeType = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) SetId(v int64) *ListSilencePoliciesResponseBodyPageBeanSilencePolicies {
	s.Id = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) SetMatchingRules(v []*ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules) *ListSilencePoliciesResponseBodyPageBeanSilencePolicies {
	s.MatchingRules = v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) SetName(v string) *ListSilencePoliciesResponseBodyPageBeanSilencePolicies {
	s.Name = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) SetState(v string) *ListSilencePoliciesResponseBodyPageBeanSilencePolicies {
	s.State = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) SetTimePeriod(v string) *ListSilencePoliciesResponseBodyPageBeanSilencePolicies {
	s.TimePeriod = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) SetTimeSlots(v string) *ListSilencePoliciesResponseBodyPageBeanSilencePolicies {
	s.TimeSlots = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) Validate() error {
	if s.MatchingRules != nil {
		for _, item := range s.MatchingRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules struct {
	// The matching conditions.
	MatchingConditions []*ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions `json:"MatchingConditions,omitempty" xml:"MatchingConditions,omitempty" type:"Repeated"`
}

func (s ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules) String() string {
	return dara.Prettify(s)
}

func (s ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules) GoString() string {
	return s.String()
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules) GetMatchingConditions() []*ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions {
	return s.MatchingConditions
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules) SetMatchingConditions(v []*ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules {
	s.MatchingConditions = v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules) Validate() error {
	if s.MatchingConditions != nil {
		for _, item := range s.MatchingConditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions struct {
	// The key of the matching condition.
	//
	// example:
	//
	// altertname
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The logical operator of the matching condition. Valid values:
	//
	// 	- `eq`: equal to
	//
	// 	- `neq`: not equal to
	//
	// 	- `in`: contains
	//
	// 	- `nin`: does not contain
	//
	// 	- `re`: regular expression match
	//
	// 	- `nre`: regular expression mismatch
	//
	// example:
	//
	// eq
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The value of the matching condition.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) String() string {
	return dara.Prettify(s)
}

func (s ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) GoString() string {
	return s.String()
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) GetKey() *string {
	return s.Key
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) GetOperator() *string {
	return s.Operator
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) GetValue() *string {
	return s.Value
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) SetKey(v string) *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions {
	s.Key = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) SetOperator(v string) *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions {
	s.Operator = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) SetValue(v string) *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions {
	s.Value = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) Validate() error {
	return dara.Validate(s)
}
