// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPaChatCompletionStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedTokenIds(v []*int64) *PaChatCompletionStreamRequest
	GetAllowedTokenIds() []*int64
	SetBadWords(v []*string) *PaChatCompletionStreamRequest
	GetBadWords() []*string
	SetChatTemplateKwargs(v *PaChatCompletionStreamRequestChatTemplateKwargs) *PaChatCompletionStreamRequest
	GetChatTemplateKwargs() *PaChatCompletionStreamRequestChatTemplateKwargs
	SetFrequencyPenalty(v float64) *PaChatCompletionStreamRequest
	GetFrequencyPenalty() *float64
	SetIgnoreEos(v bool) *PaChatCompletionStreamRequest
	GetIgnoreEos() *bool
	SetIncludeReasoning(v bool) *PaChatCompletionStreamRequest
	GetIncludeReasoning() *bool
	SetLogprobs(v bool) *PaChatCompletionStreamRequest
	GetLogprobs() *bool
	SetMaxCompletionTokens(v int64) *PaChatCompletionStreamRequest
	GetMaxCompletionTokens() *int64
	SetMaxTokens(v int64) *PaChatCompletionStreamRequest
	GetMaxTokens() *int64
	SetMessages(v []*PaChatCompletionStreamRequestMessages) *PaChatCompletionStreamRequest
	GetMessages() []*PaChatCompletionStreamRequestMessages
	SetMinP(v float64) *PaChatCompletionStreamRequest
	GetMinP() *float64
	SetMinTokens(v int64) *PaChatCompletionStreamRequest
	GetMinTokens() *int64
	SetMmProcessorKwargs(v *PaChatCompletionStreamRequestMmProcessorKwargs) *PaChatCompletionStreamRequest
	GetMmProcessorKwargs() *PaChatCompletionStreamRequestMmProcessorKwargs
	SetModel(v string) *PaChatCompletionStreamRequest
	GetModel() *string
	SetN(v int64) *PaChatCompletionStreamRequest
	GetN() *int64
	SetParallelToolCalls(v bool) *PaChatCompletionStreamRequest
	GetParallelToolCalls() *bool
	SetPresencePenalty(v float64) *PaChatCompletionStreamRequest
	GetPresencePenalty() *float64
	SetPromptLogprobs(v int64) *PaChatCompletionStreamRequest
	GetPromptLogprobs() *int64
	SetReasoningEffort(v string) *PaChatCompletionStreamRequest
	GetReasoningEffort() *string
	SetRepetitionPenalty(v float64) *PaChatCompletionStreamRequest
	GetRepetitionPenalty() *float64
	SetResponseFormat(v *PaChatCompletionStreamRequestResponseFormat) *PaChatCompletionStreamRequest
	GetResponseFormat() *PaChatCompletionStreamRequestResponseFormat
	SetSeed(v int64) *PaChatCompletionStreamRequest
	GetSeed() *int64
	SetSkipSpecialTokens(v bool) *PaChatCompletionStreamRequest
	GetSkipSpecialTokens() *bool
	SetStop(v []*string) *PaChatCompletionStreamRequest
	GetStop() []*string
	SetStopTokenIds(v []*int64) *PaChatCompletionStreamRequest
	GetStopTokenIds() []*int64
	SetStream(v bool) *PaChatCompletionStreamRequest
	GetStream() *bool
	SetStreamOptions(v *PaChatCompletionStreamRequestStreamOptions) *PaChatCompletionStreamRequest
	GetStreamOptions() *PaChatCompletionStreamRequestStreamOptions
	SetStructuredOutputs(v *PaChatCompletionStreamRequestStructuredOutputs) *PaChatCompletionStreamRequest
	GetStructuredOutputs() *PaChatCompletionStreamRequestStructuredOutputs
	SetTemperature(v float64) *PaChatCompletionStreamRequest
	GetTemperature() *float64
	SetToolChoice(v string) *PaChatCompletionStreamRequest
	GetToolChoice() *string
	SetTools(v []*PaChatCompletionStreamRequestTools) *PaChatCompletionStreamRequest
	GetTools() []*PaChatCompletionStreamRequestTools
	SetTopK(v int64) *PaChatCompletionStreamRequest
	GetTopK() *int64
	SetTopLogprobs(v int64) *PaChatCompletionStreamRequest
	GetTopLogprobs() *int64
	SetTopP(v float64) *PaChatCompletionStreamRequest
	GetTopP() *float64
}

