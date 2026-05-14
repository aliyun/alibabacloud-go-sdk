// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListCdrObsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkListCdrObsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkListCdrObsResponseBody
	GetCode() *string
	SetData(v *ClinkListCdrObsResponseBodyData) *ClinkListCdrObsResponseBody
	GetData() *ClinkListCdrObsResponseBodyData
	SetMessage(v string) *ClinkListCdrObsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkListCdrObsResponseBody
	GetRequestId() *string
}

type ClinkListCdrObsResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkListCdrObsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4B52E871-0F7C-59FC-9C7E-94B550953BA9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkListCdrObsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrObsResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkListCdrObsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkListCdrObsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkListCdrObsResponseBody) GetData() *ClinkListCdrObsResponseBodyData {
	return s.Data
}

func (s *ClinkListCdrObsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkListCdrObsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkListCdrObsResponseBody) SetAccessDeniedDetail(v string) *ClinkListCdrObsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkListCdrObsResponseBody) SetCode(v string) *ClinkListCdrObsResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkListCdrObsResponseBody) SetData(v *ClinkListCdrObsResponseBodyData) *ClinkListCdrObsResponseBody {
	s.Data = v
	return s
}

func (s *ClinkListCdrObsResponseBody) SetMessage(v string) *ClinkListCdrObsResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkListCdrObsResponseBody) SetRequestId(v string) *ClinkListCdrObsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkListCdrObsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkListCdrObsResponseBodyData struct {
	// [外呼通话记录列表] #外呼通话记录列表
	CdrObs []*ClinkListCdrObsResponseBodyDataCdrObs `json:"CdrObs,omitempty" xml:"CdrObs,omitempty" type:"Repeated"`
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
	// 49
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ClinkListCdrObsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrObsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkListCdrObsResponseBodyData) GetCdrObs() []*ClinkListCdrObsResponseBodyDataCdrObs {
	return s.CdrObs
}

func (s *ClinkListCdrObsResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkListCdrObsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ClinkListCdrObsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ClinkListCdrObsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ClinkListCdrObsResponseBodyData) SetCdrObs(v []*ClinkListCdrObsResponseBodyDataCdrObs) *ClinkListCdrObsResponseBodyData {
	s.CdrObs = v
	return s
}

