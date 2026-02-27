// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableIpAddressCount(v int64) *DescribeVSwitchAttributesResponseBody
	GetAvailableIpAddressCount() *int64
	SetCidrBlock(v string) *DescribeVSwitchAttributesResponseBody
	GetCidrBlock() *string
	SetCreationTime(v string) *DescribeVSwitchAttributesResponseBody
	GetCreationTime() *string
	SetDescription(v string) *DescribeVSwitchAttributesResponseBody
	GetDescription() *string
	SetEnabledIpv6(v bool) *DescribeVSwitchAttributesResponseBody
	GetEnabledIpv6() *bool
	SetIpv6CidrBlock(v string) *DescribeVSwitchAttributesResponseBody
	GetIpv6CidrBlock() *string
	SetIsDefault(v bool) *DescribeVSwitchAttributesResponseBody
	GetIsDefault() *bool
	SetNetworkAclId(v string) *DescribeVSwitchAttributesResponseBody
	GetNetworkAclId() *string
	SetOwnerId(v int64) *DescribeVSwitchAttributesResponseBody
	GetOwnerId() *int64
	SetRequestId(v string) *DescribeVSwitchAttributesResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeVSwitchAttributesResponseBody
	GetResourceGroupId() *string
	SetRouteTable(v *DescribeVSwitchAttributesResponseBodyRouteTable) *DescribeVSwitchAttributesResponseBody
	GetRouteTable() *DescribeVSwitchAttributesResponseBodyRouteTable
	SetShareType(v string) *DescribeVSwitchAttributesResponseBody
	GetShareType() *string
	SetStatus(v string) *DescribeVSwitchAttributesResponseBody
	GetStatus() *string
	SetTags(v *DescribeVSwitchAttributesResponseBodyTags) *DescribeVSwitchAttributesResponseBody
	GetTags() *DescribeVSwitchAttributesResponseBodyTags
	SetVSwitchId(v string) *DescribeVSwitchAttributesResponseBody
	GetVSwitchId() *string
	SetVSwitchName(v string) *DescribeVSwitchAttributesResponseBody
	GetVSwitchName() *string
	SetVpcId(v string) *DescribeVSwitchAttributesResponseBody
	GetVpcId() *string
	SetZoneId(v string) *DescribeVSwitchAttributesResponseBody
	GetZoneId() *string
}

