// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLLMGatewayMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListLLMGatewayMetricsRequest
	GetEndTime() *string
	SetGranularity(v string) *ListLLMGatewayMetricsRequest
	GetGranularity() *string
	SetInstanceName(v string) *ListLLMGatewayMetricsRequest
	GetInstanceName() *string
	SetMetricType(v string) *ListLLMGatewayMetricsRequest
	GetMetricType() *string
	SetModel(v string) *ListLLMGatewayMetricsRequest
	GetModel() *string
	SetRegionId(v string) *ListLLMGatewayMetricsRequest
	GetRegionId() *string
	SetStartTime(v string) *ListLLMGatewayMetricsRequest
	GetStartTime() *string
}

type ListLLMGatewayMetricsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2026-05-26T20:55:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// minute
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-7kox9010itfcn1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tokenPerMinute
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// example:
	//
	// qwen-plus,deepseek-v4-pro
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-05-26T20:40:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListLLMGatewayMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLLMGatewayMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListLLMGatewayMetricsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListLLMGatewayMetricsRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *ListLLMGatewayMetricsRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListLLMGatewayMetricsRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *ListLLMGatewayMetricsRequest) GetModel() *string {
	return s.Model
}

func (s *ListLLMGatewayMetricsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLLMGatewayMetricsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListLLMGatewayMetricsRequest) SetEndTime(v string) *ListLLMGatewayMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ListLLMGatewayMetricsRequest) SetGranularity(v string) *ListLLMGatewayMetricsRequest {
	s.Granularity = &v
	return s
}

func (s *ListLLMGatewayMetricsRequest) SetInstanceName(v string) *ListLLMGatewayMetricsRequest {
	s.InstanceName = &v
	return s
}

func (s *ListLLMGatewayMetricsRequest) SetMetricType(v string) *ListLLMGatewayMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *ListLLMGatewayMetricsRequest) SetModel(v string) *ListLLMGatewayMetricsRequest {
	s.Model = &v
	return s
}

func (s *ListLLMGatewayMetricsRequest) SetRegionId(v string) *ListLLMGatewayMetricsRequest {
	s.RegionId = &v
	return s
}

func (s *ListLLMGatewayMetricsRequest) SetStartTime(v string) *ListLLMGatewayMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *ListLLMGatewayMetricsRequest) Validate() error {
	return dara.Validate(s)
}
