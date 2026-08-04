// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *CreateAlarmRequestDeviceInfo) *CreateAlarmRequest
	GetDeviceInfo() *CreateAlarmRequestDeviceInfo
	SetPayload(v *CreateAlarmRequestPayload) *CreateAlarmRequest
	GetPayload() *CreateAlarmRequestPayload
	SetUserInfo(v *CreateAlarmRequestUserInfo) *CreateAlarmRequest
	GetUserInfo() *CreateAlarmRequestUserInfo
}

type CreateAlarmRequest struct {
	// Device identity information
	//
	// This parameter is required.
	DeviceInfo *CreateAlarmRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Input parameters for the service request
	//
	// This parameter is required.
	Payload *CreateAlarmRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// User Identifier information
	//
	// This parameter is required.
	UserInfo *CreateAlarmRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s CreateAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmRequest) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequest) GetDeviceInfo() *CreateAlarmRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *CreateAlarmRequest) GetPayload() *CreateAlarmRequestPayload {
	return s.Payload
}

func (s *CreateAlarmRequest) GetUserInfo() *CreateAlarmRequestUserInfo {
	return s.UserInfo
}

func (s *CreateAlarmRequest) SetDeviceInfo(v *CreateAlarmRequestDeviceInfo) *CreateAlarmRequest {
	s.DeviceInfo = v
	return s
}

func (s *CreateAlarmRequest) SetPayload(v *CreateAlarmRequestPayload) *CreateAlarmRequest {
	s.Payload = v
	return s
}

func (s *CreateAlarmRequest) SetUserInfo(v *CreateAlarmRequestUserInfo) *CreateAlarmRequest {
	s.UserInfo = v
	return s
}

func (s *CreateAlarmRequest) Validate() error {
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAlarmRequestDeviceInfo struct {
	// The value corresponding to the encoding type. When the encoding type is SKILL_ID, the value is the Skill ID of the application; when the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device identity for Maojing, and each method corresponds to a different encoding type: PACKAGE_NAME: APK package name, used in the Android application customer link; SKILL_ID: skill ID, used in the cloud link.
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID (deviceOpenId or deviceUnionId)
	//
	// This parameter is required.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of device ID: OPEN_ID: default device ID; UNION_ID: organization-dimension device ID, available only after applying for an organization on the Maojing Skill Application Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required if IdType is UNION_ID.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateAlarmRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CreateAlarmRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CreateAlarmRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *CreateAlarmRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *CreateAlarmRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateAlarmRequestDeviceInfo) SetEncodeKey(v string) *CreateAlarmRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreateAlarmRequestDeviceInfo) SetEncodeType(v string) *CreateAlarmRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *CreateAlarmRequestDeviceInfo) SetId(v string) *CreateAlarmRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *CreateAlarmRequestDeviceInfo) SetIdType(v string) *CreateAlarmRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *CreateAlarmRequestDeviceInfo) SetOrganizationId(v string) *CreateAlarmRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *CreateAlarmRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type CreateAlarmRequestPayload struct {
	// Ringtone information
	//
	// This parameter is required.
	MusicInfo *CreateAlarmRequestPayloadMusicInfo `json:"MusicInfo,omitempty" xml:"MusicInfo,omitempty" type:"Struct"`
	// Schedule information
	//
	// This parameter is required.
	ScheduleInfo *CreateAlarmRequestPayloadScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	// Ringtone volume
	//
	// example:
	//
	// 40
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s CreateAlarmRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmRequestPayload) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestPayload) GetMusicInfo() *CreateAlarmRequestPayloadMusicInfo {
	return s.MusicInfo
}

func (s *CreateAlarmRequestPayload) GetScheduleInfo() *CreateAlarmRequestPayloadScheduleInfo {
	return s.ScheduleInfo
}

func (s *CreateAlarmRequestPayload) GetVolume() *int32 {
	return s.Volume
}

func (s *CreateAlarmRequestPayload) SetMusicInfo(v *CreateAlarmRequestPayloadMusicInfo) *CreateAlarmRequestPayload {
	s.MusicInfo = v
	return s
}

func (s *CreateAlarmRequestPayload) SetScheduleInfo(v *CreateAlarmRequestPayloadScheduleInfo) *CreateAlarmRequestPayload {
	s.ScheduleInfo = v
	return s
}

func (s *CreateAlarmRequestPayload) SetVolume(v int32) *CreateAlarmRequestPayload {
	s.Volume = &v
	return s
}

