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
	CookieName                  *string                                    `json:"CookieName,omitempty" xml:"CookieName,omitempty"`
	CustomPageHtml              *string                                    `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	Description                 *string                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	DisableSessionRenewalEnable *string                                    `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	Enable                      *string                                    `json:"Enable,omitempty" xml:"Enable,omitempty"`
	HostNameAndPath             []*UpdateWaitingRoomRequestHostNameAndPath `json:"HostNameAndPath,omitempty" xml:"HostNameAndPath,omitempty" type:"Repeated"`
	JsonResponseEnable          *string                                    `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	Language                    *string                                    `json:"Language,omitempty" xml:"Language,omitempty"`
	Name                        *string                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	NewUsersPerMinute           *string                                    `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	QueueAllEnable              *string                                    `json:"QueueAllEnable,omitempty" xml:"QueueAllEnable,omitempty"`
	QueuingMethod               *string                                    `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	QueuingStatusCode           *string                                    `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	SessionDuration             *string                                    `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// This parameter is required.
	SiteId           *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// This parameter is required.
	WaitingRoomId   *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
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
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
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
