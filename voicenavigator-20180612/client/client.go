// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AssociateChatbotInstanceRequest struct {
	// example:
	//
	// chatbot-720edd02b66a
	ChatbotInstanceId *string `json:"ChatbotInstanceId,omitempty" xml:"ChatbotInstanceId,omitempty"`
	ChatbotName       *string `json:"ChatbotName,omitempty" xml:"ChatbotName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NluServiceParamsJson *string `json:"NluServiceParamsJson,omitempty" xml:"NluServiceParamsJson,omitempty"`
	NluServiceType       *string `json:"NluServiceType,omitempty" xml:"NluServiceType,omitempty"`
	UnionSource          *string `json:"UnionSource,omitempty" xml:"UnionSource,omitempty"`
}

func (s AssociateChatbotInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateChatbotInstanceRequest) GoString() string {
	return s.String()
}

func (s *AssociateChatbotInstanceRequest) SetChatbotInstanceId(v string) *AssociateChatbotInstanceRequest {
	s.ChatbotInstanceId = &v
	return s
}

func (s *AssociateChatbotInstanceRequest) SetChatbotName(v string) *AssociateChatbotInstanceRequest {
	s.ChatbotName = &v
	return s
}

func (s *AssociateChatbotInstanceRequest) SetInstanceId(v string) *AssociateChatbotInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateChatbotInstanceRequest) SetNluServiceParamsJson(v string) *AssociateChatbotInstanceRequest {
	s.NluServiceParamsJson = &v
	return s
}

func (s *AssociateChatbotInstanceRequest) SetNluServiceType(v string) *AssociateChatbotInstanceRequest {
	s.NluServiceType = &v
	return s
}

func (s *AssociateChatbotInstanceRequest) SetUnionSource(v string) *AssociateChatbotInstanceRequest {
	s.UnionSource = &v
	return s
}

type AssociateChatbotInstanceResponseBody struct {
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateChatbotInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateChatbotInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateChatbotInstanceResponseBody) SetRequestId(v string) *AssociateChatbotInstanceResponseBody {
	s.RequestId = &v
	return s
}

type AssociateChatbotInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateChatbotInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateChatbotInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateChatbotInstanceResponse) GoString() string {
	return s.String()
}

func (s *AssociateChatbotInstanceResponse) SetHeaders(v map[string]*string) *AssociateChatbotInstanceResponse {
	s.Headers = v
	return s
}

func (s *AssociateChatbotInstanceResponse) SetStatusCode(v int32) *AssociateChatbotInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateChatbotInstanceResponse) SetBody(v *AssociateChatbotInstanceResponseBody) *AssociateChatbotInstanceResponse {
	s.Body = v
	return s
}

type AuditTTSVoiceRequest struct {
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Engine    *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PitchRate  *string `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	SecretKey  *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	SpeechRate *string `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// This parameter is required.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aixia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Volume *string `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s AuditTTSVoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AuditTTSVoiceRequest) GoString() string {
	return s.String()
}

func (s *AuditTTSVoiceRequest) SetAccessKey(v string) *AuditTTSVoiceRequest {
	s.AccessKey = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetAppKey(v string) *AuditTTSVoiceRequest {
	s.AppKey = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetEngine(v string) *AuditTTSVoiceRequest {
	s.Engine = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetInstanceId(v string) *AuditTTSVoiceRequest {
	s.InstanceId = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetPitchRate(v string) *AuditTTSVoiceRequest {
	s.PitchRate = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetSecretKey(v string) *AuditTTSVoiceRequest {
	s.SecretKey = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetSpeechRate(v string) *AuditTTSVoiceRequest {
	s.SpeechRate = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetText(v string) *AuditTTSVoiceRequest {
	s.Text = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetVoice(v string) *AuditTTSVoiceRequest {
	s.Voice = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetVolume(v string) *AuditTTSVoiceRequest {
	s.Volume = &v
	return s
}

type AuditTTSVoiceResponseBody struct {
	// example:
	//
	// http://voicenavigator-cn-shanghai.oss-cn-shanghai.aliyuncs.com/exported_files/2020-02-20/ttsConfig-1582188148528-abd8e407de0a49b381bb591bd91fc073.wav?Expires=1582188208&OSSAccessKeyId=LTAIppQY5rofntVJ&Signature=FaBassElzqGEB0H2TvTKPJsOJHs%3D
	AuditionUrl *string `json:"AuditionUrl,omitempty" xml:"AuditionUrl,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuditTTSVoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuditTTSVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *AuditTTSVoiceResponseBody) SetAuditionUrl(v string) *AuditTTSVoiceResponseBody {
	s.AuditionUrl = &v
	return s
}

func (s *AuditTTSVoiceResponseBody) SetRequestId(v string) *AuditTTSVoiceResponseBody {
	s.RequestId = &v
	return s
}

type AuditTTSVoiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuditTTSVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuditTTSVoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AuditTTSVoiceResponse) GoString() string {
	return s.String()
}

func (s *AuditTTSVoiceResponse) SetHeaders(v map[string]*string) *AuditTTSVoiceResponse {
	s.Headers = v
	return s
}

func (s *AuditTTSVoiceResponse) SetStatusCode(v int32) *AuditTTSVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *AuditTTSVoiceResponse) SetBody(v *AuditTTSVoiceResponseBody) *AuditTTSVoiceResponse {
	s.Body = v
	return s
}

type BeginDialogueRequest struct {
	// example:
	//
	// 10086
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1358158****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c28fc549-d88f-4f6e-85ad-a0806e5e39c0
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// {\\"channelId\\":\\"fe2559d3-5fc9-4fa5-8314-32b9f762791d\\"}
	InitialContext *string `json:"InitialContext,omitempty" xml:"InitialContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4d7db6670b8e41b5adb1f21560ea9272
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1231639035307976
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
}

func (s BeginDialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s BeginDialogueRequest) GoString() string {
	return s.String()
}

func (s *BeginDialogueRequest) SetCalledNumber(v string) *BeginDialogueRequest {
	s.CalledNumber = &v
	return s
}

func (s *BeginDialogueRequest) SetCallingNumber(v string) *BeginDialogueRequest {
	s.CallingNumber = &v
	return s
}

func (s *BeginDialogueRequest) SetConversationId(v string) *BeginDialogueRequest {
	s.ConversationId = &v
	return s
}

func (s *BeginDialogueRequest) SetInitialContext(v string) *BeginDialogueRequest {
	s.InitialContext = &v
	return s
}

func (s *BeginDialogueRequest) SetInstanceId(v string) *BeginDialogueRequest {
	s.InstanceId = &v
	return s
}

func (s *BeginDialogueRequest) SetInstanceOwnerId(v int64) *BeginDialogueRequest {
	s.InstanceOwnerId = &v
	return s
}

type BeginDialogueResponseBody struct {
	// example:
	//
	// Broadcast
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// {\\"duration\\":31340,\\"endTime\\":1638243934786,\\"hangUpDirection\\":\\"ivr\\",\\"startTime\\":1638243903446}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// example:
	//
	// true
	Interruptible *bool `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TextResponse *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
}

func (s BeginDialogueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BeginDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *BeginDialogueResponseBody) SetAction(v string) *BeginDialogueResponseBody {
	s.Action = &v
	return s
}

func (s *BeginDialogueResponseBody) SetActionParams(v string) *BeginDialogueResponseBody {
	s.ActionParams = &v
	return s
}

func (s *BeginDialogueResponseBody) SetInterruptible(v bool) *BeginDialogueResponseBody {
	s.Interruptible = &v
	return s
}

func (s *BeginDialogueResponseBody) SetRequestId(v string) *BeginDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *BeginDialogueResponseBody) SetTextResponse(v string) *BeginDialogueResponseBody {
	s.TextResponse = &v
	return s
}

type BeginDialogueResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BeginDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BeginDialogueResponse) String() string {
	return tea.Prettify(s)
}

func (s BeginDialogueResponse) GoString() string {
	return s.String()
}

func (s *BeginDialogueResponse) SetHeaders(v map[string]*string) *BeginDialogueResponse {
	s.Headers = v
	return s
}

func (s *BeginDialogueResponse) SetStatusCode(v int32) *BeginDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *BeginDialogueResponse) SetBody(v *BeginDialogueResponseBody) *BeginDialogueResponse {
	s.Body = v
	return s
}

type CollectedNumberRequest struct {
	AdditionalContext *string `json:"AdditionalContext,omitempty" xml:"AdditionalContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0099b75d-60fd-4c63-8541-7fbba0ae6bb0
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0099b75d-60fd-4c63-8541-7fbba0ae6bb0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1426738157626835
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
	// example:
	//
	// 1500060224
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s CollectedNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s CollectedNumberRequest) GoString() string {
	return s.String()
}

func (s *CollectedNumberRequest) SetAdditionalContext(v string) *CollectedNumberRequest {
	s.AdditionalContext = &v
	return s
}

func (s *CollectedNumberRequest) SetConversationId(v string) *CollectedNumberRequest {
	s.ConversationId = &v
	return s
}

func (s *CollectedNumberRequest) SetInstanceId(v string) *CollectedNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *CollectedNumberRequest) SetInstanceOwnerId(v int64) *CollectedNumberRequest {
	s.InstanceOwnerId = &v
	return s
}

func (s *CollectedNumberRequest) SetNumber(v string) *CollectedNumberRequest {
	s.Number = &v
	return s
}

type CollectedNumberResponseBody struct {
	// example:
	//
	// Broadcast
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// {\\"duration\\":31340,\\"endTime\\":1638243934786,\\"hangUpDirection\\":\\"ivr\\",\\"startTime\\":1638243903446}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// example:
	//
	// false
	Interruptible *bool `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TextResponse *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
}

func (s CollectedNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollectedNumberResponseBody) GoString() string {
	return s.String()
}

func (s *CollectedNumberResponseBody) SetAction(v string) *CollectedNumberResponseBody {
	s.Action = &v
	return s
}

func (s *CollectedNumberResponseBody) SetActionParams(v string) *CollectedNumberResponseBody {
	s.ActionParams = &v
	return s
}

func (s *CollectedNumberResponseBody) SetInterruptible(v bool) *CollectedNumberResponseBody {
	s.Interruptible = &v
	return s
}

func (s *CollectedNumberResponseBody) SetRequestId(v string) *CollectedNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CollectedNumberResponseBody) SetTextResponse(v string) *CollectedNumberResponseBody {
	s.TextResponse = &v
	return s
}

type CollectedNumberResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CollectedNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CollectedNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s CollectedNumberResponse) GoString() string {
	return s.String()
}

func (s *CollectedNumberResponse) SetHeaders(v map[string]*string) *CollectedNumberResponse {
	s.Headers = v
	return s
}

func (s *CollectedNumberResponse) SetStatusCode(v int32) *CollectedNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *CollectedNumberResponse) SetBody(v *CollectedNumberResponseBody) *CollectedNumberResponse {
	s.Body = v
	return s
}

type CreateDownloadUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 073f092da0a847b9bf76eb88b5931c7a
	DownloadTaskId *string `json:"DownloadTaskId,omitempty" xml:"DownloadTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 22626c39603744f5a08d4d715315561a
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s CreateDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateDownloadUrlRequest) SetDownloadTaskId(v string) *CreateDownloadUrlRequest {
	s.DownloadTaskId = &v
	return s
}

func (s *CreateDownloadUrlRequest) SetFileId(v string) *CreateDownloadUrlRequest {
	s.FileId = &v
	return s
}

type CreateDownloadUrlResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// http://ssml-test.oss-cn-shanghai.aliyuncs.com/key
	FileHttpUrl *string `json:"FileHttpUrl,omitempty" xml:"FileHttpUrl,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// get upload tool url success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7401D698-0AAC-5909-B68E-88C68805FFB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDownloadUrlResponseBody) SetCode(v string) *CreateDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDownloadUrlResponseBody) SetFileHttpUrl(v string) *CreateDownloadUrlResponseBody {
	s.FileHttpUrl = &v
	return s
}

func (s *CreateDownloadUrlResponseBody) SetHttpStatusCode(v int32) *CreateDownloadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDownloadUrlResponseBody) SetMessage(v string) *CreateDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDownloadUrlResponseBody) SetRequestId(v string) *CreateDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDownloadUrlResponseBody) SetSuccess(v bool) *CreateDownloadUrlResponseBody {
	s.Success = &v
	return s
}

type CreateDownloadUrlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *CreateDownloadUrlResponse) SetHeaders(v map[string]*string) *CreateDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *CreateDownloadUrlResponse) SetStatusCode(v int32) *CreateDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDownloadUrlResponse) SetBody(v *CreateDownloadUrlResponseBody) *CreateDownloadUrlResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Concurrency *int64  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NluServiceParamsJson *string `json:"NluServiceParamsJson,omitempty" xml:"NluServiceParamsJson,omitempty"`
	UnionInstanceId      *string `json:"UnionInstanceId,omitempty" xml:"UnionInstanceId,omitempty"`
	UnionSource          *string `json:"UnionSource,omitempty" xml:"UnionSource,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetConcurrency(v int64) *CreateInstanceRequest {
	s.Concurrency = &v
	return s
}

func (s *CreateInstanceRequest) SetDescription(v string) *CreateInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequest) SetName(v string) *CreateInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateInstanceRequest) SetNluServiceParamsJson(v string) *CreateInstanceRequest {
	s.NluServiceParamsJson = &v
	return s
}

func (s *CreateInstanceRequest) SetUnionInstanceId(v string) *CreateInstanceRequest {
	s.UnionInstanceId = &v
	return s
}

func (s *CreateInstanceRequest) SetUnionSource(v string) *CreateInstanceRequest {
	s.UnionSource = &v
	return s
}

type CreateInstanceResponseBody struct {
	// example:
	//
	// c28fc549-d88f-4f6e-85ad-a0806e5e39c0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 9ab43460-c0b9-40e2-8447-48d82c97fc67
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type DebugBeginDialogueRequest struct {
	// example:
	//
	// 10086
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 135815*****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8a503680-815d-473e-a9b0-e010f47a64d2
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// {}
	InitialContext *string `json:"InitialContext,omitempty" xml:"InitialContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8a503680-815d-473e-a9b0-e010f47a64d2
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DebugBeginDialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s DebugBeginDialogueRequest) GoString() string {
	return s.String()
}

func (s *DebugBeginDialogueRequest) SetCalledNumber(v string) *DebugBeginDialogueRequest {
	s.CalledNumber = &v
	return s
}

func (s *DebugBeginDialogueRequest) SetCallingNumber(v string) *DebugBeginDialogueRequest {
	s.CallingNumber = &v
	return s
}

func (s *DebugBeginDialogueRequest) SetConversationId(v string) *DebugBeginDialogueRequest {
	s.ConversationId = &v
	return s
}

func (s *DebugBeginDialogueRequest) SetInitialContext(v string) *DebugBeginDialogueRequest {
	s.InitialContext = &v
	return s
}

func (s *DebugBeginDialogueRequest) SetInstanceId(v string) *DebugBeginDialogueRequest {
	s.InstanceId = &v
	return s
}

type DebugBeginDialogueResponseBody struct {
	// example:
	//
	// Broadcast
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// {}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// example:
	//
	// true
	Interruptible *bool `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TextResponse *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
}

func (s DebugBeginDialogueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DebugBeginDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *DebugBeginDialogueResponseBody) SetAction(v string) *DebugBeginDialogueResponseBody {
	s.Action = &v
	return s
}

func (s *DebugBeginDialogueResponseBody) SetActionParams(v string) *DebugBeginDialogueResponseBody {
	s.ActionParams = &v
	return s
}

func (s *DebugBeginDialogueResponseBody) SetInterruptible(v bool) *DebugBeginDialogueResponseBody {
	s.Interruptible = &v
	return s
}

func (s *DebugBeginDialogueResponseBody) SetRequestId(v string) *DebugBeginDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DebugBeginDialogueResponseBody) SetTextResponse(v string) *DebugBeginDialogueResponseBody {
	s.TextResponse = &v
	return s
}

type DebugBeginDialogueResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DebugBeginDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DebugBeginDialogueResponse) String() string {
	return tea.Prettify(s)
}

func (s DebugBeginDialogueResponse) GoString() string {
	return s.String()
}

func (s *DebugBeginDialogueResponse) SetHeaders(v map[string]*string) *DebugBeginDialogueResponse {
	s.Headers = v
	return s
}

func (s *DebugBeginDialogueResponse) SetStatusCode(v int32) *DebugBeginDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *DebugBeginDialogueResponse) SetBody(v *DebugBeginDialogueResponseBody) *DebugBeginDialogueResponse {
	s.Body = v
	return s
}

type DebugCollectedNumberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 7cefbff0-8d50-4d6f-b93c-73cee23c1555
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7cefbff0-8d50-4d6f-b93c-73cee23c1555
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 123
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DebugCollectedNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s DebugCollectedNumberRequest) GoString() string {
	return s.String()
}

func (s *DebugCollectedNumberRequest) SetConversationId(v string) *DebugCollectedNumberRequest {
	s.ConversationId = &v
	return s
}

func (s *DebugCollectedNumberRequest) SetInstanceId(v string) *DebugCollectedNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *DebugCollectedNumberRequest) SetNumber(v string) *DebugCollectedNumberRequest {
	s.Number = &v
	return s
}

type DebugCollectedNumberResponseBody struct {
	// example:
	//
	// Broadcast
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// {}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// example:
	//
	// true
	Interruptible *bool `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	// example:
	//
	// abb4aa26-3a8e-43dd-82f8-0c3898c9c67f
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TextResponse *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
}

func (s DebugCollectedNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DebugCollectedNumberResponseBody) GoString() string {
	return s.String()
}

func (s *DebugCollectedNumberResponseBody) SetAction(v string) *DebugCollectedNumberResponseBody {
	s.Action = &v
	return s
}

func (s *DebugCollectedNumberResponseBody) SetActionParams(v string) *DebugCollectedNumberResponseBody {
	s.ActionParams = &v
	return s
}

func (s *DebugCollectedNumberResponseBody) SetInterruptible(v bool) *DebugCollectedNumberResponseBody {
	s.Interruptible = &v
	return s
}

func (s *DebugCollectedNumberResponseBody) SetRequestId(v string) *DebugCollectedNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DebugCollectedNumberResponseBody) SetTextResponse(v string) *DebugCollectedNumberResponseBody {
	s.TextResponse = &v
	return s
}

type DebugCollectedNumberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DebugCollectedNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DebugCollectedNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s DebugCollectedNumberResponse) GoString() string {
	return s.String()
}

func (s *DebugCollectedNumberResponse) SetHeaders(v map[string]*string) *DebugCollectedNumberResponse {
	s.Headers = v
	return s
}

func (s *DebugCollectedNumberResponse) SetStatusCode(v int32) *DebugCollectedNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *DebugCollectedNumberResponse) SetBody(v *DebugCollectedNumberResponseBody) *DebugCollectedNumberResponse {
	s.Body = v
	return s
}

type DebugDialogueRequest struct {
	// example:
	//
	// {}
	AdditionalContext *string `json:"AdditionalContext,omitempty" xml:"AdditionalContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7cefbff0-8d50-4d6f-b93c-73cee23c1555
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abb4aa26-3a8e-43dd-82f8-0c3898c9c67f
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s DebugDialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s DebugDialogueRequest) GoString() string {
	return s.String()
}

func (s *DebugDialogueRequest) SetAdditionalContext(v string) *DebugDialogueRequest {
	s.AdditionalContext = &v
	return s
}

func (s *DebugDialogueRequest) SetConversationId(v string) *DebugDialogueRequest {
	s.ConversationId = &v
	return s
}

func (s *DebugDialogueRequest) SetInstanceId(v string) *DebugDialogueRequest {
	s.InstanceId = &v
	return s
}

func (s *DebugDialogueRequest) SetUtterance(v string) *DebugDialogueRequest {
	s.Utterance = &v
	return s
}

type DebugDialogueResponseBody struct {
	// example:
	//
	// Broadcast
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// {}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// example:
	//
	// true
	Interruptible *bool `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	// example:
	//
	// d74d6290-7cbe-4436-b5d7-014ebb0f4060
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 80d11be3-faad-4101-b4b0-59dbea28aaf0
	TextResponse *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
}

func (s DebugDialogueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DebugDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *DebugDialogueResponseBody) SetAction(v string) *DebugDialogueResponseBody {
	s.Action = &v
	return s
}

func (s *DebugDialogueResponseBody) SetActionParams(v string) *DebugDialogueResponseBody {
	s.ActionParams = &v
	return s
}

func (s *DebugDialogueResponseBody) SetInterruptible(v bool) *DebugDialogueResponseBody {
	s.Interruptible = &v
	return s
}

func (s *DebugDialogueResponseBody) SetRequestId(v string) *DebugDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DebugDialogueResponseBody) SetTextResponse(v string) *DebugDialogueResponseBody {
	s.TextResponse = &v
	return s
}

type DebugDialogueResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DebugDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DebugDialogueResponse) String() string {
	return tea.Prettify(s)
}

func (s DebugDialogueResponse) GoString() string {
	return s.String()
}

func (s *DebugDialogueResponse) SetHeaders(v map[string]*string) *DebugDialogueResponse {
	s.Headers = v
	return s
}

func (s *DebugDialogueResponse) SetStatusCode(v int32) *DebugDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *DebugDialogueResponse) SetBody(v *DebugDialogueResponseBody) *DebugDialogueResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// d74d6290-7cbe-4436-b5d7-014ebb0f4060
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

type DeleteInstanceResponseBody struct {
	// example:
	//
	// d74d6290-7cbe-4436-b5d7-014ebb0f4060
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetStatusCode(v int32) *DeleteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DescribeConversationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 15608cce-36be-43d5-9361-178cbe64127b
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5daac920-d6c1-429f-a95f-2a798f5255b5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConversationRequest) GoString() string {
	return s.String()
}

func (s *DescribeConversationRequest) SetConversationId(v string) *DescribeConversationRequest {
	s.ConversationId = &v
	return s
}

func (s *DescribeConversationRequest) SetInstanceId(v string) *DescribeConversationRequest {
	s.InstanceId = &v
	return s
}

type DescribeConversationResponseBody struct {
	// example:
	//
	// 1582103260434
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 138106*****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// example:
	//
	// 2d5aa451-661f-4f08-b0c4-28eec78decc4
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// 8
	EffectiveAnswerCount *int32 `json:"EffectiveAnswerCount,omitempty" xml:"EffectiveAnswerCount,omitempty"`
	// example:
	//
	// 1582103299434
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ABABCBAC
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// true
	TransferredToAgent *bool `json:"TransferredToAgent,omitempty" xml:"TransferredToAgent,omitempty"`
	// example:
	//
	// 10
	UserUtteranceCount *int32 `json:"UserUtteranceCount,omitempty" xml:"UserUtteranceCount,omitempty"`
}

func (s DescribeConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConversationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConversationResponseBody) SetBeginTime(v int64) *DescribeConversationResponseBody {
	s.BeginTime = &v
	return s
}

func (s *DescribeConversationResponseBody) SetCallingNumber(v string) *DescribeConversationResponseBody {
	s.CallingNumber = &v
	return s
}

func (s *DescribeConversationResponseBody) SetConversationId(v string) *DescribeConversationResponseBody {
	s.ConversationId = &v
	return s
}

func (s *DescribeConversationResponseBody) SetEffectiveAnswerCount(v int32) *DescribeConversationResponseBody {
	s.EffectiveAnswerCount = &v
	return s
}

func (s *DescribeConversationResponseBody) SetEndTime(v int64) *DescribeConversationResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeConversationResponseBody) SetRequestId(v string) *DescribeConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConversationResponseBody) SetSkillGroupId(v string) *DescribeConversationResponseBody {
	s.SkillGroupId = &v
	return s
}

func (s *DescribeConversationResponseBody) SetTransferredToAgent(v bool) *DescribeConversationResponseBody {
	s.TransferredToAgent = &v
	return s
}

func (s *DescribeConversationResponseBody) SetUserUtteranceCount(v int32) *DescribeConversationResponseBody {
	s.UserUtteranceCount = &v
	return s
}

type DescribeConversationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConversationResponse) GoString() string {
	return s.String()
}

func (s *DescribeConversationResponse) SetHeaders(v map[string]*string) *DescribeConversationResponse {
	s.Headers = v
	return s
}

func (s *DescribeConversationResponse) SetStatusCode(v int32) *DescribeConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConversationResponse) SetBody(v *DescribeConversationResponseBody) *DescribeConversationResponse {
	s.Body = v
	return s
}

type DescribeConversationContextRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 026ca0f4-483b-4252-ae1d-1f15f056f8b9
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeConversationContextRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConversationContextRequest) GoString() string {
	return s.String()
}

func (s *DescribeConversationContextRequest) SetConversationId(v string) *DescribeConversationContextRequest {
	s.ConversationId = &v
	return s
}

func (s *DescribeConversationContextRequest) SetInstanceId(v string) *DescribeConversationContextRequest {
	s.InstanceId = &v
	return s
}

type DescribeConversationContextResponseBody struct {
	// example:
	//
	// {         "CallingNumber": "135815***",         "AdditionalContext": "",         "ConversationId": "361c8a53-0e29-42f3-8aa7-c7752d010399"     }
	ConversationContext *string `json:"ConversationContext,omitempty" xml:"ConversationContext,omitempty"`
	// example:
	//
	// b19af5ce5314ac08108d1b33fe20e15
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeConversationContextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConversationContextResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConversationContextResponseBody) SetConversationContext(v string) *DescribeConversationContextResponseBody {
	s.ConversationContext = &v
	return s
}

func (s *DescribeConversationContextResponseBody) SetRequestId(v string) *DescribeConversationContextResponseBody {
	s.RequestId = &v
	return s
}

type DescribeConversationContextResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConversationContextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConversationContextResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConversationContextResponse) GoString() string {
	return s.String()
}

func (s *DescribeConversationContextResponse) SetHeaders(v map[string]*string) *DescribeConversationContextResponse {
	s.Headers = v
	return s
}

func (s *DescribeConversationContextResponse) SetStatusCode(v int32) *DescribeConversationContextResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConversationContextResponse) SetBody(v *DescribeConversationContextResponseBody) *DescribeConversationContextResponse {
	s.Body = v
	return s
}

type DescribeExportProgressRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0de8e5ccc2b645039ae6fbda443da73f
	ExportTaskId *string `json:"ExportTaskId,omitempty" xml:"ExportTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 868eef14-7515-4856-8a50-5c9a22abdbcc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeExportProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportProgressRequest) GoString() string {
	return s.String()
}

func (s *DescribeExportProgressRequest) SetExportTaskId(v string) *DescribeExportProgressRequest {
	s.ExportTaskId = &v
	return s
}

func (s *DescribeExportProgressRequest) SetInstanceId(v string) *DescribeExportProgressRequest {
	s.InstanceId = &v
	return s
}

type DescribeExportProgressResponseBody struct {
	// example:
	//
	// http://ssml-test.oss-cn-shanghai.aliyuncs.com/key
	FileHttpUrl *string `json:"FileHttpUrl,omitempty" xml:"FileHttpUrl,omitempty"`
	// example:
	//
	// b19af5ce5314ac08108d1b33fe20e15
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// FINISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeExportProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportProgressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExportProgressResponseBody) SetFileHttpUrl(v string) *DescribeExportProgressResponseBody {
	s.FileHttpUrl = &v
	return s
}

func (s *DescribeExportProgressResponseBody) SetRequestId(v string) *DescribeExportProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExportProgressResponseBody) SetStatus(v string) *DescribeExportProgressResponseBody {
	s.Status = &v
	return s
}

type DescribeExportProgressResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExportProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExportProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportProgressResponse) GoString() string {
	return s.String()
}

func (s *DescribeExportProgressResponse) SetHeaders(v map[string]*string) *DescribeExportProgressResponse {
	s.Headers = v
	return s
}

func (s *DescribeExportProgressResponse) SetStatusCode(v int32) *DescribeExportProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExportProgressResponse) SetBody(v *DescribeExportProgressResponseBody) *DescribeExportProgressResponse {
	s.Body = v
	return s
}

type DescribeInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ecbfa5e3-1838-4e8a-aa08-fa8b713b82df
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) SetInstanceId(v string) *DescribeInstanceRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceResponseBody struct {
	AbilityType          *string   `json:"AbilityType,omitempty" xml:"AbilityType,omitempty"`
	ApplicableOperations []*string `json:"ApplicableOperations,omitempty" xml:"ApplicableOperations,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Concurrency *int64  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// test1_instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1532436395329
	ModifyTime           *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ModifyUserName       *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NluServiceParamsJson *string `json:"NluServiceParamsJson,omitempty" xml:"NluServiceParamsJson,omitempty"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Drafted
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UnionInstanceId *string `json:"UnionInstanceId,omitempty" xml:"UnionInstanceId,omitempty"`
	UnionSource     *string `json:"UnionSource,omitempty" xml:"UnionSource,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetAbilityType(v string) *DescribeInstanceResponseBody {
	s.AbilityType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetApplicableOperations(v []*string) *DescribeInstanceResponseBody {
	s.ApplicableOperations = v
	return s
}

func (s *DescribeInstanceResponseBody) SetConcurrency(v int64) *DescribeInstanceResponseBody {
	s.Concurrency = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetDescription(v string) *DescribeInstanceResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceId(v string) *DescribeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetModifyTime(v int64) *DescribeInstanceResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetModifyUserName(v string) *DescribeInstanceResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetName(v string) *DescribeInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetNluServiceParamsJson(v string) *DescribeInstanceResponseBody {
	s.NluServiceParamsJson = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetStatus(v string) *DescribeInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetUnionInstanceId(v string) *DescribeInstanceResponseBody {
	s.UnionInstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetUnionSource(v string) *DescribeInstanceResponseBody {
	s.UnionSource = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetStatusCode(v int32) *DescribeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeNavigationConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 287289b6-1510-4e64-9224-39b53ad89dd2
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeNavigationConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNavigationConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigRequest) SetInstanceId(v string) *DescribeNavigationConfigRequest {
	s.InstanceId = &v
	return s
}

type DescribeNavigationConfigResponseBody struct {
	GreetingConfig *DescribeNavigationConfigResponseBodyGreetingConfig `json:"GreetingConfig,omitempty" xml:"GreetingConfig,omitempty" type:"Struct"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId            *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SilenceTimeoutConfig *DescribeNavigationConfigResponseBodySilenceTimeoutConfig `json:"SilenceTimeoutConfig,omitempty" xml:"SilenceTimeoutConfig,omitempty" type:"Struct"`
	UnrecognizingConfig  *DescribeNavigationConfigResponseBodyUnrecognizingConfig  `json:"UnrecognizingConfig,omitempty" xml:"UnrecognizingConfig,omitempty" type:"Struct"`
}

func (s DescribeNavigationConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNavigationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigResponseBody) SetGreetingConfig(v *DescribeNavigationConfigResponseBodyGreetingConfig) *DescribeNavigationConfigResponseBody {
	s.GreetingConfig = v
	return s
}

func (s *DescribeNavigationConfigResponseBody) SetRequestId(v string) *DescribeNavigationConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNavigationConfigResponseBody) SetSilenceTimeoutConfig(v *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) *DescribeNavigationConfigResponseBody {
	s.SilenceTimeoutConfig = v
	return s
}

func (s *DescribeNavigationConfigResponseBody) SetUnrecognizingConfig(v *DescribeNavigationConfigResponseBodyUnrecognizingConfig) *DescribeNavigationConfigResponseBody {
	s.UnrecognizingConfig = v
	return s
}

type DescribeNavigationConfigResponseBodyGreetingConfig struct {
	GreetingWords *string `json:"GreetingWords,omitempty" xml:"GreetingWords,omitempty"`
	IntentTrigger *string `json:"IntentTrigger,omitempty" xml:"IntentTrigger,omitempty"`
	// example:
	//
	// chatbotIntent
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeNavigationConfigResponseBodyGreetingConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeNavigationConfigResponseBodyGreetingConfig) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigResponseBodyGreetingConfig) SetGreetingWords(v string) *DescribeNavigationConfigResponseBodyGreetingConfig {
	s.GreetingWords = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyGreetingConfig) SetIntentTrigger(v string) *DescribeNavigationConfigResponseBodyGreetingConfig {
	s.IntentTrigger = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyGreetingConfig) SetSourceType(v string) *DescribeNavigationConfigResponseBodyGreetingConfig {
	s.SourceType = &v
	return s
}

type DescribeNavigationConfigResponseBodySilenceTimeoutConfig struct {
	// example:
	//
	// HangUp
	FinalAction *string `json:"FinalAction,omitempty" xml:"FinalAction,omitempty"`
	// example:
	//
	// {}
	FinalActionParams *string `json:"FinalActionParams,omitempty" xml:"FinalActionParams,omitempty"`
	FinalPrompt       *string `json:"FinalPrompt,omitempty" xml:"FinalPrompt,omitempty"`
	// ""
	//
	// example:
	//
	// ""
	IntentTrigger *string `json:"IntentTrigger,omitempty" xml:"IntentTrigger,omitempty"`
	Prompt        *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// chatbotIntent
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 3
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// example:
	//
	// 10
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s DescribeNavigationConfigResponseBodySilenceTimeoutConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeNavigationConfigResponseBodySilenceTimeoutConfig) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetFinalAction(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.FinalAction = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetFinalActionParams(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.FinalActionParams = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetFinalPrompt(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.FinalPrompt = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetIntentTrigger(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.IntentTrigger = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetPrompt(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.Prompt = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetSourceType(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.SourceType = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetThreshold(v int32) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.Threshold = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetTimeout(v int64) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.Timeout = &v
	return s
}

type DescribeNavigationConfigResponseBodyUnrecognizingConfig struct {
	// example:
	//
	// TransferToAgent
	FinalAction *string `json:"FinalAction,omitempty" xml:"FinalAction,omitempty"`
	// example:
	//
	// { \\"skillGroupId\\": \\"fallbackSkillGroup\\" }
	FinalActionParams *string `json:"FinalActionParams,omitempty" xml:"FinalActionParams,omitempty"`
	FinalPrompt       *string `json:"FinalPrompt,omitempty" xml:"FinalPrompt,omitempty"`
	Prompt            *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 3
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeNavigationConfigResponseBodyUnrecognizingConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeNavigationConfigResponseBodyUnrecognizingConfig) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) SetFinalAction(v string) *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	s.FinalAction = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) SetFinalActionParams(v string) *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	s.FinalActionParams = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) SetFinalPrompt(v string) *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	s.FinalPrompt = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) SetPrompt(v string) *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	s.Prompt = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) SetThreshold(v int32) *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	s.Threshold = &v
	return s
}

type DescribeNavigationConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNavigationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNavigationConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNavigationConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigResponse) SetHeaders(v map[string]*string) *DescribeNavigationConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeNavigationConfigResponse) SetStatusCode(v int32) *DescribeNavigationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNavigationConfigResponse) SetBody(v *DescribeNavigationConfigResponseBody) *DescribeNavigationConfigResponse {
	s.Body = v
	return s
}

type DescribeRecordingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// abb4aa26-3a8e-43dd-82f8-0c3898c9c67f
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7cefbff0-8d50-4d6f-b93c-73cee23c1555
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NeedVoiceSliceRecording *bool   `json:"NeedVoiceSliceRecording,omitempty" xml:"NeedVoiceSliceRecording,omitempty"`
}

func (s DescribeRecordingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordingRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordingRequest) SetConversationId(v string) *DescribeRecordingRequest {
	s.ConversationId = &v
	return s
}

func (s *DescribeRecordingRequest) SetInstanceId(v string) *DescribeRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRecordingRequest) SetNeedVoiceSliceRecording(v bool) *DescribeRecordingRequest {
	s.NeedVoiceSliceRecording = &v
	return s
}

type DescribeRecordingResponseBody struct {
	// example:
	//
	// 2019080913202222.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// url
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId                   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VoiceSliceRecordingListJson *string `json:"VoiceSliceRecordingListJson,omitempty" xml:"VoiceSliceRecordingListJson,omitempty"`
}

func (s DescribeRecordingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordingResponseBody) SetFileName(v string) *DescribeRecordingResponseBody {
	s.FileName = &v
	return s
}

func (s *DescribeRecordingResponseBody) SetFilePath(v string) *DescribeRecordingResponseBody {
	s.FilePath = &v
	return s
}

func (s *DescribeRecordingResponseBody) SetRequestId(v string) *DescribeRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordingResponseBody) SetVoiceSliceRecordingListJson(v string) *DescribeRecordingResponseBody {
	s.VoiceSliceRecordingListJson = &v
	return s
}

type DescribeRecordingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecordingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordingResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordingResponse) SetHeaders(v map[string]*string) *DescribeRecordingResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordingResponse) SetStatusCode(v int32) *DescribeRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecordingResponse) SetBody(v *DescribeRecordingResponseBody) *DescribeRecordingResponse {
	s.Body = v
	return s
}

type DescribeStatisticalDataRequest struct {
	// example:
	//
	// 1582283640000
	BeginTimeLeftRange *int64 `json:"BeginTimeLeftRange,omitempty" xml:"BeginTimeLeftRange,omitempty"`
	// example:
	//
	// 1582298040000
	BeginTimeRightRange *int64 `json:"BeginTimeRightRange,omitempty" xml:"BeginTimeRightRange,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c28fc549-d88f-4f6e-85ad-a0806e5e39c0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Day/Hour
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s DescribeStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeStatisticalDataRequest) SetBeginTimeLeftRange(v int64) *DescribeStatisticalDataRequest {
	s.BeginTimeLeftRange = &v
	return s
}

func (s *DescribeStatisticalDataRequest) SetBeginTimeRightRange(v int64) *DescribeStatisticalDataRequest {
	s.BeginTimeRightRange = &v
	return s
}

func (s *DescribeStatisticalDataRequest) SetInstanceId(v string) *DescribeStatisticalDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeStatisticalDataRequest) SetTimeUnit(v string) *DescribeStatisticalDataRequest {
	s.TimeUnit = &v
	return s
}

type DescribeStatisticalDataResponseBody struct {
	// example:
	//
	// 100
	ConversationTotalNum *int64 `json:"ConversationTotalNum,omitempty" xml:"ConversationTotalNum,omitempty"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 80
	ResolvedQuestionTotalNum *int64                                                       `json:"ResolvedQuestionTotalNum,omitempty" xml:"ResolvedQuestionTotalNum,omitempty"`
	StatisticalDataReports   []*DescribeStatisticalDataResponseBodyStatisticalDataReports `json:"StatisticalDataReports,omitempty" xml:"StatisticalDataReports,omitempty" type:"Repeated"`
	// example:
	//
	// 80.00%
	TotalDialoguePassRate *string `json:"TotalDialoguePassRate,omitempty" xml:"TotalDialoguePassRate,omitempty"`
	// example:
	//
	// 80.00%
	TotalKnowledgeHitRate *string `json:"TotalKnowledgeHitRate,omitempty" xml:"TotalKnowledgeHitRate,omitempty"`
	// example:
	//
	// 80.00%
	TotalResolutionRate *string `json:"TotalResolutionRate,omitempty" xml:"TotalResolutionRate,omitempty"`
	// example:
	//
	// 80.00%
	TotalValidAnswerRate *string `json:"TotalValidAnswerRate,omitempty" xml:"TotalValidAnswerRate,omitempty"`
}

func (s DescribeStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStatisticalDataResponseBody) SetConversationTotalNum(v int64) *DescribeStatisticalDataResponseBody {
	s.ConversationTotalNum = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetRequestId(v string) *DescribeStatisticalDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetResolvedQuestionTotalNum(v int64) *DescribeStatisticalDataResponseBody {
	s.ResolvedQuestionTotalNum = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetStatisticalDataReports(v []*DescribeStatisticalDataResponseBodyStatisticalDataReports) *DescribeStatisticalDataResponseBody {
	s.StatisticalDataReports = v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetTotalDialoguePassRate(v string) *DescribeStatisticalDataResponseBody {
	s.TotalDialoguePassRate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetTotalKnowledgeHitRate(v string) *DescribeStatisticalDataResponseBody {
	s.TotalKnowledgeHitRate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetTotalResolutionRate(v string) *DescribeStatisticalDataResponseBody {
	s.TotalResolutionRate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetTotalValidAnswerRate(v string) *DescribeStatisticalDataResponseBody {
	s.TotalValidAnswerRate = &v
	return s
}

type DescribeStatisticalDataResponseBodyStatisticalDataReports struct {
	// example:
	//
	// 80.00%
	DialoguePassRate *string `json:"DialoguePassRate,omitempty" xml:"DialoguePassRate,omitempty"`
	// example:
	//
	// 80.00%
	KnowledgeHitRate *string `json:"KnowledgeHitRate,omitempty" xml:"KnowledgeHitRate,omitempty"`
	// example:
	//
	// 80.00%
	ResolutionRate *string `json:"ResolutionRate,omitempty" xml:"ResolutionRate,omitempty"`
	// example:
	//
	// 80
	ResolvedQuestionNum *int32 `json:"ResolvedQuestionNum,omitempty" xml:"ResolvedQuestionNum,omitempty"`
	// example:
	//
	// 19:00:00
	StatisticalDate *string `json:"StatisticalDate,omitempty" xml:"StatisticalDate,omitempty"`
	// example:
	//
	// 100
	TotalConversationNum *int32 `json:"TotalConversationNum,omitempty" xml:"TotalConversationNum,omitempty"`
	// example:
	//
	// 80.0
	ValidAnswerRate *string `json:"ValidAnswerRate,omitempty" xml:"ValidAnswerRate,omitempty"`
}

func (s DescribeStatisticalDataResponseBodyStatisticalDataReports) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticalDataResponseBodyStatisticalDataReports) GoString() string {
	return s.String()
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetDialoguePassRate(v string) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.DialoguePassRate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetKnowledgeHitRate(v string) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.KnowledgeHitRate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetResolutionRate(v string) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.ResolutionRate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetResolvedQuestionNum(v int32) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.ResolvedQuestionNum = &v
	return s
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetStatisticalDate(v string) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.StatisticalDate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetTotalConversationNum(v int32) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.TotalConversationNum = &v
	return s
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetValidAnswerRate(v string) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.ValidAnswerRate = &v
	return s
}

type DescribeStatisticalDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeStatisticalDataResponse) SetHeaders(v map[string]*string) *DescribeStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeStatisticalDataResponse) SetStatusCode(v int32) *DescribeStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStatisticalDataResponse) SetBody(v *DescribeStatisticalDataResponseBody) *DescribeStatisticalDataResponse {
	s.Body = v
	return s
}

type DescribeTTSConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dc437bba-5a25-4bbc-b4c2-f268864bebb5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1971226538081821
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
}

func (s DescribeTTSConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTTSConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeTTSConfigRequest) SetInstanceId(v string) *DescribeTTSConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTTSConfigRequest) SetInstanceOwnerId(v int64) *DescribeTTSConfigRequest {
	s.InstanceOwnerId = &v
	return s
}

type DescribeTTSConfigResponseBody struct {
	AliCustomizedVoice *string `json:"AliCustomizedVoice,omitempty" xml:"AliCustomizedVoice,omitempty"`
	AppKey             *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Engine             *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineXunfei       *string `json:"EngineXunfei,omitempty" xml:"EngineXunfei,omitempty"`
	NlsServiceType     *string `json:"NlsServiceType,omitempty" xml:"NlsServiceType,omitempty"`
	PitchRate          *int32  `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	// example:
	//
	// F132DDBA-66C4-5BD3-BF7E-9642FE877158
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// -150
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// example:
	//
	// aixia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s DescribeTTSConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTTSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTTSConfigResponseBody) SetAliCustomizedVoice(v string) *DescribeTTSConfigResponseBody {
	s.AliCustomizedVoice = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetAppKey(v string) *DescribeTTSConfigResponseBody {
	s.AppKey = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetEngine(v string) *DescribeTTSConfigResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetEngineXunfei(v string) *DescribeTTSConfigResponseBody {
	s.EngineXunfei = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetNlsServiceType(v string) *DescribeTTSConfigResponseBody {
	s.NlsServiceType = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetPitchRate(v int32) *DescribeTTSConfigResponseBody {
	s.PitchRate = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetRequestId(v string) *DescribeTTSConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetSpeechRate(v int32) *DescribeTTSConfigResponseBody {
	s.SpeechRate = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetVoice(v string) *DescribeTTSConfigResponseBody {
	s.Voice = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetVolume(v int32) *DescribeTTSConfigResponseBody {
	s.Volume = &v
	return s
}

type DescribeTTSConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTTSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTTSConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTTSConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeTTSConfigResponse) SetHeaders(v map[string]*string) *DescribeTTSConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeTTSConfigResponse) SetStatusCode(v int32) *DescribeTTSConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTTSConfigResponse) SetBody(v *DescribeTTSConfigResponseBody) *DescribeTTSConfigResponse {
	s.Body = v
	return s
}

type DialogueRequest struct {
	// example:
	//
	// {}
	AdditionalContext *string `json:"AdditionalContext,omitempty" xml:"AdditionalContext,omitempty"`
	// example:
	//
	// 10086
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// example:
	//
	// 18851708605
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	Emotion        *string `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 21e0b2a3-168d-4ba7-9009-afc42666eb54
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1426738157626835
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
	// This parameter is required.
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s DialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s DialogueRequest) GoString() string {
	return s.String()
}

func (s *DialogueRequest) SetAdditionalContext(v string) *DialogueRequest {
	s.AdditionalContext = &v
	return s
}

func (s *DialogueRequest) SetCalledNumber(v string) *DialogueRequest {
	s.CalledNumber = &v
	return s
}

func (s *DialogueRequest) SetCallingNumber(v string) *DialogueRequest {
	s.CallingNumber = &v
	return s
}

func (s *DialogueRequest) SetConversationId(v string) *DialogueRequest {
	s.ConversationId = &v
	return s
}

func (s *DialogueRequest) SetEmotion(v string) *DialogueRequest {
	s.Emotion = &v
	return s
}

func (s *DialogueRequest) SetInstanceId(v string) *DialogueRequest {
	s.InstanceId = &v
	return s
}

func (s *DialogueRequest) SetInstanceOwnerId(v int64) *DialogueRequest {
	s.InstanceOwnerId = &v
	return s
}

func (s *DialogueRequest) SetUtterance(v string) *DialogueRequest {
	s.Utterance = &v
	return s
}

type DialogueResponseBody struct {
	// example:
	//
	// Broadcast
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// {\\"duration\\":2420,\\"endTime\\":1651717326805,\\"hangUpDirection\\":\\"client\\",\\"hasLastPlaybackCompleted\\":true,\\"startTime\\":1651717324385}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// example:
	//
	// true
	Interruptible *bool `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	// example:
	//
	// 9DB8BA95-4513-54B9-9C67-A28909CFB4AD
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TextResponse *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
}

func (s DialogueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DialogueResponseBody) GoString() string {
	return s.String()
}

func (s *DialogueResponseBody) SetAction(v string) *DialogueResponseBody {
	s.Action = &v
	return s
}

func (s *DialogueResponseBody) SetActionParams(v string) *DialogueResponseBody {
	s.ActionParams = &v
	return s
}

func (s *DialogueResponseBody) SetInterruptible(v bool) *DialogueResponseBody {
	s.Interruptible = &v
	return s
}

func (s *DialogueResponseBody) SetRequestId(v string) *DialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DialogueResponseBody) SetTextResponse(v string) *DialogueResponseBody {
	s.TextResponse = &v
	return s
}

type DialogueResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DialogueResponse) String() string {
	return tea.Prettify(s)
}

func (s DialogueResponse) GoString() string {
	return s.String()
}

func (s *DialogueResponse) SetHeaders(v map[string]*string) *DialogueResponse {
	s.Headers = v
	return s
}

func (s *DialogueResponse) SetStatusCode(v int32) *DialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *DialogueResponse) SetBody(v *DialogueResponseBody) *DialogueResponse {
	s.Body = v
	return s
}

type DisableInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableInstanceRequest) GoString() string {
	return s.String()
}

func (s *DisableInstanceRequest) SetInstanceId(v string) *DisableInstanceRequest {
	s.InstanceId = &v
	return s
}

type DisableInstanceResponseBody struct {
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DisableInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DisableInstanceResponseBody) SetRequestId(v string) *DisableInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableInstanceResponseBody) SetStatus(v string) *DisableInstanceResponseBody {
	s.Status = &v
	return s
}

type DisableInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableInstanceResponse) GoString() string {
	return s.String()
}

func (s *DisableInstanceResponse) SetHeaders(v map[string]*string) *DisableInstanceResponse {
	s.Headers = v
	return s
}

func (s *DisableInstanceResponse) SetStatusCode(v int32) *DisableInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableInstanceResponse) SetBody(v *DisableInstanceResponseBody) *DisableInstanceResponse {
	s.Body = v
	return s
}

type EnableInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableInstanceRequest) GoString() string {
	return s.String()
}

func (s *EnableInstanceRequest) SetInstanceId(v string) *EnableInstanceRequest {
	s.InstanceId = &v
	return s
}

type EnableInstanceResponseBody struct {
	// example:
	//
	// 3a530dc0-7cfa-48f6-9539-bf9001e77b16
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnableInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *EnableInstanceResponseBody) SetRequestId(v string) *EnableInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableInstanceResponseBody) SetStatus(v string) *EnableInstanceResponseBody {
	s.Status = &v
	return s
}

type EnableInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableInstanceResponse) GoString() string {
	return s.String()
}

func (s *EnableInstanceResponse) SetHeaders(v map[string]*string) *EnableInstanceResponse {
	s.Headers = v
	return s
}

func (s *EnableInstanceResponse) SetStatusCode(v int32) *EnableInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableInstanceResponse) SetBody(v *EnableInstanceResponseBody) *EnableInstanceResponse {
	s.Body = v
	return s
}

type EndDialogueRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8fb819b5-d032-48a9-ae5e-cff041b83596
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// {\\"duration\\":40,\\"endTime\\":1645082505345,\\"hangUpDirection\\":\\"ivr\\",\\"hasLastPlaybackCompleted\\":true,\\"startTime\\":1645082505305}
	HangUpParams *string `json:"HangUpParams,omitempty" xml:"HangUpParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e48e45dd-e47a-4744-a063-f08cbebb1c5a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1399572315967217
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
}

func (s EndDialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s EndDialogueRequest) GoString() string {
	return s.String()
}

func (s *EndDialogueRequest) SetConversationId(v string) *EndDialogueRequest {
	s.ConversationId = &v
	return s
}

func (s *EndDialogueRequest) SetHangUpParams(v string) *EndDialogueRequest {
	s.HangUpParams = &v
	return s
}

func (s *EndDialogueRequest) SetInstanceId(v string) *EndDialogueRequest {
	s.InstanceId = &v
	return s
}

func (s *EndDialogueRequest) SetInstanceOwnerId(v int64) *EndDialogueRequest {
	s.InstanceOwnerId = &v
	return s
}

