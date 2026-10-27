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
	// The whitelist of token IDs that are allowed to be generated.
	//
	// example:
	//
	// [10,11]
	AllowedTokenIds []*int64 `json:"allowedTokenIds,omitempty" xml:"allowedTokenIds,omitempty" type:"Repeated"`
	// The list of words or strings that are banned from being generated.
	//
	// example:
	//
	// ["blocked"]
	BadWords []*string `json:"badWords,omitempty" xml:"badWords,omitempty" type:"Repeated"`
	// The vLLM or OpenAI-compatible template parameter object. The request values override the default values with the same names in the model and are passed as-is to the primary model of the Mobile-Use Agent.
	ChatTemplateKwargs *GuiChatCompletionStreamRequestChatTemplateKwargs `json:"chatTemplateKwargs,omitempty" xml:"chatTemplateKwargs,omitempty" type:"Struct"`
	// The frequency penalty parameter, which is passed through to the primary model of the Mobile-Use Agent.
	//
	// example:
	//
	// 0.0
	FrequencyPenalty *float64 `json:"frequencyPenalty,omitempty" xml:"frequencyPenalty,omitempty"`
	// Specifies whether to ignore the end-of-sequence (EOS) token and continue generation.
	//
	// example:
	//
	// false
	IgnoreEos *bool `json:"ignoreEos,omitempty" xml:"ignoreEos,omitempty"`
	// Specifies whether to request the upstream to generate the reasoning field. The client-facing results of the Mobile-Use Agent do not directly expose internal thinking.
	//
	// example:
	//
	// true
	IncludeReasoning *bool `json:"includeReasoning,omitempty" xml:"includeReasoning,omitempty"`
	// The model compatibility parameter. The JSON field is logprobs. Whether this parameter is accepted and takes effect depends on the model used by the instance. The current Mobile-Use Agent response does not provide token log probability details.
	//
	// example:
	//
	// false
	Logprobs *bool `json:"logprobs,omitempty" xml:"logprobs,omitempty"`
	// The OpenAI-compatible maximum number of completion tokens. If provided, the value must be greater than 0.
	//
	// example:
	//
	// 2048
	MaxCompletionTokens *int64 `json:"maxCompletionTokens,omitempty" xml:"maxCompletionTokens,omitempty"`
	// The OpenAI-compatible maximum number of generated tokens. If provided, the value must be greater than 0.
	//
	// example:
	//
	// 2048
	MaxTokens *int64 `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	// The messages for the current Mobile-Use Agent step. The first step can only contain one user message. Subsequent steps can only contain one tool message followed by one user message. Complete multi-round history is not allowed.
	//
	// This parameter is required.
	Messages []*GuiChatCompletionStreamRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// The metadata for the Mobile-Use Agent step. Supports screen_width, screen_height, available_apps, and the compatibility field harness_message. The tool receipt and harness_message cannot be passed at the same time.
	Metadata *GuiChatCompletionStreamRequestMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// The vLLM Min-P sampling parameter.
	//
	// example:
	//
	// 0.05
	MinP *float64 `json:"minP,omitempty" xml:"minP,omitempty"`
	// The minimum number of tokens to generate.
	//
	// example:
	//
	// 1
	MinTokens *int64 `json:"minTokens,omitempty" xml:"minTokens,omitempty"`
	// The multimodal preprocessor parameters, passed through to the Mobile-Use Agent primary model as-is by key.
	MmProcessorKwargs *GuiChatCompletionStreamRequestMmProcessorKwargs `json:"mmProcessorKwargs,omitempty" xml:"mmProcessorKwargs,omitempty" type:"Struct"`
	// Reserved for OpenAI client compatibility. This parameter is ignored during routing. The model is determined by the published configuration of the instance, and the model field in the response is always an empty string.
	//
	// example:
	//
	// " "
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// Set this parameter to 1. Whether this parameter is accepted depends on the model. The current Mobile-Use Agent response does not return multiple candidate actions.
	//
	// example:
	//
	// 1
	N *int64 `json:"n,omitempty" xml:"n,omitempty"`
	// The model-compatible switch to invoke parallel tool calling. This parameter does not expand the set of Mobile-Use Agent actions active for this operation.
	//
	// example:
	//
	// false
	ParallelToolCalls *bool `json:"parallelToolCalls,omitempty" xml:"parallelToolCalls,omitempty"`
	// The presence penalty parameter, passed through to the Mobile-Use Agent primary model.
	//
	// example:
	//
	// 0.0
	PresencePenalty *float64 `json:"presencePenalty,omitempty" xml:"presencePenalty,omitempty"`
	// A model-compatible parameter. The JSON field is prompt_logprobs. Whether this parameter is accepted depends on the model. The current Mobile-Use Agent response does not return input token probability details.
	//
	// example:
	//
	// 2
	PromptLogprobs *int64 `json:"promptLogprobs,omitempty" xml:"promptLogprobs,omitempty"`
	// The OpenAI-compatible reasoning effort, such as low, medium, or high.
	//
	// example:
	//
	// medium
	ReasoningEffort *string `json:"reasoningEffort,omitempty" xml:"reasoningEffort,omitempty"`
	// The vLLM repetition penalty parameter.
	//
	// example:
	//
	// 1.1
	RepetitionPenalty *float64 `json:"repetitionPenalty,omitempty" xml:"repetitionPenalty,omitempty"`
	// The model-compatible structured output options. This parameter does not change the final Mobile-Use Agent tool calling structure when you invoke this operation. Avoid output constraint Settings that conflict with the action protocol.
	ResponseFormat *GuiChatCompletionStreamRequestResponseFormat `json:"responseFormat,omitempty" xml:"responseFormat,omitempty" type:"Struct"`
	// The random seed.
	//
	// example:
	//
	// 42
	Seed *int64 `json:"seed,omitempty" xml:"seed,omitempty"`
	// Specifies whether to skip special tokens.
	//
	// example:
	//
	// true
	SkipSpecialTokens *bool `json:"skipSpecialTokens,omitempty" xml:"skipSpecialTokens,omitempty"`
	// The stop sequences. This parameter is compatible with a single string or an array of strings. The array format is shown here.
	//
	// example:
	//
	// ["DONE"]
	Stop []*string `json:"stop,omitempty" xml:"stop,omitempty" type:"Repeated"`
	// The list of vLLM stop token IDs.
	//
	// example:
	//
	// [1,2]
	StopTokenIds []*int64 `json:"stopTokenIds,omitempty" xml:"stopTokenIds,omitempty" type:"Repeated"`
	// Specifies whether to use Server-Sent Events (SSE) for streaming responses. You must set this parameter to true for this operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// The OpenAI-compatible streaming options.
	StreamOptions *GuiChatCompletionStreamRequestStreamOptions `json:"streamOptions,omitempty" xml:"streamOptions,omitempty" type:"Struct"`
	// The vLLM structured output constraints.
	StructuredOutputs *GuiChatCompletionStreamRequestStructuredOutputs `json:"structuredOutputs,omitempty" xml:"structuredOutputs,omitempty" type:"Struct"`
	// The sampling temperature, which overrides the default value of the Mobile-Use Agent primary model.
	//
	// example:
	//
	// 0.2
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// The vLLM Top-K sampling parameter.
	//
	// example:
	//
	// 20
	TopK *int64 `json:"topK,omitempty" xml:"topK,omitempty"`
	// A model-compatible parameter. The JSON field is top_logprobs. Whether this parameter is accepted depends on the model. The current Mobile-Use Agent response does not return candidate token probability details.
	//
	// example:
	//
	// 3
	TopLogprobs *int64 `json:"topLogprobs,omitempty" xml:"topLogprobs,omitempty"`
	// The Top-P nucleus sampling parameter, which overrides the default value of the Mobile-Use Agent primary model.
	//
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
	// chat_template_kwargs.enable_thinking: the request-level switch for model thinking, which overrides the default value with the same name. Enabling this switch does not cause the Mobile-Use Agent response to return reasoning_content.
	//
	// example:
	//
	// true
	EnableThinking *bool `json:"enableThinking,omitempty" xml:"enableThinking,omitempty"`
	// Specifies whether to preserve the existing thinking context in subsequent rounds. This parameter is passed as-is according to the vLLM or Qwen chat_template_kwargs.
	//
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
	// The message content. For a user message, use a content array that contains exactly one current screenshot and optionally attached task text. Screenshots support public HTTPS URLs or Base64 image data URLs. When role is set to tool, fill in the actual execution result of the previous action, which can be a non-empty string, object, or array.
	Content []*GuiChatCompletionStreamRequestMessagesContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// The message role. The current input must end with a user message. If an execution receipt from the previous round exists, the preceding message must be a tool message.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// Required when role is set to tool. References the tool_call ID from the previous Mobile-Use Agent response.
	//
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
	// The compatible inline image field for input_image. You can pass a Base64 data URL.
	//
	// example:
	//
	// data:image/png;base64,...
	ImageData *string `json:"imageData,omitempty" xml:"imageData,omitempty"`
	// The screenshot reference. Supports an image_url string or an object containing a url field. The object format is shown here.
	ImageUrl *GuiChatCompletionStreamRequestMessagesContentImageUrl `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Struct"`
	// The text content, which can describe the current operation objective or supplement the current page state.
	//
	// example:
	//
	// Click the search button
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// The type of the content item. Use text or input_text for text. Use image_url or input_image for screenshots.
	//
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
	// A publicly accessible HTTPS image URL or a Base64 image data URL. Each round of user messages must and can only contain one current screenshot. Expired or inaccessible image addresses are rejected.
	//
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
	// The compatible alias for available_apps, which is also an array of app name strings.
	//
	// example:
	//
	// ["Chrome","Settings"]
	AppList []*string `json:"appList,omitempty" xml:"appList,omitempty" type:"Repeated"`
	// The supplementary list of application names for the Standard Edition of Mobile-Use Agent. The JSON field is available_apps. The newly passed list replaces the previous supplementary items, but the default applications are retained. An empty array does not clear the default applications. The Dedicated Edition of Mobile-Use Agent does not use this list.
	//
	// example:
	//
	// ["Chrome","Settings"]
	AvailableApps []*string `json:"availableApps,omitempty" xml:"availableApps,omitempty" type:"Repeated"`
	// The Mobile-Use Agent Harness receipt text for legacy clients. This cannot be passed at the same time as the preceding tool message. New clients should prioritize using [tool,user] messages.
	//
	// example:
	//
	// previous action completed
	HarnessMessage *string `json:"harnessMessage,omitempty" xml:"harnessMessage,omitempty"`
	// The screen height as a positive integer. For Base64 images, the width and height are parsed from the image itself. For URL images, this value is used for the runtime screen size.
	//
	// example:
	//
	// 1080
	ScreenHeight *int64 `json:"screenHeight,omitempty" xml:"screenHeight,omitempty"`
	// The screen width as a positive integer. For Base64 images, the width and height are parsed from the image itself. For URL images, this value is used for the runtime screen size.
	//
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
	// The maximum number of dynamic image patches allowed by the vision model.
	//
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
	// The response format type, such as text, json_object, or json_schema.
	//
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
	// Specifies whether to output the usage statistics block before data:[DONE].
	//
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
	// The set of allowed enum text values.
	//
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
