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
	// This parameter is required.
	Permissions []*CreateSecurityGroupPermissionsRequestPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
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
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 10.XX.XX.14/32
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// example:
	//
	// ::/0
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// example:
	//
	// ::/0
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 0.0.0.0/0
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
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
