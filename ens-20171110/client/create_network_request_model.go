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
	SetEnsRegionId(v string) *CreateNetworkRequest
	GetEnsRegionId() *string
	SetNetworkName(v string) *CreateNetworkRequest
	GetNetworkName() *string
	SetTag(v []*CreateNetworkRequestTag) *CreateNetworkRequest
	GetTag() []*CreateNetworkRequestTag
}

type CreateNetworkRequest struct {
	// The CIDR block of the network. You can use one of the following CIDR blocks or their subnets as the CIDR block of the network:
	//
	// 	- 10.0.0.0/8 (default)
	//
	// 	- 172.16.0.0/12
	//
	// 	- 192.168.0.0/16
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The description of the network.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
	//
	// example:
	//
	// this is my first network
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the edge node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The name of the network. The name must meet the following requirements:
	//
	// 	- The name must be 2 to 128 characters in length.
	//
	// 	- The name must start with a letter but cannot start with http:// or https://.
	//
	// 	- The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// abc
	NetworkName *string `json:"NetworkName,omitempty" xml:"NetworkName,omitempty"`
	// The resource tags.
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

func (s *CreateNetworkRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
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

func (s *CreateNetworkRequest) SetEnsRegionId(v string) *CreateNetworkRequest {
	s.EnsRegionId = &v
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
	return dara.Validate(s)
}

type CreateNetworkRequestTag struct {
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
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is added to the resource. You can specify up to 20 tag values. The tag value can be an empty string. The tag value can be up to 128 characters in length. It cannot start with acs: or contain http:// or https://.
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
