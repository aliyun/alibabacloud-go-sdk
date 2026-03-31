// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeThreatEventTopMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *DescribeThreatEventTopMetricRequest
	GetEventId() *string
	SetInstanceId(v string) *DescribeThreatEventTopMetricRequest
	GetInstanceId() *string
	SetMetric(v string) *DescribeThreatEventTopMetricRequest
	GetMetric() *string
	SetRegionId(v string) *DescribeThreatEventTopMetricRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeThreatEventTopMetricRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeThreatEventTopMetricRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0b7ab137a065aab7656986***11db
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// time
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeThreatEventTopMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatEventTopMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeThreatEventTopMetricRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeThreatEventTopMetricRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeThreatEventTopMetricRequest) GetMetric() *string {
	return s.Metric
}

func (s *DescribeThreatEventTopMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeThreatEventTopMetricRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeThreatEventTopMetricRequest) SetEventId(v string) *DescribeThreatEventTopMetricRequest {
	s.EventId = &v
	return s
}

func (s *DescribeThreatEventTopMetricRequest) SetInstanceId(v string) *DescribeThreatEventTopMetricRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeThreatEventTopMetricRequest) SetMetric(v string) *DescribeThreatEventTopMetricRequest {
	s.Metric = &v
	return s
}

func (s *DescribeThreatEventTopMetricRequest) SetRegionId(v string) *DescribeThreatEventTopMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeThreatEventTopMetricRequest) SetResourceManagerResourceGroupId(v string) *DescribeThreatEventTopMetricRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeThreatEventTopMetricRequest) Validate() error {
	return dara.Validate(s)
}
