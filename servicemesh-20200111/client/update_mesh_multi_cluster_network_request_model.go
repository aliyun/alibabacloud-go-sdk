// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeshMultiClusterNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMultiClusterNetworks(v map[string]*MultiClusterNetworksValue) *UpdateMeshMultiClusterNetworkRequest
	GetMultiClusterNetworks() map[string]*MultiClusterNetworksValue
	SetServiceMeshId(v string) *UpdateMeshMultiClusterNetworkRequest
	GetServiceMeshId() *string
}

type UpdateMeshMultiClusterNetworkRequest struct {
	// The network configuration descriptions of multiple Kubernetes clusters. The key in the map is the ID of a Kubernetes cluster, and the value is the network configuration of the cluster.
	MultiClusterNetworks map[string]*MultiClusterNetworksValue `json:"MultiClusterNetworks,omitempty" xml:"MultiClusterNetworks,omitempty"`
	// The ID of the ASM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s UpdateMeshMultiClusterNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeshMultiClusterNetworkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMeshMultiClusterNetworkRequest) GetMultiClusterNetworks() map[string]*MultiClusterNetworksValue {
	return s.MultiClusterNetworks
}

func (s *UpdateMeshMultiClusterNetworkRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateMeshMultiClusterNetworkRequest) SetMultiClusterNetworks(v map[string]*MultiClusterNetworksValue) *UpdateMeshMultiClusterNetworkRequest {
	s.MultiClusterNetworks = v
	return s
}

func (s *UpdateMeshMultiClusterNetworkRequest) SetServiceMeshId(v string) *UpdateMeshMultiClusterNetworkRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateMeshMultiClusterNetworkRequest) Validate() error {
	return dara.Validate(s)
}
