// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadResourceControlEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionCode(v string) *DownloadResourceControlEventsRequest
	GetActionCode() *string
	SetActionCodes(v []*string) *DownloadResourceControlEventsRequest
	GetActionCodes() []*string
	SetAliyunLang(v string) *DownloadResourceControlEventsRequest
	GetAliyunLang() *string
	SetBusinessCode(v string) *DownloadResourceControlEventsRequest
	GetBusinessCode() *string
	SetCaseCodesPrefix(v []*string) *DownloadResourceControlEventsRequest
	GetCaseCodesPrefix() []*string
	SetCurrent(v int32) *DownloadResourceControlEventsRequest
	GetCurrent() *int32
	SetDomain(v string) *DownloadResourceControlEventsRequest
	GetDomain() *string
	SetEventCode(v string) *DownloadResourceControlEventsRequest
	GetEventCode() *string
	SetEventCodes(v []*string) *DownloadResourceControlEventsRequest
	GetEventCodes() []*string
	SetEventId(v string) *DownloadResourceControlEventsRequest
	GetEventId() *string
	SetExcludeActionCodes(v []*string) *DownloadResourceControlEventsRequest
	GetExcludeActionCodes() []*string
	SetExcludeEventCodes(v []*string) *DownloadResourceControlEventsRequest
	GetExcludeEventCodes() []*string
	SetExcludeReasons(v []*string) *DownloadResourceControlEventsRequest
	GetExcludeReasons() []*string
	SetIncludeReasons(v []*string) *DownloadResourceControlEventsRequest
	GetIncludeReasons() []*string
	SetInstanceId(v string) *DownloadResourceControlEventsRequest
	GetInstanceId() *string
	SetIp(v string) *DownloadResourceControlEventsRequest
	GetIp() *string
	SetPageSize(v int32) *DownloadResourceControlEventsRequest
	GetPageSize() *int32
	SetPunishEndTime(v string) *DownloadResourceControlEventsRequest
	GetPunishEndTime() *string
	SetPunishStartTime(v string) *DownloadResourceControlEventsRequest
	GetPunishStartTime() *string
	SetReason(v string) *DownloadResourceControlEventsRequest
	GetReason() *string
	SetSourceCodes(v []*string) *DownloadResourceControlEventsRequest
	GetSourceCodes() []*string
	SetStatus(v string) *DownloadResourceControlEventsRequest
	GetStatus() *string
	SetStatusList(v []*string) *DownloadResourceControlEventsRequest
	GetStatusList() []*string
	SetUrl(v string) *DownloadResourceControlEventsRequest
	GetUrl() *string
}

type DownloadResourceControlEventsRequest struct {
	// The action name code.
	//
	// example:
	//
	// shutdown
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	// The collection of control action name codes.
	//
	// > Example: [\\\\\\"shutdown\\\\\\"]
	ActionCodes []*string `json:"ActionCodes,omitempty" xml:"ActionCodes,omitempty" type:"Repeated"`
	// The language. Valid values:
	//
	// - **zh**: Chinese (default).
	//
	// - **en**: English.
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
	BusinessCode *string `json:"BusinessCode,omitempty" xml:"BusinessCode,omitempty"`
	// The collection of event name code prefixes.
	//
	// > Example: [\\\\\\"BANFF\\\\\\"]
	CaseCodesPrefix []*string `json:"CaseCodesPrefix,omitempty" xml:"CaseCodesPrefix,omitempty" type:"Repeated"`
	// The current page number.
	//
	// > Must be greater than 0.
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
	// > Example: [\\\\\\"BANFF_ECS_PE_ECS_MINING_SHUTDOWN\\\\\\"]
	EventCodes []*string `json:"EventCodes,omitempty" xml:"EventCodes,omitempty" type:"Repeated"`
	// The alert event ID.
	//
	// example:
	//
	// 09C-2PpwIzkpx2zG2fuFrAH55CpJaTK
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The collection of excluded control action name codes.
	//
	// > Example: [\\\\\\"shutdown\\\\\\"]
	ExcludeActionCodes []*string `json:"ExcludeActionCodes,omitempty" xml:"ExcludeActionCodes,omitempty" type:"Repeated"`
	// The collection of excluded event name codes.
	//
	// > Example: [\\\\\\"TEST_CASE\\\\\\"]
	ExcludeEventCodes []*string `json:"ExcludeEventCodes,omitempty" xml:"ExcludeEventCodes,omitempty" type:"Repeated"`
	// The collection of excluded event reasons.
	//
	// > Example: [\\\\\\"Mining alert\\\\\\"]
	ExcludeReasons []*string `json:"ExcludeReasons,omitempty" xml:"ExcludeReasons,omitempty" type:"Repeated"`
	// The collection of included event reasons.
	//
	// > Example: [\\\\\\"Mining alert\\\\\\"]
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
	// > Format: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-03-16 15:15:00
	PunishEndTime *string `json:"PunishEndTime,omitempty" xml:"PunishEndTime,omitempty"`
	// The penalty start time.
	//
	// > Format: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-03-16 15:15:00
	PunishStartTime *string `json:"PunishStartTime,omitempty" xml:"PunishStartTime,omitempty"`
	// The event reason.
	//
	// example:
	//
	// Mining
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The collection of event source codes.
	//
	// > Example: [\\\\\\"MRM\\\\\\"]
	SourceCodes []*string `json:"SourceCodes,omitempty" xml:"SourceCodes,omitempty" type:"Repeated"`
	// The task status.
	//
	// - **Executing**: In progress.
	//
	// - **Removed**: Removed.
	//
	// - **Alerting**: Alerting.
	//
	// - **Ended**: Ended.
	//
	// - **Processed**: Processed by the user and pending platform review.
	//
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The collection of task statuses.
	//
	// - **Executing**: In progress.
	//
	// - **Removed**: Removed.
	//
	// - **Alerting**: Alerting.
	//
	// - **Ended**: Ended.
	//
	// - **Processed**: Processed by the user and pending platform review.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// The control URL.
	//
	// example:
	//
	// https://qimg.xiaohongshu.com/circe/1040g1v831qggp28ln0705oft1i6k1jil889lhso?imageView2/2/w/1080/format/jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DownloadResourceControlEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadResourceControlEventsRequest) GoString() string {
	return s.String()
}

