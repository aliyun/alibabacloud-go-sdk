// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddressBookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcls(v []*DescribeAddressBookResponseBodyAcls) *DescribeAddressBookResponseBody
	GetAcls() []*DescribeAddressBookResponseBodyAcls
	SetPageNo(v string) *DescribeAddressBookResponseBody
	GetPageNo() *string
	SetPageSize(v string) *DescribeAddressBookResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeAddressBookResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeAddressBookResponseBody
	GetTotalCount() *string
}

type DescribeAddressBookResponseBody struct {
	// The list of address books.
	Acls []*DescribeAddressBookResponseBodyAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Repeated"`
	// The page number of the current page.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of address books on each page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B36F150A-1E27-43AA-B72C-D2AC712F09DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of address books.
	//
	// example:
	//
	// 100
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAddressBookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressBookResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponseBody) GetAcls() []*DescribeAddressBookResponseBodyAcls {
	return s.Acls
}

func (s *DescribeAddressBookResponseBody) GetPageNo() *string {
	return s.PageNo
}

func (s *DescribeAddressBookResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAddressBookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAddressBookResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeAddressBookResponseBody) SetAcls(v []*DescribeAddressBookResponseBodyAcls) *DescribeAddressBookResponseBody {
	s.Acls = v
	return s
}

