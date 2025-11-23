// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithDesensitizeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesensitizationRule(v string) *ChatWithDesensitizeShrinkRequest
	GetDesensitizationRule() *string
	SetEnableThinking(v bool) *ChatWithDesensitizeShrinkRequest
	GetEnableThinking() *bool
	SetInstanceId(v int64) *ChatWithDesensitizeShrinkRequest
	GetInstanceId() *int64
	SetMaxTokens(v int32) *ChatWithDesensitizeShrinkRequest
	GetMaxTokens() *int32
	SetMessagesShrink(v string) *ChatWithDesensitizeShrinkRequest
	GetMessagesShrink() *string
	SetModel(v string) *ChatWithDesensitizeShrinkRequest
	GetModel() *string
	SetNeedDesensitization(v bool) *ChatWithDesensitizeShrinkRequest
	GetNeedDesensitization() *bool
	SetPresencePenalty(v float32) *ChatWithDesensitizeShrinkRequest
	GetPresencePenalty() *float32
	SetResponseFormat(v string) *ChatWithDesensitizeShrinkRequest
	GetResponseFormat() *string
	SetSeed(v int32) *ChatWithDesensitizeShrinkRequest
	GetSeed() *int32
	SetStopShrink(v string) *ChatWithDesensitizeShrinkRequest
	GetStopShrink() *string
	SetTemperature(v float32) *ChatWithDesensitizeShrinkRequest
	GetTemperature() *float32
	SetThinkingBudget(v int32) *ChatWithDesensitizeShrinkRequest
	GetThinkingBudget() *int32
	SetTopK(v int32) *ChatWithDesensitizeShrinkRequest
	GetTopK() *int32
	SetTopLogprobs(v int32) *ChatWithDesensitizeShrinkRequest
	GetTopLogprobs() *int32
	SetTopP(v float32) *ChatWithDesensitizeShrinkRequest
	GetTopP() *float32
}

type ChatWithDesensitizeShrinkRequest struct {
	// example:
	//
	// UserInfo
	DesensitizationRule *string `json:"DesensitizationRule,omitempty" xml:"DesensitizationRule,omitempty"`
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
	// 256
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// This parameter is required.
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
	PresencePenalty *float32 `json:"PresencePenalty,omitempty" xml:"PresencePenalty,omitempty"`
	// example:
	//
	// text
	ResponseFormat *string `json:"ResponseFormat,omitempty" xml:"ResponseFormat,omitempty"`
	// example:
	//
	// 1
	Seed       *int32  `json:"Seed,omitempty" xml:"Seed,omitempty"`
	StopShrink *string `json:"Stop,omitempty" xml:"Stop,omitempty"`
	// example:
	//
	// 1
	Temperature *float32 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
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
	TopP *float32 `json:"TopP,omitempty" xml:"TopP,omitempty"`
}

func (s ChatWithDesensitizeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatWithDesensitizeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChatWithDesensitizeShrinkRequest) GetDesensitizationRule() *string {
	return s.DesensitizationRule
}

func (s *ChatWithDesensitizeShrinkRequest) GetEnableThinking() *bool {
	return s.EnableThinking
}

func (s *ChatWithDesensitizeShrinkRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ChatWithDesensitizeShrinkRequest) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *ChatWithDesensitizeShrinkRequest) GetMessagesShrink() *string {
	return s.MessagesShrink
}

func (s *ChatWithDesensitizeShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *ChatWithDesensitizeShrinkRequest) GetNeedDesensitization() *bool {
	return s.NeedDesensitization
}

func (s *ChatWithDesensitizeShrinkRequest) GetPresencePenalty() *float32 {
	return s.PresencePenalty
}

func (s *ChatWithDesensitizeShrinkRequest) GetResponseFormat() *string {
	return s.ResponseFormat
}

func (s *ChatWithDesensitizeShrinkRequest) GetSeed() *int32 {
	return s.Seed
}

func (s *ChatWithDesensitizeShrinkRequest) GetStopShrink() *string {
	return s.StopShrink
}

func (s *ChatWithDesensitizeShrinkRequest) GetTemperature() *float32 {
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

func (s *ChatWithDesensitizeShrinkRequest) GetTopP() *float32 {
	return s.TopP
}

func (s *ChatWithDesensitizeShrinkRequest) SetDesensitizationRule(v string) *ChatWithDesensitizeShrinkRequest {
	s.DesensitizationRule = &v
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

func (s *ChatWithDesensitizeShrinkRequest) SetMaxTokens(v int32) *ChatWithDesensitizeShrinkRequest {
	s.MaxTokens = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetMessagesShrink(v string) *ChatWithDesensitizeShrinkRequest {
	s.MessagesShrink = &v
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

func (s *ChatWithDesensitizeShrinkRequest) SetPresencePenalty(v float32) *ChatWithDesensitizeShrinkRequest {
	s.PresencePenalty = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) SetResponseFormat(v string) *ChatWithDesensitizeShrinkRequest {
	s.ResponseFormat = &v
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

func (s *ChatWithDesensitizeShrinkRequest) SetTemperature(v float32) *ChatWithDesensitizeShrinkRequest {
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

func (s *ChatWithDesensitizeShrinkRequest) SetTopP(v float32) *ChatWithDesensitizeShrinkRequest {
	s.TopP = &v
	return s
}

func (s *ChatWithDesensitizeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
