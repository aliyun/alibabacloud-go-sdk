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
	SetSecurityGroupId(v string) *NetworkConfiguration
	GetSecurityGroupId() *string
	SetVpcId(v string) *NetworkConfiguration
	GetVpcId() *string
	SetVswitchIds(v []*string) *NetworkConfiguration
	GetVswitchIds() []*string
}

type NetworkConfiguration struct {
	// example:
	//
	// PRIVATE
	NetworkMode *string `json:"networkMode,omitempty" xml:"networkMode,omitempty"`
	// example:
	//
	// sg-1234567890abcdef0
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// example:
	//
	// vpc-1234567890abcdef0
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// example:
	//
	// vsw-1234567890abcdef0,vsw-abcdef1234567890
	VswitchIds []*string `json:"vswitchIds,omitempty" xml:"vswitchIds,omitempty" type:"Repeated"`
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

func (s *NetworkConfiguration) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *NetworkConfiguration) GetVpcId() *string {
	return s.VpcId
}

func (s *NetworkConfiguration) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *NetworkConfiguration) SetNetworkMode(v string) *NetworkConfiguration {
	s.NetworkMode = &v
	return s
}

func (s *NetworkConfiguration) SetSecurityGroupId(v string) *NetworkConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *NetworkConfiguration) SetVpcId(v string) *NetworkConfiguration {
	s.VpcId = &v
	return s
}

func (s *NetworkConfiguration) SetVswitchIds(v []*string) *NetworkConfiguration {
	s.VswitchIds = v
	return s
}

func (s *NetworkConfiguration) Validate() error {
	return dara.Validate(s)
}
