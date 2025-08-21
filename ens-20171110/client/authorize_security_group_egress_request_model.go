// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSecurityGroupEgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestCidrIp(v string) *AuthorizeSecurityGroupEgressRequest
	GetDestCidrIp() *string
	SetIpProtocol(v string) *AuthorizeSecurityGroupEgressRequest
	GetIpProtocol() *string
	SetPolicy(v string) *AuthorizeSecurityGroupEgressRequest
	GetPolicy() *string
	SetPortRange(v string) *AuthorizeSecurityGroupEgressRequest
	GetPortRange() *string
	SetPriority(v int32) *AuthorizeSecurityGroupEgressRequest
	GetPriority() *int32
	SetSecurityGroupId(v string) *AuthorizeSecurityGroupEgressRequest
	GetSecurityGroupId() *string
	SetSourcePortRange(v string) *AuthorizeSecurityGroupEgressRequest
	GetSourcePortRange() *string
}

type AuthorizeSecurityGroupEgressRequest struct {
	// The destination IP addresses. CIDR blocks and IPv4 addresses are supported.
	//
	// This parameter is empty by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.XX.XX/8
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
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
	// 	- all: All protocols are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The action of the security group rule. Valid values:
	//
	// 	- accept: allows access. This is the default value.
	//
	// 	- drop: denies access and does not return responses.
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
	// The priority of security group rule N. Valid values: **1 to 100**. Default value: **1**.
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
	// The range of port numbers that correspond to the transport layer protocol for the source security group. Valid values:
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

func (s AuthorizeSecurityGroupEgressRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSecurityGroupEgressRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupEgressRequest) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *AuthorizeSecurityGroupEgressRequest) GetPolicy() *string {
	return s.Policy
}

func (s *AuthorizeSecurityGroupEgressRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *AuthorizeSecurityGroupEgressRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *AuthorizeSecurityGroupEgressRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *AuthorizeSecurityGroupEgressRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *AuthorizeSecurityGroupEgressRequest) SetDestCidrIp(v string) *AuthorizeSecurityGroupEgressRequest {
	s.DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetIpProtocol(v string) *AuthorizeSecurityGroupEgressRequest {
	s.IpProtocol = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPolicy(v string) *AuthorizeSecurityGroupEgressRequest {
	s.Policy = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPortRange(v string) *AuthorizeSecurityGroupEgressRequest {
	s.PortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPriority(v int32) *AuthorizeSecurityGroupEgressRequest {
	s.Priority = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetSecurityGroupId(v string) *AuthorizeSecurityGroupEgressRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetSourcePortRange(v string) *AuthorizeSecurityGroupEgressRequest {
	s.SourcePortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) Validate() error {
	return dara.Validate(s)
}
