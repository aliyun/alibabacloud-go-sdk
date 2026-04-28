// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPermissionCondition interface {
	dara.Model
	String() string
	GoString() string
	SetBoolEquals(v *PermissionConditionBoolEquals) *PermissionCondition
	GetBoolEquals() *PermissionConditionBoolEquals
	SetBoolNotEquals(v *PermissionConditionBoolNotEquals) *PermissionCondition
	GetBoolNotEquals() *PermissionConditionBoolNotEquals
	SetIpEquals(v *PermissionConditionIpEquals) *PermissionCondition
	GetIpEquals() *PermissionConditionIpEquals
	SetIpNotEquals(v *PermissionConditionIpNotEquals) *PermissionCondition
	GetIpNotEquals() *PermissionConditionIpNotEquals
	SetStringLike(v *PermissionConditionStringLike) *PermissionCondition
	GetStringLike() *PermissionConditionStringLike
	SetStringNotLike(v *PermissionConditionStringNotLike) *PermissionCondition
	GetStringNotLike() *PermissionConditionStringNotLike
}

type PermissionCondition struct {
	BoolEquals    *PermissionConditionBoolEquals    `json:"bool_equals,omitempty" xml:"bool_equals,omitempty" type:"Struct"`
	BoolNotEquals *PermissionConditionBoolNotEquals `json:"bool_not_equals,omitempty" xml:"bool_not_equals,omitempty" type:"Struct"`
	// The IP address condition, which is true when the IP address is equal to one of the following lists.
	IpEquals *PermissionConditionIpEquals `json:"ip_equals,omitempty" xml:"ip_equals,omitempty" type:"Struct"`
	// The IP address condition. This condition is true when the IP address is not equal to any of the following list.
	IpNotEquals *PermissionConditionIpNotEquals `json:"ip_not_equals,omitempty" xml:"ip_not_equals,omitempty" type:"Struct"`
	// The string match condition, which is true when the string is equal to one of the following lists.
	StringLike *PermissionConditionStringLike `json:"string_like,omitempty" xml:"string_like,omitempty" type:"Struct"`
	// The string match condition. This condition is true when the input string is not equal to any one of the following lists.
	StringNotLike *PermissionConditionStringNotLike `json:"string_not_like,omitempty" xml:"string_not_like,omitempty" type:"Struct"`
}

func (s PermissionCondition) String() string {
	return dara.Prettify(s)
}

func (s PermissionCondition) GoString() string {
	return s.String()
}

func (s *PermissionCondition) GetBoolEquals() *PermissionConditionBoolEquals {
	return s.BoolEquals
}

func (s *PermissionCondition) GetBoolNotEquals() *PermissionConditionBoolNotEquals {
	return s.BoolNotEquals
}

func (s *PermissionCondition) GetIpEquals() *PermissionConditionIpEquals {
	return s.IpEquals
}

func (s *PermissionCondition) GetIpNotEquals() *PermissionConditionIpNotEquals {
	return s.IpNotEquals
}

func (s *PermissionCondition) GetStringLike() *PermissionConditionStringLike {
	return s.StringLike
}

func (s *PermissionCondition) GetStringNotLike() *PermissionConditionStringNotLike {
	return s.StringNotLike
}

func (s *PermissionCondition) SetBoolEquals(v *PermissionConditionBoolEquals) *PermissionCondition {
	s.BoolEquals = v
	return s
}

func (s *PermissionCondition) SetBoolNotEquals(v *PermissionConditionBoolNotEquals) *PermissionCondition {
	s.BoolNotEquals = v
	return s
}

func (s *PermissionCondition) SetIpEquals(v *PermissionConditionIpEquals) *PermissionCondition {
	s.IpEquals = v
	return s
}

func (s *PermissionCondition) SetIpNotEquals(v *PermissionConditionIpNotEquals) *PermissionCondition {
	s.IpNotEquals = v
	return s
}

func (s *PermissionCondition) SetStringLike(v *PermissionConditionStringLike) *PermissionCondition {
	s.StringLike = v
	return s
}

func (s *PermissionCondition) SetStringNotLike(v *PermissionConditionStringNotLike) *PermissionCondition {
	s.StringNotLike = v
	return s
}

