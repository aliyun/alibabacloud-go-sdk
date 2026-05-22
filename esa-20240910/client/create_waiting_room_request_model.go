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
	// This parameter is required.
	CookieName                  *string `json:"CookieName,omitempty" xml:"CookieName,omitempty"`
	CustomPageHtml              *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	Description                 *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisableSessionRenewalEnable *string `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	// This parameter is required.
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// This parameter is required.
	HostNameAndPath    []*CreateWaitingRoomRequestHostNameAndPath `json:"HostNameAndPath,omitempty" xml:"HostNameAndPath,omitempty" type:"Repeated"`
	JsonResponseEnable *string                                    `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	Language           *string                                    `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	NewUsersPerMinute *string `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	QueueAllEnable    *string `json:"QueueAllEnable,omitempty" xml:"QueueAllEnable,omitempty"`
	// This parameter is required.
	QueuingMethod *string `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	// This parameter is required.
	QueuingStatusCode *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	// This parameter is required.
	SessionDuration *string `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// This parameter is required.
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
	// This parameter is required.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
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
