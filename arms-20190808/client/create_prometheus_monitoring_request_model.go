// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusMonitoringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreatePrometheusMonitoringRequest
	GetClusterId() *string
	SetConfigYaml(v string) *CreatePrometheusMonitoringRequest
	GetConfigYaml() *string
	SetRegionId(v string) *CreatePrometheusMonitoringRequest
	GetRegionId() *string
	SetStatus(v string) *CreatePrometheusMonitoringRequest
	GetStatus() *string
	SetType(v string) *CreatePrometheusMonitoringRequest
	GetType() *string
}

type CreatePrometheusMonitoringRequest struct {
	// The ID of the Prometheus instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The monitoring configuration. Specify a YAML string.
	//
	// This parameter is required.
	//
	// example:
	//
	// Please refer to the supplementary explanation of the request parameters.
	ConfigYaml *string `json:"ConfigYaml,omitempty" xml:"ConfigYaml,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the monitoring configuration. Valid values: run and stop. Default value: run. This parameter is not available if the Type parameter is set to Probe.
	//
	// example:
	//
	// run
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// serviceMonitor
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePrometheusMonitoringRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusMonitoringRequest) GoString() string {
	return s.String()
}

func (s *CreatePrometheusMonitoringRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreatePrometheusMonitoringRequest) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *CreatePrometheusMonitoringRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePrometheusMonitoringRequest) GetStatus() *string {
	return s.Status
}

func (s *CreatePrometheusMonitoringRequest) GetType() *string {
	return s.Type
}

func (s *CreatePrometheusMonitoringRequest) SetClusterId(v string) *CreatePrometheusMonitoringRequest {
	s.ClusterId = &v
	return s
}

func (s *CreatePrometheusMonitoringRequest) SetConfigYaml(v string) *CreatePrometheusMonitoringRequest {
	s.ConfigYaml = &v
	return s
}

func (s *CreatePrometheusMonitoringRequest) SetRegionId(v string) *CreatePrometheusMonitoringRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePrometheusMonitoringRequest) SetStatus(v string) *CreatePrometheusMonitoringRequest {
	s.Status = &v
	return s
}

func (s *CreatePrometheusMonitoringRequest) SetType(v string) *CreatePrometheusMonitoringRequest {
	s.Type = &v
	return s
}

func (s *CreatePrometheusMonitoringRequest) Validate() error {
	return dara.Validate(s)
}