type EndDialogueResponseBody struct {
	// example:
	//
	// e48e45dd-e47a-4744-a063-f08cbebb1c5a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EndDialogueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EndDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *EndDialogueResponseBody) SetRequestId(v string) *EndDialogueResponseBody {
	s.RequestId = &v
	return s
}

type EndDialogueResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EndDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EndDialogueResponse) String() string {
	return tea.Prettify(s)
}

func (s EndDialogueResponse) GoString() string {
	return s.String()
}

func (s *EndDialogueResponse) SetHeaders(v map[string]*string) *EndDialogueResponse {
	s.Headers = v
	return s
}

func (s *EndDialogueResponse) SetStatusCode(v int32) *EndDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *EndDialogueResponse) SetBody(v *EndDialogueResponseBody) *EndDialogueResponse {
	s.Body = v
	return s
}

type ExportConversationDetailsRequest struct {
	// example:
	//
	// 1582266750353
	BeginTimeLeftRange *int64 `json:"BeginTimeLeftRange,omitempty" xml:"BeginTimeLeftRange,omitempty"`
	// example:
	//
	// 1640793599000
	BeginTimeRightRange *int64 `json:"BeginTimeRightRange,omitempty" xml:"BeginTimeRightRange,omitempty"`
	// example:
	//
	// 13581588**
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6c01a99f-1b72-4f75-a8bd-3875766bd19d
	InstanceId       *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Options          []*string `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
	Result           *int32    `json:"Result,omitempty" xml:"Result,omitempty"`
	RoundsLeftRange  *int32    `json:"RoundsLeftRange,omitempty" xml:"RoundsLeftRange,omitempty"`
	RoundsRightRange *int32    `json:"RoundsRightRange,omitempty" xml:"RoundsRightRange,omitempty"`
}

func (s ExportConversationDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportConversationDetailsRequest) GoString() string {
	return s.String()
}

func (s *ExportConversationDetailsRequest) SetBeginTimeLeftRange(v int64) *ExportConversationDetailsRequest {
	s.BeginTimeLeftRange = &v
	return s
}

func (s *ExportConversationDetailsRequest) SetBeginTimeRightRange(v int64) *ExportConversationDetailsRequest {
	s.BeginTimeRightRange = &v
	return s
}

func (s *ExportConversationDetailsRequest) SetCallingNumber(v string) *ExportConversationDetailsRequest {
	s.CallingNumber = &v
	return s
}

func (s *ExportConversationDetailsRequest) SetInstanceId(v string) *ExportConversationDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *ExportConversationDetailsRequest) SetOptions(v []*string) *ExportConversationDetailsRequest {
	s.Options = v
	return s
}

func (s *ExportConversationDetailsRequest) SetResult(v int32) *ExportConversationDetailsRequest {
	s.Result = &v
	return s
}

func (s *ExportConversationDetailsRequest) SetRoundsLeftRange(v int32) *ExportConversationDetailsRequest {
	s.RoundsLeftRange = &v
	return s
}

func (s *ExportConversationDetailsRequest) SetRoundsRightRange(v int32) *ExportConversationDetailsRequest {
	s.RoundsRightRange = &v
	return s
}

type ExportConversationDetailsResponseBody struct {
	// example:
	//
	// 6203fc87271a420c98eab6c2bbc2d856
	ExportTaskId *string `json:"ExportTaskId,omitempty" xml:"ExportTaskId,omitempty"`
	// example:
	//
	// 75BAAB9B-40B2-5FF5-A59A-7BCF8154C6EE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportConversationDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportConversationDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ExportConversationDetailsResponseBody) SetExportTaskId(v string) *ExportConversationDetailsResponseBody {
	s.ExportTaskId = &v
	return s
}

func (s *ExportConversationDetailsResponseBody) SetRequestId(v string) *ExportConversationDetailsResponseBody {
	s.RequestId = &v
	return s
}

type ExportConversationDetailsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportConversationDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportConversationDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportConversationDetailsResponse) GoString() string {
	return s.String()
}

func (s *ExportConversationDetailsResponse) SetHeaders(v map[string]*string) *ExportConversationDetailsResponse {
	s.Headers = v
	return s
}

func (s *ExportConversationDetailsResponse) SetStatusCode(v int32) *ExportConversationDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportConversationDetailsResponse) SetBody(v *ExportConversationDetailsResponseBody) *ExportConversationDetailsResponse {
	s.Body = v
	return s
}

type ExportStatisticalDataRequest struct {
	// example:
	//
	// 1582266750353
	BeginTimeLeftRange *int64 `json:"BeginTimeLeftRange,omitempty" xml:"BeginTimeLeftRange,omitempty"`
	// example:
	//
	// 1582266750353
	BeginTimeRightRange *int64 `json:"BeginTimeRightRange,omitempty" xml:"BeginTimeRightRange,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// statistical
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 29b52d70-d9fe-4fe0-8476-8aaacbcfdc84
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hour
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s ExportStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *ExportStatisticalDataRequest) SetBeginTimeLeftRange(v int64) *ExportStatisticalDataRequest {
	s.BeginTimeLeftRange = &v
	return s
}

func (s *ExportStatisticalDataRequest) SetBeginTimeRightRange(v int64) *ExportStatisticalDataRequest {
	s.BeginTimeRightRange = &v
	return s
}

func (s *ExportStatisticalDataRequest) SetExportType(v string) *ExportStatisticalDataRequest {
	s.ExportType = &v
	return s
}

func (s *ExportStatisticalDataRequest) SetInstanceId(v string) *ExportStatisticalDataRequest {
	s.InstanceId = &v
	return s
}

func (s *ExportStatisticalDataRequest) SetTimeUnit(v string) *ExportStatisticalDataRequest {
	s.TimeUnit = &v
	return s
}

type ExportStatisticalDataResponseBody struct {
	// example:
	//
	// 6be5a9f1-406e-424e-a17b-b6fb86ee3cc9
	ExportTaskId *string `json:"ExportTaskId,omitempty" xml:"ExportTaskId,omitempty"`
	// example:
	//
	// c62e6789-28a8-41db-941e-171a01d3b3b9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *ExportStatisticalDataResponseBody) SetExportTaskId(v string) *ExportStatisticalDataResponseBody {
	s.ExportTaskId = &v
	return s
}

func (s *ExportStatisticalDataResponseBody) SetRequestId(v string) *ExportStatisticalDataResponseBody {
	s.RequestId = &v
	return s
}

type ExportStatisticalDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportStatisticalDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *ExportStatisticalDataResponse) SetHeaders(v map[string]*string) *ExportStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *ExportStatisticalDataResponse) SetStatusCode(v int32) *ExportStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportStatisticalDataResponse) SetBody(v *ExportStatisticalDataResponseBody) *ExportStatisticalDataResponse {
	s.Body = v
	return s
}

type GenerateUploadUrlRequest struct {
	CallerBid                      *string `json:"CallerBid,omitempty" xml:"CallerBid,omitempty"`
	CallerIp                       *string `json:"CallerIp,omitempty" xml:"CallerIp,omitempty"`
	CallerParentId                 *int64  `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	CallerType                     *string `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	CallerUid                      *int64  `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	ClientIp                       *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	Environment                    *int32  `json:"Environment,omitempty" xml:"Environment,omitempty"`
	FileName                       *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	InstanceId                     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceOwnerId                *int64  `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
	Key                            *string `json:"Key,omitempty" xml:"Key,omitempty"`
	MfaPresent                     *bool   `json:"MfaPresent,omitempty" xml:"MfaPresent,omitempty"`
	ProxyOriginalSecurityTransport *bool   `json:"ProxyOriginalSecurityTransport,omitempty" xml:"ProxyOriginalSecurityTransport,omitempty"`
	ProxyOriginalSourceIp          *string `json:"ProxyOriginalSourceIp,omitempty" xml:"ProxyOriginalSourceIp,omitempty"`
	ProxyTrustTransportInfo        *bool   `json:"ProxyTrustTransportInfo,omitempty" xml:"ProxyTrustTransportInfo,omitempty"`
	RequestId                      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityToken                  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SecurityTransport              *bool   `json:"SecurityTransport,omitempty" xml:"SecurityTransport,omitempty"`
	TenantId                       *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantName                     *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	UserId                         *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName                       *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	XspaceServicerId               *int64  `json:"XspaceServicerId,omitempty" xml:"XspaceServicerId,omitempty"`
	XspaceTenantBuId               *int64  `json:"XspaceTenantBuId,omitempty" xml:"XspaceTenantBuId,omitempty"`
}

func (s GenerateUploadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GenerateUploadUrlRequest) SetCallerBid(v string) *GenerateUploadUrlRequest {
	s.CallerBid = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetCallerIp(v string) *GenerateUploadUrlRequest {
	s.CallerIp = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetCallerParentId(v int64) *GenerateUploadUrlRequest {
	s.CallerParentId = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetCallerType(v string) *GenerateUploadUrlRequest {
	s.CallerType = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetCallerUid(v int64) *GenerateUploadUrlRequest {
	s.CallerUid = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetClientIp(v string) *GenerateUploadUrlRequest {
	s.ClientIp = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetEnvironment(v int32) *GenerateUploadUrlRequest {
	s.Environment = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetFileName(v string) *GenerateUploadUrlRequest {
	s.FileName = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetInstanceId(v string) *GenerateUploadUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetInstanceOwnerId(v int64) *GenerateUploadUrlRequest {
	s.InstanceOwnerId = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetKey(v string) *GenerateUploadUrlRequest {
	s.Key = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetMfaPresent(v bool) *GenerateUploadUrlRequest {
	s.MfaPresent = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetProxyOriginalSecurityTransport(v bool) *GenerateUploadUrlRequest {
	s.ProxyOriginalSecurityTransport = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetProxyOriginalSourceIp(v string) *GenerateUploadUrlRequest {
	s.ProxyOriginalSourceIp = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetProxyTrustTransportInfo(v bool) *GenerateUploadUrlRequest {
	s.ProxyTrustTransportInfo = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetRequestId(v string) *GenerateUploadUrlRequest {
	s.RequestId = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetSecurityToken(v string) *GenerateUploadUrlRequest {
	s.SecurityToken = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetSecurityTransport(v bool) *GenerateUploadUrlRequest {
	s.SecurityTransport = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetTenantId(v int64) *GenerateUploadUrlRequest {
	s.TenantId = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetTenantName(v string) *GenerateUploadUrlRequest {
	s.TenantName = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetUserId(v int64) *GenerateUploadUrlRequest {
	s.UserId = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetUserName(v string) *GenerateUploadUrlRequest {
	s.UserName = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetXspaceServicerId(v int64) *GenerateUploadUrlRequest {
	s.XspaceServicerId = &v
	return s
}

func (s *GenerateUploadUrlRequest) SetXspaceTenantBuId(v int64) *GenerateUploadUrlRequest {
	s.XspaceTenantBuId = &v
	return s
}

type GenerateUploadUrlResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GenerateUploadUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateUploadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUploadUrlResponseBody) SetCode(v string) *GenerateUploadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateUploadUrlResponseBody) SetData(v *GenerateUploadUrlResponseBodyData) *GenerateUploadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GenerateUploadUrlResponseBody) SetHttpStatusCode(v int32) *GenerateUploadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateUploadUrlResponseBody) SetMessage(v string) *GenerateUploadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateUploadUrlResponseBody) SetRequestId(v string) *GenerateUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUploadUrlResponseBody) SetSuccess(v bool) *GenerateUploadUrlResponseBody {
	s.Success = &v
	return s
}

type GenerateUploadUrlResponseBodyData struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Expire    *int32  `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Folder    *string `json:"Folder,omitempty" xml:"Folder,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateUploadUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateUploadUrlResponseBodyData) SetAccessId(v string) *GenerateUploadUrlResponseBodyData {
	s.AccessId = &v
	return s
}

func (s *GenerateUploadUrlResponseBodyData) SetExpire(v int32) *GenerateUploadUrlResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GenerateUploadUrlResponseBodyData) SetFolder(v string) *GenerateUploadUrlResponseBodyData {
	s.Folder = &v
	return s
}

func (s *GenerateUploadUrlResponseBodyData) SetHost(v string) *GenerateUploadUrlResponseBodyData {
	s.Host = &v
	return s
}

func (s *GenerateUploadUrlResponseBodyData) SetMessage(v string) *GenerateUploadUrlResponseBodyData {
	s.Message = &v
	return s
}

func (s *GenerateUploadUrlResponseBodyData) SetPolicy(v string) *GenerateUploadUrlResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GenerateUploadUrlResponseBodyData) SetSignature(v string) *GenerateUploadUrlResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GenerateUploadUrlResponseBodyData) SetSuccess(v bool) *GenerateUploadUrlResponseBodyData {
	s.Success = &v
	return s
}

type GenerateUploadUrlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUploadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GenerateUploadUrlResponse) SetHeaders(v map[string]*string) *GenerateUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GenerateUploadUrlResponse) SetStatusCode(v int32) *GenerateUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUploadUrlResponse) SetBody(v *GenerateUploadUrlResponseBody) *GenerateUploadUrlResponse {
	s.Body = v
	return s
}

type GetAsrConfigRequest struct {
	// example:
	//
	// 1
	ConfigLevel *int32 `json:"ConfigLevel,omitempty" xml:"ConfigLevel,omitempty"`
	// example:
	//
	// 3b1d3031-8b6e-460a-8640-d330f2ca50b8
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
}

func (s GetAsrConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsrConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAsrConfigRequest) SetConfigLevel(v int32) *GetAsrConfigRequest {
	s.ConfigLevel = &v
	return s
}

func (s *GetAsrConfigRequest) SetEntryId(v string) *GetAsrConfigRequest {
	s.EntryId = &v
	return s
}

type GetAsrConfigResponseBody struct {
	// example:
	//
	// OK
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAsrConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Not Found
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAsrConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsrConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsrConfigResponseBody) SetCode(v string) *GetAsrConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsrConfigResponseBody) SetData(v *GetAsrConfigResponseBodyData) *GetAsrConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetAsrConfigResponseBody) SetErrorMsg(v string) *GetAsrConfigResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetAsrConfigResponseBody) SetHttpStatusCode(v int32) *GetAsrConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAsrConfigResponseBody) SetRequestId(v string) *GetAsrConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsrConfigResponseBody) SetSuccess(v bool) *GetAsrConfigResponseBody {
	s.Success = &v
	return s
}

type GetAsrConfigResponseBodyData struct {
	// example:
	//
	// 3b1d3031-8b6e-460a-8640-d330f2ca50b8
	AsrAcousticModelId *string `json:"AsrAcousticModelId,omitempty" xml:"AsrAcousticModelId,omitempty"`
	// example:
	//
	// 3b1d3031-8b6e-460a-8640-d330f2ca50b8
	AsrClassVocabularyId *string `json:"AsrClassVocabularyId,omitempty" xml:"AsrClassVocabularyId,omitempty"`
	// example:
	//
	// 3b1d3031-8b6e-460a-8640-d330f2ca50b8
	AsrCustomizationId *string `json:"AsrCustomizationId,omitempty" xml:"AsrCustomizationId,omitempty"`
	// example:
	//
	// 3b1d3031-8b6e-460a-8640-d330f2ca50b8
	AsrVocabularyId *string `json:"AsrVocabularyId,omitempty" xml:"AsrVocabularyId,omitempty"`
}

func (s GetAsrConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsrConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsrConfigResponseBodyData) SetAsrAcousticModelId(v string) *GetAsrConfigResponseBodyData {
	s.AsrAcousticModelId = &v
	return s
}

func (s *GetAsrConfigResponseBodyData) SetAsrClassVocabularyId(v string) *GetAsrConfigResponseBodyData {
	s.AsrClassVocabularyId = &v
	return s
}

func (s *GetAsrConfigResponseBodyData) SetAsrCustomizationId(v string) *GetAsrConfigResponseBodyData {
	s.AsrCustomizationId = &v
	return s
}

func (s *GetAsrConfigResponseBodyData) SetAsrVocabularyId(v string) *GetAsrConfigResponseBodyData {
	s.AsrVocabularyId = &v
	return s
}

type GetAsrConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsrConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsrConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsrConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAsrConfigResponse) SetHeaders(v map[string]*string) *GetAsrConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAsrConfigResponse) SetStatusCode(v int32) *GetAsrConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsrConfigResponse) SetBody(v *GetAsrConfigResponseBody) *GetAsrConfigResponse {
	s.Body = v
	return s
}

type GetRealTimeConcurrencyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c112c168ed664c0a851f9ca72d2f7999
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRealTimeConcurrencyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRealTimeConcurrencyRequest) GoString() string {
	return s.String()
}

func (s *GetRealTimeConcurrencyRequest) SetInstanceId(v string) *GetRealTimeConcurrencyRequest {
	s.InstanceId = &v
	return s
}

type GetRealTimeConcurrencyResponseBody struct {
	// example:
	//
	// 2
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// example:
	//
	// 1
	RealTimeConcurrency *int64 `json:"RealTimeConcurrency,omitempty" xml:"RealTimeConcurrency,omitempty"`
	// example:
	//
	// E6E61E1A-D2DC-5ACF-AED4-A115B6691F98
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1661584255029
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetRealTimeConcurrencyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRealTimeConcurrencyResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealTimeConcurrencyResponseBody) SetMaxConcurrency(v int64) *GetRealTimeConcurrencyResponseBody {
	s.MaxConcurrency = &v
	return s
}

