// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaitingRoomShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCookieName(v string) *UpdateWaitingRoomShrinkRequest
	GetCookieName() *string
	SetCustomPageHtml(v string) *UpdateWaitingRoomShrinkRequest
	GetCustomPageHtml() *string
	SetDescription(v string) *UpdateWaitingRoomShrinkRequest
	GetDescription() *string
	SetDisableSessionRenewalEnable(v string) *UpdateWaitingRoomShrinkRequest
	GetDisableSessionRenewalEnable() *string
	SetEnable(v string) *UpdateWaitingRoomShrinkRequest
	GetEnable() *string
	SetHostNameAndPathShrink(v string) *UpdateWaitingRoomShrinkRequest
	GetHostNameAndPathShrink() *string
	SetJsonResponseEnable(v string) *UpdateWaitingRoomShrinkRequest
	GetJsonResponseEnable() *string
	SetLanguage(v string) *UpdateWaitingRoomShrinkRequest
	GetLanguage() *string
	SetName(v string) *UpdateWaitingRoomShrinkRequest
	GetName() *string
	SetNewUsersPerMinute(v string) *UpdateWaitingRoomShrinkRequest
	GetNewUsersPerMinute() *string
	SetQueueAllEnable(v string) *UpdateWaitingRoomShrinkRequest
	GetQueueAllEnable() *string
	SetQueuingMethod(v string) *UpdateWaitingRoomShrinkRequest
	GetQueuingMethod() *string
	SetQueuingStatusCode(v string) *UpdateWaitingRoomShrinkRequest
	GetQueuingStatusCode() *string
	SetSessionDuration(v string) *UpdateWaitingRoomShrinkRequest
	GetSessionDuration() *string
	SetSiteId(v int64) *UpdateWaitingRoomShrinkRequest
	GetSiteId() *int64
	SetTotalActiveUsers(v string) *UpdateWaitingRoomShrinkRequest
	GetTotalActiveUsers() *string
	SetWaitingRoomId(v string) *UpdateWaitingRoomShrinkRequest
	GetWaitingRoomId() *string
	SetWaitingRoomType(v string) *UpdateWaitingRoomShrinkRequest
	GetWaitingRoomType() *string
}

