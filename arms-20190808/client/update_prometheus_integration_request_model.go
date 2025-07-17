// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusIntegrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdatePrometheusIntegrationRequest
	GetClusterId() *string
	SetInstanceId(v int64) *UpdatePrometheusIntegrationRequest
	GetInstanceId() *int64
	SetIntegrationType(v string) *UpdatePrometheusIntegrationRequest
	GetIntegrationType() *string
	SetParam(v string) *UpdatePrometheusIntegrationRequest
	GetParam() *string
	SetRegionId(v string) *UpdatePrometheusIntegrationRequest
	GetRegionId() *string
}

type UpdatePrometheusIntegrationRequest struct {
	// The ID of the Prometheus instance. Only a Prometheus instance for Container Service or a Prometheus instance for ECS is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The exporter ID.
	//
	// example:
	//
	// 2893
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the integration.
	//
	// This parameter is required.
	//
	// example:
	//
	// kafka, mysql, redis, snmp, emr, nubela, and tidb
	IntegrationType *string `json:"IntegrationType,omitempty" xml:"IntegrationType,omitempty"`
	// The configurations of the exporter. The value is a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//       "port": "5554",
	//
	//       "name": "kafka-test12",
	//
	//       "kafka_instance": "kafka-test",
	//
	//       "__label_value": "kafka-test",
	//
	//       "scrape_interval": 33,
	//
	//       "metrics_path": "/metrics",
	//
	//       "__label_key": "kafka-test"
	//
	// }
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdatePrometheusIntegrationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusIntegrationRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusIntegrationRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdatePrometheusIntegrationRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *UpdatePrometheusIntegrationRequest) GetIntegrationType() *string {
	return s.IntegrationType
}

func (s *UpdatePrometheusIntegrationRequest) GetParam() *string {
	return s.Param
}

func (s *UpdatePrometheusIntegrationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePrometheusIntegrationRequest) SetClusterId(v string) *UpdatePrometheusIntegrationRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdatePrometheusIntegrationRequest) SetInstanceId(v int64) *UpdatePrometheusIntegrationRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdatePrometheusIntegrationRequest) SetIntegrationType(v string) *UpdatePrometheusIntegrationRequest {
	s.IntegrationType = &v
	return s
}

func (s *UpdatePrometheusIntegrationRequest) SetParam(v string) *UpdatePrometheusIntegrationRequest {
	s.Param = &v
	return s
}

func (s *UpdatePrometheusIntegrationRequest) SetRegionId(v string) *UpdatePrometheusIntegrationRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePrometheusIntegrationRequest) Validate() error {
	return dara.Validate(s)
}
