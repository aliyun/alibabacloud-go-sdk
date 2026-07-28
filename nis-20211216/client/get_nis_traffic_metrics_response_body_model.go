// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNisTrafficMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *GetNisTrafficMetricsResponseBody
	GetMaxResults() *int32
	SetMetricStatics(v []*GetNisTrafficMetricsResponseBodyMetricStatics) *GetNisTrafficMetricsResponseBody
	GetMetricStatics() []*GetNisTrafficMetricsResponseBodyMetricStatics
	SetNextToken(v string) *GetNisTrafficMetricsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetNisTrafficMetricsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *GetNisTrafficMetricsResponseBody
	GetTotalCount() *int32
	SetUnit(v string) *GetNisTrafficMetricsResponseBody
	GetUnit() *string
}

type GetNisTrafficMetricsResponseBody struct {
	// The maximum number of entries returned per page or per query. In VPC scenarios, this represents the paging size. In TR and Internet Shared Bandwidth scenarios, this represents the SQL query limit.
	//
	// example:
	//
	// 1440
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The list of time series metric data points. Each element represents an aggregated time point and its corresponding metric value.
	MetricStatics []*GetNisTrafficMetricsResponseBodyMetricStatics `json:"MetricStatics,omitempty" xml:"MetricStatics,omitempty" type:"Repeated"`
	// The paging token for the next page. Paging is supported only in VPC scenarios. An empty value indicates that no more pages exist. This field is typically not returned in TR and Internet Shared Bandwidth scenarios.
	//
	// example:
	//
	// qqt9NJ3/AWeMXCntK4Kyhrt0QclAmfbtYB4899hEUzVNEo/F148UPCh2itDku111
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FA764-BA47-56F8-88E1-7AB503A62112
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of data points in the query result.
	//
	// example:
	//
	// 1440
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// **Unit and MetricName mapping**
	//
	// - Bandwidth
	//
	//   - Unit: Bits/Second
	//
	//   - Description: bits per second.
	//
	// - PacketsRate
	//
	//   - Unit: Packets/Second
	//
	//   - Description: packets per second.
	//
	// - RoundTripTime
	//
	//   - Unit: MicroSecond
	//
	//   - Description: TCP round-trip time.
	//
	// - BandwidthUtilization
	//
	//   - Unit: Percent
	//
	//   - Description: bandwidth utilization.
	//
	// - PacketsLostNoRouteRate
	//
	//   - Unit: PacketsLostNoRouteRate
	//
	//   - Description: rate of packets dropped due to no route.
	//
	// - PacketsLostBlackholeRate
	//
	//   - Unit: PacketsLostBlackholeRate
	//
	//   - Description: rate of packets dropped due to blackhole routing.
	//
	// - PacketsLostTTLExpiredRate
	//
	//   - Unit: PacketsLostTTLExpiredRate
	//
	//   - Description: rate of packets dropped due to TTL expiration.
	//
	// example:
	//
	// Bits/Second
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s GetNisTrafficMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNisTrafficMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetNisTrafficMetricsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetNisTrafficMetricsResponseBody) GetMetricStatics() []*GetNisTrafficMetricsResponseBodyMetricStatics {
	return s.MetricStatics
}

func (s *GetNisTrafficMetricsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetNisTrafficMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNisTrafficMetricsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetNisTrafficMetricsResponseBody) GetUnit() *string {
	return s.Unit
}

func (s *GetNisTrafficMetricsResponseBody) SetMaxResults(v int32) *GetNisTrafficMetricsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetNisTrafficMetricsResponseBody) SetMetricStatics(v []*GetNisTrafficMetricsResponseBodyMetricStatics) *GetNisTrafficMetricsResponseBody {
	s.MetricStatics = v
	return s
}

func (s *GetNisTrafficMetricsResponseBody) SetNextToken(v string) *GetNisTrafficMetricsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetNisTrafficMetricsResponseBody) SetRequestId(v string) *GetNisTrafficMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNisTrafficMetricsResponseBody) SetTotalCount(v int32) *GetNisTrafficMetricsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetNisTrafficMetricsResponseBody) SetUnit(v string) *GetNisTrafficMetricsResponseBody {
	s.Unit = &v
	return s
}

func (s *GetNisTrafficMetricsResponseBody) Validate() error {
	if s.MetricStatics != nil {
		for _, item := range s.MetricStatics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNisTrafficMetricsResponseBodyMetricStatics struct {
	// The timestamp of the data point, in milliseconds.
	//
	// example:
	//
	// 1785219000000
	TimeStamp *int64 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The metric value at the current time point. The specific meaning and unit are determined by the MetricName in the request.
	//
	// example:
	//
	// 102400.25
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetNisTrafficMetricsResponseBodyMetricStatics) String() string {
	return dara.Prettify(s)
}

func (s GetNisTrafficMetricsResponseBodyMetricStatics) GoString() string {
	return s.String()
}

func (s *GetNisTrafficMetricsResponseBodyMetricStatics) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *GetNisTrafficMetricsResponseBodyMetricStatics) GetValue() *float64 {
	return s.Value
}

func (s *GetNisTrafficMetricsResponseBodyMetricStatics) SetTimeStamp(v int64) *GetNisTrafficMetricsResponseBodyMetricStatics {
	s.TimeStamp = &v
	return s
}

func (s *GetNisTrafficMetricsResponseBodyMetricStatics) SetValue(v float64) *GetNisTrafficMetricsResponseBodyMetricStatics {
	s.Value = &v
	return s
}

func (s *GetNisTrafficMetricsResponseBodyMetricStatics) Validate() error {
	return dara.Validate(s)
}
