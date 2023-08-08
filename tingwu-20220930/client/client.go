// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateFileTransRequest struct {
	AbilityParams        map[string]interface{} `json:"AbilityParams,omitempty" xml:"AbilityParams,omitempty"`
	AppKey               *string                `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AsrParams            map[string]interface{} `json:"AsrParams,omitempty" xml:"AsrParams,omitempty"`
	AudioLanguage        *string                `json:"AudioLanguage,omitempty" xml:"AudioLanguage,omitempty"`
	AudioOssBucket       *string                `json:"AudioOssBucket,omitempty" xml:"AudioOssBucket,omitempty"`
	AudioOssPath         *string                `json:"AudioOssPath,omitempty" xml:"AudioOssPath,omitempty"`
	AudioOutputEnabled   *bool                  `json:"AudioOutputEnabled,omitempty" xml:"AudioOutputEnabled,omitempty"`
	AudioOutputOssBucket *string                `json:"AudioOutputOssBucket,omitempty" xml:"AudioOutputOssBucket,omitempty"`
	AudioOutputOssPath   *string                `json:"AudioOutputOssPath,omitempty" xml:"AudioOutputOssPath,omitempty"`
	AudioRoleNum         *string                `json:"AudioRoleNum,omitempty" xml:"AudioRoleNum,omitempty"`
	AudioSegmentsEnabled *bool                  `json:"AudioSegmentsEnabled,omitempty" xml:"AudioSegmentsEnabled,omitempty"`
	LabParams            map[string]interface{} `json:"LabParams,omitempty" xml:"LabParams,omitempty"`
	Tags                 map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TransKey             *string                `json:"TransKey,omitempty" xml:"TransKey,omitempty"`
	TransResultOssBucket *string                `json:"TransResultOssBucket,omitempty" xml:"TransResultOssBucket,omitempty"`
	TransResultOssPath   *string                `json:"TransResultOssPath,omitempty" xml:"TransResultOssPath,omitempty"`
	VideoOutputEnabled   *bool                  `json:"VideoOutputEnabled,omitempty" xml:"VideoOutputEnabled,omitempty"`
	VideoOutputOssBucket *string                `json:"VideoOutputOssBucket,omitempty" xml:"VideoOutputOssBucket,omitempty"`
	VideoOutputOssPath   *string                `json:"VideoOutputOssPath,omitempty" xml:"VideoOutputOssPath,omitempty"`
}

func (s CreateFileTransRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFileTransRequest) GoString() string {
	return s.String()
}

func (s *CreateFileTransRequest) SetAbilityParams(v map[string]interface{}) *CreateFileTransRequest {
	s.AbilityParams = v
	return s
}

func (s *CreateFileTransRequest) SetAppKey(v string) *CreateFileTransRequest {
	s.AppKey = &v
	return s
}

func (s *CreateFileTransRequest) SetAsrParams(v map[string]interface{}) *CreateFileTransRequest {
	s.AsrParams = v
	return s
}

func (s *CreateFileTransRequest) SetAudioLanguage(v string) *CreateFileTransRequest {
	s.AudioLanguage = &v
	return s
}

func (s *CreateFileTransRequest) SetAudioOssBucket(v string) *CreateFileTransRequest {
	s.AudioOssBucket = &v
	return s
}

func (s *CreateFileTransRequest) SetAudioOssPath(v string) *CreateFileTransRequest {
	s.AudioOssPath = &v
	return s
}

func (s *CreateFileTransRequest) SetAudioOutputEnabled(v bool) *CreateFileTransRequest {
	s.AudioOutputEnabled = &v
	return s
}

func (s *CreateFileTransRequest) SetAudioOutputOssBucket(v string) *CreateFileTransRequest {
	s.AudioOutputOssBucket = &v
	return s
}

func (s *CreateFileTransRequest) SetAudioOutputOssPath(v string) *CreateFileTransRequest {
	s.AudioOutputOssPath = &v
	return s
}

func (s *CreateFileTransRequest) SetAudioRoleNum(v string) *CreateFileTransRequest {
	s.AudioRoleNum = &v
	return s
}

func (s *CreateFileTransRequest) SetAudioSegmentsEnabled(v bool) *CreateFileTransRequest {
	s.AudioSegmentsEnabled = &v
	return s
}

func (s *CreateFileTransRequest) SetLabParams(v map[string]interface{}) *CreateFileTransRequest {
	s.LabParams = v
	return s
}

func (s *CreateFileTransRequest) SetTags(v map[string]interface{}) *CreateFileTransRequest {
	s.Tags = v
	return s
}

func (s *CreateFileTransRequest) SetTransKey(v string) *CreateFileTransRequest {
	s.TransKey = &v
	return s
}

func (s *CreateFileTransRequest) SetTransResultOssBucket(v string) *CreateFileTransRequest {
	s.TransResultOssBucket = &v
	return s
}

func (s *CreateFileTransRequest) SetTransResultOssPath(v string) *CreateFileTransRequest {
	s.TransResultOssPath = &v
	return s
}

func (s *CreateFileTransRequest) SetVideoOutputEnabled(v bool) *CreateFileTransRequest {
	s.VideoOutputEnabled = &v
	return s
}

func (s *CreateFileTransRequest) SetVideoOutputOssBucket(v string) *CreateFileTransRequest {
	s.VideoOutputOssBucket = &v
	return s
}

func (s *CreateFileTransRequest) SetVideoOutputOssPath(v string) *CreateFileTransRequest {
	s.VideoOutputOssPath = &v
	return s
}

type CreateFileTransResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateFileTransResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFileTransResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFileTransResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileTransResponseBody) SetCode(v string) *CreateFileTransResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFileTransResponseBody) SetData(v *CreateFileTransResponseBodyData) *CreateFileTransResponseBody {
	s.Data = v
	return s
}

func (s *CreateFileTransResponseBody) SetMessage(v string) *CreateFileTransResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFileTransResponseBody) SetRequestId(v string) *CreateFileTransResponseBody {
	s.RequestId = &v
	return s
}

type CreateFileTransResponseBodyData struct {
	TransId  *string `json:"TransId,omitempty" xml:"TransId,omitempty"`
	TransKey *string `json:"TransKey,omitempty" xml:"TransKey,omitempty"`
}

func (s CreateFileTransResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateFileTransResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateFileTransResponseBodyData) SetTransId(v string) *CreateFileTransResponseBodyData {
	s.TransId = &v
	return s
}

func (s *CreateFileTransResponseBodyData) SetTransKey(v string) *CreateFileTransResponseBodyData {
	s.TransKey = &v
	return s
}

type CreateFileTransResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFileTransResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFileTransResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFileTransResponse) GoString() string {
	return s.String()
}

func (s *CreateFileTransResponse) SetHeaders(v map[string]*string) *CreateFileTransResponse {
	s.Headers = v
	return s
}

func (s *CreateFileTransResponse) SetStatusCode(v int32) *CreateFileTransResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileTransResponse) SetBody(v *CreateFileTransResponseBody) *CreateFileTransResponse {
	s.Body = v
	return s
}

type CreateMeetingTransRequest struct {
	AbilityParams                    map[string]interface{} `json:"AbilityParams,omitempty" xml:"AbilityParams,omitempty"`
	AppKey                           *string                `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AsrParams                        map[string]interface{} `json:"AsrParams,omitempty" xml:"AsrParams,omitempty"`
	AudioBitRate                     *int32                 `json:"AudioBitRate,omitempty" xml:"AudioBitRate,omitempty"`
	AudioFormat                      *string                `json:"AudioFormat,omitempty" xml:"AudioFormat,omitempty"`
	AudioLanguage                    *string                `json:"AudioLanguage,omitempty" xml:"AudioLanguage,omitempty"`
	AudioOutputEnabled               *bool                  `json:"AudioOutputEnabled,omitempty" xml:"AudioOutputEnabled,omitempty"`
	AudioOutputOssBucket             *string                `json:"AudioOutputOssBucket,omitempty" xml:"AudioOutputOssBucket,omitempty"`
	AudioOutputOssPath               *string                `json:"AudioOutputOssPath,omitempty" xml:"AudioOutputOssPath,omitempty"`
	AudioPackage                     *string                `json:"AudioPackage,omitempty" xml:"AudioPackage,omitempty"`
	AudioSampleRate                  *int32                 `json:"AudioSampleRate,omitempty" xml:"AudioSampleRate,omitempty"`
	AudioSegmentsEnabled             *bool                  `json:"AudioSegmentsEnabled,omitempty" xml:"AudioSegmentsEnabled,omitempty"`
	DocResultEnabled                 *bool                  `json:"DocResultEnabled,omitempty" xml:"DocResultEnabled,omitempty"`
	LabParams                        map[string]interface{} `json:"LabParams,omitempty" xml:"LabParams,omitempty"`
	MeetingKey                       *string                `json:"MeetingKey,omitempty" xml:"MeetingKey,omitempty"`
	MeetingResultEnabled             *bool                  `json:"MeetingResultEnabled,omitempty" xml:"MeetingResultEnabled,omitempty"`
	MeetingResultOssBucket           *string                `json:"MeetingResultOssBucket,omitempty" xml:"MeetingResultOssBucket,omitempty"`
	MeetingResultOssPath             *string                `json:"MeetingResultOssPath,omitempty" xml:"MeetingResultOssPath,omitempty"`
	RealtimeActiveResultLevel        *int32                 `json:"RealtimeActiveResultLevel,omitempty" xml:"RealtimeActiveResultLevel,omitempty"`
	RealtimeResultEnabled            *bool                  `json:"RealtimeResultEnabled,omitempty" xml:"RealtimeResultEnabled,omitempty"`
	RealtimeResultLevel              *int32                 `json:"RealtimeResultLevel,omitempty" xml:"RealtimeResultLevel,omitempty"`
	RealtimeResultMeetingInfoEnabled *bool                  `json:"RealtimeResultMeetingInfoEnabled,omitempty" xml:"RealtimeResultMeetingInfoEnabled,omitempty"`
	RealtimeResultWordsEnabled       *bool                  `json:"RealtimeResultWordsEnabled,omitempty" xml:"RealtimeResultWordsEnabled,omitempty"`
	Tags                             map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TranslateActiveResultLevel       *int32                 `json:"TranslateActiveResultLevel,omitempty" xml:"TranslateActiveResultLevel,omitempty"`
	TranslateLanguages               *string                `json:"TranslateLanguages,omitempty" xml:"TranslateLanguages,omitempty"`
	TranslateResultEnabled           *bool                  `json:"TranslateResultEnabled,omitempty" xml:"TranslateResultEnabled,omitempty"`
	TranslateResultLevel             *int32                 `json:"TranslateResultLevel,omitempty" xml:"TranslateResultLevel,omitempty"`
}