type PaChatCompletionStreamRequest struct {
	// example:
	//
	// [10,11]
	AllowedTokenIds []*int64 `json:"allowedTokenIds,omitempty" xml:"allowedTokenIds,omitempty" type:"Repeated"`
	// example:
	//
	// ["blocked"]
	BadWords           []*string                                        `json:"badWords,omitempty" xml:"badWords,omitempty" type:"Repeated"`
	ChatTemplateKwargs *PaChatCompletionStreamRequestChatTemplateKwargs `json:"chatTemplateKwargs,omitempty" xml:"chatTemplateKwargs,omitempty" type:"Struct"`
	// example:
	//
	// 0.0
	FrequencyPenalty *float64 `json:"frequencyPenalty,omitempty" xml:"frequencyPenalty,omitempty"`
	// example:
	//
	// false
	IgnoreEos *bool `json:"ignoreEos,omitempty" xml:"ignoreEos,omitempty"`
	// example:
	//
	// true
	IncludeReasoning *bool `json:"includeReasoning,omitempty" xml:"includeReasoning,omitempty"`
	// example:
	//
	// false
	Logprobs *bool `json:"logprobs,omitempty" xml:"logprobs,omitempty"`
	// example:
	//
	// 2048
	MaxCompletionTokens *int64 `json:"maxCompletionTokens,omitempty" xml:"maxCompletionTokens,omitempty"`
	// example:
	//
	// 2048
	MaxTokens *int64 `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	// This parameter is required.
	Messages []*PaChatCompletionStreamRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// example:
	//
	// 0.05
	MinP *float64 `json:"minP,omitempty" xml:"minP,omitempty"`
	// example:
	//
	// 1
	MinTokens         *int64                                          `json:"minTokens,omitempty" xml:"minTokens,omitempty"`
	MmProcessorKwargs *PaChatCompletionStreamRequestMmProcessorKwargs `json:"mmProcessorKwargs,omitempty" xml:"mmProcessorKwargs,omitempty" type:"Struct"`
	Model             *string                                         `json:"model,omitempty" xml:"model,omitempty"`
	// example:
	//
	// 1
	N *int64 `json:"n,omitempty" xml:"n,omitempty"`
	// example:
	//
	// false
	ParallelToolCalls *bool `json:"parallelToolCalls,omitempty" xml:"parallelToolCalls,omitempty"`
	// example:
	//
	// 0.0
	PresencePenalty *float64 `json:"presencePenalty,omitempty" xml:"presencePenalty,omitempty"`
	// example:
	//
	// 2
	PromptLogprobs *int64 `json:"promptLogprobs,omitempty" xml:"promptLogprobs,omitempty"`
	// example:
	//
	// high
	ReasoningEffort *string `json:"reasoningEffort,omitempty" xml:"reasoningEffort,omitempty"`
	// example:
	//
	// 1.1
	RepetitionPenalty *float64                                     `json:"repetitionPenalty,omitempty" xml:"repetitionPenalty,omitempty"`
	ResponseFormat    *PaChatCompletionStreamRequestResponseFormat `json:"responseFormat,omitempty" xml:"responseFormat,omitempty" type:"Struct"`
	// example:
	//
	// 42
	Seed *int64 `json:"seed,omitempty" xml:"seed,omitempty"`
	// example:
	//
	// true
	SkipSpecialTokens *bool `json:"skipSpecialTokens,omitempty" xml:"skipSpecialTokens,omitempty"`
	// example:
	//
	// ["END"]
	Stop []*string `json:"stop,omitempty" xml:"stop,omitempty" type:"Repeated"`
	// example:
	//
	// [1,2]
	StopTokenIds []*int64 `json:"stopTokenIds,omitempty" xml:"stopTokenIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Stream            *bool                                           `json:"stream,omitempty" xml:"stream,omitempty"`
	StreamOptions     *PaChatCompletionStreamRequestStreamOptions     `json:"streamOptions,omitempty" xml:"streamOptions,omitempty" type:"Struct"`
	StructuredOutputs *PaChatCompletionStreamRequestStructuredOutputs `json:"structuredOutputs,omitempty" xml:"structuredOutputs,omitempty" type:"Struct"`
	// example:
	//
	// 0.2
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// example:
	//
	// auto
	ToolChoice *string                               `json:"toolChoice,omitempty" xml:"toolChoice,omitempty"`
	Tools      []*PaChatCompletionStreamRequestTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	TopK *int64 `json:"topK,omitempty" xml:"topK,omitempty"`
	// example:
	//
	// 3
	TopLogprobs *int64 `json:"topLogprobs,omitempty" xml:"topLogprobs,omitempty"`
	// example:
	//
	// 0.9
	TopP *float64 `json:"topP,omitempty" xml:"topP,omitempty"`
}

