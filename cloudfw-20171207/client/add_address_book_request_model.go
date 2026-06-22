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
	// The ACK cluster connector ID. You can obtain the value from:
	//
	// - [DescribeAckClusterConnectors](~~DescribeAckClusterConnectors~~): queries a list of ACK cluster connectors.
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
	// The address list of the address book. Multiple addresses are separated by commas, and within each address element, the address and its description are separated by a space.
	//
	// > This parameter is required when GroupType is set to `ip`, `port`, or `domain`.
	//
	// - When GroupType is set to `ip`, enter IP addresses in the address list. Example: 192.0.XX.XX/32 development network segment, 10.0.0.X/24,192.0.XX.XX/24 test network segment.
	//
	// - When GroupType is set to `port`, enter ports or port ranges in the address list. Example: 80 HTTP port, 100/200,3306 database port.
	//
	// - When GroupType is set to `domain`, enter domain names in the address list. Example: example.com test domain, aliyundoc.com,www.aliyun.com Alibaba Cloud official website.
	//
	// example:
	//
	// 192.0.XX.XX/32 ,192.0.XX.XX/24
	AddressList *string `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	// Specifies whether to automatically add the public IP addresses of ECS instances that match new tags to the address book.
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
	// The source IP address of the requester.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The list of ECS tags.
	TagList []*AddAddressBookRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The logical relationship among multiple ECS tags to be matched.
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
