// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterType(v string) *ListPrometheusInstancesRequest
	GetClusterType() *string
	SetRegionId(v string) *ListPrometheusInstancesRequest
	GetRegionId() *string
	SetShowGlobalView(v bool) *ListPrometheusInstancesRequest
	GetShowGlobalView() *bool
}

type ListPrometheusInstancesRequest struct {
	// The cluster type. If you do not specify this parameter, all cluster types are queried. Valid values:
	//
	// 	- cloud-product-prometheus: Prometheus instance for cloud services
	//
	// 	- ManagedKubernetes: ACK managed cluster
	//
	// 	- satellite: Prometheus instance for ARMS OpenTelemetry
	//
	// 	- Ask: ACK Serverless cluster
	//
	// 	- remote-write-prometheus: general-purpose Prometheus instance
	//
	// 	- cloud-monitor-cmee: Hybrid Cloud Monitoring
	//
	// 	- ExternalKubernetes: external Kubernetes cluster registered in ACK
	//
	// 	- vpc-prometheus: Prometheus instance for ECS
	//
	// 	- cloud-monitor-direct: cloud service self-monitoring
	//
	// 	- Edge Kubernetes: ACK Edge cluster
	//
	// example:
	//
	// cloud-product-prometheus
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to obtain global aggregation instances. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ShowGlobalView *bool `json:"ShowGlobalView,omitempty" xml:"ShowGlobalView,omitempty"`
}

func (s ListPrometheusInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstancesRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListPrometheusInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrometheusInstancesRequest) GetShowGlobalView() *bool {
	return s.ShowGlobalView
}

func (s *ListPrometheusInstancesRequest) SetClusterType(v string) *ListPrometheusInstancesRequest {
	s.ClusterType = &v
	return s
}

func (s *ListPrometheusInstancesRequest) SetRegionId(v string) *ListPrometheusInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPrometheusInstancesRequest) SetShowGlobalView(v bool) *ListPrometheusInstancesRequest {
	s.ShowGlobalView = &v
	return s
}

func (s *ListPrometheusInstancesRequest) Validate() error {
	return dara.Validate(s)
}
