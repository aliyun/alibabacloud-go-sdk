// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpProtocol(v string) *RevokeSecurityGroupRequest
	GetIpProtocol() *string
	SetPolicy(v string) *RevokeSecurityGroupRequest
	GetPolicy() *string
	SetPortRange(v string) *RevokeSecurityGroupRequest
	GetPortRange() *string
	SetPriority(v int32) *RevokeSecurityGroupRequest
	GetPriority() *int32
	SetSecurityGroupId(v string) *RevokeSecurityGroupRequest
	GetSecurityGroupId() *string
	SetSourceCidrIp(v string) *RevokeSecurityGroupRequest
	GetSourceCidrIp() *string
	SetSourcePortRange(v string) *RevokeSecurityGroupRequest
	GetSourcePortRange() *string
}

type RevokeSecurityGroupRequest struct {
	// The transport layer protocol. The value of this parameter is case-sensitive. Valid values:
	//
	// 	- tcp
	//
	// 	- udp
	//
	// 	- icmp
	//
	// 	- gre
	//
	// 	- all: all protocols.
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The authorization policy. Valid values:
	//
	// 	- accept: allows access. This is the default value.
	//
	// 	- drop: denies access and does not return responses.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of destination ports that correspond to the transport layer protocol for the security group rule. Valid values:
	//
	// 	- When the IpProtocol parameter is set to tcp or udp, the port number range is **1*	- to **65535**. The start port number and the end port number are separated by a forward slash (/). Correct example: **1/200**. Incorrect example: **200/1**.
	//
	// 	- When the IpProtocol parameter is set to icmp, the port number range is **-1/-1**, which indicates all ports.
	//
	// 	- When the IpProtocol parameter is set to gre, the port number range is **-1/-1**, which indicates all ports.
	//
	// 	- When the IpProtocol parameter is set to all, the port number range is **-1/-1**, which indicates all ports.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22/22
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the security group rule. Valid values: **1*	- to **100**. Default value: **1**.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The source CIDR block. CIDR blocks and IPv4 addresses are supported. Default value: 0.0.XX.XX/0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.XX.XX/8
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The range of source ports that correspond to the transport layer protocol for the security group rule. Valid values:
	//
	// 	- When the IpProtocol parameter is set to tcp or udp, the port number range is **1*	- to **65535**. The start port number and the end port number are separated by a forward slash (/). Correct example: **1/200**. Incorrect example: **200/1**.
	//
	// 	- When the IpProtocol parameter is set to icmp, the port number range is **-1/-1**, which indicates all ports.
	//
	// 	- When the IpProtocol parameter is set to gre, the port number range is **-1/-1**, which indicates all ports.
	//
	// 	- When the IpProtocol parameter is set to all, the port number range is **-1/-1**, which indicates all ports.
	//
	// example:
	//
	// 22/22
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s RevokeSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *RevokeSecurityGroupRequest) GetPolicy() *string {
	return s.Policy
}

func (s *RevokeSecurityGroupRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *RevokeSecurityGroupRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *RevokeSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *RevokeSecurityGroupRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *RevokeSecurityGroupRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *RevokeSecurityGroupRequest) SetIpProtocol(v string) *RevokeSecurityGroupRequest {
	s.IpProtocol = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetPolicy(v string) *RevokeSecurityGroupRequest {
	s.Policy = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetPortRange(v string) *RevokeSecurityGroupRequest {
	s.PortRange = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetPriority(v int32) *RevokeSecurityGroupRequest {
	s.Priority = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSecurityGroupId(v string) *RevokeSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSourceCidrIp(v string) *RevokeSecurityGroupRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSourcePortRange(v string) *RevokeSecurityGroupRequest {
	s.SourcePortRange = &v
	return s
}

func (s *RevokeSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
