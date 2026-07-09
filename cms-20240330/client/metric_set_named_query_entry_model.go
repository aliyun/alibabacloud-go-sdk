// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricSetNamedQueryEntry interface {
	dara.Model
	String() string
	GoString() string
	SetLabelFilters(v []*UmodelLabelFilter) *MetricSetNamedQueryEntry
	GetLabelFilters() []*UmodelLabelFilter
	SetMetric(v string) *MetricSetNamedQueryEntry
	GetMetric() *string
	SetMetricSet(v string) *MetricSetNamedQueryEntry
	GetMetricSet() *string
	SetName(v string) *MetricSetNamedQueryEntry
	GetName() *string
}

type MetricSetNamedQueryEntry struct {
	LabelFilters []*UmodelLabelFilter `json:"labelFilters,omitempty" xml:"labelFilters,omitempty" type:"Repeated"`
	Metric       *string              `json:"metric,omitempty" xml:"metric,omitempty"`
	MetricSet    *string              `json:"metricSet,omitempty" xml:"metricSet,omitempty"`
	Name         *string              `json:"name,omitempty" xml:"name,omitempty"`
}

func (s MetricSetNamedQueryEntry) String() string {
	return dara.Prettify(s)
}

func (s MetricSetNamedQueryEntry) GoString() string {
	return s.String()
}

func (s *MetricSetNamedQueryEntry) GetLabelFilters() []*UmodelLabelFilter {
	return s.LabelFilters
}

func (s *MetricSetNamedQueryEntry) GetMetric() *string {
	return s.Metric
}

func (s *MetricSetNamedQueryEntry) GetMetricSet() *string {
	return s.MetricSet
}

func (s *MetricSetNamedQueryEntry) GetName() *string {
	return s.Name
}

func (s *MetricSetNamedQueryEntry) SetLabelFilters(v []*UmodelLabelFilter) *MetricSetNamedQueryEntry {
	s.LabelFilters = v
	return s
}

func (s *MetricSetNamedQueryEntry) SetMetric(v string) *MetricSetNamedQueryEntry {
	s.Metric = &v
	return s
}

func (s *MetricSetNamedQueryEntry) SetMetricSet(v string) *MetricSetNamedQueryEntry {
	s.MetricSet = &v
	return s
}

func (s *MetricSetNamedQueryEntry) SetName(v string) *MetricSetNamedQueryEntry {
	s.Name = &v
	return s
}

func (s *MetricSetNamedQueryEntry) Validate() error {
	if s.LabelFilters != nil {
		for _, item := range s.LabelFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
