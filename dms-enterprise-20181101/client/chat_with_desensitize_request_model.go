// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithDesensitizeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioJson(v string) *ChatWithDesensitizeRequest
	GetAudioJson() *string
	SetDesensitizationRule(v string) *ChatWithDesensitizeRequest
	GetDesensitizationRule() *string
	SetEnableCodeInterpreter(v bool) *ChatWithDesensitizeRequest
	GetEnableCodeInterpreter() *bool
	SetEnableSearch(v bool) *ChatWithDesensitizeRequest
	GetEnableSearch() *bool
	SetEnableThinking(v bool) *ChatWithDesensitizeRequest
	GetEnableThinking() *bool
	SetInstanceId(v int64) *ChatWithDesensitizeRequest
	GetInstanceId() *int64
	SetLogprobs(v bool) *ChatWithDesensitizeRequest
	GetLogprobs() *bool
	SetMaxTokens(v int32) *ChatWithDesensitizeRequest
	GetMaxTokens() *int32
	SetMessages(v []interface{}) *ChatWithDesensitizeRequest
	GetMessages() []interface{}
	SetModalitiesList(v []*string) *ChatWithDesensitizeRequest
	GetModalitiesList() []*string
	SetModel(v string) *ChatWithDesensitizeRequest
	GetModel() *string
	SetNeedDesensitization(v bool) *ChatWithDesensitizeRequest
	GetNeedDesensitization() *bool
	SetPresencePenalty(v string) *ChatWithDesensitizeRequest
	GetPresencePenalty() *string
	SetResponseFormat(v string) *ChatWithDesensitizeRequest
	GetResponseFormat() *string
	SetSearchOptions(v map[string]*string) *ChatWithDesensitizeRequest
	GetSearchOptions() map[string]*string
	SetSeed(v int32) *ChatWithDesensitizeRequest
	GetSeed() *int32
	SetStop(v []*string) *ChatWithDesensitizeRequest
	GetStop() []*string
	SetTemperature(v string) *ChatWithDesensitizeRequest
	GetTemperature() *string
	SetThinkingBudget(v int32) *ChatWithDesensitizeRequest
	GetThinkingBudget() *int32
	SetTopK(v int32) *ChatWithDesensitizeRequest
	GetTopK() *int32
	SetTopLogprobs(v int32) *ChatWithDesensitizeRequest
	GetTopLogprobs() *int32
	SetTopP(v string) *ChatWithDesensitizeRequest
	GetTopP() *string
	SetVlHighResolutionImages(v bool) *ChatWithDesensitizeRequest
	GetVlHighResolutionImages() *bool
	SetXDashScopeDataInspection(v string) *ChatWithDesensitizeRequest
	GetXDashScopeDataInspection() *string
}

