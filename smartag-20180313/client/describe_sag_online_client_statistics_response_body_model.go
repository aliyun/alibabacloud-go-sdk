// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagOnlineClientStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSagOnlineClientStatisticsResponseBody
	GetRequestId() *string
	SetSagStatistics(v *DescribeSagOnlineClientStatisticsResponseBodySagStatistics) *DescribeSagOnlineClientStatisticsResponseBody
	GetSagStatistics() *DescribeSagOnlineClientStatisticsResponseBodySagStatistics
}

type DescribeSagOnlineClientStatisticsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9EC839B6-0EA5-4F19-A4B7-A9E465D057AE
	RequestId     *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SagStatistics *DescribeSagOnlineClientStatisticsResponseBodySagStatistics `json:"SagStatistics,omitempty" xml:"SagStatistics,omitempty" type:"Struct"`
}

func (s DescribeSagOnlineClientStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagOnlineClientStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagOnlineClientStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagOnlineClientStatisticsResponseBody) GetSagStatistics() *DescribeSagOnlineClientStatisticsResponseBodySagStatistics {
	return s.SagStatistics
}

func (s *DescribeSagOnlineClientStatisticsResponseBody) SetRequestId(v string) *DescribeSagOnlineClientStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagOnlineClientStatisticsResponseBody) SetSagStatistics(v *DescribeSagOnlineClientStatisticsResponseBodySagStatistics) *DescribeSagOnlineClientStatisticsResponseBody {
	s.SagStatistics = v
	return s
}

func (s *DescribeSagOnlineClientStatisticsResponseBody) Validate() error {
	if s.SagStatistics != nil {
		if err := s.SagStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSagOnlineClientStatisticsResponseBodySagStatistics struct {
	Statistics []*DescribeSagOnlineClientStatisticsResponseBodySagStatisticsStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeSagOnlineClientStatisticsResponseBodySagStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagOnlineClientStatisticsResponseBodySagStatistics) GoString() string {
	return s.String()
}

func (s *DescribeSagOnlineClientStatisticsResponseBodySagStatistics) GetStatistics() []*DescribeSagOnlineClientStatisticsResponseBodySagStatisticsStatistics {
	return s.Statistics
}

func (s *DescribeSagOnlineClientStatisticsResponseBodySagStatistics) SetStatistics(v []*DescribeSagOnlineClientStatisticsResponseBodySagStatisticsStatistics) *DescribeSagOnlineClientStatisticsResponseBodySagStatistics {
	s.Statistics = v
	return s
}

func (s *DescribeSagOnlineClientStatisticsResponseBodySagStatistics) Validate() error {
	if s.Statistics != nil {
		for _, item := range s.Statistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSagOnlineClientStatisticsResponseBodySagStatisticsStatistics struct {
	OnlineCount *string `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	SmartAGId   *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DescribeSagOnlineClientStatisticsResponseBodySagStatisticsStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagOnlineClientStatisticsResponseBodySagStatisticsStatistics) GoString() string {
	return s.String()
}

func (s *DescribeSagOnlineClientStatisticsResponseBodySagStatisticsStatistics) GetOnlineCount() *string {
	return s.OnlineCount
}

func (s *DescribeSagOnlineClientStatisticsResponseBodySagStatisticsStatistics) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagOnlineClientStatisticsResponseBodySagStatisticsStatistics) SetOnlineCount(v string) *DescribeSagOnlineClientStatisticsResponseBodySagStatisticsStatistics {
	s.OnlineCount = &v
	return s
}

func (s *DescribeSagOnlineClientStatisticsResponseBodySagStatisticsStatistics) SetSmartAGId(v string) *DescribeSagOnlineClientStatisticsResponseBodySagStatisticsStatistics {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagOnlineClientStatisticsResponseBodySagStatisticsStatistics) Validate() error {
	return dara.Validate(s)
}
