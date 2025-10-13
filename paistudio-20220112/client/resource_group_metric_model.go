// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceGroupMetric interface {
	dara.Model
	String() string
	GoString() string
	SetGpuType(v string) *ResourceGroupMetric
	GetGpuType() *string
	SetMetrics(v []*Metric) *ResourceGroupMetric
	GetMetrics() []*Metric
	SetResourceGroupID(v string) *ResourceGroupMetric
	GetResourceGroupID() *string
}

type ResourceGroupMetric struct {
	GpuType *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	// example:
	//
	// 23000
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// rg17tmvwiokhzaxg
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
}

func (s ResourceGroupMetric) String() string {
	return dara.Prettify(s)
}

func (s ResourceGroupMetric) GoString() string {
	return s.String()
}

func (s *ResourceGroupMetric) GetGpuType() *string {
	return s.GpuType
}

func (s *ResourceGroupMetric) GetMetrics() []*Metric {
	return s.Metrics
}

func (s *ResourceGroupMetric) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *ResourceGroupMetric) SetGpuType(v string) *ResourceGroupMetric {
	s.GpuType = &v
	return s
}

func (s *ResourceGroupMetric) SetMetrics(v []*Metric) *ResourceGroupMetric {
	s.Metrics = v
	return s
}

func (s *ResourceGroupMetric) SetResourceGroupID(v string) *ResourceGroupMetric {
	s.ResourceGroupID = &v
	return s
}

func (s *ResourceGroupMetric) Validate() error {
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
