// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorMetricResult interface {
	dara.Model
	String() string
	GoString() string
	SetMeasureData(v []*MonitorMetricMeasureData) *MonitorMetricResult
	GetMeasureData() []*MonitorMetricMeasureData
	SetRequestId(v string) *MonitorMetricResult
	GetRequestId() *string
}

type MonitorMetricResult struct {
	MeasureData []*MonitorMetricMeasureData `json:"measureData,omitempty" xml:"measureData,omitempty" type:"Repeated"`
	RequestId   *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s MonitorMetricResult) String() string {
	return dara.Prettify(s)
}

func (s MonitorMetricResult) GoString() string {
	return s.String()
}

func (s *MonitorMetricResult) GetMeasureData() []*MonitorMetricMeasureData {
	return s.MeasureData
}

func (s *MonitorMetricResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MonitorMetricResult) SetMeasureData(v []*MonitorMetricMeasureData) *MonitorMetricResult {
	s.MeasureData = v
	return s
}

func (s *MonitorMetricResult) SetRequestId(v string) *MonitorMetricResult {
	s.RequestId = &v
	return s
}

func (s *MonitorMetricResult) Validate() error {
	if s.MeasureData != nil {
		for _, item := range s.MeasureData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
