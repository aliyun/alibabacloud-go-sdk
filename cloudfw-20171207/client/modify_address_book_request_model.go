// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAddressBookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAckLabels(v []*ModifyAddressBookRequestAckLabels) *ModifyAddressBookRequest
	GetAckLabels() []*ModifyAddressBookRequestAckLabels
	SetAckNamespaces(v []*string) *ModifyAddressBookRequest
	GetAckNamespaces() []*string
	SetAddressList(v string) *ModifyAddressBookRequest
	GetAddressList() *string
	SetAssetMemberUids(v []*int64) *ModifyAddressBookRequest
	GetAssetMemberUids() []*int64
	SetAssetRegionResourceTypes(v []*ModifyAddressBookRequestAssetRegionResourceTypes) *ModifyAddressBookRequest
	GetAssetRegionResourceTypes() []*ModifyAddressBookRequestAssetRegionResourceTypes
	SetAutoAddTagEcs(v string) *ModifyAddressBookRequest
	GetAutoAddTagEcs() *string
	SetDescription(v string) *ModifyAddressBookRequest
	GetDescription() *string
	SetGroupName(v string) *ModifyAddressBookRequest
	GetGroupName() *string
	SetGroupUuid(v string) *ModifyAddressBookRequest
	GetGroupUuid() *string
	SetLang(v string) *ModifyAddressBookRequest
	GetLang() *string
	SetModifyMode(v string) *ModifyAddressBookRequest
	GetModifyMode() *string
	SetSourceIp(v string) *ModifyAddressBookRequest
	GetSourceIp() *string
	SetTagList(v []*ModifyAddressBookRequestTagList) *ModifyAddressBookRequest
	GetTagList() []*ModifyAddressBookRequestTagList
	SetTagRelation(v string) *ModifyAddressBookRequest
	GetTagRelation() *string
}

type ModifyAddressBookRequest struct {
	// The list of labels for ACK cluster pods.
	//
	// > A maximum of 10 labels are supported.
	AckLabels []*ModifyAddressBookRequestAckLabels `json:"AckLabels,omitempty" xml:"AckLabels,omitempty" type:"Repeated"`
	// The list of namespaces for ACK cluster pods.
	//
	// > A maximum of 10 namespaces are supported.
	AckNamespaces []*string `json:"AckNamespaces,omitempty" xml:"AckNamespaces,omitempty" type:"Repeated"`
	// The addresses in the address book. Separate multiple addresses with commas (,). Use a space to separate an address from its description. This parameter is required when GroupType is set to **ip**, **port**, or **domain**.
	//
	// - When GroupType is set to **ip**, specify IP addresses. Example: 1.2.XX.XX/32 Development CIDR block,10.0.0.X/24,1.2.XX.XX/24 Test CIDR block.
	//
	// - When GroupType is set to **port**, specify ports or port ranges. Example: 80/80 HTTP port,100/200,3306 Database port.
	//
	// - When GroupType is set to **domain**, specify domain names. Example: demo1.aliyun.com Test domain name,demo2.aliyun.com,www.aliyun.com Alibaba Cloud official website.
	//
	// example:
	//
	// 192.0.XX.XX/32 ,192.0.XX.XX/24
	AddressList *string `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	// The list of member accounts for the asset address book.
	AssetMemberUids []*int64 `json:"AssetMemberUids,omitempty" xml:"AssetMemberUids,omitempty" type:"Repeated"`
	// The list of regions and resource types for the asset address book.
	AssetRegionResourceTypes []*ModifyAddressBookRequestAssetRegionResourceTypes `json:"AssetRegionResourceTypes,omitempty" xml:"AssetRegionResourceTypes,omitempty" type:"Repeated"`
	// Specifies whether the public IP addresses of Elastic Compute Service (ECS) instances that match new labels is automatically added to the address book.
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
	// bj-001
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the address book.
	//
	// This parameter is required.
	//
	// example:
	//
	// bj-001
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The unique ID of the address book.
	//
	// > You can obtain the value by calling the [DescribeAddressBook](~~DescribeAddressBook~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0657ab9d-fe8b-4174-b2a6-6baf358e****
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The language type. Valid values:
	//
	// - **en**: English.
	//
	// - **zh**: Chinese (default).
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The modification mode.
	//
	// > When GroupType is set to **ip**, **ipv6**, **port**, or **domain**, the default value is **Cover*	- if this parameter is not specified.
	//
	// 	Notice: When GroupType is set to **tag**, this parameter must be left empty.</notice>
	//
	// example:
	//
	// Cover
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// Deprecated
	//
	// The source IP address of the request.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ECS tag list.
	TagList []*ModifyAddressBookRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The logical relationship among multiple ECS tags. Valid values:
	//
	// - **or**: The public IP address of an ECS instance is added to the address book if the instance matches any of the specified tags.
	//
	// - **and**: The public IP address of an ECS instance is added to the address book only if the instance matches all of the specified tags.
	//
	// example:
	//
	// and
	TagRelation *string `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
}

