// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityGroupPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPermissions(v []*CreateSecurityGroupPermissionsRequestPermissions) *CreateSecurityGroupPermissionsRequest
	GetPermissions() []*CreateSecurityGroupPermissionsRequestPermissions
	SetSecurityGroupId(v string) *CreateSecurityGroupPermissionsRequest
	GetSecurityGroupId() *string
}

type CreateSecurityGroupPermissionsRequest struct {
	// The security group rules.
	//
	// This parameter is required.
	Permissions []*CreateSecurityGroupPermissionsRequestPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The security group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s CreateSecurityGroupPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityGroupPermissionsRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupPermissionsRequest) GetPermissions() []*CreateSecurityGroupPermissionsRequestPermissions {
	return s.Permissions
}

func (s *CreateSecurityGroupPermissionsRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateSecurityGroupPermissionsRequest) SetPermissions(v []*CreateSecurityGroupPermissionsRequestPermissions) *CreateSecurityGroupPermissionsRequest {
	s.Permissions = v
	return s
}

func (s *CreateSecurityGroupPermissionsRequest) SetSecurityGroupId(v string) *CreateSecurityGroupPermissionsRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateSecurityGroupPermissionsRequest) Validate() error {
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

type CreateSecurityGroupPermissionsRequestPermissions struct {
	// The description. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination IPv4 CIDR block. Both CIDR format and IPv4 format IP address range are supported.
	//
	// example:
	//
	// 10.XX.XX.14/32
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// The authorization direction. Valid values:
	//
	// - egress: outbound.
	//
	// - ingress: inbound.
	//
	// This parameter is required.
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The protocol type. Valid values:
	//
	//
	//
	// - TCP: TCP protocol.
	//
	// - UDP: UDP protocol.
	//
	// - ICMP: ICMP protocol.
	//
	// - ICMPv6: ICMPv6 protocol.
	//
	// - ALL: All protocols are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The destination IPv6 CIDR block. Both CIDR format and IPv6 format IP address range are supported.
	//
	// > This parameter and the `DestCidrIp` parameter cannot be set at the same time. Settings for these two parameters are mutually exclusive.
	//
	// example:
	//
	// ::/0
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// The source IPv6 CIDR block. Both CIDR format and IPv6 format IP address range are supported.
	//
	// > This parameter and the `DestCidrIp` parameter cannot be set at the same time. Settings for these two parameters are mutually exclusive.
	//
	// example:
	//
	// ::/0
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The access permissions. Valid values:
	//
	//
	//
	// - Accept: Allows access.
	//
	// - Drop: Deny access.
	//
	// This parameter is required.
	//
	// example:
	//
	// Accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of destination ports that correspond to the protocols for the security group. Valid values:
	//
	//
	//
	// - TCP/UDP protocol: Valid values are 1 to 65535. Separate the start port and end port with a forward slash (/). Example: 1/200.
	//
	// - ICMP protocol: -1/-1.
	//
	// - ICMPv6: -1/-1.
	//
	// - ALL: -1/-1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the security group rule. A smaller value indicates a higher priority. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The source IPv4 CIDR block. Both CIDR format and IPv4 format IP address range are supported.
	//
	// example:
	//
	// 0.0.0.0/0
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The range of source ports that correspond to the protocols for the security group. Valid values:
	//
	//
	//
	// - TCP/UDP protocol: Valid values are 1 to 65535. Separate the start port and end port with a forward slash (/). Example: 1/200.
	//
	// - ICMP protocol: -1/-1.
	//
	// - ICMPv6: -1/-1.
	//
	// - ALL: -1/-1.
	//
	// example:
	//
	// 22/22
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s CreateSecurityGroupPermissionsRequestPermissions) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityGroupPermissionsRequestPermissions) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) GetDescription() *string {
	return s.Description
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) GetDirection() *string {
	return s.Direction
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) GetPolicy() *string {
	return s.Policy
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) GetPortRange() *string {
	return s.PortRange
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) SetDescription(v string) *CreateSecurityGroupPermissionsRequestPermissions {
	s.Description = &v
	return s
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) SetDestCidrIp(v string) *CreateSecurityGroupPermissionsRequestPermissions {
	s.DestCidrIp = &v
	return s
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) SetDirection(v string) *CreateSecurityGroupPermissionsRequestPermissions {
	s.Direction = &v
	return s
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) SetIpProtocol(v string) *CreateSecurityGroupPermissionsRequestPermissions {
	s.IpProtocol = &v
	return s
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) SetIpv6DestCidrIp(v string) *CreateSecurityGroupPermissionsRequestPermissions {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) SetIpv6SourceCidrIp(v string) *CreateSecurityGroupPermissionsRequestPermissions {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) SetPolicy(v string) *CreateSecurityGroupPermissionsRequestPermissions {
	s.Policy = &v
	return s
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) SetPortRange(v string) *CreateSecurityGroupPermissionsRequestPermissions {
	s.PortRange = &v
	return s
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) SetPriority(v int32) *CreateSecurityGroupPermissionsRequestPermissions {
	s.Priority = &v
	return s
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) SetSourceCidrIp(v string) *CreateSecurityGroupPermissionsRequestPermissions {
	s.SourceCidrIp = &v
	return s
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) SetSourcePortRange(v string) *CreateSecurityGroupPermissionsRequestPermissions {
	s.SourcePortRange = &v
	return s
}

func (s *CreateSecurityGroupPermissionsRequestPermissions) Validate() error {
	return dara.Validate(s)
}
