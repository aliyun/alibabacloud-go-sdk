// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNisTrafficMetricsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *GetNisTrafficMetricsShrinkRequest
	GetBeginTime() *int64
	SetDirection(v string) *GetNisTrafficMetricsShrinkRequest
	GetDirection() *string
	SetEndTime(v int64) *GetNisTrafficMetricsShrinkRequest
	GetEndTime() *int64
	SetFilterShrink(v string) *GetNisTrafficMetricsShrinkRequest
	GetFilterShrink() *string
	SetMaxResults(v int32) *GetNisTrafficMetricsShrinkRequest
	GetMaxResults() *int32
	SetMetricName(v string) *GetNisTrafficMetricsShrinkRequest
	GetMetricName() *string
	SetNextToken(v string) *GetNisTrafficMetricsShrinkRequest
	GetNextToken() *string
	SetRegionNo(v string) *GetNisTrafficMetricsShrinkRequest
	GetRegionNo() *string
	SetScanBy(v string) *GetNisTrafficMetricsShrinkRequest
	GetScanBy() *string
	SetStepMinutes(v int32) *GetNisTrafficMetricsShrinkRequest
	GetStepMinutes() *int32
	SetStorageInterval(v int32) *GetNisTrafficMetricsShrinkRequest
	GetStorageInterval() *int32
	SetTrafficAnalyzerId(v string) *GetNisTrafficMetricsShrinkRequest
	GetTrafficAnalyzerId() *string
	SetTrafficScenario(v string) *GetNisTrafficMetricsShrinkRequest
	GetTrafficScenario() *string
	SetTupleDimension(v string) *GetNisTrafficMetricsShrinkRequest
	GetTupleDimension() *string
}

type GetNisTrafficMetricsShrinkRequest struct {
	// The start timestamp, in milliseconds. If not specified, the most recent 1 hour is queried by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638239092000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The network traffic direction based on Alibaba Cloud resources.
	//
	// In: traffic flowing into the target resource.
	//
	// Out: traffic flowing out of the target resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// In
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end timestamp, in milliseconds. If not specified, the most recent 1 hour is queried by default. If only BeginTime is specified, the 1 hour after BeginTime is queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684373700099
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies additional filter conditions for the traffic to perform focused network traffic analysis.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// In VPC scenarios, this parameter specifies the paging size. In TR and Internet Shared Bandwidth scenarios, this parameter specifies the SQL query limit. If not specified, the backend defaults to 1440.
	//
	// example:
	//
	// 1440
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The metric name.
	//
	// Common parameters supported in network traffic analysis scenarios:
	//
	//   bps: bits per second.
	//
	//   pps: packets per second.
	//
	// Parameters specific to the Internet scenario:
	//
	//   rtt: round-trip time when establishing a TCP protocol connection.
	//
	//   RetransmitRate: retransmission rate.
	//
	// Parameters specific to the area-level bandwidth scenario:
	//
	//   RatelimitDropPps: rate of packet loss due to rate limiting.
	//
	//   BandwidthUtilization: bandwidth utilization.
	//
	// Parameters specific to the NAT scenario:
	//
	//   ActiveSessionCount: number of concurrent sessions.
	//
	//   NewSessionPerSecond: number of new sessions per second.
	//
	// This parameter is required.
	//
	// example:
	//
	// bps
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The token for the next query. You do not need to specify this parameter for the first query or when no more results exist. If a next page exists, set this parameter to the NextToken value returned by the previous API invoke. This parameter is valid only in VPC scenarios. TR and Internet Shared Bandwidth scenarios do not use this parameter.
	//
	// example:
	//
	// f7zUd3gArYj/xjPttJo5L5dK0R+gSbfHElLqi8C2IPWMQxtV8XckOg5lk7F2bhC+
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The sort order. Valid values:
	//
	// TimestampAscending: sorts by time in ascending order.
	//
	// TimestampDescending: sorts by time in descending order.
	//
	// example:
	//
	// TimestampAscending
	ScanBy *string `json:"ScanBy,omitempty" xml:"ScanBy,omitempty"`
	// The aggregation step for time series data, in minutes. The final query granularity is the larger value between StepMinutes and the underlying storage granularity. The number of data points calculated by (EndTime-BeginTime)/StepMinutes cannot exceed 1440.
	//
	// example:
	//
	// 10
	StepMinutes *int32 `json:"StepMinutes,omitempty" xml:"StepMinutes,omitempty"`
	// The storage bucket precision property.
	//
	// The storage bucket precision specifies the storage aggregation epoch to query. Two precision levels are supported: high precision (such as 1 minute) or long epoch (such as 1 day). The specific precision is determined by the network traffic analysis sampling interval configured for high-precision traffic statistics or long-epoch traffic statistics when creating or editing the network traffic analysis analyzer.
	//
	// - The storage precisions active for the corresponding tuples of the network traffic analysis analyzer are:
	//
	//   - `1`: in minutes (1 minute)
	//
	//   - `10`: in minutes (10 minutes)
	//
	//   - `60`: in minutes (60 minutes, i.e., 1 hour)
	//
	//   - `1440`: in minutes (1440 minutes, i.e., 1 day)
	//
	// - The storage bucket precision can be used for two typical purposes:
	//
	//   - High-precision traffic statistics: such as 1-minute, 10-minute, or 60-minute aggregation
	//
	//   - Long-epoch traffic statistics: such as 1440-minute (1-day) aggregation
	//
	// - Specify a value for this field during the query to select the storage aggregation epoch. For example:
	//
	//   - Pass `10`: queries short-epoch data with a 10-minute aggregation granularity
	//
	//   - Pass `1440`: queries long-epoch data with a 1-day aggregation granularity
	//
	// example:
	//
	// 10
	StorageInterval *int32 `json:"StorageInterval,omitempty" xml:"StorageInterval,omitempty"`
	// The ID of the network traffic analysis analyzer.
	//
	// This parameter is required.
	//
	// example:
	//
	// nta-e093cb80c7c047afbd1d
	TrafficAnalyzerId *string `json:"TrafficAnalyzerId,omitempty" xml:"TrafficAnalyzerId,omitempty"`
	// The supported analysis scenarios:
	//
	// - All VPC flow log analysis
	//
	// - Internet VPC flow log analysis
	//
	// - All TR flow log analysis
	//
	// - Internet Shared Bandwidth metric analysis
	//
	// This parameter is required.
	//
	// example:
	//
	// VpcFlowLogAll
	TrafficScenario *string `json:"TrafficScenario,omitempty" xml:"TrafficScenario,omitempty"`
	// The traffic storage aggregation dimension.
	//
	// Based on the TrafficScenario:
	//
	// - VpcFlowLogAll/VpcFlowLog: required. Specifies the storage aggregation view to query, which corresponds to the storage aggregation property configured in the network traffic analysis analyzer.
	//
	// - TRFlowLog/CbwpMetric: optional. Automatically adapts based on the storage aggregation property of the network traffic analysis analyzer.
	//
	// example:
	//
	// Tuple2
	TupleDimension *string `json:"TupleDimension,omitempty" xml:"TupleDimension,omitempty"`
}

