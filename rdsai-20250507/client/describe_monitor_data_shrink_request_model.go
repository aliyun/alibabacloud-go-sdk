// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyNameShrink(v string) *DescribeMonitorDataShrinkRequest
	GetApiKeyNameShrink() *string
	SetEndTime(v int64) *DescribeMonitorDataShrinkRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeMonitorDataShrinkRequest
	GetInstanceId() *string
	SetInterval(v int32) *DescribeMonitorDataShrinkRequest
	GetInterval() *int32
	SetMetric(v string) *DescribeMonitorDataShrinkRequest
	GetMetric() *string
	SetStartTime(v int64) *DescribeMonitorDataShrinkRequest
	GetStartTime() *int64
}

type DescribeMonitorDataShrinkRequest struct {
	// The names of the API keys to use for filtering the data. If this parameter is not specified, data from all keys is returned.
	ApiKeyNameShrink *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
	// The end of the query time range, specified as a Unix timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1627269085
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The aggregation interval for monitoring data, in seconds. Default: 15.
	//
	// example:
	//
	// 15
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The metric to query. Valid values:
	//
	// - `pv`
	//
	// - `uv`
	//
	// - `qps`
	//
	// - `success_rate`
	//
	// - `rt`
	//
	// - `rate_limited_count`
	//
	// This parameter is required.
	//
	// example:
	//
	// ● pv
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The start of the query time range, specified as a Unix timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1627268185
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMonitorDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorDataShrinkRequest) GetApiKeyNameShrink() *string {
	return s.ApiKeyNameShrink
}

func (s *DescribeMonitorDataShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeMonitorDataShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMonitorDataShrinkRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeMonitorDataShrinkRequest) GetMetric() *string {
	return s.Metric
}

func (s *DescribeMonitorDataShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeMonitorDataShrinkRequest) SetApiKeyNameShrink(v string) *DescribeMonitorDataShrinkRequest {
	s.ApiKeyNameShrink = &v
	return s
}

func (s *DescribeMonitorDataShrinkRequest) SetEndTime(v int64) *DescribeMonitorDataShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMonitorDataShrinkRequest) SetInstanceId(v string) *DescribeMonitorDataShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitorDataShrinkRequest) SetInterval(v int32) *DescribeMonitorDataShrinkRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMonitorDataShrinkRequest) SetMetric(v string) *DescribeMonitorDataShrinkRequest {
	s.Metric = &v
	return s
}

func (s *DescribeMonitorDataShrinkRequest) SetStartTime(v int64) *DescribeMonitorDataShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMonitorDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
