// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCdrObDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkCdrObDetailsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkCdrObDetailsResponseBody
	GetCode() *string
	SetData(v *ClinkCdrObDetailsResponseBodyData) *ClinkCdrObDetailsResponseBody
	GetData() *ClinkCdrObDetailsResponseBodyData
	SetMessage(v string) *ClinkCdrObDetailsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkCdrObDetailsResponseBody
	GetRequestId() *string
}

type ClinkCdrObDetailsResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkCdrObDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkCdrObDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkCdrObDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkCdrObDetailsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkCdrObDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkCdrObDetailsResponseBody) GetData() *ClinkCdrObDetailsResponseBodyData {
	return s.Data
}

func (s *ClinkCdrObDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkCdrObDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkCdrObDetailsResponseBody) SetAccessDeniedDetail(v string) *ClinkCdrObDetailsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBody) SetCode(v string) *ClinkCdrObDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBody) SetData(v *ClinkCdrObDetailsResponseBodyData) *ClinkCdrObDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ClinkCdrObDetailsResponseBody) SetMessage(v string) *ClinkCdrObDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBody) SetRequestId(v string) *ClinkCdrObDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkCdrObDetailsResponseBodyData struct {
	// [外呼详情] #外呼详情
	CdrObDetails []*ClinkCdrObDetailsResponseBodyDataCdrObDetails `json:"CdrObDetails,omitempty" xml:"CdrObDetails,omitempty" type:"Repeated"`
	// 请求 id
	//
	// example:
	//
	// xxx
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
}

func (s ClinkCdrObDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkCdrObDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkCdrObDetailsResponseBodyData) GetCdrObDetails() []*ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	return s.CdrObDetails
}

func (s *ClinkCdrObDetailsResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkCdrObDetailsResponseBodyData) SetCdrObDetails(v []*ClinkCdrObDetailsResponseBodyDataCdrObDetails) *ClinkCdrObDetailsResponseBodyData {
	s.CdrObDetails = v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyData) SetClinkRequestId(v string) *ClinkCdrObDetailsResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyData) Validate() error {
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

type ClinkCdrObDetailsResponseBodyDataCdrObDetails struct {
	// 客户接听时间
	//
	// example:
	//
	// 1717583928
	AnswerTime *int64 `json:"AnswerTime,omitempty" xml:"AnswerTime,omitempty"`
	// 通话时长
	//
	// example:
	//
	// 43
	BridgeDuration *int64 `json:"BridgeDuration,omitempty" xml:"BridgeDuration,omitempty"`
	// CallID
	//
	// example:
	//
	// 1d7acd5d-xxxx-xxxx-xxxx-0d3c08317140
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 通话类型，说明参考 通用字段 #通用字段
	//
	// example:
	//
	// 4
	CallType *int64 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 通话类型描述
	//
	// example:
	//
	// 示例值示例值示例值
	CallTypeDesc *string `json:"CallTypeDesc,omitempty" xml:"CallTypeDesc,omitempty"`
	// 客户响铃时间
	//
	// example:
	//
	// 1717583922
	CalleeRingingTime *int64 `json:"CalleeRingingTime,omitempty" xml:"CalleeRingingTime,omitempty"`
	// 座席名称
	//
	// example:
	//
	// xxx
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// 座席电话
	//
	// example:
	//
	// xxx
	ClientNumber *string `json:"ClientNumber,omitempty" xml:"ClientNumber,omitempty"`
	// 座席工号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 呼叫类型，说明参考 通用字段 #通用字段
	//
	// example:
	//
	// 208
	DetailCallType *int64 `json:"DetailCallType,omitempty" xml:"DetailCallType,omitempty"`
	// 呼叫类型描述
	//
	// example:
	//
	// 示例值示例值
	DetailCallTypeDesc *string `json:"DetailCallTypeDesc,omitempty" xml:"DetailCallTypeDesc,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 1717583950
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 通话ID
	//
	// example:
	//
	// medias_1-162046xxxx.12
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	// 呼叫结果
	//
	// example:
	//
	// 示例值示例值
	ObSipCause *string `json:"ObSipCause,omitempty" xml:"ObSipCause,omitempty"`
	// 被叫状态
	//
	// example:
	//
	// 示例值示例值
	ObSipCauseRaw *string `json:"ObSipCauseRaw,omitempty" xml:"ObSipCauseRaw,omitempty"`
	// 录音
	//
	// example:
	//
	// null
	RecordFile *string `json:"RecordFile,omitempty" xml:"RecordFile,omitempty"`
	// 呼叫情况
	//
	// example:
	//
	// 200
	SipCause *int64 `json:"SipCause,omitempty" xml:"SipCause,omitempty"`
	// 呼叫情况描述
	//
	// example:
	//
	// success
	SipCauseDesc *string `json:"SipCauseDesc,omitempty" xml:"SipCauseDesc,omitempty"`
	// 开始时间
	//
	// example:
	//
	// 1717583916
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 客户彩铃录音
	//
	// example:
	//
	// null
	TsiFile *string `json:"TsiFile,omitempty" xml:"TsiFile,omitempty"`
	// 通话唯一ID
	//
	// example:
	//
	// uniq_1-162046xxxx.12
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// WebRTCCallID
	//
	// example:
	//
	// ""
	WebrtcCallId *string `json:"WebrtcCallId,omitempty" xml:"WebrtcCallId,omitempty"`
}

func (s ClinkCdrObDetailsResponseBodyDataCdrObDetails) String() string {
	return dara.Prettify(s)
}

func (s ClinkCdrObDetailsResponseBodyDataCdrObDetails) GoString() string {
	return s.String()
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetAnswerTime() *int64 {
	return s.AnswerTime
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetBridgeDuration() *int64 {
	return s.BridgeDuration
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetCallId() *string {
	return s.CallId
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetCallType() *int64 {
	return s.CallType
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetCallTypeDesc() *string {
	return s.CallTypeDesc
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetCalleeRingingTime() *int64 {
	return s.CalleeRingingTime
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetClientName() *string {
	return s.ClientName
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetClientNumber() *string {
	return s.ClientNumber
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetCno() *string {
	return s.Cno
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetDetailCallType() *int64 {
	return s.DetailCallType
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetDetailCallTypeDesc() *string {
	return s.DetailCallTypeDesc
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetObSipCause() *string {
	return s.ObSipCause
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetObSipCauseRaw() *string {
	return s.ObSipCauseRaw
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetRecordFile() *string {
	return s.RecordFile
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetSipCause() *int64 {
	return s.SipCause
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetSipCauseDesc() *string {
	return s.SipCauseDesc
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetTsiFile() *string {
	return s.TsiFile
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetUniqueId() *string {
	return s.UniqueId
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) GetWebrtcCallId() *string {
	return s.WebrtcCallId
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetAnswerTime(v int64) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.AnswerTime = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetBridgeDuration(v int64) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.BridgeDuration = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetCallId(v string) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.CallId = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetCallType(v int64) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.CallType = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetCallTypeDesc(v string) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.CallTypeDesc = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetCalleeRingingTime(v int64) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.CalleeRingingTime = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetClientName(v string) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.ClientName = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetClientNumber(v string) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.ClientNumber = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetCno(v string) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.Cno = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetDetailCallType(v int64) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.DetailCallType = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetDetailCallTypeDesc(v string) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.DetailCallTypeDesc = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetEndTime(v int64) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.EndTime = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetMainUniqueId(v string) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetObSipCause(v string) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.ObSipCause = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetObSipCauseRaw(v string) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.ObSipCauseRaw = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetRecordFile(v string) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.RecordFile = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetSipCause(v int64) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.SipCause = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetSipCauseDesc(v string) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.SipCauseDesc = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetStartTime(v int64) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.StartTime = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetTsiFile(v string) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.TsiFile = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetUniqueId(v string) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.UniqueId = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) SetWebrtcCallId(v string) *ClinkCdrObDetailsResponseBodyDataCdrObDetails {
	s.WebrtcCallId = &v
	return s
}

func (s *ClinkCdrObDetailsResponseBodyDataCdrObDetails) Validate() error {
	return dara.Validate(s)
}
