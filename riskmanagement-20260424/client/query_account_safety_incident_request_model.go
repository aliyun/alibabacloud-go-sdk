// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountSafetyIncidentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *QueryAccountSafetyIncidentRequest
	GetAliyunLang() *string
	SetCaseCode(v string) *QueryAccountSafetyIncidentRequest
	GetCaseCode() *string
	SetCurrent(v string) *QueryAccountSafetyIncidentRequest
	GetCurrent() *string
	SetEventId(v string) *QueryAccountSafetyIncidentRequest
	GetEventId() *string
	SetPageSize(v string) *QueryAccountSafetyIncidentRequest
	GetPageSize() *string
	SetPunishEndTime(v string) *QueryAccountSafetyIncidentRequest
	GetPunishEndTime() *string
	SetPunishStartTime(v string) *QueryAccountSafetyIncidentRequest
	GetPunishStartTime() *string
	SetResourceId(v string) *QueryAccountSafetyIncidentRequest
	GetResourceId() *string
	SetStatus(v string) *QueryAccountSafetyIncidentRequest
	GetStatus() *string
}

type QueryAccountSafetyIncidentRequest struct {
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// example:
	//
	// RISKCONTROL_IMS_IMS_BAN_SUBUSER
	CaseCode *string `json:"CaseCode,omitempty" xml:"CaseCode,omitempty"`
	// example:
	//
	// 1
	Current *string `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// 2PTOHhN3YUeaPWzq9FLmpdZ9EOW
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	PunishEndTime *string `json:"PunishEndTime,omitempty" xml:"PunishEndTime,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	PunishStartTime *string `json:"PunishStartTime,omitempty" xml:"PunishStartTime,omitempty"`
	// example:
	//
	// RES001
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryAccountSafetyIncidentRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountSafetyIncidentRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountSafetyIncidentRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *QueryAccountSafetyIncidentRequest) GetCaseCode() *string {
	return s.CaseCode
}

func (s *QueryAccountSafetyIncidentRequest) GetCurrent() *string {
	return s.Current
}

func (s *QueryAccountSafetyIncidentRequest) GetEventId() *string {
	return s.EventId
}

func (s *QueryAccountSafetyIncidentRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *QueryAccountSafetyIncidentRequest) GetPunishEndTime() *string {
	return s.PunishEndTime
}

func (s *QueryAccountSafetyIncidentRequest) GetPunishStartTime() *string {
	return s.PunishStartTime
}

func (s *QueryAccountSafetyIncidentRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *QueryAccountSafetyIncidentRequest) GetStatus() *string {
	return s.Status
}

func (s *QueryAccountSafetyIncidentRequest) SetAliyunLang(v string) *QueryAccountSafetyIncidentRequest {
	s.AliyunLang = &v
	return s
}

func (s *QueryAccountSafetyIncidentRequest) SetCaseCode(v string) *QueryAccountSafetyIncidentRequest {
	s.CaseCode = &v
	return s
}

func (s *QueryAccountSafetyIncidentRequest) SetCurrent(v string) *QueryAccountSafetyIncidentRequest {
	s.Current = &v
	return s
}

func (s *QueryAccountSafetyIncidentRequest) SetEventId(v string) *QueryAccountSafetyIncidentRequest {
	s.EventId = &v
	return s
}

func (s *QueryAccountSafetyIncidentRequest) SetPageSize(v string) *QueryAccountSafetyIncidentRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAccountSafetyIncidentRequest) SetPunishEndTime(v string) *QueryAccountSafetyIncidentRequest {
	s.PunishEndTime = &v
	return s
}

func (s *QueryAccountSafetyIncidentRequest) SetPunishStartTime(v string) *QueryAccountSafetyIncidentRequest {
	s.PunishStartTime = &v
	return s
}

func (s *QueryAccountSafetyIncidentRequest) SetResourceId(v string) *QueryAccountSafetyIncidentRequest {
	s.ResourceId = &v
	return s
}

func (s *QueryAccountSafetyIncidentRequest) SetStatus(v string) *QueryAccountSafetyIncidentRequest {
	s.Status = &v
	return s
}

func (s *QueryAccountSafetyIncidentRequest) Validate() error {
	return dara.Validate(s)
}
