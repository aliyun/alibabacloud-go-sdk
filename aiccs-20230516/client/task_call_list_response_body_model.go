// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskCallListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *TaskCallListResponseBody
	GetCode() *int64
	SetMessage(v string) *TaskCallListResponseBody
	GetMessage() *string
	SetModel(v *TaskCallListResponseBodyModel) *TaskCallListResponseBody
	GetModel() *TaskCallListResponseBodyModel
	SetRequestId(v string) *TaskCallListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *TaskCallListResponseBody
	GetSuccess() *string
	SetTimestamp(v int64) *TaskCallListResponseBody
	GetTimestamp() *int64
}

type TaskCallListResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *TaskCallListResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s TaskCallListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TaskCallListResponseBody) GoString() string {
	return s.String()
}

func (s *TaskCallListResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *TaskCallListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TaskCallListResponseBody) GetModel() *TaskCallListResponseBodyModel {
	return s.Model
}

func (s *TaskCallListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TaskCallListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *TaskCallListResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *TaskCallListResponseBody) SetCode(v int64) *TaskCallListResponseBody {
	s.Code = &v
	return s
}

func (s *TaskCallListResponseBody) SetMessage(v string) *TaskCallListResponseBody {
	s.Message = &v
	return s
}

func (s *TaskCallListResponseBody) SetModel(v *TaskCallListResponseBodyModel) *TaskCallListResponseBody {
	s.Model = v
	return s
}

func (s *TaskCallListResponseBody) SetRequestId(v string) *TaskCallListResponseBody {
	s.RequestId = &v
	return s
}

func (s *TaskCallListResponseBody) SetSuccess(v string) *TaskCallListResponseBody {
	s.Success = &v
	return s
}

func (s *TaskCallListResponseBody) SetTimestamp(v int64) *TaskCallListResponseBody {
	s.Timestamp = &v
	return s
}

func (s *TaskCallListResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TaskCallListResponseBodyModel struct {
	List []*TaskCallListResponseBodyModelList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 62
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 95
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 80
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 39
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s TaskCallListResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s TaskCallListResponseBodyModel) GoString() string {
	return s.String()
}

func (s *TaskCallListResponseBodyModel) GetList() []*TaskCallListResponseBodyModelList {
	return s.List
}

func (s *TaskCallListResponseBodyModel) GetPageNo() *int64 {
	return s.PageNo
}

func (s *TaskCallListResponseBodyModel) GetPageSize() *int64 {
	return s.PageSize
}

func (s *TaskCallListResponseBodyModel) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *TaskCallListResponseBodyModel) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *TaskCallListResponseBodyModel) SetList(v []*TaskCallListResponseBodyModelList) *TaskCallListResponseBodyModel {
	s.List = v
	return s
}

func (s *TaskCallListResponseBodyModel) SetPageNo(v int64) *TaskCallListResponseBodyModel {
	s.PageNo = &v
	return s
}

func (s *TaskCallListResponseBodyModel) SetPageSize(v int64) *TaskCallListResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *TaskCallListResponseBodyModel) SetTotalCount(v int64) *TaskCallListResponseBodyModel {
	s.TotalCount = &v
	return s
}

func (s *TaskCallListResponseBodyModel) SetTotalPage(v int64) *TaskCallListResponseBodyModel {
	s.TotalPage = &v
	return s
}

