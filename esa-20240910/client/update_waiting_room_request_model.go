// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaitingRoomRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCookieName(v string) *UpdateWaitingRoomRequest
	GetCookieName() *string
	SetCustomPageHtml(v string) *UpdateWaitingRoomRequest
	GetCustomPageHtml() *string
	SetDescription(v string) *UpdateWaitingRoomRequest
	GetDescription() *string
	SetDisableSessionRenewalEnable(v string) *UpdateWaitingRoomRequest
	GetDisableSessionRenewalEnable() *string
	SetEnable(v string) *UpdateWaitingRoomRequest
	GetEnable() *string
	SetHostNameAndPath(v []*UpdateWaitingRoomRequestHostNameAndPath) *UpdateWaitingRoomRequest
	GetHostNameAndPath() []*UpdateWaitingRoomRequestHostNameAndPath
	SetJsonResponseEnable(v string) *UpdateWaitingRoomRequest
	GetJsonResponseEnable() *string
	SetLanguage(v string) *UpdateWaitingRoomRequest
	GetLanguage() *string
	SetName(v string) *UpdateWaitingRoomRequest
	GetName() *string
	SetNewUsersPerMinute(v string) *UpdateWaitingRoomRequest
	GetNewUsersPerMinute() *string
	SetQueueAllEnable(v string) *UpdateWaitingRoomRequest
	GetQueueAllEnable() *string
	SetQueuingMethod(v string) *UpdateWaitingRoomRequest
	GetQueuingMethod() *string
	SetQueuingStatusCode(v string) *UpdateWaitingRoomRequest
	GetQueuingStatusCode() *string
	SetSessionDuration(v string) *UpdateWaitingRoomRequest
	GetSessionDuration() *string
	SetSiteId(v int64) *UpdateWaitingRoomRequest
	GetSiteId() *int64
	SetTotalActiveUsers(v string) *UpdateWaitingRoomRequest
	GetTotalActiveUsers() *string
	SetWaitingRoomId(v string) *UpdateWaitingRoomRequest
	GetWaitingRoomId() *string
	SetWaitingRoomType(v string) *UpdateWaitingRoomRequest
	GetWaitingRoomType() *string
}

