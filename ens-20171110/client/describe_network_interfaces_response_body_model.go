// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInterfacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfaceSets(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) *DescribeNetworkInterfacesResponseBody
	GetNetworkInterfaceSets() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets
	SetPageNumber(v int32) *DescribeNetworkInterfacesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNetworkInterfacesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeNetworkInterfacesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeNetworkInterfacesResponseBody
	GetTotalCount() *int32
}

type DescribeNetworkInterfacesResponseBody struct {
	// Details about the ENIs.
	NetworkInterfaceSets *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets `json:"NetworkInterfaceSets,omitempty" xml:"NetworkInterfaceSets,omitempty" type:"Struct"`
	// The number of the page to return. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 708AF9CE-FF92-5DF9-93F8-B7754AB1061A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries in the list.
	//
	// example:
	//
	// 49
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNetworkInterfacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBody) GetNetworkInterfaceSets() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets {
	return s.NetworkInterfaceSets
}

func (s *DescribeNetworkInterfacesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNetworkInterfacesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNetworkInterfacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkInterfacesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNetworkInterfacesResponseBody) SetNetworkInterfaceSets(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) *DescribeNetworkInterfacesResponseBody {
	s.NetworkInterfaceSets = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBody) SetPageNumber(v int32) *DescribeNetworkInterfacesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBody) SetPageSize(v int32) *DescribeNetworkInterfacesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBody) SetRequestId(v string) *DescribeNetworkInterfacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBody) SetTotalCount(v int32) *DescribeNetworkInterfacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets struct {
	NetworkInterfaceSet []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet `json:"NetworkInterfaceSet,omitempty" xml:"NetworkInterfaceSet,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) GetNetworkInterfaceSet() []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	return s.NetworkInterfaceSet
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) SetNetworkInterfaceSet(v []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.NetworkInterfaceSet = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet struct {
	// The time when the ENI was created. Specify the time in the ISO 8601 standard in the yyyy-MM-ddThh:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-02-22T03:53:25Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of ENI.
	//
	// example:
	//
	// test-description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// id-jakarta-1
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the instance to which the ENI is bound.
	//
	// example:
	//
	// i-5siavnr3
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IPv6 addresses of the ENIs.
	Ipv6Sets *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets `json:"Ipv6Sets,omitempty" xml:"Ipv6Sets,omitempty" type:"Struct"`
	// The MAC address of the ENI.
	//
	// example:
	//
	// 00:16:3e:08:60:0a
	MacAddress *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	// The ID of the network.
	//
	// example:
	//
	// n-5w0qd03adw****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The ID of the ENI.
	//
	// example:
	//
	// eni-uf686a5
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The ENI name.
	//
	// example:
	//
	// primaryTest
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 12.23.3.4
	PrimaryIp *string `json:"PrimaryIp,omitempty" xml:"PrimaryIp,omitempty"`
	// The primary private IP address. Valid values:
	//
	// 	- **Public**: public IP address.
	//
	// 	- **Private**: internal IP address.
	//
	// example:
	//
	// private
	PrimaryIpType *string `json:"PrimaryIpType,omitempty" xml:"PrimaryIpType,omitempty"`
	// Details about the private IP address.
	PrivateIpSets *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets `json:"PrivateIpSets,omitempty" xml:"PrivateIpSets,omitempty" type:"Struct"`
	// The ID of the security group.
	SecurityGroupIds *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
	// The status of the ENI. Valid values:
	//
	// 	- Available: The ENI is available.
	//
	// 	- Attaching: The ENI is being attached to an instance.
	//
	// 	- InUse: The ENI is attached to an instance.
	//
	// 	- Detaching: The ENI is being detached from an instance.
	//
	// 	- Deleting: The ENI is being deleted.
	//
	// example:
	//
	// In_use
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the ENI. Valid values:
	//
	// 	- Primary
	//
	// 	- Secondary
	//
	// example:
	//
	// Primary
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-5rqswx1trlsj9
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetIpv6Sets() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets {
	return s.Ipv6Sets
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetMacAddress() *string {
	return s.MacAddress
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetPrimaryIp() *string {
	return s.PrimaryIp
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetPrimaryIpType() *string {
	return s.PrimaryIpType
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetPrivateIpSets() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets {
	return s.PrivateIpSets
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetSecurityGroupIds() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds {
	return s.SecurityGroupIds
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetType() *string {
	return s.Type
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetCreationTime(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.CreationTime = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetDescription(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.Description = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetEnsRegionId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetInstanceId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetIpv6Sets(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.Ipv6Sets = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetMacAddress(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.MacAddress = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetNetworkId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.NetworkId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetNetworkInterfaceId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetNetworkInterfaceName(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.NetworkInterfaceName = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetPrimaryIp(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.PrimaryIp = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetPrimaryIpType(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.PrimaryIpType = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetPrivateIpSets(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.PrivateIpSets = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetSecurityGroupIds(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetStatus(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.Status = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetType(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.Type = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetVSwitchId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.VSwitchId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets struct {
	Ipv6Set []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set `json:"Ipv6Set,omitempty" xml:"Ipv6Set,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets) GetIpv6Set() []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set {
	return s.Ipv6Set
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets) SetIpv6Set(v []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets {
	s.Ipv6Set = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set struct {
	// The IPv6 address of the ENI.
	//
	// example:
	//
	// 2605:340:cdb1:XXXX:XXXX:XXXX:XXXX:e2d6
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set) SetIpv6Address(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set {
	s.Ipv6Address = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets struct {
	PrivateIpSet []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet `json:"PrivateIpSet,omitempty" xml:"PrivateIpSet,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets) GetPrivateIpSet() []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet {
	return s.PrivateIpSet
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets) SetPrivateIpSet(v []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets {
	s.PrivateIpSet = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet struct {
	// Specifies whether the private IP address is the primary private IP address. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 192.168.0.130
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) GetPrimary() *bool {
	return s.Primary
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) SetPrimary(v bool) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet {
	s.Primary = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) SetPrivateIpAddress(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds struct {
	SecurityGroup []*string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds) GetSecurityGroup() []*string {
	return s.SecurityGroup
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds) SetSecurityGroup(v []*string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds {
	s.SecurityGroup = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds) Validate() error {
	return dara.Validate(s)
}
