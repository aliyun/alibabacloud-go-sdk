// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApmThresholdConfig interface {
	dara.Model
	String() string
	GoString() string
	SetMax(v float64) *ApmThresholdConfig
	GetMax() *float64
	SetMin(v float64) *ApmThresholdConfig
	GetMin() *float64
	SetSeverity(v string) *ApmThresholdConfig
	GetSeverity() *string
	SetThreshold(v float32) *ApmThresholdConfig
	GetThreshold() *float32
}

type ApmThresholdConfig struct {
	Max *float64 `json:"max,omitempty" xml:"max,omitempty"`
	Min *float64 `json:"min,omitempty" xml:"min,omitempty"`
	// This parameter is required.
	Severity  *string  `json:"severity,omitempty" xml:"severity,omitempty"`
	Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s ApmThresholdConfig) String() string {
	return dara.Prettify(s)
}

func (s ApmThresholdConfig) GoString() string {
	return s.String()
}

func (s *ApmThresholdConfig) GetMax() *float64 {
	return s.Max
}

func (s *ApmThresholdConfig) GetMin() *float64 {
	return s.Min
}

func (s *ApmThresholdConfig) GetSeverity() *string {
	return s.Severity
}

func (s *ApmThresholdConfig) GetThreshold() *float32 {
	return s.Threshold
}

func (s *ApmThresholdConfig) SetMax(v float64) *ApmThresholdConfig {
	s.Max = &v
	return s
}

func (s *ApmThresholdConfig) SetMin(v float64) *ApmThresholdConfig {
	s.Min = &v
	return s
}

func (s *ApmThresholdConfig) SetSeverity(v string) *ApmThresholdConfig {
	s.Severity = &v
	return s
}

func (s *ApmThresholdConfig) SetThreshold(v float32) *ApmThresholdConfig {
	s.Threshold = &v
	return s
}

func (s *ApmThresholdConfig) Validate() error {
	return dara.Validate(s)
}
