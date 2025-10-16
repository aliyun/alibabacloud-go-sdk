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
	// example:
	//
	// ac-7c1bad6c3cc84c33baab1
	AckClusterConnectorId *string                           `json:"AckClusterConnectorId,omitempty" xml:"AckClusterConnectorId,omitempty"`
	AckLabels             []*AddAddressBookRequestAckLabels `json:"AckLabels,omitempty" xml:"AckLabels,omitempty" type:"Repeated"`
	AckNamespaces         []*string                         `json:"AckNamespaces,omitempty" xml:"AckNamespaces,omitempty" type:"Repeated"`
	// The addresses that you want to add to the address book. Separate multiple addresses with commas (,).
	//
	// >  If you set GroupType to `ip`, `port` or `domain`, you must specify AddressList.
	//
	// 	- If you set GroupType to `ip`, you must add IP addresses to the address book. Example: 192.0.XX.XX/32,192.0.XX.XX/24.
	//
	// 	- If you set GroupType to `port`, you must add port numbers or port ranges to the address book. Example: 80,100/200.
	//
	// 	- If you set GroupType to `domain`, you must add domain names to the address book. Example: example.com,aliyundoc.com.
	//
	// example:
	//
	// 192.0.XX.XX/32, 192.0.XX.XX/24
	AddressList *string `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	// Specifies whether to automatically add public IP addresses of ECS instances to the address book if the instances match the specified tags. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0*	- (default): no
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
	// The type of the address book. Valid values:
	//
	// 	- **ip**: IP address book
	//
	// 	- **domain**: domain address book
	//
	// 	- **port**: port address book
	//
	// 	- **tag**: ECS tag-based address book
	//
	// This parameter is required.
	//
	// example:
	//
	// ip
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
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
	// The ECS tags that you want to match.
	TagList []*AddAddressBookRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The logical relation among the ECS tags that you want to match. Valid values:
	//
	// 	- **and*	- (default): Only the public IP addresses of ECS instances that match all the specified tags can be added to the address book.
	//
	// 	- **or**: The public IP addresses of ECS instances that match one of the specified tags can be added to the address book.
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
	// example:
	//
	// app
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
