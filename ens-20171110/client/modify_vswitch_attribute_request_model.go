// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVSwitchAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyVSwitchAttributeRequest
	GetDescription() *string
	SetEnableIPv6(v bool) *ModifyVSwitchAttributeRequest
	GetEnableIPv6() *bool
	SetIpv6CidrBlock(v string) *ModifyVSwitchAttributeRequest
	GetIpv6CidrBlock() *string
	SetVSwitchId(v string) *ModifyVSwitchAttributeRequest
	GetVSwitchId() *string
	SetVSwitchName(v string) *ModifyVSwitchAttributeRequest
	GetVSwitchName() *string
}

type ModifyVSwitchAttributeRequest struct {
	// The description.
	//
	// - The description must be 2 to 256 characters in length.
	//
	// - It cannot start with http:// or https://.
	//
	// example:
	//
	// this is my first vswitch
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable IPv6 for the vSwitch. Valid values:
	//
	// - true: enables IPv6. The VPC to which the vSwitch belongs must have IPv6 enabled, and you must also specify Ipv6CidrBlock to assign an IPv6 CIDR block to the vSwitch.
	//
	// - false (default): disables IPv6. Before you disable IPv6 for the vSwitch, make sure that no IPv6 addresses are in use, and do not specify Ipv6CidrBlock at the same time.
	//
	// example:
	//
	// false
	EnableIPv6 *bool `json:"EnableIPv6,omitempty" xml:"EnableIPv6,omitempty"`
	// The IPv6 CIDR block of the vSwitch.
	//
	// > You can specify this parameter to assign an IPv6 CIDR block to the vSwitch only when IPv6 is enabled for the VPC to which the vSwitch belongs. After the IPv6 CIDR block is assigned, it cannot be changed. Make sure that the IPv6 CIDR block does not overlap with the CIDR blocks of other vSwitches in the same VPC.
	//
	// example:
	//
	// 2408:XXXX:0:a600::/110
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	// The vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// n-****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The vSwitch name.
	//
	// - The name must be 2 to 128 characters in length.
	//
	// - It must start with a letter or a Chinese character and cannot start with http:// or https://.
	//
	// example:
	//
	// Test-switch
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s ModifyVSwitchAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVSwitchAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyVSwitchAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyVSwitchAttributeRequest) GetEnableIPv6() *bool {
	return s.EnableIPv6
}

func (s *ModifyVSwitchAttributeRequest) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *ModifyVSwitchAttributeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyVSwitchAttributeRequest) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *ModifyVSwitchAttributeRequest) SetDescription(v string) *ModifyVSwitchAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetEnableIPv6(v bool) *ModifyVSwitchAttributeRequest {
	s.EnableIPv6 = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetIpv6CidrBlock(v string) *ModifyVSwitchAttributeRequest {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetVSwitchId(v string) *ModifyVSwitchAttributeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetVSwitchName(v string) *ModifyVSwitchAttributeRequest {
	s.VSwitchName = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) Validate() error {
	return dara.Validate(s)
}