func (s CreateMeetingTransRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingTransRequest) GoString() string {
	return s.String()
}

func (s *CreateMeetingTransRequest) SetAbilityParams(v map[string]interface{}) *CreateMeetingTransRequest {
	s.AbilityParams = v
	return s
}

func (s *CreateMeetingTransRequest) SetAppKey(v string) *CreateMeetingTransRequest {
	s.AppKey = &v
	return s
}

func (s *CreateMeetingTransRequest) SetAsrParams(v map[string]interface{}) *CreateMeetingTransRequest {
	s.AsrParams = v
	return s
}

func (s *CreateMeetingTransRequest) SetAudioBitRate(v int32) *CreateMeetingTransRequest {
	s.AudioBitRate = &v
	return s
}

func (s *CreateMeetingTransRequest) SetAudioFormat(v string) *CreateMeetingTransRequest {
	s.AudioFormat = &v
	return s
}

func (s *CreateMeetingTransRequest) SetAudioLanguage(v string) *CreateMeetingTransRequest {
	s.AudioLanguage = &v
	return s
}

func (s *CreateMeetingTransRequest) SetAudioOutputEnabled(v bool) *CreateMeetingTransRequest {
	s.AudioOutputEnabled = &v
	return s
}

func (s *CreateMeetingTransRequest) SetAudioOutputOssBucket(v string) *CreateMeetingTransRequest {
	s.AudioOutputOssBucket = &v
	return s
}

