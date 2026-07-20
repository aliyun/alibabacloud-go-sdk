// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceControlEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetResourceControlEventResponseBody
	GetCode() *string
	SetData(v *GetResourceControlEventResponseBodyData) *GetResourceControlEventResponseBody
	GetData() *GetResourceControlEventResponseBodyData
	SetMessage(v string) *GetResourceControlEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetResourceControlEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetResourceControlEventResponseBody
	GetSuccess() *bool
}

type GetResourceControlEventResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetResourceControlEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful‌
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6B57D35D-9DAC-5393-AE39-07697E37C2E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetResourceControlEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceControlEventResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceControlEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetResourceControlEventResponseBody) GetData() *GetResourceControlEventResponseBodyData {
	return s.Data
}

func (s *GetResourceControlEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetResourceControlEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceControlEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetResourceControlEventResponseBody) SetCode(v string) *GetResourceControlEventResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourceControlEventResponseBody) SetData(v *GetResourceControlEventResponseBodyData) *GetResourceControlEventResponseBody {
	s.Data = v
	return s
}

func (s *GetResourceControlEventResponseBody) SetMessage(v string) *GetResourceControlEventResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourceControlEventResponseBody) SetRequestId(v string) *GetResourceControlEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceControlEventResponseBody) SetSuccess(v bool) *GetResourceControlEventResponseBody {
	s.Success = &v
	return s
}

