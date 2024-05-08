// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AsrRealtimeRequest struct {
	// example:
	//
	// 4a29b426-742f-4078-8386-79b440b25***
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CustomizationId *string `json:"CustomizationId,omitempty" xml:"CustomizationId,omitempty"`
	// example:
	//
	// false
	Disfluency *bool `json:"Disfluency,omitempty" xml:"Disfluency,omitempty"`
	// example:
	//
	// false
	EnableIgnoreSentenceTimeout *bool `json:"EnableIgnoreSentenceTimeout,omitempty" xml:"EnableIgnoreSentenceTimeout,omitempty"`
	// example:
	//
	// false
	EnableIntermediateResult *bool `json:"EnableIntermediateResult,omitempty" xml:"EnableIntermediateResult,omitempty"`
	// example:
	//
	// false
	EnableInverseTextNormalization *bool `json:"EnableInverseTextNormalization,omitempty" xml:"EnableInverseTextNormalization,omitempty"`
	// example:
	//
	// false
	EnablePunctuationPrediction *bool `json:"EnablePunctuationPrediction,omitempty" xml:"EnablePunctuationPrediction,omitempty"`
	// example:
	//
	// false
	EnableSemanticSentenceDetection *bool `json:"EnableSemanticSentenceDetection,omitempty" xml:"EnableSemanticSentenceDetection,omitempty"`
	// example:
	//
	// false
	EnableWords *bool `json:"EnableWords,omitempty" xml:"EnableWords,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/bmw-prod/0574ee2e-f494-45a5-820f-63aee***.wav
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// PCM
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// 800
	MaxSentenceSilence *int64 `json:"MaxSentenceSilence,omitempty" xml:"MaxSentenceSilence,omitempty"`
	// example:
	//
	// 16000
	SampleRate *int64 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// example:
	//
	// 0.3
	SpeechNoiseThreshold *float32 `json:"SpeechNoiseThreshold,omitempty" xml:"SpeechNoiseThreshold,omitempty"`
	VocabularyId         *string  `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
}

func (s AsrRealtimeRequest) String() string {
	return tea.Prettify(s)
}

func (s AsrRealtimeRequest) GoString() string {
	return s.String()
}

func (s *AsrRealtimeRequest) SetAppId(v string) *AsrRealtimeRequest {
	s.AppId = &v
	return s
}

func (s *AsrRealtimeRequest) SetCustomizationId(v string) *AsrRealtimeRequest {
	s.CustomizationId = &v
	return s
}

func (s *AsrRealtimeRequest) SetDisfluency(v bool) *AsrRealtimeRequest {
	s.Disfluency = &v
	return s
}

func (s *AsrRealtimeRequest) SetEnableIgnoreSentenceTimeout(v bool) *AsrRealtimeRequest {
	s.EnableIgnoreSentenceTimeout = &v
	return s
}

func (s *AsrRealtimeRequest) SetEnableIntermediateResult(v bool) *AsrRealtimeRequest {
	s.EnableIntermediateResult = &v
	return s
}

func (s *AsrRealtimeRequest) SetEnableInverseTextNormalization(v bool) *AsrRealtimeRequest {
	s.EnableInverseTextNormalization = &v
	return s
}

func (s *AsrRealtimeRequest) SetEnablePunctuationPrediction(v bool) *AsrRealtimeRequest {
	s.EnablePunctuationPrediction = &v
	return s
}

func (s *AsrRealtimeRequest) SetEnableSemanticSentenceDetection(v bool) *AsrRealtimeRequest {
	s.EnableSemanticSentenceDetection = &v
	return s
}

func (s *AsrRealtimeRequest) SetEnableWords(v bool) *AsrRealtimeRequest {
	s.EnableWords = &v
	return s
}

func (s *AsrRealtimeRequest) SetFileUrl(v string) *AsrRealtimeRequest {
	s.FileUrl = &v
	return s
}

func (s *AsrRealtimeRequest) SetFormat(v string) *AsrRealtimeRequest {
	s.Format = &v
	return s
}

func (s *AsrRealtimeRequest) SetMaxSentenceSilence(v int64) *AsrRealtimeRequest {
	s.MaxSentenceSilence = &v
	return s
}

func (s *AsrRealtimeRequest) SetSampleRate(v int64) *AsrRealtimeRequest {
	s.SampleRate = &v
	return s
}

func (s *AsrRealtimeRequest) SetSpeechNoiseThreshold(v float32) *AsrRealtimeRequest {
	s.SpeechNoiseThreshold = &v
	return s
}

func (s *AsrRealtimeRequest) SetVocabularyId(v string) *AsrRealtimeRequest {
	s.VocabularyId = &v
	return s
}

type AsrRealtimeResponseBody struct {
	// example:
	//
	// OK
	Code *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AsrRealtimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 84118BF0-56F7-54D2-8C1A-35BBBBD50***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AsrRealtimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AsrRealtimeResponseBody) GoString() string {
	return s.String()
}

func (s *AsrRealtimeResponseBody) SetCode(v int32) *AsrRealtimeResponseBody {
	s.Code = &v
	return s
}

func (s *AsrRealtimeResponseBody) SetData(v *AsrRealtimeResponseBodyData) *AsrRealtimeResponseBody {
	s.Data = v
	return s
}

func (s *AsrRealtimeResponseBody) SetHttpCode(v int32) *AsrRealtimeResponseBody {
	s.HttpCode = &v
	return s
}

func (s *AsrRealtimeResponseBody) SetMessage(v string) *AsrRealtimeResponseBody {
	s.Message = &v
	return s
}

func (s *AsrRealtimeResponseBody) SetRequestId(v string) *AsrRealtimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsrRealtimeResponseBody) SetSuccess(v bool) *AsrRealtimeResponseBody {
	s.Success = &v
	return s
}

type AsrRealtimeResponseBodyData struct {
	// example:
	//
	// 1649952000000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// OK
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 0.78
	Confidence *float64 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// f3bd31c0-0001-4b4b-977d-7cfa64b5***
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// default
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 368cfa55-2364-4d79-aeb4-c0956c4a4***
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 1638243477
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s AsrRealtimeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AsrRealtimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *AsrRealtimeResponseBodyData) SetBeginTime(v int64) *AsrRealtimeResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *AsrRealtimeResponseBodyData) SetCode(v int32) *AsrRealtimeResponseBodyData {
	s.Code = &v
	return s
}

func (s *AsrRealtimeResponseBodyData) SetConfidence(v float64) *AsrRealtimeResponseBodyData {
	s.Confidence = &v
	return s
}

func (s *AsrRealtimeResponseBodyData) SetId(v string) *AsrRealtimeResponseBodyData {
	s.Id = &v
	return s
}

func (s *AsrRealtimeResponseBodyData) SetMessage(v string) *AsrRealtimeResponseBodyData {
	s.Message = &v
	return s
}

func (s *AsrRealtimeResponseBodyData) SetName(v string) *AsrRealtimeResponseBodyData {
	s.Name = &v
	return s
}

func (s *AsrRealtimeResponseBodyData) SetResult(v string) *AsrRealtimeResponseBodyData {
	s.Result = &v
	return s
}

func (s *AsrRealtimeResponseBodyData) SetStatus(v string) *AsrRealtimeResponseBodyData {
	s.Status = &v
	return s
}

func (s *AsrRealtimeResponseBodyData) SetTaskId(v string) *AsrRealtimeResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AsrRealtimeResponseBodyData) SetTime(v int64) *AsrRealtimeResponseBodyData {
	s.Time = &v
	return s
}

type AsrRealtimeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsrRealtimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsrRealtimeResponse) String() string {
	return tea.Prettify(s)
}

func (s AsrRealtimeResponse) GoString() string {
	return s.String()
}

func (s *AsrRealtimeResponse) SetHeaders(v map[string]*string) *AsrRealtimeResponse {
	s.Headers = v
	return s
}

func (s *AsrRealtimeResponse) SetStatusCode(v int32) *AsrRealtimeResponse {
	s.StatusCode = &v
	return s
}

func (s *AsrRealtimeResponse) SetBody(v *AsrRealtimeResponseBody) *AsrRealtimeResponse {
	s.Body = v
	return s
}

type AsrSentenceRequest struct {
	AsrRequest *AsrSentenceRequestAsrRequest `json:"AsrRequest,omitempty" xml:"AsrRequest,omitempty" type:"Struct"`
}

func (s AsrSentenceRequest) String() string {
	return tea.Prettify(s)
}

func (s AsrSentenceRequest) GoString() string {
	return s.String()
}

func (s *AsrSentenceRequest) SetAsrRequest(v *AsrSentenceRequestAsrRequest) *AsrSentenceRequest {
	s.AsrRequest = v
	return s
}

type AsrSentenceRequestAsrRequest struct {
	// example:
	//
	// d61be709-49d2-4cf1-b219-cd6181f72db8
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// false
	EnableInverseTextNormalization *bool `json:"EnableInverseTextNormalization,omitempty" xml:"EnableInverseTextNormalization,omitempty"`
	// example:
	//
	// false
	EnablePunctuationPrediction *bool `json:"EnablePunctuationPrediction,omitempty" xml:"EnablePunctuationPrediction,omitempty"`
	// example:
	//
	// false
	EnableVoiceDetection *bool `json:"EnableVoiceDetection,omitempty" xml:"EnableVoiceDetection,omitempty"`
	// example:
	//
	// http://shuanglu-record-finance.oss-cn-shanghai.aliyuncs.com/record/4x5avhil/047730_30307_0/2022-02-12-10-20****
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// PCM
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// 16000
	SampleRate *int32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
}

func (s AsrSentenceRequestAsrRequest) String() string {
	return tea.Prettify(s)
}

func (s AsrSentenceRequestAsrRequest) GoString() string {
	return s.String()
}

func (s *AsrSentenceRequestAsrRequest) SetAppId(v string) *AsrSentenceRequestAsrRequest {
	s.AppId = &v
	return s
}

func (s *AsrSentenceRequestAsrRequest) SetEnableInverseTextNormalization(v bool) *AsrSentenceRequestAsrRequest {
	s.EnableInverseTextNormalization = &v
	return s
}

func (s *AsrSentenceRequestAsrRequest) SetEnablePunctuationPrediction(v bool) *AsrSentenceRequestAsrRequest {
	s.EnablePunctuationPrediction = &v
	return s
}

func (s *AsrSentenceRequestAsrRequest) SetEnableVoiceDetection(v bool) *AsrSentenceRequestAsrRequest {
	s.EnableVoiceDetection = &v
	return s
}

func (s *AsrSentenceRequestAsrRequest) SetFileUrl(v string) *AsrSentenceRequestAsrRequest {
	s.FileUrl = &v
	return s
}

func (s *AsrSentenceRequestAsrRequest) SetFormat(v string) *AsrSentenceRequestAsrRequest {
	s.Format = &v
	return s
}

func (s *AsrSentenceRequestAsrRequest) SetSampleRate(v int32) *AsrSentenceRequestAsrRequest {
	s.SampleRate = &v
	return s
}

type AsrSentenceShrinkRequest struct {
	AsrRequestShrink *string `json:"AsrRequest,omitempty" xml:"AsrRequest,omitempty"`
}

func (s AsrSentenceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AsrSentenceShrinkRequest) GoString() string {
	return s.String()
}

func (s *AsrSentenceShrinkRequest) SetAsrRequestShrink(v string) *AsrSentenceShrinkRequest {
	s.AsrRequestShrink = &v
	return s
}

type AsrSentenceResponseBody struct {
	// example:
	//
	// OK
	Code *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AsrSentenceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AsrSentenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AsrSentenceResponseBody) GoString() string {
	return s.String()
}

func (s *AsrSentenceResponseBody) SetCode(v int32) *AsrSentenceResponseBody {
	s.Code = &v
	return s
}

func (s *AsrSentenceResponseBody) SetData(v *AsrSentenceResponseBodyData) *AsrSentenceResponseBody {
	s.Data = v
	return s
}

func (s *AsrSentenceResponseBody) SetHttpCode(v int32) *AsrSentenceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *AsrSentenceResponseBody) SetMessage(v string) *AsrSentenceResponseBody {
	s.Message = &v
	return s
}

func (s *AsrSentenceResponseBody) SetRequestId(v string) *AsrSentenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsrSentenceResponseBody) SetSuccess(v bool) *AsrSentenceResponseBody {
	s.Success = &v
	return s
}

type AsrSentenceResponseBodyData struct {
	// example:
	//
	// 20000000
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Result  *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 368cfa55-2364-4d79-aeb4-c0956c4a45cd
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AsrSentenceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AsrSentenceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AsrSentenceResponseBodyData) SetCode(v int32) *AsrSentenceResponseBodyData {
	s.Code = &v
	return s
}

func (s *AsrSentenceResponseBodyData) SetId(v string) *AsrSentenceResponseBodyData {
	s.Id = &v
	return s
}

func (s *AsrSentenceResponseBodyData) SetMessage(v string) *AsrSentenceResponseBodyData {
	s.Message = &v
	return s
}

func (s *AsrSentenceResponseBodyData) SetName(v string) *AsrSentenceResponseBodyData {
	s.Name = &v
	return s
}

func (s *AsrSentenceResponseBodyData) SetResult(v string) *AsrSentenceResponseBodyData {
	s.Result = &v
	return s
}

func (s *AsrSentenceResponseBodyData) SetTaskId(v string) *AsrSentenceResponseBodyData {
	s.TaskId = &v
	return s
}

type AsrSentenceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsrSentenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsrSentenceResponse) String() string {
	return tea.Prettify(s)
}

func (s AsrSentenceResponse) GoString() string {
	return s.String()
}

func (s *AsrSentenceResponse) SetHeaders(v map[string]*string) *AsrSentenceResponse {
	s.Headers = v
	return s
}

func (s *AsrSentenceResponse) SetStatusCode(v int32) *AsrSentenceResponse {
	s.StatusCode = &v
	return s
}

func (s *AsrSentenceResponse) SetBody(v *AsrSentenceResponseBody) *AsrSentenceResponse {
	s.Body = v
	return s
}

type AsrTaskRequest struct {
	Request *AsrTaskRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s AsrTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s AsrTaskRequest) GoString() string {
	return s.String()
}

func (s *AsrTaskRequest) SetRequest(v *AsrTaskRequestRequest) *AsrTaskRequest {
	s.Request = v
	return s
}

type AsrTaskRequestRequest struct {
	// example:
	//
	// d9ee5df9-20bf-47bf-987a-76b26984b***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// d9ee5df9-20bf-47bf-987a-76b26984b***
	AsrTaskId *string `json:"AsrTaskId,omitempty" xml:"AsrTaskId,omitempty"`
	// example:
	//
	// START
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 662027426755***
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// example:
	//
	// 1656388156399
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s AsrTaskRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s AsrTaskRequestRequest) GoString() string {
	return s.String()
}

func (s *AsrTaskRequestRequest) SetAppId(v string) *AsrTaskRequestRequest {
	s.AppId = &v
	return s
}

func (s *AsrTaskRequestRequest) SetAsrTaskId(v string) *AsrTaskRequestRequest {
	s.AsrTaskId = &v
	return s
}

func (s *AsrTaskRequestRequest) SetEvent(v string) *AsrTaskRequestRequest {
	s.Event = &v
	return s
}

func (s *AsrTaskRequestRequest) SetRoomId(v string) *AsrTaskRequestRequest {
	s.RoomId = &v
	return s
}

func (s *AsrTaskRequestRequest) SetTimestamp(v int64) *AsrTaskRequestRequest {
	s.Timestamp = &v
	return s
}

type AsrTaskShrinkRequest struct {
	RequestShrink *string `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s AsrTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AsrTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *AsrTaskShrinkRequest) SetRequestShrink(v string) *AsrTaskShrinkRequest {
	s.RequestShrink = &v
	return s
}

type AsrTaskResponseBody struct {
	// example:
	//
	// OK
	Code *int32                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AsrTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DF4B0162-A5E0-5F85-BEFD-CAC36E876***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AsrTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AsrTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AsrTaskResponseBody) SetCode(v int32) *AsrTaskResponseBody {
	s.Code = &v
	return s
}

func (s *AsrTaskResponseBody) SetData(v *AsrTaskResponseBodyData) *AsrTaskResponseBody {
	s.Data = v
	return s
}

func (s *AsrTaskResponseBody) SetHttpCode(v int32) *AsrTaskResponseBody {
	s.HttpCode = &v
	return s
}

func (s *AsrTaskResponseBody) SetMessage(v string) *AsrTaskResponseBody {
	s.Message = &v
	return s
}

func (s *AsrTaskResponseBody) SetRequestId(v string) *AsrTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsrTaskResponseBody) SetSuccess(v bool) *AsrTaskResponseBody {
	s.Success = &v
	return s
}

type AsrTaskResponseBodyData struct {
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AsrTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AsrTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *AsrTaskResponseBodyData) SetResult(v string) *AsrTaskResponseBodyData {
	s.Result = &v
	return s
}

type AsrTaskResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsrTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsrTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s AsrTaskResponse) GoString() string {
	return s.String()
}

func (s *AsrTaskResponse) SetHeaders(v map[string]*string) *AsrTaskResponse {
	s.Headers = v
	return s
}

func (s *AsrTaskResponse) SetStatusCode(v int32) *AsrTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AsrTaskResponse) SetBody(v *AsrTaskResponseBody) *AsrTaskResponse {
	s.Body = v
	return s
}

type AssociateRoomRequest struct {
	// example:
	//
	// 5bbfb884-1186-4d48-906b-88d586770f6b
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// {"version":"1.0.0"}
	ClientBaseParam *string `json:"ClientBaseParam,omitempty" xml:"ClientBaseParam,omitempty"`
	// example:
	//
	// 1.0.002
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// 5bbfb884-1186-4d48-906b-88d586770f6b
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// example:
	//
	// 5500707344661
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s AssociateRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateRoomRequest) GoString() string {
	return s.String()
}

func (s *AssociateRoomRequest) SetAppId(v string) *AssociateRoomRequest {
	s.AppId = &v
	return s
}

func (s *AssociateRoomRequest) SetClientBaseParam(v string) *AssociateRoomRequest {
	s.ClientBaseParam = &v
	return s
}

func (s *AssociateRoomRequest) SetClientVersion(v string) *AssociateRoomRequest {
	s.ClientVersion = &v
	return s
}

func (s *AssociateRoomRequest) SetDepartmentId(v string) *AssociateRoomRequest {
	s.DepartmentId = &v
	return s
}

func (s *AssociateRoomRequest) SetRoomId(v string) *AssociateRoomRequest {
	s.RoomId = &v
	return s
}

type AssociateRoomResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {\"Name\": u\"\u4ee3\u7406\u4eba\", \"CreatedAt\": \"2021-11-11T15:27:39.449+08:00\", \"Channel\": \"063756\", \"Id\": \"5ead2d7f-9e2c-4521-bac4-e37bd44b6a56\"}
	Data   *string                            `json:"Data,omitempty" xml:"Data,omitempty"`
	Errors []*AssociateRoomResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// 84118BF0-56F7-54D2-8C1A-35BBBBD5016B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssociateRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateRoomResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateRoomResponseBody) SetCode(v string) *AssociateRoomResponseBody {
	s.Code = &v
	return s
}

func (s *AssociateRoomResponseBody) SetData(v string) *AssociateRoomResponseBody {
	s.Data = &v
	return s
}

func (s *AssociateRoomResponseBody) SetErrors(v []*AssociateRoomResponseBodyErrors) *AssociateRoomResponseBody {
	s.Errors = v
	return s
}

func (s *AssociateRoomResponseBody) SetHttpCode(v int32) *AssociateRoomResponseBody {
	s.HttpCode = &v
	return s
}

func (s *AssociateRoomResponseBody) SetMessage(v string) *AssociateRoomResponseBody {
	s.Message = &v
	return s
}

func (s *AssociateRoomResponseBody) SetRequestId(v string) *AssociateRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateRoomResponseBody) SetSuccess(v bool) *AssociateRoomResponseBody {
	s.Success = &v
	return s
}

type AssociateRoomResponseBodyErrors struct {
	// example:
	//
	// -
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AssociateRoomResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s AssociateRoomResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *AssociateRoomResponseBodyErrors) SetField(v string) *AssociateRoomResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *AssociateRoomResponseBodyErrors) SetMessage(v string) *AssociateRoomResponseBodyErrors {
	s.Message = &v
	return s
}

type AssociateRoomResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateRoomResponse) GoString() string {
	return s.String()
}

func (s *AssociateRoomResponse) SetHeaders(v map[string]*string) *AssociateRoomResponse {
	s.Headers = v
	return s
}

func (s *AssociateRoomResponse) SetStatusCode(v int32) *AssociateRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateRoomResponse) SetBody(v *AssociateRoomResponseBody) *AssociateRoomResponse {
	s.Body = v
	return s
}

type CreateAppRequest struct {
	// example:
	//
	// 4367c30a-c686-4bb2-a45d-5affb87****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 4367c30a-c686-4bb2-a45d-5affb87f7aca
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// com.test.app
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetClientToken(v string) *CreateAppRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppRequest) SetDepartmentId(v string) *CreateAppRequest {
	s.DepartmentId = &v
	return s
}

func (s *CreateAppRequest) SetName(v string) *CreateAppRequest {
	s.Name = &v
	return s
}

func (s *CreateAppRequest) SetPackageName(v string) *CreateAppRequest {
	s.PackageName = &v
	return s
}

type CreateAppResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code    *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *CreateAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetCode(v string) *CreateAppResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppResponseBody) SetData(v *CreateAppResponseBodyData) *CreateAppResponseBody {
	s.Data = v
	return s
}

func (s *CreateAppResponseBody) SetMessage(v string) *CreateAppResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppResponseBodyData struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyData) SetCreatedAt(v string) *CreateAppResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateAppResponseBodyData) SetCreatorName(v string) *CreateAppResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *CreateAppResponseBodyData) SetDisabled(v bool) *CreateAppResponseBodyData {
	s.Disabled = &v
	return s
}

func (s *CreateAppResponseBodyData) SetId(v string) *CreateAppResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateAppResponseBodyData) SetName(v string) *CreateAppResponseBodyData {
	s.Name = &v
	return s
}

type CreateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetStatusCode(v int32) *CreateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateDepartmentRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Label       *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentRequest) GoString() string {
	return s.String()
}

func (s *CreateDepartmentRequest) SetClientToken(v string) *CreateDepartmentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDepartmentRequest) SetDescription(v string) *CreateDepartmentRequest {
	s.Description = &v
	return s
}

func (s *CreateDepartmentRequest) SetLabel(v string) *CreateDepartmentRequest {
	s.Label = &v
	return s
}

func (s *CreateDepartmentRequest) SetName(v string) *CreateDepartmentRequest {
	s.Name = &v
	return s
}

type CreateDepartmentResponseBody struct {
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateDepartmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 11111111
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDepartmentResponseBody) SetCode(v string) *CreateDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDepartmentResponseBody) SetData(v *CreateDepartmentResponseBodyData) *CreateDepartmentResponseBody {
	s.Data = v
	return s
}

func (s *CreateDepartmentResponseBody) SetMessage(v string) *CreateDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDepartmentResponseBody) SetRequestId(v string) *CreateDepartmentResponseBody {
	s.RequestId = &v
	return s
}

type CreateDepartmentResponseBodyData struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 63bb629d-92bf-4cdc-ad0b-3032c926d23f
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateDepartmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDepartmentResponseBodyData) SetCreatedAt(v string) *CreateDepartmentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateDepartmentResponseBodyData) SetDescription(v string) *CreateDepartmentResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateDepartmentResponseBodyData) SetId(v string) *CreateDepartmentResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateDepartmentResponseBodyData) SetName(v string) *CreateDepartmentResponseBodyData {
	s.Name = &v
	return s
}

type CreateDepartmentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentResponse) GoString() string {
	return s.String()
}

func (s *CreateDepartmentResponse) SetHeaders(v map[string]*string) *CreateDepartmentResponse {
	s.Headers = v
	return s
}

func (s *CreateDepartmentResponse) SetStatusCode(v int32) *CreateDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDepartmentResponse) SetBody(v *CreateDepartmentResponseBody) *CreateDepartmentResponse {
	s.Body = v
	return s
}

type CreateDetectProcessRequest struct {
	// example:
	//
	// {}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// {}
	Draft *string `json:"Draft,omitempty" xml:"Draft,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// LOCAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDetectProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectProcessRequest) GoString() string {
	return s.String()
}

func (s *CreateDetectProcessRequest) SetContent(v string) *CreateDetectProcessRequest {
	s.Content = &v
	return s
}

func (s *CreateDetectProcessRequest) SetDraft(v string) *CreateDetectProcessRequest {
	s.Draft = &v
	return s
}

func (s *CreateDetectProcessRequest) SetName(v string) *CreateDetectProcessRequest {
	s.Name = &v
	return s
}

func (s *CreateDetectProcessRequest) SetType(v string) *CreateDetectProcessRequest {
	s.Type = &v
	return s
}

type CreateDetectProcessResponseBody struct {
	// example:
	//
	// OK
	Code    *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *CreateDetectProcessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0q1c45cd-3eee-1e60-b505-2e330b8755d3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDetectProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDetectProcessResponseBody) SetCode(v string) *CreateDetectProcessResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDetectProcessResponseBody) SetData(v *CreateDetectProcessResponseBodyData) *CreateDetectProcessResponseBody {
	s.Data = v
	return s
}

func (s *CreateDetectProcessResponseBody) SetMessage(v string) *CreateDetectProcessResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDetectProcessResponseBody) SetRequestId(v string) *CreateDetectProcessResponseBody {
	s.RequestId = &v
	return s
}

type CreateDetectProcessResponseBodyData struct {
	// example:
	//
	// {}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2020-12-04T14:47:59.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// example:
	//
	// {}
	Draft *string `json:"Draft,omitempty" xml:"Draft,omitempty"`
	// ID
	//
	// example:
	//
	// 0f1c45cd-3eee-4e60-b505-2e330b8755d3
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 987d563d38f5aef27feca8702c689bb1
	Md5  *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateDetectProcessResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectProcessResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDetectProcessResponseBodyData) SetContent(v string) *CreateDetectProcessResponseBodyData {
	s.Content = &v
	return s
}

func (s *CreateDetectProcessResponseBodyData) SetCreatedAt(v string) *CreateDetectProcessResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateDetectProcessResponseBodyData) SetDisabled(v bool) *CreateDetectProcessResponseBodyData {
	s.Disabled = &v
	return s
}

func (s *CreateDetectProcessResponseBodyData) SetDraft(v string) *CreateDetectProcessResponseBodyData {
	s.Draft = &v
	return s
}

func (s *CreateDetectProcessResponseBodyData) SetId(v string) *CreateDetectProcessResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateDetectProcessResponseBodyData) SetMd5(v string) *CreateDetectProcessResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *CreateDetectProcessResponseBodyData) SetName(v string) *CreateDetectProcessResponseBodyData {
	s.Name = &v
	return s
}

type CreateDetectProcessResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDetectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDetectProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectProcessResponse) GoString() string {
	return s.String()
}

func (s *CreateDetectProcessResponse) SetHeaders(v map[string]*string) *CreateDetectProcessResponse {
	s.Headers = v
	return s
}

func (s *CreateDetectProcessResponse) SetStatusCode(v int32) *CreateDetectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDetectProcessResponse) SetBody(v *CreateDetectProcessResponseBody) *CreateDetectProcessResponse {
	s.Body = v
	return s
}

type CreateRuleRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// [{"sequence":1,"actions":[{"name":"id_card_recognize"}]},{"sequence":2,"actions":[{"name":"document_title_recognize"},{"name":"flip_action_recognize"},{"name":"sign_action_recognize"}]},{"sequence":3,"actions":[{"name":"sign_recognize"}]},{"sequence":0,"actions":[{"name":"face_track"},{"name":"speech_to_text"}]}]
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleRequest) SetClientToken(v string) *CreateRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRuleRequest) SetContent(v string) *CreateRuleRequest {
	s.Content = &v
	return s
}

func (s *CreateRuleRequest) SetName(v string) *CreateRuleRequest {
	s.Name = &v
	return s
}

type CreateRuleResponseBody struct {
	// example:
	//
	// OK
	Code    *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *CreateRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBody) SetCode(v string) *CreateRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRuleResponseBody) SetData(v *CreateRuleResponseBodyData) *CreateRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateRuleResponseBody) SetMessage(v string) *CreateRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRuleResponseBody) SetRequestId(v string) *CreateRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateRuleResponseBodyData struct {
	// example:
	//
	// {}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBodyData) SetContent(v string) *CreateRuleResponseBodyData {
	s.Content = &v
	return s
}

func (s *CreateRuleResponseBodyData) SetId(v string) *CreateRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateRuleResponseBodyData) SetName(v string) *CreateRuleResponseBodyData {
	s.Name = &v
	return s
}

type CreateRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRuleResponse) SetHeaders(v map[string]*string) *CreateRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRuleResponse) SetStatusCode(v int32) *CreateRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRuleResponse) SetBody(v *CreateRuleResponseBody) *CreateRuleResponse {
	s.Body = v
	return s
}

type CreateSignatureRequest struct {
	// example:
	//
	// 5bbfb884-1186-4d48-906b-88d586770f6b
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// {"version":"1.0.0"}
	ClientBaseParam *string `json:"ClientBaseParam,omitempty" xml:"ClientBaseParam,omitempty"`
	// example:
	//
	// 1.0.001
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// 300
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 550070734466****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s CreateSignatureRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureRequest) GoString() string {
	return s.String()
}

func (s *CreateSignatureRequest) SetAppId(v string) *CreateSignatureRequest {
	s.AppId = &v
	return s
}

func (s *CreateSignatureRequest) SetClientBaseParam(v string) *CreateSignatureRequest {
	s.ClientBaseParam = &v
	return s
}

func (s *CreateSignatureRequest) SetClientVersion(v string) *CreateSignatureRequest {
	s.ClientVersion = &v
	return s
}

func (s *CreateSignatureRequest) SetExpireTime(v int64) *CreateSignatureRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateSignatureRequest) SetUid(v string) *CreateSignatureRequest {
	s.Uid = &v
	return s
}

type CreateSignatureResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code   *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data   *CreateSignatureResponseBodyData     `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Errors []*CreateSignatureResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSignatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSignatureResponseBody) SetCode(v string) *CreateSignatureResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSignatureResponseBody) SetData(v *CreateSignatureResponseBodyData) *CreateSignatureResponseBody {
	s.Data = v
	return s
}

func (s *CreateSignatureResponseBody) SetErrors(v []*CreateSignatureResponseBodyErrors) *CreateSignatureResponseBody {
	s.Errors = v
	return s
}

func (s *CreateSignatureResponseBody) SetHttpCode(v int32) *CreateSignatureResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateSignatureResponseBody) SetMessage(v string) *CreateSignatureResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSignatureResponseBody) SetRequestId(v string) *CreateSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSignatureResponseBody) SetSuccess(v bool) *CreateSignatureResponseBody {
	s.Success = &v
	return s
}

type CreateSignatureResponseBodyData struct {
	// example:
	//
	// 300
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 124325213125435
	RtcAppId   *string `json:"RtcAppId,omitempty" xml:"RtcAppId,omitempty"`
	RtcBizName *string `json:"RtcBizName,omitempty" xml:"RtcBizName,omitempty"`
	RtcSign    *string `json:"RtcSign,omitempty" xml:"RtcSign,omitempty"`
	// example:
	//
	// my_workspace
	RtcWorkspaceId *string `json:"RtcWorkspaceId,omitempty" xml:"RtcWorkspaceId,omitempty"`
}

func (s CreateSignatureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSignatureResponseBodyData) SetExpireTime(v string) *CreateSignatureResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *CreateSignatureResponseBodyData) SetRtcAppId(v string) *CreateSignatureResponseBodyData {
	s.RtcAppId = &v
	return s
}

func (s *CreateSignatureResponseBodyData) SetRtcBizName(v string) *CreateSignatureResponseBodyData {
	s.RtcBizName = &v
	return s
}

func (s *CreateSignatureResponseBodyData) SetRtcSign(v string) *CreateSignatureResponseBodyData {
	s.RtcSign = &v
	return s
}

func (s *CreateSignatureResponseBodyData) SetRtcWorkspaceId(v string) *CreateSignatureResponseBodyData {
	s.RtcWorkspaceId = &v
	return s
}

type CreateSignatureResponseBodyErrors struct {
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	Field   *string `json:"Field,omitempty" xml:"Field,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s CreateSignatureResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *CreateSignatureResponseBodyErrors) SetField(v string) *CreateSignatureResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *CreateSignatureResponseBodyErrors) SetMessage(v string) *CreateSignatureResponseBodyErrors {
	s.Message = &v
	return s
}

type CreateSignatureResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSignatureResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureResponse) GoString() string {
	return s.String()
}

func (s *CreateSignatureResponse) SetHeaders(v map[string]*string) *CreateSignatureResponse {
	s.Headers = v
	return s
}

func (s *CreateSignatureResponse) SetStatusCode(v int32) *CreateSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSignatureResponse) SetBody(v *CreateSignatureResponseBody) *CreateSignatureResponse {
	s.Body = v
	return s
}

type CreateTaskGroupRequest struct {
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 3
	Day []*int32 `json:"Day,omitempty" xml:"Day,omitempty" type:"Repeated"`
	// example:
	//
	// 2020-10-10
	ExpireAt  *string `json:"ExpireAt,omitempty" xml:"ExpireAt,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 17:00
	RunnableTimeFrom *string `json:"RunnableTimeFrom,omitempty" xml:"RunnableTimeFrom,omitempty"`
	// example:
	//
	// 18:00
	RunnableTimeTo *string `json:"RunnableTimeTo,omitempty" xml:"RunnableTimeTo,omitempty"`
	// example:
	//
	// immediately
	TriggerPeriod *string                            `json:"TriggerPeriod,omitempty" xml:"TriggerPeriod,omitempty"`
	VideoInfo     []*CreateTaskGroupRequestVideoInfo `json:"VideoInfo,omitempty" xml:"VideoInfo,omitempty" type:"Repeated"`
}

func (s CreateTaskGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupRequest) SetAppId(v string) *CreateTaskGroupRequest {
	s.AppId = &v
	return s
}

func (s *CreateTaskGroupRequest) SetClientToken(v string) *CreateTaskGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTaskGroupRequest) SetDay(v []*int32) *CreateTaskGroupRequest {
	s.Day = v
	return s
}

func (s *CreateTaskGroupRequest) SetExpireAt(v string) *CreateTaskGroupRequest {
	s.ExpireAt = &v
	return s
}

func (s *CreateTaskGroupRequest) SetGroupName(v string) *CreateTaskGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateTaskGroupRequest) SetRuleId(v string) *CreateTaskGroupRequest {
	s.RuleId = &v
	return s
}

func (s *CreateTaskGroupRequest) SetRunnableTimeFrom(v string) *CreateTaskGroupRequest {
	s.RunnableTimeFrom = &v
	return s
}

func (s *CreateTaskGroupRequest) SetRunnableTimeTo(v string) *CreateTaskGroupRequest {
	s.RunnableTimeTo = &v
	return s
}

func (s *CreateTaskGroupRequest) SetTriggerPeriod(v string) *CreateTaskGroupRequest {
	s.TriggerPeriod = &v
	return s
}

func (s *CreateTaskGroupRequest) SetVideoInfo(v []*CreateTaskGroupRequestVideoInfo) *CreateTaskGroupRequest {
	s.VideoInfo = v
	return s
}

type CreateTaskGroupRequestVideoInfo struct {
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 18/mrtc//641905591891464_record_64190559189146412713.mp4.meta
	VideoMetaUrl *string `json:"VideoMetaUrl,omitempty" xml:"VideoMetaUrl,omitempty"`
	// example:
	//
	// 18/mrtc//641905591891464_record_64190559189146412713.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s CreateTaskGroupRequestVideoInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskGroupRequestVideoInfo) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupRequestVideoInfo) SetRuleId(v string) *CreateTaskGroupRequestVideoInfo {
	s.RuleId = &v
	return s
}

func (s *CreateTaskGroupRequestVideoInfo) SetVideoMetaUrl(v string) *CreateTaskGroupRequestVideoInfo {
	s.VideoMetaUrl = &v
	return s
}

func (s *CreateTaskGroupRequestVideoInfo) SetVideoUrl(v string) *CreateTaskGroupRequestVideoInfo {
	s.VideoUrl = &v
	return s
}

type CreateTaskGroupResponseBody struct {
	// example:
	//
	// OK
	Code    *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *CreateTaskGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTaskGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupResponseBody) SetCode(v string) *CreateTaskGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskGroupResponseBody) SetData(v *CreateTaskGroupResponseBodyData) *CreateTaskGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateTaskGroupResponseBody) SetMessage(v string) *CreateTaskGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskGroupResponseBody) SetRequestId(v string) *CreateTaskGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateTaskGroupResponseBodyData struct {
	// example:
	//
	// 0
	CompletedTasks *int32 `json:"CompletedTasks,omitempty" xml:"CompletedTasks,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// default
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// runnable
	Status  *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalTasks *int32 `json:"TotalTasks,omitempty" xml:"TotalTasks,omitempty"`
}

func (s CreateTaskGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupResponseBodyData) SetCompletedTasks(v int32) *CreateTaskGroupResponseBodyData {
	s.CompletedTasks = &v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetCreatedAt(v string) *CreateTaskGroupResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetId(v string) *CreateTaskGroupResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetName(v string) *CreateTaskGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetRuleId(v string) *CreateTaskGroupResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetRuleName(v string) *CreateTaskGroupResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetStatus(v string) *CreateTaskGroupResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetTaskIds(v []*string) *CreateTaskGroupResponseBodyData {
	s.TaskIds = v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetTotalTasks(v int32) *CreateTaskGroupResponseBodyData {
	s.TotalTasks = &v
	return s
}

type CreateTaskGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTaskGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupResponse) SetHeaders(v map[string]*string) *CreateTaskGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskGroupResponse) SetStatusCode(v int32) *CreateTaskGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskGroupResponse) SetBody(v *CreateTaskGroupResponseBody) *CreateTaskGroupResponse {
	s.Body = v
	return s
}

type CreateTtsQuestionRequest struct {
	Request *CreateTtsQuestionRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s CreateTtsQuestionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTtsQuestionRequest) GoString() string {
	return s.String()
}

func (s *CreateTtsQuestionRequest) SetRequest(v *CreateTtsQuestionRequestRequest) *CreateTtsQuestionRequest {
	s.Request = v
	return s
}

type CreateTtsQuestionRequestRequest struct {
	Answer   *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// example:
	//
	// 47584ba4-9781-496b-8e6f-c8525a213***
	QuestionGroupId *string `json:"QuestionGroupId,omitempty" xml:"QuestionGroupId,omitempty"`
}

func (s CreateTtsQuestionRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTtsQuestionRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateTtsQuestionRequestRequest) SetAnswer(v string) *CreateTtsQuestionRequestRequest {
	s.Answer = &v
	return s
}

func (s *CreateTtsQuestionRequestRequest) SetQuestion(v string) *CreateTtsQuestionRequestRequest {
	s.Question = &v
	return s
}

func (s *CreateTtsQuestionRequestRequest) SetQuestionGroupId(v string) *CreateTtsQuestionRequestRequest {
	s.QuestionGroupId = &v
	return s
}

type CreateTtsQuestionShrinkRequest struct {
	RequestShrink *string `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s CreateTtsQuestionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTtsQuestionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTtsQuestionShrinkRequest) SetRequestShrink(v string) *CreateTtsQuestionShrinkRequest {
	s.RequestShrink = &v
	return s
}

type CreateTtsQuestionResponseBody struct {
	// example:
	//
	// OK
	Code *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateTtsQuestionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9AA9055-F73D-592C-832B-5AEECB093***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTtsQuestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTtsQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTtsQuestionResponseBody) SetCode(v int32) *CreateTtsQuestionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTtsQuestionResponseBody) SetData(v *CreateTtsQuestionResponseBodyData) *CreateTtsQuestionResponseBody {
	s.Data = v
	return s
}

func (s *CreateTtsQuestionResponseBody) SetHttpCode(v int32) *CreateTtsQuestionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateTtsQuestionResponseBody) SetMessage(v string) *CreateTtsQuestionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTtsQuestionResponseBody) SetRequestId(v string) *CreateTtsQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTtsQuestionResponseBody) SetSuccess(v bool) *CreateTtsQuestionResponseBody {
	s.Success = &v
	return s
}

type CreateTtsQuestionResponseBodyData struct {
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4ba***
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateTtsQuestionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTtsQuestionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTtsQuestionResponseBodyData) SetId(v string) *CreateTtsQuestionResponseBodyData {
	s.Id = &v
	return s
}

type CreateTtsQuestionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTtsQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTtsQuestionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTtsQuestionResponse) GoString() string {
	return s.String()
}

func (s *CreateTtsQuestionResponse) SetHeaders(v map[string]*string) *CreateTtsQuestionResponse {
	s.Headers = v
	return s
}

func (s *CreateTtsQuestionResponse) SetStatusCode(v int32) *CreateTtsQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTtsQuestionResponse) SetBody(v *CreateTtsQuestionResponseBody) *CreateTtsQuestionResponse {
	s.Body = v
	return s
}

type CreateTtsQuestionGroupRequest struct {
	Request *CreateTtsQuestionGroupRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s CreateTtsQuestionGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTtsQuestionGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateTtsQuestionGroupRequest) SetRequest(v *CreateTtsQuestionGroupRequestRequest) *CreateTtsQuestionGroupRequest {
	s.Request = v
	return s
}

type CreateTtsQuestionGroupRequestRequest struct {
	// example:
	//
	// PCM
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// 50
	PitchRate *int32 `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// example:
	//
	// 16000
	SampleRate *int32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// example:
	//
	// 50
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// example:
	//
	// xiaoyun
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s CreateTtsQuestionGroupRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTtsQuestionGroupRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateTtsQuestionGroupRequestRequest) SetFormat(v string) *CreateTtsQuestionGroupRequestRequest {
	s.Format = &v
	return s
}

func (s *CreateTtsQuestionGroupRequestRequest) SetPitchRate(v int32) *CreateTtsQuestionGroupRequestRequest {
	s.PitchRate = &v
	return s
}

func (s *CreateTtsQuestionGroupRequestRequest) SetSampleRate(v int32) *CreateTtsQuestionGroupRequestRequest {
	s.SampleRate = &v
	return s
}

func (s *CreateTtsQuestionGroupRequestRequest) SetSpeechRate(v int32) *CreateTtsQuestionGroupRequestRequest {
	s.SpeechRate = &v
	return s
}

func (s *CreateTtsQuestionGroupRequestRequest) SetVoice(v string) *CreateTtsQuestionGroupRequestRequest {
	s.Voice = &v
	return s
}

func (s *CreateTtsQuestionGroupRequestRequest) SetVolume(v int32) *CreateTtsQuestionGroupRequestRequest {
	s.Volume = &v
	return s
}

type CreateTtsQuestionGroupShrinkRequest struct {
	RequestShrink *string `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s CreateTtsQuestionGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTtsQuestionGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTtsQuestionGroupShrinkRequest) SetRequestShrink(v string) *CreateTtsQuestionGroupShrinkRequest {
	s.RequestShrink = &v
	return s
}

type CreateTtsQuestionGroupResponseBody struct {
	// example:
	//
	// OK
	Code *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateTtsQuestionGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B2AADC9E-2A58-5918-AE4E-FF59E19D7***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTtsQuestionGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTtsQuestionGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTtsQuestionGroupResponseBody) SetCode(v int32) *CreateTtsQuestionGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTtsQuestionGroupResponseBody) SetData(v *CreateTtsQuestionGroupResponseBodyData) *CreateTtsQuestionGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateTtsQuestionGroupResponseBody) SetHttpCode(v int32) *CreateTtsQuestionGroupResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateTtsQuestionGroupResponseBody) SetMessage(v string) *CreateTtsQuestionGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTtsQuestionGroupResponseBody) SetRequestId(v string) *CreateTtsQuestionGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTtsQuestionGroupResponseBody) SetSuccess(v bool) *CreateTtsQuestionGroupResponseBody {
	s.Success = &v
	return s
}

type CreateTtsQuestionGroupResponseBodyData struct {
	// example:
	//
	// 63bb629d-92bf-4cdc-ad0b-3032c926d***
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateTtsQuestionGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTtsQuestionGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTtsQuestionGroupResponseBodyData) SetId(v string) *CreateTtsQuestionGroupResponseBodyData {
	s.Id = &v
	return s
}

type CreateTtsQuestionGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTtsQuestionGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTtsQuestionGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTtsQuestionGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateTtsQuestionGroupResponse) SetHeaders(v map[string]*string) *CreateTtsQuestionGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateTtsQuestionGroupResponse) SetStatusCode(v int32) *CreateTtsQuestionGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTtsQuestionGroupResponse) SetBody(v *CreateTtsQuestionGroupResponseBody) *CreateTtsQuestionGroupResponse {
	s.Body = v
	return s
}

type CreateUserDepartmentsRequest struct {
	ClientToken  *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DepartmentId []*string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty" type:"Repeated"`
	UserId       []*string `json:"UserId,omitempty" xml:"UserId,omitempty" type:"Repeated"`
}

func (s CreateUserDepartmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDepartmentsRequest) GoString() string {
	return s.String()
}

func (s *CreateUserDepartmentsRequest) SetClientToken(v string) *CreateUserDepartmentsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateUserDepartmentsRequest) SetDepartmentId(v []*string) *CreateUserDepartmentsRequest {
	s.DepartmentId = v
	return s
}

func (s *CreateUserDepartmentsRequest) SetUserId(v []*string) *CreateUserDepartmentsRequest {
	s.UserId = v
	return s
}

type CreateUserDepartmentsResponseBody struct {
	// example:
	//
	// OK
	Code    *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0B576AAB-A638-5029-9A54-A7C1DB5AC0B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUserDepartmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDepartmentsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserDepartmentsResponseBody) SetCode(v string) *CreateUserDepartmentsResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUserDepartmentsResponseBody) SetData(v map[string]interface{}) *CreateUserDepartmentsResponseBody {
	s.Data = v
	return s
}

func (s *CreateUserDepartmentsResponseBody) SetMessage(v string) *CreateUserDepartmentsResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserDepartmentsResponseBody) SetRequestId(v string) *CreateUserDepartmentsResponseBody {
	s.RequestId = &v
	return s
}

type CreateUserDepartmentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserDepartmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserDepartmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDepartmentsResponse) GoString() string {
	return s.String()
}

func (s *CreateUserDepartmentsResponse) SetHeaders(v map[string]*string) *CreateUserDepartmentsResponse {
	s.Headers = v
	return s
}

func (s *CreateUserDepartmentsResponse) SetStatusCode(v int32) *CreateUserDepartmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserDepartmentsResponse) SetBody(v *CreateUserDepartmentsResponseBody) *CreateUserDepartmentsResponse {
	s.Body = v
	return s
}

type CreateVideoMergeTaskRequest struct {
	VideoMergeRequest *CreateVideoMergeTaskRequestVideoMergeRequest `json:"VideoMergeRequest,omitempty" xml:"VideoMergeRequest,omitempty" type:"Struct"`
}

func (s CreateVideoMergeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoMergeTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoMergeTaskRequest) SetVideoMergeRequest(v *CreateVideoMergeTaskRequestVideoMergeRequest) *CreateVideoMergeTaskRequest {
	s.VideoMergeRequest = v
	return s
}

type CreateVideoMergeTaskRequestVideoMergeRequest struct {
	// example:
	//
	// https://h5-api.neoclub.cn/v1/bff/service/other/rtc/callback
	CallbackUrl  *string                                                     `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	LayoutStyles []*CreateVideoMergeTaskRequestVideoMergeRequestLayoutStyles `json:"LayoutStyles,omitempty" xml:"LayoutStyles,omitempty" type:"Repeated"`
	VideoList    []*CreateVideoMergeTaskRequestVideoMergeRequestVideoList    `json:"VideoList,omitempty" xml:"VideoList,omitempty" type:"Repeated"`
	Watermark    *CreateVideoMergeTaskRequestVideoMergeRequestWatermark      `json:"Watermark,omitempty" xml:"Watermark,omitempty" type:"Struct"`
}

func (s CreateVideoMergeTaskRequestVideoMergeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoMergeTaskRequestVideoMergeRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequest) SetCallbackUrl(v string) *CreateVideoMergeTaskRequestVideoMergeRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequest) SetLayoutStyles(v []*CreateVideoMergeTaskRequestVideoMergeRequestLayoutStyles) *CreateVideoMergeTaskRequestVideoMergeRequest {
	s.LayoutStyles = v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequest) SetVideoList(v []*CreateVideoMergeTaskRequestVideoMergeRequestVideoList) *CreateVideoMergeTaskRequestVideoMergeRequest {
	s.VideoList = v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequest) SetWatermark(v *CreateVideoMergeTaskRequestVideoMergeRequestWatermark) *CreateVideoMergeTaskRequestVideoMergeRequest {
	s.Watermark = v
	return s
}

type CreateVideoMergeTaskRequestVideoMergeRequestLayoutStyles struct {
	// example:
	//
	// 3037
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 2
	InputNum    *int64                                                                 `json:"InputNum,omitempty" xml:"InputNum,omitempty"`
	VideoStyles []*CreateVideoMergeTaskRequestVideoMergeRequestLayoutStylesVideoStyles `json:"VideoStyles,omitempty" xml:"VideoStyles,omitempty" type:"Repeated"`
	// example:
	//
	// 800
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s CreateVideoMergeTaskRequestVideoMergeRequestLayoutStyles) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoMergeTaskRequestVideoMergeRequestLayoutStyles) GoString() string {
	return s.String()
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStyles) SetHeight(v int64) *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStyles {
	s.Height = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStyles) SetInputNum(v int64) *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStyles {
	s.InputNum = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStyles) SetVideoStyles(v []*CreateVideoMergeTaskRequestVideoMergeRequestLayoutStylesVideoStyles) *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStyles {
	s.VideoStyles = v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStyles) SetWidth(v int64) *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStyles {
	s.Width = &v
	return s
}

type CreateVideoMergeTaskRequestVideoMergeRequestLayoutStylesVideoStyles struct {
	// example:
	//
	// http://xxx.xxx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 888
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 100
	PositionX *int64 `json:"PositionX,omitempty" xml:"PositionX,omitempty"`
	// example:
	//
	// 400
	PositionY *int64 `json:"PositionY,omitempty" xml:"PositionY,omitempty"`
	// example:
	//
	// 100
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s CreateVideoMergeTaskRequestVideoMergeRequestLayoutStylesVideoStyles) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoMergeTaskRequestVideoMergeRequestLayoutStylesVideoStyles) GoString() string {
	return s.String()
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStylesVideoStyles) SetFileName(v string) *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStylesVideoStyles {
	s.FileName = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStylesVideoStyles) SetHeight(v int64) *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStylesVideoStyles {
	s.Height = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStylesVideoStyles) SetPositionX(v int64) *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStylesVideoStyles {
	s.PositionX = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStylesVideoStyles) SetPositionY(v int64) *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStylesVideoStyles {
	s.PositionY = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStylesVideoStyles) SetWidth(v int64) *CreateVideoMergeTaskRequestVideoMergeRequestLayoutStylesVideoStyles {
	s.Width = &v
	return s
}

type CreateVideoMergeTaskRequestVideoMergeRequestVideoList struct {
	// example:
	//
	// 2021-12-28
	EndTime  *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// http://shuanglu-record-finance.oss-cn-shanghai.aliyuncs.com/record/4x5avhil/264516_33430_1/2022-03-21-13-56-38_2022-03-21-14-17-22.mp4
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// 5
	MergeBeginTime *int64 `json:"MergeBeginTime,omitempty" xml:"MergeBeginTime,omitempty"`
	// example:
	//
	// 15
	MergeEndTime *int64 `json:"MergeEndTime,omitempty" xml:"MergeEndTime,omitempty"`
	// example:
	//
	// true
	PrimeVideo *bool `json:"PrimeVideo,omitempty" xml:"PrimeVideo,omitempty"`
	// example:
	//
	// 2022-03-05
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateVideoMergeTaskRequestVideoMergeRequestVideoList) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoMergeTaskRequestVideoMergeRequestVideoList) GoString() string {
	return s.String()
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestVideoList) SetEndTime(v int64) *CreateVideoMergeTaskRequestVideoMergeRequestVideoList {
	s.EndTime = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestVideoList) SetFileName(v string) *CreateVideoMergeTaskRequestVideoMergeRequestVideoList {
	s.FileName = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestVideoList) SetFileUrl(v string) *CreateVideoMergeTaskRequestVideoMergeRequestVideoList {
	s.FileUrl = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestVideoList) SetMergeBeginTime(v int64) *CreateVideoMergeTaskRequestVideoMergeRequestVideoList {
	s.MergeBeginTime = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestVideoList) SetMergeEndTime(v int64) *CreateVideoMergeTaskRequestVideoMergeRequestVideoList {
	s.MergeEndTime = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestVideoList) SetPrimeVideo(v bool) *CreateVideoMergeTaskRequestVideoMergeRequestVideoList {
	s.PrimeVideo = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestVideoList) SetStartTime(v int64) *CreateVideoMergeTaskRequestVideoMergeRequestVideoList {
	s.StartTime = &v
	return s
}

type CreateVideoMergeTaskRequestVideoMergeRequestWatermark struct {
	// example:
	//
	// 14803425
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// example:
	//
	// 20
	FontSize *int64 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// example:
	//
	// 100
	PositionX *int64 `json:"PositionX,omitempty" xml:"PositionX,omitempty"`
	// example:
	//
	// 400
	PositionY *int64  `json:"PositionY,omitempty" xml:"PositionY,omitempty"`
	Text      *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 1617245014
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s CreateVideoMergeTaskRequestVideoMergeRequestWatermark) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoMergeTaskRequestVideoMergeRequestWatermark) GoString() string {
	return s.String()
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestWatermark) SetFontColor(v string) *CreateVideoMergeTaskRequestVideoMergeRequestWatermark {
	s.FontColor = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestWatermark) SetFontSize(v int64) *CreateVideoMergeTaskRequestVideoMergeRequestWatermark {
	s.FontSize = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestWatermark) SetPositionX(v int64) *CreateVideoMergeTaskRequestVideoMergeRequestWatermark {
	s.PositionX = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestWatermark) SetPositionY(v int64) *CreateVideoMergeTaskRequestVideoMergeRequestWatermark {
	s.PositionY = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestWatermark) SetText(v string) *CreateVideoMergeTaskRequestVideoMergeRequestWatermark {
	s.Text = &v
	return s
}

func (s *CreateVideoMergeTaskRequestVideoMergeRequestWatermark) SetTimestamp(v int64) *CreateVideoMergeTaskRequestVideoMergeRequestWatermark {
	s.Timestamp = &v
	return s
}

type CreateVideoMergeTaskShrinkRequest struct {
	VideoMergeRequestShrink *string `json:"VideoMergeRequest,omitempty" xml:"VideoMergeRequest,omitempty"`
}

func (s CreateVideoMergeTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoMergeTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoMergeTaskShrinkRequest) SetVideoMergeRequestShrink(v string) *CreateVideoMergeTaskShrinkRequest {
	s.VideoMergeRequestShrink = &v
	return s
}

type CreateVideoMergeTaskResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// AC140413004421B8D17C549B31B20000
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// success
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateVideoMergeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoMergeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoMergeTaskResponseBody) SetCode(v int32) *CreateVideoMergeTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVideoMergeTaskResponseBody) SetData(v string) *CreateVideoMergeTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateVideoMergeTaskResponseBody) SetMessage(v string) *CreateVideoMergeTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVideoMergeTaskResponseBody) SetStatus(v bool) *CreateVideoMergeTaskResponseBody {
	s.Status = &v
	return s
}

type CreateVideoMergeTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVideoMergeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVideoMergeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoMergeTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoMergeTaskResponse) SetHeaders(v map[string]*string) *CreateVideoMergeTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoMergeTaskResponse) SetStatusCode(v int32) *CreateVideoMergeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoMergeTaskResponse) SetBody(v *CreateVideoMergeTaskResponseBody) *CreateVideoMergeTaskResponse {
	s.Body = v
	return s
}

type CreateWatermarkRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWatermarkRequest) GoString() string {
	return s.String()
}

func (s *CreateWatermarkRequest) SetName(v string) *CreateWatermarkRequest {
	s.Name = &v
	return s
}

func (s *CreateWatermarkRequest) SetValue(v string) *CreateWatermarkRequest {
	s.Value = &v
	return s
}

type CreateWatermarkResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code   *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data   *CreateWatermarkResponseBodyData     `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Errors []*CreateWatermarkResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWatermarkResponseBody) SetCode(v string) *CreateWatermarkResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWatermarkResponseBody) SetData(v *CreateWatermarkResponseBodyData) *CreateWatermarkResponseBody {
	s.Data = v
	return s
}

func (s *CreateWatermarkResponseBody) SetErrors(v []*CreateWatermarkResponseBodyErrors) *CreateWatermarkResponseBody {
	s.Errors = v
	return s
}

func (s *CreateWatermarkResponseBody) SetHttpCode(v int32) *CreateWatermarkResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateWatermarkResponseBody) SetMessage(v string) *CreateWatermarkResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWatermarkResponseBody) SetRequestId(v string) *CreateWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWatermarkResponseBody) SetSuccess(v bool) *CreateWatermarkResponseBody {
	s.Success = &v
	return s
}

type CreateWatermarkResponseBodyData struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateWatermarkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateWatermarkResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWatermarkResponseBodyData) SetCreatedAt(v string) *CreateWatermarkResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateWatermarkResponseBodyData) SetId(v string) *CreateWatermarkResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateWatermarkResponseBodyData) SetName(v string) *CreateWatermarkResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateWatermarkResponseBodyData) SetValue(v string) *CreateWatermarkResponseBodyData {
	s.Value = &v
	return s
}

type CreateWatermarkResponseBodyErrors struct {
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s CreateWatermarkResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s CreateWatermarkResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *CreateWatermarkResponseBodyErrors) SetField(v string) *CreateWatermarkResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *CreateWatermarkResponseBodyErrors) SetMessage(v string) *CreateWatermarkResponseBodyErrors {
	s.Message = &v
	return s
}

type CreateWatermarkResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWatermarkResponse) GoString() string {
	return s.String()
}

func (s *CreateWatermarkResponse) SetHeaders(v map[string]*string) *CreateWatermarkResponse {
	s.Headers = v
	return s
}

func (s *CreateWatermarkResponse) SetStatusCode(v int32) *CreateWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWatermarkResponse) SetBody(v *CreateWatermarkResponseBody) *CreateWatermarkResponse {
	s.Body = v
	return s
}

type DeleteAppRequest struct {
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) SetId(v string) *DeleteAppRequest {
	s.Id = &v
	return s
}

type DeleteAppResponseBody struct {
	// example:
	//
	// OK
	Code    *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) SetCode(v string) *DeleteAppResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppResponseBody) SetData(v map[string]interface{}) *DeleteAppResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAppResponseBody) SetMessage(v string) *DeleteAppResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResponse) SetHeaders(v map[string]*string) *DeleteAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppResponse) SetStatusCode(v int32) *DeleteAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppResponse) SetBody(v *DeleteAppResponseBody) *DeleteAppResponse {
	s.Body = v
	return s
}

type DeleteDepartmentRequest struct {
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDepartmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteDepartmentRequest) SetId(v string) *DeleteDepartmentRequest {
	s.Id = &v
	return s
}

type DeleteDepartmentResponseBody struct {
	// example:
	//
	// OK
	Code    *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDepartmentResponseBody) SetCode(v string) *DeleteDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetData(v map[string]interface{}) *DeleteDepartmentResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDepartmentResponseBody) SetMessage(v string) *DeleteDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetRequestId(v string) *DeleteDepartmentResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDepartmentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDepartmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteDepartmentResponse) SetHeaders(v map[string]*string) *DeleteDepartmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteDepartmentResponse) SetStatusCode(v int32) *DeleteDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDepartmentResponse) SetBody(v *DeleteDepartmentResponseBody) *DeleteDepartmentResponse {
	s.Body = v
	return s
}

type DeleteDetectProcessRequest struct {
	// ID
	//
	// example:
	//
	// 0f1c45cd-3eee-4e60-b505-2e330b8755d3
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDetectProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDetectProcessRequest) GoString() string {
	return s.String()
}

func (s *DeleteDetectProcessRequest) SetId(v string) *DeleteDetectProcessRequest {
	s.Id = &v
	return s
}

type DeleteDetectProcessResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// -
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1fdc45cd-3eee-4e60-b505-2e330b8755d3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDetectProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDetectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDetectProcessResponseBody) SetCode(v string) *DeleteDetectProcessResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDetectProcessResponseBody) SetData(v map[string]interface{}) *DeleteDetectProcessResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDetectProcessResponseBody) SetMessage(v string) *DeleteDetectProcessResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDetectProcessResponseBody) SetRequestId(v string) *DeleteDetectProcessResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDetectProcessResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDetectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDetectProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDetectProcessResponse) GoString() string {
	return s.String()
}

func (s *DeleteDetectProcessResponse) SetHeaders(v map[string]*string) *DeleteDetectProcessResponse {
	s.Headers = v
	return s
}

func (s *DeleteDetectProcessResponse) SetStatusCode(v int32) *DeleteDetectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDetectProcessResponse) SetBody(v *DeleteDetectProcessResponseBody) *DeleteDetectProcessResponse {
	s.Body = v
	return s
}

type DeleteRuleRequest struct {
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) SetId(v string) *DeleteRuleRequest {
	s.Id = &v
	return s
}

type DeleteRuleResponseBody struct {
	// example:
	//
	// OK
	Code    *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *DeleteRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBody) SetCode(v string) *DeleteRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRuleResponseBody) SetData(v *DeleteRuleResponseBodyData) *DeleteRuleResponseBody {
	s.Data = v
	return s
}

func (s *DeleteRuleResponseBody) SetMessage(v string) *DeleteRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRuleResponseBody) SetRequestId(v string) *DeleteRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRuleResponseBodyData struct {
	// example:
	//
	// {}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBodyData) SetContent(v string) *DeleteRuleResponseBodyData {
	s.Content = &v
	return s
}

func (s *DeleteRuleResponseBodyData) SetCreatedAt(v string) *DeleteRuleResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *DeleteRuleResponseBodyData) SetId(v string) *DeleteRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteRuleResponseBodyData) SetName(v string) *DeleteRuleResponseBodyData {
	s.Name = &v
	return s
}

type DeleteRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponse) SetHeaders(v map[string]*string) *DeleteRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRuleResponse) SetStatusCode(v int32) *DeleteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRuleResponse) SetBody(v *DeleteRuleResponseBody) *DeleteRuleResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetId(v string) *DeleteUserRequest {
	s.Id = &v
	return s
}

type DeleteUserResponseBody struct {
	// example:
	//
	// OK
	Code    *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    map[string]interface{}          `json:"Data,omitempty" xml:"Data,omitempty"`
	Errors  []*DeleteUserResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	Message *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) SetCode(v string) *DeleteUserResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUserResponseBody) SetData(v map[string]interface{}) *DeleteUserResponseBody {
	s.Data = v
	return s
}

func (s *DeleteUserResponseBody) SetErrors(v []*DeleteUserResponseBodyErrors) *DeleteUserResponseBody {
	s.Errors = v
	return s
}

