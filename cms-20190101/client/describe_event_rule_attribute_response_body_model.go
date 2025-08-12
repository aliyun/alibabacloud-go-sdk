// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventRuleAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeEventRuleAttributeResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeEventRuleAttributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEventRuleAttributeResponseBody
	GetRequestId() *string
	SetResult(v *DescribeEventRuleAttributeResponseBodyResult) *DescribeEventRuleAttributeResponseBody
	GetResult() *DescribeEventRuleAttributeResponseBodyResult
	SetSuccess(v bool) *DescribeEventRuleAttributeResponseBody
	GetSuccess() *bool
}

type DescribeEventRuleAttributeResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The alert does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AA3F210-C03D-4C86-8DB6-21C84FF692A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the event-triggered alert rule.
	Result *DescribeEventRuleAttributeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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
}

func (s DescribeEventRuleAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeEventRuleAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEventRuleAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventRuleAttributeResponseBody) GetResult() *DescribeEventRuleAttributeResponseBodyResult {
	return s.Result
}

func (s *DescribeEventRuleAttributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEventRuleAttributeResponseBody) SetCode(v string) *DescribeEventRuleAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBody) SetMessage(v string) *DescribeEventRuleAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBody) SetRequestId(v string) *DescribeEventRuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBody) SetResult(v *DescribeEventRuleAttributeResponseBodyResult) *DescribeEventRuleAttributeResponseBody {
	s.Result = v
	return s
}

func (s *DescribeEventRuleAttributeResponseBody) SetSuccess(v bool) *DescribeEventRuleAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleAttributeResponseBodyResult struct {
	// The description of the event-triggered alert rule.
	//
	// example:
	//
	// Default group event rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The event pattern. This parameter describes the trigger conditions of an event.
	EventPattern *DescribeEventRuleAttributeResponseBodyResultEventPattern `json:"EventPattern,omitempty" xml:"EventPattern,omitempty" type:"Struct"`
	// The event type. Valid values:
	//
	// 	- SYSTEM: system event
	//
	// 	- CUSTOM: custom event
	//
	// example:
	//
	// SYSTEM
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 3607****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the event-triggered alert rule.
	//
	// example:
	//
	// test_DefaultEventRule_7378****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s DescribeEventRuleAttributeResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleAttributeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventRuleAttributeResponseBodyResult) GetEventPattern() *DescribeEventRuleAttributeResponseBodyResultEventPattern {
	return s.EventPattern
}

func (s *DescribeEventRuleAttributeResponseBodyResult) GetEventType() *string {
	return s.EventType
}

func (s *DescribeEventRuleAttributeResponseBodyResult) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeEventRuleAttributeResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *DescribeEventRuleAttributeResponseBodyResult) GetState() *string {
	return s.State
}

