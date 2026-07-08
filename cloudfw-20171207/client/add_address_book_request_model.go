// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAddressBookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAckClusterConnectorId(v string) *AddAddressBookRequest
	GetAckClusterConnectorId() *string
	SetAckLabels(v []*AddAddressBookRequestAckLabels) *AddAddressBookRequest
	GetAckLabels() []*AddAddressBookRequestAckLabels
	SetAckNamespaces(v []*string) *AddAddressBookRequest
	GetAckNamespaces() []*string
	SetAddressList(v string) *AddAddressBookRequest
	GetAddressList() *string
	SetAssetMemberUids(v []*int64) *AddAddressBookRequest
	GetAssetMemberUids() []*int64
	SetAssetRegionResourceTypes(v []*AddAddressBookRequestAssetRegionResourceTypes) *AddAddressBookRequest
	GetAssetRegionResourceTypes() []*AddAddressBookRequestAssetRegionResourceTypes
	SetAutoAddTagEcs(v string) *AddAddressBookRequest
	GetAutoAddTagEcs() *string
	SetDescription(v string) *AddAddressBookRequest
	GetDescription() *string
	SetGroupName(v string) *AddAddressBookRequest
	GetGroupName() *string
	SetGroupType(v string) *AddAddressBookRequest
	GetGroupType() *string
	SetLang(v string) *AddAddressBookRequest
	GetLang() *string
	SetSourceIp(v string) *AddAddressBookRequest
	GetSourceIp() *string
	SetTagList(v []*AddAddressBookRequestTagList) *AddAddressBookRequest
	GetTagList() []*AddAddressBookRequestTagList
	SetTagRelation(v string) *AddAddressBookRequest
	GetTagRelation() *string
}

