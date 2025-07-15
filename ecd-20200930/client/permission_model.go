// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPermission interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *Permission
	GetCreateTime() *string
	SetDescription(v string) *Permission
	GetDescription() *string
	SetDestCidrIp(v string) *Permission
	GetDestCidrIp() *string
	SetIpProtocol(v string) *Permission
	GetIpProtocol() *string
	SetNicType(v string) *Permission
	GetNicType() *string
	SetPolicy(v string) *Permission
	GetPolicy() *string
	SetPortRange(v string) *Permission
	GetPortRange() *string
	SetPriority(v string) *Permission
	GetPriority() *string
	SetSourceCidrIp(v string) *Permission
	GetSourceCidrIp() *string
	SetSourcePortRange(v string) *Permission
	GetSourcePortRange() *string
}

type Permission struct {
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestCidrIp      *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	NicType         *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	Priority        *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SourceCidrIp    *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s Permission) String() string {
	return dara.Prettify(s)
}

func (s Permission) GoString() string {
	return s.String()
}

func (s *Permission) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Permission) GetDescription() *string {
	return s.Description
}

func (s *Permission) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *Permission) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *Permission) GetNicType() *string {
	return s.NicType
}

func (s *Permission) GetPolicy() *string {
	return s.Policy
}

func (s *Permission) GetPortRange() *string {
	return s.PortRange
}

func (s *Permission) GetPriority() *string {
	return s.Priority
}

func (s *Permission) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *Permission) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *Permission) SetCreateTime(v string) *Permission {
	s.CreateTime = &v
	return s
}

func (s *Permission) SetDescription(v string) *Permission {
	s.Description = &v
	return s
}

func (s *Permission) SetDestCidrIp(v string) *Permission {
	s.DestCidrIp = &v
	return s
}

func (s *Permission) SetIpProtocol(v string) *Permission {
	s.IpProtocol = &v
	return s
}

func (s *Permission) SetNicType(v string) *Permission {
	s.NicType = &v
	return s
}

func (s *Permission) SetPolicy(v string) *Permission {
	s.Policy = &v
	return s
}

func (s *Permission) SetPortRange(v string) *Permission {
	s.PortRange = &v
	return s
}

func (s *Permission) SetPriority(v string) *Permission {
	s.Priority = &v
	return s
}

func (s *Permission) SetSourceCidrIp(v string) *Permission {
	s.SourceCidrIp = &v
	return s
}

func (s *Permission) SetSourcePortRange(v string) *Permission {
	s.SourcePortRange = &v
	return s
}

func (s *Permission) Validate() error {
	return dara.Validate(s)
}