func (s *DeleteUserResponseBody) SetMessage(v string) *DeleteUserResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserResponseBodyErrors struct {
	Field   *string `json:"Field,omitempty" xml:"Field,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DeleteUserResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBodyErrors) SetField(v string) *DeleteUserResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *DeleteUserResponseBodyErrors) SetMessage(v string) *DeleteUserResponseBodyErrors {
	s.Message = &v
	return s
}

type DeleteUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetStatusCode(v int32) *DeleteUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type DeleteUserDepartmentsRequest struct {
	DepartmentId []*string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty" type:"Repeated"`
	UserId       []*string `json:"UserId,omitempty" xml:"UserId,omitempty" type:"Repeated"`
}

func (s DeleteUserDepartmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDepartmentsRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserDepartmentsRequest) SetDepartmentId(v []*string) *DeleteUserDepartmentsRequest {
	s.DepartmentId = v
	return s
}

func (s *DeleteUserDepartmentsRequest) SetUserId(v []*string) *DeleteUserDepartmentsRequest {
	s.UserId = v
	return s
}

type DeleteUserDepartmentsResponseBody struct {
	// example:
	//
	// OK
	Code    *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 48A2B2E2-9995-5220-B77C-871119CB05CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserDepartmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDepartmentsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserDepartmentsResponseBody) SetCode(v string) *DeleteUserDepartmentsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUserDepartmentsResponseBody) SetData(v map[string]interface{}) *DeleteUserDepartmentsResponseBody {
	s.Data = v
	return s
}

func (s *DeleteUserDepartmentsResponseBody) SetMessage(v string) *DeleteUserDepartmentsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserDepartmentsResponseBody) SetRequestId(v string) *DeleteUserDepartmentsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserDepartmentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserDepartmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserDepartmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDepartmentsResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserDepartmentsResponse) SetHeaders(v map[string]*string) *DeleteUserDepartmentsResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserDepartmentsResponse) SetStatusCode(v int32) *DeleteUserDepartmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserDepartmentsResponse) SetBody(v *DeleteUserDepartmentsResponseBody) *DeleteUserDepartmentsResponse {
	s.Body = v
	return s
}

type DeleteWatermarkRequest struct {
	// example:
	//
	// e5a923e0e727f212813a195e274b02c6
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s DeleteWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWatermarkRequest) GoString() string {
	return s.String()
}

func (s *DeleteWatermarkRequest) SetWatermarkId(v string) *DeleteWatermarkRequest {
	s.WatermarkId = &v
	return s
}

type DeleteWatermarkResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code   *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data   *DeleteWatermarkResponseBodyData     `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Errors []*DeleteWatermarkResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// 84118BF0-56F7-54D2-8C1A-35BBBBD5016B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWatermarkResponseBody) SetCode(v string) *DeleteWatermarkResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWatermarkResponseBody) SetData(v *DeleteWatermarkResponseBodyData) *DeleteWatermarkResponseBody {
	s.Data = v
	return s
}

func (s *DeleteWatermarkResponseBody) SetErrors(v []*DeleteWatermarkResponseBodyErrors) *DeleteWatermarkResponseBody {
	s.Errors = v
	return s
}

func (s *DeleteWatermarkResponseBody) SetHttpCode(v int32) *DeleteWatermarkResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteWatermarkResponseBody) SetMessage(v string) *DeleteWatermarkResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWatermarkResponseBody) SetRequestId(v string) *DeleteWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWatermarkResponseBody) SetSuccess(v bool) *DeleteWatermarkResponseBody {
	s.Success = &v
	return s
}

type DeleteWatermarkResponseBodyData struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 0f1c45cd-3eee-4e60-b505-2e330b8755d3
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DeleteWatermarkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteWatermarkResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteWatermarkResponseBodyData) SetCreatedAt(v string) *DeleteWatermarkResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *DeleteWatermarkResponseBodyData) SetId(v string) *DeleteWatermarkResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteWatermarkResponseBodyData) SetName(v string) *DeleteWatermarkResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteWatermarkResponseBodyData) SetValue(v string) *DeleteWatermarkResponseBodyData {
	s.Value = &v
	return s
}

type DeleteWatermarkResponseBodyErrors struct {
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DeleteWatermarkResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s DeleteWatermarkResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *DeleteWatermarkResponseBodyErrors) SetField(v string) *DeleteWatermarkResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *DeleteWatermarkResponseBodyErrors) SetMessage(v string) *DeleteWatermarkResponseBodyErrors {
	s.Message = &v
	return s
}

type DeleteWatermarkResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWatermarkResponse) GoString() string {
	return s.String()
}

func (s *DeleteWatermarkResponse) SetHeaders(v map[string]*string) *DeleteWatermarkResponse {
	s.Headers = v
	return s
}

func (s *DeleteWatermarkResponse) SetStatusCode(v int32) *DeleteWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWatermarkResponse) SetBody(v *DeleteWatermarkResponseBody) *DeleteWatermarkResponse {
	s.Body = v
	return s
}

type FaceCompareRequest struct {
	FaceRequest *FaceCompareRequestFaceRequest `json:"FaceRequest,omitempty" xml:"FaceRequest,omitempty" type:"Struct"`
}

func (s FaceCompareRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareRequest) GoString() string {
	return s.String()
}

func (s *FaceCompareRequest) SetFaceRequest(v *FaceCompareRequestFaceRequest) *FaceCompareRequest {
	s.FaceRequest = v
	return s
}

type FaceCompareRequestFaceRequest struct {
	// example:
	//
	// d61be709-49d2-4cf1-b219-cd6181f72***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	SourceImage *string `json:"SourceImage,omitempty" xml:"SourceImage,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	TargetImage *string `json:"TargetImage,omitempty" xml:"TargetImage,omitempty"`
}

func (s FaceCompareRequestFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareRequestFaceRequest) GoString() string {
	return s.String()
}

func (s *FaceCompareRequestFaceRequest) SetAppId(v string) *FaceCompareRequestFaceRequest {
	s.AppId = &v
	return s
}

func (s *FaceCompareRequestFaceRequest) SetSourceImage(v string) *FaceCompareRequestFaceRequest {
	s.SourceImage = &v
	return s
}

func (s *FaceCompareRequestFaceRequest) SetTargetImage(v string) *FaceCompareRequestFaceRequest {
	s.TargetImage = &v
	return s
}

type FaceCompareShrinkRequest struct {
	FaceRequestShrink *string `json:"FaceRequest,omitempty" xml:"FaceRequest,omitempty"`
}

func (s FaceCompareShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareShrinkRequest) GoString() string {
	return s.String()
}

func (s *FaceCompareShrinkRequest) SetFaceRequestShrink(v string) *FaceCompareShrinkRequest {
	s.FaceRequestShrink = &v
	return s
}

type FaceCompareResponseBody struct {
	// example:
	//
	// OK
	Code *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *FaceCompareResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4ba***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FaceCompareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareResponseBody) GoString() string {
	return s.String()
}

func (s *FaceCompareResponseBody) SetCode(v int32) *FaceCompareResponseBody {
	s.Code = &v
	return s
}

func (s *FaceCompareResponseBody) SetData(v *FaceCompareResponseBodyData) *FaceCompareResponseBody {
	s.Data = v
	return s
}

func (s *FaceCompareResponseBody) SetHttpCode(v int32) *FaceCompareResponseBody {
	s.HttpCode = &v
	return s
}

func (s *FaceCompareResponseBody) SetMessage(v string) *FaceCompareResponseBody {
	s.Message = &v
	return s
}

func (s *FaceCompareResponseBody) SetRequestId(v string) *FaceCompareResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceCompareResponseBody) SetSuccess(v bool) *FaceCompareResponseBody {
	s.Success = &v
	return s
}

type FaceCompareResponseBodyData struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// T
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 60.86
	VerifyScore *float32 `json:"VerifyScore,omitempty" xml:"VerifyScore,omitempty"`
}

func (s FaceCompareResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareResponseBodyData) GoString() string {
	return s.String()
}

func (s *FaceCompareResponseBodyData) SetCode(v string) *FaceCompareResponseBodyData {
	s.Code = &v
	return s
}

func (s *FaceCompareResponseBodyData) SetId(v string) *FaceCompareResponseBodyData {
	s.Id = &v
	return s
}

func (s *FaceCompareResponseBodyData) SetMessage(v string) *FaceCompareResponseBodyData {
	s.Message = &v
	return s
}

func (s *FaceCompareResponseBodyData) SetPassed(v string) *FaceCompareResponseBodyData {
	s.Passed = &v
	return s
}

func (s *FaceCompareResponseBodyData) SetRequestId(v string) *FaceCompareResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *FaceCompareResponseBodyData) SetStatus(v string) *FaceCompareResponseBodyData {
	s.Status = &v
	return s
}

func (s *FaceCompareResponseBodyData) SetVerifyScore(v float32) *FaceCompareResponseBodyData {
	s.VerifyScore = &v
	return s
}

type FaceCompareResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceCompareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceCompareResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareResponse) GoString() string {
	return s.String()
}

func (s *FaceCompareResponse) SetHeaders(v map[string]*string) *FaceCompareResponse {
	s.Headers = v
	return s
}

func (s *FaceCompareResponse) SetStatusCode(v int32) *FaceCompareResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceCompareResponse) SetBody(v *FaceCompareResponseBody) *FaceCompareResponse {
	s.Body = v
	return s
}

type FaceLivenessRequest struct {
	FaceRequest *FaceLivenessRequestFaceRequest `json:"FaceRequest,omitempty" xml:"FaceRequest,omitempty" type:"Struct"`
}

func (s FaceLivenessRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessRequest) GoString() string {
	return s.String()
}

func (s *FaceLivenessRequest) SetFaceRequest(v *FaceLivenessRequestFaceRequest) *FaceLivenessRequest {
	s.FaceRequest = v
	return s
}

type FaceLivenessRequestFaceRequest struct {
	// example:
	//
	// d61be709-49d2-4cf1-b219-cd6181f72***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	SourceImage *string `json:"SourceImage,omitempty" xml:"SourceImage,omitempty"`
}

func (s FaceLivenessRequestFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessRequestFaceRequest) GoString() string {
	return s.String()
}

func (s *FaceLivenessRequestFaceRequest) SetAppId(v string) *FaceLivenessRequestFaceRequest {
	s.AppId = &v
	return s
}

func (s *FaceLivenessRequestFaceRequest) SetSourceImage(v string) *FaceLivenessRequestFaceRequest {
	s.SourceImage = &v
	return s
}

type FaceLivenessShrinkRequest struct {
	FaceRequestShrink *string `json:"FaceRequest,omitempty" xml:"FaceRequest,omitempty"`
}

func (s FaceLivenessShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessShrinkRequest) GoString() string {
	return s.String()
}

func (s *FaceLivenessShrinkRequest) SetFaceRequestShrink(v string) *FaceLivenessShrinkRequest {
	s.FaceRequestShrink = &v
	return s
}

type FaceLivenessResponseBody struct {
	// example:
	//
	// OK
	Code *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *FaceLivenessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7AC54DAF-38F8-58C7-8383-3131FEDDA***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FaceLivenessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessResponseBody) GoString() string {
	return s.String()
}

func (s *FaceLivenessResponseBody) SetCode(v int32) *FaceLivenessResponseBody {
	s.Code = &v
	return s
}

func (s *FaceLivenessResponseBody) SetData(v *FaceLivenessResponseBodyData) *FaceLivenessResponseBody {
	s.Data = v
	return s
}

func (s *FaceLivenessResponseBody) SetHttpCode(v int32) *FaceLivenessResponseBody {
	s.HttpCode = &v
	return s
}

func (s *FaceLivenessResponseBody) SetMessage(v string) *FaceLivenessResponseBody {
	s.Message = &v
	return s
}

func (s *FaceLivenessResponseBody) SetRequestId(v string) *FaceLivenessResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceLivenessResponseBody) SetSuccess(v bool) *FaceLivenessResponseBody {
	s.Success = &v
	return s
}

type FaceLivenessResponseBodyData struct {
	// example:
	//
	// 2000000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// T
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// d61be709-49d2-4cf1-b219-cd6181f72***
	PublicId *string `json:"PublicId,omitempty" xml:"PublicId,omitempty"`
	// example:
	//
	// 3.24324324324
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s FaceLivenessResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessResponseBodyData) GoString() string {
	return s.String()
}

func (s *FaceLivenessResponseBodyData) SetCode(v string) *FaceLivenessResponseBodyData {
	s.Code = &v
	return s
}

func (s *FaceLivenessResponseBodyData) SetMessage(v string) *FaceLivenessResponseBodyData {
	s.Message = &v
	return s
}

func (s *FaceLivenessResponseBodyData) SetPassed(v string) *FaceLivenessResponseBodyData {
	s.Passed = &v
	return s
}

func (s *FaceLivenessResponseBodyData) SetPublicId(v string) *FaceLivenessResponseBodyData {
	s.PublicId = &v
	return s
}

func (s *FaceLivenessResponseBodyData) SetScore(v float32) *FaceLivenessResponseBodyData {
	s.Score = &v
	return s
}

func (s *FaceLivenessResponseBodyData) SetStatus(v string) *FaceLivenessResponseBodyData {
	s.Status = &v
	return s
}

type FaceLivenessResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceLivenessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceLivenessResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceLivenessResponse) GoString() string {
	return s.String()
}

func (s *FaceLivenessResponse) SetHeaders(v map[string]*string) *FaceLivenessResponse {
	s.Headers = v
	return s
}

func (s *FaceLivenessResponse) SetStatusCode(v int32) *FaceLivenessResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceLivenessResponse) SetBody(v *FaceLivenessResponseBody) *FaceLivenessResponse {
	s.Body = v
	return s
}

type FaceRecognizeRequest struct {
	FaceRequest *FaceRecognizeRequestFaceRequest `json:"FaceRequest,omitempty" xml:"FaceRequest,omitempty" type:"Struct"`
}

func (s FaceRecognizeRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceRecognizeRequest) GoString() string {
	return s.String()
}

func (s *FaceRecognizeRequest) SetFaceRequest(v *FaceRecognizeRequestFaceRequest) *FaceRecognizeRequest {
	s.FaceRequest = v
	return s
}

type FaceRecognizeRequestFaceRequest struct {
	// example:
	//
	// d61be709-49d2-4cf1-b219-cd6181f72***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// true
	Liveness *bool `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	SourceImage *string `json:"SourceImage,omitempty" xml:"SourceImage,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	TargetImage *string `json:"TargetImage,omitempty" xml:"TargetImage,omitempty"`
}

func (s FaceRecognizeRequestFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceRecognizeRequestFaceRequest) GoString() string {
	return s.String()
}

func (s *FaceRecognizeRequestFaceRequest) SetAppId(v string) *FaceRecognizeRequestFaceRequest {
	s.AppId = &v
	return s
}

func (s *FaceRecognizeRequestFaceRequest) SetLiveness(v bool) *FaceRecognizeRequestFaceRequest {
	s.Liveness = &v
	return s
}

func (s *FaceRecognizeRequestFaceRequest) SetSourceImage(v string) *FaceRecognizeRequestFaceRequest {
	s.SourceImage = &v
	return s
}

func (s *FaceRecognizeRequestFaceRequest) SetTargetImage(v string) *FaceRecognizeRequestFaceRequest {
	s.TargetImage = &v
	return s
}

type FaceRecognizeShrinkRequest struct {
	FaceRequestShrink *string `json:"FaceRequest,omitempty" xml:"FaceRequest,omitempty"`
}

func (s FaceRecognizeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceRecognizeShrinkRequest) GoString() string {
	return s.String()
}

func (s *FaceRecognizeShrinkRequest) SetFaceRequestShrink(v string) *FaceRecognizeShrinkRequest {
	s.FaceRequestShrink = &v
	return s
}

type FaceRecognizeResponseBody struct {
	// example:
	//
	// OK
	Code *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *FaceRecognizeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4ba***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FaceRecognizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceRecognizeResponseBody) GoString() string {
	return s.String()
}

func (s *FaceRecognizeResponseBody) SetCode(v int32) *FaceRecognizeResponseBody {
	s.Code = &v
	return s
}

func (s *FaceRecognizeResponseBody) SetData(v *FaceRecognizeResponseBodyData) *FaceRecognizeResponseBody {
	s.Data = v
	return s
}

func (s *FaceRecognizeResponseBody) SetHttpCode(v int32) *FaceRecognizeResponseBody {
	s.HttpCode = &v
	return s
}

func (s *FaceRecognizeResponseBody) SetMessage(v string) *FaceRecognizeResponseBody {
	s.Message = &v
	return s
}

func (s *FaceRecognizeResponseBody) SetRequestId(v string) *FaceRecognizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceRecognizeResponseBody) SetSuccess(v bool) *FaceRecognizeResponseBody {
	s.Success = &v
	return s
}

type FaceRecognizeResponseBodyData struct {
	// example:
	//
	// T
	ComparePassed *string `json:"ComparePassed,omitempty" xml:"ComparePassed,omitempty"`
	// example:
	//
	// 32.435
	CompareScore *float32 `json:"CompareScore,omitempty" xml:"CompareScore,omitempty"`
	// example:
	//
	// T
	LivenessPassed *string `json:"LivenessPassed,omitempty" xml:"LivenessPassed,omitempty"`
	// example:
	//
	// 56.34
	LivenessScore *float32 `json:"LivenessScore,omitempty" xml:"LivenessScore,omitempty"`
}

func (s FaceRecognizeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FaceRecognizeResponseBodyData) GoString() string {
	return s.String()
}

func (s *FaceRecognizeResponseBodyData) SetComparePassed(v string) *FaceRecognizeResponseBodyData {
	s.ComparePassed = &v
	return s
}

func (s *FaceRecognizeResponseBodyData) SetCompareScore(v float32) *FaceRecognizeResponseBodyData {
	s.CompareScore = &v
	return s
}

func (s *FaceRecognizeResponseBodyData) SetLivenessPassed(v string) *FaceRecognizeResponseBodyData {
	s.LivenessPassed = &v
	return s
}

func (s *FaceRecognizeResponseBodyData) SetLivenessScore(v float32) *FaceRecognizeResponseBodyData {
	s.LivenessScore = &v
	return s
}

type FaceRecognizeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceRecognizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceRecognizeResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceRecognizeResponse) GoString() string {
	return s.String()
}

func (s *FaceRecognizeResponse) SetHeaders(v map[string]*string) *FaceRecognizeResponse {
	s.Headers = v
	return s
}

func (s *FaceRecognizeResponse) SetStatusCode(v int32) *FaceRecognizeResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceRecognizeResponse) SetBody(v *FaceRecognizeResponseBody) *FaceRecognizeResponse {
	s.Body = v
	return s
}

type GetAppRequest struct {
	// example:
	//
	// {"version":"1.0.0"}
	ClientBaseParam *string `json:"ClientBaseParam,omitempty" xml:"ClientBaseParam,omitempty"`
	// example:
	//
	// 1.0.002
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// xxx-xxx-xxx
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// com.a.test
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
}

func (s GetAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppRequest) GoString() string {
	return s.String()
}

func (s *GetAppRequest) SetClientBaseParam(v string) *GetAppRequest {
	s.ClientBaseParam = &v
	return s
}

func (s *GetAppRequest) SetClientVersion(v string) *GetAppRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetAppRequest) SetDeviceId(v string) *GetAppRequest {
	s.DeviceId = &v
	return s
}

func (s *GetAppRequest) SetId(v string) *GetAppRequest {
	s.Id = &v
	return s
}

func (s *GetAppRequest) SetPackageName(v string) *GetAppRequest {
	s.PackageName = &v
	return s
}

type GetAppResponseBody struct {
	// example:
	//
	// OK
	Code    *string                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                 `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppResponseBody) SetCode(v string) *GetAppResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppResponseBody) SetData(v *GetAppResponseBodyData) *GetAppResponseBody {
	s.Data = v
	return s
}

func (s *GetAppResponseBody) SetMessage(v string) *GetAppResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppResponseBody) SetRequestId(v string) *GetAppResponseBody {
	s.RequestId = &v
	return s
}

type GetAppResponseBodyData struct {
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 1
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// false
	Disabled *string `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// example:
	//
	// ff1d7783-e087-4d62-92df-3a163eca7c07
	FeeId *string `json:"FeeId,omitempty" xml:"FeeId,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyData) SetConfig(v string) *GetAppResponseBodyData {
	s.Config = &v
	return s
}

func (s *GetAppResponseBodyData) SetCreatedAt(v string) *GetAppResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetAppResponseBodyData) SetDisabled(v string) *GetAppResponseBodyData {
	s.Disabled = &v
	return s
}

func (s *GetAppResponseBodyData) SetFeeId(v string) *GetAppResponseBodyData {
	s.FeeId = &v
	return s
}

func (s *GetAppResponseBodyData) SetName(v string) *GetAppResponseBodyData {
	s.Name = &v
	return s
}

type GetAppResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponse) GoString() string {
	return s.String()
}

func (s *GetAppResponse) SetHeaders(v map[string]*string) *GetAppResponse {
	s.Headers = v
	return s
}

func (s *GetAppResponse) SetStatusCode(v int32) *GetAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppResponse) SetBody(v *GetAppResponseBody) *GetAppResponse {
	s.Body = v
	return s
}

type GetAsrResultRequest struct {
	// example:
	//
	// B2AADC9E-2A58-5918-AE4E-FF59E19D7***
	AsrTaskId *string `json:"AsrTaskId,omitempty" xml:"AsrTaskId,omitempty"`
}

func (s GetAsrResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsrResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsrResultRequest) SetAsrTaskId(v string) *GetAsrResultRequest {
	s.AsrTaskId = &v
	return s
}

type GetAsrResultResponseBody struct {
	// example:
	//
	// OK
	Code *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAsrResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C14ED32C-B9E4-54E7-BA85-C2B562C5B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAsrResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsrResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsrResultResponseBody) SetCode(v int32) *GetAsrResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsrResultResponseBody) SetData(v *GetAsrResultResponseBodyData) *GetAsrResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAsrResultResponseBody) SetHttpCode(v int32) *GetAsrResultResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetAsrResultResponseBody) SetMessage(v string) *GetAsrResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsrResultResponseBody) SetRequestId(v string) *GetAsrResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsrResultResponseBody) SetSuccess(v bool) *GetAsrResultResponseBody {
	s.Success = &v
	return s
}

type GetAsrResultResponseBodyData struct {
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetAsrResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsrResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsrResultResponseBodyData) SetResult(v string) *GetAsrResultResponseBodyData {
	s.Result = &v
	return s
}

type GetAsrResultResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsrResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsrResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsrResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsrResultResponse) SetHeaders(v map[string]*string) *GetAsrResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsrResultResponse) SetStatusCode(v int32) *GetAsrResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsrResultResponse) SetBody(v *GetAsrResultResponseBody) *GetAsrResultResponse {
	s.Body = v
	return s
}

type GetDepartmentRequest struct {
	// example:
	//
	// {"version":"1.0.0"}
	ClientBaseParam *string `json:"ClientBaseParam,omitempty" xml:"ClientBaseParam,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDepartmentRequest) GoString() string {
	return s.String()
}

func (s *GetDepartmentRequest) SetClientBaseParam(v string) *GetDepartmentRequest {
	s.ClientBaseParam = &v
	return s
}

func (s *GetDepartmentRequest) SetId(v string) *GetDepartmentRequest {
	s.Id = &v
	return s
}

type GetDepartmentResponseBody struct {
	// example:
	//
	// OK
	Code    *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetDepartmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDepartmentResponseBody) SetCode(v string) *GetDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetDepartmentResponseBody) SetData(v *GetDepartmentResponseBodyData) *GetDepartmentResponseBody {
	s.Data = v
	return s
}

func (s *GetDepartmentResponseBody) SetMessage(v string) *GetDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetDepartmentResponseBody) SetRequestId(v string) *GetDepartmentResponseBody {
	s.RequestId = &v
	return s
}

type GetDepartmentResponseBodyData struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s GetDepartmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDepartmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDepartmentResponseBodyData) SetCreatedAt(v string) *GetDepartmentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetDepartmentResponseBodyData) SetDescription(v string) *GetDepartmentResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetDepartmentResponseBodyData) SetId(v string) *GetDepartmentResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetDepartmentResponseBodyData) SetName(v string) *GetDepartmentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetDepartmentResponseBodyData) SetUpdatedAt(v string) *GetDepartmentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

type GetDepartmentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDepartmentResponse) GoString() string {
	return s.String()
}

func (s *GetDepartmentResponse) SetHeaders(v map[string]*string) *GetDepartmentResponse {
	s.Headers = v
	return s
}

func (s *GetDepartmentResponse) SetStatusCode(v int32) *GetDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDepartmentResponse) SetBody(v *GetDepartmentResponseBody) *GetDepartmentResponse {
	s.Body = v
	return s
}

type GetDetectProcessRequest struct {
	// example:
	//
	// {"version":"1.0.0"}
	ClientBaseParam *string `json:"ClientBaseParam,omitempty" xml:"ClientBaseParam,omitempty"`
	// ID
	//
	// example:
	//
	// 0f1c45cd-3eee-4e60-b505-2e330b87****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDetectProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessRequest) GoString() string {
	return s.String()
}

func (s *GetDetectProcessRequest) SetClientBaseParam(v string) *GetDetectProcessRequest {
	s.ClientBaseParam = &v
	return s
}

func (s *GetDetectProcessRequest) SetId(v string) *GetDetectProcessRequest {
	s.Id = &v
	return s
}

type GetDetectProcessResponseBody struct {
	// example:
	//
	// OK
	Code    *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetDetectProcessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0q1c45cd-3eee-1e60-b505-2e330b8755d3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDetectProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectProcessResponseBody) SetCode(v string) *GetDetectProcessResponseBody {
	s.Code = &v
	return s
}

func (s *GetDetectProcessResponseBody) SetData(v *GetDetectProcessResponseBodyData) *GetDetectProcessResponseBody {
	s.Data = v
	return s
}

func (s *GetDetectProcessResponseBody) SetMessage(v string) *GetDetectProcessResponseBody {
	s.Message = &v
	return s
}

func (s *GetDetectProcessResponseBody) SetRequestId(v string) *GetDetectProcessResponseBody {
	s.RequestId = &v
	return s
}

type GetDetectProcessResponseBodyData struct {
	// example:
	//
	// {}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2020-12-04T14:47:59.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// {}
	Draft *string `json:"Draft,omitempty" xml:"Draft,omitempty"`
	// ID
	//
	// example:
	//
	// 0f1c45cd-3eee-4e60-b505-2e330b8755d3
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 987d563d38f5aef27feca8702c689bb1
	Md5  *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	NewVersion *bool `json:"NewVersion,omitempty" xml:"NewVersion,omitempty"`
	// example:
	//
	// 2020-12-04T14:47:59.000+08:00
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s GetDetectProcessResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDetectProcessResponseBodyData) SetContent(v string) *GetDetectProcessResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetDetectProcessResponseBodyData) SetCreatedAt(v string) *GetDetectProcessResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetDetectProcessResponseBodyData) SetDraft(v string) *GetDetectProcessResponseBodyData {
	s.Draft = &v
	return s
}

func (s *GetDetectProcessResponseBodyData) SetId(v string) *GetDetectProcessResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetDetectProcessResponseBodyData) SetMd5(v string) *GetDetectProcessResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *GetDetectProcessResponseBodyData) SetName(v string) *GetDetectProcessResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetDetectProcessResponseBodyData) SetNewVersion(v bool) *GetDetectProcessResponseBodyData {
	s.NewVersion = &v
	return s
}

func (s *GetDetectProcessResponseBodyData) SetUpdatedAt(v string) *GetDetectProcessResponseBodyData {
	s.UpdatedAt = &v
	return s
}

type GetDetectProcessResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDetectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDetectProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessResponse) GoString() string {
	return s.String()
}

func (s *GetDetectProcessResponse) SetHeaders(v map[string]*string) *GetDetectProcessResponse {
	s.Headers = v
	return s
}

func (s *GetDetectProcessResponse) SetStatusCode(v int32) *GetDetectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDetectProcessResponse) SetBody(v *GetDetectProcessResponseBody) *GetDetectProcessResponse {
	s.Body = v
	return s
}

type GetDetectProcessJsonFileRequest struct {
	// example:
	//
	// {"version":"1.0.0"}
	ClientBaseParam *string `json:"ClientBaseParam,omitempty" xml:"ClientBaseParam,omitempty"`
	// ID
	//
	// example:
	//
	// hpsk3wdo-2020122319
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDetectProcessJsonFileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessJsonFileRequest) GoString() string {
	return s.String()
}

func (s *GetDetectProcessJsonFileRequest) SetClientBaseParam(v string) *GetDetectProcessJsonFileRequest {
	s.ClientBaseParam = &v
	return s
}

func (s *GetDetectProcessJsonFileRequest) SetId(v string) *GetDetectProcessJsonFileRequest {
	s.Id = &v
	return s
}

type GetDetectProcessJsonFileResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// -
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B2695011-0604-5557-9E00-B74F58AB3F2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDetectProcessJsonFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessJsonFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectProcessJsonFileResponseBody) SetCode(v string) *GetDetectProcessJsonFileResponseBody {
	s.Code = &v
	return s
}

func (s *GetDetectProcessJsonFileResponseBody) SetData(v string) *GetDetectProcessJsonFileResponseBody {
	s.Data = &v
	return s
}

func (s *GetDetectProcessJsonFileResponseBody) SetMessage(v string) *GetDetectProcessJsonFileResponseBody {
	s.Message = &v
	return s
}

func (s *GetDetectProcessJsonFileResponseBody) SetRequestId(v string) *GetDetectProcessJsonFileResponseBody {
	s.RequestId = &v
	return s
}

type GetDetectProcessJsonFileResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDetectProcessJsonFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDetectProcessJsonFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessJsonFileResponse) GoString() string {
	return s.String()
}

func (s *GetDetectProcessJsonFileResponse) SetHeaders(v map[string]*string) *GetDetectProcessJsonFileResponse {
	s.Headers = v
	return s
}

func (s *GetDetectProcessJsonFileResponse) SetStatusCode(v int32) *GetDetectProcessJsonFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDetectProcessJsonFileResponse) SetBody(v *GetDetectProcessJsonFileResponseBody) *GetDetectProcessJsonFileResponse {
	s.Body = v
	return s
}

type GetDetectionRequest struct {
	// example:
	//
	// {"version":"1.0.0"}
	ClientBaseParam *string `json:"ClientBaseParam,omitempty" xml:"ClientBaseParam,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDetectionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionRequest) GoString() string {
	return s.String()
}

func (s *GetDetectionRequest) SetClientBaseParam(v string) *GetDetectionRequest {
	s.ClientBaseParam = &v
	return s
}

func (s *GetDetectionRequest) SetId(v string) *GetDetectionRequest {
	s.Id = &v
	return s
}

type GetDetectionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *GetDetectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDetectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectionResponseBody) SetCode(v string) *GetDetectionResponseBody {
	s.Code = &v
	return s
}

func (s *GetDetectionResponseBody) SetData(v *GetDetectionResponseBodyData) *GetDetectionResponseBody {
	s.Data = v
	return s
}

func (s *GetDetectionResponseBody) SetMessage(v string) *GetDetectionResponseBody {
	s.Message = &v
	return s
}

func (s *GetDetectionResponseBody) SetRequestId(v string) *GetDetectionResponseBody {
	s.RequestId = &v
	return s
}

type GetDetectionResponseBodyData struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	DepartmentId   *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// local
	RecordingType *string `json:"RecordingType,omitempty" xml:"RecordingType,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// default
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// runnable
	Status *string                              `json:"Status,omitempty" xml:"Status,omitempty"`
	Tasks  []*GetDetectionResponseBodyDataTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s GetDetectionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDetectionResponseBodyData) SetCreatedAt(v string) *GetDetectionResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetDetectionResponseBodyData) SetDepartmentId(v string) *GetDetectionResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetDetectionResponseBodyData) SetDepartmentName(v string) *GetDetectionResponseBodyData {
	s.DepartmentName = &v
	return s
}

