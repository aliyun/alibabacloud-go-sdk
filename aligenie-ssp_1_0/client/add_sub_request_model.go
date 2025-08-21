// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSubRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddSubscriptionInfoRequest(v *AddSubRequestAddSubscriptionInfoRequest) *AddSubRequest
	GetAddSubscriptionInfoRequest() *AddSubRequestAddSubscriptionInfoRequest
	SetDeviceInfo(v *AddSubRequestDeviceInfo) *AddSubRequest
	GetDeviceInfo() *AddSubRequestDeviceInfo
	SetUserInfo(v *AddSubRequestUserInfo) *AddSubRequest
	GetUserInfo() *AddSubRequestUserInfo
}

type AddSubRequest struct {
	AddSubscriptionInfoRequest *AddSubRequestAddSubscriptionInfoRequest `json:"AddSubscriptionInfoRequest,omitempty" xml:"AddSubscriptionInfoRequest,omitempty" type:"Struct"`
	DeviceInfo                 *AddSubRequestDeviceInfo                 `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	UserInfo                   *AddSubRequestUserInfo                   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s AddSubRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSubRequest) GoString() string {
	return s.String()
}

func (s *AddSubRequest) GetAddSubscriptionInfoRequest() *AddSubRequestAddSubscriptionInfoRequest {
	return s.AddSubscriptionInfoRequest
}

func (s *AddSubRequest) GetDeviceInfo() *AddSubRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *AddSubRequest) GetUserInfo() *AddSubRequestUserInfo {
	return s.UserInfo
}

func (s *AddSubRequest) SetAddSubscriptionInfoRequest(v *AddSubRequestAddSubscriptionInfoRequest) *AddSubRequest {
	s.AddSubscriptionInfoRequest = v
	return s
}

func (s *AddSubRequest) SetDeviceInfo(v *AddSubRequestDeviceInfo) *AddSubRequest {
	s.DeviceInfo = v
	return s
}

func (s *AddSubRequest) SetUserInfo(v *AddSubRequestUserInfo) *AddSubRequest {
	s.UserInfo = v
	return s
}

func (s *AddSubRequest) Validate() error {
	return dara.Validate(s)
}

type AddSubRequestAddSubscriptionInfoRequest struct {
	// example:
	//
	// 51999575
	AlbumId *string `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	// example:
	//
	// 2
	DailyStudyCnt *int32 `json:"DailyStudyCnt,omitempty" xml:"DailyStudyCnt,omitempty"`
	// example:
	//
	// sequence
	PlayMode     *string                                              `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
	ScheduleInfo *AddSubRequestAddSubscriptionInfoRequestScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
}

func (s AddSubRequestAddSubscriptionInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSubRequestAddSubscriptionInfoRequest) GoString() string {
	return s.String()
}

func (s *AddSubRequestAddSubscriptionInfoRequest) GetAlbumId() *string {
	return s.AlbumId
}

func (s *AddSubRequestAddSubscriptionInfoRequest) GetDailyStudyCnt() *int32 {
	return s.DailyStudyCnt
}

func (s *AddSubRequestAddSubscriptionInfoRequest) GetPlayMode() *string {
	return s.PlayMode
}

func (s *AddSubRequestAddSubscriptionInfoRequest) GetScheduleInfo() *AddSubRequestAddSubscriptionInfoRequestScheduleInfo {
	return s.ScheduleInfo
}

func (s *AddSubRequestAddSubscriptionInfoRequest) SetAlbumId(v string) *AddSubRequestAddSubscriptionInfoRequest {
	s.AlbumId = &v
	return s
}

func (s *AddSubRequestAddSubscriptionInfoRequest) SetDailyStudyCnt(v int32) *AddSubRequestAddSubscriptionInfoRequest {
	s.DailyStudyCnt = &v
	return s
}

func (s *AddSubRequestAddSubscriptionInfoRequest) SetPlayMode(v string) *AddSubRequestAddSubscriptionInfoRequest {
	s.PlayMode = &v
	return s
}

func (s *AddSubRequestAddSubscriptionInfoRequest) SetScheduleInfo(v *AddSubRequestAddSubscriptionInfoRequestScheduleInfo) *AddSubRequestAddSubscriptionInfoRequest {
	s.ScheduleInfo = v
	return s
}

func (s *AddSubRequestAddSubscriptionInfoRequest) Validate() error {
	return dara.Validate(s)
}

type AddSubRequestAddSubscriptionInfoRequestScheduleInfo struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 23
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s AddSubRequestAddSubscriptionInfoRequestScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s AddSubRequestAddSubscriptionInfoRequestScheduleInfo) GoString() string {
	return s.String()
}

func (s *AddSubRequestAddSubscriptionInfoRequestScheduleInfo) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *AddSubRequestAddSubscriptionInfoRequestScheduleInfo) GetHour() *int32 {
	return s.Hour
}

func (s *AddSubRequestAddSubscriptionInfoRequestScheduleInfo) GetMinute() *int32 {
	return s.Minute
}

func (s *AddSubRequestAddSubscriptionInfoRequestScheduleInfo) SetDaysOfWeek(v []*int32) *AddSubRequestAddSubscriptionInfoRequestScheduleInfo {
	s.DaysOfWeek = v
	return s
}

func (s *AddSubRequestAddSubscriptionInfoRequestScheduleInfo) SetHour(v int32) *AddSubRequestAddSubscriptionInfoRequestScheduleInfo {
	s.Hour = &v
	return s
}

func (s *AddSubRequestAddSubscriptionInfoRequestScheduleInfo) SetMinute(v int32) *AddSubRequestAddSubscriptionInfoRequestScheduleInfo {
	s.Minute = &v
	return s
}

func (s *AddSubRequestAddSubscriptionInfoRequestScheduleInfo) Validate() error {
	return dara.Validate(s)
}

type AddSubRequestDeviceInfo struct {
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// OPEN_ID
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s AddSubRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s AddSubRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *AddSubRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *AddSubRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *AddSubRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *AddSubRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *AddSubRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *AddSubRequestDeviceInfo) SetEncodeKey(v string) *AddSubRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *AddSubRequestDeviceInfo) SetEncodeType(v string) *AddSubRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *AddSubRequestDeviceInfo) SetId(v string) *AddSubRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *AddSubRequestDeviceInfo) SetIdType(v string) *AddSubRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *AddSubRequestDeviceInfo) SetOrganizationId(v string) *AddSubRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *AddSubRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type AddSubRequestUserInfo struct {
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// 123
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s AddSubRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s AddSubRequestUserInfo) GoString() string {
	return s.String()
}

func (s *AddSubRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *AddSubRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *AddSubRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *AddSubRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *AddSubRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *AddSubRequestUserInfo) SetEncodeKey(v string) *AddSubRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *AddSubRequestUserInfo) SetEncodeType(v string) *AddSubRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *AddSubRequestUserInfo) SetId(v string) *AddSubRequestUserInfo {
	s.Id = &v
	return s
}

func (s *AddSubRequestUserInfo) SetIdType(v string) *AddSubRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *AddSubRequestUserInfo) SetOrganizationId(v string) *AddSubRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *AddSubRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