func (s *GetRealTimeConcurrencyResponseBody) SetRealTimeConcurrency(v int64) *GetRealTimeConcurrencyResponseBody {
	s.RealTimeConcurrency = &v
	return s
}

func (s *GetRealTimeConcurrencyResponseBody) SetRequestId(v string) *GetRealTimeConcurrencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRealTimeConcurrencyResponseBody) SetTimestamp(v int64) *GetRealTimeConcurrencyResponseBody {
	s.Timestamp = &v
	return s
}

type GetRealTimeConcurrencyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRealTimeConcurrencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRealTimeConcurrencyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRealTimeConcurrencyResponse) GoString() string {
	return s.String()
}

func (s *GetRealTimeConcurrencyResponse) SetHeaders(v map[string]*string) *GetRealTimeConcurrencyResponse {
	s.Headers = v
	return s
}

func (s *GetRealTimeConcurrencyResponse) SetStatusCode(v int32) *GetRealTimeConcurrencyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealTimeConcurrencyResponse) SetBody(v *GetRealTimeConcurrencyResponseBody) *GetRealTimeConcurrencyResponse {
	s.Body = v
	return s
}

type ListChatbotInstancesRequest struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NluServiceParamsJson *string `json:"NluServiceParamsJson,omitempty" xml:"NluServiceParamsJson,omitempty"`
	NluServiceType       *string `json:"NluServiceType,omitempty" xml:"NluServiceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UnionSource *string `json:"UnionSource,omitempty" xml:"UnionSource,omitempty"`
}

func (s ListChatbotInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChatbotInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListChatbotInstancesRequest) SetInstanceId(v string) *ListChatbotInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListChatbotInstancesRequest) SetNluServiceParamsJson(v string) *ListChatbotInstancesRequest {
	s.NluServiceParamsJson = &v
	return s
}

func (s *ListChatbotInstancesRequest) SetNluServiceType(v string) *ListChatbotInstancesRequest {
	s.NluServiceType = &v
	return s
}

func (s *ListChatbotInstancesRequest) SetPageNumber(v int32) *ListChatbotInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListChatbotInstancesRequest) SetPageSize(v int32) *ListChatbotInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListChatbotInstancesRequest) SetUnionSource(v string) *ListChatbotInstancesRequest {
	s.UnionSource = &v
	return s
}

type ListChatbotInstancesResponseBody struct {
	Bots []*ListChatbotInstancesResponseBodyBots `json:"Bots,omitempty" xml:"Bots,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChatbotInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChatbotInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatbotInstancesResponseBody) SetBots(v []*ListChatbotInstancesResponseBodyBots) *ListChatbotInstancesResponseBody {
	s.Bots = v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetPageNumber(v int32) *ListChatbotInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetPageSize(v int64) *ListChatbotInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetRequestId(v string) *ListChatbotInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetTotalCount(v int64) *ListChatbotInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListChatbotInstancesResponseBodyBots struct {
	// example:
	//
	// https://dss0.ali.com/70cFuHS.jpg
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// example:
	//
	// 1582266750353
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ‘’
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// example:
	//
	// zh-cn
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// UTC+8
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s ListChatbotInstancesResponseBodyBots) String() string {
	return tea.Prettify(s)
}

func (s ListChatbotInstancesResponseBodyBots) GoString() string {
	return s.String()
}

func (s *ListChatbotInstancesResponseBodyBots) SetAvatar(v string) *ListChatbotInstancesResponseBodyBots {
	s.Avatar = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) SetCreateTime(v string) *ListChatbotInstancesResponseBodyBots {
	s.CreateTime = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) SetInstanceId(v string) *ListChatbotInstancesResponseBodyBots {
	s.InstanceId = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) SetIntroduction(v string) *ListChatbotInstancesResponseBodyBots {
	s.Introduction = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) SetLanguageCode(v string) *ListChatbotInstancesResponseBodyBots {
	s.LanguageCode = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) SetName(v string) *ListChatbotInstancesResponseBodyBots {
	s.Name = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) SetTimeZone(v string) *ListChatbotInstancesResponseBodyBots {
	s.TimeZone = &v
	return s
}

type ListChatbotInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChatbotInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChatbotInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChatbotInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListChatbotInstancesResponse) SetHeaders(v map[string]*string) *ListChatbotInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListChatbotInstancesResponse) SetStatusCode(v int32) *ListChatbotInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChatbotInstancesResponse) SetBody(v *ListChatbotInstancesResponseBody) *ListChatbotInstancesResponse {
	s.Body = v
	return s
}

type ListConversationDetailsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a2c26e67-5984-4935-984e-bcee52971993
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 82b2eaae-ce5c-45f8-8231-f15b6b27e55c
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListConversationDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConversationDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListConversationDetailsRequest) SetConversationId(v string) *ListConversationDetailsRequest {
	s.ConversationId = &v
	return s
}

func (s *ListConversationDetailsRequest) SetInstanceId(v string) *ListConversationDetailsRequest {
	s.InstanceId = &v
	return s
}

type ListConversationDetailsResponseBody struct {
	ConversationDetails []*ListConversationDetailsResponseBodyConversationDetails `json:"ConversationDetails,omitempty" xml:"ConversationDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConversationDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConversationDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConversationDetailsResponseBody) SetConversationDetails(v []*ListConversationDetailsResponseBodyConversationDetails) *ListConversationDetailsResponseBody {
	s.ConversationDetails = v
	return s
}

func (s *ListConversationDetailsResponseBody) SetRequestId(v string) *ListConversationDetailsResponseBody {
	s.RequestId = &v
	return s
}

type ListConversationDetailsResponseBodyConversationDetails struct {
	// example:
	//
	// Dialogue
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// {}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// 1582266750353
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	SequenceId *string `json:"SequenceId,omitempty" xml:"SequenceId,omitempty"`
	// example:
	//
	// Chatbot
	Speaker   *string `json:"Speaker,omitempty" xml:"Speaker,omitempty"`
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s ListConversationDetailsResponseBodyConversationDetails) String() string {
	return tea.Prettify(s)
}

func (s ListConversationDetailsResponseBodyConversationDetails) GoString() string {
	return s.String()
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetAction(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.Action = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetActionParams(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.ActionParams = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetConversationId(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.ConversationId = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetCreateTime(v int64) *ListConversationDetailsResponseBodyConversationDetails {
	s.CreateTime = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetSequenceId(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.SequenceId = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetSpeaker(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.Speaker = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetUtterance(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.Utterance = &v
	return s
}

type ListConversationDetailsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConversationDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConversationDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConversationDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListConversationDetailsResponse) SetHeaders(v map[string]*string) *ListConversationDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListConversationDetailsResponse) SetStatusCode(v int32) *ListConversationDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConversationDetailsResponse) SetBody(v *ListConversationDetailsResponseBody) *ListConversationDetailsResponse {
	s.Body = v
	return s
}

type ListConversationsRequest struct {
	// example:
	//
	// 1638288000000
	BeginTimeLeftRange *int64 `json:"BeginTimeLeftRange,omitempty" xml:"BeginTimeLeftRange,omitempty"`
	// example:
	//
	// 1637547875311
	BeginTimeRightRange *int64 `json:"BeginTimeRightRange,omitempty" xml:"BeginTimeRightRange,omitempty"`
	// example:
	//
	// 138106*****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// false
	IsSandBox *string `json:"IsSandBox,omitempty" xml:"IsSandBox,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 13788914724
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 0
	Result           *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
	RoundsLeftRange  *int32 `json:"RoundsLeftRange,omitempty" xml:"RoundsLeftRange,omitempty"`
	RoundsRightRange *int32 `json:"RoundsRightRange,omitempty" xml:"RoundsRightRange,omitempty"`
}

func (s ListConversationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConversationsRequest) GoString() string {
	return s.String()
}

func (s *ListConversationsRequest) SetBeginTimeLeftRange(v int64) *ListConversationsRequest {
	s.BeginTimeLeftRange = &v
	return s
}

func (s *ListConversationsRequest) SetBeginTimeRightRange(v int64) *ListConversationsRequest {
	s.BeginTimeRightRange = &v
	return s
}

func (s *ListConversationsRequest) SetCallingNumber(v string) *ListConversationsRequest {
	s.CallingNumber = &v
	return s
}

func (s *ListConversationsRequest) SetInstanceId(v string) *ListConversationsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListConversationsRequest) SetIsSandBox(v string) *ListConversationsRequest {
	s.IsSandBox = &v
	return s
}

func (s *ListConversationsRequest) SetPageNumber(v int32) *ListConversationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConversationsRequest) SetPageSize(v int32) *ListConversationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListConversationsRequest) SetQuery(v string) *ListConversationsRequest {
	s.Query = &v
	return s
}

func (s *ListConversationsRequest) SetResult(v int64) *ListConversationsRequest {
	s.Result = &v
	return s
}

func (s *ListConversationsRequest) SetRoundsLeftRange(v int32) *ListConversationsRequest {
	s.RoundsLeftRange = &v
	return s
}

func (s *ListConversationsRequest) SetRoundsRightRange(v int32) *ListConversationsRequest {
	s.RoundsRightRange = &v
	return s
}

type ListConversationsResponseBody struct {
	Conversations []*ListConversationsResponseBodyConversations `json:"Conversations,omitempty" xml:"Conversations,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// a2c26e67-5984-4935-984e-bcee52971993
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListConversationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConversationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConversationsResponseBody) SetConversations(v []*ListConversationsResponseBodyConversations) *ListConversationsResponseBody {
	s.Conversations = v
	return s
}

func (s *ListConversationsResponseBody) SetPageNumber(v int32) *ListConversationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListConversationsResponseBody) SetPageSize(v int32) *ListConversationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListConversationsResponseBody) SetRequestId(v string) *ListConversationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConversationsResponseBody) SetTotalCount(v int64) *ListConversationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListConversationsResponseBodyConversations struct {
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// example:
	//
	// 135815884***
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// example:
	//
	// 82b2eaae-ce5c-45f8-8231-f15b6b27e55c
	ConversationId *string   `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	DsReport       *string   `json:"DsReport,omitempty" xml:"DsReport,omitempty"`
	DsReportTitles []*string `json:"DsReportTitles,omitempty" xml:"DsReportTitles,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	EndReason *int32 `json:"EndReason,omitempty" xml:"EndReason,omitempty"`
	// example:
	//
	// 1582266750353
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// true
	HasLastPlaybackCompleted *bool `json:"HasLastPlaybackCompleted,omitempty" xml:"HasLastPlaybackCompleted,omitempty"`
	HasToAgent               *bool `json:"HasToAgent,omitempty" xml:"HasToAgent,omitempty"`
	// example:
	//
	// 2
	Rounds *int32 `json:"Rounds,omitempty" xml:"Rounds,omitempty"`
	// example:
	//
	// true
	SandBox    *bool   `json:"SandBox,omitempty" xml:"SandBox,omitempty"`
	SkillGroup *string `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
	// example:
	//
	// 1641625694286
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListConversationsResponseBodyConversations) String() string {
	return tea.Prettify(s)
}

func (s ListConversationsResponseBodyConversations) GoString() string {
	return s.String()
}

