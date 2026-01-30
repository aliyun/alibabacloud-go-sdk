// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceCreateAndDeleteStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceCreateAndDeleteStatistics(v *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatistics) *DescribeInstanceCreateAndDeleteStatisticsResponseBody
	GetInstanceCreateAndDeleteStatistics() *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatistics
	SetRequestId(v string) *DescribeInstanceCreateAndDeleteStatisticsResponseBody
	GetRequestId() *string
}

type DescribeInstanceCreateAndDeleteStatisticsResponseBody struct {
	// Metrics for instance creation and deletion.
	InstanceCreateAndDeleteStatistics *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatistics `json:"InstanceCreateAndDeleteStatistics,omitempty" xml:"InstanceCreateAndDeleteStatistics,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 086EFCD4-C76F-4DC6-9EE9-0D9B711E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceCreateAndDeleteStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceCreateAndDeleteStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBody) GetInstanceCreateAndDeleteStatistics() *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatistics {
	return s.InstanceCreateAndDeleteStatistics
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBody) SetInstanceCreateAndDeleteStatistics(v *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatistics) *DescribeInstanceCreateAndDeleteStatisticsResponseBody {
	s.InstanceCreateAndDeleteStatistics = v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBody) SetRequestId(v string) *DescribeInstanceCreateAndDeleteStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBody) Validate() error {
	if s.InstanceCreateAndDeleteStatistics != nil {
		if err := s.InstanceCreateAndDeleteStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatistics struct {
	Statistic []*DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
}

func (s DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatistics) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatistics) GetStatistic() []*DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic {
	return s.Statistic
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatistics) SetStatistic(v []*DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic) *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatistics {
	s.Statistic = v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatistics) Validate() error {
	if s.Statistic != nil {
		for _, item := range s.Statistic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic struct {
	// The number of new instances.
	//
	// example:
	//
	// 12
	CreatedVmCount *int32 `json:"CreatedVmCount,omitempty" xml:"CreatedVmCount,omitempty"`
	// The number of released instances.
	//
	// example:
	//
	// 34
	DestroyedVmCount *int32 `json:"DestroyedVmCount,omitempty" xml:"DestroyedVmCount,omitempty"`
	// The number of instances that are started from economical mode.
	//
	// example:
	//
	// 5
	StartedVmCount *int32 `json:"StartedVmCount,omitempty" xml:"StartedVmCount,omitempty"`
	// The number of instances that are stopped in economical mode.
	//
	// example:
	//
	// 10
	StoppedVmCount *int32 `json:"StoppedVmCount,omitempty" xml:"StoppedVmCount,omitempty"`
	// The time when the statistics are generated. The time is in UTC. Format: yyyyMMddHH.
	//
	// example:
	//
	// 2025121623
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic) GetCreatedVmCount() *int32 {
	return s.CreatedVmCount
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic) GetDestroyedVmCount() *int32 {
	return s.DestroyedVmCount
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic) GetStartedVmCount() *int32 {
	return s.StartedVmCount
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic) GetStoppedVmCount() *int32 {
	return s.StoppedVmCount
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic) GetTime() *string {
	return s.Time
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic) SetCreatedVmCount(v int32) *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic {
	s.CreatedVmCount = &v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic) SetDestroyedVmCount(v int32) *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic {
	s.DestroyedVmCount = &v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic) SetStartedVmCount(v int32) *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic {
	s.StartedVmCount = &v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic) SetStoppedVmCount(v int32) *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic {
	s.StoppedVmCount = &v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic) SetTime(v string) *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic {
	s.Time = &v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsResponseBodyInstanceCreateAndDeleteStatisticsStatistic) Validate() error {
	return dara.Validate(s)
}
