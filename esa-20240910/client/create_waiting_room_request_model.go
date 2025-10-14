// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWaitingRoomRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCookieName(v string) *CreateWaitingRoomRequest
	GetCookieName() *string
	SetCustomPageHtml(v string) *CreateWaitingRoomRequest
	GetCustomPageHtml() *string
	SetDescription(v string) *CreateWaitingRoomRequest
	GetDescription() *string
	SetDisableSessionRenewalEnable(v string) *CreateWaitingRoomRequest
	GetDisableSessionRenewalEnable() *string
	SetEnable(v string) *CreateWaitingRoomRequest
	GetEnable() *string
	SetHostNameAndPath(v []*CreateWaitingRoomRequestHostNameAndPath) *CreateWaitingRoomRequest
	GetHostNameAndPath() []*CreateWaitingRoomRequestHostNameAndPath
	SetJsonResponseEnable(v string) *CreateWaitingRoomRequest
	GetJsonResponseEnable() *string
	SetLanguage(v string) *CreateWaitingRoomRequest
	GetLanguage() *string
	SetName(v string) *CreateWaitingRoomRequest
	GetName() *string
	SetNewUsersPerMinute(v string) *CreateWaitingRoomRequest
	GetNewUsersPerMinute() *string
	SetQueueAllEnable(v string) *CreateWaitingRoomRequest
	GetQueueAllEnable() *string
	SetQueuingMethod(v string) *CreateWaitingRoomRequest
	GetQueuingMethod() *string
	SetQueuingStatusCode(v string) *CreateWaitingRoomRequest
	GetQueuingStatusCode() *string
	SetSessionDuration(v string) *CreateWaitingRoomRequest
	GetSessionDuration() *string
	SetSiteId(v int64) *CreateWaitingRoomRequest
	GetSiteId() *int64
	SetTotalActiveUsers(v string) *CreateWaitingRoomRequest
	GetTotalActiveUsers() *string
	SetWaitingRoomType(v string) *CreateWaitingRoomRequest
	GetWaitingRoomType() *string
}

type CreateWaitingRoomRequest struct {
	// The name of the custom cookie.
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The hostname and path.
	//
	// This parameter is required.
	HostNameAndPath []*CreateWaitingRoomRequestHostNameAndPath `json:"HostNameAndPath,omitempty" xml:"HostNameAndPath,omitempty" type:"Repeated"`
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
	// enus
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the waiting room.
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
	// This parameter is required.
	//
	// example:
	//
	// fifo
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
	// 200
	QueuingStatusCode *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
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
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The maximum number of active users.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
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

func (s CreateWaitingRoomRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWaitingRoomRequest) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomRequest) GetCookieName() *string {
	return s.CookieName
}

func (s *CreateWaitingRoomRequest) GetCustomPageHtml() *string {
	return s.CustomPageHtml
}

func (s *CreateWaitingRoomRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWaitingRoomRequest) GetDisableSessionRenewalEnable() *string {
	return s.DisableSessionRenewalEnable
}

func (s *CreateWaitingRoomRequest) GetEnable() *string {
	return s.Enable
}

func (s *CreateWaitingRoomRequest) GetHostNameAndPath() []*CreateWaitingRoomRequestHostNameAndPath {
	return s.HostNameAndPath
}

func (s *CreateWaitingRoomRequest) GetJsonResponseEnable() *string {
	return s.JsonResponseEnable
}

func (s *CreateWaitingRoomRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateWaitingRoomRequest) GetName() *string {
	return s.Name
}

func (s *CreateWaitingRoomRequest) GetNewUsersPerMinute() *string {
	return s.NewUsersPerMinute
}

func (s *CreateWaitingRoomRequest) GetQueueAllEnable() *string {
	return s.QueueAllEnable
}

func (s *CreateWaitingRoomRequest) GetQueuingMethod() *string {
	return s.QueuingMethod
}

