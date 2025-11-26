// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAdviseSummary interface {
	dara.Model
	String() string
	GoString() string
	SetMemoryUtilizationRate(v *DoubleMetric) *AdviseSummary
	GetMemoryUtilizationRate() *DoubleMetric
	SetVcoreUtilizationRate(v *DoubleMetric) *AdviseSummary
	GetVcoreUtilizationRate() *DoubleMetric
}

type AdviseSummary struct {
	MemoryUtilizationRate *DoubleMetric `json:"MemoryUtilizationRate,omitempty" xml:"MemoryUtilizationRate,omitempty"`
	VcoreUtilizationRate  *DoubleMetric `json:"VcoreUtilizationRate,omitempty" xml:"VcoreUtilizationRate,omitempty"`
}

func (s AdviseSummary) String() string {
	return dara.Prettify(s)
}

func (s AdviseSummary) GoString() string {
	return s.String()
}

func (s *AdviseSummary) GetMemoryUtilizationRate() *DoubleMetric {
	return s.MemoryUtilizationRate
}

func (s *AdviseSummary) GetVcoreUtilizationRate() *DoubleMetric {
	return s.VcoreUtilizationRate
}

func (s *AdviseSummary) SetMemoryUtilizationRate(v *DoubleMetric) *AdviseSummary {
	s.MemoryUtilizationRate = v
	return s
}

func (s *AdviseSummary) SetVcoreUtilizationRate(v *DoubleMetric) *AdviseSummary {
	s.VcoreUtilizationRate = v
	return s
}

func (s *AdviseSummary) Validate() error {
	if s.MemoryUtilizationRate != nil {
		if err := s.MemoryUtilizationRate.Validate(); err != nil {
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
