// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithDesensitizeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioJson(v string) *ChatWithDesensitizeShrinkRequest
	GetAudioJson() *string
	SetDesensitizationRule(v string) *ChatWithDesensitizeShrinkRequest
	GetDesensitizationRule() *string
	SetEnableCodeInterpreter(v bool) *ChatWithDesensitizeShrinkRequest
	GetEnableCodeInterpreter() *bool
	SetEnableSearch(v bool) *ChatWithDesensitizeShrinkRequest
	GetEnableSearch() *bool
	SetEnableThinking(v bool) *ChatWithDesensitizeShrinkRequest
	GetEnableThinking() *bool
	SetInstanceId(v int64) *ChatWithDesensitizeShrinkRequest
	GetInstanceId() *int64
	SetLogprobs(v bool) *ChatWithDesensitizeShrinkRequest
	GetLogprobs() *bool
	SetMaxTokens(v int32) *ChatWithDesensitizeShrinkRequest
	GetMaxTokens() *int32
	SetMessagesShrink(v string) *ChatWithDesensitizeShrinkRequest
	GetMessagesShrink() *string
	SetModalitiesListShrink(v string) *ChatWithDesensitizeShrinkRequest
	GetModalitiesListShrink() *string
	SetModel(v string) *ChatWithDesensitizeShrinkRequest
	GetModel() *string
	SetNeedDesensitization(v bool) *ChatWithDesensitizeShrinkRequest
	GetNeedDesensitization() *bool
	SetPresencePenalty(v string) *ChatWithDesensitizeShrinkRequest
	GetPresencePenalty() *string
	SetResponseFormat(v string) *ChatWithDesensitizeShrinkRequest
	GetResponseFormat() *string
	SetSearchOptionsShrink(v string) *ChatWithDesensitizeShrinkRequest
	GetSearchOptionsShrink() *string
	SetSeed(v int32) *ChatWithDesensitizeShrinkRequest
	GetSeed() *int32
	SetStopShrink(v string) *ChatWithDesensitizeShrinkRequest
	GetStopShrink() *string
	SetTemperature(v string) *ChatWithDesensitizeShrinkRequest
	GetTemperature() *string
	SetThinkingBudget(v int32) *ChatWithDesensitizeShrinkRequest
	GetThinkingBudget() *int32
	SetTopK(v int32) *ChatWithDesensitizeShrinkRequest
	GetTopK() *int32
	SetTopLogprobs(v int32) *ChatWithDesensitizeShrinkRequest
	GetTopLogprobs() *int32
	SetTopP(v string) *ChatWithDesensitizeShrinkRequest
	GetTopP() *string
	SetVlHighResolutionImages(v bool) *ChatWithDesensitizeShrinkRequest
	GetVlHighResolutionImages() *bool
	SetXDashScopeDataInspection(v string) *ChatWithDesensitizeShrinkRequest
	GetXDashScopeDataInspection() *string
}

type ChatWithDesensitizeShrinkRequest struct {
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
	// Specifies whether to enable Thinking Mode when using hybrid thinking models.
	//
	// example:
	//
	// true
	EnableThinking *bool `json:"EnableThinking,omitempty" xml:"EnableThinking,omitempty"`
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
	// The model name. Supported Models: Qwen series text-only Large Language Models (Commercial and Open-source). Multi-modal models are not supported.
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
	//
	// example:
	//
	// ["\\n"]
	StopShrink *string `json:"Stop,omitempty" xml:"Stop,omitempty"`
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

func (s ChatWithDesensitizeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatWithDesensitizeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChatWithDesensitizeShrinkRequest) GetAudioJson() *string {
	return s.AudioJson
}

func (s *ChatWithDesensitizeShrinkRequest) GetDesensitizationRule() *string {
	return s.DesensitizationRule
}

func (s *ChatWithDesensitizeShrinkRequest) GetEnableCodeInterpreter() *bool {
	return s.EnableCodeInterpreter
}

func (s *ChatWithDesensitizeShrinkRequest) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *ChatWithDesensitizeShrinkRequest) GetEnableThinking() *bool {
	return s.EnableThinking
}

func (s *ChatWithDesensitizeShrinkRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ChatWithDesensitizeShrinkRequest) GetLogprobs() *bool {
	return s.Logprobs
}

func (s *ChatWithDesensitizeShrinkRequest) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *ChatWithDesensitizeShrinkRequest) GetMessagesShrink() *string {
	return s.MessagesShrink
}

