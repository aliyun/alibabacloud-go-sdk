// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResourceControlEventsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionCode(v string) *QueryResourceControlEventsShrinkRequest
	GetActionCode() *string
	SetActionCodesShrink(v string) *QueryResourceControlEventsShrinkRequest
	GetActionCodesShrink() *string
	SetAliyunLang(v string) *QueryResourceControlEventsShrinkRequest
	GetAliyunLang() *string
	SetBusinessCode(v string) *QueryResourceControlEventsShrinkRequest
	GetBusinessCode() *string
	SetCaseCodesPrefixShrink(v string) *QueryResourceControlEventsShrinkRequest
	GetCaseCodesPrefixShrink() *string
	SetCurrent(v int32) *QueryResourceControlEventsShrinkRequest
	GetCurrent() *int32
	SetDomain(v string) *QueryResourceControlEventsShrinkRequest
	GetDomain() *string
	SetEventCode(v string) *QueryResourceControlEventsShrinkRequest
	GetEventCode() *string
	SetEventCodesShrink(v string) *QueryResourceControlEventsShrinkRequest
	GetEventCodesShrink() *string
	SetEventId(v string) *QueryResourceControlEventsShrinkRequest
	GetEventId() *string
	SetEventIdListShrink(v string) *QueryResourceControlEventsShrinkRequest
	GetEventIdListShrink() *string
	SetExcludeActionCodesShrink(v string) *QueryResourceControlEventsShrinkRequest
	GetExcludeActionCodesShrink() *string
	SetExcludeEventCodesShrink(v string) *QueryResourceControlEventsShrinkRequest
	GetExcludeEventCodesShrink() *string
	SetExcludeReasonsShrink(v string) *QueryResourceControlEventsShrinkRequest
	GetExcludeReasonsShrink() *string
	SetIncludeReasonsShrink(v string) *QueryResourceControlEventsShrinkRequest
	GetIncludeReasonsShrink() *string
	SetInstanceId(v string) *QueryResourceControlEventsShrinkRequest
	GetInstanceId() *string
	SetIp(v string) *QueryResourceControlEventsShrinkRequest
	GetIp() *string
	SetPageSize(v int32) *QueryResourceControlEventsShrinkRequest
	GetPageSize() *int32
	SetPunishEndTime(v string) *QueryResourceControlEventsShrinkRequest
	GetPunishEndTime() *string
	SetPunishStartTime(v string) *QueryResourceControlEventsShrinkRequest
	GetPunishStartTime() *string
	SetReason(v string) *QueryResourceControlEventsShrinkRequest
	GetReason() *string
	SetSourceCodesShrink(v string) *QueryResourceControlEventsShrinkRequest
	GetSourceCodesShrink() *string
	SetStatus(v string) *QueryResourceControlEventsShrinkRequest
	GetStatus() *string
	SetStatusListShrink(v string) *QueryResourceControlEventsShrinkRequest
	GetStatusListShrink() *string
	SetUrl(v string) *QueryResourceControlEventsShrinkRequest
	GetUrl() *string
}