func (s *CreateMeetingTransRequest) SetAudioOutputOssPath(v string) *CreateMeetingTransRequest {
	s.AudioOutputOssPath = &v
	return s
}

func (s *CreateMeetingTransRequest) SetAudioPackage(v string) *CreateMeetingTransRequest {
	s.AudioPackage = &v
	return s
}

func (s *CreateMeetingTransRequest) SetAudioSampleRate(v int32) *CreateMeetingTransRequest {
	s.AudioSampleRate = &v
	return s
}

func (s *CreateMeetingTransRequest) SetAudioSegmentsEnabled(v bool) *CreateMeetingTransRequest {
	s.AudioSegmentsEnabled = &v
	return s
}

func (s *CreateMeetingTransRequest) SetDocResultEnabled(v bool) *CreateMeetingTransRequest {
	s.DocResultEnabled = &v
	return s
}

func (s *CreateMeetingTransRequest) SetLabParams(v map[string]interface{}) *CreateMeetingTransRequest {
	s.LabParams = v
	return s
}

func (s *CreateMeetingTransRequest) SetMeetingKey(v string) *CreateMeetingTransRequest {
	s.MeetingKey = &v
	return s
}

func (s *CreateMeetingTransRequest) SetMeetingResultEnabled(v bool) *CreateMeetingTransRequest {
	s.MeetingResultEnabled = &v
	return s
}

