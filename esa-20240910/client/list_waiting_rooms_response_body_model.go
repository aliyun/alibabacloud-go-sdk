// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaitingRoomsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWaitingRoomsResponseBody
	GetRequestId() *string
	SetWaitingRooms(v []*ListWaitingRoomsResponseBodyWaitingRooms) *ListWaitingRoomsResponseBody
	GetWaitingRooms() []*ListWaitingRoomsResponseBodyWaitingRooms
}

type ListWaitingRoomsResponseBody struct {
	// The request ID, which is used to trace API calls.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of waiting rooms.
	WaitingRooms []*ListWaitingRoomsResponseBodyWaitingRooms `json:"WaitingRooms,omitempty" xml:"WaitingRooms,omitempty" type:"Repeated"`
}

func (s ListWaitingRoomsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingRoomsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWaitingRoomsResponseBody) GetWaitingRooms() []*ListWaitingRoomsResponseBodyWaitingRooms {
	return s.WaitingRooms
}

func (s *ListWaitingRoomsResponseBody) SetRequestId(v string) *ListWaitingRoomsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWaitingRoomsResponseBody) SetWaitingRooms(v []*ListWaitingRoomsResponseBodyWaitingRooms) *ListWaitingRoomsResponseBody {
	s.WaitingRooms = v
	return s
}

func (s *ListWaitingRoomsResponseBody) Validate() error {
	if s.WaitingRooms != nil {
		for _, item := range s.WaitingRooms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWaitingRoomsResponseBodyWaitingRooms struct {
	// The custom cookie name.
	//
	// example:
	//
	// __aliwaitingroom_example
	CookieName *string `json:"CookieName,omitempty" xml:"CookieName,omitempty"`
	// The HTML content or identifier of the custom queuing page. This parameter is valid only when WaitingRoomType is set to custom. The content must be URL-encoded.
	//
	// example:
	//
	// Custom HTML content
	CustomPageHtml *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	// The waiting room description.
	//
	// example:
	//
	// Test waiting room
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
	// The enabled status. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The hostname and path configurations.
	HostNameAndPath []*ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath `json:"HostNameAndPath,omitempty" xml:"HostNameAndPath,omitempty" type:"Repeated"`
	// Specifies whether to enable JSON response. When enabled, requests with an Accept header containing "application/json" return JSON data. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// on
	JsonResponseEnable *string `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	// The language of the waiting room page. This parameter is required when the waiting room type is default. Valid values:
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
	// The waiting room name.
	//
	// example:
	//
	// Test waiting room
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
	// - **random**: Random.
	//
	// - **fifo**: First in, first out.
	//
	// - **passthrough**: Passthrough.
	//
	// - **reject-all**: Reject all.
	//
	// example:
	//
	// random
	QueuingMethod *string `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	// The waiting room status code. Valid values:
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
	// The session duration, in minutes.
	//
	// example:
	//
	// 3600
	SessionDuration *string `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// The total number of active users.
	//
	// example:
	//
	// 300
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// The waiting room ID, which uniquely identifies a waiting room.
	//
	// example:
	//
	// 6a51d5bc6460887abd1291dc7d4d****
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
	// The waiting room type. Valid values:
	//
	// - **default**: Default type.
	//
	// - **custom**: Custom type.
	//
	// example:
	//
	// default
	WaitingRoomType *string `json:"WaitingRoomType,omitempty" xml:"WaitingRoomType,omitempty"`
}

func (s ListWaitingRoomsResponseBodyWaitingRooms) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingRoomsResponseBodyWaitingRooms) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetCookieName() *string {
	return s.CookieName
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetCustomPageHtml() *string {
	return s.CustomPageHtml
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetDescription() *string {
	return s.Description
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetDisableSessionRenewalEnable() *string {
	return s.DisableSessionRenewalEnable
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetEnable() *string {
	return s.Enable
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetHostNameAndPath() []*ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath {
	return s.HostNameAndPath
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetJsonResponseEnable() *string {
	return s.JsonResponseEnable
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetLanguage() *string {
	return s.Language
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetName() *string {
	return s.Name
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetNewUsersPerMinute() *string {
	return s.NewUsersPerMinute
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetQueueAllEnable() *string {
	return s.QueueAllEnable
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetQueuingMethod() *string {
	return s.QueuingMethod
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetQueuingStatusCode() *string {
	return s.QueuingStatusCode
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetSessionDuration() *string {
	return s.SessionDuration
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetTotalActiveUsers() *string {
	return s.TotalActiveUsers
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetWaitingRoomId() *string {
	return s.WaitingRoomId
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) GetWaitingRoomType() *string {
	return s.WaitingRoomType
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetCookieName(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.CookieName = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetCustomPageHtml(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.CustomPageHtml = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetDescription(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.Description = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetDisableSessionRenewalEnable(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetEnable(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.Enable = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetHostNameAndPath(v []*ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.HostNameAndPath = v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetJsonResponseEnable(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.JsonResponseEnable = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetLanguage(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.Language = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetName(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.Name = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetNewUsersPerMinute(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.NewUsersPerMinute = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetQueueAllEnable(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.QueueAllEnable = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetQueuingMethod(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.QueuingMethod = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetQueuingStatusCode(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.QueuingStatusCode = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetSessionDuration(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.SessionDuration = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetTotalActiveUsers(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.TotalActiveUsers = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetWaitingRoomId(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.WaitingRoomId = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) SetWaitingRoomType(v string) *ListWaitingRoomsResponseBodyWaitingRooms {
	s.WaitingRoomType = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRooms) Validate() error {
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

type ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath struct {
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

func (s ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) GetDomain() *string {
	return s.Domain
}

func (s *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) GetPath() *string {
	return s.Path
}

func (s *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) GetSubdomain() *string {
	return s.Subdomain
}

func (s *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) SetDomain(v string) *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath {
	s.Domain = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) SetPath(v string) *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath {
	s.Path = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) SetSubdomain(v string) *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath {
	s.Subdomain = &v
	return s
}

func (s *ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath) Validate() error {
	return dara.Validate(s)
}
