// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithDesensitizeSSEShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioJson(v string) *ChatWithDesensitizeSSEShrinkRequest
	GetAudioJson() *string
	SetDesensitizationRule(v string) *ChatWithDesensitizeSSEShrinkRequest
	GetDesensitizationRule() *string
	SetDimensions(v int32) *ChatWithDesensitizeSSEShrinkRequest
	GetDimensions() *int32
	SetEnableCodeInterpreter(v bool) *ChatWithDesensitizeSSEShrinkRequest
	GetEnableCodeInterpreter() *bool
	SetEnableSearch(v bool) *ChatWithDesensitizeSSEShrinkRequest
	GetEnableSearch() *bool
	SetEnableThinking(v bool) *ChatWithDesensitizeSSEShrinkRequest
	GetEnableThinking() *bool
	SetIncludeUsage(v bool) *ChatWithDesensitizeSSEShrinkRequest
	GetIncludeUsage() *bool
	SetInput(v string) *ChatWithDesensitizeSSEShrinkRequest
	GetInput() *string
	SetInstanceId(v int64) *ChatWithDesensitizeSSEShrinkRequest
	GetInstanceId() *int64
	SetLogprobs(v bool) *ChatWithDesensitizeSSEShrinkRequest
	GetLogprobs() *bool
	SetMaxTokens(v int32) *ChatWithDesensitizeSSEShrinkRequest
	GetMaxTokens() *int32
	SetMessagesShrink(v string) *ChatWithDesensitizeSSEShrinkRequest
	GetMessagesShrink() *string
	SetModalitiesListShrink(v string) *ChatWithDesensitizeSSEShrinkRequest
	GetModalitiesListShrink() *string
	SetModel(v string) *ChatWithDesensitizeSSEShrinkRequest
	GetModel() *string
	SetNeedDesensitization(v bool) *ChatWithDesensitizeSSEShrinkRequest
	GetNeedDesensitization() *bool
	SetParameters(v string) *ChatWithDesensitizeSSEShrinkRequest
	GetParameters() *string
	SetPresencePenalty(v string) *ChatWithDesensitizeSSEShrinkRequest
	GetPresencePenalty() *string
	SetResponseFormat(v string) *ChatWithDesensitizeSSEShrinkRequest
	GetResponseFormat() *string
	SetSearchOptionsShrink(v string) *ChatWithDesensitizeSSEShrinkRequest
	GetSearchOptionsShrink() *string
	SetSeed(v int32) *ChatWithDesensitizeSSEShrinkRequest
	GetSeed() *int32
	SetStopShrink(v string) *ChatWithDesensitizeSSEShrinkRequest
	GetStopShrink() *string
	SetStream(v bool) *ChatWithDesensitizeSSEShrinkRequest
	GetStream() *bool
	SetTemperature(v string) *ChatWithDesensitizeSSEShrinkRequest
	GetTemperature() *string
	SetThinkingBudget(v int32) *ChatWithDesensitizeSSEShrinkRequest
	GetThinkingBudget() *int32
	SetTopK(v int32) *ChatWithDesensitizeSSEShrinkRequest
	GetTopK() *int32
	SetTopLogprobs(v int32) *ChatWithDesensitizeSSEShrinkRequest
	GetTopLogprobs() *int32
	SetTopP(v string) *ChatWithDesensitizeSSEShrinkRequest
	GetTopP() *string
	SetVlHighResolutionImages(v bool) *ChatWithDesensitizeSSEShrinkRequest
	GetVlHighResolutionImages() *bool
	SetXDashScopeDataInspection(v string) *ChatWithDesensitizeSSEShrinkRequest
	GetXDashScopeDataInspection() *string
}