func (s *CreateAlarmRequestPayload) Validate() error {
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

type CreateAlarmRequestPayloadMusicInfo struct {
	// Ringtone ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MusicId *int64 `json:"MusicId,omitempty" xml:"MusicId,omitempty"`
	// Ringtone name
	//
	// This parameter is required.
	//
	// example:
	//
	// xx铃声
	MusicName *string `json:"MusicName,omitempty" xml:"MusicName,omitempty"`
	// Ringtone category ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MusicType *int64 `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	// Ringtone category name
	//
	// This parameter is required.
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

func (s CreateAlarmRequestPayloadMusicInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmRequestPayloadMusicInfo) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestPayloadMusicInfo) GetMusicId() *int64 {
	return s.MusicId
}

func (s *CreateAlarmRequestPayloadMusicInfo) GetMusicName() *string {
	return s.MusicName
}

func (s *CreateAlarmRequestPayloadMusicInfo) GetMusicType() *int64 {
	return s.MusicType
}

func (s *CreateAlarmRequestPayloadMusicInfo) GetMusicTypeName() *string {
	return s.MusicTypeName
}

func (s *CreateAlarmRequestPayloadMusicInfo) GetMusicUrl() *string {
	return s.MusicUrl
}

func (s *CreateAlarmRequestPayloadMusicInfo) SetMusicId(v int64) *CreateAlarmRequestPayloadMusicInfo {
	s.MusicId = &v
	return s
}

func (s *CreateAlarmRequestPayloadMusicInfo) SetMusicName(v string) *CreateAlarmRequestPayloadMusicInfo {
	s.MusicName = &v
	return s
}

func (s *CreateAlarmRequestPayloadMusicInfo) SetMusicType(v int64) *CreateAlarmRequestPayloadMusicInfo {
	s.MusicType = &v
	return s
}

func (s *CreateAlarmRequestPayloadMusicInfo) SetMusicTypeName(v string) *CreateAlarmRequestPayloadMusicInfo {
	s.MusicTypeName = &v
	return s
}

func (s *CreateAlarmRequestPayloadMusicInfo) SetMusicUrl(v string) *CreateAlarmRequestPayloadMusicInfo {
	s.MusicUrl = &v
	return s
}

func (s *CreateAlarmRequestPayloadMusicInfo) Validate() error {
	return dara.Validate(s)
}

type CreateAlarmRequestPayloadScheduleInfo struct {
	// One-time: This property is active when the loop type is ONCE.
	Once *CreateAlarmRequestPayloadScheduleInfoOnce `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	// Statutory working day: This property is active when the loop Type is STATUTORY_WORKING_DAY.
	StatutoryWorkingDay *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay `json:"StatutoryWorkingDay,omitempty" xml:"StatutoryWorkingDay,omitempty" type:"Struct"`
	// Schedule Type / Loop Type:
	//
	// ONCE -> One-time, WEEKLY -> Weekly loop, STATUTORY_WORKING_DAY -> Statutory working day
	//
	// This parameter is required.
	//
	// example:
	//
	// ONCE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Weekly Loop: This property is active when the loop Type is WEEKLY.
	Weekly *CreateAlarmRequestPayloadScheduleInfoWeekly `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s CreateAlarmRequestPayloadScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmRequestPayloadScheduleInfo) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestPayloadScheduleInfo) GetOnce() *CreateAlarmRequestPayloadScheduleInfoOnce {
	return s.Once
}

func (s *CreateAlarmRequestPayloadScheduleInfo) GetStatutoryWorkingDay() *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay {
	return s.StatutoryWorkingDay
}

func (s *CreateAlarmRequestPayloadScheduleInfo) GetType() *string {
	return s.Type
}

func (s *CreateAlarmRequestPayloadScheduleInfo) GetWeekly() *CreateAlarmRequestPayloadScheduleInfoWeekly {
	return s.Weekly
}

