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
	// The transport-layer protocol. The value is case-sensitive. Valid values:
	//
	// - tcp
	//
	// - udp
	//
	// - icmp
	//
	// - gre
	//
	// - all: all protocols are supported
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// Settings for access permissions. Valid values:
	//
	// - accept: accepts access.
	//
	// - drop: deny access without returning a rejection message.
	//
	// Default value: accept.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of port numbers for the transport-layer protocol of the security group. Valid values:
	//
	// - TCP/UDP protocol: valid values are **1*	- to **65535**. Separate the start port and the end port with a forward slash (/). Correct example: **1/200**. Incorrect example: **200/1**.
	//
	// - ICMP protocol: **-1/-1**.
	//
	// - GRE protocol: **-1/-1**.
	//
	// - If IpProtocol is set to all: **-1/-1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22/22
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the security group rule. Valid values: **1*	- to **100**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The security group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4ph***
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The source IP address range. CIDR blocks and IPv4 addresses are supported.
	//
	// Default value: none.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.XX.XX/8
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The range of port numbers for the transport-layer protocol of the source security group. Valid values:
	//
	// - TCP/UDP protocol: valid values are **1*	- to **65535**. Separate the start port and the end port with a forward slash (/). Correct example: **1/200**. Incorrect example: **200/1**.
	//
	// - ICMP protocol: **-1/-1**.
	//
	// - GRE protocol: **-1/-1**.
	//
	// - If IpProtocol is set to all: **-1/-1**.
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
