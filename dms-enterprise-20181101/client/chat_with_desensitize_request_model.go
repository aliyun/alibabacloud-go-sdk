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
	Messages []interface{} `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// example:
	//
	// ["text","audio"]
	ModalitiesList []*string `json:"ModalitiesList,omitempty" xml:"ModalitiesList,omitempty" type:"Repeated"`
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
	SearchOptions map[string]*string `json:"SearchOptions,omitempty" xml:"SearchOptions,omitempty"`
	// example:
	//
	// 1
	Seed *int32    `json:"Seed,omitempty" xml:"Seed,omitempty"`
	Stop []*string `json:"Stop,omitempty" xml:"Stop,omitempty" type:"Repeated"`
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