func (s *PermissionCondition) Validate() error {
	if s.BoolEquals != nil {
		if err := s.BoolEquals.Validate(); err != nil {
			return err
		}
	}
	if s.BoolNotEquals != nil {
		if err := s.BoolNotEquals.Validate(); err != nil {
			return err
		}
	}
	if s.IpEquals != nil {
		if err := s.IpEquals.Validate(); err != nil {
			return err
		}
	}
	if s.IpNotEquals != nil {
		if err := s.IpNotEquals.Validate(); err != nil {
			return err
		}
	}
	if s.StringLike != nil {
		if err := s.StringLike.Validate(); err != nil {
			return err
		}
	}
	if s.StringNotLike != nil {
		if err := s.StringNotLike.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PermissionConditionBoolEquals struct {
	IsShareLink *bool `json:"is_share_link,omitempty" xml:"is_share_link,omitempty"`
}

func (s PermissionConditionBoolEquals) String() string {
	return dara.Prettify(s)
}

func (s PermissionConditionBoolEquals) GoString() string {
	return s.String()
}

func (s *PermissionConditionBoolEquals) GetIsShareLink() *bool {
	return s.IsShareLink
}

func (s *PermissionConditionBoolEquals) SetIsShareLink(v bool) *PermissionConditionBoolEquals {
	s.IsShareLink = &v
	return s
}

func (s *PermissionConditionBoolEquals) Validate() error {
	return dara.Validate(s)
}

type PermissionConditionBoolNotEquals struct {
	IsShareLink *bool `json:"is_share_link,omitempty" xml:"is_share_link,omitempty"`
}

func (s PermissionConditionBoolNotEquals) String() string {
	return dara.Prettify(s)
}

func (s PermissionConditionBoolNotEquals) GoString() string {
	return s.String()
}

func (s *PermissionConditionBoolNotEquals) GetIsShareLink() *bool {
	return s.IsShareLink
}

func (s *PermissionConditionBoolNotEquals) SetIsShareLink(v bool) *PermissionConditionBoolNotEquals {
	s.IsShareLink = &v
	return s
}

func (s *PermissionConditionBoolNotEquals) Validate() error {
	return dara.Validate(s)
}

type PermissionConditionIpEquals struct {
	// The IP address of the client.
	ClientIp []*string `json:"client_ip,omitempty" xml:"client_ip,omitempty" type:"Repeated"`
}

func (s PermissionConditionIpEquals) String() string {
	return dara.Prettify(s)
}

func (s PermissionConditionIpEquals) GoString() string {
	return s.String()
}

func (s *PermissionConditionIpEquals) GetClientIp() []*string {
	return s.ClientIp
}

func (s *PermissionConditionIpEquals) SetClientIp(v []*string) *PermissionConditionIpEquals {
	s.ClientIp = v
	return s
}

func (s *PermissionConditionIpEquals) Validate() error {
	return dara.Validate(s)
}

type PermissionConditionIpNotEquals struct {
	// The IP address of the client.
	ClientIp []*string `json:"client_ip,omitempty" xml:"client_ip,omitempty" type:"Repeated"`
}

func (s PermissionConditionIpNotEquals) String() string {
	return dara.Prettify(s)
}

func (s PermissionConditionIpNotEquals) GoString() string {
	return s.String()
}

func (s *PermissionConditionIpNotEquals) GetClientIp() []*string {
	return s.ClientIp
}

func (s *PermissionConditionIpNotEquals) SetClientIp(v []*string) *PermissionConditionIpNotEquals {
	s.ClientIp = v
	return s
}

func (s *PermissionConditionIpNotEquals) Validate() error {
	return dara.Validate(s)
}

type PermissionConditionStringLike struct {
	// The vpcID of the client as a string match condition.
	VpcId []*string `json:"vpc_id,omitempty" xml:"vpc_id,omitempty" type:"Repeated"`
}

func (s PermissionConditionStringLike) String() string {
	return dara.Prettify(s)
}

func (s PermissionConditionStringLike) GoString() string {
	return s.String()
}

func (s *PermissionConditionStringLike) GetVpcId() []*string {
	return s.VpcId
}

func (s *PermissionConditionStringLike) SetVpcId(v []*string) *PermissionConditionStringLike {
	s.VpcId = v
	return s
}

func (s *PermissionConditionStringLike) Validate() error {
	return dara.Validate(s)
}

type PermissionConditionStringNotLike struct {
	// The vpcID of the client as a string match condition.
	VpcId []*string `json:"vpc_id,omitempty" xml:"vpc_id,omitempty" type:"Repeated"`
}

func (s PermissionConditionStringNotLike) String() string {
	return dara.Prettify(s)
}

func (s PermissionConditionStringNotLike) GoString() string {
	return s.String()
}

func (s *PermissionConditionStringNotLike) GetVpcId() []*string {
	return s.VpcId
}

func (s *PermissionConditionStringNotLike) SetVpcId(v []*string) *PermissionConditionStringNotLike {
	s.VpcId = v
	return s
}

func (s *PermissionConditionStringNotLike) Validate() error {
	return dara.Validate(s)
}
