// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterVpcEndpointConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *CreateClusterVpcEndpointConnectionRequest
	GetXDebugId() *string
	SetClusterId(v string) *CreateClusterVpcEndpointConnectionRequest
	GetClusterId() *string
	SetDryRun(v bool) *CreateClusterVpcEndpointConnectionRequest
	GetDryRun() *bool
	SetRegion(v string) *CreateClusterVpcEndpointConnectionRequest
	GetRegion() *string
	SetXSysomInvokeSource(v string) *CreateClusterVpcEndpointConnectionRequest
	GetXSysomInvokeSource() *string
}

type CreateClusterVpcEndpointConnectionRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The ID of the ACK cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// ray-bzxw7g2r7301r3f2
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The region to which the cluster belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region             *string `json:"region,omitempty" xml:"region,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s CreateClusterVpcEndpointConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterVpcEndpointConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterVpcEndpointConnectionRequest) GetXDebugId() *string {
	return s.XDebugId
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

func (s *CreateClusterVpcEndpointConnectionRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *CreateClusterVpcEndpointConnectionRequest) SetXDebugId(v string) *CreateClusterVpcEndpointConnectionRequest {
	s.XDebugId = &v
	return s
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

func (s *CreateClusterVpcEndpointConnectionRequest) SetXSysomInvokeSource(v string) *CreateClusterVpcEndpointConnectionRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *CreateClusterVpcEndpointConnectionRequest) Validate() error {
	return dara.Validate(s)
}