func (s *DescribeEventRuleAttributeResponseBodyResult) SetDescription(v string) *DescribeEventRuleAttributeResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResult) SetEventPattern(v *DescribeEventRuleAttributeResponseBodyResultEventPattern) *DescribeEventRuleAttributeResponseBodyResult {
	s.EventPattern = v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResult) SetEventType(v string) *DescribeEventRuleAttributeResponseBodyResult {
	s.EventType = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResult) SetGroupId(v string) *DescribeEventRuleAttributeResponseBodyResult {
	s.GroupId = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResult) SetName(v string) *DescribeEventRuleAttributeResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResult) SetState(v string) *DescribeEventRuleAttributeResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleAttributeResponseBodyResultEventPattern struct {
	// The types of the event-triggered alert rules.
	EventTypeList *DescribeEventRuleAttributeResponseBodyResultEventPatternEventTypeList `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty" type:"Struct"`
	// The keyword for filtering.
	KeywordFilterObj *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObj `json:"KeywordFilterObj,omitempty" xml:"KeywordFilterObj,omitempty" type:"Struct"`
	LevelList        *DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList        `json:"LevelList,omitempty" xml:"LevelList,omitempty" type:"Struct"`
	NameList         *DescribeEventRuleAttributeResponseBodyResultEventPatternNameList         `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Struct"`
	// The name of the cloud service.
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
	SQLFilter  *string                                                             `json:"SQLFilter,omitempty" xml:"SQLFilter,omitempty"`
	StatusList *DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Struct"`
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPattern) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPattern) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) GetEventTypeList() *DescribeEventRuleAttributeResponseBodyResultEventPatternEventTypeList {
	return s.EventTypeList
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) GetKeywordFilterObj() *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObj {
	return s.KeywordFilterObj
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) GetLevelList() *DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList {
	return s.LevelList
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) GetNameList() *DescribeEventRuleAttributeResponseBodyResultEventPatternNameList {
	return s.NameList
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) GetProduct() *string {
	return s.Product
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) GetSQLFilter() *string {
	return s.SQLFilter
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) GetStatusList() *DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList {
	return s.StatusList
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) SetEventTypeList(v *DescribeEventRuleAttributeResponseBodyResultEventPatternEventTypeList) *DescribeEventRuleAttributeResponseBodyResultEventPattern {
	s.EventTypeList = v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) SetKeywordFilterObj(v *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObj) *DescribeEventRuleAttributeResponseBodyResultEventPattern {
	s.KeywordFilterObj = v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) SetLevelList(v *DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList) *DescribeEventRuleAttributeResponseBodyResultEventPattern {
	s.LevelList = v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) SetNameList(v *DescribeEventRuleAttributeResponseBodyResultEventPatternNameList) *DescribeEventRuleAttributeResponseBodyResultEventPattern {
	s.NameList = v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) SetProduct(v string) *DescribeEventRuleAttributeResponseBodyResultEventPattern {
	s.Product = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) SetSQLFilter(v string) *DescribeEventRuleAttributeResponseBodyResultEventPattern {
	s.SQLFilter = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) SetStatusList(v *DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList) *DescribeEventRuleAttributeResponseBodyResultEventPattern {
	s.StatusList = v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleAttributeResponseBodyResultEventPatternEventTypeList struct {
	EventTypeList []*string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternEventTypeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternEventTypeList) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternEventTypeList) GetEventTypeList() []*string {
	return s.EventTypeList
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternEventTypeList) SetEventTypeList(v []*string) *DescribeEventRuleAttributeResponseBodyResultEventPatternEventTypeList {
	s.EventTypeList = v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternEventTypeList) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObj struct {
	// The keywords that are used to match events.
	Keywords *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObjKeywords `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Struct"`
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

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObj) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObj) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObj) GetKeywords() *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObjKeywords {
	return s.Keywords
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObj) GetRelation() *string {
	return s.Relation
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObj) SetKeywords(v *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObjKeywords) *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObj {
	s.Keywords = v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObj) SetRelation(v string) *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObj {
	s.Relation = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObj) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObjKeywords struct {
	Keyword []*string `json:"keyword,omitempty" xml:"keyword,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObjKeywords) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObjKeywords) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObjKeywords) GetKeyword() []*string {
	return s.Keyword
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObjKeywords) SetKeyword(v []*string) *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObjKeywords {
	s.Keyword = v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternKeywordFilterObjKeywords) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList struct {
	LevelList []*string `json:"LevelList,omitempty" xml:"LevelList,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList) GetLevelList() []*string {
	return s.LevelList
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList) SetLevelList(v []*string) *DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList {
	s.LevelList = v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleAttributeResponseBodyResultEventPatternNameList struct {
	NameList []*string `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternNameList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternNameList) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternNameList) GetNameList() []*string {
	return s.NameList
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternNameList) SetNameList(v []*string) *DescribeEventRuleAttributeResponseBodyResultEventPatternNameList {
	s.NameList = v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternNameList) Validate() error {
	return dara.Validate(s)
}

type DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList struct {
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList) GetStatusList() []*string {
	return s.StatusList
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList) SetStatusList(v []*string) *DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList {
	s.StatusList = v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList) Validate() error {
	return dara.Validate(s)
}
