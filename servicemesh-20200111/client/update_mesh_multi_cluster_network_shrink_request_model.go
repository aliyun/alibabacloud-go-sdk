// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeshMultiClusterNetworkShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMultiClusterNetworksShrink(v string) *UpdateMeshMultiClusterNetworkShrinkRequest
	GetMultiClusterNetworksShrink() *string
	SetServiceMeshId(v string) *UpdateMeshMultiClusterNetworkShrinkRequest
	GetServiceMeshId() *string
}

type UpdateMeshMultiClusterNetworkShrinkRequest struct {
	// The network configuration descriptions of multiple Kubernetes clusters. The key in the map is the ID of a Kubernetes cluster, and the value is the network configuration of the cluster.
	MultiClusterNetworksShrink *string `json:"MultiClusterNetworks,omitempty" xml:"MultiClusterNetworks,omitempty"`
	// The ID of the ASM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s UpdateMeshMultiClusterNetworkShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeshMultiClusterNetworkShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMeshMultiClusterNetworkShrinkRequest) GetMultiClusterNetworksShrink() *string {
	return s.MultiClusterNetworksShrink
}

func (s *UpdateMeshMultiClusterNetworkShrinkRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateMeshMultiClusterNetworkShrinkRequest) SetMultiClusterNetworksShrink(v string) *UpdateMeshMultiClusterNetworkShrinkRequest {
	s.MultiClusterNetworksShrink = &v
	return s
}

func (s *UpdateMeshMultiClusterNetworkShrinkRequest) SetServiceMeshId(v string) *UpdateMeshMultiClusterNetworkShrinkRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateMeshMultiClusterNetworkShrinkRequest) Validate() error {
	return dara.Validate(s)
}
