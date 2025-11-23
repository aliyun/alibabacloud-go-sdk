// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithDesensitizeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesensitizationRule(v string) *ChatWithDesensitizeRequest
	GetDesensitizationRule() *string
	SetEnableThinking(v bool) *ChatWithDesensitizeRequest
	GetEnableThinking() *bool
	SetInstanceId(v int64) *ChatWithDesensitizeRequest
	GetInstanceId() *int64
	SetMaxTokens(v int32) *ChatWithDesensitizeRequest
	GetMaxTokens() *int32
	SetMessages(v []map[string]interface{}) *ChatWithDesensitizeRequest
	GetMessages() []map[string]interface{}
	SetModel(v string) *ChatWithDesensitizeRequest
	GetModel() *string
	SetNeedDesensitization(v bool) *ChatWithDesensitizeRequest
	GetNeedDesensitization() *bool
	SetPresencePenalty(v float32) *ChatWithDesensitizeRequest
	GetPresencePenalty() *float32
	SetResponseFormat(v string) *ChatWithDesensitizeRequest
	GetResponseFormat() *string
	SetSeed(v int32) *ChatWithDesensitizeRequest
	GetSeed() *int32
	SetStop(v []*string) *ChatWithDesensitizeRequest
	GetStop() []*string
	SetTemperature(v float32) *ChatWithDesensitizeRequest
	GetTemperature() *float32
	SetThinkingBudget(v int32) *ChatWithDesensitizeRequest
	GetThinkingBudget() *int32
	SetTopK(v int32) *ChatWithDesensitizeRequest
	GetTopK() *int32
	SetTopLogprobs(v int32) *ChatWithDesensitizeRequest
	GetTopLogprobs() *int32
	SetTopP(v float32) *ChatWithDesensitizeRequest
	GetTopP() *float32
}

type ChatWithDesensitizeRequest struct {
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
	Messages []map[string]interface{} `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
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
	Seed *int32    `json:"Seed,omitempty" xml:"Seed,omitempty"`
	Stop []*string `json:"Stop,omitempty" xml:"Stop,omitempty" type:"Repeated"`
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

func (s ChatWithDesensitizeRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatWithDesensitizeRequest) GoString() string {
	return s.String()
}

func (s *ChatWithDesensitizeRequest) GetDesensitizationRule() *string {
	return s.DesensitizationRule
}

func (s *ChatWithDesensitizeRequest) GetEnableThinking() *bool {
	return s.EnableThinking
}

func (s *ChatWithDesensitizeRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ChatWithDesensitizeRequest) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *ChatWithDesensitizeRequest) GetMessages() []map[string]interface{} {
	return s.Messages
}

func (s *ChatWithDesensitizeRequest) GetModel() *string {
	return s.Model
}

func (s *ChatWithDesensitizeRequest) GetNeedDesensitization() *bool {
	return s.NeedDesensitization
}

func (s *ChatWithDesensitizeRequest) GetPresencePenalty() *float32 {
	return s.PresencePenalty
}

func (s *ChatWithDesensitizeRequest) GetResponseFormat() *string {
	return s.ResponseFormat
}

func (s *ChatWithDesensitizeRequest) GetSeed() *int32 {
	return s.Seed
}

func (s *ChatWithDesensitizeRequest) GetStop() []*string {
	return s.Stop
}

func (s *ChatWithDesensitizeRequest) GetTemperature() *float32 {
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

func (s *ChatWithDesensitizeRequest) GetTopP() *float32 {
	return s.TopP
}

func (s *ChatWithDesensitizeRequest) SetDesensitizationRule(v string) *ChatWithDesensitizeRequest {
	s.DesensitizationRule = &v
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

func (s *ChatWithDesensitizeRequest) SetMaxTokens(v int32) *ChatWithDesensitizeRequest {
	s.MaxTokens = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetMessages(v []map[string]interface{}) *ChatWithDesensitizeRequest {
	s.Messages = v
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

func (s *ChatWithDesensitizeRequest) SetPresencePenalty(v float32) *ChatWithDesensitizeRequest {
	s.PresencePenalty = &v
	return s
}

func (s *ChatWithDesensitizeRequest) SetResponseFormat(v string) *ChatWithDesensitizeRequest {
	s.ResponseFormat = &v
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

func (s *ChatWithDesensitizeRequest) SetTemperature(v float32) *ChatWithDesensitizeRequest {
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

func (s *ChatWithDesensitizeRequest) SetTopP(v float32) *ChatWithDesensitizeRequest {
	s.TopP = &v
	return s
}

func (s *ChatWithDesensitizeRequest) Validate() error {
	return dara.Validate(s)
}
