// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartNisTrafficRankingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *StartNisTrafficRankingShrinkRequest
	GetBeginTime() *int64
	SetDirection(v string) *StartNisTrafficRankingShrinkRequest
	GetDirection() *string
	SetEndTime(v int64) *StartNisTrafficRankingShrinkRequest
	GetEndTime() *int64
	SetFilterShrink(v string) *StartNisTrafficRankingShrinkRequest
	GetFilterShrink() *string
	SetGroupByShrink(v string) *StartNisTrafficRankingShrinkRequest
	GetGroupByShrink() *string
	SetLanguage(v string) *StartNisTrafficRankingShrinkRequest
	GetLanguage() *string
	SetMaxResults(v int32) *StartNisTrafficRankingShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *StartNisTrafficRankingShrinkRequest
	GetNextToken() *string
	SetOrderBy(v string) *StartNisTrafficRankingShrinkRequest
	GetOrderBy() *string
	SetRegionNo(v string) *StartNisTrafficRankingShrinkRequest
	GetRegionNo() *string
	SetSort(v string) *StartNisTrafficRankingShrinkRequest
	GetSort() *string
	SetStorageInterval(v int32) *StartNisTrafficRankingShrinkRequest
	GetStorageInterval() *int32
	SetTopN(v int32) *StartNisTrafficRankingShrinkRequest
	GetTopN() *int32
	SetTrafficAnalyzerId(v string) *StartNisTrafficRankingShrinkRequest
	GetTrafficAnalyzerId() *string
	SetTrafficScenario(v string) *StartNisTrafficRankingShrinkRequest
	GetTrafficScenario() *string
	SetTupleDimension(v string) *StartNisTrafficRankingShrinkRequest
	GetTupleDimension() *string
}

type StartNisTrafficRankingShrinkRequest struct {
	// The start timestamp of the query, in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638239092000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The network traffic direction based on Alibaba Cloud resources.
	//
	// In: Traffic flowing into the target resource.
	//
	// Out: Traffic flowing out of the target resource.
	//
	// - VPC flow log scenario (`TraffficScenario = VpcFlowLogAll` / `VpcFlowLogInternet`):
	//
	//   - In: Traffic flowing into the ENI.
	//
	//   - Out: Traffic flowing out of the ENI.
	//
	// - TR flow log scenario (`TraffficScenario = TRFlowlog`):
	//
	//   - In: Traffic flowing into the TR.
	//
	//   - Out: Traffic flowing out of the TR.
	//
	// - Internet Shared Bandwidth metric analysis scenario (`TraffficScenario = CbwpMetric`):
	//
	//   - In: Traffic flowing into the EIP.
	//
	//   - Out: Traffic flowing out of the EIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// Out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end timestamp of the query, in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684373700099
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies additional filter conditions for focused network traffic analysis.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies multiple traffic dimensions for aggregation and sorting.
	GroupByShrink *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// The language. Valid values: zh-CN, en-US.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The page size. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. Leave this parameter empty for the first query or when no more results are available. If a next query exists, set this value to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Based on the `TrafficScenario` field, the following metrics are supported for ranking traffic:
	//
	// - `TrafficScenario = VpcFlowLogAll` / `VpcFlowLogInternet` (VPC flow log scenario):
	//
	//   - `Bytes`: Bandwidth
	//
	//   - `Packets`: Packets
	//
	//   - `RoundTripTime`: TCP RTT
	//
	// - `TrafficScenario = TRFlowlog` (TR flow log scenario):
	//
	//   - `Bytes`: Bandwidth
	//
	//   - `Packets`: Packets
	//
	//   - `PacketsLostNoRoute`: Packet loss due to no routing
	//
	//   - `PacketsLostBlackhole`: Packet loss due to blackhole routing
	//
	//   - `PacketsLostTTLExpired`: Packet loss due to TTL timeout
	//
	//   - `BytesIncrease`: Bandwidth increase
	//
	//   - `BytesIncreaseRatio`: Bandwidth increase ratio
	//
	// - `TrafficScenario = CbwpMetric` (Internet Shared Bandwidth metric analysis scenario):
	//
	//   - `Bytes`: Bandwidth
	//
	//   - `Packets`: Packets
	//
	// This parameter is required.
	//
	// example:
	//
	// Bytes
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The region where the resource resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The sorting method for network traffic analysis. Valid values:
	//
	// - ASC: Sorts in ascending order.
	//
	// - DESC: Sorts in descending order.
	//
	// example:
	//
	// Desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The storage bucket precision property.
	//
	// The storage bucket precision specifies the storage aggregation epoch to query. Two precision levels are supported: high precision (such as 1 minute) and long epoch (such as 1 day). The specific precision is determined by the network traffic analysis sampling interval configured for high-precision traffic statistics or long-epoch traffic statistics when creating or editing the network traffic analysis instance.
	//
	// - The storage precision supported by the corresponding tuple of the network traffic analysis instance:
	//
	//   - `1`: In minutes (1 minute)
	//
	//   - `10`: In minutes (10 minutes)
	//
	//   - `60`: In minutes (60 minutes, or 1 hour)
	//
	//   - `1440`: In minutes (1440 minutes, or 1 day)
	//
	// - The storage bucket precision can be used for two typical purposes:
	//
	//   - High-precision traffic statistics: Aggregation at 1-minute, 10-minute, or 60-minute intervals.
	//
	//   - Long-epoch traffic statistics: Aggregation at 1440-minute (1-day) intervals.
	//
	// - Pass a value for this field during the query to specify the storage aggregation epoch. For example:
	//
	//   - Pass `10`: Queries short-epoch data aggregated at 10-minute granularity.
	//
	//   - Pass `1440`: Queries long-epoch data aggregated at 1-day granularity.
	//
	// Note: The active storage precision values depend on the configuration of the network traffic analysis instance.
	//
	// example:
	//
	// 10
	StorageInterval *int32 `json:"StorageInterval,omitempty" xml:"StorageInterval,omitempty"`
	// The number of entries for the network traffic analysis sorting query.
	//
	// You can specify a custom number. If this field is not specified, all traffic data that meets the specified conditions is sorted and analyzed within the performance capacity of the network traffic analysis feature.
	//
	// example:
	//
	// 10
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
	// The ID of the network traffic analysis instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// nta-262****ca07f
	TrafficAnalyzerId *string `json:"TrafficAnalyzerId,omitempty" xml:"TrafficAnalyzerId,omitempty"`
	// Supported analysis scenarios:
	//
	// - All VPC flow log analysis
	//
	// - Public VPC flow log analysis
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
	// The storage aggregation dimension of the network traffic analysis instance.
	//
	// Based on the TraffficScenario:
	//
	// - VpcFlowLogAll/VpcFlowLog: Required. Specifies the storage aggregation view to query, which corresponds to the storage aggregation property configured in the network traffic analysis instance.
	//
	// - TRFlowLog/CbwpMetric: Optional. Automatically adapts based on the storage aggregation property of the network traffic analysis instance.
	//
	// example:
	//
	// Tuple1
	TupleDimension *string `json:"TupleDimension,omitempty" xml:"TupleDimension,omitempty"`
}

