// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetObCdrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudGetObCdrResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudGetObCdrResponseBody
	GetCode() *string
	SetData(v *CloudGetObCdrResponseBodyData) *CloudGetObCdrResponseBody
	GetData() *CloudGetObCdrResponseBodyData
	SetMessage(v string) *CloudGetObCdrResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudGetObCdrResponseBody
	GetRequestId() *string
}

type CloudGetObCdrResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudGetObCdrResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9FF70B74-1B3C-44C1-ACDF-8DF962988F0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudGetObCdrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudGetObCdrResponseBody) GoString() string {
	return s.String()
}

func (s *CloudGetObCdrResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudGetObCdrResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudGetObCdrResponseBody) GetData() *CloudGetObCdrResponseBodyData {
	return s.Data
}

func (s *CloudGetObCdrResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudGetObCdrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudGetObCdrResponseBody) SetAccessDeniedDetail(v string) *CloudGetObCdrResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudGetObCdrResponseBody) SetCode(v string) *CloudGetObCdrResponseBody {
	s.Code = &v
	return s
}

func (s *CloudGetObCdrResponseBody) SetData(v *CloudGetObCdrResponseBodyData) *CloudGetObCdrResponseBody {
	s.Data = v
	return s
}

func (s *CloudGetObCdrResponseBody) SetMessage(v string) *CloudGetObCdrResponseBody {
	s.Message = &v
	return s
}

func (s *CloudGetObCdrResponseBody) SetRequestId(v string) *CloudGetObCdrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudGetObCdrResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudGetObCdrResponseBodyData struct {
	// 座席外呼通话记录详情
	List []*CloudGetObCdrResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudGetObCdrResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudGetObCdrResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudGetObCdrResponseBodyData) GetList() []*CloudGetObCdrResponseBodyDataList {
	return s.List
}

func (s *CloudGetObCdrResponseBodyData) SetList(v []*CloudGetObCdrResponseBodyDataList) *CloudGetObCdrResponseBodyData {
	s.List = v
	return s
}

