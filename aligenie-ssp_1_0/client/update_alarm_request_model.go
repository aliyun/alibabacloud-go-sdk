// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *UpdateAlarmRequestDeviceInfo) *UpdateAlarmRequest
	GetDeviceInfo() *UpdateAlarmRequestDeviceInfo
	SetPayload(v *UpdateAlarmRequestPayload) *UpdateAlarmRequest
	GetPayload() *UpdateAlarmRequestPayload
	SetUserInfo(v *UpdateAlarmRequestUserInfo) *UpdateAlarmRequest
	GetUserInfo() *UpdateAlarmRequestUserInfo
}

type UpdateAlarmRequest struct {
	// This parameter is required.
	DeviceInfo *UpdateAlarmRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Payload *UpdateAlarmRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *UpdateAlarmRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s UpdateAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequest) GetDeviceInfo() *UpdateAlarmRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *UpdateAlarmRequest) GetPayload() *UpdateAlarmRequestPayload {
	return s.Payload
}

func (s *UpdateAlarmRequest) GetUserInfo() *UpdateAlarmRequestUserInfo {
	return s.UserInfo
}

func (s *UpdateAlarmRequest) SetDeviceInfo(v *UpdateAlarmRequestDeviceInfo) *UpdateAlarmRequest {
	s.DeviceInfo = v
	return s
}

func (s *UpdateAlarmRequest) SetPayload(v *UpdateAlarmRequestPayload) *UpdateAlarmRequest {
	s.Payload = v
	return s
}

func (s *UpdateAlarmRequest) SetUserInfo(v *UpdateAlarmRequestUserInfo) *UpdateAlarmRequest {
	s.UserInfo = v
	return s
}

func (s *UpdateAlarmRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateAlarmRequestDeviceInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UpdateAlarmRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *UpdateAlarmRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *UpdateAlarmRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *UpdateAlarmRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *UpdateAlarmRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateAlarmRequestDeviceInfo) SetEncodeKey(v string) *UpdateAlarmRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *UpdateAlarmRequestDeviceInfo) SetEncodeType(v string) *UpdateAlarmRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *UpdateAlarmRequestDeviceInfo) SetId(v string) *UpdateAlarmRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *UpdateAlarmRequestDeviceInfo) SetIdType(v string) *UpdateAlarmRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *UpdateAlarmRequestDeviceInfo) SetOrganizationId(v string) *UpdateAlarmRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *UpdateAlarmRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateAlarmRequestPayload struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// This parameter is required.
	MusicInfo *UpdateAlarmRequestPayloadMusicInfo `json:"MusicInfo,omitempty" xml:"MusicInfo,omitempty" type:"Struct"`
	// This parameter is required.
	ScheduleInfo *UpdateAlarmRequestPayloadScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	// example:
	//
	// 40
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s UpdateAlarmRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmRequestPayload) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestPayload) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *UpdateAlarmRequestPayload) GetMusicInfo() *UpdateAlarmRequestPayloadMusicInfo {
	return s.MusicInfo
}

func (s *UpdateAlarmRequestPayload) GetScheduleInfo() *UpdateAlarmRequestPayloadScheduleInfo {
	return s.ScheduleInfo
}

func (s *UpdateAlarmRequestPayload) GetVolume() *int32 {
	return s.Volume
}

func (s *UpdateAlarmRequestPayload) SetAlarmId(v int64) *UpdateAlarmRequestPayload {
	s.AlarmId = &v
	return s
}

func (s *UpdateAlarmRequestPayload) SetMusicInfo(v *UpdateAlarmRequestPayloadMusicInfo) *UpdateAlarmRequestPayload {
	s.MusicInfo = v
	return s
}

func (s *UpdateAlarmRequestPayload) SetScheduleInfo(v *UpdateAlarmRequestPayloadScheduleInfo) *UpdateAlarmRequestPayload {
	s.ScheduleInfo = v
	return s
}

