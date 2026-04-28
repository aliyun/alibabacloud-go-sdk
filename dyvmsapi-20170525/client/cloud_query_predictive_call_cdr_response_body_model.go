// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryPredictiveCallCdrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudQueryPredictiveCallCdrResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudQueryPredictiveCallCdrResponseBody
	GetCode() *string
	SetData(v *CloudQueryPredictiveCallCdrResponseBodyData) *CloudQueryPredictiveCallCdrResponseBody
	GetData() *CloudQueryPredictiveCallCdrResponseBodyData
	SetMessage(v string) *CloudQueryPredictiveCallCdrResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudQueryPredictiveCallCdrResponseBody
	GetRequestId() *string
}

type CloudQueryPredictiveCallCdrResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudQueryPredictiveCallCdrResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudQueryPredictiveCallCdrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryPredictiveCallCdrResponseBody) GoString() string {
	return s.String()
}

func (s *CloudQueryPredictiveCallCdrResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudQueryPredictiveCallCdrResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudQueryPredictiveCallCdrResponseBody) GetData() *CloudQueryPredictiveCallCdrResponseBodyData {
	return s.Data
}

func (s *CloudQueryPredictiveCallCdrResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudQueryPredictiveCallCdrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudQueryPredictiveCallCdrResponseBody) SetAccessDeniedDetail(v string) *CloudQueryPredictiveCallCdrResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBody) SetCode(v string) *CloudQueryPredictiveCallCdrResponseBody {
	s.Code = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBody) SetData(v *CloudQueryPredictiveCallCdrResponseBodyData) *CloudQueryPredictiveCallCdrResponseBody {
	s.Data = v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBody) SetMessage(v string) *CloudQueryPredictiveCallCdrResponseBody {
	s.Message = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBody) SetRequestId(v string) *CloudQueryPredictiveCallCdrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudQueryPredictiveCallCdrResponseBodyData struct {
	// 预测式外呼通话记录数组
	List []*CloudQueryPredictiveCallCdrResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudQueryPredictiveCallCdrResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryPredictiveCallCdrResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudQueryPredictiveCallCdrResponseBodyData) GetList() []*CloudQueryPredictiveCallCdrResponseBodyDataList {
	return s.List
}

