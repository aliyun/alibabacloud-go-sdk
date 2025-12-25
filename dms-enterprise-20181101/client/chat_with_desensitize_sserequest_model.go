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
	SearchOptions map[string]*string `json:"SearchOptions,omitempty" xml:"SearchOptions,omitempty"`
	// example:
	//
	// 1
	Seed *int32    `json:"Seed,omitempty" xml:"Seed,omitempty"`
	Stop []*string `json:"Stop,omitempty" xml:"Stop,omitempty" type:"Repeated"`
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
