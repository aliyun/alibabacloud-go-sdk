// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceSummary interface {
	dara.Model
	String() string
	GoString() string
	SetInefficientTaskRate(v *DoubleMetric) *ResourceSummary
	GetInefficientTaskRate() *DoubleMetric
	SetMemoryUtilizationRate(v *DoubleMetric) *ResourceSummary
	GetMemoryUtilizationRate() *DoubleMetric
	SetOriginalTotalVcore(v *IntegerMetric) *ResourceSummary
	GetOriginalTotalVcore() *IntegerMetric
	SetVcoreUtilizationRate(v *DoubleMetric) *ResourceSummary
	GetVcoreUtilizationRate() *DoubleMetric
}

type ResourceSummary struct {
	InefficientTaskRate   *DoubleMetric  `json:"InefficientTaskRate,omitempty" xml:"InefficientTaskRate,omitempty"`
	MemoryUtilizationRate *DoubleMetric  `json:"MemoryUtilizationRate,omitempty" xml:"MemoryUtilizationRate,omitempty"`
	OriginalTotalVcore    *IntegerMetric `json:"OriginalTotalVcore,omitempty" xml:"OriginalTotalVcore,omitempty"`
	VcoreUtilizationRate  *DoubleMetric  `json:"VcoreUtilizationRate,omitempty" xml:"VcoreUtilizationRate,omitempty"`
}

func (s ResourceSummary) String() string {
	return dara.Prettify(s)
}

func (s ResourceSummary) GoString() string {
	return s.String()
}

func (s *ResourceSummary) GetInefficientTaskRate() *DoubleMetric {
	return s.InefficientTaskRate
}

func (s *ResourceSummary) GetMemoryUtilizationRate() *DoubleMetric {
	return s.MemoryUtilizationRate
}

func (s *ResourceSummary) GetOriginalTotalVcore() *IntegerMetric {
	return s.OriginalTotalVcore
}

func (s *ResourceSummary) GetVcoreUtilizationRate() *DoubleMetric {
	return s.VcoreUtilizationRate
}

func (s *ResourceSummary) SetInefficientTaskRate(v *DoubleMetric) *ResourceSummary {
	s.InefficientTaskRate = v
	return s
}

func (s *ResourceSummary) SetMemoryUtilizationRate(v *DoubleMetric) *ResourceSummary {
	s.MemoryUtilizationRate = v
	return s
}

func (s *ResourceSummary) SetOriginalTotalVcore(v *IntegerMetric) *ResourceSummary {
	s.OriginalTotalVcore = v
	return s
}

func (s *ResourceSummary) SetVcoreUtilizationRate(v *DoubleMetric) *ResourceSummary {
	s.VcoreUtilizationRate = v
	return s
}

func (s *ResourceSummary) Validate() error {
	if s.InefficientTaskRate != nil {
		if err := s.InefficientTaskRate.Validate(); err != nil {
			return err
		}
	}
	if s.MemoryUtilizationRate != nil {
		if err := s.MemoryUtilizationRate.Validate(); err != nil {
			return err
		}
	}
	if s.OriginalTotalVcore != nil {
		if err := s.OriginalTotalVcore.Validate(); err != nil {
			return err
		}
	}
	if s.VcoreUtilizationRate != nil {
		if err := s.VcoreUtilizationRate.Validate(); err != nil {
			return err
		}
	}
	return nil
}
