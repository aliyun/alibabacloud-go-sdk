// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeCdrIbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkDescribeCdrIbResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkDescribeCdrIbResponseBody
	GetCode() *string
	SetData(v *ClinkDescribeCdrIbResponseBodyData) *ClinkDescribeCdrIbResponseBody
	GetData() *ClinkDescribeCdrIbResponseBodyData
	SetMessage(v string) *ClinkDescribeCdrIbResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkDescribeCdrIbResponseBody
	GetRequestId() *string
}

type ClinkDescribeCdrIbResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkDescribeCdrIbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 088A4C31-A613-5339-B9E7-E4C4B8BE0F00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkDescribeCdrIbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrIbResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrIbResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkDescribeCdrIbResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkDescribeCdrIbResponseBody) GetData() *ClinkDescribeCdrIbResponseBodyData {
	return s.Data
}

func (s *ClinkDescribeCdrIbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkDescribeCdrIbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkDescribeCdrIbResponseBody) SetAccessDeniedDetail(v string) *ClinkDescribeCdrIbResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBody) SetCode(v string) *ClinkDescribeCdrIbResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBody) SetData(v *ClinkDescribeCdrIbResponseBodyData) *ClinkDescribeCdrIbResponseBody {
	s.Data = v
	return s
}

func (s *ClinkDescribeCdrIbResponseBody) SetMessage(v string) *ClinkDescribeCdrIbResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBody) SetRequestId(v string) *ClinkDescribeCdrIbResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkDescribeCdrIbResponseBodyData struct {
	// [呼入通话记录] #呼入通话记录
	CdrIb *ClinkDescribeCdrIbResponseBodyDataCdrIb `json:"CdrIb,omitempty" xml:"CdrIb,omitempty" type:"Struct"`
	// 请求 id
	//
	// example:
	//
	// xxxx
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
}

func (s ClinkDescribeCdrIbResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrIbResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrIbResponseBodyData) GetCdrIb() *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	return s.CdrIb
}

func (s *ClinkDescribeCdrIbResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkDescribeCdrIbResponseBodyData) SetCdrIb(v *ClinkDescribeCdrIbResponseBodyDataCdrIb) *ClinkDescribeCdrIbResponseBodyData {
	s.CdrIb = v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyData) SetClinkRequestId(v string) *ClinkDescribeCdrIbResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyData) Validate() error {
	if s.CdrIb != nil {
		if err := s.CdrIb.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkDescribeCdrIbResponseBodyDataCdrIb struct {
	// 接通时长
	//
	// example:
	//
	// 33
	BridgeDuration *int64 `json:"BridgeDuration,omitempty" xml:"BridgeDuration,omitempty"`
	// 接通时间
	//
	// example:
	//
	// 1775024775
	BridgeTime *int64 `json:"BridgeTime,omitempty" xml:"BridgeTime,omitempty"`
	// 通话记录ID
	//
	// example:
	//
	// b1651313-0e70-487c-99fd-3ea342b35b00
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 呼入类型
	//
	// example:
	//
	// 4
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 座席名称
	//
	// example:
	//
	// 示例值
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// 座席电话
	//
	// example:
	//
	// xxx
	ClientNumber *string `json:"ClientNumber,omitempty" xml:"ClientNumber,omitempty"`
	// 座席接起时间
	//
	// example:
	//
	// 1775024775
	ClientOffhookTime *int64 `json:"ClientOffhookTime,omitempty" xml:"ClientOffhookTime,omitempty"`
	// 座席响铃时间
	//
	// example:
	//
	// 1775024775
	ClientRingingTime *int64 `json:"ClientRingingTime,omitempty" xml:"ClientRingingTime,omitempty"`
	// 座席号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 客户来电城市
	//
	// example:
	//
	// 示例值示例值
	CustomerCity *string `json:"CustomerCity,omitempty" xml:"CustomerCity,omitempty"`
	// 客户来电号码，带区号
	//
	// example:
	//
	// 010xxx
	CustomerNumber *string `json:"CustomerNumber,omitempty" xml:"CustomerNumber,omitempty"`
	// 客户来电号码加密串
	//
	// example:
	//
	// xxx
	CustomerNumberEncrypt *string `json:"CustomerNumberEncrypt,omitempty" xml:"CustomerNumberEncrypt,omitempty"`
	// 客户来电省份
	//
	// example:
	//
	// xxx
	CustomerProvince *string `json:"CustomerProvince,omitempty" xml:"CustomerProvince,omitempty"`
	// 挂机方
	//
	// example:
	//
	// 示例值示例值
	EndReason *string `json:"EndReason,omitempty" xml:"EndReason,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 1775024775
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 是否邀评 0: 否 1: 是
	//
	// example:
	//
	// 39
	Evaluation *int64 `json:"Evaluation,omitempty" xml:"Evaluation,omitempty"`
	// 客户速挂，true：是；false：否
	//
	// example:
	//
	// true
	FastHangUp *bool `json:"FastHangUp,omitempty" xml:"FastHangUp,omitempty"`
	// 首次进入队列时间
	//
	// example:
	//
	// 1775024775
	FirstJoinQueueTime *int64 `json:"FirstJoinQueueTime,omitempty" xml:"FirstJoinQueueTime,omitempty"`
	// 来电热线号码
	//
	// example:
	//
	// 02788xxx888
	Hotline *string `json:"Hotline,omitempty" xml:"Hotline,omitempty"`
	// 热线别名
	//
	// example:
	//
	// 示例值
	HotlineName *string `json:"HotlineName,omitempty" xml:"HotlineName,omitempty"`
	// 座席振铃时长
	//
	// example:
	//
	// 3
	IbRingingDuration *int64 `json:"IbRingingDuration,omitempty" xml:"IbRingingDuration,omitempty"`
	// 排队时长
	//
	// example:
	//
	// 96
	IbWaitDuration *int64 `json:"IbWaitDuration,omitempty" xml:"IbWaitDuration,omitempty"`
	// [满意度记录] #满意度记录
	Investigation *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation `json:"Investigation,omitempty" xml:"Investigation,omitempty" type:"Struct"`
	// [路由和IVR] #路由和IVR
	IvrFlows []*ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows `json:"IvrFlows,omitempty" xml:"IvrFlows,omitempty" type:"Repeated"`
	// IVR名称
	//
	// example:
	//
	// 示例值示例值示例值
	IvrName *string `json:"IvrName,omitempty" xml:"IvrName,omitempty"`
	// 进入队列时间
	//
	// example:
	//
	// 1775024775
	JoinQueueTime *int64 `json:"JoinQueueTime,omitempty" xml:"JoinQueueTime,omitempty"`
	// 离开队列时间
	//
	// example:
	//
	// 1775024775
	LeaveQueueTime *int64 `json:"LeaveQueueTime,omitempty" xml:"LeaveQueueTime,omitempty"`
	// 标记
	//
	// example:
	//
	// 77
	Mark *int64 `json:"Mark,omitempty" xml:"Mark,omitempty"`
	// 备注
	//
	// example:
	//
	// 示例值示例值
	MarkData *string `json:"MarkData,omitempty" xml:"MarkData,omitempty"`
	// 来电队列号
	//
	// example:
	//
	// 1122
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// 队列及时应答
	//
	// example:
	//
	// 1775024775
	QueueAnswerInTime *int64 `json:"QueueAnswerInTime,omitempty" xml:"QueueAnswerInTime,omitempty"`
	// 录音文件名
	//
	// example:
	//
	// null
	RecordFile *string `json:"RecordFile,omitempty" xml:"RecordFile,omitempty"`
	// 语音播报时长
	//
	// example:
	//
	// 3
	SayVoiceDuration *int64 `json:"SayVoiceDuration,omitempty" xml:"SayVoiceDuration,omitempty"`
	// 呼叫结果
	//
	// example:
	//
	// 示例值
	SipCause *string `json:"SipCause,omitempty" xml:"SipCause,omitempty"`
	// 开始时间
	//
	// example:
	//
	// 1775024775
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 接听状态
	//
	// example:
	//
	// 示例值示例值
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 接听状态映射
	//
	// example:
	//
	// 示例值
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// 展示通话标签详情,当请求参数fields中包含tagNames时返回
	TagNames []*string `json:"TagNames,omitempty" xml:"TagNames,omitempty" type:"Repeated"`
	// 标签
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 总时长
	//
	// example:
	//
	// 87
	TotalDuration *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// 通话记录唯一标识
	//
	// example:
	//
	// 示例值
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s ClinkDescribeCdrIbResponseBodyDataCdrIb) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrIbResponseBodyDataCdrIb) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetBridgeDuration() *int64 {
	return s.BridgeDuration
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetBridgeTime() *int64 {
	return s.BridgeTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetCallId() *string {
	return s.CallId
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetCallType() *string {
	return s.CallType
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetClientName() *string {
	return s.ClientName
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetClientNumber() *string {
	return s.ClientNumber
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetClientOffhookTime() *int64 {
	return s.ClientOffhookTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetClientRingingTime() *int64 {
	return s.ClientRingingTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetCno() *string {
	return s.Cno
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetCustomerCity() *string {
	return s.CustomerCity
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetCustomerNumberEncrypt() *string {
	return s.CustomerNumberEncrypt
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetCustomerProvince() *string {
	return s.CustomerProvince
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetEndReason() *string {
	return s.EndReason
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetEvaluation() *int64 {
	return s.Evaluation
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetFastHangUp() *bool {
	return s.FastHangUp
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetFirstJoinQueueTime() *int64 {
	return s.FirstJoinQueueTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetHotline() *string {
	return s.Hotline
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetHotlineName() *string {
	return s.HotlineName
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetIbRingingDuration() *int64 {
	return s.IbRingingDuration
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetIbWaitDuration() *int64 {
	return s.IbWaitDuration
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetInvestigation() *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation {
	return s.Investigation
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetIvrFlows() []*ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows {
	return s.IvrFlows
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetIvrName() *string {
	return s.IvrName
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetJoinQueueTime() *int64 {
	return s.JoinQueueTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetLeaveQueueTime() *int64 {
	return s.LeaveQueueTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetMark() *int64 {
	return s.Mark
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetMarkData() *string {
	return s.MarkData
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetQno() *string {
	return s.Qno
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetQueueAnswerInTime() *int64 {
	return s.QueueAnswerInTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetRecordFile() *string {
	return s.RecordFile
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetSayVoiceDuration() *int64 {
	return s.SayVoiceDuration
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetSipCause() *string {
	return s.SipCause
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetStatus() *string {
	return s.Status
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetTagNames() []*string {
	return s.TagNames
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetTags() []*string {
	return s.Tags
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) GetUniqueId() *string {
	return s.UniqueId
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetBridgeDuration(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.BridgeDuration = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetBridgeTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.BridgeTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetCallId(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.CallId = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetCallType(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.CallType = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetClientName(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.ClientName = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetClientNumber(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.ClientNumber = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetClientOffhookTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.ClientOffhookTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetClientRingingTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.ClientRingingTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetCno(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.Cno = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetCustomerCity(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.CustomerCity = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetCustomerNumber(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.CustomerNumber = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetCustomerNumberEncrypt(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.CustomerNumberEncrypt = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetCustomerProvince(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.CustomerProvince = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetEndReason(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.EndReason = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetEndTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.EndTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetEvaluation(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.Evaluation = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetFastHangUp(v bool) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.FastHangUp = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetFirstJoinQueueTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.FirstJoinQueueTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetHotline(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.Hotline = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetHotlineName(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.HotlineName = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetIbRingingDuration(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.IbRingingDuration = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetIbWaitDuration(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.IbWaitDuration = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetInvestigation(v *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.Investigation = v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetIvrFlows(v []*ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.IvrFlows = v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetIvrName(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.IvrName = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetJoinQueueTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.JoinQueueTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetLeaveQueueTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.LeaveQueueTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetMark(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.Mark = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetMarkData(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.MarkData = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetQno(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.Qno = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetQueueAnswerInTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.QueueAnswerInTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetRecordFile(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.RecordFile = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetSayVoiceDuration(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.SayVoiceDuration = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetSipCause(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.SipCause = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetStartTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.StartTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetStatus(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.Status = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetStatusCode(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.StatusCode = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetTagNames(v []*string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.TagNames = v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetTags(v []*string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.Tags = v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetTotalDuration(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.TotalDuration = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) SetUniqueId(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIb {
	s.UniqueId = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIb) Validate() error {
	if s.Investigation != nil {
		if err := s.Investigation.Validate(); err != nil {
			return err
		}
	}
	if s.IvrFlows != nil {
		for _, item := range s.IvrFlows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation struct {
	// 座席名称
	//
	// example:
	//
	// xxx
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// 座席号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 1775024775
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 按键值
	//
	// example:
	//
	// 9
	Keys *int64 `json:"Keys,omitempty" xml:"Keys,omitempty"`
	// 多按键值
	//
	// example:
	//
	// xxx
	MultiKeys *string `json:"MultiKeys,omitempty" xml:"MultiKeys,omitempty"`
	// 开始时间
	//
	// example:
	//
	// 1775024775
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) GetClientName() *string {
	return s.ClientName
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) GetCno() *string {
	return s.Cno
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) GetKeys() *int64 {
	return s.Keys
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) GetMultiKeys() *string {
	return s.MultiKeys
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) SetClientName(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation {
	s.ClientName = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) SetCno(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation {
	s.Cno = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) SetEndTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation {
	s.EndTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) SetKeys(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation {
	s.Keys = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) SetMultiKeys(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation {
	s.MultiKeys = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) SetStartTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation {
	s.StartTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbInvestigation) Validate() error {
	return dara.Validate(s)
}

type ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows struct {
	// 执行动作
	//
	// example:
	//
	// 1
	Action *int64 `json:"Action,omitempty" xml:"Action,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 1775024775
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// ivr 名称
	//
	// example:
	//
	// --
	IvrName *int64 `json:"IvrName,omitempty" xml:"IvrName,omitempty"`
	// 节点 id
	//
	// example:
	//
	// 13321
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 节点名称
	//
	// example:
	//
	// -
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
	// 按键值
	//
	// example:
	//
	// 6
	PressKey *string `json:"PressKey,omitempty" xml:"PressKey,omitempty"`
	// 按键时间
	//
	// example:
	//
	// 1775024775
	PressTime *int64 `json:"PressTime,omitempty" xml:"PressTime,omitempty"`
	// 路由结束时间
	//
	// example:
	//
	// 1775024775
	RouterEndTime *int64 `json:"RouterEndTime,omitempty" xml:"RouterEndTime,omitempty"`
	// 路由名称
	//
	// example:
	//
	// 8
	RouterName *int64 `json:"RouterName,omitempty" xml:"RouterName,omitempty"`
	// 路由开始时间
	//
	// example:
	//
	// 1775024775
	RouterStartTime *int64 `json:"RouterStartTime,omitempty" xml:"RouterStartTime,omitempty"`
	// 开始时间
	//
	// example:
	//
	// 1775024775
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) GetAction() *int64 {
	return s.Action
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) GetIvrName() *int64 {
	return s.IvrName
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) GetPath() *string {
	return s.Path
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) GetPathName() *string {
	return s.PathName
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) GetPressKey() *string {
	return s.PressKey
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) GetPressTime() *int64 {
	return s.PressTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) GetRouterEndTime() *int64 {
	return s.RouterEndTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) GetRouterName() *int64 {
	return s.RouterName
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) GetRouterStartTime() *int64 {
	return s.RouterStartTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) SetAction(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows {
	s.Action = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) SetEndTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows {
	s.EndTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) SetIvrName(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows {
	s.IvrName = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) SetPath(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows {
	s.Path = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) SetPathName(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows {
	s.PathName = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) SetPressKey(v string) *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows {
	s.PressKey = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) SetPressTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows {
	s.PressTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) SetRouterEndTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows {
	s.RouterEndTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) SetRouterName(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows {
	s.RouterName = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) SetRouterStartTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows {
	s.RouterStartTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) SetStartTime(v int64) *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows {
	s.StartTime = &v
	return s
}

func (s *ClinkDescribeCdrIbResponseBodyDataCdrIbIvrFlows) Validate() error {
	return dara.Validate(s)
}
