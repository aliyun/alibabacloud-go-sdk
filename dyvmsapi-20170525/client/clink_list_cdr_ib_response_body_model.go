// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListCdrIbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkListCdrIbResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkListCdrIbResponseBody
	GetCode() *string
	SetData(v *ClinkListCdrIbResponseBodyData) *ClinkListCdrIbResponseBody
	GetData() *ClinkListCdrIbResponseBodyData
	SetMessage(v string) *ClinkListCdrIbResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkListCdrIbResponseBody
	GetRequestId() *string
}

type ClinkListCdrIbResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkListCdrIbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkListCdrIbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrIbResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkListCdrIbResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkListCdrIbResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkListCdrIbResponseBody) GetData() *ClinkListCdrIbResponseBodyData {
	return s.Data
}

func (s *ClinkListCdrIbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkListCdrIbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkListCdrIbResponseBody) SetAccessDeniedDetail(v string) *ClinkListCdrIbResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkListCdrIbResponseBody) SetCode(v string) *ClinkListCdrIbResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkListCdrIbResponseBody) SetData(v *ClinkListCdrIbResponseBodyData) *ClinkListCdrIbResponseBody {
	s.Data = v
	return s
}

func (s *ClinkListCdrIbResponseBody) SetMessage(v string) *ClinkListCdrIbResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkListCdrIbResponseBody) SetRequestId(v string) *ClinkListCdrIbResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkListCdrIbResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkListCdrIbResponseBodyData struct {
	// [客户来电记录列表] #客户来电记录列表
	CdrIb []*ClinkListCdrIbResponseBodyDataCdrIb `json:"CdrIb,omitempty" xml:"CdrIb,omitempty" type:"Repeated"`
	// 请求 id
	//
	// example:
	//
	// 1713848594
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
	// 69
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ClinkListCdrIbResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrIbResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkListCdrIbResponseBodyData) GetCdrIb() []*ClinkListCdrIbResponseBodyDataCdrIb {
	return s.CdrIb
}

func (s *ClinkListCdrIbResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkListCdrIbResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ClinkListCdrIbResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ClinkListCdrIbResponseBodyData) GetScrollId() *string {
	return s.ScrollId
}

func (s *ClinkListCdrIbResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ClinkListCdrIbResponseBodyData) SetCdrIb(v []*ClinkListCdrIbResponseBodyDataCdrIb) *ClinkListCdrIbResponseBodyData {
	s.CdrIb = v
	return s
}

