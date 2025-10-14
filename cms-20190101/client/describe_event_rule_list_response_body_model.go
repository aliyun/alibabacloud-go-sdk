// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeEventRuleListResponseBody
	GetCode() *string
	SetEventRules(v *DescribeEventRuleListResponseBodyEventRules) *DescribeEventRuleListResponseBody
	GetEventRules() *DescribeEventRuleListResponseBodyEventRules
	SetMessage(v string) *DescribeEventRuleListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEventRuleListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeEventRuleListResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeEventRuleListResponseBody
	GetTotal() *int32
}

type DescribeEventRuleListResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The event-triggered alert rule.
	EventRules *DescribeEventRuleListResponseBodyEventRules `json:"EventRules,omitempty" xml:"EventRules,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D0E6D82B-16B5-422A-8136-EE5BDC01E415
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 21
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeEventRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeEventRuleListResponseBody) GetEventRules() *DescribeEventRuleListResponseBodyEventRules {
	return s.EventRules
}

func (s *DescribeEventRuleListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEventRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventRuleListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEventRuleListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeEventRuleListResponseBody) SetCode(v string) *DescribeEventRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventRuleListResponseBody) SetEventRules(v *DescribeEventRuleListResponseBodyEventRules) *DescribeEventRuleListResponseBody {
	s.EventRules = v
	return s
}

func (s *DescribeEventRuleListResponseBody) SetMessage(v string) *DescribeEventRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventRuleListResponseBody) SetRequestId(v string) *DescribeEventRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventRuleListResponseBody) SetSuccess(v bool) *DescribeEventRuleListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEventRuleListResponseBody) SetTotal(v int32) *DescribeEventRuleListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeEventRuleListResponseBody) Validate() error {
	if s.EventRules != nil {
		if err := s.EventRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEventRuleListResponseBodyEventRules struct {
	EventRule []*DescribeEventRuleListResponseBodyEventRulesEventRule `json:"EventRule,omitempty" xml:"EventRule,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleListResponseBodyEventRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRules) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRules) GetEventRule() []*DescribeEventRuleListResponseBodyEventRulesEventRule {
	return s.EventRule
}

