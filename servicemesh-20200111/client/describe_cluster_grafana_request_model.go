// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterGrafanaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetK8sClusterId(v string) *DescribeClusterGrafanaRequest
	GetK8sClusterId() *string
	SetReAddPrometheusIntegration(v string) *DescribeClusterGrafanaRequest
	GetReAddPrometheusIntegration() *string
	SetServiceMeshId(v string) *DescribeClusterGrafanaRequest
	GetServiceMeshId() *string
}

type DescribeClusterGrafanaRequest struct {
	// The ID of the cluster on the data plane.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce3c25e247da24f3aab9b7edfae83****
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// Specifies whether to integrate Managed Service for Prometheus for the cluster on the data plane.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// false
	ReAddPrometheusIntegration *string `json:"ReAddPrometheusIntegration,omitempty" xml:"ReAddPrometheusIntegration,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeClusterGrafanaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterGrafanaRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterGrafanaRequest) GetK8sClusterId() *string {
	return s.K8sClusterId
}

func (s *DescribeClusterGrafanaRequest) GetReAddPrometheusIntegration() *string {
	return s.ReAddPrometheusIntegration
}

func (s *DescribeClusterGrafanaRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeClusterGrafanaRequest) SetK8sClusterId(v string) *DescribeClusterGrafanaRequest {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeClusterGrafanaRequest) SetReAddPrometheusIntegration(v string) *DescribeClusterGrafanaRequest {
	s.ReAddPrometheusIntegration = &v
	return s
}

func (s *DescribeClusterGrafanaRequest) SetServiceMeshId(v string) *DescribeClusterGrafanaRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeClusterGrafanaRequest) Validate() error {
	return dara.Validate(s)
}
