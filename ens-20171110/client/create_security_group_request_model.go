// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSecurityGroupRequest
	GetDescription() *string
	SetPermissions(v []*CreateSecurityGroupRequestPermissions) *CreateSecurityGroupRequest
	GetPermissions() []*CreateSecurityGroupRequestPermissions
	SetSecurityGroupName(v string) *CreateSecurityGroupRequest
	GetSecurityGroupName() *string
}

type CreateSecurityGroupRequest struct {
	// The description of the security group. The description must be 2 to 256 characters in length. It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The rules of the security group. Valid values of N: 0 to 200.
	Permissions []*CreateSecurityGroupRequestPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The name of the security group. The name must be 2 to 128 characters in length, and can contain digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`. By default, this parameter is empty.
	//
	// example:
	//
	// Dcdn1:2_3-4
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s CreateSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSecurityGroupRequest) GetPermissions() []*CreateSecurityGroupRequestPermissions {
	return s.Permissions
}

func (s *CreateSecurityGroupRequest) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *CreateSecurityGroupRequest) SetDescription(v string) *CreateSecurityGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateSecurityGroupRequest) SetPermissions(v []*CreateSecurityGroupRequestPermissions) *CreateSecurityGroupRequest {
	s.Permissions = v
	return s
}

func (s *CreateSecurityGroupRequest) SetSecurityGroupName(v string) *CreateSecurityGroupRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *CreateSecurityGroupRequest) Validate() error {
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

type CreateSecurityGroupRequestPermissions struct {
	// The description of the security group rule. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination IPv4 CIDR block. CIDR blocks and IPv4 addresses are supported.
	//
	// example:
	//
	// 0.0.0.0/0
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// The direction of the security group rule. Valid values:
	//
	// - egress: outbound
	//
	// - ingress: inbound
	//
	// This parameter is required.
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The transport layer protocol. Valid values:
	//
	// - TCP
	//
	// - UDP
	//
	// - ICMP
	//
	// - ALL
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	IpProtocol       *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Ipv6DestCidrIp   *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The policy of the security group rule. Valid values:
	//
	// - Accept: allows access.
	//
	// - Drop: denies access.
	//
	// This parameter is required.
	//
	// example:
	//
	// Accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of the destination ports for the security group rule. Valid values:
	//
	// - If you set the IpProtocol to TCP or UDP, the port range is 1 to 65535. You must use a forward slash (/) to separate the start port and the end port, for example, 1/200.
	//
	// - If you set the IpProtocol to ICMP, the port range is -1/-1.
	//
	// - If you set the IpProtocol to ALL, the port range is -1/-1.
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
	// The source IPv4 CIDR block. CIDR blocks and IPv4 addresses are supported.
	//
	// example:
	//
	// 0.0.0.0/0
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The range of the source ports for the security group rule. Valid values:
	//
	// - If you set the IpProtocol to TCP or UDP, the port range is 1 to 65535. You must use a forward slash (/) to separate the start port and the end port, for example, 1/200.
	//
	// - If you set the IpProtocol to ICMP, the port range is -1/-1.
	//
	// - If you set the IpProtocol to ALL, the port range is -1/-1.
	//
	// example:
	//
	// 22/22
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s CreateSecurityGroupRequestPermissions) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityGroupRequestPermissions) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupRequestPermissions) GetDescription() *string {
	return s.Description
}

func (s *CreateSecurityGroupRequestPermissions) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *CreateSecurityGroupRequestPermissions) GetDirection() *string {
	return s.Direction
}

func (s *CreateSecurityGroupRequestPermissions) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *CreateSecurityGroupRequestPermissions) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *CreateSecurityGroupRequestPermissions) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *CreateSecurityGroupRequestPermissions) GetPolicy() *string {
	return s.Policy
}

func (s *CreateSecurityGroupRequestPermissions) GetPortRange() *string {
	return s.PortRange
}

func (s *CreateSecurityGroupRequestPermissions) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateSecurityGroupRequestPermissions) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *CreateSecurityGroupRequestPermissions) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *CreateSecurityGroupRequestPermissions) SetDescription(v string) *CreateSecurityGroupRequestPermissions {
	s.Description = &v
	return s
}

func (s *CreateSecurityGroupRequestPermissions) SetDestCidrIp(v string) *CreateSecurityGroupRequestPermissions {
	s.DestCidrIp = &v
	return s
}

func (s *CreateSecurityGroupRequestPermissions) SetDirection(v string) *CreateSecurityGroupRequestPermissions {
	s.Direction = &v
	return s
}

func (s *CreateSecurityGroupRequestPermissions) SetIpProtocol(v string) *CreateSecurityGroupRequestPermissions {
	s.IpProtocol = &v
	return s
}

func (s *CreateSecurityGroupRequestPermissions) SetIpv6DestCidrIp(v string) *CreateSecurityGroupRequestPermissions {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *CreateSecurityGroupRequestPermissions) SetIpv6SourceCidrIp(v string) *CreateSecurityGroupRequestPermissions {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *CreateSecurityGroupRequestPermissions) SetPolicy(v string) *CreateSecurityGroupRequestPermissions {
	s.Policy = &v
	return s
}

func (s *CreateSecurityGroupRequestPermissions) SetPortRange(v string) *CreateSecurityGroupRequestPermissions {
	s.PortRange = &v
	return s
}

func (s *CreateSecurityGroupRequestPermissions) SetPriority(v int32) *CreateSecurityGroupRequestPermissions {
	s.Priority = &v
	return s
}

func (s *CreateSecurityGroupRequestPermissions) SetSourceCidrIp(v string) *CreateSecurityGroupRequestPermissions {
	s.SourceCidrIp = &v
	return s
}

func (s *CreateSecurityGroupRequestPermissions) SetSourcePortRange(v string) *CreateSecurityGroupRequestPermissions {
	s.SourcePortRange = &v
	return s
}

func (s *CreateSecurityGroupRequestPermissions) Validate() error {
	return dara.Validate(s)
}