func (s *ListConversationsResponseBodyConversations) SetCalledNumber(v string) *ListConversationsResponseBodyConversations {
	s.CalledNumber = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetCallingNumber(v string) *ListConversationsResponseBodyConversations {
	s.CallingNumber = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetConversationId(v string) *ListConversationsResponseBodyConversations {
	s.ConversationId = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetDsReport(v string) *ListConversationsResponseBodyConversations {
	s.DsReport = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetDsReportTitles(v []*string) *ListConversationsResponseBodyConversations {
	s.DsReportTitles = v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetEndReason(v int32) *ListConversationsResponseBodyConversations {
	s.EndReason = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetEndTime(v int64) *ListConversationsResponseBodyConversations {
	s.EndTime = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetHasLastPlaybackCompleted(v bool) *ListConversationsResponseBodyConversations {
	s.HasLastPlaybackCompleted = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetHasToAgent(v bool) *ListConversationsResponseBodyConversations {
	s.HasToAgent = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetRounds(v int32) *ListConversationsResponseBodyConversations {
	s.Rounds = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetSandBox(v bool) *ListConversationsResponseBodyConversations {
	s.SandBox = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetSkillGroup(v string) *ListConversationsResponseBodyConversations {
	s.SkillGroup = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetStartTime(v int64) *ListConversationsResponseBodyConversations {
	s.StartTime = &v
	return s
}

type ListConversationsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConversationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConversationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConversationsResponse) GoString() string {
	return s.String()
}

func (s *ListConversationsResponse) SetHeaders(v map[string]*string) *ListConversationsResponse {
	s.Headers = v
	return s
}

func (s *ListConversationsResponse) SetStatusCode(v int32) *ListConversationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConversationsResponse) SetBody(v *ListConversationsResponseBody) *ListConversationsResponse {
	s.Body = v
	return s
}

type ListDownloadTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDownloadTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDownloadTasksRequest) GoString() string {
	return s.String()
}

func (s *ListDownloadTasksRequest) SetPageNumber(v int32) *ListDownloadTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDownloadTasksRequest) SetPageSize(v int32) *ListDownloadTasksRequest {
	s.PageSize = &v
	return s
}

type ListDownloadTasksResponseBody struct {
	// example:
	//
	// OK
	Code          *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	DownloadTasks *ListDownloadTasksResponseBodyDownloadTasks `json:"DownloadTasks,omitempty" xml:"DownloadTasks,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D24E0148-6D40-550E-9471-B2C5A34C3D12
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDownloadTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDownloadTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDownloadTasksResponseBody) SetCode(v string) *ListDownloadTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListDownloadTasksResponseBody) SetDownloadTasks(v *ListDownloadTasksResponseBodyDownloadTasks) *ListDownloadTasksResponseBody {
	s.DownloadTasks = v
	return s
}

func (s *ListDownloadTasksResponseBody) SetHttpStatusCode(v int32) *ListDownloadTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDownloadTasksResponseBody) SetMessage(v string) *ListDownloadTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListDownloadTasksResponseBody) SetRequestId(v string) *ListDownloadTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDownloadTasksResponseBody) SetSuccess(v bool) *ListDownloadTasksResponseBody {
	s.Success = &v
	return s
}

type ListDownloadTasksResponseBodyDownloadTasks struct {
	List []*ListDownloadTasksResponseBodyDownloadTasksList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDownloadTasksResponseBodyDownloadTasks) String() string {
	return tea.Prettify(s)
}

func (s ListDownloadTasksResponseBodyDownloadTasks) GoString() string {
	return s.String()
}

func (s *ListDownloadTasksResponseBodyDownloadTasks) SetList(v []*ListDownloadTasksResponseBodyDownloadTasksList) *ListDownloadTasksResponseBodyDownloadTasks {
	s.List = v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasks) SetPageNumber(v int32) *ListDownloadTasksResponseBodyDownloadTasks {
	s.PageNumber = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasks) SetPageSize(v int32) *ListDownloadTasksResponseBodyDownloadTasks {
	s.PageSize = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasks) SetTotalCount(v int32) *ListDownloadTasksResponseBodyDownloadTasks {
	s.TotalCount = &v
	return s
}

type ListDownloadTasksResponseBodyDownloadTasksList struct {
	DownloadTaskFiles []*ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles `json:"DownloadTaskFiles,omitempty" xml:"DownloadTaskFiles,omitempty" type:"Repeated"`
	// example:
	//
	// 1637119221702
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// Expired
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cb9aba69-f578-42b2-aa2f-3e5a41947db8
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListDownloadTasksResponseBodyDownloadTasksList) String() string {
	return tea.Prettify(s)
}

func (s ListDownloadTasksResponseBodyDownloadTasksList) GoString() string {
	return s.String()
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) SetDownloadTaskFiles(v []*ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) *ListDownloadTasksResponseBodyDownloadTasksList {
	s.DownloadTaskFiles = v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) SetExpireTime(v int64) *ListDownloadTasksResponseBodyDownloadTasksList {
	s.ExpireTime = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) SetStatus(v string) *ListDownloadTasksResponseBodyDownloadTasksList {
	s.Status = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) SetTaskId(v string) *ListDownloadTasksResponseBodyDownloadTasksList {
	s.TaskId = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksList) SetTitle(v string) *ListDownloadTasksResponseBodyDownloadTasksList {
	s.Title = &v
	return s
}

type ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles struct {
	// example:
	//
	// c32bf5675b704dc5b19200a89d2e85f1
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// 70
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// Published
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) String() string {
	return tea.Prettify(s)
}

func (s ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) GoString() string {
	return s.String()
}

func (s *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) SetFileId(v string) *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles {
	s.FileId = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) SetProgress(v int32) *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles {
	s.Progress = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) SetStatus(v string) *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles {
	s.Status = &v
	return s
}

func (s *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles) SetTitle(v string) *ListDownloadTasksResponseBodyDownloadTasksListDownloadTaskFiles {
	s.Title = &v
	return s
}

type ListDownloadTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDownloadTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDownloadTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDownloadTasksResponse) GoString() string {
	return s.String()
}

func (s *ListDownloadTasksResponse) SetHeaders(v map[string]*string) *ListDownloadTasksResponse {
	s.Headers = v
	return s
}

func (s *ListDownloadTasksResponse) SetStatusCode(v int32) *ListDownloadTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDownloadTasksResponse) SetBody(v *ListDownloadTasksResponseBody) *ListDownloadTasksResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	InstanceIdListJsonString *string `json:"InstanceIdListJsonString,omitempty" xml:"InstanceIdListJsonString,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// [Managed]
	NluServiceTypeListJsonString *string `json:"NluServiceTypeListJsonString,omitempty" xml:"NluServiceTypeListJsonString,omitempty"`
	Number                       *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UnionInstanceId *string `json:"UnionInstanceId,omitempty" xml:"UnionInstanceId,omitempty"`
	UnionSource     *string `json:"UnionSource,omitempty" xml:"UnionSource,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetInstanceIdListJsonString(v string) *ListInstancesRequest {
	s.InstanceIdListJsonString = &v
	return s
}

func (s *ListInstancesRequest) SetName(v string) *ListInstancesRequest {
	s.Name = &v
	return s
}

func (s *ListInstancesRequest) SetNluServiceTypeListJsonString(v string) *ListInstancesRequest {
	s.NluServiceTypeListJsonString = &v
	return s
}

func (s *ListInstancesRequest) SetNumber(v string) *ListInstancesRequest {
	s.Number = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesRequest) SetUnionInstanceId(v string) *ListInstancesRequest {
	s.UnionInstanceId = &v
	return s
}

func (s *ListInstancesRequest) SetUnionSource(v string) *ListInstancesRequest {
	s.UnionSource = &v
	return s
}

type ListInstancesResponseBody struct {
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// A8AED3C8-F57B-5D71-9A34-4A170287533F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetPageNumber(v int32) *ListInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesResponseBody) SetPageSize(v int32) *ListInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstancesResponseBodyInstances struct {
	ApplicableOperations []*string `json:"ApplicableOperations,omitempty" xml:"ApplicableOperations,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Concurrency *int64  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// dc437bba-5a25-4bbc-b4c2-f268864bebb5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1582266750353
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// xxx
	ModifyUserName       *string   `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Name                 *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NluServiceParamsJson *string   `json:"NluServiceParamsJson,omitempty" xml:"NluServiceParamsJson,omitempty"`
	Numbers              []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	// example:
	//
	// Published
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UnionInstanceId *string `json:"UnionInstanceId,omitempty" xml:"UnionInstanceId,omitempty"`
	UnionSource     *string `json:"UnionSource,omitempty" xml:"UnionSource,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) SetApplicableOperations(v []*string) *ListInstancesResponseBodyInstances {
	s.ApplicableOperations = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetConcurrency(v int64) *ListInstancesResponseBodyInstances {
	s.Concurrency = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCreateTime(v int64) *ListInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDescription(v string) *ListInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetModifyTime(v int64) *ListInstancesResponseBodyInstances {
	s.ModifyTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetModifyUserName(v string) *ListInstancesResponseBodyInstances {
	s.ModifyUserName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetName(v string) *ListInstancesResponseBodyInstances {
	s.Name = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetNluServiceParamsJson(v string) *ListInstancesResponseBodyInstances {
	s.NluServiceParamsJson = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetNumbers(v []*string) *ListInstancesResponseBodyInstances {
	s.Numbers = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUnionInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.UnionInstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUnionSource(v string) *ListInstancesResponseBodyInstances {
	s.UnionSource = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ModifyAsrConfigRequest struct {
	// example:
	//
	// 6cc9f5ca-2cb6-4cc7-a46b-2bbfd3e61b22
	AsrAcousticModelId *string `json:"AsrAcousticModelId,omitempty" xml:"AsrAcousticModelId,omitempty"`
	// example:
	//
	// 6cc9f5ca-2cb6-4cc7-a46b-2bbfd3e61b22
	AsrClassVocabularyId *string `json:"AsrClassVocabularyId,omitempty" xml:"AsrClassVocabularyId,omitempty"`
	// example:
	//
	// 6cc9f5ca-2cb6-4cc7-a46b-2bbfd3e61b22
	AsrCustomizationId *string `json:"AsrCustomizationId,omitempty" xml:"AsrCustomizationId,omitempty"`
	// example:
	//
	// 6cc9f5ca-2cb6-4cc7-a46b-2bbfd3e61b22
	AsrVocabularyId *string `json:"AsrVocabularyId,omitempty" xml:"AsrVocabularyId,omitempty"`
	// example:
	//
	// 0
	ConfigLevel *int32 `json:"ConfigLevel,omitempty" xml:"ConfigLevel,omitempty"`
	// example:
	//
	// 6cc9f5ca-2cb6-4cc7-a46b-2bbfd3e61b22
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
}

func (s ModifyAsrConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAsrConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAsrConfigRequest) SetAsrAcousticModelId(v string) *ModifyAsrConfigRequest {
	s.AsrAcousticModelId = &v
	return s
}

func (s *ModifyAsrConfigRequest) SetAsrClassVocabularyId(v string) *ModifyAsrConfigRequest {
	s.AsrClassVocabularyId = &v
	return s
}

func (s *ModifyAsrConfigRequest) SetAsrCustomizationId(v string) *ModifyAsrConfigRequest {
	s.AsrCustomizationId = &v
	return s
}

func (s *ModifyAsrConfigRequest) SetAsrVocabularyId(v string) *ModifyAsrConfigRequest {
	s.AsrVocabularyId = &v
	return s
}

func (s *ModifyAsrConfigRequest) SetConfigLevel(v int32) *ModifyAsrConfigRequest {
	s.ConfigLevel = &v
	return s
}

func (s *ModifyAsrConfigRequest) SetEntryId(v string) *ModifyAsrConfigRequest {
	s.EntryId = &v
	return s
}

type ModifyAsrConfigResponseBody struct {
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ModifyAsrConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Not Found
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAsrConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAsrConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAsrConfigResponseBody) SetCode(v string) *ModifyAsrConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAsrConfigResponseBody) SetData(v *ModifyAsrConfigResponseBodyData) *ModifyAsrConfigResponseBody {
	s.Data = v
	return s
}

func (s *ModifyAsrConfigResponseBody) SetErrorMsg(v string) *ModifyAsrConfigResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ModifyAsrConfigResponseBody) SetHttpStatusCode(v int32) *ModifyAsrConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyAsrConfigResponseBody) SetRequestId(v string) *ModifyAsrConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAsrConfigResponseBody) SetSuccess(v bool) *ModifyAsrConfigResponseBody {
	s.Success = &v
	return s
}

type ModifyAsrConfigResponseBodyData struct {
	AffectedRows *int32 `json:"AffectedRows,omitempty" xml:"AffectedRows,omitempty"`
}

func (s ModifyAsrConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyAsrConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyAsrConfigResponseBodyData) SetAffectedRows(v int32) *ModifyAsrConfigResponseBodyData {
	s.AffectedRows = &v
	return s
}

type ModifyAsrConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAsrConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAsrConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAsrConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyAsrConfigResponse) SetHeaders(v map[string]*string) *ModifyAsrConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyAsrConfigResponse) SetStatusCode(v int32) *ModifyAsrConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAsrConfigResponse) SetBody(v *ModifyAsrConfigResponseBody) *ModifyAsrConfigResponse {
	s.Body = v
	return s
}

type ModifyGreetingConfigRequest struct {
	// This parameter is required.
	GreetingWords *string `json:"GreetingWords,omitempty" xml:"GreetingWords,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12f407b22cbe4890ac595f09985848d5
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentTrigger *string `json:"IntentTrigger,omitempty" xml:"IntentTrigger,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbotIntent
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ModifyGreetingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGreetingConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyGreetingConfigRequest) SetGreetingWords(v string) *ModifyGreetingConfigRequest {
	s.GreetingWords = &v
	return s
}

func (s *ModifyGreetingConfigRequest) SetInstanceId(v string) *ModifyGreetingConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyGreetingConfigRequest) SetIntentTrigger(v string) *ModifyGreetingConfigRequest {
	s.IntentTrigger = &v
	return s
}

func (s *ModifyGreetingConfigRequest) SetSourceType(v string) *ModifyGreetingConfigRequest {
	s.SourceType = &v
	return s
}

type ModifyGreetingConfigResponseBody struct {
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGreetingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyGreetingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGreetingConfigResponseBody) SetRequestId(v string) *ModifyGreetingConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyGreetingConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGreetingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGreetingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGreetingConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyGreetingConfigResponse) SetHeaders(v map[string]*string) *ModifyGreetingConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyGreetingConfigResponse) SetStatusCode(v int32) *ModifyGreetingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGreetingConfigResponse) SetBody(v *ModifyGreetingConfigResponseBody) *ModifyGreetingConfigResponse {
	s.Body = v
	return s
}

type ModifyInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Concurrency *int64  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12f407b22cbe4890ac595f09985848d5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRequest) SetConcurrency(v int64) *ModifyInstanceRequest {
	s.Concurrency = &v
	return s
}

func (s *ModifyInstanceRequest) SetDescription(v string) *ModifyInstanceRequest {
	s.Description = &v
	return s
}

func (s *ModifyInstanceRequest) SetInstanceId(v string) *ModifyInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRequest) SetName(v string) *ModifyInstanceRequest {
	s.Name = &v
	return s
}

type ModifyInstanceResponseBody struct {
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponseBody) SetRequestId(v string) *ModifyInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponse) SetHeaders(v map[string]*string) *ModifyInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceResponse) SetStatusCode(v int32) *ModifyInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceResponse) SetBody(v *ModifyInstanceResponseBody) *ModifyInstanceResponse {
	s.Body = v
	return s
}

type ModifySilenceTimeoutConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TransferToAgent
	FinalAction *string `json:"FinalAction,omitempty" xml:"FinalAction,omitempty"`
	// example:
	//
	// {}
	FinalActionParams *string `json:"FinalActionParams,omitempty" xml:"FinalActionParams,omitempty"`
	// This parameter is required.
	FinalPrompt *string `json:"FinalPrompt,omitempty" xml:"FinalPrompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12f407b22cbe4890ac595f09985848d5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Timeout
	IntentTrigger *string `json:"IntentTrigger,omitempty" xml:"IntentTrigger,omitempty"`
	// This parameter is required.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbotIntent
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ModifySilenceTimeoutConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySilenceTimeoutConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifySilenceTimeoutConfigRequest) SetFinalAction(v string) *ModifySilenceTimeoutConfigRequest {
	s.FinalAction = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetFinalActionParams(v string) *ModifySilenceTimeoutConfigRequest {
	s.FinalActionParams = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetFinalPrompt(v string) *ModifySilenceTimeoutConfigRequest {
	s.FinalPrompt = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetInstanceId(v string) *ModifySilenceTimeoutConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetIntentTrigger(v string) *ModifySilenceTimeoutConfigRequest {
	s.IntentTrigger = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetPrompt(v string) *ModifySilenceTimeoutConfigRequest {
	s.Prompt = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetSourceType(v string) *ModifySilenceTimeoutConfigRequest {
	s.SourceType = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetThreshold(v int32) *ModifySilenceTimeoutConfigRequest {
	s.Threshold = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetTimeout(v int64) *ModifySilenceTimeoutConfigRequest {
	s.Timeout = &v
	return s
}

type ModifySilenceTimeoutConfigResponseBody struct {
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySilenceTimeoutConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySilenceTimeoutConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySilenceTimeoutConfigResponseBody) SetRequestId(v string) *ModifySilenceTimeoutConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifySilenceTimeoutConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySilenceTimeoutConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySilenceTimeoutConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySilenceTimeoutConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifySilenceTimeoutConfigResponse) SetHeaders(v map[string]*string) *ModifySilenceTimeoutConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifySilenceTimeoutConfigResponse) SetStatusCode(v int32) *ModifySilenceTimeoutConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySilenceTimeoutConfigResponse) SetBody(v *ModifySilenceTimeoutConfigResponseBody) *ModifySilenceTimeoutConfigResponse {
	s.Body = v
	return s
}

type ModifyTTSConfigRequest struct {
	AliCustomizedVoice *string `json:"AliCustomizedVoice,omitempty" xml:"AliCustomizedVoice,omitempty"`
	AppKey             *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Engine             *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineXunfei       *string `json:"EngineXunfei,omitempty" xml:"EngineXunfei,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12f407b22cbe4890ac595f09985848d5
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NlsServiceType *string `json:"NlsServiceType,omitempty" xml:"NlsServiceType,omitempty"`
	// example:
	//
	// 100
	SpeechRate *string `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// example:
	//
	// aixia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 10
	Volume *string `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s ModifyTTSConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTTSConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyTTSConfigRequest) SetAliCustomizedVoice(v string) *ModifyTTSConfigRequest {
	s.AliCustomizedVoice = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetAppKey(v string) *ModifyTTSConfigRequest {
	s.AppKey = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetEngine(v string) *ModifyTTSConfigRequest {
	s.Engine = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetEngineXunfei(v string) *ModifyTTSConfigRequest {
	s.EngineXunfei = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetInstanceId(v string) *ModifyTTSConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetNlsServiceType(v string) *ModifyTTSConfigRequest {
	s.NlsServiceType = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetSpeechRate(v string) *ModifyTTSConfigRequest {
	s.SpeechRate = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetVoice(v string) *ModifyTTSConfigRequest {
	s.Voice = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetVolume(v string) *ModifyTTSConfigRequest {
	s.Volume = &v
	return s
}

type ModifyTTSConfigResponseBody struct {
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTTSConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTTSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTTSConfigResponseBody) SetRequestId(v string) *ModifyTTSConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTTSConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTTSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTTSConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTTSConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyTTSConfigResponse) SetHeaders(v map[string]*string) *ModifyTTSConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyTTSConfigResponse) SetStatusCode(v int32) *ModifyTTSConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTTSConfigResponse) SetBody(v *ModifyTTSConfigResponseBody) *ModifyTTSConfigResponse {
	s.Body = v
	return s
}

type ModifyUnrecognizingConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TransferToAgent
	FinalAction *string `json:"FinalAction,omitempty" xml:"FinalAction,omitempty"`
	// example:
	//
	// {}
	FinalActionParams *string `json:"FinalActionParams,omitempty" xml:"FinalActionParams,omitempty"`
	// This parameter is required.
	FinalPrompt *string `json:"FinalPrompt,omitempty" xml:"FinalPrompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12f407b22cbe4890ac595f09985848d5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifyUnrecognizingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUnrecognizingConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyUnrecognizingConfigRequest) SetFinalAction(v string) *ModifyUnrecognizingConfigRequest {
	s.FinalAction = &v
	return s
}

func (s *ModifyUnrecognizingConfigRequest) SetFinalActionParams(v string) *ModifyUnrecognizingConfigRequest {
	s.FinalActionParams = &v
	return s
}

func (s *ModifyUnrecognizingConfigRequest) SetFinalPrompt(v string) *ModifyUnrecognizingConfigRequest {
	s.FinalPrompt = &v
	return s
}

func (s *ModifyUnrecognizingConfigRequest) SetInstanceId(v string) *ModifyUnrecognizingConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUnrecognizingConfigRequest) SetPrompt(v string) *ModifyUnrecognizingConfigRequest {
	s.Prompt = &v
	return s
}

func (s *ModifyUnrecognizingConfigRequest) SetThreshold(v int32) *ModifyUnrecognizingConfigRequest {
	s.Threshold = &v
	return s
}

type ModifyUnrecognizingConfigResponseBody struct {
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUnrecognizingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUnrecognizingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUnrecognizingConfigResponseBody) SetRequestId(v string) *ModifyUnrecognizingConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUnrecognizingConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUnrecognizingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUnrecognizingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUnrecognizingConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyUnrecognizingConfigResponse) SetHeaders(v map[string]*string) *ModifyUnrecognizingConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyUnrecognizingConfigResponse) SetStatusCode(v int32) *ModifyUnrecognizingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUnrecognizingConfigResponse) SetBody(v *ModifyUnrecognizingConfigResponseBody) *ModifyUnrecognizingConfigResponse {
	s.Body = v
	return s
}

type QueryConversationsRequest struct {
	// example:
	//
	// 1582183381000
	BeginTimeLeftRange *int64 `json:"BeginTimeLeftRange,omitempty" xml:"BeginTimeLeftRange,omitempty"`
	// example:
	//
	// 1582356181000
	BeginTimeRightRange *int64 `json:"BeginTimeRightRange,omitempty" xml:"BeginTimeRightRange,omitempty"`
	// example:
	//
	// 02811111111
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12f407b22cbe4890ac595f09985848d5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryConversationsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationsRequest) GoString() string {
	return s.String()
}

func (s *QueryConversationsRequest) SetBeginTimeLeftRange(v int64) *QueryConversationsRequest {
	s.BeginTimeLeftRange = &v
	return s
}

func (s *QueryConversationsRequest) SetBeginTimeRightRange(v int64) *QueryConversationsRequest {
	s.BeginTimeRightRange = &v
	return s
}

func (s *QueryConversationsRequest) SetCallingNumber(v string) *QueryConversationsRequest {
	s.CallingNumber = &v
	return s
}

func (s *QueryConversationsRequest) SetInstanceId(v string) *QueryConversationsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryConversationsRequest) SetPageNumber(v int32) *QueryConversationsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryConversationsRequest) SetPageSize(v int32) *QueryConversationsRequest {
	s.PageSize = &v
	return s
}

type QueryConversationsResponseBody struct {
	Conversations []*QueryConversationsResponseBodyConversations `json:"Conversations,omitempty" xml:"Conversations,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryConversationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConversationsResponseBody) SetConversations(v []*QueryConversationsResponseBodyConversations) *QueryConversationsResponseBody {
	s.Conversations = v
	return s
}

func (s *QueryConversationsResponseBody) SetPageNumber(v int32) *QueryConversationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryConversationsResponseBody) SetPageSize(v int32) *QueryConversationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryConversationsResponseBody) SetRequestId(v string) *QueryConversationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConversationsResponseBody) SetTotalCount(v int64) *QueryConversationsResponseBody {
	s.TotalCount = &v
	return s
}

