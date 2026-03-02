// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorMetricMeasureData interface {
	dara.Model
	String() string
	GoString() string
	SetMeasure(v string) *MonitorMetricMeasureData
	GetMeasure() *string
	SetMeasurePointData(v []*MonitorMetricMeasurePointData) *MonitorMetricMeasureData
	GetMeasurePointData() []*MonitorMetricMeasurePointData
}

type MonitorMetricMeasureData struct {
	// example:
	//
	// youngGcCount
	Measure          *string                          `json:"measure,omitempty" xml:"measure,omitempty"`
	MeasurePointData []*MonitorMetricMeasurePointData `json:"measurePointData,omitempty" xml:"measurePointData,omitempty" type:"Repeated"`
}

func (s MonitorMetricMeasureData) String() string {
	return dara.Prettify(s)
}

func (s MonitorMetricMeasureData) GoString() string {
	return s.String()
}

func (s *MonitorMetricMeasureData) GetMeasure() *string {
	return s.Measure
}

func (s *MonitorMetricMeasureData) GetMeasurePointData() []*MonitorMetricMeasurePointData {
	return s.MeasurePointData
}

func (s *MonitorMetricMeasureData) SetMeasure(v string) *MonitorMetricMeasureData {
	s.Measure = &v
	return s
}

func (s *MonitorMetricMeasureData) SetMeasurePointData(v []*MonitorMetricMeasurePointData) *MonitorMetricMeasureData {
	s.MeasurePointData = v
	return s
}

func (s *MonitorMetricMeasureData) Validate() error {
	if s.MeasurePointData != nil {
		for _, item := range s.MeasurePointData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
