// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusIntegrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListPrometheusIntegrationRequest
	GetClusterId() *string
	SetIntegrationType(v string) *ListPrometheusIntegrationRequest
	GetIntegrationType() *string
	SetRegionId(v string) *ListPrometheusIntegrationRequest
	GetRegionId() *string
}

type ListPrometheusIntegrationRequest struct {
	// The ID of the Prometheus instance. Only aliyun-cs and ecs instances are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// c77f6f2397ea74672872acf5e31374a27
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

func (s ListPrometheusIntegrationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusIntegrationRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusIntegrationRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListPrometheusIntegrationRequest) GetIntegrationType() *string {
	return s.IntegrationType
}

func (s *ListPrometheusIntegrationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrometheusIntegrationRequest) SetClusterId(v string) *ListPrometheusIntegrationRequest {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusIntegrationRequest) SetIntegrationType(v string) *ListPrometheusIntegrationRequest {
	s.IntegrationType = &v
	return s
}

func (s *ListPrometheusIntegrationRequest) SetRegionId(v string) *ListPrometheusIntegrationRequest {
	s.RegionId = &v
	return s
}

func (s *ListPrometheusIntegrationRequest) Validate() error {
	return dara.Validate(s)
}
