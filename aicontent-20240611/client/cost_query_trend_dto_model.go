// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostQueryTrendDTO interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultMetric(v string) *CostQueryTrendDTO
	GetDefaultMetric() *string
	SetGranularity(v string) *CostQueryTrendDTO
	GetGranularity() *string
	SetMetrics(v []*MetricDefRespDTO) *CostQueryTrendDTO
	GetMetrics() []*MetricDefRespDTO
	SetPoints(v []*TrendPointDTO) *CostQueryTrendDTO
	GetPoints() []*TrendPointDTO
}

type CostQueryTrendDTO struct {
	// example:
	//
	// total_amount
	DefaultMetric *string `json:"defaultMetric,omitempty" xml:"defaultMetric,omitempty"`
	// example:
	//
	// hourly
	Granularity *string             `json:"granularity,omitempty" xml:"granularity,omitempty"`
	Metrics     []*MetricDefRespDTO `json:"metrics,omitempty" xml:"metrics,omitempty" type:"Repeated"`
	Points      []*TrendPointDTO    `json:"points,omitempty" xml:"points,omitempty" type:"Repeated"`
}

func (s CostQueryTrendDTO) String() string {
	return dara.Prettify(s)
}

func (s CostQueryTrendDTO) GoString() string {
	return s.String()
}

func (s *CostQueryTrendDTO) GetDefaultMetric() *string {
	return s.DefaultMetric
}

func (s *CostQueryTrendDTO) GetGranularity() *string {
	return s.Granularity
}

func (s *CostQueryTrendDTO) GetMetrics() []*MetricDefRespDTO {
	return s.Metrics
}

func (s *CostQueryTrendDTO) GetPoints() []*TrendPointDTO {
	return s.Points
}

func (s *CostQueryTrendDTO) SetDefaultMetric(v string) *CostQueryTrendDTO {
	s.DefaultMetric = &v
	return s
}

func (s *CostQueryTrendDTO) SetGranularity(v string) *CostQueryTrendDTO {
	s.Granularity = &v
	return s
}

func (s *CostQueryTrendDTO) SetMetrics(v []*MetricDefRespDTO) *CostQueryTrendDTO {
	s.Metrics = v
	return s
}

func (s *CostQueryTrendDTO) SetPoints(v []*TrendPointDTO) *CostQueryTrendDTO {
	s.Points = v
	return s
}

func (s *CostQueryTrendDTO) Validate() error {
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Points != nil {
		for _, item := range s.Points {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
