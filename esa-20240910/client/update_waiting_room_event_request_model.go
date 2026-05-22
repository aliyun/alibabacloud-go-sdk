// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaitingRoomEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomPageHtml(v string) *UpdateWaitingRoomEventRequest
	GetCustomPageHtml() *string
	SetDescription(v string) *UpdateWaitingRoomEventRequest
	GetDescription() *string
	SetDisableSessionRenewalEnable(v string) *UpdateWaitingRoomEventRequest
	GetDisableSessionRenewalEnable() *string
	SetEnable(v string) *UpdateWaitingRoomEventRequest
	GetEnable() *string
	SetEndTime(v string) *UpdateWaitingRoomEventRequest
	GetEndTime() *string
	SetJsonResponseEnable(v string) *UpdateWaitingRoomEventRequest
	GetJsonResponseEnable() *string
	SetLanguage(v string) *UpdateWaitingRoomEventRequest
	GetLanguage() *string
	SetName(v string) *UpdateWaitingRoomEventRequest
	GetName() *string
	SetNewUsersPerMinute(v string) *UpdateWaitingRoomEventRequest
	GetNewUsersPerMinute() *string
	SetPreQueueEnable(v string) *UpdateWaitingRoomEventRequest
	GetPreQueueEnable() *string
	SetPreQueueStartTime(v string) *UpdateWaitingRoomEventRequest
	GetPreQueueStartTime() *string
	SetQueuingMethod(v string) *UpdateWaitingRoomEventRequest
	GetQueuingMethod() *string
	SetQueuingStatusCode(v string) *UpdateWaitingRoomEventRequest
	GetQueuingStatusCode() *string
	SetRandomPreQueueEnable(v string) *UpdateWaitingRoomEventRequest
	GetRandomPreQueueEnable() *string
	SetSessionDuration(v string) *UpdateWaitingRoomEventRequest
	GetSessionDuration() *string
	SetSiteId(v int64) *UpdateWaitingRoomEventRequest
	GetSiteId() *int64
	SetStartTime(v string) *UpdateWaitingRoomEventRequest
	GetStartTime() *string
	SetTotalActiveUsers(v string) *UpdateWaitingRoomEventRequest
	GetTotalActiveUsers() *string
	SetWaitingRoomEventId(v int64) *UpdateWaitingRoomEventRequest
	GetWaitingRoomEventId() *int64
	SetWaitingRoomType(v string) *UpdateWaitingRoomEventRequest
	GetWaitingRoomType() *string
}