type UpdateWaitingRoomShrinkRequest struct {
	// The name of the custom cookie.
	//
	// example:
	//
	// __aliwaitingroom_example
	CookieName *string `json:"CookieName,omitempty" xml:"CookieName,omitempty"`
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
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The hostname and path.
	HostNameAndPathShrink *string `json:"HostNameAndPath,omitempty" xml:"HostNameAndPath,omitempty"`
	// Specifies whether to enable JSON response. If JSON response is enabled, a JSON body is returned for requests to the waiting room with the header Accept: application/json. Valid values:
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
	// The name of the waiting room.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The maximum number of new users per minute.
	//
	// example:
	//
	// 200
	NewUsersPerMinute *string `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	// Specifies whether to queue all requests. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	QueueAllEnable *string `json:"QueueAllEnable,omitempty" xml:"QueueAllEnable,omitempty"`
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
	// example:
	//
	// 200
	QueuingStatusCode *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	// The maximum duration for which a session remains valid after a user leaves the origin. Unit: minutes.
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
	// 7096621098****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The maximum number of active users.
	//
	// example:
	//
	// 300
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// The ID of the waiting room, which can be obtained by calling the [ListWaitingRooms](https://help.aliyun.com/document_detail/2850279.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6a51d5bc6460887abd129****
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
	// The type of the waiting room. Valid values:
	//
	// 	- default
	//
	// 	- custom
	//
	// example:
	//
	// default
	WaitingRoomType *string `json:"WaitingRoomType,omitempty" xml:"WaitingRoomType,omitempty"`
}

func (s UpdateWaitingRoomShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaitingRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomShrinkRequest) GetCookieName() *string {
	return s.CookieName
}

func (s *UpdateWaitingRoomShrinkRequest) GetCustomPageHtml() *string {
	return s.CustomPageHtml
}

func (s *UpdateWaitingRoomShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWaitingRoomShrinkRequest) GetDisableSessionRenewalEnable() *string {
	return s.DisableSessionRenewalEnable
}

func (s *UpdateWaitingRoomShrinkRequest) GetEnable() *string {
	return s.Enable
}

func (s *UpdateWaitingRoomShrinkRequest) GetHostNameAndPathShrink() *string {
	return s.HostNameAndPathShrink
}

func (s *UpdateWaitingRoomShrinkRequest) GetJsonResponseEnable() *string {
	return s.JsonResponseEnable
}

func (s *UpdateWaitingRoomShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *UpdateWaitingRoomShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWaitingRoomShrinkRequest) GetNewUsersPerMinute() *string {
	return s.NewUsersPerMinute
}

func (s *UpdateWaitingRoomShrinkRequest) GetQueueAllEnable() *string {
	return s.QueueAllEnable
}

func (s *UpdateWaitingRoomShrinkRequest) GetQueuingMethod() *string {
	return s.QueuingMethod
}

func (s *UpdateWaitingRoomShrinkRequest) GetQueuingStatusCode() *string {
	return s.QueuingStatusCode
}

func (s *UpdateWaitingRoomShrinkRequest) GetSessionDuration() *string {
	return s.SessionDuration
}

func (s *UpdateWaitingRoomShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateWaitingRoomShrinkRequest) GetTotalActiveUsers() *string {
	return s.TotalActiveUsers
}

func (s *UpdateWaitingRoomShrinkRequest) GetWaitingRoomId() *string {
	return s.WaitingRoomId
}

func (s *UpdateWaitingRoomShrinkRequest) GetWaitingRoomType() *string {
	return s.WaitingRoomType
}

func (s *UpdateWaitingRoomShrinkRequest) SetCookieName(v string) *UpdateWaitingRoomShrinkRequest {
	s.CookieName = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetCustomPageHtml(v string) *UpdateWaitingRoomShrinkRequest {
	s.CustomPageHtml = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetDescription(v string) *UpdateWaitingRoomShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetDisableSessionRenewalEnable(v string) *UpdateWaitingRoomShrinkRequest {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetEnable(v string) *UpdateWaitingRoomShrinkRequest {
	s.Enable = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetHostNameAndPathShrink(v string) *UpdateWaitingRoomShrinkRequest {
	s.HostNameAndPathShrink = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetJsonResponseEnable(v string) *UpdateWaitingRoomShrinkRequest {
	s.JsonResponseEnable = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetLanguage(v string) *UpdateWaitingRoomShrinkRequest {
	s.Language = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetName(v string) *UpdateWaitingRoomShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetNewUsersPerMinute(v string) *UpdateWaitingRoomShrinkRequest {
	s.NewUsersPerMinute = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetQueueAllEnable(v string) *UpdateWaitingRoomShrinkRequest {
	s.QueueAllEnable = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetQueuingMethod(v string) *UpdateWaitingRoomShrinkRequest {
	s.QueuingMethod = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetQueuingStatusCode(v string) *UpdateWaitingRoomShrinkRequest {
	s.QueuingStatusCode = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetSessionDuration(v string) *UpdateWaitingRoomShrinkRequest {
	s.SessionDuration = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetSiteId(v int64) *UpdateWaitingRoomShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetTotalActiveUsers(v string) *UpdateWaitingRoomShrinkRequest {
	s.TotalActiveUsers = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetWaitingRoomId(v string) *UpdateWaitingRoomShrinkRequest {
	s.WaitingRoomId = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) SetWaitingRoomType(v string) *UpdateWaitingRoomShrinkRequest {
	s.WaitingRoomType = &v
	return s
}

func (s *UpdateWaitingRoomShrinkRequest) Validate() error {
	return dara.Validate(s)
}
