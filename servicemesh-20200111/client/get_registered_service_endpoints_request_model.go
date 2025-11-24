// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegisteredServiceEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIds(v string) *GetRegisteredServiceEndpointsRequest
	GetClusterIds() *string
	SetName(v string) *GetRegisteredServiceEndpointsRequest
	GetName() *string
	SetNamespace(v string) *GetRegisteredServiceEndpointsRequest
	GetNamespace() *string
	SetServiceMeshId(v string) *GetRegisteredServiceEndpointsRequest
	GetServiceMeshId() *string
	SetServiceType(v string) *GetRegisteredServiceEndpointsRequest
	GetServiceType() *string
}

type GetRegisteredServiceEndpointsRequest struct {
	// The name of the registered service.
	//
	// example:
	//
	// c8b054ee8c3914d079b5ce9733328****,c58faedb8a78640d3aeb0372e4c02****
	ClusterIds *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	// The type of the registered service. Valid values:
	//
	// 	- `ServiceEntry`: indicates that the service is registered by creating a service entry.
	//
	// 	- `Kubernetes`: indicates that the service is registered on a Kubernetes cluster on the data plane.
	//
	// example:
	//
	// reviews
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IDs of clusters in the ASM instance. Separate multiple cluster IDs with commas (,).
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The endpoints of the registered service.
	//
	// example:
	//
	// Kubernetes
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s GetRegisteredServiceEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRegisteredServiceEndpointsRequest) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceEndpointsRequest) GetClusterIds() *string {
	return s.ClusterIds
}

func (s *GetRegisteredServiceEndpointsRequest) GetName() *string {
	return s.Name
}

func (s *GetRegisteredServiceEndpointsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetRegisteredServiceEndpointsRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *GetRegisteredServiceEndpointsRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetRegisteredServiceEndpointsRequest) SetClusterIds(v string) *GetRegisteredServiceEndpointsRequest {
	s.ClusterIds = &v
	return s
}

func (s *GetRegisteredServiceEndpointsRequest) SetName(v string) *GetRegisteredServiceEndpointsRequest {
	s.Name = &v
	return s
}

func (s *GetRegisteredServiceEndpointsRequest) SetNamespace(v string) *GetRegisteredServiceEndpointsRequest {
	s.Namespace = &v
	return s
}

func (s *GetRegisteredServiceEndpointsRequest) SetServiceMeshId(v string) *GetRegisteredServiceEndpointsRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *GetRegisteredServiceEndpointsRequest) SetServiceType(v string) *GetRegisteredServiceEndpointsRequest {
	s.ServiceType = &v
	return s
}

func (s *GetRegisteredServiceEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
