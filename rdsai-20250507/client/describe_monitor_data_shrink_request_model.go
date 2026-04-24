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
	ApiKeyNameShrink *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1627269085
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 15
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ● pv
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
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
