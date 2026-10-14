// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeSecurityGroupEgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestCidrIp(v string) *RevokeSecurityGroupEgressRequest
	GetDestCidrIp() *string
	SetIpProtocol(v string) *RevokeSecurityGroupEgressRequest
	GetIpProtocol() *string
	SetPolicy(v string) *RevokeSecurityGroupEgressRequest
	GetPolicy() *string
	SetPortRange(v string) *RevokeSecurityGroupEgressRequest
	GetPortRange() *string
	SetPriority(v int32) *RevokeSecurityGroupEgressRequest
	GetPriority() *int32
	SetSecurityGroupId(v string) *RevokeSecurityGroupEgressRequest
	GetSecurityGroupId() *string
	SetSourcePortRange(v string) *RevokeSecurityGroupEgressRequest
	GetSourcePortRange() *string
}

type RevokeSecurityGroupEgressRequest struct {
	// The destination IP address range. CIDR blocks and IPv4 address ranges are supported.
	//
	// Default value: null.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/8
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// The transport layer protocol. The value is case-sensitive. Valid values:
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
	// - **accept**: accepts access.
	//
	// - **drop**: deny access without returning any denial information.
	//
	// Default value: **accept**.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of destination ports that correspond to the transport layer protocol. Valid values:
	//
	// - TCP/UDP protocol: Valid values are **1*	- to **65535**. Separate the start port and the end port with a forward slash (/). Correct example: **1/200**. Incorrect example: **200/1**.
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
	// The priority of the security group rule. Valid values: **1*	- to **100**. A smaller value indicates a higher priority.
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
	// The range of source ports that correspond to the transport layer protocol. Valid values:
	//
	// - TCP/UDP protocol: Valid values are **1*	- to **65535**. Separate the start port and the end port with a forward slash (/). Correct example: **1/200**. Incorrect example: **200/1**.
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

func (s RevokeSecurityGroupEgressRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeSecurityGroupEgressRequest) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupEgressRequest) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *RevokeSecurityGroupEgressRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *RevokeSecurityGroupEgressRequest) GetPolicy() *string {
	return s.Policy
}

func (s *RevokeSecurityGroupEgressRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *RevokeSecurityGroupEgressRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *RevokeSecurityGroupEgressRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *RevokeSecurityGroupEgressRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *RevokeSecurityGroupEgressRequest) SetDestCidrIp(v string) *RevokeSecurityGroupEgressRequest {
	s.DestCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetIpProtocol(v string) *RevokeSecurityGroupEgressRequest {
	s.IpProtocol = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetPolicy(v string) *RevokeSecurityGroupEgressRequest {
	s.Policy = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetPortRange(v string) *RevokeSecurityGroupEgressRequest {
	s.PortRange = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetPriority(v int32) *RevokeSecurityGroupEgressRequest {
	s.Priority = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetSecurityGroupId(v string) *RevokeSecurityGroupEgressRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetSourcePortRange(v string) *RevokeSecurityGroupEgressRequest {
	s.SourcePortRange = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) Validate() error {
	return dara.Validate(s)
}
