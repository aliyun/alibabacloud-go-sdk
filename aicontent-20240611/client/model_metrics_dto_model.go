// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelMetricsDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAvgResponseTime(v float64) *ModelMetricsDTO
	GetAvgResponseTime() *float64
	SetInputTokens(v int64) *ModelMetricsDTO
	GetInputTokens() *int64
	SetOutputTokens(v int64) *ModelMetricsDTO
	GetOutputTokens() *int64
	SetSuccessRate(v float64) *ModelMetricsDTO
	GetSuccessRate() *float64
	SetTotalCalls(v int64) *ModelMetricsDTO
	GetTotalCalls() *int64
}

type ModelMetricsDTO struct {
	// example:
	//
	// 200.5
	AvgResponseTime *float64 `json:"avgResponseTime,omitempty" xml:"avgResponseTime,omitempty"`
	// example:
	//
	// 500000
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 300000
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 99.5
	SuccessRate *float64 `json:"successRate,omitempty" xml:"successRate,omitempty"`
	// example:
	//
	// 1000
	TotalCalls *int64 `json:"totalCalls,omitempty" xml:"totalCalls,omitempty"`
}

func (s ModelMetricsDTO) String() string {
	return dara.Prettify(s)
}

func (s ModelMetricsDTO) GoString() string {
	return s.String()
}

func (s *ModelMetricsDTO) GetAvgResponseTime() *float64 {
	return s.AvgResponseTime
}

func (s *ModelMetricsDTO) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *ModelMetricsDTO) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *ModelMetricsDTO) GetSuccessRate() *float64 {
	return s.SuccessRate
}

func (s *ModelMetricsDTO) GetTotalCalls() *int64 {
	return s.TotalCalls
}

func (s *ModelMetricsDTO) SetAvgResponseTime(v float64) *ModelMetricsDTO {
	s.AvgResponseTime = &v
	return s
}

func (s *ModelMetricsDTO) SetInputTokens(v int64) *ModelMetricsDTO {
	s.InputTokens = &v
	return s
}

func (s *ModelMetricsDTO) SetOutputTokens(v int64) *ModelMetricsDTO {
	s.OutputTokens = &v
	return s
}

func (s *ModelMetricsDTO) SetSuccessRate(v float64) *ModelMetricsDTO {
	s.SuccessRate = &v
	return s
}

func (s *ModelMetricsDTO) SetTotalCalls(v int64) *ModelMetricsDTO {
	s.TotalCalls = &v
	return s
}

func (s *ModelMetricsDTO) Validate() error {
	return dara.Validate(s)
}
