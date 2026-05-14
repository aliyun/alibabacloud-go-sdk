// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeCdrObDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkDescribeCdrObDetailsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkDescribeCdrObDetailsResponseBody
	GetCode() *string
	SetData(v *ClinkDescribeCdrObDetailsResponseBodyData) *ClinkDescribeCdrObDetailsResponseBody
	GetData() *ClinkDescribeCdrObDetailsResponseBodyData
	SetMessage(v string) *ClinkDescribeCdrObDetailsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkDescribeCdrObDetailsResponseBody
	GetRequestId() *string
}

type ClinkDescribeCdrObDetailsResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkDescribeCdrObDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkDescribeCdrObDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrObDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrObDetailsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkDescribeCdrObDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkDescribeCdrObDetailsResponseBody) GetData() *ClinkDescribeCdrObDetailsResponseBodyData {
	return s.Data
}

func (s *ClinkDescribeCdrObDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkDescribeCdrObDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkDescribeCdrObDetailsResponseBody) SetAccessDeniedDetail(v string) *ClinkDescribeCdrObDetailsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBody) SetCode(v string) *ClinkDescribeCdrObDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBody) SetData(v *ClinkDescribeCdrObDetailsResponseBodyData) *ClinkDescribeCdrObDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBody) SetMessage(v string) *ClinkDescribeCdrObDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBody) SetRequestId(v string) *ClinkDescribeCdrObDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkDescribeCdrObDetailsResponseBodyData struct {
	// [外呼通话记录明细] #外呼通话记录明细
	CdrObDetails []*ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails `json:"CdrObDetails,omitempty" xml:"CdrObDetails,omitempty" type:"Repeated"`
	// 请求 id
	//
	// example:
	//
	// xxx
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
}

func (s ClinkDescribeCdrObDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrObDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrObDetailsResponseBodyData) GetCdrObDetails() []*ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails {
	return s.CdrObDetails
}

func (s *ClinkDescribeCdrObDetailsResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkDescribeCdrObDetailsResponseBodyData) SetCdrObDetails(v []*ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) *ClinkDescribeCdrObDetailsResponseBodyData {
	s.CdrObDetails = v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyData) SetClinkRequestId(v string) *ClinkDescribeCdrObDetailsResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyData) Validate() error {
	if s.CdrObDetails != nil {
		for _, item := range s.CdrObDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails struct {
	// 接听时间
	//
	// example:
	//
	// 1536115382
	AnswerTime *int64 `json:"AnswerTime,omitempty" xml:"AnswerTime,omitempty"`
	// 呼叫类型
	//
	// example:
	//
	// 示例值
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 座席名称
	//
	// example:
	//
	// zhangsan
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
	// 2005
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 1536115391
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 通话记录主通道唯一标识
	//
	// example:
	//
	// AWS_DEV_MEDIA_SERVER_8-1536115322.0
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	// 录音文件名
	//
	// example:
	//
	// xxx
	RecordFile *string `json:"RecordFile,omitempty" xml:"RecordFile,omitempty"`
	// 是否开启主叫记忆
	//
	// example:
	//
	// 示例值示例值
	Remember *string `json:"Remember,omitempty" xml:"Remember,omitempty"`
	// 呼叫情况
	//
	// example:
	//
	// 示例值示例值示例值
	SipCause *string `json:"SipCause,omitempty" xml:"SipCause,omitempty"`
	// 接起时间
	//
	// example:
	//
	// 1536115379
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 呼叫结果
	//
	// example:
	//
	// 示例值示例值示例值
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 通话时长
	//
	// example:
	//
	// 12
	TotalDuration *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// 通话记录详情唯一标识
	//
	// example:
	//
	// AWS_DEV_MEDIA_SERVER_8-1536115379.4
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) GetAnswerTime() *int64 {
	return s.AnswerTime
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) GetCallType() *string {
	return s.CallType
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) GetClientName() *string {
	return s.ClientName
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) GetClientNumber() *string {
	return s.ClientNumber
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) GetCno() *string {
	return s.Cno
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) GetRecordFile() *string {
	return s.RecordFile
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) GetRemember() *string {
	return s.Remember
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) GetSipCause() *string {
	return s.SipCause
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) GetStatus() *string {
	return s.Status
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) GetUniqueId() *string {
	return s.UniqueId
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) SetAnswerTime(v int64) *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails {
	s.AnswerTime = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) SetCallType(v string) *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails {
	s.CallType = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) SetClientName(v string) *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails {
	s.ClientName = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) SetClientNumber(v string) *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails {
	s.ClientNumber = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) SetCno(v string) *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails {
	s.Cno = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) SetEndTime(v int64) *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails {
	s.EndTime = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) SetMainUniqueId(v string) *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) SetRecordFile(v string) *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails {
	s.RecordFile = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) SetRemember(v string) *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails {
	s.Remember = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) SetSipCause(v string) *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails {
	s.SipCause = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) SetStartTime(v int64) *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails {
	s.StartTime = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) SetStatus(v string) *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails {
	s.Status = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) SetTotalDuration(v int64) *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails {
	s.TotalDuration = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) SetUniqueId(v string) *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails {
	s.UniqueId = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsResponseBodyDataCdrObDetails) Validate() error {
	return dara.Validate(s)
}
