// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryObCdrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudQueryObCdrResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudQueryObCdrResponseBody
	GetCode() *string
	SetData(v *CloudQueryObCdrResponseBodyData) *CloudQueryObCdrResponseBody
	GetData() *CloudQueryObCdrResponseBodyData
	SetMessage(v string) *CloudQueryObCdrResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudQueryObCdrResponseBody
	GetRequestId() *string
}

type CloudQueryObCdrResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudQueryObCdrResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6086693B-2250-17CE-A41F-06259AB6DB1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudQueryObCdrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryObCdrResponseBody) GoString() string {
	return s.String()
}

func (s *CloudQueryObCdrResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudQueryObCdrResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudQueryObCdrResponseBody) GetData() *CloudQueryObCdrResponseBodyData {
	return s.Data
}

func (s *CloudQueryObCdrResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudQueryObCdrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudQueryObCdrResponseBody) SetAccessDeniedDetail(v string) *CloudQueryObCdrResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudQueryObCdrResponseBody) SetCode(v string) *CloudQueryObCdrResponseBody {
	s.Code = &v
	return s
}

func (s *CloudQueryObCdrResponseBody) SetData(v *CloudQueryObCdrResponseBodyData) *CloudQueryObCdrResponseBody {
	s.Data = v
	return s
}

func (s *CloudQueryObCdrResponseBody) SetMessage(v string) *CloudQueryObCdrResponseBody {
	s.Message = &v
	return s
}

func (s *CloudQueryObCdrResponseBody) SetRequestId(v string) *CloudQueryObCdrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudQueryObCdrResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudQueryObCdrResponseBodyData struct {
	// 座席外呼通话记录数组
	List []*CloudQueryObCdrResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudQueryObCdrResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryObCdrResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudQueryObCdrResponseBodyData) GetList() []*CloudQueryObCdrResponseBodyDataList {
	return s.List
}

func (s *CloudQueryObCdrResponseBodyData) SetList(v []*CloudQueryObCdrResponseBodyDataList) *CloudQueryObCdrResponseBodyData {
	s.List = v
	return s
}

