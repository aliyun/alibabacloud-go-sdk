// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterVpcEndpointConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateClusterVpcEndpointConnectionRequest
	GetClusterId() *string
	SetDryRun(v bool) *CreateClusterVpcEndpointConnectionRequest
	GetDryRun() *bool
	SetRegion(v string) *CreateClusterVpcEndpointConnectionRequest
	GetRegion() *string
}

type CreateClusterVpcEndpointConnectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ray-bzxw7g2r7301r3f2
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s CreateClusterVpcEndpointConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterVpcEndpointConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterVpcEndpointConnectionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateClusterVpcEndpointConnectionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateClusterVpcEndpointConnectionRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateClusterVpcEndpointConnectionRequest) SetClusterId(v string) *CreateClusterVpcEndpointConnectionRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterVpcEndpointConnectionRequest) SetDryRun(v bool) *CreateClusterVpcEndpointConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *CreateClusterVpcEndpointConnectionRequest) SetRegion(v string) *CreateClusterVpcEndpointConnectionRequest {
	s.Region = &v
	return s
}

func (s *CreateClusterVpcEndpointConnectionRequest) Validate() error {
	return dara.Validate(s)
}
