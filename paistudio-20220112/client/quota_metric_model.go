// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaMetric interface {
	dara.Model
	String() string
	GoString() string
	SetGPUType(v string) *QuotaMetric
	GetGPUType() *string
	SetMetrics(v []*Metric) *QuotaMetric
	GetMetrics() []*Metric
}

type QuotaMetric struct {
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// 23000
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
}

func (s QuotaMetric) String() string {
	return dara.Prettify(s)
}

func (s QuotaMetric) GoString() string {
	return s.String()
}

func (s *QuotaMetric) GetGPUType() *string {
	return s.GPUType
}

func (s *QuotaMetric) GetMetrics() []*Metric {
	return s.Metrics
}

func (s *QuotaMetric) SetGPUType(v string) *QuotaMetric {
	s.GPUType = &v
	return s
}

func (s *QuotaMetric) SetMetrics(v []*Metric) *QuotaMetric {
	s.Metrics = v
	return s
}

func (s *QuotaMetric) Validate() error {
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