func (s *CloudQueryObCdrResponseBodyData) Validate() error {
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

type CloudQueryObCdrResponseBodyDataList struct {
	// 座席侧外显号码
	//
	// example:
	//
	// 02138276106
	AgentClid *string `json:"AgentClid,omitempty" xml:"AgentClid,omitempty"`
	// 座席名称
	//
	// example:
	//
	// name3
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// 座席电话 区号 +7或8位的固话号码，区号与固话号码之间无\\"-\\"；或11位长度的手机号码。0104100596
	//
	// example:
	//
	// 0104100596
	AgentNumber *string `json:"AgentNumber,omitempty" xml:"AgentNumber,omitempty"`
	// 座席接听时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1775024775
	AnswerTime *string `json:"AnswerTime,omitempty" xml:"AnswerTime,omitempty"`
	// 服务处理时长秒数，单位是s
	//
	// example:
	//
	// 10
	BridgeDuration *string `json:"BridgeDuration,omitempty" xml:"BridgeDuration,omitempty"`
	// 双方接听时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1775024775
	BridgeTime *string `json:"BridgeTime,omitempty" xml:"BridgeTime,omitempty"`
	// 通话唯一标识
	//
	// example:
	//
	// uniq_1-162046xxxx.12
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 呼叫类型：4、预览外呼; 6、主叫外呼
	//
	// example:
	//
	// 4
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 客户侧响铃时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1775024775
	CalleeRingingTime *string `json:"CalleeRingingTime,omitempty" xml:"CalleeRingingTime,omitempty"`
	// 客户侧外显号码，当使用虚拟号进行AXB外呼且开关打开时，实际返回的客户侧外显号码是虚拟号
	//
	// example:
	//
	// 41008502
	Clid *string `json:"Clid,omitempty" xml:"Clid,omitempty"`
	// 座席工号，如2000
	//
	// example:
	//
	// 2000
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 客户号码区号
	//
	// example:
	//
	// 028
	CustomerAreaCode *string `json:"CustomerAreaCode,omitempty" xml:"CustomerAreaCode,omitempty"`
	// 客户号码归属城市,北京
	//
	// example:
	//
	// 示例值示例值
	CustomerCity *string `json:"CustomerCity,omitempty" xml:"CustomerCity,omitempty"`
	// 客户号码 国内固话或手机，区号 + 7或8位的固话号码，区号与固话号码之间无\\"-\\"；或11位长度的手机号码。如 01041005968或18701051984 或虚拟号-分机号场景格式，如 18701051984-7538
	//
	// example:
	//
	// 01041005968
	CustomerNumber *string `json:"CustomerNumber,omitempty" xml:"CustomerNumber,omitempty"`
	// 客户号码归属省份，如 北京
	//
	// example:
	//
	// 示例值示例值
	CustomerProvince *string `json:"CustomerProvince,omitempty" xml:"CustomerProvince,omitempty"`
	// 座席侧真实外显号码，当使用AXB场景进行外呼时，座席侧真实外显号码为虚拟号
	//
	// example:
	//
	// 41008502
	DisplayNumber *string `json:"DisplayNumber,omitempty" xml:"DisplayNumber,omitempty"`
	// 是否呼叫降级, 0-否, 1-是
	//
	// example:
	//
	// 1
	DownGrade *string `json:"DownGrade,omitempty" xml:"DownGrade,omitempty"`
	// 降级次数
	//
	// example:
	//
	// 1
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
	// 1775024875
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 133
	Gno *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	// example:
	//
	// 0
	HangupReason *string `json:"HangupReason,omitempty" xml:"HangupReason,omitempty"`
	// Id
	//
	// example:
	//
	// 9075217226
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 是否真人接听：1 - 是、0 - 否、空值
	//
	// example:
	//
	// 1
	IsRealAnswer *string `json:"IsRealAnswer,omitempty" xml:"IsRealAnswer,omitempty"`
	// 客户侧真实外显号码，当使用AXB场景进行外呼时，客户侧真实外显号码为虚拟号
	//
	// example:
	//
	// 13322225555
	LeftDisplayNumber *string `json:"LeftDisplayNumber,omitempty" xml:"LeftDisplayNumber,omitempty"`
	// 分机外呼，话机响铃时间
	//
	// example:
	//
	// 1775024775
	MainRingingTime *string `json:"MainRingingTime,omitempty" xml:"MainRingingTime,omitempty"`
	// 真人接听时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1775024775
	RealAnswerTime *string `json:"RealAnswerTime,omitempty" xml:"RealAnswerTime,omitempty"`
	// 通话记录录音数组，json格式
	RecordFile []*CloudQueryObCdrResponseBodyDataListRecordFile `json:"RecordFile,omitempty" xml:"RecordFile,omitempty" type:"Repeated"`
	// 请求唯一id
	//
	// example:
	//
	// uniq_1-162046xxxx.12
	RequestUniqueId *string `json:"RequestUniqueId,omitempty" xml:"RequestUniqueId,omitempty"`
	// RTC总通话时长秒数，单位是s，在呼叫类型为主叫外呼(RTC)时有效，计算规则：endTime-startTime
	//
	// example:
	//
	// 133
	RtcTotalDuration *string `json:"RtcTotalDuration,omitempty" xml:"RtcTotalDuration,omitempty"`
	// SIP响应码：200 ~ 699 详见 注：主叫外呼RTC座席侧没有SIP协商的概念，并未发起呼叫，因此该值默认为0
	//
	// example:
	//
	// 200
	SipCause *string `json:"SipCause,omitempty" xml:"SipCause,omitempty"`
	// 座席开始外呼时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1775024775
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 外呼结果：30 座席未接听; 31 座席接听,未呼叫客户; 32 座席接听,客户未接听; 33 双方接听; 50 主叫外呼接听; 51 主叫外呼,客户未接听; 52 主叫外呼,双方接听。
	//
	// example:
	//
	// 33
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 33
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// 通话状态原因（主叫外呼）描述：1 - 企业黑名单; 2 - 风控系统拦截; 20001 - 呼叫限制；20002 - 黑名单或非白名单；20003 - 禁拨时间段； 3 - 无效的外显号码; 4 - 错误的语音导航; 5 - 企业停机; 6 - 无权限外呼; 7 - 超出呼叫次数限制; 8 - 错误号码; 9 - 其他错误；
	//
	// example:
	//
	// 9
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// 座席自动外呼任务id
	//
	// example:
	//
	// 333
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 总通话时长秒数，单位是s
	//
	// example:
	//
	// 30
	TotalDuration *string `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// 中继群组代号
	//
	// example:
	//
	// 30020
	TrunkGroupKey *string `json:"TrunkGroupKey,omitempty" xml:"TrunkGroupKey,omitempty"`
	// 彩铃文件名称
	//
	// example:
	//
	// name2
	TsiFile *string `json:"TsiFile,omitempty" xml:"TsiFile,omitempty"`
	// 彩铃文件类型
	//
	// example:
	//
	// mp3
	TsiFileFormat *string `json:"TsiFileFormat,omitempty" xml:"TsiFileFormat,omitempty"`
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
	// 通话质量百分比,数据是0-100之间，in方向说话占用的百分比
	//
	// example:
	//
	// 30
	VadIn *string `json:"VadIn,omitempty" xml:"VadIn,omitempty"`
	// 通话质量百分比,数据是0-100之间，out方向说话占用的百分比
	//
	// example:
	//
	// 70
	VadOut *string `json:"VadOut,omitempty" xml:"VadOut,omitempty"`
	// 客户等待时长秒数，单位是s
	//
	// example:
	//
	// 5
	WaitDuration *string `json:"WaitDuration,omitempty" xml:"WaitDuration,omitempty"`
	// 虚拟号码
	//
	// example:
	//
	// 13322444
	XNumber *string `json:"XNumber,omitempty" xml:"XNumber,omitempty"`
}

func (s CloudQueryObCdrResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryObCdrResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudQueryObCdrResponseBodyDataList) GetAgentClid() *string {
	return s.AgentClid
}

func (s *CloudQueryObCdrResponseBodyDataList) GetAgentName() *string {
	return s.AgentName
}

func (s *CloudQueryObCdrResponseBodyDataList) GetAgentNumber() *string {
	return s.AgentNumber
}

func (s *CloudQueryObCdrResponseBodyDataList) GetAnswerTime() *string {
	return s.AnswerTime
}

func (s *CloudQueryObCdrResponseBodyDataList) GetBridgeDuration() *string {
	return s.BridgeDuration
}

func (s *CloudQueryObCdrResponseBodyDataList) GetBridgeTime() *string {
	return s.BridgeTime
}

func (s *CloudQueryObCdrResponseBodyDataList) GetCallId() *string {
	return s.CallId
}

func (s *CloudQueryObCdrResponseBodyDataList) GetCallType() *string {
	return s.CallType
}

func (s *CloudQueryObCdrResponseBodyDataList) GetCalleeRingingTime() *string {
	return s.CalleeRingingTime
}

func (s *CloudQueryObCdrResponseBodyDataList) GetClid() *string {
	return s.Clid
}

func (s *CloudQueryObCdrResponseBodyDataList) GetCno() *string {
	return s.Cno
}

func (s *CloudQueryObCdrResponseBodyDataList) GetCustomerAreaCode() *string {
	return s.CustomerAreaCode
}

func (s *CloudQueryObCdrResponseBodyDataList) GetCustomerCity() *string {
	return s.CustomerCity
}

func (s *CloudQueryObCdrResponseBodyDataList) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *CloudQueryObCdrResponseBodyDataList) GetCustomerProvince() *string {
	return s.CustomerProvince
}

func (s *CloudQueryObCdrResponseBodyDataList) GetDisplayNumber() *string {
	return s.DisplayNumber
}

func (s *CloudQueryObCdrResponseBodyDataList) GetDownGrade() *string {
	return s.DownGrade
}

func (s *CloudQueryObCdrResponseBodyDataList) GetDownGradeCounts() *string {
	return s.DownGradeCounts
}

func (s *CloudQueryObCdrResponseBodyDataList) GetEndReason() *string {
	return s.EndReason
}

func (s *CloudQueryObCdrResponseBodyDataList) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudQueryObCdrResponseBodyDataList) GetGno() *string {
	return s.Gno
}

func (s *CloudQueryObCdrResponseBodyDataList) GetHangupReason() *string {
	return s.HangupReason
}

func (s *CloudQueryObCdrResponseBodyDataList) GetId() *string {
	return s.Id
}

func (s *CloudQueryObCdrResponseBodyDataList) GetIsRealAnswer() *string {
	return s.IsRealAnswer
}

func (s *CloudQueryObCdrResponseBodyDataList) GetLeftDisplayNumber() *string {
	return s.LeftDisplayNumber
}

func (s *CloudQueryObCdrResponseBodyDataList) GetMainRingingTime() *string {
	return s.MainRingingTime
}

func (s *CloudQueryObCdrResponseBodyDataList) GetRealAnswerTime() *string {
	return s.RealAnswerTime
}

func (s *CloudQueryObCdrResponseBodyDataList) GetRecordFile() []*CloudQueryObCdrResponseBodyDataListRecordFile {
	return s.RecordFile
}

func (s *CloudQueryObCdrResponseBodyDataList) GetRequestUniqueId() *string {
	return s.RequestUniqueId
}

func (s *CloudQueryObCdrResponseBodyDataList) GetRtcTotalDuration() *string {
	return s.RtcTotalDuration
}

func (s *CloudQueryObCdrResponseBodyDataList) GetSipCause() *string {
	return s.SipCause
}

func (s *CloudQueryObCdrResponseBodyDataList) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudQueryObCdrResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *CloudQueryObCdrResponseBodyDataList) GetStatusCode() *string {
	return s.StatusCode
}

func (s *CloudQueryObCdrResponseBodyDataList) GetStatusReason() *string {
	return s.StatusReason
}

func (s *CloudQueryObCdrResponseBodyDataList) GetTaskId() *string {
	return s.TaskId
}

func (s *CloudQueryObCdrResponseBodyDataList) GetTotalDuration() *string {
	return s.TotalDuration
}

func (s *CloudQueryObCdrResponseBodyDataList) GetTrunkGroupKey() *string {
	return s.TrunkGroupKey
}

func (s *CloudQueryObCdrResponseBodyDataList) GetTsiFile() *string {
	return s.TsiFile
}

func (s *CloudQueryObCdrResponseBodyDataList) GetTsiFileFormat() *string {
	return s.TsiFileFormat
}

func (s *CloudQueryObCdrResponseBodyDataList) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudQueryObCdrResponseBodyDataList) GetUserField() map[string]interface{} {
	return s.UserField
}

func (s *CloudQueryObCdrResponseBodyDataList) GetVadIn() *string {
	return s.VadIn
}

func (s *CloudQueryObCdrResponseBodyDataList) GetVadOut() *string {
	return s.VadOut
}

func (s *CloudQueryObCdrResponseBodyDataList) GetWaitDuration() *string {
	return s.WaitDuration
}

func (s *CloudQueryObCdrResponseBodyDataList) GetXNumber() *string {
	return s.XNumber
}

func (s *CloudQueryObCdrResponseBodyDataList) SetAgentClid(v string) *CloudQueryObCdrResponseBodyDataList {
	s.AgentClid = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetAgentName(v string) *CloudQueryObCdrResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetAgentNumber(v string) *CloudQueryObCdrResponseBodyDataList {
	s.AgentNumber = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetAnswerTime(v string) *CloudQueryObCdrResponseBodyDataList {
	s.AnswerTime = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetBridgeDuration(v string) *CloudQueryObCdrResponseBodyDataList {
	s.BridgeDuration = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetBridgeTime(v string) *CloudQueryObCdrResponseBodyDataList {
	s.BridgeTime = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetCallId(v string) *CloudQueryObCdrResponseBodyDataList {
	s.CallId = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetCallType(v string) *CloudQueryObCdrResponseBodyDataList {
	s.CallType = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetCalleeRingingTime(v string) *CloudQueryObCdrResponseBodyDataList {
	s.CalleeRingingTime = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetClid(v string) *CloudQueryObCdrResponseBodyDataList {
	s.Clid = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetCno(v string) *CloudQueryObCdrResponseBodyDataList {
	s.Cno = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetCustomerAreaCode(v string) *CloudQueryObCdrResponseBodyDataList {
	s.CustomerAreaCode = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetCustomerCity(v string) *CloudQueryObCdrResponseBodyDataList {
	s.CustomerCity = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetCustomerNumber(v string) *CloudQueryObCdrResponseBodyDataList {
	s.CustomerNumber = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetCustomerProvince(v string) *CloudQueryObCdrResponseBodyDataList {
	s.CustomerProvince = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetDisplayNumber(v string) *CloudQueryObCdrResponseBodyDataList {
	s.DisplayNumber = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetDownGrade(v string) *CloudQueryObCdrResponseBodyDataList {
	s.DownGrade = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetDownGradeCounts(v string) *CloudQueryObCdrResponseBodyDataList {
	s.DownGradeCounts = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetEndReason(v string) *CloudQueryObCdrResponseBodyDataList {
	s.EndReason = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetEndTime(v string) *CloudQueryObCdrResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetGno(v string) *CloudQueryObCdrResponseBodyDataList {
	s.Gno = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetHangupReason(v string) *CloudQueryObCdrResponseBodyDataList {
	s.HangupReason = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetId(v string) *CloudQueryObCdrResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetIsRealAnswer(v string) *CloudQueryObCdrResponseBodyDataList {
	s.IsRealAnswer = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetLeftDisplayNumber(v string) *CloudQueryObCdrResponseBodyDataList {
	s.LeftDisplayNumber = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetMainRingingTime(v string) *CloudQueryObCdrResponseBodyDataList {
	s.MainRingingTime = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetRealAnswerTime(v string) *CloudQueryObCdrResponseBodyDataList {
	s.RealAnswerTime = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetRecordFile(v []*CloudQueryObCdrResponseBodyDataListRecordFile) *CloudQueryObCdrResponseBodyDataList {
	s.RecordFile = v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetRequestUniqueId(v string) *CloudQueryObCdrResponseBodyDataList {
	s.RequestUniqueId = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetRtcTotalDuration(v string) *CloudQueryObCdrResponseBodyDataList {
	s.RtcTotalDuration = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetSipCause(v string) *CloudQueryObCdrResponseBodyDataList {
	s.SipCause = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetStartTime(v string) *CloudQueryObCdrResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetStatus(v string) *CloudQueryObCdrResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetStatusCode(v string) *CloudQueryObCdrResponseBodyDataList {
	s.StatusCode = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetStatusReason(v string) *CloudQueryObCdrResponseBodyDataList {
	s.StatusReason = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetTaskId(v string) *CloudQueryObCdrResponseBodyDataList {
	s.TaskId = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetTotalDuration(v string) *CloudQueryObCdrResponseBodyDataList {
	s.TotalDuration = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetTrunkGroupKey(v string) *CloudQueryObCdrResponseBodyDataList {
	s.TrunkGroupKey = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetTsiFile(v string) *CloudQueryObCdrResponseBodyDataList {
	s.TsiFile = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetTsiFileFormat(v string) *CloudQueryObCdrResponseBodyDataList {
	s.TsiFileFormat = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetUniqueId(v string) *CloudQueryObCdrResponseBodyDataList {
	s.UniqueId = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetUserField(v map[string]interface{}) *CloudQueryObCdrResponseBodyDataList {
	s.UserField = v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetVadIn(v string) *CloudQueryObCdrResponseBodyDataList {
	s.VadIn = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetVadOut(v string) *CloudQueryObCdrResponseBodyDataList {
	s.VadOut = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetWaitDuration(v string) *CloudQueryObCdrResponseBodyDataList {
	s.WaitDuration = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) SetXNumber(v string) *CloudQueryObCdrResponseBodyDataList {
	s.XNumber = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataList) Validate() error {
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

type CloudQueryObCdrResponseBodyDataListRecordFile struct {
	// 录音文件名
	//
	// example:
	//
	// record_file_name.
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

func (s CloudQueryObCdrResponseBodyDataListRecordFile) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryObCdrResponseBodyDataListRecordFile) GoString() string {
	return s.String()
}

func (s *CloudQueryObCdrResponseBodyDataListRecordFile) GetFile() *string {
	return s.File
}

func (s *CloudQueryObCdrResponseBodyDataListRecordFile) GetMonitorType() *int64 {
	return s.MonitorType
}

func (s *CloudQueryObCdrResponseBodyDataListRecordFile) GetType() *string {
	return s.Type
}

func (s *CloudQueryObCdrResponseBodyDataListRecordFile) SetFile(v string) *CloudQueryObCdrResponseBodyDataListRecordFile {
	s.File = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataListRecordFile) SetMonitorType(v int64) *CloudQueryObCdrResponseBodyDataListRecordFile {
	s.MonitorType = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataListRecordFile) SetType(v string) *CloudQueryObCdrResponseBodyDataListRecordFile {
	s.Type = &v
	return s
}

func (s *CloudQueryObCdrResponseBodyDataListRecordFile) Validate() error {
	return dara.Validate(s)
}
