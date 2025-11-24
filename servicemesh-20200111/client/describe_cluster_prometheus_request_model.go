// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterPrometheusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetK8sClusterId(v string) *DescribeClusterPrometheusRequest
	GetK8sClusterId() *string
	SetK8sClusterRegionId(v string) *DescribeClusterPrometheusRequest
	GetK8sClusterRegionId() *string
	SetServiceMeshId(v string) *DescribeClusterPrometheusRequest
	GetServiceMeshId() *string
}

type DescribeClusterPrometheusRequest struct {
	// The ID of the cluster on the data plane.
	//
	// example:
	//
	// ce3c25e247da24f3aab9b7edfae83****
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// The ID of the region where the cluster on the data plane resides.
	//
	// example:
	//
	// cn-hangzhou
	K8sClusterRegionId *string `json:"K8sClusterRegionId,omitempty" xml:"K8sClusterRegionId,omitempty"`
	// The ASM instance ID.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeClusterPrometheusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterPrometheusRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterPrometheusRequest) GetK8sClusterId() *string {
	return s.K8sClusterId
}

func (s *DescribeClusterPrometheusRequest) GetK8sClusterRegionId() *string {
	return s.K8sClusterRegionId
}

func (s *DescribeClusterPrometheusRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeClusterPrometheusRequest) SetK8sClusterId(v string) *DescribeClusterPrometheusRequest {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeClusterPrometheusRequest) SetK8sClusterRegionId(v string) *DescribeClusterPrometheusRequest {
	s.K8sClusterRegionId = &v
	return s
}

func (s *DescribeClusterPrometheusRequest) SetServiceMeshId(v string) *DescribeClusterPrometheusRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeClusterPrometheusRequest) Validate() error {
	return dara.Validate(s)
}
