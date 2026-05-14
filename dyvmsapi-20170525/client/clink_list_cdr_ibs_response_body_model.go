// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListCdrIbsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkListCdrIbsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkListCdrIbsResponseBody
	GetCode() *string
	SetData(v *ClinkListCdrIbsResponseBodyData) *ClinkListCdrIbsResponseBody
	GetData() *ClinkListCdrIbsResponseBodyData
	SetMessage(v string) *ClinkListCdrIbsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkListCdrIbsResponseBody
	GetRequestId() *string
}

type ClinkListCdrIbsResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkListCdrIbsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkListCdrIbsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrIbsResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkListCdrIbsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkListCdrIbsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkListCdrIbsResponseBody) GetData() *ClinkListCdrIbsResponseBodyData {
	return s.Data
}

func (s *ClinkListCdrIbsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkListCdrIbsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkListCdrIbsResponseBody) SetAccessDeniedDetail(v string) *ClinkListCdrIbsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkListCdrIbsResponseBody) SetCode(v string) *ClinkListCdrIbsResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkListCdrIbsResponseBody) SetData(v *ClinkListCdrIbsResponseBodyData) *ClinkListCdrIbsResponseBody {
	s.Data = v
	return s
}

func (s *ClinkListCdrIbsResponseBody) SetMessage(v string) *ClinkListCdrIbsResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkListCdrIbsResponseBody) SetRequestId(v string) *ClinkListCdrIbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkListCdrIbsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkListCdrIbsResponseBodyData struct {
	// [呼入通话记录列表] #呼入通话记录列表
	CdrIbs []*ClinkListCdrIbsResponseBodyDataCdrIbs `json:"CdrIbs,omitempty" xml:"CdrIbs,omitempty" type:"Repeated"`
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
	// 总条数
	//
	// example:
	//
	// 57
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ClinkListCdrIbsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrIbsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkListCdrIbsResponseBodyData) GetCdrIbs() []*ClinkListCdrIbsResponseBodyDataCdrIbs {
	return s.CdrIbs
}

func (s *ClinkListCdrIbsResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkListCdrIbsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ClinkListCdrIbsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ClinkListCdrIbsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ClinkListCdrIbsResponseBodyData) SetCdrIbs(v []*ClinkListCdrIbsResponseBodyDataCdrIbs) *ClinkListCdrIbsResponseBodyData {
	s.CdrIbs = v
	return s
}

