// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryIbCdrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudQueryIbCdrResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudQueryIbCdrResponseBody
	GetCode() *string
	SetData(v *CloudQueryIbCdrResponseBodyData) *CloudQueryIbCdrResponseBody
	GetData() *CloudQueryIbCdrResponseBodyData
	SetMessage(v string) *CloudQueryIbCdrResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudQueryIbCdrResponseBody
	GetRequestId() *string
}

type CloudQueryIbCdrResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudQueryIbCdrResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudQueryIbCdrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryIbCdrResponseBody) GoString() string {
	return s.String()
}

func (s *CloudQueryIbCdrResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudQueryIbCdrResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudQueryIbCdrResponseBody) GetData() *CloudQueryIbCdrResponseBodyData {
	return s.Data
}

func (s *CloudQueryIbCdrResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudQueryIbCdrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudQueryIbCdrResponseBody) SetAccessDeniedDetail(v string) *CloudQueryIbCdrResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudQueryIbCdrResponseBody) SetCode(v string) *CloudQueryIbCdrResponseBody {
	s.Code = &v
	return s
}

func (s *CloudQueryIbCdrResponseBody) SetData(v *CloudQueryIbCdrResponseBodyData) *CloudQueryIbCdrResponseBody {
	s.Data = v
	return s
}

func (s *CloudQueryIbCdrResponseBody) SetMessage(v string) *CloudQueryIbCdrResponseBody {
	s.Message = &v
	return s
}

func (s *CloudQueryIbCdrResponseBody) SetRequestId(v string) *CloudQueryIbCdrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudQueryIbCdrResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudQueryIbCdrResponseBodyData struct {
	// 来电通话记录数组
	List []*CloudQueryIbCdrResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudQueryIbCdrResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryIbCdrResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudQueryIbCdrResponseBodyData) GetList() []*CloudQueryIbCdrResponseBodyDataList {
	return s.List
}

func (s *CloudQueryIbCdrResponseBodyData) SetList(v []*CloudQueryIbCdrResponseBodyDataList) *CloudQueryIbCdrResponseBodyData {
	s.List = v
	return s
}

