// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeSecurityGroupAttributeResponseBody
	GetDescription() *string
	SetPermissions(v *DescribeSecurityGroupAttributeResponseBodyPermissions) *DescribeSecurityGroupAttributeResponseBody
	GetPermissions() *DescribeSecurityGroupAttributeResponseBodyPermissions
	SetRequestId(v string) *DescribeSecurityGroupAttributeResponseBody
	GetRequestId() *string
	SetSecurityGroupId(v string) *DescribeSecurityGroupAttributeResponseBody
	GetSecurityGroupId() *string
	SetSecurityGroupName(v string) *DescribeSecurityGroupAttributeResponseBody
	GetSecurityGroupName() *string
}

type DescribeSecurityGroupAttributeResponseBody struct {
	// The description of the security group.
	//
	// example:
	//
	// testDescription1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Details about the rules.
	Permissions *DescribeSecurityGroupAttributeResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the destination security group.
	//
	// example:
	//
	// sg-bp1gxw6bznjjvhu3****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the destination security group.
	//
	// example:
	//
	// testSecurityGroupName2
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s DescribeSecurityGroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetPermissions() *DescribeSecurityGroupAttributeResponseBodyPermissions {
	return s.Permissions
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetDescription(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetPermissions(v *DescribeSecurityGroupAttributeResponseBodyPermissions) *DescribeSecurityGroupAttributeResponseBody {
	s.Permissions = v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetRequestId(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetSecurityGroupId(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetSecurityGroupName(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.SecurityGroupName = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) Validate() error {
	if s.Permissions != nil {
		if err := s.Permissions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityGroupAttributeResponseBodyPermissions struct {
	Permission []*DescribeSecurityGroupAttributeResponseBodyPermissionsPermission `json:"Permission,omitempty" xml:"Permission,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupAttributeResponseBodyPermissions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissions) GetPermission() []*DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	return s.Permission
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissions) SetPermission(v []*DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) *DescribeSecurityGroupAttributeResponseBodyPermissions {
	s.Permission = v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissions) Validate() error {
	if s.Permission != nil {
		for _, item := range s.Permission {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityGroupAttributeResponseBodyPermissionsPermission struct {
	// The time at which the security group rule was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-12-12T07:28:38Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description.
	//
	// example:
	//
	// testDescription1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The range of destination IP addresses for outbound access control.
	//
	// example:
	//
	// 0.0.0.0/0
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// The direction in which the security group rule is applied.
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The transport layer protocol.
	//
	// example:
	//
	// TCP
	IpProtocol       *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Ipv6DestCidrIp   *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The policy.
	//
	// example:
	//
	// Accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The source port range.
	//
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the rule.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The range of source IP addresses for inbound access control.
	//
	// example:
	//
	// 0.0.0.0/0
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The source port number range for the security group.
	//
	// example:
	//
	// 22/22
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetDescription() *string {
	return s.Description
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetDirection() *string {
	return s.Direction
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetPortRange() *string {
	return s.PortRange
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetCreationTime(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.CreationTime = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetDescription(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.Description = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetDestCidrIp(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.DestCidrIp = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetDirection(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.Direction = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetIpProtocol(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.IpProtocol = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetIpv6DestCidrIp(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetIpv6SourceCidrIp(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetPolicy(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.Policy = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetPortRange(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.PortRange = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetPriority(v int32) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.Priority = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetSourceCidrIp(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetSourcePortRange(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.SourcePortRange = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) Validate() error {
	return dara.Validate(s)
}
