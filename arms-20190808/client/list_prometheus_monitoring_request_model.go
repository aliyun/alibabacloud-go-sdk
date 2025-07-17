// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusMonitoringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListPrometheusMonitoringRequest
	GetClusterId() *string
	SetRegionId(v string) *ListPrometheusMonitoringRequest
	GetRegionId() *string
	SetType(v string) *ListPrometheusMonitoringRequest
	GetType() *string
}

type ListPrometheusMonitoringRequest struct {
	// The ID of the Prometheus instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID. Default value: `cn-hangzhou`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the monitoring configuration. Valid values for a Prometheus instance for Container Service: ServiceMonitor, PodMonitor, CustomJob, and Probe. Valid values for a Prometheus instance for ECS: CustomJob and Probe.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// serviceMonitor
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPrometheusMonitoringRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusMonitoringRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusMonitoringRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListPrometheusMonitoringRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrometheusMonitoringRequest) GetType() *string {
	return s.Type
}

func (s *ListPrometheusMonitoringRequest) SetClusterId(v string) *ListPrometheusMonitoringRequest {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusMonitoringRequest) SetRegionId(v string) *ListPrometheusMonitoringRequest {
	s.RegionId = &v
	return s
}

func (s *ListPrometheusMonitoringRequest) SetType(v string) *ListPrometheusMonitoringRequest {
	s.Type = &v
	return s
}

func (s *ListPrometheusMonitoringRequest) Validate() error {
	return dara.Validate(s)
}
