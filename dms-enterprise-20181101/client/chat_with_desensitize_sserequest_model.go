// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithDesensitizeSSERequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioJson(v string) *ChatWithDesensitizeSSERequest
	GetAudioJson() *string
	SetDesensitizationRule(v string) *ChatWithDesensitizeSSERequest
	GetDesensitizationRule() *string
	SetDimensions(v int32) *ChatWithDesensitizeSSERequest
	GetDimensions() *int32
	SetEnableCodeInterpreter(v bool) *ChatWithDesensitizeSSERequest
	GetEnableCodeInterpreter() *bool
	SetEnableSearch(v bool) *ChatWithDesensitizeSSERequest
	GetEnableSearch() *bool
	SetEnableThinking(v bool) *ChatWithDesensitizeSSERequest
	GetEnableThinking() *bool
	SetIncludeUsage(v bool) *ChatWithDesensitizeSSERequest
	GetIncludeUsage() *bool
	SetInput(v string) *ChatWithDesensitizeSSERequest
	GetInput() *string
	SetInstanceId(v int64) *ChatWithDesensitizeSSERequest
	GetInstanceId() *int64
	SetLogprobs(v bool) *ChatWithDesensitizeSSERequest
	GetLogprobs() *bool
	SetMaxTokens(v int32) *ChatWithDesensitizeSSERequest
	GetMaxTokens() *int32
	SetMessages(v []interface{}) *ChatWithDesensitizeSSERequest
	GetMessages() []interface{}
	SetModalitiesList(v []*string) *ChatWithDesensitizeSSERequest
	GetModalitiesList() []*string
	SetModel(v string) *ChatWithDesensitizeSSERequest
	GetModel() *string
	SetNeedDesensitization(v bool) *ChatWithDesensitizeSSERequest
	GetNeedDesensitization() *bool
	SetParameters(v string) *ChatWithDesensitizeSSERequest
	GetParameters() *string
	SetPresencePenalty(v string) *ChatWithDesensitizeSSERequest
	GetPresencePenalty() *string
	SetResponseFormat(v string) *ChatWithDesensitizeSSERequest
	GetResponseFormat() *string
	SetSearchOptions(v map[string]*string) *ChatWithDesensitizeSSERequest
	GetSearchOptions() map[string]*string
	SetSeed(v int32) *ChatWithDesensitizeSSERequest
	GetSeed() *int32
	SetStop(v []*string) *ChatWithDesensitizeSSERequest
	GetStop() []*string
	SetStream(v bool) *ChatWithDesensitizeSSERequest
	GetStream() *bool
	SetTemperature(v string) *ChatWithDesensitizeSSERequest
	GetTemperature() *string
	SetThinkingBudget(v int32) *ChatWithDesensitizeSSERequest
	GetThinkingBudget() *int32
	SetTopK(v int32) *ChatWithDesensitizeSSERequest
	GetTopK() *int32
	SetTopLogprobs(v int32) *ChatWithDesensitizeSSERequest
	GetTopLogprobs() *int32
	SetTopP(v string) *ChatWithDesensitizeSSERequest
	GetTopP() *string
	SetVlHighResolutionImages(v bool) *ChatWithDesensitizeSSERequest
	GetVlHighResolutionImages() *bool
	SetXDashScopeDataInspection(v string) *ChatWithDesensitizeSSERequest
	GetXDashScopeDataInspection() *string
}

