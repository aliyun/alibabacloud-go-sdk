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
	CustomPageHtml              *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	Description                 *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisableSessionRenewalEnable *string `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	// This parameter is required.
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// This parameter is required.
	EndTime            *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	JsonResponseEnable *string `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	Language           *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	NewUsersPerMinute *string `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	PreQueueEnable    *string `json:"PreQueueEnable,omitempty" xml:"PreQueueEnable,omitempty"`
	PreQueueStartTime *string `json:"PreQueueStartTime,omitempty" xml:"PreQueueStartTime,omitempty"`
	// This parameter is required.
	QueuingMethod *string `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	// This parameter is required.
	QueuingStatusCode    *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	RandomPreQueueEnable *string `json:"RandomPreQueueEnable,omitempty" xml:"RandomPreQueueEnable,omitempty"`
	// This parameter is required.
	SessionDuration *string `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	WaitingRoomId    *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
	// This parameter is required.
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
