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
	NetworkMode     *string   `json:"networkMode,omitempty" xml:"networkMode,omitempty"`
	SecurityGroupID *string   `json:"securityGroupID,omitempty" xml:"securityGroupID,omitempty"`
	VpcID           *string   `json:"vpcID,omitempty" xml:"vpcID,omitempty"`
	VswitchIDs      []*string `json:"vswitchIDs,omitempty" xml:"vswitchIDs,omitempty" type:"Repeated"`
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