func (s PaChatCompletionStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequest) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequest) GetAllowedTokenIds() []*int64 {
	return s.AllowedTokenIds
}

func (s *PaChatCompletionStreamRequest) GetBadWords() []*string {
	return s.BadWords
}

func (s *PaChatCompletionStreamRequest) GetChatTemplateKwargs() *PaChatCompletionStreamRequestChatTemplateKwargs {
	return s.ChatTemplateKwargs
}

func (s *PaChatCompletionStreamRequest) GetFrequencyPenalty() *float64 {
	return s.FrequencyPenalty
}

func (s *PaChatCompletionStreamRequest) GetIgnoreEos() *bool {
	return s.IgnoreEos
}

func (s *PaChatCompletionStreamRequest) GetIncludeReasoning() *bool {
	return s.IncludeReasoning
}

func (s *PaChatCompletionStreamRequest) GetLogprobs() *bool {
	return s.Logprobs
}

func (s *PaChatCompletionStreamRequest) GetMaxCompletionTokens() *int64 {
	return s.MaxCompletionTokens
}

func (s *PaChatCompletionStreamRequest) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *PaChatCompletionStreamRequest) GetMessages() []*PaChatCompletionStreamRequestMessages {
	return s.Messages
}

func (s *PaChatCompletionStreamRequest) GetMinP() *float64 {
	return s.MinP
}

func (s *PaChatCompletionStreamRequest) GetMinTokens() *int64 {
	return s.MinTokens
}

func (s *PaChatCompletionStreamRequest) GetMmProcessorKwargs() *PaChatCompletionStreamRequestMmProcessorKwargs {
	return s.MmProcessorKwargs
}

func (s *PaChatCompletionStreamRequest) GetModel() *string {
	return s.Model
}

func (s *PaChatCompletionStreamRequest) GetN() *int64 {
	return s.N
}

func (s *PaChatCompletionStreamRequest) GetParallelToolCalls() *bool {
	return s.ParallelToolCalls
}

func (s *PaChatCompletionStreamRequest) GetPresencePenalty() *float64 {
	return s.PresencePenalty
}

func (s *PaChatCompletionStreamRequest) GetPromptLogprobs() *int64 {
	return s.PromptLogprobs
}

func (s *PaChatCompletionStreamRequest) GetReasoningEffort() *string {
	return s.ReasoningEffort
}

func (s *PaChatCompletionStreamRequest) GetRepetitionPenalty() *float64 {
	return s.RepetitionPenalty
}

func (s *PaChatCompletionStreamRequest) GetResponseFormat() *PaChatCompletionStreamRequestResponseFormat {
	return s.ResponseFormat
}

func (s *PaChatCompletionStreamRequest) GetSeed() *int64 {
	return s.Seed
}

func (s *PaChatCompletionStreamRequest) GetSkipSpecialTokens() *bool {
	return s.SkipSpecialTokens
}

func (s *PaChatCompletionStreamRequest) GetStop() []*string {
	return s.Stop
}

func (s *PaChatCompletionStreamRequest) GetStopTokenIds() []*int64 {
	return s.StopTokenIds
}

func (s *PaChatCompletionStreamRequest) GetStream() *bool {
	return s.Stream
}

func (s *PaChatCompletionStreamRequest) GetStreamOptions() *PaChatCompletionStreamRequestStreamOptions {
	return s.StreamOptions
}

func (s *PaChatCompletionStreamRequest) GetStructuredOutputs() *PaChatCompletionStreamRequestStructuredOutputs {
	return s.StructuredOutputs
}

func (s *PaChatCompletionStreamRequest) GetTemperature() *float64 {
	return s.Temperature
}

func (s *PaChatCompletionStreamRequest) GetToolChoice() *string {
	return s.ToolChoice
}

func (s *PaChatCompletionStreamRequest) GetTools() []*PaChatCompletionStreamRequestTools {
	return s.Tools
}

func (s *PaChatCompletionStreamRequest) GetTopK() *int64 {
	return s.TopK
}

func (s *PaChatCompletionStreamRequest) GetTopLogprobs() *int64 {
	return s.TopLogprobs
}

func (s *PaChatCompletionStreamRequest) GetTopP() *float64 {
	return s.TopP
}

func (s *PaChatCompletionStreamRequest) SetAllowedTokenIds(v []*int64) *PaChatCompletionStreamRequest {
	s.AllowedTokenIds = v
	return s
}

