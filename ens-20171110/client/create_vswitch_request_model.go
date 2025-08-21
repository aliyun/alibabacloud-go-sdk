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
	SetNetworkId(v string) *CreateVSwitchRequest
	GetNetworkId() *string
	SetTag(v []*CreateVSwitchRequestTag) *CreateVSwitchRequest
	GetTag() []*CreateVSwitchRequestTag
	SetVSwitchName(v string) *CreateVSwitchRequest
	GetVSwitchName() *string
}

type CreateVSwitchRequest struct {
	// The CIDR block of the vSwitch. Take note of the following limits:
	//
	// 	- The subnet mask must be 16 to 29 bits in length.
	//
	// 	- The CIDR block of the vSwitch must fall within the CIDR block of the VPC to which the vSwitch belongs.
	//
	// 	- The CIDR block of the vSwitch cannot be the same as the destination CIDR block in a route entry of the VPC. However, it can be a subset of the destination CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The description of the vSwitch.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
	//
	// example:
	//
	// This is my vswitch.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the edge node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-xian-unicom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the network to which the vSwitch that you want to create belongs.
	//
	// example:
	//
	// n-257gqcdfvx6n****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The tags.
	Tag []*CreateVSwitchRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the vSwitch. The name must meet the following requirements:
	//
	// 	- The name must be 2 to 128 characters in length.
	//
	// 	- The name must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// Default value: null.
	//
	// example:
	//
	// test
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
	return dara.Validate(s)
}

type CreateVSwitchRequestTag struct {
	// The key of the tag that are to add to the instance. Valid values of N: **1*	- to **20**.
	//
	// 	- The key cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// 	- The key must be up to 64 characters in length.
	//
	// 	- The tag key cannot be an empty string.
	//
	// example:
	//
	// team
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is added to the resource. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length. It cannot start with acs: or contain http:// or https://.
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