func (s ModifyAddressBookRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAddressBookRequest) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookRequest) GetAckLabels() []*ModifyAddressBookRequestAckLabels {
	return s.AckLabels
}

func (s *ModifyAddressBookRequest) GetAckNamespaces() []*string {
	return s.AckNamespaces
}

func (s *ModifyAddressBookRequest) GetAddressList() *string {
	return s.AddressList
}

func (s *ModifyAddressBookRequest) GetAssetMemberUids() []*int64 {
	return s.AssetMemberUids
}

func (s *ModifyAddressBookRequest) GetAssetRegionResourceTypes() []*ModifyAddressBookRequestAssetRegionResourceTypes {
	return s.AssetRegionResourceTypes
}

func (s *ModifyAddressBookRequest) GetAutoAddTagEcs() *string {
	return s.AutoAddTagEcs
}

func (s *ModifyAddressBookRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAddressBookRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyAddressBookRequest) GetGroupUuid() *string {
	return s.GroupUuid
}

func (s *ModifyAddressBookRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyAddressBookRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyAddressBookRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyAddressBookRequest) GetTagList() []*ModifyAddressBookRequestTagList {
	return s.TagList
}

func (s *ModifyAddressBookRequest) GetTagRelation() *string {
	return s.TagRelation
}

func (s *ModifyAddressBookRequest) SetAckLabels(v []*ModifyAddressBookRequestAckLabels) *ModifyAddressBookRequest {
	s.AckLabels = v
	return s
}

func (s *ModifyAddressBookRequest) SetAckNamespaces(v []*string) *ModifyAddressBookRequest {
	s.AckNamespaces = v
	return s
}

func (s *ModifyAddressBookRequest) SetAddressList(v string) *ModifyAddressBookRequest {
	s.AddressList = &v
	return s
}

func (s *ModifyAddressBookRequest) SetAssetMemberUids(v []*int64) *ModifyAddressBookRequest {
	s.AssetMemberUids = v
	return s
}

func (s *ModifyAddressBookRequest) SetAssetRegionResourceTypes(v []*ModifyAddressBookRequestAssetRegionResourceTypes) *ModifyAddressBookRequest {
	s.AssetRegionResourceTypes = v
	return s
}

func (s *ModifyAddressBookRequest) SetAutoAddTagEcs(v string) *ModifyAddressBookRequest {
	s.AutoAddTagEcs = &v
	return s
}

func (s *ModifyAddressBookRequest) SetDescription(v string) *ModifyAddressBookRequest {
	s.Description = &v
	return s
}

func (s *ModifyAddressBookRequest) SetGroupName(v string) *ModifyAddressBookRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyAddressBookRequest) SetGroupUuid(v string) *ModifyAddressBookRequest {
	s.GroupUuid = &v
	return s
}

