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
	// Specifies whether to change the network type from AnyTunnel to SingleTunnel. This parameter is invalid for new instances. For new instances, this parameter is set to null by default.
	//
	// Valid values:
	//
	// 	- others/null: The network type is not changed from AnyTunnel to SingleTunnel.
	//
	// 	- true: The network type is changed from AnyTunnel to SingleTunnel.
	//
	// example:
	//
	// true
	AnyTunnelToSingleTunnel *string `json:"anyTunnelToSingleTunnel,omitempty" xml:"anyTunnelToSingleTunnel,omitempty"`
	// A list of network types that you want to enable. The network types are randomly ordered in the list. For example, the Internet, Intranet, and VPCSingleTunnel network types are enabled. If you want to disable the Internet type, set this parameter to Intranet,VPCSingleTunnel.
	//
	// Valid values:
	//
	// 	- VPCSingleTunnel: virtual private cloud (VPC).
	//
	// 	- Intranet: internal network.
	//
	// 	- VPCAnyTunnel: compatibility requirements. This value is not supported by new instances.
	//
	// 	- Internet: Internet.
	//
	// example:
	//
	// Internet,VPCSingleTunnel
	NetworkTypes *string `json:"networkTypes,omitempty" xml:"networkTypes,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-2vccsiymtqr9aavew0vo3
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-t4netc3y5etlondfb5ra7
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The ID of the Alibaba Cloud account to which the VPC belongs.
	//
	// example:
	//
	// 1999365732646672
	VpcOwnerId *string `json:"vpcOwnerId,omitempty" xml:"vpcOwnerId,omitempty"`
	// The region in which the VPC resides.
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
