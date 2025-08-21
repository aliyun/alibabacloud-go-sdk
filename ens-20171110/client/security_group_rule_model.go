// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSecurityGroupRule interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SecurityGroupRule
	GetDescription() *string
	SetDestCidrIp(v string) *SecurityGroupRule
	GetDestCidrIp() *string
	SetDirection(v string) *SecurityGroupRule
	GetDirection() *string
	SetIpProtocol(v string) *SecurityGroupRule
	GetIpProtocol() *string
	SetPolicy(v string) *SecurityGroupRule
	GetPolicy() *string
	SetPortRange(v string) *SecurityGroupRule
	GetPortRange() *string
	SetSourceCidrIp(v string) *SecurityGroupRule
	GetSourceCidrIp() *string
	SetSourcePortRange(v string) *SecurityGroupRule
	GetSourcePortRange() *string
	SetPriority(v int32) *SecurityGroupRule
	GetPriority() *int32
}

type SecurityGroupRule struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestCidrIp      *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	Direction       *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	SourceCidrIp    *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	Priority        *int32  `json:"priority,omitempty" xml:"priority,omitempty"`
}

func (s SecurityGroupRule) String() string {
	return dara.Prettify(s)
}

func (s SecurityGroupRule) GoString() string {
	return s.String()
}

func (s *SecurityGroupRule) GetDescription() *string {
	return s.Description
}

func (s *SecurityGroupRule) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *SecurityGroupRule) GetDirection() *string {
	return s.Direction
}

func (s *SecurityGroupRule) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *SecurityGroupRule) GetPolicy() *string {
	return s.Policy
}

func (s *SecurityGroupRule) GetPortRange() *string {
	return s.PortRange
}

func (s *SecurityGroupRule) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *SecurityGroupRule) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *SecurityGroupRule) GetPriority() *int32 {
	return s.Priority
}

func (s *SecurityGroupRule) SetDescription(v string) *SecurityGroupRule {
	s.Description = &v
	return s
}

func (s *SecurityGroupRule) SetDestCidrIp(v string) *SecurityGroupRule {
	s.DestCidrIp = &v
	return s
}

func (s *SecurityGroupRule) SetDirection(v string) *SecurityGroupRule {
	s.Direction = &v
	return s
}

func (s *SecurityGroupRule) SetIpProtocol(v string) *SecurityGroupRule {
	s.IpProtocol = &v
	return s
}

func (s *SecurityGroupRule) SetPolicy(v string) *SecurityGroupRule {
	s.Policy = &v
	return s
}

func (s *SecurityGroupRule) SetPortRange(v string) *SecurityGroupRule {
	s.PortRange = &v
	return s
}

func (s *SecurityGroupRule) SetSourceCidrIp(v string) *SecurityGroupRule {
	s.SourceCidrIp = &v
	return s
}

func (s *SecurityGroupRule) SetSourcePortRange(v string) *SecurityGroupRule {
	s.SourcePortRange = &v
	return s
}

func (s *SecurityGroupRule) SetPriority(v int32) *SecurityGroupRule {
	s.Priority = &v
	return s
}

func (s *SecurityGroupRule) Validate() error {
	return dara.Validate(s)
}