func (s *CreateMeetingTransRequest) SetMeetingResultOssBucket(v string) *CreateMeetingTransRequest {
	s.MeetingResultOssBucket = &v
	return s
}

func (s *CreateMeetingTransRequest) SetMeetingResultOssPath(v string) *CreateMeetingTransRequest {
	s.MeetingResultOssPath = &v
	return s
}

func (s *CreateMeetingTransRequest) SetRealtimeActiveResultLevel(v int32) *CreateMeetingTransRequest {
	s.RealtimeActiveResultLevel = &v
	return s
}

func (s *CreateMeetingTransRequest) SetRealtimeResultEnabled(v bool) *CreateMeetingTransRequest {
	s.RealtimeResultEnabled = &v
	return s
}

func (s *CreateMeetingTransRequest) SetRealtimeResultLevel(v int32) *CreateMeetingTransRequest {
	s.RealtimeResultLevel = &v
	return s
}

func (s *CreateMeetingTransRequest) SetRealtimeResultMeetingInfoEnabled(v bool) *CreateMeetingTransRequest {
	s.RealtimeResultMeetingInfoEnabled = &v
	return s
}

func (s *CreateMeetingTransRequest) SetRealtimeResultWordsEnabled(v bool) *CreateMeetingTransRequest {
	s.RealtimeResultWordsEnabled = &v
	return s
}

func (s *CreateMeetingTransRequest) SetTags(v map[string]interface{}) *CreateMeetingTransRequest {
	s.Tags = v
	return s
}

func (s *CreateMeetingTransRequest) SetTranslateActiveResultLevel(v int32) *CreateMeetingTransRequest {
	s.TranslateActiveResultLevel = &v
	return s
}

func (s *CreateMeetingTransRequest) SetTranslateLanguages(v string) *CreateMeetingTransRequest {
	s.TranslateLanguages = &v
	return s
}

func (s *CreateMeetingTransRequest) SetTranslateResultEnabled(v bool) *CreateMeetingTransRequest {
	s.TranslateResultEnabled = &v
	return s
}

func (s *CreateMeetingTransRequest) SetTranslateResultLevel(v int32) *CreateMeetingTransRequest {
	s.TranslateResultLevel = &v
	return s
}

type CreateMeetingTransResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateMeetingTransResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMeetingTransResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingTransResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMeetingTransResponseBody) SetCode(v string) *CreateMeetingTransResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMeetingTransResponseBody) SetData(v *CreateMeetingTransResponseBodyData) *CreateMeetingTransResponseBody {
	s.Data = v
	return s
}

func (s *CreateMeetingTransResponseBody) SetMessage(v string) *CreateMeetingTransResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMeetingTransResponseBody) SetRequestId(v string) *CreateMeetingTransResponseBody {
	s.RequestId = &v
	return s
}

