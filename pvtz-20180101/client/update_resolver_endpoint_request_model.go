// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResolverEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *UpdateResolverEndpointRequest
	GetEndpointId() *string
	SetIpConfig(v []*UpdateResolverEndpointRequestIpConfig) *UpdateResolverEndpointRequest
	GetIpConfig() []*UpdateResolverEndpointRequestIpConfig
	SetLang(v string) *UpdateResolverEndpointRequest
	GetLang() *string
	SetName(v string) *UpdateResolverEndpointRequest
	GetName() *string
}

type UpdateResolverEndpointRequest struct {
	// The endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hr****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The source IP addresses of outbound traffic. You can add two to six IP addresses.
	//
	// >  You must add at least two source IP addresses for outbound traffic to ensure high availability. We recommend that you add two IP addresses that reside in different zones. You can add up to six source IP addresses.
	IpConfig []*UpdateResolverEndpointRequestIpConfig `json:"IpConfig,omitempty" xml:"IpConfig,omitempty" type:"Repeated"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The endpoint name.
	//
	// example:
	//
	// endpoint-test-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateResolverEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResolverEndpointRequest) GoString() string {
	return s.String()
}

func (s *UpdateResolverEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *UpdateResolverEndpointRequest) GetIpConfig() []*UpdateResolverEndpointRequestIpConfig {
	return s.IpConfig
}

func (s *UpdateResolverEndpointRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateResolverEndpointRequest) GetName() *string {
	return s.Name
}

func (s *UpdateResolverEndpointRequest) SetEndpointId(v string) *UpdateResolverEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *UpdateResolverEndpointRequest) SetIpConfig(v []*UpdateResolverEndpointRequestIpConfig) *UpdateResolverEndpointRequest {
	s.IpConfig = v
	return s
}

func (s *UpdateResolverEndpointRequest) SetLang(v string) *UpdateResolverEndpointRequest {
	s.Lang = &v
	return s
}

func (s *UpdateResolverEndpointRequest) SetName(v string) *UpdateResolverEndpointRequest {
	s.Name = &v
	return s
}

func (s *UpdateResolverEndpointRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateResolverEndpointRequestIpConfig struct {
	// The ID of the zone to which the vSwitch belongs.
	//
	// example:
	//
	// cn-hangzhou-a
	AzId *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	// The IPv4 CIDR block of the vSwitch.
	//
	// example:
	//
	// 172.16.XX.XX/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The source IP address of outbound traffic. The IP address must be within the specified CIDR block. If you leave this parameter empty, the system automatically allocates an IP address.
	//
	// example:
	//
	// 172.16.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-0jlgeyq4oazkh5xue****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s UpdateResolverEndpointRequestIpConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateResolverEndpointRequestIpConfig) GoString() string {
	return s.String()
}

func (s *UpdateResolverEndpointRequestIpConfig) GetAzId() *string {
	return s.AzId
}

func (s *UpdateResolverEndpointRequestIpConfig) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *UpdateResolverEndpointRequestIpConfig) GetIp() *string {
	return s.Ip
}

func (s *UpdateResolverEndpointRequestIpConfig) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateResolverEndpointRequestIpConfig) SetAzId(v string) *UpdateResolverEndpointRequestIpConfig {
	s.AzId = &v
	return s
}

func (s *UpdateResolverEndpointRequestIpConfig) SetCidrBlock(v string) *UpdateResolverEndpointRequestIpConfig {
	s.CidrBlock = &v
	return s
}

func (s *UpdateResolverEndpointRequestIpConfig) SetIp(v string) *UpdateResolverEndpointRequestIpConfig {
	s.Ip = &v
	return s
}

func (s *UpdateResolverEndpointRequestIpConfig) SetVSwitchId(v string) *UpdateResolverEndpointRequestIpConfig {
	s.VSwitchId = &v
	return s
}

func (s *UpdateResolverEndpointRequestIpConfig) Validate() error {
	return dara.Validate(s)
}
