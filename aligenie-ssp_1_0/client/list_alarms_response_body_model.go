// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlarmsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListAlarmsResponseBody
	GetCode() *int32
	SetMessage(v string) *ListAlarmsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAlarmsResponseBody
	GetRequestId() *string
	SetResult(v *ListAlarmsResponseBodyResult) *ListAlarmsResponseBody
	GetResult() *ListAlarmsResponseBodyResult
}

type ListAlarmsResponseBody struct {
	// example:
	//
	// 200
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43***28C-A810-5***-8747-EC226A086881
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListAlarmsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListAlarmsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListAlarmsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAlarmsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlarmsResponseBody) GetResult() *ListAlarmsResponseBodyResult {
	return s.Result
}

func (s *ListAlarmsResponseBody) SetCode(v int32) *ListAlarmsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlarmsResponseBody) SetMessage(v string) *ListAlarmsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlarmsResponseBody) SetRequestId(v string) *ListAlarmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlarmsResponseBody) SetResult(v *ListAlarmsResponseBodyResult) *ListAlarmsResponseBody {
	s.Result = v
	return s
}

func (s *ListAlarmsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAlarmsResponseBodyResult struct {
	// example:
	//
	// 1
	CurrentPage *int32                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Model       []*ListAlarmsResponseBodyResultModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlarmsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBodyResult) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAlarmsResponseBodyResult) GetModel() []*ListAlarmsResponseBodyResultModel {
	return s.Model
}

func (s *ListAlarmsResponseBodyResult) GetPageCount() *int32 {
	return s.PageCount
}

func (s *ListAlarmsResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlarmsResponseBodyResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAlarmsResponseBodyResult) SetCurrentPage(v int32) *ListAlarmsResponseBodyResult {
	s.CurrentPage = &v
	return s
}

func (s *ListAlarmsResponseBodyResult) SetModel(v []*ListAlarmsResponseBodyResultModel) *ListAlarmsResponseBodyResult {
	s.Model = v
	return s
}

func (s *ListAlarmsResponseBodyResult) SetPageCount(v int32) *ListAlarmsResponseBodyResult {
	s.PageCount = &v
	return s
}

func (s *ListAlarmsResponseBodyResult) SetPageSize(v int32) *ListAlarmsResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListAlarmsResponseBodyResult) SetTotalCount(v int32) *ListAlarmsResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListAlarmsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListAlarmsResponseBodyResultModel struct {
	// example:
	//
	// 1234567
	AlarmId          *int64                                         `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	MusicInfo        *ListAlarmsResponseBodyResultModelMusicInfo    `json:"MusicInfo,omitempty" xml:"MusicInfo,omitempty" type:"Struct"`
	ScheduleInfo     *ListAlarmsResponseBodyResultModelScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	ScheduleTypeDesc *string                                        `json:"ScheduleTypeDesc,omitempty" xml:"ScheduleTypeDesc,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2022-07-29
	TriggerDateDesc *string `json:"TriggerDateDesc,omitempty" xml:"TriggerDateDesc,omitempty"`
	// example:
	//
	// 10:00
	TriggerTimeDesc *string `json:"TriggerTimeDesc,omitempty" xml:"TriggerTimeDesc,omitempty"`
	// example:
	//
	// 40
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s ListAlarmsResponseBodyResultModel) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsResponseBodyResultModel) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBodyResultModel) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *ListAlarmsResponseBodyResultModel) GetMusicInfo() *ListAlarmsResponseBodyResultModelMusicInfo {
	return s.MusicInfo
}

func (s *ListAlarmsResponseBodyResultModel) GetScheduleInfo() *ListAlarmsResponseBodyResultModelScheduleInfo {
	return s.ScheduleInfo
}

func (s *ListAlarmsResponseBodyResultModel) GetScheduleTypeDesc() *string {
	return s.ScheduleTypeDesc
}

func (s *ListAlarmsResponseBodyResultModel) GetStatus() *int32 {
	return s.Status
}

func (s *ListAlarmsResponseBodyResultModel) GetTriggerDateDesc() *string {
	return s.TriggerDateDesc
}

func (s *ListAlarmsResponseBodyResultModel) GetTriggerTimeDesc() *string {
	return s.TriggerTimeDesc
}

func (s *ListAlarmsResponseBodyResultModel) GetVolume() *int32 {
	return s.Volume
}

func (s *ListAlarmsResponseBodyResultModel) SetAlarmId(v int64) *ListAlarmsResponseBodyResultModel {
	s.AlarmId = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModel) SetMusicInfo(v *ListAlarmsResponseBodyResultModelMusicInfo) *ListAlarmsResponseBodyResultModel {
	s.MusicInfo = v
	return s
}

func (s *ListAlarmsResponseBodyResultModel) SetScheduleInfo(v *ListAlarmsResponseBodyResultModelScheduleInfo) *ListAlarmsResponseBodyResultModel {
	s.ScheduleInfo = v
	return s
}

func (s *ListAlarmsResponseBodyResultModel) SetScheduleTypeDesc(v string) *ListAlarmsResponseBodyResultModel {
	s.ScheduleTypeDesc = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModel) SetStatus(v int32) *ListAlarmsResponseBodyResultModel {
	s.Status = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModel) SetTriggerDateDesc(v string) *ListAlarmsResponseBodyResultModel {
	s.TriggerDateDesc = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModel) SetTriggerTimeDesc(v string) *ListAlarmsResponseBodyResultModel {
	s.TriggerTimeDesc = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModel) SetVolume(v int32) *ListAlarmsResponseBodyResultModel {
	s.Volume = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModel) Validate() error {
	return dara.Validate(s)
}

