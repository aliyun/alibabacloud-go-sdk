// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrometheusIntegrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AddPrometheusIntegrationRequest
	GetClusterId() *string
	SetIntegrationType(v string) *AddPrometheusIntegrationRequest
	GetIntegrationType() *string
	SetParam(v string) *AddPrometheusIntegrationRequest
	GetParam() *string
	SetRegionId(v string) *AddPrometheusIntegrationRequest
	GetRegionId() *string
}

type AddPrometheusIntegrationRequest struct {
	// The ID of the Prometheus instance. Only a Prometheus instance for Container Service or a Prometheus instance for ECS is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
	// {"port":"5554","name":"kafka-test12","kafka_instance":"kafka-test","__label_value":"kafka-test","scrape_interval":33,"metrics_path":"/metrics","__label_key":"kafka-test"}
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

func (s AddPrometheusIntegrationRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusIntegrationRequest) GoString() string {
	return s.String()
}

func (s *AddPrometheusIntegrationRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AddPrometheusIntegrationRequest) GetIntegrationType() *string {
	return s.IntegrationType
}

func (s *AddPrometheusIntegrationRequest) GetParam() *string {
	return s.Param
}

func (s *AddPrometheusIntegrationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddPrometheusIntegrationRequest) SetClusterId(v string) *AddPrometheusIntegrationRequest {
	s.ClusterId = &v
	return s
}

func (s *AddPrometheusIntegrationRequest) SetIntegrationType(v string) *AddPrometheusIntegrationRequest {
	s.IntegrationType = &v
	return s
}

func (s *AddPrometheusIntegrationRequest) SetParam(v string) *AddPrometheusIntegrationRequest {
	s.Param = &v
	return s
}

func (s *AddPrometheusIntegrationRequest) SetRegionId(v string) *AddPrometheusIntegrationRequest {
	s.RegionId = &v
	return s
}

func (s *AddPrometheusIntegrationRequest) Validate() error {
	return dara.Validate(s)
}
