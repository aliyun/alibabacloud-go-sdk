// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVSwitchesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVSwitchesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVSwitchesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVSwitchesResponseBody
	GetTotalCount() *int32
	SetVSwitches(v *DescribeVSwitchesResponseBodyVSwitches) *DescribeVSwitchesResponseBody
	GetVSwitches() *DescribeVSwitchesResponseBodyVSwitches
}

type DescribeVSwitchesResponseBody struct {
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A572171-4E27-40D1-BD36-D26C9E71E29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The details about the vSwitch.
	VSwitches *DescribeVSwitchesResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Struct"`
}

func (s DescribeVSwitchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVSwitchesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVSwitchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVSwitchesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVSwitchesResponseBody) GetVSwitches() *DescribeVSwitchesResponseBodyVSwitches {
	return s.VSwitches
}

func (s *DescribeVSwitchesResponseBody) SetPageNumber(v int32) *DescribeVSwitchesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetPageSize(v int32) *DescribeVSwitchesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetRequestId(v string) *DescribeVSwitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetTotalCount(v int32) *DescribeVSwitchesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetVSwitches(v *DescribeVSwitchesResponseBodyVSwitches) *DescribeVSwitchesResponseBody {
	s.VSwitches = v
	return s
}

func (s *DescribeVSwitchesResponseBody) Validate() error {
	if s.VSwitches != nil {
		if err := s.VSwitches.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVSwitchesResponseBodyVSwitches struct {
	VSwitch []*DescribeVSwitchesResponseBodyVSwitchesVSwitch `json:"VSwitch,omitempty" xml:"VSwitch,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchesResponseBodyVSwitches) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitches) GetVSwitch() []*DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	return s.VSwitch
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetVSwitch(v []*DescribeVSwitchesResponseBodyVSwitchesVSwitch) *DescribeVSwitchesResponseBodyVSwitches {
	s.VSwitch = v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitches) Validate() error {
	if s.VSwitch != nil {
		for _, item := range s.VSwitch {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVSwitchesResponseBodyVSwitchesVSwitch struct {
	// The number of available IP addresses in the vSwitch.
	//
	// example:
	//
	// 1
	AvailableIpAddressCount *int64 `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	// The IPv4 CIDR block of the vSwitch.
	//
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The time when the vSwitch was created.
	//
	// example:
	//
	// 2022-01-18T12:43:57Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the vSwitch.
	//
	// example:
	//
	// VSwitchDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether IPv6 is enabled for the vSwitch. If you enable IPv6, you must configure an IPv6 CIDR block for the vSwitch. Valid values:
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
	// 2408:4002:10c4:4e03::/64
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	// Indicates whether the vSwitch is the default vSwitch. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the network access control list (ACL).
	//
	// example:
	//
	// nacl-a2do9e413e0spzasx****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 25346073170691****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group to which the vSwitch belongs.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The information about the route table.
	RouteTable *DescribeVSwitchesResponseBodyVSwitchesVSwitchRouteTable `json:"RouteTable,omitempty" xml:"RouteTable,omitempty" type:"Struct"`
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
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the vSwitch.
	Tags *DescribeVSwitchesResponseBodyVSwitchesVSwitchTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-25bcdxs7pv1****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The vSwitch name.
	//
	// example:
	//
	// vSwitch
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The ID of the VPC to which the vSwitch belongs.
	//
	// example:
	//
	// vpc-257gcdcdq64****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone to which the vSwitch belongs.
	//
	// example:
	//
	// cn-hangzhou-d
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitch) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetAvailableIpAddressCount() *int64 {
	return s.AvailableIpAddressCount
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetDescription() *string {
	return s.Description
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetEnabledIpv6() *bool {
	return s.EnabledIpv6
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetRouteTable() *DescribeVSwitchesResponseBodyVSwitchesVSwitchRouteTable {
	return s.RouteTable
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetShareType() *string {
	return s.ShareType
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetStatus() *string {
	return s.Status
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetTags() *DescribeVSwitchesResponseBodyVSwitchesVSwitchTags {
	return s.Tags
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetAvailableIpAddressCount(v int64) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetCidrBlock(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetCreationTime(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.CreationTime = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetDescription(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Description = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetEnabledIpv6(v bool) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.EnabledIpv6 = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetIpv6CidrBlock(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetIsDefault(v bool) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.IsDefault = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetNetworkAclId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.NetworkAclId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetOwnerId(v int64) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.OwnerId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetResourceGroupId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetRouteTable(v *DescribeVSwitchesResponseBodyVSwitchesVSwitchRouteTable) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.RouteTable = v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetShareType(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.ShareType = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetStatus(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetTags(v *DescribeVSwitchesResponseBodyVSwitchesVSwitchTags) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Tags = v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetVSwitchId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetVSwitchName(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetVpcId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.VpcId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetZoneId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.ZoneId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) Validate() error {
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

type DescribeVSwitchesResponseBodyVSwitchesVSwitchRouteTable struct {
	// The ID of the route table that is associated with the vSwitch.
	//
	// example:
	//
	// vrt-bp145q7glnuzdv****
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

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitchRouteTable) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitchRouteTable) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchRouteTable) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchRouteTable) GetRouteTableType() *string {
	return s.RouteTableType
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchRouteTable) SetRouteTableId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitchRouteTable {
	s.RouteTableId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchRouteTable) SetRouteTableType(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitchRouteTable {
	s.RouteTableType = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchRouteTable) Validate() error {
	return dara.Validate(s)
}

type DescribeVSwitchesResponseBodyVSwitchesVSwitchTags struct {
	Tag []*DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitchTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitchTags) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTags) GetTag() []*DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag {
	return s.Tag
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTags) SetTag(v []*DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) *DescribeVSwitchesResponseBodyVSwitchesVSwitchTags {
	s.Tag = v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTags) Validate() error {
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

type DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// dev
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) SetKey(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) SetValue(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) Validate() error {
	return dara.Validate(s)
}
