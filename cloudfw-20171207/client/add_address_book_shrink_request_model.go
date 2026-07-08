// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAddressBookShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAckClusterConnectorId(v string) *AddAddressBookShrinkRequest
	GetAckClusterConnectorId() *string
	SetAckLabels(v []*AddAddressBookShrinkRequestAckLabels) *AddAddressBookShrinkRequest
	GetAckLabels() []*AddAddressBookShrinkRequestAckLabels
	SetAckNamespaces(v []*string) *AddAddressBookShrinkRequest
	GetAckNamespaces() []*string
	SetAddressList(v string) *AddAddressBookShrinkRequest
	GetAddressList() *string
	SetAssetMemberUidsShrink(v string) *AddAddressBookShrinkRequest
	GetAssetMemberUidsShrink() *string
	SetAssetRegionResourceTypesShrink(v string) *AddAddressBookShrinkRequest
	GetAssetRegionResourceTypesShrink() *string
	SetAutoAddTagEcs(v string) *AddAddressBookShrinkRequest
	GetAutoAddTagEcs() *string
	SetDescription(v string) *AddAddressBookShrinkRequest
	GetDescription() *string
	SetGroupName(v string) *AddAddressBookShrinkRequest
	GetGroupName() *string
	SetGroupType(v string) *AddAddressBookShrinkRequest
	GetGroupType() *string
	SetLang(v string) *AddAddressBookShrinkRequest
	GetLang() *string
	SetSourceIp(v string) *AddAddressBookShrinkRequest
	GetSourceIp() *string
	SetTagList(v []*AddAddressBookShrinkRequestTagList) *AddAddressBookShrinkRequest
	GetTagList() []*AddAddressBookShrinkRequestTagList
	SetTagRelation(v string) *AddAddressBookShrinkRequest
	GetTagRelation() *string
}