type ChatWithDesensitizeSSERequest struct {
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
	Messages []interface{} `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// Output data modality; only applicable to the Qwen-Omni model.
	//
	// example:
	//
	// ["text","audio"]
	ModalitiesList []*string `json:"ModalitiesList,omitempty" xml:"ModalitiesList,omitempty" type:"Repeated"`
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
	SearchOptions map[string]*string `json:"SearchOptions,omitempty" xml:"SearchOptions,omitempty"`
	// Random seed. Used to ensure the reproducibility of results under the same input and parameters. Valid values: [0, 2^31−1].
	//
	// example:
	//
	// 1
	Seed *int32 `json:"Seed,omitempty" xml:"Seed,omitempty"`
	// Stop sequences.
	Stop []*string `json:"Stop,omitempty" xml:"Stop,omitempty" type:"Repeated"`
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

func (s ChatWithDesensitizeSSERequest) String() string {
	return dara.Prettify(s)
}

func (s ChatWithDesensitizeSSERequest) GoString() string {
	return s.String()
}

func (s *ChatWithDesensitizeSSERequest) GetAudioJson() *string {
	return s.AudioJson
}

func (s *ChatWithDesensitizeSSERequest) GetDesensitizationRule() *string {
	return s.DesensitizationRule
}

func (s *ChatWithDesensitizeSSERequest) GetDimensions() *int32 {
	return s.Dimensions
}

func (s *ChatWithDesensitizeSSERequest) GetEnableCodeInterpreter() *bool {
	return s.EnableCodeInterpreter
}

func (s *ChatWithDesensitizeSSERequest) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *ChatWithDesensitizeSSERequest) GetEnableThinking() *bool {
	return s.EnableThinking
}

func (s *ChatWithDesensitizeSSERequest) GetIncludeUsage() *bool {
	return s.IncludeUsage
}

func (s *ChatWithDesensitizeSSERequest) GetInput() *string {
	return s.Input
}

func (s *ChatWithDesensitizeSSERequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ChatWithDesensitizeSSERequest) GetLogprobs() *bool {
	return s.Logprobs
}

func (s *ChatWithDesensitizeSSERequest) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *ChatWithDesensitizeSSERequest) GetMessages() []interface{} {
	return s.Messages
}

func (s *ChatWithDesensitizeSSERequest) GetModalitiesList() []*string {
	return s.ModalitiesList
}

func (s *ChatWithDesensitizeSSERequest) GetModel() *string {
	return s.Model
}

func (s *ChatWithDesensitizeSSERequest) GetNeedDesensitization() *bool {
	return s.NeedDesensitization
}

func (s *ChatWithDesensitizeSSERequest) GetParameters() *string {
	return s.Parameters
}

func (s *ChatWithDesensitizeSSERequest) GetPresencePenalty() *string {
	return s.PresencePenalty
}

func (s *ChatWithDesensitizeSSERequest) GetResponseFormat() *string {
	return s.ResponseFormat
}

func (s *ChatWithDesensitizeSSERequest) GetSearchOptions() map[string]*string {
	return s.SearchOptions
}

func (s *ChatWithDesensitizeSSERequest) GetSeed() *int32 {
	return s.Seed
}

func (s *ChatWithDesensitizeSSERequest) GetStop() []*string {
	return s.Stop
}

func (s *ChatWithDesensitizeSSERequest) GetStream() *bool {
	return s.Stream
}

func (s *ChatWithDesensitizeSSERequest) GetTemperature() *string {
	return s.Temperature
}

func (s *ChatWithDesensitizeSSERequest) GetThinkingBudget() *int32 {
	return s.ThinkingBudget
}

func (s *ChatWithDesensitizeSSERequest) GetTopK() *int32 {
	return s.TopK
}

func (s *ChatWithDesensitizeSSERequest) GetTopLogprobs() *int32 {
	return s.TopLogprobs
}

func (s *ChatWithDesensitizeSSERequest) GetTopP() *string {
	return s.TopP
}

func (s *ChatWithDesensitizeSSERequest) GetVlHighResolutionImages() *bool {
	return s.VlHighResolutionImages
}

func (s *ChatWithDesensitizeSSERequest) GetXDashScopeDataInspection() *string {
	return s.XDashScopeDataInspection
}

func (s *ChatWithDesensitizeSSERequest) SetAudioJson(v string) *ChatWithDesensitizeSSERequest {
	s.AudioJson = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetDesensitizationRule(v string) *ChatWithDesensitizeSSERequest {
	s.DesensitizationRule = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetDimensions(v int32) *ChatWithDesensitizeSSERequest {
	s.Dimensions = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetEnableCodeInterpreter(v bool) *ChatWithDesensitizeSSERequest {
	s.EnableCodeInterpreter = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetEnableSearch(v bool) *ChatWithDesensitizeSSERequest {
	s.EnableSearch = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetEnableThinking(v bool) *ChatWithDesensitizeSSERequest {
	s.EnableThinking = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetIncludeUsage(v bool) *ChatWithDesensitizeSSERequest {
	s.IncludeUsage = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetInput(v string) *ChatWithDesensitizeSSERequest {
	s.Input = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetInstanceId(v int64) *ChatWithDesensitizeSSERequest {
	s.InstanceId = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetLogprobs(v bool) *ChatWithDesensitizeSSERequest {
	s.Logprobs = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetMaxTokens(v int32) *ChatWithDesensitizeSSERequest {
	s.MaxTokens = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetMessages(v []interface{}) *ChatWithDesensitizeSSERequest {
	s.Messages = v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetModalitiesList(v []*string) *ChatWithDesensitizeSSERequest {
	s.ModalitiesList = v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetModel(v string) *ChatWithDesensitizeSSERequest {
	s.Model = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetNeedDesensitization(v bool) *ChatWithDesensitizeSSERequest {
	s.NeedDesensitization = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetParameters(v string) *ChatWithDesensitizeSSERequest {
	s.Parameters = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetPresencePenalty(v string) *ChatWithDesensitizeSSERequest {
	s.PresencePenalty = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetResponseFormat(v string) *ChatWithDesensitizeSSERequest {
	s.ResponseFormat = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetSearchOptions(v map[string]*string) *ChatWithDesensitizeSSERequest {
	s.SearchOptions = v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetSeed(v int32) *ChatWithDesensitizeSSERequest {
	s.Seed = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetStop(v []*string) *ChatWithDesensitizeSSERequest {
	s.Stop = v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetStream(v bool) *ChatWithDesensitizeSSERequest {
	s.Stream = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetTemperature(v string) *ChatWithDesensitizeSSERequest {
	s.Temperature = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetThinkingBudget(v int32) *ChatWithDesensitizeSSERequest {
	s.ThinkingBudget = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetTopK(v int32) *ChatWithDesensitizeSSERequest {
	s.TopK = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetTopLogprobs(v int32) *ChatWithDesensitizeSSERequest {
	s.TopLogprobs = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetTopP(v string) *ChatWithDesensitizeSSERequest {
	s.TopP = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetVlHighResolutionImages(v bool) *ChatWithDesensitizeSSERequest {
	s.VlHighResolutionImages = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) SetXDashScopeDataInspection(v string) *ChatWithDesensitizeSSERequest {
	s.XDashScopeDataInspection = &v
	return s
}

func (s *ChatWithDesensitizeSSERequest) Validate() error {
	return dara.Validate(s)
}
