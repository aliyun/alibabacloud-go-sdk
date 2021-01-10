// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AssociateChatbotInstanceRequest struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ChatbotInstanceId *string `json:"ChatbotInstanceId,omitempty" xml:"ChatbotInstanceId,omitempty"`
	ChatbotName       *string `json:"ChatbotName,omitempty" xml:"ChatbotName,omitempty"`
}

func (s AssociateChatbotInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateChatbotInstanceRequest) GoString() string {
	return s.String()
}

func (s *AssociateChatbotInstanceRequest) SetInstanceId(v string) *AssociateChatbotInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateChatbotInstanceRequest) SetChatbotInstanceId(v string) *AssociateChatbotInstanceRequest {
	s.ChatbotInstanceId = &v
	return s
}

func (s *AssociateChatbotInstanceRequest) SetChatbotName(v string) *AssociateChatbotInstanceRequest {
	s.ChatbotName = &v
	return s
}

type AssociateChatbotInstanceResponseBody struct {
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssociateChatbotInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AssociateChatbotInstanceResponse) SetBody(v *AssociateChatbotInstanceResponseBody) *AssociateChatbotInstanceResponse {
	s.Body = v
	return s
}

type AuditTTSVoiceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Voice      *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	Text       *string `json:"Text,omitempty" xml:"Text,omitempty"`
	SpeechRate *string `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Volume     *string `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s AuditTTSVoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AuditTTSVoiceRequest) GoString() string {
	return s.String()
}

func (s *AuditTTSVoiceRequest) SetInstanceId(v string) *AuditTTSVoiceRequest {
	s.InstanceId = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetVoice(v string) *AuditTTSVoiceRequest {
	s.Voice = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetText(v string) *AuditTTSVoiceRequest {
	s.Text = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetSpeechRate(v string) *AuditTTSVoiceRequest {
	s.SpeechRate = &v
	return s
}

func (s *AuditTTSVoiceRequest) SetVolume(v string) *AuditTTSVoiceRequest {
	s.Volume = &v
	return s
}

type AuditTTSVoiceResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AuditionUrl *string `json:"AuditionUrl,omitempty" xml:"AuditionUrl,omitempty"`
}

func (s AuditTTSVoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuditTTSVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *AuditTTSVoiceResponseBody) SetRequestId(v string) *AuditTTSVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuditTTSVoiceResponseBody) SetAuditionUrl(v string) *AuditTTSVoiceResponseBody {
	s.AuditionUrl = &v
	return s
}

type AuditTTSVoiceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AuditTTSVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AuditTTSVoiceResponse) SetBody(v *AuditTTSVoiceResponseBody) *AuditTTSVoiceResponse {
	s.Body = v
	return s
}

type BeginDialogueRequest struct {
	CalledNumber   *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallingNumber  *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	InitialContext *string `json:"InitialContext,omitempty" xml:"InitialContext,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

type BeginDialogueResponseBody struct {
	Action        *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Interruptible *bool   `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ActionParams  *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	TextResponse  *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
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

func (s *BeginDialogueResponseBody) SetInterruptible(v bool) *BeginDialogueResponseBody {
	s.Interruptible = &v
	return s
}

func (s *BeginDialogueResponseBody) SetRequestId(v string) *BeginDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *BeginDialogueResponseBody) SetActionParams(v string) *BeginDialogueResponseBody {
	s.ActionParams = &v
	return s
}

func (s *BeginDialogueResponseBody) SetTextResponse(v string) *BeginDialogueResponseBody {
	s.TextResponse = &v
	return s
}

type BeginDialogueResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BeginDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BeginDialogueResponse) SetBody(v *BeginDialogueResponseBody) *BeginDialogueResponse {
	s.Body = v
	return s
}

type CollectedNumberRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	Number         *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s CollectedNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s CollectedNumberRequest) GoString() string {
	return s.String()
}

func (s *CollectedNumberRequest) SetInstanceId(v string) *CollectedNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *CollectedNumberRequest) SetConversationId(v string) *CollectedNumberRequest {
	s.ConversationId = &v
	return s
}

func (s *CollectedNumberRequest) SetNumber(v string) *CollectedNumberRequest {
	s.Number = &v
	return s
}

type CollectedNumberResponseBody struct {
	Action        *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Interruptible *bool   `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ActionParams  *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	TextResponse  *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
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

func (s *CollectedNumberResponseBody) SetInterruptible(v bool) *CollectedNumberResponseBody {
	s.Interruptible = &v
	return s
}

func (s *CollectedNumberResponseBody) SetRequestId(v string) *CollectedNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CollectedNumberResponseBody) SetActionParams(v string) *CollectedNumberResponseBody {
	s.ActionParams = &v
	return s
}

func (s *CollectedNumberResponseBody) SetTextResponse(v string) *CollectedNumberResponseBody {
	s.TextResponse = &v
	return s
}

type CollectedNumberResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CollectedNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CollectedNumberResponse) SetBody(v *CollectedNumberResponseBody) *CollectedNumberResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Concurrency       *int64  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	ChatbotInstanceId *string `json:"ChatbotInstanceId,omitempty" xml:"ChatbotInstanceId,omitempty"`
	NluServiceType    *string `json:"NluServiceType,omitempty" xml:"NluServiceType,omitempty"`
	ChatbotName       *string `json:"ChatbotName,omitempty" xml:"ChatbotName,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetName(v string) *CreateInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateInstanceRequest) SetDescription(v string) *CreateInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequest) SetConcurrency(v int64) *CreateInstanceRequest {
	s.Concurrency = &v
	return s
}

func (s *CreateInstanceRequest) SetChatbotInstanceId(v string) *CreateInstanceRequest {
	s.ChatbotInstanceId = &v
	return s
}

func (s *CreateInstanceRequest) SetNluServiceType(v string) *CreateInstanceRequest {
	s.NluServiceType = &v
	return s
}

func (s *CreateInstanceRequest) SetChatbotName(v string) *CreateInstanceRequest {
	s.ChatbotName = &v
	return s
}

type CreateInstanceResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type DebugBeginDialogueRequest struct {
	CalledNumber   *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallingNumber  *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	InitialContext *string `json:"InitialContext,omitempty" xml:"InitialContext,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	Action        *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Interruptible *bool   `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ActionParams  *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	TextResponse  *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
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

func (s *DebugBeginDialogueResponseBody) SetInterruptible(v bool) *DebugBeginDialogueResponseBody {
	s.Interruptible = &v
	return s
}

func (s *DebugBeginDialogueResponseBody) SetRequestId(v string) *DebugBeginDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DebugBeginDialogueResponseBody) SetActionParams(v string) *DebugBeginDialogueResponseBody {
	s.ActionParams = &v
	return s
}

func (s *DebugBeginDialogueResponseBody) SetTextResponse(v string) *DebugBeginDialogueResponseBody {
	s.TextResponse = &v
	return s
}

type DebugBeginDialogueResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DebugBeginDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DebugBeginDialogueResponse) SetBody(v *DebugBeginDialogueResponseBody) *DebugBeginDialogueResponse {
	s.Body = v
	return s
}

type DebugCollectedNumberRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	Number         *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DebugCollectedNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s DebugCollectedNumberRequest) GoString() string {
	return s.String()
}

