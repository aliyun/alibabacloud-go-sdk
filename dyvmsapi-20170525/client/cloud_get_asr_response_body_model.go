// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetAsrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudGetAsrResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudGetAsrResponseBody
	GetCode() *string
	SetData(v *CloudGetAsrResponseBodyData) *CloudGetAsrResponseBody
	GetData() *CloudGetAsrResponseBodyData
	SetMessage(v string) *CloudGetAsrResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudGetAsrResponseBody
	GetRequestId() *string
}

type CloudGetAsrResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudGetAsrResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudGetAsrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAsrResponseBody) GoString() string {
	return s.String()
}

func (s *CloudGetAsrResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudGetAsrResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudGetAsrResponseBody) GetData() *CloudGetAsrResponseBodyData {
	return s.Data
}

func (s *CloudGetAsrResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudGetAsrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudGetAsrResponseBody) SetAccessDeniedDetail(v string) *CloudGetAsrResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudGetAsrResponseBody) SetCode(v string) *CloudGetAsrResponseBody {
	s.Code = &v
	return s
}

func (s *CloudGetAsrResponseBody) SetData(v *CloudGetAsrResponseBodyData) *CloudGetAsrResponseBody {
	s.Data = v
	return s
}

func (s *CloudGetAsrResponseBody) SetMessage(v string) *CloudGetAsrResponseBody {
	s.Message = &v
	return s
}

func (s *CloudGetAsrResponseBody) SetRequestId(v string) *CloudGetAsrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudGetAsrResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudGetAsrResponseBodyData struct {
	// 企业编号
	//
	// example:
	//
	// 7000002
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 转写类型 1：混音 2：分轨
	//
	// example:
	//
	// 1
	MonitorType *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// 转写文件个数
	//
	// example:
	//
	// 1
	RecordFileCount *string `json:"RecordFileCount,omitempty" xml:"RecordFileCount,omitempty"`
	// 第一侧转写文本结果
	RecordFileText1In *CloudGetAsrResponseBodyDataRecordFileText1In `json:"RecordFileText1In,omitempty" xml:"RecordFileText1In,omitempty" type:"Struct"`
	// 第二侧转写文本结果
	RecordFileText1Out *CloudGetAsrResponseBodyDataRecordFileText1Out `json:"RecordFileText1Out,omitempty" xml:"RecordFileText1Out,omitempty" type:"Struct"`
	// 转写结果, 当all=true时返回
	Result []*CloudGetAsrResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s CloudGetAsrResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAsrResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudGetAsrResponseBodyData) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudGetAsrResponseBodyData) GetMonitorType() *string {
	return s.MonitorType
}

func (s *CloudGetAsrResponseBodyData) GetRecordFileCount() *string {
	return s.RecordFileCount
}

func (s *CloudGetAsrResponseBodyData) GetRecordFileText1In() *CloudGetAsrResponseBodyDataRecordFileText1In {
	return s.RecordFileText1In
}

func (s *CloudGetAsrResponseBodyData) GetRecordFileText1Out() *CloudGetAsrResponseBodyDataRecordFileText1Out {
	return s.RecordFileText1Out
}

func (s *CloudGetAsrResponseBodyData) GetResult() []*CloudGetAsrResponseBodyDataResult {
	return s.Result
}

func (s *CloudGetAsrResponseBodyData) SetEnterpriseId(v string) *CloudGetAsrResponseBodyData {
	s.EnterpriseId = &v
	return s
}

func (s *CloudGetAsrResponseBodyData) SetMonitorType(v string) *CloudGetAsrResponseBodyData {
	s.MonitorType = &v
	return s
}

func (s *CloudGetAsrResponseBodyData) SetRecordFileCount(v string) *CloudGetAsrResponseBodyData {
	s.RecordFileCount = &v
	return s
}

