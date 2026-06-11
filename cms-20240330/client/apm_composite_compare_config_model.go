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
}

type ApmCompositeCompareConfig struct {
	// The aggregation method for metric data. For example, `AVG`, `SUM`, or `MAX`.
	//
	// This parameter is required.
	Aggregate *string `json:"aggregate,omitempty" xml:"aggregate,omitempty"`
	// The operator for comparing the aggregated metric data against the `threshold`. For example, `GREATER_THAN` or `LESS_THAN`.
	//
	// This parameter is required.
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The value to compare the aggregated metric data against. An alert is triggered when the metric data meets the condition defined by the `operator`.
	//
	// This parameter is required.
	Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
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

func (s *ApmCompositeCompareConfig) Validate() error {
	return dara.Validate(s)
}