func (s *CloudQueryPredictiveCallCdrResponseBodyData) SetList(v []*CloudQueryPredictiveCallCdrResponseBodyDataList) *CloudQueryPredictiveCallCdrResponseBodyData {
	s.List = v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyData) Validate() error {
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

type CloudQueryPredictiveCallCdrResponseBodyDataList struct {
	// 座席姓名, 如\\"张三1010\\"
	//
	// example:
	//
	// 示例值
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// 系统接听时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1775024775
	AnswerTime *string `json:"AnswerTime,omitempty" xml:"AnswerTime,omitempty"`
	// 服务处理时长秒数，单位是s
	//
	// example:
	//
	// 50
	BridgeDuration *string `json:"BridgeDuration,omitempty" xml:"BridgeDuration,omitempty"`
	// 座席接听时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1774822305
	BridgeTime *string `json:"BridgeTime,omitempty" xml:"BridgeTime,omitempty"`
	// 通话唯一标识
	//
	// example:
	//
	// 3cdf8d08-0216-42df-bba5-0fd45c93c9b3
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 呼叫类型 webcall
	//
	// example:
	//
	// 示例值
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 第二侧外显号码
	//
	// example:
	//
	// 15553333
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
	// 1775024775
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
	// 1111
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
	// 话机后话机响铃时间，时间戳，精确到s；无或未转接时可能为 0；如1480904471
	//
	// example:
	//
	// 1775024775
	DetailRingingTime *string `json:"DetailRingingTime,omitempty" xml:"DetailRingingTime,omitempty"`
	// 客户侧真实外显号码，当使用AXB场景进行外呼时，客户侧的真实外显号码为虚拟号
	//
	// example:
	//
	// 1775024
	DisplayNumber *string `json:"DisplayNumber,omitempty" xml:"DisplayNumber,omitempty"`
	// 是否呼叫降级, 0-否, 1-是
	//
	// example:
	//
	// 0
	DownGrade *string `json:"DownGrade,omitempty" xml:"DownGrade,omitempty"`
	// 降级次数
	//
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
	// 1775024775
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 外呼组号：座席所属外呼组编号，如2000
	//
	// example:
	//
	// 2000
	Gno *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	// 通话记录主键标识；
	//
	// example:
	//
	// 123546546
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
	// IvrName
	IvrName *string `json:"IvrName,omitempty" xml:"IvrName,omitempty"`
	// 中继号码
	//
	// example:
	//
	// 15322355
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
	// 1774904471
	RealAnswerTime *string `json:"RealAnswerTime,omitempty" xml:"RealAnswerTime,omitempty"`
	// 通话记录录音数组，json格式
	RecordFile []*CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile `json:"RecordFile,omitempty" xml:"RecordFile,omitempty" type:"Repeated"`
	// 用户自定义通话标识字段
	//
	// example:
	//
	// 543252344
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
	// 示例值示例值示例值
	SipCauseStr *string `json:"SipCauseStr,omitempty" xml:"SipCauseStr,omitempty"`
	// 电话进入系统时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1774822300
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 外呼结果， 如客户接听
	//
	// example:
	//
	// 示例值
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 外呼结果状态码：40 客户未接听；41 客户接听；42 已呼叫；43 双方接听；详见接口文档请求参数中 status 说明
	//
	// example:
	//
	// 41
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// 任务文件id
	//
	// example:
	//
	// 111
	TaskFileId *string `json:"TaskFileId,omitempty" xml:"TaskFileId,omitempty"`
	// 任务id
	//
	// example:
	//
	// 1234
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 任务名称
	//
	// example:
	//
	// name2
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 总通话时长秒数，单位是s
	//
	// example:
	//
	// 55
	TotalDuration *string `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// 中继群组代号
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
	// 通话记录自定义字段，json格式
	//
	// example:
	//
	// {}
	UserField map[string]interface{} `json:"UserField,omitempty" xml:"UserField,omitempty"`
	// 通话质量百分比，0–100，in/out 方向说话占用百分比
	//
	// example:
	//
	// 23
	VadIn *string `json:"VadIn,omitempty" xml:"VadIn,omitempty"`
	// 通话质量百分比，0–100，in/out 方向说话占用百分比
	//
	// example:
	//
	// 77
	VadOut *string `json:"VadOut,omitempty" xml:"VadOut,omitempty"`
	// 客户等待时长秒数，单位是s
	//
	// example:
	//
	// 10
	WaitDuration *string `json:"WaitDuration,omitempty" xml:"WaitDuration,omitempty"`
	// 虚拟号
	//
	// example:
	//
	// 6677333
	XNumber *string `json:"XNumber,omitempty" xml:"XNumber,omitempty"`
}

func (s CloudQueryPredictiveCallCdrResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryPredictiveCallCdrResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetAgentName() *string {
	return s.AgentName
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetAnswerTime() *string {
	return s.AnswerTime
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetBridgeDuration() *string {
	return s.BridgeDuration
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetBridgeTime() *string {
	return s.BridgeTime
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetCallId() *string {
	return s.CallId
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetCallType() *string {
	return s.CallType
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetCalleeClid() *string {
	return s.CalleeClid
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetCalleeDisplayNumber() *string {
	return s.CalleeDisplayNumber
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetCalleeNumber() *string {
	return s.CalleeNumber
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetCalleeRingingTime() *string {
	return s.CalleeRingingTime
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetClid() *string {
	return s.Clid
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetCno() *string {
	return s.Cno
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetCustomerCity() *string {
	return s.CustomerCity
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetCustomerProvince() *string {
	return s.CustomerProvince
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetDetailRingingTime() *string {
	return s.DetailRingingTime
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetDisplayNumber() *string {
	return s.DisplayNumber
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetDownGrade() *string {
	return s.DownGrade
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetDownGradeCounts() *string {
	return s.DownGradeCounts
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetEndReason() *string {
	return s.EndReason
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetGno() *string {
	return s.Gno
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetId() *string {
	return s.Id
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetIsRealAnswer() *string {
	return s.IsRealAnswer
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetIvrName() *string {
	return s.IvrName
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetNumberTrunk() *string {
	return s.NumberTrunk
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetQno() *string {
	return s.Qno
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetRealAnswerTime() *string {
	return s.RealAnswerTime
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetRecordFile() []*CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile {
	return s.RecordFile
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetRequestUniqueId() *string {
	return s.RequestUniqueId
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetSipCause() *string {
	return s.SipCause
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetSipCauseStr() *string {
	return s.SipCauseStr
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetStatusCode() *string {
	return s.StatusCode
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetTaskFileId() *string {
	return s.TaskFileId
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetTaskId() *string {
	return s.TaskId
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetTaskName() *string {
	return s.TaskName
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetTotalDuration() *string {
	return s.TotalDuration
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetTrunkGroupKey() *string {
	return s.TrunkGroupKey
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetUserField() map[string]interface{} {
	return s.UserField
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetVadIn() *string {
	return s.VadIn
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetVadOut() *string {
	return s.VadOut
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetWaitDuration() *string {
	return s.WaitDuration
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) GetXNumber() *string {
	return s.XNumber
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetAgentName(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetAnswerTime(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.AnswerTime = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetBridgeDuration(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.BridgeDuration = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetBridgeTime(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.BridgeTime = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetCallId(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.CallId = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetCallType(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.CallType = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetCalleeClid(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.CalleeClid = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetCalleeDisplayNumber(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.CalleeDisplayNumber = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetCalleeNumber(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.CalleeNumber = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetCalleeRingingTime(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.CalleeRingingTime = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetClid(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.Clid = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetCno(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.Cno = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetCustomerCity(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.CustomerCity = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetCustomerNumber(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.CustomerNumber = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetCustomerProvince(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.CustomerProvince = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetDetailRingingTime(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.DetailRingingTime = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetDisplayNumber(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.DisplayNumber = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetDownGrade(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.DownGrade = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetDownGradeCounts(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.DownGradeCounts = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetEndReason(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.EndReason = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetEndTime(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetGno(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.Gno = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetId(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetIsRealAnswer(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.IsRealAnswer = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetIvrName(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.IvrName = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetNumberTrunk(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.NumberTrunk = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetQno(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.Qno = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetRealAnswerTime(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.RealAnswerTime = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetRecordFile(v []*CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.RecordFile = v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetRequestUniqueId(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.RequestUniqueId = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetSipCause(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.SipCause = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetSipCauseStr(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.SipCauseStr = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetStartTime(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetStatus(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetStatusCode(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.StatusCode = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetTaskFileId(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.TaskFileId = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetTaskId(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.TaskId = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetTaskName(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.TaskName = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetTotalDuration(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.TotalDuration = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetTrunkGroupKey(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.TrunkGroupKey = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetUniqueId(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.UniqueId = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetUserField(v map[string]interface{}) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.UserField = v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetVadIn(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.VadIn = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetVadOut(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.VadOut = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetWaitDuration(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.WaitDuration = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) SetXNumber(v string) *CloudQueryPredictiveCallCdrResponseBodyDataList {
	s.XNumber = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataList) Validate() error {
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

type CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile struct {
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
	// 取值：record,voicemail 说明：record是录音， voicemail是留言
	//
	// example:
	//
	// record
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile) GoString() string {
	return s.String()
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile) GetFile() *string {
	return s.File
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile) GetMonitorType() *int64 {
	return s.MonitorType
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile) GetType() *string {
	return s.Type
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile) SetFile(v string) *CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile {
	s.File = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile) SetMonitorType(v int64) *CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile {
	s.MonitorType = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile) SetType(v string) *CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile {
	s.Type = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponseBodyDataListRecordFile) Validate() error {
	return dara.Validate(s)
}
