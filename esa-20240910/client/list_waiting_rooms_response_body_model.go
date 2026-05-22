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
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CookieName                  *string                                                    `json:"CookieName,omitempty" xml:"CookieName,omitempty"`
	CustomPageHtml              *string                                                    `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	Description                 *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	DisableSessionRenewalEnable *string                                                    `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	Enable                      *string                                                    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	HostNameAndPath             []*ListWaitingRoomsResponseBodyWaitingRoomsHostNameAndPath `json:"HostNameAndPath,omitempty" xml:"HostNameAndPath,omitempty" type:"Repeated"`
	JsonResponseEnable          *string                                                    `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	Language                    *string                                                    `json:"Language,omitempty" xml:"Language,omitempty"`
	Name                        *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	NewUsersPerMinute           *string                                                    `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	QueueAllEnable              *string                                                    `json:"QueueAllEnable,omitempty" xml:"QueueAllEnable,omitempty"`
	QueuingMethod               *string                                                    `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	QueuingStatusCode           *string                                                    `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	SessionDuration             *string                                                    `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	TotalActiveUsers            *string                                                    `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	WaitingRoomId               *string                                                    `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
	WaitingRoomType             *string                                                    `json:"WaitingRoomType,omitempty" xml:"WaitingRoomType,omitempty"`
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
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
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
