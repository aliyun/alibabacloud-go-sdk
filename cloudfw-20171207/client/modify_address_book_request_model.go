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
	// A list of ACK cluster pod labels.
	//
	// > Up to 10 labels are allowed.
	AckLabels []*ModifyAddressBookRequestAckLabels `json:"AckLabels,omitempty" xml:"AckLabels,omitempty" type:"Repeated"`
	// A list of ACK cluster pod namespaces.
	//
	// > Up to 10 namespaces are allowed.
	AckNamespaces []*string `json:"AckNamespaces,omitempty" xml:"AckNamespaces,omitempty" type:"Repeated"`
	// A list of addresses in the address book. Separate multiple addresses with commas. Within each address element, separate the address and its description with a space. You must specify this parameter when GroupType is **ip**, **port**, or **domain**.
	//
	// - When GroupType is **ip**, specify IP addresses. Example: 1.2.XX.XX/32 development CIDR block, 10.0.0.X/24,1.2.XX.XX/24 test CIDR block.
	//
	// - When GroupType is **port**, specify ports or port ranges. Example: 80/80 HTTP port, 100/200,3306 database port.
	//
	// - When GroupType is **domain**, specify domain names. Example: demo1.aliyun.com test domain, demo2.aliyun.com,www\\.aliyun.com Alibaba Cloud official website.
	//
	// example:
	//
	// 192.0.XX.XX/32 ,192.0.XX.XX/24
	AddressList *string `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	// Specifies whether to automatically add public IP addresses of new ECS instances that match the specified tags to the address book.
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
	// > Obtain this value from [DescribeAddressBook](~~DescribeAddressBook~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0657ab9d-fe8b-4174-b2a6-6baf358e****
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The language type.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The modification mode.
	//
	// > When GroupType is **ip**, **ipv6**, **port**, or **domain**, the default mode is **Cover*	- if this parameter is not specified.
	//
	// > 	Notice: When GroupType is **tag**, this parameter must be empty.
	//
	// example:
	//
	// Cover
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// Deprecated
	//
	// The source IP address of the requester.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// A list of ECS tags.
	TagList []*ModifyAddressBookRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The relationship between multiple ECS tags.
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

type ModifyAddressBookRequestTagList struct {
	// The tag key of the ECS instance.
	//
	// example:
	//
	// TXY
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the ECS instance.
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