func (s StartNisTrafficRankingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartNisTrafficRankingShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartNisTrafficRankingShrinkRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *StartNisTrafficRankingShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *StartNisTrafficRankingShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *StartNisTrafficRankingShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *StartNisTrafficRankingShrinkRequest) GetGroupByShrink() *string {
	return s.GroupByShrink
}

func (s *StartNisTrafficRankingShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *StartNisTrafficRankingShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *StartNisTrafficRankingShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *StartNisTrafficRankingShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *StartNisTrafficRankingShrinkRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *StartNisTrafficRankingShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *StartNisTrafficRankingShrinkRequest) GetStorageInterval() *int32 {
	return s.StorageInterval
}

func (s *StartNisTrafficRankingShrinkRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *StartNisTrafficRankingShrinkRequest) GetTrafficAnalyzerId() *string {
	return s.TrafficAnalyzerId
}

func (s *StartNisTrafficRankingShrinkRequest) GetTrafficScenario() *string {
	return s.TrafficScenario
}

func (s *StartNisTrafficRankingShrinkRequest) GetTupleDimension() *string {
	return s.TupleDimension
}

func (s *StartNisTrafficRankingShrinkRequest) SetBeginTime(v int64) *StartNisTrafficRankingShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetDirection(v string) *StartNisTrafficRankingShrinkRequest {
	s.Direction = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetEndTime(v int64) *StartNisTrafficRankingShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetFilterShrink(v string) *StartNisTrafficRankingShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetGroupByShrink(v string) *StartNisTrafficRankingShrinkRequest {
	s.GroupByShrink = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetLanguage(v string) *StartNisTrafficRankingShrinkRequest {
	s.Language = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetMaxResults(v int32) *StartNisTrafficRankingShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetNextToken(v string) *StartNisTrafficRankingShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetOrderBy(v string) *StartNisTrafficRankingShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetRegionNo(v string) *StartNisTrafficRankingShrinkRequest {
	s.RegionNo = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetSort(v string) *StartNisTrafficRankingShrinkRequest {
	s.Sort = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetStorageInterval(v int32) *StartNisTrafficRankingShrinkRequest {
	s.StorageInterval = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetTopN(v int32) *StartNisTrafficRankingShrinkRequest {
	s.TopN = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetTrafficAnalyzerId(v string) *StartNisTrafficRankingShrinkRequest {
	s.TrafficAnalyzerId = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetTrafficScenario(v string) *StartNisTrafficRankingShrinkRequest {
	s.TrafficScenario = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetTupleDimension(v string) *StartNisTrafficRankingShrinkRequest {
	s.TupleDimension = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
