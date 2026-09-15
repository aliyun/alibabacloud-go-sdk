// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGuiChatCompletionStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedTokenIds(v []*int64) *GuiChatCompletionStreamRequest
	GetAllowedTokenIds() []*int64
	SetBadWords(v []*string) *GuiChatCompletionStreamRequest
	GetBadWords() []*string
	SetChatTemplateKwargs(v *GuiChatCompletionStreamRequestChatTemplateKwargs) *GuiChatCompletionStreamRequest
	GetChatTemplateKwargs() *GuiChatCompletionStreamRequestChatTemplateKwargs
	SetFrequencyPenalty(v float64) *GuiChatCompletionStreamRequest
	GetFrequencyPenalty() *float64
	SetIgnoreEos(v bool) *GuiChatCompletionStreamRequest
	GetIgnoreEos() *bool
	SetIncludeReasoning(v bool) *GuiChatCompletionStreamRequest
	GetIncludeReasoning() *bool
	SetLogprobs(v bool) *GuiChatCompletionStreamRequest
	GetLogprobs() *bool
	SetMaxCompletionTokens(v int64) *GuiChatCompletionStreamRequest
	GetMaxCompletionTokens() *int64
	SetMaxTokens(v int64) *GuiChatCompletionStreamRequest
	GetMaxTokens() *int64
	SetMessages(v []*GuiChatCompletionStreamRequestMessages) *GuiChatCompletionStreamRequest
	GetMessages() []*GuiChatCompletionStreamRequestMessages
	SetMetadata(v *GuiChatCompletionStreamRequestMetadata) *GuiChatCompletionStreamRequest
	GetMetadata() *GuiChatCompletionStreamRequestMetadata
	SetMinP(v float64) *GuiChatCompletionStreamRequest
	GetMinP() *float64
	SetMinTokens(v int64) *GuiChatCompletionStreamRequest
	GetMinTokens() *int64
	SetMmProcessorKwargs(v *GuiChatCompletionStreamRequestMmProcessorKwargs) *GuiChatCompletionStreamRequest
	GetMmProcessorKwargs() *GuiChatCompletionStreamRequestMmProcessorKwargs
	SetModel(v string) *GuiChatCompletionStreamRequest
	GetModel() *string
	SetN(v int64) *GuiChatCompletionStreamRequest
	GetN() *int64
	SetParallelToolCalls(v bool) *GuiChatCompletionStreamRequest
	GetParallelToolCalls() *bool
	SetPresencePenalty(v float64) *GuiChatCompletionStreamRequest
	GetPresencePenalty() *float64
	SetPromptLogprobs(v int64) *GuiChatCompletionStreamRequest
	GetPromptLogprobs() *int64
	SetReasoningEffort(v string) *GuiChatCompletionStreamRequest
	GetReasoningEffort() *string
	SetRepetitionPenalty(v float64) *GuiChatCompletionStreamRequest
	GetRepetitionPenalty() *float64
	SetResponseFormat(v *GuiChatCompletionStreamRequestResponseFormat) *GuiChatCompletionStreamRequest
	GetResponseFormat() *GuiChatCompletionStreamRequestResponseFormat
	SetSeed(v int64) *GuiChatCompletionStreamRequest
	GetSeed() *int64
	SetSkipSpecialTokens(v bool) *GuiChatCompletionStreamRequest
	GetSkipSpecialTokens() *bool
	SetStop(v []*string) *GuiChatCompletionStreamRequest
	GetStop() []*string
	SetStopTokenIds(v []*int64) *GuiChatCompletionStreamRequest
	GetStopTokenIds() []*int64
	SetStream(v bool) *GuiChatCompletionStreamRequest
	GetStream() *bool
	SetStreamOptions(v *GuiChatCompletionStreamRequestStreamOptions) *GuiChatCompletionStreamRequest
	GetStreamOptions() *GuiChatCompletionStreamRequestStreamOptions
	SetStructuredOutputs(v *GuiChatCompletionStreamRequestStructuredOutputs) *GuiChatCompletionStreamRequest
	GetStructuredOutputs() *GuiChatCompletionStreamRequestStructuredOutputs
	SetTemperature(v float64) *GuiChatCompletionStreamRequest
	GetTemperature() *float64
	SetTopK(v int64) *GuiChatCompletionStreamRequest
	GetTopK() *int64
	SetTopLogprobs(v int64) *GuiChatCompletionStreamRequest
	GetTopLogprobs() *int64
	SetTopP(v float64) *GuiChatCompletionStreamRequest
	GetTopP() *float64
}

