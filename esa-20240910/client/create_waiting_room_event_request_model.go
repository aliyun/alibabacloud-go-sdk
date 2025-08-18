// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWaitingRoomEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomPageHtml(v string) *CreateWaitingRoomEventRequest
	GetCustomPageHtml() *string
	SetDescription(v string) *CreateWaitingRoomEventRequest
	GetDescription() *string
	SetDisableSessionRenewalEnable(v string) *CreateWaitingRoomEventRequest
	GetDisableSessionRenewalEnable() *string
	SetEnable(v string) *CreateWaitingRoomEventRequest
	GetEnable() *string
	SetEndTime(v string) *CreateWaitingRoomEventRequest
	GetEndTime() *string
	SetJsonResponseEnable(v string) *CreateWaitingRoomEventRequest
	GetJsonResponseEnable() *string
	SetLanguage(v string) *CreateWaitingRoomEventRequest
	GetLanguage() *string
	SetName(v string) *CreateWaitingRoomEventRequest
	GetName() *string
	SetNewUsersPerMinute(v string) *CreateWaitingRoomEventRequest
	GetNewUsersPerMinute() *string
	SetPreQueueEnable(v string) *CreateWaitingRoomEventRequest
	GetPreQueueEnable() *string
	SetPreQueueStartTime(v string) *CreateWaitingRoomEventRequest
	GetPreQueueStartTime() *string
	SetQueuingMethod(v string) *CreateWaitingRoomEventRequest
	GetQueuingMethod() *string
	SetQueuingStatusCode(v string) *CreateWaitingRoomEventRequest
	GetQueuingStatusCode() *string
	SetRandomPreQueueEnable(v string) *CreateWaitingRoomEventRequest
	GetRandomPreQueueEnable() *string
	SetSessionDuration(v string) *CreateWaitingRoomEventRequest
	GetSessionDuration() *string
	SetSiteId(v int64) *CreateWaitingRoomEventRequest
	GetSiteId() *int64
	SetStartTime(v string) *CreateWaitingRoomEventRequest
	GetStartTime() *string
	SetTotalActiveUsers(v string) *CreateWaitingRoomEventRequest
	GetTotalActiveUsers() *string
	SetWaitingRoomId(v string) *CreateWaitingRoomEventRequest
	GetWaitingRoomId() *string
	SetWaitingRoomType(v string) *CreateWaitingRoomEventRequest
	GetWaitingRoomType() *string
}

type CreateWaitingRoomEventRequest struct {
	// The content of the custom waiting room page. You must specify this parameter if you set WaitingRoomType to custom. The content must be Base64-encoded.
	//
	// example:
	//
	// Hello%20world!
	CustomPageHtml *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	// The description of the waiting room.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to disable session renewal. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	DisableSessionRenewalEnable *string `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	// Specifies whether to enable the waiting room. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The end time of the event. This value is a UNIX timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1719849600
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to enable JSON response. If you set this parameter to on, a JSON body is returned for requests to the waiting room with the header Accept: application/json. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	JsonResponseEnable *string `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	// The language of the waiting room page. You must specify this parameter if you set WaitingRoomType to default. Valid values:
	//
	// 	- enus: English.
	//
	// 	- zhcn: Simplified Chinese.
	//
	// 	- zhhk: Traditional Chinese.
	//
	// example:
	//
	// zhcn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the waiting room event.
	//
	// This parameter is required.
	//
	// example:
	//
	// waitingroom_example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The maximum number of new users per minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	NewUsersPerMinute *string `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	// Specifies whether to enable pre-queuing.
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	PreQueueEnable *string `json:"PreQueueEnable,omitempty" xml:"PreQueueEnable,omitempty"`
	// The start time for pre-queuing.
	//
	// example:
	//
	// 1719763200
	PreQueueStartTime *string `json:"PreQueueStartTime,omitempty" xml:"PreQueueStartTime,omitempty"`
	// The queuing method. Valid values:
	//
	// 	- random: Users gain access to the origin randomly, regardless of the arrival time.
	//
	// 	- fifo: Users gain access to the origin in order of arrival.
	//
	// 	- passthrough: Users pass through the waiting room and go straight to the origin.
	//
	// 	- reject-all: Users are blocked from reaching the origin.
	//
	// This parameter is required.
	//
	// example:
	//
	// random
	QueuingMethod *string `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	// The HTTP status code to return while a user is in the queue. Valid values:
	//
	// 	- 200
	//
	// 	- 202
	//
	// 	- 429
	//
	// This parameter is required.
	//
	// example:
	//
	// 202
	QueuingStatusCode *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	// Specifies whether to enable random queuing.
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	RandomPreQueueEnable *string `json:"RandomPreQueueEnable,omitempty" xml:"RandomPreQueueEnable,omitempty"`
	// The maximum duration for which a session remains valid after a user leaves the origin. Unit: minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	SessionDuration *string `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The start time of the event. This value is a UNIX timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1719763200
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The maximum number of active users.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// The ID of the waiting room, which can be obtained by calling the [ListWaitingRooms](https://help.aliyun.com/document_detail/2850279.html) operation.
	//
	// example:
	//
	// 6a51d5bc6460887abd1291dc7d4db28b
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
	// The type of the waiting room. Valid values:
	//
	// 	- default
	//
	// 	- custom
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	WaitingRoomType *string `json:"WaitingRoomType,omitempty" xml:"WaitingRoomType,omitempty"`
}

func (s CreateWaitingRoomEventRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWaitingRoomEventRequest) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomEventRequest) GetCustomPageHtml() *string {
	return s.CustomPageHtml
}

func (s *CreateWaitingRoomEventRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWaitingRoomEventRequest) GetDisableSessionRenewalEnable() *string {
	return s.DisableSessionRenewalEnable
}

func (s *CreateWaitingRoomEventRequest) GetEnable() *string {
	return s.Enable
}

func (s *CreateWaitingRoomEventRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateWaitingRoomEventRequest) GetJsonResponseEnable() *string {
	return s.JsonResponseEnable
}

func (s *CreateWaitingRoomEventRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateWaitingRoomEventRequest) GetName() *string {
	return s.Name
}

func (s *CreateWaitingRoomEventRequest) GetNewUsersPerMinute() *string {
	return s.NewUsersPerMinute
}

func (s *CreateWaitingRoomEventRequest) GetPreQueueEnable() *string {
	return s.PreQueueEnable
}

func (s *CreateWaitingRoomEventRequest) GetPreQueueStartTime() *string {
	return s.PreQueueStartTime
}

func (s *CreateWaitingRoomEventRequest) GetQueuingMethod() *string {
	return s.QueuingMethod
}

func (s *CreateWaitingRoomEventRequest) GetQueuingStatusCode() *string {
	return s.QueuingStatusCode
}

func (s *CreateWaitingRoomEventRequest) GetRandomPreQueueEnable() *string {
	return s.RandomPreQueueEnable
}

func (s *CreateWaitingRoomEventRequest) GetSessionDuration() *string {
	return s.SessionDuration
}

func (s *CreateWaitingRoomEventRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateWaitingRoomEventRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateWaitingRoomEventRequest) GetTotalActiveUsers() *string {
	return s.TotalActiveUsers
}

func (s *CreateWaitingRoomEventRequest) GetWaitingRoomId() *string {
	return s.WaitingRoomId
}

func (s *CreateWaitingRoomEventRequest) GetWaitingRoomType() *string {
	return s.WaitingRoomType
}

func (s *CreateWaitingRoomEventRequest) SetCustomPageHtml(v string) *CreateWaitingRoomEventRequest {
	s.CustomPageHtml = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetDescription(v string) *CreateWaitingRoomEventRequest {
	s.Description = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetDisableSessionRenewalEnable(v string) *CreateWaitingRoomEventRequest {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetEnable(v string) *CreateWaitingRoomEventRequest {
	s.Enable = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetEndTime(v string) *CreateWaitingRoomEventRequest {
	s.EndTime = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetJsonResponseEnable(v string) *CreateWaitingRoomEventRequest {
	s.JsonResponseEnable = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetLanguage(v string) *CreateWaitingRoomEventRequest {
	s.Language = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetName(v string) *CreateWaitingRoomEventRequest {
	s.Name = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetNewUsersPerMinute(v string) *CreateWaitingRoomEventRequest {
	s.NewUsersPerMinute = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetPreQueueEnable(v string) *CreateWaitingRoomEventRequest {
	s.PreQueueEnable = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetPreQueueStartTime(v string) *CreateWaitingRoomEventRequest {
	s.PreQueueStartTime = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetQueuingMethod(v string) *CreateWaitingRoomEventRequest {
	s.QueuingMethod = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetQueuingStatusCode(v string) *CreateWaitingRoomEventRequest {
	s.QueuingStatusCode = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetRandomPreQueueEnable(v string) *CreateWaitingRoomEventRequest {
	s.RandomPreQueueEnable = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetSessionDuration(v string) *CreateWaitingRoomEventRequest {
	s.SessionDuration = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetSiteId(v int64) *CreateWaitingRoomEventRequest {
	s.SiteId = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetStartTime(v string) *CreateWaitingRoomEventRequest {
	s.StartTime = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetTotalActiveUsers(v string) *CreateWaitingRoomEventRequest {
	s.TotalActiveUsers = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetWaitingRoomId(v string) *CreateWaitingRoomEventRequest {
	s.WaitingRoomId = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) SetWaitingRoomType(v string) *CreateWaitingRoomEventRequest {
	s.WaitingRoomType = &v
	return s
}

func (s *CreateWaitingRoomEventRequest) Validate() error {
	return dara.Validate(s)
}
