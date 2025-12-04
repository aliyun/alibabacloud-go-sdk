// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubnetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidr(v string) *CreateSubnetRequest
	GetCidr() *string
	SetRegionId(v string) *CreateSubnetRequest
	GetRegionId() *string
	SetSubnetName(v string) *CreateSubnetRequest
	GetSubnetName() *string
	SetTag(v []*CreateSubnetRequestTag) *CreateSubnetRequest
	GetTag() []*CreateSubnetRequestTag
	SetType(v string) *CreateSubnetRequest
	GetType() *string
	SetVpdId(v string) *CreateSubnetRequest
	GetVpdId() *string
	SetZoneId(v string) *CreateSubnetRequest
	GetZoneId() *string
}

type CreateSubnetRequest struct {
	// The CIDR block of the Subnet.
	//
	// 	- The network segment of the subnet must be a proper subset of the network segment of Lingjun to which it belongs, and the mask must be between 16 bits and 29 bits, which can provide 8 to 65536 addresses. For example, the CIDR block of the Lingjun CIDR block is 192.168.0.0/16, and the CIDR blocks of the subnets under the Lingjun CIDR block are 192.168.0.0/17 to 192.168.0.0/29.
	//
	// 	- The first and last three IP addresses of each subnet segment are reserved by the system. For example, the CIDR blocks of the subnet are 192.168.1.0/24,192.168.1.0, 192.168.1.253, 192.168.1.254, and 192.168.1.255.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/16
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Lingjun subnet instance name
	//
	// This parameter is required.
	//
	// example:
	//
	// subnet-1
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tag []*CreateSubnetRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Lingjun Subnet Usage Type; optional; optional. Valid values:
	//
	// 	- **If you do not set this field for a common type**
	//
	// 	- **OOB*	- :OOB type
	//
	// 	- **LB**: LB type
	//
	// example:
	//
	// OOB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the Lingjun CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-xcuhjyrj
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateSubnetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSubnetRequest) GoString() string {
	return s.String()
}

func (s *CreateSubnetRequest) GetCidr() *string {
	return s.Cidr
}

func (s *CreateSubnetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSubnetRequest) GetSubnetName() *string {
	return s.SubnetName
}

func (s *CreateSubnetRequest) GetTag() []*CreateSubnetRequestTag {
	return s.Tag
}

func (s *CreateSubnetRequest) GetType() *string {
	return s.Type
}

func (s *CreateSubnetRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *CreateSubnetRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateSubnetRequest) SetCidr(v string) *CreateSubnetRequest {
	s.Cidr = &v
	return s
}

func (s *CreateSubnetRequest) SetRegionId(v string) *CreateSubnetRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSubnetRequest) SetSubnetName(v string) *CreateSubnetRequest {
	s.SubnetName = &v
	return s
}

func (s *CreateSubnetRequest) SetTag(v []*CreateSubnetRequestTag) *CreateSubnetRequest {
	s.Tag = v
	return s
}

func (s *CreateSubnetRequest) SetType(v string) *CreateSubnetRequest {
	s.Type = &v
	return s
}

func (s *CreateSubnetRequest) SetVpdId(v string) *CreateSubnetRequest {
	s.VpdId = &v
	return s
}

func (s *CreateSubnetRequest) SetZoneId(v string) *CreateSubnetRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateSubnetRequest) Validate() error {
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

type CreateSubnetRequestTag struct {
	// The tag key of the VPN attachment.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-subnet
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the VPN connection.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// subnet-tag-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSubnetRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateSubnetRequestTag) GoString() string {
	return s.String()
}

func (s *CreateSubnetRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateSubnetRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateSubnetRequestTag) SetKey(v string) *CreateSubnetRequestTag {
	s.Key = &v
	return s
}

func (s *CreateSubnetRequestTag) SetValue(v string) *CreateSubnetRequestTag {
	s.Value = &v
	return s
}

func (s *CreateSubnetRequestTag) Validate() error {
	return dara.Validate(s)
}
