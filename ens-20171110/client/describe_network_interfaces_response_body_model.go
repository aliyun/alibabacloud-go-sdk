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
	if s.NetworkInterfaceSets != nil {
		if err := s.NetworkInterfaceSets.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.NetworkInterfaceSet != nil {
		for _, item := range s.NetworkInterfaceSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet struct {
	CreationTime         *string                                                                                       `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description          *string                                                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	EnsRegionId          *string                                                                                       `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	InstanceId           *string                                                                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ipv6Sets             *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets         `json:"Ipv6Sets,omitempty" xml:"Ipv6Sets,omitempty" type:"Struct"`
	MacAddress           *string                                                                                       `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	NetworkId            *string                                                                                       `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	NetworkInterfaceId   *string                                                                                       `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	NetworkInterfaceName *string                                                                                       `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	PrimaryIp            *string                                                                                       `json:"PrimaryIp,omitempty" xml:"PrimaryIp,omitempty"`
	PrimaryIpType        *string                                                                                       `json:"PrimaryIpType,omitempty" xml:"PrimaryIpType,omitempty"`
	PrivateIpSets        *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets    `json:"PrivateIpSets,omitempty" xml:"PrivateIpSets,omitempty" type:"Struct"`
	SecurityGroupIds     *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
	Status               *string                                                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                 *string                                                                                       `json:"Type,omitempty" xml:"Type,omitempty"`
	VSwitchId            *string                                                                                       `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VmncLearn            *bool                                                                                         `json:"VmncLearn,omitempty" xml:"VmncLearn,omitempty"`
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

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetVmncLearn() *bool {
	return s.VmncLearn
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

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetVmncLearn(v bool) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.VmncLearn = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) Validate() error {
	if s.Ipv6Sets != nil {
		if err := s.Ipv6Sets.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateIpSets != nil {
		if err := s.PrivateIpSets.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityGroupIds != nil {
		if err := s.SecurityGroupIds.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.Ipv6Set != nil {
		for _, item := range s.Ipv6Set {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set struct {
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
	if s.PrivateIpSet != nil {
		for _, item := range s.PrivateIpSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet struct {
	Primary          *bool   `json:"Primary,omitempty" xml:"Primary,omitempty"`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	PrivateIpStatus  *string `json:"PrivateIpStatus,omitempty" xml:"PrivateIpStatus,omitempty"`
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

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) GetPrivateIpStatus() *string {
	return s.PrivateIpStatus
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) SetPrimary(v bool) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet {
	s.Primary = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) SetPrivateIpAddress(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) SetPrivateIpStatus(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet {
	s.PrivateIpStatus = &v
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