func (s *ClinkListCdrIbResponseBodyData) SetClinkRequestId(v string) *ClinkListCdrIbResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyData) SetPageNumber(v int64) *ClinkListCdrIbResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyData) SetPageSize(v int64) *ClinkListCdrIbResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyData) SetScrollId(v string) *ClinkListCdrIbResponseBodyData {
	s.ScrollId = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyData) SetTotalCount(v int64) *ClinkListCdrIbResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyData) Validate() error {
	if s.CdrIb != nil {
		for _, item := range s.CdrIb {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkListCdrIbResponseBodyDataCdrIb struct {
	// 首次及时应答
	//
	// example:
	//
	// 1775024775
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
	// 32
	BridgeDuration *int64 `json:"BridgeDuration,omitempty" xml:"BridgeDuration,omitempty"`
	// 首次接听时间
	//
	// example:
	//
	// 6
	BridgeTime *int64 `json:"BridgeTime,omitempty" xml:"BridgeTime,omitempty"`
	// 业务记录ID
	BusinessIds []*string `json:"BusinessIds,omitempty" xml:"BusinessIds,omitempty" type:"Repeated"`
	// CallID
	//
	// example:
	//
	// xxx
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 客户号码城市
	//
	// example:
	//
	// 示例值示例值
	CustomerCity *string `json:"CustomerCity,omitempty" xml:"CustomerCity,omitempty"`
	// 客户号码
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
	// xxx
	CustomerProvince *string `json:"CustomerProvince,omitempty" xml:"CustomerProvince,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 1775024775
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
	// xxx
	FirstCallCname *string `json:"FirstCallCname,omitempty" xml:"FirstCallCname,omitempty"`
	// 首呼座席工号
	//
	// example:
	//
	// 1212
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
	// 示例值示例值
	FirstCallQname *string `json:"FirstCallQname,omitempty" xml:"FirstCallQname,omitempty"`
	// 首呼队列号
	//
	// example:
	//
	// 1111
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
	// 19
	FirstLeaveQueueTime *int64 `json:"FirstLeaveQueueTime,omitempty" xml:"FirstLeaveQueueTime,omitempty"`
	// 首呼队列排队时长
	//
	// example:
	//
	// 30
	FirstQueueDuration *int64 `json:"FirstQueueDuration,omitempty" xml:"FirstQueueDuration,omitempty"`
	// 热线号码
	//
	// example:
	//
	// 示例值示例值示例值
	Hotline *string `json:"Hotline,omitempty" xml:"Hotline,omitempty"`
	// 热线别名
	//
	// example:
	//
	// 示例值示例值
	HotlineName *string `json:"HotlineName,omitempty" xml:"HotlineName,omitempty"`
	// 满意度评价
	//
	// example:
	//
	// 59
	InvestigationKeys *int64 `json:"InvestigationKeys,omitempty" xml:"InvestigationKeys,omitempty"`
	// IVR名称
	//
	// example:
	//
	// 示例值示例值示例值
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
	// 示例值
	MarkData *string `json:"MarkData,omitempty" xml:"MarkData,omitempty"`
	// 挂断方
	//
	// example:
	//
	// xxx
	OnHookSource *string `json:"OnHookSource,omitempty" xml:"OnHookSource,omitempty"`
	// 录音文件
	//
	// example:
	//
	// xxx
	RecordFile *string `json:"RecordFile,omitempty" xml:"RecordFile,omitempty"`
	// rtcUid
	//
	// example:
	//
	// null
	RtcUid *string `json:"RtcUid,omitempty" xml:"RtcUid,omitempty"`
	// 语音播报时长
	//
	// example:
	//
	// 43
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
	// 示例值
	StatusResult *string `json:"StatusResult,omitempty" xml:"StatusResult,omitempty"`
	// 机器人接听状态
	//
	// example:
	//
	// null
	StatusRobot *string `json:"StatusRobot,omitempty" xml:"StatusRobot,omitempty"`
	// 通话标签
	TagNames []*string `json:"TagNames,omitempty" xml:"TagNames,omitempty" type:"Repeated"`
	// 工单ID
	TicketIds []*string `json:"TicketIds,omitempty" xml:"TicketIds,omitempty" type:"Repeated"`
	// 总时长
	//
	// example:
	//
	// 73
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
	// xxx
	WebrtcCallId *string `json:"WebrtcCallId,omitempty" xml:"WebrtcCallId,omitempty"`
	// 虚拟号码
	//
	// example:
	//
	// 示例值示例值
	Xnumber *string `json:"Xnumber,omitempty" xml:"Xnumber,omitempty"`
}

func (s ClinkListCdrIbResponseBodyDataCdrIb) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrIbResponseBodyDataCdrIb) GoString() string {
	return s.String()
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetAgentAnswerInTime() *string {
	return s.AgentAnswerInTime
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetBindType() *int64 {
	return s.BindType
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetBridgeDuration() *int64 {
	return s.BridgeDuration
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetBridgeTime() *int64 {
	return s.BridgeTime
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetBusinessIds() []*string {
	return s.BusinessIds
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetCallId() *string {
	return s.CallId
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetCustomerCity() *string {
	return s.CustomerCity
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetCustomerNumberEncrypt() *string {
	return s.CustomerNumberEncrypt
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetCustomerProvince() *string {
	return s.CustomerProvince
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetEvaluation() *string {
	return s.Evaluation
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetFirstCallCname() *string {
	return s.FirstCallCname
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetFirstCallCno() *string {
	return s.FirstCallCno
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetFirstCallNumber() *string {
	return s.FirstCallNumber
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetFirstCallQname() *string {
	return s.FirstCallQname
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetFirstCallQno() *string {
	return s.FirstCallQno
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetFirstJoinQueueTime() *int64 {
	return s.FirstJoinQueueTime
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetFirstLeaveQueueTime() *int64 {
	return s.FirstLeaveQueueTime
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetFirstQueueDuration() *int64 {
	return s.FirstQueueDuration
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetHotline() *string {
	return s.Hotline
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetHotlineName() *string {
	return s.HotlineName
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetInvestigationKeys() *int64 {
	return s.InvestigationKeys
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetIvrName() *string {
	return s.IvrName
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetMarkData() *string {
	return s.MarkData
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetOnHookSource() *string {
	return s.OnHookSource
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetRecordFile() *string {
	return s.RecordFile
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetRtcUid() *string {
	return s.RtcUid
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetSayVoiceDuration() *int64 {
	return s.SayVoiceDuration
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetStatus() *string {
	return s.Status
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetStatusResult() *string {
	return s.StatusResult
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetStatusRobot() *string {
	return s.StatusRobot
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetTagNames() []*string {
	return s.TagNames
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetTicketIds() []*string {
	return s.TicketIds
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetUserField() *string {
	return s.UserField
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetWebrtcCallId() *string {
	return s.WebrtcCallId
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) GetXnumber() *string {
	return s.Xnumber
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetAgentAnswerInTime(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.AgentAnswerInTime = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetBindType(v int64) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.BindType = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetBridgeDuration(v int64) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.BridgeDuration = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetBridgeTime(v int64) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.BridgeTime = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetBusinessIds(v []*string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.BusinessIds = v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetCallId(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.CallId = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetCustomerCity(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.CustomerCity = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetCustomerNumber(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.CustomerNumber = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetCustomerNumberEncrypt(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.CustomerNumberEncrypt = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetCustomerProvince(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.CustomerProvince = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetEndTime(v int64) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.EndTime = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetEvaluation(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.Evaluation = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetFirstCallCname(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.FirstCallCname = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetFirstCallCno(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.FirstCallCno = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetFirstCallNumber(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.FirstCallNumber = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetFirstCallQname(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.FirstCallQname = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetFirstCallQno(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.FirstCallQno = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetFirstJoinQueueTime(v int64) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.FirstJoinQueueTime = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetFirstLeaveQueueTime(v int64) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.FirstLeaveQueueTime = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetFirstQueueDuration(v int64) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.FirstQueueDuration = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetHotline(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.Hotline = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetHotlineName(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.HotlineName = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetInvestigationKeys(v int64) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.InvestigationKeys = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetIvrName(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.IvrName = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetMainUniqueId(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetMarkData(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.MarkData = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetOnHookSource(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.OnHookSource = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetRecordFile(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.RecordFile = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetRtcUid(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.RtcUid = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetSayVoiceDuration(v int64) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.SayVoiceDuration = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetStartTime(v int64) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.StartTime = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetStatus(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.Status = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetStatusResult(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.StatusResult = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetStatusRobot(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.StatusRobot = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetTagNames(v []*string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.TagNames = v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetTicketIds(v []*string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.TicketIds = v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetTotalDuration(v int64) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.TotalDuration = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetUserField(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.UserField = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetWebrtcCallId(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.WebrtcCallId = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) SetXnumber(v string) *ClinkListCdrIbResponseBodyDataCdrIb {
	s.Xnumber = &v
	return s
}

func (s *ClinkListCdrIbResponseBodyDataCdrIb) Validate() error {
	return dara.Validate(s)
}