func (s *CreateWaitingRoomRequest) GetQueuingStatusCode() *string {
	return s.QueuingStatusCode
}

func (s *CreateWaitingRoomRequest) GetSessionDuration() *string {
	return s.SessionDuration
}

func (s *CreateWaitingRoomRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateWaitingRoomRequest) GetTotalActiveUsers() *string {
	return s.TotalActiveUsers
}

func (s *CreateWaitingRoomRequest) GetWaitingRoomType() *string {
	return s.WaitingRoomType
}

func (s *CreateWaitingRoomRequest) SetCookieName(v string) *CreateWaitingRoomRequest {
	s.CookieName = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetCustomPageHtml(v string) *CreateWaitingRoomRequest {
	s.CustomPageHtml = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetDescription(v string) *CreateWaitingRoomRequest {
	s.Description = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetDisableSessionRenewalEnable(v string) *CreateWaitingRoomRequest {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetEnable(v string) *CreateWaitingRoomRequest {
	s.Enable = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetHostNameAndPath(v []*CreateWaitingRoomRequestHostNameAndPath) *CreateWaitingRoomRequest {
	s.HostNameAndPath = v
	return s
}

func (s *CreateWaitingRoomRequest) SetJsonResponseEnable(v string) *CreateWaitingRoomRequest {
	s.JsonResponseEnable = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetLanguage(v string) *CreateWaitingRoomRequest {
	s.Language = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetName(v string) *CreateWaitingRoomRequest {
	s.Name = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetNewUsersPerMinute(v string) *CreateWaitingRoomRequest {
	s.NewUsersPerMinute = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetQueueAllEnable(v string) *CreateWaitingRoomRequest {
	s.QueueAllEnable = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetQueuingMethod(v string) *CreateWaitingRoomRequest {
	s.QueuingMethod = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetQueuingStatusCode(v string) *CreateWaitingRoomRequest {
	s.QueuingStatusCode = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetSessionDuration(v string) *CreateWaitingRoomRequest {
	s.SessionDuration = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetSiteId(v int64) *CreateWaitingRoomRequest {
	s.SiteId = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetTotalActiveUsers(v string) *CreateWaitingRoomRequest {
	s.TotalActiveUsers = &v
	return s
}

func (s *CreateWaitingRoomRequest) SetWaitingRoomType(v string) *CreateWaitingRoomRequest {
	s.WaitingRoomType = &v
	return s
}

func (s *CreateWaitingRoomRequest) Validate() error {
	if s.HostNameAndPath != nil {
		for _, item := range s.HostNameAndPath {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateWaitingRoomRequestHostNameAndPath struct {
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The path.
	//
	// This parameter is required.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The subdomain.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.
	Subdomain *string `json:"Subdomain,omitempty" xml:"Subdomain,omitempty"`
}

func (s CreateWaitingRoomRequestHostNameAndPath) String() string {
	return dara.Prettify(s)
}

func (s CreateWaitingRoomRequestHostNameAndPath) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomRequestHostNameAndPath) GetDomain() *string {
	return s.Domain
}

func (s *CreateWaitingRoomRequestHostNameAndPath) GetPath() *string {
	return s.Path
}

func (s *CreateWaitingRoomRequestHostNameAndPath) GetSubdomain() *string {
	return s.Subdomain
}

func (s *CreateWaitingRoomRequestHostNameAndPath) SetDomain(v string) *CreateWaitingRoomRequestHostNameAndPath {
	s.Domain = &v
	return s
}

func (s *CreateWaitingRoomRequestHostNameAndPath) SetPath(v string) *CreateWaitingRoomRequestHostNameAndPath {
	s.Path = &v
	return s
}

func (s *CreateWaitingRoomRequestHostNameAndPath) SetSubdomain(v string) *CreateWaitingRoomRequestHostNameAndPath {
	s.Subdomain = &v
	return s
}

func (s *CreateWaitingRoomRequestHostNameAndPath) Validate() error {
	return dara.Validate(s)
}