func (s *CloudGetAsrResponseBodyData) SetRecordFileText1In(v *CloudGetAsrResponseBodyDataRecordFileText1In) *CloudGetAsrResponseBodyData {
	s.RecordFileText1In = v
	return s
}

func (s *CloudGetAsrResponseBodyData) SetRecordFileText1Out(v *CloudGetAsrResponseBodyDataRecordFileText1Out) *CloudGetAsrResponseBodyData {
	s.RecordFileText1Out = v
	return s
}

func (s *CloudGetAsrResponseBodyData) SetResult(v []*CloudGetAsrResponseBodyDataResult) *CloudGetAsrResponseBodyData {
	s.Result = v
	return s
}

func (s *CloudGetAsrResponseBodyData) Validate() error {
	if s.RecordFileText1In != nil {
		if err := s.RecordFileText1In.Validate(); err != nil {
			return err
		}
	}
	if s.RecordFileText1Out != nil {
		if err := s.RecordFileText1Out.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudGetAsrResponseBodyDataRecordFileText1In struct {
	// example:
	//
	// 33340
	BizDuration *int64 `json:"BizDuration,omitempty" xml:"BizDuration,omitempty"`
	// 错误信息
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// id
	//
	// example:
	//
	// 8f2ff56a7cfe458c942595b2f271d715
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 转写结果
	Result []*CloudGetAsrResponseBodyDataRecordFileText1InResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// 状态
	//
	// example:
	//
	// 示例值示例值
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 状态码
	//
	// example:
	//
	// SUCCEED
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s CloudGetAsrResponseBodyDataRecordFileText1In) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAsrResponseBodyDataRecordFileText1In) GoString() string {
	return s.String()
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1In) GetBizDuration() *int64 {
	return s.BizDuration
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1In) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1In) GetId() *string {
	return s.Id
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1In) GetResult() []*CloudGetAsrResponseBodyDataRecordFileText1InResult {
	return s.Result
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1In) GetStatus() *string {
	return s.Status
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1In) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1In) SetBizDuration(v int64) *CloudGetAsrResponseBodyDataRecordFileText1In {
	s.BizDuration = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1In) SetErrorMessage(v string) *CloudGetAsrResponseBodyDataRecordFileText1In {
	s.ErrorMessage = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1In) SetId(v string) *CloudGetAsrResponseBodyDataRecordFileText1In {
	s.Id = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1In) SetResult(v []*CloudGetAsrResponseBodyDataRecordFileText1InResult) *CloudGetAsrResponseBodyDataRecordFileText1In {
	s.Result = v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1In) SetStatus(v string) *CloudGetAsrResponseBodyDataRecordFileText1In {
	s.Status = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1In) SetStatusCode(v int64) *CloudGetAsrResponseBodyDataRecordFileText1In {
	s.StatusCode = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1In) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudGetAsrResponseBodyDataRecordFileText1InResult struct {
	// 开始时间
	//
	// example:
	//
	// 30140
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// 通道id
	//
	// example:
	//
	// 0
	ChannelId *int64 `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// 情感值
	//
	// example:
	//
	// 6.0
	EmotionValue *float64 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 31805
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 沉默时间
	//
	// example:
	//
	// 1
	SilenceDuration *int64 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	// 语速
	//
	// example:
	//
	// 108
	SpeechRate *int64 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// 转写文本
	//
	// example:
	//
	// 示例值示例值
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s CloudGetAsrResponseBodyDataRecordFileText1InResult) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAsrResponseBodyDataRecordFileText1InResult) GoString() string {
	return s.String()
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1InResult) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1InResult) GetChannelId() *int64 {
	return s.ChannelId
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1InResult) GetEmotionValue() *float64 {
	return s.EmotionValue
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1InResult) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1InResult) GetSilenceDuration() *int64 {
	return s.SilenceDuration
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1InResult) GetSpeechRate() *int64 {
	return s.SpeechRate
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1InResult) GetText() *string {
	return s.Text
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1InResult) SetBeginTime(v int64) *CloudGetAsrResponseBodyDataRecordFileText1InResult {
	s.BeginTime = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1InResult) SetChannelId(v int64) *CloudGetAsrResponseBodyDataRecordFileText1InResult {
	s.ChannelId = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1InResult) SetEmotionValue(v float64) *CloudGetAsrResponseBodyDataRecordFileText1InResult {
	s.EmotionValue = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1InResult) SetEndTime(v int64) *CloudGetAsrResponseBodyDataRecordFileText1InResult {
	s.EndTime = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1InResult) SetSilenceDuration(v int64) *CloudGetAsrResponseBodyDataRecordFileText1InResult {
	s.SilenceDuration = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1InResult) SetSpeechRate(v int64) *CloudGetAsrResponseBodyDataRecordFileText1InResult {
	s.SpeechRate = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1InResult) SetText(v string) *CloudGetAsrResponseBodyDataRecordFileText1InResult {
	s.Text = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1InResult) Validate() error {
	return dara.Validate(s)
}

type CloudGetAsrResponseBodyDataRecordFileText1Out struct {
	// example:
	//
	// 33340
	BizDuration *int64 `json:"BizDuration,omitempty" xml:"BizDuration,omitempty"`
	// 错误信息
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// id
	//
	// example:
	//
	// 145a13227e3249258cd3b5232a56d634
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 转写结果
	Result []*CloudGetAsrResponseBodyDataRecordFileText1OutResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// 状态
	//
	// example:
	//
	// SUCCEED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 状态码
	//
	// example:
	//
	// 0
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s CloudGetAsrResponseBodyDataRecordFileText1Out) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAsrResponseBodyDataRecordFileText1Out) GoString() string {
	return s.String()
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1Out) GetBizDuration() *int64 {
	return s.BizDuration
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1Out) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1Out) GetId() *string {
	return s.Id
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1Out) GetResult() []*CloudGetAsrResponseBodyDataRecordFileText1OutResult {
	return s.Result
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1Out) GetStatus() *string {
	return s.Status
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1Out) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1Out) SetBizDuration(v int64) *CloudGetAsrResponseBodyDataRecordFileText1Out {
	s.BizDuration = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1Out) SetErrorMessage(v string) *CloudGetAsrResponseBodyDataRecordFileText1Out {
	s.ErrorMessage = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1Out) SetId(v string) *CloudGetAsrResponseBodyDataRecordFileText1Out {
	s.Id = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1Out) SetResult(v []*CloudGetAsrResponseBodyDataRecordFileText1OutResult) *CloudGetAsrResponseBodyDataRecordFileText1Out {
	s.Result = v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1Out) SetStatus(v string) *CloudGetAsrResponseBodyDataRecordFileText1Out {
	s.Status = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1Out) SetStatusCode(v int64) *CloudGetAsrResponseBodyDataRecordFileText1Out {
	s.StatusCode = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1Out) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudGetAsrResponseBodyDataRecordFileText1OutResult struct {
	// 开始时间
	//
	// example:
	//
	// 30550
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// 通道id
	//
	// example:
	//
	// 0
	ChannelId *int64 `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// 情感值
	//
	// example:
	//
	// 6.0
	EmotionValue *float64 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 31515
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 沉默时间
	//
	// example:
	//
	// 0
	SilenceDuration *int64 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	// 语速
	//
	// example:
	//
	// 186
	SpeechRate *int64 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// 转写文本
	//
	// example:
	//
	// 示例值示例值示例值
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s CloudGetAsrResponseBodyDataRecordFileText1OutResult) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAsrResponseBodyDataRecordFileText1OutResult) GoString() string {
	return s.String()
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1OutResult) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1OutResult) GetChannelId() *int64 {
	return s.ChannelId
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1OutResult) GetEmotionValue() *float64 {
	return s.EmotionValue
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1OutResult) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1OutResult) GetSilenceDuration() *int64 {
	return s.SilenceDuration
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1OutResult) GetSpeechRate() *int64 {
	return s.SpeechRate
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1OutResult) GetText() *string {
	return s.Text
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1OutResult) SetBeginTime(v int64) *CloudGetAsrResponseBodyDataRecordFileText1OutResult {
	s.BeginTime = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1OutResult) SetChannelId(v int64) *CloudGetAsrResponseBodyDataRecordFileText1OutResult {
	s.ChannelId = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1OutResult) SetEmotionValue(v float64) *CloudGetAsrResponseBodyDataRecordFileText1OutResult {
	s.EmotionValue = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1OutResult) SetEndTime(v int64) *CloudGetAsrResponseBodyDataRecordFileText1OutResult {
	s.EndTime = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1OutResult) SetSilenceDuration(v int64) *CloudGetAsrResponseBodyDataRecordFileText1OutResult {
	s.SilenceDuration = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1OutResult) SetSpeechRate(v int64) *CloudGetAsrResponseBodyDataRecordFileText1OutResult {
	s.SpeechRate = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1OutResult) SetText(v string) *CloudGetAsrResponseBodyDataRecordFileText1OutResult {
	s.Text = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataRecordFileText1OutResult) Validate() error {
	return dara.Validate(s)
}

