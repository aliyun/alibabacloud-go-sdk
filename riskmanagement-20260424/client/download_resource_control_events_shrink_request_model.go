// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadResourceControlEventsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionCode(v string) *DownloadResourceControlEventsShrinkRequest
	GetActionCode() *string
	SetActionCodesShrink(v string) *DownloadResourceControlEventsShrinkRequest
	GetActionCodesShrink() *string
	SetAliyunLang(v string) *DownloadResourceControlEventsShrinkRequest
	GetAliyunLang() *string
	SetBusinessCode(v string) *DownloadResourceControlEventsShrinkRequest
	GetBusinessCode() *string
	SetCaseCodesPrefixShrink(v string) *DownloadResourceControlEventsShrinkRequest
	GetCaseCodesPrefixShrink() *string
	SetCurrent(v int32) *DownloadResourceControlEventsShrinkRequest
	GetCurrent() *int32
	SetDomain(v string) *DownloadResourceControlEventsShrinkRequest
	GetDomain() *string
	SetEventCode(v string) *DownloadResourceControlEventsShrinkRequest
	GetEventCode() *string
	SetEventCodesShrink(v string) *DownloadResourceControlEventsShrinkRequest
	GetEventCodesShrink() *string
	SetEventId(v string) *DownloadResourceControlEventsShrinkRequest
	GetEventId() *string
	SetExcludeActionCodesShrink(v string) *DownloadResourceControlEventsShrinkRequest
	GetExcludeActionCodesShrink() *string
	SetExcludeEventCodesShrink(v string) *DownloadResourceControlEventsShrinkRequest
	GetExcludeEventCodesShrink() *string
	SetExcludeReasonsShrink(v string) *DownloadResourceControlEventsShrinkRequest
	GetExcludeReasonsShrink() *string
	SetIncludeReasonsShrink(v string) *DownloadResourceControlEventsShrinkRequest
	GetIncludeReasonsShrink() *string
	SetInstanceId(v string) *DownloadResourceControlEventsShrinkRequest
	GetInstanceId() *string
	SetIp(v string) *DownloadResourceControlEventsShrinkRequest
	GetIp() *string
	SetPageSize(v int32) *DownloadResourceControlEventsShrinkRequest
	GetPageSize() *int32
	SetPunishEndTime(v string) *DownloadResourceControlEventsShrinkRequest
	GetPunishEndTime() *string
	SetPunishStartTime(v string) *DownloadResourceControlEventsShrinkRequest
	GetPunishStartTime() *string
	SetReason(v string) *DownloadResourceControlEventsShrinkRequest
	GetReason() *string
	SetSourceCodesShrink(v string) *DownloadResourceControlEventsShrinkRequest
	GetSourceCodesShrink() *string
	SetStatus(v string) *DownloadResourceControlEventsShrinkRequest
	GetStatus() *string
	SetStatusListShrink(v string) *DownloadResourceControlEventsShrinkRequest
	GetStatusListShrink() *string
	SetUrl(v string) *DownloadResourceControlEventsShrinkRequest
	GetUrl() *string
}

type DownloadResourceControlEventsShrinkRequest struct {
	// The action name code.
	//
	// example:
	//
	// shutdown
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	// The collection of control action name codes.
	//
	// > Example: [\\\\\\"shutdown\\\\\\"]
	ActionCodesShrink *string `json:"ActionCodes,omitempty" xml:"ActionCodes,omitempty"`
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
	CaseCodesPrefixShrink *string `json:"CaseCodesPrefix,omitempty" xml:"CaseCodesPrefix,omitempty"`
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
	EventCodesShrink *string `json:"EventCodes,omitempty" xml:"EventCodes,omitempty"`
	// The alert event ID.
	//
	// example:
	//
	// 09C-2PpwIzkpx2zG2fuFrAH55CpJaTK
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The collection of excluded control action name codes.
	//
	// > Example: [\\\\\\"shutdown\\\\\\"]
	ExcludeActionCodesShrink *string `json:"ExcludeActionCodes,omitempty" xml:"ExcludeActionCodes,omitempty"`
	// The collection of excluded event name codes.
	//
	// > Example: [\\\\\\"TEST_CASE\\\\\\"]
	ExcludeEventCodesShrink *string `json:"ExcludeEventCodes,omitempty" xml:"ExcludeEventCodes,omitempty"`
	// The collection of excluded event reasons.
	//
	// > Example: [\\\\\\"Mining alert\\\\\\"]
	ExcludeReasonsShrink *string `json:"ExcludeReasons,omitempty" xml:"ExcludeReasons,omitempty"`
	// The collection of included event reasons.
	//
	// > Example: [\\\\\\"Mining alert\\\\\\"]
	IncludeReasonsShrink *string `json:"IncludeReasons,omitempty" xml:"IncludeReasons,omitempty"`
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
	SourceCodesShrink *string `json:"SourceCodes,omitempty" xml:"SourceCodes,omitempty"`
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
	StatusListShrink *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	// The control URL.
	//
	// example:
	//
	// https://qimg.xiaohongshu.com/circe/1040g1v831qggp28ln0705oft1i6k1jil889lhso?imageView2/2/w/1080/format/jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DownloadResourceControlEventsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadResourceControlEventsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DownloadResourceControlEventsShrinkRequest) GetActionCode() *string {
	return s.ActionCode
}

