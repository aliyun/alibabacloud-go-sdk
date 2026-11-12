// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReminderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *CreateReminderRequestDeviceInfo) *CreateReminderRequest
	GetDeviceInfo() *CreateReminderRequestDeviceInfo
	SetPayload(v *CreateReminderRequestPayload) *CreateReminderRequest
	GetPayload() *CreateReminderRequestPayload
	SetUserInfo(v *CreateReminderRequestUserInfo) *CreateReminderRequest
	GetUserInfo() *CreateReminderRequestUserInfo
}

type CreateReminderRequest struct {
	// Device identification information.
	//
	// This parameter is required.
	DeviceInfo *CreateReminderRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// The input parameters for this service request.
	//
	// This parameter is required.
	Payload *CreateReminderRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// User identifier information.
	//
	// This parameter is required.
	UserInfo *CreateReminderRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s CreateReminderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReminderRequest) GoString() string {
	return s.String()
}

func (s *CreateReminderRequest) GetDeviceInfo() *CreateReminderRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *CreateReminderRequest) GetPayload() *CreateReminderRequestPayload {
	return s.Payload
}

func (s *CreateReminderRequest) GetUserInfo() *CreateReminderRequestUserInfo {
	return s.UserInfo
}

func (s *CreateReminderRequest) SetDeviceInfo(v *CreateReminderRequestDeviceInfo) *CreateReminderRequest {
	s.DeviceInfo = v
	return s
}

func (s *CreateReminderRequest) SetPayload(v *CreateReminderRequestPayload) *CreateReminderRequest {
	s.Payload = v
	return s
}

func (s *CreateReminderRequest) SetUserInfo(v *CreateReminderRequestUserInfo) *CreateReminderRequest {
	s.UserInfo = v
	return s
}

func (s *CreateReminderRequest) Validate() error {
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

type CreateReminderRequestDeviceInfo struct {
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
	// Encoding type. There are multiple ways to obtain the device ID for Tmall Genie, and each method corresponds to a different encoding type.
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
	// Type of the device ID:
	//
	// - OPEN_ID: Default device ID identifier.
	//
	// - UNION_ID: Organization-level device ID identifier. This value is available only after an organization has been registered on the Tmall Genie Skill Application Open Platform.
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

func (s CreateReminderRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateReminderRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *CreateReminderRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CreateReminderRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CreateReminderRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *CreateReminderRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *CreateReminderRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateReminderRequestDeviceInfo) SetEncodeKey(v string) *CreateReminderRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreateReminderRequestDeviceInfo) SetEncodeType(v string) *CreateReminderRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *CreateReminderRequestDeviceInfo) SetId(v string) *CreateReminderRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *CreateReminderRequestDeviceInfo) SetIdType(v string) *CreateReminderRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *CreateReminderRequestDeviceInfo) SetOrganizationId(v string) *CreateReminderRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *CreateReminderRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type CreateReminderRequestPayload struct {
	// The detailed content of the reminder message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 提醒内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Indicates whether this invocation is for debug purposes.
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
	RecurrenceRule *CreateReminderRequestPayloadRecurrenceRule `json:"RecurrenceRule,omitempty" xml:"RecurrenceRule,omitempty" type:"Struct"`
}

func (s CreateReminderRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s CreateReminderRequestPayload) GoString() string {
	return s.String()
}

func (s *CreateReminderRequestPayload) GetContent() *string {
	return s.Content
}

func (s *CreateReminderRequestPayload) GetIsDebug() *bool {
	return s.IsDebug
}

func (s *CreateReminderRequestPayload) GetRecurrenceRule() *CreateReminderRequestPayloadRecurrenceRule {
	return s.RecurrenceRule
}

func (s *CreateReminderRequestPayload) SetContent(v string) *CreateReminderRequestPayload {
	s.Content = &v
	return s
}

func (s *CreateReminderRequestPayload) SetIsDebug(v bool) *CreateReminderRequestPayload {
	s.IsDebug = &v
	return s
}

func (s *CreateReminderRequestPayload) SetRecurrenceRule(v *CreateReminderRequestPayloadRecurrenceRule) *CreateReminderRequestPayload {
	s.RecurrenceRule = v
	return s
}

