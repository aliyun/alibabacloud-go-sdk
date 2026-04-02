// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApmThresholdConfig interface {
	dara.Model
	String() string
	GoString() string
	SetSeverity(v string) *ApmThresholdConfig
	GetSeverity() *string
	SetThreshold(v float32) *ApmThresholdConfig
	GetThreshold() *float32
}

type ApmThresholdConfig struct {
	// 告警等级
	//
	// This parameter is required.
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// 阈值
	//
	// This parameter is required.
	Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s ApmThresholdConfig) String() string {
	return dara.Prettify(s)
}

func (s ApmThresholdConfig) GoString() string {
	return s.String()
}

func (s *ApmThresholdConfig) GetSeverity() *string {
	return s.Severity
}

func (s *ApmThresholdConfig) GetThreshold() *float32 {
	return s.Threshold
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
