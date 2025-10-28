// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEcsNotInClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkMode(v int32) *ListEcsNotInClusterRequest
	GetNetworkMode() *int32
	SetVpcId(v string) *ListEcsNotInClusterRequest
	GetVpcId() *string
}

type ListEcsNotInClusterRequest struct {
	// The network type. Valid values:
	//
	// 	- 1: classic network
	//
	// 	- 2: virtual private cloud (VPC)
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	NetworkMode *int32 `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// The ID of the VPC. This parameter is required if the NetworkMode parameter is set to 2.
	//
	// example:
	//
	// vpc-2zef6ob8****v8x3q46kp
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListEcsNotInClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEcsNotInClusterRequest) GoString() string {
	return s.String()
}

func (s *ListEcsNotInClusterRequest) GetNetworkMode() *int32 {
	return s.NetworkMode
}

func (s *ListEcsNotInClusterRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEcsNotInClusterRequest) SetNetworkMode(v int32) *ListEcsNotInClusterRequest {
	s.NetworkMode = &v
	return s
}

func (s *ListEcsNotInClusterRequest) SetVpcId(v string) *ListEcsNotInClusterRequest {
	s.VpcId = &v
	return s
}

func (s *ListEcsNotInClusterRequest) Validate() error {
	return dara.Validate(s)
}
