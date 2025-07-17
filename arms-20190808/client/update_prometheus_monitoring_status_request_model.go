// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusMonitoringStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdatePrometheusMonitoringStatusRequest
	GetClusterId() *string
	SetMonitoringName(v string) *UpdatePrometheusMonitoringStatusRequest
	GetMonitoringName() *string
	SetRegionId(v string) *UpdatePrometheusMonitoringStatusRequest
	GetRegionId() *string
	SetStatus(v string) *UpdatePrometheusMonitoringStatusRequest
	GetStatus() *string
	SetType(v string) *UpdatePrometheusMonitoringStatusRequest
	GetType() *string
}

type UpdatePrometheusMonitoringStatusRequest struct {
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
	// The status of the monitoring configuration. Valid values: run and stop. The status of Probe cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// run
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the monitoring configuration.
	//
	// Valid values for a Prometheus instance for Container Service: serviceMonitor, podMonitor, and customJob.
	//
	// Valid value for a Prometheus instance for ECS: customJob.
	//
	// The status of probe cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// customJob
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdatePrometheusMonitoringStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusMonitoringStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusMonitoringStatusRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdatePrometheusMonitoringStatusRequest) GetMonitoringName() *string {
	return s.MonitoringName
}

func (s *UpdatePrometheusMonitoringStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePrometheusMonitoringStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdatePrometheusMonitoringStatusRequest) GetType() *string {
	return s.Type
}

func (s *UpdatePrometheusMonitoringStatusRequest) SetClusterId(v string) *UpdatePrometheusMonitoringStatusRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdatePrometheusMonitoringStatusRequest) SetMonitoringName(v string) *UpdatePrometheusMonitoringStatusRequest {
	s.MonitoringName = &v
	return s
}

func (s *UpdatePrometheusMonitoringStatusRequest) SetRegionId(v string) *UpdatePrometheusMonitoringStatusRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePrometheusMonitoringStatusRequest) SetStatus(v string) *UpdatePrometheusMonitoringStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdatePrometheusMonitoringStatusRequest) SetType(v string) *UpdatePrometheusMonitoringStatusRequest {
	s.Type = &v
	return s
}

func (s *UpdatePrometheusMonitoringStatusRequest) Validate() error {
	return dara.Validate(s)
}