func (s *CloudGetObCdrResponseBodyData) Validate() error {
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

type CloudGetObCdrResponseBodyDataList struct {
	// example:
	//
	// name1
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// 客户接听时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1775024775
	AnswerTime *string `json:"AnswerTime,omitempty" xml:"AnswerTime,omitempty"`
	// 通话唯一标识
	//
	// example:
	//
	// uniq_1-162046xxxx.12
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 呼叫类型 转移
	//
	// example:
	//
	// 示例值示例值
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 呼叫类型编码，如102
	//
	// example:
	//
	// 102
	CallTypeCode *string `json:"CallTypeCode,omitempty" xml:"CallTypeCode,omitempty"`
	// 被叫号码区号
	//
	// example:
	//
	// 010
	CalleeAreaCode *string `json:"CalleeAreaCode,omitempty" xml:"CalleeAreaCode,omitempty"`
	// 呼叫的号码，可能是座席也可能是普通电话
	//
	// example:
	//
	// 01041003090
	CalleeNumber *string `json:"CalleeNumber,omitempty" xml:"CalleeNumber,omitempty"`
	// 外显号码
	//
	// example:
	//
	// 02138276106
	Clid *string `json:"Clid,omitempty" xml:"Clid,omitempty"`
	// 座席工号，如2000
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 客户侧真实外显号码，当使用AXB场景进行外呼时，客户侧真实外显号码为虚拟号
	//
	// example:
	//
	// 02138276106
	DisplayNumber *string `json:"DisplayNumber,omitempty" xml:"DisplayNumber,omitempty"`
	// 结束原因,接听之后的挂机原因 1000:主通道挂机 1001:非主通道挂机 1002:被强拆
	//
	// example:
	//
	// 1000
	EndReason *string `json:"EndReason,omitempty" xml:"EndReason,omitempty"`
	// 电话结束时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1775024775
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// WH514097
	Gno *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	// Id
	//
	// example:
	//
	// 797880898
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 入队列时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1775024775
	JoinQueueTime *string `json:"JoinQueueTime,omitempty" xml:"JoinQueueTime,omitempty"`
	// 主通话记录来电类型 主叫外呼
	//
	// example:
	//
	// 示例值示例值
	MainCallType *string `json:"MainCallType,omitempty" xml:"MainCallType,omitempty"`
	// 主外呼通话记录唯一标识
	//
	// example:
	//
	// uniq_1-162046xxxx.12
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	// 队列号
	//
	// example:
	//
	// 131
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// 进入队列时间
	//
	// example:
	//
	// 1775024775
	QueueName  *string   `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RecordFile []*string `json:"RecordFile,omitempty" xml:"RecordFile,omitempty" type:"Repeated"`
	// 响铃时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1775024775
	RingingTime *string `json:"RingingTime,omitempty" xml:"RingingTime,omitempty"`
	// SIP响应码：200 ~ 699
	//
	// example:
	//
	// 200
	SipCause *string `json:"SipCause,omitempty" xml:"SipCause,omitempty"`
	// 呼叫结果
	//
	// example:
	//
	// 示例值示例值示例值
	SipCauseStr *string `json:"SipCauseStr,omitempty" xml:"SipCauseStr,omitempty"`
	// 呼叫客户时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1775024775
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 外呼结果，如接听
	//
	// example:
	//
	// 示例值示例值
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// example:
	//
	// 示例值
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 总通话时长秒数，单位是s
	//
	// example:
	//
	// 50
	TotalDuration *string `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// 中继群组代号
	//
	// example:
	//
	// 30019
	TrunkGroupKey *string `json:"TrunkGroupKey,omitempty" xml:"TrunkGroupKey,omitempty"`
	// 号码状态识别录音文件名
	//
	// example:
	//
	// name5
	TsiFile *string `json:"TsiFile,omitempty" xml:"TsiFile,omitempty"`
	// 从通道唯一标识
	//
	// example:
	//
	// uniq_1-162046xxxx.12
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// example:
	//
	// 示例值
	XNumber *string `json:"XNumber,omitempty" xml:"XNumber,omitempty"`
}

func (s CloudGetObCdrResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudGetObCdrResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudGetObCdrResponseBodyDataList) GetAgentName() *string {
	return s.AgentName
}

func (s *CloudGetObCdrResponseBodyDataList) GetAnswerTime() *string {
	return s.AnswerTime
}

func (s *CloudGetObCdrResponseBodyDataList) GetCallId() *string {
	return s.CallId
}

func (s *CloudGetObCdrResponseBodyDataList) GetCallType() *string {
	return s.CallType
}

func (s *CloudGetObCdrResponseBodyDataList) GetCallTypeCode() *string {
	return s.CallTypeCode
}

func (s *CloudGetObCdrResponseBodyDataList) GetCalleeAreaCode() *string {
	return s.CalleeAreaCode
}

func (s *CloudGetObCdrResponseBodyDataList) GetCalleeNumber() *string {
	return s.CalleeNumber
}

func (s *CloudGetObCdrResponseBodyDataList) GetClid() *string {
	return s.Clid
}

func (s *CloudGetObCdrResponseBodyDataList) GetCno() *string {
	return s.Cno
}

func (s *CloudGetObCdrResponseBodyDataList) GetDisplayNumber() *string {
	return s.DisplayNumber
}

func (s *CloudGetObCdrResponseBodyDataList) GetEndReason() *string {
	return s.EndReason
}

func (s *CloudGetObCdrResponseBodyDataList) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudGetObCdrResponseBodyDataList) GetGno() *string {
	return s.Gno
}

func (s *CloudGetObCdrResponseBodyDataList) GetId() *string {
	return s.Id
}

func (s *CloudGetObCdrResponseBodyDataList) GetJoinQueueTime() *string {
	return s.JoinQueueTime
}

func (s *CloudGetObCdrResponseBodyDataList) GetMainCallType() *string {
	return s.MainCallType
}

func (s *CloudGetObCdrResponseBodyDataList) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *CloudGetObCdrResponseBodyDataList) GetQno() *string {
	return s.Qno
}

func (s *CloudGetObCdrResponseBodyDataList) GetQueueName() *string {
	return s.QueueName
}

func (s *CloudGetObCdrResponseBodyDataList) GetRecordFile() []*string {
	return s.RecordFile
}

func (s *CloudGetObCdrResponseBodyDataList) GetRingingTime() *string {
	return s.RingingTime
}

