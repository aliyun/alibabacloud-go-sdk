// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *CreateVSwitchRequest
	GetCidrBlock() *string
	SetDescription(v string) *CreateVSwitchRequest
	GetDescription() *string
	SetEnsRegionId(v string) *CreateVSwitchRequest
	GetEnsRegionId() *string
	SetIpv6CidrBlock(v string) *CreateVSwitchRequest
	GetIpv6CidrBlock() *string
	SetNetworkId(v string) *CreateVSwitchRequest
	GetNetworkId() *string
	SetTag(v []*CreateVSwitchRequestTag) *CreateVSwitchRequest
	GetTag() []*CreateVSwitchRequestTag
	SetVSwitchName(v string) *CreateVSwitchRequest
	GetVSwitchName() *string
}

type CreateVSwitchRequest struct {
	// The CIDR block of the VSwitch. The following requirements apply:
	//
	// - The prefix length must be an integer from 16 to 29.
	//
	// - The CIDR block of the VSwitch must be a subset of the CIDR block of the VPC to which the VSwitch belongs.
	//
	// - The CIDR block of the VSwitch cannot be the same as the destination CIDR block of an existing route in the VPC, but it can be a subset of the destination CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The description of the VSwitch.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter or a Chinese character, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my vswitch.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the ENS node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-xian-unicom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// example:
	//
	// 2408:XXXX:0:a600::/110
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	// The ID of the network to which the VSwitch belongs.
	//
	// example:
	//
	// n-257gqcdfvx6n****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The tags of the resource. You can specify up to 20 tags.
	Tag []*CreateVSwitchRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the VSwitch. The following requirements apply:
	//
	// - The name must be 2 to 128 characters in length.
	//
	// - The name must start with a letter or a Chinese character, and cannot start with `http://` or `https://`. The name can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// Default value: an empty string.
	//
	// example:
	//
	// Test vSwitch
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s CreateVSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVSwitchRequest) GoString() string {
	return s.String()
}

func (s *CreateVSwitchRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateVSwitchRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVSwitchRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateVSwitchRequest) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *CreateVSwitchRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *CreateVSwitchRequest) GetTag() []*CreateVSwitchRequestTag {
	return s.Tag
}

func (s *CreateVSwitchRequest) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *CreateVSwitchRequest) SetCidrBlock(v string) *CreateVSwitchRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateVSwitchRequest) SetDescription(v string) *CreateVSwitchRequest {
	s.Description = &v
	return s
}

func (s *CreateVSwitchRequest) SetEnsRegionId(v string) *CreateVSwitchRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateVSwitchRequest) SetIpv6CidrBlock(v string) *CreateVSwitchRequest {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *CreateVSwitchRequest) SetNetworkId(v string) *CreateVSwitchRequest {
	s.NetworkId = &v
	return s
}

func (s *CreateVSwitchRequest) SetTag(v []*CreateVSwitchRequestTag) *CreateVSwitchRequest {
	s.Tag = v
	return s
}

func (s *CreateVSwitchRequest) SetVSwitchName(v string) *CreateVSwitchRequest {
	s.VSwitchName = &v
	return s
}

func (s *CreateVSwitchRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateVSwitchRequestTag struct {
	// The key of the tag.
	//
	// - The tag key cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// - The tag key can be up to 64 characters long.
	//
	// - The tag key cannot be an empty string.
	//
	// example:
	//
	// team
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. It can be up to 128 characters long, can be an empty string, and cannot start with `acs:` or contain `http://` or `https://`.
	//
	// example:
	//
	// Deep
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVSwitchRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVSwitchRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVSwitchRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVSwitchRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVSwitchRequestTag) SetKey(v string) *CreateVSwitchRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVSwitchRequestTag) SetValue(v string) *CreateVSwitchRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVSwitchRequestTag) Validate() error {
	return dara.Validate(s)
}