type ListAlarmsResponseBodyResultModelMusicInfo struct {
	// example:
	//
	// 1
	MusicId   *int64  `json:"MusicId,omitempty" xml:"MusicId,omitempty"`
	MusicName *string `json:"MusicName,omitempty" xml:"MusicName,omitempty"`
	// example:
	//
	// 1
	MusicType     *int64  `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	MusicTypeName *string `json:"MusicTypeName,omitempty" xml:"MusicTypeName,omitempty"`
	// example:
	//
	// http://xx
	MusicUrl *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
}

func (s ListAlarmsResponseBodyResultModelMusicInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsResponseBodyResultModelMusicInfo) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) GetMusicId() *int64 {
	return s.MusicId
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) GetMusicName() *string {
	return s.MusicName
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) GetMusicType() *int64 {
	return s.MusicType
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) GetMusicTypeName() *string {
	return s.MusicTypeName
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) GetMusicUrl() *string {
	return s.MusicUrl
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) SetMusicId(v int64) *ListAlarmsResponseBodyResultModelMusicInfo {
	s.MusicId = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) SetMusicName(v string) *ListAlarmsResponseBodyResultModelMusicInfo {
	s.MusicName = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) SetMusicType(v int64) *ListAlarmsResponseBodyResultModelMusicInfo {
	s.MusicType = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) SetMusicTypeName(v string) *ListAlarmsResponseBodyResultModelMusicInfo {
	s.MusicTypeName = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) SetMusicUrl(v string) *ListAlarmsResponseBodyResultModelMusicInfo {
	s.MusicUrl = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelMusicInfo) Validate() error {
	return dara.Validate(s)
}

type ListAlarmsResponseBodyResultModelScheduleInfo struct {
	Once                *ListAlarmsResponseBodyResultModelScheduleInfoOnce                `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	StatutoryWorkingDay *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay `json:"StatutoryWorkingDay,omitempty" xml:"StatutoryWorkingDay,omitempty" type:"Struct"`
	// example:
	//
	// ONCE
	Type   *string                                              `json:"Type,omitempty" xml:"Type,omitempty"`
	Weekly *ListAlarmsResponseBodyResultModelScheduleInfoWeekly `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s ListAlarmsResponseBodyResultModelScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsResponseBodyResultModelScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfo) GetOnce() *ListAlarmsResponseBodyResultModelScheduleInfoOnce {
	return s.Once
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfo) GetStatutoryWorkingDay() *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay {
	return s.StatutoryWorkingDay
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfo) GetType() *string {
	return s.Type
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfo) GetWeekly() *ListAlarmsResponseBodyResultModelScheduleInfoWeekly {
	return s.Weekly
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfo) SetOnce(v *ListAlarmsResponseBodyResultModelScheduleInfoOnce) *ListAlarmsResponseBodyResultModelScheduleInfo {
	s.Once = v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfo) SetStatutoryWorkingDay(v *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay) *ListAlarmsResponseBodyResultModelScheduleInfo {
	s.StatutoryWorkingDay = v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfo) SetType(v string) *ListAlarmsResponseBodyResultModelScheduleInfo {
	s.Type = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfo) SetWeekly(v *ListAlarmsResponseBodyResultModelScheduleInfoWeekly) *ListAlarmsResponseBodyResultModelScheduleInfo {
	s.Weekly = v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfo) Validate() error {
	return dara.Validate(s)
}

type ListAlarmsResponseBodyResultModelScheduleInfoOnce struct {
	// example:
	//
	// 29
	Day *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 0
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	// example:
	//
	// 7
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// example:
	//
	// 2022
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s ListAlarmsResponseBodyResultModelScheduleInfoOnce) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsResponseBodyResultModelScheduleInfoOnce) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) GetDay() *int32 {
	return s.Day
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) GetHour() *int32 {
	return s.Hour
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) GetMinute() *int32 {
	return s.Minute
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) GetMonth() *int32 {
	return s.Month
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) GetYear() *int32 {
	return s.Year
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) SetDay(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoOnce {
	s.Day = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) SetHour(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoOnce {
	s.Hour = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) SetMinute(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoOnce {
	s.Minute = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) SetMonth(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoOnce {
	s.Month = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) SetYear(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoOnce {
	s.Year = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoOnce) Validate() error {
	return dara.Validate(s)
}

type ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay struct {
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 0
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay) GetHour() *int32 {
	return s.Hour
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay) GetMinute() *int32 {
	return s.Minute
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay) SetHour(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay {
	s.Hour = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay) SetMinute(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay {
	s.Minute = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoStatutoryWorkingDay) Validate() error {
	return dara.Validate(s)
}

type ListAlarmsResponseBodyResultModelScheduleInfoWeekly struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 0
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s ListAlarmsResponseBodyResultModelScheduleInfoWeekly) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsResponseBodyResultModelScheduleInfoWeekly) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoWeekly) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoWeekly) GetHour() *int32 {
	return s.Hour
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoWeekly) GetMinute() *int32 {
	return s.Minute
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoWeekly) SetDaysOfWeek(v []*int32) *ListAlarmsResponseBodyResultModelScheduleInfoWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoWeekly) SetHour(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoWeekly {
	s.Hour = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoWeekly) SetMinute(v int32) *ListAlarmsResponseBodyResultModelScheduleInfoWeekly {
	s.Minute = &v
	return s
}

func (s *ListAlarmsResponseBodyResultModelScheduleInfoWeekly) Validate() error {
	return dara.Validate(s)
}
