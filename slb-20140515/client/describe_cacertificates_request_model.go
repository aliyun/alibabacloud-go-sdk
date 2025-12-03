// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCACertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCACertificateId(v string) *DescribeCACertificatesRequest
	GetCACertificateId() *string
	SetOwnerAccount(v string) *DescribeCACertificatesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCACertificatesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCACertificatesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeCACertificatesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeCACertificatesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCACertificatesRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeCACertificatesRequestTag) *DescribeCACertificatesRequest
	GetTag() []*DescribeCACertificatesRequestTag
}

type DescribeCACertificatesRequest struct {
	// The CA certificate ID.
	//
	// example:
	//
	// 139a00604bd-cn-east-hangzho****
	CACertificateId *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region of the CA certificates.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
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
	// rg-atstuj3rtop****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the CA certificates.
	Tag []*DescribeCACertificatesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCACertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCACertificatesRequest) GetCACertificateId() *string {
	return s.CACertificateId
}

func (s *DescribeCACertificatesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCACertificatesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCACertificatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCACertificatesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCACertificatesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCACertificatesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCACertificatesRequest) GetTag() []*DescribeCACertificatesRequestTag {
	return s.Tag
}

func (s *DescribeCACertificatesRequest) SetCACertificateId(v string) *DescribeCACertificatesRequest {
	s.CACertificateId = &v
	return s
}

func (s *DescribeCACertificatesRequest) SetOwnerAccount(v string) *DescribeCACertificatesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCACertificatesRequest) SetOwnerId(v int64) *DescribeCACertificatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCACertificatesRequest) SetRegionId(v string) *DescribeCACertificatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCACertificatesRequest) SetResourceGroupId(v string) *DescribeCACertificatesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCACertificatesRequest) SetResourceOwnerAccount(v string) *DescribeCACertificatesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCACertificatesRequest) SetResourceOwnerId(v int64) *DescribeCACertificatesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCACertificatesRequest) SetTag(v []*DescribeCACertificatesRequestTag) *DescribeCACertificatesRequest {
	s.Tag = v
	return s
}

func (s *DescribeCACertificatesRequest) Validate() error {
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

type DescribeCACertificatesRequestTag struct {
	// The key of tag N. Valid values of N: **1 to 20**. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length, and cannot contain `http://` or `https://`. It must not start with `aliyun` or `acs:`.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Valid values of N: **1 to 20**. The tag value can be an empty string. The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. It must not start with `aliyun` or `acs:`.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCACertificatesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificatesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeCACertificatesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCACertificatesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCACertificatesRequestTag) SetKey(v string) *DescribeCACertificatesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeCACertificatesRequestTag) SetValue(v string) *DescribeCACertificatesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeCACertificatesRequestTag) Validate() error {
	return dara.Validate(s)
}