func (s *PaChatCompletionStreamRequest) SetBadWords(v []*string) *PaChatCompletionStreamRequest {
	s.BadWords = v
	return s
}

func (s *PaChatCompletionStreamRequest) SetChatTemplateKwargs(v *PaChatCompletionStreamRequestChatTemplateKwargs) *PaChatCompletionStreamRequest {
	s.ChatTemplateKwargs = v
	return s
}

func (s *PaChatCompletionStreamRequest) SetFrequencyPenalty(v float64) *PaChatCompletionStreamRequest {
	s.FrequencyPenalty = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetIgnoreEos(v bool) *PaChatCompletionStreamRequest {
	s.IgnoreEos = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetIncludeReasoning(v bool) *PaChatCompletionStreamRequest {
	s.IncludeReasoning = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetLogprobs(v bool) *PaChatCompletionStreamRequest {
	s.Logprobs = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetMaxCompletionTokens(v int64) *PaChatCompletionStreamRequest {
	s.MaxCompletionTokens = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetMaxTokens(v int64) *PaChatCompletionStreamRequest {
	s.MaxTokens = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetMessages(v []*PaChatCompletionStreamRequestMessages) *PaChatCompletionStreamRequest {
	s.Messages = v
	return s
}

func (s *PaChatCompletionStreamRequest) SetMinP(v float64) *PaChatCompletionStreamRequest {
	s.MinP = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetMinTokens(v int64) *PaChatCompletionStreamRequest {
	s.MinTokens = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetMmProcessorKwargs(v *PaChatCompletionStreamRequestMmProcessorKwargs) *PaChatCompletionStreamRequest {
	s.MmProcessorKwargs = v
	return s
}

func (s *PaChatCompletionStreamRequest) SetModel(v string) *PaChatCompletionStreamRequest {
	s.Model = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetN(v int64) *PaChatCompletionStreamRequest {
	s.N = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetParallelToolCalls(v bool) *PaChatCompletionStreamRequest {
	s.ParallelToolCalls = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetPresencePenalty(v float64) *PaChatCompletionStreamRequest {
	s.PresencePenalty = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetPromptLogprobs(v int64) *PaChatCompletionStreamRequest {
	s.PromptLogprobs = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetReasoningEffort(v string) *PaChatCompletionStreamRequest {
	s.ReasoningEffort = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetRepetitionPenalty(v float64) *PaChatCompletionStreamRequest {
	s.RepetitionPenalty = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetResponseFormat(v *PaChatCompletionStreamRequestResponseFormat) *PaChatCompletionStreamRequest {
	s.ResponseFormat = v
	return s
}

func (s *PaChatCompletionStreamRequest) SetSeed(v int64) *PaChatCompletionStreamRequest {
	s.Seed = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetSkipSpecialTokens(v bool) *PaChatCompletionStreamRequest {
	s.SkipSpecialTokens = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetStop(v []*string) *PaChatCompletionStreamRequest {
	s.Stop = v
	return s
}

func (s *PaChatCompletionStreamRequest) SetStopTokenIds(v []*int64) *PaChatCompletionStreamRequest {
	s.StopTokenIds = v
	return s
}

func (s *PaChatCompletionStreamRequest) SetStream(v bool) *PaChatCompletionStreamRequest {
	s.Stream = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetStreamOptions(v *PaChatCompletionStreamRequestStreamOptions) *PaChatCompletionStreamRequest {
	s.StreamOptions = v
	return s
}

func (s *PaChatCompletionStreamRequest) SetStructuredOutputs(v *PaChatCompletionStreamRequestStructuredOutputs) *PaChatCompletionStreamRequest {
	s.StructuredOutputs = v
	return s
}

func (s *PaChatCompletionStreamRequest) SetTemperature(v float64) *PaChatCompletionStreamRequest {
	s.Temperature = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetToolChoice(v string) *PaChatCompletionStreamRequest {
	s.ToolChoice = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetTools(v []*PaChatCompletionStreamRequestTools) *PaChatCompletionStreamRequest {
	s.Tools = v
	return s
}

func (s *PaChatCompletionStreamRequest) SetTopK(v int64) *PaChatCompletionStreamRequest {
	s.TopK = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetTopLogprobs(v int64) *PaChatCompletionStreamRequest {
	s.TopLogprobs = &v
	return s
}

func (s *PaChatCompletionStreamRequest) SetTopP(v float64) *PaChatCompletionStreamRequest {
	s.TopP = &v
	return s
}

func (s *PaChatCompletionStreamRequest) Validate() error {
	if s.ChatTemplateKwargs != nil {
		if err := s.ChatTemplateKwargs.Validate(); err != nil {
			return err
		}
	}
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MmProcessorKwargs != nil {
		if err := s.MmProcessorKwargs.Validate(); err != nil {
			return err
		}
	}
	if s.ResponseFormat != nil {
		if err := s.ResponseFormat.Validate(); err != nil {
			return err
		}
	}
	if s.StreamOptions != nil {
		if err := s.StreamOptions.Validate(); err != nil {
			return err
		}
	}
	if s.StructuredOutputs != nil {
		if err := s.StructuredOutputs.Validate(); err != nil {
			return err
		}
	}
	if s.Tools != nil {
		for _, item := range s.Tools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PaChatCompletionStreamRequestChatTemplateKwargs struct {
	// example:
	//
	// true
	EnableThinking *bool `json:"enableThinking,omitempty" xml:"enableThinking,omitempty"`
	// example:
	//
	// true
	PreserveThinking *bool `json:"preserveThinking,omitempty" xml:"preserveThinking,omitempty"`
}

func (s PaChatCompletionStreamRequestChatTemplateKwargs) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequestChatTemplateKwargs) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequestChatTemplateKwargs) GetEnableThinking() *bool {
	return s.EnableThinking
}

func (s *PaChatCompletionStreamRequestChatTemplateKwargs) GetPreserveThinking() *bool {
	return s.PreserveThinking
}

func (s *PaChatCompletionStreamRequestChatTemplateKwargs) SetEnableThinking(v bool) *PaChatCompletionStreamRequestChatTemplateKwargs {
	s.EnableThinking = &v
	return s
}

func (s *PaChatCompletionStreamRequestChatTemplateKwargs) SetPreserveThinking(v bool) *PaChatCompletionStreamRequestChatTemplateKwargs {
	s.PreserveThinking = &v
	return s
}

func (s *PaChatCompletionStreamRequestChatTemplateKwargs) Validate() error {
	return dara.Validate(s)
}

type PaChatCompletionStreamRequestMessages struct {
	// example:
	//
	// 请分析图片内容
	Content []*PaChatCompletionStreamRequestMessagesContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// call_weather_1
	ToolCallId *string                                           `json:"toolCallId,omitempty" xml:"toolCallId,omitempty"`
	ToolCalls  []*PaChatCompletionStreamRequestMessagesToolCalls `json:"toolCalls,omitempty" xml:"toolCalls,omitempty" type:"Repeated"`
}

func (s PaChatCompletionStreamRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequestMessages) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequestMessages) GetContent() []*PaChatCompletionStreamRequestMessagesContent {
	return s.Content
}

func (s *PaChatCompletionStreamRequestMessages) GetRole() *string {
	return s.Role
}

func (s *PaChatCompletionStreamRequestMessages) GetToolCallId() *string {
	return s.ToolCallId
}

func (s *PaChatCompletionStreamRequestMessages) GetToolCalls() []*PaChatCompletionStreamRequestMessagesToolCalls {
	return s.ToolCalls
}

func (s *PaChatCompletionStreamRequestMessages) SetContent(v []*PaChatCompletionStreamRequestMessagesContent) *PaChatCompletionStreamRequestMessages {
	s.Content = v
	return s
}

func (s *PaChatCompletionStreamRequestMessages) SetRole(v string) *PaChatCompletionStreamRequestMessages {
	s.Role = &v
	return s
}

func (s *PaChatCompletionStreamRequestMessages) SetToolCallId(v string) *PaChatCompletionStreamRequestMessages {
	s.ToolCallId = &v
	return s
}

func (s *PaChatCompletionStreamRequestMessages) SetToolCalls(v []*PaChatCompletionStreamRequestMessagesToolCalls) *PaChatCompletionStreamRequestMessages {
	s.ToolCalls = v
	return s
}

func (s *PaChatCompletionStreamRequestMessages) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ToolCalls != nil {
		for _, item := range s.ToolCalls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PaChatCompletionStreamRequestMessagesContent struct {
	ImageUrl *PaChatCompletionStreamRequestMessagesContentImageUrl `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Struct"`
	// example:
	//
	// 请分析这张图片
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PaChatCompletionStreamRequestMessagesContent) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequestMessagesContent) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequestMessagesContent) GetImageUrl() *PaChatCompletionStreamRequestMessagesContentImageUrl {
	return s.ImageUrl
}

func (s *PaChatCompletionStreamRequestMessagesContent) GetText() *string {
	return s.Text
}

func (s *PaChatCompletionStreamRequestMessagesContent) GetType() *string {
	return s.Type
}

func (s *PaChatCompletionStreamRequestMessagesContent) SetImageUrl(v *PaChatCompletionStreamRequestMessagesContentImageUrl) *PaChatCompletionStreamRequestMessagesContent {
	s.ImageUrl = v
	return s
}

func (s *PaChatCompletionStreamRequestMessagesContent) SetText(v string) *PaChatCompletionStreamRequestMessagesContent {
	s.Text = &v
	return s
}

func (s *PaChatCompletionStreamRequestMessagesContent) SetType(v string) *PaChatCompletionStreamRequestMessagesContent {
	s.Type = &v
	return s
}

func (s *PaChatCompletionStreamRequestMessagesContent) Validate() error {
	if s.ImageUrl != nil {
		if err := s.ImageUrl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PaChatCompletionStreamRequestMessagesContentImageUrl struct {
	// example:
	//
	// https://example.com/image.png
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PaChatCompletionStreamRequestMessagesContentImageUrl) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequestMessagesContentImageUrl) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequestMessagesContentImageUrl) GetUrl() *string {
	return s.Url
}

func (s *PaChatCompletionStreamRequestMessagesContentImageUrl) SetUrl(v string) *PaChatCompletionStreamRequestMessagesContentImageUrl {
	s.Url = &v
	return s
}

func (s *PaChatCompletionStreamRequestMessagesContentImageUrl) Validate() error {
	return dara.Validate(s)
}

type PaChatCompletionStreamRequestMessagesToolCalls struct {
	Function *PaChatCompletionStreamRequestMessagesToolCallsFunction `json:"function,omitempty" xml:"function,omitempty" type:"Struct"`
	// example:
	//
	// call_weather_1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// function
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PaChatCompletionStreamRequestMessagesToolCalls) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequestMessagesToolCalls) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequestMessagesToolCalls) GetFunction() *PaChatCompletionStreamRequestMessagesToolCallsFunction {
	return s.Function
}

func (s *PaChatCompletionStreamRequestMessagesToolCalls) GetId() *string {
	return s.Id
}

func (s *PaChatCompletionStreamRequestMessagesToolCalls) GetType() *string {
	return s.Type
}

func (s *PaChatCompletionStreamRequestMessagesToolCalls) SetFunction(v *PaChatCompletionStreamRequestMessagesToolCallsFunction) *PaChatCompletionStreamRequestMessagesToolCalls {
	s.Function = v
	return s
}

func (s *PaChatCompletionStreamRequestMessagesToolCalls) SetId(v string) *PaChatCompletionStreamRequestMessagesToolCalls {
	s.Id = &v
	return s
}

func (s *PaChatCompletionStreamRequestMessagesToolCalls) SetType(v string) *PaChatCompletionStreamRequestMessagesToolCalls {
	s.Type = &v
	return s
}

func (s *PaChatCompletionStreamRequestMessagesToolCalls) Validate() error {
	if s.Function != nil {
		if err := s.Function.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PaChatCompletionStreamRequestMessagesToolCallsFunction struct {
	// example:
	//
	// {"city":"杭州"}
	Arguments *string `json:"arguments,omitempty" xml:"arguments,omitempty"`
	// example:
	//
	// get_weather
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s PaChatCompletionStreamRequestMessagesToolCallsFunction) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequestMessagesToolCallsFunction) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequestMessagesToolCallsFunction) GetArguments() *string {
	return s.Arguments
}

func (s *PaChatCompletionStreamRequestMessagesToolCallsFunction) GetName() *string {
	return s.Name
}

func (s *PaChatCompletionStreamRequestMessagesToolCallsFunction) SetArguments(v string) *PaChatCompletionStreamRequestMessagesToolCallsFunction {
	s.Arguments = &v
	return s
}

func (s *PaChatCompletionStreamRequestMessagesToolCallsFunction) SetName(v string) *PaChatCompletionStreamRequestMessagesToolCallsFunction {
	s.Name = &v
	return s
}

func (s *PaChatCompletionStreamRequestMessagesToolCallsFunction) Validate() error {
	return dara.Validate(s)
}

type PaChatCompletionStreamRequestMmProcessorKwargs struct {
	// example:
	//
	// 8
	MaxDynamicPatch *int64 `json:"maxDynamicPatch,omitempty" xml:"maxDynamicPatch,omitempty"`
}

func (s PaChatCompletionStreamRequestMmProcessorKwargs) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequestMmProcessorKwargs) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequestMmProcessorKwargs) GetMaxDynamicPatch() *int64 {
	return s.MaxDynamicPatch
}

func (s *PaChatCompletionStreamRequestMmProcessorKwargs) SetMaxDynamicPatch(v int64) *PaChatCompletionStreamRequestMmProcessorKwargs {
	s.MaxDynamicPatch = &v
	return s
}

func (s *PaChatCompletionStreamRequestMmProcessorKwargs) Validate() error {
	return dara.Validate(s)
}

type PaChatCompletionStreamRequestResponseFormat struct {
	// example:
	//
	// json_object
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PaChatCompletionStreamRequestResponseFormat) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequestResponseFormat) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequestResponseFormat) GetType() *string {
	return s.Type
}

func (s *PaChatCompletionStreamRequestResponseFormat) SetType(v string) *PaChatCompletionStreamRequestResponseFormat {
	s.Type = &v
	return s
}

func (s *PaChatCompletionStreamRequestResponseFormat) Validate() error {
	return dara.Validate(s)
}

type PaChatCompletionStreamRequestStreamOptions struct {
	// example:
	//
	// true
	IncludeUsage *bool `json:"includeUsage,omitempty" xml:"includeUsage,omitempty"`
}

func (s PaChatCompletionStreamRequestStreamOptions) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequestStreamOptions) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequestStreamOptions) GetIncludeUsage() *bool {
	return s.IncludeUsage
}

func (s *PaChatCompletionStreamRequestStreamOptions) SetIncludeUsage(v bool) *PaChatCompletionStreamRequestStreamOptions {
	s.IncludeUsage = &v
	return s
}

func (s *PaChatCompletionStreamRequestStreamOptions) Validate() error {
	return dara.Validate(s)
}

type PaChatCompletionStreamRequestStructuredOutputs struct {
	// example:
	//
	// ["yes","no"]
	Choice []*string `json:"choice,omitempty" xml:"choice,omitempty" type:"Repeated"`
}

func (s PaChatCompletionStreamRequestStructuredOutputs) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequestStructuredOutputs) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequestStructuredOutputs) GetChoice() []*string {
	return s.Choice
}

func (s *PaChatCompletionStreamRequestStructuredOutputs) SetChoice(v []*string) *PaChatCompletionStreamRequestStructuredOutputs {
	s.Choice = v
	return s
}

func (s *PaChatCompletionStreamRequestStructuredOutputs) Validate() error {
	return dara.Validate(s)
}

type PaChatCompletionStreamRequestTools struct {
	Function *PaChatCompletionStreamRequestToolsFunction `json:"function,omitempty" xml:"function,omitempty" type:"Struct"`
	// example:
	//
	// function
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PaChatCompletionStreamRequestTools) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequestTools) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequestTools) GetFunction() *PaChatCompletionStreamRequestToolsFunction {
	return s.Function
}

func (s *PaChatCompletionStreamRequestTools) GetType() *string {
	return s.Type
}

func (s *PaChatCompletionStreamRequestTools) SetFunction(v *PaChatCompletionStreamRequestToolsFunction) *PaChatCompletionStreamRequestTools {
	s.Function = v
	return s
}

func (s *PaChatCompletionStreamRequestTools) SetType(v string) *PaChatCompletionStreamRequestTools {
	s.Type = &v
	return s
}

func (s *PaChatCompletionStreamRequestTools) Validate() error {
	if s.Function != nil {
		if err := s.Function.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PaChatCompletionStreamRequestToolsFunction struct {
	// example:
	//
	// 查询指定城市天气
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// get_weather
	Name       *string                                               `json:"name,omitempty" xml:"name,omitempty"`
	Parameters *PaChatCompletionStreamRequestToolsFunctionParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	// example:
	//
	// true
	Strict *bool `json:"strict,omitempty" xml:"strict,omitempty"`
}

func (s PaChatCompletionStreamRequestToolsFunction) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequestToolsFunction) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequestToolsFunction) GetDescription() *string {
	return s.Description
}

func (s *PaChatCompletionStreamRequestToolsFunction) GetName() *string {
	return s.Name
}

func (s *PaChatCompletionStreamRequestToolsFunction) GetParameters() *PaChatCompletionStreamRequestToolsFunctionParameters {
	return s.Parameters
}

func (s *PaChatCompletionStreamRequestToolsFunction) GetStrict() *bool {
	return s.Strict
}

func (s *PaChatCompletionStreamRequestToolsFunction) SetDescription(v string) *PaChatCompletionStreamRequestToolsFunction {
	s.Description = &v
	return s
}

func (s *PaChatCompletionStreamRequestToolsFunction) SetName(v string) *PaChatCompletionStreamRequestToolsFunction {
	s.Name = &v
	return s
}

func (s *PaChatCompletionStreamRequestToolsFunction) SetParameters(v *PaChatCompletionStreamRequestToolsFunctionParameters) *PaChatCompletionStreamRequestToolsFunction {
	s.Parameters = v
	return s
}

func (s *PaChatCompletionStreamRequestToolsFunction) SetStrict(v bool) *PaChatCompletionStreamRequestToolsFunction {
	s.Strict = &v
	return s
}

func (s *PaChatCompletionStreamRequestToolsFunction) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PaChatCompletionStreamRequestToolsFunctionParameters struct {
	Properties *PaChatCompletionStreamRequestToolsFunctionParametersProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	Required   []*string                                                       `json:"required,omitempty" xml:"required,omitempty" type:"Repeated"`
	// example:
	//
	// object
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PaChatCompletionStreamRequestToolsFunctionParameters) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequestToolsFunctionParameters) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequestToolsFunctionParameters) GetProperties() *PaChatCompletionStreamRequestToolsFunctionParametersProperties {
	return s.Properties
}

func (s *PaChatCompletionStreamRequestToolsFunctionParameters) GetRequired() []*string {
	return s.Required
}

func (s *PaChatCompletionStreamRequestToolsFunctionParameters) GetType() *string {
	return s.Type
}

func (s *PaChatCompletionStreamRequestToolsFunctionParameters) SetProperties(v *PaChatCompletionStreamRequestToolsFunctionParametersProperties) *PaChatCompletionStreamRequestToolsFunctionParameters {
	s.Properties = v
	return s
}

func (s *PaChatCompletionStreamRequestToolsFunctionParameters) SetRequired(v []*string) *PaChatCompletionStreamRequestToolsFunctionParameters {
	s.Required = v
	return s
}

func (s *PaChatCompletionStreamRequestToolsFunctionParameters) SetType(v string) *PaChatCompletionStreamRequestToolsFunctionParameters {
	s.Type = &v
	return s
}

func (s *PaChatCompletionStreamRequestToolsFunctionParameters) Validate() error {
	if s.Properties != nil {
		if err := s.Properties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PaChatCompletionStreamRequestToolsFunctionParametersProperties struct {
	City *PaChatCompletionStreamRequestToolsFunctionParametersPropertiesCity `json:"city,omitempty" xml:"city,omitempty" type:"Struct"`
}

func (s PaChatCompletionStreamRequestToolsFunctionParametersProperties) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequestToolsFunctionParametersProperties) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequestToolsFunctionParametersProperties) GetCity() *PaChatCompletionStreamRequestToolsFunctionParametersPropertiesCity {
	return s.City
}

func (s *PaChatCompletionStreamRequestToolsFunctionParametersProperties) SetCity(v *PaChatCompletionStreamRequestToolsFunctionParametersPropertiesCity) *PaChatCompletionStreamRequestToolsFunctionParametersProperties {
	s.City = v
	return s
}

func (s *PaChatCompletionStreamRequestToolsFunctionParametersProperties) Validate() error {
	if s.City != nil {
		if err := s.City.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PaChatCompletionStreamRequestToolsFunctionParametersPropertiesCity struct {
	// example:
	//
	// 要查询天气的城市名称
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PaChatCompletionStreamRequestToolsFunctionParametersPropertiesCity) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamRequestToolsFunctionParametersPropertiesCity) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamRequestToolsFunctionParametersPropertiesCity) GetDescription() *string {
	return s.Description
}

func (s *PaChatCompletionStreamRequestToolsFunctionParametersPropertiesCity) GetType() *string {
	return s.Type
}

func (s *PaChatCompletionStreamRequestToolsFunctionParametersPropertiesCity) SetDescription(v string) *PaChatCompletionStreamRequestToolsFunctionParametersPropertiesCity {
	s.Description = &v
	return s
}

func (s *PaChatCompletionStreamRequestToolsFunctionParametersPropertiesCity) SetType(v string) *PaChatCompletionStreamRequestToolsFunctionParametersPropertiesCity {
	s.Type = &v
	return s
}

func (s *PaChatCompletionStreamRequestToolsFunctionParametersPropertiesCity) Validate() error {
	return dara.Validate(s)
}
