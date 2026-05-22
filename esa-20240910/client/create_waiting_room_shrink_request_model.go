// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWaitingRoomShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCookieName(v string) *CreateWaitingRoomShrinkRequest
	GetCookieName() *string
	SetCustomPageHtml(v string) *CreateWaitingRoomShrinkRequest
	GetCustomPageHtml() *string
	SetDescription(v string) *CreateWaitingRoomShrinkRequest
	GetDescription() *string
	SetDisableSessionRenewalEnable(v string) *CreateWaitingRoomShrinkRequest
	GetDisableSessionRenewalEnable() *string
	SetEnable(v string) *CreateWaitingRoomShrinkRequest
	GetEnable() *string
	SetHostNameAndPathShrink(v string) *CreateWaitingRoomShrinkRequest
	GetHostNameAndPathShrink() *string
	SetJsonResponseEnable(v string) *CreateWaitingRoomShrinkRequest
	GetJsonResponseEnable() *string
	SetLanguage(v string) *CreateWaitingRoomShrinkRequest
	GetLanguage() *string
	SetName(v string) *CreateWaitingRoomShrinkRequest
	GetName() *string
	SetNewUsersPerMinute(v string) *CreateWaitingRoomShrinkRequest
	GetNewUsersPerMinute() *string
	SetQueueAllEnable(v string) *CreateWaitingRoomShrinkRequest
	GetQueueAllEnable() *string
	SetQueuingMethod(v string) *CreateWaitingRoomShrinkRequest
	GetQueuingMethod() *string
	SetQueuingStatusCode(v string) *CreateWaitingRoomShrinkRequest
	GetQueuingStatusCode() *string
	SetSessionDuration(v string) *CreateWaitingRoomShrinkRequest
	GetSessionDuration() *string
	SetSiteId(v int64) *CreateWaitingRoomShrinkRequest
	GetSiteId() *int64
	SetTotalActiveUsers(v string) *CreateWaitingRoomShrinkRequest
	GetTotalActiveUsers() *string
	SetWaitingRoomType(v string) *CreateWaitingRoomShrinkRequest
	GetWaitingRoomType() *string
}

type CreateWaitingRoomShrinkRequest struct {
	// This parameter is required.
	CookieName                  *string `json:"CookieName,omitempty" xml:"CookieName,omitempty"`
	CustomPageHtml              *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	Description                 *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisableSessionRenewalEnable *string `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	// This parameter is required.
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// This parameter is required.
	HostNameAndPathShrink *string `json:"HostNameAndPath,omitempty" xml:"HostNameAndPath,omitempty"`
	JsonResponseEnable    *string `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	Language              *string `json:"Language,omitempty" xml:"Language,omitempty"`
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

func (s CreateWaitingRoomShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWaitingRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomShrinkRequest) GetCookieName() *string {
	return s.CookieName
}

func (s *CreateWaitingRoomShrinkRequest) GetCustomPageHtml() *string {
	return s.CustomPageHtml
}

func (s *CreateWaitingRoomShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWaitingRoomShrinkRequest) GetDisableSessionRenewalEnable() *string {
	return s.DisableSessionRenewalEnable
}

func (s *CreateWaitingRoomShrinkRequest) GetEnable() *string {
	return s.Enable
}

func (s *CreateWaitingRoomShrinkRequest) GetHostNameAndPathShrink() *string {
	return s.HostNameAndPathShrink
}

func (s *CreateWaitingRoomShrinkRequest) GetJsonResponseEnable() *string {
	return s.JsonResponseEnable
}

func (s *CreateWaitingRoomShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateWaitingRoomShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateWaitingRoomShrinkRequest) GetNewUsersPerMinute() *string {
	return s.NewUsersPerMinute
}

func (s *CreateWaitingRoomShrinkRequest) GetQueueAllEnable() *string {
	return s.QueueAllEnable
}

func (s *CreateWaitingRoomShrinkRequest) GetQueuingMethod() *string {
	return s.QueuingMethod
}

func (s *CreateWaitingRoomShrinkRequest) GetQueuingStatusCode() *string {
	return s.QueuingStatusCode
}

func (s *CreateWaitingRoomShrinkRequest) GetSessionDuration() *string {
	return s.SessionDuration
}

func (s *CreateWaitingRoomShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateWaitingRoomShrinkRequest) GetTotalActiveUsers() *string {
	return s.TotalActiveUsers
}

func (s *CreateWaitingRoomShrinkRequest) GetWaitingRoomType() *string {
	return s.WaitingRoomType
}

func (s *CreateWaitingRoomShrinkRequest) SetCookieName(v string) *CreateWaitingRoomShrinkRequest {
	s.CookieName = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetCustomPageHtml(v string) *CreateWaitingRoomShrinkRequest {
	s.CustomPageHtml = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetDescription(v string) *CreateWaitingRoomShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetDisableSessionRenewalEnable(v string) *CreateWaitingRoomShrinkRequest {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetEnable(v string) *CreateWaitingRoomShrinkRequest {
	s.Enable = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetHostNameAndPathShrink(v string) *CreateWaitingRoomShrinkRequest {
	s.HostNameAndPathShrink = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetJsonResponseEnable(v string) *CreateWaitingRoomShrinkRequest {
	s.JsonResponseEnable = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetLanguage(v string) *CreateWaitingRoomShrinkRequest {
	s.Language = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetName(v string) *CreateWaitingRoomShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetNewUsersPerMinute(v string) *CreateWaitingRoomShrinkRequest {
	s.NewUsersPerMinute = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetQueueAllEnable(v string) *CreateWaitingRoomShrinkRequest {
	s.QueueAllEnable = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetQueuingMethod(v string) *CreateWaitingRoomShrinkRequest {
	s.QueuingMethod = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetQueuingStatusCode(v string) *CreateWaitingRoomShrinkRequest {
	s.QueuingStatusCode = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetSessionDuration(v string) *CreateWaitingRoomShrinkRequest {
	s.SessionDuration = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetSiteId(v int64) *CreateWaitingRoomShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetTotalActiveUsers(v string) *CreateWaitingRoomShrinkRequest {
	s.TotalActiveUsers = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) SetWaitingRoomType(v string) *CreateWaitingRoomShrinkRequest {
	s.WaitingRoomType = &v
	return s
}

func (s *CreateWaitingRoomShrinkRequest) Validate() error {
	return dara.Validate(s)
}
