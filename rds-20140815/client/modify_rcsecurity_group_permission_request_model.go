// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCSecurityGroupPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestCidrIp(v string) *ModifyRCSecurityGroupPermissionRequest
	GetDestCidrIp() *string
	SetDirection(v string) *ModifyRCSecurityGroupPermissionRequest
	GetDirection() *string
	SetIpProtocol(v string) *ModifyRCSecurityGroupPermissionRequest
	GetIpProtocol() *string
	SetPolicy(v string) *ModifyRCSecurityGroupPermissionRequest
	GetPolicy() *string
	SetPortRange(v string) *ModifyRCSecurityGroupPermissionRequest
	GetPortRange() *string
	SetPriority(v string) *ModifyRCSecurityGroupPermissionRequest
	GetPriority() *string
	SetRegionId(v string) *ModifyRCSecurityGroupPermissionRequest
	GetRegionId() *string
	SetSecurityGroupId(v string) *ModifyRCSecurityGroupPermissionRequest
	GetSecurityGroupId() *string
	SetSecurityGroupRuleId(v string) *ModifyRCSecurityGroupPermissionRequest
	GetSecurityGroupRuleId() *string
	SetSourceCidrIp(v string) *ModifyRCSecurityGroupPermissionRequest
	GetSourceCidrIp() *string
	SetSourcePortRange(v string) *ModifyRCSecurityGroupPermissionRequest
	GetSourcePortRange() *string
}

type ModifyRCSecurityGroupPermissionRequest struct {
	// example:
	//
	// 172.16.0.0/24
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// example:
	//
	// Accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// sg-2ze27hs990o2hn9****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// sgr-uf6ffg4du0e9jis2****
	SecurityGroupRuleId *string `json:"SecurityGroupRuleId,omitempty" xml:"SecurityGroupRuleId,omitempty"`
	// example:
	//
	// 192.168.0.0/0
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// example:
	//
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s ModifyRCSecurityGroupPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCSecurityGroupPermissionRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCSecurityGroupPermissionRequest) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *ModifyRCSecurityGroupPermissionRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModifyRCSecurityGroupPermissionRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifyRCSecurityGroupPermissionRequest) GetPolicy() *string {
	return s.Policy
}

func (s *ModifyRCSecurityGroupPermissionRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *ModifyRCSecurityGroupPermissionRequest) GetPriority() *string {
	return s.Priority
}

func (s *ModifyRCSecurityGroupPermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCSecurityGroupPermissionRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ModifyRCSecurityGroupPermissionRequest) GetSecurityGroupRuleId() *string {
	return s.SecurityGroupRuleId
}

func (s *ModifyRCSecurityGroupPermissionRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *ModifyRCSecurityGroupPermissionRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *ModifyRCSecurityGroupPermissionRequest) SetDestCidrIp(v string) *ModifyRCSecurityGroupPermissionRequest {
	s.DestCidrIp = &v
	return s
}

func (s *ModifyRCSecurityGroupPermissionRequest) SetDirection(v string) *ModifyRCSecurityGroupPermissionRequest {
	s.Direction = &v
	return s
}

func (s *ModifyRCSecurityGroupPermissionRequest) SetIpProtocol(v string) *ModifyRCSecurityGroupPermissionRequest {
	s.IpProtocol = &v
	return s
}

func (s *ModifyRCSecurityGroupPermissionRequest) SetPolicy(v string) *ModifyRCSecurityGroupPermissionRequest {
	s.Policy = &v
	return s
}

func (s *ModifyRCSecurityGroupPermissionRequest) SetPortRange(v string) *ModifyRCSecurityGroupPermissionRequest {
	s.PortRange = &v
	return s
}

func (s *ModifyRCSecurityGroupPermissionRequest) SetPriority(v string) *ModifyRCSecurityGroupPermissionRequest {
	s.Priority = &v
	return s
}

func (s *ModifyRCSecurityGroupPermissionRequest) SetRegionId(v string) *ModifyRCSecurityGroupPermissionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCSecurityGroupPermissionRequest) SetSecurityGroupId(v string) *ModifyRCSecurityGroupPermissionRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifyRCSecurityGroupPermissionRequest) SetSecurityGroupRuleId(v string) *ModifyRCSecurityGroupPermissionRequest {
	s.SecurityGroupRuleId = &v
	return s
}

func (s *ModifyRCSecurityGroupPermissionRequest) SetSourceCidrIp(v string) *ModifyRCSecurityGroupPermissionRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *ModifyRCSecurityGroupPermissionRequest) SetSourcePortRange(v string) *ModifyRCSecurityGroupPermissionRequest {
	s.SourcePortRange = &v
	return s
}

func (s *ModifyRCSecurityGroupPermissionRequest) Validate() error {
	return dara.Validate(s)
}