func (s *ModifyAddressBookRequest) SetLang(v string) *ModifyAddressBookRequest {
	s.Lang = &v
	return s
}

func (s *ModifyAddressBookRequest) SetModifyMode(v string) *ModifyAddressBookRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyAddressBookRequest) SetSourceIp(v string) *ModifyAddressBookRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyAddressBookRequest) SetTagList(v []*ModifyAddressBookRequestTagList) *ModifyAddressBookRequest {
	s.TagList = v
	return s
}

func (s *ModifyAddressBookRequest) SetTagRelation(v string) *ModifyAddressBookRequest {
	s.TagRelation = &v
	return s
}

func (s *ModifyAddressBookRequest) Validate() error {
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

type ModifyAddressBookRequestAckLabels struct {
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

func (s ModifyAddressBookRequestAckLabels) String() string {
	return dara.Prettify(s)
}

func (s ModifyAddressBookRequestAckLabels) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookRequestAckLabels) GetKey() *string {
	return s.Key
}

func (s *ModifyAddressBookRequestAckLabels) GetValue() *string {
	return s.Value
}

func (s *ModifyAddressBookRequestAckLabels) SetKey(v string) *ModifyAddressBookRequestAckLabels {
	s.Key = &v
	return s
}

func (s *ModifyAddressBookRequestAckLabels) SetValue(v string) *ModifyAddressBookRequestAckLabels {
	s.Value = &v
	return s
}

func (s *ModifyAddressBookRequestAckLabels) Validate() error {
	return dara.Validate(s)
}

type ModifyAddressBookRequestAssetRegionResourceTypes struct {
	// The region ID of the asset.
	//
	// example:
	//
	// all
	AssetRegionId *string `json:"AssetRegionId,omitempty" xml:"AssetRegionId,omitempty"`
	// The asset type.
	ResourceType *ModifyAddressBookRequestAssetRegionResourceTypesResourceType `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" type:"Struct"`
}

func (s ModifyAddressBookRequestAssetRegionResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s ModifyAddressBookRequestAssetRegionResourceTypes) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypes) GetAssetRegionId() *string {
	return s.AssetRegionId
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypes) GetResourceType() *ModifyAddressBookRequestAssetRegionResourceTypesResourceType {
	return s.ResourceType
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypes) SetAssetRegionId(v string) *ModifyAddressBookRequestAssetRegionResourceTypes {
	s.AssetRegionId = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypes) SetResourceType(v *ModifyAddressBookRequestAssetRegionResourceTypesResourceType) *ModifyAddressBookRequestAssetRegionResourceTypes {
	s.ResourceType = v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypes) Validate() error {
	if s.ResourceType != nil {
		if err := s.ResourceType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAddressBookRequestAssetRegionResourceTypesResourceType struct {
	// The IPv4 asset type.
	Ipv4 *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 `json:"Ipv4,omitempty" xml:"Ipv4,omitempty" type:"Struct"`
	// The IPv6 asset type.
	Ipv6 *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 `json:"Ipv6,omitempty" xml:"Ipv6,omitempty" type:"Struct"`
}

func (s ModifyAddressBookRequestAssetRegionResourceTypesResourceType) String() string {
	return dara.Prettify(s)
}

func (s ModifyAddressBookRequestAssetRegionResourceTypesResourceType) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceType) GetIpv4() *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	return s.Ipv4
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceType) GetIpv6() *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	return s.Ipv6
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceType) SetIpv4(v *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) *ModifyAddressBookRequestAssetRegionResourceTypesResourceType {
	s.Ipv4 = v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceType) SetIpv6(v *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) *ModifyAddressBookRequestAssetRegionResourceTypesResourceType {
	s.Ipv6 = v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceType) Validate() error {
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

type ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 struct {
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

func (s ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) String() string {
	return dara.Prettify(s)
}

func (s ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetAiGatewayEIP() *bool {
	return s.AiGatewayEIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetAlbEIP() *bool {
	return s.AlbEIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetApiGatewayEIP() *bool {
	return s.ApiGatewayEIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetBastionHostEgressIP() *bool {
	return s.BastionHostEgressIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetBastionHostIP() *bool {
	return s.BastionHostIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetBastionHostIngressIP() *bool {
	return s.BastionHostIngressIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetEIP() *bool {
	return s.EIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetEcsEIP() *bool {
	return s.EcsEIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetEcsPublicIP() *bool {
	return s.EcsPublicIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetEniEIP() *bool {
	return s.EniEIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetGaEIP() *bool {
	return s.GaEIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetHAVIP() *bool {
	return s.HAVIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetNatEIP() *bool {
	return s.NatEIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetNatPublicIP() *bool {
	return s.NatPublicIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetNlbEIP() *bool {
	return s.NlbEIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetSlbEIP() *bool {
	return s.SlbEIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) GetSlbPublicIP() *bool {
	return s.SlbPublicIP
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetAiGatewayEIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.AiGatewayEIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetAlbEIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.AlbEIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetApiGatewayEIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.ApiGatewayEIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetBastionHostEgressIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.BastionHostEgressIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetBastionHostIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.BastionHostIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetBastionHostIngressIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.BastionHostIngressIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetEIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.EIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetEcsEIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.EcsEIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetEcsPublicIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.EcsPublicIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetEniEIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.EniEIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetGaEIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.GaEIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetHAVIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.HAVIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetNatEIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.NatEIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetNatPublicIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.NatPublicIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetNlbEIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.NlbEIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetSlbEIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.SlbEIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) SetSlbPublicIP(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4 {
	s.SlbPublicIP = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv4) Validate() error {
	return dara.Validate(s)
}

type ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 struct {
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

func (s ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) String() string {
	return dara.Prettify(s)
}

func (s ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetAiGatewayEIPv6() *bool {
	return s.AiGatewayEIPv6
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetAlbIPv6() *bool {
	return s.AlbIPv6
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetApiGatewayEIPv6() *bool {
	return s.ApiGatewayEIPv6
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetEcsIPv6() *bool {
	return s.EcsIPv6
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetEniEIPv6() *bool {
	return s.EniEIPv6
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetGaEIPv6() *bool {
	return s.GaEIPv6
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetNlbIPv6() *bool {
	return s.NlbIPv6
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) GetSlbIPv6() *bool {
	return s.SlbIPv6
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetAiGatewayEIPv6(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.AiGatewayEIPv6 = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetAlbIPv6(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.AlbIPv6 = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetApiGatewayEIPv6(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.ApiGatewayEIPv6 = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetEcsIPv6(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.EcsIPv6 = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetEniEIPv6(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.EniEIPv6 = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetGaEIPv6(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.GaEIPv6 = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetNlbIPv6(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.NlbIPv6 = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) SetSlbIPv6(v bool) *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6 {
	s.SlbIPv6 = &v
	return s
}

func (s *ModifyAddressBookRequestAssetRegionResourceTypesResourceTypeIpv6) Validate() error {
	return dara.Validate(s)
}

type ModifyAddressBookRequestTagList struct {
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

func (s ModifyAddressBookRequestTagList) String() string {
	return dara.Prettify(s)
}

func (s ModifyAddressBookRequestTagList) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookRequestTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *ModifyAddressBookRequestTagList) GetTagValue() *string {
	return s.TagValue
}

func (s *ModifyAddressBookRequestTagList) SetTagKey(v string) *ModifyAddressBookRequestTagList {
	s.TagKey = &v
	return s
}

func (s *ModifyAddressBookRequestTagList) SetTagValue(v string) *ModifyAddressBookRequestTagList {
	s.TagValue = &v
	return s
}

func (s *ModifyAddressBookRequestTagList) Validate() error {
	return dara.Validate(s)
}