func (s *ClinkListCdrIbsResponseBodyData) SetClinkRequestId(v string) *ClinkListCdrIbsResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyData) SetPageNumber(v int64) *ClinkListCdrIbsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyData) SetPageSize(v int64) *ClinkListCdrIbsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyData) SetTotalCount(v int64) *ClinkListCdrIbsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyData) Validate() error {
	if s.CdrIbs != nil {
		for _, item := range s.CdrIbs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkListCdrIbsResponseBodyDataCdrIbs struct {
	// 接通时长
	//
	// example:
	//
	// 19
	BridgeDuration *int64 `json:"BridgeDuration,omitempty" xml:"BridgeDuration,omitempty"`
	// 接通时间
	//
	// example:
	//
	// 1775024775
	BridgeTime *int64 `json:"BridgeTime,omitempty" xml:"BridgeTime,omitempty"`
	// 通话记录 Id
	//
	// example:
	//
	// b1651313-0e70-487c-99fd-3ea342b35b00
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 呼入类型
	//
	// example:
	//
	// 示例值
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
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
	// 1537329247
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
	// xxx
	CustomerNumber *string `json:"CustomerNumber,omitempty" xml:"CustomerNumber,omitempty"`
	// 客户来电号码加密串
	//
	// example:
	//
	// TNjljZjZlNjgxYmIwNjMxGEwMzA3MmQ0MDQzYWEyMjY
	CustomerNumberEncrypt *string `json:"CustomerNumberEncrypt,omitempty" xml:"CustomerNumberEncrypt,omitempty"`
	// 客户来电省份
	//
	// example:
	//
	// 示例值
	CustomerProvince *string `json:"CustomerProvince,omitempty" xml:"CustomerProvince,omitempty"`
	// 挂机方
	//
	// example:
	//
	// 示例值
	EndReason *string `json:"EndReason,omitempty" xml:"EndReason,omitempty"`
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
	// 0
	Evaluation *int64 `json:"Evaluation,omitempty" xml:"Evaluation,omitempty"`
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
	// xxx
	Hotline *string `json:"Hotline,omitempty" xml:"Hotline,omitempty"`
	// 热线别名
	//
	// example:
	//
	// HotlineName
	HotlineName *string `json:"HotlineName,omitempty" xml:"HotlineName,omitempty"`
	// 座席振铃时长
	//
	// example:
	//
	// 10
	IbRingingDuration *int64 `json:"IbRingingDuration,omitempty" xml:"IbRingingDuration,omitempty"`
	// 排队时长
	//
	// example:
	//
	// 18
	IbWaitDuration *int64 `json:"IbWaitDuration,omitempty" xml:"IbWaitDuration,omitempty"`
	// IVR名称
	//
	// example:
	//
	// IvrName
	IvrName *string `json:"IvrName,omitempty" xml:"IvrName,omitempty"`
	// 加入队列时间
	//
	// example:
	//
	// 1647255885
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
	// 3
	Mark *int64 `json:"Mark,omitempty" xml:"Mark,omitempty"`
	// 备注
	//
	// example:
	//
	// 3
	MarkData *string `json:"MarkData,omitempty" xml:"MarkData,omitempty"`
	// 来电队列名称
	//
	// example:
	//
	// Qname
	Qname *string `json:"Qname,omitempty" xml:"Qname,omitempty"`
	// 来电队列号
	//
	// example:
	//
	// 2332
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// 队列及时应答
	//
	// example:
	//
	// 0
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
	// 30
	SayVoiceDuration *int64 `json:"SayVoiceDuration,omitempty" xml:"SayVoiceDuration,omitempty"`
	// 呼叫结果
	//
	// example:
	//
	// 示例值示例值
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
	// 示例值示例值示例值
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// 展示通话标签详情,当请求参数fields中包含tagNames时返回
	TagNames []*string `json:"TagNames,omitempty" xml:"TagNames,omitempty" type:"Repeated"`
	// 标签
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 总时长
	//
	// example:
	//
	// 56
	TotalDuration *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// 通话记录唯一标识
	//
	// example:
	//
	// AWS_DEV_MEDIA_SERVER_8-1537329247.6
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s ClinkListCdrIbsResponseBodyDataCdrIbs) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrIbsResponseBodyDataCdrIbs) GoString() string {
	return s.String()
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetBridgeDuration() *int64 {
	return s.BridgeDuration
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetBridgeTime() *int64 {
	return s.BridgeTime
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetCallId() *string {
	return s.CallId
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetCallType() *string {
	return s.CallType
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetClientName() *string {
	return s.ClientName
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetClientNumber() *string {
	return s.ClientNumber
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetClientOffhookTime() *int64 {
	return s.ClientOffhookTime
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetClientRingingTime() *int64 {
	return s.ClientRingingTime
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetCno() *string {
	return s.Cno
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetCustomerCity() *string {
	return s.CustomerCity
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetCustomerNumberEncrypt() *string {
	return s.CustomerNumberEncrypt
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetCustomerProvince() *string {
	return s.CustomerProvince
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetEndReason() *string {
	return s.EndReason
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetEvaluation() *int64 {
	return s.Evaluation
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetFirstJoinQueueTime() *int64 {
	return s.FirstJoinQueueTime
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetHotline() *string {
	return s.Hotline
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetHotlineName() *string {
	return s.HotlineName
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetIbRingingDuration() *int64 {
	return s.IbRingingDuration
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetIbWaitDuration() *int64 {
	return s.IbWaitDuration
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetIvrName() *string {
	return s.IvrName
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetJoinQueueTime() *int64 {
	return s.JoinQueueTime
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetLeaveQueueTime() *int64 {
	return s.LeaveQueueTime
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetMark() *int64 {
	return s.Mark
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetMarkData() *string {
	return s.MarkData
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetQname() *string {
	return s.Qname
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetQno() *string {
	return s.Qno
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetQueueAnswerInTime() *int64 {
	return s.QueueAnswerInTime
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetRecordFile() *string {
	return s.RecordFile
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetSayVoiceDuration() *int64 {
	return s.SayVoiceDuration
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetSipCause() *string {
	return s.SipCause
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetStatus() *string {
	return s.Status
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetTagNames() []*string {
	return s.TagNames
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetTags() []*string {
	return s.Tags
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) GetUniqueId() *string {
	return s.UniqueId
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetBridgeDuration(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.BridgeDuration = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetBridgeTime(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.BridgeTime = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetCallId(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.CallId = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetCallType(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.CallType = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetClientName(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.ClientName = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetClientNumber(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.ClientNumber = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetClientOffhookTime(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.ClientOffhookTime = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetClientRingingTime(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.ClientRingingTime = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetCno(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.Cno = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetCustomerCity(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.CustomerCity = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetCustomerNumber(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.CustomerNumber = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetCustomerNumberEncrypt(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.CustomerNumberEncrypt = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetCustomerProvince(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.CustomerProvince = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetEndReason(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.EndReason = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetEndTime(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.EndTime = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetEvaluation(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.Evaluation = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetFirstJoinQueueTime(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.FirstJoinQueueTime = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetHotline(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.Hotline = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetHotlineName(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.HotlineName = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetIbRingingDuration(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.IbRingingDuration = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetIbWaitDuration(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.IbWaitDuration = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetIvrName(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.IvrName = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetJoinQueueTime(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.JoinQueueTime = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetLeaveQueueTime(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.LeaveQueueTime = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetMark(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.Mark = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetMarkData(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.MarkData = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetQname(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.Qname = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetQno(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.Qno = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetQueueAnswerInTime(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.QueueAnswerInTime = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetRecordFile(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.RecordFile = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetSayVoiceDuration(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.SayVoiceDuration = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetSipCause(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.SipCause = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetStartTime(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.StartTime = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetStatus(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.Status = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetStatusCode(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.StatusCode = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetTagNames(v []*string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.TagNames = v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetTags(v []*string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.Tags = v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetTotalDuration(v int64) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.TotalDuration = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) SetUniqueId(v string) *ClinkListCdrIbsResponseBodyDataCdrIbs {
	s.UniqueId = &v
	return s
}

func (s *ClinkListCdrIbsResponseBodyDataCdrIbs) Validate() error {
	return dara.Validate(s)
}
