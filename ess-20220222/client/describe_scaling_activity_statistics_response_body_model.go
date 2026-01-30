// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingActivityStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeScalingActivityStatisticsResponseBody
	GetRequestId() *string
	SetScalingActivityErrorCodeStatistics(v *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatistics) *DescribeScalingActivityStatisticsResponseBody
	GetScalingActivityErrorCodeStatistics() *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatistics
	SetScalingActivityStatusStatistics(v *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatistics) *DescribeScalingActivityStatisticsResponseBody
	GetScalingActivityStatusStatistics() *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatistics
}

type DescribeScalingActivityStatisticsResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A8F44B4D-BAB4-5176-8705-5984xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error message statistics of the scaling activity.
	ScalingActivityErrorCodeStatistics *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatistics `json:"ScalingActivityErrorCodeStatistics,omitempty" xml:"ScalingActivityErrorCodeStatistics,omitempty" type:"Struct"`
	// The statistical metrics of the scaling activity status.
	ScalingActivityStatusStatistics *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatistics `json:"ScalingActivityStatusStatistics,omitempty" xml:"ScalingActivityStatusStatistics,omitempty" type:"Struct"`
}

func (s DescribeScalingActivityStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivityStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScalingActivityStatisticsResponseBody) GetScalingActivityErrorCodeStatistics() *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatistics {
	return s.ScalingActivityErrorCodeStatistics
}

func (s *DescribeScalingActivityStatisticsResponseBody) GetScalingActivityStatusStatistics() *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatistics {
	return s.ScalingActivityStatusStatistics
}

func (s *DescribeScalingActivityStatisticsResponseBody) SetRequestId(v string) *DescribeScalingActivityStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingActivityStatisticsResponseBody) SetScalingActivityErrorCodeStatistics(v *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatistics) *DescribeScalingActivityStatisticsResponseBody {
	s.ScalingActivityErrorCodeStatistics = v
	return s
}

func (s *DescribeScalingActivityStatisticsResponseBody) SetScalingActivityStatusStatistics(v *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatistics) *DescribeScalingActivityStatisticsResponseBody {
	s.ScalingActivityStatusStatistics = v
	return s
}

func (s *DescribeScalingActivityStatisticsResponseBody) Validate() error {
	if s.ScalingActivityErrorCodeStatistics != nil {
		if err := s.ScalingActivityErrorCodeStatistics.Validate(); err != nil {
			return err
		}
	}
	if s.ScalingActivityStatusStatistics != nil {
		if err := s.ScalingActivityStatusStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatistics struct {
	ErrorStatistic []*DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic `json:"ErrorStatistic,omitempty" xml:"ErrorStatistic,omitempty" type:"Repeated"`
}

func (s DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatistics) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatistics) GetErrorStatistic() []*DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic {
	return s.ErrorStatistic
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatistics) SetErrorStatistic(v []*DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic) *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatistics {
	s.ErrorStatistic = v
	return s
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatistics) Validate() error {
	if s.ErrorStatistic != nil {
		for _, item := range s.ErrorStatistic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic struct {
	// The number of failed scaling activities.
	//
	// example:
	//
	// 4
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Scaling Activity Error Codes
	//
	// example:
	//
	// QuotaExceeded.PrivateIpAddress
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The time when the statistics are generated. The time is in UTC. Format: yyyyMMddHH.
	//
	// example:
	//
	// 2025121623
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic) GetCount() *int32 {
	return s.Count
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic) GetTime() *string {
	return s.Time
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic) SetCount(v int32) *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic {
	s.Count = &v
	return s
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic) SetErrorCode(v string) *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic {
	s.ErrorCode = &v
	return s
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic) SetTime(v string) *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic {
	s.Time = &v
	return s
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityErrorCodeStatisticsErrorStatistic) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatistics struct {
	Statistic []*DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
}

func (s DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatistics) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatistics) GetStatistic() []*DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic {
	return s.Statistic
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatistics) SetStatistic(v []*DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic) *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatistics {
	s.Statistic = v
	return s
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatistics) Validate() error {
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

type DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic struct {
	// The number of failed scaling activities.
	//
	// example:
	//
	// 1
	FailedActivityCount *int32 `json:"FailedActivityCount,omitempty" xml:"FailedActivityCount,omitempty"`
	// The number of successful scaling activities.
	//
	// example:
	//
	// 10
	SuccessActivityCount *int32 `json:"SuccessActivityCount,omitempty" xml:"SuccessActivityCount,omitempty"`
	// The time when the statistics are generated. The time is in UTC. Format: yyyyMMddHH.
	//
	// example:
	//
	// 2025121623
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The number of partially executed scaling activities.
	//
	// example:
	//
	// 2
	WarningActivityCount *int32 `json:"WarningActivityCount,omitempty" xml:"WarningActivityCount,omitempty"`
}

func (s DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic) GetFailedActivityCount() *int32 {
	return s.FailedActivityCount
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic) GetSuccessActivityCount() *int32 {
	return s.SuccessActivityCount
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic) GetTime() *string {
	return s.Time
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic) GetWarningActivityCount() *int32 {
	return s.WarningActivityCount
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic) SetFailedActivityCount(v int32) *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic {
	s.FailedActivityCount = &v
	return s
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic) SetSuccessActivityCount(v int32) *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic {
	s.SuccessActivityCount = &v
	return s
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic) SetTime(v string) *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic {
	s.Time = &v
	return s
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic) SetWarningActivityCount(v int32) *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic {
	s.WarningActivityCount = &v
	return s
}

func (s *DescribeScalingActivityStatisticsResponseBodyScalingActivityStatusStatisticsStatistic) Validate() error {
	return dara.Validate(s)
}