func (s *TaskCallListResponseBodyModel) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TaskCallListResponseBodyModelList struct {
	// 加微
	//
	// example:
	//
	// 0
	AddWx *int64 `json:"AddWx,omitempty" xml:"AddWx,omitempty"`
	// 加微进度
	//
	// example:
	//
	// 示例值示例值
	AddWxStatus *string `json:"AddWxStatus,omitempty" xml:"AddWxStatus,omitempty"`
	// 坐席分机
	//
	// example:
	//
	// 112
	AgentExtension *string `json:"AgentExtension,omitempty" xml:"AgentExtension,omitempty"`
	// 坐席ID
	//
	// example:
	//
	// 87
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 人工通话时长
	//
	// example:
	//
	// 98
	AgentSpeakingDuration *int64 `json:"AgentSpeakingDuration,omitempty" xml:"AgentSpeakingDuration,omitempty"`
	// 人工通话时长
	//
	// example:
	//
	// 示例值示例值
	AgentSpeakingTime *string `json:"AgentSpeakingTime,omitempty" xml:"AgentSpeakingTime,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// A
	AgentTag *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	// 是否接通重呼
	//
	// example:
	//
	// 24
	AnswerRecall *int64 `json:"AnswerRecall,omitempty" xml:"AnswerRecall,omitempty"`
	// 接通时间
	//
	// example:
	//
	// 2022-01-26 18:58:25
	AnswerTime *string `json:"AnswerTime,omitempty" xml:"AnswerTime,omitempty"`
	// 批次ID
	//
	// example:
	//
	// 1
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 开始通话时间
	//
	// example:
	//
	// 2022-01-26 18:58:25
	CallBeginTime *string `json:"CallBeginTime,omitempty" xml:"CallBeginTime,omitempty"`
	// 外呼ID
	//
	// example:
	//
	// 9197ed9e-ceda-42a5-b683-823b23ef208e
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 呼叫次数
	//
	// example:
	//
	// 1
	CallTimes *string `json:"CallTimes,omitempty" xml:"CallTimes,omitempty"`
	// 外呼类型
	//
	// example:
	//
	// 1001
	CallType *int64 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 对话录音
	//
	// example:
	//
	// 示例值示例值
	ChatRecord *string `json:"ChatRecord,omitempty" xml:"ChatRecord,omitempty"`
	// 外呼网关
	//
	// example:
	//
	// 123
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// 挂断时间
	//
	// example:
	//
	// 2022-01-26 18:58:25
	HangupTime *string `json:"HangupTime,omitempty" xml:"HangupTime,omitempty"`
	// 挂机方式
	//
	// example:
	//
	// 1
	HangupType *int64 `json:"HangupType,omitempty" xml:"HangupType,omitempty"`
	// 导入时间
	//
	// example:
	//
	// 2022-01-26 18:58:25
	ImportTime *string `json:"ImportTime,omitempty" xml:"ImportTime,omitempty"`
	// 个性标签
	//
	// example:
	//
	// A
	IndividualTag *string `json:"IndividualTag,omitempty" xml:"IndividualTag,omitempty"`
	// 意向说明
	//
	// example:
	//
	// 示例值示例值示例值
	IntentDescription *string `json:"IntentDescription,omitempty" xml:"IntentDescription,omitempty"`
	// 意向标签
	//
	// example:
	//
	// “C”
	IntentTag *string `json:"IntentTag,omitempty" xml:"IntentTag,omitempty"`
	// 拦截原因
	//
	// example:
	//
	// 示例值
	InterceptReason *string `json:"InterceptReason,omitempty" xml:"InterceptReason,omitempty"`
	// 回复关键词
	//
	// example:
	//
	// 示例值示例值
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// 外呼号码
	//
	// example:
	//
	// 138*****123
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// 外呼号码MD5
	//
	// example:
	//
	// 75916b635568954583783d
	NumberMD5 *string `json:"NumberMD5,omitempty" xml:"NumberMD5,omitempty"`
	// 参数
	//
	// example:
	//
	// 示例值示例值
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// 备注
	//
	// example:
	//
	// 示例值示例值
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 振铃时长
	//
	// example:
	//
	// 66
	RingTime *int64 `json:"RingTime,omitempty" xml:"RingTime,omitempty"`
	// 挂机短信
	//
	// example:
	//
	// 示例值
	Sms *string `json:"Sms,omitempty" xml:"Sms,omitempty"`
	// AI通话时长
	//
	// example:
	//
	// 45
	SpeakingDuration *int64 `json:"SpeakingDuration,omitempty" xml:"SpeakingDuration,omitempty"`
	// AI通话时长
	//
	// example:
	//
	// 示例值示例值示例值
	SpeakingTime *string `json:"SpeakingTime,omitempty" xml:"SpeakingTime,omitempty"`
	// 对话轮次
	//
	// example:
	//
	// 0
	SpeakingTurns *string `json:"SpeakingTurns,omitempty" xml:"SpeakingTurns,omitempty"`
	// 外呼状态
	//
	// example:
	//
	// 示例值示例值示例值
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 外呼状态编码
	//
	// example:
	//
	// 1
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// 外呼状态描述
	//
	// example:
	//
	// 示例值示例值
	StatusDescription *string `json:"StatusDescription,omitempty" xml:"StatusDescription,omitempty"`
	// 用户自定义标签
	//
	// example:
	//
	// 示例值
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// 外呼任务ID
	//
	// example:
	//
	// 70
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// AI话术ID
	//
	// example:
	//
	// 66
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 转人工状态
	//
	// example:
	//
	// 示例值
	TransferStatus *string `json:"TransferStatus,omitempty" xml:"TransferStatus,omitempty"`
	// 转人工状态编码
	//
	// example:
	//
	// 1
	TransferStatusCode *int64 `json:"TransferStatusCode,omitempty" xml:"TransferStatusCode,omitempty"`
}

func (s TaskCallListResponseBodyModelList) String() string {
	return dara.Prettify(s)
}

func (s TaskCallListResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *TaskCallListResponseBodyModelList) GetAddWx() *int64 {
	return s.AddWx
}

func (s *TaskCallListResponseBodyModelList) GetAddWxStatus() *string {
	return s.AddWxStatus
}

func (s *TaskCallListResponseBodyModelList) GetAgentExtension() *string {
	return s.AgentExtension
}

func (s *TaskCallListResponseBodyModelList) GetAgentId() *int64 {
	return s.AgentId
}

func (s *TaskCallListResponseBodyModelList) GetAgentSpeakingDuration() *int64 {
	return s.AgentSpeakingDuration
}

func (s *TaskCallListResponseBodyModelList) GetAgentSpeakingTime() *string {
	return s.AgentSpeakingTime
}

func (s *TaskCallListResponseBodyModelList) GetAgentTag() *string {
	return s.AgentTag
}

func (s *TaskCallListResponseBodyModelList) GetAnswerRecall() *int64 {
	return s.AnswerRecall
}

func (s *TaskCallListResponseBodyModelList) GetAnswerTime() *string {
	return s.AnswerTime
}

func (s *TaskCallListResponseBodyModelList) GetBatchId() *string {
	return s.BatchId
}

func (s *TaskCallListResponseBodyModelList) GetCallBeginTime() *string {
	return s.CallBeginTime
}

func (s *TaskCallListResponseBodyModelList) GetCallId() *string {
	return s.CallId
}

func (s *TaskCallListResponseBodyModelList) GetCallTimes() *string {
	return s.CallTimes
}

func (s *TaskCallListResponseBodyModelList) GetCallType() *int64 {
	return s.CallType
}

func (s *TaskCallListResponseBodyModelList) GetChatRecord() *string {
	return s.ChatRecord
}

func (s *TaskCallListResponseBodyModelList) GetGateway() *string {
	return s.Gateway
}

func (s *TaskCallListResponseBodyModelList) GetHangupTime() *string {
	return s.HangupTime
}

func (s *TaskCallListResponseBodyModelList) GetHangupType() *int64 {
	return s.HangupType
}

func (s *TaskCallListResponseBodyModelList) GetImportTime() *string {
	return s.ImportTime
}

func (s *TaskCallListResponseBodyModelList) GetIndividualTag() *string {
	return s.IndividualTag
}

func (s *TaskCallListResponseBodyModelList) GetIntentDescription() *string {
	return s.IntentDescription
}

func (s *TaskCallListResponseBodyModelList) GetIntentTag() *string {
	return s.IntentTag
}

func (s *TaskCallListResponseBodyModelList) GetInterceptReason() *string {
	return s.InterceptReason
}

func (s *TaskCallListResponseBodyModelList) GetKeywords() *string {
	return s.Keywords
}

func (s *TaskCallListResponseBodyModelList) GetNumber() *string {
	return s.Number
}

func (s *TaskCallListResponseBodyModelList) GetNumberMD5() *string {
	return s.NumberMD5
}

func (s *TaskCallListResponseBodyModelList) GetProperties() *string {
	return s.Properties
}

func (s *TaskCallListResponseBodyModelList) GetRemark() *string {
	return s.Remark
}

func (s *TaskCallListResponseBodyModelList) GetRingTime() *int64 {
	return s.RingTime
}

func (s *TaskCallListResponseBodyModelList) GetSms() *string {
	return s.Sms
}

func (s *TaskCallListResponseBodyModelList) GetSpeakingDuration() *int64 {
	return s.SpeakingDuration
}

func (s *TaskCallListResponseBodyModelList) GetSpeakingTime() *string {
	return s.SpeakingTime
}

func (s *TaskCallListResponseBodyModelList) GetSpeakingTurns() *string {
	return s.SpeakingTurns
}

func (s *TaskCallListResponseBodyModelList) GetStatus() *string {
	return s.Status
}

func (s *TaskCallListResponseBodyModelList) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *TaskCallListResponseBodyModelList) GetStatusDescription() *string {
	return s.StatusDescription
}

func (s *TaskCallListResponseBodyModelList) GetTag() *string {
	return s.Tag
}

func (s *TaskCallListResponseBodyModelList) GetTaskId() *int64 {
	return s.TaskId
}

func (s *TaskCallListResponseBodyModelList) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *TaskCallListResponseBodyModelList) GetTransferStatus() *string {
	return s.TransferStatus
}

func (s *TaskCallListResponseBodyModelList) GetTransferStatusCode() *int64 {
	return s.TransferStatusCode
}

func (s *TaskCallListResponseBodyModelList) SetAddWx(v int64) *TaskCallListResponseBodyModelList {
	s.AddWx = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAddWxStatus(v string) *TaskCallListResponseBodyModelList {
	s.AddWxStatus = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAgentExtension(v string) *TaskCallListResponseBodyModelList {
	s.AgentExtension = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAgentId(v int64) *TaskCallListResponseBodyModelList {
	s.AgentId = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAgentSpeakingDuration(v int64) *TaskCallListResponseBodyModelList {
	s.AgentSpeakingDuration = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAgentSpeakingTime(v string) *TaskCallListResponseBodyModelList {
	s.AgentSpeakingTime = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAgentTag(v string) *TaskCallListResponseBodyModelList {
	s.AgentTag = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAnswerRecall(v int64) *TaskCallListResponseBodyModelList {
	s.AnswerRecall = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAnswerTime(v string) *TaskCallListResponseBodyModelList {
	s.AnswerTime = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetBatchId(v string) *TaskCallListResponseBodyModelList {
	s.BatchId = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetCallBeginTime(v string) *TaskCallListResponseBodyModelList {
	s.CallBeginTime = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetCallId(v string) *TaskCallListResponseBodyModelList {
	s.CallId = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetCallTimes(v string) *TaskCallListResponseBodyModelList {
	s.CallTimes = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetCallType(v int64) *TaskCallListResponseBodyModelList {
	s.CallType = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetChatRecord(v string) *TaskCallListResponseBodyModelList {
	s.ChatRecord = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetGateway(v string) *TaskCallListResponseBodyModelList {
	s.Gateway = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetHangupTime(v string) *TaskCallListResponseBodyModelList {
	s.HangupTime = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetHangupType(v int64) *TaskCallListResponseBodyModelList {
	s.HangupType = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetImportTime(v string) *TaskCallListResponseBodyModelList {
	s.ImportTime = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetIndividualTag(v string) *TaskCallListResponseBodyModelList {
	s.IndividualTag = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetIntentDescription(v string) *TaskCallListResponseBodyModelList {
	s.IntentDescription = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetIntentTag(v string) *TaskCallListResponseBodyModelList {
	s.IntentTag = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetInterceptReason(v string) *TaskCallListResponseBodyModelList {
	s.InterceptReason = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetKeywords(v string) *TaskCallListResponseBodyModelList {
	s.Keywords = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetNumber(v string) *TaskCallListResponseBodyModelList {
	s.Number = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetNumberMD5(v string) *TaskCallListResponseBodyModelList {
	s.NumberMD5 = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetProperties(v string) *TaskCallListResponseBodyModelList {
	s.Properties = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetRemark(v string) *TaskCallListResponseBodyModelList {
	s.Remark = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetRingTime(v int64) *TaskCallListResponseBodyModelList {
	s.RingTime = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetSms(v string) *TaskCallListResponseBodyModelList {
	s.Sms = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetSpeakingDuration(v int64) *TaskCallListResponseBodyModelList {
	s.SpeakingDuration = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetSpeakingTime(v string) *TaskCallListResponseBodyModelList {
	s.SpeakingTime = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetSpeakingTurns(v string) *TaskCallListResponseBodyModelList {
	s.SpeakingTurns = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetStatus(v string) *TaskCallListResponseBodyModelList {
	s.Status = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetStatusCode(v int64) *TaskCallListResponseBodyModelList {
	s.StatusCode = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetStatusDescription(v string) *TaskCallListResponseBodyModelList {
	s.StatusDescription = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetTag(v string) *TaskCallListResponseBodyModelList {
	s.Tag = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetTaskId(v int64) *TaskCallListResponseBodyModelList {
	s.TaskId = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetTemplateId(v int64) *TaskCallListResponseBodyModelList {
	s.TemplateId = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetTransferStatus(v string) *TaskCallListResponseBodyModelList {
	s.TransferStatus = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetTransferStatusCode(v int64) *TaskCallListResponseBodyModelList {
	s.TransferStatusCode = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) Validate() error {
	return dara.Validate(s)
}
