// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerCertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeServerCertificatesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeServerCertificatesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeServerCertificatesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeServerCertificatesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeServerCertificatesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeServerCertificatesRequest
	GetResourceOwnerId() *int64
	SetServerCertificateId(v string) *DescribeServerCertificatesRequest
	GetServerCertificateId() *string
	SetTag(v []*DescribeServerCertificatesRequestTag) *DescribeServerCertificatesRequest
	GetTag() []*DescribeServerCertificatesRequestTag
}

type DescribeServerCertificatesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the CLB instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// >  If the endpoint of the selected region is slb.aliyuncs.com, you must specify `RegionId`.
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
	// The server certificate ID.
	//
	// example:
	//
	// 12315790*******_166f8204689_1714763408_709981430
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	// The tags.
	Tag []*DescribeServerCertificatesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeServerCertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerCertificatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeServerCertificatesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeServerCertificatesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeServerCertificatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServerCertificatesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeServerCertificatesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeServerCertificatesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeServerCertificatesRequest) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *DescribeServerCertificatesRequest) GetTag() []*DescribeServerCertificatesRequestTag {
	return s.Tag
}

func (s *DescribeServerCertificatesRequest) SetOwnerAccount(v string) *DescribeServerCertificatesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeServerCertificatesRequest) SetOwnerId(v int64) *DescribeServerCertificatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeServerCertificatesRequest) SetRegionId(v string) *DescribeServerCertificatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeServerCertificatesRequest) SetResourceGroupId(v string) *DescribeServerCertificatesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeServerCertificatesRequest) SetResourceOwnerAccount(v string) *DescribeServerCertificatesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeServerCertificatesRequest) SetResourceOwnerId(v int64) *DescribeServerCertificatesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeServerCertificatesRequest) SetServerCertificateId(v string) *DescribeServerCertificatesRequest {
	s.ServerCertificateId = &v
	return s
}

func (s *DescribeServerCertificatesRequest) SetTag(v []*DescribeServerCertificatesRequestTag) *DescribeServerCertificatesRequest {
	s.Tag = v
	return s
}

func (s *DescribeServerCertificatesRequest) Validate() error {
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

type DescribeServerCertificatesRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys.
	//
	// The tag key cannot be an empty string. The tag key must be 1 to 64 characters in length and cannot start with `aliyun` or `acs`:. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value cannot be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeServerCertificatesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerCertificatesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeServerCertificatesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeServerCertificatesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeServerCertificatesRequestTag) SetKey(v string) *DescribeServerCertificatesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeServerCertificatesRequestTag) SetValue(v string) *DescribeServerCertificatesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeServerCertificatesRequestTag) Validate() error {
	return dara.Validate(s)
}
