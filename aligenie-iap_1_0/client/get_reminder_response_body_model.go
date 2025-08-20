// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReminderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v int32) *GetReminderResponseBody
	GetErrorCode() *int32
	SetErrorMsg(v string) *GetReminderResponseBody
	GetErrorMsg() *string
	SetModel(v *GetReminderResponseBodyModel) *GetReminderResponseBody
	GetModel() *GetReminderResponseBodyModel
	SetSuccess(v bool) *GetReminderResponseBody
	GetSuccess() *bool
}

type GetReminderResponseBody struct {
	// example:
	//
	// 400
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 参数不合法。
	ErrorMsg *string                       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Model    *GetReminderResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetReminderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetReminderResponseBody) GoString() string {
	return s.String()
}

func (s *GetReminderResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetReminderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetReminderResponseBody) GetModel() *GetReminderResponseBodyModel {
	return s.Model
}

func (s *GetReminderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetReminderResponseBody) SetErrorCode(v int32) *GetReminderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetReminderResponseBody) SetErrorMsg(v string) *GetReminderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetReminderResponseBody) SetModel(v *GetReminderResponseBodyModel) *GetReminderResponseBody {
	s.Model = v
	return s
}

func (s *GetReminderResponseBody) SetSuccess(v bool) *GetReminderResponseBody {
	s.Success = &v
	return s
}

func (s *GetReminderResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetReminderResponseBodyModel struct {
	RemindResponses []*GetReminderResponseBodyModelRemindResponses `json:"RemindResponses,omitempty" xml:"RemindResponses,omitempty" type:"Repeated"`
}

func (s GetReminderResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s GetReminderResponseBodyModel) GoString() string {
	return s.String()
}

func (s *GetReminderResponseBodyModel) GetRemindResponses() []*GetReminderResponseBodyModelRemindResponses {
	return s.RemindResponses
}

func (s *GetReminderResponseBodyModel) SetRemindResponses(v []*GetReminderResponseBodyModelRemindResponses) *GetReminderResponseBodyModel {
	s.RemindResponses = v
	return s
}

func (s *GetReminderResponseBodyModel) Validate() error {
	return dara.Validate(s)
}

type GetReminderResponseBodyModelRemindResponses struct {
	// example:
	//
	// 宝宝快去刷牙
	ActionTopic *string `json:"ActionTopic,omitempty" xml:"ActionTopic,omitempty"`
	// example:
	//
	// 每天
	DayDesc        *string                                                    `json:"DayDesc,omitempty" xml:"DayDesc,omitempty"`
	RecurrenceRule *GetReminderResponseBodyModelRemindResponsesRecurrenceRule `json:"RecurrenceRule,omitempty" xml:"RecurrenceRule,omitempty" type:"Struct"`
	// example:
	//
	// 20*****1
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

func (s GetReminderResponseBodyModelRemindResponses) String() string {
	return dara.Prettify(s)
}

func (s GetReminderResponseBodyModelRemindResponses) GoString() string {
	return s.String()
}

func (s *GetReminderResponseBodyModelRemindResponses) GetActionTopic() *string {
	return s.ActionTopic
}

func (s *GetReminderResponseBodyModelRemindResponses) GetDayDesc() *string {
	return s.DayDesc
}

func (s *GetReminderResponseBodyModelRemindResponses) GetRecurrenceRule() *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	return s.RecurrenceRule
}

func (s *GetReminderResponseBodyModelRemindResponses) GetRemindId() *int64 {
	return s.RemindId
}

func (s *GetReminderResponseBodyModelRemindResponses) GetRemindTime() *string {
	return s.RemindTime
}

func (s *GetReminderResponseBodyModelRemindResponses) GetRepeatCount() *int32 {
	return s.RepeatCount
}

func (s *GetReminderResponseBodyModelRemindResponses) GetWeek() *string {
	return s.Week
}

func (s *GetReminderResponseBodyModelRemindResponses) SetActionTopic(v string) *GetReminderResponseBodyModelRemindResponses {
	s.ActionTopic = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponses) SetDayDesc(v string) *GetReminderResponseBodyModelRemindResponses {
	s.DayDesc = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponses) SetRecurrenceRule(v *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) *GetReminderResponseBodyModelRemindResponses {
	s.RecurrenceRule = v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponses) SetRemindId(v int64) *GetReminderResponseBodyModelRemindResponses {
	s.RemindId = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponses) SetRemindTime(v string) *GetReminderResponseBodyModelRemindResponses {
	s.RemindTime = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponses) SetRepeatCount(v int32) *GetReminderResponseBodyModelRemindResponses {
	s.RepeatCount = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponses) SetWeek(v string) *GetReminderResponseBodyModelRemindResponses {
	s.Week = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponses) Validate() error {
	return dara.Validate(s)
}

type GetReminderResponseBodyModelRemindResponsesRecurrenceRule struct {
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

func (s GetReminderResponseBodyModelRemindResponsesRecurrenceRule) String() string {
	return dara.Prettify(s)
}

func (s GetReminderResponseBodyModelRemindResponsesRecurrenceRule) GoString() string {
	return s.String()
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) GetDay() *int32 {
	return s.Day
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) GetDaysOfMonth() []*int32 {
	return s.DaysOfMonth
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) GetEndDateTime() *string {
	return s.EndDateTime
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) GetFreq() *string {
	return s.Freq
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) GetHour() *int32 {
	return s.Hour
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) GetMinute() *int32 {
	return s.Minute
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) GetMonth() *int32 {
	return s.Month
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) GetSecond() *int32 {
	return s.Second
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) GetStartDateTime() *string {
	return s.StartDateTime
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) GetYear() *int32 {
	return s.Year
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetDay(v int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.Day = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetDaysOfMonth(v []*int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.DaysOfMonth = v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetDaysOfWeek(v []*int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.DaysOfWeek = v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetEndDateTime(v string) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.EndDateTime = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetFreq(v string) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.Freq = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetHour(v int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.Hour = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetMinute(v int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.Minute = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetMonth(v int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.Month = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetSecond(v int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.Second = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetStartDateTime(v string) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.StartDateTime = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetYear(v int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.Year = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) Validate() error {
	return dara.Validate(s)
}