func (s GetNisTrafficMetricsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNisTrafficMetricsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetNisTrafficMetricsShrinkRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetNisTrafficMetricsShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *GetNisTrafficMetricsShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetNisTrafficMetricsShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *GetNisTrafficMetricsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetNisTrafficMetricsShrinkRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *GetNisTrafficMetricsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetNisTrafficMetricsShrinkRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *GetNisTrafficMetricsShrinkRequest) GetScanBy() *string {
	return s.ScanBy
}

func (s *GetNisTrafficMetricsShrinkRequest) GetStepMinutes() *int32 {
	return s.StepMinutes
}

func (s *GetNisTrafficMetricsShrinkRequest) GetStorageInterval() *int32 {
	return s.StorageInterval
}

func (s *GetNisTrafficMetricsShrinkRequest) GetTrafficAnalyzerId() *string {
	return s.TrafficAnalyzerId
}

func (s *GetNisTrafficMetricsShrinkRequest) GetTrafficScenario() *string {
	return s.TrafficScenario
}

func (s *GetNisTrafficMetricsShrinkRequest) GetTupleDimension() *string {
	return s.TupleDimension
}

func (s *GetNisTrafficMetricsShrinkRequest) SetBeginTime(v int64) *GetNisTrafficMetricsShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetNisTrafficMetricsShrinkRequest) SetDirection(v string) *GetNisTrafficMetricsShrinkRequest {
	s.Direction = &v
	return s
}

func (s *GetNisTrafficMetricsShrinkRequest) SetEndTime(v int64) *GetNisTrafficMetricsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetNisTrafficMetricsShrinkRequest) SetFilterShrink(v string) *GetNisTrafficMetricsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *GetNisTrafficMetricsShrinkRequest) SetMaxResults(v int32) *GetNisTrafficMetricsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *GetNisTrafficMetricsShrinkRequest) SetMetricName(v string) *GetNisTrafficMetricsShrinkRequest {
	s.MetricName = &v
	return s
}

func (s *GetNisTrafficMetricsShrinkRequest) SetNextToken(v string) *GetNisTrafficMetricsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *GetNisTrafficMetricsShrinkRequest) SetRegionNo(v string) *GetNisTrafficMetricsShrinkRequest {
	s.RegionNo = &v
	return s
}

func (s *GetNisTrafficMetricsShrinkRequest) SetScanBy(v string) *GetNisTrafficMetricsShrinkRequest {
	s.ScanBy = &v
	return s
}

func (s *GetNisTrafficMetricsShrinkRequest) SetStepMinutes(v int32) *GetNisTrafficMetricsShrinkRequest {
	s.StepMinutes = &v
	return s
}

func (s *GetNisTrafficMetricsShrinkRequest) SetStorageInterval(v int32) *GetNisTrafficMetricsShrinkRequest {
	s.StorageInterval = &v
	return s
}

func (s *GetNisTrafficMetricsShrinkRequest) SetTrafficAnalyzerId(v string) *GetNisTrafficMetricsShrinkRequest {
	s.TrafficAnalyzerId = &v
	return s
}

func (s *GetNisTrafficMetricsShrinkRequest) SetTrafficScenario(v string) *GetNisTrafficMetricsShrinkRequest {
	s.TrafficScenario = &v
	return s
}

func (s *GetNisTrafficMetricsShrinkRequest) SetTupleDimension(v string) *GetNisTrafficMetricsShrinkRequest {
	s.TupleDimension = &v
	return s
}

func (s *GetNisTrafficMetricsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
