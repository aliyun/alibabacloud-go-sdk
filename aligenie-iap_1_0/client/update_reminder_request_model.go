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
	// This parameter is required.
	DeviceInfo *UpdateReminderRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Payload *UpdateReminderRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type UpdateReminderRequestDeviceInfo struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 更新提醒内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20***34
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	IsDebug *bool `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
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
	return dara.Validate(s)
}

type UpdateReminderRequestPayloadRecurrenceRule struct {
	// example:
	//
	// 25
	Day         *int32   `json:"Day,omitempty" xml:"Day,omitempty"`
	DaysOfMonth []*int32 `json:"DaysOfMonth,omitempty" xml:"DaysOfMonth,omitempty" type:"Repeated"`
	DaysOfWeek  []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1635134700000
	EndDateTime *int64 `json:"EndDateTime,omitempty" xml:"EndDateTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ONCE
	Freq *string `json:"Freq,omitempty" xml:"Freq,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 3
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	// example:
	//
	// 10
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// example:
	//
	// 3
	Second *int32 `json:"Second,omitempty" xml:"Second,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1635134400000
	StartDateTime *int64 `json:"StartDateTime,omitempty" xml:"StartDateTime,omitempty"`
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
