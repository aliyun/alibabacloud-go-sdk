// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAlarmResponseBody
	GetCode() *int32
	SetMessage(v string) *GetAlarmResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAlarmResponseBody
	GetRequestId() *string
	SetResult(v *GetAlarmResponseBodyResult) *GetAlarmResponseBody
	GetResult() *GetAlarmResponseBodyResult
}

type GetAlarmResponseBody struct {
	// Status code returned by the alarm service
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// error message
	//
	// example:
	//
	// id为空
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 43***28C-A810-5***-8747-EC226A086881
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Alarm details
	Result *GetAlarmResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlarmResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAlarmResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlarmResponseBody) GetResult() *GetAlarmResponseBodyResult {
	return s.Result
}

func (s *GetAlarmResponseBody) SetCode(v int32) *GetAlarmResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlarmResponseBody) SetMessage(v string) *GetAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *GetAlarmResponseBody) SetRequestId(v string) *GetAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlarmResponseBody) SetResult(v *GetAlarmResponseBodyResult) *GetAlarmResponseBody {
	s.Result = v
	return s
}

func (s *GetAlarmResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlarmResponseBodyResult struct {
	// Alarm ID
	//
	// example:
	//
	// 1234567
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// Ringtone Information
	MusicInfo *GetAlarmResponseBodyResultMusicInfo `json:"MusicInfo,omitempty" xml:"MusicInfo,omitempty" type:"Struct"`
	// Schedule Information
	ScheduleInfo *GetAlarmResponseBodyResultScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	// Chinese description of the loop type
	//
	// example:
	//
	// 单次
	ScheduleTypeDesc *string `json:"ScheduleTypeDesc,omitempty" xml:"ScheduleTypeDesc,omitempty"`
	// status: 0 Normal, 1 deleted, 2 shutdown
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Trigger date description (one-time)
	//
	// example:
	//
	// 2022-07-29
	TriggerDateDesc *string `json:"TriggerDateDesc,omitempty" xml:"TriggerDateDesc,omitempty"`
	// Trigger time description
	//
	// example:
	//
	// 10:00
	TriggerTimeDesc *string `json:"TriggerTimeDesc,omitempty" xml:"TriggerTimeDesc,omitempty"`
	// Ringtone volume
	//
	// example:
	//
	// 40
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s GetAlarmResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAlarmResponseBodyResult) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *GetAlarmResponseBodyResult) GetMusicInfo() *GetAlarmResponseBodyResultMusicInfo {
	return s.MusicInfo
}

func (s *GetAlarmResponseBodyResult) GetScheduleInfo() *GetAlarmResponseBodyResultScheduleInfo {
	return s.ScheduleInfo
}

func (s *GetAlarmResponseBodyResult) GetScheduleTypeDesc() *string {
	return s.ScheduleTypeDesc
}

func (s *GetAlarmResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *GetAlarmResponseBodyResult) GetTriggerDateDesc() *string {
	return s.TriggerDateDesc
}

func (s *GetAlarmResponseBodyResult) GetTriggerTimeDesc() *string {
	return s.TriggerTimeDesc
}

func (s *GetAlarmResponseBodyResult) GetVolume() *int32 {
	return s.Volume
}

func (s *GetAlarmResponseBodyResult) SetAlarmId(v int64) *GetAlarmResponseBodyResult {
	s.AlarmId = &v
	return s
}

func (s *GetAlarmResponseBodyResult) SetMusicInfo(v *GetAlarmResponseBodyResultMusicInfo) *GetAlarmResponseBodyResult {
	s.MusicInfo = v
	return s
}

func (s *GetAlarmResponseBodyResult) SetScheduleInfo(v *GetAlarmResponseBodyResultScheduleInfo) *GetAlarmResponseBodyResult {
	s.ScheduleInfo = v
	return s
}

func (s *GetAlarmResponseBodyResult) SetScheduleTypeDesc(v string) *GetAlarmResponseBodyResult {
	s.ScheduleTypeDesc = &v
	return s
}

func (s *GetAlarmResponseBodyResult) SetStatus(v int32) *GetAlarmResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetAlarmResponseBodyResult) SetTriggerDateDesc(v string) *GetAlarmResponseBodyResult {
	s.TriggerDateDesc = &v
	return s
}

func (s *GetAlarmResponseBodyResult) SetTriggerTimeDesc(v string) *GetAlarmResponseBodyResult {
	s.TriggerTimeDesc = &v
	return s
}

