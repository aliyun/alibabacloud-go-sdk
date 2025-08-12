// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricStat interface {
	dara.Model
	String() string
	GoString() string
	SetAssociated(v map[string]*string) *MetricStat
	GetAssociated() map[string]*string
	SetDimensions(v []*Dimension) *MetricStat
	GetDimensions() []*Dimension
	SetLogTime(v int64) *MetricStat
	GetLogTime() *int64
	SetMeasurements(v map[string]interface{}) *MetricStat
	GetMeasurements() map[string]interface{}
	SetMetric(v string) *MetricStat
	GetMetric() *string
	SetNamespace(v string) *MetricStat
	GetNamespace() *string
	SetPeriod(v int32) *MetricStat
	GetPeriod() *int32
	SetTimestamp(v int64) *MetricStat
	GetTimestamp() *int64
}

type MetricStat struct {
	Associated   map[string]*string     `json:"Associated,omitempty" xml:"Associated,omitempty"`
	Dimensions   []*Dimension           `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	LogTime      *int64                 `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	Measurements map[string]interface{} `json:"Measurements,omitempty" xml:"Measurements,omitempty"`
	Metric       *string                `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Namespace    *string                `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Period       *int32                 `json:"Period,omitempty" xml:"Period,omitempty"`
	Timestamp    *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s MetricStat) String() string {
	return dara.Prettify(s)
}

func (s MetricStat) GoString() string {
	return s.String()
}

func (s *MetricStat) GetAssociated() map[string]*string {
	return s.Associated
}

func (s *MetricStat) GetDimensions() []*Dimension {
	return s.Dimensions
}

func (s *MetricStat) GetLogTime() *int64 {
	return s.LogTime
}

func (s *MetricStat) GetMeasurements() map[string]interface{} {
	return s.Measurements
}

func (s *MetricStat) GetMetric() *string {
	return s.Metric
}

func (s *MetricStat) GetNamespace() *string {
	return s.Namespace
}

func (s *MetricStat) GetPeriod() *int32 {
	return s.Period
}

func (s *MetricStat) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *MetricStat) SetAssociated(v map[string]*string) *MetricStat {
	s.Associated = v
	return s
}

func (s *MetricStat) SetDimensions(v []*Dimension) *MetricStat {
	s.Dimensions = v
	return s
}

func (s *MetricStat) SetLogTime(v int64) *MetricStat {
	s.LogTime = &v
	return s
}

func (s *MetricStat) SetMeasurements(v map[string]interface{}) *MetricStat {
	s.Measurements = v
	return s
}

func (s *MetricStat) SetMetric(v string) *MetricStat {
	s.Metric = &v
	return s
}

func (s *MetricStat) SetNamespace(v string) *MetricStat {
	s.Namespace = &v
	return s
}

func (s *MetricStat) SetPeriod(v int32) *MetricStat {
	s.Period = &v
	return s
}

func (s *MetricStat) SetTimestamp(v int64) *MetricStat {
	s.Timestamp = &v
	return s
}

func (s *MetricStat) Validate() error {
	return dara.Validate(s)
}
