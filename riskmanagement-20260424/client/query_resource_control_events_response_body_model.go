// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResourceControlEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryResourceControlEventsResponseBody
	GetCode() *string
	SetData(v *QueryResourceControlEventsResponseBodyData) *QueryResourceControlEventsResponseBody
	GetData() *QueryResourceControlEventsResponseBodyData
	SetMessage(v string) *QueryResourceControlEventsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryResourceControlEventsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryResourceControlEventsResponseBody
	GetSuccess() *bool
}

type QueryResourceControlEventsResponseBody struct {
	// example:
	//
	// 200
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryResourceControlEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 739705BB-B0EF-554B-B3A8-383F4F93E067
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryResourceControlEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryResourceControlEventsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryResourceControlEventsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryResourceControlEventsResponseBody) GetData() *QueryResourceControlEventsResponseBodyData {
	return s.Data
}

func (s *QueryResourceControlEventsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryResourceControlEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryResourceControlEventsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryResourceControlEventsResponseBody) SetCode(v string) *QueryResourceControlEventsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryResourceControlEventsResponseBody) SetData(v *QueryResourceControlEventsResponseBodyData) *QueryResourceControlEventsResponseBody {
	s.Data = v
	return s
}

func (s *QueryResourceControlEventsResponseBody) SetMessage(v string) *QueryResourceControlEventsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryResourceControlEventsResponseBody) SetRequestId(v string) *QueryResourceControlEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryResourceControlEventsResponseBody) SetSuccess(v bool) *QueryResourceControlEventsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryResourceControlEventsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryResourceControlEventsResponseBodyData struct {
	List     []*QueryResourceControlEventsResponseBodyDataList   `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageInfo *QueryResourceControlEventsResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
}

func (s QueryResourceControlEventsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryResourceControlEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryResourceControlEventsResponseBodyData) GetList() []*QueryResourceControlEventsResponseBodyDataList {
	return s.List
}

func (s *QueryResourceControlEventsResponseBodyData) GetPageInfo() *QueryResourceControlEventsResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *QueryResourceControlEventsResponseBodyData) SetList(v []*QueryResourceControlEventsResponseBodyDataList) *QueryResourceControlEventsResponseBodyData {
	s.List = v
	return s
}

func (s *QueryResourceControlEventsResponseBodyData) SetPageInfo(v *QueryResourceControlEventsResponseBodyDataPageInfo) *QueryResourceControlEventsResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *QueryResourceControlEventsResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryResourceControlEventsResponseBodyDataList struct {
	// example:
	//
	// DEPLOY_STAGE_REBOOT_TASK
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	// example:
	//
	// 挖矿告警
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	AlertEndTime *string `json:"AlertEndTime,omitempty" xml:"AlertEndTime,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	AlertStartTime *string `json:"AlertStartTime,omitempty" xml:"AlertStartTime,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	AntiPunishTime *string `json:"AntiPunishTime,omitempty" xml:"AntiPunishTime,omitempty"`
	// example:
	//
	// 1
	ApplyRecordCount *int32 `json:"ApplyRecordCount,omitempty" xml:"ApplyRecordCount,omitempty"`
	// example:
	//
	// AUDIT
	ApplyStatus *string `json:"ApplyStatus,omitempty" xml:"ApplyStatus,omitempty"`
	// example:
	//
	// false
	ApplyTrial *bool `json:"ApplyTrial,omitempty" xml:"ApplyTrial,omitempty"`
	// example:
	//
	// e\\"c\\"s
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// example:
	//
	// TEST_IMS_ACCOUNT_PUNISH_WHITE_TEST
	CaseCode *string `json:"CaseCode,omitempty" xml:"CaseCode,omitempty"`
	// example:
	//
	// ubs-mm-nwwss-ddos.purchern.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// e791c08281b41e8240f897a424c188ae
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// 挖矿管控事件
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// {\\"createAt\\":\\"2025-08-03 11:18:59\\",\\"updatedAt\\":\\"2025-08-03 11:18:59\\"}
	Extras *string `json:"Extras,omitempty" xml:"Extras,omitempty"`
	// example:
	//
	// 3
	FormType *string `json:"FormType,omitempty" xml:"FormType,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	GmtLatest *string `json:"GmtLatest,omitempty" xml:"GmtLatest,omitempty"`
	// example:
	//
	// rm-3nsvwmt67pn72py1z
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 10.0.158.58
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	LastCheckTime *string `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	PreCloseTime *string `json:"PreCloseTime,omitempty" xml:"PreCloseTime,omitempty"`
	// example:
	//
	// MRM
	PunishFrom *string `json:"PunishFrom,omitempty" xml:"PunishFrom,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	PunishTime *string `json:"PunishTime,omitempty" xml:"PunishTime,omitempty"`
	// example:
	//
	// 挖矿告警
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Suggestion
	Reinforcement *string `json:"Reinforcement,omitempty" xml:"Reinforcement,omitempty"`
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	SupportBatchApply *bool `json:"SupportBatchApply,omitempty" xml:"SupportBatchApply,omitempty"`
	// example:
	//
	// true
	SupportSingleApply *bool `json:"SupportSingleApply,omitempty" xml:"SupportSingleApply,omitempty"`
	// example:
	//
	// miner
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// example:
	//
	// https://pm.alicdn.com/quali/bc98e42b619ad4127bf6437b87045597.jpg?auth_key=1758682451-0-0-897be72852503566bd6775cd9914f5aa
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryResourceControlEventsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryResourceControlEventsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetActionCode() *string {
	return s.ActionCode
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetActionName() *string {
	return s.ActionName
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetAlertEndTime() *string {
	return s.AlertEndTime
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetAlertStartTime() *string {
	return s.AlertStartTime
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetAntiPunishTime() *string {
	return s.AntiPunishTime
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetApplyRecordCount() *int32 {
	return s.ApplyRecordCount
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetApplyStatus() *string {
	return s.ApplyStatus
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetApplyTrial() *bool {
	return s.ApplyTrial
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetBusinessName() *string {
	return s.BusinessName
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetCaseCode() *string {
	return s.CaseCode
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetDomain() *string {
	return s.Domain
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetEventId() *string {
	return s.EventId
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetEventName() *string {
	return s.EventName
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetExtras() *string {
	return s.Extras
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetFormType() *string {
	return s.FormType
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetGmtLatest() *string {
	return s.GmtLatest
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetIp() *string {
	return s.Ip
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetLastCheckTime() *string {
	return s.LastCheckTime
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetPreCloseTime() *string {
	return s.PreCloseTime
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetPunishFrom() *string {
	return s.PunishFrom
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetPunishTime() *string {
	return s.PunishTime
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetReason() *string {
	return s.Reason
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetRegion() *string {
	return s.Region
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetReinforcement() *string {
	return s.Reinforcement
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetSupportBatchApply() *bool {
	return s.SupportBatchApply
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetSupportSingleApply() *bool {
	return s.SupportSingleApply
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetTriggerType() *string {
	return s.TriggerType
}

func (s *QueryResourceControlEventsResponseBodyDataList) GetUrl() *string {
	return s.Url
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetActionCode(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.ActionCode = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetActionName(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.ActionName = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetAlertEndTime(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.AlertEndTime = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetAlertStartTime(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.AlertStartTime = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetAntiPunishTime(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.AntiPunishTime = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetApplyRecordCount(v int32) *QueryResourceControlEventsResponseBodyDataList {
	s.ApplyRecordCount = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetApplyStatus(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.ApplyStatus = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetApplyTrial(v bool) *QueryResourceControlEventsResponseBodyDataList {
	s.ApplyTrial = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetBusinessName(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.BusinessName = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetCaseCode(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.CaseCode = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetDomain(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.Domain = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetEventId(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.EventId = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetEventName(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.EventName = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetExtras(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.Extras = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetFormType(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.FormType = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetGmtLatest(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.GmtLatest = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetInstanceId(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetIp(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.Ip = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetLastCheckTime(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.LastCheckTime = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetPreCloseTime(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.PreCloseTime = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetPunishFrom(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.PunishFrom = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetPunishTime(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.PunishTime = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetReason(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.Reason = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetRegion(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.Region = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetRegionId(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetReinforcement(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.Reinforcement = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetStatus(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetSupportBatchApply(v bool) *QueryResourceControlEventsResponseBodyDataList {
	s.SupportBatchApply = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetSupportSingleApply(v bool) *QueryResourceControlEventsResponseBodyDataList {
	s.SupportSingleApply = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetTriggerType(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.TriggerType = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) SetUrl(v string) *QueryResourceControlEventsResponseBodyDataList {
	s.Url = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type QueryResourceControlEventsResponseBodyDataPageInfo struct {
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// 24
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryResourceControlEventsResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryResourceControlEventsResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *QueryResourceControlEventsResponseBodyDataPageInfo) GetCurrent() *int32 {
	return s.Current
}

func (s *QueryResourceControlEventsResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryResourceControlEventsResponseBodyDataPageInfo) GetTotal() *int32 {
	return s.Total
}

func (s *QueryResourceControlEventsResponseBodyDataPageInfo) SetCurrent(v int32) *QueryResourceControlEventsResponseBodyDataPageInfo {
	s.Current = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataPageInfo) SetPageSize(v int32) *QueryResourceControlEventsResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataPageInfo) SetTotal(v int32) *QueryResourceControlEventsResponseBodyDataPageInfo {
	s.Total = &v
	return s
}

func (s *QueryResourceControlEventsResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}