type GuiChatCompletionStreamRequest struct {
	// example:
	//
	// [10,11]
	AllowedTokenIds []*int64 `json:"allowedTokenIds,omitempty" xml:"allowedTokenIds,omitempty" type:"Repeated"`
	// example:
	//
	// ["blocked"]
	BadWords           []*string                                         `json:"badWords,omitempty" xml:"badWords,omitempty" type:"Repeated"`
	ChatTemplateKwargs *GuiChatCompletionStreamRequestChatTemplateKwargs `json:"chatTemplateKwargs,omitempty" xml:"chatTemplateKwargs,omitempty" type:"Struct"`
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
	Messages []*GuiChatCompletionStreamRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	Metadata *GuiChatCompletionStreamRequestMetadata   `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// example:
	//
	// 0.05
	MinP *float64 `json:"minP,omitempty" xml:"minP,omitempty"`
	// example:
	//
	// 1
	MinTokens         *int64                                           `json:"minTokens,omitempty" xml:"minTokens,omitempty"`
	MmProcessorKwargs *GuiChatCompletionStreamRequestMmProcessorKwargs `json:"mmProcessorKwargs,omitempty" xml:"mmProcessorKwargs,omitempty" type:"Struct"`
	Model             *string                                          `json:"model,omitempty" xml:"model,omitempty"`
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
	// medium
	ReasoningEffort *string `json:"reasoningEffort,omitempty" xml:"reasoningEffort,omitempty"`
	// example:
	//
	// 1.1
	RepetitionPenalty *float64                                      `json:"repetitionPenalty,omitempty" xml:"repetitionPenalty,omitempty"`
	ResponseFormat    *GuiChatCompletionStreamRequestResponseFormat `json:"responseFormat,omitempty" xml:"responseFormat,omitempty" type:"Struct"`
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
	// ["DONE"]
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
	Stream            *bool                                            `json:"stream,omitempty" xml:"stream,omitempty"`
	StreamOptions     *GuiChatCompletionStreamRequestStreamOptions     `json:"streamOptions,omitempty" xml:"streamOptions,omitempty" type:"Struct"`
	StructuredOutputs *GuiChatCompletionStreamRequestStructuredOutputs `json:"structuredOutputs,omitempty" xml:"structuredOutputs,omitempty" type:"Struct"`
	// example:
	//
	// 0.2
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
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

func (s GuiChatCompletionStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s GuiChatCompletionStreamRequest) GoString() string {
	return s.String()
}

func (s *GuiChatCompletionStreamRequest) GetAllowedTokenIds() []*int64 {
	return s.AllowedTokenIds
}

func (s *GuiChatCompletionStreamRequest) GetBadWords() []*string {
	return s.BadWords
}

func (s *GuiChatCompletionStreamRequest) GetChatTemplateKwargs() *GuiChatCompletionStreamRequestChatTemplateKwargs {
	return s.ChatTemplateKwargs
}

func (s *GuiChatCompletionStreamRequest) GetFrequencyPenalty() *float64 {
	return s.FrequencyPenalty
}

func (s *GuiChatCompletionStreamRequest) GetIgnoreEos() *bool {
	return s.IgnoreEos
}

func (s *GuiChatCompletionStreamRequest) GetIncludeReasoning() *bool {
	return s.IncludeReasoning
}

func (s *GuiChatCompletionStreamRequest) GetLogprobs() *bool {
	return s.Logprobs
}

func (s *GuiChatCompletionStreamRequest) GetMaxCompletionTokens() *int64 {
	return s.MaxCompletionTokens
}

func (s *GuiChatCompletionStreamRequest) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *GuiChatCompletionStreamRequest) GetMessages() []*GuiChatCompletionStreamRequestMessages {
	return s.Messages
}

func (s *GuiChatCompletionStreamRequest) GetMetadata() *GuiChatCompletionStreamRequestMetadata {
	return s.Metadata
}

func (s *GuiChatCompletionStreamRequest) GetMinP() *float64 {
	return s.MinP
}

func (s *GuiChatCompletionStreamRequest) GetMinTokens() *int64 {
	return s.MinTokens
}

func (s *GuiChatCompletionStreamRequest) GetMmProcessorKwargs() *GuiChatCompletionStreamRequestMmProcessorKwargs {
	return s.MmProcessorKwargs
}

func (s *GuiChatCompletionStreamRequest) GetModel() *string {
	return s.Model
}

func (s *GuiChatCompletionStreamRequest) GetN() *int64 {
	return s.N
}

func (s *GuiChatCompletionStreamRequest) GetParallelToolCalls() *bool {
	return s.ParallelToolCalls
}

func (s *GuiChatCompletionStreamRequest) GetPresencePenalty() *float64 {
	return s.PresencePenalty
}

func (s *GuiChatCompletionStreamRequest) GetPromptLogprobs() *int64 {
	return s.PromptLogprobs
}

func (s *GuiChatCompletionStreamRequest) GetReasoningEffort() *string {
	return s.ReasoningEffort
}

func (s *GuiChatCompletionStreamRequest) GetRepetitionPenalty() *float64 {
	return s.RepetitionPenalty
}

func (s *GuiChatCompletionStreamRequest) GetResponseFormat() *GuiChatCompletionStreamRequestResponseFormat {
	return s.ResponseFormat
}

func (s *GuiChatCompletionStreamRequest) GetSeed() *int64 {
	return s.Seed
}

func (s *GuiChatCompletionStreamRequest) GetSkipSpecialTokens() *bool {
	return s.SkipSpecialTokens
}

func (s *GuiChatCompletionStreamRequest) GetStop() []*string {
	return s.Stop
}

func (s *GuiChatCompletionStreamRequest) GetStopTokenIds() []*int64 {
	return s.StopTokenIds
}

func (s *GuiChatCompletionStreamRequest) GetStream() *bool {
	return s.Stream
}

func (s *GuiChatCompletionStreamRequest) GetStreamOptions() *GuiChatCompletionStreamRequestStreamOptions {
	return s.StreamOptions
}

func (s *GuiChatCompletionStreamRequest) GetStructuredOutputs() *GuiChatCompletionStreamRequestStructuredOutputs {
	return s.StructuredOutputs
}

func (s *GuiChatCompletionStreamRequest) GetTemperature() *float64 {
	return s.Temperature
}

func (s *GuiChatCompletionStreamRequest) GetTopK() *int64 {
	return s.TopK
}

func (s *GuiChatCompletionStreamRequest) GetTopLogprobs() *int64 {
	return s.TopLogprobs
}

func (s *GuiChatCompletionStreamRequest) GetTopP() *float64 {
	return s.TopP
}

func (s *GuiChatCompletionStreamRequest) SetAllowedTokenIds(v []*int64) *GuiChatCompletionStreamRequest {
	s.AllowedTokenIds = v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetBadWords(v []*string) *GuiChatCompletionStreamRequest {
	s.BadWords = v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetChatTemplateKwargs(v *GuiChatCompletionStreamRequestChatTemplateKwargs) *GuiChatCompletionStreamRequest {
	s.ChatTemplateKwargs = v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetFrequencyPenalty(v float64) *GuiChatCompletionStreamRequest {
	s.FrequencyPenalty = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetIgnoreEos(v bool) *GuiChatCompletionStreamRequest {
	s.IgnoreEos = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetIncludeReasoning(v bool) *GuiChatCompletionStreamRequest {
	s.IncludeReasoning = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetLogprobs(v bool) *GuiChatCompletionStreamRequest {
	s.Logprobs = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetMaxCompletionTokens(v int64) *GuiChatCompletionStreamRequest {
	s.MaxCompletionTokens = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetMaxTokens(v int64) *GuiChatCompletionStreamRequest {
	s.MaxTokens = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetMessages(v []*GuiChatCompletionStreamRequestMessages) *GuiChatCompletionStreamRequest {
	s.Messages = v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetMetadata(v *GuiChatCompletionStreamRequestMetadata) *GuiChatCompletionStreamRequest {
	s.Metadata = v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetMinP(v float64) *GuiChatCompletionStreamRequest {
	s.MinP = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetMinTokens(v int64) *GuiChatCompletionStreamRequest {
	s.MinTokens = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetMmProcessorKwargs(v *GuiChatCompletionStreamRequestMmProcessorKwargs) *GuiChatCompletionStreamRequest {
	s.MmProcessorKwargs = v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetModel(v string) *GuiChatCompletionStreamRequest {
	s.Model = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetN(v int64) *GuiChatCompletionStreamRequest {
	s.N = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetParallelToolCalls(v bool) *GuiChatCompletionStreamRequest {
	s.ParallelToolCalls = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetPresencePenalty(v float64) *GuiChatCompletionStreamRequest {
	s.PresencePenalty = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetPromptLogprobs(v int64) *GuiChatCompletionStreamRequest {
	s.PromptLogprobs = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetReasoningEffort(v string) *GuiChatCompletionStreamRequest {
	s.ReasoningEffort = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetRepetitionPenalty(v float64) *GuiChatCompletionStreamRequest {
	s.RepetitionPenalty = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetResponseFormat(v *GuiChatCompletionStreamRequestResponseFormat) *GuiChatCompletionStreamRequest {
	s.ResponseFormat = v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetSeed(v int64) *GuiChatCompletionStreamRequest {
	s.Seed = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetSkipSpecialTokens(v bool) *GuiChatCompletionStreamRequest {
	s.SkipSpecialTokens = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetStop(v []*string) *GuiChatCompletionStreamRequest {
	s.Stop = v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetStopTokenIds(v []*int64) *GuiChatCompletionStreamRequest {
	s.StopTokenIds = v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetStream(v bool) *GuiChatCompletionStreamRequest {
	s.Stream = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetStreamOptions(v *GuiChatCompletionStreamRequestStreamOptions) *GuiChatCompletionStreamRequest {
	s.StreamOptions = v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetStructuredOutputs(v *GuiChatCompletionStreamRequestStructuredOutputs) *GuiChatCompletionStreamRequest {
	s.StructuredOutputs = v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetTemperature(v float64) *GuiChatCompletionStreamRequest {
	s.Temperature = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetTopK(v int64) *GuiChatCompletionStreamRequest {
	s.TopK = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetTopLogprobs(v int64) *GuiChatCompletionStreamRequest {
	s.TopLogprobs = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) SetTopP(v float64) *GuiChatCompletionStreamRequest {
	s.TopP = &v
	return s
}

func (s *GuiChatCompletionStreamRequest) Validate() error {
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
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
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
	return nil
}

type GuiChatCompletionStreamRequestChatTemplateKwargs struct {
	// example:
	//
	// true
	EnableThinking *bool `json:"enableThinking,omitempty" xml:"enableThinking,omitempty"`
	// example:
	//
	// true
	PreserveThinking *bool `json:"preserveThinking,omitempty" xml:"preserveThinking,omitempty"`
}

func (s GuiChatCompletionStreamRequestChatTemplateKwargs) String() string {
	return dara.Prettify(s)
}

func (s GuiChatCompletionStreamRequestChatTemplateKwargs) GoString() string {
	return s.String()
}

func (s *GuiChatCompletionStreamRequestChatTemplateKwargs) GetEnableThinking() *bool {
	return s.EnableThinking
}

func (s *GuiChatCompletionStreamRequestChatTemplateKwargs) GetPreserveThinking() *bool {
	return s.PreserveThinking
}

func (s *GuiChatCompletionStreamRequestChatTemplateKwargs) SetEnableThinking(v bool) *GuiChatCompletionStreamRequestChatTemplateKwargs {
	s.EnableThinking = &v
	return s
}

func (s *GuiChatCompletionStreamRequestChatTemplateKwargs) SetPreserveThinking(v bool) *GuiChatCompletionStreamRequestChatTemplateKwargs {
	s.PreserveThinking = &v
	return s
}

func (s *GuiChatCompletionStreamRequestChatTemplateKwargs) Validate() error {
	return dara.Validate(s)
}

type GuiChatCompletionStreamRequestMessages struct {
	Content []*GuiChatCompletionStreamRequestMessagesContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// call_gui_1
	ToolCallId *string `json:"toolCallId,omitempty" xml:"toolCallId,omitempty"`
}

func (s GuiChatCompletionStreamRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s GuiChatCompletionStreamRequestMessages) GoString() string {
	return s.String()
}

func (s *GuiChatCompletionStreamRequestMessages) GetContent() []*GuiChatCompletionStreamRequestMessagesContent {
	return s.Content
}

func (s *GuiChatCompletionStreamRequestMessages) GetRole() *string {
	return s.Role
}

func (s *GuiChatCompletionStreamRequestMessages) GetToolCallId() *string {
	return s.ToolCallId
}

func (s *GuiChatCompletionStreamRequestMessages) SetContent(v []*GuiChatCompletionStreamRequestMessagesContent) *GuiChatCompletionStreamRequestMessages {
	s.Content = v
	return s
}

func (s *GuiChatCompletionStreamRequestMessages) SetRole(v string) *GuiChatCompletionStreamRequestMessages {
	s.Role = &v
	return s
}

func (s *GuiChatCompletionStreamRequestMessages) SetToolCallId(v string) *GuiChatCompletionStreamRequestMessages {
	s.ToolCallId = &v
	return s
}

func (s *GuiChatCompletionStreamRequestMessages) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GuiChatCompletionStreamRequestMessagesContent struct {
	// example:
	//
	// data:image/png;base64,...
	ImageData *string                                                `json:"imageData,omitempty" xml:"imageData,omitempty"`
	ImageUrl  *GuiChatCompletionStreamRequestMessagesContentImageUrl `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Struct"`
	// example:
	//
	// 点击搜索按钮
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GuiChatCompletionStreamRequestMessagesContent) String() string {
	return dara.Prettify(s)
}