func (s *DescribeEventRuleListResponseBodyEventRules) SetEventRule(v []*DescribeEventRuleListResponseBodyEventRulesEventRule) *DescribeEventRuleListResponseBodyEventRules {
	s.EventRule = v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRules) Validate() error {
	if s.EventRule != nil {
		for _, item := range s.EventRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEventRuleListResponseBodyEventRulesEventRule struct {
	// The description of the event-triggered alert rule.
	//
	// example:
	//
	// Default group event rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mode of the event-triggered alert rule.
	EventPattern *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern `json:"EventPattern,omitempty" xml:"EventPattern,omitempty" type:"Struct"`
	// The type of the event-triggered alert rule. Valid values:
	//
	// 	- SYSTEM: system event-triggered alert rule
	//
	// 	- CUSTOM: custom event-triggered alert rule
	//
	// example:
	//
	// SYSTEM
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 7378****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the event-triggered alert rule.
	//
	// example:
	//
	// test_DefaultEventRule_7378****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The mute period during which new alert notifications are not sent even if the trigger conditions are met.
	//
	// example:
	//
	// 86400
	SilenceTime *int64 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The status of the event-triggered alert rule. Valid values:
	//
	// 	- ENABLED
	//
	// 	- DISABLED
	//
	// example:
	//
	// ENABLED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRule) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) GetEventPattern() *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern {
	return s.EventPattern
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) GetEventType() *string {
	return s.EventType
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) GetName() *string {
	return s.Name
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) GetSilenceTime() *int64 {
	return s.SilenceTime
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) GetState() *string {
	return s.State
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) SetDescription(v string) *DescribeEventRuleListResponseBodyEventRulesEventRule {
	s.Description = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) SetEventPattern(v *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern) *DescribeEventRuleListResponseBodyEventRulesEventRule {
	s.EventPattern = v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) SetEventType(v string) *DescribeEventRuleListResponseBodyEventRulesEventRule {
	s.EventType = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) SetGroupId(v string) *DescribeEventRuleListResponseBodyEventRulesEventRule {
	s.GroupId = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) SetName(v string) *DescribeEventRuleListResponseBodyEventRulesEventRule {
	s.Name = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) SetSilenceTime(v int64) *DescribeEventRuleListResponseBodyEventRulesEventRule {
	s.SilenceTime = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) SetState(v string) *DescribeEventRuleListResponseBodyEventRulesEventRule {
	s.State = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) Validate() error {
	if s.EventPattern != nil {
		if err := s.EventPattern.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern struct {
	EventPattern []*DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern `json:"EventPattern,omitempty" xml:"EventPattern,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern) GetEventPattern() []*DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern {
	return s.EventPattern
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern) SetEventPattern(v []*DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern {
	s.EventPattern = v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern) Validate() error {
	if s.EventPattern != nil {
		for _, item := range s.EventPattern {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern struct {
	// The custom filter conditions.
	//
	// example:
	//
	// ECS123
	CustomFilters *string `json:"CustomFilters,omitempty" xml:"CustomFilters,omitempty"`
	// The types of the event-triggered alert rules.
	EventTypeList *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty" type:"Struct"`
	// The keyword for filtering.
	KeywordFilter *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilter `json:"KeywordFilter,omitempty" xml:"KeywordFilter,omitempty" type:"Struct"`
	// The levels of the event-triggered alerts.
	LevelList *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList `json:"LevelList,omitempty" xml:"LevelList,omitempty" type:"Struct"`
	// The event names.
	NameList *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Struct"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// CloudMonitor
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// Indicates that logs are filtered based on the specified SQL statement. If the specified conditions are met, an alert is triggered.
	//
	// example:
	//
	// ycccluster1 and (i-23ij0o82612 or Executed1) or Asimulated not 222
	SQLFilter *string `json:"SQLFilter,omitempty" xml:"SQLFilter,omitempty"`
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) GetCustomFilters() *string {
	return s.CustomFilters
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) GetEventTypeList() *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList {
	return s.EventTypeList
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) GetKeywordFilter() *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilter {
	return s.KeywordFilter
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) GetLevelList() *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList {
	return s.LevelList
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) GetNameList() *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList {
	return s.NameList
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) GetProduct() *string {
	return s.Product
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) GetSQLFilter() *string {
	return s.SQLFilter
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) SetCustomFilters(v string) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern {
	s.CustomFilters = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) SetEventTypeList(v *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern {
	s.EventTypeList = v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) SetKeywordFilter(v *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilter) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern {
	s.KeywordFilter = v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) SetLevelList(v *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern {
	s.LevelList = v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) SetNameList(v *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern {
	s.NameList = v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) SetProduct(v string) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern {
	s.Product = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) SetSQLFilter(v string) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern {
	s.SQLFilter = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) Validate() error {
	if s.EventTypeList != nil {
		if err := s.EventTypeList.Validate(); err != nil {
			return err
		}
	}
	if s.KeywordFilter != nil {
		if err := s.KeywordFilter.Validate(); err != nil {
			return err
		}
	}
	if s.LevelList != nil {
		if err := s.LevelList.Validate(); err != nil {
			return err
		}
	}
	if s.NameList != nil {
		if err := s.NameList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList struct {
	EventTypeList []*string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList) GetEventTypeList() []*string {
	return s.EventTypeList
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList) SetEventTypeList(v []*string) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList {
	s.EventTypeList = v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilter struct {
	// The keywords that are used to match events.
	Keywords *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilterKeywords `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Struct"`
	// The relationship between multiple keywords in a condition. Valid values:
	//
	// 	- OR: The relationship between keywords is OR.
	//
	// 	- NOT: The keyword is excluded. The value NOT indicates that all events that do not contain the keywords are matched.
	//
	// example:
	//
	// OR
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilter) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilter) GetKeywords() *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilterKeywords {
	return s.Keywords
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilter) GetRelation() *string {
	return s.Relation
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilter) SetKeywords(v *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilterKeywords) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilter {
	s.Keywords = v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilter) SetRelation(v string) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilter {
	s.Relation = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilter) Validate() error {
	if s.Keywords != nil {
		if err := s.Keywords.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilterKeywords struct {
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilterKeywords) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilterKeywords) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilterKeywords) GetKeywords() []*string {
	return s.Keywords
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilterKeywords) SetKeywords(v []*string) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilterKeywords {
	s.Keywords = v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternKeywordFilterKeywords) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList struct {
	LevelList []*string `json:"LevelList,omitempty" xml:"LevelList,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList) GetLevelList() []*string {
	return s.LevelList
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList) SetLevelList(v []*string) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList {
	s.LevelList = v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList struct {
	NameList []*string `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList) GetNameList() []*string {
	return s.NameList
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList) SetNameList(v []*string) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList {
	s.NameList = v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList) Validate() error {
	return dara.Validate(s)
}