func (s *CloudQueryIbCdrResponseBodyData) Validate() error {
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

type CloudQueryIbCdrResponseBodyDataList struct {
	// 座席名称
	//
	// example:
	//
	// name1
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
	// 10
	BridgeDuration *string `json:"BridgeDuration,omitempty" xml:"BridgeDuration,omitempty"`
	// 座席接听时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1775024775
	BridgeTime *string `json:"BridgeTime,omitempty" xml:"BridgeTime,omitempty"`
	// 呼叫类型 呼入
	//
	// example:
	//
	// 示例值示例值示例值
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 座席电话 区号 +7或8位的固话号码，区号与固话号码之间无\\"-\\"；或11位长度的手机号码。0104100596
	//
	// example:
	//
	// 0104100596
	CalleeNumber *string `json:"CalleeNumber,omitempty" xml:"CalleeNumber,omitempty"`
	// 座席工号，如2000
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 客户号码区号
	//
	// example:
	//
	// 010
	CustomerAreaCode *string `json:"CustomerAreaCode,omitempty" xml:"CustomerAreaCode,omitempty"`
	// 客户号码归属城市,北京
	//
	// example:
	//
	// 示例值示例值
	CustomerCity *string `json:"CustomerCity,omitempty" xml:"CustomerCity,omitempty"`
	// 客户号码 国内固话或手机，区号 + 7或8位的固话号码，区号与固话号码之间无\\"-\\"；或11位长度的手机号码。如 01041005968或18701051984
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
	// 热线号码
	//
	// example:
	//
	// 89193631
	Hotline *string `json:"Hotline,omitempty" xml:"Hotline,omitempty"`
	// IVR名称，默认自定义
	//
	// example:
	//
	// ivrname
	IvrName *string `json:"IvrName,omitempty" xml:"IvrName,omitempty"`
	// 离开队列原因 参数说明：-1:该字段没值, 1:呼转座席接听 , 2: 队列中放弃 , 3: 队列中超时溢出 , 4: 队列中无座席溢出
	//
	// example:
	//
	// -1
	LeaveQueueCode *int64 `json:"LeaveQueueCode,omitempty" xml:"LeaveQueueCode,omitempty"`
	// 中继号码
	//
	// example:
	//
	// 89193631
	NumberTrunk *string `json:"NumberTrunk,omitempty" xml:"NumberTrunk,omitempty"`
	// 中继号码区号
	//
	// example:
	//
	// 010
	NumberTrunkAreaCode *string `json:"NumberTrunkAreaCode,omitempty" xml:"NumberTrunkAreaCode,omitempty"`
	// 队列号，如1000
	//
	// example:
	//
	// 1000
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// 通话记录录音数组，json格式
	RecordFile []*CloudQueryIbCdrResponseBodyDataListRecordFile `json:"RecordFile,omitempty" xml:"RecordFile,omitempty" type:"Repeated"`
	// 电话进入系统时间，时间戳，精确到s，如1480904471
	//
	// example:
	//
	// 1775024775
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 外呼结果， 如人工接听
	//
	// example:
	//
	// 示例值示例值
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// status对应的状态码
	//
	// example:
	//
	// 3
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// 总通话时长秒数，单位是s
	//
	// example:
	//
	// 133
	TotalDuration *string `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
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
}

func (s CloudQueryIbCdrResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryIbCdrResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetAgentName() *string {
	return s.AgentName
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetAnswerTime() *string {
	return s.AnswerTime
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetBridgeDuration() *string {
	return s.BridgeDuration
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetBridgeTime() *string {
	return s.BridgeTime
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetCallType() *string {
	return s.CallType
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetCalleeNumber() *string {
	return s.CalleeNumber
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetCno() *string {
	return s.Cno
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetCustomerAreaCode() *string {
	return s.CustomerAreaCode
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetCustomerCity() *string {
	return s.CustomerCity
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetCustomerProvince() *string {
	return s.CustomerProvince
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetEndReason() *string {
	return s.EndReason
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetHotline() *string {
	return s.Hotline
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetIvrName() *string {
	return s.IvrName
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetLeaveQueueCode() *int64 {
	return s.LeaveQueueCode
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetNumberTrunk() *string {
	return s.NumberTrunk
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetNumberTrunkAreaCode() *string {
	return s.NumberTrunkAreaCode
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetQno() *string {
	return s.Qno
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetRecordFile() []*CloudQueryIbCdrResponseBodyDataListRecordFile {
	return s.RecordFile
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetStatusCode() *string {
	return s.StatusCode
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetTotalDuration() *string {
	return s.TotalDuration
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudQueryIbCdrResponseBodyDataList) GetUserField() map[string]interface{} {
	return s.UserField
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetAgentName(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetAnswerTime(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.AnswerTime = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetBridgeDuration(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.BridgeDuration = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetBridgeTime(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.BridgeTime = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetCallType(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.CallType = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetCalleeNumber(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.CalleeNumber = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetCno(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.Cno = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetCustomerAreaCode(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.CustomerAreaCode = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetCustomerCity(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.CustomerCity = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetCustomerNumber(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.CustomerNumber = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetCustomerProvince(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.CustomerProvince = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetEndReason(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.EndReason = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetEndTime(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetHotline(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.Hotline = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetIvrName(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.IvrName = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetLeaveQueueCode(v int64) *CloudQueryIbCdrResponseBodyDataList {
	s.LeaveQueueCode = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetNumberTrunk(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.NumberTrunk = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetNumberTrunkAreaCode(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.NumberTrunkAreaCode = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetQno(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.Qno = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetRecordFile(v []*CloudQueryIbCdrResponseBodyDataListRecordFile) *CloudQueryIbCdrResponseBodyDataList {
	s.RecordFile = v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetStartTime(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetStatus(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetStatusCode(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.StatusCode = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetTotalDuration(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.TotalDuration = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetUniqueId(v string) *CloudQueryIbCdrResponseBodyDataList {
	s.UniqueId = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) SetUserField(v map[string]interface{}) *CloudQueryIbCdrResponseBodyDataList {
	s.UserField = v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataList) Validate() error {
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

type CloudQueryIbCdrResponseBodyDataListRecordFile struct {
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

func (s CloudQueryIbCdrResponseBodyDataListRecordFile) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryIbCdrResponseBodyDataListRecordFile) GoString() string {
	return s.String()
}

func (s *CloudQueryIbCdrResponseBodyDataListRecordFile) GetFile() *string {
	return s.File
}

func (s *CloudQueryIbCdrResponseBodyDataListRecordFile) GetMonitorType() *int64 {
	return s.MonitorType
}

func (s *CloudQueryIbCdrResponseBodyDataListRecordFile) GetType() *string {
	return s.Type
}

func (s *CloudQueryIbCdrResponseBodyDataListRecordFile) SetFile(v string) *CloudQueryIbCdrResponseBodyDataListRecordFile {
	s.File = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataListRecordFile) SetMonitorType(v int64) *CloudQueryIbCdrResponseBodyDataListRecordFile {
	s.MonitorType = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataListRecordFile) SetType(v string) *CloudQueryIbCdrResponseBodyDataListRecordFile {
	s.Type = &v
	return s
}

func (s *CloudQueryIbCdrResponseBodyDataListRecordFile) Validate() error {
	return dara.Validate(s)
}
