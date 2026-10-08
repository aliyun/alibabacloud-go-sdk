// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCSecurityGroupPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInnerAccessPolicy(v string) *DescribeRCSecurityGroupPermissionResponseBody
	GetInnerAccessPolicy() *string
	SetRegionId(v string) *DescribeRCSecurityGroupPermissionResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeRCSecurityGroupPermissionResponseBody
	GetRequestId() *string
	SetSecurityGroupId(v string) *DescribeRCSecurityGroupPermissionResponseBody
	GetSecurityGroupId() *string
	SetSecurityGroupPermissions(v []*DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) *DescribeRCSecurityGroupPermissionResponseBody
	GetSecurityGroupPermissions() []*DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions
	SetVpcId(v string) *DescribeRCSecurityGroupPermissionResponseBody
	GetVpcId() *string
}

type DescribeRCSecurityGroupPermissionResponseBody struct {
	InnerAccessPolicy        *string                                                                  `json:"InnerAccessPolicy,omitempty" xml:"InnerAccessPolicy,omitempty"`
	RegionId                 *string                                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId                *string                                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityGroupId          *string                                                                  `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupPermissions []*DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions `json:"SecurityGroupPermissions,omitempty" xml:"SecurityGroupPermissions,omitempty" type:"Repeated"`
	VpcId                    *string                                                                  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRCSecurityGroupPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCSecurityGroupPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCSecurityGroupPermissionResponseBody) GetInnerAccessPolicy() *string {
	return s.InnerAccessPolicy
}

func (s *DescribeRCSecurityGroupPermissionResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCSecurityGroupPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCSecurityGroupPermissionResponseBody) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeRCSecurityGroupPermissionResponseBody) GetSecurityGroupPermissions() []*DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions {
	return s.SecurityGroupPermissions
}

func (s *DescribeRCSecurityGroupPermissionResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRCSecurityGroupPermissionResponseBody) SetInnerAccessPolicy(v string) *DescribeRCSecurityGroupPermissionResponseBody {
	s.InnerAccessPolicy = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBody) SetRegionId(v string) *DescribeRCSecurityGroupPermissionResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBody) SetRequestId(v string) *DescribeRCSecurityGroupPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBody) SetSecurityGroupId(v string) *DescribeRCSecurityGroupPermissionResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBody) SetSecurityGroupPermissions(v []*DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) *DescribeRCSecurityGroupPermissionResponseBody {
	s.SecurityGroupPermissions = v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBody) SetVpcId(v string) *DescribeRCSecurityGroupPermissionResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBody) Validate() error {
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

type DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DestCidrIp          *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	Direction           *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	IpProtocol          *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Policy              *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PortRange           *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	Priority            *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SecurityGroupRuleId *string `json:"SecurityGroupRuleId,omitempty" xml:"SecurityGroupRuleId,omitempty"`
	SourceCidrIp        *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	SourcePortRange     *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) GoString() string {
	return s.String()
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) GetDirection() *string {
	return s.Direction
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) GetPortRange() *string {
	return s.PortRange
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) GetPriority() *string {
	return s.Priority
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) GetSecurityGroupRuleId() *string {
	return s.SecurityGroupRuleId
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) SetCreateTime(v string) *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions {
	s.CreateTime = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) SetDestCidrIp(v string) *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions {
	s.DestCidrIp = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) SetDirection(v string) *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions {
	s.Direction = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) SetIpProtocol(v string) *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions {
	s.IpProtocol = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) SetPolicy(v string) *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions {
	s.Policy = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) SetPortRange(v string) *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions {
	s.PortRange = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) SetPriority(v string) *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions {
	s.Priority = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) SetSecurityGroupRuleId(v string) *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions {
	s.SecurityGroupRuleId = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) SetSourceCidrIp(v string) *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) SetSourcePortRange(v string) *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions {
	s.SourcePortRange = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponseBodySecurityGroupPermissions) Validate() error {
	return dara.Validate(s)
}
