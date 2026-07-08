// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAddressBookShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAckLabels(v []*ModifyAddressBookShrinkRequestAckLabels) *ModifyAddressBookShrinkRequest
	GetAckLabels() []*ModifyAddressBookShrinkRequestAckLabels
	SetAckNamespaces(v []*string) *ModifyAddressBookShrinkRequest
	GetAckNamespaces() []*string
	SetAddressList(v string) *ModifyAddressBookShrinkRequest
	GetAddressList() *string
	SetAssetMemberUidsShrink(v string) *ModifyAddressBookShrinkRequest
	GetAssetMemberUidsShrink() *string
	SetAssetRegionResourceTypesShrink(v string) *ModifyAddressBookShrinkRequest
	GetAssetRegionResourceTypesShrink() *string
	SetAutoAddTagEcs(v string) *ModifyAddressBookShrinkRequest
	GetAutoAddTagEcs() *string
	SetDescription(v string) *ModifyAddressBookShrinkRequest
	GetDescription() *string
	SetGroupName(v string) *ModifyAddressBookShrinkRequest
	GetGroupName() *string
	SetGroupUuid(v string) *ModifyAddressBookShrinkRequest
	GetGroupUuid() *string
	SetLang(v string) *ModifyAddressBookShrinkRequest
	GetLang() *string
	SetModifyMode(v string) *ModifyAddressBookShrinkRequest
	GetModifyMode() *string
	SetSourceIp(v string) *ModifyAddressBookShrinkRequest
	GetSourceIp() *string
	SetTagList(v []*ModifyAddressBookShrinkRequestTagList) *ModifyAddressBookShrinkRequest
	GetTagList() []*ModifyAddressBookShrinkRequestTagList
	SetTagRelation(v string) *ModifyAddressBookShrinkRequest
	GetTagRelation() *string
}

type ModifyAddressBookShrinkRequest struct {
	// The list of labels for ACK cluster pods.
	//
	// > A maximum of 10 labels are supported.
	AckLabels []*ModifyAddressBookShrinkRequestAckLabels `json:"AckLabels,omitempty" xml:"AckLabels,omitempty" type:"Repeated"`
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
	AssetMemberUidsShrink *string `json:"AssetMemberUids,omitempty" xml:"AssetMemberUids,omitempty"`
	// The list of regions and resource types for the asset address book.
	AssetRegionResourceTypesShrink *string `json:"AssetRegionResourceTypes,omitempty" xml:"AssetRegionResourceTypes,omitempty"`
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
	TagList []*ModifyAddressBookShrinkRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
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

func (s ModifyAddressBookShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAddressBookShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookShrinkRequest) GetAckLabels() []*ModifyAddressBookShrinkRequestAckLabels {
	return s.AckLabels
}

func (s *ModifyAddressBookShrinkRequest) GetAckNamespaces() []*string {
	return s.AckNamespaces
}

func (s *ModifyAddressBookShrinkRequest) GetAddressList() *string {
	return s.AddressList
}

func (s *ModifyAddressBookShrinkRequest) GetAssetMemberUidsShrink() *string {
	return s.AssetMemberUidsShrink
}

func (s *ModifyAddressBookShrinkRequest) GetAssetRegionResourceTypesShrink() *string {
	return s.AssetRegionResourceTypesShrink
}

func (s *ModifyAddressBookShrinkRequest) GetAutoAddTagEcs() *string {
	return s.AutoAddTagEcs
}

func (s *ModifyAddressBookShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAddressBookShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyAddressBookShrinkRequest) GetGroupUuid() *string {
	return s.GroupUuid
}

func (s *ModifyAddressBookShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyAddressBookShrinkRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyAddressBookShrinkRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyAddressBookShrinkRequest) GetTagList() []*ModifyAddressBookShrinkRequestTagList {
	return s.TagList
}

func (s *ModifyAddressBookShrinkRequest) GetTagRelation() *string {
	return s.TagRelation
}

func (s *ModifyAddressBookShrinkRequest) SetAckLabels(v []*ModifyAddressBookShrinkRequestAckLabels) *ModifyAddressBookShrinkRequest {
	s.AckLabels = v
	return s
}

func (s *ModifyAddressBookShrinkRequest) SetAckNamespaces(v []*string) *ModifyAddressBookShrinkRequest {
	s.AckNamespaces = v
	return s
}

func (s *ModifyAddressBookShrinkRequest) SetAddressList(v string) *ModifyAddressBookShrinkRequest {
	s.AddressList = &v
	return s
}

func (s *ModifyAddressBookShrinkRequest) SetAssetMemberUidsShrink(v string) *ModifyAddressBookShrinkRequest {
	s.AssetMemberUidsShrink = &v
	return s
}

func (s *ModifyAddressBookShrinkRequest) SetAssetRegionResourceTypesShrink(v string) *ModifyAddressBookShrinkRequest {
	s.AssetRegionResourceTypesShrink = &v
	return s
}

func (s *ModifyAddressBookShrinkRequest) SetAutoAddTagEcs(v string) *ModifyAddressBookShrinkRequest {
	s.AutoAddTagEcs = &v
	return s
}

func (s *ModifyAddressBookShrinkRequest) SetDescription(v string) *ModifyAddressBookShrinkRequest {
	s.Description = &v
	return s
}

func (s *ModifyAddressBookShrinkRequest) SetGroupName(v string) *ModifyAddressBookShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyAddressBookShrinkRequest) SetGroupUuid(v string) *ModifyAddressBookShrinkRequest {
	s.GroupUuid = &v
	return s
}

