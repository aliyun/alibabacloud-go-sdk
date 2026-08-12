// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResourceControlEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionCode(v string) *QueryResourceControlEventsRequest
	GetActionCode() *string
	SetActionCodes(v []*string) *QueryResourceControlEventsRequest
	GetActionCodes() []*string
	SetAliyunLang(v string) *QueryResourceControlEventsRequest
	GetAliyunLang() *string
	SetBusinessCode(v string) *QueryResourceControlEventsRequest
	GetBusinessCode() *string
	SetBusinessCodes(v []*string) *QueryResourceControlEventsRequest
	GetBusinessCodes() []*string
	SetCaseCodesPrefix(v []*string) *QueryResourceControlEventsRequest
	GetCaseCodesPrefix() []*string
	SetCurrent(v int32) *QueryResourceControlEventsRequest
	GetCurrent() *int32
	SetDomain(v string) *QueryResourceControlEventsRequest
	GetDomain() *string
	SetEventCode(v string) *QueryResourceControlEventsRequest
	GetEventCode() *string
	SetEventCodes(v []*string) *QueryResourceControlEventsRequest
	GetEventCodes() []*string
	SetEventId(v string) *QueryResourceControlEventsRequest
	GetEventId() *string
	SetEventIdList(v []*string) *QueryResourceControlEventsRequest
	GetEventIdList() []*string
	SetExcludeActionCodes(v []*string) *QueryResourceControlEventsRequest
	GetExcludeActionCodes() []*string
	SetExcludeEventCodes(v []*string) *QueryResourceControlEventsRequest
	GetExcludeEventCodes() []*string
	SetExcludeReasons(v []*string) *QueryResourceControlEventsRequest
	GetExcludeReasons() []*string
	SetIncludeReasons(v []*string) *QueryResourceControlEventsRequest
	GetIncludeReasons() []*string
	SetInstanceId(v string) *QueryResourceControlEventsRequest
	GetInstanceId() *string
	SetIp(v string) *QueryResourceControlEventsRequest
	GetIp() *string
	SetPageSize(v int32) *QueryResourceControlEventsRequest
	GetPageSize() *int32
	SetPunishEndTime(v string) *QueryResourceControlEventsRequest
	GetPunishEndTime() *string
	SetPunishStartTime(v string) *QueryResourceControlEventsRequest
	GetPunishStartTime() *string
	SetReason(v string) *QueryResourceControlEventsRequest
	GetReason() *string
	SetSourceCodes(v []*string) *QueryResourceControlEventsRequest
	GetSourceCodes() []*string
	SetStatus(v string) *QueryResourceControlEventsRequest
	GetStatus() *string
	SetStatusList(v []*string) *QueryResourceControlEventsRequest
	GetStatusList() []*string
	SetUrl(v string) *QueryResourceControlEventsRequest
	GetUrl() *string
}

