// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReminderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *UpdateReminderRequestDeviceInfo) *UpdateReminderRequest
	GetDeviceInfo() *UpdateReminderRequestDeviceInfo
	SetPayload(v *UpdateReminderRequestPayload) *UpdateReminderRequest
	GetPayload() *UpdateReminderRequestPayload
	SetUserInfo(v *UpdateReminderRequestUserInfo) *UpdateReminderRequest
	GetUserInfo() *UpdateReminderRequestUserInfo
}

type UpdateReminderRequest struct {
	// Device identity information.
	//
	// This parameter is required.
	DeviceInfo *UpdateReminderRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// The input parameters for this service request.
	//
	// This parameter is required.
	Payload *UpdateReminderRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// User identifier information.
	//
	// This parameter is required.
	UserInfo *UpdateReminderRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s UpdateReminderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateReminderRequest) GoString() string {
	return s.String()
}

func (s *UpdateReminderRequest) GetDeviceInfo() *UpdateReminderRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *UpdateReminderRequest) GetPayload() *UpdateReminderRequestPayload {
	return s.Payload
}

func (s *UpdateReminderRequest) GetUserInfo() *UpdateReminderRequestUserInfo {
	return s.UserInfo
}

func (s *UpdateReminderRequest) SetDeviceInfo(v *UpdateReminderRequestDeviceInfo) *UpdateReminderRequest {
	s.DeviceInfo = v
	return s
}

func (s *UpdateReminderRequest) SetPayload(v *UpdateReminderRequestPayload) *UpdateReminderRequest {
	s.Payload = v
	return s
}

func (s *UpdateReminderRequest) SetUserInfo(v *UpdateReminderRequestUserInfo) *UpdateReminderRequest {
	s.UserInfo = v
	return s
}

