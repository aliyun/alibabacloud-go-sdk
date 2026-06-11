// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketModelFeature interface {
	dara.Model
	String() string
	GoString() string
	SetEnableMultiModal(v bool) *HiMarketModelFeature
	GetEnableMultiModal() *bool
	SetEnableThinking(v bool) *HiMarketModelFeature
	GetEnableThinking() *bool
	SetMaxTokens(v int32) *HiMarketModelFeature
	GetMaxTokens() *int32
	SetModel(v string) *HiMarketModelFeature
	GetModel() *string
	SetStreaming(v bool) *HiMarketModelFeature
	GetStreaming() *bool
	SetTemperature(v float32) *HiMarketModelFeature
	GetTemperature() *float32
	SetWebSearch(v bool) *HiMarketModelFeature
	GetWebSearch() *bool
}

type HiMarketModelFeature struct {
	// Indicates whether to enable multi-modal capabilities. If set to `true`, the model can process requests that include multiple data types, such as text and images.
	EnableMultiModal *bool `json:"enableMultiModal,omitempty" xml:"enableMultiModal,omitempty"`
	// Indicates whether to include the model\\"s reasoning process in the response. If set to `true`, the output may contain intermediate steps that show how the model arrived at a conclusion.
	EnableThinking *bool `json:"enableThinking,omitempty" xml:"enableThinking,omitempty"`
	// The maximum number of tokens to generate in the response. A token is a unit of text processed by the model.
	MaxTokens *int32 `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	// The identifier of the model to use for inference.
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// Indicates whether to deliver the response as a continuous stream. If set to `true`, results are sent incrementally.
	Streaming *bool `json:"streaming,omitempty" xml:"streaming,omitempty"`
	// Controls the randomness of the output. Valid values range from 0 to 1. Higher values, such as 0.8, make the output more random. Lower values, such as 0.2, make the output more deterministic.
	Temperature *float32 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// Indicates whether the model can search the web to provide more up-to-date responses.
	WebSearch *bool `json:"webSearch,omitempty" xml:"webSearch,omitempty"`
}

func (s HiMarketModelFeature) String() string {
	return dara.Prettify(s)
}

func (s HiMarketModelFeature) GoString() string {
	return s.String()
}

func (s *HiMarketModelFeature) GetEnableMultiModal() *bool {
	return s.EnableMultiModal
}

func (s *HiMarketModelFeature) GetEnableThinking() *bool {
	return s.EnableThinking
}

func (s *HiMarketModelFeature) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *HiMarketModelFeature) GetModel() *string {
	return s.Model
}

func (s *HiMarketModelFeature) GetStreaming() *bool {
	return s.Streaming
}

func (s *HiMarketModelFeature) GetTemperature() *float32 {
	return s.Temperature
}

func (s *HiMarketModelFeature) GetWebSearch() *bool {
	return s.WebSearch
}

func (s *HiMarketModelFeature) SetEnableMultiModal(v bool) *HiMarketModelFeature {
	s.EnableMultiModal = &v
	return s
}

func (s *HiMarketModelFeature) SetEnableThinking(v bool) *HiMarketModelFeature {
	s.EnableThinking = &v
	return s
}

func (s *HiMarketModelFeature) SetMaxTokens(v int32) *HiMarketModelFeature {
	s.MaxTokens = &v
	return s
}

func (s *HiMarketModelFeature) SetModel(v string) *HiMarketModelFeature {
	s.Model = &v
	return s
}

func (s *HiMarketModelFeature) SetStreaming(v bool) *HiMarketModelFeature {
	s.Streaming = &v
	return s
}

func (s *HiMarketModelFeature) SetTemperature(v float32) *HiMarketModelFeature {
	s.Temperature = &v
	return s
}

func (s *HiMarketModelFeature) SetWebSearch(v bool) *HiMarketModelFeature {
	s.WebSearch = &v
	return s
}

func (s *HiMarketModelFeature) Validate() error {
	return dara.Validate(s)
}
