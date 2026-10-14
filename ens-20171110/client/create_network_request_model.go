// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *CreateNetworkRequest
	GetCidrBlock() *string
	SetDescription(v string) *CreateNetworkRequest
	GetDescription() *string
	SetEnableIpv6(v bool) *CreateNetworkRequest
	GetEnableIpv6() *bool
	SetEnsRegionId(v string) *CreateNetworkRequest
	GetEnsRegionId() *string
	SetIpv6Isp(v string) *CreateNetworkRequest
	GetIpv6Isp() *string
	SetNetworkName(v string) *CreateNetworkRequest
	GetNetworkName() *string
	SetTag(v []*CreateNetworkRequestTag) *CreateNetworkRequest
	GetTag() []*CreateNetworkRequestTag
}

type CreateNetworkRequest struct {
	// The CIDR block of the network. You can use one of the following CIDR blocks or their subsets:
	//
	// - 10.0.0.0/8 (default).
	//
	// - 172.16.0.0/12.
	//
	// - 192.168.0.0/16.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// A description of the network.
	//
	// Must be 2 to 256 characters long, start with a letter or a Chinese character, and cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// this is my first network
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableIpv6  *bool   `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// The node ID of ENS.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// example:
	//
	// telecom
	Ipv6Isp *string `json:"Ipv6Isp,omitempty" xml:"Ipv6Isp,omitempty"`
	// The network name. Follow these naming rules:
	//
	// - Must be 2 to 128 characters long, using English or Chinese characters.
	//
	// - Must start with a letter or a Chinese character. Cannot start with http\\:// or https\\://.
	//
	// - Can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// example:
	//
	// abc
	NetworkName *string `json:"NetworkName,omitempty" xml:"NetworkName,omitempty"`
	// Resource tags.
	Tag []*CreateNetworkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateNetworkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNetworkRequest) GetEnableIpv6() *bool {
	return s.EnableIpv6
}

func (s *CreateNetworkRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateNetworkRequest) GetIpv6Isp() *string {
	return s.Ipv6Isp
}

func (s *CreateNetworkRequest) GetNetworkName() *string {
	return s.NetworkName
}

func (s *CreateNetworkRequest) GetTag() []*CreateNetworkRequestTag {
	return s.Tag
}

func (s *CreateNetworkRequest) SetCidrBlock(v string) *CreateNetworkRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateNetworkRequest) SetDescription(v string) *CreateNetworkRequest {
	s.Description = &v
	return s
}

func (s *CreateNetworkRequest) SetEnableIpv6(v bool) *CreateNetworkRequest {
	s.EnableIpv6 = &v
	return s
}

func (s *CreateNetworkRequest) SetEnsRegionId(v string) *CreateNetworkRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateNetworkRequest) SetIpv6Isp(v string) *CreateNetworkRequest {
	s.Ipv6Isp = &v
	return s
}

func (s *CreateNetworkRequest) SetNetworkName(v string) *CreateNetworkRequest {
	s.NetworkName = &v
	return s
}

func (s *CreateNetworkRequest) SetTag(v []*CreateNetworkRequestTag) *CreateNetworkRequest {
	s.Tag = v
	return s
}

func (s *CreateNetworkRequest) Validate() error {
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

type CreateNetworkRequestTag struct {
	// The tag key for the instance. N ranges from **1*	- to **20**.
	//
	// - Cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// - Supports up to 64 characters.
	//
	// - Cannot be an empty string.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value for the resource. N ranges from 1 to 20. This value can be an empty string once provided. It supports up to 128 characters, cannot start with acs:, and cannot contain http\\:// or https\\://.
	//
	// example:
	//
	// tagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNetworkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateNetworkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateNetworkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateNetworkRequestTag) SetKey(v string) *CreateNetworkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateNetworkRequestTag) SetValue(v string) *CreateNetworkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateNetworkRequestTag) Validate() error {
	return dara.Validate(s)
}