func (s *UpdateReminderRequest) Validate() error {
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

type UpdateReminderRequestDeviceInfo struct {
	// The value corresponding to the encoding type.
	//
	// - When the encoding type is set to SKILL_ID, the value is the application\\"s Skill ID.
	//
	// - When the encoding type is set to PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device identity for Tmall Genie, and each method corresponds to a different encoding type.
	//
	// - PACKAGE_NAME: APK package name, used in the Android application customer flow.
	//
	// - SKILL_ID: Skill ID, used in the cloud-based flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID, set to either deviceOpenId or deviceUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of the device ID.
	//
	// - OPEN_ID: Default device ID.
	//
	// - UNION_ID: Organization-level device ID. This value is available only after an organization has been registered on the Tmall Genie skill application platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. This parameter is required when **IdType*	- is set to **UNION_ID**.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UpdateReminderRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateReminderRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *UpdateReminderRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *UpdateReminderRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *UpdateReminderRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *UpdateReminderRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *UpdateReminderRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateReminderRequestDeviceInfo) SetEncodeKey(v string) *UpdateReminderRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *UpdateReminderRequestDeviceInfo) SetEncodeType(v string) *UpdateReminderRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *UpdateReminderRequestDeviceInfo) SetId(v string) *UpdateReminderRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *UpdateReminderRequestDeviceInfo) SetIdType(v string) *UpdateReminderRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *UpdateReminderRequestDeviceInfo) SetOrganizationId(v string) *UpdateReminderRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *UpdateReminderRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateReminderRequestPayload struct {
	// The detailed content of the reminder message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 更新提醒内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the reminder to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20***34
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether this invocation is for debugging.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	IsDebug *bool `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
	// Scheduling information for the reminder.
	//
	// This parameter is required.
	RecurrenceRule *UpdateReminderRequestPayloadRecurrenceRule `json:"RecurrenceRule,omitempty" xml:"RecurrenceRule,omitempty" type:"Struct"`
}

func (s UpdateReminderRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s UpdateReminderRequestPayload) GoString() string {
	return s.String()
}

func (s *UpdateReminderRequestPayload) GetContent() *string {
	return s.Content
}

func (s *UpdateReminderRequestPayload) GetId() *int64 {
	return s.Id
}

func (s *UpdateReminderRequestPayload) GetIsDebug() *bool {
	return s.IsDebug
}

func (s *UpdateReminderRequestPayload) GetRecurrenceRule() *UpdateReminderRequestPayloadRecurrenceRule {
	return s.RecurrenceRule
}

func (s *UpdateReminderRequestPayload) SetContent(v string) *UpdateReminderRequestPayload {
	s.Content = &v
	return s
}

func (s *UpdateReminderRequestPayload) SetId(v int64) *UpdateReminderRequestPayload {
	s.Id = &v
	return s
}

func (s *UpdateReminderRequestPayload) SetIsDebug(v bool) *UpdateReminderRequestPayload {
	s.IsDebug = &v
	return s
}

func (s *UpdateReminderRequestPayload) SetRecurrenceRule(v *UpdateReminderRequestPayloadRecurrenceRule) *UpdateReminderRequestPayload {
	s.RecurrenceRule = v
	return s
}

func (s *UpdateReminderRequestPayload) Validate() error {
	if s.RecurrenceRule != nil {
		if err := s.RecurrenceRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateReminderRequestPayloadRecurrenceRule struct {
	// The date on which the reminder is triggered.
	//
	// example:
	//
	// 25
	Day *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	// Configure monthly reminders, specifying on which day(s) of each month the reminder is triggered.
	DaysOfMonth []*int32 `json:"DaysOfMonth,omitempty" xml:"DaysOfMonth,omitempty" type:"Repeated"`
	// Configure weekly reminders, indicating on which days of the week the reminder is triggered.
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	// The expiration time when the reminder ends, in milliseconds as a UNIX timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1635134700000
	EndDateTime *int64 `json:"EndDateTime,omitempty" xml:"EndDateTime,omitempty"`
	// Recurrence type:
	//
	// - ONCE: One-time
	//
	// - DAILY: Daily
	//
	// - WEEKLY: Weekly
	//
	// - MONTHLY: Monthly
	//
	// This parameter is required.
	//
	// example:
	//
	// ONCE
	Freq *string `json:"Freq,omitempty" xml:"Freq,omitempty"`
	// The hour of the time when the reminder is triggered.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// The minute of the time when the reminder is triggered.
	//
	// example:
	//
	// 3
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	// The month in which the reminder is triggered.
	//
	// example:
	//
	// 10
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// The second of the time when the reminder is triggered.
	//
	// example:
	//
	// 3
	Second *int32 `json:"Second,omitempty" xml:"Second,omitempty"`
	// The effective start time of the reminder, in milliseconds as a UNIX timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1635134400000
	StartDateTime *int64 `json:"StartDateTime,omitempty" xml:"StartDateTime,omitempty"`
	// The year of the time when the reminder is triggered.
	//
	// example:
	//
	// 2021
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s UpdateReminderRequestPayloadRecurrenceRule) String() string {
	return dara.Prettify(s)
}

func (s UpdateReminderRequestPayloadRecurrenceRule) GoString() string {
	return s.String()
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) GetDay() *int32 {
	return s.Day
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) GetDaysOfMonth() []*int32 {
	return s.DaysOfMonth
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) GetEndDateTime() *int64 {
	return s.EndDateTime
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) GetFreq() *string {
	return s.Freq
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) GetHour() *int32 {
	return s.Hour
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) GetMinute() *int32 {
	return s.Minute
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) GetMonth() *int32 {
	return s.Month
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) GetSecond() *int32 {
	return s.Second
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) GetStartDateTime() *int64 {
	return s.StartDateTime
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) GetYear() *int32 {
	return s.Year
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetDay(v int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.Day = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetDaysOfMonth(v []*int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.DaysOfMonth = v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetDaysOfWeek(v []*int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.DaysOfWeek = v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetEndDateTime(v int64) *UpdateReminderRequestPayloadRecurrenceRule {
	s.EndDateTime = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetFreq(v string) *UpdateReminderRequestPayloadRecurrenceRule {
	s.Freq = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetHour(v int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.Hour = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetMinute(v int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.Minute = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetMonth(v int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.Month = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetSecond(v int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.Second = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetStartDateTime(v int64) *UpdateReminderRequestPayloadRecurrenceRule {
	s.StartDateTime = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetYear(v int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.Year = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) Validate() error {
	return dara.Validate(s)
}

type UpdateReminderRequestUserInfo struct {
	// Value corresponding to the encoding type.
	//
	// - When the encoding type is set to SKILL_ID, the value is the application\\"s Skill ID.
	//
	// - When the encoding type is set to PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the user identifier for Maojing, and each method corresponds to a different encoding type.
	//
	// - PACKAGE_NAME: APK package name, used in the Android application customer link.
	//
	// - SKILL_ID: Skill ID, used in the cloud link.
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User identifier, set to either userOpenId or userUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of User ID
	//
	// - OPEN_ID: The default User ID identifier.
	//
	// - UNION_ID: The organization-dimension User ID identifier. This value is available only after an organization has been requested on the Tmall Genie Skill Application Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. This parameter is required when **IdType*	- is set to **UNION_ID**.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UpdateReminderRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateReminderRequestUserInfo) GoString() string {
	return s.String()
}

func (s *UpdateReminderRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *UpdateReminderRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *UpdateReminderRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *UpdateReminderRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *UpdateReminderRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateReminderRequestUserInfo) SetEncodeKey(v string) *UpdateReminderRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *UpdateReminderRequestUserInfo) SetEncodeType(v string) *UpdateReminderRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *UpdateReminderRequestUserInfo) SetId(v string) *UpdateReminderRequestUserInfo {
	s.Id = &v
	return s
}

func (s *UpdateReminderRequestUserInfo) SetIdType(v string) *UpdateReminderRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *UpdateReminderRequestUserInfo) SetOrganizationId(v string) *UpdateReminderRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *UpdateReminderRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
