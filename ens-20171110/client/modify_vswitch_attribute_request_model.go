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
	// Description.
	//
	// - Must be 2 to 256 characters in length and can contain English or Chinese characters.
	//
	// - Cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// this is my first vswitch
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableIPv6  *bool   `json:"EnableIPv6,omitempty" xml:"EnableIPv6,omitempty"`
	// example:
	//
	// 2408:XXXX:0:a600::/110
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	// vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// n-****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// vSwitch name.
	//
	// - Must be 2 to 128 characters long. Can contain letters, digits, Chinese characters, underscores (_), and hyphens (-).
	//
	// - Must start with a letter or a Chinese character. Cannot start with http\\:// or https\\://.
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
