// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTrafficStatisticsResponseBody
	GetRequestId() *string
	SetTrafficStatistics(v []*ListTrafficStatisticsResponseBodyTrafficStatistics) *ListTrafficStatisticsResponseBody
	GetTrafficStatistics() []*ListTrafficStatisticsResponseBodyTrafficStatistics
}

type ListTrafficStatisticsResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficStatistics []*ListTrafficStatisticsResponseBodyTrafficStatistics `json:"TrafficStatistics,omitempty" xml:"TrafficStatistics,omitempty" type:"Repeated"`
}

func (s ListTrafficStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrafficStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrafficStatisticsResponseBody) GetTrafficStatistics() []*ListTrafficStatisticsResponseBodyTrafficStatistics {
	return s.TrafficStatistics
}

func (s *ListTrafficStatisticsResponseBody) SetRequestId(v string) *ListTrafficStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrafficStatisticsResponseBody) SetTrafficStatistics(v []*ListTrafficStatisticsResponseBodyTrafficStatistics) *ListTrafficStatisticsResponseBody {
	s.TrafficStatistics = v
	return s
}

func (s *ListTrafficStatisticsResponseBody) Validate() error {
	if s.TrafficStatistics != nil {
		for _, item := range s.TrafficStatistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTrafficStatisticsResponseBodyTrafficStatistics struct {
	TrafficStatisticData []*ListTrafficStatisticsResponseBodyTrafficStatisticsTrafficStatisticData `json:"TrafficStatisticData,omitempty" xml:"TrafficStatisticData,omitempty" type:"Repeated"`
	// example:
	//
	// all。
	TrafficStatisticTarget *string `json:"TrafficStatisticTarget,omitempty" xml:"TrafficStatisticTarget,omitempty"`
}

func (s ListTrafficStatisticsResponseBodyTrafficStatistics) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficStatisticsResponseBodyTrafficStatistics) GoString() string {
	return s.String()
}

func (s *ListTrafficStatisticsResponseBodyTrafficStatistics) GetTrafficStatisticData() []*ListTrafficStatisticsResponseBodyTrafficStatisticsTrafficStatisticData {
	return s.TrafficStatisticData
}

func (s *ListTrafficStatisticsResponseBodyTrafficStatistics) GetTrafficStatisticTarget() *string {
	return s.TrafficStatisticTarget
}

func (s *ListTrafficStatisticsResponseBodyTrafficStatistics) SetTrafficStatisticData(v []*ListTrafficStatisticsResponseBodyTrafficStatisticsTrafficStatisticData) *ListTrafficStatisticsResponseBodyTrafficStatistics {
	s.TrafficStatisticData = v
	return s
}

func (s *ListTrafficStatisticsResponseBodyTrafficStatistics) SetTrafficStatisticTarget(v string) *ListTrafficStatisticsResponseBodyTrafficStatistics {
	s.TrafficStatisticTarget = &v
	return s
}

func (s *ListTrafficStatisticsResponseBodyTrafficStatistics) Validate() error {
	if s.TrafficStatisticData != nil {
		for _, item := range s.TrafficStatisticData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTrafficStatisticsResponseBodyTrafficStatisticsTrafficStatisticData struct {
	// example:
	//
	// 20250815。
	TrafficStatisticTime *int64 `json:"TrafficStatisticTime,omitempty" xml:"TrafficStatisticTime,omitempty"`
	// example:
	//
	// 1.699814。
	TrafficStatisticValue *float64 `json:"TrafficStatisticValue,omitempty" xml:"TrafficStatisticValue,omitempty"`
}

func (s ListTrafficStatisticsResponseBodyTrafficStatisticsTrafficStatisticData) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficStatisticsResponseBodyTrafficStatisticsTrafficStatisticData) GoString() string {
	return s.String()
}

func (s *ListTrafficStatisticsResponseBodyTrafficStatisticsTrafficStatisticData) GetTrafficStatisticTime() *int64 {
	return s.TrafficStatisticTime
}

func (s *ListTrafficStatisticsResponseBodyTrafficStatisticsTrafficStatisticData) GetTrafficStatisticValue() *float64 {
	return s.TrafficStatisticValue
}

func (s *ListTrafficStatisticsResponseBodyTrafficStatisticsTrafficStatisticData) SetTrafficStatisticTime(v int64) *ListTrafficStatisticsResponseBodyTrafficStatisticsTrafficStatisticData {
	s.TrafficStatisticTime = &v
	return s
}

func (s *ListTrafficStatisticsResponseBodyTrafficStatisticsTrafficStatisticData) SetTrafficStatisticValue(v float64) *ListTrafficStatisticsResponseBodyTrafficStatisticsTrafficStatisticData {
	s.TrafficStatisticValue = &v
	return s
}

func (s *ListTrafficStatisticsResponseBodyTrafficStatisticsTrafficStatisticData) Validate() error {
	return dara.Validate(s)
}
