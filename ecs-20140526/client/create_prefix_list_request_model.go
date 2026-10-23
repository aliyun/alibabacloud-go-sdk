// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrefixListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressFamily(v string) *CreatePrefixListRequest
	GetAddressFamily() *string
	SetClientToken(v string) *CreatePrefixListRequest
	GetClientToken() *string
	SetDescription(v string) *CreatePrefixListRequest
	GetDescription() *string
	SetEntry(v []*CreatePrefixListRequestEntry) *CreatePrefixListRequest
	GetEntry() []*CreatePrefixListRequestEntry
	SetMaxEntries(v int32) *CreatePrefixListRequest
	GetMaxEntries() *int32
	SetOwnerAccount(v string) *CreatePrefixListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreatePrefixListRequest
	GetOwnerId() *int64
	SetPrefixListName(v string) *CreatePrefixListRequest
	GetPrefixListName() *string
	SetRegionId(v string) *CreatePrefixListRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreatePrefixListRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreatePrefixListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreatePrefixListRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreatePrefixListRequestTag) *CreatePrefixListRequest
	GetTag() []*CreatePrefixListRequestTag
}

type CreatePrefixListRequest struct {
	// The address family of the prefix list. Valid values:
	//
	// - IPv4.
	//
	// - IPv6.
	//
	// This parameter is required.
	//
	// example:
	//
	// IPv4
	AddressFamily *string `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the prefix list. The description must be 2 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The prefix list entry information.
	Entry []*CreatePrefixListRequestEntry `json:"Entry,omitempty" xml:"Entry,omitempty" type:"Repeated"`
	// The maximum number of entries supported by the prefix list. Valid values: 1 to 200.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxEntries   *int32  `json:"MaxEntries,omitempty" xml:"MaxEntries,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the prefix list. The name must be 2 to 128 characters in length, and must start with a letter or a Chinese character. It cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain Chinese characters, letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// PrefixListNameSample
	PrefixListName *string `json:"PrefixListName,omitempty" xml:"PrefixListName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the prefix list belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*CreatePrefixListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreatePrefixListRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrefixListRequest) GoString() string {
	return s.String()
}

func (s *CreatePrefixListRequest) GetAddressFamily() *string {
	return s.AddressFamily
}

func (s *CreatePrefixListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreatePrefixListRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePrefixListRequest) GetEntry() []*CreatePrefixListRequestEntry {
	return s.Entry
}

func (s *CreatePrefixListRequest) GetMaxEntries() *int32 {
	return s.MaxEntries
}

func (s *CreatePrefixListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreatePrefixListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreatePrefixListRequest) GetPrefixListName() *string {
	return s.PrefixListName
}

func (s *CreatePrefixListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePrefixListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePrefixListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreatePrefixListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreatePrefixListRequest) GetTag() []*CreatePrefixListRequestTag {
	return s.Tag
}

func (s *CreatePrefixListRequest) SetAddressFamily(v string) *CreatePrefixListRequest {
	s.AddressFamily = &v
	return s
}

func (s *CreatePrefixListRequest) SetClientToken(v string) *CreatePrefixListRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePrefixListRequest) SetDescription(v string) *CreatePrefixListRequest {
	s.Description = &v
	return s
}

func (s *CreatePrefixListRequest) SetEntry(v []*CreatePrefixListRequestEntry) *CreatePrefixListRequest {
	s.Entry = v
	return s
}

func (s *CreatePrefixListRequest) SetMaxEntries(v int32) *CreatePrefixListRequest {
	s.MaxEntries = &v
	return s
}

func (s *CreatePrefixListRequest) SetOwnerAccount(v string) *CreatePrefixListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreatePrefixListRequest) SetOwnerId(v int64) *CreatePrefixListRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePrefixListRequest) SetPrefixListName(v string) *CreatePrefixListRequest {
	s.PrefixListName = &v
	return s
}

func (s *CreatePrefixListRequest) SetRegionId(v string) *CreatePrefixListRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePrefixListRequest) SetResourceGroupId(v string) *CreatePrefixListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePrefixListRequest) SetResourceOwnerAccount(v string) *CreatePrefixListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePrefixListRequest) SetResourceOwnerId(v int64) *CreatePrefixListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreatePrefixListRequest) SetTag(v []*CreatePrefixListRequestTag) *CreatePrefixListRequest {
	s.Tag = v
	return s
}

func (s *CreatePrefixListRequest) Validate() error {
	if s.Entry != nil {
		for _, item := range s.Entry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePrefixListRequestEntry struct {
	// The Classless Inter-Domain Routing of the prefix list entry. Valid values of N: 0 to 200. Precautions:
	//
	// - The number of entries in the prefix list cannot exceed the maximum number of entries (`MaxEntries`).
	//
	// - The type of the Classless Inter-Domain Routing in an entry is determined by the address family. A prefix list cannot contain both IPv4 and IPv6 Classless Inter-Domain Routing.
	//
	// - Classless Inter-Domain Routing in multiple entries cannot be duplicate. For example, you cannot specify two Settings such as 192.168.1.0/24.
	//
	// - You can specify a single IP address in the Settings, and the system automatically transforms it into a Classless Inter-Domain Routing. For example, if you specify 192.168.1.100 in the Settings, the system automatically transforms it into 192.168.1.100/32.
	//
	// - If you use an IPv6 Classless Inter-Domain Routing, the system automatically transforms it into a zero-compressed format and transforms letters to lowercase. For example, if you specify 2001:0DB8:0000:0000:0000:0000:0000:0000/32 in the Settings, the system automatically transforms it into 2001:db8::/32.
	//
	// For more information about Classless Inter-Domain Routing, see [What is CIDR](~~185311#598efe6ef1v00~~).
	//
	// Default value: empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The description of the prefix list entry. The description must be 2 to 32 characters in length, and cannot start with `http://` or `https://`. Valid values of N: 0 to 200.
	//
	// example:
	//
	// Description Sample 01
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreatePrefixListRequestEntry) String() string {
	return dara.Prettify(s)
}

func (s CreatePrefixListRequestEntry) GoString() string {
	return s.String()
}

func (s *CreatePrefixListRequestEntry) GetCidr() *string {
	return s.Cidr
}

func (s *CreatePrefixListRequestEntry) GetDescription() *string {
	return s.Description
}

func (s *CreatePrefixListRequestEntry) SetCidr(v string) *CreatePrefixListRequestEntry {
	s.Cidr = &v
	return s
}

func (s *CreatePrefixListRequestEntry) SetDescription(v string) *CreatePrefixListRequestEntry {
	s.Description = &v
	return s
}

func (s *CreatePrefixListRequestEntry) Validate() error {
	return dara.Validate(s)
}

type CreatePrefixListRequestTag struct {
	// The tag key of the prefix list. Valid values of N: 1 to 20. If you specify this parameter, it cannot be an empty string. The tag key can be up to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the prefix list. Valid values of N: 1 to 20. This parameter can be an empty string.
	//
	// The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrefixListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreatePrefixListRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePrefixListRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreatePrefixListRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreatePrefixListRequestTag) SetKey(v string) *CreatePrefixListRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePrefixListRequestTag) SetValue(v string) *CreatePrefixListRequestTag {
	s.Value = &v
	return s
}

func (s *CreatePrefixListRequestTag) Validate() error {
	return dara.Validate(s)
}
