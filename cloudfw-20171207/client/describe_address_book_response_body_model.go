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
	// The information about the address book.
	Acls []*DescribeAddressBookResponseBodyAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
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
	// The total number of the returned address books.
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
	return dara.Validate(s)
}

type DescribeAddressBookResponseBodyAcls struct {
	// The addresses in the address book.
	AddressList []*string `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	// The number of addresses in the address book.
	//
	// example:
	//
	// 2
	AddressListCount *int32 `json:"AddressListCount,omitempty" xml:"AddressListCount,omitempty"`
	// A list of addresses in the address book, each with a single address description.
	Addresses []*DescribeAddressBookResponseBodyAclsAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// Indicates whether the public IP addresses of ECS instances are automatically added to the address book if the instances match the specified tags. The setting takes effect on both newly purchased ECS instances whose tag settings are complete and ECS instances whose tag settings are modified. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	AutoAddTagEcs *int32 `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	// The description of the address book.
	//
	// example:
	//
	// my address book
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the address book.
	//
	// example:
	//
	// demo_address_book
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
	// 	- **allCloud**: cloud service address book
	//
	// 	- **threat**: threat intelligence address book
	//
	// example:
	//
	// ip
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The UUID of the address book.
	//
	// example:
	//
	// f04ac7ce-628b-4cb7-be61-310222b7****
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The number of times that the address book is referenced.
	//
	// example:
	//
	// 3
	ReferenceCount *int32 `json:"ReferenceCount,omitempty" xml:"ReferenceCount,omitempty"`
	// The details about the ECS tags that can be automatically added to the address book.
	TagList []*DescribeAddressBookResponseBodyAclsTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The logical relationship among ECS tags. Valid values:
	//
	// 	- **and**: Only the public IP addresses of ECS instances that match all the specified tags can be added to the address book.
	//
	// 	- **or**: The public IP addresses of ECS instances that match any of the specified tags can be added to the address book.
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

func (s *DescribeAddressBookResponseBodyAcls) GetAddressList() []*string {
	return s.AddressList
}

func (s *DescribeAddressBookResponseBodyAcls) GetAddressListCount() *int32 {
	return s.AddressListCount
}

func (s *DescribeAddressBookResponseBodyAcls) GetAddresses() []*DescribeAddressBookResponseBodyAclsAddresses {
	return s.Addresses
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

func (s *DescribeAddressBookResponseBodyAcls) GetTagList() []*DescribeAddressBookResponseBodyAclsTagList {
	return s.TagList
}

func (s *DescribeAddressBookResponseBodyAcls) GetTagRelation() *string {
	return s.TagRelation
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

func (s *DescribeAddressBookResponseBodyAcls) SetTagList(v []*DescribeAddressBookResponseBodyAclsTagList) *DescribeAddressBookResponseBodyAcls {
	s.TagList = v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetTagRelation(v string) *DescribeAddressBookResponseBodyAcls {
	s.TagRelation = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) Validate() error {
	return dara.Validate(s)
}

type DescribeAddressBookResponseBodyAclsAddresses struct {
	// Address information in the address book.
	//
	// example:
	//
	// 192.168.0.1/32
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// Single address description.
	//
	// example:
	//
	// description
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
	// admin
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
