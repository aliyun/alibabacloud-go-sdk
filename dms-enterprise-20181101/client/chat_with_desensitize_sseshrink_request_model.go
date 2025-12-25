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
	// example:
	//
	// {}
	AudioJson *string `json:"AudioJson,omitempty" xml:"AudioJson,omitempty"`
	// example:
	//
	// UserInfo
	DesensitizationRule *string `json:"DesensitizationRule,omitempty" xml:"DesensitizationRule,omitempty"`
	// example:
	//
	// 256
	Dimensions *int32 `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// example:
	//
	// false
	EnableCodeInterpreter *bool `json:"EnableCodeInterpreter,omitempty" xml:"EnableCodeInterpreter,omitempty"`
	// example:
	//
	// false
	EnableSearch *bool `json:"EnableSearch,omitempty" xml:"EnableSearch,omitempty"`
	// example:
	//
	// true
	EnableThinking *bool `json:"EnableThinking,omitempty" xml:"EnableThinking,omitempty"`
	// example:
	//
	// true
	IncludeUsage *bool `json:"IncludeUsage,omitempty" xml:"IncludeUsage,omitempty"`
	// example:
	//
	// test
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123***
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// false
	Logprobs *bool `json:"Logprobs,omitempty" xml:"Logprobs,omitempty"`
	// example:
	//
	// 256
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
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
	// example:
	//
	// ["text","audio"]
	ModalitiesListShrink *string `json:"ModalitiesList,omitempty" xml:"ModalitiesList,omitempty"`
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// false
	NeedDesensitization *bool `json:"NeedDesensitization,omitempty" xml:"NeedDesensitization,omitempty"`
	// example:
	//
	// {}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// 0.0
	PresencePenalty *string `json:"PresencePenalty,omitempty" xml:"PresencePenalty,omitempty"`
	// example:
	//
	// text
	ResponseFormat *string `json:"ResponseFormat,omitempty" xml:"ResponseFormat,omitempty"`
	// example:
	//
	// {}
	SearchOptionsShrink *string `json:"SearchOptions,omitempty" xml:"SearchOptions,omitempty"`
	// example:
	//
	// 1
	Seed       *int32  `json:"Seed,omitempty" xml:"Seed,omitempty"`
	StopShrink *string `json:"Stop,omitempty" xml:"Stop,omitempty"`
	// example:
	//
	// 1-68f11da7e2b826dcc63c5877-hd
	Stream *bool `json:"Stream,omitempty" xml:"Stream,omitempty"`
	// example:
	//
	// 1
	Temperature *string `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// example:
	//
	// 256
	ThinkingBudget *int32 `json:"ThinkingBudget,omitempty" xml:"ThinkingBudget,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// example:
	//
	// 1
	TopLogprobs *int32 `json:"TopLogprobs,omitempty" xml:"TopLogprobs,omitempty"`
	// example:
	//
	// 0.5
	TopP *string `json:"TopP,omitempty" xml:"TopP,omitempty"`
	// example:
	//
	// false
	VlHighResolutionImages *bool `json:"VlHighResolutionImages,omitempty" xml:"VlHighResolutionImages,omitempty"`
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
