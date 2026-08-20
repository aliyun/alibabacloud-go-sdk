// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountSafetyIncidentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionCodesShrink(v string) *QueryAccountSafetyIncidentShrinkRequest
	GetActionCodesShrink() *string
	SetAliyunLang(v string) *QueryAccountSafetyIncidentShrinkRequest
	GetAliyunLang() *string
	SetCaseCode(v string) *QueryAccountSafetyIncidentShrinkRequest
	GetCaseCode() *string
	SetCaseCodesShrink(v string) *QueryAccountSafetyIncidentShrinkRequest
	GetCaseCodesShrink() *string
	SetCurrent(v string) *QueryAccountSafetyIncidentShrinkRequest
	GetCurrent() *string
	SetEventId(v string) *QueryAccountSafetyIncidentShrinkRequest
	GetEventId() *string
	SetEventIdsShrink(v string) *QueryAccountSafetyIncidentShrinkRequest
	GetEventIdsShrink() *string
	SetPageSize(v string) *QueryAccountSafetyIncidentShrinkRequest
	GetPageSize() *string
	SetPunishEndTime(v string) *QueryAccountSafetyIncidentShrinkRequest
	GetPunishEndTime() *string
	SetPunishStartTime(v string) *QueryAccountSafetyIncidentShrinkRequest
	GetPunishStartTime() *string
	SetResourceId(v string) *QueryAccountSafetyIncidentShrinkRequest
	GetResourceId() *string
	SetStatus(v string) *QueryAccountSafetyIncidentShrinkRequest
	GetStatus() *string
	SetStatusesShrink(v string) *QueryAccountSafetyIncidentShrinkRequest
	GetStatusesShrink() *string
}

type QueryAccountSafetyIncidentShrinkRequest struct {
	// The list of control action codes.
	ActionCodesShrink *string `json:"ActionCodes,omitempty" xml:"ActionCodes,omitempty"`
	// The internationalization language. Default value: zh. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The event name code.
	//
	// example:
	//
	// RISKCONTROL_IMS_IMS_BAN_SUBUSER
	CaseCode *string `json:"CaseCode,omitempty" xml:"CaseCode,omitempty"`
	// The list of event name codes.
	CaseCodesShrink *string `json:"CaseCodes,omitempty" xml:"CaseCodes,omitempty"`
	// The current page number. The value must be greater than 0.
	//
	// example:
	//
	// 1
	Current *string `json:"Current,omitempty" xml:"Current,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 2PTOHhN3YUeaPWzq9FLmpdZ9EOW
	EventId        *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventIdsShrink *string `json:"EventIds,omitempty" xml:"EventIds,omitempty"`
	// The number of records per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The control end time.
	//
	// > Format: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-03-16 15:15:00
	PunishEndTime *string `json:"PunishEndTime,omitempty" xml:"PunishEndTime,omitempty"`
	// The control start time.
	//
	// > Format: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-03-16 15:15:00
	PunishStartTime *string `json:"PunishStartTime,omitempty" xml:"PunishStartTime,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// RES001
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The event status. Valid values:
	//
	// - **Executing**: In progress.
	//
	// - **Removed**: Removed.
	//
	// - **Alerting**: Alerting.
	//
	// - **Ended**: Ended.
	//
	// example:
	//
	// Executing
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusesShrink *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
}

func (s QueryAccountSafetyIncidentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountSafetyIncidentShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountSafetyIncidentShrinkRequest) GetActionCodesShrink() *string {
	return s.ActionCodesShrink
}

func (s *QueryAccountSafetyIncidentShrinkRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *QueryAccountSafetyIncidentShrinkRequest) GetCaseCode() *string {
	return s.CaseCode
}

func (s *QueryAccountSafetyIncidentShrinkRequest) GetCaseCodesShrink() *string {
	return s.CaseCodesShrink
}

func (s *QueryAccountSafetyIncidentShrinkRequest) GetCurrent() *string {
	return s.Current
}

func (s *QueryAccountSafetyIncidentShrinkRequest) GetEventId() *string {
	return s.EventId
}

func (s *QueryAccountSafetyIncidentShrinkRequest) GetEventIdsShrink() *string {
	return s.EventIdsShrink
}

func (s *QueryAccountSafetyIncidentShrinkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *QueryAccountSafetyIncidentShrinkRequest) GetPunishEndTime() *string {
	return s.PunishEndTime
}

func (s *QueryAccountSafetyIncidentShrinkRequest) GetPunishStartTime() *string {
	return s.PunishStartTime
}

func (s *QueryAccountSafetyIncidentShrinkRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *QueryAccountSafetyIncidentShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *QueryAccountSafetyIncidentShrinkRequest) GetStatusesShrink() *string {
	return s.StatusesShrink
}

func (s *QueryAccountSafetyIncidentShrinkRequest) SetActionCodesShrink(v string) *QueryAccountSafetyIncidentShrinkRequest {
	s.ActionCodesShrink = &v
	return s
}

func (s *QueryAccountSafetyIncidentShrinkRequest) SetAliyunLang(v string) *QueryAccountSafetyIncidentShrinkRequest {
	s.AliyunLang = &v
	return s
}

func (s *QueryAccountSafetyIncidentShrinkRequest) SetCaseCode(v string) *QueryAccountSafetyIncidentShrinkRequest {
	s.CaseCode = &v
	return s
}

func (s *QueryAccountSafetyIncidentShrinkRequest) SetCaseCodesShrink(v string) *QueryAccountSafetyIncidentShrinkRequest {
	s.CaseCodesShrink = &v
	return s
}

func (s *QueryAccountSafetyIncidentShrinkRequest) SetCurrent(v string) *QueryAccountSafetyIncidentShrinkRequest {
	s.Current = &v
	return s
}

func (s *QueryAccountSafetyIncidentShrinkRequest) SetEventId(v string) *QueryAccountSafetyIncidentShrinkRequest {
	s.EventId = &v
	return s
}

func (s *QueryAccountSafetyIncidentShrinkRequest) SetEventIdsShrink(v string) *QueryAccountSafetyIncidentShrinkRequest {
	s.EventIdsShrink = &v
	return s
}

func (s *QueryAccountSafetyIncidentShrinkRequest) SetPageSize(v string) *QueryAccountSafetyIncidentShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAccountSafetyIncidentShrinkRequest) SetPunishEndTime(v string) *QueryAccountSafetyIncidentShrinkRequest {
	s.PunishEndTime = &v
	return s
}

func (s *QueryAccountSafetyIncidentShrinkRequest) SetPunishStartTime(v string) *QueryAccountSafetyIncidentShrinkRequest {
	s.PunishStartTime = &v
	return s
}

func (s *QueryAccountSafetyIncidentShrinkRequest) SetResourceId(v string) *QueryAccountSafetyIncidentShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *QueryAccountSafetyIncidentShrinkRequest) SetStatus(v string) *QueryAccountSafetyIncidentShrinkRequest {
	s.Status = &v
	return s
}

func (s *QueryAccountSafetyIncidentShrinkRequest) SetStatusesShrink(v string) *QueryAccountSafetyIncidentShrinkRequest {
	s.StatusesShrink = &v
	return s
}

func (s *QueryAccountSafetyIncidentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