func (s *UpdateAlarmRequestPayload) SetVolume(v int32) *UpdateAlarmRequestPayload {
	s.Volume = &v
	return s
}

func (s *UpdateAlarmRequestPayload) Validate() error {
	return dara.Validate(s)
}

type UpdateAlarmRequestPayloadMusicInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	MusicId *int64 `json:"MusicId,omitempty" xml:"MusicId,omitempty"`
	// This parameter is required.
	MusicName *string `json:"MusicName,omitempty" xml:"MusicName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12
	MusicType *int64 `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	// This parameter is required.
	MusicTypeName *string `json:"MusicTypeName,omitempty" xml:"MusicTypeName,omitempty"`
	// example:
	//
	// http://music-url.mp3
	MusicUrl *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
}

func (s UpdateAlarmRequestPayloadMusicInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmRequestPayloadMusicInfo) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestPayloadMusicInfo) GetMusicId() *int64 {
	return s.MusicId
}

func (s *UpdateAlarmRequestPayloadMusicInfo) GetMusicName() *string {
	return s.MusicName
}

func (s *UpdateAlarmRequestPayloadMusicInfo) GetMusicType() *int64 {
	return s.MusicType
}

func (s *UpdateAlarmRequestPayloadMusicInfo) GetMusicTypeName() *string {
	return s.MusicTypeName
}

func (s *UpdateAlarmRequestPayloadMusicInfo) GetMusicUrl() *string {
	return s.MusicUrl
}

func (s *UpdateAlarmRequestPayloadMusicInfo) SetMusicId(v int64) *UpdateAlarmRequestPayloadMusicInfo {
	s.MusicId = &v
	return s
}

func (s *UpdateAlarmRequestPayloadMusicInfo) SetMusicName(v string) *UpdateAlarmRequestPayloadMusicInfo {
	s.MusicName = &v
	return s
}

func (s *UpdateAlarmRequestPayloadMusicInfo) SetMusicType(v int64) *UpdateAlarmRequestPayloadMusicInfo {
	s.MusicType = &v
	return s
}

func (s *UpdateAlarmRequestPayloadMusicInfo) SetMusicTypeName(v string) *UpdateAlarmRequestPayloadMusicInfo {
	s.MusicTypeName = &v
	return s
}

func (s *UpdateAlarmRequestPayloadMusicInfo) SetMusicUrl(v string) *UpdateAlarmRequestPayloadMusicInfo {
	s.MusicUrl = &v
	return s
}

