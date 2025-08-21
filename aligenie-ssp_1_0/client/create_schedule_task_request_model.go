// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *CreateScheduleTaskRequestDeviceInfo) *CreateScheduleTaskRequest
	GetDeviceInfo() *CreateScheduleTaskRequestDeviceInfo
	SetPayload(v *CreateScheduleTaskRequestPayload) *CreateScheduleTaskRequest
	GetPayload() *CreateScheduleTaskRequestPayload
	SetUserInfo(v *CreateScheduleTaskRequestUserInfo) *CreateScheduleTaskRequest
	GetUserInfo() *CreateScheduleTaskRequestUserInfo
}

type CreateScheduleTaskRequest struct {
	// This parameter is required.
	DeviceInfo *CreateScheduleTaskRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Payload *CreateScheduleTaskRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *CreateScheduleTaskRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s CreateScheduleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequest) GetDeviceInfo() *CreateScheduleTaskRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *CreateScheduleTaskRequest) GetPayload() *CreateScheduleTaskRequestPayload {
	return s.Payload
}

func (s *CreateScheduleTaskRequest) GetUserInfo() *CreateScheduleTaskRequestUserInfo {
	return s.UserInfo
}

func (s *CreateScheduleTaskRequest) SetDeviceInfo(v *CreateScheduleTaskRequestDeviceInfo) *CreateScheduleTaskRequest {
	s.DeviceInfo = v
	return s
}

func (s *CreateScheduleTaskRequest) SetPayload(v *CreateScheduleTaskRequestPayload) *CreateScheduleTaskRequest {
	s.Payload = v
	return s
}

func (s *CreateScheduleTaskRequest) SetUserInfo(v *CreateScheduleTaskRequestUserInfo) *CreateScheduleTaskRequest {
	s.UserInfo = v
	return s
}

