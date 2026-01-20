// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadOutboundTaskCallListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadOutboundTaskCallListResponseBody
	GetCode() *string
	SetCurrent(v int32) *ReadOutboundTaskCallListResponseBody
	GetCurrent() *int32
	SetMessage(v string) *ReadOutboundTaskCallListResponseBody
	GetMessage() *string
	SetRecords(v []*ReadOutboundTaskCallListResponseBodyRecords) *ReadOutboundTaskCallListResponseBody
	GetRecords() []*ReadOutboundTaskCallListResponseBodyRecords
	SetRequestId(v string) *ReadOutboundTaskCallListResponseBody
	GetRequestId() *string
	SetSize(v int32) *ReadOutboundTaskCallListResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ReadOutboundTaskCallListResponseBody
	GetSuccess() *bool
	SetTimestamp(v string) *ReadOutboundTaskCallListResponseBody
	GetTimestamp() *string
	SetTotal(v int64) *ReadOutboundTaskCallListResponseBody
	GetTotal() *int64
	SetTraceId(v string) *ReadOutboundTaskCallListResponseBody
	GetTraceId() *string
}

type ReadOutboundTaskCallListResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// successful
	Message *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Records []*ReadOutboundTaskCallListResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 202BFA44-28D8-571E-B992-BA70F2E92FB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1743387963
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// F47D4976-FC5A-5687-A890-B7923D3B429B
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ReadOutboundTaskCallListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadOutboundTaskCallListResponseBody) GoString() string {
	return s.String()
}

func (s *ReadOutboundTaskCallListResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadOutboundTaskCallListResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ReadOutboundTaskCallListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadOutboundTaskCallListResponseBody) GetRecords() []*ReadOutboundTaskCallListResponseBodyRecords {
	return s.Records
}

func (s *ReadOutboundTaskCallListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadOutboundTaskCallListResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ReadOutboundTaskCallListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadOutboundTaskCallListResponseBody) GetTimestamp() *string {
	return s.Timestamp
}

func (s *ReadOutboundTaskCallListResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ReadOutboundTaskCallListResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ReadOutboundTaskCallListResponseBody) SetCode(v string) *ReadOutboundTaskCallListResponseBody {
	s.Code = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBody) SetCurrent(v int32) *ReadOutboundTaskCallListResponseBody {
	s.Current = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBody) SetMessage(v string) *ReadOutboundTaskCallListResponseBody {
	s.Message = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBody) SetRecords(v []*ReadOutboundTaskCallListResponseBodyRecords) *ReadOutboundTaskCallListResponseBody {
	s.Records = v
	return s
}

