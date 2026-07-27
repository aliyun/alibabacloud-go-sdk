// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueries interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v int64) *Queries
	GetEnd() *int64
	SetExpr(v string) *Queries
	GetExpr() *string
	SetLabelFilters(v []*LabelFilters) *Queries
	GetLabelFilters() []*LabelFilters
	SetMetric(v string) *Queries
	GetMetric() *string
	SetMetricSet(v string) *Queries
	GetMetricSet() *string
	SetName(v string) *Queries
	GetName() *string
	SetStart(v int64) *Queries
	GetStart() *int64
	SetTimeUnit(v string) *Queries
	GetTimeUnit() *string
	SetWindow(v int64) *Queries
	GetWindow() *int64
}

type Queries struct {
	End          *int64          `json:"end,omitempty" xml:"end,omitempty"`
	Expr         *string         `json:"expr,omitempty" xml:"expr,omitempty"`
	LabelFilters []*LabelFilters `json:"labelFilters,omitempty" xml:"labelFilters,omitempty" type:"Repeated"`
	Metric       *string         `json:"metric,omitempty" xml:"metric,omitempty"`
	MetricSet    *string         `json:"metricSet,omitempty" xml:"metricSet,omitempty"`
	Name         *string         `json:"name,omitempty" xml:"name,omitempty"`
	Start        *int64          `json:"start,omitempty" xml:"start,omitempty"`
	TimeUnit     *string         `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	Window       *int64          `json:"window,omitempty" xml:"window,omitempty"`
}

func (s Queries) String() string {
	return dara.Prettify(s)
}

func (s Queries) GoString() string {
	return s.String()
}

func (s *Queries) GetEnd() *int64 {
	return s.End
}

func (s *Queries) GetExpr() *string {
	return s.Expr
}

func (s *Queries) GetLabelFilters() []*LabelFilters {
	return s.LabelFilters
}

func (s *Queries) GetMetric() *string {
	return s.Metric
}

func (s *Queries) GetMetricSet() *string {
	return s.MetricSet
}

func (s *Queries) GetName() *string {
	return s.Name
}

func (s *Queries) GetStart() *int64 {
	return s.Start
}

func (s *Queries) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *Queries) GetWindow() *int64 {
	return s.Window
}

func (s *Queries) SetEnd(v int64) *Queries {
	s.End = &v
	return s
}

func (s *Queries) SetExpr(v string) *Queries {
	s.Expr = &v
	return s
}

func (s *Queries) SetLabelFilters(v []*LabelFilters) *Queries {
	s.LabelFilters = v
	return s
}

func (s *Queries) SetMetric(v string) *Queries {
	s.Metric = &v
	return s
}

func (s *Queries) SetMetricSet(v string) *Queries {
	s.MetricSet = &v
	return s
}

func (s *Queries) SetName(v string) *Queries {
	s.Name = &v
	return s
}

func (s *Queries) SetStart(v int64) *Queries {
	s.Start = &v
	return s
}

func (s *Queries) SetTimeUnit(v string) *Queries {
	s.TimeUnit = &v
	return s
}

func (s *Queries) SetWindow(v int64) *Queries {
	s.Window = &v
	return s
}

func (s *Queries) Validate() error {
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
