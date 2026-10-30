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
	// The destination IP address range. CIDR blocks and IPv4 address ranges are supported.
	//
	// Default value: none.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.XX.XX/8
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// The transport layer protocol. The value is case-sensitive. Valid values:
	//
	// - tcp: TCP.
	//
	// - udp: UDP.
	//
	// - icmp: ICMP.
	//
	// - gre: GRE.
	//
	// - all: All protocols.
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// Settings for access permissions. Valid values:
	//
	// - accept (default): accept access.
	//
	// - drop: deny access without returning a deny message.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of ports enabled for the transport layer protocol in the security group. Valid values:
	//
	// - For TCP or UDP: **1*	- to **65535**. Use a forward slash (/) to separate the start port and the end port. Correct example: **1/200**. Incorrect example: **200/1**.
	//
	// - For ICMP: **-1/-1**.
	//
	// - For GRE: **-1/-1**.
	//
	// - If IpProtocol is set to all: **-1/-1**.
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
	// The range of ports enabled for the transport layer protocol in the source security group. Valid values:
	//
	// - For TCP or UDP: **1*	- to **65535**. Use a forward slash (/) to separate the start port and the end port. Correct example: **1/200**. Incorrect example: **200/1**.
	//
	// - For ICMP: **-1/-1**.
	//
	// - For GRE: **-1/-1**.
	//
	// - If IpProtocol is set to all: **-1/-1**.
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