type QueryResourceControlEventsRequest struct {
	// The action name code.
	//
	// example:
	//
	// shutdown
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	// The collection of control action name codes.
	//
	// example:
	//
	// shutdown
	ActionCodes []*string `json:"ActionCodes,omitempty" xml:"ActionCodes,omitempty" type:"Repeated"`
	// The internationalization language.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The product.
	//
	// example:
	//
	// ecs
	BusinessCode  *string   `json:"BusinessCode,omitempty" xml:"BusinessCode,omitempty"`
	BusinessCodes []*string `json:"BusinessCodes,omitempty" xml:"BusinessCodes,omitempty" type:"Repeated"`
	// The collection of event name code prefixes.
	//
	// example:
	//
	// [\\"BANFF\\"]
	CaseCodesPrefix []*string `json:"CaseCodesPrefix,omitempty" xml:"CaseCodesPrefix,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// The domain name.
	//
	// example:
	//
	// short.industry.taobao.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The event name code.
	//
	// example:
	//
	// BANFF_ECS_PE_ECS_MINING_SHUTDOWN
	EventCode *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	// The collection of event name codes.
	//
	// example:
	//
	// ["BANFF_ECS_PE_ECS_MINING_SHUTDOWN"]
	EventCodes []*string `json:"EventCodes,omitempty" xml:"EventCodes,omitempty" type:"Repeated"`
	// The event ID.
	//
	// example:
	//
	// 2PTOHhN3YUeaPWzq9FLmpdZ9EOW
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The collection of event IDs.
	//
	// example:
	//
	// [\\"7ac74fbfe06b2b85bb470083b7a05fb7\\",\\"1180c5bbff0a385b00d2cf73e3371d11\\"]
	EventIdList []*string `json:"EventIdList,omitempty" xml:"EventIdList,omitempty" type:"Repeated"`
	// The collection of excluded control action name codes.
	//
	// example:
	//
	// [\\"shutdown\\"]
	ExcludeActionCodes []*string `json:"ExcludeActionCodes,omitempty" xml:"ExcludeActionCodes,omitempty" type:"Repeated"`
	// The collection of excluded event name codes.
	//
	// example:
	//
	// [\\"TEST_CASE\\"]
	ExcludeEventCodes []*string `json:"ExcludeEventCodes,omitempty" xml:"ExcludeEventCodes,omitempty" type:"Repeated"`
	// The collection of excluded event reasons.
	//
	// example:
	//
	// [\\"Cryptomining alert\\",\\"Cryptomining control event\\",\\"Cryptomining\\"]
	ExcludeReasons []*string `json:"ExcludeReasons,omitempty" xml:"ExcludeReasons,omitempty" type:"Repeated"`
	// The collection of included event reasons.
	//
	// example:
	//
	// [\\"Cryptomining alert\\",\\"Cryptomining control event\\",\\"Cryptomining\\"]
	IncludeReasons []*string `json:"IncludeReasons,omitempty" xml:"IncludeReasons,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// rm-0iw73ro05vcwn6ntq
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// IP
	//
	// example:
	//
	// 12.3*.22.11
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The penalty end time.
	//
	// example:
	//
	// 2026-03-16 15:15:00
	PunishEndTime *string `json:"PunishEndTime,omitempty" xml:"PunishEndTime,omitempty"`
	// The penalty start time.
	//
	// example:
	//
	// 2026-03-16 15:15:00
	PunishStartTime *string `json:"PunishStartTime,omitempty" xml:"PunishStartTime,omitempty"`
	// The event reason.
	//
	// example:
	//
	// Cryptomining.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The collection of event source codes.
	//
	// example:
	//
	// [\\"MRM\\"]
	SourceCodes []*string `json:"SourceCodes,omitempty" xml:"SourceCodes,omitempty" type:"Repeated"`
	// The task status.
	//
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The collection of task statuses.
	//
	// example:
	//
	// [\\"Executing\\"]
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// The control URL.
	//
	// example:
	//
	// https://qimg.xiaohongshu.com/circe/1040g1v831qggp28ln0705oft1i6k1jil889lhso?imageView2/2/w/1080/format/jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryResourceControlEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryResourceControlEventsRequest) GoString() string {
	return s.String()
}

func (s *QueryResourceControlEventsRequest) GetActionCode() *string {
	return s.ActionCode
}

func (s *QueryResourceControlEventsRequest) GetActionCodes() []*string {
	return s.ActionCodes
}

func (s *QueryResourceControlEventsRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *QueryResourceControlEventsRequest) GetBusinessCode() *string {
	return s.BusinessCode
}

func (s *QueryResourceControlEventsRequest) GetBusinessCodes() []*string {
	return s.BusinessCodes
}

func (s *QueryResourceControlEventsRequest) GetCaseCodesPrefix() []*string {
	return s.CaseCodesPrefix
}

func (s *QueryResourceControlEventsRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *QueryResourceControlEventsRequest) GetDomain() *string {
	return s.Domain
}

func (s *QueryResourceControlEventsRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *QueryResourceControlEventsRequest) GetEventCodes() []*string {
	return s.EventCodes
}

func (s *QueryResourceControlEventsRequest) GetEventId() *string {
	return s.EventId
}

func (s *QueryResourceControlEventsRequest) GetEventIdList() []*string {
	return s.EventIdList
}

func (s *QueryResourceControlEventsRequest) GetExcludeActionCodes() []*string {
	return s.ExcludeActionCodes
}

func (s *QueryResourceControlEventsRequest) GetExcludeEventCodes() []*string {
	return s.ExcludeEventCodes
}

func (s *QueryResourceControlEventsRequest) GetExcludeReasons() []*string {
	return s.ExcludeReasons
}

func (s *QueryResourceControlEventsRequest) GetIncludeReasons() []*string {
	return s.IncludeReasons
}

func (s *QueryResourceControlEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryResourceControlEventsRequest) GetIp() *string {
	return s.Ip
}

func (s *QueryResourceControlEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryResourceControlEventsRequest) GetPunishEndTime() *string {
	return s.PunishEndTime
}

func (s *QueryResourceControlEventsRequest) GetPunishStartTime() *string {
	return s.PunishStartTime
}

func (s *QueryResourceControlEventsRequest) GetReason() *string {
	return s.Reason
}

