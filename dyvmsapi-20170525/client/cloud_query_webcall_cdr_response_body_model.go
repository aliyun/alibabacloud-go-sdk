// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryWebcallCdrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudQueryWebcallCdrResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudQueryWebcallCdrResponseBody
	GetCode() *string
	SetData(v *CloudQueryWebcallCdrResponseBodyData) *CloudQueryWebcallCdrResponseBody
	GetData() *CloudQueryWebcallCdrResponseBodyData
	SetMessage(v string) *CloudQueryWebcallCdrResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudQueryWebcallCdrResponseBody
	GetRequestId() *string
}

type CloudQueryWebcallCdrResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudQueryWebcallCdrResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudQueryWebcallCdrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryWebcallCdrResponseBody) GoString() string {
	return s.String()
}

func (s *CloudQueryWebcallCdrResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudQueryWebcallCdrResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudQueryWebcallCdrResponseBody) GetData() *CloudQueryWebcallCdrResponseBodyData {
	return s.Data
}

func (s *CloudQueryWebcallCdrResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudQueryWebcallCdrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudQueryWebcallCdrResponseBody) SetAccessDeniedDetail(v string) *CloudQueryWebcallCdrResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBody) SetCode(v string) *CloudQueryWebcallCdrResponseBody {
	s.Code = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBody) SetData(v *CloudQueryWebcallCdrResponseBodyData) *CloudQueryWebcallCdrResponseBody {
	s.Data = v
	return s
}

func (s *CloudQueryWebcallCdrResponseBody) SetMessage(v string) *CloudQueryWebcallCdrResponseBody {
	s.Message = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBody) SetRequestId(v string) *CloudQueryWebcallCdrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudQueryWebcallCdrResponseBodyData struct {
	// webcall通话记录数组
	List []*CloudQueryWebcallCdrResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudQueryWebcallCdrResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryWebcallCdrResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudQueryWebcallCdrResponseBodyData) GetList() []*CloudQueryWebcallCdrResponseBodyDataList {
	return s.List
}

