// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusIntegrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeletePrometheusIntegrationRequest
	GetClusterId() *string
	SetInstanceId(v int64) *DeletePrometheusIntegrationRequest
	GetInstanceId() *int64
	SetIntegrationType(v string) *DeletePrometheusIntegrationRequest
	GetIntegrationType() *string
	SetRegionId(v string) *DeletePrometheusIntegrationRequest
	GetRegionId() *string
}

type DeletePrometheusIntegrationRequest struct {
	// The ID of the Prometheus instance. Only a Prometheus instance for Container Service or a Prometheus instance for ECS is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the exporter.
	//
	// example:
	//
	// 2875
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the integration.
	//
	// This parameter is required.
	//
	// example:
	//
	// kafka and mysql.
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

func (s DeletePrometheusIntegrationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusIntegrationRequest) GoString() string {
	return s.String()
}

func (s *DeletePrometheusIntegrationRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeletePrometheusIntegrationRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DeletePrometheusIntegrationRequest) GetIntegrationType() *string {
	return s.IntegrationType
}

func (s *DeletePrometheusIntegrationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePrometheusIntegrationRequest) SetClusterId(v string) *DeletePrometheusIntegrationRequest {
	s.ClusterId = &v
	return s
}

func (s *DeletePrometheusIntegrationRequest) SetInstanceId(v int64) *DeletePrometheusIntegrationRequest {
	s.InstanceId = &v
	return s
}

func (s *DeletePrometheusIntegrationRequest) SetIntegrationType(v string) *DeletePrometheusIntegrationRequest {
	s.IntegrationType = &v
	return s
}

func (s *DeletePrometheusIntegrationRequest) SetRegionId(v string) *DeletePrometheusIntegrationRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePrometheusIntegrationRequest) Validate() error {
	return dara.Validate(s)
}
