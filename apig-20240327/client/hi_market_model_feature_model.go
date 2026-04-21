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
	EnableMultiModal *bool    `json:"enableMultiModal,omitempty" xml:"enableMultiModal,omitempty"`
	EnableThinking   *bool    `json:"enableThinking,omitempty" xml:"enableThinking,omitempty"`
	MaxTokens        *int32   `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	Model            *string  `json:"model,omitempty" xml:"model,omitempty"`
	Streaming        *bool    `json:"streaming,omitempty" xml:"streaming,omitempty"`
	Temperature      *float32 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	WebSearch        *bool    `json:"webSearch,omitempty" xml:"webSearch,omitempty"`
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
