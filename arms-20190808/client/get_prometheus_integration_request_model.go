// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusIntegrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetPrometheusIntegrationRequest
	GetClusterId() *string
	SetInstanceId(v int64) *GetPrometheusIntegrationRequest
	GetInstanceId() *int64
	SetIntegrationType(v string) *GetPrometheusIntegrationRequest
	GetIntegrationType() *string
	SetRegionId(v string) *GetPrometheusIntegrationRequest
	GetRegionId() *string
}

type GetPrometheusIntegrationRequest struct {
	// The ID of the Prometheus instance. Valid values: aliyun-cs and ecs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the exporter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2893
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The integration type. Valid values: kafka, mysql, redis, snmp, emr, nubela, and tidb.
	//
	// This parameter is required.
	//
	// example:
	//
	// kafka, mysql, redis, snmp, emr, nubela, and tidb
	IntegrationType *string `json:"IntegrationType,omitempty" xml:"IntegrationType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPrometheusIntegrationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusIntegrationRequest) GoString() string {
	return s.String()
}

func (s *GetPrometheusIntegrationRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetPrometheusIntegrationRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetPrometheusIntegrationRequest) GetIntegrationType() *string {
	return s.IntegrationType
}

func (s *GetPrometheusIntegrationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPrometheusIntegrationRequest) SetClusterId(v string) *GetPrometheusIntegrationRequest {
	s.ClusterId = &v
	return s
}

func (s *GetPrometheusIntegrationRequest) SetInstanceId(v int64) *GetPrometheusIntegrationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPrometheusIntegrationRequest) SetIntegrationType(v string) *GetPrometheusIntegrationRequest {
	s.IntegrationType = &v
	return s
}

func (s *GetPrometheusIntegrationRequest) SetRegionId(v string) *GetPrometheusIntegrationRequest {
	s.RegionId = &v
	return s
}

func (s *GetPrometheusIntegrationRequest) Validate() error {
	return dara.Validate(s)
}
