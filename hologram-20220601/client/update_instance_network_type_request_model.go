// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceNetworkTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnyTunnelToSingleTunnel(v string) *UpdateInstanceNetworkTypeRequest
	GetAnyTunnelToSingleTunnel() *string
	SetNetworkTypes(v string) *UpdateInstanceNetworkTypeRequest
	GetNetworkTypes() *string
	SetVSwitchId(v string) *UpdateInstanceNetworkTypeRequest
	GetVSwitchId() *string
	SetVpcId(v string) *UpdateInstanceNetworkTypeRequest
	GetVpcId() *string
	SetVpcOwnerId(v string) *UpdateInstanceNetworkTypeRequest
	GetVpcOwnerId() *string
	SetVpcRegionId(v string) *UpdateInstanceNetworkTypeRequest
	GetVpcRegionId() *string
}

type UpdateInstanceNetworkTypeRequest struct {
	// Specifies whether to convert AnyTunnel to SingleTunnel. This parameter is invalid for new instances. Set it to null by default.
	//
	// example:
	//
	// true
	AnyTunnelToSingleTunnel *string `json:"anyTunnelToSingleTunnel,omitempty" xml:"anyTunnelToSingleTunnel,omitempty"`
	// The list of network types to enable. Specify a combination of existing network types. The order does not matter. For example, if the Internet, internal network, and VPCSingleTunnel network types are enabled and you want to disable the Internet network type, remove the Internet type and pass "Intranet, VPCSingleTunnel".
	//
	// example:
	//
	// Internet,VPCSingleTunnel
	NetworkTypes *string `json:"networkTypes,omitempty" xml:"networkTypes,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2vccsiymtqr9aavew0vo3
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-t4netc3y5etlondfb5ra7
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The owner ID of the VPC, which is the Alibaba Cloud account ID.
	//
	// example:
	//
	// 1999365732646672
	VpcOwnerId *string `json:"vpcOwnerId,omitempty" xml:"vpcOwnerId,omitempty"`
	// The region ID of the VPC.
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"vpcRegionId,omitempty" xml:"vpcRegionId,omitempty"`
}

func (s UpdateInstanceNetworkTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceNetworkTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkTypeRequest) GetAnyTunnelToSingleTunnel() *string {
	return s.AnyTunnelToSingleTunnel
}

func (s *UpdateInstanceNetworkTypeRequest) GetNetworkTypes() *string {
	return s.NetworkTypes
}

func (s *UpdateInstanceNetworkTypeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateInstanceNetworkTypeRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateInstanceNetworkTypeRequest) GetVpcOwnerId() *string {
	return s.VpcOwnerId
}

func (s *UpdateInstanceNetworkTypeRequest) GetVpcRegionId() *string {
	return s.VpcRegionId
}

func (s *UpdateInstanceNetworkTypeRequest) SetAnyTunnelToSingleTunnel(v string) *UpdateInstanceNetworkTypeRequest {
	s.AnyTunnelToSingleTunnel = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetNetworkTypes(v string) *UpdateInstanceNetworkTypeRequest {
	s.NetworkTypes = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetVSwitchId(v string) *UpdateInstanceNetworkTypeRequest {
	s.VSwitchId = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetVpcId(v string) *UpdateInstanceNetworkTypeRequest {
	s.VpcId = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetVpcOwnerId(v string) *UpdateInstanceNetworkTypeRequest {
	s.VpcOwnerId = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetVpcRegionId(v string) *UpdateInstanceNetworkTypeRequest {
	s.VpcRegionId = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) Validate() error {
	return dara.Validate(s)
}
