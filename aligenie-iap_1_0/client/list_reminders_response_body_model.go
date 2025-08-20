// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemindersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v int32) *ListRemindersResponseBody
	GetErrorCode() *int32
	SetErrorMsg(v string) *ListRemindersResponseBody
	GetErrorMsg() *string
	SetModel(v *ListRemindersResponseBodyModel) *ListRemindersResponseBody
	GetModel() *ListRemindersResponseBodyModel
	SetSuccess(v bool) *ListRemindersResponseBody
	GetSuccess() *bool
}

type ListRemindersResponseBody struct {
	// example:
	//
	// 400
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 参数不合法。
	ErrorMsg *string                         `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Model    *ListRemindersResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRemindersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRemindersResponseBody) GoString() string {
	return s.String()
}

func (s *ListRemindersResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *ListRemindersResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListRemindersResponseBody) GetModel() *ListRemindersResponseBodyModel {
	return s.Model
}

func (s *ListRemindersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRemindersResponseBody) SetErrorCode(v int32) *ListRemindersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRemindersResponseBody) SetErrorMsg(v string) *ListRemindersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListRemindersResponseBody) SetModel(v *ListRemindersResponseBodyModel) *ListRemindersResponseBody {
	s.Model = v
	return s
}

func (s *ListRemindersResponseBody) SetSuccess(v bool) *ListRemindersResponseBody {
	s.Success = &v
	return s
}

func (s *ListRemindersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRemindersResponseBodyModel struct {
	RemindResponses []*ListRemindersResponseBodyModelRemindResponses `json:"RemindResponses,omitempty" xml:"RemindResponses,omitempty" type:"Repeated"`
}

func (s ListRemindersResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s ListRemindersResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ListRemindersResponseBodyModel) GetRemindResponses() []*ListRemindersResponseBodyModelRemindResponses {
	return s.RemindResponses
}

func (s *ListRemindersResponseBodyModel) SetRemindResponses(v []*ListRemindersResponseBodyModelRemindResponses) *ListRemindersResponseBodyModel {
	s.RemindResponses = v
	return s
}

func (s *ListRemindersResponseBodyModel) Validate() error {
	return dara.Validate(s)
}

type ListRemindersResponseBodyModelRemindResponses struct {
	// example:
	//
	// 宝宝快去刷牙
	ActionTopic *string `json:"ActionTopic,omitempty" xml:"ActionTopic,omitempty"`
	// example:
	//
	// 每天
	DayDesc        *string                                                      `json:"DayDesc,omitempty" xml:"DayDesc,omitempty"`
	RecurrenceRule *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule `json:"RecurrenceRule,omitempty" xml:"RecurrenceRule,omitempty" type:"Struct"`
	// example:
	//
	// 20****1
	RemindId *int64 `json:"RemindId,omitempty" xml:"RemindId,omitempty"`
	// example:
	//
	// 1629850800000
	RemindTime *string `json:"RemindTime,omitempty" xml:"RemindTime,omitempty"`
	// example:
	//
	// 1
	RepeatCount *int32 `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	// example:
	//
	// 周三
	Week *string `json:"Week,omitempty" xml:"Week,omitempty"`
}

func (s ListRemindersResponseBodyModelRemindResponses) String() string {
	return dara.Prettify(s)
}

func (s ListRemindersResponseBodyModelRemindResponses) GoString() string {
	return s.String()
}

func (s *ListRemindersResponseBodyModelRemindResponses) GetActionTopic() *string {
	return s.ActionTopic
}

func (s *ListRemindersResponseBodyModelRemindResponses) GetDayDesc() *string {
	return s.DayDesc
}

func (s *ListRemindersResponseBodyModelRemindResponses) GetRecurrenceRule() *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	return s.RecurrenceRule
}

func (s *ListRemindersResponseBodyModelRemindResponses) GetRemindId() *int64 {
	return s.RemindId
}

func (s *ListRemindersResponseBodyModelRemindResponses) GetRemindTime() *string {
	return s.RemindTime
}

func (s *ListRemindersResponseBodyModelRemindResponses) GetRepeatCount() *int32 {
	return s.RepeatCount
}

func (s *ListRemindersResponseBodyModelRemindResponses) GetWeek() *string {
	return s.Week
}

func (s *ListRemindersResponseBodyModelRemindResponses) SetActionTopic(v string) *ListRemindersResponseBodyModelRemindResponses {
	s.ActionTopic = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponses) SetDayDesc(v string) *ListRemindersResponseBodyModelRemindResponses {
	s.DayDesc = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponses) SetRecurrenceRule(v *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) *ListRemindersResponseBodyModelRemindResponses {
	s.RecurrenceRule = v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponses) SetRemindId(v int64) *ListRemindersResponseBodyModelRemindResponses {
	s.RemindId = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponses) SetRemindTime(v string) *ListRemindersResponseBodyModelRemindResponses {
	s.RemindTime = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponses) SetRepeatCount(v int32) *ListRemindersResponseBodyModelRemindResponses {
	s.RepeatCount = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponses) SetWeek(v string) *ListRemindersResponseBodyModelRemindResponses {
	s.Week = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponses) Validate() error {
	return dara.Validate(s)
}

type ListRemindersResponseBodyModelRemindResponsesRecurrenceRule struct {
	// example:
	//
	// 27
	Day         *int32   `json:"Day,omitempty" xml:"Day,omitempty"`
	DaysOfMonth []*int32 `json:"DaysOfMonth,omitempty" xml:"DaysOfMonth,omitempty" type:"Repeated"`
	DaysOfWeek  []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	// example:
	//
	// 1661598000000
	EndDateTime *string `json:"EndDateTime,omitempty" xml:"EndDateTime,omitempty"`
	// example:
	//
	// WEEKLY
	Freq *string `json:"Freq,omitempty" xml:"Freq,omitempty"`
	// example:
	//
	// 18
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 0
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	// example:
	//
	// 8
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// example:
	//
	// 0
	Second *int32 `json:"Second,omitempty" xml:"Second,omitempty"`
	// example:
	//
	// 1630054800000
	StartDateTime *string `json:"StartDateTime,omitempty" xml:"StartDateTime,omitempty"`
	// example:
	//
	// 2021
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) String() string {
	return dara.Prettify(s)
}

func (s ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) GoString() string {
	return s.String()
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) GetDay() *int32 {
	return s.Day
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) GetDaysOfMonth() []*int32 {
	return s.DaysOfMonth
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) GetEndDateTime() *string {
	return s.EndDateTime
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) GetFreq() *string {
	return s.Freq
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) GetHour() *int32 {
	return s.Hour
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) GetMinute() *int32 {
	return s.Minute
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) GetMonth() *int32 {
	return s.Month
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) GetSecond() *int32 {
	return s.Second
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) GetStartDateTime() *string {
	return s.StartDateTime
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) GetYear() *int32 {
	return s.Year
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetDay(v int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.Day = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetDaysOfMonth(v []*int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.DaysOfMonth = v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetDaysOfWeek(v []*int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.DaysOfWeek = v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetEndDateTime(v string) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.EndDateTime = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetFreq(v string) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.Freq = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetHour(v int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.Hour = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetMinute(v int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.Minute = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetMonth(v int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.Month = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetSecond(v int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.Second = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetStartDateTime(v string) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.StartDateTime = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetYear(v int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.Year = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) Validate() error {
	return dara.Validate(s)
}