type AddAddressBookRequest struct {
	// The ACK cluster connector ID. You can obtain the value from the following operation:
	//
	// - [DescribeAckClusterConnectors](~~DescribeAckClusterConnectors~~): Lists ACK cluster connectors.
	//
	// example:
	//
	// ac-7c1bad6c3cc84c33baab1
	AckClusterConnectorId *string `json:"AckClusterConnectorId,omitempty" xml:"AckClusterConnectorId,omitempty"`
	// The list of ACK cluster pod labels.
	//
	// > A maximum of 10 labels are supported.
	AckLabels []*AddAddressBookRequestAckLabels `json:"AckLabels,omitempty" xml:"AckLabels,omitempty" type:"Repeated"`
	// The list of ACK cluster pod namespaces.
	//
	// > A maximum of 10 namespaces are supported.
	AckNamespaces []*string `json:"AckNamespaces,omitempty" xml:"AckNamespaces,omitempty" type:"Repeated"`
	// The addresses in the address book. Separate multiple addresses with commas (,). Use a space to separate an address from its description within a single address element.
	//
	// > This parameter is required when GroupType is set to `ip`, `port`, or `domain`.
	//
	// - When GroupType is set to `ip`, enter IP addresses in the address list. Example: 192.0.XX.XX/32 Development CIDR block,10.0.0.X/24,192.0.XX.XX/24 Test CIDR block.
	//
	// - When GroupType is set to `port`, enter ports or port ranges in the address list. Example: 80 HTTP port,100/200,3306 Database port.
	//
	// - When GroupType is set to `domain`, enter domain names in the address list. Example: example.com Test domain name,aliyundoc.com,www.aliyun.com Alibaba Cloud official website.
	//
	// example:
	//
	// 192.0.XX.XX/32 ,192.0.XX.XX/24
	AddressList *string `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	// The list of member accounts for the asset address book.
	AssetMemberUids []*int64 `json:"AssetMemberUids,omitempty" xml:"AssetMemberUids,omitempty" type:"Repeated"`
	// The list of regions and resource types for the asset address book.
	AssetRegionResourceTypes []*AddAddressBookRequestAssetRegionResourceTypes `json:"AssetRegionResourceTypes,omitempty" xml:"AssetRegionResourceTypes,omitempty" type:"Repeated"`
	// Indicates whether to automatically add the public IP addresses of Elastic Compute Service (ECS) instances that match the specified tags to the address book.
	//
	// example:
	//
	// 1
	AutoAddTagEcs *string `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	// The description of the address book.
	//
	// This parameter is required.
	//
	// example:
	//
	// sz-001
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the address book.
	//
	// This parameter is required.
	//
	// example:
	//
	// sz-001
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the address book.
	//
	// This parameter is required.
	//
	// example:
	//
	// ip
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The language type of the address book description.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	//
	// The source IP address of the request.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ECS tag list.
	TagList []*AddAddressBookRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The logical relationship among multiple ECS tags to match.
	//
	// example:
	//
	// and
	TagRelation *string `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
}

func (s AddAddressBookRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAddressBookRequest) GoString() string {
	return s.String()
}

func (s *AddAddressBookRequest) GetAckClusterConnectorId() *string {
	return s.AckClusterConnectorId
}

func (s *AddAddressBookRequest) GetAckLabels() []*AddAddressBookRequestAckLabels {
	return s.AckLabels
}

func (s *AddAddressBookRequest) GetAckNamespaces() []*string {
	return s.AckNamespaces
}

func (s *AddAddressBookRequest) GetAddressList() *string {
	return s.AddressList
}

func (s *AddAddressBookRequest) GetAssetMemberUids() []*int64 {
	return s.AssetMemberUids
}

func (s *AddAddressBookRequest) GetAssetRegionResourceTypes() []*AddAddressBookRequestAssetRegionResourceTypes {
	return s.AssetRegionResourceTypes
}

func (s *AddAddressBookRequest) GetAutoAddTagEcs() *string {
	return s.AutoAddTagEcs
}

func (s *AddAddressBookRequest) GetDescription() *string {
	return s.Description
}

func (s *AddAddressBookRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *AddAddressBookRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *AddAddressBookRequest) GetLang() *string {
	return s.Lang
}

func (s *AddAddressBookRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *AddAddressBookRequest) GetTagList() []*AddAddressBookRequestTagList {
	return s.TagList
}

func (s *AddAddressBookRequest) GetTagRelation() *string {
	return s.TagRelation
}

func (s *AddAddressBookRequest) SetAckClusterConnectorId(v string) *AddAddressBookRequest {
	s.AckClusterConnectorId = &v
	return s
}

func (s *AddAddressBookRequest) SetAckLabels(v []*AddAddressBookRequestAckLabels) *AddAddressBookRequest {
	s.AckLabels = v
	return s
}

func (s *AddAddressBookRequest) SetAckNamespaces(v []*string) *AddAddressBookRequest {
	s.AckNamespaces = v
	return s
}

func (s *AddAddressBookRequest) SetAddressList(v string) *AddAddressBookRequest {
	s.AddressList = &v
	return s
}

func (s *AddAddressBookRequest) SetAssetMemberUids(v []*int64) *AddAddressBookRequest {
	s.AssetMemberUids = v
	return s
}

func (s *AddAddressBookRequest) SetAssetRegionResourceTypes(v []*AddAddressBookRequestAssetRegionResourceTypes) *AddAddressBookRequest {
	s.AssetRegionResourceTypes = v
	return s
}

func (s *AddAddressBookRequest) SetAutoAddTagEcs(v string) *AddAddressBookRequest {
	s.AutoAddTagEcs = &v
	return s
}

func (s *AddAddressBookRequest) SetDescription(v string) *AddAddressBookRequest {
	s.Description = &v
	return s
}

func (s *AddAddressBookRequest) SetGroupName(v string) *AddAddressBookRequest {
	s.GroupName = &v
	return s
}

func (s *AddAddressBookRequest) SetGroupType(v string) *AddAddressBookRequest {
	s.GroupType = &v
	return s
}

func (s *AddAddressBookRequest) SetLang(v string) *AddAddressBookRequest {
	s.Lang = &v
	return s
}

func (s *AddAddressBookRequest) SetSourceIp(v string) *AddAddressBookRequest {
	s.SourceIp = &v
	return s
}

func (s *AddAddressBookRequest) SetTagList(v []*AddAddressBookRequestTagList) *AddAddressBookRequest {
	s.TagList = v
	return s
}

func (s *AddAddressBookRequest) SetTagRelation(v string) *AddAddressBookRequest {
	s.TagRelation = &v
	return s
}

func (s *AddAddressBookRequest) Validate() error {
	if s.AckLabels != nil {
		for _, item := range s.AckLabels {
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

type AddAddressBookRequestAckLabels struct {
	// The key of the ACK cluster pod label.
	//
	// example:
	//
	// app
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the ACK cluster pod label.
	//
	// example:
	//
	// storage-operator
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddAddressBookRequestAckLabels) String() string {
	return dara.Prettify(s)
}

func (s AddAddressBookRequestAckLabels) GoString() string {
	return s.String()
}

func (s *AddAddressBookRequestAckLabels) GetKey() *string {
	return s.Key
}

func (s *AddAddressBookRequestAckLabels) GetValue() *string {
	return s.Value
}

func (s *AddAddressBookRequestAckLabels) SetKey(v string) *AddAddressBookRequestAckLabels {
	s.Key = &v
	return s
}

func (s *AddAddressBookRequestAckLabels) SetValue(v string) *AddAddressBookRequestAckLabels {
	s.Value = &v
	return s
}

func (s *AddAddressBookRequestAckLabels) Validate() error {
	return dara.Validate(s)
}

type AddAddressBookRequestAssetRegionResourceTypes struct {
	// The region ID of the asset.
	//
	// example:
	//
	// all
	AssetRegionId *string `json:"AssetRegionId,omitempty" xml:"AssetRegionId,omitempty"`
	// The asset type.
	ResourceType *AddAddressBookRequestAssetRegionResourceTypesResourceType `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" type:"Struct"`
}

func (s AddAddressBookRequestAssetRegionResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s AddAddressBookRequestAssetRegionResourceTypes) GoString() string {
	return s.String()
}

func (s *AddAddressBookRequestAssetRegionResourceTypes) GetAssetRegionId() *string {
	return s.AssetRegionId
}

func (s *AddAddressBookRequestAssetRegionResourceTypes) GetResourceType() *AddAddressBookRequestAssetRegionResourceTypesResourceType {
	return s.ResourceType
}

func (s *AddAddressBookRequestAssetRegionResourceTypes) SetAssetRegionId(v string) *AddAddressBookRequestAssetRegionResourceTypes {
	s.AssetRegionId = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypes) SetResourceType(v *AddAddressBookRequestAssetRegionResourceTypesResourceType) *AddAddressBookRequestAssetRegionResourceTypes {
	s.ResourceType = v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypes) Validate() error {
	if s.ResourceType != nil {
		if err := s.ResourceType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddAddressBookRequestAssetRegionResourceTypesResourceType struct {
	// The IPv4 asset type.
	Ipv4 *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 `json:"Ipv4,omitempty" xml:"Ipv4,omitempty" type:"Struct"`
	// The IPv6 asset type.
	Ipv6 *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 `json:"Ipv6,omitempty" xml:"Ipv6,omitempty" type:"Struct"`
}

func (s AddAddressBookRequestAssetRegionResourceTypesResourceType) String() string {
	return dara.Prettify(s)
}

func (s AddAddressBookRequestAssetRegionResourceTypesResourceType) GoString() string {
	return s.String()
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceType) GetIpv4() *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	return s.Ipv4
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceType) GetIpv6() *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	return s.Ipv6
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceType) SetIpv4(v *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) *AddAddressBookRequestAssetRegionResourceTypesResourceType {
	s.Ipv4 = v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceType) SetIpv6(v *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) *AddAddressBookRequestAssetRegionResourceTypesResourceType {
	s.Ipv6 = v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceType) Validate() error {
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

type AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 struct {
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

func (s AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) String() string {
	return dara.Prettify(s)
}

func (s AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GoString() string {
	return s.String()
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetAiGatewayEIP() *bool {
	return s.AiGatewayEIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetAlbEIP() *bool {
	return s.AlbEIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetApiGatewayEIP() *bool {
	return s.ApiGatewayEIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetBastionHostEgressIP() *bool {
	return s.BastionHostEgressIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetBastionHostIP() *bool {
	return s.BastionHostIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetBastionHostIngressIP() *bool {
	return s.BastionHostIngressIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetEIP() *bool {
	return s.EIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetEcsEIP() *bool {
	return s.EcsEIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetEcsPublicIP() *bool {
	return s.EcsPublicIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetEniEIP() *bool {
	return s.EniEIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetGaEIP() *bool {
	return s.GaEIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetHAVIP() *bool {
	return s.HAVIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetNatEIP() *bool {
	return s.NatEIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetNatPublicIP() *bool {
	return s.NatPublicIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetNlbEIP() *bool {
	return s.NlbEIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetSlbEIP() *bool {
	return s.SlbEIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetSlbPublicIP() *bool {
	return s.SlbPublicIP
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetAiGatewayEIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.AiGatewayEIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetAlbEIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.AlbEIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetApiGatewayEIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.ApiGatewayEIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetBastionHostEgressIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.BastionHostEgressIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetBastionHostIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.BastionHostIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetBastionHostIngressIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.BastionHostIngressIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetEIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.EIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetEcsEIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.EcsEIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetEcsPublicIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.EcsPublicIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetEniEIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.EniEIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetGaEIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.GaEIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetHAVIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.HAVIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetNatEIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.NatEIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetNatPublicIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.NatPublicIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetNlbEIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.NlbEIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetSlbEIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.SlbEIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetSlbPublicIP(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.SlbPublicIP = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) Validate() error {
	return dara.Validate(s)
}

type AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 struct {
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

func (s AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) String() string {
	return dara.Prettify(s)
}

func (s AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GoString() string {
	return s.String()
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetAiGatewayEIPv6() *bool {
	return s.AiGatewayEIPv6
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetAlbIPv6() *bool {
	return s.AlbIPv6
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetApiGatewayEIPv6() *bool {
	return s.ApiGatewayEIPv6
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetEcsIPv6() *bool {
	return s.EcsIPv6
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetEniEIPv6() *bool {
	return s.EniEIPv6
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetGaEIPv6() *bool {
	return s.GaEIPv6
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetNlbIPv6() *bool {
	return s.NlbIPv6
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetSlbIPv6() *bool {
	return s.SlbIPv6
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetAiGatewayEIPv6(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.AiGatewayEIPv6 = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetAlbIPv6(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.AlbIPv6 = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetApiGatewayEIPv6(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.ApiGatewayEIPv6 = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetEcsIPv6(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.EcsIPv6 = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetEniEIPv6(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.EniEIPv6 = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetGaEIPv6(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.GaEIPv6 = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetNlbIPv6(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.NlbIPv6 = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetSlbIPv6(v bool) *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.SlbIPv6 = &v
	return s
}

func (s *AddAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) Validate() error {
	return dara.Validate(s)
}

type AddAddressBookRequestTagList struct {
	// The key of the ECS tag.
	//
	// example:
	//
	// TXY
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the ECS tag.
	//
	// example:
	//
	// 1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s AddAddressBookRequestTagList) String() string {
	return dara.Prettify(s)
}

func (s AddAddressBookRequestTagList) GoString() string {
	return s.String()
}

func (s *AddAddressBookRequestTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *AddAddressBookRequestTagList) GetTagValue() *string {
	return s.TagValue
}

func (s *AddAddressBookRequestTagList) SetTagKey(v string) *AddAddressBookRequestTagList {
	s.TagKey = &v
	return s
}

func (s *AddAddressBookRequestTagList) SetTagValue(v string) *AddAddressBookRequestTagList {
	s.TagValue = &v
	return s
}

func (s *AddAddressBookRequestTagList) Validate() error {
	return dara.Validate(s)
}