func (s *ModifyAddressBookShrinkRequest) SetLang(v string) *ModifyAddressBookShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ModifyAddressBookShrinkRequest) SetModifyMode(v string) *ModifyAddressBookShrinkRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyAddressBookShrinkRequest) SetSourceIp(v string) *ModifyAddressBookShrinkRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyAddressBookShrinkRequest) SetTagList(v []*ModifyAddressBookShrinkRequestTagList) *ModifyAddressBookShrinkRequest {
	s.TagList = v
	return s
}

func (s *ModifyAddressBookShrinkRequest) SetTagRelation(v string) *ModifyAddressBookShrinkRequest {
	s.TagRelation = &v
	return s
}

func (s *ModifyAddressBookShrinkRequest) Validate() error {
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

type ModifyAddressBookShrinkRequestAckLabels struct {
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

func (s ModifyAddressBookShrinkRequestAckLabels) String() string {
	return dara.Prettify(s)
}

func (s ModifyAddressBookShrinkRequestAckLabels) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookShrinkRequestAckLabels) GetKey() *string {
	return s.Key
}

func (s *ModifyAddressBookShrinkRequestAckLabels) GetValue() *string {
	return s.Value
}

func (s *ModifyAddressBookShrinkRequestAckLabels) SetKey(v string) *ModifyAddressBookShrinkRequestAckLabels {
	s.Key = &v
	return s
}

func (s *ModifyAddressBookShrinkRequestAckLabels) SetValue(v string) *ModifyAddressBookShrinkRequestAckLabels {
	s.Value = &v
	return s
}

func (s *ModifyAddressBookShrinkRequestAckLabels) Validate() error {
	return dara.Validate(s)
}

type ModifyAddressBookShrinkRequestTagList struct {
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

func (s ModifyAddressBookShrinkRequestTagList) String() string {
	return dara.Prettify(s)
}

func (s ModifyAddressBookShrinkRequestTagList) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookShrinkRequestTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *ModifyAddressBookShrinkRequestTagList) GetTagValue() *string {
	return s.TagValue
}

func (s *ModifyAddressBookShrinkRequestTagList) SetTagKey(v string) *ModifyAddressBookShrinkRequestTagList {
	s.TagKey = &v
	return s
}

func (s *ModifyAddressBookShrinkRequestTagList) SetTagValue(v string) *ModifyAddressBookShrinkRequestTagList {
	s.TagValue = &v
	return s
}

func (s *ModifyAddressBookShrinkRequestTagList) Validate() error {
	return dara.Validate(s)
}