func (s *DebugCollectedNumberRequest) SetInstanceId(v string) *DebugCollectedNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *DebugCollectedNumberRequest) SetConversationId(v string) *DebugCollectedNumberRequest {
	s.ConversationId = &v
	return s
}

func (s *DebugCollectedNumberRequest) SetNumber(v string) *DebugCollectedNumberRequest {
	s.Number = &v
	return s
}

type DebugCollectedNumberResponseBody struct {
	Action        *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Interruptible *bool   `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ActionParams  *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	TextResponse  *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
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

func (s *DebugCollectedNumberResponseBody) SetInterruptible(v bool) *DebugCollectedNumberResponseBody {
	s.Interruptible = &v
	return s
}

func (s *DebugCollectedNumberResponseBody) SetRequestId(v string) *DebugCollectedNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DebugCollectedNumberResponseBody) SetActionParams(v string) *DebugCollectedNumberResponseBody {
	s.ActionParams = &v
	return s
}

func (s *DebugCollectedNumberResponseBody) SetTextResponse(v string) *DebugCollectedNumberResponseBody {
	s.TextResponse = &v
	return s
}

type DebugCollectedNumberResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DebugCollectedNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DebugCollectedNumberResponse) SetBody(v *DebugCollectedNumberResponseBody) *DebugCollectedNumberResponse {
	s.Body = v
	return s
}

type DebugDialogueRequest struct {
	ConversationId    *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	AdditionalContext *string `json:"AdditionalContext,omitempty" xml:"AdditionalContext,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Utterance         *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s DebugDialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s DebugDialogueRequest) GoString() string {
	return s.String()
}

func (s *DebugDialogueRequest) SetConversationId(v string) *DebugDialogueRequest {
	s.ConversationId = &v
	return s
}

func (s *DebugDialogueRequest) SetAdditionalContext(v string) *DebugDialogueRequest {
	s.AdditionalContext = &v
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
	Action        *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Interruptible *bool   `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ActionParams  *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	TextResponse  *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
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

func (s *DebugDialogueResponseBody) SetInterruptible(v bool) *DebugDialogueResponseBody {
	s.Interruptible = &v
	return s
}

func (s *DebugDialogueResponseBody) SetRequestId(v string) *DebugDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DebugDialogueResponseBody) SetActionParams(v string) *DebugDialogueResponseBody {
	s.ActionParams = &v
	return s
}

func (s *DebugDialogueResponseBody) SetTextResponse(v string) *DebugDialogueResponseBody {
	s.TextResponse = &v
	return s
}

type DebugDialogueResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DebugDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DebugDialogueResponse) SetBody(v *DebugDialogueResponseBody) *DebugDialogueResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DescribeConversationRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
}

func (s DescribeConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConversationRequest) GoString() string {
	return s.String()
}

func (s *DescribeConversationRequest) SetInstanceId(v string) *DescribeConversationRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeConversationRequest) SetConversationId(v string) *DescribeConversationRequest {
	s.ConversationId = &v
	return s
}

type DescribeConversationResponseBody struct {
	EffectiveAnswerCount *int32  `json:"EffectiveAnswerCount,omitempty" xml:"EffectiveAnswerCount,omitempty"`
	ConversationId       *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	TransferredToAgent   *bool   `json:"TransferredToAgent,omitempty" xml:"TransferredToAgent,omitempty"`
	EndTime              *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BeginTime            *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	SkillGroupId         *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	CallingNumber        *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	UserUtteranceCount   *int32  `json:"UserUtteranceCount,omitempty" xml:"UserUtteranceCount,omitempty"`
}

func (s DescribeConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConversationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConversationResponseBody) SetEffectiveAnswerCount(v int32) *DescribeConversationResponseBody {
	s.EffectiveAnswerCount = &v
	return s
}

func (s *DescribeConversationResponseBody) SetConversationId(v string) *DescribeConversationResponseBody {
	s.ConversationId = &v
	return s
}

func (s *DescribeConversationResponseBody) SetTransferredToAgent(v bool) *DescribeConversationResponseBody {
	s.TransferredToAgent = &v
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

func (s *DescribeConversationResponseBody) SetBeginTime(v int64) *DescribeConversationResponseBody {
	s.BeginTime = &v
	return s
}

func (s *DescribeConversationResponseBody) SetSkillGroupId(v string) *DescribeConversationResponseBody {
	s.SkillGroupId = &v
	return s
}

func (s *DescribeConversationResponseBody) SetCallingNumber(v string) *DescribeConversationResponseBody {
	s.CallingNumber = &v
	return s
}

func (s *DescribeConversationResponseBody) SetUserUtteranceCount(v int32) *DescribeConversationResponseBody {
	s.UserUtteranceCount = &v
	return s
}

type DescribeConversationResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConversationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeConversationResponse) SetBody(v *DescribeConversationResponseBody) *DescribeConversationResponse {
	s.Body = v
	return s
}

type DescribeConversationContextRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
}

func (s DescribeConversationContextRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConversationContextRequest) GoString() string {
	return s.String()
}

func (s *DescribeConversationContextRequest) SetInstanceId(v string) *DescribeConversationContextRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeConversationContextRequest) SetConversationId(v string) *DescribeConversationContextRequest {
	s.ConversationId = &v
	return s
}

type DescribeConversationContextResponseBody struct {
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConversationContext *string `json:"ConversationContext,omitempty" xml:"ConversationContext,omitempty"`
}

func (s DescribeConversationContextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConversationContextResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConversationContextResponseBody) SetRequestId(v string) *DescribeConversationContextResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConversationContextResponseBody) SetConversationContext(v string) *DescribeConversationContextResponseBody {
	s.ConversationContext = &v
	return s
}

type DescribeConversationContextResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConversationContextResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeConversationContextResponse) SetBody(v *DescribeConversationContextResponseBody) *DescribeConversationContextResponse {
	s.Body = v
	return s
}

type DescribeExportProgressRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ExportTaskId *string `json:"ExportTaskId,omitempty" xml:"ExportTaskId,omitempty"`
}

func (s DescribeExportProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportProgressRequest) GoString() string {
	return s.String()
}

func (s *DescribeExportProgressRequest) SetInstanceId(v string) *DescribeExportProgressRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeExportProgressRequest) SetExportTaskId(v string) *DescribeExportProgressRequest {
	s.ExportTaskId = &v
	return s
}

type DescribeExportProgressResponseBody struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FileHttpUrl *string `json:"FileHttpUrl,omitempty" xml:"FileHttpUrl,omitempty"`
}

func (s DescribeExportProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportProgressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExportProgressResponseBody) SetStatus(v string) *DescribeExportProgressResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeExportProgressResponseBody) SetRequestId(v string) *DescribeExportProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExportProgressResponseBody) SetFileHttpUrl(v string) *DescribeExportProgressResponseBody {
	s.FileHttpUrl = &v
	return s
}

type DescribeExportProgressResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExportProgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeExportProgressResponse) SetBody(v *DescribeExportProgressResponseBody) *DescribeExportProgressResponse {
	s.Body = v
	return s
}

type DescribeInstanceRequest struct {
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
	Status               *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	ModifyTime           *int64    `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId            *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId           *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Concurrency          *int64    `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	ApplicableOperations []*string `json:"ApplicableOperations,omitempty" xml:"ApplicableOperations,omitempty" type:"Repeated"`
	ModifyUserName       *string   `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Name                 *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetStatus(v string) *DescribeInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetModifyTime(v int64) *DescribeInstanceResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetDescription(v string) *DescribeInstanceResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceId(v string) *DescribeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetConcurrency(v int64) *DescribeInstanceResponseBody {
	s.Concurrency = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetApplicableOperations(v []*string) *DescribeInstanceResponseBody {
	s.ApplicableOperations = v
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

type DescribeInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeNavigationConfigRequest struct {
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
	RequestId            *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SilenceTimeoutConfig *DescribeNavigationConfigResponseBodySilenceTimeoutConfig `json:"SilenceTimeoutConfig,omitempty" xml:"SilenceTimeoutConfig,omitempty" type:"Struct"`
	GreetingConfig       *DescribeNavigationConfigResponseBodyGreetingConfig       `json:"GreetingConfig,omitempty" xml:"GreetingConfig,omitempty" type:"Struct"`
	UnrecognizingConfig  *DescribeNavigationConfigResponseBodyUnrecognizingConfig  `json:"UnrecognizingConfig,omitempty" xml:"UnrecognizingConfig,omitempty" type:"Struct"`
}

func (s DescribeNavigationConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNavigationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigResponseBody) SetRequestId(v string) *DescribeNavigationConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNavigationConfigResponseBody) SetSilenceTimeoutConfig(v *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) *DescribeNavigationConfigResponseBody {
	s.SilenceTimeoutConfig = v
	return s
}

func (s *DescribeNavigationConfigResponseBody) SetGreetingConfig(v *DescribeNavigationConfigResponseBodyGreetingConfig) *DescribeNavigationConfigResponseBody {
	s.GreetingConfig = v
	return s
}

func (s *DescribeNavigationConfigResponseBody) SetUnrecognizingConfig(v *DescribeNavigationConfigResponseBodyUnrecognizingConfig) *DescribeNavigationConfigResponseBody {
	s.UnrecognizingConfig = v
	return s
}

type DescribeNavigationConfigResponseBodySilenceTimeoutConfig struct {
	Timeout           *int64  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	IntentTrigger     *string `json:"IntentTrigger,omitempty" xml:"IntentTrigger,omitempty"`
	FinalPrompt       *string `json:"FinalPrompt,omitempty" xml:"FinalPrompt,omitempty"`
	SourceType        *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	FinalAction       *string `json:"FinalAction,omitempty" xml:"FinalAction,omitempty"`
	Prompt            *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	Threshold         *int32  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	FinalActionParams *string `json:"FinalActionParams,omitempty" xml:"FinalActionParams,omitempty"`
}

func (s DescribeNavigationConfigResponseBodySilenceTimeoutConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeNavigationConfigResponseBodySilenceTimeoutConfig) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetTimeout(v int64) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.Timeout = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetIntentTrigger(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.IntentTrigger = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetFinalPrompt(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.FinalPrompt = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetSourceType(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.SourceType = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetFinalAction(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.FinalAction = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetPrompt(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.Prompt = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetThreshold(v int32) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.Threshold = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetFinalActionParams(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.FinalActionParams = &v
	return s
}

type DescribeNavigationConfigResponseBodyGreetingConfig struct {
	IntentTrigger *string `json:"IntentTrigger,omitempty" xml:"IntentTrigger,omitempty"`
	GreetingWords *string `json:"GreetingWords,omitempty" xml:"GreetingWords,omitempty"`
	SourceType    *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeNavigationConfigResponseBodyGreetingConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeNavigationConfigResponseBodyGreetingConfig) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigResponseBodyGreetingConfig) SetIntentTrigger(v string) *DescribeNavigationConfigResponseBodyGreetingConfig {
	s.IntentTrigger = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyGreetingConfig) SetGreetingWords(v string) *DescribeNavigationConfigResponseBodyGreetingConfig {
	s.GreetingWords = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyGreetingConfig) SetSourceType(v string) *DescribeNavigationConfigResponseBodyGreetingConfig {
	s.SourceType = &v
	return s
}

type DescribeNavigationConfigResponseBodyUnrecognizingConfig struct {
	FinalPrompt       *string `json:"FinalPrompt,omitempty" xml:"FinalPrompt,omitempty"`
	FinalAction       *string `json:"FinalAction,omitempty" xml:"FinalAction,omitempty"`
	FinalActionParams *string `json:"FinalActionParams,omitempty" xml:"FinalActionParams,omitempty"`
	Threshold         *int32  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Prompt            *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
}

func (s DescribeNavigationConfigResponseBodyUnrecognizingConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeNavigationConfigResponseBodyUnrecognizingConfig) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) SetFinalPrompt(v string) *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	s.FinalPrompt = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) SetFinalAction(v string) *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	s.FinalAction = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) SetFinalActionParams(v string) *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	s.FinalActionParams = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) SetThreshold(v int32) *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	s.Threshold = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) SetPrompt(v string) *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	s.Prompt = &v
	return s
}

type DescribeNavigationConfigResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNavigationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeNavigationConfigResponse) SetBody(v *DescribeNavigationConfigResponseBody) *DescribeNavigationConfigResponse {
	s.Body = v
	return s
}

type DescribeRecordingRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
}

func (s DescribeRecordingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordingRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordingRequest) SetInstanceId(v string) *DescribeRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRecordingRequest) SetConversationId(v string) *DescribeRecordingRequest {
	s.ConversationId = &v
	return s
}

type DescribeRecordingResponseBody struct {
	FilePath  *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s DescribeRecordingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordingResponseBody) SetFilePath(v string) *DescribeRecordingResponseBody {
	s.FilePath = &v
	return s
}

func (s *DescribeRecordingResponseBody) SetRequestId(v string) *DescribeRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordingResponseBody) SetFileName(v string) *DescribeRecordingResponseBody {
	s.FileName = &v
	return s
}

type DescribeRecordingResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRecordingResponse) SetBody(v *DescribeRecordingResponseBody) *DescribeRecordingResponse {
	s.Body = v
	return s
}

type DescribeStatisticalDataRequest struct {
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TimeUnit            *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
	BeginTimeLeftRange  *int64  `json:"BeginTimeLeftRange,omitempty" xml:"BeginTimeLeftRange,omitempty"`
	BeginTimeRightRange *int64  `json:"BeginTimeRightRange,omitempty" xml:"BeginTimeRightRange,omitempty"`
}

func (s DescribeStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeStatisticalDataRequest) SetInstanceId(v string) *DescribeStatisticalDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeStatisticalDataRequest) SetTimeUnit(v string) *DescribeStatisticalDataRequest {
	s.TimeUnit = &v
	return s
}

func (s *DescribeStatisticalDataRequest) SetBeginTimeLeftRange(v int64) *DescribeStatisticalDataRequest {
	s.BeginTimeLeftRange = &v
	return s
}

func (s *DescribeStatisticalDataRequest) SetBeginTimeRightRange(v int64) *DescribeStatisticalDataRequest {
	s.BeginTimeRightRange = &v
	return s
}