func (s *GetDetectionResponseBodyData) SetId(v string) *GetDetectionResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetDetectionResponseBodyData) SetRecordingType(v string) *GetDetectionResponseBodyData {
	s.RecordingType = &v
	return s
}

func (s *GetDetectionResponseBodyData) SetRuleId(v string) *GetDetectionResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *GetDetectionResponseBodyData) SetRuleName(v string) *GetDetectionResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetDetectionResponseBodyData) SetStatus(v string) *GetDetectionResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDetectionResponseBodyData) SetTasks(v []*GetDetectionResponseBodyDataTasks) *GetDetectionResponseBodyData {
	s.Tasks = v
	return s
}

type GetDetectionResponseBodyDataTasks struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// runnable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// http://oss.aliyun.com/1.mp4.meta
	VideoMetaUrl *string `json:"VideoMetaUrl,omitempty" xml:"VideoMetaUrl,omitempty"`
	// example:
	//
	// http://oss.aliyun.com/1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GetDetectionResponseBodyDataTasks) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionResponseBodyDataTasks) GoString() string {
	return s.String()
}

func (s *GetDetectionResponseBodyDataTasks) SetCreatedAt(v string) *GetDetectionResponseBodyDataTasks {
	s.CreatedAt = &v
	return s
}

func (s *GetDetectionResponseBodyDataTasks) SetId(v string) *GetDetectionResponseBodyDataTasks {
	s.Id = &v
	return s
}

func (s *GetDetectionResponseBodyDataTasks) SetStatus(v string) *GetDetectionResponseBodyDataTasks {
	s.Status = &v
	return s
}

func (s *GetDetectionResponseBodyDataTasks) SetVideoMetaUrl(v string) *GetDetectionResponseBodyDataTasks {
	s.VideoMetaUrl = &v
	return s
}

func (s *GetDetectionResponseBodyDataTasks) SetVideoUrl(v string) *GetDetectionResponseBodyDataTasks {
	s.VideoUrl = &v
	return s
}

type GetDetectionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDetectionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionResponse) GoString() string {
	return s.String()
}

func (s *GetDetectionResponse) SetHeaders(v map[string]*string) *GetDetectionResponse {
	s.Headers = v
	return s
}

func (s *GetDetectionResponse) SetStatusCode(v int32) *GetDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDetectionResponse) SetBody(v *GetDetectionResponseBody) *GetDetectionResponse {
	s.Body = v
	return s
}

type GetPreSignedUrlRequest struct {
	// example:
	//
	// {"version":"1.0.0"}
	ClientBaseParam *string `json:"ClientBaseParam,omitempty" xml:"ClientBaseParam,omitempty"`
	// example:
	//
	// 1.0.001
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// test.mp4
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s GetPreSignedUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPreSignedUrlRequest) GoString() string {
	return s.String()
}

func (s *GetPreSignedUrlRequest) SetClientBaseParam(v string) *GetPreSignedUrlRequest {
	s.ClientBaseParam = &v
	return s
}

func (s *GetPreSignedUrlRequest) SetClientVersion(v string) *GetPreSignedUrlRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetPreSignedUrlRequest) SetPrefix(v string) *GetPreSignedUrlRequest {
	s.Prefix = &v
	return s
}

type GetPreSignedUrlResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// http://shuanglu-pre.oss-cn-beijing.aliyuncs.com/13ba4081-84f3-42b0-af93-10a64319f8ef/test.txt
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPreSignedUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPreSignedUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetPreSignedUrlResponseBody) SetCode(v string) *GetPreSignedUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetPreSignedUrlResponseBody) SetData(v string) *GetPreSignedUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetPreSignedUrlResponseBody) SetMessage(v string) *GetPreSignedUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetPreSignedUrlResponseBody) SetRequestId(v string) *GetPreSignedUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetPreSignedUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPreSignedUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPreSignedUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPreSignedUrlResponse) GoString() string {
	return s.String()
}

func (s *GetPreSignedUrlResponse) SetHeaders(v map[string]*string) *GetPreSignedUrlResponse {
	s.Headers = v
	return s
}

func (s *GetPreSignedUrlResponse) SetStatusCode(v int32) *GetPreSignedUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPreSignedUrlResponse) SetBody(v *GetPreSignedUrlResponseBody) *GetPreSignedUrlResponse {
	s.Body = v
	return s
}

type GetRecordResultRequest struct {
	// example:
	//
	// {"version":"1.0.0"}
	ClientBaseParam *string `json:"ClientBaseParam,omitempty" xml:"ClientBaseParam,omitempty"`
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s GetRecordResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecordResultRequest) GoString() string {
	return s.String()
}

func (s *GetRecordResultRequest) SetClientBaseParam(v string) *GetRecordResultRequest {
	s.ClientBaseParam = &v
	return s
}

func (s *GetRecordResultRequest) SetRecordId(v string) *GetRecordResultRequest {
	s.RecordId = &v
	return s
}

type GetRecordResultResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code   *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data   *GetRecordResultResponseBodyData     `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Errors []*GetRecordResultResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// 84118BF0-56F7-54D2-8C1A-35BBBBD5016B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRecordResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecordResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordResultResponseBody) SetCode(v string) *GetRecordResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetRecordResultResponseBody) SetData(v *GetRecordResultResponseBodyData) *GetRecordResultResponseBody {
	s.Data = v
	return s
}

func (s *GetRecordResultResponseBody) SetErrors(v []*GetRecordResultResponseBodyErrors) *GetRecordResultResponseBody {
	s.Errors = v
	return s
}

func (s *GetRecordResultResponseBody) SetHttpCode(v int32) *GetRecordResultResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetRecordResultResponseBody) SetMessage(v string) *GetRecordResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetRecordResultResponseBody) SetRequestId(v string) *GetRecordResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecordResultResponseBody) SetSuccess(v bool) *GetRecordResultResponseBody {
	s.Success = &v
	return s
}

type GetRecordResultResponseBodyData struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt         *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DepartmentName    *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	DetectProcessName *string `json:"DetectProcessName,omitempty" xml:"DetectProcessName,omitempty"`
	// example:
	//
	// 22
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 0f1c45cd-3eee-4e60-b505-2e330b8755d3
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.meta
	MetaUrl *string `json:"MetaUrl,omitempty" xml:"MetaUrl,omitempty"`
	// example:
	//
	// ads32efef43
	OuterBusinessId *string `json:"OuterBusinessId,omitempty" xml:"OuterBusinessId,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	RecordAt       *string                                          `json:"RecordAt,omitempty" xml:"RecordAt,omitempty"`
	RecordRoomList []*GetRecordResultResponseBodyDataRecordRoomList `json:"RecordRoomList,omitempty" xml:"RecordRoomList,omitempty" type:"Repeated"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.json
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
	// example:
	//
	// 641981583353***
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GetRecordResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRecordResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRecordResultResponseBodyData) SetAppName(v string) *GetRecordResultResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetRecordResultResponseBodyData) SetCreatedAt(v string) *GetRecordResultResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetRecordResultResponseBodyData) SetDepartmentName(v string) *GetRecordResultResponseBodyData {
	s.DepartmentName = &v
	return s
}

func (s *GetRecordResultResponseBodyData) SetDetectProcessName(v string) *GetRecordResultResponseBodyData {
	s.DetectProcessName = &v
	return s
}

func (s *GetRecordResultResponseBodyData) SetDuration(v int64) *GetRecordResultResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetRecordResultResponseBodyData) SetId(v string) *GetRecordResultResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetRecordResultResponseBodyData) SetMetaUrl(v string) *GetRecordResultResponseBodyData {
	s.MetaUrl = &v
	return s
}

func (s *GetRecordResultResponseBodyData) SetOuterBusinessId(v string) *GetRecordResultResponseBodyData {
	s.OuterBusinessId = &v
	return s
}

func (s *GetRecordResultResponseBodyData) SetRecordAt(v string) *GetRecordResultResponseBodyData {
	s.RecordAt = &v
	return s
}

func (s *GetRecordResultResponseBodyData) SetRecordRoomList(v []*GetRecordResultResponseBodyDataRecordRoomList) *GetRecordResultResponseBodyData {
	s.RecordRoomList = v
	return s
}

func (s *GetRecordResultResponseBodyData) SetResultUrl(v string) *GetRecordResultResponseBodyData {
	s.ResultUrl = &v
	return s
}

func (s *GetRecordResultResponseBodyData) SetRoomId(v string) *GetRecordResultResponseBodyData {
	s.RoomId = &v
	return s
}

func (s *GetRecordResultResponseBodyData) SetStatus(v string) *GetRecordResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetRecordResultResponseBodyData) SetVideoUrl(v string) *GetRecordResultResponseBodyData {
	s.VideoUrl = &v
	return s
}

type GetRecordResultResponseBodyDataRecordRoomList struct {
	// example:
	//
	// ads32efef43
	OuterBusinessId *string `json:"OuterBusinessId,omitempty" xml:"OuterBusinessId,omitempty"`
	// example:
	//
	// Mix
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// example:
	//
	// 21343
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.meta
	RoomMetaUrl *string `json:"RoomMetaUrl,omitempty" xml:"RoomMetaUrl,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	RoomRecordAt *string `json:"RoomRecordAt,omitempty" xml:"RoomRecordAt,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.json
	RoomResultUrl *string `json:"RoomResultUrl,omitempty" xml:"RoomResultUrl,omitempty"`
	// example:
	//
	// runnable
	RoomStatus *string `json:"RoomStatus,omitempty" xml:"RoomStatus,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4
	RoomVideoUrl *string `json:"RoomVideoUrl,omitempty" xml:"RoomVideoUrl,omitempty"`
	// example:
	//
	// record_65703154805715668342
	RtcRecordId *string `json:"RtcRecordId,omitempty" xml:"RtcRecordId,omitempty"`
}

func (s GetRecordResultResponseBodyDataRecordRoomList) String() string {
	return tea.Prettify(s)
}

func (s GetRecordResultResponseBodyDataRecordRoomList) GoString() string {
	return s.String()
}

func (s *GetRecordResultResponseBodyDataRecordRoomList) SetOuterBusinessId(v string) *GetRecordResultResponseBodyDataRecordRoomList {
	s.OuterBusinessId = &v
	return s
}

func (s *GetRecordResultResponseBodyDataRecordRoomList) SetRecordType(v string) *GetRecordResultResponseBodyDataRecordRoomList {
	s.RecordType = &v
	return s
}

func (s *GetRecordResultResponseBodyDataRecordRoomList) SetRole(v string) *GetRecordResultResponseBodyDataRecordRoomList {
	s.Role = &v
	return s
}

func (s *GetRecordResultResponseBodyDataRecordRoomList) SetRoomMetaUrl(v string) *GetRecordResultResponseBodyDataRecordRoomList {
	s.RoomMetaUrl = &v
	return s
}

func (s *GetRecordResultResponseBodyDataRecordRoomList) SetRoomRecordAt(v string) *GetRecordResultResponseBodyDataRecordRoomList {
	s.RoomRecordAt = &v
	return s
}

func (s *GetRecordResultResponseBodyDataRecordRoomList) SetRoomResultUrl(v string) *GetRecordResultResponseBodyDataRecordRoomList {
	s.RoomResultUrl = &v
	return s
}

func (s *GetRecordResultResponseBodyDataRecordRoomList) SetRoomStatus(v string) *GetRecordResultResponseBodyDataRecordRoomList {
	s.RoomStatus = &v
	return s
}

func (s *GetRecordResultResponseBodyDataRecordRoomList) SetRoomVideoUrl(v string) *GetRecordResultResponseBodyDataRecordRoomList {
	s.RoomVideoUrl = &v
	return s
}

func (s *GetRecordResultResponseBodyDataRecordRoomList) SetRtcRecordId(v string) *GetRecordResultResponseBodyDataRecordRoomList {
	s.RtcRecordId = &v
	return s
}

type GetRecordResultResponseBodyErrors struct {
	// example:
	//
	// 0f1c45cd-3eee-4e60-b505-2e330b8755d3
	Field   *string `json:"Field,omitempty" xml:"Field,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetRecordResultResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s GetRecordResultResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *GetRecordResultResponseBodyErrors) SetField(v string) *GetRecordResultResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *GetRecordResultResponseBodyErrors) SetMessage(v string) *GetRecordResultResponseBodyErrors {
	s.Message = &v
	return s
}

type GetRecordResultResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecordResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecordResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecordResultResponse) GoString() string {
	return s.String()
}

func (s *GetRecordResultResponse) SetHeaders(v map[string]*string) *GetRecordResultResponse {
	s.Headers = v
	return s
}

func (s *GetRecordResultResponse) SetStatusCode(v int32) *GetRecordResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecordResultResponse) SetBody(v *GetRecordResultResponseBody) *GetRecordResultResponse {
	s.Body = v
	return s
}

type GetRecordsByFeeIdRequest struct {
	// example:
	//
	// 6c94f2a7-632d-4ea0-aa06-a97800a9060f
	FeeId *string `json:"FeeId,omitempty" xml:"FeeId,omitempty"`
}

func (s GetRecordsByFeeIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsByFeeIdRequest) GoString() string {
	return s.String()
}

func (s *GetRecordsByFeeIdRequest) SetFeeId(v string) *GetRecordsByFeeIdRequest {
	s.FeeId = &v
	return s
}

type GetRecordsByFeeIdResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code   *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data   []*GetRecordsByFeeIdResponseBodyData   `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Errors []*GetRecordsByFeeIdResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// 0B576AAB-A638-5029-9A54-A7C1DB5AC0B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRecordsByFeeIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsByFeeIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordsByFeeIdResponseBody) SetCode(v string) *GetRecordsByFeeIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBody) SetData(v []*GetRecordsByFeeIdResponseBodyData) *GetRecordsByFeeIdResponseBody {
	s.Data = v
	return s
}

func (s *GetRecordsByFeeIdResponseBody) SetErrors(v []*GetRecordsByFeeIdResponseBodyErrors) *GetRecordsByFeeIdResponseBody {
	s.Errors = v
	return s
}

func (s *GetRecordsByFeeIdResponseBody) SetHttpCode(v int32) *GetRecordsByFeeIdResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBody) SetMessage(v string) *GetRecordsByFeeIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBody) SetRequestId(v string) *GetRecordsByFeeIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBody) SetSuccess(v bool) *GetRecordsByFeeIdResponseBody {
	s.Success = &v
	return s
}

type GetRecordsByFeeIdResponseBodyData struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt         *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DepartmentName    *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	DetectProcessName *string `json:"DetectProcessName,omitempty" xml:"DetectProcessName,omitempty"`
	// example:
	//
	// 22
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 63bb629d-92bf-4cdc-ad0b-3032c926d23f
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.meta
	MetaUrl *string `json:"MetaUrl,omitempty" xml:"MetaUrl,omitempty"`
	// example:
	//
	// ads32efef43
	OuterBusinessId *string `json:"OuterBusinessId,omitempty" xml:"OuterBusinessId,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	RecordAt       *string                                            `json:"RecordAt,omitempty" xml:"RecordAt,omitempty"`
	RecordRoomList []*GetRecordsByFeeIdResponseBodyDataRecordRoomList `json:"RecordRoomList,omitempty" xml:"RecordRoomList,omitempty" type:"Repeated"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.json
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
	// example:
	//
	// 654078150345590
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GetRecordsByFeeIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsByFeeIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRecordsByFeeIdResponseBodyData) SetAppName(v string) *GetRecordsByFeeIdResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyData) SetCreatedAt(v string) *GetRecordsByFeeIdResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyData) SetDepartmentName(v string) *GetRecordsByFeeIdResponseBodyData {
	s.DepartmentName = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyData) SetDetectProcessName(v string) *GetRecordsByFeeIdResponseBodyData {
	s.DetectProcessName = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyData) SetDuration(v int64) *GetRecordsByFeeIdResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyData) SetId(v string) *GetRecordsByFeeIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyData) SetMetaUrl(v string) *GetRecordsByFeeIdResponseBodyData {
	s.MetaUrl = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyData) SetOuterBusinessId(v string) *GetRecordsByFeeIdResponseBodyData {
	s.OuterBusinessId = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyData) SetRecordAt(v string) *GetRecordsByFeeIdResponseBodyData {
	s.RecordAt = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyData) SetRecordRoomList(v []*GetRecordsByFeeIdResponseBodyDataRecordRoomList) *GetRecordsByFeeIdResponseBodyData {
	s.RecordRoomList = v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyData) SetResultUrl(v string) *GetRecordsByFeeIdResponseBodyData {
	s.ResultUrl = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyData) SetRoomId(v string) *GetRecordsByFeeIdResponseBodyData {
	s.RoomId = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyData) SetStatus(v string) *GetRecordsByFeeIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyData) SetVideoUrl(v string) *GetRecordsByFeeIdResponseBodyData {
	s.VideoUrl = &v
	return s
}

type GetRecordsByFeeIdResponseBodyDataRecordRoomList struct {
	// example:
	//
	// ads32efef43
	OuterBusinessId *string `json:"OuterBusinessId,omitempty" xml:"OuterBusinessId,omitempty"`
	// example:
	//
	// Mix
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// example:
	//
	// 21343
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.meta
	RoomMetaUrl *string `json:"RoomMetaUrl,omitempty" xml:"RoomMetaUrl,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	RoomRecordAt *string `json:"RoomRecordAt,omitempty" xml:"RoomRecordAt,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.json
	RoomResultUrl *string `json:"RoomResultUrl,omitempty" xml:"RoomResultUrl,omitempty"`
	// example:
	//
	// runnable
	RoomStatus *string `json:"RoomStatus,omitempty" xml:"RoomStatus,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4
	RoomVideoUrl *string `json:"RoomVideoUrl,omitempty" xml:"RoomVideoUrl,omitempty"`
	// example:
	//
	// record_65703154805715668342
	RtcRecordId *string `json:"RtcRecordId,omitempty" xml:"RtcRecordId,omitempty"`
}

func (s GetRecordsByFeeIdResponseBodyDataRecordRoomList) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsByFeeIdResponseBodyDataRecordRoomList) GoString() string {
	return s.String()
}

func (s *GetRecordsByFeeIdResponseBodyDataRecordRoomList) SetOuterBusinessId(v string) *GetRecordsByFeeIdResponseBodyDataRecordRoomList {
	s.OuterBusinessId = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyDataRecordRoomList) SetRecordType(v string) *GetRecordsByFeeIdResponseBodyDataRecordRoomList {
	s.RecordType = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyDataRecordRoomList) SetRole(v string) *GetRecordsByFeeIdResponseBodyDataRecordRoomList {
	s.Role = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyDataRecordRoomList) SetRoomMetaUrl(v string) *GetRecordsByFeeIdResponseBodyDataRecordRoomList {
	s.RoomMetaUrl = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyDataRecordRoomList) SetRoomRecordAt(v string) *GetRecordsByFeeIdResponseBodyDataRecordRoomList {
	s.RoomRecordAt = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyDataRecordRoomList) SetRoomResultUrl(v string) *GetRecordsByFeeIdResponseBodyDataRecordRoomList {
	s.RoomResultUrl = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyDataRecordRoomList) SetRoomStatus(v string) *GetRecordsByFeeIdResponseBodyDataRecordRoomList {
	s.RoomStatus = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyDataRecordRoomList) SetRoomVideoUrl(v string) *GetRecordsByFeeIdResponseBodyDataRecordRoomList {
	s.RoomVideoUrl = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyDataRecordRoomList) SetRtcRecordId(v string) *GetRecordsByFeeIdResponseBodyDataRecordRoomList {
	s.RtcRecordId = &v
	return s
}

type GetRecordsByFeeIdResponseBodyErrors struct {
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetRecordsByFeeIdResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsByFeeIdResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *GetRecordsByFeeIdResponseBodyErrors) SetField(v string) *GetRecordsByFeeIdResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *GetRecordsByFeeIdResponseBodyErrors) SetMessage(v string) *GetRecordsByFeeIdResponseBodyErrors {
	s.Message = &v
	return s
}

type GetRecordsByFeeIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecordsByFeeIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecordsByFeeIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsByFeeIdResponse) GoString() string {
	return s.String()
}

func (s *GetRecordsByFeeIdResponse) SetHeaders(v map[string]*string) *GetRecordsByFeeIdResponse {
	s.Headers = v
	return s
}

func (s *GetRecordsByFeeIdResponse) SetStatusCode(v int32) *GetRecordsByFeeIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecordsByFeeIdResponse) SetBody(v *GetRecordsByFeeIdResponseBody) *GetRecordsByFeeIdResponse {
	s.Body = v
	return s
}

type GetRecordsByOuterBusinessIdRequest struct {
	// example:
	//
	// ads32efef43
	OuterBusinessId *string `json:"OuterBusinessId,omitempty" xml:"OuterBusinessId,omitempty"`
}

func (s GetRecordsByOuterBusinessIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsByOuterBusinessIdRequest) GoString() string {
	return s.String()
}

func (s *GetRecordsByOuterBusinessIdRequest) SetOuterBusinessId(v string) *GetRecordsByOuterBusinessIdRequest {
	s.OuterBusinessId = &v
	return s
}

type GetRecordsByOuterBusinessIdResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code   *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data   []*GetRecordsByOuterBusinessIdResponseBodyData   `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Errors []*GetRecordsByOuterBusinessIdResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// 0f1c45cd-3eee-4e60-b505-2e330b875***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRecordsByOuterBusinessIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsByOuterBusinessIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordsByOuterBusinessIdResponseBody) SetCode(v string) *GetRecordsByOuterBusinessIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBody) SetData(v []*GetRecordsByOuterBusinessIdResponseBodyData) *GetRecordsByOuterBusinessIdResponseBody {
	s.Data = v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBody) SetErrors(v []*GetRecordsByOuterBusinessIdResponseBodyErrors) *GetRecordsByOuterBusinessIdResponseBody {
	s.Errors = v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBody) SetHttpCode(v int32) *GetRecordsByOuterBusinessIdResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBody) SetMessage(v string) *GetRecordsByOuterBusinessIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBody) SetRequestId(v string) *GetRecordsByOuterBusinessIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBody) SetSuccess(v bool) *GetRecordsByOuterBusinessIdResponseBody {
	s.Success = &v
	return s
}

type GetRecordsByOuterBusinessIdResponseBodyData struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt         *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DepartmentName    *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	DetectProcessName *string `json:"DetectProcessName,omitempty" xml:"DetectProcessName,omitempty"`
	// example:
	//
	// 22
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 0f1c45cd-3eee-4e60-b505-2e330b875***
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.meta
	MetaUrl *string `json:"MetaUrl,omitempty" xml:"MetaUrl,omitempty"`
	// example:
	//
	// ads32efef43
	OuterBusinessId *string `json:"OuterBusinessId,omitempty" xml:"OuterBusinessId,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	RecordAt       *string                                                      `json:"RecordAt,omitempty" xml:"RecordAt,omitempty"`
	RecordRoomList []*GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList `json:"RecordRoomList,omitempty" xml:"RecordRoomList,omitempty" type:"Repeated"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.json
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
	// example:
	//
	// 641981583353***
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GetRecordsByOuterBusinessIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsByOuterBusinessIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRecordsByOuterBusinessIdResponseBodyData) SetAppName(v string) *GetRecordsByOuterBusinessIdResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyData) SetCreatedAt(v string) *GetRecordsByOuterBusinessIdResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyData) SetDepartmentName(v string) *GetRecordsByOuterBusinessIdResponseBodyData {
	s.DepartmentName = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyData) SetDetectProcessName(v string) *GetRecordsByOuterBusinessIdResponseBodyData {
	s.DetectProcessName = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyData) SetDuration(v int64) *GetRecordsByOuterBusinessIdResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyData) SetId(v string) *GetRecordsByOuterBusinessIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyData) SetMetaUrl(v string) *GetRecordsByOuterBusinessIdResponseBodyData {
	s.MetaUrl = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyData) SetOuterBusinessId(v string) *GetRecordsByOuterBusinessIdResponseBodyData {
	s.OuterBusinessId = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyData) SetRecordAt(v string) *GetRecordsByOuterBusinessIdResponseBodyData {
	s.RecordAt = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyData) SetRecordRoomList(v []*GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList) *GetRecordsByOuterBusinessIdResponseBodyData {
	s.RecordRoomList = v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyData) SetResultUrl(v string) *GetRecordsByOuterBusinessIdResponseBodyData {
	s.ResultUrl = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyData) SetRoomId(v string) *GetRecordsByOuterBusinessIdResponseBodyData {
	s.RoomId = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyData) SetStatus(v string) *GetRecordsByOuterBusinessIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyData) SetVideoUrl(v string) *GetRecordsByOuterBusinessIdResponseBodyData {
	s.VideoUrl = &v
	return s
}

type GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList struct {
	// example:
	//
	// ads32efef43
	OuterBusinessId *string `json:"OuterBusinessId,omitempty" xml:"OuterBusinessId,omitempty"`
	// example:
	//
	// Mix
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.meta
	RoomMetaUrl *string `json:"RoomMetaUrl,omitempty" xml:"RoomMetaUrl,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	RoomRecordAt *string `json:"RoomRecordAt,omitempty" xml:"RoomRecordAt,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.json
	RoomResultUrl *string `json:"RoomResultUrl,omitempty" xml:"RoomResultUrl,omitempty"`
	// example:
	//
	// 1
	RoomStatus *string `json:"RoomStatus,omitempty" xml:"RoomStatus,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4
	RoomVideoUrl *string `json:"RoomVideoUrl,omitempty" xml:"RoomVideoUrl,omitempty"`
	// example:
	//
	// record_65703154805715668342
	RtcRecordId *string `json:"RtcRecordId,omitempty" xml:"RtcRecordId,omitempty"`
}

func (s GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList) GoString() string {
	return s.String()
}

func (s *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList) SetOuterBusinessId(v string) *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList {
	s.OuterBusinessId = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList) SetRecordType(v string) *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList {
	s.RecordType = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList) SetRole(v string) *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList {
	s.Role = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList) SetRoomMetaUrl(v string) *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList {
	s.RoomMetaUrl = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList) SetRoomRecordAt(v string) *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList {
	s.RoomRecordAt = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList) SetRoomResultUrl(v string) *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList {
	s.RoomResultUrl = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList) SetRoomStatus(v string) *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList {
	s.RoomStatus = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList) SetRoomVideoUrl(v string) *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList {
	s.RoomVideoUrl = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList) SetRtcRecordId(v string) *GetRecordsByOuterBusinessIdResponseBodyDataRecordRoomList {
	s.RtcRecordId = &v
	return s
}

type GetRecordsByOuterBusinessIdResponseBodyErrors struct {
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0***
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetRecordsByOuterBusinessIdResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsByOuterBusinessIdResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *GetRecordsByOuterBusinessIdResponseBodyErrors) SetField(v string) *GetRecordsByOuterBusinessIdResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponseBodyErrors) SetMessage(v string) *GetRecordsByOuterBusinessIdResponseBodyErrors {
	s.Message = &v
	return s
}

type GetRecordsByOuterBusinessIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecordsByOuterBusinessIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecordsByOuterBusinessIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsByOuterBusinessIdResponse) GoString() string {
	return s.String()
}

func (s *GetRecordsByOuterBusinessIdResponse) SetHeaders(v map[string]*string) *GetRecordsByOuterBusinessIdResponse {
	s.Headers = v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponse) SetStatusCode(v int32) *GetRecordsByOuterBusinessIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecordsByOuterBusinessIdResponse) SetBody(v *GetRecordsByOuterBusinessIdResponseBody) *GetRecordsByOuterBusinessIdResponse {
	s.Body = v
	return s
}

type GetRuleRequest struct {
	// example:
	//
	// {"version":"1.0.0"}
	ClientBaseParam *string `json:"ClientBaseParam,omitempty" xml:"ClientBaseParam,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRuleRequest) GoString() string {
	return s.String()
}

func (s *GetRuleRequest) SetClientBaseParam(v string) *GetRuleRequest {
	s.ClientBaseParam = &v
	return s
}

func (s *GetRuleRequest) SetId(v string) *GetRuleRequest {
	s.Id = &v
	return s
}

type GetRuleResponseBody struct {
	// example:
	//
	// OK
	Code    *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBody) SetCode(v string) *GetRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleResponseBody) SetData(v *GetRuleResponseBodyData) *GetRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleResponseBody) SetMessage(v string) *GetRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleResponseBody) SetRequestId(v string) *GetRuleResponseBody {
	s.RequestId = &v
	return s
}

type GetRuleResponseBodyData struct {
	// example:
	//
	// {}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyData) SetContent(v string) *GetRuleResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetRuleResponseBodyData) SetCreatedAt(v string) *GetRuleResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetRuleResponseBodyData) SetId(v string) *GetRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetRuleResponseBodyData) SetName(v string) *GetRuleResponseBodyData {
	s.Name = &v
	return s
}

type GetRuleResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponse) GoString() string {
	return s.String()
}

func (s *GetRuleResponse) SetHeaders(v map[string]*string) *GetRuleResponse {
	s.Headers = v
	return s
}

func (s *GetRuleResponse) SetStatusCode(v int32) *GetRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuleResponse) SetBody(v *GetRuleResponseBody) *GetRuleResponse {
	s.Body = v
	return s
}

type GetStatisticsRecordsByFeeIdRequest struct {
	// example:
	//
	// ab7ce59f-a68a-4a6f-82a6-6460fadc9a70
	FeeId *string `json:"FeeId,omitempty" xml:"FeeId,omitempty"`
}

func (s GetStatisticsRecordsByFeeIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStatisticsRecordsByFeeIdRequest) GoString() string {
	return s.String()
}

func (s *GetStatisticsRecordsByFeeIdRequest) SetFeeId(v string) *GetStatisticsRecordsByFeeIdRequest {
	s.FeeId = &v
	return s
}

type GetStatisticsRecordsByFeeIdResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code   *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data   []*GetStatisticsRecordsByFeeIdResponseBodyData   `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Errors []*GetStatisticsRecordsByFeeIdResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// CAD43A24-D34C-5848-BEB0-3D184F6687B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStatisticsRecordsByFeeIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStatisticsRecordsByFeeIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetStatisticsRecordsByFeeIdResponseBody) SetCode(v string) *GetStatisticsRecordsByFeeIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBody) SetData(v []*GetStatisticsRecordsByFeeIdResponseBodyData) *GetStatisticsRecordsByFeeIdResponseBody {
	s.Data = v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBody) SetErrors(v []*GetStatisticsRecordsByFeeIdResponseBodyErrors) *GetStatisticsRecordsByFeeIdResponseBody {
	s.Errors = v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBody) SetHttpCode(v int32) *GetStatisticsRecordsByFeeIdResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBody) SetMessage(v string) *GetStatisticsRecordsByFeeIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBody) SetRequestId(v string) *GetStatisticsRecordsByFeeIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBody) SetSuccess(v bool) *GetStatisticsRecordsByFeeIdResponseBody {
	s.Success = &v
	return s
}

type GetStatisticsRecordsByFeeIdResponseBodyData struct {
	// appid
	//
	// example:
	//
	// 12
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 2022-05-23 10:33:12
	BeginAt *string `json:"BeginAt,omitempty" xml:"BeginAt,omitempty"`
	// example:
	//
	// 12
	ChargeDuration *int32 `json:"ChargeDuration,omitempty" xml:"ChargeDuration,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 12
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// example:
	//
	// 1200
	DetectionDuration *int32 `json:"DetectionDuration,omitempty" xml:"DetectionDuration,omitempty"`
	// example:
	//
	// 7f3dfb9aa0dd0261
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 0
	DeviceType *int32 `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 2022-03-24 11:39:46
	EndAt *string `json:"EndAt,omitempty" xml:"EndAt,omitempty"`
	// example:
	//
	// ff1d7783-e087-4d62-92df-3a163eca7c07
	FeeId *string `json:"FeeId,omitempty" xml:"FeeId,omitempty"`
	// example:
	//
	// 2020050811
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 1
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// REMOTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s GetStatisticsRecordsByFeeIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetStatisticsRecordsByFeeIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyData) SetAppId(v int64) *GetStatisticsRecordsByFeeIdResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyData) SetBeginAt(v string) *GetStatisticsRecordsByFeeIdResponseBodyData {
	s.BeginAt = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyData) SetChargeDuration(v int32) *GetStatisticsRecordsByFeeIdResponseBodyData {
	s.ChargeDuration = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyData) SetCreatedAt(v string) *GetStatisticsRecordsByFeeIdResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyData) SetDepartmentId(v int64) *GetStatisticsRecordsByFeeIdResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyData) SetDetectionDuration(v int32) *GetStatisticsRecordsByFeeIdResponseBodyData {
	s.DetectionDuration = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyData) SetDeviceId(v string) *GetStatisticsRecordsByFeeIdResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyData) SetDeviceType(v int32) *GetStatisticsRecordsByFeeIdResponseBodyData {
	s.DeviceType = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyData) SetEndAt(v string) *GetStatisticsRecordsByFeeIdResponseBodyData {
	s.EndAt = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyData) SetFeeId(v string) *GetStatisticsRecordsByFeeIdResponseBodyData {
	s.FeeId = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyData) SetHour(v int32) *GetStatisticsRecordsByFeeIdResponseBodyData {
	s.Hour = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyData) SetTenantId(v int64) *GetStatisticsRecordsByFeeIdResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyData) SetType(v string) *GetStatisticsRecordsByFeeIdResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyData) SetUpdatedAt(v string) *GetStatisticsRecordsByFeeIdResponseBodyData {
	s.UpdatedAt = &v
	return s
}

type GetStatisticsRecordsByFeeIdResponseBodyErrors struct {
	// example:
	//
	// -
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetStatisticsRecordsByFeeIdResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s GetStatisticsRecordsByFeeIdResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyErrors) SetField(v string) *GetStatisticsRecordsByFeeIdResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponseBodyErrors) SetMessage(v string) *GetStatisticsRecordsByFeeIdResponseBodyErrors {
	s.Message = &v
	return s
}

type GetStatisticsRecordsByFeeIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStatisticsRecordsByFeeIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStatisticsRecordsByFeeIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStatisticsRecordsByFeeIdResponse) GoString() string {
	return s.String()
}

func (s *GetStatisticsRecordsByFeeIdResponse) SetHeaders(v map[string]*string) *GetStatisticsRecordsByFeeIdResponse {
	s.Headers = v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponse) SetStatusCode(v int32) *GetStatisticsRecordsByFeeIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStatisticsRecordsByFeeIdResponse) SetBody(v *GetStatisticsRecordsByFeeIdResponseBody) *GetStatisticsRecordsByFeeIdResponse {
	s.Body = v
	return s
}

type GetTaskRequest struct {
	// example:
	//
	// {"version":"1.0.0"}
	ClientBaseParam *string `json:"ClientBaseParam,omitempty" xml:"ClientBaseParam,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTaskRequest) SetClientBaseParam(v string) *GetTaskRequest {
	s.ClientBaseParam = &v
	return s
}

func (s *GetTaskRequest) SetTaskId(v string) *GetTaskRequest {
	s.TaskId = &v
	return s
}

type GetTaskResponseBody struct {
	// example:
	//
	// OK
	Code    *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) SetCode(v string) *GetTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskResponseBody) SetData(v *GetTaskResponseBodyData) *GetTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskResponseBody) SetMessage(v string) *GetTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

type GetTaskResponseBodyData struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// runnable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GetTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyData) SetCreatedAt(v string) *GetTaskResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetTaskResponseBodyData) SetId(v string) *GetTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetTaskResponseBodyData) SetStatus(v string) *GetTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBodyData) SetVideoUrl(v string) *GetTaskResponseBodyData {
	s.VideoUrl = &v
	return s
}

type GetTaskResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResponse) SetHeaders(v map[string]*string) *GetTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTaskResponse) SetStatusCode(v int32) *GetTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskResponse) SetBody(v *GetTaskResponseBody) *GetTaskResponse {
	s.Body = v
	return s
}

type GetTaskGroupRequest struct {
	// example:
	//
	// {"version":"1.0.0"}
	ClientBaseParam *string `json:"ClientBaseParam,omitempty" xml:"ClientBaseParam,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskGroupRequest) GoString() string {
	return s.String()
}

func (s *GetTaskGroupRequest) SetClientBaseParam(v string) *GetTaskGroupRequest {
	s.ClientBaseParam = &v
	return s
}

func (s *GetTaskGroupRequest) SetId(v string) *GetTaskGroupRequest {
	s.Id = &v
	return s
}

type GetTaskGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *GetTaskGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTaskGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskGroupResponseBody) SetCode(v string) *GetTaskGroupResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskGroupResponseBody) SetData(v *GetTaskGroupResponseBodyData) *GetTaskGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskGroupResponseBody) SetMessage(v string) *GetTaskGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskGroupResponseBody) SetRequestId(v string) *GetTaskGroupResponseBody {
	s.RequestId = &v
	return s
}

type GetTaskGroupResponseBodyData struct {
	// example:
	//
	// 1
	CompletedTasks *int32 `json:"CompletedTasks,omitempty" xml:"CompletedTasks,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// default
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// runnable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// -
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalTasks *int32 `json:"TotalTasks,omitempty" xml:"TotalTasks,omitempty"`
}

func (s GetTaskGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTaskGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskGroupResponseBodyData) SetCompletedTasks(v int32) *GetTaskGroupResponseBodyData {
	s.CompletedTasks = &v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetCreatedAt(v string) *GetTaskGroupResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetId(v string) *GetTaskGroupResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetName(v string) *GetTaskGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetRuleId(v string) *GetTaskGroupResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetRuleName(v string) *GetTaskGroupResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetStatus(v string) *GetTaskGroupResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetTaskIds(v []*string) *GetTaskGroupResponseBodyData {
	s.TaskIds = v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetTotalTasks(v int32) *GetTaskGroupResponseBodyData {
	s.TotalTasks = &v
	return s
}

type GetTaskGroupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskGroupResponse) GoString() string {
	return s.String()
}

func (s *GetTaskGroupResponse) SetHeaders(v map[string]*string) *GetTaskGroupResponse {
	s.Headers = v
	return s
}

func (s *GetTaskGroupResponse) SetStatusCode(v int32) *GetTaskGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskGroupResponse) SetBody(v *GetTaskGroupResponseBody) *GetTaskGroupResponse {
	s.Body = v
	return s
}

type GetTtsQuestionByGroupIdRequest struct {
	// example:
	//
	// 47584ba4-9781-496b-8e6f-c8525a213***
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s GetTtsQuestionByGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTtsQuestionByGroupIdRequest) GoString() string {
	return s.String()
}

func (s *GetTtsQuestionByGroupIdRequest) SetGroupId(v string) *GetTtsQuestionByGroupIdRequest {
	s.GroupId = &v
	return s
}

type GetTtsQuestionByGroupIdResponseBody struct {
	// example:
	//
	// OK
	Code *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTtsQuestionByGroupIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D2DFCDC3-1ECF-599D-8EDA-8F598E5A9***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTtsQuestionByGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTtsQuestionByGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTtsQuestionByGroupIdResponseBody) SetCode(v int32) *GetTtsQuestionByGroupIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetTtsQuestionByGroupIdResponseBody) SetData(v *GetTtsQuestionByGroupIdResponseBodyData) *GetTtsQuestionByGroupIdResponseBody {
	s.Data = v
	return s
}

func (s *GetTtsQuestionByGroupIdResponseBody) SetHttpCode(v int32) *GetTtsQuestionByGroupIdResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetTtsQuestionByGroupIdResponseBody) SetMessage(v string) *GetTtsQuestionByGroupIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetTtsQuestionByGroupIdResponseBody) SetRequestId(v string) *GetTtsQuestionByGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTtsQuestionByGroupIdResponseBody) SetSuccess(v bool) *GetTtsQuestionByGroupIdResponseBody {
	s.Success = &v
	return s
}

type GetTtsQuestionByGroupIdResponseBodyData struct {
	// example:
	//
	// 47584ba4-9781-496b-8e6f-c8525a213***
	GroupId   *string                                             `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Questions []*GetTtsQuestionByGroupIdResponseBodyDataQuestions `json:"Questions,omitempty" xml:"Questions,omitempty" type:"Repeated"`
}

func (s GetTtsQuestionByGroupIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTtsQuestionByGroupIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTtsQuestionByGroupIdResponseBodyData) SetGroupId(v string) *GetTtsQuestionByGroupIdResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *GetTtsQuestionByGroupIdResponseBodyData) SetQuestions(v []*GetTtsQuestionByGroupIdResponseBodyDataQuestions) *GetTtsQuestionByGroupIdResponseBodyData {
	s.Questions = v
	return s
}

type GetTtsQuestionByGroupIdResponseBodyDataQuestions struct {
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// example:
	//
	// 674
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// https://pts-data-f***.pcm
	OssUrl   *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// example:
	//
	// 47584ba4-9781-496b-8e6f-c8525a213***
	QuestionGroupId *int64 `json:"QuestionGroupId,omitempty" xml:"QuestionGroupId,omitempty"`
	// example:
	//
	// 47584ba4-9781-496b-8e6f-c8525a213***
	QuestionKey *string `json:"QuestionKey,omitempty" xml:"QuestionKey,omitempty"`
	// example:
	//
	// 2
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetTtsQuestionByGroupIdResponseBodyDataQuestions) String() string {
	return tea.Prettify(s)
}

func (s GetTtsQuestionByGroupIdResponseBodyDataQuestions) GoString() string {
	return s.String()
}

func (s *GetTtsQuestionByGroupIdResponseBodyDataQuestions) SetAnswer(v string) *GetTtsQuestionByGroupIdResponseBodyDataQuestions {
	s.Answer = &v
	return s
}

func (s *GetTtsQuestionByGroupIdResponseBodyDataQuestions) SetDuration(v float64) *GetTtsQuestionByGroupIdResponseBodyDataQuestions {
	s.Duration = &v
	return s
}

func (s *GetTtsQuestionByGroupIdResponseBodyDataQuestions) SetId(v int64) *GetTtsQuestionByGroupIdResponseBodyDataQuestions {
	s.Id = &v
	return s
}

func (s *GetTtsQuestionByGroupIdResponseBodyDataQuestions) SetOssUrl(v string) *GetTtsQuestionByGroupIdResponseBodyDataQuestions {
	s.OssUrl = &v
	return s
}

func (s *GetTtsQuestionByGroupIdResponseBodyDataQuestions) SetQuestion(v string) *GetTtsQuestionByGroupIdResponseBodyDataQuestions {
	s.Question = &v
	return s
}

func (s *GetTtsQuestionByGroupIdResponseBodyDataQuestions) SetQuestionGroupId(v int64) *GetTtsQuestionByGroupIdResponseBodyDataQuestions {
	s.QuestionGroupId = &v
	return s
}

func (s *GetTtsQuestionByGroupIdResponseBodyDataQuestions) SetQuestionKey(v string) *GetTtsQuestionByGroupIdResponseBodyDataQuestions {
	s.QuestionKey = &v
	return s
}

func (s *GetTtsQuestionByGroupIdResponseBodyDataQuestions) SetTenantId(v int64) *GetTtsQuestionByGroupIdResponseBodyDataQuestions {
	s.TenantId = &v
	return s
}

type GetTtsQuestionByGroupIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTtsQuestionByGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTtsQuestionByGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTtsQuestionByGroupIdResponse) GoString() string {
	return s.String()
}

func (s *GetTtsQuestionByGroupIdResponse) SetHeaders(v map[string]*string) *GetTtsQuestionByGroupIdResponse {
	s.Headers = v
	return s
}

func (s *GetTtsQuestionByGroupIdResponse) SetStatusCode(v int32) *GetTtsQuestionByGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTtsQuestionByGroupIdResponse) SetBody(v *GetTtsQuestionByGroupIdResponseBody) *GetTtsQuestionByGroupIdResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	// example:
	//
	// {"version":"1.0.0"}
	ClientBaseParam *string `json:"ClientBaseParam,omitempty" xml:"ClientBaseParam,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetClientBaseParam(v string) *GetUserRequest {
	s.ClientBaseParam = &v
	return s
}

func (s *GetUserRequest) SetId(v string) *GetUserRequest {
	s.Id = &v
	return s
}

type GetUserResponseBody struct {
	// example:
	//
	// OK
	Code    *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetCode(v string) *GetUserResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserResponseBody) SetData(v *GetUserResponseBodyData) *GetUserResponseBody {
	s.Data = v
	return s
}

func (s *GetUserResponseBody) SetMessage(v string) *GetUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

type GetUserResponseBodyData struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt   *string                               `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Departments []*GetUserResponseBodyDataDepartments `json:"Departments,omitempty" xml:"Departments,omitempty" type:"Repeated"`
	// example:
	//
	// xxx@aa.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 187000023323
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// admin
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// ram
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyData) SetCreatedAt(v string) *GetUserResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetUserResponseBodyData) SetDepartments(v []*GetUserResponseBodyDataDepartments) *GetUserResponseBodyData {
	s.Departments = v
	return s
}

func (s *GetUserResponseBodyData) SetEmail(v string) *GetUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyData) SetId(v string) *GetUserResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetUserResponseBodyData) SetName(v string) *GetUserResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetUserResponseBodyData) SetPhoneNumber(v string) *GetUserResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *GetUserResponseBodyData) SetRole(v string) *GetUserResponseBodyData {
	s.Role = &v
	return s
}

func (s *GetUserResponseBodyData) SetSource(v string) *GetUserResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetUserResponseBodyData) SetUpdatedAt(v string) *GetUserResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetUserResponseBodyData) SetUsername(v string) *GetUserResponseBodyData {
	s.Username = &v
	return s
}

type GetUserResponseBodyDataDepartments struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetUserResponseBodyDataDepartments) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyDataDepartments) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyDataDepartments) SetDescription(v string) *GetUserResponseBodyDataDepartments {
	s.Description = &v
	return s
}

func (s *GetUserResponseBodyDataDepartments) SetGmtCreate(v string) *GetUserResponseBodyDataDepartments {
	s.GmtCreate = &v
	return s
}

func (s *GetUserResponseBodyDataDepartments) SetGmtModified(v string) *GetUserResponseBodyDataDepartments {
	s.GmtModified = &v
	return s
}

func (s *GetUserResponseBodyDataDepartments) SetId(v string) *GetUserResponseBodyDataDepartments {
	s.Id = &v
	return s
}

func (s *GetUserResponseBodyDataDepartments) SetName(v string) *GetUserResponseBodyDataDepartments {
	s.Name = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type GetVideoMergeTaskRequest struct {
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetVideoMergeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoMergeTaskRequest) GoString() string {
	return s.String()
}

func (s *GetVideoMergeTaskRequest) SetTaskId(v string) *GetVideoMergeTaskRequest {
	s.TaskId = &v
	return s
}

type GetVideoMergeTaskResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code   *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data   *GetVideoMergeTaskResponseBodyData     `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Errors []*GetVideoMergeTaskResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// 89A61EB9-2FF8-595D-89D3-845C8B615011
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVideoMergeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoMergeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoMergeTaskResponseBody) SetCode(v string) *GetVideoMergeTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoMergeTaskResponseBody) SetData(v *GetVideoMergeTaskResponseBodyData) *GetVideoMergeTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetVideoMergeTaskResponseBody) SetErrors(v []*GetVideoMergeTaskResponseBodyErrors) *GetVideoMergeTaskResponseBody {
	s.Errors = v
	return s
}

func (s *GetVideoMergeTaskResponseBody) SetHttpCode(v int32) *GetVideoMergeTaskResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetVideoMergeTaskResponseBody) SetMessage(v string) *GetVideoMergeTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoMergeTaskResponseBody) SetRequestId(v string) *GetVideoMergeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoMergeTaskResponseBody) SetSuccess(v bool) *GetVideoMergeTaskResponseBody {
	s.Success = &v
	return s
}

type GetVideoMergeTaskResponseBodyData struct {
	// example:
	//
	// 2fr4fa55-2364-4d79-aeb4-c093ec4a4fd4
	ClientTraceId *string `json:"ClientTraceId,omitempty" xml:"ClientTraceId,omitempty"`
	// example:
	//
	// 22
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 0.6
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// http://xxx.mp4
	MergeFileId *string `json:"MergeFileId,omitempty" xml:"MergeFileId,omitempty"`
	// example:
	//
	// 368cfa55-2364-4d79-aeb4-c0956c4a45cd
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 0.6
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetVideoMergeTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetVideoMergeTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVideoMergeTaskResponseBodyData) SetClientTraceId(v string) *GetVideoMergeTaskResponseBodyData {
	s.ClientTraceId = &v
	return s
}

func (s *GetVideoMergeTaskResponseBodyData) SetDuration(v float32) *GetVideoMergeTaskResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetVideoMergeTaskResponseBodyData) SetHeight(v int32) *GetVideoMergeTaskResponseBodyData {
	s.Height = &v
	return s
}

func (s *GetVideoMergeTaskResponseBodyData) SetMergeFileId(v string) *GetVideoMergeTaskResponseBodyData {
	s.MergeFileId = &v
	return s
}

func (s *GetVideoMergeTaskResponseBodyData) SetTaskId(v string) *GetVideoMergeTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetVideoMergeTaskResponseBodyData) SetWidth(v int32) *GetVideoMergeTaskResponseBodyData {
	s.Width = &v
	return s
}

type GetVideoMergeTaskResponseBodyErrors struct {
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetVideoMergeTaskResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s GetVideoMergeTaskResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *GetVideoMergeTaskResponseBodyErrors) SetField(v string) *GetVideoMergeTaskResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *GetVideoMergeTaskResponseBodyErrors) SetMessage(v string) *GetVideoMergeTaskResponseBodyErrors {
	s.Message = &v
	return s
}

type GetVideoMergeTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoMergeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoMergeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoMergeTaskResponse) GoString() string {
	return s.String()
}

func (s *GetVideoMergeTaskResponse) SetHeaders(v map[string]*string) *GetVideoMergeTaskResponse {
	s.Headers = v
	return s
}

func (s *GetVideoMergeTaskResponse) SetStatusCode(v int32) *GetVideoMergeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoMergeTaskResponse) SetBody(v *GetVideoMergeTaskResponseBody) *GetVideoMergeTaskResponse {
	s.Body = v
	return s
}

type GetWatermarkRequest struct {
	// example:
	//
	// {"version":"1.0.0"}
	ClientBaseParam *string `json:"ClientBaseParam,omitempty" xml:"ClientBaseParam,omitempty"`
	// example:
	//
	// 1.0.003
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// 728fd812a876ec04818858982baebe6f
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s GetWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWatermarkRequest) GoString() string {
	return s.String()
}

func (s *GetWatermarkRequest) SetClientBaseParam(v string) *GetWatermarkRequest {
	s.ClientBaseParam = &v
	return s
}

func (s *GetWatermarkRequest) SetClientVersion(v string) *GetWatermarkRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetWatermarkRequest) SetWatermarkId(v string) *GetWatermarkRequest {
	s.WatermarkId = &v
	return s
}

type GetWatermarkResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code   *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data   *GetWatermarkResponseBodyData     `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Errors []*GetWatermarkResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// 84118BF0-56F7-54D2-8C1A-35BBBBD5016B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *GetWatermarkResponseBody) SetCode(v string) *GetWatermarkResponseBody {
	s.Code = &v
	return s
}

func (s *GetWatermarkResponseBody) SetData(v *GetWatermarkResponseBodyData) *GetWatermarkResponseBody {
	s.Data = v
	return s
}

func (s *GetWatermarkResponseBody) SetErrors(v []*GetWatermarkResponseBodyErrors) *GetWatermarkResponseBody {
	s.Errors = v
	return s
}

func (s *GetWatermarkResponseBody) SetHttpCode(v int32) *GetWatermarkResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetWatermarkResponseBody) SetMessage(v string) *GetWatermarkResponseBody {
	s.Message = &v
	return s
}

func (s *GetWatermarkResponseBody) SetRequestId(v string) *GetWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWatermarkResponseBody) SetSuccess(v bool) *GetWatermarkResponseBody {
	s.Success = &v
	return s
}

type GetWatermarkResponseBodyData struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetWatermarkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWatermarkResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWatermarkResponseBodyData) SetCreatedAt(v string) *GetWatermarkResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetWatermarkResponseBodyData) SetId(v string) *GetWatermarkResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetWatermarkResponseBodyData) SetName(v string) *GetWatermarkResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetWatermarkResponseBodyData) SetValue(v string) *GetWatermarkResponseBodyData {
	s.Value = &v
	return s
}

type GetWatermarkResponseBodyErrors struct {
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetWatermarkResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s GetWatermarkResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *GetWatermarkResponseBodyErrors) SetField(v string) *GetWatermarkResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *GetWatermarkResponseBodyErrors) SetMessage(v string) *GetWatermarkResponseBodyErrors {
	s.Message = &v
	return s
}

type GetWatermarkResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWatermarkResponse) GoString() string {
	return s.String()
}

func (s *GetWatermarkResponse) SetHeaders(v map[string]*string) *GetWatermarkResponse {
	s.Headers = v
	return s
}

func (s *GetWatermarkResponse) SetStatusCode(v int32) *GetWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWatermarkResponse) SetBody(v *GetWatermarkResponseBody) *GetWatermarkResponse {
	s.Body = v
	return s
}

type JoinRoomRequest struct {
	// example:
	//
	// 4a29b426-742f-4078-8386-79b440b***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 66194***681868
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// example:
	//
	// 006f4***b269
	RoomToken *string `json:"RoomToken,omitempty" xml:"RoomToken,omitempty"`
	// example:
	//
	// 12312***2412
	StreamId *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
}

func (s JoinRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinRoomRequest) GoString() string {
	return s.String()
}

func (s *JoinRoomRequest) SetAppId(v string) *JoinRoomRequest {
	s.AppId = &v
	return s
}

func (s *JoinRoomRequest) SetRoomId(v string) *JoinRoomRequest {
	s.RoomId = &v
	return s
}

func (s *JoinRoomRequest) SetRoomToken(v string) *JoinRoomRequest {
	s.RoomToken = &v
	return s
}

func (s *JoinRoomRequest) SetStreamId(v string) *JoinRoomRequest {
	s.StreamId = &v
	return s
}

type JoinRoomResponseBody struct {
	// example:
	//
	// OK
	Code *int32                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *JoinRoomResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D15744FC-97D2-5D61-A18C-8CC4FA0A3***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s JoinRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinRoomResponseBody) GoString() string {
	return s.String()
}

func (s *JoinRoomResponseBody) SetCode(v int32) *JoinRoomResponseBody {
	s.Code = &v
	return s
}

func (s *JoinRoomResponseBody) SetData(v *JoinRoomResponseBodyData) *JoinRoomResponseBody {
	s.Data = v
	return s
}

func (s *JoinRoomResponseBody) SetHttpCode(v int32) *JoinRoomResponseBody {
	s.HttpCode = &v
	return s
}

func (s *JoinRoomResponseBody) SetMessage(v string) *JoinRoomResponseBody {
	s.Message = &v
	return s
}

func (s *JoinRoomResponseBody) SetRequestId(v string) *JoinRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinRoomResponseBody) SetSuccess(v bool) *JoinRoomResponseBody {
	s.Success = &v
	return s
}

type JoinRoomResponseBodyData struct {
	// example:
	//
	// 641981583353***
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// example:
	//
	// 12312***412
	StreamId *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
}

func (s JoinRoomResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s JoinRoomResponseBodyData) GoString() string {
	return s.String()
}

func (s *JoinRoomResponseBodyData) SetRoomId(v string) *JoinRoomResponseBodyData {
	s.RoomId = &v
	return s
}

func (s *JoinRoomResponseBodyData) SetStreamId(v string) *JoinRoomResponseBodyData {
	s.StreamId = &v
	return s
}

type JoinRoomResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JoinRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JoinRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinRoomResponse) GoString() string {
	return s.String()
}

func (s *JoinRoomResponse) SetHeaders(v map[string]*string) *JoinRoomResponse {
	s.Headers = v
	return s
}

func (s *JoinRoomResponse) SetStatusCode(v int32) *JoinRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinRoomResponse) SetBody(v *JoinRoomResponseBody) *JoinRoomResponse {
	s.Body = v
	return s
}

type LeaveRoomRequest struct {
	// example:
	//
	// 661940620884***
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s LeaveRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s LeaveRoomRequest) GoString() string {
	return s.String()
}

func (s *LeaveRoomRequest) SetRoomId(v string) *LeaveRoomRequest {
	s.RoomId = &v
	return s
}

type LeaveRoomResponseBody struct {
	// example:
	//
	// OK
	Code *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 8F5F961D-07AA-5293-BFFE-32CDC251***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LeaveRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LeaveRoomResponseBody) GoString() string {
	return s.String()
}

func (s *LeaveRoomResponseBody) SetCode(v int32) *LeaveRoomResponseBody {
	s.Code = &v
	return s
}

func (s *LeaveRoomResponseBody) SetData(v string) *LeaveRoomResponseBody {
	s.Data = &v
	return s
}

func (s *LeaveRoomResponseBody) SetHttpCode(v int32) *LeaveRoomResponseBody {
	s.HttpCode = &v
	return s
}

func (s *LeaveRoomResponseBody) SetMessage(v string) *LeaveRoomResponseBody {
	s.Message = &v
	return s
}

func (s *LeaveRoomResponseBody) SetRequestId(v string) *LeaveRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *LeaveRoomResponseBody) SetSuccess(v bool) *LeaveRoomResponseBody {
	s.Success = &v
	return s
}

type LeaveRoomResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LeaveRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LeaveRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s LeaveRoomResponse) GoString() string {
	return s.String()
}

func (s *LeaveRoomResponse) SetHeaders(v map[string]*string) *LeaveRoomResponse {
	s.Headers = v
	return s
}

func (s *LeaveRoomResponse) SetStatusCode(v int32) *LeaveRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *LeaveRoomResponse) SetBody(v *LeaveRoomResponseBody) *LeaveRoomResponse {
	s.Body = v
	return s
}

type ListAppsRequest struct {
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) SetPageIndex(v int32) *ListAppsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListAppsRequest) SetPageSize(v int32) *ListAppsRequest {
	s.PageSize = &v
	return s
}

type ListAppsResponseBody struct {
	// example:
	//
	// OK
	Code    *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *ListAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBody) SetCode(v string) *ListAppsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppsResponseBody) SetData(v *ListAppsResponseBodyData) *ListAppsResponseBody {
	s.Data = v
	return s
}

func (s *ListAppsResponseBody) SetMessage(v string) *ListAppsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppsResponseBody) SetRequestId(v string) *ListAppsResponseBody {
	s.RequestId = &v
	return s
}

type ListAppsResponseBodyData struct {
	Items []*ListAppsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// example:
	//
	// 11
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListAppsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyData) SetItems(v []*ListAppsResponseBodyDataItems) *ListAppsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListAppsResponseBodyData) SetTotalElements(v int64) *ListAppsResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListAppsResponseBodyData) SetTotalPages(v int32) *ListAppsResponseBodyData {
	s.TotalPages = &v
	return s
}

type ListAppsResponseBodyDataItems struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	DepartmentId   *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// com.aaa.test
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
}

func (s ListAppsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyDataItems) SetCreatedAt(v string) *ListAppsResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetDepartmentId(v string) *ListAppsResponseBodyDataItems {
	s.DepartmentId = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetDepartmentName(v string) *ListAppsResponseBodyDataItems {
	s.DepartmentName = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetDisabled(v bool) *ListAppsResponseBodyDataItems {
	s.Disabled = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetId(v string) *ListAppsResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetName(v string) *ListAppsResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetPackageName(v string) *ListAppsResponseBodyDataItems {
	s.PackageName = &v
	return s
}

type ListAppsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponse) GoString() string {
	return s.String()
}

func (s *ListAppsResponse) SetHeaders(v map[string]*string) *ListAppsResponse {
	s.Headers = v
	return s
}

func (s *ListAppsResponse) SetStatusCode(v int32) *ListAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppsResponse) SetBody(v *ListAppsResponseBody) *ListAppsResponse {
	s.Body = v
	return s
}

type ListDepartmentsRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListDepartmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDepartmentsRequest) GoString() string {
	return s.String()
}

func (s *ListDepartmentsRequest) SetName(v string) *ListDepartmentsRequest {
	s.Name = &v
	return s
}

func (s *ListDepartmentsRequest) SetPageIndex(v int32) *ListDepartmentsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDepartmentsRequest) SetPageSize(v int32) *ListDepartmentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDepartmentsRequest) SetUserId(v string) *ListDepartmentsRequest {
	s.UserId = &v
	return s
}

type ListDepartmentsResponseBody struct {
	// example:
	//
	// OK
	Code    *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *ListDepartmentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDepartmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDepartmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDepartmentsResponseBody) SetCode(v string) *ListDepartmentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDepartmentsResponseBody) SetData(v *ListDepartmentsResponseBodyData) *ListDepartmentsResponseBody {
	s.Data = v
	return s
}

func (s *ListDepartmentsResponseBody) SetMessage(v string) *ListDepartmentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDepartmentsResponseBody) SetRequestId(v string) *ListDepartmentsResponseBody {
	s.RequestId = &v
	return s
}

type ListDepartmentsResponseBodyData struct {
	Items []*ListDepartmentsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// example:
	//
	// 11
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListDepartmentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDepartmentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDepartmentsResponseBodyData) SetItems(v []*ListDepartmentsResponseBodyDataItems) *ListDepartmentsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListDepartmentsResponseBodyData) SetTotalElements(v int64) *ListDepartmentsResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListDepartmentsResponseBodyData) SetTotalPages(v int32) *ListDepartmentsResponseBodyData {
	s.TotalPages = &v
	return s
}

type ListDepartmentsResponseBodyDataItems struct {
	Administrators []*ListDepartmentsResponseBodyDataItemsAdministrators `json:"Administrators,omitempty" xml:"Administrators,omitempty" type:"Repeated"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListDepartmentsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListDepartmentsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListDepartmentsResponseBodyDataItems) SetAdministrators(v []*ListDepartmentsResponseBodyDataItemsAdministrators) *ListDepartmentsResponseBodyDataItems {
	s.Administrators = v
	return s
}

func (s *ListDepartmentsResponseBodyDataItems) SetCreatedAt(v string) *ListDepartmentsResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListDepartmentsResponseBodyDataItems) SetDescription(v string) *ListDepartmentsResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListDepartmentsResponseBodyDataItems) SetId(v string) *ListDepartmentsResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListDepartmentsResponseBodyDataItems) SetName(v string) *ListDepartmentsResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListDepartmentsResponseBodyDataItems) SetUpdatedAt(v string) *ListDepartmentsResponseBodyDataItems {
	s.UpdatedAt = &v
	return s
}

type ListDepartmentsResponseBodyDataItemsAdministrators struct {
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// user
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDepartmentsResponseBodyDataItemsAdministrators) String() string {
	return tea.Prettify(s)
}

func (s ListDepartmentsResponseBodyDataItemsAdministrators) GoString() string {
	return s.String()
}

func (s *ListDepartmentsResponseBodyDataItemsAdministrators) SetId(v string) *ListDepartmentsResponseBodyDataItemsAdministrators {
	s.Id = &v
	return s
}

func (s *ListDepartmentsResponseBodyDataItemsAdministrators) SetName(v string) *ListDepartmentsResponseBodyDataItemsAdministrators {
	s.Name = &v
	return s
}

type ListDepartmentsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDepartmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDepartmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDepartmentsResponse) GoString() string {
	return s.String()
}

func (s *ListDepartmentsResponse) SetHeaders(v map[string]*string) *ListDepartmentsResponse {
	s.Headers = v
	return s
}

func (s *ListDepartmentsResponse) SetStatusCode(v int32) *ListDepartmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDepartmentsResponse) SetBody(v *ListDepartmentsResponseBody) *ListDepartmentsResponse {
	s.Body = v
	return s
}

type ListDetectProcessesRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	PublishStatus *bool `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	// example:
	//
	// desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// CONTENT
	SortKey *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	// example:
	//
	// LOCAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDetectProcessesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDetectProcessesRequest) GoString() string {
	return s.String()
}

func (s *ListDetectProcessesRequest) SetName(v string) *ListDetectProcessesRequest {
	s.Name = &v
	return s
}

func (s *ListDetectProcessesRequest) SetPageIndex(v int32) *ListDetectProcessesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDetectProcessesRequest) SetPageSize(v int32) *ListDetectProcessesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDetectProcessesRequest) SetPublishStatus(v bool) *ListDetectProcessesRequest {
	s.PublishStatus = &v
	return s
}

func (s *ListDetectProcessesRequest) SetSort(v string) *ListDetectProcessesRequest {
	s.Sort = &v
	return s
}

func (s *ListDetectProcessesRequest) SetSortKey(v string) *ListDetectProcessesRequest {
	s.SortKey = &v
	return s
}

func (s *ListDetectProcessesRequest) SetType(v string) *ListDetectProcessesRequest {
	s.Type = &v
	return s
}

type ListDetectProcessesResponseBody struct {
	// example:
	//
	// OK
	Code    *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *ListDetectProcessesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// bf1c45cd-3eee-4e60-b505-2e330bf755d3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDetectProcessesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDetectProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDetectProcessesResponseBody) SetCode(v string) *ListDetectProcessesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDetectProcessesResponseBody) SetData(v *ListDetectProcessesResponseBodyData) *ListDetectProcessesResponseBody {
	s.Data = v
	return s
}

func (s *ListDetectProcessesResponseBody) SetMessage(v string) *ListDetectProcessesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDetectProcessesResponseBody) SetRequestId(v string) *ListDetectProcessesResponseBody {
	s.RequestId = &v
	return s
}

type ListDetectProcessesResponseBodyData struct {
	Items []*ListDetectProcessesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// example:
	//
	// 10
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListDetectProcessesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDetectProcessesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDetectProcessesResponseBodyData) SetItems(v []*ListDetectProcessesResponseBodyDataItems) *ListDetectProcessesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListDetectProcessesResponseBodyData) SetTotalElements(v int64) *ListDetectProcessesResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListDetectProcessesResponseBodyData) SetTotalPages(v int32) *ListDetectProcessesResponseBodyData {
	s.TotalPages = &v
	return s
}

type ListDetectProcessesResponseBodyDataItems struct {
	// example:
	//
	// {}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2020-12-04T14:48:59.000+08:00
	ContentAt *string `json:"ContentAt,omitempty" xml:"ContentAt,omitempty"`
	// example:
	//
	// 2020-12-04T14:47:59.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// {}
	Draft *string `json:"Draft,omitempty" xml:"Draft,omitempty"`
	// example:
	//
	// 2020-12-04T14:48:59.000+08:00
	DraftAt *string `json:"DraftAt,omitempty" xml:"DraftAt,omitempty"`
	// example:
	//
	// http://abc.com/123.json
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// ID
	//
	// example:
	//
	// 0f1c45cd-3eee-4e60-b505-2e330b8755d3
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 987d563d38f5aef27feca8702c689bb1
	Md5  *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	NewVersion *bool `json:"NewVersion,omitempty" xml:"NewVersion,omitempty"`
	// example:
	//
	// True
	Published *bool `json:"Published,omitempty" xml:"Published,omitempty"`
	// example:
	//
	// REMOTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2020-12-04T14:48:59.000+08:00
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListDetectProcessesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListDetectProcessesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListDetectProcessesResponseBodyDataItems) SetContent(v string) *ListDetectProcessesResponseBodyDataItems {
	s.Content = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetContentAt(v string) *ListDetectProcessesResponseBodyDataItems {
	s.ContentAt = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetCreatedAt(v string) *ListDetectProcessesResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetDraft(v string) *ListDetectProcessesResponseBodyDataItems {
	s.Draft = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetDraftAt(v string) *ListDetectProcessesResponseBodyDataItems {
	s.DraftAt = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetFileUrl(v string) *ListDetectProcessesResponseBodyDataItems {
	s.FileUrl = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetId(v string) *ListDetectProcessesResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetMd5(v string) *ListDetectProcessesResponseBodyDataItems {
	s.Md5 = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetName(v string) *ListDetectProcessesResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetNewVersion(v bool) *ListDetectProcessesResponseBodyDataItems {
	s.NewVersion = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetPublished(v bool) *ListDetectProcessesResponseBodyDataItems {
	s.Published = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetType(v string) *ListDetectProcessesResponseBodyDataItems {
	s.Type = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetUpdatedAt(v string) *ListDetectProcessesResponseBodyDataItems {
	s.UpdatedAt = &v
	return s
}

type ListDetectProcessesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDetectProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDetectProcessesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDetectProcessesResponse) GoString() string {
	return s.String()
}

func (s *ListDetectProcessesResponse) SetHeaders(v map[string]*string) *ListDetectProcessesResponse {
	s.Headers = v
	return s
}

func (s *ListDetectProcessesResponse) SetStatusCode(v int32) *ListDetectProcessesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDetectProcessesResponse) SetBody(v *ListDetectProcessesResponseBody) *ListDetectProcessesResponse {
	s.Body = v
	return s
}

type ListDetectionsRequest struct {
	// example:
	//
	// 2020-10-10
	CreateDateFrom *string `json:"CreateDateFrom,omitempty" xml:"CreateDateFrom,omitempty"`
	// example:
	//
	// 2020-10-11
	CreateDateTo *string `json:"CreateDateTo,omitempty" xml:"CreateDateTo,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// local
	RecordingType *string `json:"RecordingType,omitempty" xml:"RecordingType,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListDetectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionsRequest) GoString() string {
	return s.String()
}

func (s *ListDetectionsRequest) SetCreateDateFrom(v string) *ListDetectionsRequest {
	s.CreateDateFrom = &v
	return s
}

func (s *ListDetectionsRequest) SetCreateDateTo(v string) *ListDetectionsRequest {
	s.CreateDateTo = &v
	return s
}

func (s *ListDetectionsRequest) SetDepartmentId(v string) *ListDetectionsRequest {
	s.DepartmentId = &v
	return s
}

func (s *ListDetectionsRequest) SetPageIndex(v int32) *ListDetectionsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDetectionsRequest) SetPageSize(v int32) *ListDetectionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDetectionsRequest) SetRecordingType(v string) *ListDetectionsRequest {
	s.RecordingType = &v
	return s
}

func (s *ListDetectionsRequest) SetRuleId(v string) *ListDetectionsRequest {
	s.RuleId = &v
	return s
}

type ListDetectionsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ListDetectionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// -
	Errors []*ListDetectionsResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDetectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDetectionsResponseBody) SetCode(v string) *ListDetectionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDetectionsResponseBody) SetData(v *ListDetectionsResponseBodyData) *ListDetectionsResponseBody {
	s.Data = v
	return s
}

func (s *ListDetectionsResponseBody) SetErrors(v []*ListDetectionsResponseBodyErrors) *ListDetectionsResponseBody {
	s.Errors = v
	return s
}

func (s *ListDetectionsResponseBody) SetMessage(v string) *ListDetectionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDetectionsResponseBody) SetRequestId(v string) *ListDetectionsResponseBody {
	s.RequestId = &v
	return s
}

type ListDetectionsResponseBodyData struct {
	Items []*ListDetectionsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// example:
	//
	// 11
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListDetectionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDetectionsResponseBodyData) SetItems(v []*ListDetectionsResponseBodyDataItems) *ListDetectionsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListDetectionsResponseBodyData) SetTotalElements(v int64) *ListDetectionsResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListDetectionsResponseBodyData) SetTotalPages(v int32) *ListDetectionsResponseBodyData {
	s.TotalPages = &v
	return s
}

type ListDetectionsResponseBodyDataItems struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	DepartmentId   *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// local
	RecordingType *string `json:"RecordingType,omitempty" xml:"RecordingType,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// default
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// runnable
	Status *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Tasks  []*ListDetectionsResponseBodyDataItemsTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListDetectionsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListDetectionsResponseBodyDataItems) SetCreatedAt(v string) *ListDetectionsResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetDepartmentId(v string) *ListDetectionsResponseBodyDataItems {
	s.DepartmentId = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetDepartmentName(v string) *ListDetectionsResponseBodyDataItems {
	s.DepartmentName = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetId(v string) *ListDetectionsResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetRecordingType(v string) *ListDetectionsResponseBodyDataItems {
	s.RecordingType = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetRuleId(v string) *ListDetectionsResponseBodyDataItems {
	s.RuleId = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetRuleName(v string) *ListDetectionsResponseBodyDataItems {
	s.RuleName = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetStatus(v string) *ListDetectionsResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetTasks(v []*ListDetectionsResponseBodyDataItemsTasks) *ListDetectionsResponseBodyDataItems {
	s.Tasks = v
	return s
}

type ListDetectionsResponseBodyDataItemsTasks struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// runnable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.meta
	VideoMetaUrl *string `json:"VideoMetaUrl,omitempty" xml:"VideoMetaUrl,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s ListDetectionsResponseBodyDataItemsTasks) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionsResponseBodyDataItemsTasks) GoString() string {
	return s.String()
}

func (s *ListDetectionsResponseBodyDataItemsTasks) SetCreatedAt(v string) *ListDetectionsResponseBodyDataItemsTasks {
	s.CreatedAt = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItemsTasks) SetId(v string) *ListDetectionsResponseBodyDataItemsTasks {
	s.Id = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItemsTasks) SetStatus(v string) *ListDetectionsResponseBodyDataItemsTasks {
	s.Status = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItemsTasks) SetVideoMetaUrl(v string) *ListDetectionsResponseBodyDataItemsTasks {
	s.VideoMetaUrl = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItemsTasks) SetVideoUrl(v string) *ListDetectionsResponseBodyDataItemsTasks {
	s.VideoUrl = &v
	return s
}

type ListDetectionsResponseBodyErrors struct {
	// example:
	//
	// status
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListDetectionsResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionsResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *ListDetectionsResponseBodyErrors) SetField(v string) *ListDetectionsResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *ListDetectionsResponseBodyErrors) SetMessage(v string) *ListDetectionsResponseBodyErrors {
	s.Message = &v
	return s
}

type ListDetectionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDetectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDetectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionsResponse) GoString() string {
	return s.String()
}

func (s *ListDetectionsResponse) SetHeaders(v map[string]*string) *ListDetectionsResponse {
	s.Headers = v
	return s
}

func (s *ListDetectionsResponse) SetStatusCode(v int32) *ListDetectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDetectionsResponse) SetBody(v *ListDetectionsResponseBody) *ListDetectionsResponse {
	s.Body = v
	return s
}

type ListFilesRequest struct {
	// example:
	//
	// 500
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// 20220110/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFilesRequest) GoString() string {
	return s.String()
}

func (s *ListFilesRequest) SetLimit(v int32) *ListFilesRequest {
	s.Limit = &v
	return s
}

func (s *ListFilesRequest) SetPrefix(v string) *ListFilesRequest {
	s.Prefix = &v
	return s
}

type ListFilesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// [\"d0388a41-30c6-4c74-af83-a1f923ce1725/process/\", \"d0388a41-30c6-4c74-af83-a1f923ce1725/video_12_14_09_05_41.mp4\", \"d0388a41-30c6-4c74-af83-a1f923ce1725/video_12_14_09_05_41.mp4.meta\"]
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFilesResponseBody) SetCode(v string) *ListFilesResponseBody {
	s.Code = &v
	return s
}

func (s *ListFilesResponseBody) SetData(v string) *ListFilesResponseBody {
	s.Data = &v
	return s
}

func (s *ListFilesResponseBody) SetMessage(v string) *ListFilesResponseBody {
	s.Message = &v
	return s
}

func (s *ListFilesResponseBody) SetRequestId(v string) *ListFilesResponseBody {
	s.RequestId = &v
	return s
}

type ListFilesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFilesResponse) GoString() string {
	return s.String()
}

func (s *ListFilesResponse) SetHeaders(v map[string]*string) *ListFilesResponse {
	s.Headers = v
	return s
}

func (s *ListFilesResponse) SetStatusCode(v int32) *ListFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFilesResponse) SetBody(v *ListFilesResponseBody) *ListFilesResponse {
	s.Body = v
	return s
}

type ListRecordResultsRequest struct {
	// example:
	//
	// 2020-10-10
	CreateDateFrom *string `json:"CreateDateFrom,omitempty" xml:"CreateDateFrom,omitempty"`
	// example:
	//
	// 2020-10-11
	CreateDateTo *string `json:"CreateDateTo,omitempty" xml:"CreateDateTo,omitempty"`
	// example:
	//
	// 90c2459a-4c73-4ab7-8a6b-e550d62fbd8c
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// example:
	//
	// ads32efef43
	OuterBusinessId *string `json:"OuterBusinessId,omitempty" xml:"OuterBusinessId,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// LOCAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRecordResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecordResultsRequest) GoString() string {
	return s.String()
}

func (s *ListRecordResultsRequest) SetCreateDateFrom(v string) *ListRecordResultsRequest {
	s.CreateDateFrom = &v
	return s
}

func (s *ListRecordResultsRequest) SetCreateDateTo(v string) *ListRecordResultsRequest {
	s.CreateDateTo = &v
	return s
}

func (s *ListRecordResultsRequest) SetDepartmentId(v string) *ListRecordResultsRequest {
	s.DepartmentId = &v
	return s
}

func (s *ListRecordResultsRequest) SetOuterBusinessId(v string) *ListRecordResultsRequest {
	s.OuterBusinessId = &v
	return s
}

func (s *ListRecordResultsRequest) SetPageIndex(v int32) *ListRecordResultsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListRecordResultsRequest) SetPageSize(v int32) *ListRecordResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecordResultsRequest) SetRecordId(v string) *ListRecordResultsRequest {
	s.RecordId = &v
	return s
}

func (s *ListRecordResultsRequest) SetType(v string) *ListRecordResultsRequest {
	s.Type = &v
	return s
}

type ListRecordResultsResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code   *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data   *ListRecordResultsResponseBodyData     `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Errors []*ListRecordResultsResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRecordResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecordResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecordResultsResponseBody) SetCode(v string) *ListRecordResultsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRecordResultsResponseBody) SetData(v *ListRecordResultsResponseBodyData) *ListRecordResultsResponseBody {
	s.Data = v
	return s
}

func (s *ListRecordResultsResponseBody) SetErrors(v []*ListRecordResultsResponseBodyErrors) *ListRecordResultsResponseBody {
	s.Errors = v
	return s
}

func (s *ListRecordResultsResponseBody) SetHttpCode(v int32) *ListRecordResultsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListRecordResultsResponseBody) SetMessage(v string) *ListRecordResultsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRecordResultsResponseBody) SetRequestId(v string) *ListRecordResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecordResultsResponseBody) SetSuccess(v bool) *ListRecordResultsResponseBody {
	s.Success = &v
	return s
}

type ListRecordResultsResponseBodyData struct {
	Items []*ListRecordResultsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// example:
	//
	// 11
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListRecordResultsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRecordResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRecordResultsResponseBodyData) SetItems(v []*ListRecordResultsResponseBodyDataItems) *ListRecordResultsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListRecordResultsResponseBodyData) SetTotalElements(v int64) *ListRecordResultsResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListRecordResultsResponseBodyData) SetTotalPages(v int32) *ListRecordResultsResponseBodyData {
	s.TotalPages = &v
	return s
}

type ListRecordResultsResponseBodyDataItems struct {
	// example:
	//
	// myapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt      *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// example:
	//
	// process name
	DetectProcessName *string `json:"DetectProcessName,omitempty" xml:"DetectProcessName,omitempty"`
	// example:
	//
	// 22
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.meta
	MetaUrl *string `json:"MetaUrl,omitempty" xml:"MetaUrl,omitempty"`
	// example:
	//
	// fasfasda
	OuterBusinessId *string `json:"OuterBusinessId,omitempty" xml:"OuterBusinessId,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	RecordAt *string `json:"RecordAt,omitempty" xml:"RecordAt,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.json
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
	// example:
	//
	// 642662467657798
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// example:
	//
	// record_65703154805715668342
	RtcRecordId *string `json:"RtcRecordId,omitempty" xml:"RtcRecordId,omitempty"`
	// example:
	//
	// runnable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s ListRecordResultsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListRecordResultsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListRecordResultsResponseBodyDataItems) SetAppName(v string) *ListRecordResultsResponseBodyDataItems {
	s.AppName = &v
	return s
}

func (s *ListRecordResultsResponseBodyDataItems) SetCreatedAt(v string) *ListRecordResultsResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListRecordResultsResponseBodyDataItems) SetDepartmentName(v string) *ListRecordResultsResponseBodyDataItems {
	s.DepartmentName = &v
	return s
}

func (s *ListRecordResultsResponseBodyDataItems) SetDetectProcessName(v string) *ListRecordResultsResponseBodyDataItems {
	s.DetectProcessName = &v
	return s
}

func (s *ListRecordResultsResponseBodyDataItems) SetDuration(v int64) *ListRecordResultsResponseBodyDataItems {
	s.Duration = &v
	return s
}

func (s *ListRecordResultsResponseBodyDataItems) SetId(v string) *ListRecordResultsResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListRecordResultsResponseBodyDataItems) SetMetaUrl(v string) *ListRecordResultsResponseBodyDataItems {
	s.MetaUrl = &v
	return s
}

func (s *ListRecordResultsResponseBodyDataItems) SetOuterBusinessId(v string) *ListRecordResultsResponseBodyDataItems {
	s.OuterBusinessId = &v
	return s
}

func (s *ListRecordResultsResponseBodyDataItems) SetRecordAt(v string) *ListRecordResultsResponseBodyDataItems {
	s.RecordAt = &v
	return s
}

func (s *ListRecordResultsResponseBodyDataItems) SetResultUrl(v string) *ListRecordResultsResponseBodyDataItems {
	s.ResultUrl = &v
	return s
}

func (s *ListRecordResultsResponseBodyDataItems) SetRoomId(v string) *ListRecordResultsResponseBodyDataItems {
	s.RoomId = &v
	return s
}

func (s *ListRecordResultsResponseBodyDataItems) SetRtcRecordId(v string) *ListRecordResultsResponseBodyDataItems {
	s.RtcRecordId = &v
	return s
}

func (s *ListRecordResultsResponseBodyDataItems) SetStatus(v string) *ListRecordResultsResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListRecordResultsResponseBodyDataItems) SetVideoUrl(v string) *ListRecordResultsResponseBodyDataItems {
	s.VideoUrl = &v
	return s
}

type ListRecordResultsResponseBodyErrors struct {
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	Field   *string `json:"Field,omitempty" xml:"Field,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListRecordResultsResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s ListRecordResultsResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *ListRecordResultsResponseBodyErrors) SetField(v string) *ListRecordResultsResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *ListRecordResultsResponseBodyErrors) SetMessage(v string) *ListRecordResultsResponseBodyErrors {
	s.Message = &v
	return s
}

type ListRecordResultsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecordResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecordResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecordResultsResponse) GoString() string {
	return s.String()
}

func (s *ListRecordResultsResponse) SetHeaders(v map[string]*string) *ListRecordResultsResponse {
	s.Headers = v
	return s
}

func (s *ListRecordResultsResponse) SetStatusCode(v int32) *ListRecordResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecordResultsResponse) SetBody(v *ListRecordResultsResponseBody) *ListRecordResultsResponse {
	s.Body = v
	return s
}

type ListRulesRequest struct {
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRulesRequest) SetPageIndex(v int32) *ListRulesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListRulesRequest) SetPageSize(v int32) *ListRulesRequest {
	s.PageSize = &v
	return s
}

type ListRulesResponseBody struct {
	// example:
	//
	// OK
	Code    *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *ListRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBody) SetCode(v string) *ListRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRulesResponseBody) SetData(v *ListRulesResponseBodyData) *ListRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListRulesResponseBody) SetMessage(v string) *ListRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRulesResponseBody) SetRequestId(v string) *ListRulesResponseBody {
	s.RequestId = &v
	return s
}

type ListRulesResponseBodyData struct {
	Items []*ListRulesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyData) SetItems(v []*ListRulesResponseBodyDataItems) *ListRulesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListRulesResponseBodyData) SetTotalElements(v int64) *ListRulesResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListRulesResponseBodyData) SetTotalPages(v int32) *ListRulesResponseBodyData {
	s.TotalPages = &v
	return s
}

type ListRulesResponseBodyDataItems struct {
	// example:
	//
	// {}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListRulesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyDataItems) SetContent(v string) *ListRulesResponseBodyDataItems {
	s.Content = &v
	return s
}

func (s *ListRulesResponseBodyDataItems) SetCreatedAt(v string) *ListRulesResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListRulesResponseBodyDataItems) SetId(v string) *ListRulesResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListRulesResponseBodyDataItems) SetName(v string) *ListRulesResponseBodyDataItems {
	s.Name = &v
	return s
}

type ListRulesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRulesResponse) SetHeaders(v map[string]*string) *ListRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRulesResponse) SetStatusCode(v int32) *ListRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRulesResponse) SetBody(v *ListRulesResponseBody) *ListRulesResponse {
	s.Body = v
	return s
}

type ListTaskGroupsRequest struct {
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// runnable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListTaskGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaskGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListTaskGroupsRequest) SetPageIndex(v int32) *ListTaskGroupsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListTaskGroupsRequest) SetPageSize(v int32) *ListTaskGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskGroupsRequest) SetStatus(v string) *ListTaskGroupsRequest {
	s.Status = &v
	return s
}

type ListTaskGroupsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ListTaskGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTaskGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTaskGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskGroupsResponseBody) SetCode(v string) *ListTaskGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTaskGroupsResponseBody) SetData(v *ListTaskGroupsResponseBodyData) *ListTaskGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListTaskGroupsResponseBody) SetMessage(v string) *ListTaskGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTaskGroupsResponseBody) SetRequestId(v string) *ListTaskGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ListTaskGroupsResponseBodyData struct {
	Items []*ListTaskGroupsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListTaskGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTaskGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTaskGroupsResponseBodyData) SetItems(v []*ListTaskGroupsResponseBodyDataItems) *ListTaskGroupsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListTaskGroupsResponseBodyData) SetTotalElements(v int64) *ListTaskGroupsResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListTaskGroupsResponseBodyData) SetTotalPages(v int32) *ListTaskGroupsResponseBodyData {
	s.TotalPages = &v
	return s
}

type ListTaskGroupsResponseBodyDataItems struct {
	// example:
	//
	// 0
	CompletedTasks *int32 `json:"CompletedTasks,omitempty" xml:"CompletedTasks,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// default
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// runnable
	Status  *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalTasks *int32 `json:"TotalTasks,omitempty" xml:"TotalTasks,omitempty"`
}

func (s ListTaskGroupsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListTaskGroupsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListTaskGroupsResponseBodyDataItems) SetCompletedTasks(v int32) *ListTaskGroupsResponseBodyDataItems {
	s.CompletedTasks = &v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetCreatedAt(v string) *ListTaskGroupsResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetId(v string) *ListTaskGroupsResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetName(v string) *ListTaskGroupsResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetRuleId(v string) *ListTaskGroupsResponseBodyDataItems {
	s.RuleId = &v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetRuleName(v string) *ListTaskGroupsResponseBodyDataItems {
	s.RuleName = &v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetStatus(v string) *ListTaskGroupsResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetTaskIds(v []*string) *ListTaskGroupsResponseBodyDataItems {
	s.TaskIds = v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetTotalTasks(v int32) *ListTaskGroupsResponseBodyDataItems {
	s.TotalTasks = &v
	return s
}

type ListTaskGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaskGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListTaskGroupsResponse) SetHeaders(v map[string]*string) *ListTaskGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListTaskGroupsResponse) SetStatusCode(v int32) *ListTaskGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskGroupsResponse) SetBody(v *ListTaskGroupsResponseBody) *ListTaskGroupsResponse {
	s.Body = v
	return s
}

type ListTaskItemsRequest struct {
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListTaskItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaskItemsRequest) GoString() string {
	return s.String()
}

func (s *ListTaskItemsRequest) SetTaskId(v string) *ListTaskItemsRequest {
	s.TaskId = &v
	return s
}

type ListTaskItemsResponseBody struct {
	// example:
	//
	// OK
	Code    *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    []*ListTaskItemsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTaskItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTaskItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskItemsResponseBody) SetCode(v string) *ListTaskItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTaskItemsResponseBody) SetData(v []*ListTaskItemsResponseBodyData) *ListTaskItemsResponseBody {
	s.Data = v
	return s
}

func (s *ListTaskItemsResponseBody) SetMessage(v string) *ListTaskItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTaskItemsResponseBody) SetRequestId(v string) *ListTaskItemsResponseBody {
	s.RequestId = &v
	return s
}

type ListTaskItemsResponseBodyData struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// example:
	//
	// 1
	SegmentSeq *int64 `json:"SegmentSeq,omitempty" xml:"SegmentSeq,omitempty"`
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListTaskItemsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTaskItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTaskItemsResponseBodyData) SetCreatedAt(v string) *ListTaskItemsResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListTaskItemsResponseBodyData) SetMessage(v string) *ListTaskItemsResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListTaskItemsResponseBodyData) SetName(v string) *ListTaskItemsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListTaskItemsResponseBodyData) SetOutput(v string) *ListTaskItemsResponseBodyData {
	s.Output = &v
	return s
}

func (s *ListTaskItemsResponseBodyData) SetSegmentSeq(v int64) *ListTaskItemsResponseBodyData {
	s.SegmentSeq = &v
	return s
}

func (s *ListTaskItemsResponseBodyData) SetStatus(v string) *ListTaskItemsResponseBodyData {
	s.Status = &v
	return s
}

type ListTaskItemsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaskItemsResponse) GoString() string {
	return s.String()
}

func (s *ListTaskItemsResponse) SetHeaders(v map[string]*string) *ListTaskItemsResponse {
	s.Headers = v
	return s
}

func (s *ListTaskItemsResponse) SetStatusCode(v int32) *ListTaskItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskItemsResponse) SetBody(v *ListTaskItemsResponseBody) *ListTaskItemsResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TaskGroupId *string `json:"TaskGroupId,omitempty" xml:"TaskGroupId,omitempty"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetPageIndex(v int32) *ListTasksRequest {
	s.PageIndex = &v
	return s
}

func (s *ListTasksRequest) SetPageSize(v int32) *ListTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksRequest) SetTaskGroupId(v string) *ListTasksRequest {
	s.TaskGroupId = &v
	return s
}

type ListTasksResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ListTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) SetCode(v string) *ListTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListTasksResponseBody) SetData(v *ListTasksResponseBodyData) *ListTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListTasksResponseBody) SetMessage(v string) *ListTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

type ListTasksResponseBodyData struct {
	Items []*ListTasksResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyData) SetItems(v []*ListTasksResponseBodyDataItems) *ListTasksResponseBodyData {
	s.Items = v
	return s
}

func (s *ListTasksResponseBodyData) SetTotalElements(v int64) *ListTasksResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTotalPages(v int32) *ListTasksResponseBodyData {
	s.TotalPages = &v
	return s
}

type ListTasksResponseBodyDataItems struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// runnable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.meta?expire=2010323&sign=sf2324
	VideoMetaUrl *string `json:"VideoMetaUrl,omitempty" xml:"VideoMetaUrl,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4?expire=2010323&sign=sf2324
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s ListTasksResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyDataItems) SetCreatedAt(v string) *ListTasksResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListTasksResponseBodyDataItems) SetId(v string) *ListTasksResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListTasksResponseBodyDataItems) SetStatus(v string) *ListTasksResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListTasksResponseBodyDataItems) SetVideoMetaUrl(v string) *ListTasksResponseBodyDataItems {
	s.VideoMetaUrl = &v
	return s
}

func (s *ListTasksResponseBodyDataItems) SetVideoUrl(v string) *ListTasksResponseBodyDataItems {
	s.VideoUrl = &v
	return s
}

type ListTasksResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTasksResponse) SetHeaders(v map[string]*string) *ListTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTasksResponse) SetStatusCode(v int32) *ListTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTasksResponse) SetBody(v *ListTasksResponseBody) *ListTasksResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// tom
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetDepartmentId(v string) *ListUsersRequest {
	s.DepartmentId = &v
	return s
}

func (s *ListUsersRequest) SetPageIndex(v int32) *ListUsersRequest {
	s.PageIndex = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) SetUsername(v string) *ListUsersRequest {
	s.Username = &v
	return s
}

type ListUsersResponseBody struct {
	// example:
	//
	// OK
	Code    *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *ListUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetCode(v string) *ListUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUsersResponseBody) SetData(v *ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetMessage(v string) *ListUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

type ListUsersResponseBodyData struct {
	Items []*ListUsersResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) SetItems(v []*ListUsersResponseBodyDataItems) *ListUsersResponseBodyData {
	s.Items = v
	return s
}

func (s *ListUsersResponseBodyData) SetTotalElements(v int64) *ListUsersResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListUsersResponseBodyData) SetTotalPages(v int32) *ListUsersResponseBodyData {
	s.TotalPages = &v
	return s
}

type ListUsersResponseBodyDataItems struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt   *string                                      `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Departments []*ListUsersResponseBodyDataItemsDepartments `json:"Departments,omitempty" xml:"Departments,omitempty" type:"Repeated"`
	// example:
	//
	// a@a.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// ID
	//
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 186000000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// xxx
	RamUsername *string `json:"RamUsername,omitempty" xml:"RamUsername,omitempty"`
	// example:
	//
	// admin
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// ram
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// example:
	//
	// xxx
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListUsersResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataItems) SetCreatedAt(v string) *ListUsersResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetDepartments(v []*ListUsersResponseBodyDataItemsDepartments) *ListUsersResponseBodyDataItems {
	s.Departments = v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetEmail(v string) *ListUsersResponseBodyDataItems {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetId(v string) *ListUsersResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetName(v string) *ListUsersResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetPhoneNumber(v string) *ListUsersResponseBodyDataItems {
	s.PhoneNumber = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetRamUsername(v string) *ListUsersResponseBodyDataItems {
	s.RamUsername = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetRole(v string) *ListUsersResponseBodyDataItems {
	s.Role = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetSource(v string) *ListUsersResponseBodyDataItems {
	s.Source = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetUpdatedAt(v string) *ListUsersResponseBodyDataItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetUsername(v string) *ListUsersResponseBodyDataItems {
	s.Username = &v
	return s
}

type ListUsersResponseBodyDataItemsDepartments struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListUsersResponseBodyDataItemsDepartments) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyDataItemsDepartments) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataItemsDepartments) SetCreatedAt(v string) *ListUsersResponseBodyDataItemsDepartments {
	s.CreatedAt = &v
	return s
}

func (s *ListUsersResponseBodyDataItemsDepartments) SetDescription(v string) *ListUsersResponseBodyDataItemsDepartments {
	s.Description = &v
	return s
}

func (s *ListUsersResponseBodyDataItemsDepartments) SetId(v string) *ListUsersResponseBodyDataItemsDepartments {
	s.Id = &v
	return s
}

func (s *ListUsersResponseBodyDataItemsDepartments) SetName(v string) *ListUsersResponseBodyDataItemsDepartments {
	s.Name = &v
	return s
}

func (s *ListUsersResponseBodyDataItemsDepartments) SetUpdatedAt(v string) *ListUsersResponseBodyDataItemsDepartments {
	s.UpdatedAt = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ListWatermarksRequest struct {
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListWatermarksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWatermarksRequest) GoString() string {
	return s.String()
}

func (s *ListWatermarksRequest) SetPageIndex(v int32) *ListWatermarksRequest {
	s.PageIndex = &v
	return s
}

func (s *ListWatermarksRequest) SetPageSize(v int32) *ListWatermarksRequest {
	s.PageSize = &v
	return s
}

type ListWatermarksResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code   *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data   *ListWatermarksResponseBodyData     `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Errors []*ListWatermarksResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWatermarksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWatermarksResponseBody) GoString() string {
	return s.String()
}