type CreateMeetingTransResponseBodyData struct {
	MeetingId      *string `json:"MeetingId,omitempty" xml:"MeetingId,omitempty"`
	MeetingJoinUrl *string `json:"MeetingJoinUrl,omitempty" xml:"MeetingJoinUrl,omitempty"`
	MeetingKey     *string `json:"MeetingKey,omitempty" xml:"MeetingKey,omitempty"`
}

func (s CreateMeetingTransResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingTransResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMeetingTransResponseBodyData) SetMeetingId(v string) *CreateMeetingTransResponseBodyData {
	s.MeetingId = &v
	return s
}

func (s *CreateMeetingTransResponseBodyData) SetMeetingJoinUrl(v string) *CreateMeetingTransResponseBodyData {
	s.MeetingJoinUrl = &v
	return s
}

func (s *CreateMeetingTransResponseBodyData) SetMeetingKey(v string) *CreateMeetingTransResponseBodyData {
	s.MeetingKey = &v
	return s
}

type CreateMeetingTransResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMeetingTransResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMeetingTransResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingTransResponse) GoString() string {
	return s.String()
}

func (s *CreateMeetingTransResponse) SetHeaders(v map[string]*string) *CreateMeetingTransResponse {
	s.Headers = v
	return s
}

func (s *CreateMeetingTransResponse) SetStatusCode(v int32) *CreateMeetingTransResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMeetingTransResponse) SetBody(v *CreateMeetingTransResponseBody) *CreateMeetingTransResponse {
	s.Body = v
	return s
}

type GetFileTransResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetFileTransResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileTransResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileTransResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileTransResponseBody) SetCode(v string) *GetFileTransResponseBody {
	s.Code = &v
	return s
}

func (s *GetFileTransResponseBody) SetData(v *GetFileTransResponseBodyData) *GetFileTransResponseBody {
	s.Data = v
	return s
}

func (s *GetFileTransResponseBody) SetMessage(v string) *GetFileTransResponseBody {
	s.Message = &v
	return s
}

func (s *GetFileTransResponseBody) SetRequestId(v string) *GetFileTransResponseBody {
	s.RequestId = &v
	return s
}

type GetFileTransResponseBodyData struct {
	TransId     *string `json:"TransId,omitempty" xml:"TransId,omitempty"`
	TransKey    *string `json:"TransKey,omitempty" xml:"TransKey,omitempty"`
	TransStatus *string `json:"TransStatus,omitempty" xml:"TransStatus,omitempty"`
}

func (s GetFileTransResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFileTransResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileTransResponseBodyData) SetTransId(v string) *GetFileTransResponseBodyData {
	s.TransId = &v
	return s
}

func (s *GetFileTransResponseBodyData) SetTransKey(v string) *GetFileTransResponseBodyData {
	s.TransKey = &v
	return s
}

func (s *GetFileTransResponseBodyData) SetTransStatus(v string) *GetFileTransResponseBodyData {
	s.TransStatus = &v
	return s
}

type GetFileTransResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFileTransResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileTransResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileTransResponse) GoString() string {
	return s.String()
}

func (s *GetFileTransResponse) SetHeaders(v map[string]*string) *GetFileTransResponse {
	s.Headers = v
	return s
}

func (s *GetFileTransResponse) SetStatusCode(v int32) *GetFileTransResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileTransResponse) SetBody(v *GetFileTransResponseBody) *GetFileTransResponse {
	s.Body = v
	return s
}

type GetMeetingTransResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetMeetingTransResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMeetingTransResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingTransResponseBody) GoString() string {
	return s.String()
}

func (s *GetMeetingTransResponseBody) SetCode(v string) *GetMeetingTransResponseBody {
	s.Code = &v
	return s
}

func (s *GetMeetingTransResponseBody) SetData(v *GetMeetingTransResponseBodyData) *GetMeetingTransResponseBody {
	s.Data = v
	return s
}

func (s *GetMeetingTransResponseBody) SetMessage(v string) *GetMeetingTransResponseBody {
	s.Message = &v
	return s
}