func (s *CloudQueryWebcallCdrResponseBodyData) SetList(v []*CloudQueryWebcallCdrResponseBodyDataList) *CloudQueryWebcallCdrResponseBodyData {
	s.List = v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyData) Validate() error {
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

type CloudQueryWebcallCdrResponseBodyDataList struct {
	// example:
	//
	// 示例值示例值
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// 系统接听时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1774822302
	AnswerTime *string `json:"AnswerTime,omitempty" xml:"AnswerTime,omitempty"`
	// 服务处理时长秒数，单位是s
	//
	// example:
	//
	// 4
	BridgeDuration *string `json:"BridgeDuration,omitempty" xml:"BridgeDuration,omitempty"`
	// 座席接听时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1774822308
	BridgeTime *string `json:"BridgeTime,omitempty" xml:"BridgeTime,omitempty"`
	// 通话唯一标识
	//
	// example:
	//
	// 4e18e73a-af6a-4a77-92ea-628a6effd64b
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 呼叫类型 webcall
	//
	// example:
	//
	// webcall
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 第二侧外显号码
	//
	// example:
	//
	// 02138276106
	CalleeClid *string `json:"CalleeClid,omitempty" xml:"CalleeClid,omitempty"`
	// 第二侧真实外显号码
	//
	// example:
	//
	// 02138276106
	CalleeDisplayNumber *string `json:"CalleeDisplayNumber,omitempty" xml:"CalleeDisplayNumber,omitempty"`
	// 座席电话 区号 +7或8位的固话号码，区号与固话号码之间无\\"-\\"；或11位长度的手机号码。0104100596
	//
	// example:
	//
	// 0104100596
	CalleeNumber *string `json:"CalleeNumber,omitempty" xml:"CalleeNumber,omitempty"`
	// 客户响铃时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1774822302
	CalleeRingingTime *string `json:"CalleeRingingTime,omitempty" xml:"CalleeRingingTime,omitempty"`
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
	// 2000
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 客户号码归属城市,北京
	//
	// example:
	//
	// 示例值示例值示例值
	CustomerCity *string `json:"CustomerCity,omitempty" xml:"CustomerCity,omitempty"`
	// 客户号码 国内固话或手机，区号 + 7或8位的固话号码，区号与固话号码之间无\\"-\\"；或11位长度的手机号码。如 01041005968或18701051984 或虚拟号-分机号场景格式，如 18701051984-7538
	//
	// example:
	//
	// 18701051984
	CustomerNumber *string `json:"CustomerNumber,omitempty" xml:"CustomerNumber,omitempty"`
	// 客户号码归属省份，如 北京
	//
	// example:
	//
	// 示例值示例值示例值
	CustomerProvince *string `json:"CustomerProvince,omitempty" xml:"CustomerProvince,omitempty"`
	// 转话机后,话机响铃时间
	//
	// example:
	//
	// 1774822303
	DetailRingingTime *string `json:"DetailRingingTime,omitempty" xml:"DetailRingingTime,omitempty"`
	// 真实外显号码，当使用AXB场景进行外呼时，真实外显号码为虚拟号
	//
	// example:
	//
	// 02138276106
	DisplayNumber *string `json:"DisplayNumber,omitempty" xml:"DisplayNumber,omitempty"`
	// example:
	//
	// 0
	DownGrade *string `json:"DownGrade,omitempty" xml:"DownGrade,omitempty"`
	// example:
	//
	// 0
	DownGradeCounts *string `json:"DownGradeCounts,omitempty" xml:"DownGradeCounts,omitempty"`
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
	// 1774822312
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// WH333
	Gno *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	// Id
	//
	// example:
	//
	// 14454
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 是否真人接听：1 - 是、0 - 否、空值
	//
	// example:
	//
	// 1
	IsRealAnswer *string `json:"IsRealAnswer,omitempty" xml:"IsRealAnswer,omitempty"`
	// IVR名称，默认自定义
	//
	// example:
	//
	// webcall
	IvrName *string `json:"IvrName,omitempty" xml:"IvrName,omitempty"`
	// 中继号码
	//
	// example:
	//
	// 42003255
	NumberTrunk *string `json:"NumberTrunk,omitempty" xml:"NumberTrunk,omitempty"`
	// 队列号，如1000
	//
	// example:
	//
	// 1000
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// 真人接听时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1775024775
	RealAnswerTime *string `json:"RealAnswerTime,omitempty" xml:"RealAnswerTime,omitempty"`
	// 通话记录录音数组
	RecordFile []*CloudQueryWebcallCdrResponseBodyDataListRecordFile `json:"RecordFile,omitempty" xml:"RecordFile,omitempty" type:"Repeated"`
	// 用户自定义通话标识字段
	//
	// example:
	//
	// 214b76605a43d02f4aab4e61cfa63128
	RequestUniqueId *string `json:"RequestUniqueId,omitempty" xml:"RequestUniqueId,omitempty"`
	// 号码状态识别编码，详见
	//
	// example:
	//
	// 200
	SipCause *string `json:"SipCause,omitempty" xml:"SipCause,omitempty"`
	// 呼叫结果
	//
	// example:
	//
	// 示例值示例值
	SipCauseStr *string `json:"SipCauseStr,omitempty" xml:"SipCauseStr,omitempty"`
	// 电话进入系统时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1774822302
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 外呼结果， 如客户接听
	//
	// example:
	//
	// 示例值示例值
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// status对应的状态码
	//
	// example:
	//
	// 24
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// 任务文件id
	//
	// example:
	//
	// 12144
	TaskFileId *string `json:"TaskFileId,omitempty" xml:"TaskFileId,omitempty"`
	// 任务id
	//
	// example:
	//
	// 13354
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 总通话时长秒数，单位是s
	//
	// example:
	//
	// 10
	TotalDuration *string `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// 中继组标识
	//
	// example:
	//
	// 30019
	TrunkGroupKey *string `json:"TrunkGroupKey,omitempty" xml:"TrunkGroupKey,omitempty"`
	// 电话唯一标识
	//
	// example:
	//
	// uniq_1-162046xxxx.12
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// example:
	//
	// 6
	WaitDuration *string `json:"WaitDuration,omitempty" xml:"WaitDuration,omitempty"`
}

func (s CloudQueryWebcallCdrResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryWebcallCdrResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetAgentName() *string {
	return s.AgentName
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetAnswerTime() *string {
	return s.AnswerTime
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetBridgeDuration() *string {
	return s.BridgeDuration
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetBridgeTime() *string {
	return s.BridgeTime
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetCallId() *string {
	return s.CallId
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetCallType() *string {
	return s.CallType
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetCalleeClid() *string {
	return s.CalleeClid
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetCalleeDisplayNumber() *string {
	return s.CalleeDisplayNumber
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetCalleeNumber() *string {
	return s.CalleeNumber
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetCalleeRingingTime() *string {
	return s.CalleeRingingTime
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetClid() *string {
	return s.Clid
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetCno() *string {
	return s.Cno
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetCustomerCity() *string {
	return s.CustomerCity
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetCustomerProvince() *string {
	return s.CustomerProvince
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetDetailRingingTime() *string {
	return s.DetailRingingTime
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetDisplayNumber() *string {
	return s.DisplayNumber
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetDownGrade() *string {
	return s.DownGrade
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetDownGradeCounts() *string {
	return s.DownGradeCounts
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetEndReason() *string {
	return s.EndReason
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetGno() *string {
	return s.Gno
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetId() *string {
	return s.Id
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetIsRealAnswer() *string {
	return s.IsRealAnswer
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetIvrName() *string {
	return s.IvrName
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetNumberTrunk() *string {
	return s.NumberTrunk
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetQno() *string {
	return s.Qno
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetRealAnswerTime() *string {
	return s.RealAnswerTime
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetRecordFile() []*CloudQueryWebcallCdrResponseBodyDataListRecordFile {
	return s.RecordFile
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetRequestUniqueId() *string {
	return s.RequestUniqueId
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetSipCause() *string {
	return s.SipCause
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetSipCauseStr() *string {
	return s.SipCauseStr
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetStatusCode() *string {
	return s.StatusCode
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetTaskFileId() *string {
	return s.TaskFileId
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetTaskId() *string {
	return s.TaskId
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetTotalDuration() *string {
	return s.TotalDuration
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetTrunkGroupKey() *string {
	return s.TrunkGroupKey
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) GetWaitDuration() *string {
	return s.WaitDuration
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetAgentName(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetAnswerTime(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.AnswerTime = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetBridgeDuration(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.BridgeDuration = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetBridgeTime(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.BridgeTime = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetCallId(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.CallId = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetCallType(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.CallType = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetCalleeClid(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.CalleeClid = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetCalleeDisplayNumber(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.CalleeDisplayNumber = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetCalleeNumber(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.CalleeNumber = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetCalleeRingingTime(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.CalleeRingingTime = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetClid(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.Clid = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetCno(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.Cno = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetCustomerCity(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.CustomerCity = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetCustomerNumber(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.CustomerNumber = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetCustomerProvince(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.CustomerProvince = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetDetailRingingTime(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.DetailRingingTime = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetDisplayNumber(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.DisplayNumber = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetDownGrade(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.DownGrade = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetDownGradeCounts(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.DownGradeCounts = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetEndReason(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.EndReason = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetEndTime(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetGno(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.Gno = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetId(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetIsRealAnswer(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.IsRealAnswer = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetIvrName(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.IvrName = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetNumberTrunk(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.NumberTrunk = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetQno(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.Qno = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetRealAnswerTime(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.RealAnswerTime = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetRecordFile(v []*CloudQueryWebcallCdrResponseBodyDataListRecordFile) *CloudQueryWebcallCdrResponseBodyDataList {
	s.RecordFile = v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetRequestUniqueId(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.RequestUniqueId = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetSipCause(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.SipCause = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetSipCauseStr(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.SipCauseStr = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetStartTime(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetStatus(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetStatusCode(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.StatusCode = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetTaskFileId(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.TaskFileId = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetTaskId(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.TaskId = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetTotalDuration(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.TotalDuration = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetTrunkGroupKey(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.TrunkGroupKey = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetUniqueId(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.UniqueId = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) SetWaitDuration(v string) *CloudQueryWebcallCdrResponseBodyDataList {
	s.WaitDuration = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataList) Validate() error {
	if s.RecordFile != nil {
		for _, item := range s.RecordFile {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudQueryWebcallCdrResponseBodyDataListRecordFile struct {
	// 录音文件名
	//
	// example:
	//
	// name3
	File *string `json:"File,omitempty" xml:"File,omitempty"`
	// 1. 混音 2. 分线录音（保留2个单声道WAV） 3.留言 5. 分线录音（保留1个双声道WAV）
	//
	// example:
	//
	// 1
	MonitorType *int64 `json:"Monitor_type,omitempty" xml:"Monitor_type,omitempty"`
	// 取值：record,voicemail说明：record是录音， voicemail是留言
	//
	// example:
	//
	// record
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CloudQueryWebcallCdrResponseBodyDataListRecordFile) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryWebcallCdrResponseBodyDataListRecordFile) GoString() string {
	return s.String()
}

func (s *CloudQueryWebcallCdrResponseBodyDataListRecordFile) GetFile() *string {
	return s.File
}

func (s *CloudQueryWebcallCdrResponseBodyDataListRecordFile) GetMonitorType() *int64 {
	return s.MonitorType
}

func (s *CloudQueryWebcallCdrResponseBodyDataListRecordFile) GetType() *string {
	return s.Type
}

func (s *CloudQueryWebcallCdrResponseBodyDataListRecordFile) SetFile(v string) *CloudQueryWebcallCdrResponseBodyDataListRecordFile {
	s.File = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataListRecordFile) SetMonitorType(v int64) *CloudQueryWebcallCdrResponseBodyDataListRecordFile {
	s.MonitorType = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataListRecordFile) SetType(v string) *CloudQueryWebcallCdrResponseBodyDataListRecordFile {
	s.Type = &v
	return s
}

func (s *CloudQueryWebcallCdrResponseBodyDataListRecordFile) Validate() error {
	return dara.Validate(s)
}