func (s *ListWatermarksResponseBody) SetCode(v string) *ListWatermarksResponseBody {
	s.Code = &v
	return s
}

func (s *ListWatermarksResponseBody) SetData(v *ListWatermarksResponseBodyData) *ListWatermarksResponseBody {
	s.Data = v
	return s
}

func (s *ListWatermarksResponseBody) SetErrors(v []*ListWatermarksResponseBodyErrors) *ListWatermarksResponseBody {
	s.Errors = v
	return s
}

func (s *ListWatermarksResponseBody) SetHttpCode(v int32) *ListWatermarksResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListWatermarksResponseBody) SetMessage(v string) *ListWatermarksResponseBody {
	s.Message = &v
	return s
}

func (s *ListWatermarksResponseBody) SetRequestId(v string) *ListWatermarksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWatermarksResponseBody) SetSuccess(v bool) *ListWatermarksResponseBody {
	s.Success = &v
	return s
}

type ListWatermarksResponseBodyData struct {
	Items []*ListWatermarksResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListWatermarksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWatermarksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWatermarksResponseBodyData) SetItems(v []*ListWatermarksResponseBodyDataItems) *ListWatermarksResponseBodyData {
	s.Items = v
	return s
}

func (s *ListWatermarksResponseBodyData) SetTotalElements(v int64) *ListWatermarksResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListWatermarksResponseBodyData) SetTotalPages(v int32) *ListWatermarksResponseBodyData {
	s.TotalPages = &v
	return s
}

type ListWatermarksResponseBodyDataItems struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListWatermarksResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListWatermarksResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListWatermarksResponseBodyDataItems) SetCreatedAt(v string) *ListWatermarksResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListWatermarksResponseBodyDataItems) SetId(v string) *ListWatermarksResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListWatermarksResponseBodyDataItems) SetName(v string) *ListWatermarksResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListWatermarksResponseBodyDataItems) SetValue(v string) *ListWatermarksResponseBodyDataItems {
	s.Value = &v
	return s
}

type ListWatermarksResponseBodyErrors struct {
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListWatermarksResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s ListWatermarksResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *ListWatermarksResponseBodyErrors) SetField(v string) *ListWatermarksResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *ListWatermarksResponseBodyErrors) SetMessage(v string) *ListWatermarksResponseBodyErrors {
	s.Message = &v
	return s
}

type ListWatermarksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWatermarksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWatermarksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWatermarksResponse) GoString() string {
	return s.String()
}

func (s *ListWatermarksResponse) SetHeaders(v map[string]*string) *ListWatermarksResponse {
	s.Headers = v
	return s
}

func (s *ListWatermarksResponse) SetStatusCode(v int32) *ListWatermarksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWatermarksResponse) SetBody(v *ListWatermarksResponseBody) *ListWatermarksResponse {
	s.Body = v
	return s
}

type RenameDetectProcessRequest struct {
	// ID
	//
	// example:
	//
	// 1q1c45cd-3eee-1e60-b505-2e330b8755d2
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s RenameDetectProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameDetectProcessRequest) GoString() string {
	return s.String()
}

func (s *RenameDetectProcessRequest) SetId(v string) *RenameDetectProcessRequest {
	s.Id = &v
	return s
}

func (s *RenameDetectProcessRequest) SetName(v string) *RenameDetectProcessRequest {
	s.Name = &v
	return s
}

type RenameDetectProcessResponseBody struct {
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *RenameDetectProcessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// db1c45cd-3eee-1e60-b505-2e330b8755d2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenameDetectProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameDetectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *RenameDetectProcessResponseBody) SetCode(v string) *RenameDetectProcessResponseBody {
	s.Code = &v
	return s
}

func (s *RenameDetectProcessResponseBody) SetData(v *RenameDetectProcessResponseBodyData) *RenameDetectProcessResponseBody {
	s.Data = v
	return s
}

func (s *RenameDetectProcessResponseBody) SetMessage(v string) *RenameDetectProcessResponseBody {
	s.Message = &v
	return s
}

func (s *RenameDetectProcessResponseBody) SetRequestId(v string) *RenameDetectProcessResponseBody {
	s.RequestId = &v
	return s
}

type RenameDetectProcessResponseBodyData struct {
	// example:
	//
	// {}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2020-12-04T14:47:59.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// {}
	Draft *string `json:"Draft,omitempty" xml:"Draft,omitempty"`
	// ID
	//
	// example:
	//
	// 0f1c45cd-3eee-4e60-b505-2e330b8755d3
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 987d563d38f5aef27feca8702c689bb1
	Md5  *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s RenameDetectProcessResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RenameDetectProcessResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenameDetectProcessResponseBodyData) SetContent(v string) *RenameDetectProcessResponseBodyData {
	s.Content = &v
	return s
}

func (s *RenameDetectProcessResponseBodyData) SetCreatedAt(v string) *RenameDetectProcessResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *RenameDetectProcessResponseBodyData) SetDraft(v string) *RenameDetectProcessResponseBodyData {
	s.Draft = &v
	return s
}

func (s *RenameDetectProcessResponseBodyData) SetId(v string) *RenameDetectProcessResponseBodyData {
	s.Id = &v
	return s
}

func (s *RenameDetectProcessResponseBodyData) SetMd5(v string) *RenameDetectProcessResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *RenameDetectProcessResponseBodyData) SetName(v string) *RenameDetectProcessResponseBodyData {
	s.Name = &v
	return s
}

type RenameDetectProcessResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameDetectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameDetectProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameDetectProcessResponse) GoString() string {
	return s.String()
}

func (s *RenameDetectProcessResponse) SetHeaders(v map[string]*string) *RenameDetectProcessResponse {
	s.Headers = v
	return s
}

func (s *RenameDetectProcessResponse) SetStatusCode(v int32) *RenameDetectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameDetectProcessResponse) SetBody(v *RenameDetectProcessResponseBody) *RenameDetectProcessResponse {
	s.Body = v
	return s
}

type TtsCommonRequest struct {
	TtsRequest *TtsCommonRequestTtsRequest `json:"TtsRequest,omitempty" xml:"TtsRequest,omitempty" type:"Struct"`
}

func (s TtsCommonRequest) String() string {
	return tea.Prettify(s)
}

func (s TtsCommonRequest) GoString() string {
	return s.String()
}

func (s *TtsCommonRequest) SetTtsRequest(v *TtsCommonRequestTtsRequest) *TtsCommonRequest {
	s.TtsRequest = v
	return s
}

type TtsCommonRequestTtsRequest struct {
	// appid
	//
	// example:
	//
	// d61be709-49d2-4cf1-b219-cd6181f72***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// WAV
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// 50
	PitchRate *int32 `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// example:
	//
	// 16000
	SampleRate *int32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// example:
	//
	// 50
	SpeechRate *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Text       *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// Xiaoyun
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s TtsCommonRequestTtsRequest) String() string {
	return tea.Prettify(s)
}

func (s TtsCommonRequestTtsRequest) GoString() string {
	return s.String()
}

func (s *TtsCommonRequestTtsRequest) SetAppId(v string) *TtsCommonRequestTtsRequest {
	s.AppId = &v
	return s
}

func (s *TtsCommonRequestTtsRequest) SetFormat(v string) *TtsCommonRequestTtsRequest {
	s.Format = &v
	return s
}

func (s *TtsCommonRequestTtsRequest) SetPitchRate(v int32) *TtsCommonRequestTtsRequest {
	s.PitchRate = &v
	return s
}

func (s *TtsCommonRequestTtsRequest) SetSampleRate(v int32) *TtsCommonRequestTtsRequest {
	s.SampleRate = &v
	return s
}

func (s *TtsCommonRequestTtsRequest) SetSpeechRate(v int32) *TtsCommonRequestTtsRequest {
	s.SpeechRate = &v
	return s
}

func (s *TtsCommonRequestTtsRequest) SetText(v string) *TtsCommonRequestTtsRequest {
	s.Text = &v
	return s
}

func (s *TtsCommonRequestTtsRequest) SetVoice(v string) *TtsCommonRequestTtsRequest {
	s.Voice = &v
	return s
}

func (s *TtsCommonRequestTtsRequest) SetVolume(v int32) *TtsCommonRequestTtsRequest {
	s.Volume = &v
	return s
}

type TtsCommonShrinkRequest struct {
	TtsRequestShrink *string `json:"TtsRequest,omitempty" xml:"TtsRequest,omitempty"`
}

func (s TtsCommonShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s TtsCommonShrinkRequest) GoString() string {
	return s.String()
}

func (s *TtsCommonShrinkRequest) SetTtsRequestShrink(v string) *TtsCommonShrinkRequest {
	s.TtsRequestShrink = &v
	return s
}

type TtsCommonResponseBody struct {
	// example:
	//
	// OK
	Code *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TtsCommonResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// c761f0ec-107c-43dc-8565-45330e10e11b
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TtsCommonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TtsCommonResponseBody) GoString() string {
	return s.String()
}

func (s *TtsCommonResponseBody) SetCode(v int32) *TtsCommonResponseBody {
	s.Code = &v
	return s
}

func (s *TtsCommonResponseBody) SetData(v *TtsCommonResponseBodyData) *TtsCommonResponseBody {
	s.Data = v
	return s
}

func (s *TtsCommonResponseBody) SetHttpCode(v int32) *TtsCommonResponseBody {
	s.HttpCode = &v
	return s
}

func (s *TtsCommonResponseBody) SetMessage(v string) *TtsCommonResponseBody {
	s.Message = &v
	return s
}

func (s *TtsCommonResponseBody) SetRequestId(v string) *TtsCommonResponseBody {
	s.RequestId = &v
	return s
}

func (s *TtsCommonResponseBody) SetSuccess(v bool) *TtsCommonResponseBody {
	s.Success = &v
	return s
}

type TtsCommonResponseBodyData struct {
	// example:
	//
	// 20000000
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 63bb629d-92bf-4cdc-ad0b-3032c926****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// http://idrs-daily.oss-cn-beijing.aliyuncs.com/mode/android/model.zip?Expires=1607004612&OSSAccessKeyId=LTAI4FcsdDhFc7h78gqB****&Signature=XXXXX
	PublicUrl *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
	// example:
	//
	// 368cfa55-2364-4d79-aeb4-c0956c4a****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TtsCommonResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TtsCommonResponseBodyData) GoString() string {
	return s.String()
}

func (s *TtsCommonResponseBodyData) SetCode(v int32) *TtsCommonResponseBodyData {
	s.Code = &v
	return s
}

func (s *TtsCommonResponseBodyData) SetId(v string) *TtsCommonResponseBodyData {
	s.Id = &v
	return s
}

func (s *TtsCommonResponseBodyData) SetMessage(v string) *TtsCommonResponseBodyData {
	s.Message = &v
	return s
}

func (s *TtsCommonResponseBodyData) SetName(v string) *TtsCommonResponseBodyData {
	s.Name = &v
	return s
}

func (s *TtsCommonResponseBodyData) SetPublicUrl(v string) *TtsCommonResponseBodyData {
	s.PublicUrl = &v
	return s
}

func (s *TtsCommonResponseBodyData) SetTaskId(v string) *TtsCommonResponseBodyData {
	s.TaskId = &v
	return s
}

type TtsCommonResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TtsCommonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TtsCommonResponse) String() string {
	return tea.Prettify(s)
}

func (s TtsCommonResponse) GoString() string {
	return s.String()
}

func (s *TtsCommonResponse) SetHeaders(v map[string]*string) *TtsCommonResponse {
	s.Headers = v
	return s
}

func (s *TtsCommonResponse) SetStatusCode(v int32) *TtsCommonResponse {
	s.StatusCode = &v
	return s
}

func (s *TtsCommonResponse) SetBody(v *TtsCommonResponseBody) *TtsCommonResponse {
	s.Body = v
	return s
}

type TtsTaskRequest struct {
	Request *TtsTaskRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s TtsTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s TtsTaskRequest) GoString() string {
	return s.String()
}

func (s *TtsTaskRequest) SetRequest(v *TtsTaskRequestRequest) *TtsTaskRequest {
	s.Request = v
	return s
}

type TtsTaskRequestRequest struct {
	// example:
	//
	// d9ee5df9-20bf-47bf-987a-76b26984b***
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 655259585579***
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// example:
	//
	// 1662271315039
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s TtsTaskRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s TtsTaskRequestRequest) GoString() string {
	return s.String()
}

func (s *TtsTaskRequestRequest) SetKey(v string) *TtsTaskRequestRequest {
	s.Key = &v
	return s
}

func (s *TtsTaskRequestRequest) SetRoomId(v string) *TtsTaskRequestRequest {
	s.RoomId = &v
	return s
}

func (s *TtsTaskRequestRequest) SetTimestamp(v int64) *TtsTaskRequestRequest {
	s.Timestamp = &v
	return s
}

type TtsTaskShrinkRequest struct {
	RequestShrink *string `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s TtsTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s TtsTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *TtsTaskShrinkRequest) SetRequestShrink(v string) *TtsTaskShrinkRequest {
	s.RequestShrink = &v
	return s
}

type TtsTaskResponseBody struct {
	// example:
	//
	// OK
	Code *int32                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TtsTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 16A4A091-98BE-55F4-92D1-21590658***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TtsTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TtsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *TtsTaskResponseBody) SetCode(v int32) *TtsTaskResponseBody {
	s.Code = &v
	return s
}

func (s *TtsTaskResponseBody) SetData(v *TtsTaskResponseBodyData) *TtsTaskResponseBody {
	s.Data = v
	return s
}

func (s *TtsTaskResponseBody) SetHttpCode(v int32) *TtsTaskResponseBody {
	s.HttpCode = &v
	return s
}

func (s *TtsTaskResponseBody) SetMessage(v string) *TtsTaskResponseBody {
	s.Message = &v
	return s
}

func (s *TtsTaskResponseBody) SetRequestId(v string) *TtsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *TtsTaskResponseBody) SetSuccess(v bool) *TtsTaskResponseBody {
	s.Success = &v
	return s
}

type TtsTaskResponseBodyData struct {
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// example:
	//
	// 727
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Question *string  `json:"Question,omitempty" xml:"Question,omitempty"`
	// example:
	//
	// 64
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
}

func (s TtsTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TtsTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *TtsTaskResponseBodyData) SetAnswer(v string) *TtsTaskResponseBodyData {
	s.Answer = &v
	return s
}

func (s *TtsTaskResponseBodyData) SetDuration(v float64) *TtsTaskResponseBodyData {
	s.Duration = &v
	return s
}

func (s *TtsTaskResponseBodyData) SetQuestion(v string) *TtsTaskResponseBodyData {
	s.Question = &v
	return s
}

func (s *TtsTaskResponseBodyData) SetSpeechRate(v int32) *TtsTaskResponseBodyData {
	s.SpeechRate = &v
	return s
}

type TtsTaskResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TtsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TtsTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s TtsTaskResponse) GoString() string {
	return s.String()
}

func (s *TtsTaskResponse) SetHeaders(v map[string]*string) *TtsTaskResponse {
	s.Headers = v
	return s
}

func (s *TtsTaskResponse) SetStatusCode(v int32) *TtsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *TtsTaskResponse) SetBody(v *TtsTaskResponseBody) *TtsTaskResponse {
	s.Body = v
	return s
}

type UpdateAppRequest struct {
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// com.aliyun.idrs.sdk
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
}

func (s UpdateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppRequest) SetDepartmentId(v string) *UpdateAppRequest {
	s.DepartmentId = &v
	return s
}

func (s *UpdateAppRequest) SetDisabled(v bool) *UpdateAppRequest {
	s.Disabled = &v
	return s
}

func (s *UpdateAppRequest) SetId(v string) *UpdateAppRequest {
	s.Id = &v
	return s
}

func (s *UpdateAppRequest) SetName(v string) *UpdateAppRequest {
	s.Name = &v
	return s
}

func (s *UpdateAppRequest) SetPackageName(v string) *UpdateAppRequest {
	s.PackageName = &v
	return s
}

type UpdateAppResponseBody struct {
	// example:
	//
	// OK
	Code    *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 84118BF0-56F7-54D2-8C1A-35BBBBD5016B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBody) SetCode(v string) *UpdateAppResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAppResponseBody) SetData(v map[string]interface{}) *UpdateAppResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAppResponseBody) SetMessage(v string) *UpdateAppResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppResponseBody) SetRequestId(v string) *UpdateAppResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppResponse) SetHeaders(v map[string]*string) *UpdateAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppResponse) SetStatusCode(v int32) *UpdateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppResponse) SetBody(v *UpdateAppResponseBody) *UpdateAppResponse {
	s.Body = v
	return s
}

type UpdateDepartmentRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDepartmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentRequest) SetDescription(v string) *UpdateDepartmentRequest {
	s.Description = &v
	return s
}

func (s *UpdateDepartmentRequest) SetId(v string) *UpdateDepartmentRequest {
	s.Id = &v
	return s
}

func (s *UpdateDepartmentRequest) SetLabel(v string) *UpdateDepartmentRequest {
	s.Label = &v
	return s
}

func (s *UpdateDepartmentRequest) SetName(v string) *UpdateDepartmentRequest {
	s.Name = &v
	return s
}

type UpdateDepartmentResponseBody struct {
	// example:
	//
	// OK
	Code    *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentResponseBody) SetCode(v string) *UpdateDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetData(v map[string]interface{}) *UpdateDepartmentResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDepartmentResponseBody) SetMessage(v string) *UpdateDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetRequestId(v string) *UpdateDepartmentResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDepartmentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDepartmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentResponse) SetHeaders(v map[string]*string) *UpdateDepartmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateDepartmentResponse) SetStatusCode(v int32) *UpdateDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDepartmentResponse) SetBody(v *UpdateDepartmentResponseBody) *UpdateDepartmentResponse {
	s.Body = v
	return s
}

type UpdateDetectProcessRequest struct {
	// example:
	//
	// {}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// {}
	Draft *string `json:"Draft,omitempty" xml:"Draft,omitempty"`
	// ID
	//
	// example:
	//
	// 0f1c45cd-3eee-4e60-b505-2e330b8755d3
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateDetectProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectProcessRequest) GoString() string {
	return s.String()
}

func (s *UpdateDetectProcessRequest) SetContent(v string) *UpdateDetectProcessRequest {
	s.Content = &v
	return s
}

func (s *UpdateDetectProcessRequest) SetDraft(v string) *UpdateDetectProcessRequest {
	s.Draft = &v
	return s
}

func (s *UpdateDetectProcessRequest) SetId(v string) *UpdateDetectProcessRequest {
	s.Id = &v
	return s
}

func (s *UpdateDetectProcessRequest) SetName(v string) *UpdateDetectProcessRequest {
	s.Name = &v
	return s
}

type UpdateDetectProcessResponseBody struct {
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateDetectProcessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0c1c45cd-3eee-4e60-b505-2e330b8755d3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDetectProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDetectProcessResponseBody) SetCode(v string) *UpdateDetectProcessResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDetectProcessResponseBody) SetData(v *UpdateDetectProcessResponseBodyData) *UpdateDetectProcessResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDetectProcessResponseBody) SetMessage(v string) *UpdateDetectProcessResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDetectProcessResponseBody) SetRequestId(v string) *UpdateDetectProcessResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDetectProcessResponseBodyData struct {
	// example:
	//
	// {}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2020-12-04T14:47:59.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// {}
	Draft *string `json:"Draft,omitempty" xml:"Draft,omitempty"`
	// ID
	//
	// example:
	//
	// cd1c45cd-3eee-4e60-b505-2e330b8755d3
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 987d563d38f5aef27feca8702c689bb1
	Md5  *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateDetectProcessResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectProcessResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDetectProcessResponseBodyData) SetContent(v string) *UpdateDetectProcessResponseBodyData {
	s.Content = &v
	return s
}

func (s *UpdateDetectProcessResponseBodyData) SetCreatedAt(v string) *UpdateDetectProcessResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateDetectProcessResponseBodyData) SetDraft(v string) *UpdateDetectProcessResponseBodyData {
	s.Draft = &v
	return s
}

func (s *UpdateDetectProcessResponseBodyData) SetId(v string) *UpdateDetectProcessResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateDetectProcessResponseBodyData) SetMd5(v string) *UpdateDetectProcessResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *UpdateDetectProcessResponseBodyData) SetName(v string) *UpdateDetectProcessResponseBodyData {
	s.Name = &v
	return s
}

type UpdateDetectProcessResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDetectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDetectProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectProcessResponse) GoString() string {
	return s.String()
}

func (s *UpdateDetectProcessResponse) SetHeaders(v map[string]*string) *UpdateDetectProcessResponse {
	s.Headers = v
	return s
}

func (s *UpdateDetectProcessResponse) SetStatusCode(v int32) *UpdateDetectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDetectProcessResponse) SetBody(v *UpdateDetectProcessResponseBody) *UpdateDetectProcessResponse {
	s.Body = v
	return s
}

type UpdateRuleRequest struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// ID
	//
	// example:
	//
	// Id
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleRequest) SetContent(v string) *UpdateRuleRequest {
	s.Content = &v
	return s
}

func (s *UpdateRuleRequest) SetId(v string) *UpdateRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateRuleRequest) SetName(v string) *UpdateRuleRequest {
	s.Name = &v
	return s
}

type UpdateRuleResponseBody struct {
	// example:
	//
	// OK
	Code    *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *UpdateRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponseBody) SetCode(v string) *UpdateRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRuleResponseBody) SetData(v *UpdateRuleResponseBodyData) *UpdateRuleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateRuleResponseBody) SetMessage(v string) *UpdateRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRuleResponseBody) SetRequestId(v string) *UpdateRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRuleResponseBodyData struct {
	// example:
	//
	// {}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// ID
	//
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponseBodyData) SetContent(v string) *UpdateRuleResponseBodyData {
	s.Content = &v
	return s
}

func (s *UpdateRuleResponseBodyData) SetCreatedAt(v string) *UpdateRuleResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateRuleResponseBodyData) SetId(v string) *UpdateRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateRuleResponseBodyData) SetName(v string) *UpdateRuleResponseBodyData {
	s.Name = &v
	return s
}

type UpdateRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponse) SetHeaders(v map[string]*string) *UpdateRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRuleResponse) SetStatusCode(v int32) *UpdateRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRuleResponse) SetBody(v *UpdateRuleResponseBody) *UpdateRuleResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	// example:
	//
	// xxx@xxx.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetEmail(v string) *UpdateUserRequest {
	s.Email = &v
	return s
}

func (s *UpdateUserRequest) SetId(v string) *UpdateUserRequest {
	s.Id = &v
	return s
}

func (s *UpdateUserRequest) SetName(v string) *UpdateUserRequest {
	s.Name = &v
	return s
}

func (s *UpdateUserRequest) SetPhoneNumber(v string) *UpdateUserRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateUserRequest) SetRole(v string) *UpdateUserRequest {
	s.Role = &v
	return s
}

type UpdateUserResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {\"Id\": \"5073e5f4-99c6-4bb1-bd6c-30ab12f11059\", \"CreatedAt\": \"2021-12-29T11:31:53.072+08:00\"}
	Data    map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 84118BF0-56F7-54D2-8C1A-35BBBBD5016B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) SetCode(v string) *UpdateUserResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserResponseBody) SetData(v map[string]interface{}) *UpdateUserResponseBody {
	s.Data = v
	return s
}

func (s *UpdateUserResponseBody) SetMessage(v string) *UpdateUserResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetStatusCode(v int32) *UpdateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
	s.Body = v
	return s
}

type UpdateWatermarkRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// d4ba1e0428a8df069316060cef41948a
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s UpdateWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWatermarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkRequest) SetName(v string) *UpdateWatermarkRequest {
	s.Name = &v
	return s
}

func (s *UpdateWatermarkRequest) SetValue(v string) *UpdateWatermarkRequest {
	s.Value = &v
	return s
}

func (s *UpdateWatermarkRequest) SetWatermarkId(v string) *UpdateWatermarkRequest {
	s.WatermarkId = &v
	return s
}

type UpdateWatermarkResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code   *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data   *UpdateWatermarkResponseBodyData     `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Errors []*UpdateWatermarkResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// 84118BF0-56F7-54D2-8C1A-35BBBBD5016B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkResponseBody) SetCode(v string) *UpdateWatermarkResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWatermarkResponseBody) SetData(v *UpdateWatermarkResponseBodyData) *UpdateWatermarkResponseBody {
	s.Data = v
	return s
}

func (s *UpdateWatermarkResponseBody) SetErrors(v []*UpdateWatermarkResponseBodyErrors) *UpdateWatermarkResponseBody {
	s.Errors = v
	return s
}

func (s *UpdateWatermarkResponseBody) SetHttpCode(v int32) *UpdateWatermarkResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateWatermarkResponseBody) SetMessage(v string) *UpdateWatermarkResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWatermarkResponseBody) SetRequestId(v string) *UpdateWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWatermarkResponseBody) SetSuccess(v bool) *UpdateWatermarkResponseBody {
	s.Success = &v
	return s
}

type UpdateWatermarkResponseBodyData struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// f3bd31c0-0001-4b4b-977d-7cfa64b585f5
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateWatermarkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateWatermarkResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkResponseBodyData) SetCreatedAt(v string) *UpdateWatermarkResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateWatermarkResponseBodyData) SetId(v string) *UpdateWatermarkResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateWatermarkResponseBodyData) SetName(v string) *UpdateWatermarkResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateWatermarkResponseBodyData) SetValue(v string) *UpdateWatermarkResponseBodyData {
	s.Value = &v
	return s
}

type UpdateWatermarkResponseBodyErrors struct {
	// example:
	//
	// A1899517-BB99-5D3E-A71B-97524DCB0942
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UpdateWatermarkResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s UpdateWatermarkResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkResponseBodyErrors) SetField(v string) *UpdateWatermarkResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *UpdateWatermarkResponseBodyErrors) SetMessage(v string) *UpdateWatermarkResponseBodyErrors {
	s.Message = &v
	return s
}

type UpdateWatermarkResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWatermarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkResponse) SetHeaders(v map[string]*string) *UpdateWatermarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateWatermarkResponse) SetStatusCode(v int32) *UpdateWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWatermarkResponse) SetBody(v *UpdateWatermarkResponseBody) *UpdateWatermarkResponse {
	s.Body = v
	return s
}

type UploadReportRequest struct {
	// example:
	//
	// 90c2459a-4c73-4ab7-8a6b-e550d62fbd8c
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// {"version":"1.0.0"}
	ClientBaseParam *string `json:"ClientBaseParam,omitempty" xml:"ClientBaseParam,omitempty"`
	// example:
	//
	// 1.0.003
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad38
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// example:
	//
	// iuhptk3w-2021122017
	DetectProcessId *string `json:"DetectProcessId,omitempty" xml:"DetectProcessId,omitempty"`
	// example:
	//
	// 12
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 6c94f2a7-632d-4ea0-aa06-a97800a9060f
	FeeId *string `json:"FeeId,omitempty" xml:"FeeId,omitempty"`
	// example:
	//
	// http://abc.oss.aliyuncs.com/1/12_03_20_03_36toubaoren.mp4.meta
	MetaUrl *string `json:"MetaUrl,omitempty" xml:"MetaUrl,omitempty"`
	// example:
	//
	// ads32efef43
	OuterBusinessId *string `json:"OuterBusinessId,omitempty" xml:"OuterBusinessId,omitempty"`
	// example:
	//
	// 2022-01-04T08%3A45%3A37Z
	RecordAt *string `json:"RecordAt,omitempty" xml:"RecordAt,omitempty"`
	// example:
	//
	// http://abc.oss.aliyuncs.com/1/12_03_20_03_36toubaoren.mp4.json
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
	Role      *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 6000028888875
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// example:
	//
	// record_6570315480571566****
	RtcRecordId *string `json:"RtcRecordId,omitempty" xml:"RtcRecordId,omitempty"`
	// example:
	//
	// LOCAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce70b4bad34
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VideoType *string `json:"VideoType,omitempty" xml:"VideoType,omitempty"`
	// example:
	//
	// https://mogo-apps-sh.oss-cn-shanghai-internal.aliyuncs.com/tmp/d2720028b53d384c6d3fca32e45209d1_20_src.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s UploadReportRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadReportRequest) GoString() string {
	return s.String()
}

func (s *UploadReportRequest) SetAppId(v string) *UploadReportRequest {
	s.AppId = &v
	return s
}

func (s *UploadReportRequest) SetClientBaseParam(v string) *UploadReportRequest {
	s.ClientBaseParam = &v
	return s
}

func (s *UploadReportRequest) SetClientVersion(v string) *UploadReportRequest {
	s.ClientVersion = &v
	return s
}

func (s *UploadReportRequest) SetDepartmentId(v string) *UploadReportRequest {
	s.DepartmentId = &v
	return s
}

func (s *UploadReportRequest) SetDetectProcessId(v string) *UploadReportRequest {
	s.DetectProcessId = &v
	return s
}

func (s *UploadReportRequest) SetDuration(v int64) *UploadReportRequest {
	s.Duration = &v
	return s
}

func (s *UploadReportRequest) SetFeeId(v string) *UploadReportRequest {
	s.FeeId = &v
	return s
}

func (s *UploadReportRequest) SetMetaUrl(v string) *UploadReportRequest {
	s.MetaUrl = &v
	return s
}

func (s *UploadReportRequest) SetOuterBusinessId(v string) *UploadReportRequest {
	s.OuterBusinessId = &v
	return s
}

func (s *UploadReportRequest) SetRecordAt(v string) *UploadReportRequest {
	s.RecordAt = &v
	return s
}

func (s *UploadReportRequest) SetResultUrl(v string) *UploadReportRequest {
	s.ResultUrl = &v
	return s
}

func (s *UploadReportRequest) SetRole(v string) *UploadReportRequest {
	s.Role = &v
	return s
}

func (s *UploadReportRequest) SetRoomId(v string) *UploadReportRequest {
	s.RoomId = &v
	return s
}

func (s *UploadReportRequest) SetRtcRecordId(v string) *UploadReportRequest {
	s.RtcRecordId = &v
	return s
}

func (s *UploadReportRequest) SetType(v string) *UploadReportRequest {
	s.Type = &v
	return s
}

func (s *UploadReportRequest) SetUserId(v string) *UploadReportRequest {
	s.UserId = &v
	return s
}

func (s *UploadReportRequest) SetVideoType(v string) *UploadReportRequest {
	s.VideoType = &v
	return s
}

func (s *UploadReportRequest) SetVideoUrl(v string) *UploadReportRequest {
	s.VideoUrl = &v
	return s
}

type UploadReportResponseBody struct {
	// code
	//
	// example:
	//
	// OK
	Code   *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data   *UploadReportResponseBodyData     `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Errors []*UploadReportResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// 84118BF0-56F7-54D2-8C1A-35BBBB*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadReportResponseBody) GoString() string {
	return s.String()
}

func (s *UploadReportResponseBody) SetCode(v string) *UploadReportResponseBody {
	s.Code = &v
	return s
}

func (s *UploadReportResponseBody) SetData(v *UploadReportResponseBodyData) *UploadReportResponseBody {
	s.Data = v
	return s
}

func (s *UploadReportResponseBody) SetErrors(v []*UploadReportResponseBodyErrors) *UploadReportResponseBody {
	s.Errors = v
	return s
}

func (s *UploadReportResponseBody) SetHttpCode(v int32) *UploadReportResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UploadReportResponseBody) SetMessage(v string) *UploadReportResponseBody {
	s.Message = &v
	return s
}

func (s *UploadReportResponseBody) SetRequestId(v string) *UploadReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadReportResponseBody) SetSuccess(v bool) *UploadReportResponseBody {
	s.Success = &v
	return s
}