type DescribeStatisticalDataResponseBody struct {
	TotalDialoguePassRate    *string                                                      `json:"TotalDialoguePassRate,omitempty" xml:"TotalDialoguePassRate,omitempty"`
	TotalKnowledgeHitRate    *string                                                      `json:"TotalKnowledgeHitRate,omitempty" xml:"TotalKnowledgeHitRate,omitempty"`
	TotalResolutionRate      *string                                                      `json:"TotalResolutionRate,omitempty" xml:"TotalResolutionRate,omitempty"`
	TotalValidAnswerRate     *string                                                      `json:"TotalValidAnswerRate,omitempty" xml:"TotalValidAnswerRate,omitempty"`
	RequestId                *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResolvedQuestionTotalNum *int64                                                       `json:"ResolvedQuestionTotalNum,omitempty" xml:"ResolvedQuestionTotalNum,omitempty"`
	ConversationTotalNum     *int64                                                       `json:"ConversationTotalNum,omitempty" xml:"ConversationTotalNum,omitempty"`
	StatisticalDataReports   []*DescribeStatisticalDataResponseBodyStatisticalDataReports `json:"StatisticalDataReports,omitempty" xml:"StatisticalDataReports,omitempty" type:"Repeated"`
}

func (s DescribeStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticalDataResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeStatisticalDataResponseBody) SetRequestId(v string) *DescribeStatisticalDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetResolvedQuestionTotalNum(v int64) *DescribeStatisticalDataResponseBody {
	s.ResolvedQuestionTotalNum = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetConversationTotalNum(v int64) *DescribeStatisticalDataResponseBody {
	s.ConversationTotalNum = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetStatisticalDataReports(v []*DescribeStatisticalDataResponseBodyStatisticalDataReports) *DescribeStatisticalDataResponseBody {
	s.StatisticalDataReports = v
	return s
}

type DescribeStatisticalDataResponseBodyStatisticalDataReports struct {
	KnowledgeHitRate     *string `json:"KnowledgeHitRate,omitempty" xml:"KnowledgeHitRate,omitempty"`
	ResolvedQuestionNum  *int32  `json:"ResolvedQuestionNum,omitempty" xml:"ResolvedQuestionNum,omitempty"`
	ResolutionRate       *string `json:"ResolutionRate,omitempty" xml:"ResolutionRate,omitempty"`
	StatisticalDate      *string `json:"StatisticalDate,omitempty" xml:"StatisticalDate,omitempty"`
	TotalConversationNum *int32  `json:"TotalConversationNum,omitempty" xml:"TotalConversationNum,omitempty"`
	ValidAnswerRate      *string `json:"ValidAnswerRate,omitempty" xml:"ValidAnswerRate,omitempty"`
	DialoguePassRate     *string `json:"DialoguePassRate,omitempty" xml:"DialoguePassRate,omitempty"`
}

func (s DescribeStatisticalDataResponseBodyStatisticalDataReports) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticalDataResponseBodyStatisticalDataReports) GoString() string {
	return s.String()
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetKnowledgeHitRate(v string) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.KnowledgeHitRate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetResolvedQuestionNum(v int32) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.ResolvedQuestionNum = &v
	return s
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetResolutionRate(v string) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.ResolutionRate = &v
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

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetDialoguePassRate(v string) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.DialoguePassRate = &v
	return s
}

type DescribeStatisticalDataResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeStatisticalDataResponse) SetBody(v *DescribeStatisticalDataResponseBody) *DescribeStatisticalDataResponse {
	s.Body = v
	return s
}

type DescribeTTSConfigRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

type DescribeTTSConfigResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Volume     *int32  `json:"Volume,omitempty" xml:"Volume,omitempty"`
	Voice      *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	SpeechRate *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
}

func (s DescribeTTSConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTTSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTTSConfigResponseBody) SetRequestId(v string) *DescribeTTSConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetVolume(v int32) *DescribeTTSConfigResponseBody {
	s.Volume = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetVoice(v string) *DescribeTTSConfigResponseBody {
	s.Voice = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetSpeechRate(v int32) *DescribeTTSConfigResponseBody {
	s.SpeechRate = &v
	return s
}

type DescribeTTSConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTTSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeTTSConfigResponse) SetBody(v *DescribeTTSConfigResponseBody) *DescribeTTSConfigResponse {
	s.Body = v
	return s
}

type DialogueRequest struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConversationId    *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	Utterance         *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
	CalledNumber      *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallingNumber     *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	AdditionalContext *string `json:"AdditionalContext,omitempty" xml:"AdditionalContext,omitempty"`
}

func (s DialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s DialogueRequest) GoString() string {
	return s.String()
}

func (s *DialogueRequest) SetInstanceId(v string) *DialogueRequest {
	s.InstanceId = &v
	return s
}

func (s *DialogueRequest) SetConversationId(v string) *DialogueRequest {
	s.ConversationId = &v
	return s
}

