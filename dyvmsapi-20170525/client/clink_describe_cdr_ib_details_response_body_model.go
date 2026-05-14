// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeCdrIbDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkDescribeCdrIbDetailsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkDescribeCdrIbDetailsResponseBody
	GetCode() *string
	SetData(v *ClinkDescribeCdrIbDetailsResponseBodyData) *ClinkDescribeCdrIbDetailsResponseBody
	GetData() *ClinkDescribeCdrIbDetailsResponseBodyData
	SetMessage(v string) *ClinkDescribeCdrIbDetailsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkDescribeCdrIbDetailsResponseBody
	GetRequestId() *string
}

type ClinkDescribeCdrIbDetailsResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkDescribeCdrIbDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkDescribeCdrIbDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrIbDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrIbDetailsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkDescribeCdrIbDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkDescribeCdrIbDetailsResponseBody) GetData() *ClinkDescribeCdrIbDetailsResponseBodyData {
	return s.Data
}

func (s *ClinkDescribeCdrIbDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkDescribeCdrIbDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkDescribeCdrIbDetailsResponseBody) SetAccessDeniedDetail(v string) *ClinkDescribeCdrIbDetailsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBody) SetCode(v string) *ClinkDescribeCdrIbDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBody) SetData(v *ClinkDescribeCdrIbDetailsResponseBodyData) *ClinkDescribeCdrIbDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBody) SetMessage(v string) *ClinkDescribeCdrIbDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBody) SetRequestId(v string) *ClinkDescribeCdrIbDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkDescribeCdrIbDetailsResponseBodyData struct {
	// [呼入通话记录明细] #呼入通话记录明细
	CdrIbDetails []*ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails `json:"CdrIbDetails,omitempty" xml:"CdrIbDetails,omitempty" type:"Repeated"`
	// 请求 id
	//
	// example:
	//
	// xxx
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
}

func (s ClinkDescribeCdrIbDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrIbDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyData) GetCdrIbDetails() []*ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	return s.CdrIbDetails
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyData) SetCdrIbDetails(v []*ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) *ClinkDescribeCdrIbDetailsResponseBodyData {
	s.CdrIbDetails = v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyData) SetClinkRequestId(v string) *ClinkDescribeCdrIbDetailsResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyData) Validate() error {
	if s.CdrIbDetails != nil {
		for _, item := range s.CdrIbDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails struct {
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
	// 示例值示例值
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 座席名称
	//
	// example:
	//
	// 示例值示例值
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
	// 1536115379
	ClientRingingTime *int64 `json:"ClientRingingTime,omitempty" xml:"ClientRingingTime,omitempty"`
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
	// 队列号
	//
	// example:
	//
	// 2233
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// 录音文件
	//
	// example:
	//
	// 8888888-20220406155839-15108207489-6666--record-medias_6-777788888.666
	RecordFile *string `json:"RecordFile,omitempty" xml:"RecordFile,omitempty"`
	// 是否开启主叫记忆
	//
	// example:
	//
	// 示例值示例值示例值
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
	// 示例值示例值
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 通话时长
	//
	// example:
	//
	// 47
	TotalDuration *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// 通话记录详情唯一标识
	//
	// example:
	//
	// AWS_DEV_MEDIA_SERVER_8-1536115379.4
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetAnswerTime() *int64 {
	return s.AnswerTime
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetCallType() *string {
	return s.CallType
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetClientName() *string {
	return s.ClientName
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetClientNumber() *string {
	return s.ClientNumber
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetClientRingingTime() *int64 {
	return s.ClientRingingTime
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetCno() *string {
	return s.Cno
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetQno() *string {
	return s.Qno
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetRecordFile() *string {
	return s.RecordFile
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetRemember() *string {
	return s.Remember
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetSipCause() *string {
	return s.SipCause
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetStatus() *string {
	return s.Status
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) GetUniqueId() *string {
	return s.UniqueId
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetAnswerTime(v int64) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.AnswerTime = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetCallType(v string) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.CallType = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetClientName(v string) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.ClientName = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetClientNumber(v string) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.ClientNumber = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetClientRingingTime(v int64) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.ClientRingingTime = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetCno(v string) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.Cno = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetEndTime(v int64) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.EndTime = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetMainUniqueId(v string) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetQno(v string) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.Qno = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetRecordFile(v string) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.RecordFile = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetRemember(v string) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.Remember = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetSipCause(v string) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.SipCause = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetStartTime(v int64) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.StartTime = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetStatus(v string) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.Status = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetTotalDuration(v int64) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.TotalDuration = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) SetUniqueId(v string) *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails {
	s.UniqueId = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsResponseBodyDataCdrIbDetails) Validate() error {
	return dara.Validate(s)
}
