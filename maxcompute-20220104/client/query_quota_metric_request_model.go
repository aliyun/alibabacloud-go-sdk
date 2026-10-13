// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQuotaMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInterval(v int64) *QueryQuotaMetricRequest
	GetInterval() *int64
	SetNickname(v string) *QueryQuotaMetricRequest
	GetNickname() *string
	SetSubMetric(v string) *QueryQuotaMetricRequest
	GetSubMetric() *string
	SetSubQuotaNickname(v string) *QueryQuotaMetricRequest
	GetSubQuotaNickname() *string
	SetEndTime(v int64) *QueryQuotaMetricRequest
	GetEndTime() *int64
	SetStartTime(v int64) *QueryQuotaMetricRequest
	GetStartTime() *int64
	SetStrategy(v string) *QueryQuotaMetricRequest
	GetStrategy() *string
}

type QueryQuotaMetricRequest struct {
	// The fixed interval in seconds. By default, this parameter is empty, which indicates that the adaptive interval policy of the system is used.
	//
	// - Adaptive interval policy: The interval is 60 seconds if the query time range is within 6 hours, 300 seconds if within 24 hours, 900 seconds if within 72 hours, and 1800 seconds if exceeding 72 hours.
	//
	// - Specified interval: Only 60, 300, and 900 are supported. The query time range must be within 72 hours.
	//
	// example:
	//
	// 60
	Interval *int64 `json:"interval,omitempty" xml:"interval,omitempty"`
	// The nickname of the level-1 quota. This parameter is required.
	//
	// example:
	//
	// os_sns_p
	Nickname  *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	SubMetric *string `json:"subMetric,omitempty" xml:"subMetric,omitempty"`
	// The nickname of the level-2 quota.
	//
	// example:
	//
	// os_sns
	SubQuotaNickname *string `json:"subQuotaNickname,omitempty" xml:"subQuotaNickname,omitempty"`
	// The end time of the query time range. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1735536322
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start time of the query time range. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1735534322
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The data aggregation policy. Default value: max. Valid values: max and avg.
	//
	// Data is collected every 1 minute. If the query time range is large, the adaptive step size for data display exceeds 1 minute, and the metrics are aggregated. This parameter specifies the data aggregation logic.
	//
	// example:
	//
	// max
	Strategy *string `json:"strategy,omitempty" xml:"strategy,omitempty"`
}

func (s QueryQuotaMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaMetricRequest) GoString() string {
	return s.String()
}

func (s *QueryQuotaMetricRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *QueryQuotaMetricRequest) GetNickname() *string {
	return s.Nickname
}

func (s *QueryQuotaMetricRequest) GetSubMetric() *string {
	return s.SubMetric
}

func (s *QueryQuotaMetricRequest) GetSubQuotaNickname() *string {
	return s.SubQuotaNickname
}

func (s *QueryQuotaMetricRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryQuotaMetricRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryQuotaMetricRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *QueryQuotaMetricRequest) SetInterval(v int64) *QueryQuotaMetricRequest {
	s.Interval = &v
	return s
}

func (s *QueryQuotaMetricRequest) SetNickname(v string) *QueryQuotaMetricRequest {
	s.Nickname = &v
	return s
}

func (s *QueryQuotaMetricRequest) SetSubMetric(v string) *QueryQuotaMetricRequest {
	s.SubMetric = &v
	return s
}

func (s *QueryQuotaMetricRequest) SetSubQuotaNickname(v string) *QueryQuotaMetricRequest {
	s.SubQuotaNickname = &v
	return s
}

func (s *QueryQuotaMetricRequest) SetEndTime(v int64) *QueryQuotaMetricRequest {
	s.EndTime = &v
	return s
}

func (s *QueryQuotaMetricRequest) SetStartTime(v int64) *QueryQuotaMetricRequest {
	s.StartTime = &v
	return s
}

func (s *QueryQuotaMetricRequest) SetStrategy(v string) *QueryQuotaMetricRequest {
	s.Strategy = &v
	return s
}

func (s *QueryQuotaMetricRequest) Validate() error {
	return dara.Validate(s)
}