func (s *ChatWithDesensitizeShrinkRequest) GetModalitiesListShrink() *string {
	return s.ModalitiesListShrink
}

func (s *ChatWithDesensitizeShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *ChatWithDesensitizeShrinkRequest) GetNeedDesensitization() *bool {
	return s.NeedDesensitization
}

func (s *ChatWithDesensitizeShrinkRequest) GetPresencePenalty() *string {
	return s.PresencePenalty
}

func (s *ChatWithDesensitizeShrinkRequest) GetResponseFormat() *string {
	return s.ResponseFormat
}

func (s *ChatWithDesensitizeShrinkRequest) GetSearchOptionsShrink() *string {
	return s.SearchOptionsShrink
}

func (s *ChatWithDesensitizeShrinkRequest) GetSeed() *int32 {
	return s.Seed
}

func (s *ChatWithDesensitizeShrinkRequest) GetStopShrink() *string {
	return s.StopShrink
}

func (s *ChatWithDesensitizeShrinkRequest) GetTemperature() *string {
	return s.Temperature
}

func (s *ChatWithDesensitizeShrinkRequest) GetThinkingBudget() *int32 {
	return s.ThinkingBudget
}

func (s *ChatWithDesensitizeShrinkRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *ChatWithDesensitizeShrinkRequest) GetTopLogprobs() *int32 {
	return s.TopLogprobs
}

func (s *ChatWithDesensitizeShrinkRequest) GetTopP() *string {
	return s.TopP
}

func (s *ChatWithDesensitizeShrinkRequest) GetVlHighResolutionImages() *bool {
	return s.VlHighResolutionImages
}

func (s *ChatWithDesensitizeShrinkRequest) GetXDashScopeDataInspection() *string {
	return s.XDashScopeDataInspection
}

func (s *ChatWithDesensitizeShrinkRequest) SetAudioJson(v string) *ChatWithDesensitizeShrinkRequest {
	s.AudioJson = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetDesensitizationRule(v string) *ChatWithDesensitizeShrinkRequest {
	s.DesensitizationRule = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetEnableCodeInterpreter(v bool) *ChatWithDesensitizeShrinkRequest {
	s.EnableCodeInterpreter = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetEnableSearch(v bool) *ChatWithDesensitizeShrinkRequest {
	s.EnableSearch = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetEnableThinking(v bool) *ChatWithDesensitizeShrinkRequest {
	s.EnableThinking = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetInstanceId(v int64) *ChatWithDesensitizeShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetLogprobs(v bool) *ChatWithDesensitizeShrinkRequest {
	s.Logprobs = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetMaxTokens(v int32) *ChatWithDesensitizeShrinkRequest {
	s.MaxTokens = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetMessagesShrink(v string) *ChatWithDesensitizeShrinkRequest {
	s.MessagesShrink = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetModalitiesListShrink(v string) *ChatWithDesensitizeShrinkRequest {
	s.ModalitiesListShrink = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetModel(v string) *ChatWithDesensitizeShrinkRequest {
	s.Model = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetNeedDesensitization(v bool) *ChatWithDesensitizeShrinkRequest {
	s.NeedDesensitization = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetPresencePenalty(v string) *ChatWithDesensitizeShrinkRequest {
	s.PresencePenalty = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetResponseFormat(v string) *ChatWithDesensitizeShrinkRequest {
	s.ResponseFormat = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetSearchOptionsShrink(v string) *ChatWithDesensitizeShrinkRequest {
	s.SearchOptionsShrink = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetSeed(v int32) *ChatWithDesensitizeShrinkRequest {
	s.Seed = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetStopShrink(v string) *ChatWithDesensitizeShrinkRequest {
	s.StopShrink = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetTemperature(v string) *ChatWithDesensitizeShrinkRequest {
	s.Temperature = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetThinkingBudget(v int32) *ChatWithDesensitizeShrinkRequest {
	s.ThinkingBudget = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetTopK(v int32) *ChatWithDesensitizeShrinkRequest {
	s.TopK = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetTopLogprobs(v int32) *ChatWithDesensitizeShrinkRequest {
	s.TopLogprobs = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetTopP(v string) *ChatWithDesensitizeShrinkRequest {
	s.TopP = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetVlHighResolutionImages(v bool) *ChatWithDesensitizeShrinkRequest {
	s.VlHighResolutionImages = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetXDashScopeDataInspection(v string) *ChatWithDesensitizeShrinkRequest {
	s.XDashScopeDataInspection = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
