// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDetailCdrIbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkDetailCdrIbResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkDetailCdrIbResponseBody
	GetCode() *string
	SetData(v *ClinkDetailCdrIbResponseBodyData) *ClinkDetailCdrIbResponseBody
	GetData() *ClinkDetailCdrIbResponseBodyData
	SetMessage(v string) *ClinkDetailCdrIbResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkDetailCdrIbResponseBody
	GetRequestId() *string
}

type ClinkDetailCdrIbResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkDetailCdrIbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkDetailCdrIbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkDetailCdrIbResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkDetailCdrIbResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkDetailCdrIbResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkDetailCdrIbResponseBody) GetData() *ClinkDetailCdrIbResponseBodyData {
	return s.Data
}

func (s *ClinkDetailCdrIbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkDetailCdrIbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkDetailCdrIbResponseBody) SetAccessDeniedDetail(v string) *ClinkDetailCdrIbResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBody) SetCode(v string) *ClinkDetailCdrIbResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBody) SetData(v *ClinkDetailCdrIbResponseBodyData) *ClinkDetailCdrIbResponseBody {
	s.Data = v
	return s
}

func (s *ClinkDetailCdrIbResponseBody) SetMessage(v string) *ClinkDetailCdrIbResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBody) SetRequestId(v string) *ClinkDetailCdrIbResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkDetailCdrIbResponseBodyData struct {
	// [客户来电记录详情] #客户来电记录详情
	CdrIbDetail []*ClinkDetailCdrIbResponseBodyDataCdrIbDetail `json:"CdrIbDetail,omitempty" xml:"CdrIbDetail,omitempty" type:"Repeated"`
	// 请求 id
	//
	// example:
	//
	// 示例值
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
}

func (s ClinkDetailCdrIbResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkDetailCdrIbResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkDetailCdrIbResponseBodyData) GetCdrIbDetail() []*ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	return s.CdrIbDetail
}

func (s *ClinkDetailCdrIbResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkDetailCdrIbResponseBodyData) SetCdrIbDetail(v []*ClinkDetailCdrIbResponseBodyDataCdrIbDetail) *ClinkDetailCdrIbResponseBodyData {
	s.CdrIbDetail = v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyData) SetClinkRequestId(v string) *ClinkDetailCdrIbResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyData) Validate() error {
	if s.CdrIbDetail != nil {
		for _, item := range s.CdrIbDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkDetailCdrIbResponseBodyDataCdrIbDetail struct {
	// 首次及时应答
	//
	// example:
	//
	// ""
	AgentAnswerInTime *string `json:"AgentAnswerInTime,omitempty" xml:"AgentAnswerInTime,omitempty"`
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
	// 17
	BridgeDuration *int64 `json:"BridgeDuration,omitempty" xml:"BridgeDuration,omitempty"`
	// 首次接听时间
	//
	// example:
	//
	// 1713848599
	BridgeTime *int64 `json:"BridgeTime,omitempty" xml:"BridgeTime,omitempty"`
	// CallID
	//
	// example:
	//
	// null
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 客户号码城市
	//
	// example:
	//
	// 示例值示例值示例值
	CustomerCity *string `json:"CustomerCity,omitempty" xml:"CustomerCity,omitempty"`
	// 客户号码
	//
	// example:
	//
	// 138xxxx8888
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
	// 示例值
	CustomerProvince *string `json:"CustomerProvince,omitempty" xml:"CustomerProvince,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 1713848599
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 是否邀评
	//
	// example:
	//
	// null
	Evaluation *string `json:"Evaluation,omitempty" xml:"Evaluation,omitempty"`
	// 首呼座席姓名
	//
	// example:
	//
	// zhangsan
	FirstCallCname *string `json:"FirstCallCname,omitempty" xml:"FirstCallCname,omitempty"`
	// 首呼座席工号
	//
	// example:
	//
	// 2233
	FirstCallCno *string `json:"FirstCallCno,omitempty" xml:"FirstCallCno,omitempty"`
	// 首呼座席电话
	//
	// example:
	//
	// xxx
	FirstCallNumber *string `json:"FirstCallNumber,omitempty" xml:"FirstCallNumber,omitempty"`
	// 首呼队列名称
	//
	// example:
	//
	// xxx
	FirstCallQname *string `json:"FirstCallQname,omitempty" xml:"FirstCallQname,omitempty"`
	// 首呼队列号
	//
	// example:
	//
	// 223
	FirstCallQno *string `json:"FirstCallQno,omitempty" xml:"FirstCallQno,omitempty"`
	// 首次进入队列时间
	//
	// example:
	//
	// 1713848594
	FirstJoinQueueTime *int64 `json:"FirstJoinQueueTime,omitempty" xml:"FirstJoinQueueTime,omitempty"`
	// 首次离开队列时间
	//
	// example:
	//
	// 1713848595
	FirstLeaveQueueTime *int64 `json:"FirstLeaveQueueTime,omitempty" xml:"FirstLeaveQueueTime,omitempty"`
	// 首呼队列排队时长
	//
	// example:
	//
	// 12233
	FirstQueueDuration *int64 `json:"FirstQueueDuration,omitempty" xml:"FirstQueueDuration,omitempty"`
	// 热线号码
	//
	// example:
	//
	// xxx
	Hotline *string `json:"Hotline,omitempty" xml:"Hotline,omitempty"`
	// 热线别名
	//
	// example:
	//
	// hotline3
	HotlineName *string `json:"HotlineName,omitempty" xml:"HotlineName,omitempty"`
	// [满意度记录] #满意度记录
	Investigation *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation `json:"Investigation,omitempty" xml:"Investigation,omitempty" type:"Struct"`
	// 满意度评价
	//
	// example:
	//
	// 38
	InvestigationKeys *int64 `json:"InvestigationKeys,omitempty" xml:"InvestigationKeys,omitempty"`
	// [路由和IVR] #路由和IVR
	IvrFlows []*ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows `json:"IvrFlows,omitempty" xml:"IvrFlows,omitempty" type:"Repeated"`
	// IVR名称
	//
	// example:
	//
	// xxx
	IvrName *string `json:"IvrName,omitempty" xml:"IvrName,omitempty"`
	// 通话ID
	//
	// example:
	//
	// medias_1-162046xxxx.12
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	// 备注
	//
	// example:
	//
	// mark
	MarkData *string `json:"MarkData,omitempty" xml:"MarkData,omitempty"`
	// 挂断方
	//
	// example:
	//
	// 示例值示例值
	OnHookSource *string `json:"OnHookSource,omitempty" xml:"OnHookSource,omitempty"`
	// 录音文件
	//
	// example:
	//
	// null
	RecordFile *string `json:"RecordFile,omitempty" xml:"RecordFile,omitempty"`
	// rtcUid
	//
	// example:
	//
	// xxx
	RtcUid *string `json:"RtcUid,omitempty" xml:"RtcUid,omitempty"`
	// 语音播报时长
	//
	// example:
	//
	// 10
	SayVoiceDuration *int64 `json:"SayVoiceDuration,omitempty" xml:"SayVoiceDuration,omitempty"`
	// 开始时间
	//
	// example:
	//
	// 1713848570
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 接听状态
	//
	// example:
	//
	// 示例值示例值
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 呼叫结果
	//
	// example:
	//
	// 示例值示例值
	StatusResult *string `json:"StatusResult,omitempty" xml:"StatusResult,omitempty"`
	// 机器人接听状态
	//
	// example:
	//
	// null
	StatusRobot *string `json:"StatusRobot,omitempty" xml:"StatusRobot,omitempty"`
	// 通话标签
	TagNames []*string `json:"TagNames,omitempty" xml:"TagNames,omitempty" type:"Repeated"`
	// 总时长
	//
	// example:
	//
	// 98
	TotalDuration *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// 自定义字段
	//
	// example:
	//
	// xxx
	UserField *string `json:"UserField,omitempty" xml:"UserField,omitempty"`
	// WebRTCCallID
	//
	// example:
	//
	// null
	WebrtcCallId *string `json:"WebrtcCallId,omitempty" xml:"WebrtcCallId,omitempty"`
	// 虚拟号码
	//
	// example:
	//
	// xxx
	Xnumber *string `json:"Xnumber,omitempty" xml:"Xnumber,omitempty"`
}

func (s ClinkDetailCdrIbResponseBodyDataCdrIbDetail) String() string {
	return dara.Prettify(s)
}

func (s ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GoString() string {
	return s.String()
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetAgentAnswerInTime() *string {
	return s.AgentAnswerInTime
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetBindType() *int64 {
	return s.BindType
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetBridgeDuration() *int64 {
	return s.BridgeDuration
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetBridgeTime() *int64 {
	return s.BridgeTime
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetCallId() *string {
	return s.CallId
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetCustomerCity() *string {
	return s.CustomerCity
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetCustomerNumberEncrypt() *string {
	return s.CustomerNumberEncrypt
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetCustomerProvince() *string {
	return s.CustomerProvince
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetEvaluation() *string {
	return s.Evaluation
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetFirstCallCname() *string {
	return s.FirstCallCname
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetFirstCallCno() *string {
	return s.FirstCallCno
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetFirstCallNumber() *string {
	return s.FirstCallNumber
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetFirstCallQname() *string {
	return s.FirstCallQname
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetFirstCallQno() *string {
	return s.FirstCallQno
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetFirstJoinQueueTime() *int64 {
	return s.FirstJoinQueueTime
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetFirstLeaveQueueTime() *int64 {
	return s.FirstLeaveQueueTime
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetFirstQueueDuration() *int64 {
	return s.FirstQueueDuration
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetHotline() *string {
	return s.Hotline
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetHotlineName() *string {
	return s.HotlineName
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetInvestigation() *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation {
	return s.Investigation
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetInvestigationKeys() *int64 {
	return s.InvestigationKeys
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetIvrFlows() []*ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows {
	return s.IvrFlows
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetIvrName() *string {
	return s.IvrName
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetMarkData() *string {
	return s.MarkData
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetOnHookSource() *string {
	return s.OnHookSource
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetRecordFile() *string {
	return s.RecordFile
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetRtcUid() *string {
	return s.RtcUid
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetSayVoiceDuration() *int64 {
	return s.SayVoiceDuration
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetStatus() *string {
	return s.Status
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetStatusResult() *string {
	return s.StatusResult
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetStatusRobot() *string {
	return s.StatusRobot
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetTagNames() []*string {
	return s.TagNames
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetUserField() *string {
	return s.UserField
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetWebrtcCallId() *string {
	return s.WebrtcCallId
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) GetXnumber() *string {
	return s.Xnumber
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetAgentAnswerInTime(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.AgentAnswerInTime = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetBindType(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.BindType = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetBridgeDuration(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.BridgeDuration = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetBridgeTime(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.BridgeTime = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetCallId(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.CallId = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetCustomerCity(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.CustomerCity = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetCustomerNumber(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.CustomerNumber = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetCustomerNumberEncrypt(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.CustomerNumberEncrypt = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetCustomerProvince(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.CustomerProvince = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetEndTime(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.EndTime = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetEvaluation(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.Evaluation = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetFirstCallCname(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.FirstCallCname = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetFirstCallCno(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.FirstCallCno = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetFirstCallNumber(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.FirstCallNumber = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetFirstCallQname(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.FirstCallQname = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetFirstCallQno(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.FirstCallQno = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetFirstJoinQueueTime(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.FirstJoinQueueTime = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetFirstLeaveQueueTime(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.FirstLeaveQueueTime = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetFirstQueueDuration(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.FirstQueueDuration = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetHotline(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.Hotline = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetHotlineName(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.HotlineName = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetInvestigation(v *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.Investigation = v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetInvestigationKeys(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.InvestigationKeys = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetIvrFlows(v []*ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.IvrFlows = v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetIvrName(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.IvrName = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetMainUniqueId(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetMarkData(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.MarkData = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetOnHookSource(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.OnHookSource = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetRecordFile(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.RecordFile = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetRtcUid(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.RtcUid = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetSayVoiceDuration(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.SayVoiceDuration = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetStartTime(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.StartTime = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetStatus(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.Status = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetStatusResult(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.StatusResult = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetStatusRobot(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.StatusRobot = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetTagNames(v []*string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.TagNames = v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetTotalDuration(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.TotalDuration = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetUserField(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.UserField = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetWebrtcCallId(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.WebrtcCallId = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) SetXnumber(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetail {
	s.Xnumber = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetail) Validate() error {
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

type ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation struct {
	// 座席名称
	//
	// example:
	//
	// 示例值示例值
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// 座席号
	//
	// example:
	//
	// 11221
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 1731483979
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 满意度导航
	//
	// example:
	//
	// 示例值
	InvestigationName *string `json:"InvestigationName,omitempty" xml:"InvestigationName,omitempty"`
	// 按键值
	//
	// example:
	//
	// 7
	Keys *int64 `json:"Keys,omitempty" xml:"Keys,omitempty"`
	// 多按键值
	//
	// example:
	//
	// 3
	MultiKeys *string `json:"MultiKeys,omitempty" xml:"MultiKeys,omitempty"`
	// 开始时间
	//
	// example:
	//
	// 1731483974
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) String() string {
	return dara.Prettify(s)
}

func (s ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) GoString() string {
	return s.String()
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) GetClientName() *string {
	return s.ClientName
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) GetCno() *string {
	return s.Cno
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) GetInvestigationName() *string {
	return s.InvestigationName
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) GetKeys() *int64 {
	return s.Keys
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) GetMultiKeys() *string {
	return s.MultiKeys
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) SetClientName(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation {
	s.ClientName = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) SetCno(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation {
	s.Cno = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) SetEndTime(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation {
	s.EndTime = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) SetInvestigationName(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation {
	s.InvestigationName = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) SetKeys(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation {
	s.Keys = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) SetMultiKeys(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation {
	s.MultiKeys = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) SetStartTime(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation {
	s.StartTime = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailInvestigation) Validate() error {
	return dara.Validate(s)
}

type ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows struct {
	// 执行动作
	//
	// example:
	//
	// 3
	Action *int64 `json:"Action,omitempty" xml:"Action,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 1731483927
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// ivr 名称
	//
	// example:
	//
	// 78
	Name *int64 `json:"Name,omitempty" xml:"Name,omitempty"`
	// 节点 id
	//
	// example:
	//
	// xxx
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 节点名称
	//
	// example:
	//
	// 示例值示例值
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
	// 按键值
	//
	// example:
	//
	// -
	PressKey *string `json:"PressKey,omitempty" xml:"PressKey,omitempty"`
	// 按键时间
	//
	// example:
	//
	// 1731483814
	PressTime *int64 `json:"PressTime,omitempty" xml:"PressTime,omitempty"`
	// 路由结束时间
	//
	// example:
	//
	// 1731483897
	RouterEndTime *int64 `json:"RouterEndTime,omitempty" xml:"RouterEndTime,omitempty"`
	// 路由名称
	//
	// example:
	//
	// 78
	RouterName *int64 `json:"RouterName,omitempty" xml:"RouterName,omitempty"`
	// 路由开始时间
	//
	// example:
	//
	// 1731483814
	RouterStartTime *int64 `json:"RouterStartTime,omitempty" xml:"RouterStartTime,omitempty"`
	// 开始时间
	//
	// example:
	//
	// 1731483814
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) String() string {
	return dara.Prettify(s)
}

func (s ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) GoString() string {
	return s.String()
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) GetAction() *int64 {
	return s.Action
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) GetName() *int64 {
	return s.Name
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) GetPath() *string {
	return s.Path
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) GetPathName() *string {
	return s.PathName
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) GetPressKey() *string {
	return s.PressKey
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) GetPressTime() *int64 {
	return s.PressTime
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) GetRouterEndTime() *int64 {
	return s.RouterEndTime
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) GetRouterName() *int64 {
	return s.RouterName
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) GetRouterStartTime() *int64 {
	return s.RouterStartTime
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) SetAction(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows {
	s.Action = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) SetEndTime(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows {
	s.EndTime = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) SetName(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows {
	s.Name = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) SetPath(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows {
	s.Path = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) SetPathName(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows {
	s.PathName = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) SetPressKey(v string) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows {
	s.PressKey = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) SetPressTime(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows {
	s.PressTime = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) SetRouterEndTime(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows {
	s.RouterEndTime = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) SetRouterName(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows {
	s.RouterName = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) SetRouterStartTime(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows {
	s.RouterStartTime = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) SetStartTime(v int64) *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows {
	s.StartTime = &v
	return s
}

func (s *ClinkDetailCdrIbResponseBodyDataCdrIbDetailIvrFlows) Validate() error {
	return dara.Validate(s)
}
