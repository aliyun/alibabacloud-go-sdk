// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddResolverEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpConfig(v []*AddResolverEndpointRequestIpConfig) *AddResolverEndpointRequest
	GetIpConfig() []*AddResolverEndpointRequestIpConfig
	SetLang(v string) *AddResolverEndpointRequest
	GetLang() *string
	SetName(v string) *AddResolverEndpointRequest
	GetName() *string
	SetSecurityGroupId(v string) *AddResolverEndpointRequest
	GetSecurityGroupId() *string
	SetVpcId(v string) *AddResolverEndpointRequest
	GetVpcId() *string
	SetVpcRegionId(v string) *AddResolverEndpointRequest
	GetVpcRegionId() *string
}

type AddResolverEndpointRequest struct {
	// The list of source IP addresses for outbound traffic. You must add at least two IP addresses. You can add up to six IP addresses.
	//
	// > To ensure high availability (HA), add at least two source IP addresses for the outbound endpoint. We recommend that you allocate these IP addresses in different zones. You can add a maximum of six source IP addresses.
	//
	// This parameter is required.
	IpConfig []*AddResolverEndpointRequestIpConfig `json:"IpConfig,omitempty" xml:"IpConfig,omitempty" type:"Repeated"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// Default value: en
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the endpoint. The name can be up to 20 characters long. An error is reported if the limit is exceeded.
	//
	// This parameter is required.
	//
	// example:
	//
	// endpoint-test-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the security group. The rules in the security group are applied to the outbound VPC.
	//
	// > To prevent service interruptions, you cannot change this value after you create the outbound endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-0jld3m9yq7l2cw12****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the outbound virtual private cloud (VPC). All outbound DNS query traffic from the Resolver is forwarded through this VPC.
	//
	// > To prevent service interruptions, you cannot change this value after you create the outbound endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-0jl96awrjt75ezglc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The region ID of the outbound VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
}

func (s AddResolverEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s AddResolverEndpointRequest) GoString() string {
	return s.String()
}

func (s *AddResolverEndpointRequest) GetIpConfig() []*AddResolverEndpointRequestIpConfig {
	return s.IpConfig
}

func (s *AddResolverEndpointRequest) GetLang() *string {
	return s.Lang
}

func (s *AddResolverEndpointRequest) GetName() *string {
	return s.Name
}

func (s *AddResolverEndpointRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *AddResolverEndpointRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *AddResolverEndpointRequest) GetVpcRegionId() *string {
	return s.VpcRegionId
}

func (s *AddResolverEndpointRequest) SetIpConfig(v []*AddResolverEndpointRequestIpConfig) *AddResolverEndpointRequest {
	s.IpConfig = v
	return s
}

func (s *AddResolverEndpointRequest) SetLang(v string) *AddResolverEndpointRequest {
	s.Lang = &v
	return s
}

func (s *AddResolverEndpointRequest) SetName(v string) *AddResolverEndpointRequest {
	s.Name = &v
	return s
}

func (s *AddResolverEndpointRequest) SetSecurityGroupId(v string) *AddResolverEndpointRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AddResolverEndpointRequest) SetVpcId(v string) *AddResolverEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *AddResolverEndpointRequest) SetVpcRegionId(v string) *AddResolverEndpointRequest {
	s.VpcRegionId = &v
	return s
}

func (s *AddResolverEndpointRequest) Validate() error {
	if s.IpConfig != nil {
		for _, item := range s.IpConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddResolverEndpointRequestIpConfig struct {
	// The ID of the zone where the vSwitch resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-a
	AzId *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	// The IPv4 CIDR block of the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.XX.XX/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The IP address. The IP address must be within the specified CIDR block. If you leave this parameter empty, the system automatically assigns an IP address.
	//
	// example:
	//
	// 172.16.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-0jlgeyq4oazkh5xue****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s AddResolverEndpointRequestIpConfig) String() string {
	return dara.Prettify(s)
}

func (s AddResolverEndpointRequestIpConfig) GoString() string {
	return s.String()
}

func (s *AddResolverEndpointRequestIpConfig) GetAzId() *string {
	return s.AzId
}

func (s *AddResolverEndpointRequestIpConfig) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *AddResolverEndpointRequestIpConfig) GetIp() *string {
	return s.Ip
}

func (s *AddResolverEndpointRequestIpConfig) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *AddResolverEndpointRequestIpConfig) SetAzId(v string) *AddResolverEndpointRequestIpConfig {
	s.AzId = &v
	return s
}

func (s *AddResolverEndpointRequestIpConfig) SetCidrBlock(v string) *AddResolverEndpointRequestIpConfig {
	s.CidrBlock = &v
	return s
}

func (s *AddResolverEndpointRequestIpConfig) SetIp(v string) *AddResolverEndpointRequestIpConfig {
	s.Ip = &v
	return s
}

func (s *AddResolverEndpointRequestIpConfig) SetVSwitchId(v string) *AddResolverEndpointRequestIpConfig {
	s.VSwitchId = &v
	return s
}

func (s *AddResolverEndpointRequestIpConfig) Validate() error {
	return dara.Validate(s)
}