type QueryConversationsResponseBodyConversations struct {
	// example:
	//
	// 1582183381000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 02811111111
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// 8
	EffectiveAnswerCount *int32 `json:"EffectiveAnswerCount,omitempty" xml:"EffectiveAnswerCount,omitempty"`
	// example:
	//
	// 1582183481000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// AAA
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// true
	TransferredToAgent *bool `json:"TransferredToAgent,omitempty" xml:"TransferredToAgent,omitempty"`
	// example:
	//
	// 10
	UserUtteranceCount *int32 `json:"UserUtteranceCount,omitempty" xml:"UserUtteranceCount,omitempty"`
}

func (s QueryConversationsResponseBodyConversations) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationsResponseBodyConversations) GoString() string {
	return s.String()
}

func (s *QueryConversationsResponseBodyConversations) SetBeginTime(v int64) *QueryConversationsResponseBodyConversations {
	s.BeginTime = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetCallingNumber(v string) *QueryConversationsResponseBodyConversations {
	s.CallingNumber = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetConversationId(v string) *QueryConversationsResponseBodyConversations {
	s.ConversationId = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetEffectiveAnswerCount(v int32) *QueryConversationsResponseBodyConversations {
	s.EffectiveAnswerCount = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetEndTime(v int64) *QueryConversationsResponseBodyConversations {
	s.EndTime = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetSkillGroupId(v string) *QueryConversationsResponseBodyConversations {
	s.SkillGroupId = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetTransferredToAgent(v bool) *QueryConversationsResponseBodyConversations {
	s.TransferredToAgent = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetUserUtteranceCount(v int32) *QueryConversationsResponseBodyConversations {
	s.UserUtteranceCount = &v
	return s
}

type QueryConversationsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConversationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConversationsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationsResponse) GoString() string {
	return s.String()
}

func (s *QueryConversationsResponse) SetHeaders(v map[string]*string) *QueryConversationsResponse {
	s.Headers = v
	return s
}

func (s *QueryConversationsResponse) SetStatusCode(v int32) *QueryConversationsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConversationsResponse) SetBody(v *QueryConversationsResponseBody) *QueryConversationsResponse {
	s.Body = v
	return s
}

type SaveRecordingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 390515b5-6115-4ccf-83e2-52d5bfaf2ddf
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e6bef0db439d4048bfcf45322491becf.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://test/record/
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1971226538081821
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1582267398628
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Source
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VoiceSliceRecordingList *string `json:"VoiceSliceRecordingList,omitempty" xml:"VoiceSliceRecordingList,omitempty"`
}

func (s SaveRecordingRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveRecordingRequest) GoString() string {
	return s.String()
}

func (s *SaveRecordingRequest) SetConversationId(v string) *SaveRecordingRequest {
	s.ConversationId = &v
	return s
}

func (s *SaveRecordingRequest) SetDuration(v string) *SaveRecordingRequest {
	s.Duration = &v
	return s
}

func (s *SaveRecordingRequest) SetFileName(v string) *SaveRecordingRequest {
	s.FileName = &v
	return s
}

func (s *SaveRecordingRequest) SetFilePath(v string) *SaveRecordingRequest {
	s.FilePath = &v
	return s
}

func (s *SaveRecordingRequest) SetInstanceId(v string) *SaveRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveRecordingRequest) SetInstanceOwnerId(v int64) *SaveRecordingRequest {
	s.InstanceOwnerId = &v
	return s
}

func (s *SaveRecordingRequest) SetStartTime(v int64) *SaveRecordingRequest {
	s.StartTime = &v
	return s
}

func (s *SaveRecordingRequest) SetType(v string) *SaveRecordingRequest {
	s.Type = &v
	return s
}

func (s *SaveRecordingRequest) SetVoiceSliceRecordingList(v string) *SaveRecordingRequest {
	s.VoiceSliceRecordingList = &v
	return s
}

type SaveRecordingResponseBody struct {
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveRecordingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *SaveRecordingResponseBody) SetRequestId(v string) *SaveRecordingResponseBody {
	s.RequestId = &v
	return s
}

type SaveRecordingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveRecordingResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveRecordingResponse) GoString() string {
	return s.String()
}

func (s *SaveRecordingResponse) SetHeaders(v map[string]*string) *SaveRecordingResponse {
	s.Headers = v
	return s
}

func (s *SaveRecordingResponse) SetStatusCode(v int32) *SaveRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveRecordingResponse) SetBody(v *SaveRecordingResponseBody) *SaveRecordingResponse {
	s.Body = v
	return s
}

type SilenceTimeoutRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0099b75d-60fd-4c63-8541-7fbba0ae6bb0
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// {}
	InitialContext *string `json:"InitialContext,omitempty" xml:"InitialContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0099b75d-60fd-4c63-8541-7fbba0ae6bb0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1231639035307976
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
}

func (s SilenceTimeoutRequest) String() string {
	return tea.Prettify(s)
}

func (s SilenceTimeoutRequest) GoString() string {
	return s.String()
}

func (s *SilenceTimeoutRequest) SetConversationId(v string) *SilenceTimeoutRequest {
	s.ConversationId = &v
	return s
}

func (s *SilenceTimeoutRequest) SetInitialContext(v string) *SilenceTimeoutRequest {
	s.InitialContext = &v
	return s
}

func (s *SilenceTimeoutRequest) SetInstanceId(v string) *SilenceTimeoutRequest {
	s.InstanceId = &v
	return s
}

func (s *SilenceTimeoutRequest) SetInstanceOwnerId(v int64) *SilenceTimeoutRequest {
	s.InstanceOwnerId = &v
	return s
}

type SilenceTimeoutResponseBody struct {
	// example:
	//
	// TransferToAgent
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// { "skillGroupId": "ABC"}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// example:
	//
	// false
	Interruptible *bool `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TextResponse *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
}

func (s SilenceTimeoutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SilenceTimeoutResponseBody) GoString() string {
	return s.String()
}

func (s *SilenceTimeoutResponseBody) SetAction(v string) *SilenceTimeoutResponseBody {
	s.Action = &v
	return s
}

func (s *SilenceTimeoutResponseBody) SetActionParams(v string) *SilenceTimeoutResponseBody {
	s.ActionParams = &v
	return s
}

func (s *SilenceTimeoutResponseBody) SetInterruptible(v bool) *SilenceTimeoutResponseBody {
	s.Interruptible = &v
	return s
}

func (s *SilenceTimeoutResponseBody) SetRequestId(v string) *SilenceTimeoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *SilenceTimeoutResponseBody) SetTextResponse(v string) *SilenceTimeoutResponseBody {
	s.TextResponse = &v
	return s
}

type SilenceTimeoutResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SilenceTimeoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SilenceTimeoutResponse) String() string {
	return tea.Prettify(s)
}

func (s SilenceTimeoutResponse) GoString() string {
	return s.String()
}

func (s *SilenceTimeoutResponse) SetHeaders(v map[string]*string) *SilenceTimeoutResponse {
	s.Headers = v
	return s
}

func (s *SilenceTimeoutResponse) SetStatusCode(v int32) *SilenceTimeoutResponse {
	s.StatusCode = &v
	return s
}

