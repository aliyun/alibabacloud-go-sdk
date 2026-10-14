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
	// - all: All protocols are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// Settings for access permissions. Valid values:
	//
	// - accept (default): Accepts access.
	//
	// - drop: Deny access without returning a deny message.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of destination port numbers relevant to the transport layer protocol for the security group. Valid values:
	//
	// - TCP/UDP: Valid values are **1*	- to **65535**. Separate the start port and end port with a forward slash (/). Correct example: **1/200**. Incorrect example: **200/1**.
	//
	// - ICMP: **-1/-1**.
	//
	// - GRE: **-1/-1**.
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
	// The security group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The source IP address range. CIDR format and IPv4 format are supported. Default value: 0.0.XX.XX/0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.XX.XX/8
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The range of source port numbers relevant to the transport layer protocol for the security group. Valid values:
	//
	// - TCP/UDP: Valid values are **1*	- to **65535**. Separate the start port and end port with a forward slash (/). Correct example: **1/200**. Incorrect example: **200/1**.
	//
	// - ICMP: **-1/-1**.
	//
	// - GRE: **-1/-1**.
	//
	// - If IpProtocol is set to all: **-1/-1**.
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