func (s *CreateReminderRequestPayload) Validate() error {
	if s.RecurrenceRule != nil {
		if err := s.RecurrenceRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateReminderRequestPayloadRecurrenceRule struct {
	// The date on which the reminder is triggered.
	//
	// example:
	//
	// 25
	Day *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	// Configure monthly reminders, indicating the day of the month when the reminder is triggered.
	DaysOfMonth []*int32 `json:"DaysOfMonth,omitempty" xml:"DaysOfMonth,omitempty" type:"Repeated"`
	// Configure weekly reminders, indicating on which days of the week the reminder is triggered.
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	// The expiration time when the reminder ends, in milliseconds UNIX timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1635134700000
	EndDateTime *int64 `json:"EndDateTime,omitempty" xml:"EndDateTime,omitempty"`
	// Loop type:
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

func (s CreateReminderRequestPayloadRecurrenceRule) String() string {
	return dara.Prettify(s)
}

func (s CreateReminderRequestPayloadRecurrenceRule) GoString() string {
	return s.String()
}

func (s *CreateReminderRequestPayloadRecurrenceRule) GetDay() *int32 {
	return s.Day
}

func (s *CreateReminderRequestPayloadRecurrenceRule) GetDaysOfMonth() []*int32 {
	return s.DaysOfMonth
}

func (s *CreateReminderRequestPayloadRecurrenceRule) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *CreateReminderRequestPayloadRecurrenceRule) GetEndDateTime() *int64 {
	return s.EndDateTime
}

func (s *CreateReminderRequestPayloadRecurrenceRule) GetFreq() *string {
	return s.Freq
}

func (s *CreateReminderRequestPayloadRecurrenceRule) GetHour() *int32 {
	return s.Hour
}

func (s *CreateReminderRequestPayloadRecurrenceRule) GetMinute() *int32 {
	return s.Minute
}

func (s *CreateReminderRequestPayloadRecurrenceRule) GetMonth() *int32 {
	return s.Month
}

func (s *CreateReminderRequestPayloadRecurrenceRule) GetSecond() *int32 {
	return s.Second
}

func (s *CreateReminderRequestPayloadRecurrenceRule) GetStartDateTime() *int64 {
	return s.StartDateTime
}

func (s *CreateReminderRequestPayloadRecurrenceRule) GetYear() *int32 {
	return s.Year
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetDay(v int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.Day = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetDaysOfMonth(v []*int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.DaysOfMonth = v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetDaysOfWeek(v []*int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.DaysOfWeek = v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetEndDateTime(v int64) *CreateReminderRequestPayloadRecurrenceRule {
	s.EndDateTime = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetFreq(v string) *CreateReminderRequestPayloadRecurrenceRule {
	s.Freq = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetHour(v int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.Hour = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetMinute(v int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.Minute = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetMonth(v int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.Month = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetSecond(v int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.Second = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetStartDateTime(v int64) *CreateReminderRequestPayloadRecurrenceRule {
	s.StartDateTime = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetYear(v int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.Year = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) Validate() error {
	return dara.Validate(s)
}

type CreateReminderRequestUserInfo struct {
	// The value corresponding to the encoding type:
	//
	// - When the encoding type is SKILL_ID, the value is the Skill ID of the application.
	//
	// - When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the user identifier for Tmall Genie, and each method corresponds to a different encoding type.
	//
	// - PACKAGE_NAME: APK package name, used in the Android application customer flow
	//
	// - SKILL_ID: Skill ID, used in the cloud-based flow
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
	// Type of the User ID:
	//
	// - OPEN_ID: The default User ID identity.
	//
	// - UNION_ID: The User ID identity at the organization dimension. This value is available only after an organization has been requested on the Tmall Genie Skills Application Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. This parameter is required when **IdType*	- is configured as **UNION_ID**.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateReminderRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateReminderRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CreateReminderRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CreateReminderRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CreateReminderRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *CreateReminderRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *CreateReminderRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateReminderRequestUserInfo) SetEncodeKey(v string) *CreateReminderRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreateReminderRequestUserInfo) SetEncodeType(v string) *CreateReminderRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *CreateReminderRequestUserInfo) SetId(v string) *CreateReminderRequestUserInfo {
	s.Id = &v
	return s
}

func (s *CreateReminderRequestUserInfo) SetIdType(v string) *CreateReminderRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *CreateReminderRequestUserInfo) SetOrganizationId(v string) *CreateReminderRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *CreateReminderRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