type DescribeVSwitchAttributesResponseBody struct {
	// The number of available IP addresses.
	//
	// example:
	//
	// 12
	AvailableIpAddressCount *int64 `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	// The CIDR block of the vSwitch.
	//
	// example:
	//
	// 192.168.0.1/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The time when the vSwitch was created.
	//
	// example:
	//
	// 2021-08-22T10:40:25Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the vSwitch.
	//
	// example:
	//
	// abc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether IPv6 is enabled for the vSwitch. If you enable IPv6, you must configure the IPv6 CIDR block of the vSwitch. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnabledIpv6 *bool `json:"EnabledIpv6,omitempty" xml:"EnabledIpv6,omitempty"`
	// The IPv6 CIDR block of the vSwitch.
	//
	// example:
	//
	// 2408:XXXX:3c5:44e::/64
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	// Indicates whether the vSwitch is the default vSwitch. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The network access control list (ACL) rules.
	//
	// example:
	//
	// 1
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 287683832402436789
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7B48B4B9-1EAD-469F-B488-594DAB4B6A1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the ACL belongs.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The information about the route table that is associated with the vSwitch.
	RouteTable *DescribeVSwitchAttributesResponseBodyRouteTable `json:"RouteTable,omitempty" xml:"RouteTable,omitempty" type:"Struct"`
	// Indicates whether the vSwitch is shared.
	//
	// 	- If no value is returned, the vSwitch is a regular vSwitch.
	//
	// 	- If **Shared*	- is returned, the vSwitch is shared.
	//
	// 	- If **Sharing*	- is returned, the vSwitch is being shared.
	//
	// example:
	//
	// Shared
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The status of the vSwitch. Valid values:
	//
	// 	- **Pending**
	//
	// 	- **Available**
	//
	// example:
	//
	// Pending
	Status *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeVSwitchAttributesResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-25b7pv15t****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The vSwitch name.
	//
	// example:
	//
	// test
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The ID of the VPC to which the vSwitch belongs.
	//
	// example:
	//
	// vpc-257gq642n****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone to which the vSwitch belongs.
	//
	// example:
	//
	// cn-beijing-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVSwitchAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchAttributesResponseBody) GetAvailableIpAddressCount() *int64 {
	return s.AvailableIpAddressCount
}

func (s *DescribeVSwitchAttributesResponseBody) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVSwitchAttributesResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeVSwitchAttributesResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeVSwitchAttributesResponseBody) GetEnabledIpv6() *bool {
	return s.EnabledIpv6
}

func (s *DescribeVSwitchAttributesResponseBody) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *DescribeVSwitchAttributesResponseBody) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVSwitchAttributesResponseBody) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *DescribeVSwitchAttributesResponseBody) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVSwitchAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVSwitchAttributesResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVSwitchAttributesResponseBody) GetRouteTable() *DescribeVSwitchAttributesResponseBodyRouteTable {
	return s.RouteTable
}

func (s *DescribeVSwitchAttributesResponseBody) GetShareType() *string {
	return s.ShareType
}

func (s *DescribeVSwitchAttributesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeVSwitchAttributesResponseBody) GetTags() *DescribeVSwitchAttributesResponseBodyTags {
	return s.Tags
}

func (s *DescribeVSwitchAttributesResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchAttributesResponseBody) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchAttributesResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVSwitchAttributesResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVSwitchAttributesResponseBody) SetAvailableIpAddressCount(v int64) *DescribeVSwitchAttributesResponseBody {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetCidrBlock(v string) *DescribeVSwitchAttributesResponseBody {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetCreationTime(v string) *DescribeVSwitchAttributesResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetDescription(v string) *DescribeVSwitchAttributesResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetEnabledIpv6(v bool) *DescribeVSwitchAttributesResponseBody {
	s.EnabledIpv6 = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetIpv6CidrBlock(v string) *DescribeVSwitchAttributesResponseBody {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetIsDefault(v bool) *DescribeVSwitchAttributesResponseBody {
	s.IsDefault = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetNetworkAclId(v string) *DescribeVSwitchAttributesResponseBody {
	s.NetworkAclId = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetOwnerId(v int64) *DescribeVSwitchAttributesResponseBody {
	s.OwnerId = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetRequestId(v string) *DescribeVSwitchAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetResourceGroupId(v string) *DescribeVSwitchAttributesResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetRouteTable(v *DescribeVSwitchAttributesResponseBodyRouteTable) *DescribeVSwitchAttributesResponseBody {
	s.RouteTable = v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetShareType(v string) *DescribeVSwitchAttributesResponseBody {
	s.ShareType = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetStatus(v string) *DescribeVSwitchAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetTags(v *DescribeVSwitchAttributesResponseBodyTags) *DescribeVSwitchAttributesResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetVSwitchId(v string) *DescribeVSwitchAttributesResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetVSwitchName(v string) *DescribeVSwitchAttributesResponseBody {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetVpcId(v string) *DescribeVSwitchAttributesResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetZoneId(v string) *DescribeVSwitchAttributesResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) Validate() error {
	if s.RouteTable != nil {
		if err := s.RouteTable.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVSwitchAttributesResponseBodyRouteTable struct {
	// The ID of the route table that is associated with the vSwitch.
	//
	// example:
	//
	// vtb-bp145q7glnuzdv****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The type of the route table. Valid values:
	//
	// 	- **System**
	//
	// 	- **Custom**
	//
	// example:
	//
	// System
	RouteTableType *string `json:"RouteTableType,omitempty" xml:"RouteTableType,omitempty"`
}

func (s DescribeVSwitchAttributesResponseBodyRouteTable) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchAttributesResponseBodyRouteTable) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchAttributesResponseBodyRouteTable) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeVSwitchAttributesResponseBodyRouteTable) GetRouteTableType() *string {
	return s.RouteTableType
}

func (s *DescribeVSwitchAttributesResponseBodyRouteTable) SetRouteTableId(v string) *DescribeVSwitchAttributesResponseBodyRouteTable {
	s.RouteTableId = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBodyRouteTable) SetRouteTableType(v string) *DescribeVSwitchAttributesResponseBodyRouteTable {
	s.RouteTableType = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBodyRouteTable) Validate() error {
	return dara.Validate(s)
}

type DescribeVSwitchAttributesResponseBodyTags struct {
	Tag []*DescribeVSwitchAttributesResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchAttributesResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchAttributesResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchAttributesResponseBodyTags) GetTag() []*DescribeVSwitchAttributesResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeVSwitchAttributesResponseBodyTags) SetTag(v []*DescribeVSwitchAttributesResponseBodyTagsTag) *DescribeVSwitchAttributesResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeVSwitchAttributesResponseBodyTags) Validate() error {
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

type DescribeVSwitchAttributesResponseBodyTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVSwitchAttributesResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchAttributesResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchAttributesResponseBodyTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVSwitchAttributesResponseBodyTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVSwitchAttributesResponseBodyTagsTag) SetKey(v string) *DescribeVSwitchAttributesResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBodyTagsTag) SetValue(v string) *DescribeVSwitchAttributesResponseBodyTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
