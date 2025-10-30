// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPopTrafficStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListPopTrafficStatisticsResponseBody
	GetRequestId() *string
	SetTrafficData(v []*ListPopTrafficStatisticsResponseBodyTrafficData) *ListPopTrafficStatisticsResponseBody
	GetTrafficData() []*ListPopTrafficStatisticsResponseBodyTrafficData
}

type ListPopTrafficStatisticsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// EFE7EBB2-449D-5BBB-B381-CA7839BC1649
	RequestId   *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficData []*ListPopTrafficStatisticsResponseBodyTrafficData `json:"TrafficData,omitempty" xml:"TrafficData,omitempty" type:"Repeated"`
}

func (s ListPopTrafficStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPopTrafficStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPopTrafficStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPopTrafficStatisticsResponseBody) GetTrafficData() []*ListPopTrafficStatisticsResponseBodyTrafficData {
	return s.TrafficData
}

func (s *ListPopTrafficStatisticsResponseBody) SetRequestId(v string) *ListPopTrafficStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPopTrafficStatisticsResponseBody) SetTrafficData(v []*ListPopTrafficStatisticsResponseBodyTrafficData) *ListPopTrafficStatisticsResponseBody {
	s.TrafficData = v
	return s
}

func (s *ListPopTrafficStatisticsResponseBody) Validate() error {
	if s.TrafficData != nil {
		for _, item := range s.TrafficData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPopTrafficStatisticsResponseBodyTrafficData struct {
	Datapoints []*ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints `json:"Datapoints,omitempty" xml:"Datapoints,omitempty" type:"Repeated"`
	// example:
	//
	// InternetTx
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
}

func (s ListPopTrafficStatisticsResponseBodyTrafficData) String() string {
	return dara.Prettify(s)
}

func (s ListPopTrafficStatisticsResponseBodyTrafficData) GoString() string {
	return s.String()
}

func (s *ListPopTrafficStatisticsResponseBodyTrafficData) GetDatapoints() []*ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints {
	return s.Datapoints
}

func (s *ListPopTrafficStatisticsResponseBodyTrafficData) GetMetricName() *string {
	return s.MetricName
}

func (s *ListPopTrafficStatisticsResponseBodyTrafficData) SetDatapoints(v []*ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints) *ListPopTrafficStatisticsResponseBodyTrafficData {
	s.Datapoints = v
	return s
}

func (s *ListPopTrafficStatisticsResponseBodyTrafficData) SetMetricName(v string) *ListPopTrafficStatisticsResponseBodyTrafficData {
	s.MetricName = &v
	return s
}

func (s *ListPopTrafficStatisticsResponseBodyTrafficData) Validate() error {
	if s.Datapoints != nil {
		for _, item := range s.Datapoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints struct {
	// example:
	//
	// 15325
	Average *float64 `json:"Average,omitempty" xml:"Average,omitempty"`
	// example:
	//
	// 2023-12-06 15:29:00
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
}

func (s ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints) String() string {
	return dara.Prettify(s)
}

func (s ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints) GoString() string {
	return s.String()
}

func (s *ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints) GetAverage() *float64 {
	return s.Average
}

func (s *ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints) GetDateTime() *string {
	return s.DateTime
}

func (s *ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints) SetAverage(v float64) *ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints {
	s.Average = &v
	return s
}

func (s *ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints) SetDateTime(v string) *ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints {
	s.DateTime = &v
	return s
}

func (s *ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints) Validate() error {
	return dara.Validate(s)
}
