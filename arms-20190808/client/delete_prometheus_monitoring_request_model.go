// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusMonitoringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeletePrometheusMonitoringRequest
	GetClusterId() *string
	SetMonitoringName(v string) *DeletePrometheusMonitoringRequest
	GetMonitoringName() *string
	SetRegionId(v string) *DeletePrometheusMonitoringRequest
	GetRegionId() *string
	SetType(v string) *DeletePrometheusMonitoringRequest
	GetType() *string
}

type DeletePrometheusMonitoringRequest struct {
	// The ID of the Prometheus instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the monitoring configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// customJob1
	MonitoringName *string `json:"MonitoringName,omitempty" xml:"MonitoringName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the monitoring configuration.
	//
	// Valid values for a Prometheus instance for Container Service: serviceMonitor, podMonitor, customJob, and probe.
	//
	// Valid values for a Prometheus instance for ECS: customJob and probe.
	//
	// This parameter is required.
	//
	// example:
	//
	// probe
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeletePrometheusMonitoringRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusMonitoringRequest) GoString() string {
	return s.String()
}

func (s *DeletePrometheusMonitoringRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeletePrometheusMonitoringRequest) GetMonitoringName() *string {
	return s.MonitoringName
}

func (s *DeletePrometheusMonitoringRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePrometheusMonitoringRequest) GetType() *string {
	return s.Type
}

func (s *DeletePrometheusMonitoringRequest) SetClusterId(v string) *DeletePrometheusMonitoringRequest {
	s.ClusterId = &v
	return s
}

func (s *DeletePrometheusMonitoringRequest) SetMonitoringName(v string) *DeletePrometheusMonitoringRequest {
	s.MonitoringName = &v
	return s
}

func (s *DeletePrometheusMonitoringRequest) SetRegionId(v string) *DeletePrometheusMonitoringRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePrometheusMonitoringRequest) SetType(v string) *DeletePrometheusMonitoringRequest {
	s.Type = &v
	return s
}

func (s *DeletePrometheusMonitoringRequest) Validate() error {
	return dara.Validate(s)
}