func (s GuiChatCompletionStreamRequestMessagesContent) GoString() string {
	return s.String()
}

func (s *GuiChatCompletionStreamRequestMessagesContent) GetImageData() *string {
	return s.ImageData
}

func (s *GuiChatCompletionStreamRequestMessagesContent) GetImageUrl() *GuiChatCompletionStreamRequestMessagesContentImageUrl {
	return s.ImageUrl
}

func (s *GuiChatCompletionStreamRequestMessagesContent) GetText() *string {
	return s.Text
}

func (s *GuiChatCompletionStreamRequestMessagesContent) GetType() *string {
	return s.Type
}

func (s *GuiChatCompletionStreamRequestMessagesContent) SetImageData(v string) *GuiChatCompletionStreamRequestMessagesContent {
	s.ImageData = &v
	return s
}

func (s *GuiChatCompletionStreamRequestMessagesContent) SetImageUrl(v *GuiChatCompletionStreamRequestMessagesContentImageUrl) *GuiChatCompletionStreamRequestMessagesContent {
	s.ImageUrl = v
	return s
}

func (s *GuiChatCompletionStreamRequestMessagesContent) SetText(v string) *GuiChatCompletionStreamRequestMessagesContent {
	s.Text = &v
	return s
}

func (s *GuiChatCompletionStreamRequestMessagesContent) SetType(v string) *GuiChatCompletionStreamRequestMessagesContent {
	s.Type = &v
	return s
}