func (s *CreateScheduleTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreateScheduleTaskRequestDeviceInfo struct {
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateScheduleTaskRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleTaskRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CreateScheduleTaskRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CreateScheduleTaskRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *CreateScheduleTaskRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *CreateScheduleTaskRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateScheduleTaskRequestDeviceInfo) SetEncodeKey(v string) *CreateScheduleTaskRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreateScheduleTaskRequestDeviceInfo) SetEncodeType(v string) *CreateScheduleTaskRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *CreateScheduleTaskRequestDeviceInfo) SetId(v string) *CreateScheduleTaskRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *CreateScheduleTaskRequestDeviceInfo) SetIdType(v string) *CreateScheduleTaskRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *CreateScheduleTaskRequestDeviceInfo) SetOrganizationId(v string) *CreateScheduleTaskRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *CreateScheduleTaskRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type CreateScheduleTaskRequestPayload struct {
	// This parameter is required.
	ActionDTOs []*CreateScheduleTaskRequestPayloadActionDTOs `json:"ActionDTOs,omitempty" xml:"ActionDTOs,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	IdempotentId *string `json:"IdempotentId,omitempty" xml:"IdempotentId,omitempty"`
	// This parameter is required.
	ScheduleDTO *CreateScheduleTaskRequestPayloadScheduleDTO `json:"ScheduleDTO,omitempty" xml:"ScheduleDTO,omitempty" type:"Struct"`
}

func (s CreateScheduleTaskRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleTaskRequestPayload) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestPayload) GetActionDTOs() []*CreateScheduleTaskRequestPayloadActionDTOs {
	return s.ActionDTOs
}

func (s *CreateScheduleTaskRequestPayload) GetIdempotentId() *string {
	return s.IdempotentId
}

func (s *CreateScheduleTaskRequestPayload) GetScheduleDTO() *CreateScheduleTaskRequestPayloadScheduleDTO {
	return s.ScheduleDTO
}

func (s *CreateScheduleTaskRequestPayload) SetActionDTOs(v []*CreateScheduleTaskRequestPayloadActionDTOs) *CreateScheduleTaskRequestPayload {
	s.ActionDTOs = v
	return s
}

func (s *CreateScheduleTaskRequestPayload) SetIdempotentId(v string) *CreateScheduleTaskRequestPayload {
	s.IdempotentId = &v
	return s
}

func (s *CreateScheduleTaskRequestPayload) SetScheduleDTO(v *CreateScheduleTaskRequestPayloadScheduleDTO) *CreateScheduleTaskRequestPayload {
	s.ScheduleDTO = v
	return s
}

func (s *CreateScheduleTaskRequestPayload) Validate() error {
	return dara.Validate(s)
}

type CreateScheduleTaskRequestPayloadActionDTOs struct {
	// example:
	//
	// {"k1":"v1","k2":{"key":1}}
	CustomAction map[string]interface{} `json:"customAction,omitempty" xml:"customAction,omitempty"`
}

func (s CreateScheduleTaskRequestPayloadActionDTOs) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleTaskRequestPayloadActionDTOs) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestPayloadActionDTOs) GetCustomAction() map[string]interface{} {
	return s.CustomAction
}

func (s *CreateScheduleTaskRequestPayloadActionDTOs) SetCustomAction(v map[string]interface{}) *CreateScheduleTaskRequestPayloadActionDTOs {
	s.CustomAction = v
	return s
}

func (s *CreateScheduleTaskRequestPayloadActionDTOs) Validate() error {
	return dara.Validate(s)
}

type CreateScheduleTaskRequestPayloadScheduleDTO struct {
	Once *CreateScheduleTaskRequestPayloadScheduleDTOOnce `json:"Once,omitempty" xml:"Once,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1661589255000
	ScheduleEndTime *int64 `json:"ScheduleEndTime,omitempty" xml:"ScheduleEndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1656318855000
	ScheduleStartTime *int64 `json:"ScheduleStartTime,omitempty" xml:"ScheduleStartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ONCE
	ScheduleType        *string                                                         `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	StatutoryWorkingDay *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay `json:"StatutoryWorkingDay,omitempty" xml:"StatutoryWorkingDay,omitempty" type:"Struct"`
	Weekly              *CreateScheduleTaskRequestPayloadScheduleDTOWeekly              `json:"Weekly,omitempty" xml:"Weekly,omitempty" type:"Struct"`
}

func (s CreateScheduleTaskRequestPayloadScheduleDTO) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleTaskRequestPayloadScheduleDTO) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) GetOnce() *CreateScheduleTaskRequestPayloadScheduleDTOOnce {
	return s.Once
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) GetScheduleEndTime() *int64 {
	return s.ScheduleEndTime
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) GetScheduleStartTime() *int64 {
	return s.ScheduleStartTime
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) GetStatutoryWorkingDay() *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay {
	return s.StatutoryWorkingDay
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) GetWeekly() *CreateScheduleTaskRequestPayloadScheduleDTOWeekly {
	return s.Weekly
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) SetOnce(v *CreateScheduleTaskRequestPayloadScheduleDTOOnce) *CreateScheduleTaskRequestPayloadScheduleDTO {
	s.Once = v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) SetScheduleEndTime(v int64) *CreateScheduleTaskRequestPayloadScheduleDTO {
	s.ScheduleEndTime = &v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) SetScheduleStartTime(v int64) *CreateScheduleTaskRequestPayloadScheduleDTO {
	s.ScheduleStartTime = &v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) SetScheduleType(v string) *CreateScheduleTaskRequestPayloadScheduleDTO {
	s.ScheduleType = &v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) SetStatutoryWorkingDay(v *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay) *CreateScheduleTaskRequestPayloadScheduleDTO {
	s.StatutoryWorkingDay = v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) SetWeekly(v *CreateScheduleTaskRequestPayloadScheduleDTOWeekly) *CreateScheduleTaskRequestPayloadScheduleDTO {
	s.Weekly = v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTO) Validate() error {
	return dara.Validate(s)
}

type CreateScheduleTaskRequestPayloadScheduleDTOOnce struct {
	// example:
	//
	// 26
	Day *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	// example:
	//
	// 12
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 30
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

func (s CreateScheduleTaskRequestPayloadScheduleDTOOnce) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleTaskRequestPayloadScheduleDTOOnce) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) GetDay() *int32 {
	return s.Day
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) GetHour() *int32 {
	return s.Hour
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) GetMinute() *int32 {
	return s.Minute
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) GetMonth() *int32 {
	return s.Month
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) GetYear() *int32 {
	return s.Year
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) SetDay(v int32) *CreateScheduleTaskRequestPayloadScheduleDTOOnce {
	s.Day = &v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) SetHour(v int32) *CreateScheduleTaskRequestPayloadScheduleDTOOnce {
	s.Hour = &v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) SetMinute(v int32) *CreateScheduleTaskRequestPayloadScheduleDTOOnce {
	s.Minute = &v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) SetMonth(v int32) *CreateScheduleTaskRequestPayloadScheduleDTOOnce {
	s.Month = &v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) SetYear(v int32) *CreateScheduleTaskRequestPayloadScheduleDTOOnce {
	s.Year = &v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOOnce) Validate() error {
	return dara.Validate(s)
}

type CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay struct {
	Hours   []*int32 `json:"Hours,omitempty" xml:"Hours,omitempty" type:"Repeated"`
	Minutes []*int32 `json:"Minutes,omitempty" xml:"Minutes,omitempty" type:"Repeated"`
}

func (s CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay) GetHours() []*int32 {
	return s.Hours
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay) GetMinutes() []*int32 {
	return s.Minutes
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay) SetHours(v []*int32) *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay {
	s.Hours = v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay) SetMinutes(v []*int32) *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay {
	s.Minutes = v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOStatutoryWorkingDay) Validate() error {
	return dara.Validate(s)
}

type CreateScheduleTaskRequestPayloadScheduleDTOWeekly struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	Hours      []*int32 `json:"Hours,omitempty" xml:"Hours,omitempty" type:"Repeated"`
	Minutes    []*int32 `json:"Minutes,omitempty" xml:"Minutes,omitempty" type:"Repeated"`
}

func (s CreateScheduleTaskRequestPayloadScheduleDTOWeekly) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleTaskRequestPayloadScheduleDTOWeekly) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOWeekly) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOWeekly) GetHours() []*int32 {
	return s.Hours
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOWeekly) GetMinutes() []*int32 {
	return s.Minutes
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOWeekly) SetDaysOfWeek(v []*int32) *CreateScheduleTaskRequestPayloadScheduleDTOWeekly {
	s.DaysOfWeek = v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOWeekly) SetHours(v []*int32) *CreateScheduleTaskRequestPayloadScheduleDTOWeekly {
	s.Hours = v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOWeekly) SetMinutes(v []*int32) *CreateScheduleTaskRequestPayloadScheduleDTOWeekly {
	s.Minutes = v
	return s
}

func (s *CreateScheduleTaskRequestPayloadScheduleDTOWeekly) Validate() error {
	return dara.Validate(s)
}

type CreateScheduleTaskRequestUserInfo struct {
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateScheduleTaskRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleTaskRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CreateScheduleTaskRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CreateScheduleTaskRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *CreateScheduleTaskRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *CreateScheduleTaskRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateScheduleTaskRequestUserInfo) SetEncodeKey(v string) *CreateScheduleTaskRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreateScheduleTaskRequestUserInfo) SetEncodeType(v string) *CreateScheduleTaskRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *CreateScheduleTaskRequestUserInfo) SetId(v string) *CreateScheduleTaskRequestUserInfo {
	s.Id = &v
	return s
}

func (s *CreateScheduleTaskRequestUserInfo) SetIdType(v string) *CreateScheduleTaskRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *CreateScheduleTaskRequestUserInfo) SetOrganizationId(v string) *CreateScheduleTaskRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *CreateScheduleTaskRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