func (s *GetAlarmResponseBodyResult) SetVolume(v int32) *GetAlarmResponseBodyResult {
	s.Volume = &v
	return s
}

func (s *GetAlarmResponseBodyResult) Validate() error {
	if s.MusicInfo != nil {
		if err := s.MusicInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduleInfo != nil {
		if err := s.ScheduleInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlarmResponseBodyResultMusicInfo struct {
	// Ringtone ID
	//
	// example:
	//
	// 1
	MusicId *int64 `json:"MusicId,omitempty" xml:"MusicId,omitempty"`
	// Ringtone Name
	//
	// example:
	//
	// xx铃声
	MusicName *string `json:"MusicName,omitempty" xml:"MusicName,omitempty"`
	// Ringtone Category ID
	//
	// example:
	//
	// 1
	MusicType *int64 `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	// Ringtone Category Name
	//
	// example:
	//
	// xx音乐
	MusicTypeName *string `json:"MusicTypeName,omitempty" xml:"MusicTypeName,omitempty"`
	// Ringtone URL
	//
	// example:
	//
	// http://xx
	MusicUrl *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
}

func (s GetAlarmResponseBodyResultMusicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmResponseBodyResultMusicInfo) GoString() string {
	return s.String()
}

func (s *GetAlarmResponseBodyResultMusicInfo) GetMusicId() *int64 {
	return s.MusicId
}

func (s *GetAlarmResponseBodyResultMusicInfo) GetMusicName() *string {
	return s.MusicName
}

func (s *GetAlarmResponseBodyResultMusicInfo) GetMusicType() *int64 {
	return s.MusicType
}

func (s *GetAlarmResponseBodyResultMusicInfo) GetMusicTypeName() *string {
	return s.MusicTypeName
}

func (s *GetAlarmResponseBodyResultMusicInfo) GetMusicUrl() *string {
	return s.MusicUrl
}

func (s *GetAlarmResponseBodyResultMusicInfo) SetMusicId(v int64) *GetAlarmResponseBodyResultMusicInfo {
	s.MusicId = &v
	return s
}

func (s *GetAlarmResponseBodyResultMusicInfo) SetMusicName(v string) *GetAlarmResponseBodyResultMusicInfo {
	s.MusicName = &v
	return s
}

func (s *GetAlarmResponseBodyResultMusicInfo) SetMusicType(v int64) *GetAlarmResponseBodyResultMusicInfo {
	s.MusicType = &v
	return s
}

func (s *GetAlarmResponseBodyResultMusicInfo) SetMusicTypeName(v string) *GetAlarmResponseBodyResultMusicInfo {
	s.MusicTypeName = &v
	return s
}

func (s *GetAlarmResponseBodyResultMusicInfo) SetMusicUrl(v string) *GetAlarmResponseBodyResultMusicInfo {
	s.MusicUrl = &v
	return s
}

func (s *GetAlarmResponseBodyResultMusicInfo) Validate() error {
	return dara.Validate(s)
}

type GetAlarmResponseBodyResultScheduleInfo struct {
	// One-time: This property is active when the loop type is ONCE.
	Once *GetAlarmResponseBodyResultScheduleInfoOnce `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	// Statutory working day: This property is active when the loop Type is STATUTORYWORKINGDAY.
	StatutoryWorkingDay *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay `json:"StatutoryWorkingDay,omitempty" xml:"StatutoryWorkingDay,omitempty" type:"Struct"`
	// Schedule Type / Loop Type: ONCE -> One-time, WEEKLY -> Weekly loop, STATUTORYWORKINGDAY -> Statutory working day
	//
	// example:
	//
	// ONCE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Weekly loop: This property is active when the loop Type is WEEKLY.
	Weekly *GetAlarmResponseBodyResultScheduleInfoWeekly `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s GetAlarmResponseBodyResultScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmResponseBodyResultScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetAlarmResponseBodyResultScheduleInfo) GetOnce() *GetAlarmResponseBodyResultScheduleInfoOnce {
	return s.Once
}

func (s *GetAlarmResponseBodyResultScheduleInfo) GetStatutoryWorkingDay() *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay {
	return s.StatutoryWorkingDay
}

func (s *GetAlarmResponseBodyResultScheduleInfo) GetType() *string {
	return s.Type
}

func (s *GetAlarmResponseBodyResultScheduleInfo) GetWeekly() *GetAlarmResponseBodyResultScheduleInfoWeekly {
	return s.Weekly
}

func (s *GetAlarmResponseBodyResultScheduleInfo) SetOnce(v *GetAlarmResponseBodyResultScheduleInfoOnce) *GetAlarmResponseBodyResultScheduleInfo {
	s.Once = v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfo) SetStatutoryWorkingDay(v *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay) *GetAlarmResponseBodyResultScheduleInfo {
	s.StatutoryWorkingDay = v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfo) SetType(v string) *GetAlarmResponseBodyResultScheduleInfo {
	s.Type = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfo) SetWeekly(v *GetAlarmResponseBodyResultScheduleInfoWeekly) *GetAlarmResponseBodyResultScheduleInfo {
	s.Weekly = v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfo) Validate() error {
	if s.Once != nil {
		if err := s.Once.Validate(); err != nil {
			return err
		}
	}
	if s.StatutoryWorkingDay != nil {
		if err := s.StatutoryWorkingDay.Validate(); err != nil {
			return err
		}
	}
	if s.Weekly != nil {
		if err := s.Weekly.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlarmResponseBodyResultScheduleInfoOnce struct {
	// Trigger time: Day
	//
	// example:
	//
	// 29
	Day *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	// Trigger Time: Hour
	//
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// Trigger Time: Minute
	//
	// example:
	//
	// 0
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	// Trigger Time: Month
	//
	// example:
	//
	// 7
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// Trigger Time: Year
	//
	// example:
	//
	// 2022
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s GetAlarmResponseBodyResultScheduleInfoOnce) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmResponseBodyResultScheduleInfoOnce) GoString() string {
	return s.String()
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) GetDay() *int32 {
	return s.Day
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) GetHour() *int32 {
	return s.Hour
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) GetMinute() *int32 {
	return s.Minute
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) GetMonth() *int32 {
	return s.Month
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) GetYear() *int32 {
	return s.Year
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) SetDay(v int32) *GetAlarmResponseBodyResultScheduleInfoOnce {
	s.Day = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) SetHour(v int32) *GetAlarmResponseBodyResultScheduleInfoOnce {
	s.Hour = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) SetMinute(v int32) *GetAlarmResponseBodyResultScheduleInfoOnce {
	s.Minute = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) SetMonth(v int32) *GetAlarmResponseBodyResultScheduleInfoOnce {
	s.Month = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) SetYear(v int32) *GetAlarmResponseBodyResultScheduleInfoOnce {
	s.Year = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoOnce) Validate() error {
	return dara.Validate(s)
}

type GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay struct {
	// Trigger Time: Hour
	//
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// Trigger Time: Minute
	//
	// example:
	//
	// 0
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay) GoString() string {
	return s.String()
}

func (s *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay) GetHour() *int32 {
	return s.Hour
}

func (s *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay) GetMinute() *int32 {
	return s.Minute
}

func (s *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay) SetHour(v int32) *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay {
	s.Hour = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay) SetMinute(v int32) *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay {
	s.Minute = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoStatutoryWorkingDay) Validate() error {
	return dara.Validate(s)
}

type GetAlarmResponseBodyResultScheduleInfoWeekly struct {
	// Collection of days of the week to trigger: Numeric values between 1 and 7, where each number corresponds to a specific day of the week. If triggered every day, include all numbers.
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	// Trigger time: Hour
	//
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// Trigger time: Minute
	//
	// example:
	//
	// 0
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s GetAlarmResponseBodyResultScheduleInfoWeekly) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmResponseBodyResultScheduleInfoWeekly) GoString() string {
	return s.String()
}

func (s *GetAlarmResponseBodyResultScheduleInfoWeekly) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *GetAlarmResponseBodyResultScheduleInfoWeekly) GetHour() *int32 {
	return s.Hour
}

func (s *GetAlarmResponseBodyResultScheduleInfoWeekly) GetMinute() *int32 {
	return s.Minute
}

func (s *GetAlarmResponseBodyResultScheduleInfoWeekly) SetDaysOfWeek(v []*int32) *GetAlarmResponseBodyResultScheduleInfoWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoWeekly) SetHour(v int32) *GetAlarmResponseBodyResultScheduleInfoWeekly {
	s.Hour = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoWeekly) SetMinute(v int32) *GetAlarmResponseBodyResultScheduleInfoWeekly {
	s.Minute = &v
	return s
}

func (s *GetAlarmResponseBodyResultScheduleInfoWeekly) Validate() error {
	return dara.Validate(s)
}
