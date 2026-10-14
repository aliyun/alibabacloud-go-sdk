// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelParameters interface {
	dara.Model
	String() string
	GoString() string
	SetFrequencyPenalty(v float32) *ModelParameters
	GetFrequencyPenalty() *float32
	SetMaxTokens(v int64) *ModelParameters
	GetMaxTokens() *int64
	SetPresencePenalty(v float32) *ModelParameters
	GetPresencePenalty() *float32
	SetStopSequences(v []*string) *ModelParameters
	GetStopSequences() []*string
	SetTemperature(v float32) *ModelParameters
	GetTemperature() *float32
	SetTopK(v int32) *ModelParameters
	GetTopK() *int32
	SetTopP(v float32) *ModelParameters
	GetTopP() *float32
}

type ModelParameters struct {
	// The frequency penalty. Positive values reduce repetition. Valid values: -2.0 to 2.0.
	//
	// example:
	//
	// 0.2
	FrequencyPenalty *float32 `json:"frequencyPenalty,omitempty" xml:"frequencyPenalty,omitempty"`
	// The maximum number of output tokens.
	//
	// example:
	//
	// 1024
	MaxTokens *int64 `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	// The presence penalty. Positive values encourage new topics. Valid values: -2.0 to 2.0.
	//
	// example:
	//
	// 0.1
	PresencePenalty *float32 `json:"presencePenalty,omitempty" xml:"presencePenalty,omitempty"`
	// The stop sequences. Generation stops immediately when any of these sequences is encountered.
	//
	// example:
	//
	// ["END"]
	StopSequences []*string `json:"stopSequences,omitempty" xml:"stopSequences,omitempty" type:"Repeated"`
	// The temperature parameter that controls the randomness of the output. Valid values: 0 to 2.
	//
	// example:
	//
	// 0.7
	Temperature *float32 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// The top-k sampling parameter. Valid values: 1 to 100.
	//
	// example:
	//
	// 50
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
	// The top-p sampling parameter. Valid values: 0 to 1.
	//
	// example:
	//
	// 0.9
	TopP *float32 `json:"topP,omitempty" xml:"topP,omitempty"`
}

func (s ModelParameters) String() string {
	return dara.Prettify(s)
}

func (s ModelParameters) GoString() string {
	return s.String()
}

func (s *ModelParameters) GetFrequencyPenalty() *float32 {
	return s.FrequencyPenalty
}

func (s *ModelParameters) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *ModelParameters) GetPresencePenalty() *float32 {
	return s.PresencePenalty
}

func (s *ModelParameters) GetStopSequences() []*string {
	return s.StopSequences
}

func (s *ModelParameters) GetTemperature() *float32 {
	return s.Temperature
}

func (s *ModelParameters) GetTopK() *int32 {
	return s.TopK
}

func (s *ModelParameters) GetTopP() *float32 {
	return s.TopP
}

func (s *ModelParameters) SetFrequencyPenalty(v float32) *ModelParameters {
	s.FrequencyPenalty = &v
	return s
}

func (s *ModelParameters) SetMaxTokens(v int64) *ModelParameters {
	s.MaxTokens = &v
	return s
}

func (s *ModelParameters) SetPresencePenalty(v float32) *ModelParameters {
	s.PresencePenalty = &v
	return s
}

func (s *ModelParameters) SetStopSequences(v []*string) *ModelParameters {
	s.StopSequences = v
	return s
}

func (s *ModelParameters) SetTemperature(v float32) *ModelParameters {
	s.Temperature = &v
	return s
}

func (s *ModelParameters) SetTopK(v int32) *ModelParameters {
	s.TopK = &v
	return s
}

func (s *ModelParameters) SetTopP(v float32) *ModelParameters {
	s.TopP = &v
	return s
}

func (s *ModelParameters) Validate() error {
	return dara.Validate(s)
}
