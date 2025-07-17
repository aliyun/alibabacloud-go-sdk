// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusMonitoringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdatePrometheusMonitoringRequest
	GetClusterId() *string
	SetConfigYaml(v string) *UpdatePrometheusMonitoringRequest
	GetConfigYaml() *string
	SetMonitoringName(v string) *UpdatePrometheusMonitoringRequest
	GetMonitoringName() *string
	SetRegionId(v string) *UpdatePrometheusMonitoringRequest
	GetRegionId() *string
	SetType(v string) *UpdatePrometheusMonitoringRequest
	GetType() *string
}

type UpdatePrometheusMonitoringRequest struct {
	// The ID of the Prometheus instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The monitoring configuration. The value is a YAML string.
	//
	// This parameter is required.
	//
	// example:
	//
	// apiVersion: monitoring.coreos.com/v1
	//
	// kind: ServiceMonitor
	//
	// metadata:
	//
	//   name: tomcat-demo
	//
	//   namespace: default
	//
	// spec:
	//
	//   endpoints:
	//
	//     - interval: 30s
	//
	//       path: /metrics
	//
	//       port: tomcat-monitor
	//
	//   namespaceSelector:
	//
	//     any: true
	//
	//   selector:
	//
	//     matchLabels:
	//
	//       app: tomcat
	ConfigYaml *string `json:"ConfigYaml,omitempty" xml:"ConfigYaml,omitempty"`
	// The name of the monitoring configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// podMonitor1
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
	// podMonitor
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdatePrometheusMonitoringRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusMonitoringRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusMonitoringRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdatePrometheusMonitoringRequest) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *UpdatePrometheusMonitoringRequest) GetMonitoringName() *string {
	return s.MonitoringName
}

func (s *UpdatePrometheusMonitoringRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePrometheusMonitoringRequest) GetType() *string {
	return s.Type
}

func (s *UpdatePrometheusMonitoringRequest) SetClusterId(v string) *UpdatePrometheusMonitoringRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdatePrometheusMonitoringRequest) SetConfigYaml(v string) *UpdatePrometheusMonitoringRequest {
	s.ConfigYaml = &v
	return s
}

func (s *UpdatePrometheusMonitoringRequest) SetMonitoringName(v string) *UpdatePrometheusMonitoringRequest {
	s.MonitoringName = &v
	return s
}

func (s *UpdatePrometheusMonitoringRequest) SetRegionId(v string) *UpdatePrometheusMonitoringRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePrometheusMonitoringRequest) SetType(v string) *UpdatePrometheusMonitoringRequest {
	s.Type = &v
	return s
}

func (s *UpdatePrometheusMonitoringRequest) Validate() error {
	return dara.Validate(s)
}
