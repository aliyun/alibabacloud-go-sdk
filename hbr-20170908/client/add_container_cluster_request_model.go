// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContainerClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterType(v string) *AddContainerClusterRequest
	GetClusterType() *string
	SetDescription(v string) *AddContainerClusterRequest
	GetDescription() *string
	SetIdentifier(v string) *AddContainerClusterRequest
	GetIdentifier() *string
	SetName(v string) *AddContainerClusterRequest
	GetName() *string
	SetNetworkType(v string) *AddContainerClusterRequest
	GetNetworkType() *string
}

type AddContainerClusterRequest struct {
	// The type of the cluster. Only Container Service for Kubernetes (ACK) clusters are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACK
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// description ack pv backup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the cluster that you want to register.
	//
	// This parameter is required.
	//
	// example:
	//
	// cca8f35f0e0d84540b49d994511c2c87a
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// ack_pv_backup_location
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type of the cluster. Valid values:
	//
	// 	- **CLASSIC**: the classic network
	//
	// 	- **VPC**: a virtual private cloud (VPC)
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s AddContainerClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s AddContainerClusterRequest) GoString() string {
	return s.String()
}

func (s *AddContainerClusterRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *AddContainerClusterRequest) GetDescription() *string {
	return s.Description
}

func (s *AddContainerClusterRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *AddContainerClusterRequest) GetName() *string {
	return s.Name
}

func (s *AddContainerClusterRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *AddContainerClusterRequest) SetClusterType(v string) *AddContainerClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *AddContainerClusterRequest) SetDescription(v string) *AddContainerClusterRequest {
	s.Description = &v
	return s
}

func (s *AddContainerClusterRequest) SetIdentifier(v string) *AddContainerClusterRequest {
	s.Identifier = &v
	return s
}

func (s *AddContainerClusterRequest) SetName(v string) *AddContainerClusterRequest {
	s.Name = &v
	return s
}

func (s *AddContainerClusterRequest) SetNetworkType(v string) *AddContainerClusterRequest {
	s.NetworkType = &v
	return s
}

func (s *AddContainerClusterRequest) Validate() error {
	return dara.Validate(s)
}