type UpdateWaitingRoomEventRequest struct {
	CustomPageHtml              *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
	Description                 *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisableSessionRenewalEnable *string `json:"DisableSessionRenewalEnable,omitempty" xml:"DisableSessionRenewalEnable,omitempty"`
	Enable                      *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	EndTime                     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	JsonResponseEnable          *string `json:"JsonResponseEnable,omitempty" xml:"JsonResponseEnable,omitempty"`
	Language                    *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Name                        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NewUsersPerMinute           *string `json:"NewUsersPerMinute,omitempty" xml:"NewUsersPerMinute,omitempty"`
	PreQueueEnable              *string `json:"PreQueueEnable,omitempty" xml:"PreQueueEnable,omitempty"`
	PreQueueStartTime           *string `json:"PreQueueStartTime,omitempty" xml:"PreQueueStartTime,omitempty"`
	QueuingMethod               *string `json:"QueuingMethod,omitempty" xml:"QueuingMethod,omitempty"`
	QueuingStatusCode           *string `json:"QueuingStatusCode,omitempty" xml:"QueuingStatusCode,omitempty"`
	RandomPreQueueEnable        *string `json:"RandomPreQueueEnable,omitempty" xml:"RandomPreQueueEnable,omitempty"`
	SessionDuration             *string `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// This parameter is required.
	SiteId           *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TotalActiveUsers *string `json:"TotalActiveUsers,omitempty" xml:"TotalActiveUsers,omitempty"`
	// This parameter is required.
	WaitingRoomEventId *int64  `json:"WaitingRoomEventId,omitempty" xml:"WaitingRoomEventId,omitempty"`
	WaitingRoomType    *string `json:"WaitingRoomType,omitempty" xml:"WaitingRoomType,omitempty"`
}

func (s UpdateWaitingRoomEventRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaitingRoomEventRequest) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomEventRequest) GetCustomPageHtml() *string {
	return s.CustomPageHtml
}

func (s *UpdateWaitingRoomEventRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWaitingRoomEventRequest) GetDisableSessionRenewalEnable() *string {
	return s.DisableSessionRenewalEnable
}

func (s *UpdateWaitingRoomEventRequest) GetEnable() *string {
	return s.Enable
}

func (s *UpdateWaitingRoomEventRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateWaitingRoomEventRequest) GetJsonResponseEnable() *string {
	return s.JsonResponseEnable
}

func (s *UpdateWaitingRoomEventRequest) GetLanguage() *string {
	return s.Language
}

func (s *UpdateWaitingRoomEventRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWaitingRoomEventRequest) GetNewUsersPerMinute() *string {
	return s.NewUsersPerMinute
}

func (s *UpdateWaitingRoomEventRequest) GetPreQueueEnable() *string {
	return s.PreQueueEnable
}

func (s *UpdateWaitingRoomEventRequest) GetPreQueueStartTime() *string {
	return s.PreQueueStartTime
}

func (s *UpdateWaitingRoomEventRequest) GetQueuingMethod() *string {
	return s.QueuingMethod
}

func (s *UpdateWaitingRoomEventRequest) GetQueuingStatusCode() *string {
	return s.QueuingStatusCode
}

func (s *UpdateWaitingRoomEventRequest) GetRandomPreQueueEnable() *string {
	return s.RandomPreQueueEnable
}

func (s *UpdateWaitingRoomEventRequest) GetSessionDuration() *string {
	return s.SessionDuration
}

func (s *UpdateWaitingRoomEventRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateWaitingRoomEventRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateWaitingRoomEventRequest) GetTotalActiveUsers() *string {
	return s.TotalActiveUsers
}

func (s *UpdateWaitingRoomEventRequest) GetWaitingRoomEventId() *int64 {
	return s.WaitingRoomEventId
}

func (s *UpdateWaitingRoomEventRequest) GetWaitingRoomType() *string {
	return s.WaitingRoomType
}

func (s *UpdateWaitingRoomEventRequest) SetCustomPageHtml(v string) *UpdateWaitingRoomEventRequest {
	s.CustomPageHtml = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetDescription(v string) *UpdateWaitingRoomEventRequest {
	s.Description = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetDisableSessionRenewalEnable(v string) *UpdateWaitingRoomEventRequest {
	s.DisableSessionRenewalEnable = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetEnable(v string) *UpdateWaitingRoomEventRequest {
	s.Enable = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetEndTime(v string) *UpdateWaitingRoomEventRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetJsonResponseEnable(v string) *UpdateWaitingRoomEventRequest {
	s.JsonResponseEnable = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetLanguage(v string) *UpdateWaitingRoomEventRequest {
	s.Language = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetName(v string) *UpdateWaitingRoomEventRequest {
	s.Name = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetNewUsersPerMinute(v string) *UpdateWaitingRoomEventRequest {
	s.NewUsersPerMinute = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetPreQueueEnable(v string) *UpdateWaitingRoomEventRequest {
	s.PreQueueEnable = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetPreQueueStartTime(v string) *UpdateWaitingRoomEventRequest {
	s.PreQueueStartTime = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetQueuingMethod(v string) *UpdateWaitingRoomEventRequest {
	s.QueuingMethod = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetQueuingStatusCode(v string) *UpdateWaitingRoomEventRequest {
	s.QueuingStatusCode = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetRandomPreQueueEnable(v string) *UpdateWaitingRoomEventRequest {
	s.RandomPreQueueEnable = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetSessionDuration(v string) *UpdateWaitingRoomEventRequest {
	s.SessionDuration = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetSiteId(v int64) *UpdateWaitingRoomEventRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetStartTime(v string) *UpdateWaitingRoomEventRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetTotalActiveUsers(v string) *UpdateWaitingRoomEventRequest {
	s.TotalActiveUsers = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetWaitingRoomEventId(v int64) *UpdateWaitingRoomEventRequest {
	s.WaitingRoomEventId = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) SetWaitingRoomType(v string) *UpdateWaitingRoomEventRequest {
	s.WaitingRoomType = &v
	return s
}

func (s *UpdateWaitingRoomEventRequest) Validate() error {
	return dara.Validate(s)
}
