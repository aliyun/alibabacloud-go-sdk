// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNetworkConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkMode(v string) *NetworkConfiguration
	GetNetworkMode() *string
	SetSecurityGroupID(v string) *NetworkConfiguration
	GetSecurityGroupID() *string
	SetVpcID(v string) *NetworkConfiguration
	GetVpcID() *string
	SetVswitchIDs(v []*string) *NetworkConfiguration
	GetVswitchIDs() []*string
}

type NetworkConfiguration struct {
	// The network access mode. Valid values: PUBLIC (public network only), PRIVATE (VPC only), PUBLIC_AND_PRIVATE (both public network and VPC are reachable).
	//
	// example:
	//
	// PUBLIC_AND_PRIVATE
	NetworkMode *string `json:"networkMode,omitempty" xml:"networkMode,omitempty"`
	// The ID of the security group used by the sandboxed container instances.
	//
	// example:
	//
	// sg-****
	SecurityGroupID *string `json:"securityGroupID,omitempty" xml:"securityGroupID,omitempty"`
	// The ID of the VPC to which the sandboxed container instances are connected.
	//
	// example:
	//
	// vpc-****
	VpcID *string `json:"vpcID,omitempty" xml:"vpcID,omitempty"`
	// The list of vSwitch IDs to which the sandboxed container instances are connected. Configure vSwitches in multiple zones to improve availability.
	VswitchIDs []*string `json:"vswitchIDs" xml:"vswitchIDs" type:"Repeated"`
}

func (s NetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s NetworkConfiguration) GoString() string {
	return s.String()
}

func (s *NetworkConfiguration) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *NetworkConfiguration) GetSecurityGroupID() *string {
	return s.SecurityGroupID
}

func (s *NetworkConfiguration) GetVpcID() *string {
	return s.VpcID
}

func (s *NetworkConfiguration) GetVswitchIDs() []*string {
	return s.VswitchIDs
}

func (s *NetworkConfiguration) SetNetworkMode(v string) *NetworkConfiguration {
	s.NetworkMode = &v
	return s
}

func (s *NetworkConfiguration) SetSecurityGroupID(v string) *NetworkConfiguration {
	s.SecurityGroupID = &v
	return s
}

func (s *NetworkConfiguration) SetVpcID(v string) *NetworkConfiguration {
	s.VpcID = &v
	return s
}

func (s *NetworkConfiguration) SetVswitchIDs(v []*string) *NetworkConfiguration {
	s.VswitchIDs = v
	return s
}

func (s *NetworkConfiguration) Validate() error {
	return dara.Validate(s)
}