func (s *CloudGetObCdrResponseBodyDataList) GetSipCause() *string {
	return s.SipCause
}

func (s *CloudGetObCdrResponseBodyDataList) GetSipCauseStr() *string {
	return s.SipCauseStr
}

func (s *CloudGetObCdrResponseBodyDataList) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudGetObCdrResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *CloudGetObCdrResponseBodyDataList) GetStatusCode() *string {
	return s.StatusCode
}

func (s *CloudGetObCdrResponseBodyDataList) GetTaskName() *string {
	return s.TaskName
}

func (s *CloudGetObCdrResponseBodyDataList) GetTotalDuration() *string {
	return s.TotalDuration
}

func (s *CloudGetObCdrResponseBodyDataList) GetTrunkGroupKey() *string {
	return s.TrunkGroupKey
}

func (s *CloudGetObCdrResponseBodyDataList) GetTsiFile() *string {
	return s.TsiFile
}

func (s *CloudGetObCdrResponseBodyDataList) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudGetObCdrResponseBodyDataList) GetXNumber() *string {
	return s.XNumber
}

func (s *CloudGetObCdrResponseBodyDataList) SetAgentName(v string) *CloudGetObCdrResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetAnswerTime(v string) *CloudGetObCdrResponseBodyDataList {
	s.AnswerTime = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetCallId(v string) *CloudGetObCdrResponseBodyDataList {
	s.CallId = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetCallType(v string) *CloudGetObCdrResponseBodyDataList {
	s.CallType = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetCallTypeCode(v string) *CloudGetObCdrResponseBodyDataList {
	s.CallTypeCode = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetCalleeAreaCode(v string) *CloudGetObCdrResponseBodyDataList {
	s.CalleeAreaCode = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetCalleeNumber(v string) *CloudGetObCdrResponseBodyDataList {
	s.CalleeNumber = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetClid(v string) *CloudGetObCdrResponseBodyDataList {
	s.Clid = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetCno(v string) *CloudGetObCdrResponseBodyDataList {
	s.Cno = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetDisplayNumber(v string) *CloudGetObCdrResponseBodyDataList {
	s.DisplayNumber = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetEndReason(v string) *CloudGetObCdrResponseBodyDataList {
	s.EndReason = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetEndTime(v string) *CloudGetObCdrResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetGno(v string) *CloudGetObCdrResponseBodyDataList {
	s.Gno = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetId(v string) *CloudGetObCdrResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetJoinQueueTime(v string) *CloudGetObCdrResponseBodyDataList {
	s.JoinQueueTime = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetMainCallType(v string) *CloudGetObCdrResponseBodyDataList {
	s.MainCallType = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetMainUniqueId(v string) *CloudGetObCdrResponseBodyDataList {
	s.MainUniqueId = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetQno(v string) *CloudGetObCdrResponseBodyDataList {
	s.Qno = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetQueueName(v string) *CloudGetObCdrResponseBodyDataList {
	s.QueueName = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetRecordFile(v []*string) *CloudGetObCdrResponseBodyDataList {
	s.RecordFile = v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetRingingTime(v string) *CloudGetObCdrResponseBodyDataList {
	s.RingingTime = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetSipCause(v string) *CloudGetObCdrResponseBodyDataList {
	s.SipCause = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetSipCauseStr(v string) *CloudGetObCdrResponseBodyDataList {
	s.SipCauseStr = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetStartTime(v string) *CloudGetObCdrResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetStatus(v string) *CloudGetObCdrResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetStatusCode(v string) *CloudGetObCdrResponseBodyDataList {
	s.StatusCode = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetTaskName(v string) *CloudGetObCdrResponseBodyDataList {
	s.TaskName = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetTotalDuration(v string) *CloudGetObCdrResponseBodyDataList {
	s.TotalDuration = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetTrunkGroupKey(v string) *CloudGetObCdrResponseBodyDataList {
	s.TrunkGroupKey = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetTsiFile(v string) *CloudGetObCdrResponseBodyDataList {
	s.TsiFile = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetUniqueId(v string) *CloudGetObCdrResponseBodyDataList {
	s.UniqueId = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) SetXNumber(v string) *CloudGetObCdrResponseBodyDataList {
	s.XNumber = &v
	return s
}

func (s *CloudGetObCdrResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
