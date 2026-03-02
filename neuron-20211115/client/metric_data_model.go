// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricData interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *MetricData
	GetData() *string
	SetMeasures(v string) *MetricData
	GetMeasures() *string
	SetMetric(v string) *MetricData
	GetMetric() *string
	SetTime(v int64) *MetricData
	GetTime() *int64
	SetType(v string) *MetricData
	GetType() *string
	SetTypeValue(v string) *MetricData
	GetTypeValue() *string
}

type MetricData struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// youngGcCount
	Measures *string `json:"measures,omitempty" xml:"measures,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// appstat.vm
	Metric *string `json:"metric,omitempty" xml:"metric,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1654777095632
	Time *int64 `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// rootIp
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 192.168.1.1
	TypeValue *string `json:"typeValue,omitempty" xml:"typeValue,omitempty"`
}

func (s MetricData) String() string {
	return dara.Prettify(s)
}

func (s MetricData) GoString() string {
	return s.String()
}

func (s *MetricData) GetData() *string {
	return s.Data
}

func (s *MetricData) GetMeasures() *string {
	return s.Measures
}

func (s *MetricData) GetMetric() *string {
	return s.Metric
}

func (s *MetricData) GetTime() *int64 {
	return s.Time
}

func (s *MetricData) GetType() *string {
	return s.Type
}

func (s *MetricData) GetTypeValue() *string {
	return s.TypeValue
}

func (s *MetricData) SetData(v string) *MetricData {
	s.Data = &v
	return s
}

func (s *MetricData) SetMeasures(v string) *MetricData {
	s.Measures = &v
	return s
}

func (s *MetricData) SetMetric(v string) *MetricData {
	s.Metric = &v
	return s
}

func (s *MetricData) SetTime(v int64) *MetricData {
	s.Time = &v
	return s
}

func (s *MetricData) SetType(v string) *MetricData {
	s.Type = &v
	return s
}

func (s *MetricData) SetTypeValue(v string) *MetricData {
	s.TypeValue = &v
	return s
}

func (s *MetricData) Validate() error {
	return dara.Validate(s)
}