func (s *DownloadResourceControlEventsRequest) GetActionCode() *string {
	return s.ActionCode
}

func (s *DownloadResourceControlEventsRequest) GetActionCodes() []*string {
	return s.ActionCodes
}

func (s *DownloadResourceControlEventsRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *DownloadResourceControlEventsRequest) GetBusinessCode() *string {
	return s.BusinessCode
}

func (s *DownloadResourceControlEventsRequest) GetCaseCodesPrefix() []*string {
	return s.CaseCodesPrefix
}

func (s *DownloadResourceControlEventsRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *DownloadResourceControlEventsRequest) GetDomain() *string {
	return s.Domain
}

func (s *DownloadResourceControlEventsRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DownloadResourceControlEventsRequest) GetEventCodes() []*string {
	return s.EventCodes
}

func (s *DownloadResourceControlEventsRequest) GetEventId() *string {
	return s.EventId
}

func (s *DownloadResourceControlEventsRequest) GetExcludeActionCodes() []*string {
	return s.ExcludeActionCodes
}

func (s *DownloadResourceControlEventsRequest) GetExcludeEventCodes() []*string {
	return s.ExcludeEventCodes
}

func (s *DownloadResourceControlEventsRequest) GetExcludeReasons() []*string {
	return s.ExcludeReasons
}

func (s *DownloadResourceControlEventsRequest) GetIncludeReasons() []*string {
	return s.IncludeReasons
}

func (s *DownloadResourceControlEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DownloadResourceControlEventsRequest) GetIp() *string {
	return s.Ip
}

func (s *DownloadResourceControlEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DownloadResourceControlEventsRequest) GetPunishEndTime() *string {
	return s.PunishEndTime
}

func (s *DownloadResourceControlEventsRequest) GetPunishStartTime() *string {
	return s.PunishStartTime
}

func (s *DownloadResourceControlEventsRequest) GetReason() *string {
	return s.Reason
}

func (s *DownloadResourceControlEventsRequest) GetSourceCodes() []*string {
	return s.SourceCodes
}

func (s *DownloadResourceControlEventsRequest) GetStatus() *string {
	return s.Status
}

func (s *DownloadResourceControlEventsRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *DownloadResourceControlEventsRequest) GetUrl() *string {
	return s.Url
}

func (s *DownloadResourceControlEventsRequest) SetActionCode(v string) *DownloadResourceControlEventsRequest {
	s.ActionCode = &v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetActionCodes(v []*string) *DownloadResourceControlEventsRequest {
	s.ActionCodes = v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetAliyunLang(v string) *DownloadResourceControlEventsRequest {
	s.AliyunLang = &v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetBusinessCode(v string) *DownloadResourceControlEventsRequest {
	s.BusinessCode = &v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetCaseCodesPrefix(v []*string) *DownloadResourceControlEventsRequest {
	s.CaseCodesPrefix = v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetCurrent(v int32) *DownloadResourceControlEventsRequest {
	s.Current = &v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetDomain(v string) *DownloadResourceControlEventsRequest {
	s.Domain = &v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetEventCode(v string) *DownloadResourceControlEventsRequest {
	s.EventCode = &v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetEventCodes(v []*string) *DownloadResourceControlEventsRequest {
	s.EventCodes = v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetEventId(v string) *DownloadResourceControlEventsRequest {
	s.EventId = &v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetExcludeActionCodes(v []*string) *DownloadResourceControlEventsRequest {
	s.ExcludeActionCodes = v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetExcludeEventCodes(v []*string) *DownloadResourceControlEventsRequest {
	s.ExcludeEventCodes = v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetExcludeReasons(v []*string) *DownloadResourceControlEventsRequest {
	s.ExcludeReasons = v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetIncludeReasons(v []*string) *DownloadResourceControlEventsRequest {
	s.IncludeReasons = v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetInstanceId(v string) *DownloadResourceControlEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetIp(v string) *DownloadResourceControlEventsRequest {
	s.Ip = &v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetPageSize(v int32) *DownloadResourceControlEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetPunishEndTime(v string) *DownloadResourceControlEventsRequest {
	s.PunishEndTime = &v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetPunishStartTime(v string) *DownloadResourceControlEventsRequest {
	s.PunishStartTime = &v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetReason(v string) *DownloadResourceControlEventsRequest {
	s.Reason = &v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetSourceCodes(v []*string) *DownloadResourceControlEventsRequest {
	s.SourceCodes = v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetStatus(v string) *DownloadResourceControlEventsRequest {
	s.Status = &v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetStatusList(v []*string) *DownloadResourceControlEventsRequest {
	s.StatusList = v
	return s
}

func (s *DownloadResourceControlEventsRequest) SetUrl(v string) *DownloadResourceControlEventsRequest {
	s.Url = &v
	return s
}

func (s *DownloadResourceControlEventsRequest) Validate() error {
	return dara.Validate(s)
}