func (s *DescribeAddressBookResponseBody) SetPageNo(v string) *DescribeAddressBookResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeAddressBookResponseBody) SetPageSize(v string) *DescribeAddressBookResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAddressBookResponseBody) SetRequestId(v string) *DescribeAddressBookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAddressBookResponseBody) SetTotalCount(v string) *DescribeAddressBookResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAddressBookResponseBody) Validate() error {
	if s.Acls != nil {
		for _, item := range s.Acls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAddressBookResponseBodyAcls struct {
	// The ID of the ACK cluster connector.
	//
	// example:
	//
	// ac-7c1bad6c3cc84c33baab
	AckClusterConnectorId *string `json:"AckClusterConnectorId,omitempty" xml:"AckClusterConnectorId,omitempty"`
	// The name of the ACK cluster connector.
	//
	// example:
	//
	// ack-cluster-connector-name
	AckClusterConnectorName *string `json:"AckClusterConnectorName,omitempty" xml:"AckClusterConnectorName,omitempty"`
	// The list of pod labels in the ACK cluster.
	AckLabels []*DescribeAddressBookResponseBodyAclsAckLabels `json:"AckLabels,omitempty" xml:"AckLabels,omitempty" type:"Repeated"`
	// The list of pod namespaces in the ACK cluster.
	AckNamespaces []*string `json:"AckNamespaces,omitempty" xml:"AckNamespaces,omitempty" type:"Repeated"`
	// The address list of the address book.
	AddressList []*string `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	// The number of addresses in the address book.
	//
	// example:
	//
	// 2
	AddressListCount *int32 `json:"AddressListCount,omitempty" xml:"AddressListCount,omitempty"`
	// The address list of the address book that includes descriptions for individual addresses.
	Addresses []*DescribeAddressBookResponseBodyAclsAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// The list of member accounts for the asset address book.
	AssetMemberUids []*int64 `json:"AssetMemberUids,omitempty" xml:"AssetMemberUids,omitempty" type:"Repeated"`
	// The list of regions and resource types for the asset address book.
	AssetRegionResourceTypes []*DescribeAddressBookResponseBodyAclsAssetRegionResourceTypes `json:"AssetRegionResourceTypes,omitempty" xml:"AssetRegionResourceTypes,omitempty" type:"Repeated"`
	// Indicates whether the public IP addresses of ECS instances that match new tags are automatically added to the address book. Valid values:
	//
	// - **0**: The public IP addresses are not automatically added.
	//
	// - **1**: The public IP addresses are automatically added.
	//
	// example:
	//
	// 1
	AutoAddTagEcs *int32 `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	// The description of the address book.
	//
	// example:
	//
	// DEMO
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the address book.
	//
	// example:
	//
	// demo_address_book
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the address book.
	//
	// example:
	//
	// ip
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The unique ID of the address book.
	//
	// example:
	//
	// f04ac7ce-628b-4cb7-be61-310222b7****
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The number of times the address book is referenced.
	//
	// example:
	//
	// 3
	ReferenceCount *int32 `json:"ReferenceCount,omitempty" xml:"ReferenceCount,omitempty"`
	// The region of the ACK cluster connector to which the address book belongs when GroupType is an ACK address book.
	//
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The list of ECS tags.
	TagList []*DescribeAddressBookResponseBodyAclsTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The relationship between multiple ECS tags. Valid values:
	//
	// - **or**: The relationship between multiple tags is OR. The public IP address of an ECS instance that matches any tag is added to the address book.
	//
	// - **and**: The relationship between multiple tags is AND. The public IP address of an ECS instance that matches all tags is added to the address book.
	//
	// example:
	//
	// and
	TagRelation *string `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
}

func (s DescribeAddressBookResponseBodyAcls) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressBookResponseBodyAcls) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponseBodyAcls) GetAckClusterConnectorId() *string {
	return s.AckClusterConnectorId
}

func (s *DescribeAddressBookResponseBodyAcls) GetAckClusterConnectorName() *string {
	return s.AckClusterConnectorName
}

func (s *DescribeAddressBookResponseBodyAcls) GetAckLabels() []*DescribeAddressBookResponseBodyAclsAckLabels {
	return s.AckLabels
}

func (s *DescribeAddressBookResponseBodyAcls) GetAckNamespaces() []*string {
	return s.AckNamespaces
}

func (s *DescribeAddressBookResponseBodyAcls) GetAddressList() []*string {
	return s.AddressList
}

func (s *DescribeAddressBookResponseBodyAcls) GetAddressListCount() *int32 {
	return s.AddressListCount
}

func (s *DescribeAddressBookResponseBodyAcls) GetAddresses() []*DescribeAddressBookResponseBodyAclsAddresses {
	return s.Addresses
}

func (s *DescribeAddressBookResponseBodyAcls) GetAssetMemberUids() []*int64 {
	return s.AssetMemberUids
}

func (s *DescribeAddressBookResponseBodyAcls) GetAssetRegionResourceTypes() []*DescribeAddressBookResponseBodyAclsAssetRegionResourceTypes {
	return s.AssetRegionResourceTypes
}

func (s *DescribeAddressBookResponseBodyAcls) GetAutoAddTagEcs() *int32 {
	return s.AutoAddTagEcs
}

func (s *DescribeAddressBookResponseBodyAcls) GetDescription() *string {
	return s.Description
}

func (s *DescribeAddressBookResponseBodyAcls) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeAddressBookResponseBodyAcls) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeAddressBookResponseBodyAcls) GetGroupUuid() *string {
	return s.GroupUuid
}

func (s *DescribeAddressBookResponseBodyAcls) GetReferenceCount() *int32 {
	return s.ReferenceCount
}

func (s *DescribeAddressBookResponseBodyAcls) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeAddressBookResponseBodyAcls) GetTagList() []*DescribeAddressBookResponseBodyAclsTagList {
	return s.TagList
}

func (s *DescribeAddressBookResponseBodyAcls) GetTagRelation() *string {
	return s.TagRelation
}

func (s *DescribeAddressBookResponseBodyAcls) SetAckClusterConnectorId(v string) *DescribeAddressBookResponseBodyAcls {
	s.AckClusterConnectorId = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetAckClusterConnectorName(v string) *DescribeAddressBookResponseBodyAcls {
	s.AckClusterConnectorName = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetAckLabels(v []*DescribeAddressBookResponseBodyAclsAckLabels) *DescribeAddressBookResponseBodyAcls {
	s.AckLabels = v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetAckNamespaces(v []*string) *DescribeAddressBookResponseBodyAcls {
	s.AckNamespaces = v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetAddressList(v []*string) *DescribeAddressBookResponseBodyAcls {
	s.AddressList = v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetAddressListCount(v int32) *DescribeAddressBookResponseBodyAcls {
	s.AddressListCount = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetAddresses(v []*DescribeAddressBookResponseBodyAclsAddresses) *DescribeAddressBookResponseBodyAcls {
	s.Addresses = v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetAssetMemberUids(v []*int64) *DescribeAddressBookResponseBodyAcls {
	s.AssetMemberUids = v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetAssetRegionResourceTypes(v []*DescribeAddressBookResponseBodyAclsAssetRegionResourceTypes) *DescribeAddressBookResponseBodyAcls {
	s.AssetRegionResourceTypes = v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetAutoAddTagEcs(v int32) *DescribeAddressBookResponseBodyAcls {
	s.AutoAddTagEcs = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetDescription(v string) *DescribeAddressBookResponseBodyAcls {
	s.Description = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetGroupName(v string) *DescribeAddressBookResponseBodyAcls {
	s.GroupName = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetGroupType(v string) *DescribeAddressBookResponseBodyAcls {
	s.GroupType = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetGroupUuid(v string) *DescribeAddressBookResponseBodyAcls {
	s.GroupUuid = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetReferenceCount(v int32) *DescribeAddressBookResponseBodyAcls {
	s.ReferenceCount = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetRegionNo(v string) *DescribeAddressBookResponseBodyAcls {
	s.RegionNo = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetTagList(v []*DescribeAddressBookResponseBodyAclsTagList) *DescribeAddressBookResponseBodyAcls {
	s.TagList = v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetTagRelation(v string) *DescribeAddressBookResponseBodyAcls {
	s.TagRelation = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) Validate() error {
	if s.AckLabels != nil {
		for _, item := range s.AckLabels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Addresses != nil {
		for _, item := range s.Addresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AssetRegionResourceTypes != nil {
		for _, item := range s.AssetRegionResourceTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TagList != nil {
		for _, item := range s.TagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAddressBookResponseBodyAclsAckLabels struct {
	// The key of the pod label in the ACK cluster.
	//
	// example:
	//
	// app
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the pod label in the ACK cluster.
	//
	// example:
	//
	// storage-operator
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAddressBookResponseBodyAclsAckLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressBookResponseBodyAclsAckLabels) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponseBodyAclsAckLabels) GetKey() *string {
	return s.Key
}

func (s *DescribeAddressBookResponseBodyAclsAckLabels) GetValue() *string {
	return s.Value
}

func (s *DescribeAddressBookResponseBodyAclsAckLabels) SetKey(v string) *DescribeAddressBookResponseBodyAclsAckLabels {
	s.Key = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAckLabels) SetValue(v string) *DescribeAddressBookResponseBodyAclsAckLabels {
	s.Value = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAckLabels) Validate() error {
	return dara.Validate(s)
}

type DescribeAddressBookResponseBodyAclsAddresses struct {
	// The address information of the address book.
	//
	// example:
	//
	// 192.168.0.1/32
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The description of the individual address.
	//
	// example:
	//
	// Single Address Description
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s DescribeAddressBookResponseBodyAclsAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressBookResponseBodyAclsAddresses) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponseBodyAclsAddresses) GetAddress() *string {
	return s.Address
}

func (s *DescribeAddressBookResponseBodyAclsAddresses) GetNote() *string {
	return s.Note
}

func (s *DescribeAddressBookResponseBodyAclsAddresses) SetAddress(v string) *DescribeAddressBookResponseBodyAclsAddresses {
	s.Address = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAddresses) SetNote(v string) *DescribeAddressBookResponseBodyAclsAddresses {
	s.Note = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAddresses) Validate() error {
	return dara.Validate(s)
}

type DescribeAddressBookResponseBodyAclsAssetRegionResourceTypes struct {
	// The region ID of the asset.
	//
	// example:
	//
	// all
	AssetRegionId *string `json:"AssetRegionId,omitempty" xml:"AssetRegionId,omitempty"`
	// The asset type.
	ResourceType *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceType `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" type:"Struct"`
}

func (s DescribeAddressBookResponseBodyAclsAssetRegionResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressBookResponseBodyAclsAssetRegionResourceTypes) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypes) GetAssetRegionId() *string {
	return s.AssetRegionId
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypes) GetResourceType() *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceType {
	return s.ResourceType
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypes) SetAssetRegionId(v string) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypes {
	s.AssetRegionId = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypes) SetResourceType(v *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceType) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypes {
	s.ResourceType = v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypes) Validate() error {
	if s.ResourceType != nil {
		if err := s.ResourceType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceType struct {
	// The IPv4 asset type.
	Ipv4 *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 `json:"Ipv4,omitempty" xml:"Ipv4,omitempty" type:"Struct"`
	// The IPv6 asset type.
	Ipv6 *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6 `json:"Ipv6,omitempty" xml:"Ipv6,omitempty" type:"Struct"`
}

func (s DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceType) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceType) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceType) GetIpv4() *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	return s.Ipv4
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceType) GetIpv6() *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6 {
	return s.Ipv6
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceType) SetIpv4(v *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceType {
	s.Ipv4 = v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceType) SetIpv6(v *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceType {
	s.Ipv6 = v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceType) Validate() error {
	if s.Ipv4 != nil {
		if err := s.Ipv4.Validate(); err != nil {
			return err
		}
	}
	if s.Ipv6 != nil {
		if err := s.Ipv6.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 struct {
	// The asset type: AIGatewayEIP.
	//
	// example:
	//
	// false
	AiGatewayEIP *bool `json:"AiGatewayEIP,omitempty" xml:"AiGatewayEIP,omitempty"`
	// The asset type: AlbEIP.
	//
	// example:
	//
	// false
	AlbEIP *bool `json:"AlbEIP,omitempty" xml:"AlbEIP,omitempty"`
	// The asset type: ApigEIP.
	//
	// example:
	//
	// false
	ApiGatewayEIP *bool `json:"ApiGatewayEIP,omitempty" xml:"ApiGatewayEIP,omitempty"`
	// The asset type: BastionHostEgressIP.
	//
	// example:
	//
	// false
	BastionHostEgressIP *bool `json:"BastionHostEgressIP,omitempty" xml:"BastionHostEgressIP,omitempty"`
	// The asset type: BastionHostIP.
	//
	// example:
	//
	// false
	BastionHostIP *bool `json:"BastionHostIP,omitempty" xml:"BastionHostIP,omitempty"`
	// The asset type: BastionHostIngressIP.
	//
	// example:
	//
	// false
	BastionHostIngressIP *bool `json:"BastionHostIngressIP,omitempty" xml:"BastionHostIngressIP,omitempty"`
	// The asset type: EIP.
	//
	// example:
	//
	// false
	EIP *bool `json:"EIP,omitempty" xml:"EIP,omitempty"`
	// The asset type: EcsEIP.
	//
	// example:
	//
	// false
	EcsEIP *bool `json:"EcsEIP,omitempty" xml:"EcsEIP,omitempty"`
	// The asset type: EcsPublicIP.
	//
	// example:
	//
	// false
	EcsPublicIP *bool `json:"EcsPublicIP,omitempty" xml:"EcsPublicIP,omitempty"`
	// The asset type: EniEIP.
	//
	// example:
	//
	// false
	EniEIP *bool `json:"EniEIP,omitempty" xml:"EniEIP,omitempty"`
	// The asset type: GaEIP.
	//
	// example:
	//
	// false
	GaEIP *bool `json:"GaEIP,omitempty" xml:"GaEIP,omitempty"`
	// The asset type: HAVIP.
	//
	// example:
	//
	// false
	HAVIP *bool `json:"HAVIP,omitempty" xml:"HAVIP,omitempty"`
	// The asset type: NatEIP.
	//
	// example:
	//
	// false
	NatEIP *bool `json:"NatEIP,omitempty" xml:"NatEIP,omitempty"`
	// The asset type: NatPublicIP.
	//
	// example:
	//
	// false
	NatPublicIP *bool `json:"NatPublicIP,omitempty" xml:"NatPublicIP,omitempty"`
	// The asset type: NlbEIP.
	//
	// example:
	//
	// false
	NlbEIP *bool `json:"NlbEIP,omitempty" xml:"NlbEIP,omitempty"`
	// The asset type: SlbEIP.
	//
	// example:
	//
	// true
	SlbEIP *bool `json:"SlbEIP,omitempty" xml:"SlbEIP,omitempty"`
	// The asset type: SlbPublicIP.
	//
	// example:
	//
	// false
	SlbPublicIP *bool `json:"SlbPublicIP,omitempty" xml:"SlbPublicIP,omitempty"`
}

func (s DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetAiGatewayEIP() *bool {
	return s.AiGatewayEIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetAlbEIP() *bool {
	return s.AlbEIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetApiGatewayEIP() *bool {
	return s.ApiGatewayEIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetBastionHostEgressIP() *bool {
	return s.BastionHostEgressIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetBastionHostIP() *bool {
	return s.BastionHostIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetBastionHostIngressIP() *bool {
	return s.BastionHostIngressIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetEIP() *bool {
	return s.EIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetEcsEIP() *bool {
	return s.EcsEIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetEcsPublicIP() *bool {
	return s.EcsPublicIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetEniEIP() *bool {
	return s.EniEIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetGaEIP() *bool {
	return s.GaEIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetHAVIP() *bool {
	return s.HAVIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetNatEIP() *bool {
	return s.NatEIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetNatPublicIP() *bool {
	return s.NatPublicIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetNlbEIP() *bool {
	return s.NlbEIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetSlbEIP() *bool {
	return s.SlbEIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) GetSlbPublicIP() *bool {
	return s.SlbPublicIP
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetAiGatewayEIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.AiGatewayEIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetAlbEIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.AlbEIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetApiGatewayEIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.ApiGatewayEIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetBastionHostEgressIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.BastionHostEgressIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetBastionHostIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.BastionHostIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetBastionHostIngressIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.BastionHostIngressIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetEIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.EIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetEcsEIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.EcsEIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetEcsPublicIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.EcsPublicIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetEniEIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.EniEIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetGaEIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.GaEIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetHAVIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.HAVIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetNatEIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.NatEIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetNatPublicIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.NatPublicIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetNlbEIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.NlbEIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetSlbEIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.SlbEIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) SetSlbPublicIP(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4 {
	s.SlbPublicIP = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv4) Validate() error {
	return dara.Validate(s)
}

type DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6 struct {
	// The asset type: AIGatewayEIPv6.
	//
	// example:
	//
	// false
	AiGatewayEIPv6 *bool `json:"AiGatewayEIPv6,omitempty" xml:"AiGatewayEIPv6,omitempty"`
	// The asset type: AlbIPv6.
	//
	// example:
	//
	// false
	AlbIPv6 *bool `json:"AlbIPv6,omitempty" xml:"AlbIPv6,omitempty"`
	// The asset type: ApigEIPv6.
	//
	// example:
	//
	// false
	ApiGatewayEIPv6 *bool `json:"ApiGatewayEIPv6,omitempty" xml:"ApiGatewayEIPv6,omitempty"`
	// The asset type: EcsIPv6.
	//
	// example:
	//
	// false
	EcsIPv6 *bool `json:"EcsIPv6,omitempty" xml:"EcsIPv6,omitempty"`
	// The asset type: EniEIPv6.
	//
	// example:
	//
	// false
	EniEIPv6 *bool `json:"EniEIPv6,omitempty" xml:"EniEIPv6,omitempty"`
	// The asset type: GaEIPv6.
	//
	// example:
	//
	// false
	GaEIPv6 *bool `json:"GaEIPv6,omitempty" xml:"GaEIPv6,omitempty"`
	// The asset type: NlbIPv6.
	//
	// example:
	//
	// false
	NlbIPv6 *bool `json:"NlbIPv6,omitempty" xml:"NlbIPv6,omitempty"`
	// The asset type: SlbIPv6.
	//
	// example:
	//
	// false
	SlbIPv6 *bool `json:"SlbIPv6,omitempty" xml:"SlbIPv6,omitempty"`
}

func (s DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) GetAiGatewayEIPv6() *bool {
	return s.AiGatewayEIPv6
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) GetAlbIPv6() *bool {
	return s.AlbIPv6
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) GetApiGatewayEIPv6() *bool {
	return s.ApiGatewayEIPv6
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) GetEcsIPv6() *bool {
	return s.EcsIPv6
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) GetEniEIPv6() *bool {
	return s.EniEIPv6
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) GetGaEIPv6() *bool {
	return s.GaEIPv6
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) GetNlbIPv6() *bool {
	return s.NlbIPv6
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) GetSlbIPv6() *bool {
	return s.SlbIPv6
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) SetAiGatewayEIPv6(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6 {
	s.AiGatewayEIPv6 = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) SetAlbIPv6(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6 {
	s.AlbIPv6 = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) SetApiGatewayEIPv6(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6 {
	s.ApiGatewayEIPv6 = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) SetEcsIPv6(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6 {
	s.EcsIPv6 = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) SetEniEIPv6(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6 {
	s.EniEIPv6 = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) SetGaEIPv6(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6 {
	s.GaEIPv6 = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) SetNlbIPv6(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6 {
	s.NlbIPv6 = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) SetSlbIPv6(v bool) *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6 {
	s.SlbIPv6 = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsAssetRegionResourceTypesResourceTypeIpv6) Validate() error {
	return dara.Validate(s)
}

type DescribeAddressBookResponseBodyAclsTagList struct {
	// The key of the ECS tag.
	//
	// example:
	//
	// company
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the ECS tag.
	//
	// example:
	//
	// ALL VALUE
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeAddressBookResponseBodyAclsTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressBookResponseBodyAclsTagList) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponseBodyAclsTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeAddressBookResponseBodyAclsTagList) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeAddressBookResponseBodyAclsTagList) SetTagKey(v string) *DescribeAddressBookResponseBodyAclsTagList {
	s.TagKey = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsTagList) SetTagValue(v string) *DescribeAddressBookResponseBodyAclsTagList {
	s.TagValue = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsTagList) Validate() error {
	return dara.Validate(s)
}