type AddAddressBookShrinkRequest struct {
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
	AckLabels []*AddAddressBookShrinkRequestAckLabels `json:"AckLabels,omitempty" xml:"AckLabels,omitempty" type:"Repeated"`
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
	AssetMemberUidsShrink *string `json:"AssetMemberUids,omitempty" xml:"AssetMemberUids,omitempty"`
	// The list of regions and resource types for the asset address book.
	AssetRegionResourceTypesShrink *string `json:"AssetRegionResourceTypes,omitempty" xml:"AssetRegionResourceTypes,omitempty"`
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
	TagList []*AddAddressBookShrinkRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The logical relationship among multiple ECS tags to match.
	//
	// example:
	//
	// and
	TagRelation *string `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
}

func (s AddAddressBookShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAddressBookShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddAddressBookShrinkRequest) GetAckClusterConnectorId() *string {
	return s.AckClusterConnectorId
}

func (s *AddAddressBookShrinkRequest) GetAckLabels() []*AddAddressBookShrinkRequestAckLabels {
	return s.AckLabels
}

func (s *AddAddressBookShrinkRequest) GetAckNamespaces() []*string {
	return s.AckNamespaces
}

func (s *AddAddressBookShrinkRequest) GetAddressList() *string {
	return s.AddressList
}

func (s *AddAddressBookShrinkRequest) GetAssetMemberUidsShrink() *string {
	return s.AssetMemberUidsShrink
}

func (s *AddAddressBookShrinkRequest) GetAssetRegionResourceTypesShrink() *string {
	return s.AssetRegionResourceTypesShrink
}

func (s *AddAddressBookShrinkRequest) GetAutoAddTagEcs() *string {
	return s.AutoAddTagEcs
}

func (s *AddAddressBookShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *AddAddressBookShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *AddAddressBookShrinkRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *AddAddressBookShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *AddAddressBookShrinkRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *AddAddressBookShrinkRequest) GetTagList() []*AddAddressBookShrinkRequestTagList {
	return s.TagList
}

func (s *AddAddressBookShrinkRequest) GetTagRelation() *string {
	return s.TagRelation
}

func (s *AddAddressBookShrinkRequest) SetAckClusterConnectorId(v string) *AddAddressBookShrinkRequest {
	s.AckClusterConnectorId = &v
	return s
}

func (s *AddAddressBookShrinkRequest) SetAckLabels(v []*AddAddressBookShrinkRequestAckLabels) *AddAddressBookShrinkRequest {
	s.AckLabels = v
	return s
}

func (s *AddAddressBookShrinkRequest) SetAckNamespaces(v []*string) *AddAddressBookShrinkRequest {
	s.AckNamespaces = v
	return s
}

func (s *AddAddressBookShrinkRequest) SetAddressList(v string) *AddAddressBookShrinkRequest {
	s.AddressList = &v
	return s
}

func (s *AddAddressBookShrinkRequest) SetAssetMemberUidsShrink(v string) *AddAddressBookShrinkRequest {
	s.AssetMemberUidsShrink = &v
	return s
}

func (s *AddAddressBookShrinkRequest) SetAssetRegionResourceTypesShrink(v string) *AddAddressBookShrinkRequest {
	s.AssetRegionResourceTypesShrink = &v
	return s
}

func (s *AddAddressBookShrinkRequest) SetAutoAddTagEcs(v string) *AddAddressBookShrinkRequest {
	s.AutoAddTagEcs = &v
	return s
}

func (s *AddAddressBookShrinkRequest) SetDescription(v string) *AddAddressBookShrinkRequest {
	s.Description = &v
	return s
}

func (s *AddAddressBookShrinkRequest) SetGroupName(v string) *AddAddressBookShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *AddAddressBookShrinkRequest) SetGroupType(v string) *AddAddressBookShrinkRequest {
	s.GroupType = &v
	return s
}

func (s *AddAddressBookShrinkRequest) SetLang(v string) *AddAddressBookShrinkRequest {
	s.Lang = &v
	return s
}

func (s *AddAddressBookShrinkRequest) SetSourceIp(v string) *AddAddressBookShrinkRequest {
	s.SourceIp = &v
	return s
}

func (s *AddAddressBookShrinkRequest) SetTagList(v []*AddAddressBookShrinkRequestTagList) *AddAddressBookShrinkRequest {
	s.TagList = v
	return s
}

func (s *AddAddressBookShrinkRequest) SetTagRelation(v string) *AddAddressBookShrinkRequest {
	s.TagRelation = &v
	return s
}

func (s *AddAddressBookShrinkRequest) Validate() error {
	if s.AckLabels != nil {
		for _, item := range s.AckLabels {
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

type AddAddressBookShrinkRequestAckLabels struct {
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

func (s AddAddressBookShrinkRequestAckLabels) String() string {
	return dara.Prettify(s)
}

func (s AddAddressBookShrinkRequestAckLabels) GoString() string {
	return s.String()
}

func (s *AddAddressBookShrinkRequestAckLabels) GetKey() *string {
	return s.Key
}

func (s *AddAddressBookShrinkRequestAckLabels) GetValue() *string {
	return s.Value
}

func (s *AddAddressBookShrinkRequestAckLabels) SetKey(v string) *AddAddressBookShrinkRequestAckLabels {
	s.Key = &v
	return s
}

func (s *AddAddressBookShrinkRequestAckLabels) SetValue(v string) *AddAddressBookShrinkRequestAckLabels {
	s.Value = &v
	return s
}

func (s *AddAddressBookShrinkRequestAckLabels) Validate() error {
	return dara.Validate(s)
}

type AddAddressBookShrinkRequestTagList struct {
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

func (s AddAddressBookShrinkRequestTagList) String() string {
	return dara.Prettify(s)
}

func (s AddAddressBookShrinkRequestTagList) GoString() string {
	return s.String()
}

func (s *AddAddressBookShrinkRequestTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *AddAddressBookShrinkRequestTagList) GetTagValue() *string {
	return s.TagValue
}

func (s *AddAddressBookShrinkRequestTagList) SetTagKey(v string) *AddAddressBookShrinkRequestTagList {
	s.TagKey = &v
	return s
}

func (s *AddAddressBookShrinkRequestTagList) SetTagValue(v string) *AddAddressBookShrinkRequestTagList {
	s.TagValue = &v
	return s
}

func (s *AddAddressBookShrinkRequestTagList) Validate() error {
	return dara.Validate(s)
}
