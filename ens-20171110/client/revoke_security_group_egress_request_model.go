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
	// The destination IP addresses. CIDR blocks and IPv4 addresses are supported.
	//
	// By default, this parameter is empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/8
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
	// 	- **accept**: allows access.
	//
	// 	- **drop**: denies access and returns no responses.
	//
	// Default value: **accept**.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of destination ports that correspond to the transport layer protocol for the security group rule. Valid values:
	//
	// 	- If you set the IpProtocol parameter to tcp or udp, the port number ranges from **1*	- to **65535**. The start port number and the end port number are separated by a forward slash (/). Correct example: **1/200**. Incorrect example: **200/1**.
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
	// The priority of the security group rule. Valid values: **1*	- to **100**. A smaller value indicates a higher priority.
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
	// The range of port numbers that correspond to the transport layer protocol for the source security group. Valid values:
	//
	// 	- If you set the IpProtocol parameter to tcp or udp, the port number ranges from **1*	- to **65535**. The start port number and the end port number are separated by a forward slash (/). Correct example: **1/200**. Incorrect example: **200/1**.
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
