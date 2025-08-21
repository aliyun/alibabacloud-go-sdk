// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpProtocol(v string) *AuthorizeSecurityGroupRequest
	GetIpProtocol() *string
	SetPolicy(v string) *AuthorizeSecurityGroupRequest
	GetPolicy() *string
	SetPortRange(v string) *AuthorizeSecurityGroupRequest
	GetPortRange() *string
	SetPriority(v int32) *AuthorizeSecurityGroupRequest
	GetPriority() *int32
	SetSecurityGroupId(v string) *AuthorizeSecurityGroupRequest
	GetSecurityGroupId() *string
	SetSourceCidrIp(v string) *AuthorizeSecurityGroupRequest
	GetSourceCidrIp() *string
	SetSourcePortRange(v string) *AuthorizeSecurityGroupRequest
	GetSourcePortRange() *string
}

type AuthorizeSecurityGroupRequest struct {
	// The transport layer protocol. The values of this parameter are case-sensitive. Valid values:
	//
	// 	- tcp
	//
	// 	- udp
	//
	// 	- icmp
	//
	// 	- gre
	//
	// 	- all: All protocols are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The action of security group rule N that determines whether to allow inbound access. Valid values:
	//
	// 	- accept: allows access.
	//
	// 	- drop: denies access and returns no responses.
	//
	// Default value: accept.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of destination ports that correspond to the transport layer protocol for security group rule N. Valid values:
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
	// The priority of security group rule N. Valid values: **1*	- to **100**.
	//
	// Default value: **1**.
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
	// sg-bp67acfmxazb4ph***
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The source IPv4 CIDR block. CIDR blocks and IPv4 addresses are supported.
	//
	// This parameter is empty by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.XX.XX/8
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The range of port numbers that correspond to the transport layer protocol for the source security group. Valid values:
	//
	// 	- When the IpProtocol parameter is set to tcp or udp, the port number range is **1 to 65535**. The start port number and the end port number are separated by a forward slash (/). Correct example: **1/200**. Incorrect example: **200/1**.
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

func (s AuthorizeSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *AuthorizeSecurityGroupRequest) GetPolicy() *string {
	return s.Policy
}

func (s *AuthorizeSecurityGroupRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *AuthorizeSecurityGroupRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *AuthorizeSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *AuthorizeSecurityGroupRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *AuthorizeSecurityGroupRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *AuthorizeSecurityGroupRequest) SetIpProtocol(v string) *AuthorizeSecurityGroupRequest {
	s.IpProtocol = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPolicy(v string) *AuthorizeSecurityGroupRequest {
	s.Policy = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPortRange(v string) *AuthorizeSecurityGroupRequest {
	s.PortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPriority(v int32) *AuthorizeSecurityGroupRequest {
	s.Priority = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSecurityGroupId(v string) *AuthorizeSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourceCidrIp(v string) *AuthorizeSecurityGroupRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourcePortRange(v string) *AuthorizeSecurityGroupRequest {
	s.SourcePortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