type UploadReportResponseBodyData struct {
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 12
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 59b0bbfe-929b-4a8c-9833-3ce7****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.meta
	MetaUrl *string `json:"MetaUrl,omitempty" xml:"MetaUrl,omitempty"`
	// example:
	//
	// ads32efef43
	OuterBusinessId *string `json:"OuterBusinessId,omitempty" xml:"OuterBusinessId,omitempty"`
	// example:
	//
	// 2020-07-14T14:01:41.000+08:00
	RecordAt *string `json:"RecordAt,omitempty" xml:"RecordAt,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4.json
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
	// example:
	//
	// record_6570315480571566****
	RtcRecordId *string `json:"RtcRecordId,omitempty" xml:"RtcRecordId,omitempty"`
	// example:
	//
	// http://oss.aliyuncs.com/1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s UploadReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UploadReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadReportResponseBodyData) SetCreatedAt(v string) *UploadReportResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UploadReportResponseBodyData) SetDuration(v int64) *UploadReportResponseBodyData {
	s.Duration = &v
	return s
}

func (s *UploadReportResponseBodyData) SetId(v string) *UploadReportResponseBodyData {
	s.Id = &v
	return s
}

func (s *UploadReportResponseBodyData) SetMetaUrl(v string) *UploadReportResponseBodyData {
	s.MetaUrl = &v
	return s
}

func (s *UploadReportResponseBodyData) SetOuterBusinessId(v string) *UploadReportResponseBodyData {
	s.OuterBusinessId = &v
	return s
}

func (s *UploadReportResponseBodyData) SetRecordAt(v string) *UploadReportResponseBodyData {
	s.RecordAt = &v
	return s
}

func (s *UploadReportResponseBodyData) SetResultUrl(v string) *UploadReportResponseBodyData {
	s.ResultUrl = &v
	return s
}

func (s *UploadReportResponseBodyData) SetRtcRecordId(v string) *UploadReportResponseBodyData {
	s.RtcRecordId = &v
	return s
}

func (s *UploadReportResponseBodyData) SetVideoUrl(v string) *UploadReportResponseBodyData {
	s.VideoUrl = &v
	return s
}

type UploadReportResponseBodyErrors struct {
	Field   *string `json:"Field,omitempty" xml:"Field,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UploadReportResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s UploadReportResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *UploadReportResponseBodyErrors) SetField(v string) *UploadReportResponseBodyErrors {
	s.Field = &v
	return s
}

func (s *UploadReportResponseBodyErrors) SetMessage(v string) *UploadReportResponseBodyErrors {
	s.Message = &v
	return s
}

type UploadReportResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadReportResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadReportResponse) GoString() string {
	return s.String()
}

func (s *UploadReportResponse) SetHeaders(v map[string]*string) *UploadReportResponse {
	s.Headers = v
	return s
}

func (s *UploadReportResponse) SetStatusCode(v int32) *UploadReportResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadReportResponse) SetBody(v *UploadReportResponseBody) *UploadReportResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("idrsservice.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("idrsservice.aliyuncs.com"),
		"ap-south-1":                  tea.String("idrsservice.aliyuncs.com"),
		"ap-southeast-1":              tea.String("idrsservice.aliyuncs.com"),
		"ap-southeast-2":              tea.String("idrsservice.aliyuncs.com"),
		"ap-southeast-3":              tea.String("idrsservice.aliyuncs.com"),
		"ap-southeast-5":              tea.String("idrsservice.aliyuncs.com"),
		"cn-beijing":                  tea.String("idrsservice.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("idrsservice.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("idrsservice.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("idrsservice.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("idrsservice.aliyuncs.com"),
		"cn-chengdu":                  tea.String("idrsservice.aliyuncs.com"),
		"cn-edge-1":                   tea.String("idrsservice.aliyuncs.com"),
		"cn-fujian":                   tea.String("idrsservice.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("idrsservice.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("idrsservice.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("idrsservice.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("idrsservice.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("idrsservice.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("idrsservice.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("idrsservice.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("idrsservice.aliyuncs.com"),
		"cn-hongkong":                 tea.String("idrsservice.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("idrsservice.aliyuncs.com"),
		"cn-huhehaote":                tea.String("idrsservice.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("idrsservice.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("idrsservice.aliyuncs.com"),
		"cn-qingdao":                  tea.String("idrsservice.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("idrsservice.aliyuncs.com"),
		"cn-shanghai":                 tea.String("idrsservice.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("idrsservice.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("idrsservice.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("idrsservice.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("idrsservice.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("idrsservice.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("idrsservice.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("idrsservice.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("idrsservice.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("idrsservice.aliyuncs.com"),
		"cn-wuhan":                    tea.String("idrsservice.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("idrsservice.aliyuncs.com"),
		"cn-yushanfang":               tea.String("idrsservice.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("idrsservice.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("idrsservice.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("idrsservice.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("idrsservice.aliyuncs.com"),
		"eu-central-1":                tea.String("idrsservice.aliyuncs.com"),
		"eu-west-1":                   tea.String("idrsservice.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("idrsservice.aliyuncs.com"),
		"me-east-1":                   tea.String("idrsservice.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("idrsservice.aliyuncs.com"),
		"us-east-1":                   tea.String("idrsservice.aliyuncs.com"),
		"us-west-1":                   tea.String("idrsservice.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("idrsservice"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AsrRealtimeWithOptions(request *AsrRealtimeRequest, runtime *util.RuntimeOptions) (_result *AsrRealtimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomizationId)) {
		query["CustomizationId"] = request.CustomizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Disfluency)) {
		query["Disfluency"] = request.Disfluency
	}

	if !tea.BoolValue(util.IsUnset(request.EnableIgnoreSentenceTimeout)) {
		query["EnableIgnoreSentenceTimeout"] = request.EnableIgnoreSentenceTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.EnableIntermediateResult)) {
		query["EnableIntermediateResult"] = request.EnableIntermediateResult
	}

	if !tea.BoolValue(util.IsUnset(request.EnableInverseTextNormalization)) {
		query["EnableInverseTextNormalization"] = request.EnableInverseTextNormalization
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePunctuationPrediction)) {
		query["EnablePunctuationPrediction"] = request.EnablePunctuationPrediction
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSemanticSentenceDetection)) {
		query["EnableSemanticSentenceDetection"] = request.EnableSemanticSentenceDetection
	}

	if !tea.BoolValue(util.IsUnset(request.EnableWords)) {
		query["EnableWords"] = request.EnableWords
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["FileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Format)) {
		query["Format"] = request.Format
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSentenceSilence)) {
		query["MaxSentenceSilence"] = request.MaxSentenceSilence
	}

	if !tea.BoolValue(util.IsUnset(request.SampleRate)) {
		query["SampleRate"] = request.SampleRate
	}

	if !tea.BoolValue(util.IsUnset(request.SpeechNoiseThreshold)) {
		query["SpeechNoiseThreshold"] = request.SpeechNoiseThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.VocabularyId)) {
		query["VocabularyId"] = request.VocabularyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AsrRealtime"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AsrRealtimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AsrRealtime(request *AsrRealtimeRequest) (_result *AsrRealtimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AsrRealtimeResponse{}
	_body, _err := client.AsrRealtimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AsrSentenceWithOptions(tmpReq *AsrSentenceRequest, runtime *util.RuntimeOptions) (_result *AsrSentenceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AsrSentenceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AsrRequest)) {
		request.AsrRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AsrRequest, tea.String("AsrRequest"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AsrRequestShrink)) {
		body["AsrRequest"] = request.AsrRequestShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AsrSentence"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AsrSentenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AsrSentence(request *AsrSentenceRequest) (_result *AsrSentenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AsrSentenceResponse{}
	_body, _err := client.AsrSentenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AsrTaskWithOptions(tmpReq *AsrTaskRequest, runtime *util.RuntimeOptions) (_result *AsrTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AsrTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Request)) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Request, tea.String("Request"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestShrink)) {
		body["Request"] = request.RequestShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AsrTask"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AsrTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AsrTask(request *AsrTaskRequest) (_result *AsrTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AsrTaskResponse{}
	_body, _err := client.AsrTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateRoomWithOptions(request *AssociateRoomRequest, runtime *util.RuntimeOptions) (_result *AssociateRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientBaseParam)) {
		query["ClientBaseParam"] = request.ClientBaseParam
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		query["RoomId"] = request.RoomId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateRoom"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateRoom(request *AssociateRoomRequest) (_result *AssociateRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateRoomResponse{}
	_body, _err := client.AssociateRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppWithOptions(request *CreateAppRequest, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PackageName)) {
		query["PackageName"] = request.PackageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApp"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDepartmentWithOptions(request *CreateDepartmentRequest, runtime *util.RuntimeOptions) (_result *CreateDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		query["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDepartment"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDepartmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDepartment(request *CreateDepartmentRequest) (_result *CreateDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDepartmentResponse{}
	_body, _err := client.CreateDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDetectProcessWithOptions(request *CreateDetectProcessRequest, runtime *util.RuntimeOptions) (_result *CreateDetectProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Draft)) {
		query["Draft"] = request.Draft
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDetectProcess"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDetectProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDetectProcess(request *CreateDetectProcessRequest) (_result *CreateDetectProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDetectProcessResponse{}
	_body, _err := client.CreateDetectProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRuleWithOptions(request *CreateRuleRequest, runtime *util.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRule"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRule(request *CreateRuleRequest) (_result *CreateRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRuleResponse{}
	_body, _err := client.CreateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSignatureWithOptions(request *CreateSignatureRequest, runtime *util.RuntimeOptions) (_result *CreateSignatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientBaseParam)) {
		query["ClientBaseParam"] = request.ClientBaseParam
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSignature"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSignature(request *CreateSignatureRequest) (_result *CreateSignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSignatureResponse{}
	_body, _err := client.CreateSignatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTaskGroupWithOptions(request *CreateTaskGroupRequest, runtime *util.RuntimeOptions) (_result *CreateTaskGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Day)) {
		query["Day"] = request.Day
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireAt)) {
		query["ExpireAt"] = request.ExpireAt
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RunnableTimeFrom)) {
		query["RunnableTimeFrom"] = request.RunnableTimeFrom
	}

	if !tea.BoolValue(util.IsUnset(request.RunnableTimeTo)) {
		query["RunnableTimeTo"] = request.RunnableTimeTo
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerPeriod)) {
		query["TriggerPeriod"] = request.TriggerPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.VideoInfo)) {
		query["VideoInfo"] = request.VideoInfo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTaskGroup"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTaskGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTaskGroup(request *CreateTaskGroupRequest) (_result *CreateTaskGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTaskGroupResponse{}
	_body, _err := client.CreateTaskGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTtsQuestionWithOptions(tmpReq *CreateTtsQuestionRequest, runtime *util.RuntimeOptions) (_result *CreateTtsQuestionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateTtsQuestionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Request)) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Request, tea.String("Request"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestShrink)) {
		body["Request"] = request.RequestShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTtsQuestion"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTtsQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTtsQuestion(request *CreateTtsQuestionRequest) (_result *CreateTtsQuestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTtsQuestionResponse{}
	_body, _err := client.CreateTtsQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTtsQuestionGroupWithOptions(tmpReq *CreateTtsQuestionGroupRequest, runtime *util.RuntimeOptions) (_result *CreateTtsQuestionGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateTtsQuestionGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Request)) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Request, tea.String("Request"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestShrink)) {
		body["Request"] = request.RequestShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTtsQuestionGroup"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTtsQuestionGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTtsQuestionGroup(request *CreateTtsQuestionGroupRequest) (_result *CreateTtsQuestionGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTtsQuestionGroupResponse{}
	_body, _err := client.CreateTtsQuestionGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserDepartmentsWithOptions(request *CreateUserDepartmentsRequest, runtime *util.RuntimeOptions) (_result *CreateUserDepartmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserDepartments"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserDepartmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserDepartments(request *CreateUserDepartmentsRequest) (_result *CreateUserDepartmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserDepartmentsResponse{}
	_body, _err := client.CreateUserDepartmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVideoMergeTaskWithOptions(tmpReq *CreateVideoMergeTaskRequest, runtime *util.RuntimeOptions) (_result *CreateVideoMergeTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateVideoMergeTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.VideoMergeRequest)) {
		request.VideoMergeRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoMergeRequest, tea.String("VideoMergeRequest"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VideoMergeRequestShrink)) {
		body["VideoMergeRequest"] = request.VideoMergeRequestShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVideoMergeTask"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVideoMergeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVideoMergeTask(request *CreateVideoMergeTaskRequest) (_result *CreateVideoMergeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVideoMergeTaskResponse{}
	_body, _err := client.CreateVideoMergeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWatermarkWithOptions(request *CreateWatermarkRequest, runtime *util.RuntimeOptions) (_result *CreateWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWatermark"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWatermarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWatermark(request *CreateWatermarkRequest) (_result *CreateWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWatermarkResponse{}
	_body, _err := client.CreateWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppWithOptions(request *DeleteAppRequest, runtime *util.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApp"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApp(request *DeleteAppRequest) (_result *DeleteAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDepartmentWithOptions(request *DeleteDepartmentRequest, runtime *util.RuntimeOptions) (_result *DeleteDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDepartment"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDepartmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDepartment(request *DeleteDepartmentRequest) (_result *DeleteDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDepartmentResponse{}
	_body, _err := client.DeleteDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDetectProcessWithOptions(request *DeleteDetectProcessRequest, runtime *util.RuntimeOptions) (_result *DeleteDetectProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDetectProcess"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDetectProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDetectProcess(request *DeleteDetectProcessRequest) (_result *DeleteDetectProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDetectProcessResponse{}
	_body, _err := client.DeleteDetectProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRule"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRule(request *DeleteRuleRequest) (_result *DeleteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DeleteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserDepartmentsWithOptions(request *DeleteUserDepartmentsRequest, runtime *util.RuntimeOptions) (_result *DeleteUserDepartmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["DepartmentId"] = request.DepartmentId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserDepartments"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserDepartmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserDepartments(request *DeleteUserDepartmentsRequest) (_result *DeleteUserDepartmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserDepartmentsResponse{}
	_body, _err := client.DeleteUserDepartmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWatermarkWithOptions(request *DeleteWatermarkRequest, runtime *util.RuntimeOptions) (_result *DeleteWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WatermarkId)) {
		query["WatermarkId"] = request.WatermarkId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWatermark"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWatermarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWatermark(request *DeleteWatermarkRequest) (_result *DeleteWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWatermarkResponse{}
	_body, _err := client.DeleteWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FaceCompareWithOptions(tmpReq *FaceCompareRequest, runtime *util.RuntimeOptions) (_result *FaceCompareResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &FaceCompareShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FaceRequest)) {
		request.FaceRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FaceRequest, tea.String("FaceRequest"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceRequestShrink)) {
		body["FaceRequest"] = request.FaceRequestShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceCompare"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceCompareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FaceCompare(request *FaceCompareRequest) (_result *FaceCompareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceCompareResponse{}
	_body, _err := client.FaceCompareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FaceLivenessWithOptions(tmpReq *FaceLivenessRequest, runtime *util.RuntimeOptions) (_result *FaceLivenessResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &FaceLivenessShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FaceRequest)) {
		request.FaceRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FaceRequest, tea.String("FaceRequest"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceRequestShrink)) {
		body["FaceRequest"] = request.FaceRequestShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceLiveness"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceLivenessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FaceLiveness(request *FaceLivenessRequest) (_result *FaceLivenessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceLivenessResponse{}
	_body, _err := client.FaceLivenessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FaceRecognizeWithOptions(tmpReq *FaceRecognizeRequest, runtime *util.RuntimeOptions) (_result *FaceRecognizeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &FaceRecognizeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FaceRequest)) {
		request.FaceRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FaceRequest, tea.String("FaceRequest"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceRequestShrink)) {
		body["FaceRequest"] = request.FaceRequestShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceRecognize"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceRecognizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FaceRecognize(request *FaceRecognizeRequest) (_result *FaceRecognizeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceRecognizeResponse{}
	_body, _err := client.FaceRecognizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppWithOptions(request *GetAppRequest, runtime *util.RuntimeOptions) (_result *GetAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientBaseParam)) {
		query["ClientBaseParam"] = request.ClientBaseParam
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PackageName)) {
		query["PackageName"] = request.PackageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApp"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApp(request *GetAppRequest) (_result *GetAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppResponse{}
	_body, _err := client.GetAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsrResultWithOptions(request *GetAsrResultRequest, runtime *util.RuntimeOptions) (_result *GetAsrResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AsrTaskId)) {
		query["AsrTaskId"] = request.AsrTaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsrResult"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsrResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsrResult(request *GetAsrResultRequest) (_result *GetAsrResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsrResultResponse{}
	_body, _err := client.GetAsrResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDepartmentWithOptions(request *GetDepartmentRequest, runtime *util.RuntimeOptions) (_result *GetDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientBaseParam)) {
		query["ClientBaseParam"] = request.ClientBaseParam
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDepartment"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDepartmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDepartment(request *GetDepartmentRequest) (_result *GetDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDepartmentResponse{}
	_body, _err := client.GetDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDetectProcessWithOptions(request *GetDetectProcessRequest, runtime *util.RuntimeOptions) (_result *GetDetectProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientBaseParam)) {
		query["ClientBaseParam"] = request.ClientBaseParam
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDetectProcess"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDetectProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDetectProcess(request *GetDetectProcessRequest) (_result *GetDetectProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetectProcessResponse{}
	_body, _err := client.GetDetectProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDetectProcessJsonFileWithOptions(request *GetDetectProcessJsonFileRequest, runtime *util.RuntimeOptions) (_result *GetDetectProcessJsonFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientBaseParam)) {
		query["ClientBaseParam"] = request.ClientBaseParam
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDetectProcessJsonFile"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDetectProcessJsonFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDetectProcessJsonFile(request *GetDetectProcessJsonFileRequest) (_result *GetDetectProcessJsonFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetectProcessJsonFileResponse{}
	_body, _err := client.GetDetectProcessJsonFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDetectionWithOptions(request *GetDetectionRequest, runtime *util.RuntimeOptions) (_result *GetDetectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientBaseParam)) {
		query["ClientBaseParam"] = request.ClientBaseParam
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDetection"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDetectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDetection(request *GetDetectionRequest) (_result *GetDetectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetectionResponse{}
	_body, _err := client.GetDetectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPreSignedUrlWithOptions(request *GetPreSignedUrlRequest, runtime *util.RuntimeOptions) (_result *GetPreSignedUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientBaseParam)) {
		query["ClientBaseParam"] = request.ClientBaseParam
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		body["Prefix"] = request.Prefix
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPreSignedUrl"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPreSignedUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPreSignedUrl(request *GetPreSignedUrlRequest) (_result *GetPreSignedUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPreSignedUrlResponse{}
	_body, _err := client.GetPreSignedUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRecordResultWithOptions(request *GetRecordResultRequest, runtime *util.RuntimeOptions) (_result *GetRecordResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientBaseParam)) {
		query["ClientBaseParam"] = request.ClientBaseParam
	}

	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRecordResult"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRecordResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRecordResult(request *GetRecordResultRequest) (_result *GetRecordResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRecordResultResponse{}
	_body, _err := client.GetRecordResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRecordsByFeeIdWithOptions(request *GetRecordsByFeeIdRequest, runtime *util.RuntimeOptions) (_result *GetRecordsByFeeIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeeId)) {
		body["FeeId"] = request.FeeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRecordsByFeeId"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRecordsByFeeIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRecordsByFeeId(request *GetRecordsByFeeIdRequest) (_result *GetRecordsByFeeIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRecordsByFeeIdResponse{}
	_body, _err := client.GetRecordsByFeeIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRecordsByOuterBusinessIdWithOptions(request *GetRecordsByOuterBusinessIdRequest, runtime *util.RuntimeOptions) (_result *GetRecordsByOuterBusinessIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OuterBusinessId)) {
		query["OuterBusinessId"] = request.OuterBusinessId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRecordsByOuterBusinessId"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRecordsByOuterBusinessIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRecordsByOuterBusinessId(request *GetRecordsByOuterBusinessIdRequest) (_result *GetRecordsByOuterBusinessIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRecordsByOuterBusinessIdResponse{}
	_body, _err := client.GetRecordsByOuterBusinessIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRuleWithOptions(request *GetRuleRequest, runtime *util.RuntimeOptions) (_result *GetRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientBaseParam)) {
		query["ClientBaseParam"] = request.ClientBaseParam
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRule"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRule(request *GetRuleRequest) (_result *GetRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRuleResponse{}
	_body, _err := client.GetRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStatisticsRecordsByFeeIdWithOptions(request *GetStatisticsRecordsByFeeIdRequest, runtime *util.RuntimeOptions) (_result *GetStatisticsRecordsByFeeIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeeId)) {
		body["FeeId"] = request.FeeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStatisticsRecordsByFeeId"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStatisticsRecordsByFeeIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStatisticsRecordsByFeeId(request *GetStatisticsRecordsByFeeIdRequest) (_result *GetStatisticsRecordsByFeeIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStatisticsRecordsByFeeIdResponse{}
	_body, _err := client.GetStatisticsRecordsByFeeIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskWithOptions(request *GetTaskRequest, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientBaseParam)) {
		query["ClientBaseParam"] = request.ClientBaseParam
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTask"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTask(request *GetTaskRequest) (_result *GetTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// **1**
//
// @param request - GetTaskGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskGroupResponse
func (client *Client) GetTaskGroupWithOptions(request *GetTaskGroupRequest, runtime *util.RuntimeOptions) (_result *GetTaskGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientBaseParam)) {
		query["ClientBaseParam"] = request.ClientBaseParam
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskGroup"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// **1**
//
// @param request - GetTaskGroupRequest
//
// @return GetTaskGroupResponse
func (client *Client) GetTaskGroup(request *GetTaskGroupRequest) (_result *GetTaskGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskGroupResponse{}
	_body, _err := client.GetTaskGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTtsQuestionByGroupIdWithOptions(request *GetTtsQuestionByGroupIdRequest, runtime *util.RuntimeOptions) (_result *GetTtsQuestionByGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTtsQuestionByGroupId"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTtsQuestionByGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTtsQuestionByGroupId(request *GetTtsQuestionByGroupIdRequest) (_result *GetTtsQuestionByGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTtsQuestionByGroupIdResponse{}
	_body, _err := client.GetTtsQuestionByGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientBaseParam)) {
		query["ClientBaseParam"] = request.ClientBaseParam
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVideoMergeTaskWithOptions(request *GetVideoMergeTaskRequest, runtime *util.RuntimeOptions) (_result *GetVideoMergeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVideoMergeTask"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVideoMergeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideoMergeTask(request *GetVideoMergeTaskRequest) (_result *GetVideoMergeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoMergeTaskResponse{}
	_body, _err := client.GetVideoMergeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWatermarkWithOptions(request *GetWatermarkRequest, runtime *util.RuntimeOptions) (_result *GetWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientBaseParam)) {
		query["ClientBaseParam"] = request.ClientBaseParam
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkId)) {
		query["WatermarkId"] = request.WatermarkId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWatermark"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWatermarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWatermark(request *GetWatermarkRequest) (_result *GetWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWatermarkResponse{}
	_body, _err := client.GetWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinRoomWithOptions(request *JoinRoomRequest, runtime *util.RuntimeOptions) (_result *JoinRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		query["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomToken)) {
		query["RoomToken"] = request.RoomToken
	}

	if !tea.BoolValue(util.IsUnset(request.StreamId)) {
		query["StreamId"] = request.StreamId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("JoinRoom"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &JoinRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinRoom(request *JoinRoomRequest) (_result *JoinRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinRoomResponse{}
	_body, _err := client.JoinRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LeaveRoomWithOptions(request *LeaveRoomRequest, runtime *util.RuntimeOptions) (_result *LeaveRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		query["RoomId"] = request.RoomId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LeaveRoom"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LeaveRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LeaveRoom(request *LeaveRoomRequest) (_result *LeaveRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LeaveRoomResponse{}
	_body, _err := client.LeaveRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppsWithOptions(request *ListAppsRequest, runtime *util.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApps"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApps(request *ListAppsRequest) (_result *ListAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDepartmentsWithOptions(request *ListDepartmentsRequest, runtime *util.RuntimeOptions) (_result *ListDepartmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDepartments"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDepartmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDepartments(request *ListDepartmentsRequest) (_result *ListDepartmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDepartmentsResponse{}
	_body, _err := client.ListDepartmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDetectProcessesWithOptions(request *ListDetectProcessesRequest, runtime *util.RuntimeOptions) (_result *ListDetectProcessesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PublishStatus)) {
		query["PublishStatus"] = request.PublishStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.SortKey)) {
		query["SortKey"] = request.SortKey
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDetectProcesses"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDetectProcessesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDetectProcesses(request *ListDetectProcessesRequest) (_result *ListDetectProcessesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDetectProcessesResponse{}
	_body, _err := client.ListDetectProcessesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDetectionsWithOptions(request *ListDetectionsRequest, runtime *util.RuntimeOptions) (_result *ListDetectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateDateFrom)) {
		query["CreateDateFrom"] = request.CreateDateFrom
	}

	if !tea.BoolValue(util.IsUnset(request.CreateDateTo)) {
		query["CreateDateTo"] = request.CreateDateTo
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RecordingType)) {
		query["RecordingType"] = request.RecordingType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDetections"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDetectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDetections(request *ListDetectionsRequest) (_result *ListDetectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDetectionsResponse{}
	_body, _err := client.ListDetectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFilesWithOptions(request *ListFilesRequest, runtime *util.RuntimeOptions) (_result *ListFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["Prefix"] = request.Prefix
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFiles"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFiles(request *ListFilesRequest) (_result *ListFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFilesResponse{}
	_body, _err := client.ListFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRecordResultsWithOptions(request *ListRecordResultsRequest, runtime *util.RuntimeOptions) (_result *ListRecordResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateDateFrom)) {
		query["CreateDateFrom"] = request.CreateDateFrom
	}

	if !tea.BoolValue(util.IsUnset(request.CreateDateTo)) {
		query["CreateDateTo"] = request.CreateDateTo
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.OuterBusinessId)) {
		query["OuterBusinessId"] = request.OuterBusinessId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRecordResults"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRecordResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRecordResults(request *ListRecordResultsRequest) (_result *ListRecordResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRecordResultsResponse{}
	_body, _err := client.ListRecordResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRulesWithOptions(request *ListRulesRequest, runtime *util.RuntimeOptions) (_result *ListRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRules"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRules(request *ListRulesRequest) (_result *ListRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRulesResponse{}
	_body, _err := client.ListRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTaskGroupsWithOptions(request *ListTaskGroupsRequest, runtime *util.RuntimeOptions) (_result *ListTaskGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTaskGroups"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTaskGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTaskGroups(request *ListTaskGroupsRequest) (_result *ListTaskGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTaskGroupsResponse{}
	_body, _err := client.ListTaskGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTaskItemsWithOptions(request *ListTaskItemsRequest, runtime *util.RuntimeOptions) (_result *ListTaskItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTaskItems"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTaskItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTaskItems(request *ListTaskItemsRequest) (_result *ListTaskItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTaskItemsResponse{}
	_body, _err := client.ListTaskItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTasksWithOptions(request *ListTasksRequest, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskGroupId)) {
		query["TaskGroupId"] = request.TaskGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTasks"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTasks(request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWatermarksWithOptions(request *ListWatermarksRequest, runtime *util.RuntimeOptions) (_result *ListWatermarksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWatermarks"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWatermarksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWatermarks(request *ListWatermarksRequest) (_result *ListWatermarksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWatermarksResponse{}
	_body, _err := client.ListWatermarksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenameDetectProcessWithOptions(request *RenameDetectProcessRequest, runtime *util.RuntimeOptions) (_result *RenameDetectProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameDetectProcess"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameDetectProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenameDetectProcess(request *RenameDetectProcessRequest) (_result *RenameDetectProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameDetectProcessResponse{}
	_body, _err := client.RenameDetectProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TtsCommonWithOptions(tmpReq *TtsCommonRequest, runtime *util.RuntimeOptions) (_result *TtsCommonResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &TtsCommonShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TtsRequest)) {
		request.TtsRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TtsRequest, tea.String("TtsRequest"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TtsRequestShrink)) {
		body["TtsRequest"] = request.TtsRequestShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TtsCommon"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TtsCommonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TtsCommon(request *TtsCommonRequest) (_result *TtsCommonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TtsCommonResponse{}
	_body, _err := client.TtsCommonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TtsTaskWithOptions(tmpReq *TtsTaskRequest, runtime *util.RuntimeOptions) (_result *TtsTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &TtsTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Request)) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Request, tea.String("Request"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestShrink)) {
		body["Request"] = request.RequestShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TtsTask"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TtsTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TtsTask(request *TtsTaskRequest) (_result *TtsTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TtsTaskResponse{}
	_body, _err := client.TtsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppWithOptions(request *UpdateAppRequest, runtime *util.RuntimeOptions) (_result *UpdateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.Disabled)) {
		query["Disabled"] = request.Disabled
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PackageName)) {
		query["PackageName"] = request.PackageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApp"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApp(request *UpdateAppRequest) (_result *UpdateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppResponse{}
	_body, _err := client.UpdateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDepartmentWithOptions(request *UpdateDepartmentRequest, runtime *util.RuntimeOptions) (_result *UpdateDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		body["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDepartment"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDepartmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDepartment(request *UpdateDepartmentRequest) (_result *UpdateDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDepartmentResponse{}
	_body, _err := client.UpdateDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// ********
//
// @param request - UpdateDetectProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDetectProcessResponse
func (client *Client) UpdateDetectProcessWithOptions(request *UpdateDetectProcessRequest, runtime *util.RuntimeOptions) (_result *UpdateDetectProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Draft)) {
		query["Draft"] = request.Draft
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDetectProcess"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDetectProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// ********
//
// @param request - UpdateDetectProcessRequest
//
// @return UpdateDetectProcessResponse
func (client *Client) UpdateDetectProcess(request *UpdateDetectProcessRequest) (_result *UpdateDetectProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDetectProcessResponse{}
	_body, _err := client.UpdateDetectProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRuleWithOptions(request *UpdateRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRule"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRule(request *UpdateRuleRequest) (_result *UpdateRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRuleResponse{}
	_body, _err := client.UpdateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["Role"] = request.Role
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUser"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWatermarkWithOptions(request *UpdateWatermarkRequest, runtime *util.RuntimeOptions) (_result *UpdateWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkId)) {
		query["WatermarkId"] = request.WatermarkId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWatermark"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWatermarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWatermark(request *UpdateWatermarkRequest) (_result *UpdateWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWatermarkResponse{}
	_body, _err := client.UpdateWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadReportWithOptions(request *UploadReportRequest, runtime *util.RuntimeOptions) (_result *UploadReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientBaseParam)) {
		query["ClientBaseParam"] = request.ClientBaseParam
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.DetectProcessId)) {
		query["DetectProcessId"] = request.DetectProcessId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.FeeId)) {
		query["FeeId"] = request.FeeId
	}

	if !tea.BoolValue(util.IsUnset(request.MetaUrl)) {
		query["MetaUrl"] = request.MetaUrl
	}

	if !tea.BoolValue(util.IsUnset(request.OuterBusinessId)) {
		query["OuterBusinessId"] = request.OuterBusinessId
	}

	if !tea.BoolValue(util.IsUnset(request.RecordAt)) {
		query["RecordAt"] = request.RecordAt
	}

	if !tea.BoolValue(util.IsUnset(request.ResultUrl)) {
		query["ResultUrl"] = request.ResultUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["Role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		query["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.RtcRecordId)) {
		query["RtcRecordId"] = request.RtcRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.VideoType)) {
		query["VideoType"] = request.VideoType
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		query["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadReport"),
		Version:     tea.String("2020-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadReport(request *UploadReportRequest) (_result *UploadReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadReportResponse{}
	_body, _err := client.UploadReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
