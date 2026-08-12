// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApmCompositeCompareConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAggregate(v string) *ApmCompositeCompareConfig
	GetAggregate() *string
	SetOperator(v string) *ApmCompositeCompareConfig
	GetOperator() *string
	SetThreshold(v float32) *ApmCompositeCompareConfig
	GetThreshold() *float32
	SetYoyTimeUnit(v string) *ApmCompositeCompareConfig
	GetYoyTimeUnit() *string
	SetYoyTimeValue(v int32) *ApmCompositeCompareConfig
	GetYoyTimeValue() *int32
}

type ApmCompositeCompareConfig struct {
	// The aggregate functions used for aggregation.
	//
	// This parameter is required.
	Aggregate *string `json:"aggregate,omitempty" xml:"aggregate,omitempty"`
	// The comparison operator. GTE/LTE indicates greater than or equal to/less than or equal to. YOY_UP/YOY_DOWN indicates year-over-year increase/decrease, which requires yoyTimeUnit and yoyTimeValue to be specified.
	//
	// This parameter is required.
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The threshold.
	//
	// This parameter is required.
	Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// The year-over-year time unit. This parameter takes effect only when operator is set to YOY_UP or YOY_DOWN.
	YoyTimeUnit *string `json:"yoyTimeUnit,omitempty" xml:"yoyTimeUnit,omitempty"`
	// The year-over-year time value. This parameter takes effect only when operator is set to YOY_UP or YOY_DOWN.
	YoyTimeValue *int32 `json:"yoyTimeValue,omitempty" xml:"yoyTimeValue,omitempty"`
}

func (s ApmCompositeCompareConfig) String() string {
	return dara.Prettify(s)
}

func (s ApmCompositeCompareConfig) GoString() string {
	return s.String()
}

func (s *ApmCompositeCompareConfig) GetAggregate() *string {
	return s.Aggregate
}

func (s *ApmCompositeCompareConfig) GetOperator() *string {
	return s.Operator
}

func (s *ApmCompositeCompareConfig) GetThreshold() *float32 {
	return s.Threshold
}

func (s *ApmCompositeCompareConfig) GetYoyTimeUnit() *string {
	return s.YoyTimeUnit
}

func (s *ApmCompositeCompareConfig) GetYoyTimeValue() *int32 {
	return s.YoyTimeValue
}

func (s *ApmCompositeCompareConfig) SetAggregate(v string) *ApmCompositeCompareConfig {
	s.Aggregate = &v
	return s
}

func (s *ApmCompositeCompareConfig) SetOperator(v string) *ApmCompositeCompareConfig {
	s.Operator = &v
	return s
}

func (s *ApmCompositeCompareConfig) SetThreshold(v float32) *ApmCompositeCompareConfig {
	s.Threshold = &v
	return s
}

func (s *ApmCompositeCompareConfig) SetYoyTimeUnit(v string) *ApmCompositeCompareConfig {
	s.YoyTimeUnit = &v
	return s
}

func (s *ApmCompositeCompareConfig) SetYoyTimeValue(v int32) *ApmCompositeCompareConfig {
	s.YoyTimeValue = &v
	return s
}

func (s *ApmCompositeCompareConfig) Validate() error {
	return dara.Validate(s)
}