func (s *DownloadResourceControlEventsShrinkRequest) GetActionCodesShrink() *string {
	return s.ActionCodesShrink
}

func (s *DownloadResourceControlEventsShrinkRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *DownloadResourceControlEventsShrinkRequest) GetBusinessCode() *string {
	return s.BusinessCode
}

func (s *DownloadResourceControlEventsShrinkRequest) GetCaseCodesPrefixShrink() *string {
	return s.CaseCodesPrefixShrink
}

func (s *DownloadResourceControlEventsShrinkRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *DownloadResourceControlEventsShrinkRequest) GetDomain() *string {
	return s.Domain
}

func (s *DownloadResourceControlEventsShrinkRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DownloadResourceControlEventsShrinkRequest) GetEventCodesShrink() *string {
	return s.EventCodesShrink
}

func (s *DownloadResourceControlEventsShrinkRequest) GetEventId() *string {
	return s.EventId
}

func (s *DownloadResourceControlEventsShrinkRequest) GetExcludeActionCodesShrink() *string {
	return s.ExcludeActionCodesShrink
}

func (s *DownloadResourceControlEventsShrinkRequest) GetExcludeEventCodesShrink() *string {
	return s.ExcludeEventCodesShrink
}

func (s *DownloadResourceControlEventsShrinkRequest) GetExcludeReasonsShrink() *string {
	return s.ExcludeReasonsShrink
}

func (s *DownloadResourceControlEventsShrinkRequest) GetIncludeReasonsShrink() *string {
	return s.IncludeReasonsShrink
}

func (s *DownloadResourceControlEventsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DownloadResourceControlEventsShrinkRequest) GetIp() *string {
	return s.Ip
}

func (s *DownloadResourceControlEventsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DownloadResourceControlEventsShrinkRequest) GetPunishEndTime() *string {
	return s.PunishEndTime
}

func (s *DownloadResourceControlEventsShrinkRequest) GetPunishStartTime() *string {
	return s.PunishStartTime
}

func (s *DownloadResourceControlEventsShrinkRequest) GetReason() *string {
	return s.Reason
}

func (s *DownloadResourceControlEventsShrinkRequest) GetSourceCodesShrink() *string {
	return s.SourceCodesShrink
}

func (s *DownloadResourceControlEventsShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *DownloadResourceControlEventsShrinkRequest) GetStatusListShrink() *string {
	return s.StatusListShrink
}

func (s *DownloadResourceControlEventsShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *DownloadResourceControlEventsShrinkRequest) SetActionCode(v string) *DownloadResourceControlEventsShrinkRequest {
	s.ActionCode = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetActionCodesShrink(v string) *DownloadResourceControlEventsShrinkRequest {
	s.ActionCodesShrink = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetAliyunLang(v string) *DownloadResourceControlEventsShrinkRequest {
	s.AliyunLang = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetBusinessCode(v string) *DownloadResourceControlEventsShrinkRequest {
	s.BusinessCode = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetCaseCodesPrefixShrink(v string) *DownloadResourceControlEventsShrinkRequest {
	s.CaseCodesPrefixShrink = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetCurrent(v int32) *DownloadResourceControlEventsShrinkRequest {
	s.Current = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetDomain(v string) *DownloadResourceControlEventsShrinkRequest {
	s.Domain = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetEventCode(v string) *DownloadResourceControlEventsShrinkRequest {
	s.EventCode = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetEventCodesShrink(v string) *DownloadResourceControlEventsShrinkRequest {
	s.EventCodesShrink = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetEventId(v string) *DownloadResourceControlEventsShrinkRequest {
	s.EventId = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetExcludeActionCodesShrink(v string) *DownloadResourceControlEventsShrinkRequest {
	s.ExcludeActionCodesShrink = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetExcludeEventCodesShrink(v string) *DownloadResourceControlEventsShrinkRequest {
	s.ExcludeEventCodesShrink = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetExcludeReasonsShrink(v string) *DownloadResourceControlEventsShrinkRequest {
	s.ExcludeReasonsShrink = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetIncludeReasonsShrink(v string) *DownloadResourceControlEventsShrinkRequest {
	s.IncludeReasonsShrink = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetInstanceId(v string) *DownloadResourceControlEventsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetIp(v string) *DownloadResourceControlEventsShrinkRequest {
	s.Ip = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetPageSize(v int32) *DownloadResourceControlEventsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetPunishEndTime(v string) *DownloadResourceControlEventsShrinkRequest {
	s.PunishEndTime = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetPunishStartTime(v string) *DownloadResourceControlEventsShrinkRequest {
	s.PunishStartTime = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetReason(v string) *DownloadResourceControlEventsShrinkRequest {
	s.Reason = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetSourceCodesShrink(v string) *DownloadResourceControlEventsShrinkRequest {
	s.SourceCodesShrink = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetStatus(v string) *DownloadResourceControlEventsShrinkRequest {
	s.Status = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetStatusListShrink(v string) *DownloadResourceControlEventsShrinkRequest {
	s.StatusListShrink = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) SetUrl(v string) *DownloadResourceControlEventsShrinkRequest {
	s.Url = &v
	return s
}

func (s *DownloadResourceControlEventsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
