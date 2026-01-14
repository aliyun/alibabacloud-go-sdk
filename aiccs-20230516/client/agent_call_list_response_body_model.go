// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentCallListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AgentCallListResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int64) *AgentCallListResponseBody
	GetCode() *int64
	SetMessage(v string) *AgentCallListResponseBody
	GetMessage() *string
	SetModel(v *AgentCallListResponseBodyModel) *AgentCallListResponseBody
	GetModel() *AgentCallListResponseBodyModel
	SetRequestId(v string) *AgentCallListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AgentCallListResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *AgentCallListResponseBody
	GetTimestamp() *int64
}

type AgentCallListResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 60
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *AgentCallListResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 89
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s AgentCallListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AgentCallListResponseBody) GoString() string {
	return s.String()
}

func (s *AgentCallListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AgentCallListResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *AgentCallListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AgentCallListResponseBody) GetModel() *AgentCallListResponseBodyModel {
	return s.Model
}

func (s *AgentCallListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AgentCallListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AgentCallListResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *AgentCallListResponseBody) SetAccessDeniedDetail(v string) *AgentCallListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AgentCallListResponseBody) SetCode(v int64) *AgentCallListResponseBody {
	s.Code = &v
	return s
}

func (s *AgentCallListResponseBody) SetMessage(v string) *AgentCallListResponseBody {
	s.Message = &v
	return s
}

func (s *AgentCallListResponseBody) SetModel(v *AgentCallListResponseBodyModel) *AgentCallListResponseBody {
	s.Model = v
	return s
}

func (s *AgentCallListResponseBody) SetRequestId(v string) *AgentCallListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgentCallListResponseBody) SetSuccess(v bool) *AgentCallListResponseBody {
	s.Success = &v
	return s
}

func (s *AgentCallListResponseBody) SetTimestamp(v int64) *AgentCallListResponseBody {
	s.Timestamp = &v
	return s
}

func (s *AgentCallListResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AgentCallListResponseBodyModel struct {
	List []*AgentCallListResponseBodyModelList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 44
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 40
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 27
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 49
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s AgentCallListResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s AgentCallListResponseBodyModel) GoString() string {
	return s.String()
}

func (s *AgentCallListResponseBodyModel) GetList() []*AgentCallListResponseBodyModelList {
	return s.List
}

func (s *AgentCallListResponseBodyModel) GetPageNo() *int64 {
	return s.PageNo
}

func (s *AgentCallListResponseBodyModel) GetPageSize() *int64 {
	return s.PageSize
}

func (s *AgentCallListResponseBodyModel) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *AgentCallListResponseBodyModel) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *AgentCallListResponseBodyModel) SetList(v []*AgentCallListResponseBodyModelList) *AgentCallListResponseBodyModel {
	s.List = v
	return s
}

func (s *AgentCallListResponseBodyModel) SetPageNo(v int64) *AgentCallListResponseBodyModel {
	s.PageNo = &v
	return s
}

func (s *AgentCallListResponseBodyModel) SetPageSize(v int64) *AgentCallListResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *AgentCallListResponseBodyModel) SetTotalCount(v int64) *AgentCallListResponseBodyModel {
	s.TotalCount = &v
	return s
}

func (s *AgentCallListResponseBodyModel) SetTotalPage(v int64) *AgentCallListResponseBodyModel {
	s.TotalPage = &v
	return s
}