type CloudGetAsrResponseBodyDataResult struct {
	// 开始时间
	//
	// example:
	//
	// 4350
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// 通道id
	//
	// example:
	//
	// 32
	ChannelId *int64 `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// 情感值
	//
	// example:
	//
	// 7.0
	EmotionValue *float64 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 10445
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 语音文本所属方；in: 第一侧；out: 第二侧
	//
	// example:
	//
	// out
	Side *string `json:"Side,omitempty" xml:"Side,omitempty"`
	// 沉默时间
	//
	// example:
	//
	// 0
	SilenceDuration *int64 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	// 语速
	//
	// example:
	//
	// 255
	SpeechRate *int64 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// 转写文本
	//
	// example:
	//
	// 示例值示例值
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s CloudGetAsrResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAsrResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *CloudGetAsrResponseBodyDataResult) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *CloudGetAsrResponseBodyDataResult) GetChannelId() *int64 {
	return s.ChannelId
}

func (s *CloudGetAsrResponseBodyDataResult) GetEmotionValue() *float64 {
	return s.EmotionValue
}

func (s *CloudGetAsrResponseBodyDataResult) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CloudGetAsrResponseBodyDataResult) GetSide() *string {
	return s.Side
}

func (s *CloudGetAsrResponseBodyDataResult) GetSilenceDuration() *int64 {
	return s.SilenceDuration
}

func (s *CloudGetAsrResponseBodyDataResult) GetSpeechRate() *int64 {
	return s.SpeechRate
}

func (s *CloudGetAsrResponseBodyDataResult) GetText() *string {
	return s.Text
}

func (s *CloudGetAsrResponseBodyDataResult) SetBeginTime(v int64) *CloudGetAsrResponseBodyDataResult {
	s.BeginTime = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataResult) SetChannelId(v int64) *CloudGetAsrResponseBodyDataResult {
	s.ChannelId = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataResult) SetEmotionValue(v float64) *CloudGetAsrResponseBodyDataResult {
	s.EmotionValue = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataResult) SetEndTime(v int64) *CloudGetAsrResponseBodyDataResult {
	s.EndTime = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataResult) SetSide(v string) *CloudGetAsrResponseBodyDataResult {
	s.Side = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataResult) SetSilenceDuration(v int64) *CloudGetAsrResponseBodyDataResult {
	s.SilenceDuration = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataResult) SetSpeechRate(v int64) *CloudGetAsrResponseBodyDataResult {
	s.SpeechRate = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataResult) SetText(v string) *CloudGetAsrResponseBodyDataResult {
	s.Text = &v
	return s
}

func (s *CloudGetAsrResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
