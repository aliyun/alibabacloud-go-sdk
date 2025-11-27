// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeRCSecurityGroupPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *AuthorizeRCSecurityGroupPermissionRequest
	GetDirection() *string
	SetRegionId(v string) *AuthorizeRCSecurityGroupPermissionRequest
	GetRegionId() *string
	SetSecurityGroupId(v string) *AuthorizeRCSecurityGroupPermissionRequest
	GetSecurityGroupId() *string
	SetSecurityGroupPermissions(v []*AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) *AuthorizeRCSecurityGroupPermissionRequest
	GetSecurityGroupPermissions() []*AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions
}

type AuthorizeRCSecurityGroupPermissionRequest struct {
	// The direction of the rule. Valid values:
	//
	// 	- **ingress**: the inbound security group rule.
	//
	// 	- **egress**: the outbound security group rule.
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-2ze27hs990o2hn9****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The information about the security group.
	SecurityGroupPermissions []*AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions `json:"SecurityGroupPermissions,omitempty" xml:"SecurityGroupPermissions,omitempty" type:"Repeated"`
}

func (s AuthorizeRCSecurityGroupPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeRCSecurityGroupPermissionRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeRCSecurityGroupPermissionRequest) GetDirection() *string {
	return s.Direction
}

func (s *AuthorizeRCSecurityGroupPermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AuthorizeRCSecurityGroupPermissionRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *AuthorizeRCSecurityGroupPermissionRequest) GetSecurityGroupPermissions() []*AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions {
	return s.SecurityGroupPermissions
}

func (s *AuthorizeRCSecurityGroupPermissionRequest) SetDirection(v string) *AuthorizeRCSecurityGroupPermissionRequest {
	s.Direction = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionRequest) SetRegionId(v string) *AuthorizeRCSecurityGroupPermissionRequest {
	s.RegionId = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionRequest) SetSecurityGroupId(v string) *AuthorizeRCSecurityGroupPermissionRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionRequest) SetSecurityGroupPermissions(v []*AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) *AuthorizeRCSecurityGroupPermissionRequest {
	s.SecurityGroupPermissions = v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionRequest) Validate() error {
	if s.SecurityGroupPermissions != nil {
		for _, item := range s.SecurityGroupPermissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions struct {
	// The destination CIDR block for outbound access control. CIDR blocks and IPv4 addresses are supported.
	//
	// example:
	//
	// 192.168.0.1/12
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// The protocol type supported by the rule. The value is not case-sensitive. Valid values:
	//
	// 	- **ICMP**
	//
	// 	- **GRE**
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// 	- **ALL**: All protocols are supported.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The action that you want to specify in the rule.
	//
	// example:
	//
	// Accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of destination ports that correspond to the transport layer protocol of the destination security group. Valid values:
	//
	// 	- The value is in the X/Y format when IpProtocol is set to TCP or UDP. X specifies the start port number and Y specifies the end port number. X and Y range from **1*	- to **65535**. The start port number and the end port number are separated by a forward slash (/). Correct example: **1/200**. Incorrect example: **200/1**.
	//
	// 	- Valid value when IpProtocol is set to ICMP: **-1/-1**.
	//
	// 	- Valid value when IpProtocol is set to GRE: **-1/-1**.
	//
	// 	- Valid value when IpProtocol is set to ALL: **-1/-1**.
	//
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the rule. Valid values: 1 to 100. A smaller value indicates a higher priority. When multiple security group rules have the same priority, drop rules take precedence.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The source CIDR block for inbound access control. CIDR blocks and IPv4 addresses are supported.
	//
	// example:
	//
	// 192.168.0.1/12
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The range of port numbers that correspond to the transport layer protocol for the source security group. Valid values:
	//
	// 	- The value is in the X/Y format when IpProtocol is set to TCP or UDP. X specifies the start port number and Y specifies the end port number. X and Y range from **1*	- to **65535**. The start port number and the end port number are separated by a forward slash (/). Correct example: **1/200**. Incorrect example: **200/1**.
	//
	// 	- Valid value when IpProtocol is set to ICMP: **-1/-1**.
	//
	// 	- Valid value when IpProtocol is set to GRE: **-1/-1**.
	//
	// 	- Valid value when IpProtocol is set to ALL: **-1/-1**.
	//
	// example:
	//
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) GoString() string {
	return s.String()
}

func (s *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) GetPolicy() *string {
	return s.Policy
}

func (s *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) GetPortRange() *string {
	return s.PortRange
}

func (s *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) GetPriority() *int32 {
	return s.Priority
}

func (s *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) SetDestCidrIp(v string) *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions {
	s.DestCidrIp = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) SetIpProtocol(v string) *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions {
	s.IpProtocol = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) SetPolicy(v string) *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions {
	s.Policy = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) SetPortRange(v string) *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions {
	s.PortRange = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) SetPriority(v int32) *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions {
	s.Priority = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) SetSourceCidrIp(v string) *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions {
	s.SourceCidrIp = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) SetSourcePortRange(v string) *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions {
	s.SourcePortRange = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionRequestSecurityGroupPermissions) Validate() error {
	return dara.Validate(s)
}