func (s *UpdateAlarmRequestPayloadMusicInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateAlarmRequestPayloadScheduleInfo struct {
	Once                *UpdateAlarmRequestPayloadScheduleInfoOnce                `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	StatutoryWorkingDay *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay `json:"StatutoryWorkingDay,omitempty" xml:"StatutoryWorkingDay,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// WEEKLY
	Type   *string                                      `json:"Type,omitempty" xml:"Type,omitempty"`
	Weekly *UpdateAlarmRequestPayloadScheduleInfoWeekly `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s UpdateAlarmRequestPayloadScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmRequestPayloadScheduleInfo) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestPayloadScheduleInfo) GetOnce() *UpdateAlarmRequestPayloadScheduleInfoOnce {
	return s.Once
}

func (s *UpdateAlarmRequestPayloadScheduleInfo) GetStatutoryWorkingDay() *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay {
	return s.StatutoryWorkingDay
}

func (s *UpdateAlarmRequestPayloadScheduleInfo) GetType() *string {
	return s.Type
}

func (s *UpdateAlarmRequestPayloadScheduleInfo) GetWeekly() *UpdateAlarmRequestPayloadScheduleInfoWeekly {
	return s.Weekly
}

func (s *UpdateAlarmRequestPayloadScheduleInfo) SetOnce(v *UpdateAlarmRequestPayloadScheduleInfoOnce) *UpdateAlarmRequestPayloadScheduleInfo {
	s.Once = v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfo) SetStatutoryWorkingDay(v *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) *UpdateAlarmRequestPayloadScheduleInfo {
	s.StatutoryWorkingDay = v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfo) SetType(v string) *UpdateAlarmRequestPayloadScheduleInfo {
	s.Type = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfo) SetWeekly(v *UpdateAlarmRequestPayloadScheduleInfoWeekly) *UpdateAlarmRequestPayloadScheduleInfo {
	s.Weekly = v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateAlarmRequestPayloadScheduleInfoOnce struct {
	// example:
	//
	// 1
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
	// 8
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// example:
	//
	// 2022
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s UpdateAlarmRequestPayloadScheduleInfoOnce) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmRequestPayloadScheduleInfoOnce) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) GetDay() *int32 {
	return s.Day
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) GetHour() *int32 {
	return s.Hour
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) GetMinute() *int32 {
	return s.Minute
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) GetMonth() *int32 {
	return s.Month
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) GetYear() *int32 {
	return s.Year
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) SetDay(v int32) *UpdateAlarmRequestPayloadScheduleInfoOnce {
	s.Day = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) SetHour(v int32) *UpdateAlarmRequestPayloadScheduleInfoOnce {
	s.Hour = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) SetMinute(v int32) *UpdateAlarmRequestPayloadScheduleInfoOnce {
	s.Minute = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) SetMonth(v int32) *UpdateAlarmRequestPayloadScheduleInfoOnce {
	s.Month = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) SetYear(v int32) *UpdateAlarmRequestPayloadScheduleInfoOnce {
	s.Year = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoOnce) Validate() error {
	return dara.Validate(s)
}

type UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay struct {
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 0
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) GetHour() *int32 {
	return s.Hour
}

func (s *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) GetMinute() *int32 {
	return s.Minute
}

func (s *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) SetHour(v int32) *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay {
	s.Hour = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) SetMinute(v int32) *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay {
	s.Minute = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) Validate() error {
	return dara.Validate(s)
}

type UpdateAlarmRequestPayloadScheduleInfoWeekly struct {
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

func (s UpdateAlarmRequestPayloadScheduleInfoWeekly) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmRequestPayloadScheduleInfoWeekly) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestPayloadScheduleInfoWeekly) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *UpdateAlarmRequestPayloadScheduleInfoWeekly) GetHour() *int32 {
	return s.Hour
}

func (s *UpdateAlarmRequestPayloadScheduleInfoWeekly) GetMinute() *int32 {
	return s.Minute
}

func (s *UpdateAlarmRequestPayloadScheduleInfoWeekly) SetDaysOfWeek(v []*int32) *UpdateAlarmRequestPayloadScheduleInfoWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoWeekly) SetHour(v int32) *UpdateAlarmRequestPayloadScheduleInfoWeekly {
	s.Hour = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoWeekly) SetMinute(v int32) *UpdateAlarmRequestPayloadScheduleInfoWeekly {
	s.Minute = &v
	return s
}

func (s *UpdateAlarmRequestPayloadScheduleInfoWeekly) Validate() error {
	return dara.Validate(s)
}

type UpdateAlarmRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UpdateAlarmRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmRequestUserInfo) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *UpdateAlarmRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *UpdateAlarmRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *UpdateAlarmRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *UpdateAlarmRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateAlarmRequestUserInfo) SetEncodeKey(v string) *UpdateAlarmRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *UpdateAlarmRequestUserInfo) SetEncodeType(v string) *UpdateAlarmRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *UpdateAlarmRequestUserInfo) SetId(v string) *UpdateAlarmRequestUserInfo {
	s.Id = &v
	return s
}

func (s *UpdateAlarmRequestUserInfo) SetIdType(v string) *UpdateAlarmRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *UpdateAlarmRequestUserInfo) SetOrganizationId(v string) *UpdateAlarmRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *UpdateAlarmRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
