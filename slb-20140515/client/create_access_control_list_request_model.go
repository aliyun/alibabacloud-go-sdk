// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessControlListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclName(v string) *CreateAccessControlListRequest
	GetAclName() *string
	SetAddressIPVersion(v string) *CreateAccessControlListRequest
	GetAddressIPVersion() *string
	SetOwnerAccount(v string) *CreateAccessControlListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAccessControlListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateAccessControlListRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateAccessControlListRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateAccessControlListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAccessControlListRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateAccessControlListRequestTag) *CreateAccessControlListRequest
	GetTag() []*CreateAccessControlListRequestTag
}

type CreateAccessControlListRequest struct {
	// The name of the ACL. The name must be 1 to 80 characters in length, and can contain letters, digits, periods (.), hyphens (-), forward slashes (/), and underscores (_). The name of the ACL that you create must be unique within each region.
	//
	// This parameter is required.
	//
	// example:
	//
	// rule1
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The IP version. Valid values: **ipv4*	- and **ipv6**.
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rt******
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*CreateAccessControlListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateAccessControlListRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessControlListRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessControlListRequest) GetAclName() *string {
	return s.AclName
}

func (s *CreateAccessControlListRequest) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *CreateAccessControlListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAccessControlListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAccessControlListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAccessControlListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAccessControlListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAccessControlListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAccessControlListRequest) GetTag() []*CreateAccessControlListRequestTag {
	return s.Tag
}

func (s *CreateAccessControlListRequest) SetAclName(v string) *CreateAccessControlListRequest {
	s.AclName = &v
	return s
}

func (s *CreateAccessControlListRequest) SetAddressIPVersion(v string) *CreateAccessControlListRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *CreateAccessControlListRequest) SetOwnerAccount(v string) *CreateAccessControlListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAccessControlListRequest) SetOwnerId(v int64) *CreateAccessControlListRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccessControlListRequest) SetRegionId(v string) *CreateAccessControlListRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAccessControlListRequest) SetResourceGroupId(v string) *CreateAccessControlListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAccessControlListRequest) SetResourceOwnerAccount(v string) *CreateAccessControlListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccessControlListRequest) SetResourceOwnerId(v int64) *CreateAccessControlListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAccessControlListRequest) SetTag(v []*CreateAccessControlListRequestTag) *CreateAccessControlListRequest {
	s.Tag = v
	return s
}

func (s *CreateAccessControlListRequest) Validate() error {
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

type CreateAccessControlListRequestTag struct {
	// The key of tag N. Valid values of N: **1*	- to **20**. The tag key cannot be an empty string. The tag key can be up to 128 characters in length, and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Valid values of N: **1*	- to **20**. The tag value can be an empty string. The tag value must be 0 to 128 characters in length, and cannot start with `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAccessControlListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessControlListRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAccessControlListRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateAccessControlListRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateAccessControlListRequestTag) SetKey(v string) *CreateAccessControlListRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAccessControlListRequestTag) SetValue(v string) *CreateAccessControlListRequestTag {
	s.Value = &v
	return s
}

func (s *CreateAccessControlListRequestTag) Validate() error {
	return dara.Validate(s)
}