func (s *DialogueRequest) SetUtterance(v string) *DialogueRequest {
	s.Utterance = &v
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

func (s *DialogueRequest) SetAdditionalContext(v string) *DialogueRequest {
	s.AdditionalContext = &v
	return s
}

type DialogueResponseBody struct {
	Action        *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Interruptible *bool   `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ActionParams  *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	TextResponse  *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
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

func (s *DialogueResponseBody) SetInterruptible(v bool) *DialogueResponseBody {
	s.Interruptible = &v
	return s
}

func (s *DialogueResponseBody) SetRequestId(v string) *DialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DialogueResponseBody) SetActionParams(v string) *DialogueResponseBody {
	s.ActionParams = &v
	return s
}

func (s *DialogueResponseBody) SetTextResponse(v string) *DialogueResponseBody {
	s.TextResponse = &v
	return s
}

type DialogueResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DialogueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DialogueResponse) SetBody(v *DialogueResponseBody) *DialogueResponse {
	s.Body = v
	return s
}

type DisableInstanceRequest struct {
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
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DisableInstanceResponseBody) SetStatus(v string) *DisableInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DisableInstanceResponseBody) SetRequestId(v string) *DisableInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DisableInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DisableInstanceResponse) SetBody(v *DisableInstanceResponseBody) *DisableInstanceResponse {
	s.Body = v
	return s
}

type EnableInstanceRequest struct {
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
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *EnableInstanceResponseBody) SetStatus(v string) *EnableInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *EnableInstanceResponseBody) SetRequestId(v string) *EnableInstanceResponseBody {
	s.RequestId = &v
	return s
}

type EnableInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *EnableInstanceResponse) SetBody(v *EnableInstanceResponseBody) *EnableInstanceResponse {
	s.Body = v
	return s
}

type EndDialogueRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
}

func (s EndDialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s EndDialogueRequest) GoString() string {
	return s.String()
}

func (s *EndDialogueRequest) SetInstanceId(v string) *EndDialogueRequest {
	s.InstanceId = &v
	return s
}

func (s *EndDialogueRequest) SetConversationId(v string) *EndDialogueRequest {
	s.ConversationId = &v
	return s
}

type EndDialogueResponseBody struct {
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EndDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *EndDialogueResponse) SetBody(v *EndDialogueResponseBody) *EndDialogueResponse {
	s.Body = v
	return s
}

type ExportConversationDetailsRequest struct {
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CallingNumber       *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	BeginTimeLeftRange  *int64  `json:"BeginTimeLeftRange,omitempty" xml:"BeginTimeLeftRange,omitempty"`
	BeginTimeRightRange *int64  `json:"BeginTimeRightRange,omitempty" xml:"BeginTimeRightRange,omitempty"`
}

func (s ExportConversationDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportConversationDetailsRequest) GoString() string {
	return s.String()
}

func (s *ExportConversationDetailsRequest) SetInstanceId(v string) *ExportConversationDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *ExportConversationDetailsRequest) SetCallingNumber(v string) *ExportConversationDetailsRequest {
	s.CallingNumber = &v
	return s
}

func (s *ExportConversationDetailsRequest) SetBeginTimeLeftRange(v int64) *ExportConversationDetailsRequest {
	s.BeginTimeLeftRange = &v
	return s
}

func (s *ExportConversationDetailsRequest) SetBeginTimeRightRange(v int64) *ExportConversationDetailsRequest {
	s.BeginTimeRightRange = &v
	return s
}

type ExportConversationDetailsResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExportTaskId *string `json:"ExportTaskId,omitempty" xml:"ExportTaskId,omitempty"`
}

func (s ExportConversationDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportConversationDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ExportConversationDetailsResponseBody) SetRequestId(v string) *ExportConversationDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportConversationDetailsResponseBody) SetExportTaskId(v string) *ExportConversationDetailsResponseBody {
	s.ExportTaskId = &v
	return s
}

type ExportConversationDetailsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExportConversationDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ExportConversationDetailsResponse) SetBody(v *ExportConversationDetailsResponseBody) *ExportConversationDetailsResponse {
	s.Body = v
	return s
}

type ExportStatisticalDataRequest struct {
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TimeUnit            *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
	ExportType          *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	BeginTimeLeftRange  *int64  `json:"BeginTimeLeftRange,omitempty" xml:"BeginTimeLeftRange,omitempty"`
	BeginTimeRightRange *int64  `json:"BeginTimeRightRange,omitempty" xml:"BeginTimeRightRange,omitempty"`
}

func (s ExportStatisticalDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportStatisticalDataRequest) GoString() string {
	return s.String()
}

func (s *ExportStatisticalDataRequest) SetInstanceId(v string) *ExportStatisticalDataRequest {
	s.InstanceId = &v
	return s
}

func (s *ExportStatisticalDataRequest) SetTimeUnit(v string) *ExportStatisticalDataRequest {
	s.TimeUnit = &v
	return s
}

func (s *ExportStatisticalDataRequest) SetExportType(v string) *ExportStatisticalDataRequest {
	s.ExportType = &v
	return s
}

func (s *ExportStatisticalDataRequest) SetBeginTimeLeftRange(v int64) *ExportStatisticalDataRequest {
	s.BeginTimeLeftRange = &v
	return s
}

func (s *ExportStatisticalDataRequest) SetBeginTimeRightRange(v int64) *ExportStatisticalDataRequest {
	s.BeginTimeRightRange = &v
	return s
}

type ExportStatisticalDataResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExportTaskId *string `json:"ExportTaskId,omitempty" xml:"ExportTaskId,omitempty"`
}

func (s ExportStatisticalDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *ExportStatisticalDataResponseBody) SetRequestId(v string) *ExportStatisticalDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportStatisticalDataResponseBody) SetExportTaskId(v string) *ExportStatisticalDataResponseBody {
	s.ExportTaskId = &v
	return s
}

type ExportStatisticalDataResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExportStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ExportStatisticalDataResponse) SetBody(v *ExportStatisticalDataResponseBody) *ExportStatisticalDataResponse {
	s.Body = v
	return s
}

type ListChatbotInstancesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *ListChatbotInstancesRequest) SetPageNumber(v int32) *ListChatbotInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListChatbotInstancesRequest) SetPageSize(v int32) *ListChatbotInstancesRequest {
	s.PageSize = &v
	return s
}

type ListChatbotInstancesResponseBody struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int64                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Bots       []*ListChatbotInstancesResponseBodyBots `json:"Bots,omitempty" xml:"Bots,omitempty" type:"Repeated"`
}

func (s ListChatbotInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChatbotInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatbotInstancesResponseBody) SetRequestId(v string) *ListChatbotInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetPageNumber(v int32) *ListChatbotInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetTotalCount(v int64) *ListChatbotInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetPageSize(v int64) *ListChatbotInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChatbotInstancesResponseBody) SetBots(v []*ListChatbotInstancesResponseBodyBots) *ListChatbotInstancesResponseBody {
	s.Bots = v
	return s
}

type ListChatbotInstancesResponseBodyBots struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TimeZone     *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	Avatar       *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s ListChatbotInstancesResponseBodyBots) String() string {
	return tea.Prettify(s)
}

func (s ListChatbotInstancesResponseBodyBots) GoString() string {
	return s.String()
}

func (s *ListChatbotInstancesResponseBodyBots) SetInstanceId(v string) *ListChatbotInstancesResponseBodyBots {
	s.InstanceId = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) SetTimeZone(v string) *ListChatbotInstancesResponseBodyBots {
	s.TimeZone = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) SetAvatar(v string) *ListChatbotInstancesResponseBodyBots {
	s.Avatar = &v
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

func (s *ListChatbotInstancesResponseBodyBots) SetIntroduction(v string) *ListChatbotInstancesResponseBodyBots {
	s.Introduction = &v
	return s
}

func (s *ListChatbotInstancesResponseBodyBots) SetCreateTime(v string) *ListChatbotInstancesResponseBodyBots {
	s.CreateTime = &v
	return s
}

type ListChatbotInstancesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListChatbotInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListChatbotInstancesResponse) SetBody(v *ListChatbotInstancesResponseBody) *ListChatbotInstancesResponse {
	s.Body = v
	return s
}

type ListConversationDetailsRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
}

func (s ListConversationDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConversationDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListConversationDetailsRequest) SetInstanceId(v string) *ListConversationDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListConversationDetailsRequest) SetConversationId(v string) *ListConversationDetailsRequest {
	s.ConversationId = &v
	return s
}

type ListConversationDetailsResponseBody struct {
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConversationDetails []*ListConversationDetailsResponseBodyConversationDetails `json:"ConversationDetails,omitempty" xml:"ConversationDetails,omitempty" type:"Repeated"`
}

func (s ListConversationDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConversationDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConversationDetailsResponseBody) SetRequestId(v string) *ListConversationDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConversationDetailsResponseBody) SetConversationDetails(v []*ListConversationDetailsResponseBodyConversationDetails) *ListConversationDetailsResponseBody {
	s.ConversationDetails = v
	return s
}

type ListConversationDetailsResponseBodyConversationDetails struct {
	Action         *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Speaker        *string `json:"Speaker,omitempty" xml:"Speaker,omitempty"`
	CreateTime     *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	ActionParams   *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	SequenceId     *string `json:"SequenceId,omitempty" xml:"SequenceId,omitempty"`
	Utterance      *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
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

func (s *ListConversationDetailsResponseBodyConversationDetails) SetSpeaker(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.Speaker = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetCreateTime(v int64) *ListConversationDetailsResponseBodyConversationDetails {
	s.CreateTime = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetConversationId(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.ConversationId = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetActionParams(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.ActionParams = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetSequenceId(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.SequenceId = &v
	return s
}

func (s *ListConversationDetailsResponseBodyConversationDetails) SetUtterance(v string) *ListConversationDetailsResponseBodyConversationDetails {
	s.Utterance = &v
	return s
}

type ListConversationDetailsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConversationDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListConversationDetailsResponse) SetBody(v *ListConversationDetailsResponseBody) *ListConversationDetailsResponse {
	s.Body = v
	return s
}

type ListConversationsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListConversationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConversationsRequest) GoString() string {
	return s.String()
}

func (s *ListConversationsRequest) SetInstanceId(v string) *ListConversationsRequest {
	s.InstanceId = &v
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

type ListConversationsResponseBody struct {
	TotalCount    *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize      *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber    *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Conversations []*ListConversationsResponseBodyConversations `json:"Conversations,omitempty" xml:"Conversations,omitempty" type:"Repeated"`
}

func (s ListConversationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConversationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConversationsResponseBody) SetTotalCount(v int64) *ListConversationsResponseBody {
	s.TotalCount = &v
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

func (s *ListConversationsResponseBody) SetPageNumber(v int32) *ListConversationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListConversationsResponseBody) SetConversations(v []*ListConversationsResponseBodyConversations) *ListConversationsResponseBody {
	s.Conversations = v
	return s
}

type ListConversationsResponseBodyConversations struct {
	EndTime              *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EffectiveAnswerCount *int32  `json:"EffectiveAnswerCount,omitempty" xml:"EffectiveAnswerCount,omitempty"`
	TransferredToAgent   *bool   `json:"TransferredToAgent,omitempty" xml:"TransferredToAgent,omitempty"`
	BeginTime            *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	SkillGroupId         *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	ConversationId       *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	CallingNumber        *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	UserUtteranceCount   *int32  `json:"UserUtteranceCount,omitempty" xml:"UserUtteranceCount,omitempty"`
}

func (s ListConversationsResponseBodyConversations) String() string {
	return tea.Prettify(s)
}

func (s ListConversationsResponseBodyConversations) GoString() string {
	return s.String()
}

func (s *ListConversationsResponseBodyConversations) SetEndTime(v int64) *ListConversationsResponseBodyConversations {
	s.EndTime = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetEffectiveAnswerCount(v int32) *ListConversationsResponseBodyConversations {
	s.EffectiveAnswerCount = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetTransferredToAgent(v bool) *ListConversationsResponseBodyConversations {
	s.TransferredToAgent = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetBeginTime(v int64) *ListConversationsResponseBodyConversations {
	s.BeginTime = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetSkillGroupId(v string) *ListConversationsResponseBodyConversations {
	s.SkillGroupId = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetConversationId(v string) *ListConversationsResponseBodyConversations {
	s.ConversationId = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetCallingNumber(v string) *ListConversationsResponseBodyConversations {
	s.CallingNumber = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetUserUtteranceCount(v int32) *ListConversationsResponseBodyConversations {
	s.UserUtteranceCount = &v
	return s
}

type ListConversationsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConversationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListConversationsResponse) SetBody(v *ListConversationsResponseBody) *ListConversationsResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

type ListInstancesResponseBody struct {
	Instances  []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
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

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
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

func (s *ListInstancesResponseBody) SetPageNumber(v int32) *ListInstancesResponseBody {
	s.PageNumber = &v
	return s
}

type ListInstancesResponseBodyInstances struct {
	Status               *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	ModifyUserName       *string   `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	ApplicableOperations []*string `json:"ApplicableOperations,omitempty" xml:"ApplicableOperations,omitempty" type:"Repeated"`
	InstanceId           *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name                 *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Concurrency          *int64    `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	ModifyTime           *int64    `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetModifyUserName(v string) *ListInstancesResponseBodyInstances {
	s.ModifyUserName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDescription(v string) *ListInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetApplicableOperations(v []*string) *ListInstancesResponseBodyInstances {
	s.ApplicableOperations = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetName(v string) *ListInstancesResponseBodyInstances {
	s.Name = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetConcurrency(v int64) *ListInstancesResponseBodyInstances {
	s.Concurrency = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetModifyTime(v int64) *ListInstancesResponseBodyInstances {
	s.ModifyTime = &v
	return s
}

type ListInstancesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ModifyGreetingConfigRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GreetingWords *string `json:"GreetingWords,omitempty" xml:"GreetingWords,omitempty"`
	SourceType    *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	IntentTrigger *string `json:"IntentTrigger,omitempty" xml:"IntentTrigger,omitempty"`
}

func (s ModifyGreetingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGreetingConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyGreetingConfigRequest) SetInstanceId(v string) *ModifyGreetingConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyGreetingConfigRequest) SetGreetingWords(v string) *ModifyGreetingConfigRequest {
	s.GreetingWords = &v
	return s
}

func (s *ModifyGreetingConfigRequest) SetSourceType(v string) *ModifyGreetingConfigRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyGreetingConfigRequest) SetIntentTrigger(v string) *ModifyGreetingConfigRequest {
	s.IntentTrigger = &v
	return s
}

type ModifyGreetingConfigResponseBody struct {
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyGreetingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyGreetingConfigResponse) SetBody(v *ModifyGreetingConfigResponseBody) *ModifyGreetingConfigResponse {
	s.Body = v
	return s
}

type ModifyInstanceRequest struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Concurrency       *int64  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	ChatbotInstanceId *string `json:"ChatbotInstanceId,omitempty" xml:"ChatbotInstanceId,omitempty"`
}

func (s ModifyInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRequest) SetInstanceId(v string) *ModifyInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRequest) SetDescription(v string) *ModifyInstanceRequest {
	s.Description = &v
	return s
}

func (s *ModifyInstanceRequest) SetConcurrency(v int64) *ModifyInstanceRequest {
	s.Concurrency = &v
	return s
}

func (s *ModifyInstanceRequest) SetChatbotInstanceId(v string) *ModifyInstanceRequest {
	s.ChatbotInstanceId = &v
	return s
}

type ModifyInstanceResponseBody struct {
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyInstanceResponse) SetBody(v *ModifyInstanceResponseBody) *ModifyInstanceResponse {
	s.Body = v
	return s
}

type ModifySilenceTimeoutConfigRequest struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Prompt            *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	Timeout           *int64  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Threshold         *int32  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	FinalPrompt       *string `json:"FinalPrompt,omitempty" xml:"FinalPrompt,omitempty"`
	FinalAction       *string `json:"FinalAction,omitempty" xml:"FinalAction,omitempty"`
	FinalActionParams *string `json:"FinalActionParams,omitempty" xml:"FinalActionParams,omitempty"`
	SourceType        *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	IntentTrigger     *string `json:"IntentTrigger,omitempty" xml:"IntentTrigger,omitempty"`
}

func (s ModifySilenceTimeoutConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySilenceTimeoutConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifySilenceTimeoutConfigRequest) SetInstanceId(v string) *ModifySilenceTimeoutConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetPrompt(v string) *ModifySilenceTimeoutConfigRequest {
	s.Prompt = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetTimeout(v int64) *ModifySilenceTimeoutConfigRequest {
	s.Timeout = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetThreshold(v int32) *ModifySilenceTimeoutConfigRequest {
	s.Threshold = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetFinalPrompt(v string) *ModifySilenceTimeoutConfigRequest {
	s.FinalPrompt = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetFinalAction(v string) *ModifySilenceTimeoutConfigRequest {
	s.FinalAction = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetFinalActionParams(v string) *ModifySilenceTimeoutConfigRequest {
	s.FinalActionParams = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetSourceType(v string) *ModifySilenceTimeoutConfigRequest {
	s.SourceType = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetIntentTrigger(v string) *ModifySilenceTimeoutConfigRequest {
	s.IntentTrigger = &v
	return s
}

type ModifySilenceTimeoutConfigResponseBody struct {
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySilenceTimeoutConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifySilenceTimeoutConfigResponse) SetBody(v *ModifySilenceTimeoutConfigResponseBody) *ModifySilenceTimeoutConfigResponse {
	s.Body = v
	return s
}

type ModifyTTSConfigRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Voice      *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	SpeechRate *string `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Volume     *string `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s ModifyTTSConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTTSConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyTTSConfigRequest) SetInstanceId(v string) *ModifyTTSConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetVoice(v string) *ModifyTTSConfigRequest {
	s.Voice = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetSpeechRate(v string) *ModifyTTSConfigRequest {
	s.SpeechRate = &v
	return s
}

func (s *ModifyTTSConfigRequest) SetVolume(v string) *ModifyTTSConfigRequest {
	s.Volume = &v
	return s
}

type ModifyTTSConfigResponseBody struct {
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyTTSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyTTSConfigResponse) SetBody(v *ModifyTTSConfigResponseBody) *ModifyTTSConfigResponse {
	s.Body = v
	return s
}

type ModifyUnrecognizingConfigRequest struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Prompt            *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	Threshold         *int32  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	FinalPrompt       *string `json:"FinalPrompt,omitempty" xml:"FinalPrompt,omitempty"`
	FinalAction       *string `json:"FinalAction,omitempty" xml:"FinalAction,omitempty"`
	FinalActionParams *string `json:"FinalActionParams,omitempty" xml:"FinalActionParams,omitempty"`
}

func (s ModifyUnrecognizingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUnrecognizingConfigRequest) GoString() string {
	return s.String()
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

func (s *ModifyUnrecognizingConfigRequest) SetFinalPrompt(v string) *ModifyUnrecognizingConfigRequest {
	s.FinalPrompt = &v
	return s
}

func (s *ModifyUnrecognizingConfigRequest) SetFinalAction(v string) *ModifyUnrecognizingConfigRequest {
	s.FinalAction = &v
	return s
}

func (s *ModifyUnrecognizingConfigRequest) SetFinalActionParams(v string) *ModifyUnrecognizingConfigRequest {
	s.FinalActionParams = &v
	return s
}

type ModifyUnrecognizingConfigResponseBody struct {
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUnrecognizingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyUnrecognizingConfigResponse) SetBody(v *ModifyUnrecognizingConfigResponseBody) *ModifyUnrecognizingConfigResponse {
	s.Body = v
	return s
}

type QueryConversationsRequest struct {
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CallingNumber       *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	BeginTimeLeftRange  *int64  `json:"BeginTimeLeftRange,omitempty" xml:"BeginTimeLeftRange,omitempty"`
	BeginTimeRightRange *int64  `json:"BeginTimeRightRange,omitempty" xml:"BeginTimeRightRange,omitempty"`
	PageNumber          *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryConversationsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationsRequest) GoString() string {
	return s.String()
}

func (s *QueryConversationsRequest) SetInstanceId(v string) *QueryConversationsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryConversationsRequest) SetCallingNumber(v string) *QueryConversationsRequest {
	s.CallingNumber = &v
	return s
}

func (s *QueryConversationsRequest) SetBeginTimeLeftRange(v int64) *QueryConversationsRequest {
	s.BeginTimeLeftRange = &v
	return s
}

func (s *QueryConversationsRequest) SetBeginTimeRightRange(v int64) *QueryConversationsRequest {
	s.BeginTimeRightRange = &v
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
	TotalCount    *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize      *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber    *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Conversations []*QueryConversationsResponseBodyConversations `json:"Conversations,omitempty" xml:"Conversations,omitempty" type:"Repeated"`
}

func (s QueryConversationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConversationsResponseBody) SetTotalCount(v int64) *QueryConversationsResponseBody {
	s.TotalCount = &v
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

func (s *QueryConversationsResponseBody) SetPageNumber(v int32) *QueryConversationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryConversationsResponseBody) SetConversations(v []*QueryConversationsResponseBodyConversations) *QueryConversationsResponseBody {
	s.Conversations = v
	return s
}

type QueryConversationsResponseBodyConversations struct {
	EndTime              *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EffectiveAnswerCount *int32  `json:"EffectiveAnswerCount,omitempty" xml:"EffectiveAnswerCount,omitempty"`
	TransferredToAgent   *bool   `json:"TransferredToAgent,omitempty" xml:"TransferredToAgent,omitempty"`
	BeginTime            *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	SkillGroupId         *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	ConversationId       *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	CallingNumber        *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	UserUtteranceCount   *int32  `json:"UserUtteranceCount,omitempty" xml:"UserUtteranceCount,omitempty"`
}

func (s QueryConversationsResponseBodyConversations) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationsResponseBodyConversations) GoString() string {
	return s.String()
}

func (s *QueryConversationsResponseBodyConversations) SetEndTime(v int64) *QueryConversationsResponseBodyConversations {
	s.EndTime = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetEffectiveAnswerCount(v int32) *QueryConversationsResponseBodyConversations {
	s.EffectiveAnswerCount = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetTransferredToAgent(v bool) *QueryConversationsResponseBodyConversations {
	s.TransferredToAgent = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetBeginTime(v int64) *QueryConversationsResponseBodyConversations {
	s.BeginTime = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetSkillGroupId(v string) *QueryConversationsResponseBodyConversations {
	s.SkillGroupId = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetConversationId(v string) *QueryConversationsResponseBodyConversations {
	s.ConversationId = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetCallingNumber(v string) *QueryConversationsResponseBodyConversations {
	s.CallingNumber = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetUserUtteranceCount(v int32) *QueryConversationsResponseBodyConversations {
	s.UserUtteranceCount = &v
	return s
}

type QueryConversationsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryConversationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryConversationsResponse) SetBody(v *QueryConversationsResponseBody) *QueryConversationsResponse {
	s.Body = v
	return s
}

type SaveRecordingRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FileName       *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FilePath       *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SaveRecordingRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveRecordingRequest) GoString() string {
	return s.String()
}

func (s *SaveRecordingRequest) SetInstanceId(v string) *SaveRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveRecordingRequest) SetConversationId(v string) *SaveRecordingRequest {
	s.ConversationId = &v
	return s
}

func (s *SaveRecordingRequest) SetStartTime(v int64) *SaveRecordingRequest {
	s.StartTime = &v
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

func (s *SaveRecordingRequest) SetType(v string) *SaveRecordingRequest {
	s.Type = &v
	return s
}

type SaveRecordingResponseBody struct {
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SaveRecordingResponse) SetBody(v *SaveRecordingResponseBody) *SaveRecordingResponse {
	s.Body = v
	return s
}

type SilenceTimeoutRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	InitialContext *string `json:"InitialContext,omitempty" xml:"InitialContext,omitempty"`
}

func (s SilenceTimeoutRequest) String() string {
	return tea.Prettify(s)
}

func (s SilenceTimeoutRequest) GoString() string {
	return s.String()
}

func (s *SilenceTimeoutRequest) SetInstanceId(v string) *SilenceTimeoutRequest {
	s.InstanceId = &v
	return s
}

func (s *SilenceTimeoutRequest) SetConversationId(v string) *SilenceTimeoutRequest {
	s.ConversationId = &v
	return s
}

func (s *SilenceTimeoutRequest) SetInitialContext(v string) *SilenceTimeoutRequest {
	s.InitialContext = &v
	return s
}

type SilenceTimeoutResponseBody struct {
	Action        *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Interruptible *bool   `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ActionParams  *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	TextResponse  *string `json:"TextResponse,omitempty" xml:"TextResponse,omitempty"`
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

func (s *SilenceTimeoutResponseBody) SetInterruptible(v bool) *SilenceTimeoutResponseBody {
	s.Interruptible = &v
	return s
}

func (s *SilenceTimeoutResponseBody) SetRequestId(v string) *SilenceTimeoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *SilenceTimeoutResponseBody) SetActionParams(v string) *SilenceTimeoutResponseBody {
	s.ActionParams = &v
	return s
}

func (s *SilenceTimeoutResponseBody) SetTextResponse(v string) *SilenceTimeoutResponseBody {
	s.TextResponse = &v
	return s
}

type SilenceTimeoutResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SilenceTimeoutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) AssociateChatbotInstanceWithOptions(request *AssociateChatbotInstanceRequest, runtime *util.RuntimeOptions) (_result *AssociateChatbotInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssociateChatbotInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssociateChatbotInstance"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) AuditTTSVoiceWithOptions(request *AuditTTSVoiceRequest, runtime *util.RuntimeOptions) (_result *AuditTTSVoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AuditTTSVoiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AuditTTSVoice"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) BeginDialogueWithOptions(request *BeginDialogueRequest, runtime *util.RuntimeOptions) (_result *BeginDialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BeginDialogueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BeginDialogue"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CollectedNumberWithOptions(request *CollectedNumberRequest, runtime *util.RuntimeOptions) (_result *CollectedNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CollectedNumberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CollectedNumber"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateInstance"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DebugBeginDialogueWithOptions(request *DebugBeginDialogueRequest, runtime *util.RuntimeOptions) (_result *DebugBeginDialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DebugBeginDialogueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DebugBeginDialogue"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DebugCollectedNumberWithOptions(request *DebugCollectedNumberRequest, runtime *util.RuntimeOptions) (_result *DebugCollectedNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DebugCollectedNumberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DebugCollectedNumber"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DebugDialogueWithOptions(request *DebugDialogueRequest, runtime *util.RuntimeOptions) (_result *DebugDialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DebugDialogueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DebugDialogue"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteInstance"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeConversationWithOptions(request *DescribeConversationRequest, runtime *util.RuntimeOptions) (_result *DescribeConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeConversationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeConversation"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeConversationContextWithOptions(request *DescribeConversationContextRequest, runtime *util.RuntimeOptions) (_result *DescribeConversationContextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeConversationContextResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeConversationContext"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeExportProgressWithOptions(request *DescribeExportProgressRequest, runtime *util.RuntimeOptions) (_result *DescribeExportProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeExportProgressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeExportProgress"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstance"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeNavigationConfigWithOptions(request *DescribeNavigationConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeNavigationConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeNavigationConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNavigationConfig"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeRecordingWithOptions(request *DescribeRecordingRequest, runtime *util.RuntimeOptions) (_result *DescribeRecordingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeRecordingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRecording"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeStatisticalDataWithOptions(request *DescribeStatisticalDataRequest, runtime *util.RuntimeOptions) (_result *DescribeStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeStatisticalDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeStatisticalData"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescribeTTSConfigWithOptions(request *DescribeTTSConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeTTSConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeTTSConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTTSConfig"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DialogueWithOptions(request *DialogueRequest, runtime *util.RuntimeOptions) (_result *DialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DialogueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("Dialogue"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DisableInstanceWithOptions(request *DisableInstanceRequest, runtime *util.RuntimeOptions) (_result *DisableInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableInstance"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) EnableInstanceWithOptions(request *EnableInstanceRequest, runtime *util.RuntimeOptions) (_result *EnableInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableInstance"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) EndDialogueWithOptions(request *EndDialogueRequest, runtime *util.RuntimeOptions) (_result *EndDialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EndDialogueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EndDialogue"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ExportConversationDetailsWithOptions(request *ExportConversationDetailsRequest, runtime *util.RuntimeOptions) (_result *ExportConversationDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExportConversationDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExportConversationDetails"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ExportStatisticalDataWithOptions(request *ExportStatisticalDataRequest, runtime *util.RuntimeOptions) (_result *ExportStatisticalDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExportStatisticalDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExportStatisticalData"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListChatbotInstancesWithOptions(request *ListChatbotInstancesRequest, runtime *util.RuntimeOptions) (_result *ListChatbotInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListChatbotInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListChatbotInstances"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListConversationDetailsWithOptions(request *ListConversationDetailsRequest, runtime *util.RuntimeOptions) (_result *ListConversationDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListConversationDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListConversationDetails"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListConversationsWithOptions(request *ListConversationsRequest, runtime *util.RuntimeOptions) (_result *ListConversationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListConversationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListConversations"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstances"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ModifyGreetingConfigWithOptions(request *ModifyGreetingConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyGreetingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyGreetingConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyGreetingConfig"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ModifyInstanceWithOptions(request *ModifyInstanceRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstance"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ModifySilenceTimeoutConfigWithOptions(request *ModifySilenceTimeoutConfigRequest, runtime *util.RuntimeOptions) (_result *ModifySilenceTimeoutConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySilenceTimeoutConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySilenceTimeoutConfig"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ModifyTTSConfigWithOptions(request *ModifyTTSConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyTTSConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyTTSConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyTTSConfig"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ModifyUnrecognizingConfigWithOptions(request *ModifyUnrecognizingConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyUnrecognizingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyUnrecognizingConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyUnrecognizingConfig"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) QueryConversationsWithOptions(request *QueryConversationsRequest, runtime *util.RuntimeOptions) (_result *QueryConversationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &QueryConversationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryConversations"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) SaveRecordingWithOptions(request *SaveRecordingRequest, runtime *util.RuntimeOptions) (_result *SaveRecordingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveRecordingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveRecording"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) SilenceTimeoutWithOptions(request *SilenceTimeoutRequest, runtime *util.RuntimeOptions) (_result *SilenceTimeoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SilenceTimeoutResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SilenceTimeout"), tea.String("2018-06-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
