// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCNetworkInterfacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfaceSets(v []*DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) *DescribeRCNetworkInterfacesResponseBody
	GetNetworkInterfaceSets() []*DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets
	SetRequestId(v string) *DescribeRCNetworkInterfacesResponseBody
	GetRequestId() *string
}

type DescribeRCNetworkInterfacesResponseBody struct {
	NetworkInterfaceSets []*DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets `json:"NetworkInterfaceSets,omitempty" xml:"NetworkInterfaceSets,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRCNetworkInterfacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCNetworkInterfacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCNetworkInterfacesResponseBody) GetNetworkInterfaceSets() []*DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	return s.NetworkInterfaceSets
}

func (s *DescribeRCNetworkInterfacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCNetworkInterfacesResponseBody) SetNetworkInterfaceSets(v []*DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) *DescribeRCNetworkInterfacesResponseBody {
	s.NetworkInterfaceSets = v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBody) SetRequestId(v string) *DescribeRCNetworkInterfacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBody) Validate() error {
	if s.NetworkInterfaceSets != nil {
		for _, item := range s.NetworkInterfaceSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets struct {
	AssociatedPublicIp *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsAssociatedPublicIp `json:"AssociatedPublicIp,omitempty" xml:"AssociatedPublicIp,omitempty" type:"Struct"`
	// example:
	//
	// 2022-12-25T12:31:31Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// DescriptionTest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// rc-m5sc1271fv344a1r****
	InstanceId *string                                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ipv6Sets   []*DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsIpv6Sets `json:"Ipv6Sets,omitempty" xml:"Ipv6Sets,omitempty" type:"Repeated"`
	// example:
	//
	// 00:16:3e:12:**:**
	MacAddress *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	// example:
	//
	// eni-bp125p95hhdhn3ot****
	NetworkInterfaceId *string                                                                     `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	PrivateIpSets      []*DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsPrivateIpSets `json:"PrivateIpSets,omitempty" xml:"PrivateIpSets,omitempty" type:"Repeated"`
	// example:
	//
	// rg-2ze88m67qx5z****
	ResourceGroupId  *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// Available
	Status *string                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// Secondary
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// vsw-bp16usj2p27htro3****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp1j7w3gc1cexjqd****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GoString() string {
	return s.String()
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetAssociatedPublicIp() *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsAssociatedPublicIp {
	return s.AssociatedPublicIp
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetDescription() *string {
	return s.Description
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetIpv6Sets() []*DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsIpv6Sets {
	return s.Ipv6Sets
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetMacAddress() *string {
	return s.MacAddress
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetPrivateIpSets() []*DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsPrivateIpSets {
	return s.PrivateIpSets
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetStatus() *string {
	return s.Status
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetTags() []*DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsTags {
	return s.Tags
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetType() *string {
	return s.Type
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetAssociatedPublicIp(v *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsAssociatedPublicIp) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.AssociatedPublicIp = v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetCreationTime(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.CreationTime = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetDescription(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.Description = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetInstanceId(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetIpv6Sets(v []*DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsIpv6Sets) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.Ipv6Sets = v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetMacAddress(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.MacAddress = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetNetworkInterfaceId(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetPrivateIpSets(v []*DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsPrivateIpSets) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.PrivateIpSets = v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetResourceGroupId(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetSecurityGroupIds(v []*string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetStatus(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.Status = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetTags(v []*DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsTags) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.Tags = v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetType(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.Type = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetVSwitchId(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.VSwitchId = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetVpcId(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.VpcId = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) SetZoneId(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.ZoneId = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSets) Validate() error {
	if s.AssociatedPublicIp != nil {
		if err := s.AssociatedPublicIp.Validate(); err != nil {
			return err
		}
	}
	if s.Ipv6Sets != nil {
		for _, item := range s.Ipv6Sets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PrivateIpSets != nil {
		for _, item := range s.PrivateIpSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsAssociatedPublicIp struct {
	// example:
	//
	// ``116.62.**.**``
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
}

func (s DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsAssociatedPublicIp) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsAssociatedPublicIp) GoString() string {
	return s.String()
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsAssociatedPublicIp) GetPublicIpAddress() *string {
	return s.PublicIpAddress
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsAssociatedPublicIp) SetPublicIpAddress(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsAssociatedPublicIp {
	s.PublicIpAddress = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsAssociatedPublicIp) Validate() error {
	return dara.Validate(s)
}

type DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsIpv6Sets struct {
	// example:
	//
	// 2408:4321:180:1701:94c7:bc38:3bfa:****
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
}

func (s DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsIpv6Sets) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsIpv6Sets) GoString() string {
	return s.String()
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsIpv6Sets) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsIpv6Sets) SetIpv6Address(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsIpv6Sets {
	s.Ipv6Address = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsIpv6Sets) Validate() error {
	return dara.Validate(s)
}

type DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsPrivateIpSets struct {
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
	// example:
	//
	// ``172.17.**.**``
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsPrivateIpSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsPrivateIpSets) GoString() string {
	return s.String()
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsPrivateIpSets) GetPrimary() *bool {
	return s.Primary
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsPrivateIpSets) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsPrivateIpSets) SetPrimary(v bool) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsPrivateIpSets {
	s.Primary = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsPrivateIpSets) SetPrivateIpAddress(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsPrivateIpSets {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsPrivateIpSets) Validate() error {
	return dara.Validate(s)
}

type DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsTags struct {
	// example:
	//
	// TestValue
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// TestKey
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsTags) GoString() string {
	return s.String()
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsTags) SetTagKey(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsTags {
	s.TagKey = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsTags) SetTagValue(v string) *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsTags {
	s.TagValue = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponseBodyNetworkInterfaceSetsTags) Validate() error {
	return dara.Validate(s)
}