func (s *ClinkListCdrObsResponseBodyData) SetClinkRequestId(v string) *ClinkListCdrObsResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyData) SetPageNumber(v int64) *ClinkListCdrObsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyData) SetPageSize(v int64) *ClinkListCdrObsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyData) SetTotalCount(v int64) *ClinkListCdrObsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyData) Validate() error {
	if s.CdrObs != nil {
		for _, item := range s.CdrObs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkListCdrObsResponseBodyDataCdrObs struct {
	// 接通时长
	//
	// example:
	//
	// 87
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
	// CallId
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 外呼类型，预览外呼、预测外呼、直接外呼
	//
	// example:
	//
	// 示例值示例值示例值
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
	// 客户城市
	//
	// example:
	//
	// 示例值示例值
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
	// 客户省份
	//
	// example:
	//
	// 示例值
	CustomerProvince *string `json:"CustomerProvince,omitempty" xml:"CustomerProvince,omitempty"`
	// 客户响铃时间
	//
	// example:
	//
	// 1775024775
	CustomerRingingTime *int64 `json:"CustomerRingingTime,omitempty" xml:"CustomerRingingTime,omitempty"`
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
	// 是否邀评
	//
	// example:
	//
	// 0
	Evaluation *int64 `json:"Evaluation,omitempty" xml:"Evaluation,omitempty"`
	// 热线号码
	//
	// example:
	//
	// Hotline
	Hotline *string `json:"Hotline,omitempty" xml:"Hotline,omitempty"`
	// IVR名称
	//
	// example:
	//
	// IvrName
	IvrName *string `json:"IvrName,omitempty" xml:"IvrName,omitempty"`
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
	// 2222
	Qname *string `json:"Qname,omitempty" xml:"Qname,omitempty"`
	// 来电队列号
	//
	// example:
	//
	// 1212
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
	// 呼叫结果
	//
	// example:
	//
	// 示例值示例值示例值
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
	// 示例值
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 接听状态映射
	//
	// example:
	//
	// 示例值示例值
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// 展示通话标签详情,当请求参数fields中包含tagNames时返回
	TagNames []*string `json:"TagNames,omitempty" xml:"TagNames,omitempty" type:"Repeated"`
	// 标签
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 总时长
	//
	// example:
	//
	// 81
	TotalDuration *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// 通话记录唯一标识
	//
	// example:
	//
	// UniqueId
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// 虚拟号码
	//
	// example:
	//
	// xxx
	Xnumber *string `json:"Xnumber,omitempty" xml:"Xnumber,omitempty"`
}

func (s ClinkListCdrObsResponseBodyDataCdrObs) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrObsResponseBodyDataCdrObs) GoString() string {
	return s.String()
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetBridgeDuration() *int64 {
	return s.BridgeDuration
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetBridgeTime() *int64 {
	return s.BridgeTime
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetCallId() *string {
	return s.CallId
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetCallType() *string {
	return s.CallType
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetClientName() *string {
	return s.ClientName
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetClientNumber() *string {
	return s.ClientNumber
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetClientRingingTime() *int64 {
	return s.ClientRingingTime
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetCno() *string {
	return s.Cno
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetCustomerCity() *string {
	return s.CustomerCity
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetCustomerNumberEncrypt() *string {
	return s.CustomerNumberEncrypt
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetCustomerProvince() *string {
	return s.CustomerProvince
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetCustomerRingingTime() *int64 {
	return s.CustomerRingingTime
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetEndReason() *string {
	return s.EndReason
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetEvaluation() *int64 {
	return s.Evaluation
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetHotline() *string {
	return s.Hotline
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetIvrName() *string {
	return s.IvrName
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetMark() *int64 {
	return s.Mark
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetMarkData() *string {
	return s.MarkData
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetQname() *string {
	return s.Qname
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetQno() *string {
	return s.Qno
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetQueueAnswerInTime() *int64 {
	return s.QueueAnswerInTime
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetRecordFile() *string {
	return s.RecordFile
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetSipCause() *string {
	return s.SipCause
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetStatus() *string {
	return s.Status
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetTagNames() []*string {
	return s.TagNames
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetTags() []*string {
	return s.Tags
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetUniqueId() *string {
	return s.UniqueId
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) GetXnumber() *string {
	return s.Xnumber
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetBridgeDuration(v int64) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.BridgeDuration = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetBridgeTime(v int64) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.BridgeTime = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetCallId(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.CallId = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetCallType(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.CallType = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetClientName(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.ClientName = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetClientNumber(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.ClientNumber = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetClientRingingTime(v int64) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.ClientRingingTime = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetCno(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.Cno = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetCustomerCity(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.CustomerCity = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetCustomerNumber(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.CustomerNumber = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetCustomerNumberEncrypt(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.CustomerNumberEncrypt = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetCustomerProvince(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.CustomerProvince = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetCustomerRingingTime(v int64) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.CustomerRingingTime = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetEndReason(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.EndReason = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetEndTime(v int64) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.EndTime = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetEvaluation(v int64) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.Evaluation = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetHotline(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.Hotline = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetIvrName(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.IvrName = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetMark(v int64) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.Mark = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetMarkData(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.MarkData = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetQname(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.Qname = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetQno(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.Qno = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetQueueAnswerInTime(v int64) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.QueueAnswerInTime = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetRecordFile(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.RecordFile = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetSipCause(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.SipCause = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetStartTime(v int64) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.StartTime = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetStatus(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.Status = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetStatusCode(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.StatusCode = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetTagNames(v []*string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.TagNames = v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetTags(v []*string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.Tags = v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetTotalDuration(v int64) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.TotalDuration = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetUniqueId(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.UniqueId = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) SetXnumber(v string) *ClinkListCdrObsResponseBodyDataCdrObs {
	s.Xnumber = &v
	return s
}

func (s *ClinkListCdrObsResponseBodyDataCdrObs) Validate() error {
	return dara.Validate(s)
}