func (s *GuiChatCompletionStreamRequestMessagesContent) Validate() error {
	if s.ImageUrl != nil {
		if err := s.ImageUrl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GuiChatCompletionStreamRequestMessagesContentImageUrl struct {
	// example:
	//
	// https://example.com/screenshot.png
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GuiChatCompletionStreamRequestMessagesContentImageUrl) String() string {
	return dara.Prettify(s)
}

func (s GuiChatCompletionStreamRequestMessagesContentImageUrl) GoString() string {
	return s.String()
}

func (s *GuiChatCompletionStreamRequestMessagesContentImageUrl) GetUrl() *string {
	return s.Url
}

func (s *GuiChatCompletionStreamRequestMessagesContentImageUrl) SetUrl(v string) *GuiChatCompletionStreamRequestMessagesContentImageUrl {
	s.Url = &v
	return s
}

func (s *GuiChatCompletionStreamRequestMessagesContentImageUrl) Validate() error {
	return dara.Validate(s)
}

type GuiChatCompletionStreamRequestMetadata struct {
	// example:
	//
	// ["Chrome","Settings"]
	AppList []*string `json:"appList,omitempty" xml:"appList,omitempty" type:"Repeated"`
	// example:
	//
	// ["Chrome","Settings"]
	AvailableApps []*string `json:"availableApps,omitempty" xml:"availableApps,omitempty" type:"Repeated"`
	// example:
	//
	// previous action completed
	HarnessMessage *string `json:"harnessMessage,omitempty" xml:"harnessMessage,omitempty"`
	// example:
	//
	// 1080
	ScreenHeight *int64 `json:"screenHeight,omitempty" xml:"screenHeight,omitempty"`
	// example:
	//
	// 1920
	ScreenWidth *int64 `json:"screenWidth,omitempty" xml:"screenWidth,omitempty"`
}

func (s GuiChatCompletionStreamRequestMetadata) String() string {
	return dara.Prettify(s)
}

func (s GuiChatCompletionStreamRequestMetadata) GoString() string {
	return s.String()
}

func (s *GuiChatCompletionStreamRequestMetadata) GetAppList() []*string {
	return s.AppList
}

func (s *GuiChatCompletionStreamRequestMetadata) GetAvailableApps() []*string {
	return s.AvailableApps
}

func (s *GuiChatCompletionStreamRequestMetadata) GetHarnessMessage() *string {
	return s.HarnessMessage
}

func (s *GuiChatCompletionStreamRequestMetadata) GetScreenHeight() *int64 {
	return s.ScreenHeight
}

func (s *GuiChatCompletionStreamRequestMetadata) GetScreenWidth() *int64 {
	return s.ScreenWidth
}

func (s *GuiChatCompletionStreamRequestMetadata) SetAppList(v []*string) *GuiChatCompletionStreamRequestMetadata {
	s.AppList = v
	return s
}

func (s *GuiChatCompletionStreamRequestMetadata) SetAvailableApps(v []*string) *GuiChatCompletionStreamRequestMetadata {
	s.AvailableApps = v
	return s
}

func (s *GuiChatCompletionStreamRequestMetadata) SetHarnessMessage(v string) *GuiChatCompletionStreamRequestMetadata {
	s.HarnessMessage = &v
	return s
}

func (s *GuiChatCompletionStreamRequestMetadata) SetScreenHeight(v int64) *GuiChatCompletionStreamRequestMetadata {
	s.ScreenHeight = &v
	return s
}

func (s *GuiChatCompletionStreamRequestMetadata) SetScreenWidth(v int64) *GuiChatCompletionStreamRequestMetadata {
	s.ScreenWidth = &v
	return s
}

func (s *GuiChatCompletionStreamRequestMetadata) Validate() error {
	return dara.Validate(s)
}

type GuiChatCompletionStreamRequestMmProcessorKwargs struct {
	// example:
	//
	// 8
	MaxDynamicPatch *int64 `json:"maxDynamicPatch,omitempty" xml:"maxDynamicPatch,omitempty"`
}

func (s GuiChatCompletionStreamRequestMmProcessorKwargs) String() string {
	return dara.Prettify(s)
}

func (s GuiChatCompletionStreamRequestMmProcessorKwargs) GoString() string {
	return s.String()
}

func (s *GuiChatCompletionStreamRequestMmProcessorKwargs) GetMaxDynamicPatch() *int64 {
	return s.MaxDynamicPatch
}

func (s *GuiChatCompletionStreamRequestMmProcessorKwargs) SetMaxDynamicPatch(v int64) *GuiChatCompletionStreamRequestMmProcessorKwargs {
	s.MaxDynamicPatch = &v
	return s
}

func (s *GuiChatCompletionStreamRequestMmProcessorKwargs) Validate() error {
	return dara.Validate(s)
}

type GuiChatCompletionStreamRequestResponseFormat struct {
	// example:
	//
	// json_object
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GuiChatCompletionStreamRequestResponseFormat) String() string {
	return dara.Prettify(s)
}

func (s GuiChatCompletionStreamRequestResponseFormat) GoString() string {
	return s.String()
}

func (s *GuiChatCompletionStreamRequestResponseFormat) GetType() *string {
	return s.Type
}

func (s *GuiChatCompletionStreamRequestResponseFormat) SetType(v string) *GuiChatCompletionStreamRequestResponseFormat {
	s.Type = &v
	return s
}

func (s *GuiChatCompletionStreamRequestResponseFormat) Validate() error {
	return dara.Validate(s)
}

type GuiChatCompletionStreamRequestStreamOptions struct {
	// example:
	//
	// true
	IncludeUsage *bool `json:"includeUsage,omitempty" xml:"includeUsage,omitempty"`
}

func (s GuiChatCompletionStreamRequestStreamOptions) String() string {
	return dara.Prettify(s)
}

func (s GuiChatCompletionStreamRequestStreamOptions) GoString() string {
	return s.String()
}

func (s *GuiChatCompletionStreamRequestStreamOptions) GetIncludeUsage() *bool {
	return s.IncludeUsage
}

func (s *GuiChatCompletionStreamRequestStreamOptions) SetIncludeUsage(v bool) *GuiChatCompletionStreamRequestStreamOptions {
	s.IncludeUsage = &v
	return s
}

func (s *GuiChatCompletionStreamRequestStreamOptions) Validate() error {
	return dara.Validate(s)
}

type GuiChatCompletionStreamRequestStructuredOutputs struct {
	// example:
	//
	// ["tap","type"]
	Choice []*string `json:"choice,omitempty" xml:"choice,omitempty" type:"Repeated"`
}

func (s GuiChatCompletionStreamRequestStructuredOutputs) String() string {
	return dara.Prettify(s)
}

func (s GuiChatCompletionStreamRequestStructuredOutputs) GoString() string {
	return s.String()
}

func (s *GuiChatCompletionStreamRequestStructuredOutputs) GetChoice() []*string {
	return s.Choice
}

func (s *GuiChatCompletionStreamRequestStructuredOutputs) SetChoice(v []*string) *GuiChatCompletionStreamRequestStructuredOutputs {
	s.Choice = v
	return s
}

func (s *GuiChatCompletionStreamRequestStructuredOutputs) Validate() error {
	return dara.Validate(s)
}