func (s *GetMeetingTransResponseBody) SetRequestId(v string) *GetMeetingTransResponseBody {
	s.RequestId = &v
	return s
}

type GetMeetingTransResponseBodyData struct {
	MeetingId     *string `json:"MeetingId,omitempty" xml:"MeetingId,omitempty"`
	MeetingKey    *string `json:"MeetingKey,omitempty" xml:"MeetingKey,omitempty"`
	MeetingStatus *string `json:"MeetingStatus,omitempty" xml:"MeetingStatus,omitempty"`
}

func (s GetMeetingTransResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingTransResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMeetingTransResponseBodyData) SetMeetingId(v string) *GetMeetingTransResponseBodyData {
	s.MeetingId = &v
	return s
}

func (s *GetMeetingTransResponseBodyData) SetMeetingKey(v string) *GetMeetingTransResponseBodyData {
	s.MeetingKey = &v
	return s
}

func (s *GetMeetingTransResponseBodyData) SetMeetingStatus(v string) *GetMeetingTransResponseBodyData {
	s.MeetingStatus = &v
	return s
}

type GetMeetingTransResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMeetingTransResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMeetingTransResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMeetingTransResponse) GoString() string {
	return s.String()
}

func (s *GetMeetingTransResponse) SetHeaders(v map[string]*string) *GetMeetingTransResponse {
	s.Headers = v
	return s
}

func (s *GetMeetingTransResponse) SetStatusCode(v int32) *GetMeetingTransResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMeetingTransResponse) SetBody(v *GetMeetingTransResponseBody) *GetMeetingTransResponse {
	s.Body = v
	return s
}

type StopMeetingTransRequest struct {
	MeetingRoleNum      *int32 `json:"MeetingRoleNum,omitempty" xml:"MeetingRoleNum,omitempty"`
	OnlyRoleSplitResult *bool  `json:"OnlyRoleSplitResult,omitempty" xml:"OnlyRoleSplitResult,omitempty"`
}

func (s StopMeetingTransRequest) String() string {
	return tea.Prettify(s)
}

func (s StopMeetingTransRequest) GoString() string {
	return s.String()
}

func (s *StopMeetingTransRequest) SetMeetingRoleNum(v int32) *StopMeetingTransRequest {
	s.MeetingRoleNum = &v
	return s
}

func (s *StopMeetingTransRequest) SetOnlyRoleSplitResult(v bool) *StopMeetingTransRequest {
	s.OnlyRoleSplitResult = &v
	return s
}

type StopMeetingTransResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *StopMeetingTransResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopMeetingTransResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopMeetingTransResponseBody) GoString() string {
	return s.String()
}

func (s *StopMeetingTransResponseBody) SetCode(v string) *StopMeetingTransResponseBody {
	s.Code = &v
	return s
}

func (s *StopMeetingTransResponseBody) SetData(v *StopMeetingTransResponseBodyData) *StopMeetingTransResponseBody {
	s.Data = v
	return s
}

func (s *StopMeetingTransResponseBody) SetMessage(v string) *StopMeetingTransResponseBody {
	s.Message = &v
	return s
}

func (s *StopMeetingTransResponseBody) SetRequestId(v string) *StopMeetingTransResponseBody {
	s.RequestId = &v
	return s
}

type StopMeetingTransResponseBodyData struct {
	MeetingId     *string `json:"MeetingId,omitempty" xml:"MeetingId,omitempty"`
	MeetingKey    *string `json:"MeetingKey,omitempty" xml:"MeetingKey,omitempty"`
	MeetingStatus *string `json:"MeetingStatus,omitempty" xml:"MeetingStatus,omitempty"`
}

func (s StopMeetingTransResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StopMeetingTransResponseBodyData) GoString() string {
	return s.String()
}

func (s *StopMeetingTransResponseBodyData) SetMeetingId(v string) *StopMeetingTransResponseBodyData {
	s.MeetingId = &v
	return s
}

func (s *StopMeetingTransResponseBodyData) SetMeetingKey(v string) *StopMeetingTransResponseBodyData {
	s.MeetingKey = &v
	return s
}

func (s *StopMeetingTransResponseBodyData) SetMeetingStatus(v string) *StopMeetingTransResponseBodyData {
	s.MeetingStatus = &v
	return s
}

type StopMeetingTransResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopMeetingTransResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopMeetingTransResponse) String() string {
	return tea.Prettify(s)
}

func (s StopMeetingTransResponse) GoString() string {
	return s.String()
}

func (s *StopMeetingTransResponse) SetHeaders(v map[string]*string) *StopMeetingTransResponse {
	s.Headers = v
	return s
}

func (s *StopMeetingTransResponse) SetStatusCode(v int32) *StopMeetingTransResponse {
	s.StatusCode = &v
	return s
}

func (s *StopMeetingTransResponse) SetBody(v *StopMeetingTransResponseBody) *StopMeetingTransResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("tingwu"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateFileTransWithOptions(request *CreateFileTransRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFileTransResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AbilityParams)) {
		body["AbilityParams"] = request.AbilityParams
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.AsrParams)) {
		body["AsrParams"] = request.AsrParams
	}

	if !tea.BoolValue(util.IsUnset(request.AudioLanguage)) {
		body["AudioLanguage"] = request.AudioLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.AudioOssBucket)) {
		body["AudioOssBucket"] = request.AudioOssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.AudioOssPath)) {
		body["AudioOssPath"] = request.AudioOssPath
	}

	if !tea.BoolValue(util.IsUnset(request.AudioOutputEnabled)) {
		body["AudioOutputEnabled"] = request.AudioOutputEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.AudioOutputOssBucket)) {
		body["AudioOutputOssBucket"] = request.AudioOutputOssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.AudioOutputOssPath)) {
		body["AudioOutputOssPath"] = request.AudioOutputOssPath
	}

	if !tea.BoolValue(util.IsUnset(request.AudioRoleNum)) {
		body["AudioRoleNum"] = request.AudioRoleNum
	}

	if !tea.BoolValue(util.IsUnset(request.AudioSegmentsEnabled)) {
		body["AudioSegmentsEnabled"] = request.AudioSegmentsEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.LabParams)) {
		body["LabParams"] = request.LabParams
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.TransKey)) {
		body["TransKey"] = request.TransKey
	}

	if !tea.BoolValue(util.IsUnset(request.TransResultOssBucket)) {
		body["TransResultOssBucket"] = request.TransResultOssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.TransResultOssPath)) {
		body["TransResultOssPath"] = request.TransResultOssPath
	}

	if !tea.BoolValue(util.IsUnset(request.VideoOutputEnabled)) {
		body["VideoOutputEnabled"] = request.VideoOutputEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.VideoOutputOssBucket)) {
		body["VideoOutputOssBucket"] = request.VideoOutputOssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.VideoOutputOssPath)) {
		body["VideoOutputOssPath"] = request.VideoOutputOssPath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFileTrans"),
		Version:     tea.String("2022-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/file-trans"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFileTransResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFileTrans(request *CreateFileTransRequest) (_result *CreateFileTransResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFileTransResponse{}
	_body, _err := client.CreateFileTransWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMeetingTransWithOptions(request *CreateMeetingTransRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMeetingTransResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AbilityParams)) {
		body["AbilityParams"] = request.AbilityParams
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.AsrParams)) {
		body["AsrParams"] = request.AsrParams
	}

	if !tea.BoolValue(util.IsUnset(request.AudioBitRate)) {
		body["AudioBitRate"] = request.AudioBitRate
	}

	if !tea.BoolValue(util.IsUnset(request.AudioFormat)) {
		body["AudioFormat"] = request.AudioFormat
	}

	if !tea.BoolValue(util.IsUnset(request.AudioLanguage)) {
		body["AudioLanguage"] = request.AudioLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.AudioOutputEnabled)) {
		body["AudioOutputEnabled"] = request.AudioOutputEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.AudioOutputOssBucket)) {
		body["AudioOutputOssBucket"] = request.AudioOutputOssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.AudioOutputOssPath)) {
		body["AudioOutputOssPath"] = request.AudioOutputOssPath
	}

	if !tea.BoolValue(util.IsUnset(request.AudioPackage)) {
		body["AudioPackage"] = request.AudioPackage
	}

	if !tea.BoolValue(util.IsUnset(request.AudioSampleRate)) {
		body["AudioSampleRate"] = request.AudioSampleRate
	}

	if !tea.BoolValue(util.IsUnset(request.AudioSegmentsEnabled)) {
		body["AudioSegmentsEnabled"] = request.AudioSegmentsEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DocResultEnabled)) {
		body["DocResultEnabled"] = request.DocResultEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.LabParams)) {
		body["LabParams"] = request.LabParams
	}

	if !tea.BoolValue(util.IsUnset(request.MeetingKey)) {
		body["MeetingKey"] = request.MeetingKey
	}

	if !tea.BoolValue(util.IsUnset(request.MeetingResultEnabled)) {
		body["MeetingResultEnabled"] = request.MeetingResultEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.MeetingResultOssBucket)) {
		body["MeetingResultOssBucket"] = request.MeetingResultOssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.MeetingResultOssPath)) {
		body["MeetingResultOssPath"] = request.MeetingResultOssPath
	}

	if !tea.BoolValue(util.IsUnset(request.RealtimeActiveResultLevel)) {
		body["RealtimeActiveResultLevel"] = request.RealtimeActiveResultLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RealtimeResultEnabled)) {
		body["RealtimeResultEnabled"] = request.RealtimeResultEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.RealtimeResultLevel)) {
		body["RealtimeResultLevel"] = request.RealtimeResultLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RealtimeResultMeetingInfoEnabled)) {
		body["RealtimeResultMeetingInfoEnabled"] = request.RealtimeResultMeetingInfoEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.RealtimeResultWordsEnabled)) {
		body["RealtimeResultWordsEnabled"] = request.RealtimeResultWordsEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.TranslateActiveResultLevel)) {
		body["TranslateActiveResultLevel"] = request.TranslateActiveResultLevel
	}

	if !tea.BoolValue(util.IsUnset(request.TranslateLanguages)) {
		body["TranslateLanguages"] = request.TranslateLanguages
	}

	if !tea.BoolValue(util.IsUnset(request.TranslateResultEnabled)) {
		body["TranslateResultEnabled"] = request.TranslateResultEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.TranslateResultLevel)) {
		body["TranslateResultLevel"] = request.TranslateResultLevel
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMeetingTrans"),
		Version:     tea.String("2022-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/meeting-trans"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMeetingTransResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMeetingTrans(request *CreateMeetingTransRequest) (_result *CreateMeetingTransResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMeetingTransResponse{}
	_body, _err := client.CreateMeetingTransWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileTransWithOptions(TransId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFileTransResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetFileTrans"),
		Version:     tea.String("2022-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/file-trans/" + tea.StringValue(openapiutil.GetEncodeParam(TransId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileTransResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileTrans(TransId *string) (_result *GetFileTransResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFileTransResponse{}
	_body, _err := client.GetFileTransWithOptions(TransId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMeetingTransWithOptions(MeetingId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMeetingTransResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMeetingTrans"),
		Version:     tea.String("2022-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/meeting-trans/" + tea.StringValue(openapiutil.GetEncodeParam(MeetingId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMeetingTransResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMeetingTrans(MeetingId *string) (_result *GetMeetingTransResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMeetingTransResponse{}
	_body, _err := client.GetMeetingTransWithOptions(MeetingId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopMeetingTransWithOptions(MeetingId *string, request *StopMeetingTransRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopMeetingTransResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MeetingRoleNum)) {
		body["MeetingRoleNum"] = request.MeetingRoleNum
	}

	if !tea.BoolValue(util.IsUnset(request.OnlyRoleSplitResult)) {
		body["OnlyRoleSplitResult"] = request.OnlyRoleSplitResult
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopMeetingTrans"),
		Version:     tea.String("2022-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/meeting-trans/" + tea.StringValue(openapiutil.GetEncodeParam(MeetingId)) + "/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopMeetingTransResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopMeetingTrans(MeetingId *string, request *StopMeetingTransRequest) (_result *StopMeetingTransResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopMeetingTransResponse{}
	_body, _err := client.StopMeetingTransWithOptions(MeetingId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
