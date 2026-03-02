// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorMetricMeasurePointData interface {
	dara.Model
	String() string
	GoString() string
	SetMeasure(v string) *MonitorMetricMeasurePointData
	GetMeasure() *string
	SetTimeStamp(v string) *MonitorMetricMeasurePointData
	GetTimeStamp() *string
	SetValue(v string) *MonitorMetricMeasurePointData
	GetValue() *string
	SetValueWithUnit(v string) *MonitorMetricMeasurePointData
	GetValueWithUnit() *string
}

type MonitorMetricMeasurePointData struct {
	// example:
	//
	// youngGcCount
	Measure *string `json:"measure,omitempty" xml:"measure,omitempty"`
	// example:
	//
	// 08-22 17:25
	TimeStamp *string `json:"timeStamp,omitempty" xml:"timeStamp,omitempty"`
	// example:
	//
	// 1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// example:
	//
	// 1M
	ValueWithUnit *string `json:"valueWithUnit,omitempty" xml:"valueWithUnit,omitempty"`
}

func (s MonitorMetricMeasurePointData) String() string {
	return dara.Prettify(s)
}

func (s MonitorMetricMeasurePointData) GoString() string {
	return s.String()
}

func (s *MonitorMetricMeasurePointData) GetMeasure() *string {
	return s.Measure
}

func (s *MonitorMetricMeasurePointData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *MonitorMetricMeasurePointData) GetValue() *string {
	return s.Value
}

func (s *MonitorMetricMeasurePointData) GetValueWithUnit() *string {
	return s.ValueWithUnit
}

func (s *MonitorMetricMeasurePointData) SetMeasure(v string) *MonitorMetricMeasurePointData {
	s.Measure = &v
	return s
}

func (s *MonitorMetricMeasurePointData) SetTimeStamp(v string) *MonitorMetricMeasurePointData {
	s.TimeStamp = &v
	return s
}

func (s *MonitorMetricMeasurePointData) SetValue(v string) *MonitorMetricMeasurePointData {
	s.Value = &v
	return s
}

func (s *MonitorMetricMeasurePointData) SetValueWithUnit(v string) *MonitorMetricMeasurePointData {
	s.ValueWithUnit = &v
	return s
}

func (s *MonitorMetricMeasurePointData) Validate() error {
	return dara.Validate(s)
}