type QueryResourceControlEventsShrinkRequest struct {
	// example:
	//
	// shutdown
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	// example:
	//
	// shutdown
	ActionCodesShrink *string `json:"ActionCodes,omitempty" xml:"ActionCodes,omitempty"`
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// example:
	//
	// ecs
	BusinessCode *string `json:"BusinessCode,omitempty" xml:"BusinessCode,omitempty"`
	// example:
	//
	// [\\"BANFF\\"]
	CaseCodesPrefixShrink *string `json:"CaseCodesPrefix,omitempty" xml:"CaseCodesPrefix,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// short.industry.taobao.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// BANFF_ECS_PE_ECS_MINING_SHUTDOWN
	EventCode *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	// example:
	//
	// ["BANFF_ECS_PE_ECS_MINING_SHUTDOWN"]
	EventCodesShrink *string `json:"EventCodes,omitempty" xml:"EventCodes,omitempty"`
	// example:
	//
	// 2PTOHhN3YUeaPWzq9FLmpdZ9EOW
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// [\\"7ac74fbfe06b2b85bb470083b7a05fb7\\",\\"1180c5bbff0a385b00d2cf73e3371d11\\"]
	EventIdListShrink *string `json:"EventIdList,omitempty" xml:"EventIdList,omitempty"`
	// example:
	//
	// [\\"shutdown\\"]
	ExcludeActionCodesShrink *string `json:"ExcludeActionCodes,omitempty" xml:"ExcludeActionCodes,omitempty"`
	// example:
	//
	// [\\"TEST_CASE\\"]
	ExcludeEventCodesShrink *string `json:"ExcludeEventCodes,omitempty" xml:"ExcludeEventCodes,omitempty"`
	// example:
	//
	// [\\"挖矿告警\\",\\"挖矿管控事件\\",\\"挖矿\\"]
	ExcludeReasonsShrink *string `json:"ExcludeReasons,omitempty" xml:"ExcludeReasons,omitempty"`
	// example:
	//
	// [\\"挖矿告警\\",\\"挖矿管控事件\\",\\"挖矿\\"]
	IncludeReasonsShrink *string `json:"IncludeReasons,omitempty" xml:"IncludeReasons,omitempty"`
	// example:
	//
	// rm-0iw73ro05vcwn6ntq
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 12.3*.22.11
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// 挖矿
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// [\\"MRM\\"]
	SourceCodesShrink *string `json:"SourceCodes,omitempty" xml:"SourceCodes,omitempty"`
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// [\\"Executing\\"]
	StatusListShrink *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	// example:
	//
	// https://qimg.xiaohongshu.com/circe/1040g1v831qggp28ln0705oft1i6k1jil889lhso?imageView2/2/w/1080/format/jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryResourceControlEventsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryResourceControlEventsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryResourceControlEventsShrinkRequest) GetActionCode() *string {
	return s.ActionCode
}

func (s *QueryResourceControlEventsShrinkRequest) GetActionCodesShrink() *string {
	return s.ActionCodesShrink
}

func (s *QueryResourceControlEventsShrinkRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *QueryResourceControlEventsShrinkRequest) GetBusinessCode() *string {
	return s.BusinessCode
}

func (s *QueryResourceControlEventsShrinkRequest) GetCaseCodesPrefixShrink() *string {
	return s.CaseCodesPrefixShrink
}

func (s *QueryResourceControlEventsShrinkRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *QueryResourceControlEventsShrinkRequest) GetDomain() *string {
	return s.Domain
}

func (s *QueryResourceControlEventsShrinkRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *QueryResourceControlEventsShrinkRequest) GetEventCodesShrink() *string {
	return s.EventCodesShrink
}

func (s *QueryResourceControlEventsShrinkRequest) GetEventId() *string {
	return s.EventId
}

func (s *QueryResourceControlEventsShrinkRequest) GetEventIdListShrink() *string {
	return s.EventIdListShrink
}

func (s *QueryResourceControlEventsShrinkRequest) GetExcludeActionCodesShrink() *string {
	return s.ExcludeActionCodesShrink
}

func (s *QueryResourceControlEventsShrinkRequest) GetExcludeEventCodesShrink() *string {
	return s.ExcludeEventCodesShrink
}

func (s *QueryResourceControlEventsShrinkRequest) GetExcludeReasonsShrink() *string {
	return s.ExcludeReasonsShrink
}

func (s *QueryResourceControlEventsShrinkRequest) GetIncludeReasonsShrink() *string {
	return s.IncludeReasonsShrink
}

func (s *QueryResourceControlEventsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryResourceControlEventsShrinkRequest) GetIp() *string {
	return s.Ip
}

func (s *QueryResourceControlEventsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryResourceControlEventsShrinkRequest) GetPunishEndTime() *string {
	return s.PunishEndTime
}

func (s *QueryResourceControlEventsShrinkRequest) GetPunishStartTime() *string {
	return s.PunishStartTime
}

func (s *QueryResourceControlEventsShrinkRequest) GetReason() *string {
	return s.Reason
}

func (s *QueryResourceControlEventsShrinkRequest) GetSourceCodesShrink() *string {
	return s.SourceCodesShrink
}

func (s *QueryResourceControlEventsShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *QueryResourceControlEventsShrinkRequest) GetStatusListShrink() *string {
	return s.StatusListShrink
}

func (s *QueryResourceControlEventsShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *QueryResourceControlEventsShrinkRequest) SetActionCode(v string) *QueryResourceControlEventsShrinkRequest {
	s.ActionCode = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetActionCodesShrink(v string) *QueryResourceControlEventsShrinkRequest {
	s.ActionCodesShrink = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetAliyunLang(v string) *QueryResourceControlEventsShrinkRequest {
	s.AliyunLang = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetBusinessCode(v string) *QueryResourceControlEventsShrinkRequest {
	s.BusinessCode = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetCaseCodesPrefixShrink(v string) *QueryResourceControlEventsShrinkRequest {
	s.CaseCodesPrefixShrink = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetCurrent(v int32) *QueryResourceControlEventsShrinkRequest {
	s.Current = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetDomain(v string) *QueryResourceControlEventsShrinkRequest {
	s.Domain = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetEventCode(v string) *QueryResourceControlEventsShrinkRequest {
	s.EventCode = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetEventCodesShrink(v string) *QueryResourceControlEventsShrinkRequest {
	s.EventCodesShrink = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetEventId(v string) *QueryResourceControlEventsShrinkRequest {
	s.EventId = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetEventIdListShrink(v string) *QueryResourceControlEventsShrinkRequest {
	s.EventIdListShrink = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetExcludeActionCodesShrink(v string) *QueryResourceControlEventsShrinkRequest {
	s.ExcludeActionCodesShrink = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetExcludeEventCodesShrink(v string) *QueryResourceControlEventsShrinkRequest {
	s.ExcludeEventCodesShrink = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetExcludeReasonsShrink(v string) *QueryResourceControlEventsShrinkRequest {
	s.ExcludeReasonsShrink = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetIncludeReasonsShrink(v string) *QueryResourceControlEventsShrinkRequest {
	s.IncludeReasonsShrink = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetInstanceId(v string) *QueryResourceControlEventsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetIp(v string) *QueryResourceControlEventsShrinkRequest {
	s.Ip = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetPageSize(v int32) *QueryResourceControlEventsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetPunishEndTime(v string) *QueryResourceControlEventsShrinkRequest {
	s.PunishEndTime = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetPunishStartTime(v string) *QueryResourceControlEventsShrinkRequest {
	s.PunishStartTime = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetReason(v string) *QueryResourceControlEventsShrinkRequest {
	s.Reason = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetSourceCodesShrink(v string) *QueryResourceControlEventsShrinkRequest {
	s.SourceCodesShrink = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetStatus(v string) *QueryResourceControlEventsShrinkRequest {
	s.Status = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetStatusListShrink(v string) *QueryResourceControlEventsShrinkRequest {
	s.StatusListShrink = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) SetUrl(v string) *QueryResourceControlEventsShrinkRequest {
	s.Url = &v
	return s
}

func (s *QueryResourceControlEventsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