func (s *QueryResourceControlEventsRequest) GetSourceCodes() []*string {
	return s.SourceCodes
}

func (s *QueryResourceControlEventsRequest) GetStatus() *string {
	return s.Status
}

func (s *QueryResourceControlEventsRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *QueryResourceControlEventsRequest) GetUrl() *string {
	return s.Url
}

func (s *QueryResourceControlEventsRequest) SetActionCode(v string) *QueryResourceControlEventsRequest {
	s.ActionCode = &v
	return s
}

func (s *QueryResourceControlEventsRequest) SetActionCodes(v []*string) *QueryResourceControlEventsRequest {
	s.ActionCodes = v
	return s
}

func (s *QueryResourceControlEventsRequest) SetAliyunLang(v string) *QueryResourceControlEventsRequest {
	s.AliyunLang = &v
	return s
}

func (s *QueryResourceControlEventsRequest) SetBusinessCode(v string) *QueryResourceControlEventsRequest {
	s.BusinessCode = &v
	return s
}

func (s *QueryResourceControlEventsRequest) SetBusinessCodes(v []*string) *QueryResourceControlEventsRequest {
	s.BusinessCodes = v
	return s
}

func (s *QueryResourceControlEventsRequest) SetCaseCodesPrefix(v []*string) *QueryResourceControlEventsRequest {
	s.CaseCodesPrefix = v
	return s
}

func (s *QueryResourceControlEventsRequest) SetCurrent(v int32) *QueryResourceControlEventsRequest {
	s.Current = &v
	return s
}

func (s *QueryResourceControlEventsRequest) SetDomain(v string) *QueryResourceControlEventsRequest {
	s.Domain = &v
	return s
}

func (s *QueryResourceControlEventsRequest) SetEventCode(v string) *QueryResourceControlEventsRequest {
	s.EventCode = &v
	return s
}

func (s *QueryResourceControlEventsRequest) SetEventCodes(v []*string) *QueryResourceControlEventsRequest {
	s.EventCodes = v
	return s
}

func (s *QueryResourceControlEventsRequest) SetEventId(v string) *QueryResourceControlEventsRequest {
	s.EventId = &v
	return s
}

func (s *QueryResourceControlEventsRequest) SetEventIdList(v []*string) *QueryResourceControlEventsRequest {
	s.EventIdList = v
	return s
}

func (s *QueryResourceControlEventsRequest) SetExcludeActionCodes(v []*string) *QueryResourceControlEventsRequest {
	s.ExcludeActionCodes = v
	return s
}

func (s *QueryResourceControlEventsRequest) SetExcludeEventCodes(v []*string) *QueryResourceControlEventsRequest {
	s.ExcludeEventCodes = v
	return s
}

func (s *QueryResourceControlEventsRequest) SetExcludeReasons(v []*string) *QueryResourceControlEventsRequest {
	s.ExcludeReasons = v
	return s
}

func (s *QueryResourceControlEventsRequest) SetIncludeReasons(v []*string) *QueryResourceControlEventsRequest {
	s.IncludeReasons = v
	return s
}

func (s *QueryResourceControlEventsRequest) SetInstanceId(v string) *QueryResourceControlEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryResourceControlEventsRequest) SetIp(v string) *QueryResourceControlEventsRequest {
	s.Ip = &v
	return s
}

func (s *QueryResourceControlEventsRequest) SetPageSize(v int32) *QueryResourceControlEventsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryResourceControlEventsRequest) SetPunishEndTime(v string) *QueryResourceControlEventsRequest {
	s.PunishEndTime = &v
	return s
}

func (s *QueryResourceControlEventsRequest) SetPunishStartTime(v string) *QueryResourceControlEventsRequest {
	s.PunishStartTime = &v
	return s
}

func (s *QueryResourceControlEventsRequest) SetReason(v string) *QueryResourceControlEventsRequest {
	s.Reason = &v
	return s
}

func (s *QueryResourceControlEventsRequest) SetSourceCodes(v []*string) *QueryResourceControlEventsRequest {
	s.SourceCodes = v
	return s
}

func (s *QueryResourceControlEventsRequest) SetStatus(v string) *QueryResourceControlEventsRequest {
	s.Status = &v
	return s
}

func (s *QueryResourceControlEventsRequest) SetStatusList(v []*string) *QueryResourceControlEventsRequest {
	s.StatusList = v
	return s
}

func (s *QueryResourceControlEventsRequest) SetUrl(v string) *QueryResourceControlEventsRequest {
	s.Url = &v
	return s
}

func (s *QueryResourceControlEventsRequest) Validate() error {
	return dara.Validate(s)
}
