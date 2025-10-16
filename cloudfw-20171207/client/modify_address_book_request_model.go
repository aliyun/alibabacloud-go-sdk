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
	AckLabels     []*ModifyAddressBookRequestAckLabels `json:"AckLabels,omitempty" xml:"AckLabels,omitempty" type:"Repeated"`
	AckNamespaces []*string                            `json:"AckNamespaces,omitempty" xml:"AckNamespaces,omitempty" type:"Repeated"`
	// The addresses in the address book. Separate multiple addresses with commas (,). If you set GroupType to **ip**, **port**, or **domain**, you must specify this parameter.
	//
	// 	- If you set GroupType to **ip**, you must specify IP addresses for the address book. Example: 1.2.XX.XX/32,1.2.XX.XX/24.
	//
	// 	- If you set GroupType to **port**, you must specify port numbers or port ranges for the address book. Example: 80/80,100/200.
	//
	// 	- If you set GroupType to **domain**, you must specify domain names for the address book. Example: demo1.aliyun.com,demo2.aliyun.com.
	//
	// example:
	//
	// 192.0.XX.XX/32, 192.0.XX.XX/24
	AddressList *string `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	// Specifies whether to automatically add public IP addresses of Elastic Compute Service (ECS) instances to the address book if the instances match the specified tags. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
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
	// The ID of the address book.
	//
	// >  To modify the address book, you must provide the ID of the address book. You can call the [DescribeAddressBook](https://help.aliyun.com/document_detail/138869.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0657ab9d-fe8b-4174-b2a6-6baf358e****
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Modification mode with the following values:
	//
	// - **Cover**: Use the value of the AddressList parameter to overwrite the original address book.
	//
	// - **Append**: After the original address book, append addresses using the value of the AddressList parameter.
	//
	// - **Delete**: Delete addresses using the value of the AddressList parameter from the address book.
	//
	// >When GroupType is **ip**, **ipv6**, **port**, or **domain**, if this parameter is not configured, the default is to use the **Cover*	- method to modify the address book.
	//
	// 	Notice: When GroupType is **tag**, this parameter must be empty.</notice>
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
	// The ECS tags that you want to match.
	TagList []*ModifyAddressBookRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The logical relationship among ECS tags. Valid values:
	//
	// 	- **and**: Only the public IP addresses of ECS instances that match all the specified tags can be added to the address book.
	//
	// 	- **or**: The public IP addresses of ECS instances that match one of the specified tags can be added to the address book.
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
	// example:
	//
	// app
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	// The key of ECS tag N that you want to match.
	//
	// example:
	//
	// TXY
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of ECS tag N that you want to match.
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