type UpdateWaitingRoomRequest struct {
	// The custom cookie name.
	//
	// example:
	//
	// __aliwaitingroom_example
	CookieName *string `json:"CookieName,omitempty" xml:"CookieName,omitempty"`
	// The custom waiting room page content. This parameter is required when the waiting room type is set to custom. The content must be in Base64 encoding.
	//
	// example:
	//
	// SGVsbG8gd29ybGQ=
	CustomPageHtml *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	// The description of the waiting room.
	//
	// example:
	//
	// 特别活动排队页面
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to disable session renewal. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// on
	DisableSessionRenewalEnable *string `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	// The status of the waiting room. Valid values:
	//
	//  - **on**: Enabled.
	//
	//  - **off**: Disabled.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The hostname and path configurations.
	HostNameAndPath []*UpdateWaitingRoomRequestHostNameAndPath `json:"HostNameAndPath,omitempty" xml:"HostNameAndPath,omitempty" type:"Repeated"`
	// Specifies whether to enable JSON response. If enabled, requests with an Accept header containing "application/json" return JSON data. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// on
	JsonResponseEnable *string `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	// The language of the waiting room page. This parameter is required when the waiting room type is set to default. Valid values:
	//
	// - **enus**: English.
	//
	// - **zhcn**: Simplified Chinese.
	//
	// - **zhhk**: Traditional Chinese.
	//
	// example:
	//
	// zhcn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the waiting room.
	//
	// example:
	//
	// 节假日促销等候室
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of new users per minute.
	//
	// example:
	//
	// 200
	NewUsersPerMinute *string `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	// Specifies whether to queue all visitors. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// on
	QueueAllEnable *string `json:"QueueAllEnable,omitempty" xml:"QueueAllEnable,omitempty"`
	// The queuing method. Valid values:
	//
	// - **random**: random.
	//
	// - **fifo**: first-in, first-out.
	//
	// - **passthrough**: passthrough.
	//
	// - **reject-all**: reject all.
	//
	// example:
	//
	// random
	QueuingMethod *string `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	// The HTTP status code returned by the waiting room. Valid values:
	//
	// - **200**
	//
	// - **202**
	//
	// - **429**
	//
	// example:
	//
	// 200
	QueuingStatusCode *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	// The session duration in minutes.
	//
	// example:
	//
	// 5
	SessionDuration *string `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// The site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7096621098****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The total number of active users.
	//
	// example:
	//
	// 300
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// The waiting room ID, which can be obtained by calling the [ListWaitingRooms](https://help.aliyun.com/document_detail/2850279.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6a51d5bc6460887abd129****
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
	// The type of the waiting room. Valid values:
	//
	// - **default**: default type.
	//
	// - **custom**: custom type.
	//
	// example:
	//
	// default
	WaitingRoomType *string `json:"WaitingRoomType,omitempty" xml:"WaitingRoomType,omitempty"`
}

func (s UpdateWaitingRoomRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaitingRoomRequest) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomRequest) GetCookieName() *string {
	return s.CookieName
}

func (s *UpdateWaitingRoomRequest) GetCustomPageHtml() *string {
	return s.CustomPageHtml
}

func (s *UpdateWaitingRoomRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWaitingRoomRequest) GetDisableSessionRenewalEnable() *string {
	return s.DisableSessionRenewalEnable
}

func (s *UpdateWaitingRoomRequest) GetEnable() *string {
	return s.Enable
}

func (s *UpdateWaitingRoomRequest) GetHostNameAndPath() []*UpdateWaitingRoomRequestHostNameAndPath {
	return s.HostNameAndPath
}

func (s *UpdateWaitingRoomRequest) GetJsonResponseEnable() *string {
	return s.JsonResponseEnable
}

func (s *UpdateWaitingRoomRequest) GetLanguage() *string {
	return s.Language
}

func (s *UpdateWaitingRoomRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWaitingRoomRequest) GetNewUsersPerMinute() *string {
	return s.NewUsersPerMinute
}

func (s *UpdateWaitingRoomRequest) GetQueueAllEnable() *string {
	return s.QueueAllEnable
}

func (s *UpdateWaitingRoomRequest) GetQueuingMethod() *string {
	return s.QueuingMethod
}

func (s *UpdateWaitingRoomRequest) GetQueuingStatusCode() *string {
	return s.QueuingStatusCode
}

func (s *UpdateWaitingRoomRequest) GetSessionDuration() *string {
	return s.SessionDuration
}

func (s *UpdateWaitingRoomRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateWaitingRoomRequest) GetTotalActiveUsers() *string {
	return s.TotalActiveUsers
}

func (s *UpdateWaitingRoomRequest) GetWaitingRoomId() *string {
	return s.WaitingRoomId
}

func (s *UpdateWaitingRoomRequest) GetWaitingRoomType() *string {
	return s.WaitingRoomType
}

func (s *UpdateWaitingRoomRequest) SetCookieName(v string) *UpdateWaitingRoomRequest {
	s.CookieName = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetCustomPageHtml(v string) *UpdateWaitingRoomRequest {
	s.CustomPageHtml = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetDescription(v string) *UpdateWaitingRoomRequest {
	s.Description = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetDisableSessionRenewalEnable(v string) *UpdateWaitingRoomRequest {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetEnable(v string) *UpdateWaitingRoomRequest {
	s.Enable = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetHostNameAndPath(v []*UpdateWaitingRoomRequestHostNameAndPath) *UpdateWaitingRoomRequest {
	s.HostNameAndPath = v
	return s
}

func (s *UpdateWaitingRoomRequest) SetJsonResponseEnable(v string) *UpdateWaitingRoomRequest {
	s.JsonResponseEnable = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetLanguage(v string) *UpdateWaitingRoomRequest {
	s.Language = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetName(v string) *UpdateWaitingRoomRequest {
	s.Name = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetNewUsersPerMinute(v string) *UpdateWaitingRoomRequest {
	s.NewUsersPerMinute = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetQueueAllEnable(v string) *UpdateWaitingRoomRequest {
	s.QueueAllEnable = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetQueuingMethod(v string) *UpdateWaitingRoomRequest {
	s.QueuingMethod = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetQueuingStatusCode(v string) *UpdateWaitingRoomRequest {
	s.QueuingStatusCode = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetSessionDuration(v string) *UpdateWaitingRoomRequest {
	s.SessionDuration = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetSiteId(v int64) *UpdateWaitingRoomRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetTotalActiveUsers(v string) *UpdateWaitingRoomRequest {
	s.TotalActiveUsers = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetWaitingRoomId(v string) *UpdateWaitingRoomRequest {
	s.WaitingRoomId = &v
	return s
}

func (s *UpdateWaitingRoomRequest) SetWaitingRoomType(v string) *UpdateWaitingRoomRequest {
	s.WaitingRoomType = &v
	return s
}

func (s *UpdateWaitingRoomRequest) Validate() error {
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

type UpdateWaitingRoomRequestHostNameAndPath struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The path.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The subdomain.
	//
	// example:
	//
	// test.
	Subdomain *string `json:"Subdomain,omitempty" xml:"Subdomain,omitempty"`
}

func (s UpdateWaitingRoomRequestHostNameAndPath) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaitingRoomRequestHostNameAndPath) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomRequestHostNameAndPath) GetDomain() *string {
	return s.Domain
}

func (s *UpdateWaitingRoomRequestHostNameAndPath) GetPath() *string {
	return s.Path
}

func (s *UpdateWaitingRoomRequestHostNameAndPath) GetSubdomain() *string {
	return s.Subdomain
}

func (s *UpdateWaitingRoomRequestHostNameAndPath) SetDomain(v string) *UpdateWaitingRoomRequestHostNameAndPath {
	s.Domain = &v
	return s
}

func (s *UpdateWaitingRoomRequestHostNameAndPath) SetPath(v string) *UpdateWaitingRoomRequestHostNameAndPath {
	s.Path = &v
	return s
}

func (s *UpdateWaitingRoomRequestHostNameAndPath) SetSubdomain(v string) *UpdateWaitingRoomRequestHostNameAndPath {
	s.Subdomain = &v
	return s
}

func (s *UpdateWaitingRoomRequestHostNameAndPath) Validate() error {
	return dara.Validate(s)
}
