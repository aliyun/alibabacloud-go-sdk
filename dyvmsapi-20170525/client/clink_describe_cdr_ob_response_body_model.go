// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeCdrObResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkDescribeCdrObResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkDescribeCdrObResponseBody
	GetCode() *string
	SetData(v *ClinkDescribeCdrObResponseBodyData) *ClinkDescribeCdrObResponseBody
	GetData() *ClinkDescribeCdrObResponseBodyData
	SetMessage(v string) *ClinkDescribeCdrObResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkDescribeCdrObResponseBody
	GetRequestId() *string
}

type ClinkDescribeCdrObResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkDescribeCdrObResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkDescribeCdrObResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrObResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrObResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkDescribeCdrObResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkDescribeCdrObResponseBody) GetData() *ClinkDescribeCdrObResponseBodyData {
	return s.Data
}

func (s *ClinkDescribeCdrObResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkDescribeCdrObResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkDescribeCdrObResponseBody) SetAccessDeniedDetail(v string) *ClinkDescribeCdrObResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBody) SetCode(v string) *ClinkDescribeCdrObResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBody) SetData(v *ClinkDescribeCdrObResponseBodyData) *ClinkDescribeCdrObResponseBody {
	s.Data = v
	return s
}

func (s *ClinkDescribeCdrObResponseBody) SetMessage(v string) *ClinkDescribeCdrObResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBody) SetRequestId(v string) *ClinkDescribeCdrObResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkDescribeCdrObResponseBodyData struct {
	// [外呼通话记录] #外呼通话记录
	CdrOb *ClinkDescribeCdrObResponseBodyDataCdrOb `json:"CdrOb,omitempty" xml:"CdrOb,omitempty" type:"Struct"`
	// 请求 id
	//
	// example:
	//
	// 示例值
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
}

func (s ClinkDescribeCdrObResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrObResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrObResponseBodyData) GetCdrOb() *ClinkDescribeCdrObResponseBodyDataCdrOb {
	return s.CdrOb
}

func (s *ClinkDescribeCdrObResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkDescribeCdrObResponseBodyData) SetCdrOb(v *ClinkDescribeCdrObResponseBodyDataCdrOb) *ClinkDescribeCdrObResponseBodyData {
	s.CdrOb = v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyData) SetClinkRequestId(v string) *ClinkDescribeCdrObResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyData) Validate() error {
	if s.CdrOb != nil {
		if err := s.CdrOb.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkDescribeCdrObResponseBodyDataCdrOb struct {
	// 接通时长
	//
	// example:
	//
	// 100
	BridgeDuration *int64 `json:"BridgeDuration,omitempty" xml:"BridgeDuration,omitempty"`
	// 接通时间
	//
	// example:
	//
	// 82
	BridgeTime *int64 `json:"BridgeTime,omitempty" xml:"BridgeTime,omitempty"`
	// 通话记录id
	//
	// example:
	//
	// ba078ace-717c-6666-af41-d4dd5035edd6
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 呼叫类型
	//
	// example:
	//
	// 示例值示例值
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 座席名称
	//
	// example:
	//
	// 示例值
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
	// 2555
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 客户城市
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
	// 客户省份
	//
	// example:
	//
	// 示例值示例值
	CustomerProvince *string `json:"CustomerProvince,omitempty" xml:"CustomerProvince,omitempty"`
	// 客户响铃时间
	//
	// example:
	//
	// 66
	CustomerRingingTime *int64 `json:"CustomerRingingTime,omitempty" xml:"CustomerRingingTime,omitempty"`
	// 挂机方
	//
	// example:
	//
	// 示例值示例值示例值
	EndReason *string `json:"EndReason,omitempty" xml:"EndReason,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 1536892706
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 是否邀评 0: 否 1: 是
	//
	// example:
	//
	// 0
	Evaluation *int64 `json:"Evaluation,omitempty" xml:"Evaluation,omitempty"`
	// 热线号码
	//
	// example:
	//
	// xxx
	Hotline *string `json:"Hotline,omitempty" xml:"Hotline,omitempty"`
	// [满意度记录] #满意度记录
	Investigation map[string]interface{} `json:"Investigation,omitempty" xml:"Investigation,omitempty"`
	// ivr名称
	//
	// example:
	//
	// name1
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
	// 示例值示例值示例值
	MarkData *string `json:"MarkData,omitempty" xml:"MarkData,omitempty"`
	// 录音文件名
	//
	// example:
	//
	// xxx
	RecordFile *string `json:"RecordFile,omitempty" xml:"RecordFile,omitempty"`
	// 呼叫结果
	//
	// example:
	//
	// null
	SipCause *string `json:"SipCause,omitempty" xml:"SipCause,omitempty"`
	// 开始时间
	//
	// example:
	//
	// 1536115324
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
	// 64
	TotalDuration *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// 通话记录唯一标识
	//
	// example:
	//
	// AWS_DEV_MEDIA_SERVER_8-1536892698.2
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// 虚拟号码
	//
	// example:
	//
	// xxx
	Xnumber *string `json:"Xnumber,omitempty" xml:"Xnumber,omitempty"`
}

func (s ClinkDescribeCdrObResponseBodyDataCdrOb) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrObResponseBodyDataCdrOb) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetBridgeDuration() *int64 {
	return s.BridgeDuration
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetBridgeTime() *int64 {
	return s.BridgeTime
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetCallId() *string {
	return s.CallId
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetCallType() *string {
	return s.CallType
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetClientName() *string {
	return s.ClientName
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetClientNumber() *string {
	return s.ClientNumber
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetCno() *string {
	return s.Cno
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetCustomerCity() *string {
	return s.CustomerCity
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetCustomerNumberEncrypt() *string {
	return s.CustomerNumberEncrypt
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetCustomerProvince() *string {
	return s.CustomerProvince
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetCustomerRingingTime() *int64 {
	return s.CustomerRingingTime
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetEndReason() *string {
	return s.EndReason
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetEvaluation() *int64 {
	return s.Evaluation
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetHotline() *string {
	return s.Hotline
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetInvestigation() map[string]interface{} {
	return s.Investigation
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetIvrName() *string {
	return s.IvrName
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetMark() *int64 {
	return s.Mark
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetMarkData() *string {
	return s.MarkData
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetRecordFile() *string {
	return s.RecordFile
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetSipCause() *string {
	return s.SipCause
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetStatus() *string {
	return s.Status
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetTagNames() []*string {
	return s.TagNames
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetTags() []*string {
	return s.Tags
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetUniqueId() *string {
	return s.UniqueId
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) GetXnumber() *string {
	return s.Xnumber
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetBridgeDuration(v int64) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.BridgeDuration = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetBridgeTime(v int64) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.BridgeTime = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetCallId(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.CallId = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetCallType(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.CallType = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetClientName(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.ClientName = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetClientNumber(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.ClientNumber = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetCno(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.Cno = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetCustomerCity(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.CustomerCity = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetCustomerNumber(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.CustomerNumber = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetCustomerNumberEncrypt(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.CustomerNumberEncrypt = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetCustomerProvince(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.CustomerProvince = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetCustomerRingingTime(v int64) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.CustomerRingingTime = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetEndReason(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.EndReason = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetEndTime(v int64) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.EndTime = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetEvaluation(v int64) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.Evaluation = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetHotline(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.Hotline = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetInvestigation(v map[string]interface{}) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.Investigation = v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetIvrName(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.IvrName = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetMark(v int64) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.Mark = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetMarkData(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.MarkData = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetRecordFile(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.RecordFile = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetSipCause(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.SipCause = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetStartTime(v int64) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.StartTime = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetStatus(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.Status = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetStatusCode(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.StatusCode = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetTagNames(v []*string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.TagNames = v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetTags(v []*string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.Tags = v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetTotalDuration(v int64) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.TotalDuration = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetUniqueId(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.UniqueId = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) SetXnumber(v string) *ClinkDescribeCdrObResponseBodyDataCdrOb {
	s.Xnumber = &v
	return s
}

func (s *ClinkDescribeCdrObResponseBodyDataCdrOb) Validate() error {
	return dara.Validate(s)
}
