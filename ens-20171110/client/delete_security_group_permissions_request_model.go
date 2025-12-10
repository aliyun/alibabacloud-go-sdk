// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityGroupPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPermissions(v []*DeleteSecurityGroupPermissionsRequestPermissions) *DeleteSecurityGroupPermissionsRequest
	GetPermissions() []*DeleteSecurityGroupPermissionsRequestPermissions
	SetSecurityGroupId(v string) *DeleteSecurityGroupPermissionsRequest
	GetSecurityGroupId() *string
}

type DeleteSecurityGroupPermissionsRequest struct {
	// The security group rules.
	//
	// This parameter is required.
	Permissions []*DeleteSecurityGroupPermissionsRequestPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The ID of the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DeleteSecurityGroupPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityGroupPermissionsRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupPermissionsRequest) GetPermissions() []*DeleteSecurityGroupPermissionsRequestPermissions {
	return s.Permissions
}

func (s *DeleteSecurityGroupPermissionsRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeleteSecurityGroupPermissionsRequest) SetPermissions(v []*DeleteSecurityGroupPermissionsRequestPermissions) *DeleteSecurityGroupPermissionsRequest {
	s.Permissions = v
	return s
}

func (s *DeleteSecurityGroupPermissionsRequest) SetSecurityGroupId(v string) *DeleteSecurityGroupPermissionsRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DeleteSecurityGroupPermissionsRequest) Validate() error {
	if s.Permissions != nil {
		for _, item := range s.Permissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteSecurityGroupPermissionsRequestPermissions struct {
	// The destination IPv4 CIDR block. CIDR blocks and IPv4 addresses are supported.
	//
	// example:
	//
	// 10.XX.XX.91
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// The direction in which the security group rule is applied.
	//
	// 	- egress
	//
	// 	- ingress
	//
	// This parameter is required.
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// Protocol type. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// 	- ICMP: the ICMP protocol
	//
	// 	- ICMPv6: the IPv6 ICMP protocol.
	//
	// 	- ALL: All protocols are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The destination IPv6 CIDR block. IPv6 CIDR blocks and IPv6 addresses are supported.
	//
	// >  This parameter and the `DestCidrIp` parameter cannot be set at the same time.
	//
	// example:
	//
	// ::/0
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// The source IPv6 CIDR block of the security group rule. IPv6 CIDR blocks and IPv6 addresses are supported.
	//
	// >  This parameter and the `DestCidrIp` parameter cannot be set at the same time.
	//
	// example:
	//
	// ::/0
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The action specified in the security group rule. Valid values:
	//
	// 	- Accept
	//
	// 	- Drop
	//
	// This parameter is required.
	//
	// example:
	//
	// Accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The destination port range of the security group rule. Valid values:
	//
	// 	- If you set IpProtocol to TCP or UDP, the valid values of this parameter are 1 to 65535. Specify a port range in the format of \\<Start port number>/\\<End port number>. Example: 1/200.
	//
	// 	- If you set IpProtocol to ICMP, the port range is -1/-1.
	//
	// 	- If you set IpProtocol to ALL, the port number range is -1/-1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the security group rule. A smaller value specifies a higher priority. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The source IPv4 CIDR block. CIDR blocks and IPv4 addresses are supported.
	//
	// example:
	//
	// 0.XX.XX.0/0
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The range of source port numbers for the protocols specified in the security group rule. Valid values:
	//
	// 	- If you set IpProtocol to TCP or UDP, the valid values of this parameter are 1 to 65535. Specify a port range in the format of \\<Start port number>/\\<End port number>. Example: 1/200.
	//
	// 	- If you set IpProtocol to ICMP, the port range is -1/-1.
	//
	// 	- If you set IpProtocol to ALL, the port range is -1/-1.
	//
	// example:
	//
	// 22/22
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s DeleteSecurityGroupPermissionsRequestPermissions) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityGroupPermissionsRequestPermissions) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) GetDirection() *string {
	return s.Direction
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) GetPolicy() *string {
	return s.Policy
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) GetPortRange() *string {
	return s.PortRange
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) GetPriority() *int32 {
	return s.Priority
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) SetDestCidrIp(v string) *DeleteSecurityGroupPermissionsRequestPermissions {
	s.DestCidrIp = &v
	return s
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) SetDirection(v string) *DeleteSecurityGroupPermissionsRequestPermissions {
	s.Direction = &v
	return s
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) SetIpProtocol(v string) *DeleteSecurityGroupPermissionsRequestPermissions {
	s.IpProtocol = &v
	return s
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) SetIpv6DestCidrIp(v string) *DeleteSecurityGroupPermissionsRequestPermissions {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) SetIpv6SourceCidrIp(v string) *DeleteSecurityGroupPermissionsRequestPermissions {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) SetPolicy(v string) *DeleteSecurityGroupPermissionsRequestPermissions {
	s.Policy = &v
	return s
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) SetPortRange(v string) *DeleteSecurityGroupPermissionsRequestPermissions {
	s.PortRange = &v
	return s
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) SetPriority(v int32) *DeleteSecurityGroupPermissionsRequestPermissions {
	s.Priority = &v
	return s
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) SetSourceCidrIp(v string) *DeleteSecurityGroupPermissionsRequestPermissions {
	s.SourceCidrIp = &v
	return s
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) SetSourcePortRange(v string) *DeleteSecurityGroupPermissionsRequestPermissions {
	s.SourcePortRange = &v
	return s
}

func (s *DeleteSecurityGroupPermissionsRequestPermissions) Validate() error {
	return dara.Validate(s)
}