type ChatWithDesensitizeRequest struct {
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
	Messages []interface{} `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// Output data modality; only applicable to the Qwen-Omni model.
	//
	// example:
	//
	// ["text","audio"]
	ModalitiesList []*string `json:"ModalitiesList,omitempty" xml:"ModalitiesList,omitempty" type:"Repeated"`
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
	SearchOptions map[string]*string `json:"SearchOptions,omitempty" xml:"SearchOptions,omitempty"`
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
	Stop []*string `json:"Stop,omitempty" xml:"Stop,omitempty" type:"Repeated"`
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

func (s ChatWithDesensitizeRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatWithDesensitizeRequest) GoString() string {
	return s.String()
}

func (s *ChatWithDesensitizeRequest) GetAudioJson() *string {
	return s.AudioJson
}

func (s *ChatWithDesensitizeRequest) GetDesensitizationRule() *string {
	return s.DesensitizationRule
}

func (s *ChatWithDesensitizeRequest) GetEnableCodeInterpreter() *bool {
	return s.EnableCodeInterpreter
}

func (s *ChatWithDesensitizeRequest) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *ChatWithDesensitizeRequest) GetEnableThinking() *bool {
	return s.EnableThinking
}

func (s *ChatWithDesensitizeRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ChatWithDesensitizeRequest) GetLogprobs() *bool {
	return s.Logprobs
}

func (s *ChatWithDesensitizeRequest) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *ChatWithDesensitizeRequest) GetMessages() []interface{} {
	return s.Messages
}

func (s *ChatWithDesensitizeRequest) GetModalitiesList() []*string {
	return s.ModalitiesList
}

func (s *ChatWithDesensitizeRequest) GetModel() *string {
	return s.Model
}

func (s *ChatWithDesensitizeRequest) GetNeedDesensitization() *bool {
	return s.NeedDesensitization
}

func (s *ChatWithDesensitizeRequest) GetPresencePenalty() *string {
	return s.PresencePenalty
}

func (s *ChatWithDesensitizeRequest) GetResponseFormat() *string {
	return s.ResponseFormat
}

func (s *ChatWithDesensitizeRequest) GetSearchOptions() map[string]*string {
	return s.SearchOptions
}

func (s *ChatWithDesensitizeRequest) GetSeed() *int32 {
	return s.Seed
}

func (s *ChatWithDesensitizeRequest) GetStop() []*string {
	return s.Stop
}

func (s *ChatWithDesensitizeRequest) GetTemperature() *string {
	return s.Temperature
}

func (s *ChatWithDesensitizeRequest) GetThinkingBudget() *int32 {
	return s.ThinkingBudget
}

func (s *ChatWithDesensitizeRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *ChatWithDesensitizeRequest) GetTopLogprobs() *int32 {
	return s.TopLogprobs
}

func (s *ChatWithDesensitizeRequest) GetTopP() *string {
	return s.TopP
}

func (s *ChatWithDesensitizeRequest) GetVlHighResolutionImages() *bool {
	return s.VlHighResolutionImages
}

func (s *ChatWithDesensitizeRequest) GetXDashScopeDataInspection() *string {
	return s.XDashScopeDataInspection
}

func (s *ChatWithDesensitizeRequest) SetAudioJson(v string) *ChatWithDesensitizeRequest {
	s.AudioJson = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetDesensitizationRule(v string) *ChatWithDesensitizeRequest {
	s.DesensitizationRule = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetEnableCodeInterpreter(v bool) *ChatWithDesensitizeRequest {
	s.EnableCodeInterpreter = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetEnableSearch(v bool) *ChatWithDesensitizeRequest {
	s.EnableSearch = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetEnableThinking(v bool) *ChatWithDesensitizeRequest {
	s.EnableThinking = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetInstanceId(v int64) *ChatWithDesensitizeRequest {
	s.InstanceId = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetLogprobs(v bool) *ChatWithDesensitizeRequest {
	s.Logprobs = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetMaxTokens(v int32) *ChatWithDesensitizeRequest {
	s.MaxTokens = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetMessages(v []interface{}) *ChatWithDesensitizeRequest {
	s.Messages = v
	return s
}

func (s *ChatWithDesensitizeRequest) SetModalitiesList(v []*string) *ChatWithDesensitizeRequest {
	s.ModalitiesList = v
	return s
}

func (s *ChatWithDesensitizeRequest) SetModel(v string) *ChatWithDesensitizeRequest {
	s.Model = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetNeedDesensitization(v bool) *ChatWithDesensitizeRequest {
	s.NeedDesensitization = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetPresencePenalty(v string) *ChatWithDesensitizeRequest {
	s.PresencePenalty = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetResponseFormat(v string) *ChatWithDesensitizeRequest {
	s.ResponseFormat = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetSearchOptions(v map[string]*string) *ChatWithDesensitizeRequest {
	s.SearchOptions = v
	return s
}

func (s *ChatWithDesensitizeRequest) SetSeed(v int32) *ChatWithDesensitizeRequest {
	s.Seed = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetStop(v []*string) *ChatWithDesensitizeRequest {
	s.Stop = v
	return s
}

func (s *ChatWithDesensitizeRequest) SetTemperature(v string) *ChatWithDesensitizeRequest {
	s.Temperature = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetThinkingBudget(v int32) *ChatWithDesensitizeRequest {
	s.ThinkingBudget = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetTopK(v int32) *ChatWithDesensitizeRequest {
	s.TopK = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetTopLogprobs(v int32) *ChatWithDesensitizeRequest {
	s.TopLogprobs = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetTopP(v string) *ChatWithDesensitizeRequest {
	s.TopP = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetVlHighResolutionImages(v bool) *ChatWithDesensitizeRequest {
	s.VlHighResolutionImages = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetXDashScopeDataInspection(v string) *ChatWithDesensitizeRequest {
	s.XDashScopeDataInspection = &v
	return s
}

func (s *ChatWithDesensitizeRequest) Validate() error {
	return dara.Validate(s)
}
