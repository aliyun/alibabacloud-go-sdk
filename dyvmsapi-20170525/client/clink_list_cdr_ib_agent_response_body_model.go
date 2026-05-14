// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListCdrIbAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkListCdrIbAgentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkListCdrIbAgentResponseBody
	GetCode() *string
	SetData(v *ClinkListCdrIbAgentResponseBodyData) *ClinkListCdrIbAgentResponseBody
	GetData() *ClinkListCdrIbAgentResponseBodyData
	SetMessage(v string) *ClinkListCdrIbAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkListCdrIbAgentResponseBody
	GetRequestId() *string
}

type ClinkListCdrIbAgentResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkListCdrIbAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkListCdrIbAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrIbAgentResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkListCdrIbAgentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkListCdrIbAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkListCdrIbAgentResponseBody) GetData() *ClinkListCdrIbAgentResponseBodyData {
	return s.Data
}

func (s *ClinkListCdrIbAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkListCdrIbAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkListCdrIbAgentResponseBody) SetAccessDeniedDetail(v string) *ClinkListCdrIbAgentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBody) SetCode(v string) *ClinkListCdrIbAgentResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBody) SetData(v *ClinkListCdrIbAgentResponseBodyData) *ClinkListCdrIbAgentResponseBody {
	s.Data = v
	return s
}

func (s *ClinkListCdrIbAgentResponseBody) SetMessage(v string) *ClinkListCdrIbAgentResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBody) SetRequestId(v string) *ClinkListCdrIbAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkListCdrIbAgentResponseBodyData struct {
	// [座席接听记录列表] #座席接听记录列表
	CdrIbAgent []*ClinkListCdrIbAgentResponseBodyDataCdrIbAgent `json:"CdrIbAgent,omitempty" xml:"CdrIbAgent,omitempty" type:"Repeated"`
	// 请求 id
	//
	// example:
	//
	// xxx
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	// 当前页码
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 一页展示条数
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 滚动查询ID
	//
	// example:
	//
	// xxx
	ScrollId *string `json:"ScrollId,omitempty" xml:"ScrollId,omitempty"`
	// 总条数
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ClinkListCdrIbAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrIbAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkListCdrIbAgentResponseBodyData) GetCdrIbAgent() []*ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	return s.CdrIbAgent
}

func (s *ClinkListCdrIbAgentResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkListCdrIbAgentResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ClinkListCdrIbAgentResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ClinkListCdrIbAgentResponseBodyData) GetScrollId() *string {
	return s.ScrollId
}

func (s *ClinkListCdrIbAgentResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ClinkListCdrIbAgentResponseBodyData) SetCdrIbAgent(v []*ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) *ClinkListCdrIbAgentResponseBodyData {
	s.CdrIbAgent = v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyData) SetClinkRequestId(v string) *ClinkListCdrIbAgentResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyData) SetPageNumber(v int64) *ClinkListCdrIbAgentResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyData) SetPageSize(v int64) *ClinkListCdrIbAgentResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyData) SetScrollId(v string) *ClinkListCdrIbAgentResponseBodyData {
	s.ScrollId = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyData) SetTotalCount(v int64) *ClinkListCdrIbAgentResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyData) Validate() error {
	if s.CdrIbAgent != nil {
		for _, item := range s.CdrIbAgent {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkListCdrIbAgentResponseBodyDataCdrIbAgent struct {
	// 及时应答
	//
	// example:
	//
	// 1
	AgentAnswerInTime *int64 `json:"AgentAnswerInTime,omitempty" xml:"AgentAnswerInTime,omitempty"`
	// 接听时间
	//
	// example:
	//
	// 1775024775
	AnswerTime *int64 `json:"AnswerTime,omitempty" xml:"AnswerTime,omitempty"`
	// 接听设备
	//
	// example:
	//
	// 1
	BindType *int64 `json:"BindType,omitempty" xml:"BindType,omitempty"`
	// 通话时长
	//
	// example:
	//
	// 30
	BridgeDuration *int64 `json:"BridgeDuration,omitempty" xml:"BridgeDuration,omitempty"`
	// 接听状态
	//
	// example:
	//
	// 示例值示例值
	BridgeStatus *string `json:"BridgeStatus,omitempty" xml:"BridgeStatus,omitempty"`
	// 呼叫结果
	//
	// example:
	//
	// -
	BridgeStatusDetail *string `json:"BridgeStatusDetail,omitempty" xml:"BridgeStatusDetail,omitempty"`
	// CallID
	//
	// example:
	//
	// xxx
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 座席名称
	//
	// example:
	//
	// ClientName
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// 座席电话
	//
	// example:
	//
	// xxx
	ClientNumber *string `json:"ClientNumber,omitempty" xml:"ClientNumber,omitempty"`
	// 座席号
	//
	// example:
	//
	// 1122
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 客户号码城市
	//
	// example:
	//
	// 示例值
	CustomerCity *string `json:"CustomerCity,omitempty" xml:"CustomerCity,omitempty"`
	// 客户号码，带区号
	//
	// example:
	//
	// xxx
	CustomerNumber *string `json:"CustomerNumber,omitempty" xml:"CustomerNumber,omitempty"`
	// 客户号码加密串
	//
	// example:
	//
	// xxx
	CustomerNumberEncrypt *string `json:"CustomerNumberEncrypt,omitempty" xml:"CustomerNumberEncrypt,omitempty"`
	// 客户号码省份
	//
	// example:
	//
	// 示例值示例值
	CustomerProvince *string `json:"CustomerProvince,omitempty" xml:"CustomerProvince,omitempty"`
	// 呼叫类型
	//
	// example:
	//
	// 示例值示例值
	DetailCallType *string `json:"DetailCallType,omitempty" xml:"DetailCallType,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 82
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 热线号码
	//
	// example:
	//
	// Hotline
	Hotline *string `json:"Hotline,omitempty" xml:"Hotline,omitempty"`
	// 热线别名
	//
	// example:
	//
	// HotlineName
	HotlineName *string `json:"HotlineName,omitempty" xml:"HotlineName,omitempty"`
	// 振铃时长
	//
	// example:
	//
	// 3
	IbRingingDuration *int64 `json:"IbRingingDuration,omitempty" xml:"IbRingingDuration,omitempty"`
	// 通话ID
	//
	// example:
	//
	// medias_1-162046xxxx.12
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	// 座席接起时间
	//
	// example:
	//
	// 1775024775
	OffhookTime *int64 `json:"OffhookTime,omitempty" xml:"OffhookTime,omitempty"`
	// 挂断方
	//
	// example:
	//
	// 示例值
	OnHookSource *string `json:"OnHookSource,omitempty" xml:"OnHookSource,omitempty"`
	// 来电队列名称
	//
	// example:
	//
	// 示例值示例值示例值
	Qname *string `json:"Qname,omitempty" xml:"Qname,omitempty"`
	// 来电队列号
	//
	// example:
	//
	// 2233
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// 录音文件名
	//
	// example:
	//
	// null
	RecordFile *string `json:"RecordFile,omitempty" xml:"RecordFile,omitempty"`
	// 主叫记忆。obRemember: 外呼主叫记忆，ibRemember: 来电主叫记忆
	//
	// example:
	//
	// obRemember
	Remember *string `json:"Remember,omitempty" xml:"Remember,omitempty"`
	// 振铃时间
	//
	// example:
	//
	// 1775024775
	RingTime *int64 `json:"RingTime,omitempty" xml:"RingTime,omitempty"`
	// 语音播报时长
	//
	// example:
	//
	// 51
	SayVoiceDuration *int64 `json:"SayVoiceDuration,omitempty" xml:"SayVoiceDuration,omitempty"`
	// 呼叫情况
	//
	// example:
	//
	// 示例值
	SipCauseDesc *string `json:"SipCauseDesc,omitempty" xml:"SipCauseDesc,omitempty"`
	// 开始时间
	//
	// example:
	//
	// 1775024775
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 唯一标识
	//
	// example:
	//
	// xxx
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// WebRTCCallID
	//
	// example:
	//
	// xxx
	WebrtcCallId *string `json:"WebrtcCallId,omitempty" xml:"WebrtcCallId,omitempty"`
}

func (s ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GoString() string {
	return s.String()
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetAgentAnswerInTime() *int64 {
	return s.AgentAnswerInTime
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetAnswerTime() *int64 {
	return s.AnswerTime
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetBindType() *int64 {
	return s.BindType
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetBridgeDuration() *int64 {
	return s.BridgeDuration
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetBridgeStatus() *string {
	return s.BridgeStatus
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetBridgeStatusDetail() *string {
	return s.BridgeStatusDetail
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetCallId() *string {
	return s.CallId
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetClientName() *string {
	return s.ClientName
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetClientNumber() *string {
	return s.ClientNumber
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetCno() *string {
	return s.Cno
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetCustomerCity() *string {
	return s.CustomerCity
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetCustomerNumberEncrypt() *string {
	return s.CustomerNumberEncrypt
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetCustomerProvince() *string {
	return s.CustomerProvince
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetDetailCallType() *string {
	return s.DetailCallType
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetHotline() *string {
	return s.Hotline
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetHotlineName() *string {
	return s.HotlineName
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetIbRingingDuration() *int64 {
	return s.IbRingingDuration
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetOffhookTime() *int64 {
	return s.OffhookTime
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetOnHookSource() *string {
	return s.OnHookSource
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetQname() *string {
	return s.Qname
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetQno() *string {
	return s.Qno
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetRecordFile() *string {
	return s.RecordFile
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetRemember() *string {
	return s.Remember
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetRingTime() *int64 {
	return s.RingTime
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetSayVoiceDuration() *int64 {
	return s.SayVoiceDuration
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetSipCauseDesc() *string {
	return s.SipCauseDesc
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetUniqueId() *string {
	return s.UniqueId
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) GetWebrtcCallId() *string {
	return s.WebrtcCallId
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetAgentAnswerInTime(v int64) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.AgentAnswerInTime = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetAnswerTime(v int64) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.AnswerTime = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetBindType(v int64) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.BindType = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetBridgeDuration(v int64) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.BridgeDuration = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetBridgeStatus(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.BridgeStatus = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetBridgeStatusDetail(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.BridgeStatusDetail = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetCallId(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.CallId = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetClientName(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.ClientName = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetClientNumber(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.ClientNumber = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetCno(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.Cno = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetCustomerCity(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.CustomerCity = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetCustomerNumber(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.CustomerNumber = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetCustomerNumberEncrypt(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.CustomerNumberEncrypt = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetCustomerProvince(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.CustomerProvince = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetDetailCallType(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.DetailCallType = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetEndTime(v int64) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.EndTime = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetHotline(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.Hotline = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetHotlineName(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.HotlineName = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetIbRingingDuration(v int64) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.IbRingingDuration = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetMainUniqueId(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetOffhookTime(v int64) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.OffhookTime = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetOnHookSource(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.OnHookSource = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetQname(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.Qname = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetQno(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.Qno = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetRecordFile(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.RecordFile = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetRemember(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.Remember = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetRingTime(v int64) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.RingTime = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetSayVoiceDuration(v int64) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.SayVoiceDuration = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetSipCauseDesc(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.SipCauseDesc = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetStartTime(v int64) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.StartTime = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetUniqueId(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.UniqueId = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) SetWebrtcCallId(v string) *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent {
	s.WebrtcCallId = &v
	return s
}

func (s *ClinkListCdrIbAgentResponseBodyDataCdrIbAgent) Validate() error {
	return dara.Validate(s)
}