func (s *AgentCallListResponseBodyModel) Validate() error {
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

type AgentCallListResponseBodyModelList struct {
	// 坐席分机
	//
	// example:
	//
	// a
	AgentExtension *string `json:"AgentExtension,omitempty" xml:"AgentExtension,omitempty"`
	// 分配坐席ID
	//
	// example:
	//
	// 1
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 坐席通话时长，单位：秒
	//
	// example:
	//
	// 1
	AgentSpeakingDuration *int64 `json:"AgentSpeakingDuration,omitempty" xml:"AgentSpeakingDuration,omitempty"`
	// 坐席通话时长，单位：大于1分钟，显示分钟秒，小于1分钟，显示秒
	//
	// example:
	//
	// 1
	AgentSpeakingTime *string `json:"AgentSpeakingTime,omitempty" xml:"AgentSpeakingTime,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// a
	AgentTag *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	// 接通时间 格式：2019-11-23 14:47:06
	//
	// example:
	//
	// 2019-11-23 14:47:06
	AnswerTime *string `json:"AnswerTime,omitempty" xml:"AnswerTime,omitempty"`
	// 导入号码时返回的批次号
	//
	// example:
	//
	// 1
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 开始通话时间, 格式：2019-11-23 14:47:06
	//
	// example:
	//
	// 2019-11-23 14:47:06
	CallBeginTime *string `json:"CallBeginTime,omitempty" xml:"CallBeginTime,omitempty"`
	// 外呼ID
	//
	// example:
	//
	// 1
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 可选项：1001：坐席-人工外呼，1002：坐席-AI外呼-不转人工，1003：坐席-AI外呼-接通转人工，1004：坐席-AI外呼-智能转人工；
	//
	// example:
	//
	// 1001
	CallType *int64 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 对话录音
	//
	// example:
	//
	// URL
	ChatRecord *string `json:"ChatRecord,omitempty" xml:"ChatRecord,omitempty"`
	// 挂断时间 格式：2019-11-23 14:47:06
	//
	// example:
	//
	// 2019-11-23 14:47:06
	HangupTime *string `json:"HangupTime,omitempty" xml:"HangupTime,omitempty"`
	// 挂机方式  AI挂机1，坐席挂机2，客户挂机3
	//
	// example:
	//
	// 1
	HangupType *int64 `json:"HangupType,omitempty" xml:"HangupType,omitempty"`
	// 导入时间,格式：2019-11-23 14:47:06
	//
	// example:
	//
	// 2019-11-23 14:47:06
	ImportTime *string `json:"ImportTime,omitempty" xml:"ImportTime,omitempty"`
	// 个性标签
	//
	// example:
	//
	// 如投诉,非本人
	IndividualTag *string `json:"IndividualTag,omitempty" xml:"IndividualTag,omitempty"`
	// 意向说明 如：确认本人,未承诺还款
	//
	// example:
	//
	// 如：确认本人
	IntentDescription *string `json:"IntentDescription,omitempty" xml:"IntentDescription,omitempty"`
	// 意向标签, 如“C”
	//
	// example:
	//
	// C
	IntentTag *string `json:"IntentTag,omitempty" xml:"IntentTag,omitempty"`
	// 拦截原因 当状态为已拦截时，可选值：黑名单拦截，灰名单拦截，异常号码拦截
	//
	// example:
	//
	// 黑名单拦截
	InterceptReason *string `json:"InterceptReason,omitempty" xml:"InterceptReason,omitempty"`
	// 回复关键词
	//
	// example:
	//
	// 如链接,利息
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// 外呼号码
	//
	// example:
	//
	// 1
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// 外呼号码MD5
	//
	// example:
	//
	// 1
	NumberMD5 *string `json:"NumberMD5,omitempty" xml:"NumberMD5,omitempty"`
	// 导入号码时的参数值
	//
	// example:
	//
	// {"电话号码":"13100000000"}
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// 人工备注信息
	//
	// example:
	//
	// abcd
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 振铃时长 单位：毫秒
	//
	// example:
	//
	// 1
	RingTime *int64 `json:"RingTime,omitempty" xml:"RingTime,omitempty"`
	// 挂机短信 1:发送，2:不发送
	//
	// example:
	//
	// 1
	Sms *string `json:"Sms,omitempty" xml:"Sms,omitempty"`
	// 通话时长 单位：秒
	//
	// example:
	//
	// 1
	SpeakingDuration *int64 `json:"SpeakingDuration,omitempty" xml:"SpeakingDuration,omitempty"`
	// 通话时长 单位：大于1分钟，显示分钟秒，小于1分钟，显示秒
	//
	// example:
	//
	// 1
	SpeakingTime *string `json:"SpeakingTime,omitempty" xml:"SpeakingTime,omitempty"`
	// 对话轮次
	//
	// example:
	//
	// 1
	SpeakingTurns *string `json:"SpeakingTurns,omitempty" xml:"SpeakingTurns,omitempty"`
	// 外呼状态编码
	//
	// example:
	//
	// 1
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// 外呼状态，如“已接听”“拒接”，转外呼状态编码与描述对应 1-已接听,2-关机
	//
	// example:
	//
	// 已接听
	StatusDescription *string `json:"StatusDescription,omitempty" xml:"StatusDescription,omitempty"`
	// 用户自定义标签
	//
	// example:
	//
	// 1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// 外呼任务ID
	//
	// example:
	//
	// 1
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// AI话术ID
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 转人工状态 0-未触发
	//
	// example:
	//
	// 未触发
	TransferStatus *string `json:"TransferStatus,omitempty" xml:"TransferStatus,omitempty"`
	// 转人工状态编码
	//
	// example:
	//
	// 0
	TransferStatusCode *int64 `json:"TransferStatusCode,omitempty" xml:"TransferStatusCode,omitempty"`
	// 转接人工时间
	//
	// example:
	//
	// 2019-01-09 14:14:19
	TransferTime *string `json:"TransferTime,omitempty" xml:"TransferTime,omitempty"`
}

func (s AgentCallListResponseBodyModelList) String() string {
	return dara.Prettify(s)
}

func (s AgentCallListResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *AgentCallListResponseBodyModelList) GetAgentExtension() *string {
	return s.AgentExtension
}

func (s *AgentCallListResponseBodyModelList) GetAgentId() *int64 {
	return s.AgentId
}

func (s *AgentCallListResponseBodyModelList) GetAgentSpeakingDuration() *int64 {
	return s.AgentSpeakingDuration
}

func (s *AgentCallListResponseBodyModelList) GetAgentSpeakingTime() *string {
	return s.AgentSpeakingTime
}

func (s *AgentCallListResponseBodyModelList) GetAgentTag() *string {
	return s.AgentTag
}

func (s *AgentCallListResponseBodyModelList) GetAnswerTime() *string {
	return s.AnswerTime
}

func (s *AgentCallListResponseBodyModelList) GetBatchId() *string {
	return s.BatchId
}

func (s *AgentCallListResponseBodyModelList) GetCallBeginTime() *string {
	return s.CallBeginTime
}

func (s *AgentCallListResponseBodyModelList) GetCallId() *string {
	return s.CallId
}

func (s *AgentCallListResponseBodyModelList) GetCallType() *int64 {
	return s.CallType
}

func (s *AgentCallListResponseBodyModelList) GetChatRecord() *string {
	return s.ChatRecord
}

func (s *AgentCallListResponseBodyModelList) GetHangupTime() *string {
	return s.HangupTime
}

func (s *AgentCallListResponseBodyModelList) GetHangupType() *int64 {
	return s.HangupType
}

func (s *AgentCallListResponseBodyModelList) GetImportTime() *string {
	return s.ImportTime
}

func (s *AgentCallListResponseBodyModelList) GetIndividualTag() *string {
	return s.IndividualTag
}

func (s *AgentCallListResponseBodyModelList) GetIntentDescription() *string {
	return s.IntentDescription
}

func (s *AgentCallListResponseBodyModelList) GetIntentTag() *string {
	return s.IntentTag
}

func (s *AgentCallListResponseBodyModelList) GetInterceptReason() *string {
	return s.InterceptReason
}

func (s *AgentCallListResponseBodyModelList) GetKeywords() *string {
	return s.Keywords
}

func (s *AgentCallListResponseBodyModelList) GetNumber() *string {
	return s.Number
}

func (s *AgentCallListResponseBodyModelList) GetNumberMD5() *string {
	return s.NumberMD5
}

func (s *AgentCallListResponseBodyModelList) GetProperties() *string {
	return s.Properties
}

func (s *AgentCallListResponseBodyModelList) GetRemark() *string {
	return s.Remark
}

func (s *AgentCallListResponseBodyModelList) GetRingTime() *int64 {
	return s.RingTime
}

func (s *AgentCallListResponseBodyModelList) GetSms() *string {
	return s.Sms
}

func (s *AgentCallListResponseBodyModelList) GetSpeakingDuration() *int64 {
	return s.SpeakingDuration
}

func (s *AgentCallListResponseBodyModelList) GetSpeakingTime() *string {
	return s.SpeakingTime
}

func (s *AgentCallListResponseBodyModelList) GetSpeakingTurns() *string {
	return s.SpeakingTurns
}

func (s *AgentCallListResponseBodyModelList) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *AgentCallListResponseBodyModelList) GetStatusDescription() *string {
	return s.StatusDescription
}

func (s *AgentCallListResponseBodyModelList) GetTag() *string {
	return s.Tag
}

func (s *AgentCallListResponseBodyModelList) GetTaskId() *int64 {
	return s.TaskId
}

func (s *AgentCallListResponseBodyModelList) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *AgentCallListResponseBodyModelList) GetTransferStatus() *string {
	return s.TransferStatus
}

func (s *AgentCallListResponseBodyModelList) GetTransferStatusCode() *int64 {
	return s.TransferStatusCode
}

func (s *AgentCallListResponseBodyModelList) GetTransferTime() *string {
	return s.TransferTime
}

func (s *AgentCallListResponseBodyModelList) SetAgentExtension(v string) *AgentCallListResponseBodyModelList {
	s.AgentExtension = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetAgentId(v int64) *AgentCallListResponseBodyModelList {
	s.AgentId = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetAgentSpeakingDuration(v int64) *AgentCallListResponseBodyModelList {
	s.AgentSpeakingDuration = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetAgentSpeakingTime(v string) *AgentCallListResponseBodyModelList {
	s.AgentSpeakingTime = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetAgentTag(v string) *AgentCallListResponseBodyModelList {
	s.AgentTag = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetAnswerTime(v string) *AgentCallListResponseBodyModelList {
	s.AnswerTime = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetBatchId(v string) *AgentCallListResponseBodyModelList {
	s.BatchId = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetCallBeginTime(v string) *AgentCallListResponseBodyModelList {
	s.CallBeginTime = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetCallId(v string) *AgentCallListResponseBodyModelList {
	s.CallId = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetCallType(v int64) *AgentCallListResponseBodyModelList {
	s.CallType = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetChatRecord(v string) *AgentCallListResponseBodyModelList {
	s.ChatRecord = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetHangupTime(v string) *AgentCallListResponseBodyModelList {
	s.HangupTime = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetHangupType(v int64) *AgentCallListResponseBodyModelList {
	s.HangupType = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetImportTime(v string) *AgentCallListResponseBodyModelList {
	s.ImportTime = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetIndividualTag(v string) *AgentCallListResponseBodyModelList {
	s.IndividualTag = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetIntentDescription(v string) *AgentCallListResponseBodyModelList {
	s.IntentDescription = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetIntentTag(v string) *AgentCallListResponseBodyModelList {
	s.IntentTag = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetInterceptReason(v string) *AgentCallListResponseBodyModelList {
	s.InterceptReason = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetKeywords(v string) *AgentCallListResponseBodyModelList {
	s.Keywords = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetNumber(v string) *AgentCallListResponseBodyModelList {
	s.Number = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetNumberMD5(v string) *AgentCallListResponseBodyModelList {
	s.NumberMD5 = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetProperties(v string) *AgentCallListResponseBodyModelList {
	s.Properties = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetRemark(v string) *AgentCallListResponseBodyModelList {
	s.Remark = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetRingTime(v int64) *AgentCallListResponseBodyModelList {
	s.RingTime = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetSms(v string) *AgentCallListResponseBodyModelList {
	s.Sms = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetSpeakingDuration(v int64) *AgentCallListResponseBodyModelList {
	s.SpeakingDuration = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetSpeakingTime(v string) *AgentCallListResponseBodyModelList {
	s.SpeakingTime = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetSpeakingTurns(v string) *AgentCallListResponseBodyModelList {
	s.SpeakingTurns = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetStatusCode(v int64) *AgentCallListResponseBodyModelList {
	s.StatusCode = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetStatusDescription(v string) *AgentCallListResponseBodyModelList {
	s.StatusDescription = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetTag(v string) *AgentCallListResponseBodyModelList {
	s.Tag = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetTaskId(v int64) *AgentCallListResponseBodyModelList {
	s.TaskId = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetTemplateId(v int64) *AgentCallListResponseBodyModelList {
	s.TemplateId = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetTransferStatus(v string) *AgentCallListResponseBodyModelList {
	s.TransferStatus = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetTransferStatusCode(v int64) *AgentCallListResponseBodyModelList {
	s.TransferStatusCode = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) SetTransferTime(v string) *AgentCallListResponseBodyModelList {
	s.TransferTime = &v
	return s
}

func (s *AgentCallListResponseBodyModelList) Validate() error {
	return dara.Validate(s)
}
