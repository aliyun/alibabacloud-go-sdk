// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyName(v []*string) *DescribeMonitorDataRequest
	GetApiKeyName() []*string
	SetEndTime(v int64) *DescribeMonitorDataRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeMonitorDataRequest
	GetInstanceId() *string
	SetInterval(v int32) *DescribeMonitorDataRequest
	GetInterval() *int32
	SetMetric(v string) *DescribeMonitorDataRequest
	GetMetric() *string
	SetStartTime(v int64) *DescribeMonitorDataRequest
	GetStartTime() *int64
}

type DescribeMonitorDataRequest struct {
	ApiKeyName []*string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty" type:"Repeated"`
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

func (s DescribeMonitorDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorDataRequest) GetApiKeyName() []*string {
	return s.ApiKeyName
}

func (s *DescribeMonitorDataRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeMonitorDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMonitorDataRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeMonitorDataRequest) GetMetric() *string {
	return s.Metric
}

func (s *DescribeMonitorDataRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeMonitorDataRequest) SetApiKeyName(v []*string) *DescribeMonitorDataRequest {
	s.ApiKeyName = v
	return s
}

func (s *DescribeMonitorDataRequest) SetEndTime(v int64) *DescribeMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetInstanceId(v string) *DescribeMonitorDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetInterval(v int32) *DescribeMonitorDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetMetric(v string) *DescribeMonitorDataRequest {
	s.Metric = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetStartTime(v int64) *DescribeMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMonitorDataRequest) Validate() error {
	return dara.Validate(s)
}
