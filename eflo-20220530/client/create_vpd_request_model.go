// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidr(v string) *CreateVpdRequest
	GetCidr() *string
	SetRegionId(v string) *CreateVpdRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateVpdRequest
	GetResourceGroupId() *string
	SetSubnets(v []*CreateVpdRequestSubnets) *CreateVpdRequest
	GetSubnets() []*CreateVpdRequestSubnets
	SetTag(v []*CreateVpdRequestTag) *CreateVpdRequest
	GetTag() []*CreateVpdRequestTag
	SetVpdName(v string) *CreateVpdRequest
	GetVpdName() *string
}

type CreateVpdRequest struct {
	// The CIDR block of the VPD.
	//
	// 	- We recommend that you use an RFC private endpoint as the Lingjun CIDR block, such as 10.0.0.0/8,172.16.0.0/12,192.168.0.0/16. In scenarios where the Doringjun CIDR block is connected to each other or where the Lingjun CIDR block is connected to a VPC, make sure that the addresses do not conflict with each other.
	//
	// 	- You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16 and their subnets as the primary IPv4 CIDR block of the VPD.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/8
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The region ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Lingjun subnet information list
	Subnets []*CreateVpdRequestSubnets `json:"Subnets,omitempty" xml:"Subnets,omitempty" type:"Repeated"`
	// A tag.
	Tag []*CreateVpdRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Lingjun CIDR block instance name
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-1
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s CreateVpdRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpdRequest) GoString() string {
	return s.String()
}

func (s *CreateVpdRequest) GetCidr() *string {
	return s.Cidr
}

func (s *CreateVpdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpdRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVpdRequest) GetSubnets() []*CreateVpdRequestSubnets {
	return s.Subnets
}

func (s *CreateVpdRequest) GetTag() []*CreateVpdRequestTag {
	return s.Tag
}

func (s *CreateVpdRequest) GetVpdName() *string {
	return s.VpdName
}

func (s *CreateVpdRequest) SetCidr(v string) *CreateVpdRequest {
	s.Cidr = &v
	return s
}

func (s *CreateVpdRequest) SetRegionId(v string) *CreateVpdRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpdRequest) SetResourceGroupId(v string) *CreateVpdRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVpdRequest) SetSubnets(v []*CreateVpdRequestSubnets) *CreateVpdRequest {
	s.Subnets = v
	return s
}

func (s *CreateVpdRequest) SetTag(v []*CreateVpdRequestTag) *CreateVpdRequest {
	s.Tag = v
	return s
}

func (s *CreateVpdRequest) SetVpdName(v string) *CreateVpdRequest {
	s.VpdName = &v
	return s
}

func (s *CreateVpdRequest) Validate() error {
	if s.Subnets != nil {
		for _, item := range s.Subnets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type CreateVpdRequestSubnets struct {
	// The CIDR block of the Subnet.
	//
	// 	- The network segment of the subnet must be a proper subset of the network segment of Lingjun to which it belongs, and the mask must be between 16 bits and 29 bits, which can provide 8 to 65536 addresses. For example, the CIDR block of the Lingjun CIDR block is 192.168.0.0/16, and the CIDR blocks of the subnets under the Lingjun CIDR block are 192.168.0.0/17 to 192.168.0.0/29.
	//
	// 	- The first and last three IP addresses of each subnet segment are reserved by the system. For example, the CIDR blocks of the subnet are 192.168.1.0/24,192.168.1.0, 192.168.1.253, 192.168.1.254, and 192.168.1.255.
	//
	// example:
	//
	// 10.1.0.0/16
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Lingjun subnet instance name
	//
	// example:
	//
	// subnet-1
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	// Lingjun Subnet Usage Type; optional; optional. Valid values:
	//
	// 	- **Generic type is not specified**.
	//
	// 	- **OOB*	- :OOB type
	//
	// 	- **LB**: LB type
	//
	// example:
	//
	// OOB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The zone ID of the disk.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateVpdRequestSubnets) String() string {
	return dara.Prettify(s)
}

func (s CreateVpdRequestSubnets) GoString() string {
	return s.String()
}

func (s *CreateVpdRequestSubnets) GetCidr() *string {
	return s.Cidr
}

func (s *CreateVpdRequestSubnets) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpdRequestSubnets) GetSubnetName() *string {
	return s.SubnetName
}

func (s *CreateVpdRequestSubnets) GetType() *string {
	return s.Type
}

func (s *CreateVpdRequestSubnets) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateVpdRequestSubnets) SetCidr(v string) *CreateVpdRequestSubnets {
	s.Cidr = &v
	return s
}

func (s *CreateVpdRequestSubnets) SetRegionId(v string) *CreateVpdRequestSubnets {
	s.RegionId = &v
	return s
}

func (s *CreateVpdRequestSubnets) SetSubnetName(v string) *CreateVpdRequestSubnets {
	s.SubnetName = &v
	return s
}

func (s *CreateVpdRequestSubnets) SetType(v string) *CreateVpdRequestSubnets {
	s.Type = &v
	return s
}

func (s *CreateVpdRequestSubnets) SetZoneId(v string) *CreateVpdRequestSubnets {
	s.ZoneId = &v
	return s
}

func (s *CreateVpdRequestSubnets) Validate() error {
	return dara.Validate(s)
}

type CreateVpdRequestTag struct {
	// The tag key of the VPN attachment.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// vpd-wulanchabu
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the VPN connection.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each tag key corresponds to a tag value. You can enter a maximum of 20 tag values at a time.
	//
	// example:
	//
	// wulanchabu-a
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVpdRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVpdRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVpdRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVpdRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVpdRequestTag) SetKey(v string) *CreateVpdRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVpdRequestTag) SetValue(v string) *CreateVpdRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVpdRequestTag) Validate() error {
	return dara.Validate(s)
}