type ChatWithDesensitizeSSEShrinkRequest struct {
	// Output audio voice and format; only applicable to the Qwen-Omni model, provided that the modalities parameter is set to ["text", "audio"].
	//
	// example:
	//
	// {}
	AudioJson *string `json:"AudioJson,omitempty" xml:"AudioJson,omitempty"`
	// Masking category. Required when needDataMasking is true.
	//
	// example:
	//
	// UserInfo
	DesensitizationRule *string `json:"DesensitizationRule,omitempty" xml:"DesensitizationRule,omitempty"`
	// The embedding dimensions.
	//
	// example:
	//
	// 256
	Dimensions *int32 `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// Specifies whether to enable the code interpreter feature. Takes effect only when model is qwen3-max-preview and enable_thinking is true.
	//
	// example:
	//
	// false
	EnableCodeInterpreter *bool `json:"EnableCodeInterpreter,omitempty" xml:"EnableCodeInterpreter,omitempty"`
	// Whether to enable web search.
	//
	// example:
	//
	// false
	EnableSearch *bool `json:"EnableSearch,omitempty" xml:"EnableSearch,omitempty"`
	// Specifies whether to enable thinking mode when using hybrid thinking models.
	//
	// example:
	//
	// true
	EnableThinking *bool `json:"EnableThinking,omitempty" xml:"EnableThinking,omitempty"`
	// Specifies whether to include token usage information in the final chunk of the streaming response.
	//
	// example:
	//
	// true
	IncludeUsage *bool `json:"IncludeUsage,omitempty" xml:"IncludeUsage,omitempty"`
	// The input to the embedding model.
	//
	// example:
	//
	// test
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The ID of the instance, used to specify the corresponding data masking rules. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123***
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to return the log probabilities of the output tokens.
	//
	// example:
	//
	// false
	Logprobs *bool `json:"Logprobs,omitempty" xml:"Logprobs,omitempty"`
	// Limits the maximum number of tokens the model can generate. If the output exceeds this value, generation will be truncated. Suitable for scenarios where you need to control the output length.
	//
	// example:
	//
	// 256
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// The conversation context passed to the model, arranged in chronological order.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "content": "你好",
	//
	//         "role": "user"
	//
	//     }
	//
	// ]
	MessagesShrink *string `json:"Messages,omitempty" xml:"Messages,omitempty"`
	// Output data modality; only applicable to the Qwen-Omni model.
	//
	// example:
	//
	// ["text","audio"]
	ModalitiesListShrink *string `json:"ModalitiesList,omitempty" xml:"ModalitiesList,omitempty"`
	// The model name. Supported models: Qwen series text-only Large Language Models (commercial and open-source). Multi-modal models are not supported.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Whether to enable data masking. Defaults to false.
	//
	// example:
	//
	// false
	NeedDesensitization *bool `json:"NeedDesensitization,omitempty" xml:"NeedDesensitization,omitempty"`
	// Model configuration parameters.
	//
	// example:
	//
	// {}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// Controls the degree of repetition in generated text. Valid values: [-2.0, 2.0]. Positive values decrease repetition, while negative values increase it.
	//
	// example:
	//
	// 0.0
	PresencePenalty *string `json:"PresencePenalty,omitempty" xml:"PresencePenalty,omitempty"`
	// The format of the returned content. Valid values: text: Plain text response; json_object: Standardized JSON string.
	//
	// example:
	//
	// text
	ResponseFormat *string `json:"ResponseFormat,omitempty" xml:"ResponseFormat,omitempty"`
	// Web search strategy.
	//
	// example:
	//
	// {}
	SearchOptionsShrink *string `json:"SearchOptions,omitempty" xml:"SearchOptions,omitempty"`
	// Random seed. Used to ensure the reproducibility of results under the same input and parameters. Valid values: [0, 2^31−1].
	//
	// example:
	//
	// 1
	Seed *int32 `json:"Seed,omitempty" xml:"Seed,omitempty"`
	// Stop sequences.
	StopShrink *string `json:"Stop,omitempty" xml:"Stop,omitempty"`
	// Specifies whether to use streaming output.
	//
	// example:
	//
	// 1-68f11da7e2b826dcc63c5877-hd
	Stream *bool `json:"Stream,omitempty" xml:"Stream,omitempty"`
	// The sampling temperature controls the diversity of the generated text. The higher the temperature, the more diverse the generated text, and conversely, the more deterministic the generated text. Valid values: [0, 2).
	//
	// example:
	//
	// 1
	Temperature *string `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// The maximum number of tokens allowed for the model\\"s internal reasoning process.
	//
	// example:
	//
	// 256
	ThinkingBudget *int32 `json:"ThinkingBudget,omitempty" xml:"ThinkingBudget,omitempty"`
	// Specifies the number of candidate tokens to consider during sampling. Higher values increase randomness, while lower values make the output more deterministic. Set to null or a value greater than 100 to disable.
	//
	// example:
	//
	// 10
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// Specifies the number of most likely candidate tokens to return at each generation step. Valid values: [0, 5].
	//
	// example:
	//
	// 1
	TopLogprobs *int32 `json:"TopLogprobs,omitempty" xml:"TopLogprobs,omitempty"`
	// The probability threshold for nucleus sampling, used to control the diversity of the generated text. Higher Top-P values result in more diverse generated text. Valid values: (0,1.0].
	//
	// example:
	//
	// 0.5
	TopP *string `json:"TopP,omitempty" xml:"TopP,omitempty"`
	// Specifies whether to increase the maximum pixel limit of input images to the equivalent of 16,384 tokens.
	//
	// example:
	//
	// false
	VlHighResolutionImages *bool `json:"VlHighResolutionImages,omitempty" xml:"VlHighResolutionImages,omitempty"`
	// Specifies whether to further identify non-compliant information in both input and output content, building upon the built-in content safety capabilities of the Tongyi Qianwen API.
	//
	// example:
	//
	// {}
	XDashScopeDataInspection *string `json:"XDashScopeDataInspection,omitempty" xml:"XDashScopeDataInspection,omitempty"`
}

func (s ChatWithDesensitizeSSEShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatWithDesensitizeSSEShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetAudioJson() *string {
	return s.AudioJson
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetDesensitizationRule() *string {
	return s.DesensitizationRule
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetDimensions() *int32 {
	return s.Dimensions
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetEnableCodeInterpreter() *bool {
	return s.EnableCodeInterpreter
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetEnableThinking() *bool {
	return s.EnableThinking
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetIncludeUsage() *bool {
	return s.IncludeUsage
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetInput() *string {
	return s.Input
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetLogprobs() *bool {
	return s.Logprobs
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetMessagesShrink() *string {
	return s.MessagesShrink
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetModalitiesListShrink() *string {
	return s.ModalitiesListShrink
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetNeedDesensitization() *bool {
	return s.NeedDesensitization
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetPresencePenalty() *string {
	return s.PresencePenalty
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetResponseFormat() *string {
	return s.ResponseFormat
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetSearchOptionsShrink() *string {
	return s.SearchOptionsShrink
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetSeed() *int32 {
	return s.Seed
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetStopShrink() *string {
	return s.StopShrink
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetStream() *bool {
	return s.Stream
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetTemperature() *string {
	return s.Temperature
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetThinkingBudget() *int32 {
	return s.ThinkingBudget
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetTopLogprobs() *int32 {
	return s.TopLogprobs
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetTopP() *string {
	return s.TopP
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetVlHighResolutionImages() *bool {
	return s.VlHighResolutionImages
}

func (s *ChatWithDesensitizeSSEShrinkRequest) GetXDashScopeDataInspection() *string {
	return s.XDashScopeDataInspection
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetAudioJson(v string) *ChatWithDesensitizeSSEShrinkRequest {
	s.AudioJson = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetDesensitizationRule(v string) *ChatWithDesensitizeSSEShrinkRequest {
	s.DesensitizationRule = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetDimensions(v int32) *ChatWithDesensitizeSSEShrinkRequest {
	s.Dimensions = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetEnableCodeInterpreter(v bool) *ChatWithDesensitizeSSEShrinkRequest {
	s.EnableCodeInterpreter = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetEnableSearch(v bool) *ChatWithDesensitizeSSEShrinkRequest {
	s.EnableSearch = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetEnableThinking(v bool) *ChatWithDesensitizeSSEShrinkRequest {
	s.EnableThinking = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetIncludeUsage(v bool) *ChatWithDesensitizeSSEShrinkRequest {
	s.IncludeUsage = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetInput(v string) *ChatWithDesensitizeSSEShrinkRequest {
	s.Input = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetInstanceId(v int64) *ChatWithDesensitizeSSEShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetLogprobs(v bool) *ChatWithDesensitizeSSEShrinkRequest {
	s.Logprobs = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetMaxTokens(v int32) *ChatWithDesensitizeSSEShrinkRequest {
	s.MaxTokens = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetMessagesShrink(v string) *ChatWithDesensitizeSSEShrinkRequest {
	s.MessagesShrink = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetModalitiesListShrink(v string) *ChatWithDesensitizeSSEShrinkRequest {
	s.ModalitiesListShrink = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetModel(v string) *ChatWithDesensitizeSSEShrinkRequest {
	s.Model = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetNeedDesensitization(v bool) *ChatWithDesensitizeSSEShrinkRequest {
	s.NeedDesensitization = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetParameters(v string) *ChatWithDesensitizeSSEShrinkRequest {
	s.Parameters = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetPresencePenalty(v string) *ChatWithDesensitizeSSEShrinkRequest {
	s.PresencePenalty = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetResponseFormat(v string) *ChatWithDesensitizeSSEShrinkRequest {
	s.ResponseFormat = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetSearchOptionsShrink(v string) *ChatWithDesensitizeSSEShrinkRequest {
	s.SearchOptionsShrink = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetSeed(v int32) *ChatWithDesensitizeSSEShrinkRequest {
	s.Seed = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetStopShrink(v string) *ChatWithDesensitizeSSEShrinkRequest {
	s.StopShrink = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetStream(v bool) *ChatWithDesensitizeSSEShrinkRequest {
	s.Stream = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetTemperature(v string) *ChatWithDesensitizeSSEShrinkRequest {
	s.Temperature = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetThinkingBudget(v int32) *ChatWithDesensitizeSSEShrinkRequest {
	s.ThinkingBudget = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetTopK(v int32) *ChatWithDesensitizeSSEShrinkRequest {
	s.TopK = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetTopLogprobs(v int32) *ChatWithDesensitizeSSEShrinkRequest {
	s.TopLogprobs = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetTopP(v string) *ChatWithDesensitizeSSEShrinkRequest {
	s.TopP = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetVlHighResolutionImages(v bool) *ChatWithDesensitizeSSEShrinkRequest {
	s.VlHighResolutionImages = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) SetXDashScopeDataInspection(v string) *ChatWithDesensitizeSSEShrinkRequest {
	s.XDashScopeDataInspection = &v
	return s
}

func (s *ChatWithDesensitizeSSEShrinkRequest) Validate() error {
	return dara.Validate(s)
}