func (s *ReadOutboundTaskCallListResponseBody) SetRequestId(v string) *ReadOutboundTaskCallListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBody) SetSize(v int32) *ReadOutboundTaskCallListResponseBody {
	s.Size = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBody) SetSuccess(v bool) *ReadOutboundTaskCallListResponseBody {
	s.Success = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBody) SetTimestamp(v string) *ReadOutboundTaskCallListResponseBody {
	s.Timestamp = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBody) SetTotal(v int64) *ReadOutboundTaskCallListResponseBody {
	s.Total = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBody) SetTraceId(v string) *ReadOutboundTaskCallListResponseBody {
	s.TraceId = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReadOutboundTaskCallListResponseBodyRecords struct {
	// example:
	//
	// 2025-09-23 19:38:44
	CallEndTime *string `json:"CallEndTime,omitempty" xml:"CallEndTime,omitempty"`
	// example:
	//
	// call_123456
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// example:
	//
	// 2025-09-23 19:34:12
	CallStartTime *string `json:"CallStartTime,omitempty" xml:"CallStartTime,omitempty"`
	// example:
	//
	// 客户对双十一活动很感兴趣
	CallSummary *string `json:"CallSummary,omitempty" xml:"CallSummary,omitempty"`
	// example:
	//
	// normal
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// 张先生
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	// example:
	//
	// 131****9908
	CustomerPhone *string `json:"CustomerPhone,omitempty" xml:"CustomerPhone,omitempty"`
	// example:
	//
	// []
	DialogueList []*ReadOutboundTaskCallListResponseBodyRecordsDialogueList `json:"DialogueList,omitempty" xml:"DialogueList,omitempty" type:"Repeated"`
	// example:
	//
	// 已接通
	DisplayStatus *string `json:"DisplayStatus,omitempty" xml:"DisplayStatus,omitempty"`
	// example:
	//
	// 272
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 4m32s
	DurationText *string `json:"DurationText,omitempty" xml:"DurationText,omitempty"`
	// example:
	//
	// {"123":"122"}
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	// example:
	//
	// 2025-09-23 19:34:12
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-09-23 19:34:12
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// instance_001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ["有意向", "高净值"]
	LabelTags []*string `json:"LabelTags,omitempty" xml:"LabelTags,omitempty" type:"Repeated"`
	// example:
	//
	// true
	RecordDetailReady *bool `json:"RecordDetailReady,omitempty" xml:"RecordDetailReady,omitempty"`
	// example:
	//
	// https://xxx.com/record.mp3
	RecordUrl *string `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty"`
	// example:
	//
	// 0
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 3
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// tenant_001
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// xiaomei
	TtsVoiceCode *string `json:"TtsVoiceCode,omitempty" xml:"TtsVoiceCode,omitempty"`
	// example:
	//
	// 小美(声音甜美)
	TtsVoiceDesc *string `json:"TtsVoiceDesc,omitempty" xml:"TtsVoiceDesc,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ReadOutboundTaskCallListResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s ReadOutboundTaskCallListResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetCallEndTime() *string {
	return s.CallEndTime
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetCallId() *string {
	return s.CallId
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetCallStartTime() *string {
	return s.CallStartTime
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetCallSummary() *string {
	return s.CallSummary
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetChannel() *string {
	return s.Channel
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetCustomerName() *string {
	return s.CustomerName
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetCustomerPhone() *string {
	return s.CustomerPhone
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetDialogueList() []*ReadOutboundTaskCallListResponseBodyRecordsDialogueList {
	return s.DialogueList
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetDisplayStatus() *string {
	return s.DisplayStatus
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetDuration() *int64 {
	return s.Duration
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetDurationText() *string {
	return s.DurationText
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetId() *int64 {
	return s.Id
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetLabelTags() []*string {
	return s.LabelTags
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetRecordDetailReady() *bool {
	return s.RecordDetailReady
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetRecordUrl() *string {
	return s.RecordUrl
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetSceneId() *string {
	return s.SceneId
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetStatus() *string {
	return s.Status
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetTaskId() *string {
	return s.TaskId
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetTenantId() *string {
	return s.TenantId
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetTtsVoiceCode() *string {
	return s.TtsVoiceCode
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetTtsVoiceDesc() *string {
	return s.TtsVoiceDesc
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) GetUserId() *string {
	return s.UserId
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetCallEndTime(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.CallEndTime = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetCallId(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.CallId = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetCallStartTime(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.CallStartTime = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetCallSummary(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.CallSummary = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetChannel(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.Channel = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetCustomerName(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.CustomerName = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetCustomerPhone(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.CustomerPhone = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetDialogueList(v []*ReadOutboundTaskCallListResponseBodyRecordsDialogueList) *ReadOutboundTaskCallListResponseBodyRecords {
	s.DialogueList = v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetDisplayStatus(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.DisplayStatus = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetDuration(v int64) *ReadOutboundTaskCallListResponseBodyRecords {
	s.Duration = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetDurationText(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.DurationText = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetExtInfo(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.ExtInfo = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetGmtCreate(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.GmtCreate = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetGmtModified(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.GmtModified = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetId(v int64) *ReadOutboundTaskCallListResponseBodyRecords {
	s.Id = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetInstanceId(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.InstanceId = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetLabelTags(v []*string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.LabelTags = v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetRecordDetailReady(v bool) *ReadOutboundTaskCallListResponseBodyRecords {
	s.RecordDetailReady = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetRecordUrl(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.RecordUrl = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetRetryCount(v int32) *ReadOutboundTaskCallListResponseBodyRecords {
	s.RetryCount = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetSceneId(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.SceneId = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetStatus(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.Status = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetTaskId(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.TaskId = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetTenantId(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.TenantId = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetTtsVoiceCode(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.TtsVoiceCode = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetTtsVoiceDesc(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.TtsVoiceDesc = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) SetUserId(v string) *ReadOutboundTaskCallListResponseBodyRecords {
	s.UserId = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecords) Validate() error {
	if s.DialogueList != nil {
		for _, item := range s.DialogueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReadOutboundTaskCallListResponseBodyRecordsDialogueList struct {
	// example:
	//
	// 1000
	BeginTime *int32 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 5000
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// agent
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 您好，请问有什么可以帮助您？
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ReadOutboundTaskCallListResponseBodyRecordsDialogueList) String() string {
	return dara.Prettify(s)
}

func (s ReadOutboundTaskCallListResponseBodyRecordsDialogueList) GoString() string {
	return s.String()
}

func (s *ReadOutboundTaskCallListResponseBodyRecordsDialogueList) GetBeginTime() *int32 {
	return s.BeginTime
}

func (s *ReadOutboundTaskCallListResponseBodyRecordsDialogueList) GetEndTime() *int32 {
	return s.EndTime
}

func (s *ReadOutboundTaskCallListResponseBodyRecordsDialogueList) GetRole() *string {
	return s.Role
}

func (s *ReadOutboundTaskCallListResponseBodyRecordsDialogueList) GetText() *string {
	return s.Text
}

func (s *ReadOutboundTaskCallListResponseBodyRecordsDialogueList) SetBeginTime(v int32) *ReadOutboundTaskCallListResponseBodyRecordsDialogueList {
	s.BeginTime = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecordsDialogueList) SetEndTime(v int32) *ReadOutboundTaskCallListResponseBodyRecordsDialogueList {
	s.EndTime = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecordsDialogueList) SetRole(v string) *ReadOutboundTaskCallListResponseBodyRecordsDialogueList {
	s.Role = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecordsDialogueList) SetText(v string) *ReadOutboundTaskCallListResponseBodyRecordsDialogueList {
	s.Text = &v
	return s
}

func (s *ReadOutboundTaskCallListResponseBodyRecordsDialogueList) Validate() error {
	return dara.Validate(s)
}