func (s *GetResourceControlEventResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceControlEventResponseBodyData struct {
	ApplyRecordList []*GetResourceControlEventResponseBodyDataApplyRecordList `json:"ApplyRecordList,omitempty" xml:"ApplyRecordList,omitempty" type:"Repeated"`
	// example:
	//
	// assistant tip
	AssistantTip *string `json:"AssistantTip,omitempty" xml:"AssistantTip,omitempty"`
	// example:
	//
	// 196.251.81.30
	BlockIp *string `json:"BlockIp,omitempty" xml:"BlockIp,omitempty"`
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// 10.199.31.155
	DstIp *string `json:"DstIp,omitempty" xml:"DstIp,omitempty"`
	// example:
	//
	// 30629
	DstPort *string `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// example:
	//
	// 57ed8c6ddc9aafb1a3df38e6e84d2d45
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// Instance Stopped
	EventImpact *string `json:"EventImpact,omitempty" xml:"EventImpact,omitempty"`
	// example:
	//
	// Mining Management Event
	LeakName *string `json:"LeakName,omitempty" xml:"LeakName,omitempty"`
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// This instance is connecting to a Miner Pool and has likely been compromised by hackers for mining or other illicit activities.
	PunishReason *string `json:"PunishReason,omitempty" xml:"PunishReason,omitempty"`
	// example:
	//
	// https://xxx.aliyun.com/v2
	SnapshotUrl *string `json:"SnapshotUrl,omitempty" xml:"SnapshotUrl,omitempty"`
	// example:
	//
	// 36.134.124.185
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// example:
	//
	// 2168
	SrcPort *string `json:"SrcPort,omitempty" xml:"SrcPort,omitempty"`
	// example:
	//
	// Suggestion
	Tip *string `json:"Tip,omitempty" xml:"Tip,omitempty"`
}

func (s GetResourceControlEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetResourceControlEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourceControlEventResponseBodyData) GetApplyRecordList() []*GetResourceControlEventResponseBodyDataApplyRecordList {
	return s.ApplyRecordList
}

func (s *GetResourceControlEventResponseBodyData) GetAssistantTip() *string {
	return s.AssistantTip
}

func (s *GetResourceControlEventResponseBodyData) GetBlockIp() *string {
	return s.BlockIp
}

func (s *GetResourceControlEventResponseBodyData) GetDirection() *string {
	return s.Direction
}

func (s *GetResourceControlEventResponseBodyData) GetDstIp() *string {
	return s.DstIp
}

func (s *GetResourceControlEventResponseBodyData) GetDstPort() *string {
	return s.DstPort
}

func (s *GetResourceControlEventResponseBodyData) GetEventId() *string {
	return s.EventId
}

func (s *GetResourceControlEventResponseBodyData) GetEventImpact() *string {
	return s.EventImpact
}

func (s *GetResourceControlEventResponseBodyData) GetLeakName() *string {
	return s.LeakName
}

func (s *GetResourceControlEventResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetResourceControlEventResponseBodyData) GetPunishReason() *string {
	return s.PunishReason
}

func (s *GetResourceControlEventResponseBodyData) GetSnapshotUrl() *string {
	return s.SnapshotUrl
}

func (s *GetResourceControlEventResponseBodyData) GetSrcIp() *string {
	return s.SrcIp
}

func (s *GetResourceControlEventResponseBodyData) GetSrcPort() *string {
	return s.SrcPort
}

func (s *GetResourceControlEventResponseBodyData) GetTip() *string {
	return s.Tip
}

func (s *GetResourceControlEventResponseBodyData) SetApplyRecordList(v []*GetResourceControlEventResponseBodyDataApplyRecordList) *GetResourceControlEventResponseBodyData {
	s.ApplyRecordList = v
	return s
}

func (s *GetResourceControlEventResponseBodyData) SetAssistantTip(v string) *GetResourceControlEventResponseBodyData {
	s.AssistantTip = &v
	return s
}

func (s *GetResourceControlEventResponseBodyData) SetBlockIp(v string) *GetResourceControlEventResponseBodyData {
	s.BlockIp = &v
	return s
}

func (s *GetResourceControlEventResponseBodyData) SetDirection(v string) *GetResourceControlEventResponseBodyData {
	s.Direction = &v
	return s
}

func (s *GetResourceControlEventResponseBodyData) SetDstIp(v string) *GetResourceControlEventResponseBodyData {
	s.DstIp = &v
	return s
}

func (s *GetResourceControlEventResponseBodyData) SetDstPort(v string) *GetResourceControlEventResponseBodyData {
	s.DstPort = &v
	return s
}

func (s *GetResourceControlEventResponseBodyData) SetEventId(v string) *GetResourceControlEventResponseBodyData {
	s.EventId = &v
	return s
}

func (s *GetResourceControlEventResponseBodyData) SetEventImpact(v string) *GetResourceControlEventResponseBodyData {
	s.EventImpact = &v
	return s
}

func (s *GetResourceControlEventResponseBodyData) SetLeakName(v string) *GetResourceControlEventResponseBodyData {
	s.LeakName = &v
	return s
}

func (s *GetResourceControlEventResponseBodyData) SetProtocol(v string) *GetResourceControlEventResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetResourceControlEventResponseBodyData) SetPunishReason(v string) *GetResourceControlEventResponseBodyData {
	s.PunishReason = &v
	return s
}

func (s *GetResourceControlEventResponseBodyData) SetSnapshotUrl(v string) *GetResourceControlEventResponseBodyData {
	s.SnapshotUrl = &v
	return s
}

func (s *GetResourceControlEventResponseBodyData) SetSrcIp(v string) *GetResourceControlEventResponseBodyData {
	s.SrcIp = &v
	return s
}

func (s *GetResourceControlEventResponseBodyData) SetSrcPort(v string) *GetResourceControlEventResponseBodyData {
	s.SrcPort = &v
	return s
}

func (s *GetResourceControlEventResponseBodyData) SetTip(v string) *GetResourceControlEventResponseBodyData {
	s.Tip = &v
	return s
}

func (s *GetResourceControlEventResponseBodyData) Validate() error {
	if s.ApplyRecordList != nil {
		for _, item := range s.ApplyRecordList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourceControlEventResponseBodyDataApplyRecordList struct {
	// example:
	//
	// meet the requirements
	ApprovalReason  *string                                                                `json:"ApprovalReason,omitempty" xml:"ApprovalReason,omitempty"`
	EventTimeRecord *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord `json:"EventTimeRecord,omitempty" xml:"EventTimeRecord,omitempty" type:"Struct"`
	// example:
	//
	// does not meet the requirements
	RejectReason *string `json:"RejectReason,omitempty" xml:"RejectReason,omitempty"`
	// example:
	//
	// Test Desc for Draft
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetResourceControlEventResponseBodyDataApplyRecordList) String() string {
	return dara.Prettify(s)
}

func (s GetResourceControlEventResponseBodyDataApplyRecordList) GoString() string {
	return s.String()
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordList) GetApprovalReason() *string {
	return s.ApprovalReason
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordList) GetEventTimeRecord() *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	return s.EventTimeRecord
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordList) GetRejectReason() *string {
	return s.RejectReason
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordList) GetRemark() *string {
	return s.Remark
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordList) GetStatus() *string {
	return s.Status
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordList) SetApprovalReason(v string) *GetResourceControlEventResponseBodyDataApplyRecordList {
	s.ApprovalReason = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordList) SetEventTimeRecord(v *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) *GetResourceControlEventResponseBodyDataApplyRecordList {
	s.EventTimeRecord = v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordList) SetRejectReason(v string) *GetResourceControlEventResponseBodyDataApplyRecordList {
	s.RejectReason = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordList) SetRemark(v string) *GetResourceControlEventResponseBodyDataApplyRecordList {
	s.Remark = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordList) SetStatus(v string) *GetResourceControlEventResponseBodyDataApplyRecordList {
	s.Status = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordList) Validate() error {
	if s.EventTimeRecord != nil {
		if err := s.EventTimeRecord.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord struct {
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
	// 2025-08-21T02:26:50Z
	ApplyTime *string `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	IgnoreAlertTime *string `json:"IgnoreAlertTime,omitempty" xml:"IgnoreAlertTime,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	InstanceCloseTime *string `json:"InstanceCloseTime,omitempty" xml:"InstanceCloseTime,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	InstanceScanTime *string `json:"InstanceScanTime,omitempty" xml:"InstanceScanTime,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	LastCheckTime *string `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	MiningAlertProcessTime *string `json:"MiningAlertProcessTime,omitempty" xml:"MiningAlertProcessTime,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	PreCloseTime *string `json:"PreCloseTime,omitempty" xml:"PreCloseTime,omitempty"`
	// example:
	//
	// 2025-11-05 10:06:21
	ProcessTime *string `json:"ProcessTime,omitempty" xml:"ProcessTime,omitempty"`
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
	// 2026-03-16 15:15:00
	RejectTime *string `json:"RejectTime,omitempty" xml:"RejectTime,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	RemoveTime *string `json:"RemoveTime,omitempty" xml:"RemoveTime,omitempty"`
	// example:
	//
	// 2026-03-16 15:15:00
	RiskCheckSuccessTime *string `json:"RiskCheckSuccessTime,omitempty" xml:"RiskCheckSuccessTime,omitempty"`
}

func (s GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) String() string {
	return dara.Prettify(s)
}

func (s GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GoString() string {
	return s.String()
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetAlertEndTime() *string {
	return s.AlertEndTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetAlertStartTime() *string {
	return s.AlertStartTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetAntiPunishTime() *string {
	return s.AntiPunishTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetApplyTime() *string {
	return s.ApplyTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetIgnoreAlertTime() *string {
	return s.IgnoreAlertTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetInstanceCloseTime() *string {
	return s.InstanceCloseTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetInstanceScanTime() *string {
	return s.InstanceScanTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetLastCheckTime() *string {
	return s.LastCheckTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetMiningAlertProcessTime() *string {
	return s.MiningAlertProcessTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetPreCloseTime() *string {
	return s.PreCloseTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetProcessTime() *string {
	return s.ProcessTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetPunishEndTime() *string {
	return s.PunishEndTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetPunishStartTime() *string {
	return s.PunishStartTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetRejectTime() *string {
	return s.RejectTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetRemoveTime() *string {
	return s.RemoveTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) GetRiskCheckSuccessTime() *string {
	return s.RiskCheckSuccessTime
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetAlertEndTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.AlertEndTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetAlertStartTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.AlertStartTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetAntiPunishTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.AntiPunishTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetApplyTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.ApplyTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetIgnoreAlertTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.IgnoreAlertTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetInstanceCloseTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.InstanceCloseTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetInstanceScanTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.InstanceScanTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetLastCheckTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.LastCheckTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetMiningAlertProcessTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.MiningAlertProcessTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetPreCloseTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.PreCloseTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetProcessTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.ProcessTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetPunishEndTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.PunishEndTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetPunishStartTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.PunishStartTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetRejectTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.RejectTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetRemoveTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.RemoveTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) SetRiskCheckSuccessTime(v string) *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord {
	s.RiskCheckSuccessTime = &v
	return s
}

func (s *GetResourceControlEventResponseBodyDataApplyRecordListEventTimeRecord) Validate() error {
	return dara.Validate(s)
}