func (s *SilenceTimeoutResponse) SetBody(v *SilenceTimeoutResponseBody) *SilenceTimeoutResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("voicenavigator"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AssociateChatbotInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateChatbotInstanceResponse
func (client *Client) AssociateChatbotInstanceWithOptions(request *AssociateChatbotInstanceRequest, runtime *util.RuntimeOptions) (_result *AssociateChatbotInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChatbotInstanceId)) {
		query["ChatbotInstanceId"] = request.ChatbotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ChatbotName)) {
		query["ChatbotName"] = request.ChatbotName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NluServiceParamsJson)) {
		query["NluServiceParamsJson"] = request.NluServiceParamsJson
	}

	if !tea.BoolValue(util.IsUnset(request.NluServiceType)) {
		query["NluServiceType"] = request.NluServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.UnionSource)) {
		query["UnionSource"] = request.UnionSource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateChatbotInstance"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateChatbotInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AssociateChatbotInstanceRequest
//
// @return AssociateChatbotInstanceResponse
func (client *Client) AssociateChatbotInstance(request *AssociateChatbotInstanceRequest) (_result *AssociateChatbotInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateChatbotInstanceResponse{}
	_body, _err := client.AssociateChatbotInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AuditTTSVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuditTTSVoiceResponse
func (client *Client) AuditTTSVoiceWithOptions(request *AuditTTSVoiceRequest, runtime *util.RuntimeOptions) (_result *AuditTTSVoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["AccessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PitchRate)) {
		query["PitchRate"] = request.PitchRate
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKey)) {
		query["SecretKey"] = request.SecretKey
	}

	if !tea.BoolValue(util.IsUnset(request.SpeechRate)) {
		query["SpeechRate"] = request.SpeechRate
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		query["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.Voice)) {
		query["Voice"] = request.Voice
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AuditTTSVoice"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuditTTSVoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AuditTTSVoiceRequest
//
// @return AuditTTSVoiceResponse
func (client *Client) AuditTTSVoice(request *AuditTTSVoiceRequest) (_result *AuditTTSVoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuditTTSVoiceResponse{}
	_body, _err := client.AuditTTSVoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BeginDialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BeginDialogueResponse
func (client *Client) BeginDialogueWithOptions(request *BeginDialogueRequest, runtime *util.RuntimeOptions) (_result *BeginDialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CallingNumber)) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		query["ConversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.InitialContext)) {
		query["InitialContext"] = request.InitialContext
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceOwnerId)) {
		query["InstanceOwnerId"] = request.InstanceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BeginDialogue"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BeginDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - BeginDialogueRequest
//
// @return BeginDialogueResponse
func (client *Client) BeginDialogue(request *BeginDialogueRequest) (_result *BeginDialogueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BeginDialogueResponse{}
	_body, _err := client.BeginDialogueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CollectedNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CollectedNumberResponse
func (client *Client) CollectedNumberWithOptions(request *CollectedNumberRequest, runtime *util.RuntimeOptions) (_result *CollectedNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdditionalContext)) {
		query["AdditionalContext"] = request.AdditionalContext
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		query["ConversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceOwnerId)) {
		query["InstanceOwnerId"] = request.InstanceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Number)) {
		query["Number"] = request.Number
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CollectedNumber"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CollectedNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CollectedNumberRequest
//
// @return CollectedNumberResponse
func (client *Client) CollectedNumber(request *CollectedNumberRequest) (_result *CollectedNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CollectedNumberResponse{}
	_body, _err := client.CollectedNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDownloadUrlResponse
func (client *Client) CreateDownloadUrlWithOptions(request *CreateDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *CreateDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDownloadUrl"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateDownloadUrlRequest
//
// @return CreateDownloadUrlResponse
func (client *Client) CreateDownloadUrl(request *CreateDownloadUrlRequest) (_result *CreateDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDownloadUrlResponse{}
	_body, _err := client.CreateDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Concurrency)) {
		query["Concurrency"] = request.Concurrency
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NluServiceParamsJson)) {
		query["NluServiceParamsJson"] = request.NluServiceParamsJson
	}

	if !tea.BoolValue(util.IsUnset(request.UnionInstanceId)) {
		query["UnionInstanceId"] = request.UnionInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionSource)) {
		query["UnionSource"] = request.UnionSource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DebugBeginDialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DebugBeginDialogueResponse
func (client *Client) DebugBeginDialogueWithOptions(request *DebugBeginDialogueRequest, runtime *util.RuntimeOptions) (_result *DebugBeginDialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CallingNumber)) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		query["ConversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.InitialContext)) {
		query["InitialContext"] = request.InitialContext
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DebugBeginDialogue"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DebugBeginDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DebugBeginDialogueRequest
//
// @return DebugBeginDialogueResponse
func (client *Client) DebugBeginDialogue(request *DebugBeginDialogueRequest) (_result *DebugBeginDialogueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DebugBeginDialogueResponse{}
	_body, _err := client.DebugBeginDialogueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DebugCollectedNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DebugCollectedNumberResponse
func (client *Client) DebugCollectedNumberWithOptions(request *DebugCollectedNumberRequest, runtime *util.RuntimeOptions) (_result *DebugCollectedNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		query["ConversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Number)) {
		query["Number"] = request.Number
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DebugCollectedNumber"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DebugCollectedNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DebugCollectedNumberRequest
//
// @return DebugCollectedNumberResponse
func (client *Client) DebugCollectedNumber(request *DebugCollectedNumberRequest) (_result *DebugCollectedNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DebugCollectedNumberResponse{}
	_body, _err := client.DebugCollectedNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DebugDialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DebugDialogueResponse
func (client *Client) DebugDialogueWithOptions(request *DebugDialogueRequest, runtime *util.RuntimeOptions) (_result *DebugDialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdditionalContext)) {
		query["AdditionalContext"] = request.AdditionalContext
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		query["ConversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Utterance)) {
		query["Utterance"] = request.Utterance
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DebugDialogue"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DebugDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DebugDialogueRequest
//
// @return DebugDialogueResponse
func (client *Client) DebugDialogue(request *DebugDialogueRequest) (_result *DebugDialogueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DebugDialogueResponse{}
	_body, _err := client.DebugDialogueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeConversationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConversationResponse
func (client *Client) DescribeConversationWithOptions(request *DescribeConversationRequest, runtime *util.RuntimeOptions) (_result *DescribeConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeConversation"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeConversationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeConversationRequest
//
// @return DescribeConversationResponse
func (client *Client) DescribeConversation(request *DescribeConversationRequest) (_result *DescribeConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConversationResponse{}
	_body, _err := client.DescribeConversationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeConversationContextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConversationContextResponse
func (client *Client) DescribeConversationContextWithOptions(request *DescribeConversationContextRequest, runtime *util.RuntimeOptions) (_result *DescribeConversationContextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeConversationContext"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeConversationContextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeConversationContextRequest
//
// @return DescribeConversationContextResponse
func (client *Client) DescribeConversationContext(request *DescribeConversationContextRequest) (_result *DescribeConversationContextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConversationContextResponse{}
	_body, _err := client.DescribeConversationContextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeExportProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExportProgressResponse
func (client *Client) DescribeExportProgressWithOptions(request *DescribeExportProgressRequest, runtime *util.RuntimeOptions) (_result *DescribeExportProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExportProgress"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExportProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeExportProgressRequest
//
// @return DescribeExportProgressResponse
func (client *Client) DescribeExportProgress(request *DescribeExportProgressRequest) (_result *DescribeExportProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExportProgressResponse{}
	_body, _err := client.DescribeExportProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstance"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeInstanceRequest
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeNavigationConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNavigationConfigResponse
func (client *Client) DescribeNavigationConfigWithOptions(request *DescribeNavigationConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeNavigationConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNavigationConfig"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNavigationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeNavigationConfigRequest
//
// @return DescribeNavigationConfigResponse
func (client *Client) DescribeNavigationConfig(request *DescribeNavigationConfigRequest) (_result *DescribeNavigationConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNavigationConfigResponse{}
	_body, _err := client.DescribeNavigationConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordingResponse
func (client *Client) DescribeRecordingWithOptions(request *DescribeRecordingRequest, runtime *util.RuntimeOptions) (_result *DescribeRecordingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRecording"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRecordingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRecordingRequest
//
// @return DescribeRecordingResponse
func (client *Client) DescribeRecording(request *DescribeRecordingRequest) (_result *DescribeRecordingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecordingResponse{}
	_body, _err := client.DescribeRecordingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeStatisticalDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStatisticalDataResponse
func (client *Client) DescribeStatisticalDataWithOptions(request *DescribeStatisticalDataRequest, runtime *util.RuntimeOptions) (_result *DescribeStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStatisticalData"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStatisticalDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeStatisticalDataRequest
//
// @return DescribeStatisticalDataResponse
func (client *Client) DescribeStatisticalData(request *DescribeStatisticalDataRequest) (_result *DescribeStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStatisticalDataResponse{}
	_body, _err := client.DescribeStatisticalDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeTTSConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTTSConfigResponse
func (client *Client) DescribeTTSConfigWithOptions(request *DescribeTTSConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeTTSConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTTSConfig"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTTSConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeTTSConfigRequest
//
// @return DescribeTTSConfigResponse
func (client *Client) DescribeTTSConfig(request *DescribeTTSConfigRequest) (_result *DescribeTTSConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTTSConfigResponse{}
	_body, _err := client.DescribeTTSConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DialogueResponse
func (client *Client) DialogueWithOptions(request *DialogueRequest, runtime *util.RuntimeOptions) (_result *DialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdditionalContext)) {
		query["AdditionalContext"] = request.AdditionalContext
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CallingNumber)) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		query["ConversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Emotion)) {
		query["Emotion"] = request.Emotion
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceOwnerId)) {
		query["InstanceOwnerId"] = request.InstanceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Utterance)) {
		query["Utterance"] = request.Utterance
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Dialogue"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DialogueRequest
//
// @return DialogueResponse
func (client *Client) Dialogue(request *DialogueRequest) (_result *DialogueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DialogueResponse{}
	_body, _err := client.DialogueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DisableInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableInstanceResponse
func (client *Client) DisableInstanceWithOptions(request *DisableInstanceRequest, runtime *util.RuntimeOptions) (_result *DisableInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableInstance"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DisableInstanceRequest
//
// @return DisableInstanceResponse
func (client *Client) DisableInstance(request *DisableInstanceRequest) (_result *DisableInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableInstanceResponse{}
	_body, _err := client.DisableInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EnableInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableInstanceResponse
func (client *Client) EnableInstanceWithOptions(request *EnableInstanceRequest, runtime *util.RuntimeOptions) (_result *EnableInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableInstance"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableInstanceRequest
//
// @return EnableInstanceResponse
func (client *Client) EnableInstance(request *EnableInstanceRequest) (_result *EnableInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableInstanceResponse{}
	_body, _err := client.EnableInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EndDialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EndDialogueResponse
func (client *Client) EndDialogueWithOptions(request *EndDialogueRequest, runtime *util.RuntimeOptions) (_result *EndDialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		query["ConversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.HangUpParams)) {
		query["HangUpParams"] = request.HangUpParams
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceOwnerId)) {
		query["InstanceOwnerId"] = request.InstanceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EndDialogue"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EndDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - EndDialogueRequest
//
// @return EndDialogueResponse
func (client *Client) EndDialogue(request *EndDialogueRequest) (_result *EndDialogueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EndDialogueResponse{}
	_body, _err := client.EndDialogueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ExportConversationDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportConversationDetailsResponse
func (client *Client) ExportConversationDetailsWithOptions(request *ExportConversationDetailsRequest, runtime *util.RuntimeOptions) (_result *ExportConversationDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTimeLeftRange)) {
		query["BeginTimeLeftRange"] = request.BeginTimeLeftRange
	}

	if !tea.BoolValue(util.IsUnset(request.BeginTimeRightRange)) {
		query["BeginTimeRightRange"] = request.BeginTimeRightRange
	}

	if !tea.BoolValue(util.IsUnset(request.CallingNumber)) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		query["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.Result)) {
		query["Result"] = request.Result
	}

	if !tea.BoolValue(util.IsUnset(request.RoundsLeftRange)) {
		query["RoundsLeftRange"] = request.RoundsLeftRange
	}

	if !tea.BoolValue(util.IsUnset(request.RoundsRightRange)) {
		query["RoundsRightRange"] = request.RoundsRightRange
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportConversationDetails"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportConversationDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ExportConversationDetailsRequest
//
// @return ExportConversationDetailsResponse
func (client *Client) ExportConversationDetails(request *ExportConversationDetailsRequest) (_result *ExportConversationDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportConversationDetailsResponse{}
	_body, _err := client.ExportConversationDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ExportStatisticalDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportStatisticalDataResponse
func (client *Client) ExportStatisticalDataWithOptions(request *ExportStatisticalDataRequest, runtime *util.RuntimeOptions) (_result *ExportStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTimeLeftRange)) {
		query["BeginTimeLeftRange"] = request.BeginTimeLeftRange
	}

	if !tea.BoolValue(util.IsUnset(request.BeginTimeRightRange)) {
		query["BeginTimeRightRange"] = request.BeginTimeRightRange
	}

	if !tea.BoolValue(util.IsUnset(request.ExportType)) {
		query["ExportType"] = request.ExportType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TimeUnit)) {
		query["TimeUnit"] = request.TimeUnit
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportStatisticalData"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportStatisticalDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ExportStatisticalDataRequest
//
// @return ExportStatisticalDataResponse
func (client *Client) ExportStatisticalData(request *ExportStatisticalDataRequest) (_result *ExportStatisticalDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportStatisticalDataResponse{}
	_body, _err := client.ExportStatisticalDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GenerateUploadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUploadUrlResponse
func (client *Client) GenerateUploadUrlWithOptions(request *GenerateUploadUrlRequest, runtime *util.RuntimeOptions) (_result *GenerateUploadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallerBid)) {
		body["CallerBid"] = request.CallerBid
	}

	if !tea.BoolValue(util.IsUnset(request.CallerIp)) {
		body["CallerIp"] = request.CallerIp
	}

	if !tea.BoolValue(util.IsUnset(request.CallerParentId)) {
		body["CallerParentId"] = request.CallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.CallerType)) {
		body["CallerType"] = request.CallerType
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		body["CallerUid"] = request.CallerUid
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		body["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		body["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceOwnerId)) {
		body["InstanceOwnerId"] = request.InstanceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		body["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.MfaPresent)) {
		body["MfaPresent"] = request.MfaPresent
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyOriginalSecurityTransport)) {
		body["ProxyOriginalSecurityTransport"] = request.ProxyOriginalSecurityTransport
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyOriginalSourceIp)) {
		body["ProxyOriginalSourceIp"] = request.ProxyOriginalSourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyTrustTransportInfo)) {
		body["ProxyTrustTransportInfo"] = request.ProxyTrustTransportInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		body["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityTransport)) {
		body["SecurityTransport"] = request.SecurityTransport
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantName)) {
		body["TenantName"] = request.TenantName
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.XspaceServicerId)) {
		body["XspaceServicerId"] = request.XspaceServicerId
	}

	if !tea.BoolValue(util.IsUnset(request.XspaceTenantBuId)) {
		body["XspaceTenantBuId"] = request.XspaceTenantBuId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateUploadUrl"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateUploadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GenerateUploadUrlRequest
//
// @return GenerateUploadUrlResponse
func (client *Client) GenerateUploadUrl(request *GenerateUploadUrlRequest) (_result *GenerateUploadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateUploadUrlResponse{}
	_body, _err := client.GenerateUploadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAsrConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsrConfigResponse
func (client *Client) GetAsrConfigWithOptions(request *GetAsrConfigRequest, runtime *util.RuntimeOptions) (_result *GetAsrConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigLevel)) {
		query["ConfigLevel"] = request.ConfigLevel
	}

	if !tea.BoolValue(util.IsUnset(request.EntryId)) {
		query["EntryId"] = request.EntryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsrConfig"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsrConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAsrConfigRequest
//
// @return GetAsrConfigResponse
func (client *Client) GetAsrConfig(request *GetAsrConfigRequest) (_result *GetAsrConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsrConfigResponse{}
	_body, _err := client.GetAsrConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetRealTimeConcurrencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRealTimeConcurrencyResponse
func (client *Client) GetRealTimeConcurrencyWithOptions(request *GetRealTimeConcurrencyRequest, runtime *util.RuntimeOptions) (_result *GetRealTimeConcurrencyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRealTimeConcurrency"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRealTimeConcurrencyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetRealTimeConcurrencyRequest
//
// @return GetRealTimeConcurrencyResponse
func (client *Client) GetRealTimeConcurrency(request *GetRealTimeConcurrencyRequest) (_result *GetRealTimeConcurrencyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRealTimeConcurrencyResponse{}
	_body, _err := client.GetRealTimeConcurrencyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListChatbotInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChatbotInstancesResponse
func (client *Client) ListChatbotInstancesWithOptions(request *ListChatbotInstancesRequest, runtime *util.RuntimeOptions) (_result *ListChatbotInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListChatbotInstances"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListChatbotInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListChatbotInstancesRequest
//
// @return ListChatbotInstancesResponse
func (client *Client) ListChatbotInstances(request *ListChatbotInstancesRequest) (_result *ListChatbotInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListChatbotInstancesResponse{}
	_body, _err := client.ListChatbotInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListConversationDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConversationDetailsResponse
func (client *Client) ListConversationDetailsWithOptions(request *ListConversationDetailsRequest, runtime *util.RuntimeOptions) (_result *ListConversationDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConversationDetails"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConversationDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListConversationDetailsRequest
//
// @return ListConversationDetailsResponse
func (client *Client) ListConversationDetails(request *ListConversationDetailsRequest) (_result *ListConversationDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConversationDetailsResponse{}
	_body, _err := client.ListConversationDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListConversationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConversationsResponse
func (client *Client) ListConversationsWithOptions(request *ListConversationsRequest, runtime *util.RuntimeOptions) (_result *ListConversationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConversations"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConversationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListConversationsRequest
//
// @return ListConversationsResponse
func (client *Client) ListConversations(request *ListConversationsRequest) (_result *ListConversationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConversationsResponse{}
	_body, _err := client.ListConversationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListDownloadTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDownloadTasksResponse
func (client *Client) ListDownloadTasksWithOptions(request *ListDownloadTasksRequest, runtime *util.RuntimeOptions) (_result *ListDownloadTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDownloadTasks"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDownloadTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListDownloadTasksRequest
//
// @return ListDownloadTasksResponse
func (client *Client) ListDownloadTasks(request *ListDownloadTasksRequest) (_result *ListDownloadTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDownloadTasksResponse{}
	_body, _err := client.ListDownloadTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyAsrConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAsrConfigResponse
func (client *Client) ModifyAsrConfigWithOptions(request *ModifyAsrConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyAsrConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AsrAcousticModelId)) {
		query["AsrAcousticModelId"] = request.AsrAcousticModelId
	}

	if !tea.BoolValue(util.IsUnset(request.AsrClassVocabularyId)) {
		query["AsrClassVocabularyId"] = request.AsrClassVocabularyId
	}

	if !tea.BoolValue(util.IsUnset(request.AsrCustomizationId)) {
		query["AsrCustomizationId"] = request.AsrCustomizationId
	}

	if !tea.BoolValue(util.IsUnset(request.AsrVocabularyId)) {
		query["AsrVocabularyId"] = request.AsrVocabularyId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigLevel)) {
		query["ConfigLevel"] = request.ConfigLevel
	}

	if !tea.BoolValue(util.IsUnset(request.EntryId)) {
		query["EntryId"] = request.EntryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAsrConfig"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAsrConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyAsrConfigRequest
//
// @return ModifyAsrConfigResponse
func (client *Client) ModifyAsrConfig(request *ModifyAsrConfigRequest) (_result *ModifyAsrConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAsrConfigResponse{}
	_body, _err := client.ModifyAsrConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyGreetingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGreetingConfigResponse
func (client *Client) ModifyGreetingConfigWithOptions(request *ModifyGreetingConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyGreetingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GreetingWords)) {
		query["GreetingWords"] = request.GreetingWords
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IntentTrigger)) {
		query["IntentTrigger"] = request.IntentTrigger
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyGreetingConfig"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyGreetingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyGreetingConfigRequest
//
// @return ModifyGreetingConfigResponse
func (client *Client) ModifyGreetingConfig(request *ModifyGreetingConfigRequest) (_result *ModifyGreetingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGreetingConfigResponse{}
	_body, _err := client.ModifyGreetingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceResponse
func (client *Client) ModifyInstanceWithOptions(request *ModifyInstanceRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Concurrency)) {
		query["Concurrency"] = request.Concurrency
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstance"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyInstanceRequest
//
// @return ModifyInstanceResponse
func (client *Client) ModifyInstance(request *ModifyInstanceRequest) (_result *ModifyInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.ModifyInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifySilenceTimeoutConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySilenceTimeoutConfigResponse
func (client *Client) ModifySilenceTimeoutConfigWithOptions(request *ModifySilenceTimeoutConfigRequest, runtime *util.RuntimeOptions) (_result *ModifySilenceTimeoutConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FinalAction)) {
		query["FinalAction"] = request.FinalAction
	}

	if !tea.BoolValue(util.IsUnset(request.FinalActionParams)) {
		query["FinalActionParams"] = request.FinalActionParams
	}

	if !tea.BoolValue(util.IsUnset(request.FinalPrompt)) {
		query["FinalPrompt"] = request.FinalPrompt
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IntentTrigger)) {
		query["IntentTrigger"] = request.IntentTrigger
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		query["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["Timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySilenceTimeoutConfig"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySilenceTimeoutConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifySilenceTimeoutConfigRequest
//
// @return ModifySilenceTimeoutConfigResponse
func (client *Client) ModifySilenceTimeoutConfig(request *ModifySilenceTimeoutConfigRequest) (_result *ModifySilenceTimeoutConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySilenceTimeoutConfigResponse{}
	_body, _err := client.ModifySilenceTimeoutConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyTTSConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTTSConfigResponse
func (client *Client) ModifyTTSConfigWithOptions(request *ModifyTTSConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyTTSConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliCustomizedVoice)) {
		query["AliCustomizedVoice"] = request.AliCustomizedVoice
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.EngineXunfei)) {
		query["EngineXunfei"] = request.EngineXunfei
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NlsServiceType)) {
		query["NlsServiceType"] = request.NlsServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.SpeechRate)) {
		query["SpeechRate"] = request.SpeechRate
	}

	if !tea.BoolValue(util.IsUnset(request.Voice)) {
		query["Voice"] = request.Voice
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTTSConfig"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTTSConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyTTSConfigRequest
//
// @return ModifyTTSConfigResponse
func (client *Client) ModifyTTSConfig(request *ModifyTTSConfigRequest) (_result *ModifyTTSConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTTSConfigResponse{}
	_body, _err := client.ModifyTTSConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyUnrecognizingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUnrecognizingConfigResponse
func (client *Client) ModifyUnrecognizingConfigWithOptions(request *ModifyUnrecognizingConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyUnrecognizingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FinalAction)) {
		query["FinalAction"] = request.FinalAction
	}

	if !tea.BoolValue(util.IsUnset(request.FinalActionParams)) {
		query["FinalActionParams"] = request.FinalActionParams
	}

	if !tea.BoolValue(util.IsUnset(request.FinalPrompt)) {
		query["FinalPrompt"] = request.FinalPrompt
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		query["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUnrecognizingConfig"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUnrecognizingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyUnrecognizingConfigRequest
//
// @return ModifyUnrecognizingConfigResponse
func (client *Client) ModifyUnrecognizingConfig(request *ModifyUnrecognizingConfigRequest) (_result *ModifyUnrecognizingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUnrecognizingConfigResponse{}
	_body, _err := client.ModifyUnrecognizingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryConversationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConversationsResponse
func (client *Client) QueryConversationsWithOptions(request *QueryConversationsRequest, runtime *util.RuntimeOptions) (_result *QueryConversationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryConversations"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryConversationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryConversationsRequest
//
// @return QueryConversationsResponse
func (client *Client) QueryConversations(request *QueryConversationsRequest) (_result *QueryConversationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryConversationsResponse{}
	_body, _err := client.QueryConversationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SaveRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveRecordingResponse
func (client *Client) SaveRecordingWithOptions(request *SaveRecordingRequest, runtime *util.RuntimeOptions) (_result *SaveRecordingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		query["ConversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		query["FilePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceOwnerId)) {
		query["InstanceOwnerId"] = request.InstanceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceSliceRecordingList)) {
		query["VoiceSliceRecordingList"] = request.VoiceSliceRecordingList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveRecording"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveRecordingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveRecordingRequest
//
// @return SaveRecordingResponse
func (client *Client) SaveRecording(request *SaveRecordingRequest) (_result *SaveRecordingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveRecordingResponse{}
	_body, _err := client.SaveRecordingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SilenceTimeoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SilenceTimeoutResponse
func (client *Client) SilenceTimeoutWithOptions(request *SilenceTimeoutRequest, runtime *util.RuntimeOptions) (_result *SilenceTimeoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		query["ConversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.InitialContext)) {
		query["InitialContext"] = request.InitialContext
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceOwnerId)) {
		query["InstanceOwnerId"] = request.InstanceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SilenceTimeout"),
		Version:     tea.String("2018-06-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SilenceTimeoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SilenceTimeoutRequest
//
// @return SilenceTimeoutResponse
func (client *Client) SilenceTimeout(request *SilenceTimeoutRequest) (_result *SilenceTimeoutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SilenceTimeoutResponse{}
	_body, _err := client.SilenceTimeoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