func (s *CreateAlarmRequestPayloadScheduleInfo) SetOnce(v *CreateAlarmRequestPayloadScheduleInfoOnce) *CreateAlarmRequestPayloadScheduleInfo {
	s.Once = v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfo) SetStatutoryWorkingDay(v *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) *CreateAlarmRequestPayloadScheduleInfo {
	s.StatutoryWorkingDay = v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfo) SetType(v string) *CreateAlarmRequestPayloadScheduleInfo {
	s.Type = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfo) SetWeekly(v *CreateAlarmRequestPayloadScheduleInfoWeekly) *CreateAlarmRequestPayloadScheduleInfo {
	s.Weekly = v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfo) Validate() error {
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

type CreateAlarmRequestPayloadScheduleInfoOnce struct {
	// Trigger Time: Day
	//
	// example:
	//
	// 1
	Day *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	// Trigger Time: Hour
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
	// Trigger time: Month
	//
	// example:
	//
	// 8
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// Trigger time: Year
	//
	// example:
	//
	// 2022
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s CreateAlarmRequestPayloadScheduleInfoOnce) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmRequestPayloadScheduleInfoOnce) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) GetDay() *int32 {
	return s.Day
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) GetHour() *int32 {
	return s.Hour
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) GetMinute() *int32 {
	return s.Minute
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) GetMonth() *int32 {
	return s.Month
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) GetYear() *int32 {
	return s.Year
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) SetDay(v int32) *CreateAlarmRequestPayloadScheduleInfoOnce {
	s.Day = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) SetHour(v int32) *CreateAlarmRequestPayloadScheduleInfoOnce {
	s.Hour = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) SetMinute(v int32) *CreateAlarmRequestPayloadScheduleInfoOnce {
	s.Minute = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) SetMonth(v int32) *CreateAlarmRequestPayloadScheduleInfoOnce {
	s.Month = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) SetYear(v int32) *CreateAlarmRequestPayloadScheduleInfoOnce {
	s.Year = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoOnce) Validate() error {
	return dara.Validate(s)
}

type CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay struct {
	// Trigger time: hour
	//
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// Trigger time: minute
	//
	// example:
	//
	// 0
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) GetHour() *int32 {
	return s.Hour
}

func (s *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) GetMinute() *int32 {
	return s.Minute
}

func (s *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) SetHour(v int32) *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay {
	s.Hour = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) SetMinute(v int32) *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay {
	s.Minute = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoStatutoryWorkingDay) Validate() error {
	return dara.Validate(s)
}

type CreateAlarmRequestPayloadScheduleInfoWeekly struct {
	// Collection of Days of the Week to Trigger
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	// Trigger time: hour
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

func (s CreateAlarmRequestPayloadScheduleInfoWeekly) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmRequestPayloadScheduleInfoWeekly) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestPayloadScheduleInfoWeekly) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *CreateAlarmRequestPayloadScheduleInfoWeekly) GetHour() *int32 {
	return s.Hour
}

func (s *CreateAlarmRequestPayloadScheduleInfoWeekly) GetMinute() *int32 {
	return s.Minute
}

func (s *CreateAlarmRequestPayloadScheduleInfoWeekly) SetDaysOfWeek(v []*int32) *CreateAlarmRequestPayloadScheduleInfoWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoWeekly) SetHour(v int32) *CreateAlarmRequestPayloadScheduleInfoWeekly {
	s.Hour = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoWeekly) SetMinute(v int32) *CreateAlarmRequestPayloadScheduleInfoWeekly {
	s.Minute = &v
	return s
}

func (s *CreateAlarmRequestPayloadScheduleInfoWeekly) Validate() error {
	return dara.Validate(s)
}

type CreateAlarmRequestUserInfo struct {
	// Value corresponding to the encoding type. When the encoding type is SKILL_ID, the value is the application\\"s Skill ID. When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding Type. There are multiple ways to obtain the User Identifier for Maojing, and each way corresponds to a different encoding Type: PACKAGE_NAME: APK package name, used for the Android application Customer link; SKILL_ID: Skill ID, used for the cloud link.
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User Identifier (userOpenId or userUnionId)
	//
	// This parameter is required.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of User ID: OPEN_ID: default User ID identifier; UNION_ID: organization-dimension User ID identifier, available only after an organization has been requested on the Maojing Skill Application Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required when IdType is UNION_ID.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateAlarmRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CreateAlarmRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CreateAlarmRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *CreateAlarmRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *CreateAlarmRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateAlarmRequestUserInfo) SetEncodeKey(v string) *CreateAlarmRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreateAlarmRequestUserInfo) SetEncodeType(v string) *CreateAlarmRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *CreateAlarmRequestUserInfo) SetId(v string) *CreateAlarmRequestUserInfo {
	s.Id = &v
	return s
}

func (s *CreateAlarmRequestUserInfo) SetIdType(v string) *CreateAlarmRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *CreateAlarmRequestUserInfo) SetOrganizationId(v string) *CreateAlarmRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *CreateAlarmRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
