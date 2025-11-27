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
