// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusMonitoringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetPrometheusMonitoringRequest
	GetClusterId() *string
	SetMonitoringName(v string) *GetPrometheusMonitoringRequest
	GetMonitoringName() *string
	SetRegionId(v string) *GetPrometheusMonitoringRequest
	GetRegionId() *string
	SetType(v string) *GetPrometheusMonitoringRequest
	GetType() *string
}

type GetPrometheusMonitoringRequest struct {
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
	// customJob
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPrometheusMonitoringRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusMonitoringRequest) GoString() string {
	return s.String()
}

func (s *GetPrometheusMonitoringRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetPrometheusMonitoringRequest) GetMonitoringName() *string {
	return s.MonitoringName
}

func (s *GetPrometheusMonitoringRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPrometheusMonitoringRequest) GetType() *string {
	return s.Type
}

func (s *GetPrometheusMonitoringRequest) SetClusterId(v string) *GetPrometheusMonitoringRequest {
	s.ClusterId = &v
	return s
}

func (s *GetPrometheusMonitoringRequest) SetMonitoringName(v string) *GetPrometheusMonitoringRequest {
	s.MonitoringName = &v
	return s
}

func (s *GetPrometheusMonitoringRequest) SetRegionId(v string) *GetPrometheusMonitoringRequest {
	s.RegionId = &v
	return s
}

func (s *GetPrometheusMonitoringRequest) SetType(v string) *GetPrometheusMonitoringRequest {
	s.Type = &v
	return s
}

func (s *GetPrometheusMonitoringRequest) Validate() error {
	return dara.Validate(s)
}
